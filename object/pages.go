package object

type pagesWikipage struct {
	CreatorID   int    `json:"creator_id"`
	CreatorName int    `json:"creator_name"`
	EditorID    int    `json:"editor_id"`
	EditorName  string `json:"editor_name"`
	GroupID     int    `json:"group_id"`
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Views       int    `json:"views"`
	WhoCanEdit  int    `json:"who_can_edit"`
	WhoCanView  int    `json:"who_can_view"`
}

type pagesWikipageFull struct {
	Created                  int    `json:"created"`
	CreatorID                int    `json:"creator_id"`
	CurrentUserCanEdit       int    `json:"current_user_can_edit"`
	CurrentUserCanEditAccess int    `json:"current_user_can_edit_access"`
	Edited                   int    `json:"edited"`
	EditorID                 int    `json:"editor_id"`
	GroupID                  int    `json:"group_id"`
	HTML                     string `json:"html"`
	ID                       int    `json:"id"`
	Source                   string `json:"source"`
	Title                    string `json:"title"`
	ViewURL                  string `json:"view_url"`
	Views                    int    `json:"views"`
	WhoCanEdit               int    `json:"who_can_edit"`
	WhoCanView               int    `json:"who_can_view"`
}

type pagesWikipageVersion struct {
	Edited     int    `json:"edited"`
	EditorID   int    `json:"editor_id"`
	EditorName string `json:"editor_name"`
	ID         int    `json:"id"`
	Length     int    `json:"length"`
}
