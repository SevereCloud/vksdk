package object // import "github.com/SevereCloud/vksdk/object"

type statsActivity struct {
	Comments     int `json:"comments"`     // Comments number
	Copies       int `json:"copies"`       // Reposts number
	Hidden       int `json:"hidden"`       // Hidden from news count
	Likes        int `json:"likes"`        // Likes number
	Subscribed   int `json:"subscribed"`   // New subscribers count
	Unsubscribed int `json:"unsubscribed"` // Unsubscribed count
}

type statsCity struct {
	Count int    `json:"count"` // Visitors number
	Name  string `json:"name"`  // City name
	Value int    `json:"value"` // City ID
}

type statsCountry struct {
	Code  string `json:"code"`  // Country code
	Count int    `json:"count"` // Visitors number
	Name  string `json:"name"`  // Country name
	Value int    `json:"value"` // Country ID
}

// StatsPeriod struct
type StatsPeriod struct {
	Activity   statsActivity `json:"activity"`
	PeriodFrom int           `json:"period_from"` // Unix timestamp
	PeriodTo   int           `json:"period_to"`   // Unix timestamp
	Reach      statsReach    `json:"reach"`
	Visitors   statsViews    `json:"visitors"`
}

type statsReach struct {
	Age              []statsSexAge  `json:"age"`
	Cities           []statsCity    `json:"cities"`
	Countries        []statsCountry `json:"countries"`
	MobileReach      int            `json:"mobile_reach"`      // Reach count from mobile devices
	Reach            int            `json:"reach"`             // Reach count
	ReachSubscribers int            `json:"reach_subscribers"` // Subscribers reach count
	Sex              []statsSexAge  `json:"sex"`
	SexAge           []statsSexAge  `json:"sex_age"`
}

type statsSexAge struct {
	Count int    `json:"count"` // Visitors number
	Value string `json:"value"` // Sex/age value
}

type statsViews struct {
	Age         []statsSexAge  `json:"age"`
	Cities      []statsCity    `json:"cities"`
	Countries   []statsCountry `json:"countries"`
	MobileViews int            `json:"mobile_views"` // Number of views from mobile devices
	Sex         []statsSexAge  `json:"sex"`
	SexAge      []statsSexAge  `json:"sex_age"`
	Views       int            `json:"views"`    // Views number
	Visitors    int            `json:"visitors"` // Visitors number
}

// StatsWallpostStat struct
type StatsWallpostStat struct {
	Hide             int `json:"hide"`              // Hidings number
	JoinGroup        int `json:"join_group"`        // People have joined the group
	Links            int `json:"links"`             // Link clickthrough
	ReachSubscribers int `json:"reach_subscribers"` // Subscribers reach
	ReachTotal       int `json:"reach_total"`       // Total reach
	Report           int `json:"report"`            // Reports number
	ToGroup          int `json:"to_group"`          // Clickthrough to community
	Unsubscribe      int `json:"unsubscribe"`       // Unsubscribed members
}
