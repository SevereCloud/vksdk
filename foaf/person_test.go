package foaf_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/foaf"
	"github.com/stretchr/testify/assert"
)

func TestGetPerson_Verified(t *testing.T) {
	f := func(userID int) {
		t.Helper()

		person, err := foaf.GetPerson(context.Background(), userID)
		assert.NoError(t, err)
		assert.Equal(t, foaf.AccessAllowed, person.PublicAccess)
		assert.Equal(t, foaf.ProfileStateVerified, person.ProfileState)
		assert.NotEmpty(t, person.FirstName)
		assert.NotEmpty(t, person.SecondName)
		assert.NotEmpty(t, person.Name)
		// assert.NotEmpty(t, person.Weblog.Title)
		assert.NotEmpty(t, person.Weblog.Resource)
		assert.NotEmpty(t, person.Created)
		assert.NotEmpty(t, person.Modified)
		assert.NotEmpty(t, person.SubscribersCount)
		assert.NotEmpty(t, person.Img.Image)
		assert.NotEmpty(t, person.Img.Image.Width)
		assert.NotEmpty(t, person.Img.Image.Height)
		assert.NotEmpty(t, person.Img.Image.About)
	}

	f(1)
	f(301419447)
}

func TestGetPerson_Banned(t *testing.T) {
	person, err := foaf.GetPerson(context.Background(), 540036751)
	assert.NoError(t, err)
	assert.Equal(t, foaf.ProfileStateBanned, person.ProfileState)
	assert.Equal(t, foaf.AccessAllowed, person.PublicAccess)
}

func TestGetPerson_Deleted(t *testing.T) {
	person, err := foaf.GetPerson(context.Background(), 3)
	assert.NoError(t, err)
	assert.Equal(t, foaf.ProfileStateDeleted, person.ProfileState)
	assert.Equal(t, foaf.AccessAllowed, person.PublicAccess)
}

func TestGetPerson_Active(t *testing.T) {
	person, err := foaf.GetPerson(context.Background(), 5024999)
	assert.NoError(t, err)
	assert.Equal(t, foaf.ProfileStateActive, person.ProfileState)
	assert.Equal(t, foaf.AccessDisallowed, person.PublicAccess)
}

func TestGetPerson(t *testing.T) {
	f := func(userID int) {
		t.Helper()

		_, err := foaf.GetPerson(context.Background(), userID)
		assert.NoError(t, err)
	}

	// FIXME: f(319698546) // check illegal character code U+000C
	f(177512888) // check invalid character entity &#5555555;
	f(135122185) // check invalid character entity &a (no semicolon)
}
