package params // import "github.com/SevereCloud/vksdk/api/params"

import (
	"github.com/SevereCloud/vksdk/api"
)

// UsersGetBulder builder
//
// Returns detailed information on users.
//
// https://vk.com/dev/users.get
type UsersGetBulder struct {
	api.Params
}

// NewUsersGetBulder func
func NewUsersGetBulder() *UsersGetBulder {
	return &UsersGetBulder{api.Params{}}
}

// UserIDs User IDs or screen names ('screen_name'). By default, current user ID.
func (b *UsersGetBulder) UserIDs(v []string) {
	b.Params["user_ids"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online', 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post', 'universities',
func (b *UsersGetBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *UsersGetBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// UsersGetFollowersBulder builder
//
// Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// https://vk.com/dev/users.getFollowers
type UsersGetFollowersBulder struct {
	api.Params
}

// NewUsersGetFollowersBulder func
func NewUsersGetFollowersBulder() *UsersGetFollowersBulder {
	return &UsersGetFollowersBulder{api.Params{}}
}

// UserID User ID.
func (b *UsersGetFollowersBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Offset Offset needed to return a specific subset of followers.
func (b *UsersGetFollowersBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of followers to return.
func (b *UsersGetFollowersBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online'.
func (b *UsersGetFollowersBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// NameCase Case for declension of user name and surname: 'nom' — nominative (default), 'gen' — genitive , 'dat' — dative, 'acc' — accusative , 'ins' — instrumental , 'abl' — prepositional
func (b *UsersGetFollowersBulder) NameCase(v string) {
	b.Params["name_case"] = v
}

// UsersGetSubscriptionsBulder builder
//
// Returns a list of IDs of users and communities followed by the user.
//
// https://vk.com/dev/users.getSubscriptions
type UsersGetSubscriptionsBulder struct {
	api.Params
}

// NewUsersGetSubscriptionsBulder func
func NewUsersGetSubscriptionsBulder() *UsersGetSubscriptionsBulder {
	return &UsersGetSubscriptionsBulder{api.Params{}}
}

// UserID User ID.
func (b *UsersGetSubscriptionsBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Extended '1' — to return a combined list of users and communities, '0' — to return separate lists of users and communities (default)
func (b *UsersGetSubscriptionsBulder) Extended(v bool) {
	b.Params["extended"] = v
}

// Offset Offset needed to return a specific subset of subscriptions.
func (b *UsersGetSubscriptionsBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users and communities to return.
func (b *UsersGetSubscriptionsBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields parameter
func (b *UsersGetSubscriptionsBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// UsersIsAppUserBulder builder
//
// Returns information whether a user installed the application.
//
// https://vk.com/dev/users.isAppUser
type UsersIsAppUserBulder struct {
	api.Params
}

// NewUsersIsAppUserBulder func
func NewUsersIsAppUserBulder() *UsersIsAppUserBulder {
	return &UsersIsAppUserBulder{api.Params{}}
}

// UserID parameter
func (b *UsersIsAppUserBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// UsersReportBulder builder
//
// Reports (submits a complain about) a user.
//
// https://vk.com/dev/users.report
type UsersReportBulder struct {
	api.Params
}

// NewUsersReportBulder func
func NewUsersReportBulder() *UsersReportBulder {
	return &UsersReportBulder{api.Params{}}
}

// UserID ID of the user about whom a complaint is being made.
func (b *UsersReportBulder) UserID(v int) {
	b.Params["user_id"] = v
}

// Type Type of complaint: 'porn' – pornography, 'spam' – spamming, 'insult' – abusive behavior, 'advertisement' – disruptive advertisements
func (b *UsersReportBulder) Type(v string) {
	b.Params["type"] = v
}

// Comment Comment describing the complaint.
func (b *UsersReportBulder) Comment(v string) {
	b.Params["comment"] = v
}

// UsersSearchBulder builder
//
// Returns a list of users matching the search criteria.
//
// https://vk.com/dev/users.search
type UsersSearchBulder struct {
	api.Params
}

// NewUsersSearchBulder func
func NewUsersSearchBulder() *UsersSearchBulder {
	return &UsersSearchBulder{api.Params{}}
}

// Q Search query string (e.g., 'Vasya Babich').
func (b *UsersSearchBulder) Q(v string) {
	b.Params["q"] = v
}

// Sort Sort order: '1' — by date registered, '0' — by rating
func (b *UsersSearchBulder) Sort(v int) {
	b.Params["sort"] = v
}

// Offset Offset needed to return a specific subset of users.
func (b *UsersSearchBulder) Offset(v int) {
	b.Params["offset"] = v
}

// Count Number of users to return.
func (b *UsersSearchBulder) Count(v int) {
	b.Params["count"] = v
}

// Fields Profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education', 'online',
func (b *UsersSearchBulder) Fields(v []string) {
	b.Params["fields"] = v
}

// City City ID.
func (b *UsersSearchBulder) City(v int) {
	b.Params["city"] = v
}

// Country Country ID.
func (b *UsersSearchBulder) Country(v int) {
	b.Params["country"] = v
}

// Hometown City name in a string.
func (b *UsersSearchBulder) Hometown(v string) {
	b.Params["hometown"] = v
}

// UniversityCountry ID of the country where the user graduated.
func (b *UsersSearchBulder) UniversityCountry(v int) {
	b.Params["university_country"] = v
}

// University ID of the institution of higher education.
func (b *UsersSearchBulder) University(v int) {
	b.Params["university"] = v
}

// UniversityYear Year of graduation from an institution of higher education.
func (b *UsersSearchBulder) UniversityYear(v int) {
	b.Params["university_year"] = v
}

// UniversityFaculty Faculty ID.
func (b *UsersSearchBulder) UniversityFaculty(v int) {
	b.Params["university_faculty"] = v
}

// UniversityChair Chair ID.
func (b *UsersSearchBulder) UniversityChair(v int) {
	b.Params["university_chair"] = v
}

// Sex '1' — female, '2' — male, '0' — any (default)
func (b *UsersSearchBulder) Sex(v int) {
	b.Params["sex"] = v
}

// Status Relationship status: '1' — Not married, '2' — In a relationship, '3' — Engaged, '4' — Married, '5' — It's complicated, '6' — Actively searching, '7' — In love
func (b *UsersSearchBulder) Status(v int) {
	b.Params["status"] = v
}

// AgeFrom Minimum age.
func (b *UsersSearchBulder) AgeFrom(v int) {
	b.Params["age_from"] = v
}

// AgeTo Maximum age.
func (b *UsersSearchBulder) AgeTo(v int) {
	b.Params["age_to"] = v
}

// BirthDay Day of birth.
func (b *UsersSearchBulder) BirthDay(v int) {
	b.Params["birth_day"] = v
}

// BirthMonth Month of birth.
func (b *UsersSearchBulder) BirthMonth(v int) {
	b.Params["birth_month"] = v
}

// BirthYear Year of birth.
func (b *UsersSearchBulder) BirthYear(v int) {
	b.Params["birth_year"] = v
}

// Online '1' — online only, '0' — all users
func (b *UsersSearchBulder) Online(v bool) {
	b.Params["online"] = v
}

// HasPhoto '1' — with photo only, '0' — all users
func (b *UsersSearchBulder) HasPhoto(v bool) {
	b.Params["has_photo"] = v
}

// SchoolCountry ID of the country where users finished school.
func (b *UsersSearchBulder) SchoolCountry(v int) {
	b.Params["school_country"] = v
}

// SchoolCity ID of the city where users finished school.
func (b *UsersSearchBulder) SchoolCity(v int) {
	b.Params["school_city"] = v
}

// SchoolClass parameter
func (b *UsersSearchBulder) SchoolClass(v int) {
	b.Params["school_class"] = v
}

// School ID of the school.
func (b *UsersSearchBulder) School(v int) {
	b.Params["school"] = v
}

// SchoolYear School graduation year.
func (b *UsersSearchBulder) SchoolYear(v int) {
	b.Params["school_year"] = v
}

// Religion Users' religious affiliation.
func (b *UsersSearchBulder) Religion(v string) {
	b.Params["religion"] = v
}

// Interests Users' interests.
func (b *UsersSearchBulder) Interests(v string) {
	b.Params["interests"] = v
}

// Company Name of the company where users work.
func (b *UsersSearchBulder) Company(v string) {
	b.Params["company"] = v
}

// Position Job position.
func (b *UsersSearchBulder) Position(v string) {
	b.Params["position"] = v
}

// GroupID ID of a community to search in communities.
func (b *UsersSearchBulder) GroupID(v int) {
	b.Params["group_id"] = v
}

// FromList parameter
func (b *UsersSearchBulder) FromList(v []string) {
	b.Params["from_list"] = v
}
