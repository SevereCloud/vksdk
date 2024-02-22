package events_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/SevereCloud/vksdk/v2/events"
)

const GID = 123456

func TestFuncList_HandlerMessageNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageNew(func(ctx context.Context, obj events.MessageNewObject) {
		t.Helper()

		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
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
			Type:    events.EventMessageNew,
			Object:  []byte(`{"message":{"date":1,"from_id":1,"id":1,"peer_id":1,"text":"{}"},"client_info":{}}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageNew,
			Object:  []byte(`{"message":{"date":1,"from_id":1,"id":1,"peer_id":1,"text":"{}"},"client_info":{}}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageReply(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageReply(func(ctx context.Context, _ events.MessageReplyObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageReply}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageReply,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageReply,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageReply,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageEdit(func(ctx context.Context, _ events.MessageEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageAllow(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageAllow(func(ctx context.Context, _ events.MessageAllowObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageAllow}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageAllow,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageAllow,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageAllow,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageDeny(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageDeny(func(ctx context.Context, _ events.MessageDenyObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageDeny}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageDeny,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageDeny,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageDeny,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageTypingState(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageTypingState(func(ctx context.Context, _ events.MessageTypingStateObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageTypingState}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageTypingState,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageTypingState,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageTypingState,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageEvent(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageEvent(func(ctx context.Context, _ events.MessageEventObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageEvent}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageEvent,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageEvent,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageEvent,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoNew(func(ctx context.Context, _ events.PhotoNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventPhotoNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventPhotoNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventPhotoNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventPhotoNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentNew(func(ctx context.Context, _ events.PhotoCommentNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventPhotoCommentNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventPhotoCommentNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentEdit(func(ctx context.Context, _ events.PhotoCommentEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventPhotoCommentEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventPhotoCommentEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentRestore(func(ctx context.Context, _ events.PhotoCommentRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventPhotoCommentRestore}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventPhotoCommentRestore,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPhotoCommentDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PhotoCommentDelete(func(ctx context.Context, _ events.PhotoCommentDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventPhotoCommentDelete}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventPhotoCommentDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventPhotoCommentDelete,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerAudioNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.AudioNew(func(ctx context.Context, _ events.AudioNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventAudioNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventAudioNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventAudioNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventAudioNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoNew(func(ctx context.Context, _ events.VideoNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventVideoNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventVideoNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventVideoNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventVideoNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentNew(func(ctx context.Context, _ events.VideoCommentNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventVideoCommentNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventVideoCommentNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentEdit(func(ctx context.Context, _ events.VideoCommentEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventVideoCommentEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventVideoCommentEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentRestore(func(ctx context.Context, _ events.VideoCommentRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventVideoCommentRestore}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventVideoCommentRestore,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVideoCommentDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VideoCommentDelete(func(ctx context.Context, _ events.VideoCommentDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventVideoCommentDelete}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventVideoCommentDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventVideoCommentDelete,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallPostNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallPostNew(func(ctx context.Context, _ events.WallPostNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventWallPostNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventWallPostNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventWallPostNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventWallPostNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallRepost(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallRepost(func(ctx context.Context, _ events.WallRepostObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventWallRepost}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventWallRepost,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventWallRepost,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventWallRepost,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyNew(func(ctx context.Context, _ events.WallReplyNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventWallReplyNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventWallReplyNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventWallReplyNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventWallReplyNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyEdit(func(ctx context.Context, _ events.WallReplyEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventWallReplyEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventWallReplyEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventWallReplyEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventWallReplyEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyRestore(func(ctx context.Context, _ events.WallReplyRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventWallReplyRestore}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventWallReplyRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventWallReplyRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventWallReplyRestore,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerWallReplyDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.WallReplyDelete(func(ctx context.Context, _ events.WallReplyDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventWallReplyDelete}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventWallReplyDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventWallReplyDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventWallReplyDelete,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostNew(func(ctx context.Context, _ events.BoardPostNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventBoardPostNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventBoardPostNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventBoardPostNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventBoardPostNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostEdit(func(ctx context.Context, _ events.BoardPostEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventBoardPostEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventBoardPostEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventBoardPostEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventBoardPostEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostRestore(func(ctx context.Context, _ events.BoardPostRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventBoardPostRestore}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventBoardPostRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventBoardPostRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventBoardPostRestore,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerBoardPostDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.BoardPostDelete(func(ctx context.Context, _ events.BoardPostDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventBoardPostDelete}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventBoardPostDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventBoardPostDelete,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventBoardPostDelete,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentNew(func(ctx context.Context, _ events.MarketCommentNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMarketCommentNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMarketCommentNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentEdit(func(ctx context.Context, _ events.MarketCommentEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMarketCommentEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMarketCommentEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentRestore(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentRestore(func(ctx context.Context, _ events.MarketCommentRestoreObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMarketCommentRestore}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMarketCommentRestore,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMarketCommentRestore,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMarketCommentDelete(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MarketCommentDelete(func(ctx context.Context, _ events.MarketCommentDeleteObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
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
	fl.Goroutine(true)
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

	fl.MarketOrderNew(func(ctx context.Context, _ events.MarketOrderNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMarketOrderNew}, fl.ListEvents())

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
	fl.Goroutine(true)
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

	fl.MarketOrderEdit(func(ctx context.Context, _ events.MarketOrderEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
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
	fl.Goroutine(true)
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

	fl.GroupLeave(func(ctx context.Context, _ events.GroupLeaveObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventGroupLeave}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventGroupLeave,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventGroupLeave,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventGroupLeave,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupJoin(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupJoin(func(ctx context.Context, _ events.GroupJoinObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventGroupJoin}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventGroupJoin,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventGroupJoin,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventGroupJoin,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerUserBlock(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.UserBlock(func(ctx context.Context, _ events.UserBlockObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventUserBlock}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventUserBlock,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventUserBlock,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventUserBlock,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerUserUnblock(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.UserUnblock(func(ctx context.Context, _ events.UserUnblockObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventUserUnblock}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventUserUnblock,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventUserUnblock,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventUserUnblock,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerPollVoteNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.PollVoteNew(func(ctx context.Context, _ events.PollVoteNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventPollVoteNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventPollVoteNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventPollVoteNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventPollVoteNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupOfficersEdit(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupOfficersEdit(func(ctx context.Context, _ events.GroupOfficersEditObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventGroupOfficersEdit}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventGroupOfficersEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventGroupOfficersEdit,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventGroupOfficersEdit,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupChangeSettings(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupChangeSettings(func(ctx context.Context, _ events.GroupChangeSettingsObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventGroupChangeSettings}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventGroupChangeSettings,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventGroupChangeSettings,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventGroupChangeSettings,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerGroupChangePhoto(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.GroupChangePhoto(func(ctx context.Context, _ events.GroupChangePhotoObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventGroupChangePhoto}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventGroupChangePhoto,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventGroupChangePhoto,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventGroupChangePhoto,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerVkpayTransaction(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.VkpayTransaction(func(ctx context.Context, _ events.VkpayTransactionObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventVkpayTransaction}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventVkpayTransaction,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventVkpayTransaction,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventVkpayTransaction,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLeadFormsNew(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.LeadFormsNew(func(ctx context.Context, _ events.LeadFormsNewObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventLeadFormsNew}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventLeadFormsNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventLeadFormsNew,
			Object:  []byte("{}"),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventLeadFormsNew,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerAppPayload(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.AppPayload(func(ctx context.Context, _ events.AppPayloadObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventAppPayload}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventAppPayload,
			Object:  []byte(`{"user_id":117253521,"app_id":6703670,"payload":"{\"foo\":\"bar\"}"}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventAppPayload,
			Object:  []byte(`{"user_id":117253521,"app_id":6703670,"payload":"{\"foo\":\"bar\"}"}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventAppPayload,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerMessageRead(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.MessageRead(func(ctx context.Context, _ events.MessageReadObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventMessageRead}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventMessageRead,
			Object:  []byte(`{"from_id":1,"peer_id":1,"read_message_id":1}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventMessageRead,
			Object:  []byte(`{"from_id":1,"peer_id":1,"read_message_id":1}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventMessageRead,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLikeAdd(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.LikeAdd(func(ctx context.Context, _ events.LikeAddObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventLikeAdd}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventLikeAdd,
			Object:  []byte(`{"liker_id": 574423462,"object_type": "photo","object_owner_id": -178044536,"object_id": 457242474,"thread_reply_id": 0}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventLikeAdd,
			Object:  []byte(`{"liker_id": 574423462,"object_type": "photo","object_owner_id": -178044536,"object_id": 457242474,"thread_reply_id": 0}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventLikeAdd,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_HandlerLikeRemove(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.LikeRemove(func(ctx context.Context, _ events.LikeRemoveObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventLikeRemove}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventLikeRemove,
			Object:  []byte(`{"liker_id": 574423462,"object_type": "photo","object_owner_id": -178044536,"object_id": 457242474,"thread_reply_id": 0}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventLikeRemove,
			Object:  []byte(`{"liker_id": 574423462,"object_type": "photo","object_owner_id": -178044536,"object_id": 457242474,"thread_reply_id": 0}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventLikeRemove,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutSubscriptionCreate(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutSubscriptionCreate(func(ctx context.Context, _ events.DonutSubscriptionCreateObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutSubscriptionCreate}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionCreate,
			Object:  []byte(`{"amount":50,"amount_without_fee":47.5,"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionCreate,
			Object:  []byte(`{"amount":50,"amount_without_fee":47.5,"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutSubscriptionCreate,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutSubscriptionProlonged(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutSubscriptionProlonged(func(ctx context.Context, _ events.DonutSubscriptionProlongedObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutSubscriptionProlonged}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionProlonged,
			Object:  []byte(`{"amount":50,"amount_without_fee":47.5,"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionProlonged,
			Object:  []byte(`{"amount":50,"amount_without_fee":47.5,"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutSubscriptionProlonged,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutSubscriptionExpired(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutSubscriptionExpired(func(ctx context.Context, _ events.DonutSubscriptionExpiredObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutSubscriptionExpired}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionExpired,
			Object:  []byte(`{"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionExpired,
			Object:  []byte(`{"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutSubscriptionExpired,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutSubscriptionCancelled(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutSubscriptionCancelled(func(ctx context.Context, _ events.DonutSubscriptionCancelledObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutSubscriptionCancelled}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionCancelled,
			Object:  []byte(`{"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionCancelled,
			Object:  []byte(`{"user_id":117253521}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutSubscriptionCancelled,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutSubscriptionPriceChanged(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutSubscriptionPriceChanged(func(ctx context.Context, _ events.DonutSubscriptionPriceChangedObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutSubscriptionPriceChanged}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionPriceChanged,
			Object:  []byte(`{"amount_old":50,"amount_new":69}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutSubscriptionPriceChanged,
			Object:  []byte(`{"amount_old":50,"amount_new":69}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutSubscriptionPriceChanged,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutMoneyWithdraw(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutMoneyWithdraw(func(ctx context.Context, _ events.DonutMoneyWithdrawObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutMoneyWithdraw}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutMoneyWithdraw,
			Object:  []byte(`{"amount":50,"amount_without_fee":47.5}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutMoneyWithdraw,
			Object:  []byte(`{"amount":50,"amount_without_fee":47.5}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutMoneyWithdraw,
			Object: []byte(""),
		},
		true,
	)
}

func TestFuncList_DonutMoneyWithdrawError(t *testing.T) {
	t.Parallel()

	fl := events.NewFuncList()

	fl.DonutMoneyWithdrawError(func(ctx context.Context, _ events.DonutMoneyWithdrawErrorObject) {
		groupID := events.GroupIDFromContext(ctx)
		assert.Equal(t, GID, groupID)
	})
	assert.Equal(t, []events.EventType{events.EventDonutMoneyWithdrawError}, fl.ListEvents())

	f := func(e events.GroupEvent, wantErr bool) {
		if err := fl.Handler(context.Background(), e); (err != nil) != wantErr {
			t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, wantErr)
		}
	}

	f(
		events.GroupEvent{
			Type:    events.EventDonutMoneyWithdrawError,
			Object:  []byte(`{"reason":"test"}`),
			GroupID: GID,
		},
		false,
	)
	fl.Goroutine(true)
	f(
		events.GroupEvent{
			Type:    events.EventDonutMoneyWithdrawError,
			Object:  []byte(`{"reason":"test"}`),
			GroupID: GID,
		},
		false,
	)
	f(
		events.GroupEvent{
			Type:   events.EventDonutMoneyWithdrawError,
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
	assert.Equal(t, []events.EventType{"wtf_event"}, fl.ListEvents())

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
	fl.Goroutine(true)
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
