package object // import "github.com/severecloud/vksdk/object"

type pollsAnswer struct {
	ID    int     `json:"id"`
	Rate  float64 `json:"rate"`
	Text  string  `json:"text"`
	Votes int     `json:"votes"`
}

type pollsPoll struct {
	Anonymous int           `json:"anonymous"`
	AnswerID  int           `json:"answer_id"`
	Answers   []pollsAnswer `json:"answers"`
	Created   int           `json:"created"`
	ID        int           `json:"id"`
	OwnerID   int           `json:"owner_id"`
	Question  string        `json:"question"`
	Votes     string        `json:"votes"`
}

type pollsVoters struct {
	AnswerID int              `json:"answer_id"`
	Users    pollsVotersUsers `json:"users"`
}

type pollsVotersUsers struct {
	Count int   `json:"count"`
	Items []int `json:"items"`
}
