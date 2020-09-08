package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// PollsAddVoteBuilder builder.
//
// Adds the current user's vote to the selected answer in the poll.
//
// https://vk.com/dev/polls.addVote
type PollsAddVoteBuilder struct {
	api.Params
}

// NewPollsAddVoteBuilder func.
func NewPollsAddVoteBuilder() *PollsAddVoteBuilder {
	return &PollsAddVoteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsAddVoteBuilder) OwnerID(v int) *PollsAddVoteBuilder {
	b.Params["owner_id"] = v
	return b
}

// PollID parameter.
func (b *PollsAddVoteBuilder) PollID(v int) *PollsAddVoteBuilder {
	b.Params["poll_id"] = v
	return b
}

// AnswerIDs parameter.
func (b *PollsAddVoteBuilder) AnswerIDs(v []int) *PollsAddVoteBuilder {
	b.Params["answer_ids"] = v
	return b
}

// IsBoard parameter.
func (b *PollsAddVoteBuilder) IsBoard(v bool) *PollsAddVoteBuilder {
	b.Params["is_board"] = v
	return b
}

// PollsCreateBuilder builder.
//
// Creates polls that can be attached to the users' or communities' posts.
//
// https://vk.com/dev/polls.create
type PollsCreateBuilder struct {
	api.Params
}

// NewPollsCreateBuilder func.
func NewPollsCreateBuilder() *PollsCreateBuilder {
	return &PollsCreateBuilder{api.Params{}}
}

// Question question text.
func (b *PollsCreateBuilder) Question(v string) *PollsCreateBuilder {
	b.Params["question"] = v
	return b
}

// IsAnonymous '1' – anonymous poll, participants list is hidden,, '0' – public poll, participants list is available,,
// Default value is '0'.
func (b *PollsCreateBuilder) IsAnonymous(v bool) *PollsCreateBuilder {
	b.Params["is_anonymous"] = v
	return b
}

// IsMultiple parameter.
func (b *PollsCreateBuilder) IsMultiple(v bool) *PollsCreateBuilder {
	b.Params["is_multiple"] = v
	return b
}

// EndDate parameter.
func (b *PollsCreateBuilder) EndDate(v int) *PollsCreateBuilder {
	b.Params["end_date"] = v
	return b
}

// OwnerID if a poll will be added to a community it is required to send a negative group identifier.
// Current user by default.
func (b *PollsCreateBuilder) OwnerID(v int) *PollsCreateBuilder {
	b.Params["owner_id"] = v
	return b
}

// AddAnswers available answers list, for example: " ["yes","no","maybe"]", There can be from 1 to 10 answers.
func (b *PollsCreateBuilder) AddAnswers(v string) *PollsCreateBuilder {
	b.Params["add_answers"] = v
	return b
}

// PhotoID parameter.
func (b *PollsCreateBuilder) PhotoID(v int) *PollsCreateBuilder {
	b.Params["photo_id"] = v
	return b
}

// BackgroundID parameter.
func (b *PollsCreateBuilder) BackgroundID(v string) *PollsCreateBuilder {
	b.Params["background_id"] = v
	return b
}

// PollsDeleteVoteBuilder builder.
//
// Deletes the current user's vote from the selected answer in the poll.
//
// https://vk.com/dev/polls.deleteVote
type PollsDeleteVoteBuilder struct {
	api.Params
}

// NewPollsDeleteVoteBuilder func.
func NewPollsDeleteVoteBuilder() *PollsDeleteVoteBuilder {
	return &PollsDeleteVoteBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsDeleteVoteBuilder) OwnerID(v int) *PollsDeleteVoteBuilder {
	b.Params["owner_id"] = v
	return b
}

// PollID parameter.
func (b *PollsDeleteVoteBuilder) PollID(v int) *PollsDeleteVoteBuilder {
	b.Params["poll_id"] = v
	return b
}

// AnswerID parameter.
func (b *PollsDeleteVoteBuilder) AnswerID(v int) *PollsDeleteVoteBuilder {
	b.Params["answer_id"] = v
	return b
}

// IsBoard parameter.
func (b *PollsDeleteVoteBuilder) IsBoard(v bool) *PollsDeleteVoteBuilder {
	b.Params["is_board"] = v
	return b
}

// PollsEditBuilder builder.
//
// Edits created polls.
//
// https://vk.com/dev/polls.edit
type PollsEditBuilder struct {
	api.Params
}

// NewPollsEditBuilder func.
func NewPollsEditBuilder() *PollsEditBuilder {
	return &PollsEditBuilder{api.Params{}}
}

// OwnerID poll owner id.
func (b *PollsEditBuilder) OwnerID(v int) *PollsEditBuilder {
	b.Params["owner_id"] = v
	return b
}

// PollID edited poll's id.
func (b *PollsEditBuilder) PollID(v int) *PollsEditBuilder {
	b.Params["poll_id"] = v
	return b
}

// Question new question text.
func (b *PollsEditBuilder) Question(v string) *PollsEditBuilder {
	b.Params["question"] = v
	return b
}

// AddAnswers answers list, for example: , "["yes","no","maybe"]".
func (b *PollsEditBuilder) AddAnswers(v string) *PollsEditBuilder {
	b.Params["add_answers"] = v
	return b
}

// EditAnswers object containing answers that need to be edited,, key – answer id, value – new answer text.
// Example: {"382967099":"option1", "382967103":"option2"}".
func (b *PollsEditBuilder) EditAnswers(v string) *PollsEditBuilder {
	b.Params["edit_answers"] = v
	return b
}

// DeleteAnswers list of answer ids to be deleted. For example: "[382967099, 382967103]".
func (b *PollsEditBuilder) DeleteAnswers(v string) *PollsEditBuilder {
	b.Params["delete_answers"] = v
	return b
}

// EndDate parameter.
func (b *PollsEditBuilder) EndDate(v int) *PollsEditBuilder {
	b.Params["end_date"] = v
	return b
}

// PhotoID parameter.
func (b *PollsEditBuilder) PhotoID(v int) *PollsEditBuilder {
	b.Params["photo_id"] = v
	return b
}

// BackgroundID parameter.
func (b *PollsEditBuilder) BackgroundID(v string) *PollsEditBuilder {
	b.Params["background_id"] = v
	return b
}

// PollsGetByIDBuilder builder.
//
// Returns detailed information about a poll by its ID.
//
// https://vk.com/dev/polls.getById
type PollsGetByIDBuilder struct {
	api.Params
}

// NewPollsGetByIDBuilder func.
func NewPollsGetByIDBuilder() *PollsGetByIDBuilder {
	return &PollsGetByIDBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsGetByIDBuilder) OwnerID(v int) *PollsGetByIDBuilder {
	b.Params["owner_id"] = v
	return b
}

// IsBoard '1' – poll is in a board, '0' – poll is on a wall. '0' by default.
func (b *PollsGetByIDBuilder) IsBoard(v bool) *PollsGetByIDBuilder {
	b.Params["is_board"] = v
	return b
}

// PollID parameter.
func (b *PollsGetByIDBuilder) PollID(v int) *PollsGetByIDBuilder {
	b.Params["poll_id"] = v
	return b
}

// Extended parameter.
func (b *PollsGetByIDBuilder) Extended(v bool) *PollsGetByIDBuilder {
	b.Params["extended"] = v
	return b
}

// FriendsCount parameter.
func (b *PollsGetByIDBuilder) FriendsCount(v int) *PollsGetByIDBuilder {
	b.Params["friends_count"] = v
	return b
}

// Fields parameter.
func (b *PollsGetByIDBuilder) Fields(v []string) *PollsGetByIDBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase parameter.
func (b *PollsGetByIDBuilder) NameCase(v string) *PollsGetByIDBuilder {
	b.Params["name_case"] = v
	return b
}

// PollsGetVotersBuilder builder.
//
// Returns a list of IDs of users who selected specific answers in the poll.
//
// https://vk.com/dev/polls.getVoters
type PollsGetVotersBuilder struct {
	api.Params
}

// NewPollsGetVotersBuilder func.
func NewPollsGetVotersBuilder() *PollsGetVotersBuilder {
	return &PollsGetVotersBuilder{api.Params{}}
}

// OwnerID ID of the user or community that owns the poll. Use a negative value to designate a community ID.
func (b *PollsGetVotersBuilder) OwnerID(v int) *PollsGetVotersBuilder {
	b.Params["owner_id"] = v
	return b
}

// PollID parameter.
func (b *PollsGetVotersBuilder) PollID(v int) *PollsGetVotersBuilder {
	b.Params["poll_id"] = v
	return b
}

// AnswerIDs parameter.
func (b *PollsGetVotersBuilder) AnswerIDs(v []int) *PollsGetVotersBuilder {
	b.Params["answer_ids"] = v
	return b
}

// IsBoard parameter.
func (b *PollsGetVotersBuilder) IsBoard(v bool) *PollsGetVotersBuilder {
	b.Params["is_board"] = v
	return b
}

// FriendsOnly '1' — to return only current user's friends, '0' — to return all users (default).
func (b *PollsGetVotersBuilder) FriendsOnly(v bool) *PollsGetVotersBuilder {
	b.Params["friends_only"] = v
	return b
}

// Offset needed to return a specific subset of voters. '0' — (default).
func (b *PollsGetVotersBuilder) Offset(v int) *PollsGetVotersBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of user IDs to return (if the 'friends_only' parameter is not set, maximum '1000', otherwise '10').
// '100' — (default).
func (b *PollsGetVotersBuilder) Count(v int) *PollsGetVotersBuilder {
	b.Params["count"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate (birthdate)', 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education',
// 'online', 'counters'.
func (b *PollsGetVotersBuilder) Fields(v []string) *PollsGetVotersBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase case for declension of user name and surname: , 'nom' — nominative (default) , 'gen' — genitive ,
// 'dat' — dative , 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional.
func (b *PollsGetVotersBuilder) NameCase(v string) *PollsGetVotersBuilder {
	b.Params["name_case"] = v
	return b
}
