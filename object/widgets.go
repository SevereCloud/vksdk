package object

// WidgetsCommentMedia struct
type widgetsCommentMedia struct {
	ItemID   int                     `json:"item_id"`
	OwnerID  int                     `json:"owner_id"`
	ThumbSrc string                  `json:"thumb_src"`
	Type     widgetsCommentMediaType `json:"type"`
}

// WidgetsCommentMediaType struct
type widgetsCommentMediaType string

// WidgetsCommentReplies struct
type widgetsCommentReplies struct {
	CanPost int                         `json:"can_post"`
	Count   int                         `json:"count"`
	Replies []widgetsCommentRepliesItem `json:"replies"`
}

// WidgetsCommentRepliesItem struct
type widgetsCommentRepliesItem struct {
	Cid   int                `json:"cid"`
	Date  int                `json:"date"`
	Likes widgetsWidgetLikes `json:"likes"`
	Text  string             `json:"text"`
	UID   int                `json:"uid"`
	User  UsersUser          `json:"user"`
}

// WidgetsWidgetComment struct
type widgetsWidgetComment struct {
	Attachments []wallCommentAttachment `json:"attachments"`
	CanDelete   int                     `json:"can_delete"`
	Comments    widgetsCommentReplies   `json:"comments"`
	Date        int                     `json:"date"`
	FromID      int                     `json:"from_id"`
	ID          int                     `json:"id"`
	Likes       baseLikesInfo           `json:"likes"`
	Media       widgetsCommentMedia     `json:"media"`
	PostSource  wallPostSource          `json:"post_source"`
	PostType    int                     `json:"post_type"`
	Reposts     baseRepostsInfo         `json:"reposts"`
	Text        string                  `json:"text"`
	ToID        int                     `json:"to_id"`
	User        UsersUser               `json:"user"`
}

// WidgetsWidgetLikes struct
type widgetsWidgetLikes struct {
	Count int `json:"count"`
}

// WidgetsWidgetPage struct
type widgetsWidgetPage struct {
	Comments    BaseObjectCount `json:"comments"`
	Date        int             `json:"date"`
	Description string          `json:"description"`
	ID          int             `json:"id"`
	Likes       BaseObjectCount `json:"likes"`
	PageID      string          `json:"page_id"`
	Photo       string          `json:"photo"`
	Title       string          `json:"title"`
	URL         string          `json:"url"`
}
