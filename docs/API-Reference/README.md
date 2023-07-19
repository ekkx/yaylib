<!-- 参考: https://developer.twitter.com/en/docs/api-reference-index -->

<p align="center">
    <img src="https://github.com/qvco/yaylib/assets/77382767/45c45b21-d812-4cad-8f27-315ffef53201" alt="Logo" height="300px">
    <h3 align="center">API reference index</h3>
    <p align="center">
        <a href="https://github.com/qvco/yaylib/issues">Report Bug</a>
        ·
        <a href="https://github.com/qvco/yaylib/issues">Request Feature</a>
        ·
        <a href="https://discord.gg/MEuBfNtqRN">Join the discord</a>
    </p>
</p>

## Posts

#### Bookmarks

- <a href="">GET /v1/users/:user_id/bookmarks</a> (`getBookmark`)
- <a href="">PUT /v1/users/:user_id/bookmarks/:id</a> (`addBookmark`)
- <a href="">DELETE /v1/users/:user_id/bookmarks/:id</a> (`removeBookmark`)

#### Highlights

- <a href="">GET /v1/groups/:group_id/highlights</a> (`getGroupHighlightPosts`)
- <a href="">PUT /v1/groups/:group_id/highlights/:post_id</a> (`addGroupHighlightPost`)
- <a href="">DELETE /v1/groups/:group_id/highlights/:post_id</a> (`removeGroupHighlightPost`)

#### Pins

- <a href="">POST /v1/pinned/posts</a> (`createPinPost`)
- <a href="">DELETE /v1/pinned/posts/:id</a> (`deletePinPost`)
- <a href="">PUT /v2/posts/group_pinned_post</a> (`createGroupPinPost`)
- <a href="">DELETE /v2/posts/group_pinned_post</a> (`deleteGroupPinPost`)

#### Post Management

- <a href="">POST /v3/posts/new</a> (`createPost`)
- <a href="">POST /v3/posts/repost</a> (`createRepost`)
- <a href="">PUT /v3/posts/:id</a> (`updatePost`)
- <a href="">POST /v1/posts/delete_all_post</a> (`deleteAllPost`)
- <a href="">POST /v2/posts/mass_destroy</a> (`removePosts`)
- <a href="">POST /v3/posts/:post_id/report</a> (`reportPost`)
- <a href="">POST /v1/posts/validate</a> (`validatePost`)
- <a href="">GET /v2/conversations/:id</a> (`getConversation`)

#### Threads

- <a href="">POST /v1/threads/:id/posts</a> (`createThreadPost`)

#### Timelines

- <a href="">GET /v2/posts/call_timeline</a> (`getCallTimeline`)
- <a href="">GET /v2/conversations/root_posts</a> (`getConversationRootPosts`)
- <a href="">GET /v2/posts/call_followers_timeline</a> (`getFollowingCallTimeline`)
- <a href="">GET /v2/posts/following_timeline</a> (`getFollowingTimeline`)
- <a href="">GET /v2/groups/:id/posts/search</a> (`getGroupSearchPosts`)
- <a href="">GET /v2/posts/group_timeline</a> (`getGroupTimeline`)
- <a href="">GET /v2/posts/mine</a> (`getMyPosts`)
- <a href="">GET /v2/posts/:id</a> (`getPost`)
- <a href="">GET /v2/posts/:id/reposts</a> (`getPostReposts`)
- <a href="">GET /v2/posts/multiple</a> (`getPosts`)
- <a href="">POST /v1/posts/recommended_tag</a> (`getRecommendedPostTags`)
- <a href="">GET /v2/posts/recommended_timeline</a> (`getRecommendedPosts`)
- <a href="">GET /v2/posts/:noreply_modetimeline</a> (`getTimeline`)
- <a href="">GET /v2/posts/url_metadata</a> (`getUrlMetadata`)
- <a href="">GET /v2/posts/user_timeline</a> (`getUserTimeline`)

#### Search Posts

- <a href="">GET /v2/posts/search</a> (`getSearchPosts`)
- <a href="">GET /v2/posts/tags/:tag</a> (`getHashtagTimeline`)

#### Likes

- <a href="">GET /v1/posts/:id/likers</a> (`getPostLikers`)
- <a href="">POST /v2/posts/like</a> (`likePosts`)
- <a href="">POST /v1/posts/:id/unlike</a> (`unlikePost`)

#### Miscellaneous

- <a href="">POST /v2/posts/new_conference_call</a> (`createGroupCallPost`)
- <a href="">POST /v2/posts/new_share_post</a> (`createSharePost`)
- <a href="">POST /v2/posts/:id/recommendation_feedback</a> (`updateRecommendationFeedback`)
- <a href="">POST /v1/posts/videos/:videoId/view</a> (`viewVideo`)
- <a href="">POST /v2/surveys/:id/vote</a> (`voteSurvey`)

## Threads

#### Thread Management

- <a href="">POST /v1/threads</a> (`createThread`)
- <a href="">GET /v1/threads</a> (`getGroupThreadList`)
- <a href="">GET /v1/threads/:id</a> (`getThread`)
- <a href="">PUT /v1/threads/:id</a> (`updateThread`)
- <a href="">DELETE /v1/threads/:id</a> (`removeThread`)

#### Join

- <a href="">POST /v1/threads/:thread_id/members/:id</a> (`joinThread`)
- <a href="">DELETE /v1/threads/:thread_id/members/:id</a> (`leaveThread`)

#### Post Management

- <a href="">PUT /v3/posts/:id/move_to_thread/:thread_id</a> (`addPostToThread`)
- <a href="">POST /v3/posts/:id/move_to_thread</a> (`convertPostToThread`)
- <a href="">POST /v1/threads/:id/posts</a> (`getThreadPosts`)

#### User Management

- <a href="">GET /v1/threads/joined_statuses</a> (`getJoinedStatuses`)

## Users

#### Get Users

- <a href="">GET /v2/users/:id</a> (`getUser`)
- <a href="">GET /v1/users/list_id</a> (`getUsers`)
- <a href="">GET /v2/users/info/:id</a> (`getUserWithoutLeavingFootprint`)
- <a href="">GET /v1/users/qr_codes/:qr</a> (`getUserFromQr`)
- <a href="">GET /v2/users/list_uuid/</a> (`getUsersFromUuid`)
- <a href="">GET /v2/users/fresh/:id</a> (`getUserEmail`)
- <a href="">GET /v2/users/social_shared_users</a> (`getSocialSharedUsers`)

#### User Management

- <a href="">GET /v1/users/custom_definitions</a> (`getUserCustomDefinitions`)
- <a href="">GET /v1/users/interests</a> (`getUserInterests`)
- <a href="">POST /v3/users/register</a> (`createUser`)
- <a href="">POST /v2/users/destroy</a> (`destroyUser`)
- <a href="">PUT /v1/users/reset_password</a> (`resetPassword`)
- <a href="">POST /v3/users/:user_id/report</a> (`reportUser`)
- <a href="">POST /v3/users/edit</a> (`updateUser`)
- <a href="">PUT /v1/users/interests</a> (`updateUserInterests`)
- <a href="">PUT /v1/users/language</a> (`updateLanguage`)
- <a href="">POST /v2/users/remove_profile_photo</a> (`removeUserAvatar`)
- <a href="">POST /v2/users/remove_cover_image</a> (`removeUserCover`)
- <a href="">POST /v1/users/reset_counters</a> (`refreshCounter`)
- <a href="">GET /v1/users/reset_counters</a> (`getRefreshCounterRequests`)

#### Follow Management

- <a href="">POST /v2/users/:id/follow</a> (`followUser`)
- <a href="">POST /v2/users/follow</a> (`followUsers`)
- <a href="">GET /v2/users/:id/followers</a> (`getUserFollowers`)
- <a href="">GET /v2/users/:id/list_followings</a> (`getUserFollowings`)
- <a href="">GET /v2/users/follow_requests</a> (`getFollowRequest`)
- <a href="">GET /v2/users/follow_requests_count</a> (`getFollowRequestCount`)
- <a href="">GET /v1/users/following_born_today</a> (`getFollowingUsersBorn`)
- <a href="">POST /v2/users/:target_id/follow_request</a> (`takeActionFollowRequest`)
- <a href="">POST /v2/users/:id/unfollow</a> (`unfollowUser`)
- <a href="">POST /v2/users/edit</a> (`setFollowPermissionEnabled`)
- <a href="">GET /v1/users/active_followings</a> (`getActiveFollowings`)

#### Recommendations

- <a href="">GET /v1/friends</a> (`getFollowRecommendations`)
- <a href="">GET /v1/users/:id/follow_recommended</a> (`getRecommendedUsersToFollowForProfile`)
- <a href="">GET /v1/users/initial_follow_recommended</a> (`getInitialRecommendedUsersToFollow`)

#### Search Users

- <a href="">GET /v1/lobi_friends</a> (`searchLobiUsers`)
- <a href="">GET /v1/users/search</a> (`searchUsers`)

#### Hima Now

- <a href="">GET /v2/users/hima_users</a> (`getHimaUsers`)
- <a href="">POST /v1/users/hima</a> (`turnOnHima`)

#### Footprints

- <a href="">GET /v2/users/footprints</a> (`getFootprints`)
- <a href="">DELETE /v2/users/:user_id/footprints/:footprint_id</a> (`deleteFootprint`)

#### Miscellaneous

- <a href="">POST /v1/users/invite_contact</a> (`updateInviteContactStatus`)
- <a href="">POST /v1/users/contact_friends</a> (`uploadContactsFriends`)
- <a href="">DELETE /v1/users/contact_friends</a> (`deleteContactFriends`)
- <a href="">POST /v1/users/twitter_friends</a> (`uploadTwitterFriendIds`)
- <a href="">GET /v1/users/additonal_notification_setting</a> (`getAdditionalSettings`)
- <a href="">GET /v1/users/:uuid/app_review_status</a> (`getAppReviewStatus`)
- <a href="">POST /v1/users/contact_status</a> (`getContactStatus`)
- <a href="">GET /v1/users/default_settings</a> (`getDefaultSettings`)
- <a href="">GET /v2/users/timestamp</a> (`getTimestamp`)
- <a href="">POST /v2/users/social_shared</a> (`postSocialShared`)
- <a href="">POST /v1/users/:uuid/app_review_status</a> (`recordAppReviewStatus`)
- <a href="">POST /v1/users/additonal_notification_setting</a> (`setAdditionalSettingEnabled`)
- <a href="">POST /v1/users/visible_on_sns_friend_recommendation_setting</a> (`setSettingFollowRecommendationEnabled`)

## Groups

#### Group Management

- <a href="">POST /v3/groups/new</a> (`create`)
- <a href="">GET /v1/groups/:id</a> (`getGroup`)
- <a href="">GET /v2/groups</a> (`getGroups`)
- <a href="">GET /v2/groups/mine</a> (`getMyGroups`)
- <a href="">GET /v1/groups/categories</a> (`getCategories`)
- <a href="">GET /v1/groups/created_quota</a> (`getCreateQuota`)
- <a href="">POST /v1/groups/:id/set_title</a> (`setTitle`)
- <a href="">POST /v1/groups/:id/take_over</a> (`takeoverOwnership`)
- <a href="">POST /v3/groups/:id/update</a> (`update`)

#### Join and Invite

- <a href="">POST /v1/groups/:id/invite</a> (`inviteUsers`)
- <a href="">POST /v1/groups/:id/join</a> (`join`)
- <a href="">DELETE /v1/groups/:id/leave</a> (`leave`)
- <a href="">POST /v1/groups/:id/accept/:user_id</a> (`acceptUserRequest`)
- <a href="">POST /v1/groups/:id/decline/:user_id</a> (`declineUserRequest`)

#### User Management

- <a href="">GET /v1/groups/:id/ban_list</a> (`getBannedMembers`)
- <a href="">GET /v1/groups/:group_id/users/invitable</a> (`getInvitableUsers`)
- <a href="">GET /v1/groups/:id/members/:user_id</a> (`getMember`)
- <a href="">GET /v2/groups/:id/members</a> (`getMembers`)
- <a href="">GET /v1/groups/user_group_list</a> (`getUserGroups`)
- <a href="">POST /v1/groups/:id/ban/:user_id</a> (`addRelatedGroups`)
- <a href="">POST /v1/groups/:id/unban/:user_id</a> (`unbanUser`)
- <a href="">POST /v1/groups/:group_id/fire/:user_id</a> (`removeModerator`)
- <a href="">POST /v1/groups/:id/remove_cover</a> (`removeCover`)

#### Relations

- <a href="">GET /v1/groups/:id/related</a> (`getRelatedGroups`)
- <a href="">GET /v1/groups/:id/relatable</a> (`getRelatableGroups`)
- <a href="">PUT /v1/groups/:id/related</a> (`addRelatedGroups`)
- <a href="">DELETE /v1/groups/:id/related</a> (`removeRelatedGroups`)

#### Administrations

- <a href="">PUT /v1/groups/:id/deputize</a> (`acceptModeratorOffer`)
- <a href="">DELETE /v1/groups/:id/deputize</a> (`declineModeratorOffer`)
- <a href="">PUT /v1/groups/:group_id/deputize/:user_id:/withdraw</a> (`withdrawModeratorOffer`)
- <a href="">POST /v3/groups/:group_id/deputize/mass</a> (`sendModeratorOffers`)
- <a href="">PUT /v1/groups/:id/transfer</a> (`acceptOwnershipOffer`)
- <a href="">PUT /v1/groups/:id/transfer/withdraw</a> (`withdrawOwnershipOffer`)
- <a href="">DELETE /v1/groups/:id/transfer</a> (`declineOwnershipOffer`)
- <a href="">POST /v3/groups/:id/transfer</a> (`sendOwnershipOffer`)

#### Miscellaneous

- <a href="">GET /v1/groups/unread_status</a> (`checkUnreadStatus`)
- <a href="">GET /v1/groups/joined_statuses</a> (`getJoinedStatuses`)
- <a href="">POST /v1/pinned/groups</a> (`createPinGroup`)
- <a href="">DELETE /v1/pinned/groups/:id</a> (`deletePinGroup`)
- <a href="">POST /v2/groups/:group_id/social_shared</a> (`postSocialShared`)
- <a href="">POST /v1/groups/:id/visit</a> (`visit`)
- <a href="">POST /v3/groups/:group_id/report</a> (`report`)

## Chats

#### Chat rooms

- <a href="">POST /v1/chat_rooms/accept_chat_request</a> (`acceptRequest`)
- <a href="">GET /v1/chat_rooms/unread_status</a> (`checkUnreadStatus`)
- <a href="">POST /v3/chat_rooms/new</a> (`createGroup`)
- <a href="">POST /v1/chat_rooms/new</a> (`createPrivate`)
- <a href="">DELETE /v2/chat_rooms/:id/background</a> (`deleteBackground`)
- <a href="">POST /v3/chat_rooms/:id/edit</a> (`edit`)
- <a href="">GET /v1/chat_rooms/main_list</a> (`getMainRooms`)
- <a href="">GET /v1/chat_rooms/request_list</a> (`getRequestRooms`)
- <a href="">GET /v2/chat_rooms/:id</a> (`getRoom`)
- <a href="">POST /v3/chat_rooms/:chat_room_id/report</a> (`report`)

#### Chat messages

- <a href="">DELETE /v1/chat_rooms/:room_id/messages/:message_id/delete</a> (`deleteMessage`)
- <a href="">POST /v3/chat_rooms/:id/messages/new</a> (`sendMessage`)
- <a href="">GET /v2/chat_rooms/:id/messages</a> (`getMessages`)
- <a href="">POST /v1/chat_rooms/:id/attachments_read</a> (`readAttachment`)
- <a href="">POST /v2/chat_rooms/:id/messages/:message_id/read</a> (`readMessage`)
- <a href="">POST /v1/chat_rooms/:id/videos_read</a> (`readVideoMessage`)

#### Chat users

- <a href="">POST /v1/users/followings/chatable</a> (`getChatableUsers`)
- <a href="">POST /v2/chat_rooms/:id/invite</a> (`invite`)
- <a href="">POST /v2/chat_rooms/:id/kick</a> (`kickUsers`)

#### Chat settings

- <a href="">GET /v2/chat_rooms/update</a> (`refreshRooms`)
- <a href="">GET /v2/notification_settings/chat_rooms/:id</a> (`getNotificationSettings`)
- <a href="">POST /v2/notification_settings/chat_rooms/:id</a> (`setNotificationSettings`)
- <a href="">POST /v1/chat_rooms/:id/pinned</a> (`pin`)
- <a href="">DELETE /v1/chat_rooms/:id/pinned</a> (`unpin`)
- <a href="">POST /v1/chat_rooms/mass_destroy</a> (`remove`)

#### Chat notifications

- <a href="">POST /v1/chat_rooms/:id/screen_captured</a> (`sendMediaScreenshotNotification`)

#### Hidden chats

- <a href="">GET /v1/hidden/chats</a> (`getHiddenChatRooms`)
- <a href="">POST /v1/hidden/chats</a> (`hideChat`)
- <a href="">DELETE /v1/hidden/chats</a> (`unHideChat`)

#### Miscellaneous

- <a href="">GET /v2/sticker_packs</a> (`getStickerPacks`)
- <a href="">GET /v1/users/gif_data</a> (`getGifsData`)
- <a href="">GET /v1/chat_rooms/total_chat_request</a> (`getTotalRequests`)

## Reviews

#### Review Management

- <a href="">POST /v2/users/reviews/:id</a> (`createReview`)
- <a href="">POST /v1/users/reviews</a> (`createReview`)
- <a href="">DELETE /v1/users/reviews</a> (`deleteReviews`)

#### Get Reviews

- <a href="">GET /v1/users/reviews/mine</a> (`getMyReviews`)
- <a href="">GET /v2/users/reviews/:id</a> (`getReviews`)

#### Pins

- <a href="">POST /v1/pinned/reviews</a> (`pinReview`)
- <a href="">DELETE /v1/pinned/reviews/:id</a> (`unpinReview`)

## Calls

#### Call Management

- <a href="">PUT /v1/calls/:call_id</a> (`setCall`)
- <a href="">PUT /v1/calls/:call_id/users/:user_id</a> (`setUserRole`)
- <a href="">POST /v1/calls/start_conference_call</a> (`startCall`)
- <a href="">POST /v1/calls/leave_conference_call</a> (`stopCall`)
- <a href="">POST /v1/calls/:call_id/bump</a> (`bumpCall`)
- <a href="">GET /v1/calls/:call_id/users/invitable</a> (`getCallInvitableUsers`)
- <a href="">GET /v1/calls/phone_status/:opponent_id</a> (`getCallStatus`)
- <a href="">GET /v1/calls/bgm</a> (`getBgms`)
- <a href="">GET /v1/games/apps</a> (`getGames`)
- <a href="">GET /v1/genres</a> (`getGenres`)

#### Get Calls

- <a href="">GET /v1/calls/conferences/:call_id</a> (`getCall`)
- <a href="">GET /v1/posts/active_call</a> (`getUserActiveCall`)
- <a href="">GET /v1/posts/group_calls</a> (`getGroupCalls`)

#### User Management

- <a href="">POST /v1/calls/:call_id/bulk_invite</a> (`inviteToCallBulk`)
- <a href="">POST /v1/calls/conference_calls/:call_id/invite</a> (`inviteUsersToCall`)
- <a href="">POST /v1/calls/conference_calls/:call_id/kick</a> (`kickAndBanFromCall`)
- <a href="">POST /v2/calls/invite</a> (`inviteUsersToChatCall`)

#### Miscellaneous

- <a href="">POST /v1/anonymous_calls/leave_agora_channel</a> (`notifyAnonymousUserLeaveAgoraChannel`)
- <a href="">POST /v1/calls/leave_agora_channel</a> (`notifyUserLeaveAgoraChannel`)
- <a href="">PUT /v1/calls/screenshot</a> (`sendCallScreenshot`)

## Login

#### Authentication

- <a href="">POST /api/v1/oauth/token</a> (`getToken`)
- <a href="">POST /api/v1/oauth/token/migrate</a> (`migrateToken`)
- <a href="">POST /v1/users/device_tokens</a> (`revokeTokens`)
- <a href="">POST /v2/users/device_tokens/new</a> (`registerDeviceToken`)
- <a href="">POST /v3/users/login_with_email</a> (`loginWithEmail`)
- <a href="">POST /v3/users/sign_in_with_sns</a> (`loginWithSns`)
- <a href="">POST /v1/users/logout</a> (`logoutDevice`)

#### User Management

- <a href="https://github.com/qvco/yaylib/blob/master/docs/API-Reference/login/change_email.md">PUT /v1/users/change_email</a> (`changeEmail`)
- <a href="">PUT /v1/users/change_password</a> (`changePassword`)
- <a href="">POST /v2/users/sns/connect</a> (`connectAccountWithSns`)
- <a href="">POST /v2/users/sns/disconnect</a> (`disconnectAccountWithSns`)
- <a href="">POST /v2/users/resend_confirm_email</a> (`resendConfirmEmail`)
- <a href="">POST /v2/users/restore</a> (`restoreUser`)
- <a href="">POST /v3/users/login_update</a> (`saveAccountWithEmail`)

## Notifications

- <a href="">GET /api/user_activities</a> (`getUserActivities`)
- <a href="">GET /api/v2/user_activities</a> (`getMergedActivities`)
- <a href="">POST /api/received_push_notifications</a> (`receivedNotification`)

## Miscellaneous

- <a href="">POST /v1/users/policy_agreements/:type</a> (`acceptPolicyAgreement`)
- <a href="">GET /v1/sns_thumbnail/generate</a> (`generateSnsThumbnail`)
- <a href="">POST /v1/email_verification_urls</a> (`getEmailVerificationPresignedUrl`)
- <a href="">GET /v1/buckets/presigned_urls</a> (`getFileUploadPresignedUrls`)
- <a href="">GET /v1/id_check/:model/:action</a> (`getIdCheckerPresignedUrl`)
- <a href="">GET /v1/users/presigned_url</a> (`getOldFileUploadPresignedUrl`)
- <a href="">GET /v1/users/policy_agreements</a> (`getPolicyAgreements`)
- <a href="">GET /v1/promotions</a> (`getPromotions`)
- <a href="">GET /v1/skyfall/url</a> (`getVipGameRewardUrl`)
- <a href="">GET /v1/users/ws_token</a> (`getWebSocketToken`)
- <a href="">POST /v1/genuine_devices/verify</a> (`verifyDevice`)
