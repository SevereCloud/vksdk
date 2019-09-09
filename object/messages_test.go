package object

import (
	"testing"
)

func TestMessagesKeyboard_AddRow(t *testing.T) {
	var keyboard MessagesKeyboard
	t.Run("add 1 row", func(t *testing.T) {
		keyboard.AddRow()
		if len(keyboard.Buttons) != 1 {
			t.Error("Keyboard len != 1")
		}
	})
	t.Run("add 2 row", func(t *testing.T) {
		keyboard.AddRow()
		if len(keyboard.Buttons) != 2 {
			t.Error("Keyboard len != 2")
		}
	})
}

func TestMessagesKeyboard_AddButton(t *testing.T) {
	const (
		label   = "label"
		payload = "payload"
		color   = "color"
	)
	var keyboard MessagesKeyboard
	keyboard.AddRow()

	t.Run("Add button", func(t *testing.T) {
		keyboard.AddButton(label, payload, color)
		if keyboard.Buttons[0][0].Color != color {
			t.Error("Bad button color")
		}
		if keyboard.Buttons[0][0].Action.Label != label {
			t.Error("Bad button label")
		}
		if keyboard.Buttons[0][0].Action.Payload != payload {
			t.Error("Bad button payload")
		}
	})
}

func TestMessagesKeyboard_AddTextButton(t *testing.T) {
	const (
		label   = "label"
		payload = "payload"
		color   = "color"
	)
	var keyboard MessagesKeyboard
	keyboard.AddRow()

	t.Run("Add Text button", func(t *testing.T) {
		keyboard.AddTextButton(label, payload, color)
		if keyboard.Buttons[0][0].Color != color {
			t.Error("Bad button color")
		}
		if keyboard.Buttons[0][0].Action.Label != label {
			t.Error("Bad button label")
		}
		if keyboard.Buttons[0][0].Action.Payload != payload {
			t.Error("Bad button payload")
		}
	})
}

func TestMessagesKeyboard_AddLocationButton(t *testing.T) {
	const payload = "payload"
	var keyboard MessagesKeyboard
	keyboard.AddRow()

	t.Run("Add Location button", func(t *testing.T) {
		keyboard.AddLocationButton(payload)
		if keyboard.Buttons[0][0].Action.Payload != payload {
			t.Error("Bad button payload")
		}
	})
}

func TestMessagesKeyboard_AddVKPayButton(t *testing.T) {
	const (
		payload = "payload"
		hash    = "hash"
	)
	var keyboard MessagesKeyboard
	keyboard.AddRow()

	t.Run("Add VK Pay button", func(t *testing.T) {
		keyboard.AddVKPayButton(payload, hash)
		if keyboard.Buttons[0][0].Action.Payload != payload {
			t.Error("Bad button payload")
		}
		if keyboard.Buttons[0][0].Action.Hash != hash {
			t.Error("Bad button hash")
		}
	})
}

func TestMessagesKeyboard_AddVKAppsButton(t *testing.T) {
	const (
		appID   = 1
		ownerID = 2
		payload = "payload"
		label   = "label"
		hash    = "hash"
	)
	var keyboard MessagesKeyboard
	keyboard.AddRow()

	t.Run("Add VK Apps button", func(t *testing.T) {
		keyboard.AddVKAppsButton(appID, ownerID, payload, label, hash)
		if keyboard.Buttons[0][0].Action.AppID != appID {
			t.Error("Bad button appID")
		}
		if keyboard.Buttons[0][0].Action.OwnerID != ownerID {
			t.Error("Bad button ownerID")
		}
		if keyboard.Buttons[0][0].Action.Payload != payload {
			t.Error("Bad button payload")
		}
		if keyboard.Buttons[0][0].Action.Label != label {
			t.Error("Bad button label")
		}
		if keyboard.Buttons[0][0].Action.Hash != hash {
			t.Error("Bad button hash")
		}
	})
}

func TestMessagesAudioMessage_ToAttachment(t *testing.T) {
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
			doc := MessagesAudioMessage{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := doc.ToAttachment(); got != tt.want {
				t.Errorf("MessagesAudioMessage.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessagesGraffiti_ToAttachment(t *testing.T) {
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
			doc := MessagesGraffiti{
				ID:      tt.fields.ID,
				OwnerID: tt.fields.OwnerID,
			}
			if got := doc.ToAttachment(); got != tt.want {
				t.Errorf("MessagesGraffiti.ToAttachment() = %v, want %v", got, tt.want)
			}
		})
	}
}
