package api_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_FriendsAddDelete(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	resAdd, err := vkUser.FriendsAdd(api.Params{
		"user_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, resAdd)

	resDelete, err := vkUser.FriendsDelete(api.Params{
		"user_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, resDelete.Success)
	assert.NotEmpty(t, resDelete.OutRequestDeleted)
}

func TestVK_FriendsList(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	list, err := vkUser.FriendsAddList(api.Params{
		"name": "new list",
	})
	noError(t, err)
	assert.NotEmpty(t, list.ListID)

	res, err := vkUser.FriendsEditList(api.Params{
		"list_id": list.ListID,
		"name":    "edit list",
	})
	noError(t, err)
	assert.NotEmpty(t, res)

	lists, err := vkUser.FriendsGetLists(nil)
	noError(t, err)
	assert.NotEmpty(t, lists.Count)
	assert.NotEmpty(t, lists.Items)

	res, err = vkUser.FriendsDeleteList(api.Params{
		"list_id": list.ListID,
		"name":    "edit list",
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FriendsAreFriends(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.FriendsAreFriends(api.Params{
		"user_ids": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FriendsDeleteAllRequests(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.FriendsDeleteAllRequests(nil)
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FriendsEdit(t *testing.T) {
	t.Parallel()

	// NOTE: need user friends
	// NOTE: https://vk.com/bug191897
	needUserToken(t)

	_, _ = vkUser.FriendsEdit(nil)
}

func TestVK_FriendsGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.FriendsGet(api.Params{
		"user_id": 6492,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	resFields, err := vkUser.FriendsGetFields(api.Params{
		"user_id": 6492,
		"count":   10,
	})
	noError(t, err)
	assert.NotEmpty(t, resFields.Count)
	assert.NotEmpty(t, resFields.Items)
}

func TestVK_FriendsGetAppUsers(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.FriendsGetAppUsers(nil)
	noError(t, err)
}

func TestVK_FriendsGetByPhones(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.FriendsGetByPhones(api.Params{
		"phones": "79215686249",
	})
	noError(t, err)
}

func TestVK_FriendsGetMutual(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.FriendsGetMutual(api.Params{
		"source_uid": 6492,
		"target_uid": 2745,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_FriendsGetOnline(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.FriendsGetOnline(nil)
	noError(t, err)

	_, err = vkUser.FriendsGetOnlineOnlineMobile(nil)
	noError(t, err)
}

func TestVK_FriendsGetRecent(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.FriendsGetRecent(nil)
	noError(t, err)
}

func TestVK_FriendsGetRequests(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	_, err := vkUser.FriendsGetRequests(api.Params{
		"need_mutual": true,
		"extended":    true,
		"fields":      "sex",
	})
	noError(t, err)

	_, err = vkUser.FriendsGetRequestsNeedMutual(api.Params{
		"need_mutual": false,
		"extended":    true,
		"fields":      "sex",
	})
	noError(t, err)

	_, err = vkUser.FriendsGetRequestsExtended(api.Params{
		"need_mutual": true,
		"extended":    false,
		"fields":      "sex",
	})
	noError(t, err)
}

func TestVK_FriendsGetSuggestions(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.FriendsGetSuggestions(nil)
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}

func TestVK_FriendsSearch(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.FriendsSearch(api.Params{
		"user_id": 2314852,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}
