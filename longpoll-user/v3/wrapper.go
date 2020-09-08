package wrapper // import "github.com/SevereCloud/vksdk/v2/longpoll-user/v3"

import "github.com/SevereCloud/vksdk/v2/longpoll-user"

// Wrapper struct.
type Wrapper struct {
	longpoll *longpoll.LongPoll
}

// NewWrapper return *Wrapper for longpoll v3.
func NewWrapper(lp *longpoll.LongPoll) *Wrapper {
	lp.Version = 3
	return &Wrapper{longpoll: lp}
}

// MessageFlagsChangeHandler handler func for MessageFlagsChange.
type MessageFlagsChangeHandler func(m MessageFlagsChange)

// OnMessageFlagsChange handler for MessageFlagsChange.
//
// event with code 1.
//
// Replace message flags.
func (w Wrapper) OnMessageFlagsChange(f MessageFlagsChangeHandler) {
	w.longpoll.EventNew(1, func(i []interface{}) error {
		var event MessageFlagsChange
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// MessageFlagsSetHandler handler func for MessageFlagsSet.
type MessageFlagsSetHandler func(m MessageFlagsSet)

// OnMessageFlagsSet handler for MessageFlagsSet.
//
// event with code 2.
//
// Install message flags.
func (w Wrapper) OnMessageFlagsSet(f MessageFlagsSetHandler) {
	w.longpoll.EventNew(2, func(i []interface{}) error {
		var event MessageFlagsSet
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// MessageFlagsResetHandler handler func for MessageFlagsReset.
type MessageFlagsResetHandler func(m MessageFlagsReset)

// OnMessageFlagsReset handler for MessageFlagsReset.
//
// event with code 3.
//
// Reset message flags.
func (w Wrapper) OnMessageFlagsReset(f MessageFlagsResetHandler) {
	w.longpoll.EventNew(3, func(i []interface{}) error {
		var event MessageFlagsReset
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// NewMessageHandler handler func for NewMessage.
type NewMessageHandler func(m NewMessage)

// OnNewMessage handler for NewMessage.
//
// event with code 4.
//
// Add a new message.
func (w Wrapper) OnNewMessage(f NewMessageHandler) {
	w.longpoll.EventNew(4, func(i []interface{}) error {
		var event NewMessage
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// EditMessageHandler handler func for EditMessage.
type EditMessageHandler func(m EditMessage)

// OnEditMessage handler for EditMessage.
//
// event with code 5.
//
// Edit message.
func (w Wrapper) OnEditMessage(f EditMessageHandler) {
	w.longpoll.EventNew(5, func(i []interface{}) error {
		var event EditMessage
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// ReadInMessagesHandler handler func for ReadInMessages.
type ReadInMessagesHandler func(m ReadInMessages)

// OnReadInMessages handler for ReadInMessages.
//
// event with code 6.
//
// Read all incoming messages received before message with LocalID.
func (w Wrapper) OnReadInMessages(f ReadInMessagesHandler) {
	w.longpoll.EventNew(6, func(i []interface{}) error {
		var event ReadInMessages
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// ReadOutMessagesHandler handler func for ReadOutMessages.
type ReadOutMessagesHandler func(m ReadOutMessages)

// OnReadOutMessages handler for ReadOutMessages.
//
// event with code 7.
//
// Read all outgoing messages sent before message with LocalID.
func (w Wrapper) OnReadOutMessages(f ReadOutMessagesHandler) {
	w.longpoll.EventNew(7, func(i []interface{}) error {
		var event ReadOutMessages
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// FriendBecameOnlineHandler handler func for FriendBecameOnline.
type FriendBecameOnlineHandler func(m FriendBecameOnline)

// OnFriendBecameOnline handler for FriendBecameOnline.
//
// event with code 8.
//
// A friend UserID is online.
func (w Wrapper) OnFriendBecameOnline(f FriendBecameOnlineHandler) {
	w.longpoll.EventNew(8, func(i []interface{}) error {
		var event FriendBecameOnline
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// FriendBecameOfflineHandler handler func for FriendBecameOffline.
type FriendBecameOfflineHandler func(m FriendBecameOffline)

// OnFriendBecameOffline handler for FriendBecameOffline.
//
// event with code 9.
//
// A friend UserID is offline.
func (w Wrapper) OnFriendBecameOffline(f FriendBecameOfflineHandler) {
	w.longpoll.EventNew(9, func(i []interface{}) error {
		var event FriendBecameOffline
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// DialogFlagsResetHandler handler func for DialogFlagsReset.
type DialogFlagsResetHandler func(m DialogFlagsReset)

// OnDialogFlagsReset handler for DialogFlagsReset.
//
// event with code 10.
//
// Reset dialog flags.
func (w Wrapper) OnDialogFlagsReset(f DialogFlagsResetHandler) {
	w.longpoll.EventNew(10, func(i []interface{}) error {
		var event DialogFlagsReset
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// DialogFlagsReplaceHandler handler func for DialogFlagsReplace.
type DialogFlagsReplaceHandler func(m DialogFlagsReplace)

// OnDialogFlagsReplace handler for DialogFlagsReplace.
//
// event with code 11.
//
// Replace dialog flags.
func (w Wrapper) OnDialogFlagsReplace(f DialogFlagsReplaceHandler) {
	w.longpoll.EventNew(11, func(i []interface{}) error {
		var event DialogFlagsReplace
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// DialogsFlagsSetHandler handler func for DialogsFlagsSet.
type DialogsFlagsSetHandler func(m DialogsFlagsSet)

// OnDialogsFlagsSet handler for DialogsFlagsSet.
//
// event with code 12.
//
// Install dialog flags.
func (w Wrapper) OnDialogsFlagsSet(f DialogsFlagsSetHandler) {
	w.longpoll.EventNew(12, func(i []interface{}) error {
		var event DialogsFlagsSet
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// DeleteMessagesHandler handler func for DeleteMessages.
type DeleteMessagesHandler func(m DeleteMessages)

// OnDeleteMessages handler for DeleteMessages.
//
// event with code 13.
//
// Deleting all messages in PeerID dialogs with IDs up to LocalID.
func (w Wrapper) OnDeleteMessages(f DeleteMessagesHandler) {
	w.longpoll.EventNew(13, func(i []interface{}) error {
		var event DeleteMessages
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// RestoreDeletedMessagesHandler handler func for RestoreDeletedMessages.
type RestoreDeletedMessagesHandler func(m RestoreDeletedMessages)

// OnRestoreDeletedMessages handler for RestoreDeletedMessages.
//
// event with code 14.
//
// Restore message.
func (w Wrapper) OnRestoreDeletedMessages(f RestoreDeletedMessagesHandler) {
	w.longpoll.EventNew(14, func(i []interface{}) error {
		var event RestoreDeletedMessages
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// ChatParamsChangeHandler handler func for ChatParamsChange.
type ChatParamsChangeHandler func(m ChatParamsChange)

// OnChatParamsChange handler for ChatParamsChange.
//
// event with code 51.
//
// One of the parameters (content, topic) of the conversation ChatID was
// changed. Self — 1 or 0 (whether the change was caused by the user).
func (w Wrapper) OnChatParamsChange(f ChatParamsChangeHandler) {
	w.longpoll.EventNew(51, func(i []interface{}) error {
		var event ChatParamsChange
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// ChatInfoChangeHandler handler func for ChatInfoChange.
type ChatInfoChangeHandler func(m ChatInfoChange)

// OnChatInfoChange handler for ChatInfoChange.
//
// event with code 52.
//
// Chat info change.
func (w Wrapper) OnChatInfoChange(f ChatInfoChangeHandler) {
	w.longpoll.EventNew(52, func(i []interface{}) error {
		var event ChatInfoChange
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// UserTypingHandler handler func for UserTyping.
type UserTypingHandler func(m UserTyping)

// OnUserTyping handler for UserTyping.
//
// event with code 61.
//
// User UserID began typing in the dialog. The event should happen once in
// ~5 seconds during continuous typing.
func (w Wrapper) OnUserTyping(f UserTypingHandler) {
	w.longpoll.EventNew(61, func(i []interface{}) error {
		var event UserTyping
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// UserTypingChatHandler handler func for UserTypingChat.
type UserTypingChatHandler func(m UserTypingChat)

// OnUserTypingChat handler for UserTypingChat.
//
// event with code 62.
//
// User UserID began typing in the conversation ChatID.
func (w Wrapper) OnUserTypingChat(f UserTypingChatHandler) {
	w.longpoll.EventNew(62, func(i []interface{}) error {
		var event UserTypingChat
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// UsersTypingHandler handler func for UsersTyping.
type UsersTypingHandler func(m UsersTyping)

// OnUsersTyping handler for UsersTyping.
//
// event with code 63.
func (w Wrapper) OnUsersTyping(f UsersTypingHandler) {
	w.longpoll.EventNew(63, func(i []interface{}) error {
		var event UsersTyping
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// UsersRecordingAudioMessageHandler handler func for UsersRecordingAudioMessage.
type UsersRecordingAudioMessageHandler func(m UsersRecordingAudioMessage)

// OnUsersRecordingAudioMessage handler for UsersRecordingAudioMessage.
//
// event with code 64.
func (w Wrapper) OnUsersRecordingAudioMessage(f UsersRecordingAudioMessageHandler) {
	w.longpoll.EventNew(64, func(i []interface{}) error {
		var event UsersRecordingAudioMessage
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// UserCallHandler handler func for UserCall.
type UserCallHandler func(m UserCall)

// OnUserCall handler for UserCall.
//
// event with code 70.
func (w Wrapper) OnUserCall(f UserCallHandler) {
	w.longpoll.EventNew(70, func(i []interface{}) error {
		var event UserCall
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// CounterChangeHandler handler func for CounterChange.
type CounterChangeHandler func(m CounterChange)

// OnCounterChange handler for CounterChange.
//
// event with code 80.
func (w Wrapper) OnCounterChange(f CounterChangeHandler) {
	w.longpoll.EventNew(80, func(i []interface{}) error {
		var event CounterChange
		if err := event.parse(i); err != nil {
			return err
		}

		f(event)

		return nil
	})
}

// NotificationSettingsChangeHandler handler func for NotificationSettingsChange.
type NotificationSettingsChangeHandler func(m NotificationSettingsChange)

// OnNotificationSettingsChange handler for NotificationSettingsChange.
//
// event with code 114.
func (w Wrapper) OnNotificationSettingsChange(f NotificationSettingsChangeHandler) {
	w.longpoll.EventNew(114, func(i []interface{}) error {
		var event NotificationSettingsChange
		if w.longpoll.Mode&longpoll.ExtendedEvents != 0 {
			if err := event.parseMode8(i); err != nil {
				return err
			}
		} else {
			if err := event.parse(i); err != nil {
				return err
			}
		}

		f(event)

		return nil
	})
}
