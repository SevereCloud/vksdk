package handler

import (
	"testing"

	"github.com/SevereCloud/vksdk/5.92/object"
)

const GID = 123456

func TestFuncList_HandlerMessageNew(t *testing.T) {
	funcList := FuncList{
		MessageNew: []object.MessageNewFunc{
			func(obj object.MessageNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "message_new",
			argE: object.GroupEvent{
				Type:    "message_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_new error",
			argE: object.GroupEvent{
				Type:   "message_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMessageReply(t *testing.T) {
	funcList := FuncList{
		MessageReply: []object.MessageReplyFunc{
			func(obj object.MessageReplyObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageReplyFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "message_reply",
			argE: object.GroupEvent{
				Type:    "message_reply",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_reply error",
			argE: object.GroupEvent{
				Type:   "message_reply",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMessageEdit(t *testing.T) {
	funcList := FuncList{
		MessageEdit: []object.MessageEditFunc{
			func(obj object.MessageEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "message_edit",
			argE: object.GroupEvent{
				Type:    "message_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_edit error",
			argE: object.GroupEvent{
				Type:   "message_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMessageAllow(t *testing.T) {
	funcList := FuncList{
		MessageAllow: []object.MessageAllowFunc{
			func(obj object.MessageAllowObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageAllowFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "message_allow",
			argE: object.GroupEvent{
				Type:    "message_allow",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_allow error",
			argE: object.GroupEvent{
				Type:   "message_allow",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMessageDeny(t *testing.T) {
	funcList := FuncList{
		MessageDeny: []object.MessageDenyFunc{
			func(obj object.MessageDenyObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageDenyFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "messages_deny",
			argE: object.GroupEvent{
				Type:    "messages_deny",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "messages_deny error",
			argE: object.GroupEvent{
				Type:   "messages_deny",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMessageTypingState(t *testing.T) {
	funcList := FuncList{
		MessageTypingState: []object.MessageTypingStateFunc{
			func(obj object.MessageTypingStateObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MessageTypingStateFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "message_typing_state",
			argE: object.GroupEvent{
				Type:    "message_typing_state",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "message_typing_state error",
			argE: object.GroupEvent{
				Type:   "message_typing_state",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerPhotoNew(t *testing.T) {
	funcList := FuncList{
		PhotoNew: []object.PhotoNewFunc{
			func(obj object.PhotoNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "photo_new",
			argE: object.GroupEvent{
				Type:    "photo_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_new error",
			argE: object.GroupEvent{
				Type:   "photo_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerPhotoCommentNew(t *testing.T) {
	funcList := FuncList{
		PhotoCommentNew: []object.PhotoCommentNewFunc{
			func(obj object.PhotoCommentNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "photo_comment_new",
			argE: object.GroupEvent{
				Type:    "photo_comment_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_new error",
			argE: object.GroupEvent{
				Type:   "photo_comment_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerPhotoCommentEdit(t *testing.T) {
	funcList := FuncList{
		PhotoCommentEdit: []object.PhotoCommentEditFunc{
			func(obj object.PhotoCommentEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "photo_comment_edit",
			argE: object.GroupEvent{
				Type:    "photo_comment_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_edit error",
			argE: object.GroupEvent{
				Type:   "photo_comment_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerPhotoCommentRestore(t *testing.T) {
	funcList := FuncList{
		PhotoCommentRestore: []object.PhotoCommentRestoreFunc{
			func(obj object.PhotoCommentRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentRestoreFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "photo_comment_restore",
			argE: object.GroupEvent{
				Type:    "photo_comment_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_restore error",
			argE: object.GroupEvent{
				Type:   "photo_comment_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerPhotoCommentDelete(t *testing.T) {
	funcList := FuncList{
		PhotoCommentDelete: []object.PhotoCommentDeleteFunc{
			func(obj object.PhotoCommentDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PhotoCommentDeleteFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "photo_comment_delete",
			argE: object.GroupEvent{
				Type:    "photo_comment_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "photo_comment_delete error",
			argE: object.GroupEvent{
				Type:   "photo_comment_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerAudioNew(t *testing.T) {
	funcList := FuncList{
		AudioNew: []object.AudioNewFunc{
			func(obj object.AudioNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() AudioNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "audio_new",
			argE: object.GroupEvent{
				Type:    "audio_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "audio_new error",
			argE: object.GroupEvent{
				Type:   "audio_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerVideoNew(t *testing.T) {
	funcList := FuncList{
		VideoNew: []object.VideoNewFunc{
			func(obj object.VideoNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "video_new",
			argE: object.GroupEvent{
				Type:    "video_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_new error",
			argE: object.GroupEvent{
				Type:   "video_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerVideoCommentNew(t *testing.T) {
	funcList := FuncList{
		VideoCommentNew: []object.VideoCommentNewFunc{
			func(obj object.VideoCommentNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "video_comment_new",
			argE: object.GroupEvent{
				Type:    "video_comment_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_new error",
			argE: object.GroupEvent{
				Type:   "video_comment_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerVideoCommentEdit(t *testing.T) {
	funcList := FuncList{
		VideoCommentEdit: []object.VideoCommentEditFunc{
			func(obj object.VideoCommentEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "video_comment_edit",
			argE: object.GroupEvent{
				Type:    "video_comment_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_edit error",
			argE: object.GroupEvent{
				Type:   "video_comment_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerVideoCommentRestore(t *testing.T) {
	funcList := FuncList{
		VideoCommentRestore: []object.VideoCommentRestoreFunc{
			func(obj object.VideoCommentRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentRestoreFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "video_comment_restore",
			argE: object.GroupEvent{
				Type:    "video_comment_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_restore error",
			argE: object.GroupEvent{
				Type:   "video_comment_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerVideoCommentDelete(t *testing.T) {
	funcList := FuncList{
		VideoCommentDelete: []object.VideoCommentDeleteFunc{
			func(obj object.VideoCommentDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VideoCommentDeleteFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "video_comment_delete",
			argE: object.GroupEvent{
				Type:    "video_comment_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "video_comment_delete error",
			argE: object.GroupEvent{
				Type:   "video_comment_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerWallPostNew(t *testing.T) {
	funcList := FuncList{
		WallPostNew: []object.WallPostNewFunc{
			func(obj object.WallPostNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallPostNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "wall_post_new",
			argE: object.GroupEvent{
				Type:    "wall_post_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_post_new error",
			argE: object.GroupEvent{
				Type:   "wall_post_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerWallRepost(t *testing.T) {
	funcList := FuncList{
		WallRepost: []object.WallRepostFunc{
			func(obj object.WallRepostObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallRepostFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "wall_repost",
			argE: object.GroupEvent{
				Type:    "wall_repost",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_repost error",
			argE: object.GroupEvent{
				Type:   "wall_repost",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerWallReplyNew(t *testing.T) {
	funcList := FuncList{
		WallReplyNew: []object.WallReplyNewFunc{
			func(obj object.WallReplyNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "wall_reply_new",
			argE: object.GroupEvent{
				Type:    "wall_reply_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_new error",
			argE: object.GroupEvent{
				Type:   "wall_reply_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerWallReplyEdit(t *testing.T) {
	funcList := FuncList{
		WallReplyEdit: []object.WallReplyEditFunc{
			func(obj object.WallReplyEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "wall_reply_edit",
			argE: object.GroupEvent{
				Type:    "wall_reply_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_edit error",
			argE: object.GroupEvent{
				Type:   "wall_reply_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerWallReplyRestore(t *testing.T) {
	funcList := FuncList{
		WallReplyRestore: []object.WallReplyRestoreFunc{
			func(obj object.WallReplyRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyRestoreFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "wall_reply_restore",
			argE: object.GroupEvent{
				Type:    "wall_reply_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_restore error",
			argE: object.GroupEvent{
				Type:   "wall_reply_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerWallReplyDelete(t *testing.T) {
	funcList := FuncList{
		WallReplyDelete: []object.WallReplyDeleteFunc{
			func(obj object.WallReplyDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() WallReplyDeleteFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "wall_reply_delete",
			argE: object.GroupEvent{
				Type:    "wall_reply_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "wall_reply_delete error",
			argE: object.GroupEvent{
				Type:   "wall_reply_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerBoardPostNew(t *testing.T) {
	funcList := FuncList{
		BoardPostNew: []object.BoardPostNewFunc{
			func(obj object.BoardPostNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "board_post_new",
			argE: object.GroupEvent{
				Type:    "board_post_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_new error",
			argE: object.GroupEvent{
				Type:   "board_post_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerBoardPostEdit(t *testing.T) {
	funcList := FuncList{
		BoardPostEdit: []object.BoardPostEditFunc{
			func(obj object.BoardPostEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "board_post_edit",
			argE: object.GroupEvent{
				Type:    "board_post_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_edit error",
			argE: object.GroupEvent{
				Type:   "board_post_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerBoardPostRestore(t *testing.T) {
	funcList := FuncList{
		BoardPostRestore: []object.BoardPostRestoreFunc{
			func(obj object.BoardPostRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostRestoreFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "board_post_restore",
			argE: object.GroupEvent{
				Type:    "board_post_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_restore error",
			argE: object.GroupEvent{
				Type:   "board_post_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerBoardPostDelete(t *testing.T) {
	funcList := FuncList{
		BoardPostDelete: []object.BoardPostDeleteFunc{
			func(obj object.BoardPostDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() BoardPostDeleteFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "board_post_delete",
			argE: object.GroupEvent{
				Type:    "board_post_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "board_post_delete error",
			argE: object.GroupEvent{
				Type:   "board_post_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMarketCommentNew(t *testing.T) {
	funcList := FuncList{
		MarketCommentNew: []object.MarketCommentNewFunc{
			func(obj object.MarketCommentNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "market_comment_new",
			argE: object.GroupEvent{
				Type:    "market_comment_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_new error",
			argE: object.GroupEvent{
				Type:   "market_comment_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMarketCommentEdit(t *testing.T) {
	funcList := FuncList{
		MarketCommentEdit: []object.MarketCommentEditFunc{
			func(obj object.MarketCommentEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "market_comment_edit",
			argE: object.GroupEvent{
				Type:    "market_comment_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_edit error",
			argE: object.GroupEvent{
				Type:   "market_comment_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMarketCommentRestore(t *testing.T) {
	funcList := FuncList{
		MarketCommentRestore: []object.MarketCommentRestoreFunc{
			func(obj object.MarketCommentRestoreObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentRestoreFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "market_comment_restore",
			argE: object.GroupEvent{
				Type:    "market_comment_restore",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_restore error",
			argE: object.GroupEvent{
				Type:   "market_comment_restore",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerMarketCommentDelete(t *testing.T) {
	funcList := FuncList{
		MarketCommentDelete: []object.MarketCommentDeleteFunc{
			func(obj object.MarketCommentDeleteObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() MarketCommentDeleteFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "market_comment_delete",
			argE: object.GroupEvent{
				Type:    "market_comment_delete",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "market_comment_delete error",
			argE: object.GroupEvent{
				Type:   "market_comment_delete",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerGroupLeave(t *testing.T) {
	funcList := FuncList{
		GroupLeave: []object.GroupLeaveFunc{
			func(obj object.GroupLeaveObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupLeaveFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "group_leave",
			argE: object.GroupEvent{
				Type:    "group_leave",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_leave error",
			argE: object.GroupEvent{
				Type:   "group_leave",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerGroupJoin(t *testing.T) {
	funcList := FuncList{
		GroupJoin: []object.GroupJoinFunc{
			func(obj object.GroupJoinObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupJoinFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "group_join",
			argE: object.GroupEvent{
				Type:    "group_join",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_join error",
			argE: object.GroupEvent{
				Type:   "group_join",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerUserBlock(t *testing.T) {
	funcList := FuncList{
		UserBlock: []object.UserBlockFunc{
			func(obj object.UserBlockObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() UserBlockFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "user_block",
			argE: object.GroupEvent{
				Type:    "user_block",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "user_block error",
			argE: object.GroupEvent{
				Type:   "user_block",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerUserUnblock(t *testing.T) {
	funcList := FuncList{
		UserUnblock: []object.UserUnblockFunc{
			func(obj object.UserUnblockObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() UserUnblockFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "user_unblock",
			argE: object.GroupEvent{
				Type:    "user_unblock",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "user_unblock error",
			argE: object.GroupEvent{
				Type:   "user_unblock",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerPollVoteNew(t *testing.T) {
	funcList := FuncList{
		PollVoteNew: []object.PollVoteNewFunc{
			func(obj object.PollVoteNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() PollVoteNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "poll_vote_new",
			argE: object.GroupEvent{
				Type:    "poll_vote_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "poll_vote_new error",
			argE: object.GroupEvent{
				Type:   "poll_vote_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerGroupOfficersEdit(t *testing.T) {
	funcList := FuncList{
		GroupOfficersEdit: []object.GroupOfficersEditFunc{
			func(obj object.GroupOfficersEditObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupOfficersEditFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "group_officers_edit",
			argE: object.GroupEvent{
				Type:    "group_officers_edit",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_officers_edit error",
			argE: object.GroupEvent{
				Type:   "group_officers_edit",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerGroupChangeSettings(t *testing.T) {
	funcList := FuncList{
		GroupChangeSettings: []object.GroupChangeSettingsFunc{
			func(obj object.GroupChangeSettingsObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupChangeSettingsFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "group_change_settings",
			argE: object.GroupEvent{
				Type:    "group_change_settings",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_change_settings error",
			argE: object.GroupEvent{
				Type:   "group_change_settings",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerGroupChangePhoto(t *testing.T) {
	funcList := FuncList{
		GroupChangePhoto: []object.GroupChangePhotoFunc{
			func(obj object.GroupChangePhotoObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() GroupChangePhotoFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "group_change_photo",
			argE: object.GroupEvent{
				Type:    "group_change_photo",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "group_change_photo error",
			argE: object.GroupEvent{
				Type:   "group_change_photo",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerVkpayTransaction(t *testing.T) {
	funcList := FuncList{
		VkpayTransaction: []object.VkpayTransactionFunc{
			func(obj object.VkpayTransactionObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() VkpayTransactionFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "vkpay_transaction",
			argE: object.GroupEvent{
				Type:    "vkpay_transaction",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "vkpay_transaction error",
			argE: object.GroupEvent{
				Type:   "vkpay_transaction",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFuncList_HandlerLeadFormsNew(t *testing.T) {
	funcList := FuncList{
		LeadFormsNew: []object.LeadFormsNewFunc{
			func(obj object.LeadFormsNewObject, groupID int) {
				if groupID != GID {
					t.Errorf("FuncList.Handler() LeadFormsNewFunc groupID != GID")
				}
			},
		},
	}
	tests := []struct {
		name    string
		argE    object.GroupEvent
		wantErr bool
	}{
		{
			name: "lead_forms_new",
			argE: object.GroupEvent{
				Type:    "lead_forms_new",
				Object:  []byte("{}"),
				GroupID: GID,
			},
			wantErr: false,
		},
		{
			name: "lead_forms_new error",
			argE: object.GroupEvent{
				Type:   "lead_forms_new",
				Object: []byte(""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := funcList.Handler(tt.argE); (err != nil) != tt.wantErr {
				t.Errorf("FuncList.Handler() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
