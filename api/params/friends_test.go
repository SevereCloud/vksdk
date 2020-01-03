package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestFriendsAddBulder(t *testing.T) {
	b := params.NewFriendsAddBulder()

	b.UserID(1)
	b.Text("text")
	b.Follow(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["follow"], true)
}

func TestFriendsAddListBulder(t *testing.T) {
	b := params.NewFriendsAddListBulder()

	b.Name("text")
	b.UserIDs([]int{1})

	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["user_ids"], []int{1})
}

func TestFriendsAreFriendsBulder(t *testing.T) {
	b := params.NewFriendsAreFriendsBulder()

	b.UserIDs([]int{1})
	b.NeedSign(true)

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["need_sign"], true)
}

func TestFriendsDeleteBulder(t *testing.T) {
	b := params.NewFriendsDeleteBulder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestFriendsDeleteListBulder(t *testing.T) {
	b := params.NewFriendsDeleteListBulder()

	b.ListID(1)

	assert.Equal(t, b.Params["list_id"], 1)
}

func TestFriendsEditBulder(t *testing.T) {
	b := params.NewFriendsEditBulder()

	b.UserID(1)
	b.ListIDs([]int{1})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["list_ids"], []int{1})
}

func TestFriendsEditListBulder(t *testing.T) {
	b := params.NewFriendsEditListBulder()

	b.Name("text")
	b.ListID(1)
	b.UserIDs([]int{1})
	b.AddUserIDs([]int{1})
	b.DeleteUserIDs([]int{1})

	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["list_id"], 1)
	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["add_user_ids"], []int{1})
	assert.Equal(t, b.Params["delete_user_ids"], []int{1})
}

func TestFriendsGetBulder(t *testing.T) {
	b := params.NewFriendsGetBulder()

	b.UserID(1)
	b.Order("text")
	b.ListID(1)
	b.Count(1)
	b.Offset(1)
	b.Fields([]string{"test"})
	b.NameCase("text")
	b.Ref("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["order"], "text")
	assert.Equal(t, b.Params["list_id"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
	assert.Equal(t, b.Params["ref"], "text")
}

func TestFriendsGetByPhonesBulder(t *testing.T) {
	b := params.NewFriendsGetByPhonesBulder()

	b.Phones([]string{"text"})
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["phones"], []string{"text"})
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestFriendsGetListsBulder(t *testing.T) {
	b := params.NewFriendsGetListsBulder()

	b.UserID(1)
	b.ReturnSystem(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["return_system"], true)
}

func TestFriendsGetMutualBulder(t *testing.T) {
	b := params.NewFriendsGetMutualBulder()

	b.SourceUID(1)
	b.TargetUID(1)
	b.TargetUids([]int{1})
	b.Order("text")
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["source_uid"], 1)
	assert.Equal(t, b.Params["target_uid"], 1)
	assert.Equal(t, b.Params["target_uids"], []int{1})
	assert.Equal(t, b.Params["order"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestFriendsGetOnlineBulder(t *testing.T) {
	b := params.NewFriendsGetOnlineBulder()

	b.UserID(1)
	b.ListID(1)
	b.OnlineMobile(true)
	b.Order("text")
	b.Count(1)
	b.Offset(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["list_id"], 1)
	assert.Equal(t, b.Params["online_mobile"], true)
	assert.Equal(t, b.Params["order"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
}

func TestFriendsGetRecentBulder(t *testing.T) {
	b := params.NewFriendsGetRecentBulder()

	b.Count(1)

	assert.Equal(t, b.Params["count"], 1)
}

func TestFriendsGetRequestsBulder(t *testing.T) {
	b := params.NewFriendsGetRequestsBulder()

	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.NeedMutual(true)
	b.Out(true)
	b.Sort(1)
	b.NeedViewed(true)
	b.Suggested(true)
	b.Ref("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["need_mutual"], true)
	assert.Equal(t, b.Params["out"], true)
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["need_viewed"], true)
	assert.Equal(t, b.Params["suggested"], true)
	assert.Equal(t, b.Params["ref"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestFriendsGetSuggestionsBulder(t *testing.T) {
	b := params.NewFriendsGetSuggestionsBulder()

	b.Filter([]string{"text"})
	b.Count(1)
	b.Offset(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, b.Params["filter"], []string{"text"})
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}

func TestFriendsSearchBulder(t *testing.T) {
	b := params.NewFriendsSearchBulder()

	b.UserID(1)
	b.Q("text")
	b.Fields([]string{"test"})
	b.NameCase("text")
	b.Offset(1)
	b.Count(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
}
