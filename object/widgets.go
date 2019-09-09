package object // import "github.com/SevereCloud/vksdk/object"

// WidgetsCommentMedia struct
type widgetsCommentMedia struct {
	ItemID   int                     `json:"item_id"`   // Media item ID
	OwnerID  int                     `json:"owner_id"`  // Media owner's ID
	ThumbSrc string                  `json:"thumb_src"` // URL of the preview image (type=photo only)
	Type     widgetsCommentMediaType `json:"type"`
}

// WidgetsCommentMediaType struct
type widgetsCommentMediaType string

// WidgetsCommentReplies struct
type widgetsCommentReplies struct {
	CanPost       int                         `json:"can_post"` // Information whether current user can comment the post
	Count         int                         `json:"count"`    // Comments number
	Replies       []widgetsCommentRepliesItem `json:"replies"`
	GroupsCanPost bool                        `json:"groups_can_post"`
}

// WidgetsCommentRepliesItem struct
type widgetsCommentRepliesItem struct {
	Cid   int                `json:"cid"`  // Comment ID
	Date  int                `json:"date"` // Date when the comment has been added in Unixtime
	Likes widgetsWidgetLikes `json:"likes"`
	Text  string             `json:"text"` // Comment text
	UID   int                `json:"uid"`  // User ID
	User  UsersUser          `json:"user"`
}

// WidgetsWidgetComment struct
type WidgetsWidgetComment struct {
	Attachments []wallCommentAttachment `json:"attachments"`
	CanDelete   int                     `json:"can_delete"` // Information whether current user can delete the comment
	Comments    widgetsCommentReplies   `json:"comments"`
	Date        int                     `json:"date"`    // Date when the comment has been added in Unixtime
	FromID      int                     `json:"from_id"` // Comment author ID
	ID          int                     `json:"id"`      // Comment ID
	Likes       baseLikesInfo           `json:"likes"`
	Media       widgetsCommentMedia     `json:"media"`
	PostType    int                     `json:"post_type"` // Post type
	Reposts     baseRepostsInfo         `json:"reposts"`
	Text        string                  `json:"text"`  // Comment text
	ToID        int                     `json:"to_id"` // Wall owner
	PostSource  struct {
		Link struct {
			URL         string `json:"url"`
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"link"`
		Type string `json:"type"`
		Data string `json:"data"`
	} `json:"post_source"`
	IsFavorite bool `json:"is_favorite"`
}

// WidgetsWidgetLikes struct
type widgetsWidgetLikes struct {
	Count int `json:"count"` // Likes number
}

// WidgetsWidgetPage struct
type WidgetsWidgetPage struct {
	Comments    widgetsWidgetLikes `json:"comments,omitempty"`
	Date        int                `json:"date,omitempty"`        // Date when widgets on the page has been initialized firstly in Unixtime
	Description string             `json:"description,omitempty"` // Page description
	ID          int                `json:"id,omitempty"`          // Page ID
	Likes       widgetsWidgetLikes `json:"likes,omitempty"`
	PageID      string             `json:"page_id,omitempty"` // page_id parameter value
	Photo       string             `json:"photo,omitempty"`   // URL of the preview image
	Title       string             `json:"title,omitempty"`   // Page title
	URL         string             `json:"url,omitempty"`     // Page absolute URL
}
