package object // import "github.com/SevereCloud/vksdk/5.92/object"

// PagesWikipage struct
type PagesWikipage struct {
	CreatorID   int    `json:"creator_id"`   // Page creator ID
	CreatorName int    `json:"creator_name"` // Page creator name
	EditorID    int    `json:"editor_id"`    // Last editor ID
	EditorName  string `json:"editor_name"`  // Last editor name
	GroupID     int    `json:"group_id"`     // Community ID
	ID          int    `json:"id"`           // Page ID
	Title       string `json:"title"`        // Page title
	Views       int    `json:"views"`        // Views number
	WhoCanEdit  int    `json:"who_can_edit"` // Edit settings of the page
	WhoCanView  int    `json:"who_can_view"` // View settings of the page
}

// PagesWikipageFull struct
type PagesWikipageFull struct {
	Created                  int    `json:"created"`                      // Date when the page has been created in Unixtime
	CreatorID                int    `json:"creator_id"`                   // Page creator ID
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`        // Information whether current user can edit the page
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"` // Information whether current user can edit the page access settings
	Edited                   int    `json:"edited"`                       // Date when the page has been edited in Unixtime
	EditorID                 int    `json:"editor_id"`                    // Last editor ID
	GroupID                  int    `json:"group_id"`                     // Community ID
	HTML                     string `json:"html"`                         // Page content, HTML
	ID                       int    `json:"id"`                           // Page ID
	Source                   string `json:"source"`                       // Page content, wiki
	Title                    string `json:"title"`                        // Page title
	ViewURL                  string `json:"view_url"`                     // URL of the page preview
	Views                    int    `json:"views"`                        // Views number
	WhoCanEdit               int    `json:"who_can_edit"`                 // Edit settings of the page
	WhoCanView               int    `json:"who_can_view"`                 // View settings of the page
}

// PagesWikipageHistory struct
// BUG(VK): https://vk.com/dev/pages.getHistory edited and date
type PagesWikipageHistory struct {
	Date       int    `json:"date"`        // Date when the page has been edited in Unixtime
	EditorID   int    `json:"editor_id"`   // Last editor ID
	EditorName string `json:"editor_name"` // Last editor name
	ID         int    `json:"id"`          // Version ID
	Length     int    `json:"length"`      // Page size in bytes
}
