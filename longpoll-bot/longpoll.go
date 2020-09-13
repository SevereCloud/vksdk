/*
Package longpoll implements Bots Long Poll API.

See more https://vk.com/dev/bots_longpoll
*/
package longpoll // import "github.com/SevereCloud/vksdk/v2/longpoll-bot"

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/SevereCloud/vksdk/v2"
	"github.com/SevereCloud/vksdk/v2/internal"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/events"
)

// Response struct.
type Response struct {
	Ts      string              `json:"ts"`
	Updates []events.GroupEvent `json:"updates"`
	Failed  int                 `json:"failed"`
}

// LongPoll struct.
type LongPoll struct {
	GroupID int
	Server  string
	Key     string
	Ts      string
	Wait    int
	VK      *api.VK
	Client  *http.Client

	funcFullResponseList []func(Response)
	inShutdown           int32

	events.FuncList
}

// NewLongPoll returns a new LongPoll.
//
// The LongPoll will use the http.DefaultClient.
// This means that if the http.DefaultClient is modified by other components
// of your application the modifications will be picked up by the SDK as well.
func NewLongPoll(vk *api.VK, groupID int) (*LongPoll, error) {
	lp := &LongPoll{
		VK:      vk,
		GroupID: groupID,
		Wait:    25,
		Client:  http.DefaultClient,
	}
	lp.FuncList = *events.NewFuncList()

	err := lp.updateServer(true)

	return lp, err
}

// NewLongPollCommunity returns a new LongPoll for community token.
//
// The LongPoll will use the http.DefaultClient.
// This means that if the http.DefaultClient is modified by other components
// of your application the modifications will be picked up by the SDK as well.
func NewLongPollCommunity(vk *api.VK) (*LongPoll, error) {
	resp, err := vk.GroupsGetByID(nil)
	if err != nil {
		return nil, err
	}

	lp := &LongPoll{
		VK:      vk,
		GroupID: resp[0].ID,
		Wait:    25,
		Client:  http.DefaultClient,
	}
	lp.FuncList = *events.NewFuncList()

	err = lp.updateServer(true)

	return lp, err
}

func (lp *LongPoll) updateServer(updateTs bool) error {
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

func (lp *LongPoll) check() (Response, error) {
	var response Response

	u := fmt.Sprintf("%s?act=a_check&key=%s&ts=%s&wait=%d", lp.Server, lp.Key, lp.Ts, lp.Wait)

	resp, err := lp.Client.Get(u)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	err = lp.checkResponse(response)

	return response, err
}

func (lp *LongPoll) checkResponse(response Response) (err error) {
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
		err = &Failed{response.Failed}
	}

	return
}

func (lp *LongPoll) autoSetting() error {
	params := api.Params{
		"group_id":    lp.GroupID,
		"enabled":     true,
		"api_version": vksdk.API,
	}
	for _, event := range lp.ListEvents() {
		params[string(event)] = true
	}

	// Updating LongPoll settings
	_, err := lp.VK.GroupsSetLongPollSettings(params)

	return err
}

// Run handler.
func (lp *LongPoll) Run() error {
	atomic.StoreInt32(&lp.inShutdown, 0)

	err := lp.autoSetting()
	if err != nil {
		return err
	}

	for atomic.LoadInt32(&lp.inShutdown) == 0 {
		resp, err := lp.check()
		if err != nil {
			return err
		}

		ctx := context.WithValue(context.Background(), internal.LongPollTsKey, resp.Ts)

		for _, event := range resp.Updates {
			err = lp.Handler(ctx, event)
			if err != nil {
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
func (lp *LongPoll) Shutdown() {
	atomic.StoreInt32(&lp.inShutdown, 1)
}

// FullResponse handler.
func (lp *LongPoll) FullResponse(f func(Response)) {
	lp.funcFullResponseList = append(lp.funcFullResponseList, f)
}
