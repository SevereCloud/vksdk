package longpoll // import "github.com/SevereCloud/vksdk/longpoll-bot"

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/SevereCloud/vksdk/object"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/events"
)

// Longpoll struct
type Longpoll struct {
	GroupID int
	Server  string
	Key     string
	Ts      string
	Wait    int
	VK      *api.VK
	Client  *http.Client

	funcFullResponseList []func(object.LongpollBotResponse)
	inShutdown           int32

	events.FuncList
}

// Init Longpoll
func Init(vk *api.VK, groupID int) (lp Longpoll, err error) {
	lp.VK = vk
	lp.GroupID = groupID
	lp.Wait = 25
	lp.Client = &http.Client{}
	err = lp.updateServer(true)

	return
}

func (lp *Longpoll) updateServer(updateTs bool) error {
	params := api.Params{
		"group_id": lp.GroupID,
	}

	serverSetting, err := lp.VK.GroupsGetLongPollServer(params)
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

func (lp *Longpoll) check() ([]object.GroupEvent, error) {
	var response object.LongpollBotResponse

	u := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.Ts, lp.Wait)

	resp, err := lp.Client.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	err = lp.checkResponse(response)

	return response.Updates, err
}

func (lp *Longpoll) checkResponse(response object.LongpollBotResponse) (err error) {
	switch response.Failed {
	case 0:
		lp.Ts = response.Ts
	case 1:
		lp.Ts = response.Ts
	case 2:
		err = lp.updateServer(false)
	case 3:
		err = lp.updateServer(true)
	default:
		err = fmt.Errorf(`"failed":%d`, response.Failed)
	}

	return
}

// Run handler
func (lp *Longpoll) Run() error {
	atomic.StoreInt32(&lp.inShutdown, 0)

	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		events, err := lp.check()
		if err != nil {
			return err
		}

		for _, event := range events {
			err = lp.Handler(event)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Shutdown gracefully shuts down the longpoll without interrupting any active connections.
func (lp *Longpoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// FullResponse handler
func (lp *Longpoll) FullResponse(f func(object.LongpollBotResponse)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
