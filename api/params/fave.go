package params // import "github.com/SevereCloud/vksdk/v2/api/params"

import (
	"github.com/SevereCloud/vksdk/v2/api"
)

// FaveAddArticleBuilder builder.
//
// https://vk.com/dev/fave.addArticle
type FaveAddArticleBuilder struct {
	api.Params
}

// NewFaveAddArticleBuilder func.
func NewFaveAddArticleBuilder() *FaveAddArticleBuilder {
	return &FaveAddArticleBuilder{api.Params{}}
}

// URL parameter.
func (b *FaveAddArticleBuilder) URL(v string) *FaveAddArticleBuilder {
	b.Params["url"] = v
	return b
}

// FaveAddLinkBuilder builder.
//
// Adds a link to user faves.
//
// https://vk.com/dev/fave.addLink
type FaveAddLinkBuilder struct {
	api.Params
}

// NewFaveAddLinkBuilder func.
func NewFaveAddLinkBuilder() *FaveAddLinkBuilder {
	return &FaveAddLinkBuilder{api.Params{}}
}

// Link URL.
func (b *FaveAddLinkBuilder) Link(v string) *FaveAddLinkBuilder {
	b.Params["link"] = v
	return b
}

// FaveAddPageBuilder builder.
//
// https://vk.com/dev/fave.addPage
type FaveAddPageBuilder struct {
	api.Params
}

// NewFaveAddPageBuilder func.
func NewFaveAddPageBuilder() *FaveAddPageBuilder {
	return &FaveAddPageBuilder{api.Params{}}
}

// UserID parameter.
func (b *FaveAddPageBuilder) UserID(v int) *FaveAddPageBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID parameter.
func (b *FaveAddPageBuilder) GroupID(v int) *FaveAddPageBuilder {
	b.Params["group_id"] = v
	return b
}

// FaveAddPostBuilder builder.
//
// https://vk.com/dev/fave.addPost
type FaveAddPostBuilder struct {
	api.Params
}

// NewFaveAddPostBuilder func.
func NewFaveAddPostBuilder() *FaveAddPostBuilder {
	return &FaveAddPostBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *FaveAddPostBuilder) OwnerID(v int) *FaveAddPostBuilder {
	b.Params["owner_id"] = v
	return b
}

// ID parameter.
func (b *FaveAddPostBuilder) ID(v int) *FaveAddPostBuilder {
	b.Params["id"] = v
	return b
}

// AccessKey parameter.
func (b *FaveAddPostBuilder) AccessKey(v string) *FaveAddPostBuilder {
	b.Params["access_key"] = v
	return b
}

// FaveAddProductBuilder builder.
//
// https://vk.com/dev/fave.addProduct
type FaveAddProductBuilder struct {
	api.Params
}

// NewFaveAddProductBuilder func.
func NewFaveAddProductBuilder() *FaveAddProductBuilder {
	return &FaveAddProductBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *FaveAddProductBuilder) OwnerID(v int) *FaveAddProductBuilder {
	b.Params["owner_id"] = v
	return b
}

// ID parameter.
func (b *FaveAddProductBuilder) ID(v int) *FaveAddProductBuilder {
	b.Params["id"] = v
	return b
}

// AccessKey parameter.
func (b *FaveAddProductBuilder) AccessKey(v string) *FaveAddProductBuilder {
	b.Params["access_key"] = v
	return b
}

// FaveAddTagBuilder builder.
//
// https://vk.com/dev/fave.addTag
type FaveAddTagBuilder struct {
	api.Params
}

// NewFaveAddTagBuilder func.
func NewFaveAddTagBuilder() *FaveAddTagBuilder {
	return &FaveAddTagBuilder{api.Params{}}
}

// Name parameter.
func (b *FaveAddTagBuilder) Name(v string) *FaveAddTagBuilder {
	b.Params["name"] = v
	return b
}

// FaveAddVideoBuilder builder.
//
// https://vk.com/dev/fave.addVideo
type FaveAddVideoBuilder struct {
	api.Params
}

// NewFaveAddVideoBuilder func.
func NewFaveAddVideoBuilder() *FaveAddVideoBuilder {
	return &FaveAddVideoBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *FaveAddVideoBuilder) OwnerID(v int) *FaveAddVideoBuilder {
	b.Params["owner_id"] = v
	return b
}

// ID parameter.
func (b *FaveAddVideoBuilder) ID(v int) *FaveAddVideoBuilder {
	b.Params["id"] = v
	return b
}

// AccessKey parameter.
func (b *FaveAddVideoBuilder) AccessKey(v string) *FaveAddVideoBuilder {
	b.Params["access_key"] = v
	return b
}

// FaveEditTagBuilder builder.
//
// https://vk.com/dev/fave.editTag
type FaveEditTagBuilder struct {
	api.Params
}

// NewFaveEditTagBuilder func.
func NewFaveEditTagBuilder() *FaveEditTagBuilder {
	return &FaveEditTagBuilder{api.Params{}}
}

// ID parameter.
func (b *FaveEditTagBuilder) ID(v int) *FaveEditTagBuilder {
	b.Params["id"] = v
	return b
}

// Name parameter.
func (b *FaveEditTagBuilder) Name(v string) *FaveEditTagBuilder {
	b.Params["name"] = v
	return b
}

// FaveGetBuilder builder.
//
// https://vk.com/dev/fave.get
type FaveGetBuilder struct {
	api.Params
}

// NewFaveGetBuilder func.
func NewFaveGetBuilder() *FaveGetBuilder {
	return &FaveGetBuilder{api.Params{}}
}

// Extended '1' — to return additional 'wall', 'profiles', and 'groups' fields. By default: '0'.
func (b *FaveGetBuilder) Extended(v bool) *FaveGetBuilder {
	b.Params["extended"] = v
	return b
}

// ItemType parameter.
func (b *FaveGetBuilder) ItemType(v string) *FaveGetBuilder {
	b.Params["item_type"] = v
	return b
}

// TagID parameter.
func (b *FaveGetBuilder) TagID(v int) *FaveGetBuilder {
	b.Params["tag_id"] = v
	return b
}

// Offset needed to return a specific subset of users.
func (b *FaveGetBuilder) Offset(v int) *FaveGetBuilder {
	b.Params["offset"] = v
	return b
}

// Count number of users to return.
func (b *FaveGetBuilder) Count(v int) *FaveGetBuilder {
	b.Params["count"] = v
	return b
}

// Fields parameter.
func (b *FaveGetBuilder) Fields(v string) *FaveGetBuilder {
	b.Params["fields"] = v
	return b
}

// IsFromSnackbar parameter.
func (b *FaveGetBuilder) IsFromSnackbar(v bool) *FaveGetBuilder {
	b.Params["is_from_snackbar"] = v
	return b
}

// FaveGetPagesBuilder builder.
//
// https://vk.com/dev/fave.getPages
type FaveGetPagesBuilder struct {
	api.Params
}

// NewFaveGetPagesBuilder func.
func NewFaveGetPagesBuilder() *FaveGetPagesBuilder {
	return &FaveGetPagesBuilder{api.Params{}}
}

// Offset parameter.
func (b *FaveGetPagesBuilder) Offset(v int) *FaveGetPagesBuilder {
	b.Params["offset"] = v
	return b
}

// Count parameter.
func (b *FaveGetPagesBuilder) Count(v int) *FaveGetPagesBuilder {
	b.Params["count"] = v
	return b
}

// Type parameter.
func (b *FaveGetPagesBuilder) Type(v string) *FaveGetPagesBuilder {
	b.Params["type"] = v
	return b
}

// Fields parameter.
func (b *FaveGetPagesBuilder) Fields(v []string) *FaveGetPagesBuilder {
	b.Params["fields"] = v
	return b
}

// TagID parameter.
func (b *FaveGetPagesBuilder) TagID(v int) *FaveGetPagesBuilder {
	b.Params["tag_id"] = v
	return b
}

// FaveRemoveArticleBuilder builder.
//
// https://vk.com/dev/fave.removeArticle
type FaveRemoveArticleBuilder struct {
	api.Params
}

// NewFaveRemoveArticleBuilder func.
func NewFaveRemoveArticleBuilder() *FaveRemoveArticleBuilder {
	return &FaveRemoveArticleBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *FaveRemoveArticleBuilder) OwnerID(v int) *FaveRemoveArticleBuilder {
	b.Params["owner_id"] = v
	return b
}

// ArticleID parameter.
func (b *FaveRemoveArticleBuilder) ArticleID(v int) *FaveRemoveArticleBuilder {
	b.Params["article_id"] = v
	return b
}

// FaveRemoveLinkBuilder builder.
//
// Removes link from the user's faves.
//
// https://vk.com/dev/fave.removeLink
type FaveRemoveLinkBuilder struct {
	api.Params
}

// NewFaveRemoveLinkBuilder func.
func NewFaveRemoveLinkBuilder() *FaveRemoveLinkBuilder {
	return &FaveRemoveLinkBuilder{api.Params{}}
}

// LinkID parameter (can be obtained by [vk.com/dev/faves.getLinks|faves.getLinks] method).
func (b *FaveRemoveLinkBuilder) LinkID(v string) *FaveRemoveLinkBuilder {
	b.Params["link_id"] = v
	return b
}

// Link URL.
func (b *FaveRemoveLinkBuilder) Link(v string) *FaveRemoveLinkBuilder {
	b.Params["link"] = v
	return b
}

// FaveRemovePageBuilder builder.
//
// https://vk.com/dev/fave.removePage
type FaveRemovePageBuilder struct {
	api.Params
}

// NewFaveRemovePageBuilder func.
func NewFaveRemovePageBuilder() *FaveRemovePageBuilder {
	return &FaveRemovePageBuilder{api.Params{}}
}

// UserID parameter.
func (b *FaveRemovePageBuilder) UserID(v int) *FaveRemovePageBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID parameter.
func (b *FaveRemovePageBuilder) GroupID(v int) *FaveRemovePageBuilder {
	b.Params["group_id"] = v
	return b
}

// FaveRemovePostBuilder builder.
//
// https://vk.com/dev/fave.removePost
type FaveRemovePostBuilder struct {
	api.Params
}

// NewFaveRemovePostBuilder func.
func NewFaveRemovePostBuilder() *FaveRemovePostBuilder {
	return &FaveRemovePostBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *FaveRemovePostBuilder) OwnerID(v int) *FaveRemovePostBuilder {
	b.Params["owner_id"] = v
	return b
}

// ID parameter.
func (b *FaveRemovePostBuilder) ID(v int) *FaveRemovePostBuilder {
	b.Params["id"] = v
	return b
}

// FaveRemoveProductBuilder builder.
//
// https://vk.com/dev/fave.removeProduct
type FaveRemoveProductBuilder struct {
	api.Params
}

// NewFaveRemoveProductBuilder func.
func NewFaveRemoveProductBuilder() *FaveRemoveProductBuilder {
	return &FaveRemoveProductBuilder{api.Params{}}
}

// OwnerID parameter.
func (b *FaveRemoveProductBuilder) OwnerID(v int) *FaveRemoveProductBuilder {
	b.Params["owner_id"] = v
	return b
}

// ID parameter.
func (b *FaveRemoveProductBuilder) ID(v int) *FaveRemoveProductBuilder {
	b.Params["id"] = v
	return b
}

// FaveRemoveTagBuilder builder.
//
// https://vk.com/dev/fave.removeTag
type FaveRemoveTagBuilder struct {
	api.Params
}

// NewFaveRemoveTagBuilder func.
func NewFaveRemoveTagBuilder() *FaveRemoveTagBuilder {
	return &FaveRemoveTagBuilder{api.Params{}}
}

// ID parameter.
func (b *FaveRemoveTagBuilder) ID(v int) *FaveRemoveTagBuilder {
	b.Params["id"] = v
	return b
}

// FaveReorderTagsBuilder builder.
//
// https://vk.com/dev/fave.reorderTags
type FaveReorderTagsBuilder struct {
	api.Params
}

// NewFaveReorderTagsBuilder func.
func NewFaveReorderTagsBuilder() *FaveReorderTagsBuilder {
	return &FaveReorderTagsBuilder{api.Params{}}
}

// IDs parameter.
func (b *FaveReorderTagsBuilder) IDs(v []int) *FaveReorderTagsBuilder {
	b.Params["ids"] = v
	return b
}

// FaveSetPageTagsBuilder builder.
//
// https://vk.com/dev/fave.setPageTags
type FaveSetPageTagsBuilder struct {
	api.Params
}

// NewFaveSetPageTagsBuilder func.
func NewFaveSetPageTagsBuilder() *FaveSetPageTagsBuilder {
	return &FaveSetPageTagsBuilder{api.Params{}}
}

// UserID parameter.
func (b *FaveSetPageTagsBuilder) UserID(v int) *FaveSetPageTagsBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID parameter.
func (b *FaveSetPageTagsBuilder) GroupID(v int) *FaveSetPageTagsBuilder {
	b.Params["group_id"] = v
	return b
}

// TagIDs parameter.
func (b *FaveSetPageTagsBuilder) TagIDs(v []int) *FaveSetPageTagsBuilder {
	b.Params["tag_ids"] = v
	return b
}

// FaveSetTagsBuilder builder.
//
// https://vk.com/dev/fave.setTags
type FaveSetTagsBuilder struct {
	api.Params
}

// NewFaveSetTagsBuilder func.
func NewFaveSetTagsBuilder() *FaveSetTagsBuilder {
	return &FaveSetTagsBuilder{api.Params{}}
}

// ItemType parameter.
func (b *FaveSetTagsBuilder) ItemType(v string) *FaveSetTagsBuilder {
	b.Params["item_type"] = v
	return b
}

// ItemOwnerID parameter.
func (b *FaveSetTagsBuilder) ItemOwnerID(v int) *FaveSetTagsBuilder {
	b.Params["item_owner_id"] = v
	return b
}

// ItemID parameter.
func (b *FaveSetTagsBuilder) ItemID(v int) *FaveSetTagsBuilder {
	b.Params["item_id"] = v
	return b
}

// TagIDs parameter.
func (b *FaveSetTagsBuilder) TagIDs(v []int) *FaveSetTagsBuilder {
	b.Params["tag_ids"] = v
	return b
}

// LinkID parameter.
func (b *FaveSetTagsBuilder) LinkID(v string) *FaveSetTagsBuilder {
	b.Params["link_id"] = v
	return b
}

// LinkURL parameter.
func (b *FaveSetTagsBuilder) LinkURL(v string) *FaveSetTagsBuilder {
	b.Params["link_url"] = v
	return b
}

// FaveTrackPageInteractionBuilder builder.
//
// https://vk.com/dev/fave.trackPageInteraction
type FaveTrackPageInteractionBuilder struct {
	api.Params
}

// NewFaveTrackPageInteractionBuilder func.
func NewFaveTrackPageInteractionBuilder() *FaveTrackPageInteractionBuilder {
	return &FaveTrackPageInteractionBuilder{api.Params{}}
}

// UserID parameter.
func (b *FaveTrackPageInteractionBuilder) UserID(v int) *FaveTrackPageInteractionBuilder {
	b.Params["user_id"] = v
	return b
}

// GroupID parameter.
func (b *FaveTrackPageInteractionBuilder) GroupID(v int) *FaveTrackPageInteractionBuilder {
	b.Params["group_id"] = v
	return b
}
