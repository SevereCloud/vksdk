package api_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/api"
	"github.com/stretchr/testify/assert"
)

func TestVK_LeadForms(t *testing.T) {
	t.Parallel()

	needGroupToken(t)
	needUserToken(t)

	params := api.Params{
		"group_id":        vkGroupID,
		"name":            "test",
		"title":           "test",
		"description":     "test",
		"questions":       `[{"type":"first_name"},{"type":"input","label":"Кличка кота"}]`,
		"policy_link_url": "ya.ru",
	}

	lf, err := vkUser.LeadFormsCreate(params)
	noError(t, err)
	assert.NotEmpty(t, lf.FormID)
	assert.NotEmpty(t, lf.URL)

	params["form_id"] = lf.FormID

	_, err = vkUser.LeadFormsUpdate(params)
	noError(t, err)
	assert.NotEmpty(t, lf.FormID)
	assert.NotEmpty(t, lf.URL)

	list, err := vkUser.LeadFormsList(api.Params{
		"group_id": vkGroupID,
	})
	noError(t, err)

	if assert.NotEmpty(t, list) {
		assert.NotEmpty(t, list[0].FormID)
		assert.NotEmpty(t, list[0].GroupID)
		// assert.NotEmpty(t, list[0].Photo)
		assert.NotEmpty(t, list[0].Name)
		assert.NotEmpty(t, list[0].Title)
		assert.NotEmpty(t, list[0].Description)
		// assert.NotEmpty(t, list[0].Confirmation)
		// assert.NotEmpty(t, list[0].SiteLinkURL)
		assert.NotEmpty(t, list[0].PolicyLinkURL)
		assert.NotEmpty(t, list[0].Questions)
		// assert.NotEmpty(t, list[0].Active)
		// assert.NotEmpty(t, list[0].LeadsCount)
		// assert.NotEmpty(t, list[0].PixelCode)
		// assert.NotEmpty(t, list[0].OncePerUser)
		// assert.NotEmpty(t, list[0].NotifyAdmins)
		// assert.NotEmpty(t, list[0].NotifyEmails)
		assert.NotEmpty(t, list[0].URL)
	}

	lfGet, err := vkUser.LeadFormsGet(api.Params{
		"group_id": vkGroupID,
		"form_id":  lf.FormID,
	})
	noError(t, err)
	assert.NotEmpty(t, lfGet.FormID)
	assert.NotEmpty(t, lfGet.GroupID)
	// assert.NotEmpty(t, lfGet.Photo)
	assert.NotEmpty(t, lfGet.Name)
	assert.NotEmpty(t, lfGet.Title)
	assert.NotEmpty(t, lfGet.Description)
	// assert.NotEmpty(t, lfGet.Confirmation)
	// assert.NotEmpty(t, lfGet.SiteLinkURL)
	assert.NotEmpty(t, lfGet.PolicyLinkURL)
	assert.NotEmpty(t, lfGet.Questions)
	// assert.NotEmpty(t, lfGet.Active)
	// assert.NotEmpty(t, lfGet.LeadsCount)
	// assert.NotEmpty(t, lfGet.PixelCode)
	// assert.NotEmpty(t, lfGet.OncePerUser)
	// assert.NotEmpty(t, lfGet.NotifyAdmins)
	// assert.NotEmpty(t, lfGet.NotifyEmails)
	assert.NotEmpty(t, lfGet.URL)

	_, err = vkUser.LeadFormsDelete(api.Params{
		"group_id": vkGroupID,
		"form_id":  lf.FormID,
	})
	noError(t, err)
}

// func TestVK_LeadFormsGetLeads(t *testing.T) {
// TODO: write test
// }

// func TestVK_LeadFormsGetUploadURL(t *testing.T) {
// TODO: write test
// }
