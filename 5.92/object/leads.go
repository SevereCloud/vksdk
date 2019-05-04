package object // import "github.com/SevereCloud/vksdk/5.92/object"

type leadsChecked struct {
	Reason    string `json:"reason"`
	Result    string `json:"result"`
	Sid       string `json:"sid"`
	StartLink string `json:"start_link"`
}

type leadsComplete struct {
	Cost     int `json:"cost"`
	Limit    int `json:"limit"`
	Spent    int `json:"spent"`
	Success  int `json:"success"`
	TestMode int `json:"test_mode"`
}

type leadsEntry struct {
	Aid       int    `json:"aid"`
	Comment   string `json:"comment"`
	Date      int    `json:"date"`
	Sid       string `json:"sid"`
	StartDate int    `json:"start_date"`
	Status    int    `json:"status"`
	TestMode  int    `json:"test_mode"`
	UID       int    `json:"uid"`
}

type leadsLead struct {
	Completed   int           `json:"completed"`
	Cost        int           `json:"cost"`
	Days        leadsLeadDays `json:"days"`
	Impressions int           `json:"impressions"`
	Limit       int           `json:"limit"`
	Spent       int           `json:"spent"`
	Started     int           `json:"started"`
}

type leadsLeadDays struct {
	Completed   int `json:"completed"`
	Impressions int `json:"impressions"`
	Spent       int `json:"spent"`
	Started     int `json:"started"`
}

type leadsStart struct {
	TestMode int    `json:"test_mode"`
	VkSid    string `json:"vk_sid"`
}
