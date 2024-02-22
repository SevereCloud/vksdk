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

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["text"])
	assert.Equal(t, true, b.Params["follow"])
}

func TestFriendsAddListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsAddListBuilder()

	b.Name("text")
	b.UserIDs([]int{1})

	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
}

func TestFriendsAreFriendsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsAreFriendsBuilder()

	b.UserIDs([]int{1})
	b.NeedSign(true)

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, true, b.Params["need_sign"])
}

func TestFriendsDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsDeleteBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}

func TestFriendsDeleteListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsDeleteListBuilder()

	b.ListID(1)

	assert.Equal(t, 1, b.Params["list_id"])
}

func TestFriendsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsEditBuilder()

	b.UserID(1)
	b.ListIDs([]int{1})

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, []int{1}, b.Params["list_ids"])
}

func TestFriendsEditListBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsEditListBuilder()

	b.Name("text")
	b.ListID(1)
	b.UserIDs([]int{1})
	b.AddUserIDs([]int{1})
	b.DeleteUserIDs([]int{1})

	assert.Equal(t, "text", b.Params["name"])
	assert.Equal(t, 1, b.Params["list_id"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, []int{1}, b.Params["add_user_ids"])
	assert.Equal(t, []int{1}, b.Params["delete_user_ids"])
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

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["order"])
	assert.Equal(t, 1, b.Params["list_id"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
	assert.Equal(t, "text", b.Params["ref"])
}

func TestFriendsGetByPhonesBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetByPhonesBuilder()

	b.Phones([]string{"text"})
	b.Fields([]string{"test"})

	assert.Equal(t, []string{"text"}, b.Params["phones"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestFriendsGetListsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetListsBuilder()

	b.UserID(1)
	b.ReturnSystem(true)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["return_system"])
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

	assert.Equal(t, 1, b.Params["source_uid"])
	assert.Equal(t, 1, b.Params["target_uid"])
	assert.Equal(t, []int{1}, b.Params["target_uids"])
	assert.Equal(t, "text", b.Params["order"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
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

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["list_id"])
	assert.Equal(t, true, b.Params["online_mobile"])
	assert.Equal(t, "text", b.Params["order"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
}

func TestFriendsGetRecentBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetRecentBuilder()

	b.Count(1)

	assert.Equal(t, 1, b.Params["count"])
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

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, true, b.Params["need_mutual"])
	assert.Equal(t, true, b.Params["out"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, true, b.Params["need_viewed"])
	assert.Equal(t, true, b.Params["suggested"])
	assert.Equal(t, "text", b.Params["ref"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestFriendsGetSuggestionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewFriendsGetSuggestionsBuilder()

	b.Filter([]string{"text"})
	b.Count(1)
	b.Offset(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, []string{"text"}, b.Params["filter"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
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

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
}
