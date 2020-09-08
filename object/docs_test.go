package object_test

import (
	"reflect"
	"testing"

	"github.com/SevereCloud/vksdk/v2/object"
)

func TestDocsDoc_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(doc object.DocsDoc, want string) {
		if got := doc.ToAttachment(); got != want {
			t.Errorf("DocsDoc.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.DocsDoc{ID: 10, OwnerID: 20}, "doc20_10")
	f(object.DocsDoc{ID: 20, OwnerID: -10}, "doc-10_20")
}

func TestDocsDocPreviewPhoto_MaxSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.DocsDocPreviewPhoto, want object.DocsDocPreviewPhotoSizes) {
		t.Helper()

		if got := photo.MaxSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("DocsDocPreviewPhoto.MaxSize() = %v, want %v", got, want)
		}
	}

	f(object.DocsDocPreviewPhoto{}, object.DocsDocPreviewPhotoSizes{})
	f(object.DocsDocPreviewPhoto{
		Sizes: []object.DocsDocPreviewPhotoSizes{
			{
				Width: 10, Height: 20,
			},
			{
				Width: 100, Height: 200,
			},
		},
	}, object.DocsDocPreviewPhotoSizes{
		Width: 100, Height: 200,
	})
}

func TestDocsDocPreviewPhoto_MinSize(t *testing.T) {
	t.Parallel()

	f := func(photo object.DocsDocPreviewPhoto, want object.DocsDocPreviewPhotoSizes) {
		t.Helper()

		if got := photo.MinSize(); !reflect.DeepEqual(got, want) {
			t.Errorf("DocsDocPreviewPhoto.MinSize() = %v, want %v", got, want)
		}
	}

	f(object.DocsDocPreviewPhoto{}, object.DocsDocPreviewPhotoSizes{})
	f(object.DocsDocPreviewPhoto{
		Sizes: []object.DocsDocPreviewPhotoSizes{
			{
				Width: 10, Height: 20,
			},
			{
				Width: 100, Height: 200,
			},
		},
	}, object.DocsDocPreviewPhotoSizes{
		Width: 10, Height: 20,
	})
}
