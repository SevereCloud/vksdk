package api // import "github.com/SevereCloud/vksdk/api"

import (
	"github.com/SevereCloud/vksdk/object"
)

// PodcastsGetCatalogResponse struct
type PodcastsGetCatalogResponse struct {
	Items []object.PodcastsItem `json:"items"`
	object.ExtendedResponse
}

// PodcastsGetCatalog
//
// extended=0 method
//
// https://vk.com/dev/podcasts.getCatalog
func (vk *VK) PodcastsGetCatalog(params map[string]string) (response PodcastsGetCatalogResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("podcasts.getCatalog", params, &response)
	return
}

// PodcastsGetCatalogExtendedResponse struct
type PodcastsGetCatalogExtendedResponse struct {
	Items []object.PodcastsItem `json:"items"`
	object.ExtendedResponse
}

// PodcastsGetCatalogExtended
//
// extended=1 method
//
// https://vk.com/dev/podcasts.getCatalog
func (vk *VK) PodcastsGetCatalogExtended(params map[string]string) (response PodcastsGetCatalogExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("podcasts.getCatalog", params, &response)
	return
}

// PodcastsGetCategoriesResponse struct
type PodcastsGetCategoriesResponse []object.PodcastsCategory

// PodcastsGetCategories method
//
// https://vk.com/dev/podcasts.getCategories
func (vk *VK) PodcastsGetCategories(params map[string]string) (response PodcastsGetCategoriesResponse, err error) {
	err = vk.RequestUnmarshal("podcasts.getCategories", params, &response)
	return
}

// PodcastsGetEpisodesResponse struct
type PodcastsGetEpisodesResponse struct {
	Count int                      `json:"count"`
	Items []object.PodcastsEpisode `json:"items"`
}

// PodcastsGetEpisodes method
//
// https://vk.com/dev/podcasts.getEpisodes
func (vk *VK) PodcastsGetEpisodes(params map[string]string) (response PodcastsGetEpisodesResponse, err error) {
	err = vk.RequestUnmarshal("podcasts.getEpisodes", params, &response)
	return
}

// PodcastsGetFeedResponse struct
type PodcastsGetFeedResponse struct {
	Items    []object.PodcastsEpisode `json:"items"`
	NextFrom string                   `json:"next_from"`
}

// PodcastsGetFeed
//
// extended=0 method
//
// https://vk.com/dev/podcasts.getFeed
func (vk *VK) PodcastsGetFeed(params map[string]string) (response PodcastsGetFeedResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("podcasts.getFeed", params, &response)
	return
}

// PodcastsGetFeedExtendedResponse struct
type PodcastsGetFeedExtendedResponse struct {
	Items    []object.PodcastsEpisode `json:"items"`
	NextFrom string                   `json:"next_from"`
	object.ExtendedResponse
}

// PodcastsGetFeedExtended
//
// extended=1 method
//
// https://vk.com/dev/podcasts.getFeed
func (vk *VK) PodcastsGetFeedExtended(params map[string]string) (response PodcastsGetFeedExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("podcasts.getFeed", params, &response)
	return
}

// PodcastsGetStartPageResponse struct
type PodcastsGetStartPageResponse struct {
	Order               []string                  `json:"order"`
	InProgress          []object.PodcastsEpisode  `json:"in_progress"`
	Bookmarks           []object.PodcastsEpisode  `json:"bookmarks"`
	Articles            []object.Article          `json:"articles"`
	StaticHowTo         []bool                    `json:"static_how_to"`
	FriendsLiked        []object.PodcastsEpisode  `json:"friends_liked"`
	Subscriptions       []object.PodcastsEpisode  `json:"subscriptions"`
	CategoriesList      []object.PodcastsCategory `json:"categories_list"`
	RecommendedEpisodes []object.PodcastsEpisode  `json:"recommended_episodes"`
	Catalog             []struct {
		Category object.PodcastsCategory `json:"category"`
		Items    []object.PodcastsItem   `json:"items"`
	} `json:"catalog"`
}

// PodcastsGetStartPage
//
// extended=0 method
//
// https://vk.com/dev/podcasts.getStartPage
func (vk *VK) PodcastsGetStartPage(params map[string]string) (response PodcastsGetStartPageResponse, err error) {
	params["extended"] = "0"
	err = vk.RequestUnmarshal("podcasts.getStartPage", params, &response)
	return
}

// PodcastsGetStartPageExtendedResponse struct
type PodcastsGetStartPageExtendedResponse struct {
	Order               []string                  `json:"order"`
	InProgress          []object.PodcastsEpisode  `json:"in_progress"`
	Bookmarks           []object.PodcastsEpisode  `json:"bookmarks"`
	Articles            []object.Article          `json:"articles"`
	StaticHowTo         []bool                    `json:"static_how_to"`
	FriendsLiked        []object.PodcastsEpisode  `json:"friends_liked"`
	Subscriptions       []object.PodcastsEpisode  `json:"subscriptions"`
	CategoriesList      []object.PodcastsCategory `json:"categories_list"`
	RecommendedEpisodes []object.PodcastsEpisode  `json:"recommended_episodes"`
	Catalog             []struct {
		Category object.PodcastsCategory `json:"category"`
		Items    []object.PodcastsItem   `json:"items"`
	} `json:"catalog"`
	object.ExtendedResponse
}

// PodcastsGetStartPageExtended
//
// extended=1 method
//
// https://vk.com/dev/podcasts.getStartPage
func (vk *VK) PodcastsGetStartPageExtended(params map[string]string) (response PodcastsGetStartPageExtendedResponse, err error) {
	params["extended"] = "1"
	err = vk.RequestUnmarshal("podcasts.getStartPage", params, &response)
	return
}

// PodcastsMarkAsListened method
//
// https://vk.com/dev/podcasts.markAsListened
func (vk *VK) PodcastsMarkAsListened(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("podcasts.markAsListened", params, &response)
	return
}

// PodcastsSubscribe method
//
// https://vk.com/dev/podcasts.subscribe
func (vk *VK) PodcastsSubscribe(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("podcasts.subscribe", params, &response)
	return
}

// PodcastsUnsubscribe method
//
// https://vk.com/dev/podcasts.unsubscribe
func (vk *VK) PodcastsUnsubscribe(params map[string]string) (response int, err error) {
	err = vk.RequestUnmarshal("podcasts.unsubscribe", params, &response)
	return
}
