package longpoll

import (
	"os"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/5.92/api"
	"github.com/SevereCloud/vksdk/5.92/object"
)

func TestLongpoll_Shutdown(t *testing.T) {
	t.Run("Shutdown", func(t *testing.T) {
		lp := &Longpoll{}
		lp.Shutdown()
		if lp.inShutdown != 1 {
			t.Error("inShutdown != 1")
		}
	})

}

func TestLongpoll_MessageNew(t *testing.T) {
	t.Run("MessageNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MessageNew(func(obj object.MessageNewObject, groupID int) {})
	})

}

func TestLongpoll_MessageReply(t *testing.T) {
	t.Run("MessageReply", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MessageReply(func(obj object.MessageReplyObject, groupID int) {})
	})

}

func TestLongpoll_MessageEdit(t *testing.T) {
	t.Run("MessageEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MessageEdit(func(obj object.MessageEditObject, groupID int) {})
	})

}

func TestLongpoll_MessageAllow(t *testing.T) {
	t.Run("MessageAllow", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MessageAllow(func(obj object.MessageAllowObject, groupID int) {})
	})

}

func TestLongpoll_MessageDeny(t *testing.T) {
	t.Run("MessageDeny", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MessageDeny(func(obj object.MessageDenyObject, groupID int) {})
	})

}

func TestLongpoll_MessageTypingState(t *testing.T) {
	t.Run("MessageTypingState", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MessageTypingState(func(obj object.MessageTypingStateObject, groupID int) {})
	})

}

func TestLongpoll_PhotoNew(t *testing.T) {
	t.Run("PhotoNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.PhotoNew(func(obj object.PhotoNewObject, groupID int) {})
	})

}

func TestLongpoll_PhotoCommentNew(t *testing.T) {
	t.Run("PhotoCommentNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.PhotoCommentNew(func(obj object.PhotoCommentNewObject, groupID int) {})
	})

}

func TestLongpoll_PhotoCommentEdit(t *testing.T) {
	t.Run("PhotoCommentEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.PhotoCommentEdit(func(obj object.PhotoCommentEditObject, groupID int) {})
	})

}

func TestLongpoll_PhotoCommentRestore(t *testing.T) {
	t.Run("PhotoCommentRestore", func(t *testing.T) {
		lp := &Longpoll{}
		lp.PhotoCommentRestore(func(obj object.PhotoCommentRestoreObject, groupID int) {})
	})

}

func TestLongpoll_PhotoCommentDelete(t *testing.T) {
	t.Run("PhotoCommentDelete", func(t *testing.T) {
		lp := &Longpoll{}
		lp.PhotoCommentDelete(func(obj object.PhotoCommentDeleteObject, groupID int) {})
	})

}

func TestLongpoll_AudioNew(t *testing.T) {
	t.Run("AudioNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.AudioNew(func(obj object.AudioNewObject, groupID int) {})
	})

}

func TestLongpoll_VideoNew(t *testing.T) {
	t.Run("VideoNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.VideoNew(func(obj object.VideoNewObject, groupID int) {})
	})

}

func TestLongpoll_VideoCommentNew(t *testing.T) {
	t.Run("VideoCommentNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.VideoCommentNew(func(obj object.VideoCommentNewObject, groupID int) {})
	})

}

func TestLongpoll_VideoCommentEdit(t *testing.T) {
	t.Run("VideoCommentEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.VideoCommentEdit(func(obj object.VideoCommentEditObject, groupID int) {})
	})

}

func TestLongpoll_VideoCommentRestore(t *testing.T) {
	t.Run("VideoCommentRestore", func(t *testing.T) {
		lp := &Longpoll{}
		lp.VideoCommentRestore(func(obj object.VideoCommentRestoreObject, groupID int) {})
	})

}

func TestLongpoll_VideoCommentDelete(t *testing.T) {
	t.Run("VideoCommentDelete", func(t *testing.T) {
		lp := &Longpoll{}
		lp.VideoCommentDelete(func(obj object.VideoCommentDeleteObject, groupID int) {})
	})

}

func TestLongpoll_WallPostNew(t *testing.T) {
	t.Run("WallPostNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.WallPostNew(func(obj object.WallPostNewObject, groupID int) {})
	})

}

func TestLongpoll_WallRepost(t *testing.T) {
	t.Run("WallRepost", func(t *testing.T) {
		lp := &Longpoll{}
		lp.WallRepost(func(obj object.WallRepostObject, groupID int) {})
	})

}

func TestLongpoll_WallReplyNew(t *testing.T) {
	t.Run("WallReplyNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.WallReplyNew(func(obj object.WallReplyNewObject, groupID int) {})
	})

}

func TestLongpoll_WallReplyEdit(t *testing.T) {
	t.Run("WallReplyEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.WallReplyEdit(func(obj object.WallReplyEditObject, groupID int) {})
	})

}

func TestLongpoll_WallReplyRestore(t *testing.T) {
	t.Run("WallReplyRestore", func(t *testing.T) {
		lp := &Longpoll{}
		lp.WallReplyRestore(func(obj object.WallReplyRestoreObject, groupID int) {})
	})

}

func TestLongpoll_WallReplyDelete(t *testing.T) {
	t.Run("WallReplyDelete", func(t *testing.T) {
		lp := &Longpoll{}
		lp.WallReplyDelete(func(obj object.WallReplyDeleteObject, groupID int) {})
	})

}

func TestLongpoll_BoardPostNew(t *testing.T) {
	t.Run("BoardPostNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.BoardPostNew(func(obj object.BoardPostNewObject, groupID int) {})
	})

}

func TestLongpoll_BoardPostEdit(t *testing.T) {
	t.Run("BoardPostEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.BoardPostEdit(func(obj object.BoardPostEditObject, groupID int) {})
	})

}

func TestLongpoll_BoardPostRestore(t *testing.T) {
	t.Run("BoardPostRestore", func(t *testing.T) {
		lp := &Longpoll{}
		lp.BoardPostRestore(func(obj object.BoardPostRestoreObject, groupID int) {})
	})

}

func TestLongpoll_BoardPostDelete(t *testing.T) {
	t.Run("BoardPostDelete", func(t *testing.T) {
		lp := &Longpoll{}
		lp.BoardPostDelete(func(obj object.BoardPostDeleteObject, groupID int) {})
	})

}

func TestLongpoll_MarketCommentNew(t *testing.T) {
	t.Run("MarketCommentNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MarketCommentNew(func(obj object.MarketCommentNewObject, groupID int) {})
	})

}

func TestLongpoll_MarketCommentEdit(t *testing.T) {
	t.Run("MarketCommentEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MarketCommentEdit(func(obj object.MarketCommentEditObject, groupID int) {})
	})

}

func TestLongpoll_MarketCommentRestore(t *testing.T) {
	t.Run("MarketCommentRestore", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MarketCommentRestore(func(obj object.MarketCommentRestoreObject, groupID int) {})
	})

}

func TestLongpoll_MarketCommentDelete(t *testing.T) {
	t.Run("MarketCommentDelete", func(t *testing.T) {
		lp := &Longpoll{}
		lp.MarketCommentDelete(func(obj object.MarketCommentDeleteObject, groupID int) {})
	})

}

func TestLongpoll_GroupLeave(t *testing.T) {
	t.Run("GroupLeave", func(t *testing.T) {
		lp := &Longpoll{}
		lp.GroupLeave(func(obj object.GroupLeaveObject, groupID int) {})
	})

}

func TestLongpoll_GroupJoin(t *testing.T) {
	t.Run("GroupJoin", func(t *testing.T) {
		lp := &Longpoll{}
		lp.GroupJoin(func(obj object.GroupJoinObject, groupID int) {})
	})

}

func TestLongpoll_UserBlock(t *testing.T) {
	t.Run("UserBlock", func(t *testing.T) {
		lp := &Longpoll{}
		lp.UserBlock(func(obj object.UserBlockObject, groupID int) {})
	})

}

func TestLongpoll_UserUnblock(t *testing.T) {
	t.Run("UserUnblock", func(t *testing.T) {
		lp := &Longpoll{}
		lp.UserUnblock(func(obj object.UserUnblockObject, groupID int) {})
	})

}

func TestLongpoll_PollVoteNew(t *testing.T) {
	t.Run("PollVoteNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.PollVoteNew(func(obj object.PollVoteNewObject, groupID int) {})
	})

}

func TestLongpoll_GroupOfficersEdit(t *testing.T) {
	t.Run("GroupOfficersEdit", func(t *testing.T) {
		lp := &Longpoll{}
		lp.GroupOfficersEdit(func(obj object.GroupOfficersEditObject, groupID int) {})
	})

}

func TestLongpoll_GroupChangeSettings(t *testing.T) {
	t.Run("GroupChangeSettings", func(t *testing.T) {
		lp := &Longpoll{}
		lp.GroupChangeSettings(func(obj object.GroupChangeSettingsObject, groupID int) {})
	})

}

func TestLongpoll_GroupChangePhoto(t *testing.T) {
	t.Run("GroupChangePhoto", func(t *testing.T) {
		lp := &Longpoll{}
		lp.GroupChangePhoto(func(obj object.GroupChangePhotoObject, groupID int) {})
	})

}

func TestLongpoll_VkpayTransaction(t *testing.T) {
	t.Run("VkpayTransaction", func(t *testing.T) {
		lp := &Longpoll{}
		lp.VkpayTransaction(func(obj object.VkpayTransactionObject, groupID int) {})
	})

}

func TestLongpoll_LeadFormsNew(t *testing.T) {
	t.Run("LeadFormsNew", func(t *testing.T) {
		lp := &Longpoll{}
		lp.LeadFormsNew(func(obj object.LeadFormsNewObject, groupID int) {})
	})

}

func TestInit(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := api.Init(groupToken)
	badVk := api.Init("")
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))

	type args struct {
		vk      *api.VK
		groupID int
	}
	tests := []struct {
		name string
		args args
		// wantLp  Longpoll
		wantErr bool
	}{
		{
			name: "Init error",
			args: args{
				vk:      &badVk,
				groupID: 0,
			},
			wantErr: true,
		},
		{
			name: "Init good",
			args: args{
				vk:      &vk,
				groupID: groupID,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Init(tt.args.vk, tt.args.groupID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(gotLp, tt.wantLp) {
			// 	t.Errorf("Init() = %v, want %v", gotLp, tt.wantLp)
			// }
		})
	}
}

func TestLongpoll_checkResponse(t *testing.T) {
	groupToken := os.Getenv("GROUP_TOKEN")
	if groupToken == "" {
		t.Skip("GROUP_TOKEN empty")
	}
	vk := api.Init(groupToken)
	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
	lp, _ := Init(&vk, groupID)

	tests := []struct {
		name        string
		argResponse longpollResponse
		wantErr     bool
	}{
		{
			name: "ok",
		},
		{
			name:        "failed: 1",
			argResponse: longpollResponse{Failed: 1},
		},
		{
			name:        "failed: 2",
			argResponse: longpollResponse{Failed: 2},
		},
		{
			name:        "failed: 3",
			argResponse: longpollResponse{Failed: 3},
		},
		{
			name:        "failed: 4",
			argResponse: longpollResponse{Failed: 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := lp.checkResponse(tt.argResponse); (err != nil) != tt.wantErr {
				t.Errorf("Longpoll.checkResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TODO: write
// func TestLongpoll_Run(t *testing.T) {
// 	groupToken := os.Getenv("GROUP_TOKEN")
// 	if groupToken == "" {
// 		t.Skip("GROUP_TOKEN empty")
// 	}
// 	vk := api.Init(groupToken)
// 	groupID, _ := strconv.Atoi(os.Getenv("GROUP_ID"))
// 	lp, _ := Init(&vk, groupID)
// 	lp.Wait = 1

// 	t.Run("Run", func(t *testing.T) {
//
// 	})
// }
