package params_test

import (
	"testing"

	"github.com/SevereCloud/vksdk/api/params"
	"github.com/stretchr/testify/assert"
)

func TestUsersGetBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersGetBuilder()

	b.UserIDs([]string{"text"})
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, b.Params["user_ids"], []string{"text"})
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}

func TestUsersGetFollowersBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersGetFollowersBuilder()

	b.UserID(1)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})
	b.NameCase("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["name_case"], "text")
}

func TestUsersGetSubscriptionsBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersGetSubscriptionsBuilder()

	b.UserID(1)
	b.Extended(true)
	b.Offset(1)
	b.Count(1)
	b.Fields([]string{"test"})

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["extended"], true)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
}

func TestUsersIsAppUserBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersIsAppUserBuilder()

	b.UserID(1)

	assert.Equal(t, b.Params["user_id"], 1)
}

func TestUsersReportBuilder(t *testing.T) {
	t.Parallel()

	b := params.NewUsersReportBuilder()

	b.UserID(1)
	b.Type("text")
	b.Comment("text")

	assert.Equal(t, b.Params["user_id"], 1)
	assert.Equal(t, b.Params["type"], "text")
	assert.Equal(t, b.Params["comment"], "text")
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

	assert.Equal(t, b.Params["q"], "text")
	assert.Equal(t, b.Params["sort"], 1)
	assert.Equal(t, b.Params["offset"], 1)
	assert.Equal(t, b.Params["count"], 1)
	assert.Equal(t, b.Params["fields"], []string{"test"})
	assert.Equal(t, b.Params["city"], 1)
	assert.Equal(t, b.Params["country"], 1)
	assert.Equal(t, b.Params["hometown"], "text")
	assert.Equal(t, b.Params["university_country"], 1)
	assert.Equal(t, b.Params["university"], 1)
	assert.Equal(t, b.Params["university_year"], 1)
	assert.Equal(t, b.Params["university_faculty"], 1)
	assert.Equal(t, b.Params["university_chair"], 1)
	assert.Equal(t, b.Params["sex"], 1)
	assert.Equal(t, b.Params["status"], 1)
	assert.Equal(t, b.Params["age_from"], 1)
	assert.Equal(t, b.Params["age_to"], 1)
	assert.Equal(t, b.Params["birth_day"], 1)
	assert.Equal(t, b.Params["birth_month"], 1)
	assert.Equal(t, b.Params["birth_year"], 1)
	assert.Equal(t, b.Params["online"], true)
	assert.Equal(t, b.Params["has_photo"], true)
	assert.Equal(t, b.Params["school_country"], 1)
	assert.Equal(t, b.Params["school_city"], 1)
	assert.Equal(t, b.Params["school_class"], 1)
	assert.Equal(t, b.Params["school"], 1)
	assert.Equal(t, b.Params["school_year"], 1)
	assert.Equal(t, b.Params["religion"], "text")
	assert.Equal(t, b.Params["interests"], "text")
	assert.Equal(t, b.Params["company"], "text")
	assert.Equal(t, b.Params["position"], "text")
	assert.Equal(t, b.Params["group_id"], 1)
	assert.Equal(t, b.Params["from_list"], []string{"text"})
}
