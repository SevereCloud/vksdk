package v3

import "github.com/SevereCloud/vksdk/longpoll-user"

type Wrapper struct {
	longpoll *longpoll.Longpoll
}

func NewWrapper(lp *longpoll.Longpoll) *Wrapper {
	lp.Version = 3
	return &Wrapper{longpoll: lp}
}

type MessageFlagsChangeHandler func(m MessageFlagsChange)

func (w Wrapper) OnMessageFlagsChange(f MessageFlagsChangeHandler) {
	w.longpoll.EventNew(1, func(i []interface{}) {
		event := MessageFlagsChange{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type MessageFlagsSetHandler func(m MessageFlagsSet)

func (w Wrapper) OnMessageFlagsSet(f MessageFlagsSetHandler) {
	w.longpoll.EventNew(2, func(i []interface{}) {
		event := MessageFlagsSet{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type MessageFlagsResetHandler func(m MessageFlagsReset)

func (w Wrapper) OnMessageFlagsReset(f MessageFlagsResetHandler) {
	w.longpoll.EventNew(3, func(i []interface{}) {
		event := MessageFlagsReset{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type NewMessageHandler func(m NewMessage)

func (w Wrapper) OnNewMessage(f NewMessageHandler) {
	w.longpoll.EventNew(4, func(i []interface{}) {
		event := NewMessage{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type EditMessageHandler func(m EditMessage)

func (w Wrapper) OnEditMessage(f EditMessageHandler) {
	w.longpoll.EventNew(5, func(i []interface{}) {
		event := EditMessage{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type ReadInMessagesHandler func(m ReadInMessages)

func (w Wrapper) OnReadInMessages(f ReadInMessagesHandler) {
	w.longpoll.EventNew(6, func(i []interface{}) {
		event := ReadInMessages{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type ReadOutMessagesHandler func(m ReadOutMessages)

func (w Wrapper) OnReadOutMessages(f ReadOutMessagesHandler) {
	w.longpoll.EventNew(7, func(i []interface{}) {
		event := ReadOutMessages{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type FriendBecameOnlineHandler func(m FriendBecameOnline)

func (w Wrapper) OnFriendBecameOnline(f FriendBecameOnlineHandler) {
	w.longpoll.EventNew(8, func(i []interface{}) {
		event := FriendBecameOnline{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type FriendBecameOfflineHandler func(m FriendBecameOffline)

func (w Wrapper) OnFriendBecameOffline(f FriendBecameOfflineHandler) {
	w.longpoll.EventNew(9, func(i []interface{}) {
		event := FriendBecameOffline{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type DialogFlagsResetHandler func(m DialogFlagsReset)

func (w Wrapper) OnDialogFlagsReset(f DialogFlagsResetHandler) {
	w.longpoll.EventNew(10, func(i []interface{}) {
		event := DialogFlagsReset{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type DialogFlagsReplaceHandler func(m DialogFlagsReplace)

func (w Wrapper) OnDialogFlagsReplace(f DialogFlagsReplaceHandler) {
	w.longpoll.EventNew(11, func(i []interface{}) {
		event := DialogFlagsReplace{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type DialogsFlagsSetHandler func(m DialogsFlagsSet)

func (w Wrapper) OnDialogsFlagsSet(f DialogsFlagsSetHandler) {
	w.longpoll.EventNew(12, func(i []interface{}) {
		event := DialogsFlagsSet{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type DeleteMessagesHandler func(m DeleteMessages)

func (w Wrapper) OnDeleteMessages(f DeleteMessagesHandler) {
	w.longpoll.EventNew(13, func(i []interface{}) {
		event := DeleteMessages{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type RestoreDeletedMessagesHandler func(m RestoreDeletedMessages)

func (w Wrapper) OnRestoreDeletedMessages(f RestoreDeletedMessagesHandler) {
	w.longpoll.EventNew(14, func(i []interface{}) {
		event := RestoreDeletedMessages{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type ChatParamsChangeHandler func(m ChatParamsChange)

func (w Wrapper) OnChatParamsChange(f ChatParamsChangeHandler) {
	w.longpoll.EventNew(51, func(i []interface{}) {
		event := ChatParamsChange{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type ChatInfoChangeHandler func(m ChatInfoChange)

func (w Wrapper) OnChatInfoChange(f ChatInfoChangeHandler) {
	w.longpoll.EventNew(52, func(i []interface{}) {
		event := ChatInfoChange{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type UserTypingHandler func(m UserTyping)

func (w Wrapper) OnUserTyping(f UserTypingHandler) {
	w.longpoll.EventNew(61, func(i []interface{}) {
		event := UserTyping{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type UserTypingChatHandler func(m UserTypingChat)

func (w Wrapper) OnUserTypingChat(f UserTypingChatHandler) {
	w.longpoll.EventNew(62, func(i []interface{}) {
		event := UserTypingChat{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type UsersTypingHandler func(m UsersTyping)

func (w Wrapper) OnUsersTyping(f UsersTypingHandler) {
	w.longpoll.EventNew(63, func(i []interface{}) {
		event := UsersTyping{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type UsersRecordingAudioMessageHandler func(m UsersRecordingAudioMessage)

func (w Wrapper) OnUsersRecordingAudioMessage(f UsersRecordingAudioMessageHandler) {
	w.longpoll.EventNew(64, func(i []interface{}) {
		event := UsersRecordingAudioMessage{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type UserCallHandler func(m UserCall)

func (w Wrapper) OnUserCall(f UserCallHandler) {
	w.longpoll.EventNew(70, func(i []interface{}) {
		event := UserCall{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type CounterChangeHandler func(m CounterChange)

func (w Wrapper) OnCounterChange(f CounterChangeHandler) {
	w.longpoll.EventNew(80, func(i []interface{}) {
		event := CounterChange{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}

type NotificationSettingsChangeHandler func(m NotificationSettingsChange)

func (w Wrapper) OnNotificationSettingsChange(f NotificationSettingsChangeHandler) {
	w.longpoll.EventNew(114, func(i []interface{}) {
		event := NotificationSettingsChange{}
		if err := event.Parse(i); err == nil {
			f(event)
		}
	})
}
