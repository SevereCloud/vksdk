package wrapper // import "github.com/SevereCloud/vksdk/longpoll-user/v10"

import "github.com/SevereCloud/vksdk/longpoll-user"

// Wrapper struct
type Wrapper struct {
	longpoll *longpoll.Longpoll
}

// NewWrapper return *Wrapper for longpoll v10
func NewWrapper(lp *longpoll.Longpoll) *Wrapper {
	lp.Version = 10
	return &Wrapper{longpoll: lp}
}

// MessageFlagsSetHandler handler func for MessageFlagsSet
type MessageFlagsSetHandler func(m MessageFlagsSet)

// OnMessageFlagsSet func
//
// event with code 2
//
// TODO: write description
func (w Wrapper) OnMessageFlagsSet(f MessageFlagsSetHandler) {
	w.longpoll.EventNew(2, func(i []interface{}) error {
		event := MessageFlagsSet{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// MessageFlagsResetHandler handler func for MessageFlagsReset
type MessageFlagsResetHandler func(m MessageFlagsReset)

// OnMessageFlagsReset func
//
// event with code 3
//
// TODO: write description
func (w Wrapper) OnMessageFlagsReset(f MessageFlagsResetHandler) {
	w.longpoll.EventNew(3, func(i []interface{}) error {
		event := MessageFlagsReset{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// NewMessageHandler handler func for NewMessage
type NewMessageHandler func(m NewMessage)

// OnNewMessage func
//
// event with code 4
//
// TODO: write description
func (w Wrapper) OnNewMessage(f NewMessageHandler) {
	w.longpoll.EventNew(4, func(i []interface{}) error {
		event := NewMessage{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// EditMessageHandler handler func for EditMessage
type EditMessageHandler func(m EditMessage)

// OnEditMessage func
//
// event with code 5
//
// TODO: write description
func (w Wrapper) OnEditMessage(f EditMessageHandler) {
	w.longpoll.EventNew(5, func(i []interface{}) error {
		event := EditMessage{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// ReadInMessagesHandler handler func for ReadInMessages
type ReadInMessagesHandler func(m ReadInMessages)

// OnReadInMessages func
//
// event with code 6
//
// TODO: write description
func (w Wrapper) OnReadInMessages(f ReadInMessagesHandler) {
	w.longpoll.EventNew(6, func(i []interface{}) error {
		event := ReadInMessages{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// ReadOutMessagesHandler handler func for ReadOutMessages
type ReadOutMessagesHandler func(m ReadOutMessages)

// OnReadOutMessages func
//
// event with code 7
//
// TODO: write description
func (w Wrapper) OnReadOutMessages(f ReadOutMessagesHandler) {
	w.longpoll.EventNew(7, func(i []interface{}) error {
		event := ReadOutMessages{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// FriendBecameOnlineHandler handler func for FriendBecameOnline
type FriendBecameOnlineHandler func(m FriendBecameOnline)

// OnFriendBecameOnline func
//
// event with code 8
//
// TODO: write description
func (w Wrapper) OnFriendBecameOnline(f FriendBecameOnlineHandler) {
	w.longpoll.EventNew(8, func(i []interface{}) error {
		event := FriendBecameOnline{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// FriendBecameOfflineHandler handler func for FriendBecameOffline
type FriendBecameOfflineHandler func(m FriendBecameOffline)

// OnFriendBecameOffline func
//
// event with code 9
//
// TODO: write description
func (w Wrapper) OnFriendBecameOffline(f FriendBecameOfflineHandler) {
	w.longpoll.EventNew(9, func(i []interface{}) error {
		event := FriendBecameOffline{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// MentionViewedHandler handler func for MentionViewed
type MentionViewedHandler func(m MentionViewed)

// OnMentionViewed func
//
// event with code 10
//
// TODO: write description
func (w Wrapper) OnMentionViewed(f MentionViewedHandler) {
	w.longpoll.EventNew(10, func(i []interface{}) error {
		event := MentionViewed{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// UserMentionedHandler handler func for UserMentioned
type UserMentionedHandler func(m UserMentioned)

// OnUserMentioned func
//
// event with code 12
//
// TODO: write description
func (w Wrapper) OnUserMentioned(f UserMentionedHandler) {
	w.longpoll.EventNew(12, func(i []interface{}) error {
		event := UserMentioned{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// DeleteAllMessageHandler handler func for DeleteAllMessage
type DeleteAllMessageHandler func(m DeleteAllMessage)

// OnDeleteAllMessage func
//
// event with code 13
//
// TODO: write description
func (w Wrapper) OnDeleteAllMessage(f DeleteAllMessageHandler) {
	w.longpoll.EventNew(13, func(i []interface{}) error {
		event := DeleteAllMessage{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// LinkSnippetLoadedHandler handler func for LinkSnippetLoaded
type LinkSnippetLoadedHandler func(m LinkSnippetLoaded)

// OnLinkSnippetLoaded func
//
// event with code 18
//
// TODO: write description
func (w Wrapper) OnLinkSnippetLoaded(f LinkSnippetLoadedHandler) {
	w.longpoll.EventNew(18, func(i []interface{}) error {
		event := LinkSnippetLoaded{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// MessageCacheResetHandler handler func for MessageCacheReset
type MessageCacheResetHandler func(m MessageCacheReset)

// OnMessageCacheReset func
//
// event with code 19
//
// TODO: write description
func (w Wrapper) OnMessageCacheReset(f MessageCacheResetHandler) {
	w.longpoll.EventNew(19, func(i []interface{}) error {
		event := MessageCacheReset{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// ChatParamsChangeHandler handler func for ChatParamsChange
type ChatParamsChangeHandler func(m ChatParamsChange)

// OnChatParamsChange func
//
// event with code 51
//
// TODO: write description
func (w Wrapper) OnChatParamsChange(f ChatParamsChangeHandler) {
	w.longpoll.EventNew(51, func(i []interface{}) error {
		event := ChatParamsChange{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// ChatInfoChangeHandler handler func for ChatInfoChange
type ChatInfoChangeHandler func(m ChatInfoChange)

// OnChatInfoChange func
//
// event with code 52
//
// TODO: write description
func (w Wrapper) OnChatInfoChange(f ChatInfoChangeHandler) {
	w.longpoll.EventNew(52, func(i []interface{}) error {
		event := ChatInfoChange{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// UsersTypingHandler handler func for UsersTyping
type UsersTypingHandler func(m UsersTyping)

// OnUsersTyping func
//
// event with code 63
//
// TODO: write description
func (w Wrapper) OnUsersTyping(f UsersTypingHandler) {
	w.longpoll.EventNew(63, func(i []interface{}) error {
		event := UsersTyping{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// UsersRecordingAudioMessageHandler handler func for UsersRecordingAudioMessage
type UsersRecordingAudioMessageHandler func(m UsersRecordingAudioMessage)

// OnUsersRecordingAudioMessage func
//
// event with code 64
//
// TODO: write description
func (w Wrapper) OnUsersRecordingAudioMessage(f UsersRecordingAudioMessageHandler) {
	w.longpoll.EventNew(64, func(i []interface{}) error {
		event := UsersRecordingAudioMessage{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// CounterChangeHandler handler func for CounterChange
type CounterChangeHandler func(m CounterChange)

// OnCounterChange func
//
// event with code 80
//
// TODO: write description
func (w Wrapper) OnCounterChange(f CounterChangeHandler) {
	w.longpoll.EventNew(80, func(i []interface{}) error {
		event := CounterChange{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// FriendInvisibilityChangeHandler handler func for FriendInvisibilityChange
type FriendInvisibilityChangeHandler func(m FriendInvisibilityChange)

// OnFriendInvisibilityChange func
//
// event with code 81
//
// TODO: write description
func (w Wrapper) OnFriendInvisibilityChange(f FriendInvisibilityChangeHandler) {
	w.longpoll.EventNew(81, func(i []interface{}) error {
		event := FriendInvisibilityChange{}
		if err := event.Parse(i); err != nil {
			return err
		}
		f(event)

		return nil
	})
}

// NotificationSettingsChangeHandler handler func for NotificationSettingsChange
type NotificationSettingsChangeHandler func(m NotificationSettingsChange)

// OnNotificationSettingsChange func
//
// event with code 114
//
// TODO: write description
func (w Wrapper) OnNotificationSettingsChange(f NotificationSettingsChangeHandler) {
	w.longpoll.EventNew(114, func(i []interface{}) error {
		event := NotificationSettingsChange{}
		if w.longpoll.Mode&longpoll.ExtendedEvents != 0 {
			if err := event.ParseMode8(i); err != nil {
				return err
			}
		} else {
			if err := event.Parse(i); err != nil {
				return err
			}
		}
		f(event)

		return nil
	})
}
