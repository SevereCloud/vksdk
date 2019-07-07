package api // import "github.com/SevereCloud/vksdk/5.92/api"

import (
	"github.com/SevereCloud/vksdk/5.92/object"
)

// StatsGetResponse struct
type StatsGetResponse []object.StatsPeriod

// StatsGet returns statistics of a community or an application.
//
// https://vk.com/dev/stats.get
func (vk *VK) StatsGet(params map[string]string) (response StatsGetResponse, vkErr Error) {
	vk.RequestUnmarshal("stats.get", params, &response, &vkErr)
	return
}

// StatsGetPostReachResponse struct
type StatsGetPostReachResponse []object.StatsWallpostStat

// StatsGetPostReach returns stats for a wall post.
//
// https://vk.com/dev/stats.getPostReach
func (vk *VK) StatsGetPostReach(params map[string]string) (response StatsGetPostReachResponse, vkErr Error) {
	vk.RequestUnmarshal("stats.getPostReach", params, &response, &vkErr)
	return
}

// StatsTrackVisitorResponse struct
type StatsTrackVisitorResponse int

// StatsTrackVisitor adds current session's data in the application statistics.
//
// https://vk.com/dev/stats.trackVisitor
func (vk *VK) StatsTrackVisitor(params map[string]string) (response StatsTrackVisitorResponse, vkErr Error) {
	vk.RequestUnmarshal("stats.trackVisitor", params, &response, &vkErr)
	return
}
