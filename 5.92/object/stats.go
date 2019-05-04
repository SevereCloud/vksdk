package object // import "github.com/SevereCloud/vksdk/5.92/object"

type statsActivity struct {
	Comments     int `json:"comments"`
	Copies       int `json:"copies"`
	Hidden       int `json:"hidden"`
	Likes        int `json:"likes"`
	Subscribed   int `json:"subscribed"`
	Unsubscribed int `json:"unsubscribed"`
}

type statsCity struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type statsCountry struct {
	Code  string `json:"code"`
	Count int    `json:"count"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type statsPeriod struct {
	Activity   statsActivity `json:"activity"`
	PeriodFrom string        `json:"period_from"`
	PeriodTo   string        `json:"period_to"`
	Reach      statsReach    `json:"reach"`
	Visitors   statsViews    `json:"visitors"`
}

type statsReach struct {
	Age              []statsSexAge  `json:"age"`
	Cities           []statsCity    `json:"cities"`
	Countries        []statsCountry `json:"countries"`
	MobileReach      int            `json:"mobile_reach"`
	Reach            int            `json:"reach"`
	ReachSubscribers int            `json:"reach_subscribers"`
	Sex              []statsSexAge  `json:"sex"`
	SexAge           []statsSexAge  `json:"sex_age"`
}

type statsSexAge struct {
	Count int    `json:"count"`
	Value string `json:"value"`
}

type statsViews struct {
	Age         []statsSexAge  `json:"age"`
	Cities      []statsCity    `json:"cities"`
	Countries   []statsCountry `json:"countries"`
	MobileViews int            `json:"mobile_views"`
	Sex         []statsSexAge  `json:"sex"`
	SexAge      []statsSexAge  `json:"sex_age"`
	Views       int            `json:"views"`
	Visitors    int            `json:"visitors"`
}

type statsWallpostStat struct {
	Hide             int `json:"hide"`
	JoinGroup        int `json:"join_group"`
	Links            int `json:"links"`
	ReachSubscribers int `json:"reach_subscribers"`
	ReachTotal       int `json:"reach_total"`
	Report           int `json:"report"`
	ToGroup          int `json:"to_group"`
	Unsubscribe      int `json:"unsubscribe"`
}
