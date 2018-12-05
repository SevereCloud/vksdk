package object // import "github.com/severecloud/vksdk/object"

type boardTopic struct {
	Comments  int    `json:"comments"`
	Created   int    `json:"created"`
	CreatedBy int    `json:"created_by"`
	ID        int    `json:"id"`
	IsClosed  int    `json:"is_closed"`
	IsFixed   int    `json:"is_fixed"`
	Title     string `json:"title"`
	Updated   int    `json:"updated"`
	UpdatedBy int    `json:"updated_by"`
}

type boardTopicComment struct {
	Attachments  []wallCommentAttachment `json:"attachments"`
	Date         int                     `json:"date"`
	FromID       int                     `json:"from_id"`
	ID           int                     `json:"id"`
	RealOffset   int                     `json:"real_offset"`
	Text         string                  `json:"text"`
	TopicID      int                     `json:"topic_id"`
	TopicOwnerID int                     `json:"topic_owner_id"`
}

type boardTopicPoll struct {
	AnswerID int           `json:"answer_id"`
	Answers  []pollsAnswer `json:"answers"`
	Created  int           `json:"created"`
	IsClosed int           `json:"is_closed"`
	OwnerID  int           `json:"owner_id"`
	PollID   int           `json:"poll_id"`
	Question string        `json:"question"`
	Votes    string        `json:"votes"`
}
