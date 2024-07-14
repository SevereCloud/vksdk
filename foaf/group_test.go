package foaf_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/SevereCloud/vksdk/v3/foaf"
	"github.com/stretchr/testify/assert"
)

func testGroup(t *testing.T, group foaf.Group) {
	t.Helper()

	assert.NotEmpty(t, group.Name)

	if assert.NotEmpty(t, group.URI) {
		assert.NotEmpty(t, group.URI[0].Resource)
	}

	// assert.NotEmpty(t, group.Img.Image.Primary)
	// assert.NotEmpty(t, group.Img.Image.Width)
	// assert.NotEmpty(t, group.Img.Image.Height)
	// assert.NotEmpty(t, group.Img.Image.About)
	// assert.NotEmpty(t, group.Img.Image.Thumbnail)

	assert.NotEmpty(t, group.Weblog.Resource)
	assert.NotEmpty(t, group.MembersCount)
}

func TestGetGroup_Group(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 1)
	if errors.Is(err, foaf.ErrorStatusCode{Code: http.StatusTooManyRequests}) {
		t.SkipNow()
	}

	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypeGroup, group.GroupType)
	testGroup(t, group)
}

func TestGetGroup_Public(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 29534144)
	if errors.Is(err, foaf.ErrorStatusCode{Code: http.StatusTooManyRequests}) {
		t.SkipNow()
	}

	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypePublic, group.GroupType)
	testGroup(t, group)
}

func TestGetGroup_Event(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 174832349)
	if errors.Is(err, foaf.ErrorStatusCode{Code: http.StatusTooManyRequests}) {
		t.SkipNow()
	}

	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypeEvent, group.GroupType)
	testGroup(t, group)

	assert.NotEmpty(t, group.StartDate)
	assert.NotEmpty(t, group.FinishDate)
}

func TestGetGroup_Private(t *testing.T) {
	t.Parallel()

	group, err := foaf.GetGroup(context.Background(), 184580855)
	if errors.Is(err, foaf.ErrorStatusCode{Code: http.StatusTooManyRequests}) {
		t.SkipNow()
	}

	assert.NoError(t, err)
	assert.Equal(t, foaf.GroupTypeGroup, group.GroupType)
}

func TestGetGroup_Error(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := foaf.GetGroup(ctx, 1)
	assert.Error(t, err, "context canceled")
}
