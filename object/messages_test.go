package object_test

import (
	"testing"

	"github.com/Derad6709/vksdk/v2/object"

	"github.com/stretchr/testify/assert"
)

func TestMessagesKeyboard_AddRow(t *testing.T) {
	t.Parallel()

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()
	assert.Len(t, keyboard.Buttons, 1)

	keyboard.AddRow()
	assert.Len(t, keyboard.Buttons, 2)
}

func TestMessagesKeyboard_AddTextButton(t *testing.T) {
	t.Parallel()

	const (
		label   = "label"
		payload = "payload"
		color   = "color"
	)

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()

	keyboard.AddTextButton(label, payload, color)
	assert.Equal(t, keyboard.Buttons[0][0].Color, color)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, `"`+payload+`"`)
}

func TestMessagesKeyboard_AddTextButton_Panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddTextButton("label", make(chan int), "color")
}

func TestMessagesKeyboard_AddOpenLinkButton(t *testing.T) {
	t.Parallel()

	const (
		payload = "payload"
		label   = "label"
		link    = "https://vk.com"
	)

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()

	keyboard.AddOpenLinkButton(link, label, payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, `"`+payload+`"`)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Link, link)
}

func TestMessagesKeyboard_AddOpenLinkButton_Panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddOpenLinkButton("ink", "label", make(chan int))
}

func TestMessagesKeyboard_AddLocationButton(t *testing.T) {
	t.Parallel()

	const payload = "payload"

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()

	keyboard.AddLocationButton(payload)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, `"`+payload+`"`)
}

func TestMessagesKeyboard_AddLocationButton_Panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddLocationButton(make(chan int))
}

func TestMessagesKeyboard_AddVKPayButton(t *testing.T) {
	t.Parallel()

	const (
		payload = "payload"
		hash    = "hash"
	)

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()

	keyboard.AddVKPayButton(payload, hash)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, `"`+payload+`"`)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Hash, hash)
}

func TestMessagesKeyboard_AddVKPayButton_Panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddVKPayButton(make(chan int), "hash")
}

func TestMessagesKeyboard_AddVKAppsButton(t *testing.T) {
	t.Parallel()

	const (
		appID   = 1
		ownerID = 2
		payload = "payload"
		label   = "label"
		hash    = "hash"
	)

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()

	keyboard.AddVKAppsButton(appID, ownerID, payload, label, hash)
	assert.Equal(t, keyboard.Buttons[0][0].Action.AppID, appID)
	assert.Equal(t, keyboard.Buttons[0][0].Action.OwnerID, ownerID)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, `"`+payload+`"`)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Hash, hash)
}

func TestMessagesKeyboard_AddVKAppsButton_Panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddVKAppsButton(1, 1, make(chan int), "label", "hash")
}

func TestMessagesKeyboard_AddCallbackButton(t *testing.T) {
	t.Parallel()

	const (
		label   = "label"
		payload = "payload"
		color   = "color"
	)

	keyboard := object.NewMessagesKeyboard(false)

	keyboard.AddRow()

	keyboard.AddCallbackButton(label, payload, color)
	assert.Equal(t, keyboard.Buttons[0][0].Color, color)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Label, label)
	assert.Equal(t, keyboard.Buttons[0][0].Action.Payload, `"`+payload+`"`)
}

func TestMessagesKeyboard_AddCallbackButton_Panic(t *testing.T) {
	t.Parallel()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	keyboard := object.NewMessagesKeyboard(false)
	keyboard.AddRow()
	keyboard.AddCallbackButton("label", make(chan int), "color")
}

func TestMessagesAudioMessage_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(doc object.MessagesAudioMessage, want string) {
		if got := doc.ToAttachment(); got != want {
			t.Errorf("MessagesAudioMessage.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.MessagesAudioMessage{ID: 10, OwnerID: 20}, "doc20_10")
	f(object.MessagesAudioMessage{ID: 20, OwnerID: -10}, "doc-10_20")
}

func TestMessagesGraffiti_ToAttachment(t *testing.T) {
	t.Parallel()

	f := func(doc object.MessagesGraffiti, want string) {
		if got := doc.ToAttachment(); got != want {
			t.Errorf("MessagesGraffiti.ToAttachment() = %v, want %v", got, want)
		}
	}

	f(object.MessagesGraffiti{ID: 10, OwnerID: 20}, "doc20_10")
	f(object.MessagesGraffiti{ID: 20, OwnerID: -10}, "doc-10_20")
}

func TestMessagesKeyboard_ToJSON(t *testing.T) {
	t.Parallel()

	f := func(keyboard *object.MessagesKeyboard, want string) {
		t.Helper()

		got := keyboard.ToJSON()
		assert.Equal(t, got, want)
	}

	f(
		object.NewMessagesKeyboard(false),
		`{"buttons":[]}`,
	)
	f(
		object.NewMessagesKeyboard(true),
		`{"buttons":[],"one_time":true}`,
	)
	f(
		object.NewMessagesKeyboardInline(),
		`{"buttons":[],"inline":true}`,
	)
}

func TestMessagesTemplate_ToJSON(t *testing.T) {
	t.Parallel()

	f := func(template object.MessagesTemplate, want string) {
		t.Helper()

		got := template.ToJSON()
		assert.Equal(t, got, want)
	}

	f(
		object.MessagesTemplate{Type: "carousel", Elements: []object.MessagesTemplateElement{}},
		`{"type":"carousel","elements":[]}`,
	)
}

func TestMessageContentSource_ToJSON(t *testing.T) {
	t.Parallel()

	f := func(cs *object.MessageContentSource, want string) {
		t.Helper()

		got := cs.ToJSON()
		assert.Equal(t, got, want)
	}

	f(
		object.NewMessageContentSourceURL("https://vk.com"),
		`{"type":"url","url":"https://vk.com"}`,
	)
	f(
		object.NewMessageContentSourceMessage(1, 2, 3),
		`{"type":"message","owner_id":1,"peer_id":2,"conversation_message_id":3}`,
	)
}

func TestMessagesEventData_ToJSON(t *testing.T) {
	t.Parallel()

	f := func(eventData *object.MessagesEventData, want string) {
		t.Helper()

		got := eventData.ToJSON()
		assert.Equal(t, got, want)
	}

	f(
		object.NewMessagesEventDataShowSnackbar("test"),
		`{"type":"show_snackbar","text":"test"}`,
	)
	f(
		object.NewMessagesEventDataOpenLink("https://vk.com"),
		`{"type":"open_link","link":"https://vk.com"}`,
	)
	f(
		object.NewMessagesEventDataOpenApp(1, 2, "3"),
		`{"type":"open_app","app_id":1,"owner_id":2,"hash":"3"}`,
	)
	f(
		object.NewMessagesEventDataOpenApp(1, 0, "3"),
		`{"type":"open_app","app_id":1,"hash":"3"}`,
	)
}

func TestMessagesForward_ToJSON(t *testing.T) {
	t.Parallel()

	f := func(forward object.MessagesForward, want string) {
		t.Helper()

		got := forward.ToJSON()
		assert.Equal(t, got, want)
	}

	f(
		object.MessagesForward{
			OwnerID:                1,
			PeerID:                 2,
			ConversationMessageIDs: []int{3},
			IsReply:                true,
			MessageIDs:             []int{4},
		},
		`{"owner_id":1,"peer_id":2,"conversation_message_ids":[3],"message_ids":[4],"is_reply":true}`,
	)
	f(
		object.MessagesForward{},
		`{}`,
	)
}
