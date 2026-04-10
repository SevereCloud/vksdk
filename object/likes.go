package object // import "github.com/SevereCloud/vksdk/v3/object"

// LikesReactionSet struct.
type LikesReactionSet struct {
	ID    string `json:"id"`
	Items []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Asset struct {
			AnimationURL string `json:"animation_url"`
			Images       []struct {
				URL    string `json:"url"`
				Width  int    `json:"width"`
				Height int    `json:"height"`
			} `json:"images"`
			Title struct {
				Color struct {
					Foreground struct {
						Light string `json:"light"`
						Dark  string `json:"dark"`
					} `json:"foreground"`
					Background struct {
						Light string `json:"light"`
						Dark  string `json:"dark"`
					} `json:"background"`
				} `json:"color"`
			} `json:"title"`
			TitleColor struct {
				Light string `json:"light"`
				Dark  string `json:"dark"`
			} `json:"title_color"`
		} `json:"asset"`
	} `json:"items"`
}
