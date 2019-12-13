package object

import "testing"

func TestDocsDoc_ToAttachment(t *testing.T) {
	type fields struct {
		ID      int
		OwnerID int
	}

	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "doc20_10",
			fields: fields{10, 20},
			want:   "doc20_10",
		},
		{
			name:   "doc-10_20",
			fields: fields{20, -10},
			want:   "doc-10_20",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			doc := DocsDoc{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := doc.ToAttachment(); got != tt.want {
				t.Errorf("DocsDoc.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
