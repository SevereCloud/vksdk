package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/stretchr/testify/assert"
)

func TestUsersGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersGetBuilder()

	b.UserIDs([]string{"text"})
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, []string{"text"}, b.Params["user_ids"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
}

func TestUsersGetFollowersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersGetFollowersBuilder()

	b.UserID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, "text", b.Params["name_case"])
}

func TestUsersGetSubscriptionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersGetSubscriptionsBuilder()

	b.UserID(1)
	b.Extended(true)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, true, b.Params["extended"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
}

func TestUsersIsAppUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersIsAppUserBuilder()

	b.UserID(1)

	assert.Equal(t, 1, b.Params["user_id"])
}

func TestUsersReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersReportBuilder()

	b.UserID(1)
	b.Type("text")
	b.Comment("text")

	assert.Equal(t, 1, b.Params["user_id"])
	assert.Equal(t, "text", b.Params["type"])
	assert.Equal(t, "text", b.Params["comment"])
}

func TestUsersSearchBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersSearchBuilder()

	b.Q("text")
	b.Sort(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.City(1)
	b.Country(1)
	b.Hometown("text")
	b.UniversityCountry(1)
	b.University(1)
	b.UniversityYear(1)
	b.UniversityFaculty(1)
	b.UniversityChair(1)
	b.Sex(1)
	b.Status(1)
	b.AgeFrom(1)
	b.AgeTo(1)
	b.BirthDay(1)
	b.BirthMonth(1)
	b.BirthYear(1)
	b.Online(true)
	b.HasPhoto(true)
	b.SchoolCountry(1)
	b.SchoolCity(1)
	b.SchoolClass(1)
	b.School(1)
	b.SchoolYear(1)
	b.Religion("text")
	b.Interests("text")
	b.Company("text")
	b.Position("text")
	b.GroupID(1)
	b.FromList([]string{"text"})

	assert.Equal(t, "text", b.Params["q"])
	assert.Equal(t, 1, b.Params["sort"])
	assert.Equal(t, 1, b.Params["offset"])
	assert.Equal(t, 1, b.Params["count"])
	assert.Equal(t, []string{"test"}, b.Params["fields"])
	assert.Equal(t, 1, b.Params["city"])
	assert.Equal(t, 1, b.Params["country"])
	assert.Equal(t, "text", b.Params["hometown"])
	assert.Equal(t, 1, b.Params["university_country"])
	assert.Equal(t, 1, b.Params["university"])
	assert.Equal(t, 1, b.Params["university_year"])
	assert.Equal(t, 1, b.Params["university_faculty"])
	assert.Equal(t, 1, b.Params["university_chair"])
	assert.Equal(t, 1, b.Params["sex"])
	assert.Equal(t, 1, b.Params["status"])
	assert.Equal(t, 1, b.Params["age_from"])
	assert.Equal(t, 1, b.Params["age_to"])
	assert.Equal(t, 1, b.Params["birth_day"])
	assert.Equal(t, 1, b.Params["birth_month"])
	assert.Equal(t, 1, b.Params["birth_year"])
	assert.Equal(t, true, b.Params["online"])
	assert.Equal(t, true, b.Params["has_photo"])
	assert.Equal(t, 1, b.Params["school_country"])
	assert.Equal(t, 1, b.Params["school_city"])
	assert.Equal(t, 1, b.Params["school_class"])
	assert.Equal(t, 1, b.Params["school"])
	assert.Equal(t, 1, b.Params["school_year"])
	assert.Equal(t, "text", b.Params["religion"])
	assert.Equal(t, "text", b.Params["interests"])
	assert.Equal(t, "text", b.Params["company"])
	assert.Equal(t, "text", b.Params["position"])
	assert.Equal(t, 1, b.Params["group_id"])
	assert.Equal(t, []string{"text"}, b.Params["from_list"])
}
