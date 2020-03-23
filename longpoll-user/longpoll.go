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

	vk := api.Init("<TOKEN>")

And then longpoll

	mode := longpoll.ReceiveAttachments + longpoll.ExtendedEvents
	lp, err := longpoll.Init(vk, mode)

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

Wrapper for v3 https://pkg.go.dev/github.com/SevereCloud/vksdk/longpoll-user/v3

Run and shutdown

TODO: write about lp.Run() and lp.Shutdown()


VK documentation https://vk.com/dev/using_longpoll
*/
package longpoll // import "github.com/SevereCloud/vksdk/longpoll-user"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/object"
)

// Mode additional answer options
type Mode = int

// A list of necessary option codes
const (
	// receive attachments
	ReceiveAttachments Mode = 1 << 1
	// receive more events
	ExtendedEvents Mode = 1 << 3
	// receive pts (used in messages.getLongPollHistory)
	ReturnPts Mode = 1 << 5
	// extra fields in event type 8(friend become online)
	Code8ExtraFields Mode = 1 << 6
	// return random_id field
	ReturnRandomID Mode = 1 << 7
)

// Longpoll struct
type Longpoll struct {
	Key     string
	Server  string
	Ts      int
	Mode    Mode
	Version int
	Wait    int
	VK      *api.VK
	Client  *http.Client

	funcList             FuncList
	funcFullResponseList []func(object.LongpollResponse)
	inShutdown           int32
}

// Init Longpoll
//
// TODO: v2 NewLongpoll
// FIXME: v2 return *Longpoll
func Init(vk *api.VK, mode Mode) (lp Longpoll, err error) {
	// NOTE: what about group_id?
	lp.VK = vk
	lp.Mode = mode
	lp.Version = 3
	lp.Wait = 25
	lp.funcList = make(FuncList)
	lp.Client = &http.Client{}
	err = lp.updateServer(true)

	return
}

func (lp *Longpoll) updateServer(updateTs bool) error {
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

func (lp *Longpoll) check() (response object.LongpollResponse, err error) {
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

func (lp *Longpoll) checkResponse(response object.LongpollResponse) (err error) {
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
		err = fmt.Errorf("not valid version")
	default:
		err = fmt.Errorf(`"failed":%d`, response.Failed)
	}

	return
}

// Run handler
func (lp *Longpoll) Run() error {
	atomic.StoreInt32(&lp.inShutdown, 0)

	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		resp, err := lp.check()
		if err != nil {
			return err
		}

		for _, event := range resp.Updates {
			if err := lp.funcList.Handler(event); err != nil {
				return err
			}
		}

		for _, f := range lp.funcFullResponseList {
			f(resp)
		}
	}

	return nil
}

// Shutdown gracefully shuts down the longpoll without interrupting any active connections.
func (lp *Longpoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// EventNew handler
func (lp *Longpoll) EventNew(key int, f EventNewFunc) {
	lp.funcList[key] = append(lp.funcList[key], f)
}

// FullResponse handler
func (lp *Longpoll) FullResponse(f func(object.LongpollResponse)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
