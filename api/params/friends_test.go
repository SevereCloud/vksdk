package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestFriendsAddBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsAddBuilder()

	b.UserID(1)
	b.Text("text")
	b.Follow(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["text"], "text")
	assert.Equal(t, b.Params["follow"], true)
}

func TestFriendsAddListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsAddListBuilder()

	b.Name("text")
	b.UserIDs([]int{1})

	assert.Equal(t, b.Params["name"], "text")
	assert.Equal(t, b.Params["user_ids"], []int{1})
}

func TestFriendsAreFriendsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsAreFriendsBuilder()

	b.UserIDs([]int{1})
	b.NeedSign(true)

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["need_sign"], true)
}

func TestFriendsDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsDeleteBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestFriendsDeleteListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsDeleteListBuilder()

	b.ListID(1)

	assert.Equal(t, b.Params["list_id"], 1)
}

func TestFriendsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsEditBuilder()

	b.UserID(1)
	b.ListIDs([]int{1})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["list_ids"], []int{1})
}

func TestFriendsEditListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsEditListBuilder()

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

func TestFriendsGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetBuilder()

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

func TestFriendsGetByPhonesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetByPhonesBuilder()

	b.Phones([]string{"text"})
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["phones"], []string{"text"})
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestFriendsGetListsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetListsBuilder()

	b.UserID(1)
	b.ReturnSystem(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["return_system"], true)
}

func TestFriendsGetMutualBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetMutualBuilder()

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

func TestFriendsGetOnlineBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetOnlineBuilder()

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

func TestFriendsGetRecentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetRecentBuilder()

	b.Count(1)

	assert.Equal(t, b.Params["count"], 1)
}

func TestFriendsGetRequestsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetRequestsBuilder()

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

func TestFriendsGetSuggestionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetSuggestionsBuilder()

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

func TestFriendsSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsSearchBuilder()

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
