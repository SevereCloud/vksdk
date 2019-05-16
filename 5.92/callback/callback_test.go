package callback

import (
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestCallback_MessageNew(t *testing.T) {
	t.Run("MessageNew", func(t *testing.T) {
		cb := &Callback{}
		cb.MessageNew(func(obj object.MessageNewObject, groupID int) {})
	})

}

func TestCallback_MessageReply(t *testing.T) {
	t.Run("MessageReply", func(t *testing.T) {
		cb := &Callback{}
		cb.MessageReply(func(obj object.MessageReplyObject, groupID int) {})
	})

}

func TestCallback_MessageEdit(t *testing.T) {
	t.Run("MessageEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.MessageEdit(func(obj object.MessageEditObject, groupID int) {})
	})

}

func TestCallback_MessageAllow(t *testing.T) {
	t.Run("MessageAllow", func(t *testing.T) {
		cb := &Callback{}
		cb.MessageAllow(func(obj object.MessageAllowObject, groupID int) {})
	})

}

func TestCallback_MessageDeny(t *testing.T) {
	t.Run("MessageDeny", func(t *testing.T) {
		cb := &Callback{}
		cb.MessageDeny(func(obj object.MessageDenyObject, groupID int) {})
	})

}

func TestCallback_MessageTypingState(t *testing.T) {
	t.Run("MessageTypingState", func(t *testing.T) {
		cb := &Callback{}
		cb.MessageTypingState(func(obj object.MessageTypingStateObject, groupID int) {})
	})

}

func TestCallback_PhotoNew(t *testing.T) {
	t.Run("PhotoNew", func(t *testing.T) {
		cb := &Callback{}
		cb.PhotoNew(func(obj object.PhotoNewObject, groupID int) {})
	})

}

func TestCallback_PhotoCommentNew(t *testing.T) {
	t.Run("PhotoCommentNew", func(t *testing.T) {
		cb := &Callback{}
		cb.PhotoCommentNew(func(obj object.PhotoCommentNewObject, groupID int) {})
	})

}

func TestCallback_PhotoCommentEdit(t *testing.T) {
	t.Run("PhotoCommentEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.PhotoCommentEdit(func(obj object.PhotoCommentEditObject, groupID int) {})
	})

}

func TestCallback_PhotoCommentRestore(t *testing.T) {
	t.Run("PhotoCommentRestore", func(t *testing.T) {
		cb := &Callback{}
		cb.PhotoCommentRestore(func(obj object.PhotoCommentRestoreObject, groupID int) {})
	})

}

func TestCallback_PhotoCommentDelete(t *testing.T) {
	t.Run("PhotoCommentDelete", func(t *testing.T) {
		cb := &Callback{}
		cb.PhotoCommentDelete(func(obj object.PhotoCommentDeleteObject, groupID int) {})
	})

}

func TestCallback_AudioNew(t *testing.T) {
	t.Run("AudioNew", func(t *testing.T) {
		cb := &Callback{}
		cb.AudioNew(func(obj object.AudioNewObject, groupID int) {})
	})

}

func TestCallback_VideoNew(t *testing.T) {
	t.Run("VideoNew", func(t *testing.T) {
		cb := &Callback{}
		cb.VideoNew(func(obj object.VideoNewObject, groupID int) {})
	})

}

func TestCallback_VideoCommentNew(t *testing.T) {
	t.Run("VideoCommentNew", func(t *testing.T) {
		cb := &Callback{}
		cb.VideoCommentNew(func(obj object.VideoCommentNewObject, groupID int) {})
	})

}

func TestCallback_VideoCommentEdit(t *testing.T) {
	t.Run("VideoCommentEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.VideoCommentEdit(func(obj object.VideoCommentEditObject, groupID int) {})
	})

}

func TestCallback_VideoCommentRestore(t *testing.T) {
	t.Run("VideoCommentRestore", func(t *testing.T) {
		cb := &Callback{}
		cb.VideoCommentRestore(func(obj object.VideoCommentRestoreObject, groupID int) {})
	})

}

func TestCallback_VideoCommentDelete(t *testing.T) {
	t.Run("VideoCommentDelete", func(t *testing.T) {
		cb := &Callback{}
		cb.VideoCommentDelete(func(obj object.VideoCommentDeleteObject, groupID int) {})
	})

}

func TestCallback_WallPostNew(t *testing.T) {
	t.Run("WallPostNew", func(t *testing.T) {
		cb := &Callback{}
		cb.WallPostNew(func(obj object.WallPostNewObject, groupID int) {})
	})

}

func TestCallback_WallRepost(t *testing.T) {
	t.Run("WallRepost", func(t *testing.T) {
		cb := &Callback{}
		cb.WallRepost(func(obj object.WallRepostObject, groupID int) {})
	})

}

func TestCallback_WallReplyNew(t *testing.T) {
	t.Run("WallReplyNew", func(t *testing.T) {
		cb := &Callback{}
		cb.WallReplyNew(func(obj object.WallReplyNewObject, groupID int) {})
	})

}

func TestCallback_WallReplyEdit(t *testing.T) {
	t.Run("WallReplyEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.WallReplyEdit(func(obj object.WallReplyEditObject, groupID int) {})
	})

}

func TestCallback_WallReplyRestore(t *testing.T) {
	t.Run("WallReplyRestore", func(t *testing.T) {
		cb := &Callback{}
		cb.WallReplyRestore(func(obj object.WallReplyRestoreObject, groupID int) {})
	})

}

func TestCallback_WallReplyDelete(t *testing.T) {
	t.Run("WallReplyDelete", func(t *testing.T) {
		cb := &Callback{}
		cb.WallReplyDelete(func(obj object.WallReplyDeleteObject, groupID int) {})
	})

}

func TestCallback_BoardPostNew(t *testing.T) {
	t.Run("BoardPostNew", func(t *testing.T) {
		cb := &Callback{}
		cb.BoardPostNew(func(obj object.BoardPostNewObject, groupID int) {})
	})

}

func TestCallback_BoardPostEdit(t *testing.T) {
	t.Run("BoardPostEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.BoardPostEdit(func(obj object.BoardPostEditObject, groupID int) {})
	})

}

func TestCallback_BoardPostRestore(t *testing.T) {
	t.Run("BoardPostRestore", func(t *testing.T) {
		cb := &Callback{}
		cb.BoardPostRestore(func(obj object.BoardPostRestoreObject, groupID int) {})
	})

}

func TestCallback_BoardPostDelete(t *testing.T) {
	t.Run("BoardPostDelete", func(t *testing.T) {
		cb := &Callback{}
		cb.BoardPostDelete(func(obj object.BoardPostDeleteObject, groupID int) {})
	})

}

func TestCallback_MarketCommentNew(t *testing.T) {
	t.Run("MarketCommentNew", func(t *testing.T) {
		cb := &Callback{}
		cb.MarketCommentNew(func(obj object.MarketCommentNewObject, groupID int) {})
	})

}

func TestCallback_MarketCommentEdit(t *testing.T) {
	t.Run("MarketCommentEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.MarketCommentEdit(func(obj object.MarketCommentEditObject, groupID int) {})
	})

}

func TestCallback_MarketCommentRestore(t *testing.T) {
	t.Run("MarketCommentRestore", func(t *testing.T) {
		cb := &Callback{}
		cb.MarketCommentRestore(func(obj object.MarketCommentRestoreObject, groupID int) {})
	})

}

func TestCallback_MarketCommentDelete(t *testing.T) {
	t.Run("MarketCommentDelete", func(t *testing.T) {
		cb := &Callback{}
		cb.MarketCommentDelete(func(obj object.MarketCommentDeleteObject, groupID int) {})
	})

}

func TestCallback_GroupLeave(t *testing.T) {
	t.Run("GroupLeave", func(t *testing.T) {
		cb := &Callback{}
		cb.GroupLeave(func(obj object.GroupLeaveObject, groupID int) {})
	})

}

func TestCallback_GroupJoin(t *testing.T) {
	t.Run("GroupJoin", func(t *testing.T) {
		cb := &Callback{}
		cb.GroupJoin(func(obj object.GroupJoinObject, groupID int) {})
	})

}

func TestCallback_UserBlock(t *testing.T) {
	t.Run("UserBlock", func(t *testing.T) {
		cb := &Callback{}
		cb.UserBlock(func(obj object.UserBlockObject, groupID int) {})
	})

}

func TestCallback_UserUnblock(t *testing.T) {
	t.Run("UserUnblock", func(t *testing.T) {
		cb := &Callback{}
		cb.UserUnblock(func(obj object.UserUnblockObject, groupID int) {})
	})

}

func TestCallback_PollVoteNew(t *testing.T) {
	t.Run("PollVoteNew", func(t *testing.T) {
		cb := &Callback{}
		cb.PollVoteNew(func(obj object.PollVoteNewObject, groupID int) {})
	})

}

func TestCallback_GroupOfficersEdit(t *testing.T) {
	t.Run("GroupOfficersEdit", func(t *testing.T) {
		cb := &Callback{}
		cb.GroupOfficersEdit(func(obj object.GroupOfficersEditObject, groupID int) {})
	})

}

func TestCallback_GroupChangeSettings(t *testing.T) {
	t.Run("GroupChangeSettings", func(t *testing.T) {
		cb := &Callback{}
		cb.GroupChangeSettings(func(obj object.GroupChangeSettingsObject, groupID int) {})
	})

}

func TestCallback_GroupChangePhoto(t *testing.T) {
	t.Run("GroupChangePhoto", func(t *testing.T) {
		cb := &Callback{}
		cb.GroupChangePhoto(func(obj object.GroupChangePhotoObject, groupID int) {})
	})

}

func TestCallback_VkpayTransaction(t *testing.T) {
	t.Run("VkpayTransaction", func(t *testing.T) {
		cb := &Callback{}
		cb.VkpayTransaction(func(obj object.VkpayTransactionObject, groupID int) {})
	})

}

func TestCallback_LeadFormsNew(t *testing.T) {
	t.Run("LeadFormsNew", func(t *testing.T) {
		cb := &Callback{}
		cb.LeadFormsNew(func(obj object.LeadFormsNewObject, groupID int) {})
	})

}
