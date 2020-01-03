package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// UsersGetBuilder builder
//
// Returns detailed information on users.
//
// https://vk.com/dev/users.get
type UsersGetBuilder struct {
	api.Params
}

// NewUsersGetBuilder func
func NewUsersGetBuilder() *UsersGetBuilder {
	return &UsersGetBuilder{api.Params{}}
}

// UserIDs User IDs or screen names ('screen_name'). By default, current user ID.
func (b *UsersGetBuilder) UserIDs(v []string) {
	b.Params["user_ids"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities',
func (b *UsersGetBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *UsersGetBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}

// UsersGetFollowersBuilder builder
//
// Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// https://vk.com/dev/users.getFollowers
type UsersGetFollowersBuilder struct {
	api.Params
}

// NewUsersGetFollowersBuilder func
func NewUsersGetFollowersBuilder() *UsersGetFollowersBuilder {
	return &UsersGetFollowersBuilder{api.Params{}}
}

// UserID User ID.
func (b *UsersGetFollowersBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Offset Offset needed to return a specific subset of followers.
func (b *UsersGetFollowersBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of followers to return.
func (b *UsersGetFollowersBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online'.
func (b *UsersGetFollowersBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *UsersGetFollowersBuilder) NameCase(v string) {
	b.Params["name_case"] = v
}

// UsersGetSubscriptionsBuilder builder
//
// Returns a list of IDs of users and communities followed by the user.
//
// https://vk.com/dev/users.getSubscriptions
type UsersGetSubscriptionsBuilder struct {
	api.Params
}

// NewUsersGetSubscriptionsBuilder func
func NewUsersGetSubscriptionsBuilder() *UsersGetSubscriptionsBuilder {
	return &UsersGetSubscriptionsBuilder{api.Params{}}
}

// UserID User ID.
func (b *UsersGetSubscriptionsBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Extended '1' — to return a combined list of users and communities, '0' — to return separate lists of users and communities (default)
func (b *UsersGetSubscriptionsBuilder) Extended(v bool) {
	b.Params["extended"] = v
}

// Offset Offset needed to return a specific subset of subscriptions.
func (b *UsersGetSubscriptionsBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users and communities to return.
func (b *UsersGetSubscriptionsBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields parameter
func (b *UsersGetSubscriptionsBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// UsersIsAppUserBuilder builder
//
// Returns information whether a user installed the application.
//
// https://vk.com/dev/users.isAppUser
type UsersIsAppUserBuilder struct {
	api.Params
}

// NewUsersIsAppUserBuilder func
func NewUsersIsAppUserBuilder() *UsersIsAppUserBuilder {
	return &UsersIsAppUserBuilder{api.Params{}}
}

// UserID parameter
func (b *UsersIsAppUserBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// UsersReportBuilder builder
//
// Reports (submits a complain about) a user.
//
// https://vk.com/dev/users.report
type UsersReportBuilder struct {
	api.Params
}

// NewUsersReportBuilder func
func NewUsersReportBuilder() *UsersReportBuilder {
	return &UsersReportBuilder{api.Params{}}
}

// UserID ID of the user about whom a complaint is being made.
func (b *UsersReportBuilder) UserID(v int) {
	b.Params["user_id"] = v
}

// Type Type of complaint: 'porn' – pornography, 'spam' – spamming, 'insult' – abusive behavior, 'advertisement' – disruptive advertisements
func (b *UsersReportBuilder) Type(v string) {
	b.Params["type"] = v
}

// Comment Comment describing the complaint.
func (b *UsersReportBuilder) Comment(v string) {
	b.Params["comment"] = v
}

// UsersSearchBuilder builder
//
// Returns a list of users matching the search criteria.
//
// https://vk.com/dev/users.search
type UsersSearchBuilder struct {
	api.Params
}

// NewUsersSearchBuilder func
func NewUsersSearchBuilder() *UsersSearchBuilder {
	return &UsersSearchBuilder{api.Params{}}
}

// Q Search query string (e.g., 'Vasya Babich').
func (b *UsersSearchBuilder) Q(v string) {
	b.Params["q"] = v
}

// Sort Sort order: '1' — by date registered, '0' — by rating
func (b *UsersSearchBuilder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of users.
func (b *UsersSearchBuilder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users to return.
func (b *UsersSearchBuilder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
func (b *UsersSearchBuilder) Fields(v []string) {
	b.Params["fields"] = v
}

// City City ID.
func (b *UsersSearchBuilder) City(v int) {
	b.Params["city"] = v
}

// Country Country ID.
func (b *UsersSearchBuilder) Country(v int) {
	b.Params["country"] = v
}

// Hometown City name in a string.
func (b *UsersSearchBuilder) Hometown(v string) {
	b.Params["hometown"] = v
}

// UniversityCountry ID of the country where the user graduated.
func (b *UsersSearchBuilder) UniversityCountry(v int) {
	b.Params["university_country"] = v
}

// University ID of the institution of higher education.
func (b *UsersSearchBuilder) University(v int) {
	b.Params["university"] = v
}

// UniversityYear Year of graduation from an institution of higher education.
func (b *UsersSearchBuilder) UniversityYear(v int) {
	b.Params["university_year"] = v
}

// UniversityFaculty Faculty ID.
func (b *UsersSearchBuilder) UniversityFaculty(v int) {
	b.Params["university_faculty"] = v
}

// UniversityChair Chair ID.
func (b *UsersSearchBuilder) UniversityChair(v int) {
	b.Params["university_chair"] = v
}

// Sex '1' — female, '2' — male, '0' — any (default)
func (b *UsersSearchBuilder) Sex(v int) {
	b.Params["sex"] = v
}

// Status Relationship status: '1' — Not married, '2' — In a relationship, '3' — Engaged, '4' — Married, '5' — It's complicated, '6' — Actively searching, '7' — In love
func (b *UsersSearchBuilder) Status(v int) {
	b.Params["status"] = v
}

// AgeFrom Minimum age.
func (b *UsersSearchBuilder) AgeFrom(v int) {
	b.Params["age_from"] = v
}

// AgeTo Maximum age.
func (b *UsersSearchBuilder) AgeTo(v int) {
	b.Params["age_to"] = v
}

// BirthDay Day of birth.
func (b *UsersSearchBuilder) BirthDay(v int) {
	b.Params["birth_day"] = v
}

// BirthMonth Month of birth.
func (b *UsersSearchBuilder) BirthMonth(v int) {
	b.Params["birth_month"] = v
}

// BirthYear Year of birth.
func (b *UsersSearchBuilder) BirthYear(v int) {
	b.Params["birth_year"] = v
}

// Online '1' — online only, '0' — all users
func (b *UsersSearchBuilder) Online(v bool) {
	b.Params["online"] = v
}

// HasPhoto '1' — with photo only, '0' — all users
func (b *UsersSearchBuilder) HasPhoto(v bool) {
	b.Params["has_photo"] = v
}

// SchoolCountry ID of the country where users finished school.
func (b *UsersSearchBuilder) SchoolCountry(v int) {
	b.Params["school_country"] = v
}

// SchoolCity ID of the city where users finished school.
func (b *UsersSearchBuilder) SchoolCity(v int) {
	b.Params["school_city"] = v
}

// SchoolClass parameter
func (b *UsersSearchBuilder) SchoolClass(v int) {
	b.Params["school_class"] = v
}

// School ID of the school.
func (b *UsersSearchBuilder) School(v int) {
	b.Params["school"] = v
}

// SchoolYear School graduation year.
func (b *UsersSearchBuilder) SchoolYear(v int) {
	b.Params["school_year"] = v
}

// Religion Users' religious affiliation.
func (b *UsersSearchBuilder) Religion(v string) {
	b.Params["religion"] = v
}

// Interests Users' interests.
func (b *UsersSearchBuilder) Interests(v string) {
	b.Params["interests"] = v
}

// Company Name of the company where users work.
func (b *UsersSearchBuilder) Company(v string) {
	b.Params["company"] = v
}

// Position Job position.
func (b *UsersSearchBuilder) Position(v string) {
	b.Params["position"] = v
}

// GroupID ID of a community to search in communities.
func (b *UsersSearchBuilder) GroupID(v int) {
	b.Params["group_id"] = v
}

// FromList parameter
func (b *UsersSearchBuilder) FromList(v []string) {
	b.Params["from_list"] = v
}
