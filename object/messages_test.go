package object_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/object"

	"github.com/stretchr/testify/assert"
)

func TestMessagesKeyboard_AddRow(t *testing.T) {
	keyboard := object.NewMessagesKeyboard(false, false)

	keyboard.AddRow()
	assert.Len(t, keyboard.Buttons, 1)

	keyboard.AddRow()
	assert.Len(t, keyboard.Buttons, 2)
}

func TestMessagesKeyboard_AddTextButton(t *testing.T) {
	const (
		label   = "label"
		payload = "payload"
		color   = "color"
	)

	keyboard := object.NewMessagesKeyboard(false, false)

	keyboard.AddRow()

	keyboard.AddTextButton(label, payload, color)
	assert.Equal(t, keyboard.Buttons[0][0].Color, color)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, payload)
}

func TestMessagesKeyboard_AddOpenLinkButton(t *testing.T) {
	const (
		payload = "payload"
		label   = "label"
		link    = "https://vk.com"
	)

	keyboard := object.NewMessagesKeyboard(false, false)

	keyboard.AddRow()

	keyboard.AddOpenLinkButton(link, label, payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Link, link)
}

func TestMessagesKeyboard_AddLocationButton(t *testing.T) {
	const payload = "payload"

	keyboard := object.NewMessagesKeyboard(false, false)

	keyboard.AddRow()

	keyboard.AddLocationButton(payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, payload)
}

func TestMessagesKeyboard_AddVKPayButton(t *testing.T) {
	const (
		payload = "payload"
		hash    = "hash"
	)

	keyboard := object.NewMessagesKeyboard(false, false)

	keyboard.AddRow()

	keyboard.AddVKPayButton(payload, hash)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Hash, hash)
}

func TestMessagesKeyboard_AddVKAppsButton(t *testing.T) {
	const (
		appID   = 1
		ownerID = 2
		payload = "payload"
		label   = "label"
		hash    = "hash"
	)

	keyboard := object.NewMessagesKeyboard(false, false)

	keyboard.AddRow()

	keyboard.AddVKAppsButton(appID, ownerID, payload, label, hash)
	assert.Equal(t, keyboard.Buttons[0][0].Action.AppID, appID)
	assert.Equal(t, keyboard.Buttons[0][0].Action.OwnerID, ownerID)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Hash, hash)
}

func TestMessagesAudioMessage_ToAttachment(t *testing.T) {
	f := func(doc object.MessagesAudioMessage, want string) {
		if got := doc.ToAttachment(); got != want {
			t.Errorf("MessagesAudioMessage.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.MessagesAudioMessage{ID: 10, OwnerID: 20}, "doc20_10")
	f(object.MessagesAudioMessage{ID: 20, OwnerID: -10}, "doc-10_20")
}

func TestMessagesGraffiti_ToAttachment(t *testing.T) {
	f := func(doc object.MessagesGraffiti, want string) {
		if got := doc.ToAttachment(); got != want {
			t.Errorf("MessagesGraffiti.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.MessagesGraffiti{ID: 10, OwnerID: 20}, "doc20_10")
	f(object.MessagesGraffiti{ID: 20, OwnerID: -10}, "doc-10_20")
}

func TestMessagesKeyboard_ToJSON(t *testing.T) {
	f := func(keyboard object.MessagesKeyboard, want string) {
		t.Helper()

		got := keyboard.ToJSON()
		assert.Equal(t, got, want)
	}

	f(
		object.NewMessagesKeyboard(false, false),
		`{"buttons":[],"one_time":false}`,
	)
}
