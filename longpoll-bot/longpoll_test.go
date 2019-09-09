package longpoll

import (
	"os"
	"strconv"
	"testing"

	"github.com/SevereCloud/vksdk/api"
	"github.com/SevereCloud/vksdk/object"
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

func TestLongpoll_Handler(t *testing.T) { // nolint:gocyclo
	lp := &Longpoll{}
	t.Run("MessageNew", func(t *testing.T) {
		lp.MessageNew(func(obj object.MessageNewObject, groupID int) {})
		if len(lp.funcList.MessageNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageReply", func(t *testing.T) {
		lp.MessageReply(func(obj object.MessageReplyObject, groupID int) {})
		if len(lp.funcList.MessageReply) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageEdit", func(t *testing.T) {
		lp.MessageEdit(func(obj object.MessageEditObject, groupID int) {})
		if len(lp.funcList.MessageEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageAllow", func(t *testing.T) {
		lp.MessageAllow(func(obj object.MessageAllowObject, groupID int) {})
		if len(lp.funcList.MessageAllow) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageDeny", func(t *testing.T) {
		lp.MessageDeny(func(obj object.MessageDenyObject, groupID int) {})
		if len(lp.funcList.MessageDeny) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MessageTypingState", func(t *testing.T) {
		lp.MessageTypingState(func(obj object.MessageTypingStateObject, groupID int) {})
		if len(lp.funcList.MessageTypingState) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoNew", func(t *testing.T) {
		lp.PhotoNew(func(obj object.PhotoNewObject, groupID int) {})
		if len(lp.funcList.PhotoNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentNew", func(t *testing.T) {
		lp.PhotoCommentNew(func(obj object.PhotoCommentNewObject, groupID int) {})
		if len(lp.funcList.PhotoCommentNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentEdit", func(t *testing.T) {
		lp.PhotoCommentEdit(func(obj object.PhotoCommentEditObject, groupID int) {})
		if len(lp.funcList.PhotoCommentEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentRestore", func(t *testing.T) {
		lp.PhotoCommentRestore(func(obj object.PhotoCommentRestoreObject, groupID int) {})
		if len(lp.funcList.PhotoCommentRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PhotoCommentDelete", func(t *testing.T) {
		lp.PhotoCommentDelete(func(obj object.PhotoCommentDeleteObject, groupID int) {})
		if len(lp.funcList.PhotoCommentDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("AudioNew", func(t *testing.T) {
		lp.AudioNew(func(obj object.AudioNewObject, groupID int) {})
		if len(lp.funcList.AudioNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoNew", func(t *testing.T) {
		lp.VideoNew(func(obj object.VideoNewObject, groupID int) {})
		if len(lp.funcList.VideoNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentNew", func(t *testing.T) {
		lp.VideoCommentNew(func(obj object.VideoCommentNewObject, groupID int) {})
		if len(lp.funcList.VideoCommentNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentEdit", func(t *testing.T) {
		lp.VideoCommentEdit(func(obj object.VideoCommentEditObject, groupID int) {})
		if len(lp.funcList.VideoCommentEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentRestore", func(t *testing.T) {
		lp.VideoCommentRestore(func(obj object.VideoCommentRestoreObject, groupID int) {})
		if len(lp.funcList.VideoCommentRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VideoCommentDelete", func(t *testing.T) {
		lp.VideoCommentDelete(func(obj object.VideoCommentDeleteObject, groupID int) {})
		if len(lp.funcList.VideoCommentDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallPostNew", func(t *testing.T) {
		lp.WallPostNew(func(obj object.WallPostNewObject, groupID int) {})
		if len(lp.funcList.WallPostNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallRepost", func(t *testing.T) {
		lp.WallRepost(func(obj object.WallRepostObject, groupID int) {})
		if len(lp.funcList.WallRepost) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyNew", func(t *testing.T) {
		lp.WallReplyNew(func(obj object.WallReplyNewObject, groupID int) {})
		if len(lp.funcList.WallReplyNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyEdit", func(t *testing.T) {
		lp.WallReplyEdit(func(obj object.WallReplyEditObject, groupID int) {})
		if len(lp.funcList.WallReplyEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyRestore", func(t *testing.T) {
		lp.WallReplyRestore(func(obj object.WallReplyRestoreObject, groupID int) {})
		if len(lp.funcList.WallReplyRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("WallReplyDelete", func(t *testing.T) {
		lp.WallReplyDelete(func(obj object.WallReplyDeleteObject, groupID int) {})
		if len(lp.funcList.WallReplyDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostNew", func(t *testing.T) {
		lp.BoardPostNew(func(obj object.BoardPostNewObject, groupID int) {})
		if len(lp.funcList.BoardPostNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostEdit", func(t *testing.T) {
		lp.BoardPostEdit(func(obj object.BoardPostEditObject, groupID int) {})
		if len(lp.funcList.BoardPostEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostRestore", func(t *testing.T) {
		lp.BoardPostRestore(func(obj object.BoardPostRestoreObject, groupID int) {})
		if len(lp.funcList.BoardPostRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("BoardPostDelete", func(t *testing.T) {
		lp.BoardPostDelete(func(obj object.BoardPostDeleteObject, groupID int) {})
		if len(lp.funcList.BoardPostDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentNew", func(t *testing.T) {
		lp.MarketCommentNew(func(obj object.MarketCommentNewObject, groupID int) {})
		if len(lp.funcList.MarketCommentNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentEdit", func(t *testing.T) {
		lp.MarketCommentEdit(func(obj object.MarketCommentEditObject, groupID int) {})
		if len(lp.funcList.MarketCommentEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentRestore", func(t *testing.T) {
		lp.MarketCommentRestore(func(obj object.MarketCommentRestoreObject, groupID int) {})
		if len(lp.funcList.MarketCommentRestore) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("MarketCommentDelete", func(t *testing.T) {
		lp.MarketCommentDelete(func(obj object.MarketCommentDeleteObject, groupID int) {})
		if len(lp.funcList.MarketCommentDelete) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupLeave", func(t *testing.T) {
		lp.GroupLeave(func(obj object.GroupLeaveObject, groupID int) {})
		if len(lp.funcList.GroupLeave) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupJoin", func(t *testing.T) {
		lp.GroupJoin(func(obj object.GroupJoinObject, groupID int) {})
		if len(lp.funcList.GroupJoin) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("UserBlock", func(t *testing.T) {
		lp.UserBlock(func(obj object.UserBlockObject, groupID int) {})
		if len(lp.funcList.UserBlock) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("UserUnblock", func(t *testing.T) {
		lp.UserUnblock(func(obj object.UserUnblockObject, groupID int) {})
		if len(lp.funcList.UserUnblock) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("PollVoteNew", func(t *testing.T) {
		lp.PollVoteNew(func(obj object.PollVoteNewObject, groupID int) {})
		if len(lp.funcList.PollVoteNew) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupOfficersEdit", func(t *testing.T) {
		lp.GroupOfficersEdit(func(obj object.GroupOfficersEditObject, groupID int) {})
		if len(lp.funcList.GroupOfficersEdit) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupChangeSettings", func(t *testing.T) {
		lp.GroupChangeSettings(func(obj object.GroupChangeSettingsObject, groupID int) {})
		if len(lp.funcList.GroupChangeSettings) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("GroupChangePhoto", func(t *testing.T) {
		lp.GroupChangePhoto(func(obj object.GroupChangePhotoObject, groupID int) {})
		if len(lp.funcList.GroupChangePhoto) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("VkpayTransaction", func(t *testing.T) {
		lp.VkpayTransaction(func(obj object.VkpayTransactionObject, groupID int) {})
		if len(lp.funcList.VkpayTransaction) != 1 {
			t.Error("Want len = 1")
		}
	})
	t.Run("LeadFormsNew", func(t *testing.T) {
		lp.LeadFormsNew(func(obj object.LeadFormsNewObject, groupID int) {})
		if len(lp.funcList.LeadFormsNew) != 1 {
			t.Error("Want len = 1")
		}
	})

	t.Run("FullResponse", func(t *testing.T) {
		lp.FullResponse(func(resp object.LongpollBotResponse) {})
		if len(lp.funcFullResponseList) != 1 {
			t.Error("Want len = 1")
		}
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
				vk:      badVk,
				groupID: 0,
			},
			wantErr: true,
		},
		{
			name: "Init good",
			args: args{
				vk:      vk,
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
	lp, _ := Init(vk, groupID)

	tests := []struct {
		name        string
		argResponse object.LongpollBotResponse
		wantErr     bool
	}{
		{
			name: "ok",
		},
		{
			name:        "failed: 1",
			argResponse: object.LongpollBotResponse{Failed: 1},
		},
		{
			name:        "failed: 2",
			argResponse: object.LongpollBotResponse{Failed: 2},
		},
		{
			name:        "failed: 3",
			argResponse: object.LongpollBotResponse{Failed: 3},
		},
		{
			name:        "failed: 4",
			argResponse: object.LongpollBotResponse{Failed: 4},
			wantErr:     true,
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

// TODO: write TestLongpoll_Run
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

func TestLongpoll_RunError(t *testing.T) {
	vk := api.Init("")
	lp, _ := Init(vk, 0)
	lp.Wait = 1

	t.Run("Run client error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})
	lp.Server = "http://example.com"
	t.Run("Run json error", func(t *testing.T) {
		if err := lp.Run(); err == nil {
			t.Error(err)
		}
	})
}
