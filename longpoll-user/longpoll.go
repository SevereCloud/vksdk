/*
Package longpoll implements User Long Poll API.

Long polling – this is a technology that allows the receiving of information
about new events with the help of "long requests". The server receives the
request but it doesn't immediately send the answer but rather when some event
will happen (for example, receiving a new incoming message), or waiting
period is over.


By using this approach, you can instantly display in your app the most
important events. Be aware that with the help of Long Poll, you won't be able
to send messages. For this, you’ll need to use the messages.send method.

Initialization

The module can be used with the user access key obtained in the Standalone
application via Implicit Flow(access rights required: messages) or with
the community access key(access rights required: messages).

	vk := api.NewVK("<TOKEN>")

And then longpoll

	mode := longpoll.ReceiveAttachments + longpoll.ExtendedEvents
	lp, err := longpoll.NewLongPoll(vk, mode)

Setting

TODO: write about lp.Ts lp.Wait

The module has the ability to modify the HTTP client

	dialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, proxy.Direct)
	httpTransport := &http.Transport{
		Dial:              dialer.Dial,
		// DisableKeepAlives: true,
	}
	httpTransport.Dial = dialer.Dial
	lp.Client.Transport = httpTransport

Wrapper

Wrapper allows you to get ready-made structures

Wrapper for v3 https://pkg.go.dev/github.com/SevereCloud/vksdk/v2/longpoll-user/v3

Run and shutdown

TODO: write about lp.Run() and lp.Shutdown()


VK documentation https://vk.com/dev/using_longpoll
*/
package longpoll // import "github.com/SevereCloud/vksdk/v2/longpoll-user"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/object"
)

// Mode additional answer options.
type Mode = int

// A list of necessary option codes.
const (
	// receive attachments.
	ReceiveAttachments Mode = 1 << 1
	// receive more events.
	ExtendedEvents Mode = 1 << 3
	// receive pts (used in messages.getLongPollHistory).
	ReturnPts Mode = 1 << 5
	// extra fields in event type 8(friend become online).
	Code8ExtraFields Mode = 1 << 6
	// return random_id field.
	ReturnRandomID Mode = 1 << 7
)

// LongPoll struct.
type LongPoll struct {
	Key     string
	Server  string
	Ts      int
	Mode    Mode
	Version int
	Wait    int
	VK      *api.VK
	Client  *http.Client

	funcList             map[int][]EventNewFunc
	funcFullResponseList []func(object.LongPollResponse)
	inShutdown           int32
	goroutine            bool
}

// NewLongPoll returns a new LongPoll.
//
// The LongPoll will use the http.DefaultClient.
// This means that if the http.DefaultClient is modified by other components
// of your application the modifications will be picked up by the SDK as well.
func NewLongPoll(vk *api.VK, mode Mode) (*LongPoll, error) {
	lp := &LongPoll{
		VK:       vk,
		Mode:     mode,
		Version:  3,
		Wait:     25,
		funcList: make(FuncList),
		Client:   http.DefaultClient,
	}

	err := lp.updateServer(true)

	return lp, err
}

// Goroutine invoke functions in a goroutine.
func (lp *LongPoll) Goroutine(v bool) {
	lp.goroutine = v
}

func (lp *LongPoll) updateServer(updateTs bool) error {
	params := api.Params{
		"lp_version": lp.Version,
	}

	serverSetting, err := lp.VK.MessagesGetLongPollServer(params)
	if err != nil {
		return err
	}

	lp.Key = serverSetting.Key
	lp.Server = serverSetting.Server

	if updateTs {
		lp.Ts = serverSetting.Ts
	}

	return nil
}

func (lp *LongPoll) check() (response object.LongPollResponse, err error) {
	u := fmt.Sprintf(
		"https://%s?act=a_check&key=%s&ts=%d&wait=%d&mode=%d&version=%d",
		lp.Server,
		lp.Key,
		lp.Ts,
		lp.Wait,
		lp.Mode,
		lp.Version,
	)

	resp, err := lp.Client.Get(u)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return
	}

	err = lp.checkResponse(response)

	return
}

func (lp *LongPoll) checkResponse(response object.LongPollResponse) (err error) {
	switch response.Failed {
	case 0:
		lp.Ts = response.Ts
	case 1:
		lp.Ts = response.Ts
	case 2:
		err = lp.updateServer(false)
	case 3:
		err = lp.updateServer(true)
	case 4:
		err = ErrNotValidVersion
	default:
		err = &Failed{response.Failed}
	}

	return
}

func (lp LongPoll) handler(event []interface{}) error {
	key := int(event[0].(float64))

	for _, f := range lp.funcList[key] {
		if lp.goroutine {
			go func() { _ = f(event) }()
		} else {
			err := f(event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Run handler.
func (lp *LongPoll) Run() error {
	atomic.StoreInt32(&lp.inShutdown, 0)

	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		resp, err := lp.check()
		if err != nil {
			return err
		}

		for _, event := range resp.Updates {
			if err := lp.handler(event); err != nil {
				return err
			}
		}

		for _, f := range lp.funcFullResponseList {
			if lp.goroutine {
				go func() { f(resp) }()
			} else {
				f(resp)
			}
		}
	}

	return nil
}

// Shutdown gracefully shuts down the longpoll without interrupting any active connections.
func (lp *LongPoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// EventNew handler.
func (lp *LongPoll) EventNew(key int, f EventNewFunc) {
	lp.funcList[key] = append(lp.funcList[key], f)
}

// FullResponse handler.
func (lp *LongPoll) FullResponse(f func(object.LongPollResponse)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
