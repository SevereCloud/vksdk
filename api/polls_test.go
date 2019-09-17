package api

import (
	"log"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVK_PollsCreate(t *testing.T) {
	needUserToken(t)

	poll, err := vkUser.PollsCreate(map[string]string{
		"question":    "question",
		"add_answers": `["yes", "no", "maybe"]`,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = vkUser.PollsEdit(map[string]string{
		"poll_id":  strconv.Itoa(poll.ID),
		"question": "questionEdit",
	})
	assert.NoError(t, err)

	_, err = vkUser.PollsAddVote(map[string]string{
		"owner_id":   strconv.Itoa(poll.OwnerID),
		"poll_id":    strconv.Itoa(poll.ID),
		"answer_ids": strconv.Itoa(poll.Answers[0].ID),
	})
	assert.NoError(t, err)

	_, err = vkUser.PollsDeleteVote(map[string]string{
		"owner_id":   strconv.Itoa(poll.OwnerID),
		"poll_id":    strconv.Itoa(poll.ID),
		"answer_ids": strconv.Itoa(poll.Answers[0].ID),
	})
	assert.NoError(t, err)
}

func TestVK_PollsGetBackgrounds(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetBackgrounds(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_PollsGetByID(t *testing.T) {
	needUserToken(t)
	f := func(params map[string]string) {
		t.Helper()
		_, err := vkUser.PollsGetByID(params)
		if err != nil {
			t.Errorf("VK.PollsGetByID() err = %v", err)
		}
	}

	f(map[string]string{
		"owner_id": "-2158488",
		"poll_id":  "338827002",
	})
	f(map[string]string{
		"owner_id": "-169097025",
		"poll_id":  "341032442",
	})
}

func TestVK_PollsGetPhotoUploadServer(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetPhotoUploadServer(map[string]string{})
	assert.NoError(t, err)
}

func TestVK_PollsGetVoters(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetVoters(map[string]string{
		"owner_id":   "-169097025",
		"poll_id":    "341032442",
		"answer_ids": "1144979948, 1144979949, 1144979950",
	})
	assert.NoError(t, err)
}

func TestVK_PollsGetVotersFields(t *testing.T) {
	needUserToken(t)

	_, err := vkUser.PollsGetVotersFields(map[string]string{
		"owner_id":   "-169097025",
		"poll_id":    "341032442",
		"answer_ids": "1144979948, 1144979949, 1144979950",
		"fields":     "nickname, screen_name, sex",
	})
	assert.NoError(t, err)
}

// TODO: TestVK_PollsSavePhoto
