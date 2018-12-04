package callback

// MessageNew handler
func (cb *Callback) MessageNew(f func(obj MessageNewObject, groupID int)) {
	// TODO MessageNew handler
}

// MessageReply handler
func (cb *Callback) MessageReply(f func(obj MessageReplyObject, groupID int)) {
	// TODO MessageReply handler
}

// MessageAllow handler
func (cb *Callback) MessageAllow(f func(obj MessagesAllowObject, groupID int)) {
	// TODO MessageAllow handler
}

// MessagesDeny handler
func (cb *Callback) MessagesDeny(f func(obj MessagesDenyObject, groupID int)) {
	// TODO MessagesDeny handler
}

// PhotoNew handler
func (cb *Callback) PhotoNew(f func(obj PhotoNewObject, groupID int)) {
	// TODO PhotoNew handler
}

// PhotoCommentNew handler
func (cb *Callback) PhotoCommentNew(f func(obj PhotoCommentNewObject, groupID int)) {
	// TODO PhotoCommentNew handler
}

// PhotoCommentEdit handler
func (cb *Callback) PhotoCommentEdit(f func(obj PhotoCommentEditObject, groupID int)) {
	// TODO PhotoCommentEdit handler
}

// PhotoCommentRestore handler
func (cb *Callback) PhotoCommentRestore(f func(obj PhotoCommentRestoreObject, groupID int)) {
	// TODO PhotoCommentRestore handler
}

// PhotoCommentDelete handler
func (cb *Callback) PhotoCommentDelete(f func(obj PhotoCommentDeleteObject, groupID int)) {
	// TODO PhotoCommentDelete handler
}

// AudioNew handler
func (cb *Callback) AudioNew(f func(obj AudioNewObject, groupID int)) {
	// TODO AudioNew handler
}

// VideoNew handler
func (cb *Callback) VideoNew(f func(obj VideoNewObject, groupID int)) {
	// TODO VideoNew handler
}

// VideoCommentNew handler
func (cb *Callback) VideoCommentNew(f func(obj VideoCommentNewObject, groupID int)) {
	// TODO VideoCommentNew handler
}

// VideoCommentEdit handler
func (cb *Callback) VideoCommentEdit(f func(obj VideoCommentEditObject, groupID int)) {
	// TODO VideoCommentEdit handler
}

// VideoCommentRestore handler
func (cb *Callback) VideoCommentRestore(f func(obj VideoCommentRestoreObject, groupID int)) {
	// TODO VideoCommentRestore handler
}

// VideoCommentDelete handler
func (cb *Callback) VideoCommentDelete(f func(obj VideoCommentDeleteObject, groupID int)) {
	// TODO VideoCommentDelete handler
}

// WallPostNew handler
func (cb *Callback) WallPostNew(f func(obj WallPostNewObject, groupID int)) {
	// TODO WallPostNew handler
}

// WallRepost handler
func (cb *Callback) WallRepost(f func(obj WallRepostObject, groupID int)) {
	// TODO WallRepost handler
}

// WallReplyNew handler
func (cb *Callback) WallReplyNew(f func(obj WallReplyNewObject, groupID int)) {
	// TODO WallReplyNew handler
}

// WallReplyEdit handler
func (cb *Callback) WallReplyEdit(f func(obj WallReplyEditObject, groupID int)) {
	// TODO WallReplyEdit handler
}

// WallReplyRestore handler
func (cb *Callback) WallReplyRestore(f func(obj WallReplyRestoreObject, groupID int)) {
	// TODO WallReplyRestore handler
}

// WallReplyDelete handler
func (cb *Callback) WallReplyDelete(f func(obj WallReplyDeleteObject, groupID int)) {
	// TODO WallReplyDelete handler
}

// BoardPostNew handler
func (cb *Callback) BoardPostNew(f func(obj BoardPostNewObject, groupID int)) {
	// TODO BoardPostNew handler
}

// BoardPostEdit handler
func (cb *Callback) BoardPostEdit(f func(obj BoardPostEditObject, groupID int)) {
	// TODO BoardPostEdit handler
}

// BoardPostRestore handler
func (cb *Callback) BoardPostRestore(f func(obj BoardPostRestoreObject, groupID int)) {
	// TODO BoardPostRestore handler
}

// BoardPostDelete handler
func (cb *Callback) BoardPostDelete(f func(obj BoardPostDeleteObject, groupID int)) {
	// TODO BoardPostDelete handler
}

// MarketCommentNew handler
func (cb *Callback) MarketCommentNew(f func(obj MarketCommentNewObject, groupID int)) {
	// TODO MarketCommentNew handler
}

// MarketCommentEdit handler
func (cb *Callback) MarketCommentEdit(f func(obj MarketCommentEditObject, groupID int)) {
	// TODO MarketCommentEdit handler
}

// MarketCommentRestore handler
func (cb *Callback) MarketCommentRestore(f func(obj MarketCommentRestoreObject, groupID int)) {
	// TODO MarketCommentRestore handler
}

// MarketCommentDelete handler
func (cb *Callback) MarketCommentDelete(f func(obj MarketCommentDeleteObject, groupID int)) {
	// TODO MarketCommentDelete handler
}

// GroupLeave handler
func (cb *Callback) GroupLeave(f func(obj GroupLeaveObject, groupID int)) {
	// TODO GroupLeave handler
}

// GroupJoin handler
func (cb *Callback) GroupJoin(f func(obj GroupJoinObject, groupID int)) {
	// TODO GroupJoin handler
}

// UserBlock handler
func (cb *Callback) UserBlock(f func(obj UserBlockObject, groupID int)) {
	// TODO UserBlock handler
}

// UserUnblock handler
func (cb *Callback) UserUnblock(f func(obj UserUnblockObject, groupID int)) {
	// TODO UserUnblock handler
}

// PollVoteNew handler
func (cb *Callback) PollVoteNew(f func(obj PollVoteNewObject, groupID int)) {
	// TODO PollVoteNew handler
}

// GroupOfficersEdit handler
func (cb *Callback) GroupOfficersEdit(f func(obj GroupOfficersEditObject, groupID int)) {
	// TODO GroupOfficersEdit handler
}

// GroupChangeSettings handler
func (cb *Callback) GroupChangeSettings(f func(obj GroupChangeSettingsObject, groupID int)) {
	// TODO GroupChangeSettings handler
}

// GroupChangePhoto handler
func (cb *Callback) GroupChangePhoto(f func(obj GroupChangePhotoObject, groupID int)) {
	// TODO GroupChangePhoto handler
}

// VkpayTransaction handler
func (cb *Callback) VkpayTransaction(f func(obj VkpayTransactionObject, groupID int)) {
	// TODO VkpayTransaction handler
}
