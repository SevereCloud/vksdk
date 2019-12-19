package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"
)

func TestDocsDoc_ToAttachment(t *testing.T) {
	f := func(doc object.DocsDoc, want string) {
		if got := doc.ToAttachment(); got != want {
			t.Errorf("DocsDoc.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.DocsDoc{ID: 10, OwnerID: 20}, "doc20_10")
	f(object.DocsDoc{ID: 20, OwnerID: -10}, "doc-10_20")
}
