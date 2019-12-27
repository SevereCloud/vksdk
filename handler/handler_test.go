package handler_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/handler"
	"github.com/SevereCloud/vksdk/object"
)

const GID = 123456

func TestFuncList_HandlerMessageNew(t *testing.T) {
	funcList := handler.FuncList{
		MessageNew: []object.MessageNewFunc{
			func(obj object.MessageNewObject, groupID int) {
				t.Helper()

				assert.Equal(t, groupID, GID)
				assert.NotEmpty(t, obj.Message)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		t.Helper()

		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MessageReply: []object.MessageReplyFunc{
			func(obj object.MessageReplyObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MessageEdit: []object.MessageEditFunc{
			func(obj object.MessageEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MessageAllow: []object.MessageAllowFunc{
			func(obj object.MessageAllowObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MessageDeny: []object.MessageDenyFunc{
			func(obj object.MessageDenyObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		object.GroupEvent{
			Type:    "messages_deny",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		object.GroupEvent{
			Type:   "messages_deny",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageTypingState(t *testing.T) {
	funcList := handler.FuncList{
		MessageTypingState: []object.MessageTypingStateFunc{
			func(obj object.MessageTypingStateObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		PhotoNew: []object.PhotoNewFunc{
			func(obj object.PhotoNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		PhotoCommentNew: []object.PhotoCommentNewFunc{
			func(obj object.PhotoCommentNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		PhotoCommentEdit: []object.PhotoCommentEditFunc{
			func(obj object.PhotoCommentEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		PhotoCommentRestore: []object.PhotoCommentRestoreFunc{
			func(obj object.PhotoCommentRestoreObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		PhotoCommentDelete: []object.PhotoCommentDeleteFunc{
			func(obj object.PhotoCommentDeleteObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		AudioNew: []object.AudioNewFunc{
			func(obj object.AudioNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		VideoNew: []object.VideoNewFunc{
			func(obj object.VideoNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		VideoCommentNew: []object.VideoCommentNewFunc{
			func(obj object.VideoCommentNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		VideoCommentEdit: []object.VideoCommentEditFunc{
			func(obj object.VideoCommentEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		VideoCommentRestore: []object.VideoCommentRestoreFunc{
			func(obj object.VideoCommentRestoreObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		VideoCommentDelete: []object.VideoCommentDeleteFunc{
			func(obj object.VideoCommentDeleteObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		WallPostNew: []object.WallPostNewFunc{
			func(obj object.WallPostNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		WallRepost: []object.WallRepostFunc{
			func(obj object.WallRepostObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		WallReplyNew: []object.WallReplyNewFunc{
			func(obj object.WallReplyNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		WallReplyEdit: []object.WallReplyEditFunc{
			func(obj object.WallReplyEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		WallReplyRestore: []object.WallReplyRestoreFunc{
			func(obj object.WallReplyRestoreObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		WallReplyDelete: []object.WallReplyDeleteFunc{
			func(obj object.WallReplyDeleteObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		BoardPostNew: []object.BoardPostNewFunc{
			func(obj object.BoardPostNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		BoardPostEdit: []object.BoardPostEditFunc{
			func(obj object.BoardPostEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		BoardPostRestore: []object.BoardPostRestoreFunc{
			func(obj object.BoardPostRestoreObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		BoardPostDelete: []object.BoardPostDeleteFunc{
			func(obj object.BoardPostDeleteObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MarketCommentNew: []object.MarketCommentNewFunc{
			func(obj object.MarketCommentNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MarketCommentEdit: []object.MarketCommentEditFunc{
			func(obj object.MarketCommentEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MarketCommentRestore: []object.MarketCommentRestoreFunc{
			func(obj object.MarketCommentRestoreObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		MarketCommentDelete: []object.MarketCommentDeleteFunc{
			func(obj object.MarketCommentDeleteObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		GroupLeave: []object.GroupLeaveFunc{
			func(obj object.GroupLeaveObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		GroupJoin: []object.GroupJoinFunc{
			func(obj object.GroupJoinObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		UserBlock: []object.UserBlockFunc{
			func(obj object.UserBlockObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		UserUnblock: []object.UserUnblockFunc{
			func(obj object.UserUnblockObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		PollVoteNew: []object.PollVoteNewFunc{
			func(obj object.PollVoteNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		GroupOfficersEdit: []object.GroupOfficersEditFunc{
			func(obj object.GroupOfficersEditObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		GroupChangeSettings: []object.GroupChangeSettingsFunc{
			func(obj object.GroupChangeSettingsObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		GroupChangePhoto: []object.GroupChangePhotoFunc{
			func(obj object.GroupChangePhotoObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		VkpayTransaction: []object.VkpayTransactionFunc{
			func(obj object.VkpayTransactionObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
	funcList := handler.FuncList{
		LeadFormsNew: []object.LeadFormsNewFunc{
			func(obj object.LeadFormsNewObject, groupID int) {
				assert.Equal(t, groupID, GID)
			},
		},
	}

	f := func(e object.GroupEvent, wantErr bool) {
		if err := funcList.Handler(e); (err != nil) != wantErr {
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
