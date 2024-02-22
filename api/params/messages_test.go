package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMessagesAddChatUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesAddChatUserBuilder()

	b.ChatID(1)
	b.UserID(1)

	assert.Equal(t, 1, b.Params["chat_id"])
	assert.Equal(t, 1, b.Params["user_id"])
}

func TestMessagesAllowMessagesFromGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesAllowMessagesFromGroupBuilder()

	b.GroupID(1)
	b.Key("text")

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["key"])
}

func TestMessagesCreateChatBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesCreateChatBuilder()

	b.UserIDs([]int{1})
	b.Title("text")

	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, "text", b.Params["title"])
}

func TestMessagesDeleteBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDeleteBuilder()

	b.MessageIDs([]int{1})
	b.PeerID(1)
	b.ConversationMessageIDs([]int{1})
	b.Spam(true)
	b.GroupID(1)
	b.DeleteForAll(true)

	assert.Equal(t, []int{1}, b.Params["message_ids"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, []int{1}, b.Params["conversation_message_ids"])
	assert.Equal(t, true, b.Params["spam"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, true, b.Params["delete_for_all"])
}

func TestMessagesDeleteChatPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDeleteChatPhotoBuilder()

	b.ChatID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["chat_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesDeleteConversationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDeleteConversationBuilder()

	b.UserID(1)
	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesDenyMessagesFromGroupBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesDenyMessagesFromGroupBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesEditBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesEditBuilder()

	b.PeerID(1)
	b.Message("text")
	b.MessageID(1)
	b.Lat(1.1)
	b.Long(1.1)
	b.Attachment("text")
	b.KeepForwardMessages(true)
	b.KeepSnippets(true)
	b.GroupID(1)
	b.Keyboard("text")
	b.Template("text")
	b.ConversationMessageID(1)
	b.DontParseLinks(true)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, "text", b.Params["message"])
	assert.Equal(t, 1, b.Params["message_id"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, "text", b.Params["attachment"])
	assert.Equal(t, true, b.Params["keep_forward_messages"])
	assert.Equal(t, true, b.Params["keep_snippets"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["keyboard"])
	assert.Equal(t, "text", b.Params["template"])
	assert.Equal(t, 1, b.Params["conversation_message_id"])
	assert.Equal(t, true, b.Params["dont_parse_links"])
}

func TestMessagesEditChatBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesEditChatBuilder()

	b.ChatID(1)
	b.Title("text")

	assert.Equal(t, 1, b.Params["chat_id"])
	assert.Equal(t, "text", b.Params["title"])
}

func TestMessagesForceCallFinishBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesForceCallFinishBuilder()

	b.CallID("text")

	assert.Equal(t, "text", b.Params["call_id"])
}

func TestMessagesGetByConversationMessageIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetByConversationMessageIDBuilder()

	b.PeerID(1)
	b.ConversationMessageIDs([]int{1})
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, []int{1}, b.Params["conversation_message_ids"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetByIDBuilder()

	b.MessageIDs([]int{1})
	b.PreviewLength(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, []int{1}, b.Params["message_ids"])
	assert.Equal(t, 1, b.Params["preview_length"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetChatPreviewBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetChatPreviewBuilder()

	b.PeerID(1)
	b.Link("text")
	b.Fields([]string{"test"})

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, "text", b.Params["link"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestMessagesGetConversationMembersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetConversationMembersBuilder()

	b.PeerID(1)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetConversationsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetConversationsBuilder()

	b.Offset(1)
	b.Count(1)
	b.Filter("text")
	b.Extended(true)
	b.StartMessageID(1)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, "text", b.Params["filter"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["start_message_id"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetConversationsByIDBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetConversationsByIDBuilder()

	b.PeerIDs([]int{1})
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, []int{1}, b.Params["peer_ids"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetHistoryBuilder()

	b.Offset(1)
	b.Count(1)
	b.UserID(1)
	b.PeerID(1)
	b.StartMessageID(1)
	b.Rev(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["start_message_id"])
	assert.Equal(t, 1, b.Params["rev"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetHistoryAttachmentsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetHistoryAttachmentsBuilder()

	b.PeerID(1)
	b.MediaType("text")
	b.StartFrom("text")
	b.Count(1)
	b.PhotoSizes(true)
	b.Fields([]string{"test"})
	b.GroupID(1)
	b.PreserveOrder(true)
	b.MaxForwardsLevel(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, "text", b.Params["media_type"])
	assert.Equal(t, "text", b.Params["start_from"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["photo_sizes"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, true, b.Params["preserve_order"])
	assert.Equal(t, 1, b.Params["max_forwards_level"])
}

func TestMessagesGetIntentUsersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetIntentUsersBuilder()

	b.Intent("test")
	b.SubscribeID(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.NameCase([]string{"test"})
	b.Fields([]string{"test"})

	assert.Equal(t, "test", b.Params["intent"])
	assert.Equal(t, 1, b.Params["subscribe_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["name_case"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestMessagesGetInviteLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetInviteLinkBuilder()

	b.PeerID(1)
	b.Reset(true)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, true, b.Params["reset"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesGetLastActivityBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetLastActivityBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}

func TestMessagesGetLongPollHistoryBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetLongPollHistoryBuilder()

	b.Ts(1)
	b.Pts(1)
	b.PreviewLength(1)
	b.Onlines(true)
	b.Fields([]string{"test"})
	b.EventsLimit(1)
	b.MsgsLimit(1)
	b.MaxMsgID(1)
	b.GroupID(1)
	b.LpVersion(1)
	b.LastN(1)
	b.Credentials(true)

	assert.Equal(t, 1, b.Params["ts"])
	assert.Equal(t, 1, b.Params["pts"])
	assert.Equal(t, 1, b.Params["preview_length"])
	assert.Equal(t, true, b.Params["onlines"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["events_limit"])
	assert.Equal(t, 1, b.Params["msgs_limit"])
	assert.Equal(t, 1, b.Params["max_msg_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["lp_version"])
	assert.Equal(t, 1, b.Params["last_n"])
	assert.Equal(t, true, b.Params["credentials"])
}

func TestMessagesGetLongPollServerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesGetLongPollServerBuilder()

	b.NeedPts(true)
	b.GroupID(1)
	b.LpVersion(1)

	assert.Equal(t, true, b.Params["need_pts"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["lp_version"])
}

func TestMessagesIsMessagesFromGroupAllowedBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesIsMessagesFromGroupAllowedBuilder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, 1, b.Params["user_id"])
}

func TestMessagesJoinChatByInviteLinkBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesJoinChatByInviteLinkBuilder()

	b.Link("text")

	assert.Equal(t, "text", b.Params["link"])
}

func TestMessagesMarkAsAnsweredConversationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsAnsweredConversationBuilder()

	b.PeerID(1)
	b.Answered(true)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, true, b.Params["answered"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesMarkAsImportantBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsImportantBuilder()

	b.MessageIDs([]int{1})
	b.Important(1)

	assert.Equal(t, []int{1}, b.Params["message_ids"])
	assert.Equal(t, 1, b.Params["important"])
}

func TestMessagesMarkAsImportantConversationBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsImportantConversationBuilder()

	b.PeerID(1)
	b.Important(true)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, true, b.Params["important"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesMarkAsReadBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesMarkAsReadBuilder()

	b.MessageIDs([]int{1})
	b.PeerID(1)
	b.StartMessageID(1)
	b.GroupID(1)

	assert.Equal(t, []int{1}, b.Params["message_ids"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["start_message_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesPinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesPinBuilder()

	b.PeerID(1)
	b.MessageID(1)
	b.ConversationMessageID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["message_id"])
	assert.Equal(t, 1, b.Params["conversation_message_id"])
}

func TestMessagesRemoveChatUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesRemoveChatUserBuilder()

	b.ChatID(1)
	b.UserID(1)
	b.MemberID(1)

	assert.Equal(t, 1, b.Params["chat_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["member_id"])
}

func TestMessagesRestoreBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesRestoreBuilder()

	b.MessageID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["message_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSearchBuilder()

	b.Q("text")
	b.PeerID(1)
	b.Date(1)
	b.PreviewLength(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Fields([]string{"text"})
	b.GroupID(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["date"])
	assert.Equal(t, 1, b.Params["preview_length"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"text"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesSearchConversationsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSearchConversationsBuilder()

	b.Q("text")
	b.Count(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesSendBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSendBuilder()

	b.UserID(1)
	b.RandomID(1)
	b.PeerID(1)
	b.Domain("text")
	b.ChatID(1)
	b.UserIDs([]int{1})
	b.PeerIDs([]int{1})
	b.Message("text")
	b.Lat(1.1)
	b.Long(1.1)
	b.Attachment("text")
	b.ReplyTo(1)
	b.ForwardMessages([]int{1})
	b.Forward("text")
	b.StickerID(1)
	b.GroupID(1)
	b.Keyboard("text")
	b.Template("text")
	b.Payload("text")
	b.ContentSource("text")
	b.DontParseLinks(true)
	b.DisableMentions(true)
	b.Intent("text")
	b.SubscribeID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["random_id"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, "text", b.Params["domain"])
	assert.Equal(t, 1, b.Params["chat_id"])
	assert.Equal(t, []int{1}, b.Params["user_ids"])
	assert.Equal(t, []int{1}, b.Params["peer_ids"])
	assert.Equal(t, "text", b.Params["message"])
	assert.InEpsilon(t, 1.1, b.Params["lat"], 0.1)
	assert.InEpsilon(t, 1.1, b.Params["long"], 0.1)
	assert.Equal(t, "text", b.Params["attachment"])
	assert.Equal(t, 1, b.Params["reply_to"])
	assert.Equal(t, []int{1}, b.Params["forward_messages"])
	assert.Equal(t, "text", b.Params["forward"])
	assert.Equal(t, 1, b.Params["sticker_id"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, "text", b.Params["keyboard"])
	assert.Equal(t, "text", b.Params["template"])
	assert.Equal(t, "text", b.Params["payload"])
	assert.Equal(t, "text", b.Params["content_source"])
	assert.Equal(t, true, b.Params["dont_parse_links"])
	assert.Equal(t, true, b.Params["disable_mentions"])
	assert.Equal(t, "text", b.Params["intent"])
	assert.Equal(t, 1, b.Params["subscribe_id"])
}

func TestMessagesSendMessageEventAnswerBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSendMessageEventAnswerBuilder()

	b.EventID("text")
	b.UserID(1)
	b.PeerID(1)
	b.EventData("text")

	assert.Equal(t, "text", b.Params["event_id"])
	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, "text", b.Params["event_data"])
}

func TestMessagesSetActivityBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSetActivityBuilder()

	b.UserID(1)
	b.Type("text")
	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesSetChatPhotoBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesSetChatPhotoBuilder()

	b.File("text")

	assert.Equal(t, "text", b.Params["file"])
}

func TestMessagesStartCallBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesStartCallBuilder()

	b.GroupID(1)

	assert.Equal(t, 1, b.Params["group_id"])
}

func TestMessagesUnpinBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewMessagesUnpinBuilder()

	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, 1, b.Params["peer_id"])
	assert.Equal(t, 1, b.Params["group_id"])
}
