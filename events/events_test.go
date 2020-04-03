package events_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/events"
	"github.com/SevereCloud/vksdk/object"
)

const GID = 123456

func TestFuncList_HandlerMessageNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageNew(func(obj object.MessageNewObject, groupID int) {
		t.Helper()

		assert.Equal(t, groupID, GID)
		assert.NotEmpty(t, obj.Message)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		t.Helper()

		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_new",
			Object:  []byte(`{"message":{"date":1,"from_id":1,"id":1,"peer_id":1,"text":"{}"},"client_info":{}}`),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:    "message_new",
			Object:  []byte(`{"date":1,"from_id":1,"id":1,"peer_id":1,"text":"{}"}`),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageReply(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageReply(func(obj object.MessageReplyObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_reply",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_reply",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageEdit(func(obj object.MessageEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageAllow(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageAllow(func(obj object.MessageAllowObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_allow",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_allow",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageDeny(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageDeny(func(obj object.MessageDenyObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_deny",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_deny",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageTypingState(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageTypingState(func(obj object.MessageTypingStateObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_typing_state",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_typing_state",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.PhotoNew(func(obj object.PhotoNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "photo_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "photo_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.PhotoCommentNew(func(obj object.PhotoCommentNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "photo_comment_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "photo_comment_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.PhotoCommentEdit(func(obj object.PhotoCommentEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "photo_comment_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "photo_comment_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentRestore(t *testing.T) {
	fl := events.NewFuncList()

	fl.PhotoCommentRestore(func(obj object.PhotoCommentRestoreObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "photo_comment_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "photo_comment_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentDelete(t *testing.T) {
	fl := events.NewFuncList()

	fl.PhotoCommentDelete(func(obj object.PhotoCommentDeleteObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "photo_comment_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "photo_comment_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerAudioNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.AudioNew(func(obj object.AudioNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "audio_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "audio_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.VideoNew(func(obj object.VideoNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "video_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "video_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.VideoCommentNew(func(obj object.VideoCommentNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "video_comment_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "video_comment_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.VideoCommentEdit(func(obj object.VideoCommentEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "video_comment_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "video_comment_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentRestore(t *testing.T) {
	fl := events.NewFuncList()

	fl.VideoCommentRestore(func(obj object.VideoCommentRestoreObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "video_comment_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "video_comment_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentDelete(t *testing.T) {
	fl := events.NewFuncList()

	fl.VideoCommentDelete(func(obj object.VideoCommentDeleteObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "video_comment_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "video_comment_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallPostNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.WallPostNew(func(obj object.WallPostNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wall_post_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wall_post_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallRepost(t *testing.T) {
	fl := events.NewFuncList()

	fl.WallRepost(func(obj object.WallRepostObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wall_repost",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wall_repost",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.WallReplyNew(func(obj object.WallReplyNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wall_reply_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wall_reply_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.WallReplyEdit(func(obj object.WallReplyEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wall_reply_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wall_reply_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyRestore(t *testing.T) {
	fl := events.NewFuncList()

	fl.WallReplyRestore(func(obj object.WallReplyRestoreObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wall_reply_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wall_reply_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyDelete(t *testing.T) {
	fl := events.NewFuncList()

	fl.WallReplyDelete(func(obj object.WallReplyDeleteObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wall_reply_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wall_reply_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.BoardPostNew(func(obj object.BoardPostNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "board_post_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "board_post_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.BoardPostEdit(func(obj object.BoardPostEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "board_post_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "board_post_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostRestore(t *testing.T) {
	fl := events.NewFuncList()

	fl.BoardPostRestore(func(obj object.BoardPostRestoreObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "board_post_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "board_post_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostDelete(t *testing.T) {
	fl := events.NewFuncList()

	fl.BoardPostDelete(func(obj object.BoardPostDeleteObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "board_post_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "board_post_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.MarketCommentNew(func(obj object.MarketCommentNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "market_comment_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "market_comment_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.MarketCommentEdit(func(obj object.MarketCommentEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "market_comment_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "market_comment_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentRestore(t *testing.T) {
	fl := events.NewFuncList()

	fl.MarketCommentRestore(func(obj object.MarketCommentRestoreObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "market_comment_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "market_comment_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentDelete(t *testing.T) {
	fl := events.NewFuncList()

	fl.MarketCommentDelete(func(obj object.MarketCommentDeleteObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "market_comment_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "market_comment_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupLeave(t *testing.T) {
	fl := events.NewFuncList()

	fl.GroupLeave(func(obj object.GroupLeaveObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "group_leave",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "group_leave",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupJoin(t *testing.T) {
	fl := events.NewFuncList()

	fl.GroupJoin(func(obj object.GroupJoinObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "group_join",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "group_join",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerUserBlock(t *testing.T) {
	fl := events.NewFuncList()

	fl.UserBlock(func(obj object.UserBlockObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "user_block",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "user_block",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerUserUnblock(t *testing.T) {
	fl := events.NewFuncList()

	fl.UserUnblock(func(obj object.UserUnblockObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "user_unblock",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "user_unblock",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPollVoteNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.PollVoteNew(func(obj object.PollVoteNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "poll_vote_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "poll_vote_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupOfficersEdit(t *testing.T) {
	fl := events.NewFuncList()

	fl.GroupOfficersEdit(func(obj object.GroupOfficersEditObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "group_officers_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "group_officers_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupChangeSettings(t *testing.T) {
	fl := events.NewFuncList()

	fl.GroupChangeSettings(func(obj object.GroupChangeSettingsObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "group_change_settings",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "group_change_settings",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupChangePhoto(t *testing.T) {
	fl := events.NewFuncList()

	fl.GroupChangePhoto(func(obj object.GroupChangePhotoObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "group_change_photo",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "group_change_photo",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVkpayTransaction(t *testing.T) {
	fl := events.NewFuncList()

	fl.VkpayTransaction(func(obj object.VkpayTransactionObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "vkpay_transaction",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "vkpay_transaction",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLeadFormsNew(t *testing.T) {
	fl := events.NewFuncList()

	fl.LeadFormsNew(func(obj object.LeadFormsNewObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "lead_forms_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "lead_forms_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerAppPayload(t *testing.T) {
	fl := events.NewFuncList()

	fl.AppPayload(func(obj object.AppPayloadObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "app_payload",
			Object:  []byte(`{"user_id":117253521,"app_id":6703670,"payload":"{\"foo\":\"bar\"}"}`),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "app_payload",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageRead(t *testing.T) {
	fl := events.NewFuncList()

	fl.MessageRead(func(obj object.MessageReadObject, groupID int) {
		assert.Equal(t, groupID, GID)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "message_read",
			Object:  []byte(`{"from_id":1,"peer_id":1,"read_message_id":1}`),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "message_read",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_OnEvent(t *testing.T) {
	var fl events.FuncList

	fl.OnEvent("wtf_event", func(e object.GroupEvent) {
		assert.NotEmpty(t, e)
	})

	f := func(e object.GroupEvent, wantErr bool) {
		if err := fl.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "wtf_event",
			Object:  []byte(`{"from_id":1,"peer_id":1,"read_message_id":1}`),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "wtf_event",
			Object: []byte(""),
		},
		false,
	)
}
