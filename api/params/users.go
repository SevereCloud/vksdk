package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// UsersGetBuilder builder.
//
// Returns detailed information on users.
//
// https://vk.com/dev/users.get
type UsersGetBuilder struct {
	api.Params
}

// NewUsersGetBuilder func.
func NewUsersGetBuilder() *UsersGetBuilder {
	return &UsersGetBuilder{api.Params{}}
}

// UserIDs user IDs or screen names ('screen_name'). By default, current user ID.
func (b *UsersGetBuilder) UserIDs(v []string) *UsersGetBuilder {
	b.Params["user_ids"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'contacts', 'education', 'online',
// 'counters', 'relation', 'last_seen', 'activity', 'can_write_private_message', 'can_see_all_posts', 'can_post',
// 'universities'.
func (b *UsersGetBuilder) Fields(v []string) *UsersGetBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase case for declension of user name and surname:
//
// * nom — nominative (default);
//
// * gen — genitive;
//
// * dat — dative;
//
// * acc — accusative;
//
// * ins — instrumental;
//
// * abl — prepositional.
func (b *UsersGetBuilder) NameCase(v string) *UsersGetBuilder {
	b.Params["name_case"] = v
	return b
}

// UsersGetFollowersBuilder builder.
//
// Returns a list of IDs of followers of the user in question, sorted by date added, most recent first.
//
// https://vk.com/dev/users.getFollowers
type UsersGetFollowersBuilder struct {
	api.Params
}

// NewUsersGetFollowersBuilder func.
func NewUsersGetFollowersBuilder() *UsersGetFollowersBuilder {
	return &UsersGetFollowersBuilder{api.Params{}}
}

// UserID parameter.
func (b *UsersGetFollowersBuilder) UserID(v int) *UsersGetFollowersBuilder {
	b.Params["user_id"] = v
	return b
}

// Offset needed to return a specific subset of followers.
func (b *UsersGetFollowersBuilder) Offset(v int) *UsersGetFollowersBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of followers to return.
func (b *UsersGetFollowersBuilder) Count(v int) *UsersGetFollowersBuilder {
	b.Params["count"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate),
// 'city', 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts',
// 'education', 'online'.
func (b *UsersGetFollowersBuilder) Fields(v []string) *UsersGetFollowersBuilder {
	b.Params["fields"] = v
	return b
}

// NameCase case for declension of user name and surname:
//
// * nom — nominative (default);
//
// * gen — genitive;
//
// * dat — dative;
//
// * acc — accusative;
//
// * ins — instrumental;
//
// * abl — prepositional.
func (b *UsersGetFollowersBuilder) NameCase(v string) *UsersGetFollowersBuilder {
	b.Params["name_case"] = v
	return b
}

// UsersGetSubscriptionsBuilder builder.
//
// Returns a list of IDs of users and communities followed by the user.
//
// https://vk.com/dev/users.getSubscriptions
type UsersGetSubscriptionsBuilder struct {
	api.Params
}

// NewUsersGetSubscriptionsBuilder func.
func NewUsersGetSubscriptionsBuilder() *UsersGetSubscriptionsBuilder {
	return &UsersGetSubscriptionsBuilder{api.Params{}}
}

// UserID parameter.
func (b *UsersGetSubscriptionsBuilder) UserID(v int) *UsersGetSubscriptionsBuilder {
	b.Params["user_id"] = v
	return b
}

// Extended '1' — to return a combined list of users and communities, '0' — to return separate lists of users
// and communities (default).
func (b *UsersGetSubscriptionsBuilder) Extended(v bool) *UsersGetSubscriptionsBuilder {
	b.Params["extended"] = v
	return b
}

// Offset needed to return a specific subset of subscriptions.
func (b *UsersGetSubscriptionsBuilder) Offset(v int) *UsersGetSubscriptionsBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of users and communities to return.
func (b *UsersGetSubscriptionsBuilder) Count(v int) *UsersGetSubscriptionsBuilder {
	b.Params["count"] = v
	return b
}

// Fields parameter.
func (b *UsersGetSubscriptionsBuilder) Fields(v []string) *UsersGetSubscriptionsBuilder {
	b.Params["fields"] = v
	return b
}

// UsersIsAppUserBuilder builder.
//
// Returns information whether a user installed the application.
//
// https://vk.com/dev/users.isAppUser
type UsersIsAppUserBuilder struct {
	api.Params
}

// NewUsersIsAppUserBuilder func.
func NewUsersIsAppUserBuilder() *UsersIsAppUserBuilder {
	return &UsersIsAppUserBuilder{api.Params{}}
}

// UserID parameter.
func (b *UsersIsAppUserBuilder) UserID(v int) *UsersIsAppUserBuilder {
	b.Params["user_id"] = v
	return b
}

// UsersReportBuilder builder.
//
// Reports (submits a complain about) a user.
//
// https://vk.com/dev/users.report
type UsersReportBuilder struct {
	api.Params
}

// NewUsersReportBuilder func.
func NewUsersReportBuilder() *UsersReportBuilder {
	return &UsersReportBuilder{api.Params{}}
}

// UserID ID of the user about whom a complaint is being made.
func (b *UsersReportBuilder) UserID(v int) *UsersReportBuilder {
	b.Params["user_id"] = v
	return b
}

// Type Type of complaint:
// 'porn' – pornography, 'spam' – spamming, 'insult' – abusive behavior, 'advertisement' – disruptive advertisements.
func (b *UsersReportBuilder) Type(v string) *UsersReportBuilder {
	b.Params["type"] = v
	return b
}

// Comment describing the complaint.
func (b *UsersReportBuilder) Comment(v string) *UsersReportBuilder {
	b.Params["comment"] = v
	return b
}

// UsersSearchBuilder builder.
//
// Returns a list of users matching the search criteria.
//
// https://vk.com/dev/users.search
type UsersSearchBuilder struct {
	api.Params
}

// NewUsersSearchBuilder func.
func NewUsersSearchBuilder() *UsersSearchBuilder {
	return &UsersSearchBuilder{api.Params{}}
}

// Q search query string (e.g., 'Vasya Babich').
func (b *UsersSearchBuilder) Q(v string) *UsersSearchBuilder {
	b.Params["q"] = v
	return b
}

// Sort order: '1' — by date registered, '0' — by rating.
func (b *UsersSearchBuilder) Sort(v int) *UsersSearchBuilder {
	b.Params["sort"] = v
	return b
}

// Offset needed to return a specific subset of users.
func (b *UsersSearchBuilder) Offset(v int) *UsersSearchBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of users to return.
func (b *UsersSearchBuilder) Count(v int) *UsersSearchBuilder {
	b.Params["count"] = v
	return b
}

// Fields profile fields to return. Sample values: 'nickname', 'screen_name', 'sex', 'bdate' (birthdate), 'city',
// 'country', 'timezone', 'photo', 'photo_medium', 'photo_big', 'has_mobile', 'rate', 'contacts', 'education',
// 'online'.
func (b *UsersSearchBuilder) Fields(v []string) *UsersSearchBuilder {
	b.Params["fields"] = v
	return b
}

// City parameter.
func (b *UsersSearchBuilder) City(v int) *UsersSearchBuilder {
	b.Params["city"] = v
	return b
}

// Country parameter.
func (b *UsersSearchBuilder) Country(v int) *UsersSearchBuilder {
	b.Params["country"] = v
	return b
}

// Hometown city name in a string.
func (b *UsersSearchBuilder) Hometown(v string) *UsersSearchBuilder {
	b.Params["hometown"] = v
	return b
}

// UniversityCountry ID of the country where the user graduated.
func (b *UsersSearchBuilder) UniversityCountry(v int) *UsersSearchBuilder {
	b.Params["university_country"] = v
	return b
}

// University ID of the institution of higher education.
func (b *UsersSearchBuilder) University(v int) *UsersSearchBuilder {
	b.Params["university"] = v
	return b
}

// UniversityYear year of graduation from an institution of higher education.
func (b *UsersSearchBuilder) UniversityYear(v int) *UsersSearchBuilder {
	b.Params["university_year"] = v
	return b
}

// UniversityFaculty faculty ID.
func (b *UsersSearchBuilder) UniversityFaculty(v int) *UsersSearchBuilder {
	b.Params["university_faculty"] = v
	return b
}

// UniversityChair chair ID.
func (b *UsersSearchBuilder) UniversityChair(v int) *UsersSearchBuilder {
	b.Params["university_chair"] = v
	return b
}

// Sex '1' — female, '2' — male, '0' — any (default).
func (b *UsersSearchBuilder) Sex(v int) *UsersSearchBuilder {
	b.Params["sex"] = v
	return b
}

// Status relationship status:
// '1' — Not married, '2' — In a relationship, '3' — Engaged, '4' — Married,
// '5' — It's complicated, '6' — Actively searching, '7' — In love.
func (b *UsersSearchBuilder) Status(v int) *UsersSearchBuilder {
	b.Params["status"] = v
	return b
}

// AgeFrom minimum age.
func (b *UsersSearchBuilder) AgeFrom(v int) *UsersSearchBuilder {
	b.Params["age_from"] = v
	return b
}

// AgeTo maximum age.
func (b *UsersSearchBuilder) AgeTo(v int) *UsersSearchBuilder {
	b.Params["age_to"] = v
	return b
}

// BirthDay day of birth.
func (b *UsersSearchBuilder) BirthDay(v int) *UsersSearchBuilder {
	b.Params["birth_day"] = v
	return b
}

// BirthMonth month of birth.
func (b *UsersSearchBuilder) BirthMonth(v int) *UsersSearchBuilder {
	b.Params["birth_month"] = v
	return b
}

// BirthYear year of birth.
func (b *UsersSearchBuilder) BirthYear(v int) *UsersSearchBuilder {
	b.Params["birth_year"] = v
	return b
}

// Online '1' — online only, '0' — all users.
func (b *UsersSearchBuilder) Online(v bool) *UsersSearchBuilder {
	b.Params["online"] = v
	return b
}

// HasPhoto '1' — with photo only, '0' — all users.
func (b *UsersSearchBuilder) HasPhoto(v bool) *UsersSearchBuilder {
	b.Params["has_photo"] = v
	return b
}

// SchoolCountry ID of the country where users finished school.
func (b *UsersSearchBuilder) SchoolCountry(v int) *UsersSearchBuilder {
	b.Params["school_country"] = v
	return b
}

// SchoolCity ID of the city where users finished school.
func (b *UsersSearchBuilder) SchoolCity(v int) *UsersSearchBuilder {
	b.Params["school_city"] = v
	return b
}

// SchoolClass parameter.
func (b *UsersSearchBuilder) SchoolClass(v int) *UsersSearchBuilder {
	b.Params["school_class"] = v
	return b
}

// School ID of the school.
func (b *UsersSearchBuilder) School(v int) *UsersSearchBuilder {
	b.Params["school"] = v
	return b
}

// SchoolYear school graduation year.
func (b *UsersSearchBuilder) SchoolYear(v int) *UsersSearchBuilder {
	b.Params["school_year"] = v
	return b
}

// Religion users religious affiliation.
func (b *UsersSearchBuilder) Religion(v string) *UsersSearchBuilder {
	b.Params["religion"] = v
	return b
}

// Interests users interests.
func (b *UsersSearchBuilder) Interests(v string) *UsersSearchBuilder {
	b.Params["interests"] = v
	return b
}

// Company name where users work.
func (b *UsersSearchBuilder) Company(v string) *UsersSearchBuilder {
	b.Params["company"] = v
	return b
}

// Position job position.
func (b *UsersSearchBuilder) Position(v string) *UsersSearchBuilder {
	b.Params["position"] = v
	return b
}

// GroupID to search in communities.
func (b *UsersSearchBuilder) GroupID(v int) *UsersSearchBuilder {
	b.Params["group_id"] = v
	return b
}

// FromList parameter.
func (b *UsersSearchBuilder) FromList(v []string) *UsersSearchBuilder {
	b.Params["from_list"] = v
	return b
}
