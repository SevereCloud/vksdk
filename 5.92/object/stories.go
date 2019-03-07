package object // import "github.com/severecloud/vksdk/5.92/object"

type storiesReplies struct {
	Count int `json:"count"`
	New   int `json:"new"`
}

type storiesStoryLink struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type storiesStoryStats struct {
	Answer      storiesStoryStatsStat `json:"answer"`
	Bans        storiesStoryStatsStat `json:"bans"`
	OpenLink    storiesStoryStatsStat `json:"open_link"`
	Replies     storiesStoryStatsStat `json:"replies"`
	Shares      storiesStoryStatsStat `json:"shares"`
	Subscribers storiesStoryStatsStat `json:"subscribers"`
	Views       storiesStoryStatsStat `json:"views"`
}

type storiesStoryStatsStat struct {
	Count int    `json:"count"`
	State string `json:"state"`
}
