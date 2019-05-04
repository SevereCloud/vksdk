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
	var keyboard MessagesKeyboard
	keyboard.AddRow()
	t.Run("Add button", func(t *testing.T) {
		keyboard.AddButton("label", "payload", "color")
		if keyboard.Buttons[0][0].Color != "color" {
			t.Error("Bad button color")
		}
		if keyboard.Buttons[0][0].Action.Label != "label" {
			t.Error("Bad button label")
		}
		if keyboard.Buttons[0][0].Action.Payload != "payload" {
			t.Error("Bad button payload")
		}
	})
}
