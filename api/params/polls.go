package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PollsAddVoteBuilder builder
//
// Adds the current user's vote to the selected answer in the poll.
//
// https://vk.com/dev/polls.addVote
type PollsAddVoteBuilder struct {
	api.Params
}

// NewPollsAddVoteBuilder func
func NewPollsAddVoteBuilder() *PollsAddVoteBuilder {
	return &PollsAddVoteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsAddVoteBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID Poll ID.
func (b *PollsAddVoteBuilder) PollID(v int) {
	b.Params["poll_id"] = v
}

// AnswerIDs parameter
func (b *PollsAddVoteBuilder) AnswerIDs(v []int) {
	b.Params["answer_ids"] = v
}

// IsBoard parameter
func (b *PollsAddVoteBuilder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// PollsCreateBuilder builder
//
// Creates polls that can be attached to the users' or communities' posts.
//
// https://vk.com/dev/polls.create
type PollsCreateBuilder struct {
	api.Params
}

// NewPollsCreateBuilder func
func NewPollsCreateBuilder() *PollsCreateBuilder {
	return &PollsCreateBuilder{api.Params{}}
}

// Question question text
func (b *PollsCreateBuilder) Question(v string) {
	b.Params["question"] = v
}

// IsAnonymous '1' – anonymous poll, participants list is hidden,, '0' – public poll, participants list is available,,
// Default value is '0'.
func (b *PollsCreateBuilder) IsAnonymous(v bool) {
	b.Params["is_anonymous"] = v
}

// IsMultiple parameter
func (b *PollsCreateBuilder) IsMultiple(v bool) {
	b.Params["is_multiple"] = v
}

// EndDate parameter
func (b *PollsCreateBuilder) EndDate(v int) {
	b.Params["end_date"] = v
}

// OwnerID If a poll will be added to a communty it is required to send a negative group identifier.
// Current user by default.
func (b *PollsCreateBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AddAnswers available answers list, for example: " ["yes","no","maybe"]", There can be from 1 to 10 answers.
func (b *PollsCreateBuilder) AddAnswers(v string) {
	b.Params["add_answers"] = v
}

// PhotoID parameter
func (b *PollsCreateBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// BackgroundID parameter
func (b *PollsCreateBuilder) BackgroundID(v string) {
	b.Params["background_id"] = v
}

// PollsDeleteVoteBuilder builder
//
// Deletes the current user's vote from the selected answer in the poll.
//
// https://vk.com/dev/polls.deleteVote
type PollsDeleteVoteBuilder struct {
	api.Params
}

// NewPollsDeleteVoteBuilder func
func NewPollsDeleteVoteBuilder() *PollsDeleteVoteBuilder {
	return &PollsDeleteVoteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsDeleteVoteBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID Poll ID.
func (b *PollsDeleteVoteBuilder) PollID(v int) {
	b.Params["poll_id"] = v
}

// AnswerID Answer ID.
func (b *PollsDeleteVoteBuilder) AnswerID(v int) {
	b.Params["answer_id"] = v
}

// IsBoard parameter
func (b *PollsDeleteVoteBuilder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// PollsEditBuilder builder
//
// Edits created polls
//
// https://vk.com/dev/polls.edit
type PollsEditBuilder struct {
	api.Params
}

// NewPollsEditBuilder func
func NewPollsEditBuilder() *PollsEditBuilder {
	return &PollsEditBuilder{api.Params{}}
}

// OwnerID poll owner id
func (b *PollsEditBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID edited poll's id
func (b *PollsEditBuilder) PollID(v int) {
	b.Params["poll_id"] = v
}

// Question new question text
func (b *PollsEditBuilder) Question(v string) {
	b.Params["question"] = v
}

// AddAnswers answers list, for example: , "["yes","no","maybe"]"
func (b *PollsEditBuilder) AddAnswers(v string) {
	b.Params["add_answers"] = v
}

// EditAnswers object containing answers that need to be edited,, key – answer id, value – new answer text.
// Example: {"382967099":"option1", "382967103":"option2"}"
func (b *PollsEditBuilder) EditAnswers(v string) {
	b.Params["edit_answers"] = v
}

// DeleteAnswers list of answer ids to be deleted. For example: "[382967099, 382967103]"
func (b *PollsEditBuilder) DeleteAnswers(v string) {
	b.Params["delete_answers"] = v
}

// EndDate parameter
func (b *PollsEditBuilder) EndDate(v int) {
	b.Params["end_date"] = v
}

// PhotoID parameter
func (b *PollsEditBuilder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// BackgroundID parameter
func (b *PollsEditBuilder) BackgroundID(v string) {
	b.Params["background_id"] = v
}

// PollsGetByIDBuilder builder
//
// Returns detailed information about a poll by its ID.
//
// https://vk.com/dev/polls.getById
type PollsGetByIDBuilder struct {
	api.Params
}

// NewPollsGetByIDBuilder func
func NewPollsGetByIDBuilder() *PollsGetByIDBuilder {
	return &PollsGetByIDBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsGetByIDBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// IsBoard '1' – poll is in a board, '0' – poll is on a wall. '0' by default.
func (b *PollsGetByIDBuilder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// PollID Poll ID.
func (b *PollsGetByIDBuilder) PollID(v int) {
	b.Params["poll_id"] = v
}

// Extended parameter
func (b *PollsGetByIDBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// FriendsCount parameter
func (b *PollsGetByIDBuilder) FriendsCount(v int) {
	b.Params["friends_count"] = v
}

// Fields parameter
func (b *PollsGetByIDBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase parameter
func (b *PollsGetByIDBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}

// PollsGetVotersBuilder builder
//
// Returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
type PollsGetVotersBuilder struct {
	api.Params
}

// NewPollsGetVotersBuilder func
func NewPollsGetVotersBuilder() *PollsGetVotersBuilder {
	return &PollsGetVotersBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsGetVotersBuilder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID Poll ID.
func (b *PollsGetVotersBuilder) PollID(v int) {
	b.Params["poll_id"] = v
}

// AnswerIDs Answer IDs.
func (b *PollsGetVotersBuilder) AnswerIDs(v []int) {
	b.Params["answer_ids"] = v
}

// IsBoard parameter
func (b *PollsGetVotersBuilder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// FriendsOnly '1' — to return only current user's friends, '0' — to return all users (default),
func (b *PollsGetVotersBuilder) FriendsOnly(v bool) {
	b.Params["friends_only"] = v
}

// Offset Offset needed to return a specific subset of voters. '0' — (default)
func (b *PollsGetVotersBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of user IDs to return (if the 'friends_only' parameter is not set, maximum '1000', otherwise '10').
// '100' — (default)
func (b *PollsGetVotersBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate (birthdate)', 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education',
// 'online', 'counters'.
func (b *PollsGetVotersBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive ,
// 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *PollsGetVotersBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}
