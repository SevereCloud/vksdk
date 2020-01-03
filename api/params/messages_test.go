package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestMessagesAddChatUserBulder(t *testing.T) {
	b := params.NewMessagesAddChatUserBulder()

	b.ChatID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestMessagesAllowMessagesFromGroupBulder(t *testing.T) {
	b := params.NewMessagesAllowMessagesFromGroupBulder()

	b.GroupID(1)
	b.Key("text")

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["key"], "text")
}

func TestMessagesCreateChatBulder(t *testing.T) {
	b := params.NewMessagesCreateChatBulder()

	b.UserIDs([]int{1})
	b.Title("text")

	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["title"], "text")
}

func TestMessagesDeleteBulder(t *testing.T) {
	b := params.NewMessagesDeleteBulder()

	b.MessageIDs([]int{1})
	b.Spam(true)
	b.GroupID(1)
	b.DeleteForAll(true)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["spam"], true)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["delete_for_all"], true)
}

func TestMessagesDeleteChatPhotoBulder(t *testing.T) {
	b := params.NewMessagesDeleteChatPhotoBulder()

	b.ChatID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesDeleteConversationBulder(t *testing.T) {
	b := params.NewMessagesDeleteConversationBulder()

	b.UserID(1)
	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesDenyMessagesFromGroupBulder(t *testing.T) {
	b := params.NewMessagesDenyMessagesFromGroupBulder()

	b.GroupID(1)

	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesEditBulder(t *testing.T) {
	b := params.NewMessagesEditBulder()

	b.PeerID(1)
	b.Message("text")
	b.MessageID(1)
	b.Lat(1.1)
	b.Long(1.1)
	b.Attachment("text")
	b.KeepForwardMessages(true)
	b.KeepSnippets(true)
	b.GroupID(1)
	b.DontParseLinks(true)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["message_id"], 1)
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["attachment"], "text")
	assert.Equal(t, b.Params["keep_forward_messages"], true)
	assert.Equal(t, b.Params["keep_snippets"], true)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["dont_parse_links"], true)
}

func TestMessagesEditChatBulder(t *testing.T) {
	b := params.NewMessagesEditChatBulder()

	b.ChatID(1)
	b.Title("text")

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["title"], "text")
}

func TestMessagesGetByConversationMessageIDBulder(t *testing.T) {
	b := params.NewMessagesGetByConversationMessageIDBulder()

	b.PeerID(1)
	b.ConversationMessageIDs([]int{1})
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["conversation_message_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetByIDBulder(t *testing.T) {
	b := params.NewMessagesGetByIDBulder()

	b.MessageIDs([]int{1})
	b.PreviewLength(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetChatPreviewBulder(t *testing.T) {
	b := params.NewMessagesGetChatPreviewBulder()

	b.PeerID(1)
	b.Link("text")
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["link"], "text")
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestMessagesGetConversationMembersBulder(t *testing.T) {
	b := params.NewMessagesGetConversationMembersBulder()

	b.PeerID(1)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetConversationsBulder(t *testing.T) {
	b := params.NewMessagesGetConversationsBulder()

	b.Offset(1)
	b.Count(1)
	b.Filter("text")
	b.Extended(true)
	b.StartMessageID(1)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["filter"], "text")
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["start_message_id"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetConversationsByIDBulder(t *testing.T) {
	b := params.NewMessagesGetConversationsByIDBulder()

	b.PeerIDs([]int{1})
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_ids"], []int{1})
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetHistoryBulder(t *testing.T) {
	b := params.NewMessagesGetHistoryBulder()

	b.Offset(1)
	b.Count(1)
	b.UserID(1)
	b.PeerID(1)
	b.StartMessageID(1)
	b.Rev(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["start_message_id"], 1)
	assert.Equal(t, b.Params["rev"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetHistoryAttachmentsBulder(t *testing.T) {
	b := params.NewMessagesGetHistoryAttachmentsBulder()

	b.PeerID(1)
	b.MediaType("text")
	b.StartFrom("text")
	b.Count(1)
	b.PhotoSizes(true)
	b.Fields([]string{"test"})
	b.GroupID(1)
	b.PreserveOrder(true)
	b.MaxForwardsLevel(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["media_type"], "text")
	assert.Equal(t, b.Params["start_from"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["photo_sizes"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["preserve_order"], true)
	assert.Equal(t, b.Params["max_forwards_level"], 1)
}

func TestMessagesGetInviteLinkBulder(t *testing.T) {
	b := params.NewMessagesGetInviteLinkBulder()

	b.PeerID(1)
	b.Reset(true)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["reset"], true)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesGetLastActivityBulder(t *testing.T) {
	b := params.NewMessagesGetLastActivityBulder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestMessagesGetLongPollHistoryBulder(t *testing.T) {
	b := params.NewMessagesGetLongPollHistoryBulder()

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

	assert.Equal(t, b.Params["ts"], 1)
	assert.Equal(t, b.Params["pts"], 1)
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["onlines"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["events_limit"], 1)
	assert.Equal(t, b.Params["msgs_limit"], 1)
	assert.Equal(t, b.Params["max_msg_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["lp_version"], 1)
	assert.Equal(t, b.Params["last_n"], 1)
	assert.Equal(t, b.Params["credentials"], true)
}

func TestMessagesGetLongPollServerBulder(t *testing.T) {
	b := params.NewMessagesGetLongPollServerBulder()

	b.NeedPts(true)
	b.GroupID(1)
	b.LpVersion(1)

	assert.Equal(t, b.Params["need_pts"], true)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["lp_version"], 1)
}

func TestMessagesIsMessagesFromGroupAllowedBulder(t *testing.T) {
	b := params.NewMessagesIsMessagesFromGroupAllowedBulder()

	b.GroupID(1)
	b.UserID(1)

	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
}

func TestMessagesJoinChatByInviteLinkBulder(t *testing.T) {
	b := params.NewMessagesJoinChatByInviteLinkBulder()

	b.Link("text")

	assert.Equal(t, b.Params["link"], "text")
}

func TestMessagesMarkAsAnsweredConversationBulder(t *testing.T) {
	b := params.NewMessagesMarkAsAnsweredConversationBulder()

	b.PeerID(1)
	b.Answered(true)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["answered"], true)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesMarkAsImportantBulder(t *testing.T) {
	b := params.NewMessagesMarkAsImportantBulder()

	b.MessageIDs([]int{1})
	b.Important(1)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["important"], 1)
}

func TestMessagesMarkAsImportantConversationBulder(t *testing.T) {
	b := params.NewMessagesMarkAsImportantConversationBulder()

	b.PeerID(1)
	b.Important(true)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["important"], true)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesMarkAsReadBulder(t *testing.T) {
	b := params.NewMessagesMarkAsReadBulder()

	b.MessageIDs([]int{1})
	b.PeerID(1)
	b.StartMessageID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["message_ids"], []int{1})
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["start_message_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesPinBulder(t *testing.T) {
	b := params.NewMessagesPinBulder()

	b.PeerID(1)
	b.MessageID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["message_id"], 1)
}

func TestMessagesRemoveChatUserBulder(t *testing.T) {
	b := params.NewMessagesRemoveChatUserBulder()

	b.ChatID(1)
	b.UserID(1)
	b.MemberID(1)

	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["member_id"], 1)
}

func TestMessagesRestoreBulder(t *testing.T) {
	b := params.NewMessagesRestoreBulder()

	b.MessageID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["message_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSearchBulder(t *testing.T) {
	b := params.NewMessagesSearchBulder()

	b.Q("text")
	b.PeerID(1)
	b.Date(1)
	b.PreviewLength(1)
	b.Offset(1)
	b.Count(1)
	b.Extended(true)
	b.Fields([]string{"text"})
	b.GroupID(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["date"], 1)
	assert.Equal(t, b.Params["preview_length"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"text"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSearchConversationsBulder(t *testing.T) {
	b := params.NewMessagesSearchConversationsBulder()

	b.Q("text")
	b.Count(1)
	b.Extended(true)
	b.Fields([]string{"test"})
	b.GroupID(1)

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSendBulder(t *testing.T) {
	b := params.NewMessagesSendBulder()

	b.UserID(1)
	b.RandomID(1)
	b.PeerID(1)
	b.Domain("text")
	b.ChatID(1)
	b.UserIDs([]int{1})
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
	b.Payload("text")
	b.DontParseLinks(true)
	b.DisableMentions(true)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["random_id"], 1)
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["domain"], "text")
	assert.Equal(t, b.Params["chat_id"], 1)
	assert.Equal(t, b.Params["user_ids"], []int{1})
	assert.Equal(t, b.Params["message"], "text")
	assert.Equal(t, b.Params["lat"], 1.1)
	assert.Equal(t, b.Params["long"], 1.1)
	assert.Equal(t, b.Params["attachment"], "text")
	assert.Equal(t, b.Params["reply_to"], 1)
	assert.Equal(t, b.Params["forward_messages"], []int{1})
	assert.Equal(t, b.Params["forward"], "text")
	assert.Equal(t, b.Params["sticker_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["keyboard"], "text")
	assert.Equal(t, b.Params["payload"], "text")
	assert.Equal(t, b.Params["dont_parse_links"], true)
	assert.Equal(t, b.Params["disable_mentions"], true)
}

func TestMessagesSetActivityBulder(t *testing.T) {
	b := params.NewMessagesSetActivityBulder()

	b.UserID(1)
	b.Type("text")
	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}

func TestMessagesSetChatPhotoBulder(t *testing.T) {
	b := params.NewMessagesSetChatPhotoBulder()

	b.File("text")

	assert.Equal(t, b.Params["file"], "text")
}

func TestMessagesUnpinBulder(t *testing.T) {
	b := params.NewMessagesUnpinBulder()

	b.PeerID(1)
	b.GroupID(1)

	assert.Equal(t, b.Params["peer_id"], 1)
	assert.Equal(t, b.Params["group_id"], 1)
}
