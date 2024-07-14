package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v3/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPollsAddVoteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsAddVoteBuilder()

	b.OwnerID(1)
	b.PollID(1)
	b.AnswerIDs([]int{1})
	b.IsBoard(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["poll_id"])
	assert.Equal(t, []int{1}, b.Params["answer_ids"])
	assert.Equal(t, true, b.Params["is_board"])
}

func TestPollsCreateBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsCreateBuilder()

	b.Question("text")
	b.IsAnonymous(true)
	b.IsMultiple(true)
	b.EndDate(1)
	b.OwnerID(1)
	b.AddAnswers("text")
	b.PhotoID(1)
	b.BackgroundID("text")

	assert.Equal(t, "text", b.Params["question"])
	assert.Equal(t, true, b.Params["is_anonymous"])
	assert.Equal(t, true, b.Params["is_multiple"])
	assert.Equal(t, 1, b.Params["end_date"])
	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, "text", b.Params["add_answers"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, "text", b.Params["background_id"])
}

func TestPollsDeleteVoteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsDeleteVoteBuilder()

	b.OwnerID(1)
	b.PollID(1)
	b.AnswerID(1)
	b.IsBoard(true)

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["poll_id"])
	assert.Equal(t, 1, b.Params["answer_id"])
	assert.Equal(t, true, b.Params["is_board"])
}

func TestPollsEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsEditBuilder()

	b.OwnerID(1)
	b.PollID(1)
	b.Question("text")
	b.AddAnswers("text")
	b.EditAnswers("text")
	b.DeleteAnswers("text")
	b.EndDate(1)
	b.PhotoID(1)
	b.BackgroundID("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["poll_id"])
	assert.Equal(t, "text", b.Params["question"])
	assert.Equal(t, "text", b.Params["add_answers"])
	assert.Equal(t, "text", b.Params["edit_answers"])
	assert.Equal(t, "text", b.Params["delete_answers"])
	assert.Equal(t, 1, b.Params["end_date"])
	assert.Equal(t, 1, b.Params["photo_id"])
	assert.Equal(t, "text", b.Params["background_id"])
}

func TestPollsGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsGetByIDBuilder()

	b.OwnerID(1)
	b.IsBoard(true)
	b.PollID(1)
	b.Extended(true)
	b.FriendsCount(1)
	b.Fields([]string{"text"})
	b.NameCase("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, true, b.Params["is_board"])
	assert.Equal(t, 1, b.Params["poll_id"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["friends_count"])
	assert.Equal(t, []string{"text"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
}

func TestPollsGetVotersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsGetVotersBuilder()

	b.OwnerID(1)
	b.PollID(1)
	b.AnswerIDs([]int{1})
	b.IsBoard(true)
	b.FriendsOnly(true)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, 1, b.Params["owner_id"])
	assert.Equal(t, 1, b.Params["poll_id"])
	assert.Equal(t, []int{1}, b.Params["answer_ids"])
	assert.Equal(t, true, b.Params["is_board"])
	assert.Equal(t, true, b.Params["friends_only"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
}
