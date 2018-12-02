package vksdk

// Структуры ответов

// TODO appWidgets.getAppImageUploadServer response

// TODO appWidgets.getAppImages response

// TODO appWidgets.getGroupImageUploadServer response

// TODO appWidgets.getGroupImages response

// TODO appWidgets.getImagesById response

// TODO appWidgets.saveAppImage response

// TODO appWidgets.saveGroupImage response

// TODO appWidgets.update response

// TODO apps.deleteAppRequests response

// TODO apps.get response

// TODO apps.getCatalog response

// TODO apps.getFriendsList response

// TODO apps.getLeaderboard response

// TODO apps.getScopes response

// TODO apps.getScore response

// TODO apps.sendRequest response

// TODO auth.checkPhone response

// TODO auth.restore response

// TODO board.addTopic response

// TODO board.closeTopic response

// TODO board.createComment response

// TODO board.deleteComment response

// TODO board.deleteTopic response

// TODO board.editComment response

// TODO board.editTopic response

// TODO board.fixTopic response

// TODO board.getComments response

// TODO board.getTopics response

// TODO board.openTopicRe response

// TODO board.restoreComment response

// TODO board.unfixTopic response

// TODO database.getChairs response

// DatabaseGetCitiesResponse database.getCities response

// TODO database.getCitiesById response

// TODO database.getCountries response

// TODO database.getCountriesById response

// TODO database.getFaculties response

// TODO database.getRegions response

// TODO database.getSchoolClasses response

// TODO database.getSchools response

// TODO database.getUniversities response

// TODO docs.add response

// TODO docs.delete response

// TODO docs.edit response

// TODO docs.get response

// TODO docs.getById response

// TODO docs.getMessagesUploadServer response

// TODO docs.getTypes response

// TODO docs.getUploadServer response

// TODO docs.getWallUploadServer response

// TODO docs.save response

// TODO docs.search response

// TODO fave.addGroup response

// TODO fave.addLink response

// TODO fave.addUser response

// TODO fave.getLinks response

// TODO fave.getMarketItems response

// TODO fave.getPhotos response

// TODO fave.getPosts response

// TODO fave.getUsers response

// TODO fave.getVideos response

// TODO fave.removeGroup response

// TODO fave.removeLink response

// TODO fave.removeUser response

// TODO friends.add response

// TODO friends.addList response

// TODO friends.areFriends response

// TODO friends.delete response

// TODO friends.deleteAllRequests response

// TODO friends.deleteList response

// TODO friends.edit response

// TODO friends.editList response

// TODO friends.getAppUsers response

// TODO friends.getByPhones response

// TODO friends.getLists response

// TODO friends.getMutual response

// TODO friends.getOnline response

// TODO friends.getRecent response

// TODO friends.getRequests response

// TODO friends.getSuggestions response

// TODO friends.search response

// TODO gifts.get response

// TODO groups.addCallbackServerAdds response

// TODO groups.addLink response

// TODO groups.approveRequest response

// TODO groups.ban response

// TODO groups.create response

// TODO groups.deleteCallbackServerDeletes response

// TODO groups.deleteLink response

// TODO groups.disableOnline response

// TODO groups.edit response

// TODO groups.editCallbackServer response

// TODO groups.editLink response

// TODO groups.editManager response

// TODO groups.editPlace response

// TODO groups.enableOnline response

// TODO groups.get response

// TODO groups.getAddresses response

// TODO groups.getBanned response

// TODO groups.getById response

// TODO groups.getCallbackConfirmationCode response

// TODO groups.getCallbackServers response

// TODO groups.getCallbackSettings response

// TODO groups.getCatalog response

// TODO groups.getCatalogInfo response

// TODO groups.getInvitedUsers response

// TODO groups.getInvites response

// TODO groups.getLongPollServer response

// TODO groups.getLongPollSettings response

// TODO groups.getOnlineStatus response

// TODO groups.getRequests response

// TODO groups.getSettings response

// TODO groups.getTokenPermissions response

// TODO groups.invite response

// TODO groups.isMember response

// TODO groups.join response

// TODO groups.leave response

// TODO groups.removeUser response

// TODO groups.reorderLink response

// TODO groups.search response

// TODO groups.setCallbackSettings response

// TODO groups.setLongPollSettings response

// TODO groups.unban response

// TODO leadForms.create response

// TODO leadForms.delete response

// TODO leadForms.get response

// TODO leadForms.getLeads response

// TODO leadForms.getUploadURL response

// TODO leadForms.list response

// TODO leadForms.update response

// TODO likes.add response

// TODO likes.delete response

// TODO likes.getList response

// TODO likes.isLiked response

// TODO market.add response

// TODO market.addAlbum response

// TODO market.addToAlbum response

// TODO market.createComment response

// TODO market.delete response

// TODO market.deleteAlbum response

// TODO market.deleteComment response

// TODO market.edit response

// TODO market.editAlbum response

// TODO market.editComment response

// TODO market.get response

// TODO market.getAlbumById response

// TODO market.getAlbums response

// TODO market.getById response

// TODO market.getCategories response

// TODO market.getComments response

// TODO market.removeFromAlbum response

// TODO market.reorderAlbums response

// TODO market.reorderItems response

// TODO market.report response

// TODO market.reportComment response

// TODO market.restore response

// TODO market.restoreComment response

// TODO market.search response

// TODO messages.addChatUser response

// TODO messages.allowMessagesFromGroup response

// TODO messages.createChat response

// TODO messages.delete response

// TODO messages.deleteChatPhoto response

// TODO messages.deleteConversation response

// TODO messages.denyMessagesFromGroup response

// TODO messages.edit response

// TODO messages.editChat response

// TODO messages.getByConversationMessageId response

// TODO messages.getById response

// TODO messages.getChat response

// TODO messages.getChatPreview response

// TODO messages.getConversationMembers response

// TODO messages.getConversations response

// TODO messages.getConversationsById response

// TODO messages.getHistory response

// TODO messages.getHistoryAttachments response

// TODO messages.getImportantMessages response

// TODO messages.getInviteLink response

// TODO messages.getLastActivity response

// TODO messages.getLongPollHistory response

// TODO messages.getLongPollServer response

// TODO messages.isMessagesFromGroupAllowed response

// TODO messages.joinChatByInviteLink response

// TODO messages.markAsAnsweredConversation response

// TODO messages.markAsImportant response

// TODO messages.markAsImportantConversation response

// TODO messages.markAsRead response

// TODO messages.pin response

// TODO messages.removeChatUser response

// TODO messages.restore response

// TODO messages.search response

// TODO messages.searchConversations response

// TODO messages.send response

// TODO messages.setActivity response

// TODO messages.setChatPhoto response

// TODO messages.unpin response

// TODO newsfeed.addBan response

// TODO newsfeed.deleteBan response

// TODO newsfeed.deleteList response

// TODO newsfeed.get response

// TODO newsfeed.getBanned response

// TODO newsfeed.getComments response

// TODO newsfeed.getDiscoverForContestant response

// TODO newsfeed.getLists response

// TODO newsfeed.getMentions response

// TODO newsfeed.getRecommended response

// TODO newsfeed.getSuggestedSources response

// TODO newsfeed.ignoreItemHides response

// TODO newsfeed.saveList response

// TODO newsfeed.search response

// TODO newsfeed.unignoreItem response

// TODO newsfeed.unsubscribe response

// TODO notes.add response

// TODO notes.createComment response

// TODO notes.delete response

// TODO notes.deleteComment response

// TODO notes.edit response

// TODO notes.editComment response

// TODO notes.get response

// TODO notes.getById response

// TODO notes.getComments response

// TODO notes.restoreComment response

// TODO notifications.get response

// TODO notifications.markAsViewed response

// TODO notifications.sendMessage response

// TODO pages.clearCache response

// TODO pages.get response

// TODO pages.getHistory response

// TODO pages.getTitles response

// TODO pages.getVersion response

// TODO pages.parseWiki response

// TODO pages.save response

// TODO pages.saveAccess response

// TODO photos.confirmTag response

// TODO photos.copy response

// TODO photos.createAlbum response

// TODO photos.createComment response

// TODO photos.delete response

// TODO photos.deleteAlbum response

// TODO photos.deleteComment response

// TODO photos.edit response

// TODO photos.editAlbum response

// TODO photos.editComment response

// TODO photos.get response

// TODO photos.getAlbums response

// TODO photos.getAlbumsCount response

// TODO photos.getAll response

// TODO photos.getAllComments response

// TODO photos.getById response

// TODO photos.getChatUploadServer response

// TODO photos.getComments response

// TODO photos.getMarketAlbumUploadServer response

// TODO photos.getMarketUploadServer response

// TODO photos.getMessagesUploadServer response

// TODO photos.getNewTags response

// TODO photos.getOwnerCoverPhotoUploadServer response

// TODO photos.getOwnerPhotoUploadServer response

// TODO photos.getTags response

// TODO photos.getUploadServer response

// TODO photos.getUserPhotos response

// TODO photos.getWallUploadServer response

// TODO photos.makeCover response

// TODO photos.moveMoves response

// TODO photos.putTag response

// TODO photos.removeTag response

// TODO photos.reorderAlbums response

// TODO photos.reorderPhotos response

// TODO photos.report response

// TODO photos.reportComment response

// TODO photos.restore response

// TODO photos.restoreComment response

// TODO photos.save response

// TODO photos.saveMarketAlbum response

// TODO photos.saveMarketPhoto response

// TODO photos.saveMessagesPhoto response

// TODO photos.saveOwnerCoverPhoto response

// TODO photos.saveOwnerPhoto response

// TODO photos.saveWallPhoto response

// TODO photos.search response

// TODO polls.addVote response

// TODO polls.create response

// TODO polls.deleteVote response

// TODO polls.edit response

// TODO polls.getBackgrounds response

// TODO polls.getById response

// TODO polls.getPhotoUploadServer response

// TODO polls.getVoters response

// TODO polls.savePhoto response

// TODO prettyCards.create response

// TODO prettyCards.delete response

// TODO prettyCards.edit response

// TODO prettyCards.get response

// TODO prettyCards.getById response

// TODO prettyCards.getUploadURL response

// TODO search.getHints response

// TODO stats.get response

// TODO stats.getPostReach response

// TODO stats.trackVisitor response

// TODO status.get response

// TODO status.set response

// TODO storage.get response

// TODO storage.getKeys response

// TODO storage.set response

// TODO stories.banOwner response

// TODO stories.delete response

// TODO stories.get response

// TODO stories.getBanned response

// TODO stories.getById response

// TODO stories.getPhotoUploadServer response

// TODO stories.getReplies response

// TODO stories.getStats response

// TODO stories.getVideoUploadServer response

// TODO stories.getViewers response

// TODO stories.hideAllReplies response

// TODO stories.hideReply response

// TODO stories.unbanOwner response

// TODO streaming.getServerUrl response

// TODO streaming.getSettings response

// TODO streaming.getStats response

// TODO streaming.getStem response

// TODO streaming.setSettings response

// TODO users.getFollowers response

// TODO users.getSubscriptions response

// TODO users.isAppUser response

// TODO users.report response

// TODO users.search response

// TODO utils.checkLink response

// TODO utils.deleteFromLastShortened response

// TODO utils.getLastShortenedLinks response

// TODO utils.getLinkStats response

// TODO utils.getServerTime response

// TODO utils.getShortLink response

// TODO video.add response

// TODO video.addAlbum response

// TODO video.addToAlbum response

// TODO video.createComment response

// TODO video.delete response

// TODO video.deleteAlbum response

// TODO video.deleteComment response

// TODO video.edit response

// TODO video.editAlbum response

// TODO video.editComment response

// TODO video.get response

// TODO video.getAlbumById response

// TODO video.getAlbums response

// TODO video.getAlbumsByVideo response

// TODO video.getComments response

// TODO video.removeFromAlbum response

// TODO video.reorderAlbums response

// TODO video.reorderVideos response

// TODO video.report response

// TODO video.reportComment response

// TODO video.restore response

// TODO video.restoreComment response

// TODO video.save response

// TODO video.search response

// TODO wall.closeComments response

// TODO wall.createComment response

// TODO wall.delete response

// TODO wall.deleteComment response

// TODO wall.edit response

// TODO wall.editAdsStealth response

// TODO wall.editComment response

// TODO wall.get response

// TODO wall.getById response

// TODO wall.getComment response

// TODO wall.getComments response

// TODO wall.getReposts response

// TODO wall.openComments response

// TODO wall.pin response

// TODO wall.post response

// TODO wall.postAdsStealth response

// TODO wall.reportComment response

// TODO wall.reportPost response

// TODO wall.repost response

// TODO wall.restore response

// TODO wall.restoreComment response

// TODO wall.search response

// TODO wall.unpin response

// TODO widgets.getComments response

// TODO widgets.getPages response
