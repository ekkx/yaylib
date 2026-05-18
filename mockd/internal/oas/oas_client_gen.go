// Code generated; DO NOT EDIT.

package oas

import (
	"context"
	"net/url"
	"strings"
	"time"

	"github.com/go-faster/errors"
	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/otelogen"
	"github.com/ogen-go/ogen/uri"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/metric"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
	"go.opentelemetry.io/otel/trace"
)

func trimTrailingSlashes(u *url.URL) {
	u.Path = strings.TrimRight(u.Path, "/")
	u.RawPath = strings.TrimRight(u.RawPath, "/")
}

// Invoker invokes operations described by OpenAPI v3 specification.
type Invoker interface {
	// AcceptChatRequest invokes acceptChatRequest operation.
	//
	// POST /v1/chat_rooms/accept_chat_request
	AcceptChatRequest(ctx context.Context, request *AcceptChatRequestReq) error
	// AcceptGroupJoinRequest invokes acceptGroupJoinRequest operation.
	//
	// POST /v1/groups/{id}/accept/{userId}
	AcceptGroupJoinRequest(ctx context.Context, params AcceptGroupJoinRequestParams) error
	// AcceptGroupTransfer invokes acceptGroupTransfer operation.
	//
	// PUT /v1/groups/{id}/transfer
	AcceptGroupTransfer(ctx context.Context, params AcceptGroupTransferParams) error
	// AddThreadMember invokes addThreadMember operation.
	//
	// POST /v1/threads/{thread_id}/members/{id}
	AddThreadMember(ctx context.Context, params AddThreadMemberParams) error
	// AgreePolicy invokes agreePolicy operation.
	//
	// POST /v1/users/policy_agreements/{type}
	AgreePolicy(ctx context.Context, params AgreePolicyParams) error
	// BanGroupUser invokes banGroupUser operation.
	//
	// POST /v1/groups/{id}/ban/{userId}
	BanGroupUser(ctx context.Context, params BanGroupUserParams) error
	// BlockUser invokes blockUser operation.
	//
	// POST /v1/users/{id}/block
	BlockUser(ctx context.Context, params BlockUserParams) error
	// BulkInviteToCall invokes bulkInviteToCall operation.
	//
	// POST /v1/calls/{call_id}/bulk_invite
	BulkInviteToCall(ctx context.Context, params BulkInviteToCallParams) error
	// BumpCall invokes bumpCall operation.
	//
	// POST /v1/calls/{call_id}/bump
	BumpCall(ctx context.Context, params BumpCallParams) error
	// CancelGroupTransfer invokes cancelGroupTransfer operation.
	//
	// DELETE /v1/groups/{id}/transfer
	CancelGroupTransfer(ctx context.Context, params CancelGroupTransferParams) error
	// ChangeEmail invokes changeEmail operation.
	//
	// PUT /v1/users/change_email
	ChangeEmail(ctx context.Context, request *ChangeEmailReq) (*LoginUpdateResponse, error)
	// ChangePassword invokes changePassword operation.
	//
	// PUT /v1/users/change_password
	ChangePassword(ctx context.Context, request *ChangePasswordReq) (*LoginUpdateResponse, error)
	// CreateBookmark invokes createBookmark operation.
	//
	// PUT /v1/users/{user_id}/bookmarks/{id}
	CreateBookmark(ctx context.Context, params CreateBookmarkParams) (*BookmarkPostResponse, error)
	// CreateChatRoom invokes createChatRoom operation.
	//
	// POST /v3/chat_rooms/new
	CreateChatRoom(ctx context.Context, request *CreateChatRoomReq) (*CreateChatRoomResponse, error)
	// CreateChatRoomV1 invokes createChatRoomV1 operation.
	//
	// POST /v1/chat_rooms/new
	CreateChatRoomV1(ctx context.Context, request *CreateChatRoomV1Req) (*CreateChatRoomResponse, error)
	// CreateConferenceCallPost invokes createConferenceCallPost operation.
	//
	// POST /v2/posts/new_conference_call
	CreateConferenceCallPost(ctx context.Context, request *CreateConferenceCallPostReq) (*CreatePostResponse, error)
	// CreateGroup invokes createGroup operation.
	//
	// POST /v3/groups/new
	CreateGroup(ctx context.Context, request *CreateGroupReq) (*CreateGroupResponse, error)
	// CreateMuteKeyword invokes createMuteKeyword operation.
	//
	// POST /v1/hidden/words
	CreateMuteKeyword(ctx context.Context, request *MuteKeywordRequest) (*CreateMuteKeywordResponse, error)
	// CreatePost invokes createPost operation.
	//
	// POST /v3/posts/new
	CreatePost(ctx context.Context, request *CreatePostReq, params CreatePostParams) (*Post, error)
	// CreateReview invokes createReview operation.
	//
	// POST /v1/users/reviews
	CreateReview(ctx context.Context, request *CreateReviewReq) error
	// CreateSharePost invokes createSharePost operation.
	//
	// POST /v2/posts/new_share_post
	CreateSharePost(ctx context.Context, request *CreateSharePostReq) (*Post, error)
	// CreateThread invokes createThread operation.
	//
	// POST /v1/threads
	CreateThread(ctx context.Context, request *CreateGroupThreadRequest) (*ThreadInfo, error)
	// CreateThreadPost invokes createThreadPost operation.
	//
	// POST /v1/threads/{id}/posts
	CreateThreadPost(ctx context.Context, request *CreateThreadPostReq, params CreateThreadPostParams) (*Post, error)
	// DeclineGroupJoinRequest invokes declineGroupJoinRequest operation.
	//
	// POST /v1/groups/{id}/decline/{userId}
	DeclineGroupJoinRequest(ctx context.Context, params DeclineGroupJoinRequestParams) error
	// DeleteAllPosts invokes deleteAllPosts operation.
	//
	// POST /v1/posts/delete_all_post
	DeleteAllPosts(ctx context.Context) error
	// DeleteBookmark invokes deleteBookmark operation.
	//
	// DELETE /v1/users/{user_id}/bookmarks/{id}
	DeleteBookmark(ctx context.Context, params DeleteBookmarkParams) error
	// DeleteChatMessage invokes deleteChatMessage operation.
	//
	// DELETE /v1/chat_rooms/{room_id}/messages/{message_id}/delete
	DeleteChatMessage(ctx context.Context, params DeleteChatMessageParams) error
	// DeleteChatRooms invokes deleteChatRooms operation.
	//
	// POST /v1/chat_rooms/mass_destroy
	DeleteChatRooms(ctx context.Context, request *DeleteChatRoomsReq) error
	// DeleteFootprint invokes deleteFootprint operation.
	//
	// DELETE /v2/users/{user_id}/footprints/{footprint_id}
	DeleteFootprint(ctx context.Context, params DeleteFootprintParams) error
	// DeleteMuteKeyword invokes deleteMuteKeyword operation.
	//
	// DELETE /v1/hidden/words
	DeleteMuteKeyword(ctx context.Context, params DeleteMuteKeywordParams) error
	// DeleteMyReviews invokes deleteMyReviews operation.
	//
	// DELETE /v1/users/reviews
	DeleteMyReviews(ctx context.Context, params DeleteMyReviewsParams) error
	// DeletePosts invokes deletePosts operation.
	//
	// POST /v2/posts/mass_destroy
	DeletePosts(ctx context.Context, request *DeletePostsReq) error
	// DeleteThread invokes deleteThread operation.
	//
	// DELETE /v1/threads/{id}
	DeleteThread(ctx context.Context, params DeleteThreadParams) error
	// DeputizeGroupUsers invokes deputizeGroupUsers operation.
	//
	// PUT /v1/groups/{id}/deputize
	DeputizeGroupUsers(ctx context.Context, params DeputizeGroupUsersParams) error
	// DeputizeGroupUsersMass invokes deputizeGroupUsersMass operation.
	//
	// POST /v3/groups/{group_id}/deputize/mass
	DeputizeGroupUsersMass(ctx context.Context, request *DeputizeGroupUsersMassReq, params DeputizeGroupUsersMassParams) error
	// DisableTwoFactorAuth invokes disableTwoFactorAuth operation.
	//
	// PUT /v1/users/secret/disable
	DisableTwoFactorAuth(ctx context.Context, request *DisableTwoFactorAuthReq) error
	// EditUser invokes editUser operation.
	//
	// POST /v3/users/edit
	EditUser(ctx context.Context, request *EditUserReq) error
	// EditUserV2 invokes editUserV2 operation.
	//
	// POST /v2/users/edit
	EditUserV2(ctx context.Context, request *EditUserV2Req) error
	// EnableTwoFactorAuth invokes enableTwoFactorAuth operation.
	//
	// PUT /v1/users/secret/enable
	EnableTwoFactorAuth(ctx context.Context, request *EnableTwoFactorAuthReq) (*TwoStepAuthEnabledResponse, error)
	// FireGroupUser invokes fireGroupUser operation.
	//
	// POST /v1/groups/{group_id}/fire/{user_id}
	FireGroupUser(ctx context.Context, params FireGroupUserParams) error
	// FollowUser invokes followUser operation.
	//
	// POST /v2/users/{id}/follow
	FollowUser(ctx context.Context, params FollowUserParams) error
	// FollowUsers invokes followUsers operation.
	//
	// POST /v2/users/follow
	FollowUsers(ctx context.Context, params FollowUsersParams) error
	// GenerateCallActionSignature invokes generateCallActionSignature operation.
	//
	// POST /v1/calls/action_signature/generate
	GenerateCallActionSignature(ctx context.Context, request *GenerateCallActionSignatureReq) (*CallActionSignatureResponse, error)
	// GetActiveCallPost invokes getActiveCallPost operation.
	//
	// GET /v1/posts/active_call
	GetActiveCallPost(ctx context.Context, params GetActiveCallPostParams) (*PostResponse, error)
	// GetActiveFollowings invokes getActiveFollowings operation.
	//
	// GET /v1/users/active_followings
	GetActiveFollowings(ctx context.Context, params GetActiveFollowingsParams) (*ActiveFollowingsResponse, error)
	// GetAdditionalNotificationSetting invokes getAdditionalNotificationSetting operation.
	//
	// GET /v1/users/additonal_notification_setting
	GetAdditionalNotificationSetting(ctx context.Context) (*AdditionalSettingsResponse, error)
	// GetAgoraRtmToken invokes getAgoraRtmToken operation.
	//
	// GET /v3/calls/{call_id}/agora_rtm_token
	GetAgoraRtmToken(ctx context.Context, params GetAgoraRtmTokenParams) (*RtmTokenResponse, error)
	// GetAppConfig invokes getAppConfig operation.
	//
	// GET /api/apps/{app}
	GetAppConfig(ctx context.Context, params GetAppConfigParams) (*ApplicationConfigResponse, error)
	// GetBannedWords invokes getBannedWords operation.
	//
	// GET /{countryApiValue}/api/v2/banned_words
	GetBannedWords(ctx context.Context, params GetBannedWordsParams) (*BanWordsResponse, error)
	// GetBlockedUserIds invokes getBlockedUserIds operation.
	//
	// GET /v1/users/block_ids
	GetBlockedUserIds(ctx context.Context) (*BlockedUserIdsResponse, error)
	// GetBlockedUsers invokes getBlockedUsers operation.
	//
	// POST /v2/users/blocked
	GetBlockedUsers(ctx context.Context, request *SearchUsersRequest, params GetBlockedUsersParams) (*BlockedUsersResponse, error)
	// GetBookmarkedPosts invokes getBookmarkedPosts operation.
	//
	// GET /v1/users/{user_id}/bookmarks
	GetBookmarkedPosts(ctx context.Context, params GetBookmarkedPostsParams) (*PostsResponse, error)
	// GetBucketPresignedUrls invokes getBucketPresignedUrls operation.
	//
	// GET /v1/buckets/presigned_urls
	GetBucketPresignedUrls(ctx context.Context, params GetBucketPresignedUrlsParams) (*PresignedUrlsResponse, error)
	// GetCallBgms invokes getCallBgms operation.
	//
	// GET /v1/calls/bgm
	GetCallBgms(ctx context.Context) (*BgmsResponse, error)
	// GetCallFollowersTimeline invokes getCallFollowersTimeline operation.
	//
	// GET /v2/posts/call_followers_timeline
	GetCallFollowersTimeline(ctx context.Context, params GetCallFollowersTimelineParams) (*PostsResponse, error)
	// GetCallGiftHistory invokes getCallGiftHistory operation.
	//
	// GET /v1/calls/{call_id}/gift_transactions
	GetCallGiftHistory(ctx context.Context, params GetCallGiftHistoryParams) (*CallGiftHistoryResponse, error)
	// GetCallTimeline invokes getCallTimeline operation.
	//
	// GET /v2/posts/call_timeline
	GetCallTimeline(ctx context.Context, params GetCallTimelineParams) (*PostsResponse, error)
	// GetChatMessages invokes getChatMessages operation.
	//
	// GET /v2/chat_rooms/{id}/messages
	GetChatMessages(ctx context.Context, params GetChatMessagesParams) (*MessagesResponse, error)
	// GetChatRequestCount invokes getChatRequestCount operation.
	//
	// GET /v1/chat_rooms/total_chat_request
	GetChatRequestCount(ctx context.Context) (*TotalChatRequestResponse, error)
	// GetChatRequests invokes getChatRequests operation.
	//
	// GET /v1/chat_rooms/request_list
	GetChatRequests(ctx context.Context, params GetChatRequestsParams) (*ChatRoomsResponse, error)
	// GetChatRoom invokes getChatRoom operation.
	//
	// GET /v2/chat_rooms/{id}
	GetChatRoom(ctx context.Context, params GetChatRoomParams) (*ChatRoomResponse, error)
	// GetChatUnreadStatus invokes getChatUnreadStatus operation.
	//
	// GET /v1/chat_rooms/unread_status
	GetChatUnreadStatus(ctx context.Context, params GetChatUnreadStatusParams) (*UnreadStatusResponse, error)
	// GetChatableFollowings invokes getChatableFollowings operation.
	//
	// POST /v1/users/followings/chatable
	GetChatableFollowings(ctx context.Context, request *SearchUsersRequest, params GetChatableFollowingsParams) (*FollowUsersResponse, error)
	// GetConferenceCall invokes getConferenceCall operation.
	//
	// GET /v2/calls/conferences/{call_id}
	GetConferenceCall(ctx context.Context, params GetConferenceCallParams) (*ConferenceCallResponse, error)
	// GetConversation invokes getConversation operation.
	//
	// GET /v2/conversations/{id}
	GetConversation(ctx context.Context, params GetConversationParams) (*PostsResponse, error)
	// GetDefaultSettings invokes getDefaultSettings operation.
	//
	// GET /v1/users/default_settings
	GetDefaultSettings(ctx context.Context) (*DefaultSettingsResponse, error)
	// GetFollowRequestCount invokes getFollowRequestCount operation.
	//
	// GET /v2/users/follow_requests_count
	GetFollowRequestCount(ctx context.Context) (*FollowRequestCountResponse, error)
	// GetFollowRequests invokes getFollowRequests operation.
	//
	// GET /v2/users/follow_requests
	GetFollowRequests(ctx context.Context, params GetFollowRequestsParams) (*UsersByTimestampResponse, error)
	// GetFollowingTimeline invokes getFollowingTimeline operation.
	//
	// GET /v2/posts/following_timeline
	GetFollowingTimeline(ctx context.Context, params GetFollowingTimelineParams) (*PostsResponse, error)
	// GetFollowingsBornToday invokes getFollowingsBornToday operation.
	//
	// GET /v1/users/following_born_today
	GetFollowingsBornToday(ctx context.Context, params GetFollowingsBornTodayParams) (*UsersResponse, error)
	// GetFootprints invokes getFootprints operation.
	//
	// GET /v3/users/footprints
	GetFootprints(ctx context.Context, params GetFootprintsParams) (*FootprintsResponse, error)
	// GetFreshUser invokes getFreshUser operation.
	//
	// GET /v2/users/fresh/{id}
	GetFreshUser(ctx context.Context, params GetFreshUserParams) (*UserResponse, error)
	// GetGroup invokes getGroup operation.
	//
	// GET /v1/groups/{id}
	GetGroup(ctx context.Context, params GetGroupParams) (*GroupResponse, error)
	// GetGroupBanList invokes getGroupBanList operation.
	//
	// GET /v1/groups/{id}/ban_list
	GetGroupBanList(ctx context.Context, params GetGroupBanListParams) (*UsersResponse, error)
	// GetGroupCreateQuota invokes getGroupCreateQuota operation.
	//
	// GET /v1/groups/created_quota
	GetGroupCreateQuota(ctx context.Context) (*CreateQuotaResponse, error)
	// GetGroupGiftHistory invokes getGroupGiftHistory operation.
	//
	// GET /v1/groups/{group_id}/gift_history
	GetGroupGiftHistory(ctx context.Context, params GetGroupGiftHistoryParams) (*GroupGiftHistoryResponse, error)
	// GetGroupGiftTransactions invokes getGroupGiftTransactions operation.
	//
	// GET /v1/groups/{group_id}/gift_transactions
	GetGroupGiftTransactions(ctx context.Context, params GetGroupGiftTransactionsParams) (*GiftTransactionsResponse, error)
	// GetGroupHighlights invokes getGroupHighlights operation.
	//
	// GET /v1/groups/{group_id}/highlights
	GetGroupHighlights(ctx context.Context, params GetGroupHighlightsParams) (*PostsResponse, error)
	// GetGroupMember invokes getGroupMember operation.
	//
	// GET /v1/groups/{id}/members/{userId}
	GetGroupMember(ctx context.Context, params GetGroupMemberParams) (*GroupUserResponse, error)
	// GetGroupMembers invokes getGroupMembers operation.
	//
	// GET /v2/groups/{id}/members
	GetGroupMembers(ctx context.Context, params GetGroupMembersParams) (*GroupUsersResponse, error)
	// GetGroupNotificationSettings invokes getGroupNotificationSettings operation.
	//
	// GET /v2/notification_settings/groups/{id}
	GetGroupNotificationSettings(ctx context.Context, params GetGroupNotificationSettingsParams) (*GroupNotificationSettingsResponse, error)
	// GetGroupReceivedGiftSenders invokes getGroupReceivedGiftSenders operation.
	//
	// GET /v1/groups/{group_id}/received_gifts/{gift_id}/senders
	GetGroupReceivedGiftSenders(ctx context.Context, params GetGroupReceivedGiftSendersParams) (*GiftSendersResponse, error)
	// GetGroupTimeline invokes getGroupTimeline operation.
	//
	// GET /v2/posts/group_timeline
	GetGroupTimeline(ctx context.Context, params GetGroupTimelineParams) (*PostsResponse, error)
	// GetGroupUnreadStatus invokes getGroupUnreadStatus operation.
	//
	// GET /v1/groups/unread_status
	GetGroupUnreadStatus(ctx context.Context, params GetGroupUnreadStatusParams) (*UnreadStatusResponse, error)
	// GetHimaUsers invokes getHimaUsers operation.
	//
	// GET /v2/users/hima_users
	GetHimaUsers(ctx context.Context, params GetHimaUsersParams) (*HimaUsersResponse, error)
	// GetInvitableCallUsers invokes getInvitableCallUsers operation.
	//
	// GET /v1/calls/{call_id}/users/invitable
	GetInvitableCallUsers(ctx context.Context, params GetInvitableCallUsersParams) (*UsersByTimestampResponse, error)
	// GetInvitableGroupUsers invokes getInvitableGroupUsers operation.
	//
	// GET /v1/groups/{group_id}/users/invitable
	GetInvitableGroupUsers(ctx context.Context, params GetInvitableGroupUsersParams) (*UsersByTimestampResponse, error)
	// GetJoinedGroupStatuses invokes getJoinedGroupStatuses operation.
	//
	// GET /v1/groups/joined_statuses
	GetJoinedGroupStatuses(ctx context.Context, params GetJoinedGroupStatusesParams) (GetJoinedGroupStatusesOK, error)
	// GetJoinedThreadStatuses invokes getJoinedThreadStatuses operation.
	//
	// GET /v1/threads/joined_statuses
	GetJoinedThreadStatuses(ctx context.Context, params GetJoinedThreadStatusesParams) (GetJoinedThreadStatusesOK, error)
	// GetMainChatRooms invokes getMainChatRooms operation.
	//
	// GET /v1/chat_rooms/main_list
	GetMainChatRooms(ctx context.Context, params GetMainChatRoomsParams) (*ChatRoomsResponse, error)
	// GetMyPosts invokes getMyPosts operation.
	//
	// GET /v2/posts/mine
	GetMyPosts(ctx context.Context, params GetMyPostsParams) (*PostsResponse, error)
	// GetMyReviews invokes getMyReviews operation.
	//
	// GET /v1/users/reviews/mine
	GetMyReviews(ctx context.Context, params GetMyReviewsParams) (*ReviewsResponse, error)
	// GetPhoneStatus invokes getPhoneStatus operation.
	//
	// GET /v1/calls/phone_status/{opponent_id}
	GetPhoneStatus(ctx context.Context, params GetPhoneStatusParams) (*CallStatusResponse, error)
	// GetPolicyAgreements invokes getPolicyAgreements operation.
	//
	// GET /v1/users/policy_agreements
	GetPolicyAgreements(ctx context.Context) (*PolicyAgreementsResponse, error)
	// GetPopularWords invokes getPopularWords operation.
	//
	// GET /{countryApiValue}/api/apps/{app}/popular_words
	GetPopularWords(ctx context.Context, params GetPopularWordsParams) (*PopularWordsResponse, error)
	// GetPost invokes getPost operation.
	//
	// GET /v2/posts/{id}
	GetPost(ctx context.Context, params GetPostParams) (*PostResponse, error)
	// GetPostGiftTransactions invokes getPostGiftTransactions operation.
	//
	// GET /v1/posts/{post_id}/gift_transactions
	GetPostGiftTransactions(ctx context.Context, params GetPostGiftTransactionsParams) (*GiftTransactionsResponse, error)
	// GetPostLikers invokes getPostLikers operation.
	//
	// GET /v1/posts/{id}/likers
	GetPostLikers(ctx context.Context, params GetPostLikersParams) (*PostLikersResponse, error)
	// GetPostReposts invokes getPostReposts operation.
	//
	// GET /v2/posts/{id}/reposts
	GetPostReposts(ctx context.Context, params GetPostRepostsParams) (*PostsResponse, error)
	// GetPostUrlMetadata invokes getPostUrlMetadata operation.
	//
	// GET /v2/posts/url_metadata
	GetPostUrlMetadata(ctx context.Context, params GetPostUrlMetadataParams) (*SharedUrl, error)
	// GetPostsByIds invokes getPostsByIds operation.
	//
	// GET /v2/posts/multiple
	GetPostsByIds(ctx context.Context, params GetPostsByIdsParams) (*PostsResponse, error)
	// GetPostsByTag invokes getPostsByTag operation.
	//
	// GET /v2/posts/tags/{tag}
	GetPostsByTag(ctx context.Context, params GetPostsByTagParams) (*PostsResponse, error)
	// GetReceivedGiftSenders invokes getReceivedGiftSenders operation.
	//
	// GET /v1/received_gifts/{gift_id}/senders
	GetReceivedGiftSenders(ctx context.Context, params GetReceivedGiftSendersParams) (*GiftSendersResponse, error)
	// GetReceivedGiftTransaction invokes getReceivedGiftTransaction operation.
	//
	// GET /v1/received_gifts/{gift_transaction_uuid}
	GetReceivedGiftTransaction(ctx context.Context, params GetReceivedGiftTransactionParams) (*GiftReceivedTransactionResponse, error)
	// GetRecentEngagementPosts invokes getRecentEngagementPosts operation.
	//
	// GET /v2/posts/recent_engagement
	GetRecentEngagementPosts(ctx context.Context, params GetRecentEngagementPostsParams) (*PostsResponse, error)
	// GetRecommendedFollowUsers invokes getRecommendedFollowUsers operation.
	//
	// GET /v1/users/{id}/follow_recommended
	GetRecommendedFollowUsers(ctx context.Context, params GetRecommendedFollowUsersParams) (*UsersResponse, error)
	// GetRecommendedPostTags invokes getRecommendedPostTags operation.
	//
	// POST /v1/posts/recommended_tag
	GetRecommendedPostTags(ctx context.Context, request *GetRecommendedPostTagsReq) (*PostTagsResponse, error)
	// GetRecommendedTimeline invokes getRecommendedTimeline operation.
	//
	// GET /v2/posts/recommended_timeline
	GetRecommendedTimeline(ctx context.Context, params GetRecommendedTimelineParams) (*PostsResponse, error)
	// GetRelatableGroups invokes getRelatableGroups operation.
	//
	// GET /v1/groups/{id}/relatable
	GetRelatableGroups(ctx context.Context, params GetRelatableGroupsParams) (*GroupsRelatedResponse, error)
	// GetRelatedGroups invokes getRelatedGroups operation.
	//
	// GET /v1/groups/{id}/related
	GetRelatedGroups(ctx context.Context, params GetRelatedGroupsParams) (*GroupsRelatedResponse, error)
	// GetResetCounters invokes getResetCounters operation.
	//
	// GET /v1/users/reset_counters
	GetResetCounters(ctx context.Context) (*RefreshCounterRequestsResponse, error)
	// GetRootPosts invokes getRootPosts operation.
	//
	// GET /v2/conversations/root_posts
	GetRootPosts(ctx context.Context, params GetRootPostsParams) (*PostsResponse, error)
	// GetThread invokes getThread operation.
	//
	// GET /v1/threads/{id}
	GetThread(ctx context.Context, params GetThreadParams) (*ThreadInfo, error)
	// GetThreadPosts invokes getThreadPosts operation.
	//
	// GET /v1/threads/{id}/posts
	GetThreadPosts(ctx context.Context, params GetThreadPostsParams) (*PostsResponse, error)
	// GetTimeline invokes getTimeline operation.
	//
	// GET /v2/posts/{noreply_mode}timeline
	GetTimeline(ctx context.Context, params GetTimelineParams) (*PostsResponse, error)
	// GetTwoFactorAuthRequestInfo invokes getTwoFactorAuthRequestInfo operation.
	//
	// GET /v1/users/secret
	GetTwoFactorAuthRequestInfo(ctx context.Context) (*TwoStepAuthRequestInfoResponse, error)
	// GetTwoFactorAuthStatus invokes getTwoFactorAuthStatus operation.
	//
	// GET /v1/users/secret/status
	GetTwoFactorAuthStatus(ctx context.Context) (*TwoFAStatusResponse, error)
	// GetUpdatedChatRooms invokes getUpdatedChatRooms operation.
	//
	// GET /v2/chat_rooms/update
	GetUpdatedChatRooms(ctx context.Context, params GetUpdatedChatRoomsParams) (*ChatRoomsResponse, error)
	// GetUser invokes getUser operation.
	//
	// GET /v2/users/{id}
	GetUser(ctx context.Context, params GetUserParams) (*UserResponse, error)
	// GetUserActivities invokes getUserActivities operation.
	//
	// GET /api/v2/user_activities
	GetUserActivities(ctx context.Context, params GetUserActivitiesParams) (*ActivitiesResponse, error)
	// GetUserActivitiesV1 invokes getUserActivitiesV1 operation.
	//
	// GET /api/user_activities
	GetUserActivitiesV1(ctx context.Context, params GetUserActivitiesV1Params) (*ActivitiesResponse, error)
	// GetUserByQr invokes getUserByQr operation.
	//
	// GET /v1/users/qr_codes/{qr}
	GetUserByQr(ctx context.Context, params GetUserByQrParams) (*UserResponse, error)
	// GetUserCustomDefinitions invokes getUserCustomDefinitions operation.
	//
	// GET /v1/users/custom_definitions
	GetUserCustomDefinitions(ctx context.Context) (*UserCustomDefinitionsResponse, error)
	// GetUserFollowers invokes getUserFollowers operation.
	//
	// GET /v3/users/{id}/followers
	GetUserFollowers(ctx context.Context, params GetUserFollowersParams) (*FollowUsersResponse, error)
	// GetUserFollowings invokes getUserFollowings operation.
	//
	// GET /v3/users/{id}/followings
	GetUserFollowings(ctx context.Context, params GetUserFollowingsParams) (*FollowUsersResponse, error)
	// GetUserGiftTransactions invokes getUserGiftTransactions operation.
	//
	// GET /v1/users/{user_id}/gift_transactions
	GetUserGiftTransactions(ctx context.Context, params GetUserGiftTransactionsParams) (*GiftTransactionsResponse, error)
	// GetUserGroupList invokes getUserGroupList operation.
	//
	// GET /v1/groups/user_group_list
	GetUserGroupList(ctx context.Context, params GetUserGroupListParams) (*GroupsResponse, error)
	// GetUserInfo invokes getUserInfo operation.
	//
	// GET /v2/users/info/{id}
	GetUserInfo(ctx context.Context, params GetUserInfoParams) (*UserResponse, error)
	// GetUserInterests invokes getUserInterests operation.
	//
	// GET /v1/users/interests
	GetUserInterests(ctx context.Context) (*UserInterestsResponse, error)
	// GetUserPresignedUrl invokes getUserPresignedUrl operation.
	//
	// GET /v1/users/presigned_url
	GetUserPresignedUrl(ctx context.Context, params GetUserPresignedUrlParams) (*PresignedUrlResponse, error)
	// GetUserReviews invokes getUserReviews operation.
	//
	// GET /v2/users/reviews/{id}
	GetUserReviews(ctx context.Context, params GetUserReviewsParams) (*ReviewsResponse, error)
	// GetUserTimeline invokes getUserTimeline operation.
	//
	// GET /v2/posts/user_timeline
	GetUserTimeline(ctx context.Context, params GetUserTimelineParams) (*PostsResponse, error)
	// GetUserTimestamp invokes getUserTimestamp operation.
	//
	// GET /v2/users/timestamp
	GetUserTimestamp(ctx context.Context) (*UserTimestampResponse, error)
	// GetUsersByIds invokes getUsersByIds operation.
	//
	// GET /v1/users/list_id
	GetUsersByIds(ctx context.Context, params GetUsersByIdsParams) (*UsersResponse, error)
	// GetWebSocketToken invokes getWebSocketToken operation.
	//
	// GET /v1/users/ws_token
	GetWebSocketToken(ctx context.Context) (*WebSocketTokenResponse, error)
	// HideChats invokes hideChats operation.
	//
	// POST /v1/hidden/chats
	HideChats(ctx context.Context, request *HideChatsReq) error
	// HideUsers invokes hideUsers operation.
	//
	// POST /v1/hidden/users
	HideUsers(ctx context.Context, request *HideUsersReq) error
	// InviteToCall invokes inviteToCall operation.
	//
	// POST /v2/calls/invite
	InviteToCall(ctx context.Context, request *InviteToCallReq) error
	// InviteToChatRoom invokes inviteToChatRoom operation.
	//
	// POST /v2/chat_rooms/{id}/invite
	InviteToChatRoom(ctx context.Context, request *InviteToChatRoomReq, params InviteToChatRoomParams) error
	// InviteToConferenceCall invokes inviteToConferenceCall operation.
	//
	// POST /v1/calls/conference_calls/{call_id}/invite
	InviteToConferenceCall(ctx context.Context, request *InviteToConferenceCallReq, params InviteToConferenceCallParams) error
	// InviteToGroup invokes inviteToGroup operation.
	//
	// POST /v1/groups/{id}/invite
	InviteToGroup(ctx context.Context, request *InviteToGroupReq, params InviteToGroupParams) error
	// JoinGroup invokes joinGroup operation.
	//
	// POST /v1/groups/{id}/join
	JoinGroup(ctx context.Context, params JoinGroupParams) error
	// KickFromChatRoom invokes kickFromChatRoom operation.
	//
	// POST /v2/chat_rooms/{id}/kick
	KickFromChatRoom(ctx context.Context, request *KickFromChatRoomReq, params KickFromChatRoomParams) error
	// KickFromConferenceCall invokes kickFromConferenceCall operation.
	//
	// POST /v3/calls/conference_calls/{call_id}/kick
	KickFromConferenceCall(ctx context.Context, request *KickFromConferenceCallReq, params KickFromConferenceCallParams) (*CallActionSignatureResponse, error)
	// LeaveAgoraChannel invokes leaveAgoraChannel operation.
	//
	// POST /v1/calls/leave_agora_channel
	LeaveAgoraChannel(ctx context.Context, request *LeaveAgoraChannelReq) error
	// LeaveConferenceCall invokes leaveConferenceCall operation.
	//
	// POST /v1/calls/leave_conference_call
	LeaveConferenceCall(ctx context.Context, request *LeaveConferenceCallReq) error
	// LeaveGroup invokes leaveGroup operation.
	//
	// DELETE /v1/groups/{id}/leave
	LeaveGroup(ctx context.Context, params LeaveGroupParams) error
	// LikePosts invokes likePosts operation.
	//
	// POST /v2/posts/like
	LikePosts(ctx context.Context, request *LikePostsReq) (*LikePostsResponse, error)
	// ListGameApps invokes listGameApps operation.
	//
	// GET /v1/games/apps
	ListGameApps(ctx context.Context, params ListGameAppsParams) (*GamesResponse, error)
	// ListGameWalkthroughs invokes listGameWalkthroughs operation.
	//
	// GET /v1/games/apps/{app_id}/walkthroughs
	ListGameWalkthroughs(ctx context.Context, params ListGameWalkthroughsParams) ([]Walkthrough, error)
	// ListGenres invokes listGenres operation.
	//
	// GET /v1/genres
	ListGenres(ctx context.Context, params ListGenresParams) (*GenresResponse, error)
	// ListGifts invokes listGifts operation.
	//
	// GET /v2/gifts
	ListGifts(ctx context.Context, params ListGiftsParams) (*GiftsResponse, error)
	// ListGroupCategories invokes listGroupCategories operation.
	//
	// GET /v1/groups/categories
	ListGroupCategories(ctx context.Context, params ListGroupCategoriesParams) (*GroupCategoriesResponse, error)
	// ListGroups invokes listGroups operation.
	//
	// GET /v2/groups
	ListGroups(ctx context.Context, params ListGroupsParams) (*GroupsResponse, error)
	// ListHiddenChats invokes listHiddenChats operation.
	//
	// GET /v1/hidden/chats
	ListHiddenChats(ctx context.Context, params ListHiddenChatsParams) (*ChatRoomsResponse, error)
	// ListHiddenUsers invokes listHiddenUsers operation.
	//
	// GET /v1/hidden/users
	ListHiddenUsers(ctx context.Context, params ListHiddenUsersParams) (*HiddenResponse, error)
	// ListMuteKeywords invokes listMuteKeywords operation.
	//
	// GET /v1/hidden/words
	ListMuteKeywords(ctx context.Context) (*MuteKeywordResponse, error)
	// ListMutedGroupUsers invokes listMutedGroupUsers operation.
	//
	// GET /v1/group_mute/{id}/muted_users
	ListMutedGroupUsers(ctx context.Context, params ListMutedGroupUsersParams) (*GroupMuteUsersResponse, error)
	// ListMyGroups invokes listMyGroups operation.
	//
	// GET /v2/groups/mine
	ListMyGroups(ctx context.Context, params ListMyGroupsParams) (*GroupsResponse, error)
	// ListReceivedGifts invokes listReceivedGifts operation.
	//
	// GET /v2/received_gifts
	ListReceivedGifts(ctx context.Context) (*GiftReceivedResponse, error)
	// ListReceivedGiftsV1 invokes listReceivedGiftsV1 operation.
	//
	// GET /v1/received_gifts
	ListReceivedGiftsV1(ctx context.Context, params ListReceivedGiftsV1Params) (*GiftReceivedResponse, error)
	// ListStickerPacks invokes listStickerPacks operation.
	//
	// GET /v2/sticker_packs
	ListStickerPacks(ctx context.Context) (*StickerPacksResponse, error)
	// ListThreads invokes listThreads operation.
	//
	// GET /v1/threads
	ListThreads(ctx context.Context, params ListThreadsParams) (*GroupThreadListResponse, error)
	// LoginWithEmail invokes loginWithEmail operation.
	//
	// POST /v3/users/login_with_email
	LoginWithEmail(ctx context.Context, request *LoginEmailUserRequest) (*LoginUserResponse, error)
	// Logout invokes logout operation.
	//
	// POST /v1/users/logout
	Logout(ctx context.Context, request *LogoutReq) error
	// MovePostToSpecificThread invokes movePostToSpecificThread operation.
	//
	// PUT /v3/posts/{id}/move_to_thread/{thread_id}
	MovePostToSpecificThread(ctx context.Context, params MovePostToSpecificThreadParams) (*ThreadInfo, error)
	// MovePostToThread invokes movePostToThread operation.
	//
	// POST /v3/posts/{id}/move_to_thread
	MovePostToThread(ctx context.Context, request *MovePostToThreadReq, params MovePostToThreadParams) (*ThreadInfo, error)
	// MuteGroupUser invokes muteGroupUser operation.
	//
	// POST /v1/group_mute/{id}/mute/{user_id}
	MuteGroupUser(ctx context.Context, params MuteGroupUserParams) error
	// OauthToken invokes oauthToken operation.
	//
	// POST /api/v1/oauth/token
	OauthToken(ctx context.Context, request *OauthTokenReq) (*TokenResponse, error)
	// PinChatRoom invokes pinChatRoom operation.
	//
	// POST /v1/chat_rooms/{id}/pinned
	PinChatRoom(ctx context.Context, params PinChatRoomParams) error
	// PinGroup invokes pinGroup operation.
	//
	// POST /v1/pinned/groups
	PinGroup(ctx context.Context, request *PinGroupReq) error
	// PinGroupHighlightPost invokes pinGroupHighlightPost operation.
	//
	// PUT /v1/groups/{group_id}/highlights/{post_id}
	PinGroupHighlightPost(ctx context.Context, params PinGroupHighlightPostParams) error
	// PinGroupPost invokes pinGroupPost operation.
	//
	// PUT /v2/posts/group_pinned_post
	PinGroupPost(ctx context.Context, request *PinGroupPostReq) error
	// PinPost invokes pinPost operation.
	//
	// POST /v1/pinned/posts
	PinPost(ctx context.Context, request *PinPostReq) error
	// PinReview invokes pinReview operation.
	//
	// POST /v1/pinned/reviews
	PinReview(ctx context.Context, request *PinReviewReq) error
	// PingAlive invokes pingAlive operation.
	//
	// POST /v1/users/alive
	PingAlive(ctx context.Context) error
	// ReadChatAttachments invokes readChatAttachments operation.
	//
	// POST /v1/chat_rooms/{id}/attachments_read
	ReadChatAttachments(ctx context.Context, request *ReadAttachmentRequest, params ReadChatAttachmentsParams) error
	// ReadChatMessage invokes readChatMessage operation.
	//
	// POST /v2/chat_rooms/{id}/messages/{message_id}/read
	ReadChatMessage(ctx context.Context, params ReadChatMessageParams) error
	// ReadChatVideos invokes readChatVideos operation.
	//
	// POST /v1/chat_rooms/{id}/videos_read
	ReadChatVideos(ctx context.Context, request *ReadChatVideosReq, params ReadChatVideosParams) error
	// RemoveChatRoomBackground invokes removeChatRoomBackground operation.
	//
	// DELETE /v2/chat_rooms/{id}/background
	RemoveChatRoomBackground(ctx context.Context, params RemoveChatRoomBackgroundParams) error
	// RemoveCoverImage invokes removeCoverImage operation.
	//
	// POST /v2/users/remove_cover_image
	RemoveCoverImage(ctx context.Context) error
	// RemoveGroupCover invokes removeGroupCover operation.
	//
	// DELETE /v3/groups/{id}/cover
	RemoveGroupCover(ctx context.Context, params RemoveGroupCoverParams) error
	// RemoveGroupDeputies invokes removeGroupDeputies operation.
	//
	// DELETE /v1/groups/{id}/deputize
	RemoveGroupDeputies(ctx context.Context, params RemoveGroupDeputiesParams) error
	// RemoveGroupIcon invokes removeGroupIcon operation.
	//
	// DELETE /v3/groups/{id}/icon
	RemoveGroupIcon(ctx context.Context, params RemoveGroupIconParams) error
	// RemoveProfilePhoto invokes removeProfilePhoto operation.
	//
	// POST /v2/users/remove_profile_photo
	RemoveProfilePhoto(ctx context.Context) error
	// RemoveRelatedGroups invokes removeRelatedGroups operation.
	//
	// DELETE /v1/groups/{id}/related
	RemoveRelatedGroups(ctx context.Context, params RemoveRelatedGroupsParams) error
	// RemoveThreadMember invokes removeThreadMember operation.
	//
	// DELETE /v1/threads/{thread_id}/members/{id}
	RemoveThreadMember(ctx context.Context, params RemoveThreadMemberParams) error
	// ReplyToReview invokes replyToReview operation.
	//
	// POST /v2/users/reviews/{id}
	ReplyToReview(ctx context.Context, request *ReplyToReviewReq, params ReplyToReviewParams) error
	// ReportChatRoom invokes reportChatRoom operation.
	//
	// POST /v3/chat_rooms/{chat_room_id}/report
	ReportChatRoom(ctx context.Context, request *ReportChatRoomReq, params ReportChatRoomParams) error
	// ReportGroup invokes reportGroup operation.
	//
	// POST /v3/groups/{group_id}/report
	ReportGroup(ctx context.Context, request *ReportGroupReq, params ReportGroupParams) error
	// ReportPost invokes reportPost operation.
	//
	// POST /v3/posts/{post_id}/report
	ReportPost(ctx context.Context, request *ReportPostReq, params ReportPostParams) error
	// ReportUser invokes reportUser operation.
	//
	// POST /v3/users/{user_id}/report
	ReportUser(ctx context.Context, request *ReportUserReq, params ReportUserParams) error
	// Repost invokes repost operation.
	//
	// POST /v3/posts/repost
	Repost(ctx context.Context, request *RepostReq, params RepostParams) (*CreatePostResponse, error)
	// RequestEmailVerification invokes requestEmailVerification operation.
	//
	// POST /v1/email_verification_urls
	RequestEmailVerification(ctx context.Context, request *RequestEmailVerificationReq) (*CommonUrlResponse, error)
	// RequestFollow invokes requestFollow operation.
	//
	// POST /v2/users/{target_id}/follow_request
	RequestFollow(ctx context.Context, request *RequestFollowReq, params RequestFollowParams) error
	// RequestGroupWalkthrough invokes requestGroupWalkthrough operation.
	//
	// POST /v1/groups/{id}/request_walkthrough
	RequestGroupWalkthrough(ctx context.Context, params RequestGroupWalkthroughParams) error
	// ResendConfirmEmail invokes resendConfirmEmail operation.
	//
	// POST /v2/users/resend_confirm_email
	ResendConfirmEmail(ctx context.Context) error
	// ResetCounters invokes resetCounters operation.
	//
	// POST /v1/users/reset_counters
	ResetCounters(ctx context.Context, request *ResetCountersReq) error
	// ResetPassword invokes resetPassword operation.
	//
	// PUT /v1/users/reset_password
	ResetPassword(ctx context.Context, request *ResetPasswordReq) error
	// SearchGroupPosts invokes searchGroupPosts operation.
	//
	// GET /v2/groups/{id}/posts/search
	SearchGroupPosts(ctx context.Context, params SearchGroupPostsParams) (*PostsResponse, error)
	// SearchPosts invokes searchPosts operation.
	//
	// GET /v2/posts/search
	SearchPosts(ctx context.Context, params SearchPostsParams) (*PostsResponse, error)
	// SearchUsers invokes searchUsers operation.
	//
	// GET /v1/users/search
	SearchUsers(ctx context.Context, params SearchUsersParams) (*UsersResponse, error)
	// SendChatMessage invokes sendChatMessage operation.
	//
	// POST /v3/chat_rooms/{id}/messages/new
	SendChatMessage(ctx context.Context, request *SendChatMessageReq, params SendChatMessageParams) (*MessageResponse, error)
	// SetGroupTitle invokes setGroupTitle operation.
	//
	// POST /v1/groups/{id}/set_title
	SetGroupTitle(ctx context.Context, request *SetGroupTitleReq, params SetGroupTitleParams) error
	// SetHima invokes setHima operation.
	//
	// POST /v1/users/hima
	SetHima(ctx context.Context) error
	// StartConferenceCall invokes startConferenceCall operation.
	//
	// POST /v2/calls/start_conference_call
	StartConferenceCall(ctx context.Context, request *StartConferenceCallReq) (*ConferenceCallResponse, error)
	// TakeOverGroup invokes takeOverGroup operation.
	//
	// POST /v1/groups/{id}/take_over
	TakeOverGroup(ctx context.Context, params TakeOverGroupParams) error
	// TransferGroup invokes transferGroup operation.
	//
	// POST /v3/groups/{id}/transfer
	TransferGroup(ctx context.Context, request *TransferGroupReq, params TransferGroupParams) error
	// UnbanGroupUser invokes unbanGroupUser operation.
	//
	// POST /v1/groups/{id}/unban/{userId}
	UnbanGroupUser(ctx context.Context, params UnbanGroupUserParams) error
	// UnblockUser invokes unblockUser operation.
	//
	// POST /v2/users/{id}/unblock
	UnblockUser(ctx context.Context, params UnblockUserParams) error
	// UnfollowUser invokes unfollowUser operation.
	//
	// POST /v2/users/{id}/unfollow
	UnfollowUser(ctx context.Context, params UnfollowUserParams) error
	// UnhideChats invokes unhideChats operation.
	//
	// DELETE /v1/hidden/chats
	UnhideChats(ctx context.Context, params UnhideChatsParams) error
	// UnhideUsers invokes unhideUsers operation.
	//
	// DELETE /v1/hidden/users
	UnhideUsers(ctx context.Context, params UnhideUsersParams) error
	// UnlikePost invokes unlikePost operation.
	//
	// POST /v1/posts/{id}/unlike
	UnlikePost(ctx context.Context, params UnlikePostParams) error
	// UnmuteGroupUser invokes unmuteGroupUser operation.
	//
	// DELETE /v1/group_mute/{id}/unmute/{user_id}
	UnmuteGroupUser(ctx context.Context, params UnmuteGroupUserParams) error
	// UnpinChatRoom invokes unpinChatRoom operation.
	//
	// DELETE /v1/chat_rooms/{id}/pinned
	UnpinChatRoom(ctx context.Context, params UnpinChatRoomParams) error
	// UnpinGroup invokes unpinGroup operation.
	//
	// DELETE /v1/pinned/groups/{id}
	UnpinGroup(ctx context.Context, params UnpinGroupParams) error
	// UnpinGroupHighlightPost invokes unpinGroupHighlightPost operation.
	//
	// DELETE /v1/groups/{group_id}/highlights/{post_id}
	UnpinGroupHighlightPost(ctx context.Context, params UnpinGroupHighlightPostParams) error
	// UnpinGroupPost invokes unpinGroupPost operation.
	//
	// DELETE /v2/posts/group_pinned_post
	UnpinGroupPost(ctx context.Context, params UnpinGroupPostParams) error
	// UnpinReview invokes unpinReview operation.
	//
	// DELETE /v1/pinned/reviews/{id}
	UnpinReview(ctx context.Context, params UnpinReviewParams) error
	// UpdateAdditionalNotificationSetting invokes updateAdditionalNotificationSetting operation.
	//
	// POST /v1/users/additonal_notification_setting
	UpdateAdditionalNotificationSetting(ctx context.Context, request *UpdateAdditionalNotificationSettingReq) error
	// UpdateCall invokes updateCall operation.
	//
	// PUT /v1/calls/{call_id}
	UpdateCall(ctx context.Context, request *UpdateCallReq, params UpdateCallParams) error
	// UpdateCallUser invokes updateCallUser operation.
	//
	// PUT /v1/calls/{call_id}/users/{user_id}
	UpdateCallUser(ctx context.Context, request *UpdateCallUserReq, params UpdateCallUserParams) error
	// UpdateChatRoom invokes updateChatRoom operation.
	//
	// POST /v3/chat_rooms/{id}/edit
	UpdateChatRoom(ctx context.Context, request *UpdateChatRoomReq, params UpdateChatRoomParams) error
	// UpdateChatRoomNotificationSettings invokes updateChatRoomNotificationSettings operation.
	//
	// POST /v2/notification_settings/chat_rooms/{id}
	UpdateChatRoomNotificationSettings(ctx context.Context, request *UpdateChatRoomNotificationSettingsReq, params UpdateChatRoomNotificationSettingsParams) (*NotificationSettingResponse, error)
	// UpdateGroup invokes updateGroup operation.
	//
	// POST /v3/groups/{id}/update
	UpdateGroup(ctx context.Context, request *UpdateGroupReq, params UpdateGroupParams) (*GroupResponse, error)
	// UpdateGroupNotificationSettings invokes updateGroupNotificationSettings operation.
	//
	// POST /v2/notification_settings/groups/{id}
	UpdateGroupNotificationSettings(ctx context.Context, request *UpdateGroupNotificationSettingsReq, params UpdateGroupNotificationSettingsParams) (*AdditionalSettingsResponse, error)
	// UpdateLanguage invokes updateLanguage operation.
	//
	// PUT /v1/users/language
	UpdateLanguage(ctx context.Context, request *UpdateLanguageReq) error
	// UpdateLogin invokes updateLogin operation.
	//
	// POST /v3/users/login_update
	UpdateLogin(ctx context.Context, request *UpdateLoginReq) (*LoginUpdateResponse, error)
	// UpdatePost invokes updatePost operation.
	//
	// PUT /v3/posts/{id}
	UpdatePost(ctx context.Context, request *UpdatePostReq, params UpdatePostParams) (*Post, error)
	// UpdateRelatedGroups invokes updateRelatedGroups operation.
	//
	// PUT /v1/groups/{id}/related
	UpdateRelatedGroups(ctx context.Context, params UpdateRelatedGroupsParams) error
	// UpdateThread invokes updateThread operation.
	//
	// PUT /v1/threads/{id}
	UpdateThread(ctx context.Context, request *UpdateThreadReq, params UpdateThreadParams) error
	// UpdateUserInterests invokes updateUserInterests operation.
	//
	// PUT /v1/users/interests
	UpdateUserInterests(ctx context.Context, request *CommonIdsRequest) error
	// ValidateCallActionSignature invokes validateCallActionSignature operation.
	//
	// GET /v2/calls/action_signature/validate
	ValidateCallActionSignature(ctx context.Context, params ValidateCallActionSignatureParams) error
	// ValidatePost invokes validatePost operation.
	//
	// POST /v1/posts/validate
	ValidatePost(ctx context.Context, request *ValidatePostReq) (*ValidationPostResponse, error)
	// ViewPostVideo invokes viewPostVideo operation.
	//
	// POST /v1/posts/videos/{videoId}/view
	ViewPostVideo(ctx context.Context, params ViewPostVideoParams) error
	// VisitGroup invokes visitGroup operation.
	//
	// POST /v1/groups/{id}/visit
	VisitGroup(ctx context.Context, params VisitGroupParams) error
	// VoteSurvey invokes voteSurvey operation.
	//
	// POST /v2/surveys/{id}/vote
	VoteSurvey(ctx context.Context, request *VoteSurveyReq, params VoteSurveyParams) (*VoteSurveyResponse, error)
	// WithdrawGroupDeputy invokes withdrawGroupDeputy operation.
	//
	// PUT /v1/groups/{group_id}/deputize/{user_id}/withdraw
	WithdrawGroupDeputy(ctx context.Context, params WithdrawGroupDeputyParams) error
	// WithdrawGroupTransfer invokes withdrawGroupTransfer operation.
	//
	// PUT /v1/groups/{id}/transfer/withdraw
	WithdrawGroupTransfer(ctx context.Context, request *WithdrawGroupTransferReq, params WithdrawGroupTransferParams) error
}

// Client implements OAS client.
type Client struct {
	serverURL *url.URL
	baseClient
}

// NewClient initializes new Client defined by OAS.
func NewClient(serverURL string, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	trimTrailingSlashes(u)

	c, err := newClientConfig(opts...).baseClient()
	if err != nil {
		return nil, err
	}
	return &Client{
		serverURL:  u,
		baseClient: c,
	}, nil
}

type serverURLKey struct{}

// WithServerURL sets context key to override server URL.
func WithServerURL(ctx context.Context, u *url.URL) context.Context {
	return context.WithValue(ctx, serverURLKey{}, u)
}

func (c *Client) requestURL(ctx context.Context) *url.URL {
	u, ok := ctx.Value(serverURLKey{}).(*url.URL)
	if !ok {
		return c.serverURL
	}
	return u
}

// AcceptChatRequest invokes acceptChatRequest operation.
//
// POST /v1/chat_rooms/accept_chat_request
func (c *Client) AcceptChatRequest(ctx context.Context, request *AcceptChatRequestReq) error {
	_, err := c.sendAcceptChatRequest(ctx, request)
	return err
}

func (c *Client) sendAcceptChatRequest(ctx context.Context, request *AcceptChatRequestReq) (res *AcceptChatRequestNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("acceptChatRequest"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/accept_chat_request"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AcceptChatRequestOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/accept_chat_request"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeAcceptChatRequestRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeAcceptChatRequestResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AcceptGroupJoinRequest invokes acceptGroupJoinRequest operation.
//
// POST /v1/groups/{id}/accept/{userId}
func (c *Client) AcceptGroupJoinRequest(ctx context.Context, params AcceptGroupJoinRequestParams) error {
	_, err := c.sendAcceptGroupJoinRequest(ctx, params)
	return err
}

func (c *Client) sendAcceptGroupJoinRequest(ctx context.Context, params AcceptGroupJoinRequestParams) (res *AcceptGroupJoinRequestNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("acceptGroupJoinRequest"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/accept/{userId}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AcceptGroupJoinRequestOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/accept/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeAcceptGroupJoinRequestResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AcceptGroupTransfer invokes acceptGroupTransfer operation.
//
// PUT /v1/groups/{id}/transfer
func (c *Client) AcceptGroupTransfer(ctx context.Context, params AcceptGroupTransferParams) error {
	_, err := c.sendAcceptGroupTransfer(ctx, params)
	return err
}

func (c *Client) sendAcceptGroupTransfer(ctx context.Context, params AcceptGroupTransferParams) (res *AcceptGroupTransferNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("acceptGroupTransfer"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/transfer"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AcceptGroupTransferOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/transfer"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeAcceptGroupTransferResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AddThreadMember invokes addThreadMember operation.
//
// POST /v1/threads/{thread_id}/members/{id}
func (c *Client) AddThreadMember(ctx context.Context, params AddThreadMemberParams) error {
	_, err := c.sendAddThreadMember(ctx, params)
	return err
}

func (c *Client) sendAddThreadMember(ctx context.Context, params AddThreadMemberParams) (res *AddThreadMemberNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("addThreadMember"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/threads/{thread_id}/members/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AddThreadMemberOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "thread_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "thread_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ThreadID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/members/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeAddThreadMemberResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// AgreePolicy invokes agreePolicy operation.
//
// POST /v1/users/policy_agreements/{type}
func (c *Client) AgreePolicy(ctx context.Context, params AgreePolicyParams) error {
	_, err := c.sendAgreePolicy(ctx, params)
	return err
}

func (c *Client) sendAgreePolicy(ctx context.Context, params AgreePolicyParams) (res *AgreePolicyNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("agreePolicy"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/policy_agreements/{type}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, AgreePolicyOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/users/policy_agreements/"
	{
		// Encode "type" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "type",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Type))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeAgreePolicyResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// BanGroupUser invokes banGroupUser operation.
//
// POST /v1/groups/{id}/ban/{userId}
func (c *Client) BanGroupUser(ctx context.Context, params BanGroupUserParams) error {
	_, err := c.sendBanGroupUser(ctx, params)
	return err
}

func (c *Client) sendBanGroupUser(ctx context.Context, params BanGroupUserParams) (res *BanGroupUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("banGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/ban/{userId}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, BanGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/ban/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeBanGroupUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// BlockUser invokes blockUser operation.
//
// POST /v1/users/{id}/block
func (c *Client) BlockUser(ctx context.Context, params BlockUserParams) error {
	_, err := c.sendBlockUser(ctx, params)
	return err
}

func (c *Client) sendBlockUser(ctx context.Context, params BlockUserParams) (res *BlockUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("blockUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/{id}/block"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, BlockUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/block"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeBlockUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// BulkInviteToCall invokes bulkInviteToCall operation.
//
// POST /v1/calls/{call_id}/bulk_invite
func (c *Client) BulkInviteToCall(ctx context.Context, params BulkInviteToCallParams) error {
	_, err := c.sendBulkInviteToCall(ctx, params)
	return err
}

func (c *Client) sendBulkInviteToCall(ctx context.Context, params BulkInviteToCallParams) (res *BulkInviteToCallNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("bulkInviteToCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/calls/{call_id}/bulk_invite"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, BulkInviteToCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/bulk_invite"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeBulkInviteToCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// BumpCall invokes bumpCall operation.
//
// POST /v1/calls/{call_id}/bump
func (c *Client) BumpCall(ctx context.Context, params BumpCallParams) error {
	_, err := c.sendBumpCall(ctx, params)
	return err
}

func (c *Client) sendBumpCall(ctx context.Context, params BumpCallParams) (res *BumpCallNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("bumpCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/calls/{call_id}/bump"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, BumpCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/bump"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "participant_limit" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "participant_limit",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ParticipantLimit.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeBumpCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CancelGroupTransfer invokes cancelGroupTransfer operation.
//
// DELETE /v1/groups/{id}/transfer
func (c *Client) CancelGroupTransfer(ctx context.Context, params CancelGroupTransferParams) error {
	_, err := c.sendCancelGroupTransfer(ctx, params)
	return err
}

func (c *Client) sendCancelGroupTransfer(ctx context.Context, params CancelGroupTransferParams) (res *CancelGroupTransferNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("cancelGroupTransfer"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/transfer"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CancelGroupTransferOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/transfer"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCancelGroupTransferResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChangeEmail invokes changeEmail operation.
//
// PUT /v1/users/change_email
func (c *Client) ChangeEmail(ctx context.Context, request *ChangeEmailReq) (*LoginUpdateResponse, error) {
	res, err := c.sendChangeEmail(ctx, request)
	return res, err
}

func (c *Client) sendChangeEmail(ctx context.Context, request *ChangeEmailReq) (res *LoginUpdateResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("changeEmail"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/change_email"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ChangeEmailOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/change_email"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeChangeEmailRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeChangeEmailResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ChangePassword invokes changePassword operation.
//
// PUT /v1/users/change_password
func (c *Client) ChangePassword(ctx context.Context, request *ChangePasswordReq) (*LoginUpdateResponse, error) {
	res, err := c.sendChangePassword(ctx, request)
	return res, err
}

func (c *Client) sendChangePassword(ctx context.Context, request *ChangePasswordReq) (res *LoginUpdateResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("changePassword"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/change_password"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ChangePasswordOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/change_password"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeChangePasswordRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeChangePasswordResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateBookmark invokes createBookmark operation.
//
// PUT /v1/users/{user_id}/bookmarks/{id}
func (c *Client) CreateBookmark(ctx context.Context, params CreateBookmarkParams) (*BookmarkPostResponse, error) {
	res, err := c.sendCreateBookmark(ctx, params)
	return res, err
}

func (c *Client) sendCreateBookmark(ctx context.Context, params CreateBookmarkParams) (res *BookmarkPostResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createBookmark"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/{user_id}/bookmarks/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateBookmarkOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/bookmarks/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateBookmarkResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateChatRoom invokes createChatRoom operation.
//
// POST /v3/chat_rooms/new
func (c *Client) CreateChatRoom(ctx context.Context, request *CreateChatRoomReq) (*CreateChatRoomResponse, error) {
	res, err := c.sendCreateChatRoom(ctx, request)
	return res, err
}

func (c *Client) sendCreateChatRoom(ctx context.Context, request *CreateChatRoomReq) (res *CreateChatRoomResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/chat_rooms/new"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/chat_rooms/new"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateChatRoomRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateChatRoomV1 invokes createChatRoomV1 operation.
//
// POST /v1/chat_rooms/new
func (c *Client) CreateChatRoomV1(ctx context.Context, request *CreateChatRoomV1Req) (*CreateChatRoomResponse, error) {
	res, err := c.sendCreateChatRoomV1(ctx, request)
	return res, err
}

func (c *Client) sendCreateChatRoomV1(ctx context.Context, request *CreateChatRoomV1Req) (res *CreateChatRoomResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createChatRoomV1"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/new"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateChatRoomV1Operation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/new"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateChatRoomV1Request(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateChatRoomV1Response(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateConferenceCallPost invokes createConferenceCallPost operation.
//
// POST /v2/posts/new_conference_call
func (c *Client) CreateConferenceCallPost(ctx context.Context, request *CreateConferenceCallPostReq) (*CreatePostResponse, error) {
	res, err := c.sendCreateConferenceCallPost(ctx, request)
	return res, err
}

func (c *Client) sendCreateConferenceCallPost(ctx context.Context, request *CreateConferenceCallPostReq) (res *CreatePostResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createConferenceCallPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/posts/new_conference_call"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateConferenceCallPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/new_conference_call"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateConferenceCallPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateConferenceCallPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateGroup invokes createGroup operation.
//
// POST /v3/groups/new
func (c *Client) CreateGroup(ctx context.Context, request *CreateGroupReq) (*CreateGroupResponse, error) {
	res, err := c.sendCreateGroup(ctx, request)
	return res, err
}

func (c *Client) sendCreateGroup(ctx context.Context, request *CreateGroupReq) (res *CreateGroupResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/groups/new"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/groups/new"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateGroupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateMuteKeyword invokes createMuteKeyword operation.
//
// POST /v1/hidden/words
func (c *Client) CreateMuteKeyword(ctx context.Context, request *MuteKeywordRequest) (*CreateMuteKeywordResponse, error) {
	res, err := c.sendCreateMuteKeyword(ctx, request)
	return res, err
}

func (c *Client) sendCreateMuteKeyword(ctx context.Context, request *MuteKeywordRequest) (res *CreateMuteKeywordResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createMuteKeyword"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/hidden/words"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateMuteKeywordOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/words"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateMuteKeywordRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateMuteKeywordResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreatePost invokes createPost operation.
//
// POST /v3/posts/new
func (c *Client) CreatePost(ctx context.Context, request *CreatePostReq, params CreatePostParams) (*Post, error) {
	res, err := c.sendCreatePost(ctx, request, params)
	return res, err
}

func (c *Client) sendCreatePost(ctx context.Context, request *CreatePostReq, params CreatePostParams) (res *Post, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/posts/new"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreatePostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/posts/new"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreatePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.XJwt.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreatePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateReview invokes createReview operation.
//
// POST /v1/users/reviews
func (c *Client) CreateReview(ctx context.Context, request *CreateReviewReq) error {
	_, err := c.sendCreateReview(ctx, request)
	return err
}

func (c *Client) sendCreateReview(ctx context.Context, request *CreateReviewReq) (res *CreateReviewNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createReview"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/reviews"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateReviewOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/reviews"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateReviewRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateReviewResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateSharePost invokes createSharePost operation.
//
// POST /v2/posts/new_share_post
func (c *Client) CreateSharePost(ctx context.Context, request *CreateSharePostReq) (*Post, error) {
	res, err := c.sendCreateSharePost(ctx, request)
	return res, err
}

func (c *Client) sendCreateSharePost(ctx context.Context, request *CreateSharePostReq) (res *Post, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createSharePost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/posts/new_share_post"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateSharePostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/new_share_post"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateSharePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateSharePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateThread invokes createThread operation.
//
// POST /v1/threads
func (c *Client) CreateThread(ctx context.Context, request *CreateGroupThreadRequest) (*ThreadInfo, error) {
	res, err := c.sendCreateThread(ctx, request)
	return res, err
}

func (c *Client) sendCreateThread(ctx context.Context, request *CreateGroupThreadRequest) (res *ThreadInfo, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createThread"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/threads"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateThreadOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/threads"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateThreadRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateThreadResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// CreateThreadPost invokes createThreadPost operation.
//
// POST /v1/threads/{id}/posts
func (c *Client) CreateThreadPost(ctx context.Context, request *CreateThreadPostReq, params CreateThreadPostParams) (*Post, error) {
	res, err := c.sendCreateThreadPost(ctx, request, params)
	return res, err
}

func (c *Client) sendCreateThreadPost(ctx context.Context, request *CreateThreadPostReq, params CreateThreadPostParams) (res *Post, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("createThreadPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/threads/{id}/posts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, CreateThreadPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/posts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeCreateThreadPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.XJwt.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeCreateThreadPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeclineGroupJoinRequest invokes declineGroupJoinRequest operation.
//
// POST /v1/groups/{id}/decline/{userId}
func (c *Client) DeclineGroupJoinRequest(ctx context.Context, params DeclineGroupJoinRequestParams) error {
	_, err := c.sendDeclineGroupJoinRequest(ctx, params)
	return err
}

func (c *Client) sendDeclineGroupJoinRequest(ctx context.Context, params DeclineGroupJoinRequestParams) (res *DeclineGroupJoinRequestNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("declineGroupJoinRequest"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/decline/{userId}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeclineGroupJoinRequestOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/decline/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeclineGroupJoinRequestResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteAllPosts invokes deleteAllPosts operation.
//
// POST /v1/posts/delete_all_post
func (c *Client) DeleteAllPosts(ctx context.Context) error {
	_, err := c.sendDeleteAllPosts(ctx)
	return err
}

func (c *Client) sendDeleteAllPosts(ctx context.Context) (res *DeleteAllPostsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteAllPosts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/posts/delete_all_post"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteAllPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/posts/delete_all_post"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteAllPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteBookmark invokes deleteBookmark operation.
//
// DELETE /v1/users/{user_id}/bookmarks/{id}
func (c *Client) DeleteBookmark(ctx context.Context, params DeleteBookmarkParams) error {
	_, err := c.sendDeleteBookmark(ctx, params)
	return err
}

func (c *Client) sendDeleteBookmark(ctx context.Context, params DeleteBookmarkParams) (res *DeleteBookmarkNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteBookmark"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/users/{user_id}/bookmarks/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteBookmarkOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/bookmarks/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteBookmarkResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteChatMessage invokes deleteChatMessage operation.
//
// DELETE /v1/chat_rooms/{room_id}/messages/{message_id}/delete
func (c *Client) DeleteChatMessage(ctx context.Context, params DeleteChatMessageParams) error {
	_, err := c.sendDeleteChatMessage(ctx, params)
	return err
}

func (c *Client) sendDeleteChatMessage(ctx context.Context, params DeleteChatMessageParams) (res *DeleteChatMessageNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteChatMessage"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/{room_id}/messages/{message_id}/delete"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteChatMessageOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [5]string
	pathParts[0] = "/v1/chat_rooms/"
	{
		// Encode "room_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "room_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.RoomID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages/"
	{
		// Encode "message_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "message_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.MessageID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	pathParts[4] = "/delete"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteChatMessageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteChatRooms invokes deleteChatRooms operation.
//
// POST /v1/chat_rooms/mass_destroy
func (c *Client) DeleteChatRooms(ctx context.Context, request *DeleteChatRoomsReq) error {
	_, err := c.sendDeleteChatRooms(ctx, request)
	return err
}

func (c *Client) sendDeleteChatRooms(ctx context.Context, request *DeleteChatRoomsReq) (res *DeleteChatRoomsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteChatRooms"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/mass_destroy"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteChatRoomsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/mass_destroy"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeDeleteChatRoomsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteChatRoomsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteFootprint invokes deleteFootprint operation.
//
// DELETE /v2/users/{user_id}/footprints/{footprint_id}
func (c *Client) DeleteFootprint(ctx context.Context, params DeleteFootprintParams) error {
	_, err := c.sendDeleteFootprint(ctx, params)
	return err
}

func (c *Client) sendDeleteFootprint(ctx context.Context, params DeleteFootprintParams) (res *DeleteFootprintNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteFootprint"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v2/users/{user_id}/footprints/{footprint_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteFootprintOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v2/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/footprints/"
	{
		// Encode "footprint_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "footprint_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.FootprintID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteFootprintResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteMuteKeyword invokes deleteMuteKeyword operation.
//
// DELETE /v1/hidden/words
func (c *Client) DeleteMuteKeyword(ctx context.Context, params DeleteMuteKeywordParams) error {
	_, err := c.sendDeleteMuteKeyword(ctx, params)
	return err
}

func (c *Client) sendDeleteMuteKeyword(ctx context.Context, params DeleteMuteKeywordParams) (res *DeleteMuteKeywordNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteMuteKeyword"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/hidden/words"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteMuteKeywordOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/words"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.Ids != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.Ids {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteMuteKeywordResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteMyReviews invokes deleteMyReviews operation.
//
// DELETE /v1/users/reviews
func (c *Client) DeleteMyReviews(ctx context.Context, params DeleteMyReviewsParams) error {
	_, err := c.sendDeleteMyReviews(ctx, params)
	return err
}

func (c *Client) sendDeleteMyReviews(ctx context.Context, params DeleteMyReviewsParams) (res *DeleteMyReviewsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteMyReviews"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/users/reviews"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteMyReviewsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/reviews"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "review_ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "review_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.ReviewIds != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.ReviewIds {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteMyReviewsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeletePosts invokes deletePosts operation.
//
// POST /v2/posts/mass_destroy
func (c *Client) DeletePosts(ctx context.Context, request *DeletePostsReq) error {
	_, err := c.sendDeletePosts(ctx, request)
	return err
}

func (c *Client) sendDeletePosts(ctx context.Context, request *DeletePostsReq) (res *DeletePostsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deletePosts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/posts/mass_destroy"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeletePostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/mass_destroy"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeDeletePostsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeletePostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeleteThread invokes deleteThread operation.
//
// DELETE /v1/threads/{id}
func (c *Client) DeleteThread(ctx context.Context, params DeleteThreadParams) error {
	_, err := c.sendDeleteThread(ctx, params)
	return err
}

func (c *Client) sendDeleteThread(ctx context.Context, params DeleteThreadParams) (res *DeleteThreadNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deleteThread"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/threads/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeleteThreadOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeleteThreadResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeputizeGroupUsers invokes deputizeGroupUsers operation.
//
// PUT /v1/groups/{id}/deputize
func (c *Client) DeputizeGroupUsers(ctx context.Context, params DeputizeGroupUsersParams) error {
	_, err := c.sendDeputizeGroupUsers(ctx, params)
	return err
}

func (c *Client) sendDeputizeGroupUsers(ctx context.Context, params DeputizeGroupUsersParams) (res *DeputizeGroupUsersNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deputizeGroupUsers"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/deputize"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeputizeGroupUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/deputize"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeputizeGroupUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DeputizeGroupUsersMass invokes deputizeGroupUsersMass operation.
//
// POST /v3/groups/{group_id}/deputize/mass
func (c *Client) DeputizeGroupUsersMass(ctx context.Context, request *DeputizeGroupUsersMassReq, params DeputizeGroupUsersMassParams) error {
	_, err := c.sendDeputizeGroupUsersMass(ctx, request, params)
	return err
}

func (c *Client) sendDeputizeGroupUsersMass(ctx context.Context, request *DeputizeGroupUsersMassReq, params DeputizeGroupUsersMassParams) (res *DeputizeGroupUsersMassNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("deputizeGroupUsersMass"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/groups/{group_id}/deputize/mass"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DeputizeGroupUsersMassOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/deputize/mass"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeDeputizeGroupUsersMassRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDeputizeGroupUsersMassResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// DisableTwoFactorAuth invokes disableTwoFactorAuth operation.
//
// PUT /v1/users/secret/disable
func (c *Client) DisableTwoFactorAuth(ctx context.Context, request *DisableTwoFactorAuthReq) error {
	_, err := c.sendDisableTwoFactorAuth(ctx, request)
	return err
}

func (c *Client) sendDisableTwoFactorAuth(ctx context.Context, request *DisableTwoFactorAuthReq) (res *DisableTwoFactorAuthNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("disableTwoFactorAuth"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/secret/disable"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, DisableTwoFactorAuthOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/secret/disable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeDisableTwoFactorAuthRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeDisableTwoFactorAuthResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// EditUser invokes editUser operation.
//
// POST /v3/users/edit
func (c *Client) EditUser(ctx context.Context, request *EditUserReq) error {
	_, err := c.sendEditUser(ctx, request)
	return err
}

func (c *Client) sendEditUser(ctx context.Context, request *EditUserReq) (res *EditUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("editUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/users/edit"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, EditUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/users/edit"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeEditUserRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeEditUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// EditUserV2 invokes editUserV2 operation.
//
// POST /v2/users/edit
func (c *Client) EditUserV2(ctx context.Context, request *EditUserV2Req) error {
	_, err := c.sendEditUserV2(ctx, request)
	return err
}

func (c *Client) sendEditUserV2(ctx context.Context, request *EditUserV2Req) (res *EditUserV2NoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("editUserV2"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/edit"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, EditUserV2Operation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/edit"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeEditUserV2Request(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeEditUserV2Response(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// EnableTwoFactorAuth invokes enableTwoFactorAuth operation.
//
// PUT /v1/users/secret/enable
func (c *Client) EnableTwoFactorAuth(ctx context.Context, request *EnableTwoFactorAuthReq) (*TwoStepAuthEnabledResponse, error) {
	res, err := c.sendEnableTwoFactorAuth(ctx, request)
	return res, err
}

func (c *Client) sendEnableTwoFactorAuth(ctx context.Context, request *EnableTwoFactorAuthReq) (res *TwoStepAuthEnabledResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("enableTwoFactorAuth"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/secret/enable"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, EnableTwoFactorAuthOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/secret/enable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeEnableTwoFactorAuthRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeEnableTwoFactorAuthResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// FireGroupUser invokes fireGroupUser operation.
//
// POST /v1/groups/{group_id}/fire/{user_id}
func (c *Client) FireGroupUser(ctx context.Context, params FireGroupUserParams) error {
	_, err := c.sendFireGroupUser(ctx, params)
	return err
}

func (c *Client) sendFireGroupUser(ctx context.Context, params FireGroupUserParams) (res *FireGroupUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("fireGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/fire/{user_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, FireGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/fire/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeFireGroupUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// FollowUser invokes followUser operation.
//
// POST /v2/users/{id}/follow
func (c *Client) FollowUser(ctx context.Context, params FollowUserParams) error {
	_, err := c.sendFollowUser(ctx, params)
	return err
}

func (c *Client) sendFollowUser(ctx context.Context, params FollowUserParams) (res *FollowUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("followUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/{id}/follow"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, FollowUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/follow"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeFollowUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// FollowUsers invokes followUsers operation.
//
// POST /v2/users/follow
func (c *Client) FollowUsers(ctx context.Context, params FollowUsersParams) error {
	_, err := c.sendFollowUsers(ctx, params)
	return err
}

func (c *Client) sendFollowUsers(ctx context.Context, params FollowUsersParams) (res *FollowUsersNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("followUsers"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/follow"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, FollowUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/follow"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "user_ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.UserIds != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.UserIds {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeFollowUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GenerateCallActionSignature invokes generateCallActionSignature operation.
//
// POST /v1/calls/action_signature/generate
func (c *Client) GenerateCallActionSignature(ctx context.Context, request *GenerateCallActionSignatureReq) (*CallActionSignatureResponse, error) {
	res, err := c.sendGenerateCallActionSignature(ctx, request)
	return res, err
}

func (c *Client) sendGenerateCallActionSignature(ctx context.Context, request *GenerateCallActionSignatureReq) (res *CallActionSignatureResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("generateCallActionSignature"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/calls/action_signature/generate"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GenerateCallActionSignatureOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/calls/action_signature/generate"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeGenerateCallActionSignatureRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGenerateCallActionSignatureResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetActiveCallPost invokes getActiveCallPost operation.
//
// GET /v1/posts/active_call
func (c *Client) GetActiveCallPost(ctx context.Context, params GetActiveCallPostParams) (*PostResponse, error) {
	res, err := c.sendGetActiveCallPost(ctx, params)
	return res, err
}

func (c *Client) sendGetActiveCallPost(ctx context.Context, params GetActiveCallPostParams) (res *PostResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getActiveCallPost"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/posts/active_call"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetActiveCallPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/posts/active_call"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "user_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetActiveCallPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetActiveFollowings invokes getActiveFollowings operation.
//
// GET /v1/users/active_followings
func (c *Client) GetActiveFollowings(ctx context.Context, params GetActiveFollowingsParams) (*ActiveFollowingsResponse, error) {
	res, err := c.sendGetActiveFollowings(ctx, params)
	return res, err
}

func (c *Client) sendGetActiveFollowings(ctx context.Context, params GetActiveFollowingsParams) (res *ActiveFollowingsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getActiveFollowings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/active_followings"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetActiveFollowingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/active_followings"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "only_online" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "only_online",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OnlyOnline.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_loggedin_at" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_loggedin_at",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromLoggedinAt.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetActiveFollowingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetAdditionalNotificationSetting invokes getAdditionalNotificationSetting operation.
//
// GET /v1/users/additonal_notification_setting
func (c *Client) GetAdditionalNotificationSetting(ctx context.Context) (*AdditionalSettingsResponse, error) {
	res, err := c.sendGetAdditionalNotificationSetting(ctx)
	return res, err
}

func (c *Client) sendGetAdditionalNotificationSetting(ctx context.Context) (res *AdditionalSettingsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getAdditionalNotificationSetting"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/additonal_notification_setting"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetAdditionalNotificationSettingOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/additonal_notification_setting"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetAdditionalNotificationSettingResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetAgoraRtmToken invokes getAgoraRtmToken operation.
//
// GET /v3/calls/{call_id}/agora_rtm_token
func (c *Client) GetAgoraRtmToken(ctx context.Context, params GetAgoraRtmTokenParams) (*RtmTokenResponse, error) {
	res, err := c.sendGetAgoraRtmToken(ctx, params)
	return res, err
}

func (c *Client) sendGetAgoraRtmToken(ctx context.Context, params GetAgoraRtmTokenParams) (res *RtmTokenResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getAgoraRtmToken"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v3/calls/{call_id}/agora_rtm_token"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetAgoraRtmTokenOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/agora_rtm_token"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetAgoraRtmTokenResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetAppConfig invokes getAppConfig operation.
//
// GET /api/apps/{app}
func (c *Client) GetAppConfig(ctx context.Context, params GetAppConfigParams) (*ApplicationConfigResponse, error) {
	res, err := c.sendGetAppConfig(ctx, params)
	return res, err
}

func (c *Client) sendGetAppConfig(ctx context.Context, params GetAppConfigParams) (res *ApplicationConfigResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getAppConfig"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/api/apps/{app}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetAppConfigOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/api/apps/"
	{
		// Encode "app" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "app",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.App))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetAppConfigResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetBannedWords invokes getBannedWords operation.
//
// GET /{countryApiValue}/api/v2/banned_words
func (c *Client) GetBannedWords(ctx context.Context, params GetBannedWordsParams) (*BanWordsResponse, error) {
	res, err := c.sendGetBannedWords(ctx, params)
	return res, err
}

func (c *Client) sendGetBannedWords(ctx context.Context, params GetBannedWordsParams) (res *BanWordsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBannedWords"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/{countryApiValue}/api/v2/banned_words"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetBannedWordsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/"
	{
		// Encode "countryApiValue" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "countryApiValue",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.CountryApiValue))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/api/v2/banned_words"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetBannedWordsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetBlockedUserIds invokes getBlockedUserIds operation.
//
// GET /v1/users/block_ids
func (c *Client) GetBlockedUserIds(ctx context.Context) (*BlockedUserIdsResponse, error) {
	res, err := c.sendGetBlockedUserIds(ctx)
	return res, err
}

func (c *Client) sendGetBlockedUserIds(ctx context.Context) (res *BlockedUserIdsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBlockedUserIds"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/block_ids"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetBlockedUserIdsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/block_ids"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetBlockedUserIdsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetBlockedUsers invokes getBlockedUsers operation.
//
// POST /v2/users/blocked
func (c *Client) GetBlockedUsers(ctx context.Context, request *SearchUsersRequest, params GetBlockedUsersParams) (*BlockedUsersResponse, error) {
	res, err := c.sendGetBlockedUsers(ctx, request, params)
	return res, err
}

func (c *Client) sendGetBlockedUsers(ctx context.Context, request *SearchUsersRequest, params GetBlockedUsersParams) (res *BlockedUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBlockedUsers"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/blocked"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetBlockedUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/blocked"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeGetBlockedUsersRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetBlockedUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetBookmarkedPosts invokes getBookmarkedPosts operation.
//
// GET /v1/users/{user_id}/bookmarks
func (c *Client) GetBookmarkedPosts(ctx context.Context, params GetBookmarkedPostsParams) (*PostsResponse, error) {
	res, err := c.sendGetBookmarkedPosts(ctx, params)
	return res, err
}

func (c *Client) sendGetBookmarkedPosts(ctx context.Context, params GetBookmarkedPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBookmarkedPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/{user_id}/bookmarks"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetBookmarkedPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/bookmarks"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetBookmarkedPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetBucketPresignedUrls invokes getBucketPresignedUrls operation.
//
// GET /v1/buckets/presigned_urls
func (c *Client) GetBucketPresignedUrls(ctx context.Context, params GetBucketPresignedUrlsParams) (*PresignedUrlsResponse, error) {
	res, err := c.sendGetBucketPresignedUrls(ctx, params)
	return res, err
}

func (c *Client) sendGetBucketPresignedUrls(ctx context.Context, params GetBucketPresignedUrlsParams) (res *PresignedUrlsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getBucketPresignedUrls"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/buckets/presigned_urls"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetBucketPresignedUrlsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/buckets/presigned_urls"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "file_names[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "file_names[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.FileNames != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.FileNames {
						if err := func() error {
							return e.EncodeValue(conv.StringToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetBucketPresignedUrlsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetCallBgms invokes getCallBgms operation.
//
// GET /v1/calls/bgm
func (c *Client) GetCallBgms(ctx context.Context) (*BgmsResponse, error) {
	res, err := c.sendGetCallBgms(ctx)
	return res, err
}

func (c *Client) sendGetCallBgms(ctx context.Context) (res *BgmsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallBgms"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/calls/bgm"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetCallBgmsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/calls/bgm"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetCallBgmsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetCallFollowersTimeline invokes getCallFollowersTimeline operation.
//
// GET /v2/posts/call_followers_timeline
func (c *Client) GetCallFollowersTimeline(ctx context.Context, params GetCallFollowersTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetCallFollowersTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetCallFollowersTimeline(ctx context.Context, params GetCallFollowersTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallFollowersTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/call_followers_timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetCallFollowersTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/call_followers_timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "category_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CategoryID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "call_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "call_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CallType.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "include_circle_call" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "include_circle_call",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.IncludeCircleCall.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "exclude_recent_gomimushi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "exclude_recent_gomimushi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ExcludeRecentGomimushi.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetCallFollowersTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetCallGiftHistory invokes getCallGiftHistory operation.
//
// GET /v1/calls/{call_id}/gift_transactions
func (c *Client) GetCallGiftHistory(ctx context.Context, params GetCallGiftHistoryParams) (*CallGiftHistoryResponse, error) {
	res, err := c.sendGetCallGiftHistory(ctx, params)
	return res, err
}

func (c *Client) sendGetCallGiftHistory(ctx context.Context, params GetCallGiftHistoryParams) (res *CallGiftHistoryResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallGiftHistory"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/calls/{call_id}/gift_transactions"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetCallGiftHistoryOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/gift_transactions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetCallGiftHistoryResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetCallTimeline invokes getCallTimeline operation.
//
// GET /v2/posts/call_timeline
func (c *Client) GetCallTimeline(ctx context.Context, params GetCallTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetCallTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetCallTimeline(ctx context.Context, params GetCallTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getCallTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/call_timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetCallTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/call_timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "category_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CategoryID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "call_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "call_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CallType.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "include_circle_call" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "include_circle_call",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.IncludeCircleCall.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "cross_generation" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "cross_generation",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CrossGeneration.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "exclude_recent_gomimushi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "exclude_recent_gomimushi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ExcludeRecentGomimushi.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetCallTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetChatMessages invokes getChatMessages operation.
//
// GET /v2/chat_rooms/{id}/messages
func (c *Client) GetChatMessages(ctx context.Context, params GetChatMessagesParams) (*MessagesResponse, error) {
	res, err := c.sendGetChatMessages(ctx, params)
	return res, err
}

func (c *Client) sendGetChatMessages(ctx context.Context, params GetChatMessagesParams) (res *MessagesResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatMessages"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/{id}/messages"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetChatMessagesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_message_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_message_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromMessageID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "to_message_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "to_message_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ToMessageID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetChatMessagesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetChatRequestCount invokes getChatRequestCount operation.
//
// GET /v1/chat_rooms/total_chat_request
func (c *Client) GetChatRequestCount(ctx context.Context) (*TotalChatRequestResponse, error) {
	res, err := c.sendGetChatRequestCount(ctx)
	return res, err
}

func (c *Client) sendGetChatRequestCount(ctx context.Context) (res *TotalChatRequestResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatRequestCount"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/total_chat_request"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetChatRequestCountOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/total_chat_request"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetChatRequestCountResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetChatRequests invokes getChatRequests operation.
//
// GET /v1/chat_rooms/request_list
func (c *Client) GetChatRequests(ctx context.Context, params GetChatRequestsParams) (*ChatRoomsResponse, error) {
	res, err := c.sendGetChatRequests(ctx, params)
	return res, err
}

func (c *Client) sendGetChatRequests(ctx context.Context, params GetChatRequestsParams) (res *ChatRoomsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatRequests"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/request_list"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetChatRequestsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/request_list"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetChatRequestsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetChatRoom invokes getChatRoom operation.
//
// GET /v2/chat_rooms/{id}
func (c *Client) GetChatRoom(ctx context.Context, params GetChatRoomParams) (*ChatRoomResponse, error) {
	res, err := c.sendGetChatRoom(ctx, params)
	return res, err
}

func (c *Client) sendGetChatRoom(ctx context.Context, params GetChatRoomParams) (res *ChatRoomResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatRoom"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetChatUnreadStatus invokes getChatUnreadStatus operation.
//
// GET /v1/chat_rooms/unread_status
func (c *Client) GetChatUnreadStatus(ctx context.Context, params GetChatUnreadStatusParams) (*UnreadStatusResponse, error) {
	res, err := c.sendGetChatUnreadStatus(ctx, params)
	return res, err
}

func (c *Client) sendGetChatUnreadStatus(ctx context.Context, params GetChatUnreadStatusParams) (res *UnreadStatusResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatUnreadStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/unread_status"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetChatUnreadStatusOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/unread_status"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_time" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTime.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetChatUnreadStatusResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetChatableFollowings invokes getChatableFollowings operation.
//
// POST /v1/users/followings/chatable
func (c *Client) GetChatableFollowings(ctx context.Context, request *SearchUsersRequest, params GetChatableFollowingsParams) (*FollowUsersResponse, error) {
	res, err := c.sendGetChatableFollowings(ctx, request, params)
	return res, err
}

func (c *Client) sendGetChatableFollowings(ctx context.Context, request *SearchUsersRequest, params GetChatableFollowingsParams) (res *FollowUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getChatableFollowings"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/followings/chatable"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetChatableFollowingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/followings/chatable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_follow_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_follow_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromFollowID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "order_by" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OrderBy.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeGetChatableFollowingsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetChatableFollowingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetConferenceCall invokes getConferenceCall operation.
//
// GET /v2/calls/conferences/{call_id}
func (c *Client) GetConferenceCall(ctx context.Context, params GetConferenceCallParams) (*ConferenceCallResponse, error) {
	res, err := c.sendGetConferenceCall(ctx, params)
	return res, err
}

func (c *Client) sendGetConferenceCall(ctx context.Context, params GetConferenceCallParams) (res *ConferenceCallResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getConferenceCall"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/calls/conferences/{call_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/calls/conferences/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetConferenceCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetConversation invokes getConversation operation.
//
// GET /v2/conversations/{id}
func (c *Client) GetConversation(ctx context.Context, params GetConversationParams) (*PostsResponse, error) {
	res, err := c.sendGetConversation(ctx, params)
	return res, err
}

func (c *Client) sendGetConversation(ctx context.Context, params GetConversationParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getConversation"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/conversations/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetConversationOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/conversations/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "thread_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "thread_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ThreadID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "reverse" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "reverse",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Reverse.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetConversationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetDefaultSettings invokes getDefaultSettings operation.
//
// GET /v1/users/default_settings
func (c *Client) GetDefaultSettings(ctx context.Context) (*DefaultSettingsResponse, error) {
	res, err := c.sendGetDefaultSettings(ctx)
	return res, err
}

func (c *Client) sendGetDefaultSettings(ctx context.Context) (res *DefaultSettingsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getDefaultSettings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/default_settings"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetDefaultSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/default_settings"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetDefaultSettingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetFollowRequestCount invokes getFollowRequestCount operation.
//
// GET /v2/users/follow_requests_count
func (c *Client) GetFollowRequestCount(ctx context.Context) (*FollowRequestCountResponse, error) {
	res, err := c.sendGetFollowRequestCount(ctx)
	return res, err
}

func (c *Client) sendGetFollowRequestCount(ctx context.Context) (res *FollowRequestCountResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowRequestCount"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/follow_requests_count"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetFollowRequestCountOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/follow_requests_count"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetFollowRequestCountResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetFollowRequests invokes getFollowRequests operation.
//
// GET /v2/users/follow_requests
func (c *Client) GetFollowRequests(ctx context.Context, params GetFollowRequestsParams) (*UsersByTimestampResponse, error) {
	res, err := c.sendGetFollowRequests(ctx, params)
	return res, err
}

func (c *Client) sendGetFollowRequests(ctx context.Context, params GetFollowRequestsParams) (res *UsersByTimestampResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowRequests"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/follow_requests"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetFollowRequestsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/follow_requests"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetFollowRequestsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetFollowingTimeline invokes getFollowingTimeline operation.
//
// GET /v2/posts/following_timeline
func (c *Client) GetFollowingTimeline(ctx context.Context, params GetFollowingTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetFollowingTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetFollowingTimeline(ctx context.Context, params GetFollowingTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowingTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/following_timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetFollowingTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/following_timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "only_root" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "only_root",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OnlyRoot.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "order_by" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OrderBy.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "mxn" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "mxn",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Mxn.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "reduce_selfie" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "reduce_selfie",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ReduceSelfie.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "custom_generation_range" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "custom_generation_range",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CustomGenerationRange.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetFollowingTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetFollowingsBornToday invokes getFollowingsBornToday operation.
//
// GET /v1/users/following_born_today
func (c *Client) GetFollowingsBornToday(ctx context.Context, params GetFollowingsBornTodayParams) (*UsersResponse, error) {
	res, err := c.sendGetFollowingsBornToday(ctx, params)
	return res, err
}

func (c *Client) sendGetFollowingsBornToday(ctx context.Context, params GetFollowingsBornTodayParams) (res *UsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFollowingsBornToday"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/following_born_today"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetFollowingsBornTodayOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/following_born_today"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "birthdate" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "birthdate",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Birthdate.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetFollowingsBornTodayResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetFootprints invokes getFootprints operation.
//
// GET /v3/users/footprints
func (c *Client) GetFootprints(ctx context.Context, params GetFootprintsParams) (*FootprintsResponse, error) {
	res, err := c.sendGetFootprints(ctx, params)
	return res, err
}

func (c *Client) sendGetFootprints(ctx context.Context, params GetFootprintsParams) (res *FootprintsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFootprints"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v3/users/footprints"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetFootprintsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/users/footprints"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "mode" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "mode",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Mode.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetFootprintsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetFreshUser invokes getFreshUser operation.
//
// GET /v2/users/fresh/{id}
func (c *Client) GetFreshUser(ctx context.Context, params GetFreshUserParams) (*UserResponse, error) {
	res, err := c.sendGetFreshUser(ctx, params)
	return res, err
}

func (c *Client) sendGetFreshUser(ctx context.Context, params GetFreshUserParams) (res *UserResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getFreshUser"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/fresh/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetFreshUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/users/fresh/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetFreshUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroup invokes getGroup operation.
//
// GET /v1/groups/{id}
func (c *Client) GetGroup(ctx context.Context, params GetGroupParams) (*GroupResponse, error) {
	res, err := c.sendGetGroup(ctx, params)
	return res, err
}

func (c *Client) sendGetGroup(ctx context.Context, params GetGroupParams) (res *GroupResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroup"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupBanList invokes getGroupBanList operation.
//
// GET /v1/groups/{id}/ban_list
func (c *Client) GetGroupBanList(ctx context.Context, params GetGroupBanListParams) (*UsersResponse, error) {
	res, err := c.sendGetGroupBanList(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupBanList(ctx context.Context, params GetGroupBanListParams) (res *UsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupBanList"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/ban_list"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupBanListOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/ban_list"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupBanListResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupCreateQuota invokes getGroupCreateQuota operation.
//
// GET /v1/groups/created_quota
func (c *Client) GetGroupCreateQuota(ctx context.Context) (*CreateQuotaResponse, error) {
	res, err := c.sendGetGroupCreateQuota(ctx)
	return res, err
}

func (c *Client) sendGetGroupCreateQuota(ctx context.Context) (res *CreateQuotaResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupCreateQuota"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/created_quota"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupCreateQuotaOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/groups/created_quota"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupCreateQuotaResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupGiftHistory invokes getGroupGiftHistory operation.
//
// GET /v1/groups/{group_id}/gift_history
func (c *Client) GetGroupGiftHistory(ctx context.Context, params GetGroupGiftHistoryParams) (*GroupGiftHistoryResponse, error) {
	res, err := c.sendGetGroupGiftHistory(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupGiftHistory(ctx context.Context, params GetGroupGiftHistoryParams) (res *GroupGiftHistoryResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupGiftHistory"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/gift_history"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupGiftHistoryOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/gift_history"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupGiftHistoryResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupGiftTransactions invokes getGroupGiftTransactions operation.
//
// GET /v1/groups/{group_id}/gift_transactions
func (c *Client) GetGroupGiftTransactions(ctx context.Context, params GetGroupGiftTransactionsParams) (*GiftTransactionsResponse, error) {
	res, err := c.sendGetGroupGiftTransactions(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupGiftTransactions(ctx context.Context, params GetGroupGiftTransactionsParams) (res *GiftTransactionsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupGiftTransactions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/gift_transactions"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupGiftTransactionsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/gift_transactions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupGiftTransactionsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupHighlights invokes getGroupHighlights operation.
//
// GET /v1/groups/{group_id}/highlights
func (c *Client) GetGroupHighlights(ctx context.Context, params GetGroupHighlightsParams) (*PostsResponse, error) {
	res, err := c.sendGetGroupHighlights(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupHighlights(ctx context.Context, params GetGroupHighlightsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupHighlights"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/highlights"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupHighlightsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/highlights"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupHighlightsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupMember invokes getGroupMember operation.
//
// GET /v1/groups/{id}/members/{userId}
func (c *Client) GetGroupMember(ctx context.Context, params GetGroupMemberParams) (*GroupUserResponse, error) {
	res, err := c.sendGetGroupMember(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupMember(ctx context.Context, params GetGroupMemberParams) (res *GroupUserResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupMember"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/members/{userId}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupMemberOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/members/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupMemberResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupMembers invokes getGroupMembers operation.
//
// GET /v2/groups/{id}/members
func (c *Client) GetGroupMembers(ctx context.Context, params GetGroupMembersParams) (*GroupUsersResponse, error) {
	res, err := c.sendGetGroupMembers(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupMembers(ctx context.Context, params GetGroupMembersParams) (res *GroupUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupMembers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/groups/{id}/members"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupMembersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/members"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupMembersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupNotificationSettings invokes getGroupNotificationSettings operation.
//
// GET /v2/notification_settings/groups/{id}
func (c *Client) GetGroupNotificationSettings(ctx context.Context, params GetGroupNotificationSettingsParams) (*GroupNotificationSettingsResponse, error) {
	res, err := c.sendGetGroupNotificationSettings(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupNotificationSettings(ctx context.Context, params GetGroupNotificationSettingsParams) (res *GroupNotificationSettingsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupNotificationSettings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/notification_settings/groups/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupNotificationSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/notification_settings/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupNotificationSettingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupReceivedGiftSenders invokes getGroupReceivedGiftSenders operation.
//
// GET /v1/groups/{group_id}/received_gifts/{gift_id}/senders
func (c *Client) GetGroupReceivedGiftSenders(ctx context.Context, params GetGroupReceivedGiftSendersParams) (*GiftSendersResponse, error) {
	res, err := c.sendGetGroupReceivedGiftSenders(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupReceivedGiftSenders(ctx context.Context, params GetGroupReceivedGiftSendersParams) (res *GiftSendersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupReceivedGiftSenders"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/received_gifts/{gift_id}/senders"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupReceivedGiftSendersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [5]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/received_gifts/"
	{
		// Encode "gift_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "gift_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GiftID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	pathParts[4] = "/senders"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "user_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Timestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupReceivedGiftSendersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupTimeline invokes getGroupTimeline operation.
//
// GET /v2/posts/group_timeline
func (c *Client) GetGroupTimeline(ctx context.Context, params GetGroupTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetGroupTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupTimeline(ctx context.Context, params GetGroupTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/group_timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/group_timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "reverse" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "reverse",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Reverse.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "post_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.PostType.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "only_root" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "only_root",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OnlyRoot.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetGroupUnreadStatus invokes getGroupUnreadStatus operation.
//
// GET /v1/groups/unread_status
func (c *Client) GetGroupUnreadStatus(ctx context.Context, params GetGroupUnreadStatusParams) (*UnreadStatusResponse, error) {
	res, err := c.sendGetGroupUnreadStatus(ctx, params)
	return res, err
}

func (c *Client) sendGetGroupUnreadStatus(ctx context.Context, params GetGroupUnreadStatusParams) (res *UnreadStatusResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getGroupUnreadStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/unread_status"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetGroupUnreadStatusOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/groups/unread_status"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_time" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTime.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetGroupUnreadStatusResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetHimaUsers invokes getHimaUsers operation.
//
// GET /v2/users/hima_users
func (c *Client) GetHimaUsers(ctx context.Context, params GetHimaUsersParams) (*HimaUsersResponse, error) {
	res, err := c.sendGetHimaUsers(ctx, params)
	return res, err
}

func (c *Client) sendGetHimaUsers(ctx context.Context, params GetHimaUsersParams) (res *HimaUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getHimaUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/hima_users"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetHimaUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/hima_users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_hima_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_hima_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromHimaID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetHimaUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetInvitableCallUsers invokes getInvitableCallUsers operation.
//
// GET /v1/calls/{call_id}/users/invitable
func (c *Client) GetInvitableCallUsers(ctx context.Context, params GetInvitableCallUsersParams) (*UsersByTimestampResponse, error) {
	res, err := c.sendGetInvitableCallUsers(ctx, params)
	return res, err
}

func (c *Client) sendGetInvitableCallUsers(ctx context.Context, params GetInvitableCallUsersParams) (res *UsersByTimestampResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getInvitableCallUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/calls/{call_id}/users/invitable"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetInvitableCallUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/users/invitable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user[nickname]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserNickname.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetInvitableCallUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetInvitableGroupUsers invokes getInvitableGroupUsers operation.
//
// GET /v1/groups/{group_id}/users/invitable
func (c *Client) GetInvitableGroupUsers(ctx context.Context, params GetInvitableGroupUsersParams) (*UsersByTimestampResponse, error) {
	res, err := c.sendGetInvitableGroupUsers(ctx, params)
	return res, err
}

func (c *Client) sendGetInvitableGroupUsers(ctx context.Context, params GetInvitableGroupUsersParams) (res *UsersByTimestampResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getInvitableGroupUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/users/invitable"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetInvitableGroupUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/users/invitable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user[nickname]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserNickname.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetInvitableGroupUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetJoinedGroupStatuses invokes getJoinedGroupStatuses operation.
//
// GET /v1/groups/joined_statuses
func (c *Client) GetJoinedGroupStatuses(ctx context.Context, params GetJoinedGroupStatusesParams) (GetJoinedGroupStatusesOK, error) {
	res, err := c.sendGetJoinedGroupStatuses(ctx, params)
	return res, err
}

func (c *Client) sendGetJoinedGroupStatuses(ctx context.Context, params GetJoinedGroupStatusesParams) (res GetJoinedGroupStatusesOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getJoinedGroupStatuses"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/joined_statuses"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetJoinedGroupStatusesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/groups/joined_statuses"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.Ids != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.Ids {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetJoinedGroupStatusesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetJoinedThreadStatuses invokes getJoinedThreadStatuses operation.
//
// GET /v1/threads/joined_statuses
func (c *Client) GetJoinedThreadStatuses(ctx context.Context, params GetJoinedThreadStatusesParams) (GetJoinedThreadStatusesOK, error) {
	res, err := c.sendGetJoinedThreadStatuses(ctx, params)
	return res, err
}

func (c *Client) sendGetJoinedThreadStatuses(ctx context.Context, params GetJoinedThreadStatusesParams) (res GetJoinedThreadStatusesOK, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getJoinedThreadStatuses"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/threads/joined_statuses"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetJoinedThreadStatusesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/threads/joined_statuses"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.Ids != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.Ids {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetJoinedThreadStatusesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetMainChatRooms invokes getMainChatRooms operation.
//
// GET /v1/chat_rooms/main_list
func (c *Client) GetMainChatRooms(ctx context.Context, params GetMainChatRoomsParams) (*ChatRoomsResponse, error) {
	res, err := c.sendGetMainChatRooms(ctx, params)
	return res, err
}

func (c *Client) sendGetMainChatRooms(ctx context.Context, params GetMainChatRoomsParams) (res *ChatRoomsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMainChatRooms"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/main_list"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetMainChatRoomsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/chat_rooms/main_list"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetMainChatRoomsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetMyPosts invokes getMyPosts operation.
//
// GET /v2/posts/mine
func (c *Client) GetMyPosts(ctx context.Context, params GetMyPostsParams) (*PostsResponse, error) {
	res, err := c.sendGetMyPosts(ctx, params)
	return res, err
}

func (c *Client) sendGetMyPosts(ctx context.Context, params GetMyPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMyPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/mine"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetMyPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/mine"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "include_group_post" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "include_group_post",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.IncludeGroupPost.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetMyPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetMyReviews invokes getMyReviews operation.
//
// GET /v1/users/reviews/mine
func (c *Client) GetMyReviews(ctx context.Context, params GetMyReviewsParams) (*ReviewsResponse, error) {
	res, err := c.sendGetMyReviews(ctx, params)
	return res, err
}

func (c *Client) sendGetMyReviews(ctx context.Context, params GetMyReviewsParams) (res *ReviewsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getMyReviews"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/reviews/mine"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetMyReviewsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/reviews/mine"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetMyReviewsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPhoneStatus invokes getPhoneStatus operation.
//
// GET /v1/calls/phone_status/{opponent_id}
func (c *Client) GetPhoneStatus(ctx context.Context, params GetPhoneStatusParams) (*CallStatusResponse, error) {
	res, err := c.sendGetPhoneStatus(ctx, params)
	return res, err
}

func (c *Client) sendGetPhoneStatus(ctx context.Context, params GetPhoneStatusParams) (res *CallStatusResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPhoneStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/calls/phone_status/{opponent_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPhoneStatusOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/calls/phone_status/"
	{
		// Encode "opponent_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "opponent_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.OpponentID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPhoneStatusResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPolicyAgreements invokes getPolicyAgreements operation.
//
// GET /v1/users/policy_agreements
func (c *Client) GetPolicyAgreements(ctx context.Context) (*PolicyAgreementsResponse, error) {
	res, err := c.sendGetPolicyAgreements(ctx)
	return res, err
}

func (c *Client) sendGetPolicyAgreements(ctx context.Context) (res *PolicyAgreementsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPolicyAgreements"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/policy_agreements"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPolicyAgreementsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/policy_agreements"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPolicyAgreementsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPopularWords invokes getPopularWords operation.
//
// GET /{countryApiValue}/api/apps/{app}/popular_words
func (c *Client) GetPopularWords(ctx context.Context, params GetPopularWordsParams) (*PopularWordsResponse, error) {
	res, err := c.sendGetPopularWords(ctx, params)
	return res, err
}

func (c *Client) sendGetPopularWords(ctx context.Context, params GetPopularWordsParams) (res *PopularWordsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPopularWords"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/{countryApiValue}/api/apps/{app}/popular_words"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPopularWordsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [5]string
	pathParts[0] = "/"
	{
		// Encode "countryApiValue" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "countryApiValue",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.CountryApiValue))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/api/apps/"
	{
		// Encode "app" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "app",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.App))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	pathParts[4] = "/popular_words"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPopularWordsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPost invokes getPost operation.
//
// GET /v2/posts/{id}
func (c *Client) GetPost(ctx context.Context, params GetPostParams) (*PostResponse, error) {
	res, err := c.sendGetPost(ctx, params)
	return res, err
}

func (c *Client) sendGetPost(ctx context.Context, params GetPostParams) (res *PostResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPost"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "Cache-Control",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CacheControl.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPostGiftTransactions invokes getPostGiftTransactions operation.
//
// GET /v1/posts/{post_id}/gift_transactions
func (c *Client) GetPostGiftTransactions(ctx context.Context, params GetPostGiftTransactionsParams) (*GiftTransactionsResponse, error) {
	res, err := c.sendGetPostGiftTransactions(ctx, params)
	return res, err
}

func (c *Client) sendGetPostGiftTransactions(ctx context.Context, params GetPostGiftTransactionsParams) (res *GiftTransactionsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostGiftTransactions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/posts/{post_id}/gift_transactions"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostGiftTransactionsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/posts/"
	{
		// Encode "post_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "post_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PostID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/gift_transactions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostGiftTransactionsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPostLikers invokes getPostLikers operation.
//
// GET /v1/posts/{id}/likers
func (c *Client) GetPostLikers(ctx context.Context, params GetPostLikersParams) (*PostLikersResponse, error) {
	res, err := c.sendGetPostLikers(ctx, params)
	return res, err
}

func (c *Client) sendGetPostLikers(ctx context.Context, params GetPostLikersParams) (res *PostLikersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostLikers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/posts/{id}/likers"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostLikersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/likers"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostLikersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPostReposts invokes getPostReposts operation.
//
// GET /v2/posts/{id}/reposts
func (c *Client) GetPostReposts(ctx context.Context, params GetPostRepostsParams) (*PostsResponse, error) {
	res, err := c.sendGetPostReposts(ctx, params)
	return res, err
}

func (c *Client) sendGetPostReposts(ctx context.Context, params GetPostRepostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostReposts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/{id}/reposts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostRepostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/reposts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostRepostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPostUrlMetadata invokes getPostUrlMetadata operation.
//
// GET /v2/posts/url_metadata
func (c *Client) GetPostUrlMetadata(ctx context.Context, params GetPostUrlMetadataParams) (*SharedUrl, error) {
	res, err := c.sendGetPostUrlMetadata(ctx, params)
	return res, err
}

func (c *Client) sendGetPostUrlMetadata(ctx context.Context, params GetPostUrlMetadataParams) (res *SharedUrl, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostUrlMetadata"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/url_metadata"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostUrlMetadataOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/url_metadata"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "url" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "url",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.URL.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostUrlMetadataResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPostsByIds invokes getPostsByIds operation.
//
// GET /v2/posts/multiple
func (c *Client) GetPostsByIds(ctx context.Context, params GetPostsByIdsParams) (*PostsResponse, error) {
	res, err := c.sendGetPostsByIds(ctx, params)
	return res, err
}

func (c *Client) sendGetPostsByIds(ctx context.Context, params GetPostsByIdsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostsByIds"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/multiple"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostsByIdsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/multiple"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "post_ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.PostIds != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.PostIds {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostsByIdsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetPostsByTag invokes getPostsByTag operation.
//
// GET /v2/posts/tags/{tag}
func (c *Client) GetPostsByTag(ctx context.Context, params GetPostsByTagParams) (*PostsResponse, error) {
	res, err := c.sendGetPostsByTag(ctx, params)
	return res, err
}

func (c *Client) sendGetPostsByTag(ctx context.Context, params GetPostsByTagParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getPostsByTag"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/tags/{tag}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetPostsByTagOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/posts/tags/"
	{
		// Encode "tag" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "tag",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Tag))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetPostsByTagResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetReceivedGiftSenders invokes getReceivedGiftSenders operation.
//
// GET /v1/received_gifts/{gift_id}/senders
func (c *Client) GetReceivedGiftSenders(ctx context.Context, params GetReceivedGiftSendersParams) (*GiftSendersResponse, error) {
	res, err := c.sendGetReceivedGiftSenders(ctx, params)
	return res, err
}

func (c *Client) sendGetReceivedGiftSenders(ctx context.Context, params GetReceivedGiftSendersParams) (res *GiftSendersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getReceivedGiftSenders"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/received_gifts/{gift_id}/senders"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetReceivedGiftSendersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/received_gifts/"
	{
		// Encode "gift_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "gift_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GiftID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/senders"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetReceivedGiftSendersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetReceivedGiftTransaction invokes getReceivedGiftTransaction operation.
//
// GET /v1/received_gifts/{gift_transaction_uuid}
func (c *Client) GetReceivedGiftTransaction(ctx context.Context, params GetReceivedGiftTransactionParams) (*GiftReceivedTransactionResponse, error) {
	res, err := c.sendGetReceivedGiftTransaction(ctx, params)
	return res, err
}

func (c *Client) sendGetReceivedGiftTransaction(ctx context.Context, params GetReceivedGiftTransactionParams) (res *GiftReceivedTransactionResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getReceivedGiftTransaction"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/received_gifts/{gift_transaction_uuid}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetReceivedGiftTransactionOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/received_gifts/"
	{
		// Encode "gift_transaction_uuid" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "gift_transaction_uuid",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.GiftTransactionUUID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetReceivedGiftTransactionResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRecentEngagementPosts invokes getRecentEngagementPosts operation.
//
// GET /v2/posts/recent_engagement
func (c *Client) GetRecentEngagementPosts(ctx context.Context, params GetRecentEngagementPostsParams) (*PostsResponse, error) {
	res, err := c.sendGetRecentEngagementPosts(ctx, params)
	return res, err
}

func (c *Client) sendGetRecentEngagementPosts(ctx context.Context, params GetRecentEngagementPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecentEngagementPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/recent_engagement"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRecentEngagementPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/recent_engagement"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRecentEngagementPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRecommendedFollowUsers invokes getRecommendedFollowUsers operation.
//
// GET /v1/users/{id}/follow_recommended
func (c *Client) GetRecommendedFollowUsers(ctx context.Context, params GetRecommendedFollowUsersParams) (*UsersResponse, error) {
	res, err := c.sendGetRecommendedFollowUsers(ctx, params)
	return res, err
}

func (c *Client) sendGetRecommendedFollowUsers(ctx context.Context, params GetRecommendedFollowUsersParams) (res *UsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecommendedFollowUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/{id}/follow_recommended"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRecommendedFollowUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/follow_recommended"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRecommendedFollowUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRecommendedPostTags invokes getRecommendedPostTags operation.
//
// POST /v1/posts/recommended_tag
func (c *Client) GetRecommendedPostTags(ctx context.Context, request *GetRecommendedPostTagsReq) (*PostTagsResponse, error) {
	res, err := c.sendGetRecommendedPostTags(ctx, request)
	return res, err
}

func (c *Client) sendGetRecommendedPostTags(ctx context.Context, request *GetRecommendedPostTagsReq) (res *PostTagsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecommendedPostTags"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/posts/recommended_tag"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRecommendedPostTagsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/posts/recommended_tag"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeGetRecommendedPostTagsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRecommendedPostTagsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRecommendedTimeline invokes getRecommendedTimeline operation.
//
// GET /v2/posts/recommended_timeline
func (c *Client) GetRecommendedTimeline(ctx context.Context, params GetRecommendedTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetRecommendedTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetRecommendedTimeline(ctx context.Context, params GetRecommendedTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRecommendedTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/recommended_timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRecommendedTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/recommended_timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "experiment_num" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "experiment_num",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ExperimentNum.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "variant_num" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "variant_num",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.VariantNum.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRecommendedTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRelatableGroups invokes getRelatableGroups operation.
//
// GET /v1/groups/{id}/relatable
func (c *Client) GetRelatableGroups(ctx context.Context, params GetRelatableGroupsParams) (*GroupsRelatedResponse, error) {
	res, err := c.sendGetRelatableGroups(ctx, params)
	return res, err
}

func (c *Client) sendGetRelatableGroups(ctx context.Context, params GetRelatableGroupsParams) (res *GroupsRelatedResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRelatableGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/relatable"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRelatableGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/relatable"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "keyword" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Keyword.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRelatableGroupsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRelatedGroups invokes getRelatedGroups operation.
//
// GET /v1/groups/{id}/related
func (c *Client) GetRelatedGroups(ctx context.Context, params GetRelatedGroupsParams) (*GroupsRelatedResponse, error) {
	res, err := c.sendGetRelatedGroups(ctx, params)
	return res, err
}

func (c *Client) sendGetRelatedGroups(ctx context.Context, params GetRelatedGroupsParams) (res *GroupsRelatedResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRelatedGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/related"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRelatedGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/related"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "keyword" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Keyword.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRelatedGroupsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetResetCounters invokes getResetCounters operation.
//
// GET /v1/users/reset_counters
func (c *Client) GetResetCounters(ctx context.Context) (*RefreshCounterRequestsResponse, error) {
	res, err := c.sendGetResetCounters(ctx)
	return res, err
}

func (c *Client) sendGetResetCounters(ctx context.Context) (res *RefreshCounterRequestsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getResetCounters"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/reset_counters"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetResetCountersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/reset_counters"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetResetCountersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetRootPosts invokes getRootPosts operation.
//
// GET /v2/conversations/root_posts
func (c *Client) GetRootPosts(ctx context.Context, params GetRootPostsParams) (*PostsResponse, error) {
	res, err := c.sendGetRootPosts(ctx, params)
	return res, err
}

func (c *Client) sendGetRootPosts(ctx context.Context, params GetRootPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getRootPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/conversations/root_posts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetRootPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/conversations/root_posts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.Ids != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.Ids {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetRootPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetThread invokes getThread operation.
//
// GET /v1/threads/{id}
func (c *Client) GetThread(ctx context.Context, params GetThreadParams) (*ThreadInfo, error) {
	res, err := c.sendGetThread(ctx, params)
	return res, err
}

func (c *Client) sendGetThread(ctx context.Context, params GetThreadParams) (res *ThreadInfo, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getThread"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/threads/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetThreadOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetThreadResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetThreadPosts invokes getThreadPosts operation.
//
// GET /v1/threads/{id}/posts
func (c *Client) GetThreadPosts(ctx context.Context, params GetThreadPostsParams) (*PostsResponse, error) {
	res, err := c.sendGetThreadPosts(ctx, params)
	return res, err
}

func (c *Client) sendGetThreadPosts(ctx context.Context, params GetThreadPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getThreadPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/threads/{id}/posts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetThreadPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/posts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "post_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.PostType.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetThreadPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetTimeline invokes getTimeline operation.
//
// GET /v2/posts/{noreply_mode}timeline
func (c *Client) GetTimeline(ctx context.Context, params GetTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetTimeline(ctx context.Context, params GetTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/{noreply_mode}timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/posts/"
	{
		// Encode "noreply_mode" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "noreply_mode",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(string(params.NoreplyMode)))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "order_by" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OrderBy.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "experiment_older_age_rules" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "experiment_older_age_rules",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ExperimentOlderAgeRules.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "mxn" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "mxn",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Mxn.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "en" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "en",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.En.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "vn" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "vn",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Vn.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "reduce_selfie" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "reduce_selfie",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ReduceSelfie.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "custom_generation_range" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "custom_generation_range",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CustomGenerationRange.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetTwoFactorAuthRequestInfo invokes getTwoFactorAuthRequestInfo operation.
//
// GET /v1/users/secret
func (c *Client) GetTwoFactorAuthRequestInfo(ctx context.Context) (*TwoStepAuthRequestInfoResponse, error) {
	res, err := c.sendGetTwoFactorAuthRequestInfo(ctx)
	return res, err
}

func (c *Client) sendGetTwoFactorAuthRequestInfo(ctx context.Context) (res *TwoStepAuthRequestInfoResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTwoFactorAuthRequestInfo"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/secret"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetTwoFactorAuthRequestInfoOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/secret"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetTwoFactorAuthRequestInfoResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetTwoFactorAuthStatus invokes getTwoFactorAuthStatus operation.
//
// GET /v1/users/secret/status
func (c *Client) GetTwoFactorAuthStatus(ctx context.Context) (*TwoFAStatusResponse, error) {
	res, err := c.sendGetTwoFactorAuthStatus(ctx)
	return res, err
}

func (c *Client) sendGetTwoFactorAuthStatus(ctx context.Context) (res *TwoFAStatusResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getTwoFactorAuthStatus"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/secret/status"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetTwoFactorAuthStatusOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/secret/status"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetTwoFactorAuthStatusResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUpdatedChatRooms invokes getUpdatedChatRooms operation.
//
// GET /v2/chat_rooms/update
func (c *Client) GetUpdatedChatRooms(ctx context.Context, params GetUpdatedChatRoomsParams) (*ChatRoomsResponse, error) {
	res, err := c.sendGetUpdatedChatRooms(ctx, params)
	return res, err
}

func (c *Client) sendGetUpdatedChatRooms(ctx context.Context, params GetUpdatedChatRoomsParams) (res *ChatRoomsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUpdatedChatRooms"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/update"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUpdatedChatRoomsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/chat_rooms/update"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_time" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_time",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTime.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUpdatedChatRoomsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUser invokes getUser operation.
//
// GET /v2/users/{id}
func (c *Client) GetUser(ctx context.Context, params GetUserParams) (*UserResponse, error) {
	res, err := c.sendGetUser(ctx, params)
	return res, err
}

func (c *Client) sendGetUser(ctx context.Context, params GetUserParams) (res *UserResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUser"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserActivities invokes getUserActivities operation.
//
// GET /api/v2/user_activities
func (c *Client) GetUserActivities(ctx context.Context, params GetUserActivitiesParams) (*ActivitiesResponse, error) {
	res, err := c.sendGetUserActivities(ctx, params)
	return res, err
}

func (c *Client) sendGetUserActivities(ctx context.Context, params GetUserActivitiesParams) (res *ActivitiesResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserActivities"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/api/v2/user_activities"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserActivitiesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v2/user_activities"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserActivitiesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserActivitiesV1 invokes getUserActivitiesV1 operation.
//
// GET /api/user_activities
func (c *Client) GetUserActivitiesV1(ctx context.Context, params GetUserActivitiesV1Params) (*ActivitiesResponse, error) {
	res, err := c.sendGetUserActivitiesV1(ctx, params)
	return res, err
}

func (c *Client) sendGetUserActivitiesV1(ctx context.Context, params GetUserActivitiesV1Params) (res *ActivitiesResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserActivitiesV1"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/api/user_activities"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserActivitiesV1Operation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/user_activities"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "important" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "important",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Important.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserActivitiesV1Response(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserByQr invokes getUserByQr operation.
//
// GET /v1/users/qr_codes/{qr}
func (c *Client) GetUserByQr(ctx context.Context, params GetUserByQrParams) (*UserResponse, error) {
	res, err := c.sendGetUserByQr(ctx, params)
	return res, err
}

func (c *Client) sendGetUserByQr(ctx context.Context, params GetUserByQrParams) (res *UserResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserByQr"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/qr_codes/{qr}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserByQrOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/users/qr_codes/"
	{
		// Encode "qr" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "qr",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.StringToString(params.Qr))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserByQrResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserCustomDefinitions invokes getUserCustomDefinitions operation.
//
// GET /v1/users/custom_definitions
func (c *Client) GetUserCustomDefinitions(ctx context.Context) (*UserCustomDefinitionsResponse, error) {
	res, err := c.sendGetUserCustomDefinitions(ctx)
	return res, err
}

func (c *Client) sendGetUserCustomDefinitions(ctx context.Context) (res *UserCustomDefinitionsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserCustomDefinitions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/custom_definitions"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserCustomDefinitionsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/custom_definitions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserCustomDefinitionsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserFollowers invokes getUserFollowers operation.
//
// GET /v3/users/{id}/followers
func (c *Client) GetUserFollowers(ctx context.Context, params GetUserFollowersParams) (*FollowUsersResponse, error) {
	res, err := c.sendGetUserFollowers(ctx, params)
	return res, err
}

func (c *Client) sendGetUserFollowers(ctx context.Context, params GetUserFollowersParams) (res *FollowUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserFollowers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v3/users/{id}/followers"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserFollowersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/followers"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "followed_by_me" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "followed_by_me",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FollowedByMe.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user[nickname]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserNickname.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserFollowersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserFollowings invokes getUserFollowings operation.
//
// GET /v3/users/{id}/followings
func (c *Client) GetUserFollowings(ctx context.Context, params GetUserFollowingsParams) (*FollowUsersResponse, error) {
	res, err := c.sendGetUserFollowings(ctx, params)
	return res, err
}

func (c *Client) sendGetUserFollowings(ctx context.Context, params GetUserFollowingsParams) (res *FollowUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserFollowings"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v3/users/{id}/followings"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserFollowingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/followings"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "order_by" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "order_by",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OrderBy.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user[nickname]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user[nickname]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserNickname.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserFollowingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserGiftTransactions invokes getUserGiftTransactions operation.
//
// GET /v1/users/{user_id}/gift_transactions
func (c *Client) GetUserGiftTransactions(ctx context.Context, params GetUserGiftTransactionsParams) (*GiftTransactionsResponse, error) {
	res, err := c.sendGetUserGiftTransactions(ctx, params)
	return res, err
}

func (c *Client) sendGetUserGiftTransactions(ctx context.Context, params GetUserGiftTransactionsParams) (res *GiftTransactionsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserGiftTransactions"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/{user_id}/gift_transactions"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserGiftTransactionsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/gift_transactions"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserGiftTransactionsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserGroupList invokes getUserGroupList operation.
//
// GET /v1/groups/user_group_list
func (c *Client) GetUserGroupList(ctx context.Context, params GetUserGroupListParams) (*GroupsResponse, error) {
	res, err := c.sendGetUserGroupList(ctx, params)
	return res, err
}

func (c *Client) sendGetUserGroupList(ctx context.Context, params GetUserGroupListParams) (res *GroupsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserGroupList"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/user_group_list"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserGroupListOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/groups/user_group_list"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "user_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserGroupListResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserInfo invokes getUserInfo operation.
//
// GET /v2/users/info/{id}
func (c *Client) GetUserInfo(ctx context.Context, params GetUserInfoParams) (*UserResponse, error) {
	res, err := c.sendGetUserInfo(ctx, params)
	return res, err
}

func (c *Client) sendGetUserInfo(ctx context.Context, params GetUserInfoParams) (res *UserResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserInfo"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/info/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserInfoOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/users/info/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserInfoResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserInterests invokes getUserInterests operation.
//
// GET /v1/users/interests
func (c *Client) GetUserInterests(ctx context.Context) (*UserInterestsResponse, error) {
	res, err := c.sendGetUserInterests(ctx)
	return res, err
}

func (c *Client) sendGetUserInterests(ctx context.Context) (res *UserInterestsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserInterests"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/interests"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserInterestsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/interests"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserInterestsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserPresignedUrl invokes getUserPresignedUrl operation.
//
// GET /v1/users/presigned_url
func (c *Client) GetUserPresignedUrl(ctx context.Context, params GetUserPresignedUrlParams) (*PresignedUrlResponse, error) {
	res, err := c.sendGetUserPresignedUrl(ctx, params)
	return res, err
}

func (c *Client) sendGetUserPresignedUrl(ctx context.Context, params GetUserPresignedUrlParams) (res *PresignedUrlResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserPresignedUrl"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/presigned_url"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserPresignedUrlOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/presigned_url"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "video_file_name" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "video_file_name",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.VideoFileName.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserPresignedUrlResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserReviews invokes getUserReviews operation.
//
// GET /v2/users/reviews/{id}
func (c *Client) GetUserReviews(ctx context.Context, params GetUserReviewsParams) (*ReviewsResponse, error) {
	res, err := c.sendGetUserReviews(ctx, params)
	return res, err
}

func (c *Client) sendGetUserReviews(ctx context.Context, params GetUserReviewsParams) (res *ReviewsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserReviews"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/reviews/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserReviewsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/users/reviews/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserReviewsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserTimeline invokes getUserTimeline operation.
//
// GET /v2/posts/user_timeline
func (c *Client) GetUserTimeline(ctx context.Context, params GetUserTimelineParams) (*PostsResponse, error) {
	res, err := c.sendGetUserTimeline(ctx, params)
	return res, err
}

func (c *Client) sendGetUserTimeline(ctx context.Context, params GetUserTimelineParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserTimeline"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/user_timeline"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserTimelineOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/user_timeline"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "user_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.UserID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "post_type" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post_type",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.PostType.Get(); ok {
				return e.EncodeValue(conv.StringToString(string(val)))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserTimelineResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUserTimestamp invokes getUserTimestamp operation.
//
// GET /v2/users/timestamp
func (c *Client) GetUserTimestamp(ctx context.Context) (*UserTimestampResponse, error) {
	res, err := c.sendGetUserTimestamp(ctx)
	return res, err
}

func (c *Client) sendGetUserTimestamp(ctx context.Context) (res *UserTimestampResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUserTimestamp"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/users/timestamp"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUserTimestampOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/timestamp"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUserTimestampResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetUsersByIds invokes getUsersByIds operation.
//
// GET /v1/users/list_id
func (c *Client) GetUsersByIds(ctx context.Context, params GetUsersByIdsParams) (*UsersResponse, error) {
	res, err := c.sendGetUsersByIds(ctx, params)
	return res, err
}

func (c *Client) sendGetUsersByIds(ctx context.Context, params GetUsersByIdsParams) (res *UsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getUsersByIds"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/list_id"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetUsersByIdsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/list_id"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "user_ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.UserIds != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.UserIds {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.XJwt.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetUsersByIdsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// GetWebSocketToken invokes getWebSocketToken operation.
//
// GET /v1/users/ws_token
func (c *Client) GetWebSocketToken(ctx context.Context) (*WebSocketTokenResponse, error) {
	res, err := c.sendGetWebSocketToken(ctx)
	return res, err
}

func (c *Client) sendGetWebSocketToken(ctx context.Context) (res *WebSocketTokenResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("getWebSocketToken"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/ws_token"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, GetWebSocketTokenOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/ws_token"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeGetWebSocketTokenResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// HideChats invokes hideChats operation.
//
// POST /v1/hidden/chats
func (c *Client) HideChats(ctx context.Context, request *HideChatsReq) error {
	_, err := c.sendHideChats(ctx, request)
	return err
}

func (c *Client) sendHideChats(ctx context.Context, request *HideChatsReq) (res *HideChatsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("hideChats"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/hidden/chats"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, HideChatsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/chats"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeHideChatsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeHideChatsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// HideUsers invokes hideUsers operation.
//
// POST /v1/hidden/users
func (c *Client) HideUsers(ctx context.Context, request *HideUsersReq) error {
	_, err := c.sendHideUsers(ctx, request)
	return err
}

func (c *Client) sendHideUsers(ctx context.Context, request *HideUsersReq) (res *HideUsersNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("hideUsers"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/hidden/users"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, HideUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeHideUsersRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeHideUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InviteToCall invokes inviteToCall operation.
//
// POST /v2/calls/invite
func (c *Client) InviteToCall(ctx context.Context, request *InviteToCallReq) error {
	_, err := c.sendInviteToCall(ctx, request)
	return err
}

func (c *Client) sendInviteToCall(ctx context.Context, request *InviteToCallReq) (res *InviteToCallNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/calls/invite"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, InviteToCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/calls/invite"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeInviteToCallRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeInviteToCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InviteToChatRoom invokes inviteToChatRoom operation.
//
// POST /v2/chat_rooms/{id}/invite
func (c *Client) InviteToChatRoom(ctx context.Context, request *InviteToChatRoomReq, params InviteToChatRoomParams) error {
	_, err := c.sendInviteToChatRoom(ctx, request, params)
	return err
}

func (c *Client) sendInviteToChatRoom(ctx context.Context, request *InviteToChatRoomReq, params InviteToChatRoomParams) (res *InviteToChatRoomNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/{id}/invite"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, InviteToChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/invite"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeInviteToChatRoomRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeInviteToChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InviteToConferenceCall invokes inviteToConferenceCall operation.
//
// POST /v1/calls/conference_calls/{call_id}/invite
func (c *Client) InviteToConferenceCall(ctx context.Context, request *InviteToConferenceCallReq, params InviteToConferenceCallParams) error {
	_, err := c.sendInviteToConferenceCall(ctx, request, params)
	return err
}

func (c *Client) sendInviteToConferenceCall(ctx context.Context, request *InviteToConferenceCallReq, params InviteToConferenceCallParams) (res *InviteToConferenceCallNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/calls/conference_calls/{call_id}/invite"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, InviteToConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/calls/conference_calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/invite"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeInviteToConferenceCallRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeInviteToConferenceCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// InviteToGroup invokes inviteToGroup operation.
//
// POST /v1/groups/{id}/invite
func (c *Client) InviteToGroup(ctx context.Context, request *InviteToGroupReq, params InviteToGroupParams) error {
	_, err := c.sendInviteToGroup(ctx, request, params)
	return err
}

func (c *Client) sendInviteToGroup(ctx context.Context, request *InviteToGroupReq, params InviteToGroupParams) (res *InviteToGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("inviteToGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/invite"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, InviteToGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/invite"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeInviteToGroupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeInviteToGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// JoinGroup invokes joinGroup operation.
//
// POST /v1/groups/{id}/join
func (c *Client) JoinGroup(ctx context.Context, params JoinGroupParams) error {
	_, err := c.sendJoinGroup(ctx, params)
	return err
}

func (c *Client) sendJoinGroup(ctx context.Context, params JoinGroupParams) (res *JoinGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("joinGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/join"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, JoinGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/join"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeJoinGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// KickFromChatRoom invokes kickFromChatRoom operation.
//
// POST /v2/chat_rooms/{id}/kick
func (c *Client) KickFromChatRoom(ctx context.Context, request *KickFromChatRoomReq, params KickFromChatRoomParams) error {
	_, err := c.sendKickFromChatRoom(ctx, request, params)
	return err
}

func (c *Client) sendKickFromChatRoom(ctx context.Context, request *KickFromChatRoomReq, params KickFromChatRoomParams) (res *KickFromChatRoomNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("kickFromChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/{id}/kick"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, KickFromChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/kick"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeKickFromChatRoomRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeKickFromChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// KickFromConferenceCall invokes kickFromConferenceCall operation.
//
// POST /v3/calls/conference_calls/{call_id}/kick
func (c *Client) KickFromConferenceCall(ctx context.Context, request *KickFromConferenceCallReq, params KickFromConferenceCallParams) (*CallActionSignatureResponse, error) {
	res, err := c.sendKickFromConferenceCall(ctx, request, params)
	return res, err
}

func (c *Client) sendKickFromConferenceCall(ctx context.Context, request *KickFromConferenceCallReq, params KickFromConferenceCallParams) (res *CallActionSignatureResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("kickFromConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/calls/conference_calls/{call_id}/kick"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, KickFromConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/calls/conference_calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/kick"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeKickFromConferenceCallRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeKickFromConferenceCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LeaveAgoraChannel invokes leaveAgoraChannel operation.
//
// POST /v1/calls/leave_agora_channel
func (c *Client) LeaveAgoraChannel(ctx context.Context, request *LeaveAgoraChannelReq) error {
	_, err := c.sendLeaveAgoraChannel(ctx, request)
	return err
}

func (c *Client) sendLeaveAgoraChannel(ctx context.Context, request *LeaveAgoraChannelReq) (res *LeaveAgoraChannelNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("leaveAgoraChannel"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/calls/leave_agora_channel"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LeaveAgoraChannelOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/calls/leave_agora_channel"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeLeaveAgoraChannelRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeLeaveAgoraChannelResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LeaveConferenceCall invokes leaveConferenceCall operation.
//
// POST /v1/calls/leave_conference_call
func (c *Client) LeaveConferenceCall(ctx context.Context, request *LeaveConferenceCallReq) error {
	_, err := c.sendLeaveConferenceCall(ctx, request)
	return err
}

func (c *Client) sendLeaveConferenceCall(ctx context.Context, request *LeaveConferenceCallReq) (res *LeaveConferenceCallNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("leaveConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/calls/leave_conference_call"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LeaveConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/calls/leave_conference_call"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeLeaveConferenceCallRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeLeaveConferenceCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LeaveGroup invokes leaveGroup operation.
//
// DELETE /v1/groups/{id}/leave
func (c *Client) LeaveGroup(ctx context.Context, params LeaveGroupParams) error {
	_, err := c.sendLeaveGroup(ctx, params)
	return err
}

func (c *Client) sendLeaveGroup(ctx context.Context, params LeaveGroupParams) (res *LeaveGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("leaveGroup"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/leave"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LeaveGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/leave"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeLeaveGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LikePosts invokes likePosts operation.
//
// POST /v2/posts/like
func (c *Client) LikePosts(ctx context.Context, request *LikePostsReq) (*LikePostsResponse, error) {
	res, err := c.sendLikePosts(ctx, request)
	return res, err
}

func (c *Client) sendLikePosts(ctx context.Context, request *LikePostsReq) (res *LikePostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("likePosts"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/posts/like"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LikePostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/like"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeLikePostsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeLikePostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListGameApps invokes listGameApps operation.
//
// GET /v1/games/apps
func (c *Client) ListGameApps(ctx context.Context, params ListGameAppsParams) (*GamesResponse, error) {
	res, err := c.sendListGameApps(ctx, params)
	return res, err
}

func (c *Client) sendListGameApps(ctx context.Context, params ListGameAppsParams) (res *GamesResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGameApps"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/games/apps"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListGameAppsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/games/apps"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.Ids != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.Ids {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListGameAppsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListGameWalkthroughs invokes listGameWalkthroughs operation.
//
// GET /v1/games/apps/{app_id}/walkthroughs
func (c *Client) ListGameWalkthroughs(ctx context.Context, params ListGameWalkthroughsParams) ([]Walkthrough, error) {
	res, err := c.sendListGameWalkthroughs(ctx, params)
	return res, err
}

func (c *Client) sendListGameWalkthroughs(ctx context.Context, params ListGameWalkthroughsParams) (res []Walkthrough, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGameWalkthroughs"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/games/apps/{app_id}/walkthroughs"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListGameWalkthroughsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/games/apps/"
	{
		// Encode "app_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "app_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.AppID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/walkthroughs"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListGameWalkthroughsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListGenres invokes listGenres operation.
//
// GET /v1/genres
func (c *Client) ListGenres(ctx context.Context, params ListGenresParams) (*GenresResponse, error) {
	res, err := c.sendListGenres(ctx, params)
	return res, err
}

func (c *Client) sendListGenres(ctx context.Context, params ListGenresParams) (res *GenresResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGenres"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/genres"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListGenresOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/genres"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListGenresResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListGifts invokes listGifts operation.
//
// GET /v2/gifts
func (c *Client) ListGifts(ctx context.Context, params ListGiftsParams) (*GiftsResponse, error) {
	res, err := c.sendListGifts(ctx, params)
	return res, err
}

func (c *Client) sendListGifts(ctx context.Context, params ListGiftsParams) (res *GiftsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGifts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/gifts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListGiftsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/gifts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "currency" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "currency",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Currency.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListGiftsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListGroupCategories invokes listGroupCategories operation.
//
// GET /v1/groups/categories
func (c *Client) ListGroupCategories(ctx context.Context, params ListGroupCategoriesParams) (*GroupCategoriesResponse, error) {
	res, err := c.sendListGroupCategories(ctx, params)
	return res, err
}

func (c *Client) sendListGroupCategories(ctx context.Context, params ListGroupCategoriesParams) (res *GroupCategoriesResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGroupCategories"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/groups/categories"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListGroupCategoriesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/groups/categories"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "page" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "page",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Page.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListGroupCategoriesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListGroups invokes listGroups operation.
//
// GET /v2/groups
func (c *Client) ListGroups(ctx context.Context, params ListGroupsParams) (*GroupsResponse, error) {
	res, err := c.sendListGroups(ctx, params)
	return res, err
}

func (c *Client) sendListGroups(ctx context.Context, params ListGroupsParams) (res *GroupsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/groups"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/groups"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_category_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupCategoryID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "keyword" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Keyword.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "sub_category_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "sub_category_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.SubCategoryID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListGroupsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListHiddenChats invokes listHiddenChats operation.
//
// GET /v1/hidden/chats
func (c *Client) ListHiddenChats(ctx context.Context, params ListHiddenChatsParams) (*ChatRoomsResponse, error) {
	res, err := c.sendListHiddenChats(ctx, params)
	return res, err
}

func (c *Client) sendListHiddenChats(ctx context.Context, params ListHiddenChatsParams) (res *ChatRoomsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listHiddenChats"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/hidden/chats"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListHiddenChatsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/chats"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListHiddenChatsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListHiddenUsers invokes listHiddenUsers operation.
//
// GET /v1/hidden/users
func (c *Client) ListHiddenUsers(ctx context.Context, params ListHiddenUsersParams) (*HiddenResponse, error) {
	res, err := c.sendListHiddenUsers(ctx, params)
	return res, err
}

func (c *Client) sendListHiddenUsers(ctx context.Context, params ListHiddenUsersParams) (res *HiddenResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listHiddenUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/hidden/users"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListHiddenUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListHiddenUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListMuteKeywords invokes listMuteKeywords operation.
//
// GET /v1/hidden/words
func (c *Client) ListMuteKeywords(ctx context.Context) (*MuteKeywordResponse, error) {
	res, err := c.sendListMuteKeywords(ctx)
	return res, err
}

func (c *Client) sendListMuteKeywords(ctx context.Context) (res *MuteKeywordResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listMuteKeywords"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/hidden/words"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListMuteKeywordsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/words"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListMuteKeywordsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListMutedGroupUsers invokes listMutedGroupUsers operation.
//
// GET /v1/group_mute/{id}/muted_users
func (c *Client) ListMutedGroupUsers(ctx context.Context, params ListMutedGroupUsersParams) (*GroupMuteUsersResponse, error) {
	res, err := c.sendListMutedGroupUsers(ctx, params)
	return res, err
}

func (c *Client) sendListMutedGroupUsers(ctx context.Context, params ListMutedGroupUsersParams) (res *GroupMuteUsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listMutedGroupUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/group_mute/{id}/muted_users"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListMutedGroupUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/group_mute/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/muted_users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "keyword" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Keyword.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "cursor" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "cursor",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Cursor.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "size" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "size",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Size.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListMutedGroupUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListMyGroups invokes listMyGroups operation.
//
// GET /v2/groups/mine
func (c *Client) ListMyGroups(ctx context.Context, params ListMyGroupsParams) (*GroupsResponse, error) {
	res, err := c.sendListMyGroups(ctx, params)
	return res, err
}

func (c *Client) sendListMyGroups(ctx context.Context, params ListMyGroupsParams) (res *GroupsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listMyGroups"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/groups/mine"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListMyGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/groups/mine"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListMyGroupsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListReceivedGifts invokes listReceivedGifts operation.
//
// GET /v2/received_gifts
func (c *Client) ListReceivedGifts(ctx context.Context) (*GiftReceivedResponse, error) {
	res, err := c.sendListReceivedGifts(ctx)
	return res, err
}

func (c *Client) sendListReceivedGifts(ctx context.Context) (res *GiftReceivedResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listReceivedGifts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/received_gifts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListReceivedGiftsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/received_gifts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListReceivedGiftsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListReceivedGiftsV1 invokes listReceivedGiftsV1 operation.
//
// GET /v1/received_gifts
func (c *Client) ListReceivedGiftsV1(ctx context.Context, params ListReceivedGiftsV1Params) (*GiftReceivedResponse, error) {
	res, err := c.sendListReceivedGiftsV1(ctx, params)
	return res, err
}

func (c *Client) sendListReceivedGiftsV1(ctx context.Context, params ListReceivedGiftsV1Params) (res *GiftReceivedResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listReceivedGiftsV1"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/received_gifts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListReceivedGiftsV1Operation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/received_gifts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListReceivedGiftsV1Response(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListStickerPacks invokes listStickerPacks operation.
//
// GET /v2/sticker_packs
func (c *Client) ListStickerPacks(ctx context.Context) (*StickerPacksResponse, error) {
	res, err := c.sendListStickerPacks(ctx)
	return res, err
}

func (c *Client) sendListStickerPacks(ctx context.Context) (res *StickerPacksResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listStickerPacks"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/sticker_packs"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListStickerPacksOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/sticker_packs"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListStickerPacksResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ListThreads invokes listThreads operation.
//
// GET /v1/threads
func (c *Client) ListThreads(ctx context.Context, params ListThreadsParams) (*GroupThreadListResponse, error) {
	res, err := c.sendListThreads(ctx, params)
	return res, err
}

func (c *Client) sendListThreads(ctx context.Context, params ListThreadsParams) (res *GroupThreadListResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("listThreads"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/threads"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ListThreadsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/threads"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.From.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "join_status" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "join_status",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.JoinStatus.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeListThreadsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// LoginWithEmail invokes loginWithEmail operation.
//
// POST /v3/users/login_with_email
func (c *Client) LoginWithEmail(ctx context.Context, request *LoginEmailUserRequest) (*LoginUserResponse, error) {
	res, err := c.sendLoginWithEmail(ctx, request)
	return res, err
}

func (c *Client) sendLoginWithEmail(ctx context.Context, request *LoginEmailUserRequest) (res *LoginUserResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("loginWithEmail"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/users/login_with_email"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LoginWithEmailOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/users/login_with_email"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeLoginWithEmailRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeLoginWithEmailResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Logout invokes logout operation.
//
// POST /v1/users/logout
func (c *Client) Logout(ctx context.Context, request *LogoutReq) error {
	_, err := c.sendLogout(ctx, request)
	return err
}

func (c *Client) sendLogout(ctx context.Context, request *LogoutReq) (res *LogoutNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("logout"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/logout"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, LogoutOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/logout"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeLogoutRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeLogoutResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MovePostToSpecificThread invokes movePostToSpecificThread operation.
//
// PUT /v3/posts/{id}/move_to_thread/{thread_id}
func (c *Client) MovePostToSpecificThread(ctx context.Context, params MovePostToSpecificThreadParams) (*ThreadInfo, error) {
	res, err := c.sendMovePostToSpecificThread(ctx, params)
	return res, err
}

func (c *Client) sendMovePostToSpecificThread(ctx context.Context, params MovePostToSpecificThreadParams) (res *ThreadInfo, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("movePostToSpecificThread"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v3/posts/{id}/move_to_thread/{thread_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MovePostToSpecificThreadOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v3/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/move_to_thread/"
	{
		// Encode "thread_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "thread_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ThreadID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeMovePostToSpecificThreadResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MovePostToThread invokes movePostToThread operation.
//
// POST /v3/posts/{id}/move_to_thread
func (c *Client) MovePostToThread(ctx context.Context, request *MovePostToThreadReq, params MovePostToThreadParams) (*ThreadInfo, error) {
	res, err := c.sendMovePostToThread(ctx, request, params)
	return res, err
}

func (c *Client) sendMovePostToThread(ctx context.Context, request *MovePostToThreadReq, params MovePostToThreadParams) (res *ThreadInfo, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("movePostToThread"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/posts/{id}/move_to_thread"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MovePostToThreadOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/move_to_thread"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeMovePostToThreadRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeMovePostToThreadResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// MuteGroupUser invokes muteGroupUser operation.
//
// POST /v1/group_mute/{id}/mute/{user_id}
func (c *Client) MuteGroupUser(ctx context.Context, params MuteGroupUserParams) error {
	_, err := c.sendMuteGroupUser(ctx, params)
	return err
}

func (c *Client) sendMuteGroupUser(ctx context.Context, params MuteGroupUserParams) (res *MuteGroupUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("muteGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/group_mute/{id}/mute/{user_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, MuteGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/group_mute/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/mute/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeMuteGroupUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// OauthToken invokes oauthToken operation.
//
// POST /api/v1/oauth/token
func (c *Client) OauthToken(ctx context.Context, request *OauthTokenReq) (*TokenResponse, error) {
	res, err := c.sendOauthToken(ctx, request)
	return res, err
}

func (c *Client) sendOauthToken(ctx context.Context, request *OauthTokenReq) (res *TokenResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("oauthToken"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/api/v1/oauth/token"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, OauthTokenOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/api/v1/oauth/token"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeOauthTokenRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeOauthTokenResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PinChatRoom invokes pinChatRoom operation.
//
// POST /v1/chat_rooms/{id}/pinned
func (c *Client) PinChatRoom(ctx context.Context, params PinChatRoomParams) error {
	_, err := c.sendPinChatRoom(ctx, params)
	return err
}

func (c *Client) sendPinChatRoom(ctx context.Context, params PinChatRoomParams) (res *PinChatRoomNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/{id}/pinned"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PinChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/pinned"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePinChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PinGroup invokes pinGroup operation.
//
// POST /v1/pinned/groups
func (c *Client) PinGroup(ctx context.Context, request *PinGroupReq) error {
	_, err := c.sendPinGroup(ctx, request)
	return err
}

func (c *Client) sendPinGroup(ctx context.Context, request *PinGroupReq) (res *PinGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/pinned/groups"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PinGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/pinned/groups"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePinGroupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePinGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PinGroupHighlightPost invokes pinGroupHighlightPost operation.
//
// PUT /v1/groups/{group_id}/highlights/{post_id}
func (c *Client) PinGroupHighlightPost(ctx context.Context, params PinGroupHighlightPostParams) error {
	_, err := c.sendPinGroupHighlightPost(ctx, params)
	return err
}

func (c *Client) sendPinGroupHighlightPost(ctx context.Context, params PinGroupHighlightPostParams) (res *PinGroupHighlightPostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinGroupHighlightPost"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/highlights/{post_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PinGroupHighlightPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/highlights/"
	{
		// Encode "post_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "post_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PostID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePinGroupHighlightPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PinGroupPost invokes pinGroupPost operation.
//
// PUT /v2/posts/group_pinned_post
func (c *Client) PinGroupPost(ctx context.Context, request *PinGroupPostReq) error {
	_, err := c.sendPinGroupPost(ctx, request)
	return err
}

func (c *Client) sendPinGroupPost(ctx context.Context, request *PinGroupPostReq) (res *PinGroupPostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinGroupPost"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v2/posts/group_pinned_post"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PinGroupPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/group_pinned_post"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePinGroupPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePinGroupPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PinPost invokes pinPost operation.
//
// POST /v1/pinned/posts
func (c *Client) PinPost(ctx context.Context, request *PinPostReq) error {
	_, err := c.sendPinPost(ctx, request)
	return err
}

func (c *Client) sendPinPost(ctx context.Context, request *PinPostReq) (res *PinPostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/pinned/posts"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PinPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/pinned/posts"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePinPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePinPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PinReview invokes pinReview operation.
//
// POST /v1/pinned/reviews
func (c *Client) PinReview(ctx context.Context, request *PinReviewReq) error {
	_, err := c.sendPinReview(ctx, request)
	return err
}

func (c *Client) sendPinReview(ctx context.Context, request *PinReviewReq) (res *PinReviewNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pinReview"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/pinned/reviews"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PinReviewOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/pinned/reviews"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodePinReviewRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePinReviewResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// PingAlive invokes pingAlive operation.
//
// POST /v1/users/alive
func (c *Client) PingAlive(ctx context.Context) error {
	_, err := c.sendPingAlive(ctx)
	return err
}

func (c *Client) sendPingAlive(ctx context.Context) (res *PingAliveNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("pingAlive"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/alive"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, PingAliveOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/alive"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodePingAliveResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadChatAttachments invokes readChatAttachments operation.
//
// POST /v1/chat_rooms/{id}/attachments_read
func (c *Client) ReadChatAttachments(ctx context.Context, request *ReadAttachmentRequest, params ReadChatAttachmentsParams) error {
	_, err := c.sendReadChatAttachments(ctx, request, params)
	return err
}

func (c *Client) sendReadChatAttachments(ctx context.Context, request *ReadAttachmentRequest, params ReadChatAttachmentsParams) (res *ReadChatAttachmentsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readChatAttachments"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/{id}/attachments_read"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReadChatAttachmentsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/attachments_read"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReadChatAttachmentsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReadChatAttachmentsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadChatMessage invokes readChatMessage operation.
//
// POST /v2/chat_rooms/{id}/messages/{message_id}/read
func (c *Client) ReadChatMessage(ctx context.Context, params ReadChatMessageParams) error {
	_, err := c.sendReadChatMessage(ctx, params)
	return err
}

func (c *Client) sendReadChatMessage(ctx context.Context, params ReadChatMessageParams) (res *ReadChatMessageNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readChatMessage"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/{id}/messages/{message_id}/read"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReadChatMessageOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [5]string
	pathParts[0] = "/v2/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages/"
	{
		// Encode "message_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "message_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.MessageID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	pathParts[4] = "/read"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReadChatMessageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReadChatVideos invokes readChatVideos operation.
//
// POST /v1/chat_rooms/{id}/videos_read
func (c *Client) ReadChatVideos(ctx context.Context, request *ReadChatVideosReq, params ReadChatVideosParams) error {
	_, err := c.sendReadChatVideos(ctx, request, params)
	return err
}

func (c *Client) sendReadChatVideos(ctx context.Context, request *ReadChatVideosReq, params ReadChatVideosParams) (res *ReadChatVideosNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("readChatVideos"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/{id}/videos_read"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReadChatVideosOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/videos_read"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReadChatVideosRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReadChatVideosResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveChatRoomBackground invokes removeChatRoomBackground operation.
//
// DELETE /v2/chat_rooms/{id}/background
func (c *Client) RemoveChatRoomBackground(ctx context.Context, params RemoveChatRoomBackgroundParams) error {
	_, err := c.sendRemoveChatRoomBackground(ctx, params)
	return err
}

func (c *Client) sendRemoveChatRoomBackground(ctx context.Context, params RemoveChatRoomBackgroundParams) (res *RemoveChatRoomBackgroundNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeChatRoomBackground"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v2/chat_rooms/{id}/background"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveChatRoomBackgroundOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/background"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveChatRoomBackgroundResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveCoverImage invokes removeCoverImage operation.
//
// POST /v2/users/remove_cover_image
func (c *Client) RemoveCoverImage(ctx context.Context) error {
	_, err := c.sendRemoveCoverImage(ctx)
	return err
}

func (c *Client) sendRemoveCoverImage(ctx context.Context) (res *RemoveCoverImageNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeCoverImage"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/remove_cover_image"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveCoverImageOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/remove_cover_image"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveCoverImageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveGroupCover invokes removeGroupCover operation.
//
// DELETE /v3/groups/{id}/cover
func (c *Client) RemoveGroupCover(ctx context.Context, params RemoveGroupCoverParams) error {
	_, err := c.sendRemoveGroupCover(ctx, params)
	return err
}

func (c *Client) sendRemoveGroupCover(ctx context.Context, params RemoveGroupCoverParams) (res *RemoveGroupCoverNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeGroupCover"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v3/groups/{id}/cover"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveGroupCoverOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/cover"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveGroupCoverResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveGroupDeputies invokes removeGroupDeputies operation.
//
// DELETE /v1/groups/{id}/deputize
func (c *Client) RemoveGroupDeputies(ctx context.Context, params RemoveGroupDeputiesParams) error {
	_, err := c.sendRemoveGroupDeputies(ctx, params)
	return err
}

func (c *Client) sendRemoveGroupDeputies(ctx context.Context, params RemoveGroupDeputiesParams) (res *RemoveGroupDeputiesNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeGroupDeputies"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/deputize"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveGroupDeputiesOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/deputize"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveGroupDeputiesResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveGroupIcon invokes removeGroupIcon operation.
//
// DELETE /v3/groups/{id}/icon
func (c *Client) RemoveGroupIcon(ctx context.Context, params RemoveGroupIconParams) error {
	_, err := c.sendRemoveGroupIcon(ctx, params)
	return err
}

func (c *Client) sendRemoveGroupIcon(ctx context.Context, params RemoveGroupIconParams) (res *RemoveGroupIconNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeGroupIcon"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v3/groups/{id}/icon"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveGroupIconOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/icon"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveGroupIconResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveProfilePhoto invokes removeProfilePhoto operation.
//
// POST /v2/users/remove_profile_photo
func (c *Client) RemoveProfilePhoto(ctx context.Context) error {
	_, err := c.sendRemoveProfilePhoto(ctx)
	return err
}

func (c *Client) sendRemoveProfilePhoto(ctx context.Context) (res *RemoveProfilePhotoNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeProfilePhoto"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/remove_profile_photo"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveProfilePhotoOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/remove_profile_photo"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveProfilePhotoResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveRelatedGroups invokes removeRelatedGroups operation.
//
// DELETE /v1/groups/{id}/related
func (c *Client) RemoveRelatedGroups(ctx context.Context, params RemoveRelatedGroupsParams) error {
	_, err := c.sendRemoveRelatedGroups(ctx, params)
	return err
}

func (c *Client) sendRemoveRelatedGroups(ctx context.Context, params RemoveRelatedGroupsParams) (res *RemoveRelatedGroupsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeRelatedGroups"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/related"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveRelatedGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/related"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "related_group_id[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "related_group_id[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.RelatedGroupID != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.RelatedGroupID {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveRelatedGroupsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RemoveThreadMember invokes removeThreadMember operation.
//
// DELETE /v1/threads/{thread_id}/members/{id}
func (c *Client) RemoveThreadMember(ctx context.Context, params RemoveThreadMemberParams) error {
	_, err := c.sendRemoveThreadMember(ctx, params)
	return err
}

func (c *Client) sendRemoveThreadMember(ctx context.Context, params RemoveThreadMemberParams) (res *RemoveThreadMemberNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("removeThreadMember"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/threads/{thread_id}/members/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RemoveThreadMemberOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "thread_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "thread_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ThreadID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/members/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRemoveThreadMemberResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReplyToReview invokes replyToReview operation.
//
// POST /v2/users/reviews/{id}
func (c *Client) ReplyToReview(ctx context.Context, request *ReplyToReviewReq, params ReplyToReviewParams) error {
	_, err := c.sendReplyToReview(ctx, request, params)
	return err
}

func (c *Client) sendReplyToReview(ctx context.Context, request *ReplyToReviewReq, params ReplyToReviewParams) (res *ReplyToReviewNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("replyToReview"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/reviews/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReplyToReviewOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/users/reviews/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReplyToReviewRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReplyToReviewResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReportChatRoom invokes reportChatRoom operation.
//
// POST /v3/chat_rooms/{chat_room_id}/report
func (c *Client) ReportChatRoom(ctx context.Context, request *ReportChatRoomReq, params ReportChatRoomParams) error {
	_, err := c.sendReportChatRoom(ctx, request, params)
	return err
}

func (c *Client) sendReportChatRoom(ctx context.Context, request *ReportChatRoomReq, params ReportChatRoomParams) (res *ReportChatRoomNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/chat_rooms/{chat_room_id}/report"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReportChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/chat_rooms/"
	{
		// Encode "chat_room_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "chat_room_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ChatRoomID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/report"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReportChatRoomRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReportChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReportGroup invokes reportGroup operation.
//
// POST /v3/groups/{group_id}/report
func (c *Client) ReportGroup(ctx context.Context, request *ReportGroupReq, params ReportGroupParams) error {
	_, err := c.sendReportGroup(ctx, request, params)
	return err
}

func (c *Client) sendReportGroup(ctx context.Context, request *ReportGroupReq, params ReportGroupParams) (res *ReportGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/groups/{group_id}/report"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReportGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/report"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReportGroupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReportGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReportPost invokes reportPost operation.
//
// POST /v3/posts/{post_id}/report
func (c *Client) ReportPost(ctx context.Context, request *ReportPostReq, params ReportPostParams) error {
	_, err := c.sendReportPost(ctx, request, params)
	return err
}

func (c *Client) sendReportPost(ctx context.Context, request *ReportPostReq, params ReportPostParams) (res *ReportPostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportPost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/posts/{post_id}/report"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReportPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/posts/"
	{
		// Encode "post_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "post_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PostID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/report"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReportPostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReportPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ReportUser invokes reportUser operation.
//
// POST /v3/users/{user_id}/report
func (c *Client) ReportUser(ctx context.Context, request *ReportUserReq, params ReportUserParams) error {
	_, err := c.sendReportUser(ctx, request, params)
	return err
}

func (c *Client) sendReportUser(ctx context.Context, request *ReportUserReq, params ReportUserParams) (res *ReportUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("reportUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/users/{user_id}/report"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ReportUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/report"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeReportUserRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeReportUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// Repost invokes repost operation.
//
// POST /v3/posts/repost
func (c *Client) Repost(ctx context.Context, request *RepostReq, params RepostParams) (*CreatePostResponse, error) {
	res, err := c.sendRepost(ctx, request, params)
	return res, err
}

func (c *Client) sendRepost(ctx context.Context, request *RepostReq, params RepostParams) (res *CreatePostResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("repost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/posts/repost"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RepostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/posts/repost"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeRepostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "EncodeHeaderParams"
	h := uri.NewHeaderEncoder(r.Header)
	{
		cfg := uri.HeaderParameterEncodingConfig{
			Name:    "X-Jwt",
			Explode: false,
		}
		if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.XJwt.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode header")
		}
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRepostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RequestEmailVerification invokes requestEmailVerification operation.
//
// POST /v1/email_verification_urls
func (c *Client) RequestEmailVerification(ctx context.Context, request *RequestEmailVerificationReq) (*CommonUrlResponse, error) {
	res, err := c.sendRequestEmailVerification(ctx, request)
	return res, err
}

func (c *Client) sendRequestEmailVerification(ctx context.Context, request *RequestEmailVerificationReq) (res *CommonUrlResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("requestEmailVerification"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/email_verification_urls"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RequestEmailVerificationOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/email_verification_urls"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeRequestEmailVerificationRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRequestEmailVerificationResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RequestFollow invokes requestFollow operation.
//
// POST /v2/users/{target_id}/follow_request
func (c *Client) RequestFollow(ctx context.Context, request *RequestFollowReq, params RequestFollowParams) error {
	_, err := c.sendRequestFollow(ctx, request, params)
	return err
}

func (c *Client) sendRequestFollow(ctx context.Context, request *RequestFollowReq, params RequestFollowParams) (res *RequestFollowNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("requestFollow"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/{target_id}/follow_request"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RequestFollowOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/users/"
	{
		// Encode "target_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "target_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.TargetID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/follow_request"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeRequestFollowRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRequestFollowResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// RequestGroupWalkthrough invokes requestGroupWalkthrough operation.
//
// POST /v1/groups/{id}/request_walkthrough
func (c *Client) RequestGroupWalkthrough(ctx context.Context, params RequestGroupWalkthroughParams) error {
	_, err := c.sendRequestGroupWalkthrough(ctx, params)
	return err
}

func (c *Client) sendRequestGroupWalkthrough(ctx context.Context, params RequestGroupWalkthroughParams) (res *RequestGroupWalkthroughNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("requestGroupWalkthrough"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/request_walkthrough"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, RequestGroupWalkthroughOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/request_walkthrough"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeRequestGroupWalkthroughResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ResendConfirmEmail invokes resendConfirmEmail operation.
//
// POST /v2/users/resend_confirm_email
func (c *Client) ResendConfirmEmail(ctx context.Context) error {
	_, err := c.sendResendConfirmEmail(ctx)
	return err
}

func (c *Client) sendResendConfirmEmail(ctx context.Context) (res *ResendConfirmEmailNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resendConfirmEmail"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/resend_confirm_email"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ResendConfirmEmailOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/users/resend_confirm_email"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeResendConfirmEmailResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ResetCounters invokes resetCounters operation.
//
// POST /v1/users/reset_counters
func (c *Client) ResetCounters(ctx context.Context, request *ResetCountersReq) error {
	_, err := c.sendResetCounters(ctx, request)
	return err
}

func (c *Client) sendResetCounters(ctx context.Context, request *ResetCountersReq) (res *ResetCountersNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resetCounters"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/reset_counters"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ResetCountersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/reset_counters"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeResetCountersRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeResetCountersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ResetPassword invokes resetPassword operation.
//
// PUT /v1/users/reset_password
func (c *Client) ResetPassword(ctx context.Context, request *ResetPasswordReq) error {
	_, err := c.sendResetPassword(ctx, request)
	return err
}

func (c *Client) sendResetPassword(ctx context.Context, request *ResetPasswordReq) (res *ResetPasswordNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("resetPassword"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/reset_password"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ResetPasswordOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/reset_password"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeResetPasswordRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeResetPasswordResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SearchGroupPosts invokes searchGroupPosts operation.
//
// GET /v2/groups/{id}/posts/search
func (c *Client) SearchGroupPosts(ctx context.Context, params SearchGroupPostsParams) (*PostsResponse, error) {
	res, err := c.sendSearchGroupPosts(ctx, params)
	return res, err
}

func (c *Client) sendSearchGroupPosts(ctx context.Context, params SearchGroupPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchGroupPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/groups/{id}/posts/search"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SearchGroupPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/posts/search"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "keyword" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Keyword.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "only_thread_posts" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "only_thread_posts",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OnlyThreadPosts.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeSearchGroupPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SearchPosts invokes searchPosts operation.
//
// GET /v2/posts/search
func (c *Client) SearchPosts(ctx context.Context, params SearchPostsParams) (*PostsResponse, error) {
	res, err := c.sendSearchPosts(ctx, params)
	return res, err
}

func (c *Client) sendSearchPosts(ctx context.Context, params SearchPostsParams) (res *PostsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchPosts"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/posts/search"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SearchPostsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/search"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "keyword" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "keyword",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Keyword.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "post_owner_scope" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "post_owner_scope",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.PostOwnerScope.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "only_media" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "only_media",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.OnlyMedia.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_post_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_post_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromPostID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "number" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "number",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Number.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeSearchPostsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SearchUsers invokes searchUsers operation.
//
// GET /v1/users/search
func (c *Client) SearchUsers(ctx context.Context, params SearchUsersParams) (*UsersResponse, error) {
	res, err := c.sendSearchUsers(ctx, params)
	return res, err
}

func (c *Client) sendSearchUsers(ctx context.Context, params SearchUsersParams) (res *UsersResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("searchUsers"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v1/users/search"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SearchUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/search"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "gender" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "gender",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Gender.Get(); ok {
				return e.EncodeValue(conv.Int32ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "nickname" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "nickname",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Nickname.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "title" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "title",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Title.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "biography" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "biography",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Biography.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "from_timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "from_timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.FromTimestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "similar_age" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "similar_age",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.SimilarAge.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "not_recent_gomimushi" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "not_recent_gomimushi",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.NotRecentGomimushi.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "recently_created" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "recently_created",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.RecentlyCreated.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "same_prefecture" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "same_prefecture",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.SamePrefecture.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "save_recent_search" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "save_recent_search",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.SaveRecentSearch.Get(); ok {
				return e.EncodeValue(conv.BoolToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeSearchUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SendChatMessage invokes sendChatMessage operation.
//
// POST /v3/chat_rooms/{id}/messages/new
func (c *Client) SendChatMessage(ctx context.Context, request *SendChatMessageReq, params SendChatMessageParams) (*MessageResponse, error) {
	res, err := c.sendSendChatMessage(ctx, request, params)
	return res, err
}

func (c *Client) sendSendChatMessage(ctx context.Context, request *SendChatMessageReq, params SendChatMessageParams) (res *MessageResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("sendChatMessage"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/chat_rooms/{id}/messages/new"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SendChatMessageOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/messages/new"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeSendChatMessageRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeSendChatMessageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SetGroupTitle invokes setGroupTitle operation.
//
// POST /v1/groups/{id}/set_title
func (c *Client) SetGroupTitle(ctx context.Context, request *SetGroupTitleReq, params SetGroupTitleParams) error {
	_, err := c.sendSetGroupTitle(ctx, request, params)
	return err
}

func (c *Client) sendSetGroupTitle(ctx context.Context, request *SetGroupTitleReq, params SetGroupTitleParams) (res *SetGroupTitleNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setGroupTitle"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/set_title"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SetGroupTitleOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/set_title"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeSetGroupTitleRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeSetGroupTitleResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// SetHima invokes setHima operation.
//
// POST /v1/users/hima
func (c *Client) SetHima(ctx context.Context) error {
	_, err := c.sendSetHima(ctx)
	return err
}

func (c *Client) sendSetHima(ctx context.Context) (res *SetHimaNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("setHima"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/hima"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, SetHimaOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/hima"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeSetHimaResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// StartConferenceCall invokes startConferenceCall operation.
//
// POST /v2/calls/start_conference_call
func (c *Client) StartConferenceCall(ctx context.Context, request *StartConferenceCallReq) (*ConferenceCallResponse, error) {
	res, err := c.sendStartConferenceCall(ctx, request)
	return res, err
}

func (c *Client) sendStartConferenceCall(ctx context.Context, request *StartConferenceCallReq) (res *ConferenceCallResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("startConferenceCall"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/calls/start_conference_call"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, StartConferenceCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/calls/start_conference_call"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeStartConferenceCallRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeStartConferenceCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TakeOverGroup invokes takeOverGroup operation.
//
// POST /v1/groups/{id}/take_over
func (c *Client) TakeOverGroup(ctx context.Context, params TakeOverGroupParams) error {
	_, err := c.sendTakeOverGroup(ctx, params)
	return err
}

func (c *Client) sendTakeOverGroup(ctx context.Context, params TakeOverGroupParams) (res *TakeOverGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("takeOverGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/take_over"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TakeOverGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/take_over"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeTakeOverGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// TransferGroup invokes transferGroup operation.
//
// POST /v3/groups/{id}/transfer
func (c *Client) TransferGroup(ctx context.Context, request *TransferGroupReq, params TransferGroupParams) error {
	_, err := c.sendTransferGroup(ctx, request, params)
	return err
}

func (c *Client) sendTransferGroup(ctx context.Context, request *TransferGroupReq, params TransferGroupParams) (res *TransferGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("transferGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/groups/{id}/transfer"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, TransferGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/transfer"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeTransferGroupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeTransferGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnbanGroupUser invokes unbanGroupUser operation.
//
// POST /v1/groups/{id}/unban/{userId}
func (c *Client) UnbanGroupUser(ctx context.Context, params UnbanGroupUserParams) error {
	_, err := c.sendUnbanGroupUser(ctx, params)
	return err
}

func (c *Client) sendUnbanGroupUser(ctx context.Context, params UnbanGroupUserParams) (res *UnbanGroupUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unbanGroupUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/unban/{userId}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnbanGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/unban/"
	{
		// Encode "userId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "userId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnbanGroupUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnblockUser invokes unblockUser operation.
//
// POST /v2/users/{id}/unblock
func (c *Client) UnblockUser(ctx context.Context, params UnblockUserParams) error {
	_, err := c.sendUnblockUser(ctx, params)
	return err
}

func (c *Client) sendUnblockUser(ctx context.Context, params UnblockUserParams) (res *UnblockUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unblockUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/{id}/unblock"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnblockUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/unblock"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnblockUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnfollowUser invokes unfollowUser operation.
//
// POST /v2/users/{id}/unfollow
func (c *Client) UnfollowUser(ctx context.Context, params UnfollowUserParams) error {
	_, err := c.sendUnfollowUser(ctx, params)
	return err
}

func (c *Client) sendUnfollowUser(ctx context.Context, params UnfollowUserParams) (res *UnfollowUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unfollowUser"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/users/{id}/unfollow"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnfollowUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/users/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/unfollow"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnfollowUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnhideChats invokes unhideChats operation.
//
// DELETE /v1/hidden/chats
func (c *Client) UnhideChats(ctx context.Context, params UnhideChatsParams) error {
	_, err := c.sendUnhideChats(ctx, params)
	return err
}

func (c *Client) sendUnhideChats(ctx context.Context, params UnhideChatsParams) (res *UnhideChatsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unhideChats"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/hidden/chats"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnhideChatsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/chats"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "chat_room_ids" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "chat_room_ids",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ChatRoomIds.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnhideChatsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnhideUsers invokes unhideUsers operation.
//
// DELETE /v1/hidden/users
func (c *Client) UnhideUsers(ctx context.Context, params UnhideUsersParams) error {
	_, err := c.sendUnhideUsers(ctx, params)
	return err
}

func (c *Client) sendUnhideUsers(ctx context.Context, params UnhideUsersParams) (res *UnhideUsersNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unhideUsers"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/hidden/users"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnhideUsersOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/hidden/users"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "user_ids[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "user_ids[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.UserIds != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.UserIds {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnhideUsersResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnlikePost invokes unlikePost operation.
//
// POST /v1/posts/{id}/unlike
func (c *Client) UnlikePost(ctx context.Context, params UnlikePostParams) error {
	_, err := c.sendUnlikePost(ctx, params)
	return err
}

func (c *Client) sendUnlikePost(ctx context.Context, params UnlikePostParams) (res *UnlikePostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unlikePost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/posts/{id}/unlike"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnlikePostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/unlike"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnlikePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnmuteGroupUser invokes unmuteGroupUser operation.
//
// DELETE /v1/group_mute/{id}/unmute/{user_id}
func (c *Client) UnmuteGroupUser(ctx context.Context, params UnmuteGroupUserParams) error {
	_, err := c.sendUnmuteGroupUser(ctx, params)
	return err
}

func (c *Client) sendUnmuteGroupUser(ctx context.Context, params UnmuteGroupUserParams) (res *UnmuteGroupUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unmuteGroupUser"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/group_mute/{id}/unmute/{user_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnmuteGroupUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/group_mute/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/unmute/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnmuteGroupUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnpinChatRoom invokes unpinChatRoom operation.
//
// DELETE /v1/chat_rooms/{id}/pinned
func (c *Client) UnpinChatRoom(ctx context.Context, params UnpinChatRoomParams) error {
	_, err := c.sendUnpinChatRoom(ctx, params)
	return err
}

func (c *Client) sendUnpinChatRoom(ctx context.Context, params UnpinChatRoomParams) (res *UnpinChatRoomNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinChatRoom"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/chat_rooms/{id}/pinned"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnpinChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/pinned"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnpinChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnpinGroup invokes unpinGroup operation.
//
// DELETE /v1/pinned/groups/{id}
func (c *Client) UnpinGroup(ctx context.Context, params UnpinGroupParams) error {
	_, err := c.sendUnpinGroup(ctx, params)
	return err
}

func (c *Client) sendUnpinGroup(ctx context.Context, params UnpinGroupParams) (res *UnpinGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinGroup"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/pinned/groups/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnpinGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/pinned/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnpinGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnpinGroupHighlightPost invokes unpinGroupHighlightPost operation.
//
// DELETE /v1/groups/{group_id}/highlights/{post_id}
func (c *Client) UnpinGroupHighlightPost(ctx context.Context, params UnpinGroupHighlightPostParams) error {
	_, err := c.sendUnpinGroupHighlightPost(ctx, params)
	return err
}

func (c *Client) sendUnpinGroupHighlightPost(ctx context.Context, params UnpinGroupHighlightPostParams) (res *UnpinGroupHighlightPostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinGroupHighlightPost"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/highlights/{post_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnpinGroupHighlightPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/highlights/"
	{
		// Encode "post_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "post_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.PostID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnpinGroupHighlightPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnpinGroupPost invokes unpinGroupPost operation.
//
// DELETE /v2/posts/group_pinned_post
func (c *Client) UnpinGroupPost(ctx context.Context, params UnpinGroupPostParams) error {
	_, err := c.sendUnpinGroupPost(ctx, params)
	return err
}

func (c *Client) sendUnpinGroupPost(ctx context.Context, params UnpinGroupPostParams) (res *UnpinGroupPostNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinGroupPost"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v2/posts/group_pinned_post"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnpinGroupPostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/posts/group_pinned_post"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "group_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "group_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.GroupID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnpinGroupPostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UnpinReview invokes unpinReview operation.
//
// DELETE /v1/pinned/reviews/{id}
func (c *Client) UnpinReview(ctx context.Context, params UnpinReviewParams) error {
	_, err := c.sendUnpinReview(ctx, params)
	return err
}

func (c *Client) sendUnpinReview(ctx context.Context, params UnpinReviewParams) (res *UnpinReviewNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("unpinReview"),
		semconv.HTTPRequestMethodKey.String("DELETE"),
		semconv.URLTemplateKey.String("/v1/pinned/reviews/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UnpinReviewOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/pinned/reviews/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "DELETE", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUnpinReviewResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateAdditionalNotificationSetting invokes updateAdditionalNotificationSetting operation.
//
// POST /v1/users/additonal_notification_setting
func (c *Client) UpdateAdditionalNotificationSetting(ctx context.Context, request *UpdateAdditionalNotificationSettingReq) error {
	_, err := c.sendUpdateAdditionalNotificationSetting(ctx, request)
	return err
}

func (c *Client) sendUpdateAdditionalNotificationSetting(ctx context.Context, request *UpdateAdditionalNotificationSettingReq) (res *UpdateAdditionalNotificationSettingNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateAdditionalNotificationSetting"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/users/additonal_notification_setting"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateAdditionalNotificationSettingOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/additonal_notification_setting"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateAdditionalNotificationSettingRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateAdditionalNotificationSettingResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateCall invokes updateCall operation.
//
// PUT /v1/calls/{call_id}
func (c *Client) UpdateCall(ctx context.Context, request *UpdateCallReq, params UpdateCallParams) error {
	_, err := c.sendUpdateCall(ctx, request, params)
	return err
}

func (c *Client) sendUpdateCall(ctx context.Context, request *UpdateCallReq, params UpdateCallParams) (res *UpdateCallNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateCall"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/calls/{call_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateCallOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateCallRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateCallResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateCallUser invokes updateCallUser operation.
//
// PUT /v1/calls/{call_id}/users/{user_id}
func (c *Client) UpdateCallUser(ctx context.Context, request *UpdateCallUserReq, params UpdateCallUserParams) error {
	_, err := c.sendUpdateCallUser(ctx, request, params)
	return err
}

func (c *Client) sendUpdateCallUser(ctx context.Context, request *UpdateCallUserReq, params UpdateCallUserParams) (res *UpdateCallUserNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateCallUser"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/calls/{call_id}/users/{user_id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateCallUserOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [4]string
	pathParts[0] = "/v1/calls/"
	{
		// Encode "call_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "call_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.CallID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/users/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateCallUserRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateCallUserResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateChatRoom invokes updateChatRoom operation.
//
// POST /v3/chat_rooms/{id}/edit
func (c *Client) UpdateChatRoom(ctx context.Context, request *UpdateChatRoomReq, params UpdateChatRoomParams) error {
	_, err := c.sendUpdateChatRoom(ctx, request, params)
	return err
}

func (c *Client) sendUpdateChatRoom(ctx context.Context, request *UpdateChatRoomReq, params UpdateChatRoomParams) (res *UpdateChatRoomNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateChatRoom"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/chat_rooms/{id}/edit"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateChatRoomOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/edit"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateChatRoomRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateChatRoomResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateChatRoomNotificationSettings invokes updateChatRoomNotificationSettings operation.
//
// POST /v2/notification_settings/chat_rooms/{id}
func (c *Client) UpdateChatRoomNotificationSettings(ctx context.Context, request *UpdateChatRoomNotificationSettingsReq, params UpdateChatRoomNotificationSettingsParams) (*NotificationSettingResponse, error) {
	res, err := c.sendUpdateChatRoomNotificationSettings(ctx, request, params)
	return res, err
}

func (c *Client) sendUpdateChatRoomNotificationSettings(ctx context.Context, request *UpdateChatRoomNotificationSettingsReq, params UpdateChatRoomNotificationSettingsParams) (res *NotificationSettingResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateChatRoomNotificationSettings"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/notification_settings/chat_rooms/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateChatRoomNotificationSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/notification_settings/chat_rooms/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateChatRoomNotificationSettingsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateChatRoomNotificationSettingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateGroup invokes updateGroup operation.
//
// POST /v3/groups/{id}/update
func (c *Client) UpdateGroup(ctx context.Context, request *UpdateGroupReq, params UpdateGroupParams) (*GroupResponse, error) {
	res, err := c.sendUpdateGroup(ctx, request, params)
	return res, err
}

func (c *Client) sendUpdateGroup(ctx context.Context, request *UpdateGroupReq, params UpdateGroupParams) (res *GroupResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/groups/{id}/update"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v3/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/update"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateGroupRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateGroupNotificationSettings invokes updateGroupNotificationSettings operation.
//
// POST /v2/notification_settings/groups/{id}
func (c *Client) UpdateGroupNotificationSettings(ctx context.Context, request *UpdateGroupNotificationSettingsReq, params UpdateGroupNotificationSettingsParams) (*AdditionalSettingsResponse, error) {
	res, err := c.sendUpdateGroupNotificationSettings(ctx, request, params)
	return res, err
}

func (c *Client) sendUpdateGroupNotificationSettings(ctx context.Context, request *UpdateGroupNotificationSettingsReq, params UpdateGroupNotificationSettingsParams) (res *AdditionalSettingsResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateGroupNotificationSettings"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/notification_settings/groups/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateGroupNotificationSettingsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v2/notification_settings/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateGroupNotificationSettingsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateGroupNotificationSettingsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateLanguage invokes updateLanguage operation.
//
// PUT /v1/users/language
func (c *Client) UpdateLanguage(ctx context.Context, request *UpdateLanguageReq) error {
	_, err := c.sendUpdateLanguage(ctx, request)
	return err
}

func (c *Client) sendUpdateLanguage(ctx context.Context, request *UpdateLanguageReq) (res *UpdateLanguageNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateLanguage"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/language"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateLanguageOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/language"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateLanguageRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateLanguageResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateLogin invokes updateLogin operation.
//
// POST /v3/users/login_update
func (c *Client) UpdateLogin(ctx context.Context, request *UpdateLoginReq) (*LoginUpdateResponse, error) {
	res, err := c.sendUpdateLogin(ctx, request)
	return res, err
}

func (c *Client) sendUpdateLogin(ctx context.Context, request *UpdateLoginReq) (res *LoginUpdateResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateLogin"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v3/users/login_update"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateLoginOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v3/users/login_update"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateLoginRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateLoginResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdatePost invokes updatePost operation.
//
// PUT /v3/posts/{id}
func (c *Client) UpdatePost(ctx context.Context, request *UpdatePostReq, params UpdatePostParams) (*Post, error) {
	res, err := c.sendUpdatePost(ctx, request, params)
	return res, err
}

func (c *Client) sendUpdatePost(ctx context.Context, request *UpdatePostReq, params UpdatePostParams) (res *Post, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updatePost"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v3/posts/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdatePostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v3/posts/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdatePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdatePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateRelatedGroups invokes updateRelatedGroups operation.
//
// PUT /v1/groups/{id}/related
func (c *Client) UpdateRelatedGroups(ctx context.Context, params UpdateRelatedGroupsParams) error {
	_, err := c.sendUpdateRelatedGroups(ctx, params)
	return err
}

func (c *Client) sendUpdateRelatedGroups(ctx context.Context, params UpdateRelatedGroupsParams) (res *UpdateRelatedGroupsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateRelatedGroups"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/related"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateRelatedGroupsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/related"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "related_group_id[]" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "related_group_id[]",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if params.RelatedGroupID != nil {
				return e.EncodeArray(func(e uri.Encoder) error {
					for i, item := range params.RelatedGroupID {
						if err := func() error {
							return e.EncodeValue(conv.Int64ToString(item))
						}(); err != nil {
							return errors.Wrapf(err, "[%d]", i)
						}
					}
					return nil
				})
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateRelatedGroupsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateThread invokes updateThread operation.
//
// PUT /v1/threads/{id}
func (c *Client) UpdateThread(ctx context.Context, request *UpdateThreadReq, params UpdateThreadParams) error {
	_, err := c.sendUpdateThread(ctx, request, params)
	return err
}

func (c *Client) sendUpdateThread(ctx context.Context, request *UpdateThreadReq, params UpdateThreadParams) (res *UpdateThreadNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateThread"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/threads/{id}"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateThreadOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [2]string
	pathParts[0] = "/v1/threads/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateThreadRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateThreadResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// UpdateUserInterests invokes updateUserInterests operation.
//
// PUT /v1/users/interests
func (c *Client) UpdateUserInterests(ctx context.Context, request *CommonIdsRequest) error {
	_, err := c.sendUpdateUserInterests(ctx, request)
	return err
}

func (c *Client) sendUpdateUserInterests(ctx context.Context, request *CommonIdsRequest) (res *UpdateUserInterestsNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("updateUserInterests"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/users/interests"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, UpdateUserInterestsOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/users/interests"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeUpdateUserInterestsRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeUpdateUserInterestsResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ValidateCallActionSignature invokes validateCallActionSignature operation.
//
// GET /v2/calls/action_signature/validate
func (c *Client) ValidateCallActionSignature(ctx context.Context, params ValidateCallActionSignatureParams) error {
	_, err := c.sendValidateCallActionSignature(ctx, params)
	return err
}

func (c *Client) sendValidateCallActionSignature(ctx context.Context, params ValidateCallActionSignatureParams) (res *ValidateCallActionSignatureNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("validateCallActionSignature"),
		semconv.HTTPRequestMethodKey.String("GET"),
		semconv.URLTemplateKey.String("/v2/calls/action_signature/validate"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ValidateCallActionSignatureOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v2/calls/action_signature/validate"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeQueryParams"
	q := uri.NewQueryEncoder()
	{
		// Encode "call_id" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "call_id",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.CallID.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "sender_uuid" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "sender_uuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.SenderUUID.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "receiver_uuid" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "receiver_uuid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.ReceiverUUID.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "action" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "action",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Action.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "timestamp" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "timestamp",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Timestamp.Get(); ok {
				return e.EncodeValue(conv.Int64ToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	{
		// Encode "signature" parameter.
		cfg := uri.QueryParameterEncodingConfig{
			Name:    "signature",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.EncodeParam(cfg, func(e uri.Encoder) error {
			if val, ok := params.Signature.Get(); ok {
				return e.EncodeValue(conv.StringToString(val))
			}
			return nil
		}); err != nil {
			return res, errors.Wrap(err, "encode query")
		}
	}
	u.RawQuery = q.Values().Encode()

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "GET", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeValidateCallActionSignatureResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ValidatePost invokes validatePost operation.
//
// POST /v1/posts/validate
func (c *Client) ValidatePost(ctx context.Context, request *ValidatePostReq) (*ValidationPostResponse, error) {
	res, err := c.sendValidatePost(ctx, request)
	return res, err
}

func (c *Client) sendValidatePost(ctx context.Context, request *ValidatePostReq) (res *ValidationPostResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("validatePost"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/posts/validate"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ValidatePostOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [1]string
	pathParts[0] = "/v1/posts/validate"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeValidatePostRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeValidatePostResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// ViewPostVideo invokes viewPostVideo operation.
//
// POST /v1/posts/videos/{videoId}/view
func (c *Client) ViewPostVideo(ctx context.Context, params ViewPostVideoParams) error {
	_, err := c.sendViewPostVideo(ctx, params)
	return err
}

func (c *Client) sendViewPostVideo(ctx context.Context, params ViewPostVideoParams) (res *ViewPostVideoNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("viewPostVideo"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/posts/videos/{videoId}/view"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, ViewPostVideoOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/posts/videos/"
	{
		// Encode "videoId" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "videoId",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.VideoId))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/view"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeViewPostVideoResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// VisitGroup invokes visitGroup operation.
//
// POST /v1/groups/{id}/visit
func (c *Client) VisitGroup(ctx context.Context, params VisitGroupParams) error {
	_, err := c.sendVisitGroup(ctx, params)
	return err
}

func (c *Client) sendVisitGroup(ctx context.Context, params VisitGroupParams) (res *VisitGroupNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("visitGroup"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/visit"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, VisitGroupOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/visit"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeVisitGroupResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// VoteSurvey invokes voteSurvey operation.
//
// POST /v2/surveys/{id}/vote
func (c *Client) VoteSurvey(ctx context.Context, request *VoteSurveyReq, params VoteSurveyParams) (*VoteSurveyResponse, error) {
	res, err := c.sendVoteSurvey(ctx, request, params)
	return res, err
}

func (c *Client) sendVoteSurvey(ctx context.Context, request *VoteSurveyReq, params VoteSurveyParams) (res *VoteSurveyResponse, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("voteSurvey"),
		semconv.HTTPRequestMethodKey.String("POST"),
		semconv.URLTemplateKey.String("/v2/surveys/{id}/vote"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, VoteSurveyOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v2/surveys/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/vote"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "POST", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeVoteSurveyRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeVoteSurveyResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// WithdrawGroupDeputy invokes withdrawGroupDeputy operation.
//
// PUT /v1/groups/{group_id}/deputize/{user_id}/withdraw
func (c *Client) WithdrawGroupDeputy(ctx context.Context, params WithdrawGroupDeputyParams) error {
	_, err := c.sendWithdrawGroupDeputy(ctx, params)
	return err
}

func (c *Client) sendWithdrawGroupDeputy(ctx context.Context, params WithdrawGroupDeputyParams) (res *WithdrawGroupDeputyNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("withdrawGroupDeputy"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/groups/{group_id}/deputize/{user_id}/withdraw"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, WithdrawGroupDeputyOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [5]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "group_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "group_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.GroupID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/deputize/"
	{
		// Encode "user_id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "user_id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.UserID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[3] = encoded
	}
	pathParts[4] = "/withdraw"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeWithdrawGroupDeputyResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}

// WithdrawGroupTransfer invokes withdrawGroupTransfer operation.
//
// PUT /v1/groups/{id}/transfer/withdraw
func (c *Client) WithdrawGroupTransfer(ctx context.Context, request *WithdrawGroupTransferReq, params WithdrawGroupTransferParams) error {
	_, err := c.sendWithdrawGroupTransfer(ctx, request, params)
	return err
}

func (c *Client) sendWithdrawGroupTransfer(ctx context.Context, request *WithdrawGroupTransferReq, params WithdrawGroupTransferParams) (res *WithdrawGroupTransferNoContent, err error) {
	otelAttrs := []attribute.KeyValue{
		otelogen.OperationID("withdrawGroupTransfer"),
		semconv.HTTPRequestMethodKey.String("PUT"),
		semconv.URLTemplateKey.String("/v1/groups/{id}/transfer/withdraw"),
	}
	otelAttrs = append(otelAttrs, c.cfg.Attributes...)

	// Run stopwatch.
	startTime := time.Now()
	defer func() {
		// Use floating point division here for higher precision (instead of Millisecond method).
		elapsedDuration := time.Since(startTime)
		c.duration.Record(ctx, float64(elapsedDuration)/float64(time.Millisecond), metric.WithAttributes(otelAttrs...))
	}()

	// Increment request counter.
	c.requests.Add(ctx, 1, metric.WithAttributes(otelAttrs...))

	// Start a span for this request.
	ctx, span := c.cfg.Tracer.Start(ctx, WithdrawGroupTransferOperation,
		trace.WithAttributes(otelAttrs...),
		clientSpanKind,
	)
	// Track stage for error reporting.
	var stage string
	defer func() {
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, stage)
			c.errors.Add(ctx, 1, metric.WithAttributes(otelAttrs...))
		}
		span.End()
	}()

	stage = "BuildURL"
	u := uri.Clone(c.requestURL(ctx))
	var pathParts [3]string
	pathParts[0] = "/v1/groups/"
	{
		// Encode "id" parameter.
		e := uri.NewPathEncoder(uri.PathEncoderConfig{
			Param:   "id",
			Style:   uri.PathStyleSimple,
			Explode: false,
		})
		if err := func() error {
			return e.EncodeValue(conv.Int64ToString(params.ID))
		}(); err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		encoded, err := e.Result()
		if err != nil {
			return res, errors.Wrap(err, "encode path")
		}
		pathParts[1] = encoded
	}
	pathParts[2] = "/transfer/withdraw"
	uri.AddPathParts(u, pathParts[:]...)

	stage = "EncodeRequest"
	r, err := ht.NewRequest(ctx, "PUT", u)
	if err != nil {
		return res, errors.Wrap(err, "create request")
	}
	if err := encodeWithdrawGroupTransferRequest(request, r); err != nil {
		return res, errors.Wrap(err, "encode request")
	}

	stage = "SendRequest"
	resp, err := c.cfg.Client.Do(r)
	if err != nil {
		return res, errors.Wrap(err, "do request")
	}
	body := resp.Body
	defer body.Close()

	stage = "DecodeResponse"
	result, err := decodeWithdrawGroupTransferResponse(resp)
	if err != nil {
		return res, errors.Wrap(err, "decode response")
	}

	return result, nil
}
