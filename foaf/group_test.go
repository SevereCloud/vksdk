package foaf_test

import (
	"context"
	"testing"

	"github.com/SevereCloud/vksdk/v2/foaf"
	"github.com/stretchr/testify/assert"
)

func testGroup(t *testing.T, group foaf.Group) {
	assert.NotEmpty(t, group.Name)

	if assert.NotEmpty(t, group.URI) {
		assert.NotEmpty(t, group.URI[0].Resource)
	}

	assert.NotEmpty(t, group.Img.Image.Primary)
	// assert.NotEmpty(t, group.Img.Image.Width)
	// assert.NotEmpty(t, group.Img.Image.Height)
	assert.NotEmpty(t, group.Img.Image.About)
	assert.NotEmpty(t, group.Img.Image.Thumbnail)

	assert.NotEmpty(t, group.Weblog.Resource)
	assert.NotEmpty(t, group.MembersCount)
}

func TestGetGroup_Group(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypeGroup, group.GroupType)
	testGroup(t, group)
}

func TestGetGroup_Public(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 29534144)
	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypePublic, group.GroupType)
	testGroup(t, group)
}

func TestGetGroup_Event(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 86529522)
	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypeEvent, group.GroupType)
	testGroup(t, group)

	assert.NotEmpty(t, group.StartDate)
	assert.NotEmpty(t, group.FinishDate)
}

func TestGetGroup_Private(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 184580855)
	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypeGroup, group.GroupType)
}
