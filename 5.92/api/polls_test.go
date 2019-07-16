package api

import (
	"log"
	"strconv"
	"testing"
)

func TestVK_PollsCreate(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	poll, gotVkErr := vkUser.PollsCreate(map[string]string{
		"question":    "question",
		"add_answers": `["yes", "no", "maybe"]`,
	})
	if gotVkErr.Code != 0 {
		log.Fatal(gotVkErr)
	}

	_, gotVkErr = vkUser.PollsEdit(map[string]string{
		"poll_id":  strconv.Itoa(poll.ID),
		"question": "questionEdit",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PollsEdit() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PollsAddVote(map[string]string{
		"owner_id":   strconv.Itoa(poll.OwnerID),
		"poll_id":    strconv.Itoa(poll.ID),
		"answer_ids": strconv.Itoa(poll.Answers[0].ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PollsAddVote() gotVkErr = %v, want %v", gotVkErr, 0)
	}

	_, gotVkErr = vkUser.PollsDeleteVote(map[string]string{
		"owner_id":   strconv.Itoa(poll.OwnerID),
		"poll_id":    strconv.Itoa(poll.ID),
		"answer_ids": strconv.Itoa(poll.Answers[0].ID),
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PollsDeleteVote() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PollsGetBackgrounds(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PollsGetBackgrounds(map[string]string{})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PollsGetBackgrounds() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PollsGetByID(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}
	f := func(params map[string]string) {
		t.Helper()
		_, gotVkErr := vkUser.PollsGetByID(params)
		if gotVkErr.Code != 0 {
			t.Errorf("VK.PollsGetByID() gotVkErr = %v, want %v", gotVkErr, 0)
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

func TestVK_PollsGetVoters(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PollsGetVoters(map[string]string{
		"owner_id":   "-169097025",
		"poll_id":    "341032442",
		"answer_ids": "1144979948, 1144979949, 1144979950",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PollsGetVoters() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

func TestVK_PollsGetVotersFields(t *testing.T) {
	if vkUser.AccessToken == "" {
		t.Skip("USER_TOKEN empty")
	}

	_, gotVkErr := vkUser.PollsGetVotersFields(map[string]string{
		"owner_id":   "-169097025",
		"poll_id":    "341032442",
		"answer_ids": "1144979948, 1144979949, 1144979950",
		"fields":     "nickname, screen_name, sex",
	})
	if gotVkErr.Code != 0 {
		t.Errorf("VK.PollsGetVotersFields() gotVkErr = %v, want %v", gotVkErr, 0)
	}
}

// TODO: TestVK_PollsGetPhotoUploadServer
// TODO: TestVK_PollsSavePhoto
