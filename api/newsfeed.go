package api

// NewsfeedAddBanResponse struct
type NewsfeedAddBanResponse struct{}

// TODO: newsfeed.addBan Prevents news from specified users and communities from appearing in the current user's newsfeed.

// NewsfeedDeleteBanResponse struct
type NewsfeedDeleteBanResponse struct{}

// TODO: newsfeed.deleteBan Allows news from previously banned users and communities to be shown in the current user's newsfeed.

// NewsfeedDeleteListResponse struct
type NewsfeedDeleteListResponse struct{}

// TODO: newsfeed.deleteList

// NewsfeedGetResponse struct
type NewsfeedGetResponse struct{}

// TODO: newsfeed.get Returns data required to show newsfeed for the current user.

// NewsfeedGetBannedResponse struct
type NewsfeedGetBannedResponse struct{}

// TODO: newsfeed.getBanned Returns a list of users and communities banned from the current user's newsfeed.

// NewsfeedGetCommentsResponse struct
type NewsfeedGetCommentsResponse struct{}

// TODO: newsfeed.getComments Returns a list of comments in the current user's newsfeed.

// NewsfeedGetDiscoverForContestantResponse struct
type NewsfeedGetDiscoverForContestantResponse struct{}

// TODO: newsfeed.getDiscoverForContestant

// NewsfeedGetListsResponse struct
type NewsfeedGetListsResponse struct{}

// TODO: newsfeed.getLists Returns a list of newsfeeds followed by the current user.

// NewsfeedGetMentionsResponse struct
type NewsfeedGetMentionsResponse struct{}

// TODO: newsfeed.getMentions Returns a list of posts on user walls in which the current user is mentioned.

// NewsfeedGetRecommendedResponse struct
type NewsfeedGetRecommendedResponse struct{}

// TODO: newsfeed.getRecommended Returns a list of newsfeeds recommended to the current user.

// NewsfeedGetSuggestedSourcesResponse struct
type NewsfeedGetSuggestedSourcesResponse struct{}

// TODO: newsfeed.getSuggestedSources Returns communities and users that current user is suggested to follow.

// NewsfeedIgnoreItemHidesResponse struct
type NewsfeedIgnoreItemHidesResponse struct{}

// TODO: newsfeed.ignoreItemHides an item from the newsfeed.

// NewsfeedSaveListResponse struct
type NewsfeedSaveListResponse struct{}

// TODO: newsfeed.saveList Creates and edits user newsfeed lists

// NewsfeedSearchResponse struct
type NewsfeedSearchResponse struct{}

// TODO: newsfeed.search Returns search results by statuses.

// NewsfeedUnignoreItemResponse struct
type NewsfeedUnignoreItemResponse struct{}

// TODO: newsfeed.unignoreItem Returns a hidden item to the newsfeed.

// NewsfeedUnsubscribeResponse struct
type NewsfeedUnsubscribeResponse struct{}

// TODO: newsfeed.unsubscribe Unsubscribes the current user from specified newsfeeds.
