// Code generated; DO NOT EDIT.

package oas

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// AcceptChatRequest implements acceptChatRequest operation.
	//
	// POST /v1/chat_rooms/accept_chat_request
	AcceptChatRequest(ctx context.Context, req *AcceptChatRequestReq) error
	// AcceptGroupJoinRequest implements acceptGroupJoinRequest operation.
	//
	// POST /v1/groups/{id}/accept/{userId}
	AcceptGroupJoinRequest(ctx context.Context, params AcceptGroupJoinRequestParams) error
	// AcceptGroupTransfer implements acceptGroupTransfer operation.
	//
	// PUT /v1/groups/{id}/transfer
	AcceptGroupTransfer(ctx context.Context, params AcceptGroupTransferParams) error
	// AddThreadMember implements addThreadMember operation.
	//
	// POST /v1/threads/{thread_id}/members/{id}
	AddThreadMember(ctx context.Context, params AddThreadMemberParams) error
	// AgreePolicy implements agreePolicy operation.
	//
	// POST /v1/users/policy_agreements/{type}
	AgreePolicy(ctx context.Context, params AgreePolicyParams) error
	// BanGroupUser implements banGroupUser operation.
	//
	// POST /v1/groups/{id}/ban/{userId}
	BanGroupUser(ctx context.Context, params BanGroupUserParams) error
	// BlockUser implements blockUser operation.
	//
	// POST /v1/users/{id}/block
	BlockUser(ctx context.Context, params BlockUserParams) error
	// BulkInviteToCall implements bulkInviteToCall operation.
	//
	// POST /v1/calls/{call_id}/bulk_invite
	BulkInviteToCall(ctx context.Context, params BulkInviteToCallParams) error
	// BumpCall implements bumpCall operation.
	//
	// POST /v1/calls/{call_id}/bump
	BumpCall(ctx context.Context, params BumpCallParams) error
	// CancelGroupTransfer implements cancelGroupTransfer operation.
	//
	// DELETE /v1/groups/{id}/transfer
	CancelGroupTransfer(ctx context.Context, params CancelGroupTransferParams) error
	// ChangeEmail implements changeEmail operation.
	//
	// PUT /v1/users/change_email
	ChangeEmail(ctx context.Context, req *ChangeEmailReq) (*LoginUpdateResponse, error)
	// ChangePassword implements changePassword operation.
	//
	// PUT /v1/users/change_password
	ChangePassword(ctx context.Context, req *ChangePasswordReq) (*LoginUpdateResponse, error)
	// CreateBookmark implements createBookmark operation.
	//
	// PUT /v1/users/{user_id}/bookmarks/{id}
	CreateBookmark(ctx context.Context, params CreateBookmarkParams) (*BookmarkPostResponse, error)
	// CreateChatRoom implements createChatRoom operation.
	//
	// POST /v3/chat_rooms/new
	CreateChatRoom(ctx context.Context, req *CreateChatRoomReq) (*CreateChatRoomResponse, error)
	// CreateChatRoomV1 implements createChatRoomV1 operation.
	//
	// POST /v1/chat_rooms/new
	CreateChatRoomV1(ctx context.Context, req *CreateChatRoomV1Req) (*CreateChatRoomResponse, error)
	// CreateConferenceCallPost implements createConferenceCallPost operation.
	//
	// POST /v2/posts/new_conference_call
	CreateConferenceCallPost(ctx context.Context, req *CreateConferenceCallPostReq) (*CreatePostResponse, error)
	// CreateGroup implements createGroup operation.
	//
	// POST /v3/groups/new
	CreateGroup(ctx context.Context, req *CreateGroupReq) (*CreateGroupResponse, error)
	// CreateMuteKeyword implements createMuteKeyword operation.
	//
	// POST /v1/hidden/words
	CreateMuteKeyword(ctx context.Context, req *MuteKeywordRequest) (*CreateMuteKeywordResponse, error)
	// CreatePost implements createPost operation.
	//
	// POST /v3/posts/new
	CreatePost(ctx context.Context, req *CreatePostReq, params CreatePostParams) (*Post, error)
	// CreateReview implements createReview operation.
	//
	// POST /v1/users/reviews
	CreateReview(ctx context.Context, req *CreateReviewReq) error
	// CreateSharePost implements createSharePost operation.
	//
	// POST /v2/posts/new_share_post
	CreateSharePost(ctx context.Context, req *CreateSharePostReq) (*Post, error)
	// CreateThread implements createThread operation.
	//
	// POST /v1/threads
	CreateThread(ctx context.Context, req *CreateGroupThreadRequest) (*ThreadInfo, error)
	// CreateThreadPost implements createThreadPost operation.
	//
	// POST /v1/threads/{id}/posts
	CreateThreadPost(ctx context.Context, req *CreateThreadPostReq, params CreateThreadPostParams) (*Post, error)
	// DeclineGroupJoinRequest implements declineGroupJoinRequest operation.
	//
	// POST /v1/groups/{id}/decline/{userId}
	DeclineGroupJoinRequest(ctx context.Context, params DeclineGroupJoinRequestParams) error
	// DeleteAllPosts implements deleteAllPosts operation.
	//
	// POST /v1/posts/delete_all_post
	DeleteAllPosts(ctx context.Context) error
	// DeleteBookmark implements deleteBookmark operation.
	//
	// DELETE /v1/users/{user_id}/bookmarks/{id}
	DeleteBookmark(ctx context.Context, params DeleteBookmarkParams) error
	// DeleteChatMessage implements deleteChatMessage operation.
	//
	// DELETE /v1/chat_rooms/{room_id}/messages/{message_id}/delete
	DeleteChatMessage(ctx context.Context, params DeleteChatMessageParams) error
	// DeleteChatRooms implements deleteChatRooms operation.
	//
	// POST /v1/chat_rooms/mass_destroy
	DeleteChatRooms(ctx context.Context, req *DeleteChatRoomsReq) error
	// DeleteFootprint implements deleteFootprint operation.
	//
	// DELETE /v2/users/{user_id}/footprints/{footprint_id}
	DeleteFootprint(ctx context.Context, params DeleteFootprintParams) error
	// DeleteMuteKeyword implements deleteMuteKeyword operation.
	//
	// DELETE /v1/hidden/words
	DeleteMuteKeyword(ctx context.Context, params DeleteMuteKeywordParams) error
	// DeleteMyReviews implements deleteMyReviews operation.
	//
	// DELETE /v1/users/reviews
	DeleteMyReviews(ctx context.Context, params DeleteMyReviewsParams) error
	// DeletePosts implements deletePosts operation.
	//
	// POST /v2/posts/mass_destroy
	DeletePosts(ctx context.Context, req *DeletePostsReq) error
	// DeleteThread implements deleteThread operation.
	//
	// DELETE /v1/threads/{id}
	DeleteThread(ctx context.Context, params DeleteThreadParams) error
	// DeputizeGroupUsers implements deputizeGroupUsers operation.
	//
	// PUT /v1/groups/{id}/deputize
	DeputizeGroupUsers(ctx context.Context, params DeputizeGroupUsersParams) error
	// DeputizeGroupUsersMass implements deputizeGroupUsersMass operation.
	//
	// POST /v3/groups/{group_id}/deputize/mass
	DeputizeGroupUsersMass(ctx context.Context, req *DeputizeGroupUsersMassReq, params DeputizeGroupUsersMassParams) error
	// DisableTwoFactorAuth implements disableTwoFactorAuth operation.
	//
	// PUT /v1/users/secret/disable
	DisableTwoFactorAuth(ctx context.Context, req *DisableTwoFactorAuthReq) error
	// EditUser implements editUser operation.
	//
	// POST /v3/users/edit
	EditUser(ctx context.Context, req *EditUserReq) error
	// EditUserV2 implements editUserV2 operation.
	//
	// POST /v2/users/edit
	EditUserV2(ctx context.Context, req *EditUserV2Req) error
	// EnableTwoFactorAuth implements enableTwoFactorAuth operation.
	//
	// PUT /v1/users/secret/enable
	EnableTwoFactorAuth(ctx context.Context, req *EnableTwoFactorAuthReq) (*TwoStepAuthEnabledResponse, error)
	// FireGroupUser implements fireGroupUser operation.
	//
	// POST /v1/groups/{group_id}/fire/{user_id}
	FireGroupUser(ctx context.Context, params FireGroupUserParams) error
	// FollowUser implements followUser operation.
	//
	// POST /v2/users/{id}/follow
	FollowUser(ctx context.Context, params FollowUserParams) error
	// FollowUsers implements followUsers operation.
	//
	// POST /v2/users/follow
	FollowUsers(ctx context.Context, params FollowUsersParams) error
	// GenerateCallActionSignature implements generateCallActionSignature operation.
	//
	// POST /v1/calls/action_signature/generate
	GenerateCallActionSignature(ctx context.Context, req *GenerateCallActionSignatureReq) (*CallActionSignatureResponse, error)
	// GetActiveCallPost implements getActiveCallPost operation.
	//
	// GET /v1/posts/active_call
	GetActiveCallPost(ctx context.Context, params GetActiveCallPostParams) (*PostResponse, error)
	// GetActiveFollowings implements getActiveFollowings operation.
	//
	// GET /v1/users/active_followings
	GetActiveFollowings(ctx context.Context, params GetActiveFollowingsParams) (*ActiveFollowingsResponse, error)
	// GetAdditionalNotificationSetting implements getAdditionalNotificationSetting operation.
	//
	// GET /v1/users/additonal_notification_setting
	GetAdditionalNotificationSetting(ctx context.Context) (*AdditionalSettingsResponse, error)
	// GetAgoraRtmToken implements getAgoraRtmToken operation.
	//
	// GET /v3/calls/{call_id}/agora_rtm_token
	GetAgoraRtmToken(ctx context.Context, params GetAgoraRtmTokenParams) (*RtmTokenResponse, error)
	// GetAppConfig implements getAppConfig operation.
	//
	// GET /api/apps/{app}
	GetAppConfig(ctx context.Context, params GetAppConfigParams) (*ApplicationConfigResponse, error)
	// GetBannedWords implements getBannedWords operation.
	//
	// GET /{countryApiValue}/api/v2/banned_words
	GetBannedWords(ctx context.Context, params GetBannedWordsParams) (*BanWordsResponse, error)
	// GetBlockedUserIds implements getBlockedUserIds operation.
	//
	// GET /v1/users/block_ids
	GetBlockedUserIds(ctx context.Context) (*BlockedUserIdsResponse, error)
	// GetBlockedUsers implements getBlockedUsers operation.
	//
	// POST /v2/users/blocked
	GetBlockedUsers(ctx context.Context, req *SearchUsersRequest, params GetBlockedUsersParams) (*BlockedUsersResponse, error)
	// GetBookmarkedPosts implements getBookmarkedPosts operation.
	//
	// GET /v1/users/{user_id}/bookmarks
	GetBookmarkedPosts(ctx context.Context, params GetBookmarkedPostsParams) (*PostsResponse, error)
	// GetBucketPresignedUrls implements getBucketPresignedUrls operation.
	//
	// GET /v1/buckets/presigned_urls
	GetBucketPresignedUrls(ctx context.Context, params GetBucketPresignedUrlsParams) (*PresignedUrlsResponse, error)
	// GetCallBgms implements getCallBgms operation.
	//
	// GET /v1/calls/bgm
	GetCallBgms(ctx context.Context) (*BgmsResponse, error)
	// GetCallFollowersTimeline implements getCallFollowersTimeline operation.
	//
	// GET /v2/posts/call_followers_timeline
	GetCallFollowersTimeline(ctx context.Context, params GetCallFollowersTimelineParams) (*PostsResponse, error)
	// GetCallGiftHistory implements getCallGiftHistory operation.
	//
	// GET /v1/calls/{call_id}/gift_transactions
	GetCallGiftHistory(ctx context.Context, params GetCallGiftHistoryParams) (*CallGiftHistoryResponse, error)
	// GetCallTimeline implements getCallTimeline operation.
	//
	// GET /v2/posts/call_timeline
	GetCallTimeline(ctx context.Context, params GetCallTimelineParams) (*PostsResponse, error)
	// GetChatMessages implements getChatMessages operation.
	//
	// GET /v2/chat_rooms/{id}/messages
	GetChatMessages(ctx context.Context, params GetChatMessagesParams) (*MessagesResponse, error)
	// GetChatRequestCount implements getChatRequestCount operation.
	//
	// GET /v1/chat_rooms/total_chat_request
	GetChatRequestCount(ctx context.Context) (*TotalChatRequestResponse, error)
	// GetChatRequests implements getChatRequests operation.
	//
	// GET /v1/chat_rooms/request_list
	GetChatRequests(ctx context.Context, params GetChatRequestsParams) (*ChatRoomsResponse, error)
	// GetChatRoom implements getChatRoom operation.
	//
	// GET /v2/chat_rooms/{id}
	GetChatRoom(ctx context.Context, params GetChatRoomParams) (*ChatRoomResponse, error)
	// GetChatUnreadStatus implements getChatUnreadStatus operation.
	//
	// GET /v1/chat_rooms/unread_status
	GetChatUnreadStatus(ctx context.Context, params GetChatUnreadStatusParams) (*UnreadStatusResponse, error)
	// GetChatableFollowings implements getChatableFollowings operation.
	//
	// POST /v1/users/followings/chatable
	GetChatableFollowings(ctx context.Context, req *SearchUsersRequest, params GetChatableFollowingsParams) (*FollowUsersResponse, error)
	// GetConferenceCall implements getConferenceCall operation.
	//
	// GET /v2/calls/conferences/{call_id}
	GetConferenceCall(ctx context.Context, params GetConferenceCallParams) (*ConferenceCallResponse, error)
	// GetConversation implements getConversation operation.
	//
	// GET /v2/conversations/{id}
	GetConversation(ctx context.Context, params GetConversationParams) (*PostsResponse, error)
	// GetDefaultSettings implements getDefaultSettings operation.
	//
	// GET /v1/users/default_settings
	GetDefaultSettings(ctx context.Context) (*DefaultSettingsResponse, error)
	// GetFollowRequestCount implements getFollowRequestCount operation.
	//
	// GET /v2/users/follow_requests_count
	GetFollowRequestCount(ctx context.Context) (*FollowRequestCountResponse, error)
	// GetFollowRequests implements getFollowRequests operation.
	//
	// GET /v2/users/follow_requests
	GetFollowRequests(ctx context.Context, params GetFollowRequestsParams) (*UsersByTimestampResponse, error)
	// GetFollowingTimeline implements getFollowingTimeline operation.
	//
	// GET /v2/posts/following_timeline
	GetFollowingTimeline(ctx context.Context, params GetFollowingTimelineParams) (*PostsResponse, error)
	// GetFollowingsBornToday implements getFollowingsBornToday operation.
	//
	// GET /v1/users/following_born_today
	GetFollowingsBornToday(ctx context.Context, params GetFollowingsBornTodayParams) (*UsersResponse, error)
	// GetFootprints implements getFootprints operation.
	//
	// GET /v3/users/footprints
	GetFootprints(ctx context.Context, params GetFootprintsParams) (*FootprintsResponse, error)
	// GetFreshUser implements getFreshUser operation.
	//
	// GET /v2/users/fresh/{id}
	GetFreshUser(ctx context.Context, params GetFreshUserParams) (*UserResponse, error)
	// GetGroup implements getGroup operation.
	//
	// GET /v1/groups/{id}
	GetGroup(ctx context.Context, params GetGroupParams) (*GroupResponse, error)
	// GetGroupBanList implements getGroupBanList operation.
	//
	// GET /v1/groups/{id}/ban_list
	GetGroupBanList(ctx context.Context, params GetGroupBanListParams) (*UsersResponse, error)
	// GetGroupCreateQuota implements getGroupCreateQuota operation.
	//
	// GET /v1/groups/created_quota
	GetGroupCreateQuota(ctx context.Context) (*CreateQuotaResponse, error)
	// GetGroupGiftHistory implements getGroupGiftHistory operation.
	//
	// GET /v1/groups/{group_id}/gift_history
	GetGroupGiftHistory(ctx context.Context, params GetGroupGiftHistoryParams) (*GroupGiftHistoryResponse, error)
	// GetGroupGiftTransactions implements getGroupGiftTransactions operation.
	//
	// GET /v1/groups/{group_id}/gift_transactions
	GetGroupGiftTransactions(ctx context.Context, params GetGroupGiftTransactionsParams) (*GiftTransactionsResponse, error)
	// GetGroupHighlights implements getGroupHighlights operation.
	//
	// GET /v1/groups/{group_id}/highlights
	GetGroupHighlights(ctx context.Context, params GetGroupHighlightsParams) (*PostsResponse, error)
	// GetGroupMember implements getGroupMember operation.
	//
	// GET /v1/groups/{id}/members/{userId}
	GetGroupMember(ctx context.Context, params GetGroupMemberParams) (*GroupUserResponse, error)
	// GetGroupMembers implements getGroupMembers operation.
	//
	// GET /v2/groups/{id}/members
	GetGroupMembers(ctx context.Context, params GetGroupMembersParams) (*GroupUsersResponse, error)
	// GetGroupNotificationSettings implements getGroupNotificationSettings operation.
	//
	// GET /v2/notification_settings/groups/{id}
	GetGroupNotificationSettings(ctx context.Context, params GetGroupNotificationSettingsParams) (*GroupNotificationSettingsResponse, error)
	// GetGroupReceivedGiftSenders implements getGroupReceivedGiftSenders operation.
	//
	// GET /v1/groups/{group_id}/received_gifts/{gift_id}/senders
	GetGroupReceivedGiftSenders(ctx context.Context, params GetGroupReceivedGiftSendersParams) (*GiftSendersResponse, error)
	// GetGroupTimeline implements getGroupTimeline operation.
	//
	// GET /v2/posts/group_timeline
	GetGroupTimeline(ctx context.Context, params GetGroupTimelineParams) (*PostsResponse, error)
	// GetGroupUnreadStatus implements getGroupUnreadStatus operation.
	//
	// GET /v1/groups/unread_status
	GetGroupUnreadStatus(ctx context.Context, params GetGroupUnreadStatusParams) (*UnreadStatusResponse, error)
	// GetHimaUsers implements getHimaUsers operation.
	//
	// GET /v2/users/hima_users
	GetHimaUsers(ctx context.Context, params GetHimaUsersParams) (*HimaUsersResponse, error)
	// GetInvitableCallUsers implements getInvitableCallUsers operation.
	//
	// GET /v1/calls/{call_id}/users/invitable
	GetInvitableCallUsers(ctx context.Context, params GetInvitableCallUsersParams) (*UsersByTimestampResponse, error)
	// GetInvitableGroupUsers implements getInvitableGroupUsers operation.
	//
	// GET /v1/groups/{group_id}/users/invitable
	GetInvitableGroupUsers(ctx context.Context, params GetInvitableGroupUsersParams) (*UsersByTimestampResponse, error)
	// GetJoinedGroupStatuses implements getJoinedGroupStatuses operation.
	//
	// GET /v1/groups/joined_statuses
	GetJoinedGroupStatuses(ctx context.Context, params GetJoinedGroupStatusesParams) (GetJoinedGroupStatusesOK, error)
	// GetJoinedThreadStatuses implements getJoinedThreadStatuses operation.
	//
	// GET /v1/threads/joined_statuses
	GetJoinedThreadStatuses(ctx context.Context, params GetJoinedThreadStatusesParams) (GetJoinedThreadStatusesOK, error)
	// GetMainChatRooms implements getMainChatRooms operation.
	//
	// GET /v1/chat_rooms/main_list
	GetMainChatRooms(ctx context.Context, params GetMainChatRoomsParams) (*ChatRoomsResponse, error)
	// GetMyPosts implements getMyPosts operation.
	//
	// GET /v2/posts/mine
	GetMyPosts(ctx context.Context, params GetMyPostsParams) (*PostsResponse, error)
	// GetMyReviews implements getMyReviews operation.
	//
	// GET /v1/users/reviews/mine
	GetMyReviews(ctx context.Context, params GetMyReviewsParams) (*ReviewsResponse, error)
	// GetPhoneStatus implements getPhoneStatus operation.
	//
	// GET /v1/calls/phone_status/{opponent_id}
	GetPhoneStatus(ctx context.Context, params GetPhoneStatusParams) (*CallStatusResponse, error)
	// GetPolicyAgreements implements getPolicyAgreements operation.
	//
	// GET /v1/users/policy_agreements
	GetPolicyAgreements(ctx context.Context) (*PolicyAgreementsResponse, error)
	// GetPopularWords implements getPopularWords operation.
	//
	// GET /{countryApiValue}/api/apps/{app}/popular_words
	GetPopularWords(ctx context.Context, params GetPopularWordsParams) (*PopularWordsResponse, error)
	// GetPost implements getPost operation.
	//
	// GET /v2/posts/{id}
	GetPost(ctx context.Context, params GetPostParams) (*PostResponse, error)
	// GetPostGiftTransactions implements getPostGiftTransactions operation.
	//
	// GET /v1/posts/{post_id}/gift_transactions
	GetPostGiftTransactions(ctx context.Context, params GetPostGiftTransactionsParams) (*GiftTransactionsResponse, error)
	// GetPostLikers implements getPostLikers operation.
	//
	// GET /v1/posts/{id}/likers
	GetPostLikers(ctx context.Context, params GetPostLikersParams) (*PostLikersResponse, error)
	// GetPostReposts implements getPostReposts operation.
	//
	// GET /v2/posts/{id}/reposts
	GetPostReposts(ctx context.Context, params GetPostRepostsParams) (*PostsResponse, error)
	// GetPostUrlMetadata implements getPostUrlMetadata operation.
	//
	// GET /v2/posts/url_metadata
	GetPostUrlMetadata(ctx context.Context, params GetPostUrlMetadataParams) (*SharedUrl, error)
	// GetPostsByIds implements getPostsByIds operation.
	//
	// GET /v2/posts/multiple
	GetPostsByIds(ctx context.Context, params GetPostsByIdsParams) (*PostsResponse, error)
	// GetPostsByTag implements getPostsByTag operation.
	//
	// GET /v2/posts/tags/{tag}
	GetPostsByTag(ctx context.Context, params GetPostsByTagParams) (*PostsResponse, error)
	// GetReceivedGiftSenders implements getReceivedGiftSenders operation.
	//
	// GET /v1/received_gifts/{gift_id}/senders
	GetReceivedGiftSenders(ctx context.Context, params GetReceivedGiftSendersParams) (*GiftSendersResponse, error)
	// GetReceivedGiftTransaction implements getReceivedGiftTransaction operation.
	//
	// GET /v1/received_gifts/{gift_transaction_uuid}
	GetReceivedGiftTransaction(ctx context.Context, params GetReceivedGiftTransactionParams) (*GiftReceivedTransactionResponse, error)
	// GetRecentEngagementPosts implements getRecentEngagementPosts operation.
	//
	// GET /v2/posts/recent_engagement
	GetRecentEngagementPosts(ctx context.Context, params GetRecentEngagementPostsParams) (*PostsResponse, error)
	// GetRecommendedFollowUsers implements getRecommendedFollowUsers operation.
	//
	// GET /v1/users/{id}/follow_recommended
	GetRecommendedFollowUsers(ctx context.Context, params GetRecommendedFollowUsersParams) (*UsersResponse, error)
	// GetRecommendedPostTags implements getRecommendedPostTags operation.
	//
	// POST /v1/posts/recommended_tag
	GetRecommendedPostTags(ctx context.Context, req *GetRecommendedPostTagsReq) (*PostTagsResponse, error)
	// GetRecommendedTimeline implements getRecommendedTimeline operation.
	//
	// GET /v2/posts/recommended_timeline
	GetRecommendedTimeline(ctx context.Context, params GetRecommendedTimelineParams) (*PostsResponse, error)
	// GetRelatableGroups implements getRelatableGroups operation.
	//
	// GET /v1/groups/{id}/relatable
	GetRelatableGroups(ctx context.Context, params GetRelatableGroupsParams) (*GroupsRelatedResponse, error)
	// GetRelatedGroups implements getRelatedGroups operation.
	//
	// GET /v1/groups/{id}/related
	GetRelatedGroups(ctx context.Context, params GetRelatedGroupsParams) (*GroupsRelatedResponse, error)
	// GetResetCounters implements getResetCounters operation.
	//
	// GET /v1/users/reset_counters
	GetResetCounters(ctx context.Context) (*RefreshCounterRequestsResponse, error)
	// GetRootPosts implements getRootPosts operation.
	//
	// GET /v2/conversations/root_posts
	GetRootPosts(ctx context.Context, params GetRootPostsParams) (*PostsResponse, error)
	// GetThread implements getThread operation.
	//
	// GET /v1/threads/{id}
	GetThread(ctx context.Context, params GetThreadParams) (*ThreadInfo, error)
	// GetThreadPosts implements getThreadPosts operation.
	//
	// GET /v1/threads/{id}/posts
	GetThreadPosts(ctx context.Context, params GetThreadPostsParams) (*PostsResponse, error)
	// GetTimeline implements getTimeline operation.
	//
	// GET /v2/posts/{noreply_mode}timeline
	GetTimeline(ctx context.Context, params GetTimelineParams) (*PostsResponse, error)
	// GetTwoFactorAuthRequestInfo implements getTwoFactorAuthRequestInfo operation.
	//
	// GET /v1/users/secret
	GetTwoFactorAuthRequestInfo(ctx context.Context) (*TwoStepAuthRequestInfoResponse, error)
	// GetTwoFactorAuthStatus implements getTwoFactorAuthStatus operation.
	//
	// GET /v1/users/secret/status
	GetTwoFactorAuthStatus(ctx context.Context) (*TwoFAStatusResponse, error)
	// GetUpdatedChatRooms implements getUpdatedChatRooms operation.
	//
	// GET /v2/chat_rooms/update
	GetUpdatedChatRooms(ctx context.Context, params GetUpdatedChatRoomsParams) (*ChatRoomsResponse, error)
	// GetUser implements getUser operation.
	//
	// GET /v2/users/{id}
	GetUser(ctx context.Context, params GetUserParams) (*UserResponse, error)
	// GetUserActivities implements getUserActivities operation.
	//
	// GET /api/v2/user_activities
	GetUserActivities(ctx context.Context, params GetUserActivitiesParams) (*ActivitiesResponse, error)
	// GetUserActivitiesV1 implements getUserActivitiesV1 operation.
	//
	// GET /api/user_activities
	GetUserActivitiesV1(ctx context.Context, params GetUserActivitiesV1Params) (*ActivitiesResponse, error)
	// GetUserByQr implements getUserByQr operation.
	//
	// GET /v1/users/qr_codes/{qr}
	GetUserByQr(ctx context.Context, params GetUserByQrParams) (*UserResponse, error)
	// GetUserCustomDefinitions implements getUserCustomDefinitions operation.
	//
	// GET /v1/users/custom_definitions
	GetUserCustomDefinitions(ctx context.Context) (*UserCustomDefinitionsResponse, error)
	// GetUserFollowers implements getUserFollowers operation.
	//
	// GET /v3/users/{id}/followers
	GetUserFollowers(ctx context.Context, params GetUserFollowersParams) (*FollowUsersResponse, error)
	// GetUserFollowings implements getUserFollowings operation.
	//
	// GET /v3/users/{id}/followings
	GetUserFollowings(ctx context.Context, params GetUserFollowingsParams) (*FollowUsersResponse, error)
	// GetUserGiftTransactions implements getUserGiftTransactions operation.
	//
	// GET /v1/users/{user_id}/gift_transactions
	GetUserGiftTransactions(ctx context.Context, params GetUserGiftTransactionsParams) (*GiftTransactionsResponse, error)
	// GetUserGroupList implements getUserGroupList operation.
	//
	// GET /v1/groups/user_group_list
	GetUserGroupList(ctx context.Context, params GetUserGroupListParams) (*GroupsResponse, error)
	// GetUserInfo implements getUserInfo operation.
	//
	// GET /v2/users/info/{id}
	GetUserInfo(ctx context.Context, params GetUserInfoParams) (*UserResponse, error)
	// GetUserInterests implements getUserInterests operation.
	//
	// GET /v1/users/interests
	GetUserInterests(ctx context.Context) (*UserInterestsResponse, error)
	// GetUserPresignedUrl implements getUserPresignedUrl operation.
	//
	// GET /v1/users/presigned_url
	GetUserPresignedUrl(ctx context.Context, params GetUserPresignedUrlParams) (*PresignedUrlResponse, error)
	// GetUserReviews implements getUserReviews operation.
	//
	// GET /v2/users/reviews/{id}
	GetUserReviews(ctx context.Context, params GetUserReviewsParams) (*ReviewsResponse, error)
	// GetUserTimeline implements getUserTimeline operation.
	//
	// GET /v2/posts/user_timeline
	GetUserTimeline(ctx context.Context, params GetUserTimelineParams) (*PostsResponse, error)
	// GetUserTimestamp implements getUserTimestamp operation.
	//
	// GET /v2/users/timestamp
	GetUserTimestamp(ctx context.Context) (*UserTimestampResponse, error)
	// GetUsersByIds implements getUsersByIds operation.
	//
	// GET /v1/users/list_id
	GetUsersByIds(ctx context.Context, params GetUsersByIdsParams) (*UsersResponse, error)
	// GetWebSocketToken implements getWebSocketToken operation.
	//
	// GET /v1/users/ws_token
	GetWebSocketToken(ctx context.Context) (*WebSocketTokenResponse, error)
	// HideChats implements hideChats operation.
	//
	// POST /v1/hidden/chats
	HideChats(ctx context.Context, req *HideChatsReq) error
	// HideUsers implements hideUsers operation.
	//
	// POST /v1/hidden/users
	HideUsers(ctx context.Context, req *HideUsersReq) error
	// InviteToCall implements inviteToCall operation.
	//
	// POST /v2/calls/invite
	InviteToCall(ctx context.Context, req *InviteToCallReq) error
	// InviteToChatRoom implements inviteToChatRoom operation.
	//
	// POST /v2/chat_rooms/{id}/invite
	InviteToChatRoom(ctx context.Context, req *InviteToChatRoomReq, params InviteToChatRoomParams) error
	// InviteToConferenceCall implements inviteToConferenceCall operation.
	//
	// POST /v1/calls/conference_calls/{call_id}/invite
	InviteToConferenceCall(ctx context.Context, req *InviteToConferenceCallReq, params InviteToConferenceCallParams) error
	// InviteToGroup implements inviteToGroup operation.
	//
	// POST /v1/groups/{id}/invite
	InviteToGroup(ctx context.Context, req *InviteToGroupReq, params InviteToGroupParams) error
	// JoinGroup implements joinGroup operation.
	//
	// POST /v1/groups/{id}/join
	JoinGroup(ctx context.Context, params JoinGroupParams) error
	// KickFromChatRoom implements kickFromChatRoom operation.
	//
	// POST /v2/chat_rooms/{id}/kick
	KickFromChatRoom(ctx context.Context, req *KickFromChatRoomReq, params KickFromChatRoomParams) error
	// KickFromConferenceCall implements kickFromConferenceCall operation.
	//
	// POST /v3/calls/conference_calls/{call_id}/kick
	KickFromConferenceCall(ctx context.Context, req *KickFromConferenceCallReq, params KickFromConferenceCallParams) (*CallActionSignatureResponse, error)
	// LeaveAgoraChannel implements leaveAgoraChannel operation.
	//
	// POST /v1/calls/leave_agora_channel
	LeaveAgoraChannel(ctx context.Context, req *LeaveAgoraChannelReq) error
	// LeaveConferenceCall implements leaveConferenceCall operation.
	//
	// POST /v1/calls/leave_conference_call
	LeaveConferenceCall(ctx context.Context, req *LeaveConferenceCallReq) error
	// LeaveGroup implements leaveGroup operation.
	//
	// DELETE /v1/groups/{id}/leave
	LeaveGroup(ctx context.Context, params LeaveGroupParams) error
	// LikePosts implements likePosts operation.
	//
	// POST /v2/posts/like
	LikePosts(ctx context.Context, req *LikePostsReq) (*LikePostsResponse, error)
	// ListGameApps implements listGameApps operation.
	//
	// GET /v1/games/apps
	ListGameApps(ctx context.Context, params ListGameAppsParams) (*GamesResponse, error)
	// ListGameWalkthroughs implements listGameWalkthroughs operation.
	//
	// GET /v1/games/apps/{app_id}/walkthroughs
	ListGameWalkthroughs(ctx context.Context, params ListGameWalkthroughsParams) ([]Walkthrough, error)
	// ListGenres implements listGenres operation.
	//
	// GET /v1/genres
	ListGenres(ctx context.Context, params ListGenresParams) (*GenresResponse, error)
	// ListGifts implements listGifts operation.
	//
	// GET /v2/gifts
	ListGifts(ctx context.Context, params ListGiftsParams) (*GiftsResponse, error)
	// ListGroupCategories implements listGroupCategories operation.
	//
	// GET /v1/groups/categories
	ListGroupCategories(ctx context.Context, params ListGroupCategoriesParams) (*GroupCategoriesResponse, error)
	// ListGroups implements listGroups operation.
	//
	// GET /v2/groups
	ListGroups(ctx context.Context, params ListGroupsParams) (*GroupsResponse, error)
	// ListHiddenChats implements listHiddenChats operation.
	//
	// GET /v1/hidden/chats
	ListHiddenChats(ctx context.Context, params ListHiddenChatsParams) (*ChatRoomsResponse, error)
	// ListHiddenUsers implements listHiddenUsers operation.
	//
	// GET /v1/hidden/users
	ListHiddenUsers(ctx context.Context, params ListHiddenUsersParams) (*HiddenResponse, error)
	// ListMuteKeywords implements listMuteKeywords operation.
	//
	// GET /v1/hidden/words
	ListMuteKeywords(ctx context.Context) (*MuteKeywordResponse, error)
	// ListMutedGroupUsers implements listMutedGroupUsers operation.
	//
	// GET /v1/group_mute/{id}/muted_users
	ListMutedGroupUsers(ctx context.Context, params ListMutedGroupUsersParams) (*GroupMuteUsersResponse, error)
	// ListMyGroups implements listMyGroups operation.
	//
	// GET /v2/groups/mine
	ListMyGroups(ctx context.Context, params ListMyGroupsParams) (*GroupsResponse, error)
	// ListReceivedGifts implements listReceivedGifts operation.
	//
	// GET /v2/received_gifts
	ListReceivedGifts(ctx context.Context) (*GiftReceivedResponse, error)
	// ListReceivedGiftsV1 implements listReceivedGiftsV1 operation.
	//
	// GET /v1/received_gifts
	ListReceivedGiftsV1(ctx context.Context, params ListReceivedGiftsV1Params) (*GiftReceivedResponse, error)
	// ListStickerPacks implements listStickerPacks operation.
	//
	// GET /v2/sticker_packs
	ListStickerPacks(ctx context.Context) (*StickerPacksResponse, error)
	// ListThreads implements listThreads operation.
	//
	// GET /v1/threads
	ListThreads(ctx context.Context, params ListThreadsParams) (*GroupThreadListResponse, error)
	// LoginWithEmail implements loginWithEmail operation.
	//
	// POST /v3/users/login_with_email
	LoginWithEmail(ctx context.Context, req *LoginEmailUserRequest) (*LoginUserResponse, error)
	// Logout implements logout operation.
	//
	// POST /v1/users/logout
	Logout(ctx context.Context, req *LogoutReq) error
	// MovePostToSpecificThread implements movePostToSpecificThread operation.
	//
	// PUT /v3/posts/{id}/move_to_thread/{thread_id}
	MovePostToSpecificThread(ctx context.Context, params MovePostToSpecificThreadParams) (*ThreadInfo, error)
	// MovePostToThread implements movePostToThread operation.
	//
	// POST /v3/posts/{id}/move_to_thread
	MovePostToThread(ctx context.Context, req *MovePostToThreadReq, params MovePostToThreadParams) (*ThreadInfo, error)
	// MuteGroupUser implements muteGroupUser operation.
	//
	// POST /v1/group_mute/{id}/mute/{user_id}
	MuteGroupUser(ctx context.Context, params MuteGroupUserParams) error
	// OauthToken implements oauthToken operation.
	//
	// POST /api/v1/oauth/token
	OauthToken(ctx context.Context, req *OauthTokenReq) (*TokenResponse, error)
	// PinChatRoom implements pinChatRoom operation.
	//
	// POST /v1/chat_rooms/{id}/pinned
	PinChatRoom(ctx context.Context, params PinChatRoomParams) error
	// PinGroup implements pinGroup operation.
	//
	// POST /v1/pinned/groups
	PinGroup(ctx context.Context, req *PinGroupReq) error
	// PinGroupHighlightPost implements pinGroupHighlightPost operation.
	//
	// PUT /v1/groups/{group_id}/highlights/{post_id}
	PinGroupHighlightPost(ctx context.Context, params PinGroupHighlightPostParams) error
	// PinGroupPost implements pinGroupPost operation.
	//
	// PUT /v2/posts/group_pinned_post
	PinGroupPost(ctx context.Context, req *PinGroupPostReq) error
	// PinPost implements pinPost operation.
	//
	// POST /v1/pinned/posts
	PinPost(ctx context.Context, req *PinPostReq) error
	// PinReview implements pinReview operation.
	//
	// POST /v1/pinned/reviews
	PinReview(ctx context.Context, req *PinReviewReq) error
	// PingAlive implements pingAlive operation.
	//
	// POST /v1/users/alive
	PingAlive(ctx context.Context) error
	// ReadChatAttachments implements readChatAttachments operation.
	//
	// POST /v1/chat_rooms/{id}/attachments_read
	ReadChatAttachments(ctx context.Context, req *ReadAttachmentRequest, params ReadChatAttachmentsParams) error
	// ReadChatMessage implements readChatMessage operation.
	//
	// POST /v2/chat_rooms/{id}/messages/{message_id}/read
	ReadChatMessage(ctx context.Context, params ReadChatMessageParams) error
	// ReadChatVideos implements readChatVideos operation.
	//
	// POST /v1/chat_rooms/{id}/videos_read
	ReadChatVideos(ctx context.Context, req *ReadChatVideosReq, params ReadChatVideosParams) error
	// RemoveChatRoomBackground implements removeChatRoomBackground operation.
	//
	// DELETE /v2/chat_rooms/{id}/background
	RemoveChatRoomBackground(ctx context.Context, params RemoveChatRoomBackgroundParams) error
	// RemoveCoverImage implements removeCoverImage operation.
	//
	// POST /v2/users/remove_cover_image
	RemoveCoverImage(ctx context.Context) error
	// RemoveGroupCover implements removeGroupCover operation.
	//
	// DELETE /v3/groups/{id}/cover
	RemoveGroupCover(ctx context.Context, params RemoveGroupCoverParams) error
	// RemoveGroupDeputies implements removeGroupDeputies operation.
	//
	// DELETE /v1/groups/{id}/deputize
	RemoveGroupDeputies(ctx context.Context, params RemoveGroupDeputiesParams) error
	// RemoveGroupIcon implements removeGroupIcon operation.
	//
	// DELETE /v3/groups/{id}/icon
	RemoveGroupIcon(ctx context.Context, params RemoveGroupIconParams) error
	// RemoveProfilePhoto implements removeProfilePhoto operation.
	//
	// POST /v2/users/remove_profile_photo
	RemoveProfilePhoto(ctx context.Context) error
	// RemoveRelatedGroups implements removeRelatedGroups operation.
	//
	// DELETE /v1/groups/{id}/related
	RemoveRelatedGroups(ctx context.Context, params RemoveRelatedGroupsParams) error
	// RemoveThreadMember implements removeThreadMember operation.
	//
	// DELETE /v1/threads/{thread_id}/members/{id}
	RemoveThreadMember(ctx context.Context, params RemoveThreadMemberParams) error
	// ReplyToReview implements replyToReview operation.
	//
	// POST /v2/users/reviews/{id}
	ReplyToReview(ctx context.Context, req *ReplyToReviewReq, params ReplyToReviewParams) error
	// ReportChatRoom implements reportChatRoom operation.
	//
	// POST /v3/chat_rooms/{chat_room_id}/report
	ReportChatRoom(ctx context.Context, req *ReportChatRoomReq, params ReportChatRoomParams) error
	// ReportGroup implements reportGroup operation.
	//
	// POST /v3/groups/{group_id}/report
	ReportGroup(ctx context.Context, req *ReportGroupReq, params ReportGroupParams) error
	// ReportPost implements reportPost operation.
	//
	// POST /v3/posts/{post_id}/report
	ReportPost(ctx context.Context, req *ReportPostReq, params ReportPostParams) error
	// ReportUser implements reportUser operation.
	//
	// POST /v3/users/{user_id}/report
	ReportUser(ctx context.Context, req *ReportUserReq, params ReportUserParams) error
	// Repost implements repost operation.
	//
	// POST /v3/posts/repost
	Repost(ctx context.Context, req *RepostReq, params RepostParams) (*CreatePostResponse, error)
	// RequestEmailVerification implements requestEmailVerification operation.
	//
	// POST /v1/email_verification_urls
	RequestEmailVerification(ctx context.Context, req *RequestEmailVerificationReq) (*CommonUrlResponse, error)
	// RequestFollow implements requestFollow operation.
	//
	// POST /v2/users/{target_id}/follow_request
	RequestFollow(ctx context.Context, req *RequestFollowReq, params RequestFollowParams) error
	// RequestGroupWalkthrough implements requestGroupWalkthrough operation.
	//
	// POST /v1/groups/{id}/request_walkthrough
	RequestGroupWalkthrough(ctx context.Context, params RequestGroupWalkthroughParams) error
	// ResendConfirmEmail implements resendConfirmEmail operation.
	//
	// POST /v2/users/resend_confirm_email
	ResendConfirmEmail(ctx context.Context) error
	// ResetCounters implements resetCounters operation.
	//
	// POST /v1/users/reset_counters
	ResetCounters(ctx context.Context, req *ResetCountersReq) error
	// ResetPassword implements resetPassword operation.
	//
	// PUT /v1/users/reset_password
	ResetPassword(ctx context.Context, req *ResetPasswordReq) error
	// SearchGroupPosts implements searchGroupPosts operation.
	//
	// GET /v2/groups/{id}/posts/search
	SearchGroupPosts(ctx context.Context, params SearchGroupPostsParams) (*PostsResponse, error)
	// SearchPosts implements searchPosts operation.
	//
	// GET /v2/posts/search
	SearchPosts(ctx context.Context, params SearchPostsParams) (*PostsResponse, error)
	// SearchUsers implements searchUsers operation.
	//
	// GET /v1/users/search
	SearchUsers(ctx context.Context, params SearchUsersParams) (*UsersResponse, error)
	// SendChatMessage implements sendChatMessage operation.
	//
	// POST /v3/chat_rooms/{id}/messages/new
	SendChatMessage(ctx context.Context, req *SendChatMessageReq, params SendChatMessageParams) (*MessageResponse, error)
	// SetGroupTitle implements setGroupTitle operation.
	//
	// POST /v1/groups/{id}/set_title
	SetGroupTitle(ctx context.Context, req *SetGroupTitleReq, params SetGroupTitleParams) error
	// SetHima implements setHima operation.
	//
	// POST /v1/users/hima
	SetHima(ctx context.Context) error
	// StartConferenceCall implements startConferenceCall operation.
	//
	// POST /v2/calls/start_conference_call
	StartConferenceCall(ctx context.Context, req *StartConferenceCallReq) (*ConferenceCallResponse, error)
	// TakeOverGroup implements takeOverGroup operation.
	//
	// POST /v1/groups/{id}/take_over
	TakeOverGroup(ctx context.Context, params TakeOverGroupParams) error
	// TransferGroup implements transferGroup operation.
	//
	// POST /v3/groups/{id}/transfer
	TransferGroup(ctx context.Context, req *TransferGroupReq, params TransferGroupParams) error
	// UnbanGroupUser implements unbanGroupUser operation.
	//
	// POST /v1/groups/{id}/unban/{userId}
	UnbanGroupUser(ctx context.Context, params UnbanGroupUserParams) error
	// UnblockUser implements unblockUser operation.
	//
	// POST /v2/users/{id}/unblock
	UnblockUser(ctx context.Context, params UnblockUserParams) error
	// UnfollowUser implements unfollowUser operation.
	//
	// POST /v2/users/{id}/unfollow
	UnfollowUser(ctx context.Context, params UnfollowUserParams) error
	// UnhideChats implements unhideChats operation.
	//
	// DELETE /v1/hidden/chats
	UnhideChats(ctx context.Context, params UnhideChatsParams) error
	// UnhideUsers implements unhideUsers operation.
	//
	// DELETE /v1/hidden/users
	UnhideUsers(ctx context.Context, params UnhideUsersParams) error
	// UnlikePost implements unlikePost operation.
	//
	// POST /v1/posts/{id}/unlike
	UnlikePost(ctx context.Context, params UnlikePostParams) error
	// UnmuteGroupUser implements unmuteGroupUser operation.
	//
	// DELETE /v1/group_mute/{id}/unmute/{user_id}
	UnmuteGroupUser(ctx context.Context, params UnmuteGroupUserParams) error
	// UnpinChatRoom implements unpinChatRoom operation.
	//
	// DELETE /v1/chat_rooms/{id}/pinned
	UnpinChatRoom(ctx context.Context, params UnpinChatRoomParams) error
	// UnpinGroup implements unpinGroup operation.
	//
	// DELETE /v1/pinned/groups/{id}
	UnpinGroup(ctx context.Context, params UnpinGroupParams) error
	// UnpinGroupHighlightPost implements unpinGroupHighlightPost operation.
	//
	// DELETE /v1/groups/{group_id}/highlights/{post_id}
	UnpinGroupHighlightPost(ctx context.Context, params UnpinGroupHighlightPostParams) error
	// UnpinGroupPost implements unpinGroupPost operation.
	//
	// DELETE /v2/posts/group_pinned_post
	UnpinGroupPost(ctx context.Context, params UnpinGroupPostParams) error
	// UnpinReview implements unpinReview operation.
	//
	// DELETE /v1/pinned/reviews/{id}
	UnpinReview(ctx context.Context, params UnpinReviewParams) error
	// UpdateAdditionalNotificationSetting implements updateAdditionalNotificationSetting operation.
	//
	// POST /v1/users/additonal_notification_setting
	UpdateAdditionalNotificationSetting(ctx context.Context, req *UpdateAdditionalNotificationSettingReq) error
	// UpdateCall implements updateCall operation.
	//
	// PUT /v1/calls/{call_id}
	UpdateCall(ctx context.Context, req *UpdateCallReq, params UpdateCallParams) error
	// UpdateCallUser implements updateCallUser operation.
	//
	// PUT /v1/calls/{call_id}/users/{user_id}
	UpdateCallUser(ctx context.Context, req *UpdateCallUserReq, params UpdateCallUserParams) error
	// UpdateChatRoom implements updateChatRoom operation.
	//
	// POST /v3/chat_rooms/{id}/edit
	UpdateChatRoom(ctx context.Context, req *UpdateChatRoomReq, params UpdateChatRoomParams) error
	// UpdateChatRoomNotificationSettings implements updateChatRoomNotificationSettings operation.
	//
	// POST /v2/notification_settings/chat_rooms/{id}
	UpdateChatRoomNotificationSettings(ctx context.Context, req *UpdateChatRoomNotificationSettingsReq, params UpdateChatRoomNotificationSettingsParams) (*NotificationSettingResponse, error)
	// UpdateGroup implements updateGroup operation.
	//
	// POST /v3/groups/{id}/update
	UpdateGroup(ctx context.Context, req *UpdateGroupReq, params UpdateGroupParams) (*GroupResponse, error)
	// UpdateGroupNotificationSettings implements updateGroupNotificationSettings operation.
	//
	// POST /v2/notification_settings/groups/{id}
	UpdateGroupNotificationSettings(ctx context.Context, req *UpdateGroupNotificationSettingsReq, params UpdateGroupNotificationSettingsParams) (*AdditionalSettingsResponse, error)
	// UpdateLanguage implements updateLanguage operation.
	//
	// PUT /v1/users/language
	UpdateLanguage(ctx context.Context, req *UpdateLanguageReq) error
	// UpdateLogin implements updateLogin operation.
	//
	// POST /v3/users/login_update
	UpdateLogin(ctx context.Context, req *UpdateLoginReq) (*LoginUpdateResponse, error)
	// UpdatePost implements updatePost operation.
	//
	// PUT /v3/posts/{id}
	UpdatePost(ctx context.Context, req *UpdatePostReq, params UpdatePostParams) (*Post, error)
	// UpdateRelatedGroups implements updateRelatedGroups operation.
	//
	// PUT /v1/groups/{id}/related
	UpdateRelatedGroups(ctx context.Context, params UpdateRelatedGroupsParams) error
	// UpdateThread implements updateThread operation.
	//
	// PUT /v1/threads/{id}
	UpdateThread(ctx context.Context, req *UpdateThreadReq, params UpdateThreadParams) error
	// UpdateUserInterests implements updateUserInterests operation.
	//
	// PUT /v1/users/interests
	UpdateUserInterests(ctx context.Context, req *CommonIdsRequest) error
	// ValidateCallActionSignature implements validateCallActionSignature operation.
	//
	// GET /v2/calls/action_signature/validate
	ValidateCallActionSignature(ctx context.Context, params ValidateCallActionSignatureParams) error
	// ValidatePost implements validatePost operation.
	//
	// POST /v1/posts/validate
	ValidatePost(ctx context.Context, req *ValidatePostReq) (*ValidationPostResponse, error)
	// ViewPostVideo implements viewPostVideo operation.
	//
	// POST /v1/posts/videos/{videoId}/view
	ViewPostVideo(ctx context.Context, params ViewPostVideoParams) error
	// VisitGroup implements visitGroup operation.
	//
	// POST /v1/groups/{id}/visit
	VisitGroup(ctx context.Context, params VisitGroupParams) error
	// VoteSurvey implements voteSurvey operation.
	//
	// POST /v2/surveys/{id}/vote
	VoteSurvey(ctx context.Context, req *VoteSurveyReq, params VoteSurveyParams) (*VoteSurveyResponse, error)
	// WithdrawGroupDeputy implements withdrawGroupDeputy operation.
	//
	// PUT /v1/groups/{group_id}/deputize/{user_id}/withdraw
	WithdrawGroupDeputy(ctx context.Context, params WithdrawGroupDeputyParams) error
	// WithdrawGroupTransfer implements withdrawGroupTransfer operation.
	//
	// PUT /v1/groups/{id}/transfer/withdraw
	WithdrawGroupTransfer(ctx context.Context, req *WithdrawGroupTransferReq, params WithdrawGroupTransferParams) error
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
