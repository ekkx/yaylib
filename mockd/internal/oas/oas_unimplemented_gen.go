// Code generated; DO NOT EDIT.

package oas

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// AcceptChatRequest implements acceptChatRequest operation.
//
// POST /v1/chat_rooms/accept_chat_request
func (UnimplementedHandler) AcceptChatRequest(ctx context.Context, req *AcceptChatRequestReq) error {
	return ht.ErrNotImplemented
}

// AcceptGroupJoinRequest implements acceptGroupJoinRequest operation.
//
// POST /v1/groups/{id}/accept/{userId}
func (UnimplementedHandler) AcceptGroupJoinRequest(ctx context.Context, params AcceptGroupJoinRequestParams) error {
	return ht.ErrNotImplemented
}

// AcceptGroupTransfer implements acceptGroupTransfer operation.
//
// PUT /v1/groups/{id}/transfer
func (UnimplementedHandler) AcceptGroupTransfer(ctx context.Context, params AcceptGroupTransferParams) error {
	return ht.ErrNotImplemented
}

// AddThreadMember implements addThreadMember operation.
//
// POST /v1/threads/{thread_id}/members/{id}
func (UnimplementedHandler) AddThreadMember(ctx context.Context, params AddThreadMemberParams) error {
	return ht.ErrNotImplemented
}

// AgreePolicy implements agreePolicy operation.
//
// POST /v1/users/policy_agreements/{type}
func (UnimplementedHandler) AgreePolicy(ctx context.Context, params AgreePolicyParams) error {
	return ht.ErrNotImplemented
}

// BanGroupUser implements banGroupUser operation.
//
// POST /v1/groups/{id}/ban/{userId}
func (UnimplementedHandler) BanGroupUser(ctx context.Context, params BanGroupUserParams) error {
	return ht.ErrNotImplemented
}

// BlockUser implements blockUser operation.
//
// POST /v1/users/{id}/block
func (UnimplementedHandler) BlockUser(ctx context.Context, params BlockUserParams) error {
	return ht.ErrNotImplemented
}

// BulkInviteToCall implements bulkInviteToCall operation.
//
// POST /v1/calls/{call_id}/bulk_invite
func (UnimplementedHandler) BulkInviteToCall(ctx context.Context, params BulkInviteToCallParams) error {
	return ht.ErrNotImplemented
}

// BumpCall implements bumpCall operation.
//
// POST /v1/calls/{call_id}/bump
func (UnimplementedHandler) BumpCall(ctx context.Context, params BumpCallParams) error {
	return ht.ErrNotImplemented
}

// CancelGroupTransfer implements cancelGroupTransfer operation.
//
// DELETE /v1/groups/{id}/transfer
func (UnimplementedHandler) CancelGroupTransfer(ctx context.Context, params CancelGroupTransferParams) error {
	return ht.ErrNotImplemented
}

// ChangeEmail implements changeEmail operation.
//
// PUT /v1/users/change_email
func (UnimplementedHandler) ChangeEmail(ctx context.Context, req *ChangeEmailReq) (r *LoginUpdateResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ChangePassword implements changePassword operation.
//
// PUT /v1/users/change_password
func (UnimplementedHandler) ChangePassword(ctx context.Context, req *ChangePasswordReq) (r *LoginUpdateResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateBookmark implements createBookmark operation.
//
// PUT /v1/users/{user_id}/bookmarks/{id}
func (UnimplementedHandler) CreateBookmark(ctx context.Context, params CreateBookmarkParams) (r *BookmarkPostResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChatRoom implements createChatRoom operation.
//
// POST /v3/chat_rooms/new
func (UnimplementedHandler) CreateChatRoom(ctx context.Context, req *CreateChatRoomReq) (r *CreateChatRoomResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateChatRoomV1 implements createChatRoomV1 operation.
//
// POST /v1/chat_rooms/new
func (UnimplementedHandler) CreateChatRoomV1(ctx context.Context, req *CreateChatRoomV1Req) (r *CreateChatRoomResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateConferenceCallPost implements createConferenceCallPost operation.
//
// POST /v2/posts/new_conference_call
func (UnimplementedHandler) CreateConferenceCallPost(ctx context.Context, req *CreateConferenceCallPostReq) (r *CreatePostResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateGroup implements createGroup operation.
//
// POST /v3/groups/new
func (UnimplementedHandler) CreateGroup(ctx context.Context, req *CreateGroupReq) (r *CreateGroupResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateMuteKeyword implements createMuteKeyword operation.
//
// POST /v1/hidden/words
func (UnimplementedHandler) CreateMuteKeyword(ctx context.Context, req *MuteKeywordRequest) (r *CreateMuteKeywordResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// CreatePost implements createPost operation.
//
// POST /v3/posts/new
func (UnimplementedHandler) CreatePost(ctx context.Context, req *CreatePostReq, params CreatePostParams) (r *Post, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateReview implements createReview operation.
//
// POST /v1/users/reviews
func (UnimplementedHandler) CreateReview(ctx context.Context, req *CreateReviewReq) error {
	return ht.ErrNotImplemented
}

// CreateSharePost implements createSharePost operation.
//
// POST /v2/posts/new_share_post
func (UnimplementedHandler) CreateSharePost(ctx context.Context, req *CreateSharePostReq) (r *Post, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateThread implements createThread operation.
//
// POST /v1/threads
func (UnimplementedHandler) CreateThread(ctx context.Context, req *CreateGroupThreadRequest) (r *ThreadInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateThreadPost implements createThreadPost operation.
//
// POST /v1/threads/{id}/posts
func (UnimplementedHandler) CreateThreadPost(ctx context.Context, req *CreateThreadPostReq, params CreateThreadPostParams) (r *Post, _ error) {
	return r, ht.ErrNotImplemented
}

// DeclineGroupJoinRequest implements declineGroupJoinRequest operation.
//
// POST /v1/groups/{id}/decline/{userId}
func (UnimplementedHandler) DeclineGroupJoinRequest(ctx context.Context, params DeclineGroupJoinRequestParams) error {
	return ht.ErrNotImplemented
}

// DeleteAllPosts implements deleteAllPosts operation.
//
// POST /v1/posts/delete_all_post
func (UnimplementedHandler) DeleteAllPosts(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// DeleteBookmark implements deleteBookmark operation.
//
// DELETE /v1/users/{user_id}/bookmarks/{id}
func (UnimplementedHandler) DeleteBookmark(ctx context.Context, params DeleteBookmarkParams) error {
	return ht.ErrNotImplemented
}

// DeleteChatMessage implements deleteChatMessage operation.
//
// DELETE /v1/chat_rooms/{room_id}/messages/{message_id}/delete
func (UnimplementedHandler) DeleteChatMessage(ctx context.Context, params DeleteChatMessageParams) error {
	return ht.ErrNotImplemented
}

// DeleteChatRooms implements deleteChatRooms operation.
//
// POST /v1/chat_rooms/mass_destroy
func (UnimplementedHandler) DeleteChatRooms(ctx context.Context, req *DeleteChatRoomsReq) error {
	return ht.ErrNotImplemented
}

// DeleteFootprint implements deleteFootprint operation.
//
// DELETE /v2/users/{user_id}/footprints/{footprint_id}
func (UnimplementedHandler) DeleteFootprint(ctx context.Context, params DeleteFootprintParams) error {
	return ht.ErrNotImplemented
}

// DeleteMuteKeyword implements deleteMuteKeyword operation.
//
// DELETE /v1/hidden/words
func (UnimplementedHandler) DeleteMuteKeyword(ctx context.Context, params DeleteMuteKeywordParams) error {
	return ht.ErrNotImplemented
}

// DeleteMyReviews implements deleteMyReviews operation.
//
// DELETE /v1/users/reviews
func (UnimplementedHandler) DeleteMyReviews(ctx context.Context, params DeleteMyReviewsParams) error {
	return ht.ErrNotImplemented
}

// DeletePosts implements deletePosts operation.
//
// POST /v2/posts/mass_destroy
func (UnimplementedHandler) DeletePosts(ctx context.Context, req *DeletePostsReq) error {
	return ht.ErrNotImplemented
}

// DeleteThread implements deleteThread operation.
//
// DELETE /v1/threads/{id}
func (UnimplementedHandler) DeleteThread(ctx context.Context, params DeleteThreadParams) error {
	return ht.ErrNotImplemented
}

// DeputizeGroupUsers implements deputizeGroupUsers operation.
//
// PUT /v1/groups/{id}/deputize
func (UnimplementedHandler) DeputizeGroupUsers(ctx context.Context, params DeputizeGroupUsersParams) error {
	return ht.ErrNotImplemented
}

// DeputizeGroupUsersMass implements deputizeGroupUsersMass operation.
//
// POST /v3/groups/{group_id}/deputize/mass
func (UnimplementedHandler) DeputizeGroupUsersMass(ctx context.Context, req *DeputizeGroupUsersMassReq, params DeputizeGroupUsersMassParams) error {
	return ht.ErrNotImplemented
}

// DisableTwoFactorAuth implements disableTwoFactorAuth operation.
//
// PUT /v1/users/secret/disable
func (UnimplementedHandler) DisableTwoFactorAuth(ctx context.Context, req *DisableTwoFactorAuthReq) error {
	return ht.ErrNotImplemented
}

// EditUser implements editUser operation.
//
// POST /v3/users/edit
func (UnimplementedHandler) EditUser(ctx context.Context, req *EditUserReq) error {
	return ht.ErrNotImplemented
}

// EditUserV2 implements editUserV2 operation.
//
// POST /v2/users/edit
func (UnimplementedHandler) EditUserV2(ctx context.Context, req *EditUserV2Req) error {
	return ht.ErrNotImplemented
}

// EnableTwoFactorAuth implements enableTwoFactorAuth operation.
//
// PUT /v1/users/secret/enable
func (UnimplementedHandler) EnableTwoFactorAuth(ctx context.Context, req *EnableTwoFactorAuthReq) (r *TwoStepAuthEnabledResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// FireGroupUser implements fireGroupUser operation.
//
// POST /v1/groups/{group_id}/fire/{user_id}
func (UnimplementedHandler) FireGroupUser(ctx context.Context, params FireGroupUserParams) error {
	return ht.ErrNotImplemented
}

// FollowUser implements followUser operation.
//
// POST /v2/users/{id}/follow
func (UnimplementedHandler) FollowUser(ctx context.Context, params FollowUserParams) error {
	return ht.ErrNotImplemented
}

// FollowUsers implements followUsers operation.
//
// POST /v2/users/follow
func (UnimplementedHandler) FollowUsers(ctx context.Context, params FollowUsersParams) error {
	return ht.ErrNotImplemented
}

// GenerateCallActionSignature implements generateCallActionSignature operation.
//
// POST /v1/calls/action_signature/generate
func (UnimplementedHandler) GenerateCallActionSignature(ctx context.Context, req *GenerateCallActionSignatureReq) (r *CallActionSignatureResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetActiveCallPost implements getActiveCallPost operation.
//
// GET /v1/posts/active_call
func (UnimplementedHandler) GetActiveCallPost(ctx context.Context, params GetActiveCallPostParams) (r *PostResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetActiveFollowings implements getActiveFollowings operation.
//
// GET /v1/users/active_followings
func (UnimplementedHandler) GetActiveFollowings(ctx context.Context, params GetActiveFollowingsParams) (r *ActiveFollowingsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAdditionalNotificationSetting implements getAdditionalNotificationSetting operation.
//
// GET /v1/users/additonal_notification_setting
func (UnimplementedHandler) GetAdditionalNotificationSetting(ctx context.Context) (r *AdditionalSettingsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAgoraRtmToken implements getAgoraRtmToken operation.
//
// GET /v3/calls/{call_id}/agora_rtm_token
func (UnimplementedHandler) GetAgoraRtmToken(ctx context.Context, params GetAgoraRtmTokenParams) (r *RtmTokenResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetAppConfig implements getAppConfig operation.
//
// GET /api/apps/{app}
func (UnimplementedHandler) GetAppConfig(ctx context.Context, params GetAppConfigParams) (r *ApplicationConfigResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBannedWords implements getBannedWords operation.
//
// GET /{countryApiValue}/api/v2/banned_words
func (UnimplementedHandler) GetBannedWords(ctx context.Context, params GetBannedWordsParams) (r *BanWordsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockedUserIds implements getBlockedUserIds operation.
//
// GET /v1/users/block_ids
func (UnimplementedHandler) GetBlockedUserIds(ctx context.Context) (r *BlockedUserIdsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBlockedUsers implements getBlockedUsers operation.
//
// POST /v2/users/blocked
func (UnimplementedHandler) GetBlockedUsers(ctx context.Context, req *SearchUsersRequest, params GetBlockedUsersParams) (r *BlockedUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBookmarkedPosts implements getBookmarkedPosts operation.
//
// GET /v1/users/{user_id}/bookmarks
func (UnimplementedHandler) GetBookmarkedPosts(ctx context.Context, params GetBookmarkedPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetBucketPresignedUrls implements getBucketPresignedUrls operation.
//
// GET /v1/buckets/presigned_urls
func (UnimplementedHandler) GetBucketPresignedUrls(ctx context.Context, params GetBucketPresignedUrlsParams) (r *PresignedUrlsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetCallBgms implements getCallBgms operation.
//
// GET /v1/calls/bgm
func (UnimplementedHandler) GetCallBgms(ctx context.Context) (r *BgmsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetCallFollowersTimeline implements getCallFollowersTimeline operation.
//
// GET /v2/posts/call_followers_timeline
func (UnimplementedHandler) GetCallFollowersTimeline(ctx context.Context, params GetCallFollowersTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetCallGiftHistory implements getCallGiftHistory operation.
//
// GET /v1/calls/{call_id}/gift_transactions
func (UnimplementedHandler) GetCallGiftHistory(ctx context.Context, params GetCallGiftHistoryParams) (r *CallGiftHistoryResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetCallTimeline implements getCallTimeline operation.
//
// GET /v2/posts/call_timeline
func (UnimplementedHandler) GetCallTimeline(ctx context.Context, params GetCallTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatMessages implements getChatMessages operation.
//
// GET /v2/chat_rooms/{id}/messages
func (UnimplementedHandler) GetChatMessages(ctx context.Context, params GetChatMessagesParams) (r *MessagesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatRequestCount implements getChatRequestCount operation.
//
// GET /v1/chat_rooms/total_chat_request
func (UnimplementedHandler) GetChatRequestCount(ctx context.Context) (r *TotalChatRequestResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatRequests implements getChatRequests operation.
//
// GET /v1/chat_rooms/request_list
func (UnimplementedHandler) GetChatRequests(ctx context.Context, params GetChatRequestsParams) (r *ChatRoomsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatRoom implements getChatRoom operation.
//
// GET /v2/chat_rooms/{id}
func (UnimplementedHandler) GetChatRoom(ctx context.Context, params GetChatRoomParams) (r *ChatRoomResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatUnreadStatus implements getChatUnreadStatus operation.
//
// GET /v1/chat_rooms/unread_status
func (UnimplementedHandler) GetChatUnreadStatus(ctx context.Context, params GetChatUnreadStatusParams) (r *UnreadStatusResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetChatableFollowings implements getChatableFollowings operation.
//
// POST /v1/users/followings/chatable
func (UnimplementedHandler) GetChatableFollowings(ctx context.Context, req *SearchUsersRequest, params GetChatableFollowingsParams) (r *FollowUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetConferenceCall implements getConferenceCall operation.
//
// GET /v2/calls/conferences/{call_id}
func (UnimplementedHandler) GetConferenceCall(ctx context.Context, params GetConferenceCallParams) (r *ConferenceCallResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetConversation implements getConversation operation.
//
// GET /v2/conversations/{id}
func (UnimplementedHandler) GetConversation(ctx context.Context, params GetConversationParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetDefaultSettings implements getDefaultSettings operation.
//
// GET /v1/users/default_settings
func (UnimplementedHandler) GetDefaultSettings(ctx context.Context) (r *DefaultSettingsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFollowRequestCount implements getFollowRequestCount operation.
//
// GET /v2/users/follow_requests_count
func (UnimplementedHandler) GetFollowRequestCount(ctx context.Context) (r *FollowRequestCountResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFollowRequests implements getFollowRequests operation.
//
// GET /v2/users/follow_requests
func (UnimplementedHandler) GetFollowRequests(ctx context.Context, params GetFollowRequestsParams) (r *UsersByTimestampResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFollowingTimeline implements getFollowingTimeline operation.
//
// GET /v2/posts/following_timeline
func (UnimplementedHandler) GetFollowingTimeline(ctx context.Context, params GetFollowingTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFollowingsBornToday implements getFollowingsBornToday operation.
//
// GET /v1/users/following_born_today
func (UnimplementedHandler) GetFollowingsBornToday(ctx context.Context, params GetFollowingsBornTodayParams) (r *UsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFootprints implements getFootprints operation.
//
// GET /v3/users/footprints
func (UnimplementedHandler) GetFootprints(ctx context.Context, params GetFootprintsParams) (r *FootprintsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetFreshUser implements getFreshUser operation.
//
// GET /v2/users/fresh/{id}
func (UnimplementedHandler) GetFreshUser(ctx context.Context, params GetFreshUserParams) (r *UserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroup implements getGroup operation.
//
// GET /v1/groups/{id}
func (UnimplementedHandler) GetGroup(ctx context.Context, params GetGroupParams) (r *GroupResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupBanList implements getGroupBanList operation.
//
// GET /v1/groups/{id}/ban_list
func (UnimplementedHandler) GetGroupBanList(ctx context.Context, params GetGroupBanListParams) (r *UsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupCreateQuota implements getGroupCreateQuota operation.
//
// GET /v1/groups/created_quota
func (UnimplementedHandler) GetGroupCreateQuota(ctx context.Context) (r *CreateQuotaResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupGiftHistory implements getGroupGiftHistory operation.
//
// GET /v1/groups/{group_id}/gift_history
func (UnimplementedHandler) GetGroupGiftHistory(ctx context.Context, params GetGroupGiftHistoryParams) (r *GroupGiftHistoryResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupGiftTransactions implements getGroupGiftTransactions operation.
//
// GET /v1/groups/{group_id}/gift_transactions
func (UnimplementedHandler) GetGroupGiftTransactions(ctx context.Context, params GetGroupGiftTransactionsParams) (r *GiftTransactionsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupHighlights implements getGroupHighlights operation.
//
// GET /v1/groups/{group_id}/highlights
func (UnimplementedHandler) GetGroupHighlights(ctx context.Context, params GetGroupHighlightsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupMember implements getGroupMember operation.
//
// GET /v1/groups/{id}/members/{userId}
func (UnimplementedHandler) GetGroupMember(ctx context.Context, params GetGroupMemberParams) (r *GroupUserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupMembers implements getGroupMembers operation.
//
// GET /v2/groups/{id}/members
func (UnimplementedHandler) GetGroupMembers(ctx context.Context, params GetGroupMembersParams) (r *GroupUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupNotificationSettings implements getGroupNotificationSettings operation.
//
// GET /v2/notification_settings/groups/{id}
func (UnimplementedHandler) GetGroupNotificationSettings(ctx context.Context, params GetGroupNotificationSettingsParams) (r *GroupNotificationSettingsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupReceivedGiftSenders implements getGroupReceivedGiftSenders operation.
//
// GET /v1/groups/{group_id}/received_gifts/{gift_id}/senders
func (UnimplementedHandler) GetGroupReceivedGiftSenders(ctx context.Context, params GetGroupReceivedGiftSendersParams) (r *GiftSendersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupTimeline implements getGroupTimeline operation.
//
// GET /v2/posts/group_timeline
func (UnimplementedHandler) GetGroupTimeline(ctx context.Context, params GetGroupTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetGroupUnreadStatus implements getGroupUnreadStatus operation.
//
// GET /v1/groups/unread_status
func (UnimplementedHandler) GetGroupUnreadStatus(ctx context.Context, params GetGroupUnreadStatusParams) (r *UnreadStatusResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetHimaUsers implements getHimaUsers operation.
//
// GET /v2/users/hima_users
func (UnimplementedHandler) GetHimaUsers(ctx context.Context, params GetHimaUsersParams) (r *HimaUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetInvitableCallUsers implements getInvitableCallUsers operation.
//
// GET /v1/calls/{call_id}/users/invitable
func (UnimplementedHandler) GetInvitableCallUsers(ctx context.Context, params GetInvitableCallUsersParams) (r *UsersByTimestampResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetInvitableGroupUsers implements getInvitableGroupUsers operation.
//
// GET /v1/groups/{group_id}/users/invitable
func (UnimplementedHandler) GetInvitableGroupUsers(ctx context.Context, params GetInvitableGroupUsersParams) (r *UsersByTimestampResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJoinedGroupStatuses implements getJoinedGroupStatuses operation.
//
// GET /v1/groups/joined_statuses
func (UnimplementedHandler) GetJoinedGroupStatuses(ctx context.Context, params GetJoinedGroupStatusesParams) (r GetJoinedGroupStatusesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetJoinedThreadStatuses implements getJoinedThreadStatuses operation.
//
// GET /v1/threads/joined_statuses
func (UnimplementedHandler) GetJoinedThreadStatuses(ctx context.Context, params GetJoinedThreadStatusesParams) (r GetJoinedThreadStatusesOK, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMainChatRooms implements getMainChatRooms operation.
//
// GET /v1/chat_rooms/main_list
func (UnimplementedHandler) GetMainChatRooms(ctx context.Context, params GetMainChatRoomsParams) (r *ChatRoomsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMyPosts implements getMyPosts operation.
//
// GET /v2/posts/mine
func (UnimplementedHandler) GetMyPosts(ctx context.Context, params GetMyPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetMyReviews implements getMyReviews operation.
//
// GET /v1/users/reviews/mine
func (UnimplementedHandler) GetMyReviews(ctx context.Context, params GetMyReviewsParams) (r *ReviewsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPhoneStatus implements getPhoneStatus operation.
//
// GET /v1/calls/phone_status/{opponent_id}
func (UnimplementedHandler) GetPhoneStatus(ctx context.Context, params GetPhoneStatusParams) (r *CallStatusResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPolicyAgreements implements getPolicyAgreements operation.
//
// GET /v1/users/policy_agreements
func (UnimplementedHandler) GetPolicyAgreements(ctx context.Context) (r *PolicyAgreementsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPopularWords implements getPopularWords operation.
//
// GET /{countryApiValue}/api/apps/{app}/popular_words
func (UnimplementedHandler) GetPopularWords(ctx context.Context, params GetPopularWordsParams) (r *PopularWordsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPost implements getPost operation.
//
// GET /v2/posts/{id}
func (UnimplementedHandler) GetPost(ctx context.Context, params GetPostParams) (r *PostResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPostGiftTransactions implements getPostGiftTransactions operation.
//
// GET /v1/posts/{post_id}/gift_transactions
func (UnimplementedHandler) GetPostGiftTransactions(ctx context.Context, params GetPostGiftTransactionsParams) (r *GiftTransactionsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPostLikers implements getPostLikers operation.
//
// GET /v1/posts/{id}/likers
func (UnimplementedHandler) GetPostLikers(ctx context.Context, params GetPostLikersParams) (r *PostLikersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPostReposts implements getPostReposts operation.
//
// GET /v2/posts/{id}/reposts
func (UnimplementedHandler) GetPostReposts(ctx context.Context, params GetPostRepostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPostUrlMetadata implements getPostUrlMetadata operation.
//
// GET /v2/posts/url_metadata
func (UnimplementedHandler) GetPostUrlMetadata(ctx context.Context, params GetPostUrlMetadataParams) (r *SharedUrl, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPostsByIds implements getPostsByIds operation.
//
// GET /v2/posts/multiple
func (UnimplementedHandler) GetPostsByIds(ctx context.Context, params GetPostsByIdsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetPostsByTag implements getPostsByTag operation.
//
// GET /v2/posts/tags/{tag}
func (UnimplementedHandler) GetPostsByTag(ctx context.Context, params GetPostsByTagParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetReceivedGiftSenders implements getReceivedGiftSenders operation.
//
// GET /v1/received_gifts/{gift_id}/senders
func (UnimplementedHandler) GetReceivedGiftSenders(ctx context.Context, params GetReceivedGiftSendersParams) (r *GiftSendersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetReceivedGiftTransaction implements getReceivedGiftTransaction operation.
//
// GET /v1/received_gifts/{gift_transaction_uuid}
func (UnimplementedHandler) GetReceivedGiftTransaction(ctx context.Context, params GetReceivedGiftTransactionParams) (r *GiftReceivedTransactionResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRecentEngagementPosts implements getRecentEngagementPosts operation.
//
// GET /v2/posts/recent_engagement
func (UnimplementedHandler) GetRecentEngagementPosts(ctx context.Context, params GetRecentEngagementPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRecommendedFollowUsers implements getRecommendedFollowUsers operation.
//
// GET /v1/users/{id}/follow_recommended
func (UnimplementedHandler) GetRecommendedFollowUsers(ctx context.Context, params GetRecommendedFollowUsersParams) (r *UsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRecommendedPostTags implements getRecommendedPostTags operation.
//
// POST /v1/posts/recommended_tag
func (UnimplementedHandler) GetRecommendedPostTags(ctx context.Context, req *GetRecommendedPostTagsReq) (r *PostTagsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRecommendedTimeline implements getRecommendedTimeline operation.
//
// GET /v2/posts/recommended_timeline
func (UnimplementedHandler) GetRecommendedTimeline(ctx context.Context, params GetRecommendedTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRelatableGroups implements getRelatableGroups operation.
//
// GET /v1/groups/{id}/relatable
func (UnimplementedHandler) GetRelatableGroups(ctx context.Context, params GetRelatableGroupsParams) (r *GroupsRelatedResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRelatedGroups implements getRelatedGroups operation.
//
// GET /v1/groups/{id}/related
func (UnimplementedHandler) GetRelatedGroups(ctx context.Context, params GetRelatedGroupsParams) (r *GroupsRelatedResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetResetCounters implements getResetCounters operation.
//
// GET /v1/users/reset_counters
func (UnimplementedHandler) GetResetCounters(ctx context.Context) (r *RefreshCounterRequestsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetRootPosts implements getRootPosts operation.
//
// GET /v2/conversations/root_posts
func (UnimplementedHandler) GetRootPosts(ctx context.Context, params GetRootPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetThread implements getThread operation.
//
// GET /v1/threads/{id}
func (UnimplementedHandler) GetThread(ctx context.Context, params GetThreadParams) (r *ThreadInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// GetThreadPosts implements getThreadPosts operation.
//
// GET /v1/threads/{id}/posts
func (UnimplementedHandler) GetThreadPosts(ctx context.Context, params GetThreadPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTimeline implements getTimeline operation.
//
// GET /v2/posts/{noreply_mode}timeline
func (UnimplementedHandler) GetTimeline(ctx context.Context, params GetTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTwoFactorAuthRequestInfo implements getTwoFactorAuthRequestInfo operation.
//
// GET /v1/users/secret
func (UnimplementedHandler) GetTwoFactorAuthRequestInfo(ctx context.Context) (r *TwoStepAuthRequestInfoResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetTwoFactorAuthStatus implements getTwoFactorAuthStatus operation.
//
// GET /v1/users/secret/status
func (UnimplementedHandler) GetTwoFactorAuthStatus(ctx context.Context) (r *TwoFAStatusResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUpdatedChatRooms implements getUpdatedChatRooms operation.
//
// GET /v2/chat_rooms/update
func (UnimplementedHandler) GetUpdatedChatRooms(ctx context.Context, params GetUpdatedChatRoomsParams) (r *ChatRoomsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUser implements getUser operation.
//
// GET /v2/users/{id}
func (UnimplementedHandler) GetUser(ctx context.Context, params GetUserParams) (r *UserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserActivities implements getUserActivities operation.
//
// GET /api/v2/user_activities
func (UnimplementedHandler) GetUserActivities(ctx context.Context, params GetUserActivitiesParams) (r *ActivitiesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserActivitiesV1 implements getUserActivitiesV1 operation.
//
// GET /api/user_activities
func (UnimplementedHandler) GetUserActivitiesV1(ctx context.Context, params GetUserActivitiesV1Params) (r *ActivitiesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserByQr implements getUserByQr operation.
//
// GET /v1/users/qr_codes/{qr}
func (UnimplementedHandler) GetUserByQr(ctx context.Context, params GetUserByQrParams) (r *UserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserCustomDefinitions implements getUserCustomDefinitions operation.
//
// GET /v1/users/custom_definitions
func (UnimplementedHandler) GetUserCustomDefinitions(ctx context.Context) (r *UserCustomDefinitionsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserFollowers implements getUserFollowers operation.
//
// GET /v3/users/{id}/followers
func (UnimplementedHandler) GetUserFollowers(ctx context.Context, params GetUserFollowersParams) (r *FollowUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserFollowings implements getUserFollowings operation.
//
// GET /v3/users/{id}/followings
func (UnimplementedHandler) GetUserFollowings(ctx context.Context, params GetUserFollowingsParams) (r *FollowUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserGiftTransactions implements getUserGiftTransactions operation.
//
// GET /v1/users/{user_id}/gift_transactions
func (UnimplementedHandler) GetUserGiftTransactions(ctx context.Context, params GetUserGiftTransactionsParams) (r *GiftTransactionsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserGroupList implements getUserGroupList operation.
//
// GET /v1/groups/user_group_list
func (UnimplementedHandler) GetUserGroupList(ctx context.Context, params GetUserGroupListParams) (r *GroupsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserInfo implements getUserInfo operation.
//
// GET /v2/users/info/{id}
func (UnimplementedHandler) GetUserInfo(ctx context.Context, params GetUserInfoParams) (r *UserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserInterests implements getUserInterests operation.
//
// GET /v1/users/interests
func (UnimplementedHandler) GetUserInterests(ctx context.Context) (r *UserInterestsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserPresignedUrl implements getUserPresignedUrl operation.
//
// GET /v1/users/presigned_url
func (UnimplementedHandler) GetUserPresignedUrl(ctx context.Context, params GetUserPresignedUrlParams) (r *PresignedUrlResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserReviews implements getUserReviews operation.
//
// GET /v2/users/reviews/{id}
func (UnimplementedHandler) GetUserReviews(ctx context.Context, params GetUserReviewsParams) (r *ReviewsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserTimeline implements getUserTimeline operation.
//
// GET /v2/posts/user_timeline
func (UnimplementedHandler) GetUserTimeline(ctx context.Context, params GetUserTimelineParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUserTimestamp implements getUserTimestamp operation.
//
// GET /v2/users/timestamp
func (UnimplementedHandler) GetUserTimestamp(ctx context.Context) (r *UserTimestampResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetUsersByIds implements getUsersByIds operation.
//
// GET /v1/users/list_id
func (UnimplementedHandler) GetUsersByIds(ctx context.Context, params GetUsersByIdsParams) (r *UsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// GetWebSocketToken implements getWebSocketToken operation.
//
// GET /v1/users/ws_token
func (UnimplementedHandler) GetWebSocketToken(ctx context.Context) (r *WebSocketTokenResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// HideChats implements hideChats operation.
//
// POST /v1/hidden/chats
func (UnimplementedHandler) HideChats(ctx context.Context, req *HideChatsReq) error {
	return ht.ErrNotImplemented
}

// HideUsers implements hideUsers operation.
//
// POST /v1/hidden/users
func (UnimplementedHandler) HideUsers(ctx context.Context, req *HideUsersReq) error {
	return ht.ErrNotImplemented
}

// InviteToCall implements inviteToCall operation.
//
// POST /v2/calls/invite
func (UnimplementedHandler) InviteToCall(ctx context.Context, req *InviteToCallReq) error {
	return ht.ErrNotImplemented
}

// InviteToChatRoom implements inviteToChatRoom operation.
//
// POST /v2/chat_rooms/{id}/invite
func (UnimplementedHandler) InviteToChatRoom(ctx context.Context, req *InviteToChatRoomReq, params InviteToChatRoomParams) error {
	return ht.ErrNotImplemented
}

// InviteToConferenceCall implements inviteToConferenceCall operation.
//
// POST /v1/calls/conference_calls/{call_id}/invite
func (UnimplementedHandler) InviteToConferenceCall(ctx context.Context, req *InviteToConferenceCallReq, params InviteToConferenceCallParams) error {
	return ht.ErrNotImplemented
}

// InviteToGroup implements inviteToGroup operation.
//
// POST /v1/groups/{id}/invite
func (UnimplementedHandler) InviteToGroup(ctx context.Context, req *InviteToGroupReq, params InviteToGroupParams) error {
	return ht.ErrNotImplemented
}

// JoinGroup implements joinGroup operation.
//
// POST /v1/groups/{id}/join
func (UnimplementedHandler) JoinGroup(ctx context.Context, params JoinGroupParams) error {
	return ht.ErrNotImplemented
}

// KickFromChatRoom implements kickFromChatRoom operation.
//
// POST /v2/chat_rooms/{id}/kick
func (UnimplementedHandler) KickFromChatRoom(ctx context.Context, req *KickFromChatRoomReq, params KickFromChatRoomParams) error {
	return ht.ErrNotImplemented
}

// KickFromConferenceCall implements kickFromConferenceCall operation.
//
// POST /v3/calls/conference_calls/{call_id}/kick
func (UnimplementedHandler) KickFromConferenceCall(ctx context.Context, req *KickFromConferenceCallReq, params KickFromConferenceCallParams) (r *CallActionSignatureResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// LeaveAgoraChannel implements leaveAgoraChannel operation.
//
// POST /v1/calls/leave_agora_channel
func (UnimplementedHandler) LeaveAgoraChannel(ctx context.Context, req *LeaveAgoraChannelReq) error {
	return ht.ErrNotImplemented
}

// LeaveConferenceCall implements leaveConferenceCall operation.
//
// POST /v1/calls/leave_conference_call
func (UnimplementedHandler) LeaveConferenceCall(ctx context.Context, req *LeaveConferenceCallReq) error {
	return ht.ErrNotImplemented
}

// LeaveGroup implements leaveGroup operation.
//
// DELETE /v1/groups/{id}/leave
func (UnimplementedHandler) LeaveGroup(ctx context.Context, params LeaveGroupParams) error {
	return ht.ErrNotImplemented
}

// LikePosts implements likePosts operation.
//
// POST /v2/posts/like
func (UnimplementedHandler) LikePosts(ctx context.Context, req *LikePostsReq) (r *LikePostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGameApps implements listGameApps operation.
//
// GET /v1/games/apps
func (UnimplementedHandler) ListGameApps(ctx context.Context, params ListGameAppsParams) (r *GamesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGameWalkthroughs implements listGameWalkthroughs operation.
//
// GET /v1/games/apps/{app_id}/walkthroughs
func (UnimplementedHandler) ListGameWalkthroughs(ctx context.Context, params ListGameWalkthroughsParams) (r []Walkthrough, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGenres implements listGenres operation.
//
// GET /v1/genres
func (UnimplementedHandler) ListGenres(ctx context.Context, params ListGenresParams) (r *GenresResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGifts implements listGifts operation.
//
// GET /v2/gifts
func (UnimplementedHandler) ListGifts(ctx context.Context, params ListGiftsParams) (r *GiftsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGroupCategories implements listGroupCategories operation.
//
// GET /v1/groups/categories
func (UnimplementedHandler) ListGroupCategories(ctx context.Context, params ListGroupCategoriesParams) (r *GroupCategoriesResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListGroups implements listGroups operation.
//
// GET /v2/groups
func (UnimplementedHandler) ListGroups(ctx context.Context, params ListGroupsParams) (r *GroupsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHiddenChats implements listHiddenChats operation.
//
// GET /v1/hidden/chats
func (UnimplementedHandler) ListHiddenChats(ctx context.Context, params ListHiddenChatsParams) (r *ChatRoomsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListHiddenUsers implements listHiddenUsers operation.
//
// GET /v1/hidden/users
func (UnimplementedHandler) ListHiddenUsers(ctx context.Context, params ListHiddenUsersParams) (r *HiddenResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMuteKeywords implements listMuteKeywords operation.
//
// GET /v1/hidden/words
func (UnimplementedHandler) ListMuteKeywords(ctx context.Context) (r *MuteKeywordResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMutedGroupUsers implements listMutedGroupUsers operation.
//
// GET /v1/group_mute/{id}/muted_users
func (UnimplementedHandler) ListMutedGroupUsers(ctx context.Context, params ListMutedGroupUsersParams) (r *GroupMuteUsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListMyGroups implements listMyGroups operation.
//
// GET /v2/groups/mine
func (UnimplementedHandler) ListMyGroups(ctx context.Context, params ListMyGroupsParams) (r *GroupsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListReceivedGifts implements listReceivedGifts operation.
//
// GET /v2/received_gifts
func (UnimplementedHandler) ListReceivedGifts(ctx context.Context) (r *GiftReceivedResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListReceivedGiftsV1 implements listReceivedGiftsV1 operation.
//
// GET /v1/received_gifts
func (UnimplementedHandler) ListReceivedGiftsV1(ctx context.Context, params ListReceivedGiftsV1Params) (r *GiftReceivedResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListStickerPacks implements listStickerPacks operation.
//
// GET /v2/sticker_packs
func (UnimplementedHandler) ListStickerPacks(ctx context.Context) (r *StickerPacksResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ListThreads implements listThreads operation.
//
// GET /v1/threads
func (UnimplementedHandler) ListThreads(ctx context.Context, params ListThreadsParams) (r *GroupThreadListResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// LoginWithEmail implements loginWithEmail operation.
//
// POST /v3/users/login_with_email
func (UnimplementedHandler) LoginWithEmail(ctx context.Context, req *LoginEmailUserRequest) (r *LoginUserResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// Logout implements logout operation.
//
// POST /v1/users/logout
func (UnimplementedHandler) Logout(ctx context.Context, req *LogoutReq) error {
	return ht.ErrNotImplemented
}

// MovePostToSpecificThread implements movePostToSpecificThread operation.
//
// PUT /v3/posts/{id}/move_to_thread/{thread_id}
func (UnimplementedHandler) MovePostToSpecificThread(ctx context.Context, params MovePostToSpecificThreadParams) (r *ThreadInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// MovePostToThread implements movePostToThread operation.
//
// POST /v3/posts/{id}/move_to_thread
func (UnimplementedHandler) MovePostToThread(ctx context.Context, req *MovePostToThreadReq, params MovePostToThreadParams) (r *ThreadInfo, _ error) {
	return r, ht.ErrNotImplemented
}

// MuteGroupUser implements muteGroupUser operation.
//
// POST /v1/group_mute/{id}/mute/{user_id}
func (UnimplementedHandler) MuteGroupUser(ctx context.Context, params MuteGroupUserParams) error {
	return ht.ErrNotImplemented
}

// OauthToken implements oauthToken operation.
//
// POST /api/v1/oauth/token
func (UnimplementedHandler) OauthToken(ctx context.Context, req *OauthTokenReq) (r *TokenResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// PinChatRoom implements pinChatRoom operation.
//
// POST /v1/chat_rooms/{id}/pinned
func (UnimplementedHandler) PinChatRoom(ctx context.Context, params PinChatRoomParams) error {
	return ht.ErrNotImplemented
}

// PinGroup implements pinGroup operation.
//
// POST /v1/pinned/groups
func (UnimplementedHandler) PinGroup(ctx context.Context, req *PinGroupReq) error {
	return ht.ErrNotImplemented
}

// PinGroupHighlightPost implements pinGroupHighlightPost operation.
//
// PUT /v1/groups/{group_id}/highlights/{post_id}
func (UnimplementedHandler) PinGroupHighlightPost(ctx context.Context, params PinGroupHighlightPostParams) error {
	return ht.ErrNotImplemented
}

// PinGroupPost implements pinGroupPost operation.
//
// PUT /v2/posts/group_pinned_post
func (UnimplementedHandler) PinGroupPost(ctx context.Context, req *PinGroupPostReq) error {
	return ht.ErrNotImplemented
}

// PinPost implements pinPost operation.
//
// POST /v1/pinned/posts
func (UnimplementedHandler) PinPost(ctx context.Context, req *PinPostReq) error {
	return ht.ErrNotImplemented
}

// PinReview implements pinReview operation.
//
// POST /v1/pinned/reviews
func (UnimplementedHandler) PinReview(ctx context.Context, req *PinReviewReq) error {
	return ht.ErrNotImplemented
}

// PingAlive implements pingAlive operation.
//
// POST /v1/users/alive
func (UnimplementedHandler) PingAlive(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// ReadChatAttachments implements readChatAttachments operation.
//
// POST /v1/chat_rooms/{id}/attachments_read
func (UnimplementedHandler) ReadChatAttachments(ctx context.Context, req *ReadAttachmentRequest, params ReadChatAttachmentsParams) error {
	return ht.ErrNotImplemented
}

// ReadChatMessage implements readChatMessage operation.
//
// POST /v2/chat_rooms/{id}/messages/{message_id}/read
func (UnimplementedHandler) ReadChatMessage(ctx context.Context, params ReadChatMessageParams) error {
	return ht.ErrNotImplemented
}

// ReadChatVideos implements readChatVideos operation.
//
// POST /v1/chat_rooms/{id}/videos_read
func (UnimplementedHandler) ReadChatVideos(ctx context.Context, req *ReadChatVideosReq, params ReadChatVideosParams) error {
	return ht.ErrNotImplemented
}

// RemoveChatRoomBackground implements removeChatRoomBackground operation.
//
// DELETE /v2/chat_rooms/{id}/background
func (UnimplementedHandler) RemoveChatRoomBackground(ctx context.Context, params RemoveChatRoomBackgroundParams) error {
	return ht.ErrNotImplemented
}

// RemoveCoverImage implements removeCoverImage operation.
//
// POST /v2/users/remove_cover_image
func (UnimplementedHandler) RemoveCoverImage(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// RemoveGroupCover implements removeGroupCover operation.
//
// DELETE /v3/groups/{id}/cover
func (UnimplementedHandler) RemoveGroupCover(ctx context.Context, params RemoveGroupCoverParams) error {
	return ht.ErrNotImplemented
}

// RemoveGroupDeputies implements removeGroupDeputies operation.
//
// DELETE /v1/groups/{id}/deputize
func (UnimplementedHandler) RemoveGroupDeputies(ctx context.Context, params RemoveGroupDeputiesParams) error {
	return ht.ErrNotImplemented
}

// RemoveGroupIcon implements removeGroupIcon operation.
//
// DELETE /v3/groups/{id}/icon
func (UnimplementedHandler) RemoveGroupIcon(ctx context.Context, params RemoveGroupIconParams) error {
	return ht.ErrNotImplemented
}

// RemoveProfilePhoto implements removeProfilePhoto operation.
//
// POST /v2/users/remove_profile_photo
func (UnimplementedHandler) RemoveProfilePhoto(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// RemoveRelatedGroups implements removeRelatedGroups operation.
//
// DELETE /v1/groups/{id}/related
func (UnimplementedHandler) RemoveRelatedGroups(ctx context.Context, params RemoveRelatedGroupsParams) error {
	return ht.ErrNotImplemented
}

// RemoveThreadMember implements removeThreadMember operation.
//
// DELETE /v1/threads/{thread_id}/members/{id}
func (UnimplementedHandler) RemoveThreadMember(ctx context.Context, params RemoveThreadMemberParams) error {
	return ht.ErrNotImplemented
}

// ReplyToReview implements replyToReview operation.
//
// POST /v2/users/reviews/{id}
func (UnimplementedHandler) ReplyToReview(ctx context.Context, req *ReplyToReviewReq, params ReplyToReviewParams) error {
	return ht.ErrNotImplemented
}

// ReportChatRoom implements reportChatRoom operation.
//
// POST /v3/chat_rooms/{chat_room_id}/report
func (UnimplementedHandler) ReportChatRoom(ctx context.Context, req *ReportChatRoomReq, params ReportChatRoomParams) error {
	return ht.ErrNotImplemented
}

// ReportGroup implements reportGroup operation.
//
// POST /v3/groups/{group_id}/report
func (UnimplementedHandler) ReportGroup(ctx context.Context, req *ReportGroupReq, params ReportGroupParams) error {
	return ht.ErrNotImplemented
}

// ReportPost implements reportPost operation.
//
// POST /v3/posts/{post_id}/report
func (UnimplementedHandler) ReportPost(ctx context.Context, req *ReportPostReq, params ReportPostParams) error {
	return ht.ErrNotImplemented
}

// ReportUser implements reportUser operation.
//
// POST /v3/users/{user_id}/report
func (UnimplementedHandler) ReportUser(ctx context.Context, req *ReportUserReq, params ReportUserParams) error {
	return ht.ErrNotImplemented
}

// Repost implements repost operation.
//
// POST /v3/posts/repost
func (UnimplementedHandler) Repost(ctx context.Context, req *RepostReq, params RepostParams) (r *CreatePostResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RequestEmailVerification implements requestEmailVerification operation.
//
// POST /v1/email_verification_urls
func (UnimplementedHandler) RequestEmailVerification(ctx context.Context, req *RequestEmailVerificationReq) (r *CommonUrlResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// RequestFollow implements requestFollow operation.
//
// POST /v2/users/{target_id}/follow_request
func (UnimplementedHandler) RequestFollow(ctx context.Context, req *RequestFollowReq, params RequestFollowParams) error {
	return ht.ErrNotImplemented
}

// RequestGroupWalkthrough implements requestGroupWalkthrough operation.
//
// POST /v1/groups/{id}/request_walkthrough
func (UnimplementedHandler) RequestGroupWalkthrough(ctx context.Context, params RequestGroupWalkthroughParams) error {
	return ht.ErrNotImplemented
}

// ResendConfirmEmail implements resendConfirmEmail operation.
//
// POST /v2/users/resend_confirm_email
func (UnimplementedHandler) ResendConfirmEmail(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// ResetCounters implements resetCounters operation.
//
// POST /v1/users/reset_counters
func (UnimplementedHandler) ResetCounters(ctx context.Context, req *ResetCountersReq) error {
	return ht.ErrNotImplemented
}

// ResetPassword implements resetPassword operation.
//
// PUT /v1/users/reset_password
func (UnimplementedHandler) ResetPassword(ctx context.Context, req *ResetPasswordReq) error {
	return ht.ErrNotImplemented
}

// SearchGroupPosts implements searchGroupPosts operation.
//
// GET /v2/groups/{id}/posts/search
func (UnimplementedHandler) SearchGroupPosts(ctx context.Context, params SearchGroupPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SearchPosts implements searchPosts operation.
//
// GET /v2/posts/search
func (UnimplementedHandler) SearchPosts(ctx context.Context, params SearchPostsParams) (r *PostsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SearchUsers implements searchUsers operation.
//
// GET /v1/users/search
func (UnimplementedHandler) SearchUsers(ctx context.Context, params SearchUsersParams) (r *UsersResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SendChatMessage implements sendChatMessage operation.
//
// POST /v3/chat_rooms/{id}/messages/new
func (UnimplementedHandler) SendChatMessage(ctx context.Context, req *SendChatMessageReq, params SendChatMessageParams) (r *MessageResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// SetGroupTitle implements setGroupTitle operation.
//
// POST /v1/groups/{id}/set_title
func (UnimplementedHandler) SetGroupTitle(ctx context.Context, req *SetGroupTitleReq, params SetGroupTitleParams) error {
	return ht.ErrNotImplemented
}

// SetHima implements setHima operation.
//
// POST /v1/users/hima
func (UnimplementedHandler) SetHima(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// StartConferenceCall implements startConferenceCall operation.
//
// POST /v2/calls/start_conference_call
func (UnimplementedHandler) StartConferenceCall(ctx context.Context, req *StartConferenceCallReq) (r *ConferenceCallResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// TakeOverGroup implements takeOverGroup operation.
//
// POST /v1/groups/{id}/take_over
func (UnimplementedHandler) TakeOverGroup(ctx context.Context, params TakeOverGroupParams) error {
	return ht.ErrNotImplemented
}

// TransferGroup implements transferGroup operation.
//
// POST /v3/groups/{id}/transfer
func (UnimplementedHandler) TransferGroup(ctx context.Context, req *TransferGroupReq, params TransferGroupParams) error {
	return ht.ErrNotImplemented
}

// UnbanGroupUser implements unbanGroupUser operation.
//
// POST /v1/groups/{id}/unban/{userId}
func (UnimplementedHandler) UnbanGroupUser(ctx context.Context, params UnbanGroupUserParams) error {
	return ht.ErrNotImplemented
}

// UnblockUser implements unblockUser operation.
//
// POST /v2/users/{id}/unblock
func (UnimplementedHandler) UnblockUser(ctx context.Context, params UnblockUserParams) error {
	return ht.ErrNotImplemented
}

// UnfollowUser implements unfollowUser operation.
//
// POST /v2/users/{id}/unfollow
func (UnimplementedHandler) UnfollowUser(ctx context.Context, params UnfollowUserParams) error {
	return ht.ErrNotImplemented
}

// UnhideChats implements unhideChats operation.
//
// DELETE /v1/hidden/chats
func (UnimplementedHandler) UnhideChats(ctx context.Context, params UnhideChatsParams) error {
	return ht.ErrNotImplemented
}

// UnhideUsers implements unhideUsers operation.
//
// DELETE /v1/hidden/users
func (UnimplementedHandler) UnhideUsers(ctx context.Context, params UnhideUsersParams) error {
	return ht.ErrNotImplemented
}

// UnlikePost implements unlikePost operation.
//
// POST /v1/posts/{id}/unlike
func (UnimplementedHandler) UnlikePost(ctx context.Context, params UnlikePostParams) error {
	return ht.ErrNotImplemented
}

// UnmuteGroupUser implements unmuteGroupUser operation.
//
// DELETE /v1/group_mute/{id}/unmute/{user_id}
func (UnimplementedHandler) UnmuteGroupUser(ctx context.Context, params UnmuteGroupUserParams) error {
	return ht.ErrNotImplemented
}

// UnpinChatRoom implements unpinChatRoom operation.
//
// DELETE /v1/chat_rooms/{id}/pinned
func (UnimplementedHandler) UnpinChatRoom(ctx context.Context, params UnpinChatRoomParams) error {
	return ht.ErrNotImplemented
}

// UnpinGroup implements unpinGroup operation.
//
// DELETE /v1/pinned/groups/{id}
func (UnimplementedHandler) UnpinGroup(ctx context.Context, params UnpinGroupParams) error {
	return ht.ErrNotImplemented
}

// UnpinGroupHighlightPost implements unpinGroupHighlightPost operation.
//
// DELETE /v1/groups/{group_id}/highlights/{post_id}
func (UnimplementedHandler) UnpinGroupHighlightPost(ctx context.Context, params UnpinGroupHighlightPostParams) error {
	return ht.ErrNotImplemented
}

// UnpinGroupPost implements unpinGroupPost operation.
//
// DELETE /v2/posts/group_pinned_post
func (UnimplementedHandler) UnpinGroupPost(ctx context.Context, params UnpinGroupPostParams) error {
	return ht.ErrNotImplemented
}

// UnpinReview implements unpinReview operation.
//
// DELETE /v1/pinned/reviews/{id}
func (UnimplementedHandler) UnpinReview(ctx context.Context, params UnpinReviewParams) error {
	return ht.ErrNotImplemented
}

// UpdateAdditionalNotificationSetting implements updateAdditionalNotificationSetting operation.
//
// POST /v1/users/additonal_notification_setting
func (UnimplementedHandler) UpdateAdditionalNotificationSetting(ctx context.Context, req *UpdateAdditionalNotificationSettingReq) error {
	return ht.ErrNotImplemented
}

// UpdateCall implements updateCall operation.
//
// PUT /v1/calls/{call_id}
func (UnimplementedHandler) UpdateCall(ctx context.Context, req *UpdateCallReq, params UpdateCallParams) error {
	return ht.ErrNotImplemented
}

// UpdateCallUser implements updateCallUser operation.
//
// PUT /v1/calls/{call_id}/users/{user_id}
func (UnimplementedHandler) UpdateCallUser(ctx context.Context, req *UpdateCallUserReq, params UpdateCallUserParams) error {
	return ht.ErrNotImplemented
}

// UpdateChatRoom implements updateChatRoom operation.
//
// POST /v3/chat_rooms/{id}/edit
func (UnimplementedHandler) UpdateChatRoom(ctx context.Context, req *UpdateChatRoomReq, params UpdateChatRoomParams) error {
	return ht.ErrNotImplemented
}

// UpdateChatRoomNotificationSettings implements updateChatRoomNotificationSettings operation.
//
// POST /v2/notification_settings/chat_rooms/{id}
func (UnimplementedHandler) UpdateChatRoomNotificationSettings(ctx context.Context, req *UpdateChatRoomNotificationSettingsReq, params UpdateChatRoomNotificationSettingsParams) (r *NotificationSettingResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateGroup implements updateGroup operation.
//
// POST /v3/groups/{id}/update
func (UnimplementedHandler) UpdateGroup(ctx context.Context, req *UpdateGroupReq, params UpdateGroupParams) (r *GroupResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateGroupNotificationSettings implements updateGroupNotificationSettings operation.
//
// POST /v2/notification_settings/groups/{id}
func (UnimplementedHandler) UpdateGroupNotificationSettings(ctx context.Context, req *UpdateGroupNotificationSettingsReq, params UpdateGroupNotificationSettingsParams) (r *AdditionalSettingsResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateLanguage implements updateLanguage operation.
//
// PUT /v1/users/language
func (UnimplementedHandler) UpdateLanguage(ctx context.Context, req *UpdateLanguageReq) error {
	return ht.ErrNotImplemented
}

// UpdateLogin implements updateLogin operation.
//
// POST /v3/users/login_update
func (UnimplementedHandler) UpdateLogin(ctx context.Context, req *UpdateLoginReq) (r *LoginUpdateResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdatePost implements updatePost operation.
//
// PUT /v3/posts/{id}
func (UnimplementedHandler) UpdatePost(ctx context.Context, req *UpdatePostReq, params UpdatePostParams) (r *Post, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateRelatedGroups implements updateRelatedGroups operation.
//
// PUT /v1/groups/{id}/related
func (UnimplementedHandler) UpdateRelatedGroups(ctx context.Context, params UpdateRelatedGroupsParams) error {
	return ht.ErrNotImplemented
}

// UpdateThread implements updateThread operation.
//
// PUT /v1/threads/{id}
func (UnimplementedHandler) UpdateThread(ctx context.Context, req *UpdateThreadReq, params UpdateThreadParams) error {
	return ht.ErrNotImplemented
}

// UpdateUserInterests implements updateUserInterests operation.
//
// PUT /v1/users/interests
func (UnimplementedHandler) UpdateUserInterests(ctx context.Context, req *CommonIdsRequest) error {
	return ht.ErrNotImplemented
}

// ValidateCallActionSignature implements validateCallActionSignature operation.
//
// GET /v2/calls/action_signature/validate
func (UnimplementedHandler) ValidateCallActionSignature(ctx context.Context, params ValidateCallActionSignatureParams) error {
	return ht.ErrNotImplemented
}

// ValidatePost implements validatePost operation.
//
// POST /v1/posts/validate
func (UnimplementedHandler) ValidatePost(ctx context.Context, req *ValidatePostReq) (r *ValidationPostResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// ViewPostVideo implements viewPostVideo operation.
//
// POST /v1/posts/videos/{videoId}/view
func (UnimplementedHandler) ViewPostVideo(ctx context.Context, params ViewPostVideoParams) error {
	return ht.ErrNotImplemented
}

// VisitGroup implements visitGroup operation.
//
// POST /v1/groups/{id}/visit
func (UnimplementedHandler) VisitGroup(ctx context.Context, params VisitGroupParams) error {
	return ht.ErrNotImplemented
}

// VoteSurvey implements voteSurvey operation.
//
// POST /v2/surveys/{id}/vote
func (UnimplementedHandler) VoteSurvey(ctx context.Context, req *VoteSurveyReq, params VoteSurveyParams) (r *VoteSurveyResponse, _ error) {
	return r, ht.ErrNotImplemented
}

// WithdrawGroupDeputy implements withdrawGroupDeputy operation.
//
// PUT /v1/groups/{group_id}/deputize/{user_id}/withdraw
func (UnimplementedHandler) WithdrawGroupDeputy(ctx context.Context, params WithdrawGroupDeputyParams) error {
	return ht.ErrNotImplemented
}

// WithdrawGroupTransfer implements withdrawGroupTransfer operation.
//
// PUT /v1/groups/{id}/transfer/withdraw
func (UnimplementedHandler) WithdrawGroupTransfer(ctx context.Context, req *WithdrawGroupTransferReq, params WithdrawGroupTransferParams) error {
	return ht.ErrNotImplemented
}
