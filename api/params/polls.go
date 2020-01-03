package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// PollsAddVoteBulder builder
//
// Adds the current user's vote to the selected answer in the poll.
//
// https://vk.com/dev/polls.addVote
type PollsAddVoteBulder struct {
	api.Params
}

// NewPollsAddVoteBulder func
func NewPollsAddVoteBulder() *PollsAddVoteBulder {
	return &PollsAddVoteBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsAddVoteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID Poll ID.
func (b *PollsAddVoteBulder) PollID(v int) {
	b.Params["poll_id"] = v
}

// AnswerIDs parameter
func (b *PollsAddVoteBulder) AnswerIDs(v []int) {
	b.Params["answer_ids"] = v
}

// IsBoard parameter
func (b *PollsAddVoteBulder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// PollsCreateBulder builder
//
// Creates polls that can be attached to the users' or communities' posts.
//
// https://vk.com/dev/polls.create
type PollsCreateBulder struct {
	api.Params
}

// NewPollsCreateBulder func
func NewPollsCreateBulder() *PollsCreateBulder {
	return &PollsCreateBulder{api.Params{}}
}

// Question question text
func (b *PollsCreateBulder) Question(v string) {
	b.Params["question"] = v
}

// IsAnonymous '1' – anonymous poll, participants list is hidden,, '0' – public poll, participants list is available,, Default value is '0'.
func (b *PollsCreateBulder) IsAnonymous(v bool) {
	b.Params["is_anonymous"] = v
}

// IsMultiple parameter
func (b *PollsCreateBulder) IsMultiple(v bool) {
	b.Params["is_multiple"] = v
}

// EndDate parameter
func (b *PollsCreateBulder) EndDate(v int) {
	b.Params["end_date"] = v
}

// OwnerID If a poll will be added to a communty it is required to send a negative group identifier. Current user by default.
func (b *PollsCreateBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// AddAnswers available answers list, for example: " ["yes","no","maybe"]", There can be from 1 to 10 answers.
func (b *PollsCreateBulder) AddAnswers(v string) {
	b.Params["add_answers"] = v
}

// PhotoID parameter
func (b *PollsCreateBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// BackgroundID parameter
func (b *PollsCreateBulder) BackgroundID(v string) {
	b.Params["background_id"] = v
}

// PollsDeleteVoteBulder builder
//
// Deletes the current user's vote from the selected answer in the poll.
//
// https://vk.com/dev/polls.deleteVote
type PollsDeleteVoteBulder struct {
	api.Params
}

// NewPollsDeleteVoteBulder func
func NewPollsDeleteVoteBulder() *PollsDeleteVoteBulder {
	return &PollsDeleteVoteBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsDeleteVoteBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID Poll ID.
func (b *PollsDeleteVoteBulder) PollID(v int) {
	b.Params["poll_id"] = v
}

// AnswerID Answer ID.
func (b *PollsDeleteVoteBulder) AnswerID(v int) {
	b.Params["answer_id"] = v
}

// IsBoard parameter
func (b *PollsDeleteVoteBulder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// PollsEditBulder builder
//
// Edits created polls
//
// https://vk.com/dev/polls.edit
type PollsEditBulder struct {
	api.Params
}

// NewPollsEditBulder func
func NewPollsEditBulder() *PollsEditBulder {
	return &PollsEditBulder{api.Params{}}
}

// OwnerID poll owner id
func (b *PollsEditBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID edited poll's id
func (b *PollsEditBulder) PollID(v int) {
	b.Params["poll_id"] = v
}

// Question new question text
func (b *PollsEditBulder) Question(v string) {
	b.Params["question"] = v
}

// AddAnswers answers list, for example: , "["yes","no","maybe"]"
func (b *PollsEditBulder) AddAnswers(v string) {
	b.Params["add_answers"] = v
}

// EditAnswers object containing answers that need to be edited,, key – answer id, value – new answer text. Example: {"382967099":"option1", "382967103":"option2"}"
func (b *PollsEditBulder) EditAnswers(v string) {
	b.Params["edit_answers"] = v
}

// DeleteAnswers list of answer ids to be deleted. For example: "[382967099, 382967103]"
func (b *PollsEditBulder) DeleteAnswers(v string) {
	b.Params["delete_answers"] = v
}

// EndDate parameter
func (b *PollsEditBulder) EndDate(v int) {
	b.Params["end_date"] = v
}

// PhotoID parameter
func (b *PollsEditBulder) PhotoID(v int) {
	b.Params["photo_id"] = v
}

// BackgroundID parameter
func (b *PollsEditBulder) BackgroundID(v string) {
	b.Params["background_id"] = v
}

// PollsGetByIDBulder builder
//
// Returns detailed information about a poll by its ID.
//
// https://vk.com/dev/polls.getById
type PollsGetByIDBulder struct {
	api.Params
}

// NewPollsGetByIDBulder func
func NewPollsGetByIDBulder() *PollsGetByIDBulder {
	return &PollsGetByIDBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsGetByIDBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// IsBoard '1' – poll is in a board, '0' – poll is on a wall. '0' by default.
func (b *PollsGetByIDBulder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// PollID Poll ID.
func (b *PollsGetByIDBulder) PollID(v int) {
	b.Params["poll_id"] = v
}

// Extended parameter
func (b *PollsGetByIDBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// FriendsCount parameter
func (b *PollsGetByIDBulder) FriendsCount(v int) {
	b.Params["friends_count"] = v
}

// Fields parameter
func (b *PollsGetByIDBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase parameter
func (b *PollsGetByIDBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// PollsGetVotersBulder builder
//
// Returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
type PollsGetVotersBulder struct {
	api.Params
}

// NewPollsGetVotersBulder func
func NewPollsGetVotersBulder() *PollsGetVotersBulder {
	return &PollsGetVotersBulder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsGetVotersBulder) OwnerID(v int) {
	b.Params["owner_id"] = v
}

// PollID Poll ID.
func (b *PollsGetVotersBulder) PollID(v int) {
	b.Params["poll_id"] = v
}

// AnswerIDs Answer IDs.
func (b *PollsGetVotersBulder) AnswerIDs(v []int) {
	b.Params["answer_ids"] = v
}

// IsBoard parameter
func (b *PollsGetVotersBulder) IsBoard(v bool) {
	b.Params["is_board"] = v
}

// FriendsOnly '1' — to return only current user's friends, '0' — to return all users (default),
func (b *PollsGetVotersBulder) FriendsOnly(v bool) {
	b.Params["friends_only"] = v
}

// Offset Offset needed to return a specific subset of voters. '0' — (default)
func (b *PollsGetVotersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of user IDs to return (if the 'friends_only' parameter is not set, maximum '1000', otherwise '10'). '100' — (default)
func (b *PollsGetVotersBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate (birthdate)', 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online', 'counters'.
func (b *PollsGetVotersBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive , 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *PollsGetVotersBulder) NameCase(v string) {
	b.Params["name_case"] = v
}
