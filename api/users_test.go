package api_test

import (
	"errors"
	"testing"

	"github.com/SevereCloud/vksdk/v2/api"

	"github.com/stretchr/testify/assert"
)

func TestVK_UsersGet(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.UsersGet(api.Params{
		"user_ids": 1,
		"fields": []string{
			"photo_id",
			"verified",
			"sex",
			"bdate",
			"city",
			"country",
			"home_town",
			"has_photo",
			"photo_50",
			"photo_100",
			"photo_200_orig",
			"photo_200",
			"photo_400_orig",
			"photo_max",
			"photo_max_orig",
			"online",
			"domain",
			"has_mobile",
			"contacts",
			"site",
			"education",
			"universities",
			"schools",
			"status",
			"last_seen",
			"followers_count",
			"common_count",
			"occupation",
			"nickname",
			"relatives",
			"relation",
			"personal",
			"connections",
			"exports",
			"activities",
			"interests",
			"music",
			"movies",
			"tv",
			"books",
			"games",
			"about",
			"quotes",
			"can_post",
			"can_see_all_posts",
			"can_see_audio",
			"can_write_private_message",
			"can_send_friend_request",
			"is_favorite",
			"is_hidden_from_feed",
			"timezone",
			"screen_name",
			"maiden_name",
			"crop_photo",
			"is_friend",
			"friend_status",
			"career",
			"military",
			"blacklisted",
			"blacklisted_by_me",
			"can_be_invited_group",
		},
	})
	noError(t, err)

	if assert.NotEmpty(t, res) {
		for _, user := range res {
			assert.NotEmpty(t, user.ID)
			assert.NotEmpty(t, user.FirstName)
			assert.NotEmpty(t, user.LastName)
			assert.NotEmpty(t, user.Sex)
			assert.NotEmpty(t, user.Domain)
			assert.NotEmpty(t, user.ScreenName)
			assert.NotEmpty(t, user.Bdate)
			assert.NotEmpty(t, user.City)
			assert.NotEmpty(t, user.Country)
			assert.NotEmpty(t, user.Photo50)
			assert.NotEmpty(t, user.Photo100)
			assert.NotEmpty(t, user.Photo200)
			assert.NotEmpty(t, user.PhotoMax)
			assert.NotEmpty(t, user.Photo200Orig)
			assert.NotEmpty(t, user.Photo400Orig)
			assert.NotEmpty(t, user.PhotoMaxOrig)
			assert.NotEmpty(t, user.PhotoID)
			assert.NotEmpty(t, user.HasPhoto)
			assert.NotEmpty(t, user.HasMobile)
			// assert.NotEmpty(t, user.IsFriend)
			// assert.NotEmpty(t, user.FriendStatus)
			// assert.NotEmpty(t, user.Online)
			// assert.NotEmpty(t, user.CanPost)
			// assert.NotEmpty(t, user.CanSeeAllPosts)
			// assert.NotEmpty(t, user.CanSeeAudio)
			// assert.NotEmpty(t, user.CanWritePrivateMessage)
			// assert.NotEmpty(t, user.CanSendFriendRequest)
			assert.NotEmpty(t, user.Facebook)
			assert.NotEmpty(t, user.FacebookName)
			assert.NotEmpty(t, user.Twitter)
			assert.NotEmpty(t, user.Instagram)
			assert.NotEmpty(t, user.Site)
			assert.NotEmpty(t, user.Status)
			// assert.NotEmpty(t, user.LastSeen)
			assert.NotEmpty(t, user.CropPhoto)
			assert.NotEmpty(t, user.Verified)
			// assert.NotEmpty(t, user.CanBeInvitedGroup)
			assert.NotEmpty(t, user.FollowersCount)
			// assert.NotEmpty(t, user.Blacklisted)
			// assert.NotEmpty(t, user.BlacklistedByMe)
			// assert.NotEmpty(t, user.IsFavorite)
			// assert.NotEmpty(t, user.IsHiddenFromFeed)
			// assert.NotEmpty(t, user.CommonCount)
			// assert.NotEmpty(t, user.Career)
			// assert.NotEmpty(t, user.Military)
			// assert.NotEmpty(t, user.University)
			// assert.NotEmpty(t, user.UniversityName)
			// assert.NotEmpty(t, user.Faculty)
			// assert.NotEmpty(t, user.FacultyName)
			// assert.NotEmpty(t, user.Graduation)
			// assert.NotEmpty(t, user.HomeTown)
			// assert.NotEmpty(t, user.Relation)
			// if assert.NotEmpty(t, user.Personal) {
			// 	assert.NotEmpty(t, user.Personal.Alcohol)
			// 	assert.NotEmpty(t, user.Personal.Political)
			// 	assert.NotEmpty(t, user.Personal.Langs)
			// 	assert.NotEmpty(t, user.Personal.Religion)
			// 	assert.NotEmpty(t, user.Personal.InspiredBy)
			// 	assert.NotEmpty(t, user.Personal.PeopleMain)
			// 	assert.NotEmpty(t, user.Personal.LifeMain)
			// 	assert.NotEmpty(t, user.Personal.Smoking)
			// 	// assert.NotEmpty(t, user.Personal.ReligionID)
			// 	assert.NotEmpty(t, user.Personal.Alcohol)
			// }
			// assert.NotEmpty(t, user.About)
			// assert.NotEmpty(t, user.Relatives)
			// assert.NotEmpty(t, user.Quotes)
			assert.NotEmpty(t, user.Occupation)
		}
	}
}

func TestVK_UsersGetFollowers(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.UsersGetFollowers(api.Params{
		"user_id": 1,
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)

	_, err = vkUser.UsersGetFollowersFields(api.Params{
		"user_id": 1,
	})
	noError(t, err)
}

func TestVK_UsersGetSubscriptions(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.UsersGetSubscriptions(api.Params{
		"user_id": 117253521,
	})
	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_UsersReport(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.UsersReport(api.Params{
		"user_id": 1,
		"type":    "spam",
		"comment": "Тестовый репорт - github.com/SevereCloud/vksdk",
	})

	if errors.Is(err, api.ErrAccess) {
		t.Skip("Access denied")
	}

	noError(t, err)
	assert.NotEmpty(t, res)
}

func TestVK_UsersSearch(t *testing.T) {
	t.Parallel()

	needUserToken(t)

	res, err := vkUser.UsersSearch(api.Params{
		"q": "Vasya Babich",
		"fields": []string{
			"about",
			"activities",
			"bdate",
			"blacklisted_by_me",
			"blacklisted",
			"books",
			"can_be_invited_group",
			"can_post",
			"can_see_all_posts",
			"can_see_audio",
			"can_send_friend_request",
			"can_write_private_message",
			"career",
			"city",
			"common_count",
			"connections",
			"contacts",
			"counters",
			"country",
			"crop_photo",
			"domain",
			"education",
			"exports", // FIXME: VK not return
			"first_name_abl",
			"first_name_acc",
			"first_name_dat",
			"first_name_gen",
			"first_name_ins",
			"first_name_nom",
			"followers_count",
			"friend_status",
			"games",
			"has_mobile",
			"has_photo",
			"home_town",
			"interests",
			"is_favorite",
			"is_friend",
			"is_hidden_from_feed",
			"last_name_abl",
			"last_name_acc",
			"last_name_dat",
			"last_name_gen",
			"last_name_ins",
			"last_name_nom",
			"last_seen",
			"lists", // For friends.get only
			"maiden_name",
			"military",
			"movies",
			"music",
			"mutual",
			"nickname",
			"occupation",
			"online",
			"personal",
			"photo_100",
			"photo_200_orig",
			"photo_200",
			"photo_400_orig",
			"photo_50",
			"photo_id",
			"photo_max_orig",
			"photo_max",
			"quotes",
			"relation",
			"relatives",
			"schools",
			"screen_name",
			"sex",
			"site",
			"status",
			"timezone",
			"trending",
			"tv",
			"universities",
			"verified",
			"wall_default",
		},
	})
	noError(t, err)
	assert.NotEmpty(t, res.Count)
	assert.NotEmpty(t, res.Items)
}
