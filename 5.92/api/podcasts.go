package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// PodcastsGetCatalogResponse struct
type PodcastsGetCatalogResponse struct {
	Items    []object.PodcastsItem `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// PodcastsGetCatalog
//
// extended=0 method
//
// https://vk.com/dev/podcasts.getCatalog
func (vk *VK) PodcastsGetCatalog(params map[string]string) (response PodcastsGetCatalogResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("podcasts.getCatalog", params, &response, &vkErr)
	return
}

// PodcastsGetCatalogExtendedResponse struct
type PodcastsGetCatalogExtendedResponse struct {
	Items    []object.PodcastsItem `json:"items"`
	Profiles []object.UsersUser    `json:"profiles"`
	Groups   []object.GroupsGroup  `json:"groups"`
}

// PodcastsGetCatalogExtended
//
// extended=1 method
//
// https://vk.com/dev/podcasts.getCatalog
func (vk *VK) PodcastsGetCatalogExtended(params map[string]string) (response PodcastsGetCatalogExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("podcasts.getCatalog", params, &response, &vkErr)
	return
}

// PodcastsGetCategoriesResponse struct
type PodcastsGetCategoriesResponse []object.PodcastsCategory

// PodcastsGetCategories method
//
// https://vk.com/dev/podcasts.getCategories
func (vk *VK) PodcastsGetCategories(params map[string]string) (response PodcastsGetCategoriesResponse, vkErr Error) {
	vk.RequestUnmarshal("podcasts.getCategories", params, &response, &vkErr)
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
func (vk *VK) PodcastsGetEpisodes(params map[string]string) (response PodcastsGetEpisodesResponse, vkErr Error) {
	vk.RequestUnmarshal("podcasts.getEpisodes", params, &response, &vkErr)
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
func (vk *VK) PodcastsGetFeed(params map[string]string) (response PodcastsGetFeedResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("podcasts.getFeed", params, &response, &vkErr)
	return
}

// PodcastsGetFeedExtendedResponse struct
type PodcastsGetFeedExtendedResponse struct {
	Items    []object.PodcastsEpisode `json:"items"`
	NextFrom string                   `json:"next_from"`
	Profiles []object.UsersUser       `json:"profiles"`
	Groups   []object.GroupsGroup     `json:"groups"`
}

// PodcastsGetFeedExtended
//
// extended=1 method
//
// https://vk.com/dev/podcasts.getFeed
func (vk *VK) PodcastsGetFeedExtended(params map[string]string) (response PodcastsGetFeedExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("podcasts.getFeed", params, &response, &vkErr)
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
func (vk *VK) PodcastsGetStartPage(params map[string]string) (response PodcastsGetStartPageResponse, vkErr Error) {
	params["extended"] = "0"
	vk.RequestUnmarshal("podcasts.getStartPage", params, &response, &vkErr)
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
	Profiles []object.UsersUser   `json:"profiles"`
	Groups   []object.GroupsGroup `json:"groups"`
}

// PodcastsGetStartPageExtended
//
// extended=1 method
//
// https://vk.com/dev/podcasts.getStartPage
func (vk *VK) PodcastsGetStartPageExtended(params map[string]string) (response PodcastsGetStartPageExtendedResponse, vkErr Error) {
	params["extended"] = "1"
	vk.RequestUnmarshal("podcasts.getStartPage", params, &response, &vkErr)
	return
}

// PodcastsMarkAsListenedResponse struct
type PodcastsMarkAsListenedResponse int

// PodcastsMarkAsListened method
//
// https://vk.com/dev/podcasts.markAsListened
func (vk *VK) PodcastsMarkAsListened(params map[string]string) (response PodcastsMarkAsListenedResponse, vkErr Error) {
	vk.RequestUnmarshal("podcasts.markAsListened", params, &response, &vkErr)
	return
}

// PodcastsSubscribeResponse struct
type PodcastsSubscribeResponse int

// PodcastsSubscribe method
//
// https://vk.com/dev/podcasts.subscribe
func (vk *VK) PodcastsSubscribe(params map[string]string) (response PodcastsSubscribeResponse, vkErr Error) {
	vk.RequestUnmarshal("podcasts.subscribe", params, &response, &vkErr)
	return
}

// PodcastsUnsubscribeResponse struct
type PodcastsUnsubscribeResponse int

// PodcastsUnsubscribe method
//
// https://vk.com/dev/podcasts.unsubscribe
func (vk *VK) PodcastsUnsubscribe(params map[string]string) (response PodcastsUnsubscribeResponse, vkErr Error) {
	vk.RequestUnmarshal("podcasts.unsubscribe", params, &response, &vkErr)
	return
}
