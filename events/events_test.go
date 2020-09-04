package events_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/events"
)

const GID = 123456

func TestFuncList_HandlerMessageNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		t.Helper()

		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
		assert.NotEmpty(t, obj.Message)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		t.Helper()

		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_new",
			Object:  []byte(`{"message":{"date":1,"from_id":1,"id":1,"peer_id":1,"text":"{}"},"client_info":{}}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageReply(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageReply(func(ctx context.Context, obj events.MessageReplyObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_reply",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_reply",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageEdit(func(ctx context.Context, obj events.MessageEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageAllow(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageAllow(func(ctx context.Context, obj events.MessageAllowObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_allow",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_allow",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageDeny(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageDeny(func(ctx context.Context, obj events.MessageDenyObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_deny",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_deny",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageTypingState(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageTypingState(func(ctx context.Context, obj events.MessageTypingStateObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_typing_state",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_typing_state",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageEvent(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageEvent(func(ctx context.Context, obj events.MessageEventObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_event",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_event",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoNew(func(ctx context.Context, obj events.PhotoNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "photo_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "photo_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentNew(func(ctx context.Context, obj events.PhotoCommentNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "photo_comment_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "photo_comment_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentEdit(func(ctx context.Context, obj events.PhotoCommentEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "photo_comment_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "photo_comment_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentRestore(func(ctx context.Context, obj events.PhotoCommentRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "photo_comment_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "photo_comment_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentDelete(func(ctx context.Context, obj events.PhotoCommentDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "photo_comment_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "photo_comment_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerAudioNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.AudioNew(func(ctx context.Context, obj events.AudioNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "audio_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "audio_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoNew(func(ctx context.Context, obj events.VideoNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "video_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "video_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentNew(func(ctx context.Context, obj events.VideoCommentNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "video_comment_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "video_comment_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentEdit(func(ctx context.Context, obj events.VideoCommentEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "video_comment_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "video_comment_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentRestore(func(ctx context.Context, obj events.VideoCommentRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "video_comment_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "video_comment_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentDelete(func(ctx context.Context, obj events.VideoCommentDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "video_comment_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "video_comment_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallPostNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallPostNew(func(ctx context.Context, obj events.WallPostNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wall_post_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wall_post_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallRepost(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallRepost(func(ctx context.Context, obj events.WallRepostObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wall_repost",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wall_repost",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyNew(func(ctx context.Context, obj events.WallReplyNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wall_reply_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wall_reply_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyEdit(func(ctx context.Context, obj events.WallReplyEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wall_reply_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wall_reply_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyRestore(func(ctx context.Context, obj events.WallReplyRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wall_reply_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wall_reply_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyDelete(func(ctx context.Context, obj events.WallReplyDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wall_reply_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wall_reply_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostNew(func(ctx context.Context, obj events.BoardPostNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "board_post_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "board_post_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostEdit(func(ctx context.Context, obj events.BoardPostEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "board_post_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "board_post_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostRestore(func(ctx context.Context, obj events.BoardPostRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "board_post_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "board_post_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostDelete(func(ctx context.Context, obj events.BoardPostDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "board_post_delete",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "board_post_delete",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentNew(func(ctx context.Context, obj events.MarketCommentNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "market_comment_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "market_comment_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentEdit(func(ctx context.Context, obj events.MarketCommentEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "market_comment_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "market_comment_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentRestore(func(ctx context.Context, obj events.MarketCommentRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "market_comment_restore",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "market_comment_restore",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentDelete(func(ctx context.Context, obj events.MarketCommentDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})
	assert.Equal(t, []events.EventType{events.EventMarketCommentDelete}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMarketCommentDelete,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketOrderNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketOrderNew(func(ctx context.Context, obj events.MarketOrderNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMarketOrderNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMarketOrderNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketOrderEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketOrderEdit(func(ctx context.Context, obj events.MarketOrderEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})
	assert.Equal(t, []events.EventType{events.EventMarketOrderEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMarketOrderEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMarketOrderEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupLeave(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupLeave(func(ctx context.Context, obj events.GroupLeaveObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "group_leave",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "group_leave",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupJoin(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupJoin(func(ctx context.Context, obj events.GroupJoinObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "group_join",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "group_join",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerUserBlock(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.UserBlock(func(ctx context.Context, obj events.UserBlockObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "user_block",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "user_block",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerUserUnblock(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.UserUnblock(func(ctx context.Context, obj events.UserUnblockObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "user_unblock",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "user_unblock",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPollVoteNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PollVoteNew(func(ctx context.Context, obj events.PollVoteNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "poll_vote_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "poll_vote_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupOfficersEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupOfficersEdit(func(ctx context.Context, obj events.GroupOfficersEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "group_officers_edit",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "group_officers_edit",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupChangeSettings(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupChangeSettings(func(ctx context.Context, obj events.GroupChangeSettingsObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "group_change_settings",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "group_change_settings",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupChangePhoto(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupChangePhoto(func(ctx context.Context, obj events.GroupChangePhotoObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "group_change_photo",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "group_change_photo",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVkpayTransaction(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VkpayTransaction(func(ctx context.Context, obj events.VkpayTransactionObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "vkpay_transaction",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "vkpay_transaction",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLeadFormsNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.LeadFormsNew(func(ctx context.Context, obj events.LeadFormsNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "lead_forms_new",
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "lead_forms_new",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerAppPayload(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.AppPayload(func(ctx context.Context, obj events.AppPayloadObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "app_payload",
			Object:  []byte(`{"user_id":117253521,"app_id":6703670,"payload":"{\"foo\":\"bar\"}"}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "app_payload",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageRead(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageRead(func(ctx context.Context, obj events.MessageReadObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "message_read",
			Object:  []byte(`{"from_id":1,"peer_id":1,"read_message_id":1}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "message_read",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLikeAdd(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.LikeAdd(func(ctx context.Context, obj events.LikeAddObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "like_add",
			Object:  []byte(`{"liker_id": 574423462,"object_type": "photo","object_owner_id": -178044536,"object_id": 457242474,"thread_reply_id": 0}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "like_add",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLikeRemove(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.LikeRemove(func(ctx context.Context, obj events.LikeRemoveObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, groupID, GID)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "like_remove",
			Object:  []byte(`{"liker_id": 574423462,"object_type": "photo","object_owner_id": -178044536,"object_id": 457242474,"thread_reply_id": 0}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "like_remove",
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_OnEvent(t *testing.T) {
	t.Parallel()

	var fl events.FuncList

	fl.OnEvent("wtf_event", func(_ context.Context, e events.GroupEvent) {
		assert.NotEmpty(t, e)
	})

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    "wtf_event",
			Object:  []byte(`{"from_id":1,"peer_id":1,"read_message_id":1}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   "wtf_event",
			Object: []byte(""),
		},
		false,
	)
}
