package params_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestPollsAddVoteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsAddVoteBuilder()

	b.OwnerID(1)
	b.PollID(1)
	b.AnswerIDs([]int{1})
	b.IsBoard(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["poll_id"], 1)
	assert.Equal(t, b.Params["answer_ids"], []int{1})
	assert.Equal(t, b.Params["is_board"], true)
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

	assert.Equal(t, b.Params["question"], "text")
	assert.Equal(t, b.Params["is_anonymous"], true)
	assert.Equal(t, b.Params["is_multiple"], true)
	assert.Equal(t, b.Params["end_date"], 1)
	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["add_answers"], "text")
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["background_id"], "text")
}

func TestPollsDeleteVoteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewPollsDeleteVoteBuilder()

	b.OwnerID(1)
	b.PollID(1)
	b.AnswerID(1)
	b.IsBoard(true)

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["poll_id"], 1)
	assert.Equal(t, b.Params["answer_id"], 1)
	assert.Equal(t, b.Params["is_board"], true)
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["poll_id"], 1)
	assert.Equal(t, b.Params["question"], "text")
	assert.Equal(t, b.Params["add_answers"], "text")
	assert.Equal(t, b.Params["edit_answers"], "text")
	assert.Equal(t, b.Params["delete_answers"], "text")
	assert.Equal(t, b.Params["end_date"], 1)
	assert.Equal(t, b.Params["photo_id"], 1)
	assert.Equal(t, b.Params["background_id"], "text")
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["is_board"], true)
	assert.Equal(t, b.Params["poll_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["friends_count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"text"})
	assert.Equal(t, b.Params["name_case"], "text")
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

	assert.Equal(t, b.Params["owner_id"], 1)
	assert.Equal(t, b.Params["poll_id"], 1)
	assert.Equal(t, b.Params["answer_ids"], []int{1})
	assert.Equal(t, b.Params["is_board"], true)
	assert.Equal(t, b.Params["friends_only"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}
