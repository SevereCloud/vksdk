package api

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PollsCreate(t *testing.T) {
	needUserToken(t)

	poll, err := vkUser.PollsCreate(Params{
		"question":    "question",
		"add_answers": `["yes", "no", "maybe"]`,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = vkUser.PollsEdit(Params{
		"poll_id":  poll.ID,
		"question": "questionEdit",
	})
	assert.NoError(t, err)

	_, err = vkUser.PollsAddVote(Params{
		"owner_id":   poll.OwnerID,
		"poll_id":    poll.ID,
		"answer_ids": poll.Answers[0].ID,
	})
	assert.NoError(t, err)

	_, err = vkUser.PollsDeleteVote(Params{
		"owner_id":   poll.OwnerID,
		"poll_id":    poll.ID,
		"answer_ids": poll.Answers[0].ID,
	})
	assert.NoError(t, err)
}

func TestVK_PollsGetBackgrounds(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetBackgrounds(Params{})
	assert.NoError(t, err)
}

func TestVK_PollsGetByID(t *testing.T) {
	needUserToken(t)

	f := func(params Params) {
		t.Helper()

		_, err := vkUser.PollsGetByID(params)
		if err != nil {
			t.Errorf("VK.PollsGetByID() err = %v", err)
		}
	}

	f(Params{
		"owner_id": -2158488,
		"poll_id":  338827002,
	})
	f(Params{
		"owner_id": -169097025,
		"poll_id":  341032442,
	})
}

func TestVK_PollsGetPhotoUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetPhotoUploadServer(Params{})
	assert.NoError(t, err)
}

func TestVK_PollsGetVoters(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetVoters(Params{
		"owner_id":   -169097025,
		"poll_id":    341032442,
		"answer_ids": []int{1144979948, 1144979949, 1144979950},
	})
	assert.NoError(t, err)
}

func TestVK_PollsGetVotersFields(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetVotersFields(Params{
		"owner_id":   -169097025,
		"poll_id":    341032442,
		"answer_ids": []int{1144979948, 1144979949, 1144979950},
		"fields":     "nickname, screen_name, sex",
	})
	assert.NoError(t, err)
}

// TODO: TestVK_PollsSavePhoto
