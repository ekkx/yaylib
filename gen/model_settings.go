
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Settings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Settings{}

// Settings struct for Settings
type Settings struct {
	AgeRestrictedOnReview NullableBool `json:"age_restricted_on_review,omitempty"`
	AllowReposts NullableBool `json:"allow_reposts,omitempty"`
	CautionUserChat NullableBool `json:"caution_user_chat,omitempty"`
	FollowingOnlyCallInvite NullableBool `json:"following_only_call_invite,omitempty"`
	FollowingOnlyGroupInvite NullableBool `json:"following_only_group_invite,omitempty"`
	FollowingRestrictedOnReview NullableBool `json:"following_restricted_on_review,omitempty"`
	HideActiveCall NullableBool `json:"hide_active_call,omitempty"`
	HideGiftsReceived NullableBool `json:"hide_gifts_received,omitempty"`
	HideHotPost NullableBool `json:"hide_hot_post,omitempty"`
	HideOnInvitable NullableBool `json:"hide_on_invitable,omitempty"`
	HideOnlineStatus NullableBool `json:"hide_online_status,omitempty"`
	HideReplyFollowingTimeline NullableBool `json:"hide_reply_following_timeline,omitempty"`
	HideReplyPublicTimeline NullableBool `json:"hide_reply_public_timeline,omitempty"`
	HideVip NullableBool `json:"hide_vip,omitempty"`
	InvisibleOnUserSearch NullableBool `json:"invisible_on_user_search,omitempty"`
	NoReplyGroupTimeline NullableBool `json:"no_reply_group_timeline,omitempty"`
	NotificationBirthdayToFollowers NullableBool `json:"notification_birthday_to_followers,omitempty"`
	NotificationBulkCallInvite NullableBool `json:"notification_bulk_call_invite,omitempty"`
	NotificationCallInvite NullableBool `json:"notification_call_invite,omitempty"`
	NotificationChat NullableBool `json:"notification_chat,omitempty"`
	NotificationChatDelete NullableBool `json:"notification_chat_delete,omitempty"`
	NotificationContactFriend NullableBool `json:"notification_contact_friend,omitempty"`
	NotificationDailyQuest NullableBool `json:"notification_daily_quest,omitempty"`
	NotificationDailySummary NullableBool `json:"notification_daily_summary,omitempty"`
	NotificationFollow NullableBool `json:"notification_follow,omitempty"`
	NotificationFollowAccept NullableBool `json:"notification_follow_accept,omitempty"`
	NotificationFollowRequest NullableBool `json:"notification_follow_request,omitempty"`
	NotificationFollowerConferenceCall NullableBool `json:"notification_follower_conference_call,omitempty"`
	NotificationFollowerCreateGroup NullableBool `json:"notification_follower_create_group,omitempty"`
	NotificationFollowingBirthdateOn NullableBool `json:"notification_following_birthdate_on,omitempty"`
	NotificationFollowingPostAfterBreak NullableBool `json:"notification_following_post_after_break,omitempty"`
	NotificationFollowingsInCall NullableBool `json:"notification_followings_in_call,omitempty"`
	NotificationFootprint NullableBool `json:"notification_footprint,omitempty"`
	NotificationGiftReceive NullableBool `json:"notification_gift_receive,omitempty"`
	NotificationGroupAccept NullableBool `json:"notification_group_accept,omitempty"`
	NotificationGroupConferenceCall NullableBool `json:"notification_group_conference_call,omitempty"`
	NotificationGroupInvite NullableBool `json:"notification_group_invite,omitempty"`
	NotificationGroupJoin NullableBool `json:"notification_group_join,omitempty"`
	NotificationGroupMessageTagAll NullableBool `json:"notification_group_message_tag_all,omitempty"`
	NotificationGroupModerator NullableBool `json:"notification_group_moderator,omitempty"`
	NotificationGroupPost NullableBool `json:"notification_group_post,omitempty"`
	NotificationGroupRequest NullableBool `json:"notification_group_request,omitempty"`
	NotificationHimaNow NullableBool `json:"notification_hima_now,omitempty"`
	NotificationInviteeCreateAccount NullableBool `json:"notification_invitee_create_account,omitempty"`
	NotificationLatestNews NullableBool `json:"notification_latest_news,omitempty"`
	NotificationLike NullableBool `json:"notification_like,omitempty"`
	NotificationMessageTag NullableBool `json:"notification_message_tag,omitempty"`
	NotificationPopularPost NullableBool `json:"notification_popular_post,omitempty"`
	NotificationProfileScreenshot NullableBool `json:"notification_profile_screenshot,omitempty"`
	NotificationReply NullableBool `json:"notification_reply,omitempty"`
	NotificationRepost NullableBool `json:"notification_repost,omitempty"`
	NotificationReview NullableBool `json:"notification_review,omitempty"`
	NotificationSecurityWarning NullableBool `json:"notification_security_warning,omitempty"`
	NotificationTwitterFriend NullableBool `json:"notification_twitter_friend,omitempty"`
	PrivacyMode NullableBool `json:"privacy_mode,omitempty"`
	PrivatePost NullableBool `json:"private_post,omitempty"`
	PrivateUserTimeline NullableBool `json:"private_user_timeline,omitempty"`
	ShowTotalGiftsReceivedCount NullableBool `json:"show_total_gifts_received_count,omitempty"`
	VipInvisibleFootprintMode NullableBool `json:"vip_invisible_footprint_mode,omitempty"`
	VisibleOnSnsFriendRecommendation NullableBool `json:"visible_on_sns_friend_recommendation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Settings Settings

// NewSettings instantiates a new Settings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettings() *Settings {
	this := Settings{}
	return &this
}

// NewSettingsWithDefaults instantiates a new Settings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsWithDefaults() *Settings {
	this := Settings{}
	return &this
}

// GetAgeRestrictedOnReview returns the AgeRestrictedOnReview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetAgeRestrictedOnReview() bool {
	if o == nil || IsNil(o.AgeRestrictedOnReview.Get()) {
		var ret bool
		return ret
	}
	return *o.AgeRestrictedOnReview.Get()
}

// GetAgeRestrictedOnReviewOk returns a tuple with the AgeRestrictedOnReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetAgeRestrictedOnReviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgeRestrictedOnReview.Get(), o.AgeRestrictedOnReview.IsSet()
}

// HasAgeRestrictedOnReview returns a boolean if a field has been set.
func (o *Settings) HasAgeRestrictedOnReview() bool {
	if o != nil && o.AgeRestrictedOnReview.IsSet() {
		return true
	}

	return false
}

// SetAgeRestrictedOnReview gets a reference to the given NullableBool and assigns it to the AgeRestrictedOnReview field.
func (o *Settings) SetAgeRestrictedOnReview(v bool) {
	o.AgeRestrictedOnReview.Set(&v)
}
// SetAgeRestrictedOnReviewNil sets the value for AgeRestrictedOnReview to be an explicit nil
func (o *Settings) SetAgeRestrictedOnReviewNil() {
	o.AgeRestrictedOnReview.Set(nil)
}

// UnsetAgeRestrictedOnReview ensures that no value is present for AgeRestrictedOnReview, not even an explicit nil
func (o *Settings) UnsetAgeRestrictedOnReview() {
	o.AgeRestrictedOnReview.Unset()
}

// GetAllowReposts returns the AllowReposts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetAllowReposts() bool {
	if o == nil || IsNil(o.AllowReposts.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowReposts.Get()
}

// GetAllowRepostsOk returns a tuple with the AllowReposts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetAllowRepostsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowReposts.Get(), o.AllowReposts.IsSet()
}

// HasAllowReposts returns a boolean if a field has been set.
func (o *Settings) HasAllowReposts() bool {
	if o != nil && o.AllowReposts.IsSet() {
		return true
	}

	return false
}

// SetAllowReposts gets a reference to the given NullableBool and assigns it to the AllowReposts field.
func (o *Settings) SetAllowReposts(v bool) {
	o.AllowReposts.Set(&v)
}
// SetAllowRepostsNil sets the value for AllowReposts to be an explicit nil
func (o *Settings) SetAllowRepostsNil() {
	o.AllowReposts.Set(nil)
}

// UnsetAllowReposts ensures that no value is present for AllowReposts, not even an explicit nil
func (o *Settings) UnsetAllowReposts() {
	o.AllowReposts.Unset()
}

// GetCautionUserChat returns the CautionUserChat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetCautionUserChat() bool {
	if o == nil || IsNil(o.CautionUserChat.Get()) {
		var ret bool
		return ret
	}
	return *o.CautionUserChat.Get()
}

// GetCautionUserChatOk returns a tuple with the CautionUserChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetCautionUserChatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CautionUserChat.Get(), o.CautionUserChat.IsSet()
}

// HasCautionUserChat returns a boolean if a field has been set.
func (o *Settings) HasCautionUserChat() bool {
	if o != nil && o.CautionUserChat.IsSet() {
		return true
	}

	return false
}

// SetCautionUserChat gets a reference to the given NullableBool and assigns it to the CautionUserChat field.
func (o *Settings) SetCautionUserChat(v bool) {
	o.CautionUserChat.Set(&v)
}
// SetCautionUserChatNil sets the value for CautionUserChat to be an explicit nil
func (o *Settings) SetCautionUserChatNil() {
	o.CautionUserChat.Set(nil)
}

// UnsetCautionUserChat ensures that no value is present for CautionUserChat, not even an explicit nil
func (o *Settings) UnsetCautionUserChat() {
	o.CautionUserChat.Unset()
}

// GetFollowingOnlyCallInvite returns the FollowingOnlyCallInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetFollowingOnlyCallInvite() bool {
	if o == nil || IsNil(o.FollowingOnlyCallInvite.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowingOnlyCallInvite.Get()
}

// GetFollowingOnlyCallInviteOk returns a tuple with the FollowingOnlyCallInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetFollowingOnlyCallInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingOnlyCallInvite.Get(), o.FollowingOnlyCallInvite.IsSet()
}

// HasFollowingOnlyCallInvite returns a boolean if a field has been set.
func (o *Settings) HasFollowingOnlyCallInvite() bool {
	if o != nil && o.FollowingOnlyCallInvite.IsSet() {
		return true
	}

	return false
}

// SetFollowingOnlyCallInvite gets a reference to the given NullableBool and assigns it to the FollowingOnlyCallInvite field.
func (o *Settings) SetFollowingOnlyCallInvite(v bool) {
	o.FollowingOnlyCallInvite.Set(&v)
}
// SetFollowingOnlyCallInviteNil sets the value for FollowingOnlyCallInvite to be an explicit nil
func (o *Settings) SetFollowingOnlyCallInviteNil() {
	o.FollowingOnlyCallInvite.Set(nil)
}

// UnsetFollowingOnlyCallInvite ensures that no value is present for FollowingOnlyCallInvite, not even an explicit nil
func (o *Settings) UnsetFollowingOnlyCallInvite() {
	o.FollowingOnlyCallInvite.Unset()
}

// GetFollowingOnlyGroupInvite returns the FollowingOnlyGroupInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetFollowingOnlyGroupInvite() bool {
	if o == nil || IsNil(o.FollowingOnlyGroupInvite.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowingOnlyGroupInvite.Get()
}

// GetFollowingOnlyGroupInviteOk returns a tuple with the FollowingOnlyGroupInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetFollowingOnlyGroupInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingOnlyGroupInvite.Get(), o.FollowingOnlyGroupInvite.IsSet()
}

// HasFollowingOnlyGroupInvite returns a boolean if a field has been set.
func (o *Settings) HasFollowingOnlyGroupInvite() bool {
	if o != nil && o.FollowingOnlyGroupInvite.IsSet() {
		return true
	}

	return false
}

// SetFollowingOnlyGroupInvite gets a reference to the given NullableBool and assigns it to the FollowingOnlyGroupInvite field.
func (o *Settings) SetFollowingOnlyGroupInvite(v bool) {
	o.FollowingOnlyGroupInvite.Set(&v)
}
// SetFollowingOnlyGroupInviteNil sets the value for FollowingOnlyGroupInvite to be an explicit nil
func (o *Settings) SetFollowingOnlyGroupInviteNil() {
	o.FollowingOnlyGroupInvite.Set(nil)
}

// UnsetFollowingOnlyGroupInvite ensures that no value is present for FollowingOnlyGroupInvite, not even an explicit nil
func (o *Settings) UnsetFollowingOnlyGroupInvite() {
	o.FollowingOnlyGroupInvite.Unset()
}

// GetFollowingRestrictedOnReview returns the FollowingRestrictedOnReview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetFollowingRestrictedOnReview() bool {
	if o == nil || IsNil(o.FollowingRestrictedOnReview.Get()) {
		var ret bool
		return ret
	}
	return *o.FollowingRestrictedOnReview.Get()
}

// GetFollowingRestrictedOnReviewOk returns a tuple with the FollowingRestrictedOnReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetFollowingRestrictedOnReviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowingRestrictedOnReview.Get(), o.FollowingRestrictedOnReview.IsSet()
}

// HasFollowingRestrictedOnReview returns a boolean if a field has been set.
func (o *Settings) HasFollowingRestrictedOnReview() bool {
	if o != nil && o.FollowingRestrictedOnReview.IsSet() {
		return true
	}

	return false
}

// SetFollowingRestrictedOnReview gets a reference to the given NullableBool and assigns it to the FollowingRestrictedOnReview field.
func (o *Settings) SetFollowingRestrictedOnReview(v bool) {
	o.FollowingRestrictedOnReview.Set(&v)
}
// SetFollowingRestrictedOnReviewNil sets the value for FollowingRestrictedOnReview to be an explicit nil
func (o *Settings) SetFollowingRestrictedOnReviewNil() {
	o.FollowingRestrictedOnReview.Set(nil)
}

// UnsetFollowingRestrictedOnReview ensures that no value is present for FollowingRestrictedOnReview, not even an explicit nil
func (o *Settings) UnsetFollowingRestrictedOnReview() {
	o.FollowingRestrictedOnReview.Unset()
}

// GetHideActiveCall returns the HideActiveCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideActiveCall() bool {
	if o == nil || IsNil(o.HideActiveCall.Get()) {
		var ret bool
		return ret
	}
	return *o.HideActiveCall.Get()
}

// GetHideActiveCallOk returns a tuple with the HideActiveCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideActiveCallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideActiveCall.Get(), o.HideActiveCall.IsSet()
}

// HasHideActiveCall returns a boolean if a field has been set.
func (o *Settings) HasHideActiveCall() bool {
	if o != nil && o.HideActiveCall.IsSet() {
		return true
	}

	return false
}

// SetHideActiveCall gets a reference to the given NullableBool and assigns it to the HideActiveCall field.
func (o *Settings) SetHideActiveCall(v bool) {
	o.HideActiveCall.Set(&v)
}
// SetHideActiveCallNil sets the value for HideActiveCall to be an explicit nil
func (o *Settings) SetHideActiveCallNil() {
	o.HideActiveCall.Set(nil)
}

// UnsetHideActiveCall ensures that no value is present for HideActiveCall, not even an explicit nil
func (o *Settings) UnsetHideActiveCall() {
	o.HideActiveCall.Unset()
}

// GetHideGiftsReceived returns the HideGiftsReceived field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideGiftsReceived() bool {
	if o == nil || IsNil(o.HideGiftsReceived.Get()) {
		var ret bool
		return ret
	}
	return *o.HideGiftsReceived.Get()
}

// GetHideGiftsReceivedOk returns a tuple with the HideGiftsReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideGiftsReceivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideGiftsReceived.Get(), o.HideGiftsReceived.IsSet()
}

// HasHideGiftsReceived returns a boolean if a field has been set.
func (o *Settings) HasHideGiftsReceived() bool {
	if o != nil && o.HideGiftsReceived.IsSet() {
		return true
	}

	return false
}

// SetHideGiftsReceived gets a reference to the given NullableBool and assigns it to the HideGiftsReceived field.
func (o *Settings) SetHideGiftsReceived(v bool) {
	o.HideGiftsReceived.Set(&v)
}
// SetHideGiftsReceivedNil sets the value for HideGiftsReceived to be an explicit nil
func (o *Settings) SetHideGiftsReceivedNil() {
	o.HideGiftsReceived.Set(nil)
}

// UnsetHideGiftsReceived ensures that no value is present for HideGiftsReceived, not even an explicit nil
func (o *Settings) UnsetHideGiftsReceived() {
	o.HideGiftsReceived.Unset()
}

// GetHideHotPost returns the HideHotPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideHotPost() bool {
	if o == nil || IsNil(o.HideHotPost.Get()) {
		var ret bool
		return ret
	}
	return *o.HideHotPost.Get()
}

// GetHideHotPostOk returns a tuple with the HideHotPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideHotPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideHotPost.Get(), o.HideHotPost.IsSet()
}

// HasHideHotPost returns a boolean if a field has been set.
func (o *Settings) HasHideHotPost() bool {
	if o != nil && o.HideHotPost.IsSet() {
		return true
	}

	return false
}

// SetHideHotPost gets a reference to the given NullableBool and assigns it to the HideHotPost field.
func (o *Settings) SetHideHotPost(v bool) {
	o.HideHotPost.Set(&v)
}
// SetHideHotPostNil sets the value for HideHotPost to be an explicit nil
func (o *Settings) SetHideHotPostNil() {
	o.HideHotPost.Set(nil)
}

// UnsetHideHotPost ensures that no value is present for HideHotPost, not even an explicit nil
func (o *Settings) UnsetHideHotPost() {
	o.HideHotPost.Unset()
}

// GetHideOnInvitable returns the HideOnInvitable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideOnInvitable() bool {
	if o == nil || IsNil(o.HideOnInvitable.Get()) {
		var ret bool
		return ret
	}
	return *o.HideOnInvitable.Get()
}

// GetHideOnInvitableOk returns a tuple with the HideOnInvitable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideOnInvitableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideOnInvitable.Get(), o.HideOnInvitable.IsSet()
}

// HasHideOnInvitable returns a boolean if a field has been set.
func (o *Settings) HasHideOnInvitable() bool {
	if o != nil && o.HideOnInvitable.IsSet() {
		return true
	}

	return false
}

// SetHideOnInvitable gets a reference to the given NullableBool and assigns it to the HideOnInvitable field.
func (o *Settings) SetHideOnInvitable(v bool) {
	o.HideOnInvitable.Set(&v)
}
// SetHideOnInvitableNil sets the value for HideOnInvitable to be an explicit nil
func (o *Settings) SetHideOnInvitableNil() {
	o.HideOnInvitable.Set(nil)
}

// UnsetHideOnInvitable ensures that no value is present for HideOnInvitable, not even an explicit nil
func (o *Settings) UnsetHideOnInvitable() {
	o.HideOnInvitable.Unset()
}

// GetHideOnlineStatus returns the HideOnlineStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideOnlineStatus() bool {
	if o == nil || IsNil(o.HideOnlineStatus.Get()) {
		var ret bool
		return ret
	}
	return *o.HideOnlineStatus.Get()
}

// GetHideOnlineStatusOk returns a tuple with the HideOnlineStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideOnlineStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideOnlineStatus.Get(), o.HideOnlineStatus.IsSet()
}

// HasHideOnlineStatus returns a boolean if a field has been set.
func (o *Settings) HasHideOnlineStatus() bool {
	if o != nil && o.HideOnlineStatus.IsSet() {
		return true
	}

	return false
}

// SetHideOnlineStatus gets a reference to the given NullableBool and assigns it to the HideOnlineStatus field.
func (o *Settings) SetHideOnlineStatus(v bool) {
	o.HideOnlineStatus.Set(&v)
}
// SetHideOnlineStatusNil sets the value for HideOnlineStatus to be an explicit nil
func (o *Settings) SetHideOnlineStatusNil() {
	o.HideOnlineStatus.Set(nil)
}

// UnsetHideOnlineStatus ensures that no value is present for HideOnlineStatus, not even an explicit nil
func (o *Settings) UnsetHideOnlineStatus() {
	o.HideOnlineStatus.Unset()
}

// GetHideReplyFollowingTimeline returns the HideReplyFollowingTimeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideReplyFollowingTimeline() bool {
	if o == nil || IsNil(o.HideReplyFollowingTimeline.Get()) {
		var ret bool
		return ret
	}
	return *o.HideReplyFollowingTimeline.Get()
}

// GetHideReplyFollowingTimelineOk returns a tuple with the HideReplyFollowingTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideReplyFollowingTimelineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideReplyFollowingTimeline.Get(), o.HideReplyFollowingTimeline.IsSet()
}

// HasHideReplyFollowingTimeline returns a boolean if a field has been set.
func (o *Settings) HasHideReplyFollowingTimeline() bool {
	if o != nil && o.HideReplyFollowingTimeline.IsSet() {
		return true
	}

	return false
}

// SetHideReplyFollowingTimeline gets a reference to the given NullableBool and assigns it to the HideReplyFollowingTimeline field.
func (o *Settings) SetHideReplyFollowingTimeline(v bool) {
	o.HideReplyFollowingTimeline.Set(&v)
}
// SetHideReplyFollowingTimelineNil sets the value for HideReplyFollowingTimeline to be an explicit nil
func (o *Settings) SetHideReplyFollowingTimelineNil() {
	o.HideReplyFollowingTimeline.Set(nil)
}

// UnsetHideReplyFollowingTimeline ensures that no value is present for HideReplyFollowingTimeline, not even an explicit nil
func (o *Settings) UnsetHideReplyFollowingTimeline() {
	o.HideReplyFollowingTimeline.Unset()
}

// GetHideReplyPublicTimeline returns the HideReplyPublicTimeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideReplyPublicTimeline() bool {
	if o == nil || IsNil(o.HideReplyPublicTimeline.Get()) {
		var ret bool
		return ret
	}
	return *o.HideReplyPublicTimeline.Get()
}

// GetHideReplyPublicTimelineOk returns a tuple with the HideReplyPublicTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideReplyPublicTimelineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideReplyPublicTimeline.Get(), o.HideReplyPublicTimeline.IsSet()
}

// HasHideReplyPublicTimeline returns a boolean if a field has been set.
func (o *Settings) HasHideReplyPublicTimeline() bool {
	if o != nil && o.HideReplyPublicTimeline.IsSet() {
		return true
	}

	return false
}

// SetHideReplyPublicTimeline gets a reference to the given NullableBool and assigns it to the HideReplyPublicTimeline field.
func (o *Settings) SetHideReplyPublicTimeline(v bool) {
	o.HideReplyPublicTimeline.Set(&v)
}
// SetHideReplyPublicTimelineNil sets the value for HideReplyPublicTimeline to be an explicit nil
func (o *Settings) SetHideReplyPublicTimelineNil() {
	o.HideReplyPublicTimeline.Set(nil)
}

// UnsetHideReplyPublicTimeline ensures that no value is present for HideReplyPublicTimeline, not even an explicit nil
func (o *Settings) UnsetHideReplyPublicTimeline() {
	o.HideReplyPublicTimeline.Unset()
}

// GetHideVip returns the HideVip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetHideVip() bool {
	if o == nil || IsNil(o.HideVip.Get()) {
		var ret bool
		return ret
	}
	return *o.HideVip.Get()
}

// GetHideVipOk returns a tuple with the HideVip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetHideVipOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideVip.Get(), o.HideVip.IsSet()
}

// HasHideVip returns a boolean if a field has been set.
func (o *Settings) HasHideVip() bool {
	if o != nil && o.HideVip.IsSet() {
		return true
	}

	return false
}

// SetHideVip gets a reference to the given NullableBool and assigns it to the HideVip field.
func (o *Settings) SetHideVip(v bool) {
	o.HideVip.Set(&v)
}
// SetHideVipNil sets the value for HideVip to be an explicit nil
func (o *Settings) SetHideVipNil() {
	o.HideVip.Set(nil)
}

// UnsetHideVip ensures that no value is present for HideVip, not even an explicit nil
func (o *Settings) UnsetHideVip() {
	o.HideVip.Unset()
}

// GetInvisibleOnUserSearch returns the InvisibleOnUserSearch field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetInvisibleOnUserSearch() bool {
	if o == nil || IsNil(o.InvisibleOnUserSearch.Get()) {
		var ret bool
		return ret
	}
	return *o.InvisibleOnUserSearch.Get()
}

// GetInvisibleOnUserSearchOk returns a tuple with the InvisibleOnUserSearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetInvisibleOnUserSearchOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvisibleOnUserSearch.Get(), o.InvisibleOnUserSearch.IsSet()
}

// HasInvisibleOnUserSearch returns a boolean if a field has been set.
func (o *Settings) HasInvisibleOnUserSearch() bool {
	if o != nil && o.InvisibleOnUserSearch.IsSet() {
		return true
	}

	return false
}

// SetInvisibleOnUserSearch gets a reference to the given NullableBool and assigns it to the InvisibleOnUserSearch field.
func (o *Settings) SetInvisibleOnUserSearch(v bool) {
	o.InvisibleOnUserSearch.Set(&v)
}
// SetInvisibleOnUserSearchNil sets the value for InvisibleOnUserSearch to be an explicit nil
func (o *Settings) SetInvisibleOnUserSearchNil() {
	o.InvisibleOnUserSearch.Set(nil)
}

// UnsetInvisibleOnUserSearch ensures that no value is present for InvisibleOnUserSearch, not even an explicit nil
func (o *Settings) UnsetInvisibleOnUserSearch() {
	o.InvisibleOnUserSearch.Unset()
}

// GetNoReplyGroupTimeline returns the NoReplyGroupTimeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNoReplyGroupTimeline() bool {
	if o == nil || IsNil(o.NoReplyGroupTimeline.Get()) {
		var ret bool
		return ret
	}
	return *o.NoReplyGroupTimeline.Get()
}

// GetNoReplyGroupTimelineOk returns a tuple with the NoReplyGroupTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNoReplyGroupTimelineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoReplyGroupTimeline.Get(), o.NoReplyGroupTimeline.IsSet()
}

// HasNoReplyGroupTimeline returns a boolean if a field has been set.
func (o *Settings) HasNoReplyGroupTimeline() bool {
	if o != nil && o.NoReplyGroupTimeline.IsSet() {
		return true
	}

	return false
}

// SetNoReplyGroupTimeline gets a reference to the given NullableBool and assigns it to the NoReplyGroupTimeline field.
func (o *Settings) SetNoReplyGroupTimeline(v bool) {
	o.NoReplyGroupTimeline.Set(&v)
}
// SetNoReplyGroupTimelineNil sets the value for NoReplyGroupTimeline to be an explicit nil
func (o *Settings) SetNoReplyGroupTimelineNil() {
	o.NoReplyGroupTimeline.Set(nil)
}

// UnsetNoReplyGroupTimeline ensures that no value is present for NoReplyGroupTimeline, not even an explicit nil
func (o *Settings) UnsetNoReplyGroupTimeline() {
	o.NoReplyGroupTimeline.Unset()
}

// GetNotificationBirthdayToFollowers returns the NotificationBirthdayToFollowers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationBirthdayToFollowers() bool {
	if o == nil || IsNil(o.NotificationBirthdayToFollowers.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationBirthdayToFollowers.Get()
}

// GetNotificationBirthdayToFollowersOk returns a tuple with the NotificationBirthdayToFollowers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationBirthdayToFollowersOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationBirthdayToFollowers.Get(), o.NotificationBirthdayToFollowers.IsSet()
}

// HasNotificationBirthdayToFollowers returns a boolean if a field has been set.
func (o *Settings) HasNotificationBirthdayToFollowers() bool {
	if o != nil && o.NotificationBirthdayToFollowers.IsSet() {
		return true
	}

	return false
}

// SetNotificationBirthdayToFollowers gets a reference to the given NullableBool and assigns it to the NotificationBirthdayToFollowers field.
func (o *Settings) SetNotificationBirthdayToFollowers(v bool) {
	o.NotificationBirthdayToFollowers.Set(&v)
}
// SetNotificationBirthdayToFollowersNil sets the value for NotificationBirthdayToFollowers to be an explicit nil
func (o *Settings) SetNotificationBirthdayToFollowersNil() {
	o.NotificationBirthdayToFollowers.Set(nil)
}

// UnsetNotificationBirthdayToFollowers ensures that no value is present for NotificationBirthdayToFollowers, not even an explicit nil
func (o *Settings) UnsetNotificationBirthdayToFollowers() {
	o.NotificationBirthdayToFollowers.Unset()
}

// GetNotificationBulkCallInvite returns the NotificationBulkCallInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationBulkCallInvite() bool {
	if o == nil || IsNil(o.NotificationBulkCallInvite.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationBulkCallInvite.Get()
}

// GetNotificationBulkCallInviteOk returns a tuple with the NotificationBulkCallInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationBulkCallInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationBulkCallInvite.Get(), o.NotificationBulkCallInvite.IsSet()
}

// HasNotificationBulkCallInvite returns a boolean if a field has been set.
func (o *Settings) HasNotificationBulkCallInvite() bool {
	if o != nil && o.NotificationBulkCallInvite.IsSet() {
		return true
	}

	return false
}

// SetNotificationBulkCallInvite gets a reference to the given NullableBool and assigns it to the NotificationBulkCallInvite field.
func (o *Settings) SetNotificationBulkCallInvite(v bool) {
	o.NotificationBulkCallInvite.Set(&v)
}
// SetNotificationBulkCallInviteNil sets the value for NotificationBulkCallInvite to be an explicit nil
func (o *Settings) SetNotificationBulkCallInviteNil() {
	o.NotificationBulkCallInvite.Set(nil)
}

// UnsetNotificationBulkCallInvite ensures that no value is present for NotificationBulkCallInvite, not even an explicit nil
func (o *Settings) UnsetNotificationBulkCallInvite() {
	o.NotificationBulkCallInvite.Unset()
}

// GetNotificationCallInvite returns the NotificationCallInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationCallInvite() bool {
	if o == nil || IsNil(o.NotificationCallInvite.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationCallInvite.Get()
}

// GetNotificationCallInviteOk returns a tuple with the NotificationCallInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationCallInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationCallInvite.Get(), o.NotificationCallInvite.IsSet()
}

// HasNotificationCallInvite returns a boolean if a field has been set.
func (o *Settings) HasNotificationCallInvite() bool {
	if o != nil && o.NotificationCallInvite.IsSet() {
		return true
	}

	return false
}

// SetNotificationCallInvite gets a reference to the given NullableBool and assigns it to the NotificationCallInvite field.
func (o *Settings) SetNotificationCallInvite(v bool) {
	o.NotificationCallInvite.Set(&v)
}
// SetNotificationCallInviteNil sets the value for NotificationCallInvite to be an explicit nil
func (o *Settings) SetNotificationCallInviteNil() {
	o.NotificationCallInvite.Set(nil)
}

// UnsetNotificationCallInvite ensures that no value is present for NotificationCallInvite, not even an explicit nil
func (o *Settings) UnsetNotificationCallInvite() {
	o.NotificationCallInvite.Unset()
}

// GetNotificationChat returns the NotificationChat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationChat() bool {
	if o == nil || IsNil(o.NotificationChat.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationChat.Get()
}

// GetNotificationChatOk returns a tuple with the NotificationChat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationChatOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationChat.Get(), o.NotificationChat.IsSet()
}

// HasNotificationChat returns a boolean if a field has been set.
func (o *Settings) HasNotificationChat() bool {
	if o != nil && o.NotificationChat.IsSet() {
		return true
	}

	return false
}

// SetNotificationChat gets a reference to the given NullableBool and assigns it to the NotificationChat field.
func (o *Settings) SetNotificationChat(v bool) {
	o.NotificationChat.Set(&v)
}
// SetNotificationChatNil sets the value for NotificationChat to be an explicit nil
func (o *Settings) SetNotificationChatNil() {
	o.NotificationChat.Set(nil)
}

// UnsetNotificationChat ensures that no value is present for NotificationChat, not even an explicit nil
func (o *Settings) UnsetNotificationChat() {
	o.NotificationChat.Unset()
}

// GetNotificationChatDelete returns the NotificationChatDelete field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationChatDelete() bool {
	if o == nil || IsNil(o.NotificationChatDelete.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationChatDelete.Get()
}

// GetNotificationChatDeleteOk returns a tuple with the NotificationChatDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationChatDeleteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationChatDelete.Get(), o.NotificationChatDelete.IsSet()
}

// HasNotificationChatDelete returns a boolean if a field has been set.
func (o *Settings) HasNotificationChatDelete() bool {
	if o != nil && o.NotificationChatDelete.IsSet() {
		return true
	}

	return false
}

// SetNotificationChatDelete gets a reference to the given NullableBool and assigns it to the NotificationChatDelete field.
func (o *Settings) SetNotificationChatDelete(v bool) {
	o.NotificationChatDelete.Set(&v)
}
// SetNotificationChatDeleteNil sets the value for NotificationChatDelete to be an explicit nil
func (o *Settings) SetNotificationChatDeleteNil() {
	o.NotificationChatDelete.Set(nil)
}

// UnsetNotificationChatDelete ensures that no value is present for NotificationChatDelete, not even an explicit nil
func (o *Settings) UnsetNotificationChatDelete() {
	o.NotificationChatDelete.Unset()
}

// GetNotificationContactFriend returns the NotificationContactFriend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationContactFriend() bool {
	if o == nil || IsNil(o.NotificationContactFriend.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationContactFriend.Get()
}

// GetNotificationContactFriendOk returns a tuple with the NotificationContactFriend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationContactFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationContactFriend.Get(), o.NotificationContactFriend.IsSet()
}

// HasNotificationContactFriend returns a boolean if a field has been set.
func (o *Settings) HasNotificationContactFriend() bool {
	if o != nil && o.NotificationContactFriend.IsSet() {
		return true
	}

	return false
}

// SetNotificationContactFriend gets a reference to the given NullableBool and assigns it to the NotificationContactFriend field.
func (o *Settings) SetNotificationContactFriend(v bool) {
	o.NotificationContactFriend.Set(&v)
}
// SetNotificationContactFriendNil sets the value for NotificationContactFriend to be an explicit nil
func (o *Settings) SetNotificationContactFriendNil() {
	o.NotificationContactFriend.Set(nil)
}

// UnsetNotificationContactFriend ensures that no value is present for NotificationContactFriend, not even an explicit nil
func (o *Settings) UnsetNotificationContactFriend() {
	o.NotificationContactFriend.Unset()
}

// GetNotificationDailyQuest returns the NotificationDailyQuest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationDailyQuest() bool {
	if o == nil || IsNil(o.NotificationDailyQuest.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationDailyQuest.Get()
}

// GetNotificationDailyQuestOk returns a tuple with the NotificationDailyQuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationDailyQuestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationDailyQuest.Get(), o.NotificationDailyQuest.IsSet()
}

// HasNotificationDailyQuest returns a boolean if a field has been set.
func (o *Settings) HasNotificationDailyQuest() bool {
	if o != nil && o.NotificationDailyQuest.IsSet() {
		return true
	}

	return false
}

// SetNotificationDailyQuest gets a reference to the given NullableBool and assigns it to the NotificationDailyQuest field.
func (o *Settings) SetNotificationDailyQuest(v bool) {
	o.NotificationDailyQuest.Set(&v)
}
// SetNotificationDailyQuestNil sets the value for NotificationDailyQuest to be an explicit nil
func (o *Settings) SetNotificationDailyQuestNil() {
	o.NotificationDailyQuest.Set(nil)
}

// UnsetNotificationDailyQuest ensures that no value is present for NotificationDailyQuest, not even an explicit nil
func (o *Settings) UnsetNotificationDailyQuest() {
	o.NotificationDailyQuest.Unset()
}

// GetNotificationDailySummary returns the NotificationDailySummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationDailySummary() bool {
	if o == nil || IsNil(o.NotificationDailySummary.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationDailySummary.Get()
}

// GetNotificationDailySummaryOk returns a tuple with the NotificationDailySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationDailySummaryOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationDailySummary.Get(), o.NotificationDailySummary.IsSet()
}

// HasNotificationDailySummary returns a boolean if a field has been set.
func (o *Settings) HasNotificationDailySummary() bool {
	if o != nil && o.NotificationDailySummary.IsSet() {
		return true
	}

	return false
}

// SetNotificationDailySummary gets a reference to the given NullableBool and assigns it to the NotificationDailySummary field.
func (o *Settings) SetNotificationDailySummary(v bool) {
	o.NotificationDailySummary.Set(&v)
}
// SetNotificationDailySummaryNil sets the value for NotificationDailySummary to be an explicit nil
func (o *Settings) SetNotificationDailySummaryNil() {
	o.NotificationDailySummary.Set(nil)
}

// UnsetNotificationDailySummary ensures that no value is present for NotificationDailySummary, not even an explicit nil
func (o *Settings) UnsetNotificationDailySummary() {
	o.NotificationDailySummary.Unset()
}

// GetNotificationFollow returns the NotificationFollow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollow() bool {
	if o == nil || IsNil(o.NotificationFollow.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollow.Get()
}

// GetNotificationFollowOk returns a tuple with the NotificationFollow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollow.Get(), o.NotificationFollow.IsSet()
}

// HasNotificationFollow returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollow() bool {
	if o != nil && o.NotificationFollow.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollow gets a reference to the given NullableBool and assigns it to the NotificationFollow field.
func (o *Settings) SetNotificationFollow(v bool) {
	o.NotificationFollow.Set(&v)
}
// SetNotificationFollowNil sets the value for NotificationFollow to be an explicit nil
func (o *Settings) SetNotificationFollowNil() {
	o.NotificationFollow.Set(nil)
}

// UnsetNotificationFollow ensures that no value is present for NotificationFollow, not even an explicit nil
func (o *Settings) UnsetNotificationFollow() {
	o.NotificationFollow.Unset()
}

// GetNotificationFollowAccept returns the NotificationFollowAccept field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowAccept() bool {
	if o == nil || IsNil(o.NotificationFollowAccept.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowAccept.Get()
}

// GetNotificationFollowAcceptOk returns a tuple with the NotificationFollowAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowAcceptOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowAccept.Get(), o.NotificationFollowAccept.IsSet()
}

// HasNotificationFollowAccept returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowAccept() bool {
	if o != nil && o.NotificationFollowAccept.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowAccept gets a reference to the given NullableBool and assigns it to the NotificationFollowAccept field.
func (o *Settings) SetNotificationFollowAccept(v bool) {
	o.NotificationFollowAccept.Set(&v)
}
// SetNotificationFollowAcceptNil sets the value for NotificationFollowAccept to be an explicit nil
func (o *Settings) SetNotificationFollowAcceptNil() {
	o.NotificationFollowAccept.Set(nil)
}

// UnsetNotificationFollowAccept ensures that no value is present for NotificationFollowAccept, not even an explicit nil
func (o *Settings) UnsetNotificationFollowAccept() {
	o.NotificationFollowAccept.Unset()
}

// GetNotificationFollowRequest returns the NotificationFollowRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowRequest() bool {
	if o == nil || IsNil(o.NotificationFollowRequest.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowRequest.Get()
}

// GetNotificationFollowRequestOk returns a tuple with the NotificationFollowRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowRequest.Get(), o.NotificationFollowRequest.IsSet()
}

// HasNotificationFollowRequest returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowRequest() bool {
	if o != nil && o.NotificationFollowRequest.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowRequest gets a reference to the given NullableBool and assigns it to the NotificationFollowRequest field.
func (o *Settings) SetNotificationFollowRequest(v bool) {
	o.NotificationFollowRequest.Set(&v)
}
// SetNotificationFollowRequestNil sets the value for NotificationFollowRequest to be an explicit nil
func (o *Settings) SetNotificationFollowRequestNil() {
	o.NotificationFollowRequest.Set(nil)
}

// UnsetNotificationFollowRequest ensures that no value is present for NotificationFollowRequest, not even an explicit nil
func (o *Settings) UnsetNotificationFollowRequest() {
	o.NotificationFollowRequest.Unset()
}

// GetNotificationFollowerConferenceCall returns the NotificationFollowerConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowerConferenceCall() bool {
	if o == nil || IsNil(o.NotificationFollowerConferenceCall.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowerConferenceCall.Get()
}

// GetNotificationFollowerConferenceCallOk returns a tuple with the NotificationFollowerConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowerConferenceCallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowerConferenceCall.Get(), o.NotificationFollowerConferenceCall.IsSet()
}

// HasNotificationFollowerConferenceCall returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowerConferenceCall() bool {
	if o != nil && o.NotificationFollowerConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowerConferenceCall gets a reference to the given NullableBool and assigns it to the NotificationFollowerConferenceCall field.
func (o *Settings) SetNotificationFollowerConferenceCall(v bool) {
	o.NotificationFollowerConferenceCall.Set(&v)
}
// SetNotificationFollowerConferenceCallNil sets the value for NotificationFollowerConferenceCall to be an explicit nil
func (o *Settings) SetNotificationFollowerConferenceCallNil() {
	o.NotificationFollowerConferenceCall.Set(nil)
}

// UnsetNotificationFollowerConferenceCall ensures that no value is present for NotificationFollowerConferenceCall, not even an explicit nil
func (o *Settings) UnsetNotificationFollowerConferenceCall() {
	o.NotificationFollowerConferenceCall.Unset()
}

// GetNotificationFollowerCreateGroup returns the NotificationFollowerCreateGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowerCreateGroup() bool {
	if o == nil || IsNil(o.NotificationFollowerCreateGroup.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowerCreateGroup.Get()
}

// GetNotificationFollowerCreateGroupOk returns a tuple with the NotificationFollowerCreateGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowerCreateGroupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowerCreateGroup.Get(), o.NotificationFollowerCreateGroup.IsSet()
}

// HasNotificationFollowerCreateGroup returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowerCreateGroup() bool {
	if o != nil && o.NotificationFollowerCreateGroup.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowerCreateGroup gets a reference to the given NullableBool and assigns it to the NotificationFollowerCreateGroup field.
func (o *Settings) SetNotificationFollowerCreateGroup(v bool) {
	o.NotificationFollowerCreateGroup.Set(&v)
}
// SetNotificationFollowerCreateGroupNil sets the value for NotificationFollowerCreateGroup to be an explicit nil
func (o *Settings) SetNotificationFollowerCreateGroupNil() {
	o.NotificationFollowerCreateGroup.Set(nil)
}

// UnsetNotificationFollowerCreateGroup ensures that no value is present for NotificationFollowerCreateGroup, not even an explicit nil
func (o *Settings) UnsetNotificationFollowerCreateGroup() {
	o.NotificationFollowerCreateGroup.Unset()
}

// GetNotificationFollowingBirthdateOn returns the NotificationFollowingBirthdateOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowingBirthdateOn() bool {
	if o == nil || IsNil(o.NotificationFollowingBirthdateOn.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowingBirthdateOn.Get()
}

// GetNotificationFollowingBirthdateOnOk returns a tuple with the NotificationFollowingBirthdateOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowingBirthdateOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowingBirthdateOn.Get(), o.NotificationFollowingBirthdateOn.IsSet()
}

// HasNotificationFollowingBirthdateOn returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowingBirthdateOn() bool {
	if o != nil && o.NotificationFollowingBirthdateOn.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowingBirthdateOn gets a reference to the given NullableBool and assigns it to the NotificationFollowingBirthdateOn field.
func (o *Settings) SetNotificationFollowingBirthdateOn(v bool) {
	o.NotificationFollowingBirthdateOn.Set(&v)
}
// SetNotificationFollowingBirthdateOnNil sets the value for NotificationFollowingBirthdateOn to be an explicit nil
func (o *Settings) SetNotificationFollowingBirthdateOnNil() {
	o.NotificationFollowingBirthdateOn.Set(nil)
}

// UnsetNotificationFollowingBirthdateOn ensures that no value is present for NotificationFollowingBirthdateOn, not even an explicit nil
func (o *Settings) UnsetNotificationFollowingBirthdateOn() {
	o.NotificationFollowingBirthdateOn.Unset()
}

// GetNotificationFollowingPostAfterBreak returns the NotificationFollowingPostAfterBreak field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowingPostAfterBreak() bool {
	if o == nil || IsNil(o.NotificationFollowingPostAfterBreak.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowingPostAfterBreak.Get()
}

// GetNotificationFollowingPostAfterBreakOk returns a tuple with the NotificationFollowingPostAfterBreak field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowingPostAfterBreakOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowingPostAfterBreak.Get(), o.NotificationFollowingPostAfterBreak.IsSet()
}

// HasNotificationFollowingPostAfterBreak returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowingPostAfterBreak() bool {
	if o != nil && o.NotificationFollowingPostAfterBreak.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowingPostAfterBreak gets a reference to the given NullableBool and assigns it to the NotificationFollowingPostAfterBreak field.
func (o *Settings) SetNotificationFollowingPostAfterBreak(v bool) {
	o.NotificationFollowingPostAfterBreak.Set(&v)
}
// SetNotificationFollowingPostAfterBreakNil sets the value for NotificationFollowingPostAfterBreak to be an explicit nil
func (o *Settings) SetNotificationFollowingPostAfterBreakNil() {
	o.NotificationFollowingPostAfterBreak.Set(nil)
}

// UnsetNotificationFollowingPostAfterBreak ensures that no value is present for NotificationFollowingPostAfterBreak, not even an explicit nil
func (o *Settings) UnsetNotificationFollowingPostAfterBreak() {
	o.NotificationFollowingPostAfterBreak.Unset()
}

// GetNotificationFollowingsInCall returns the NotificationFollowingsInCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFollowingsInCall() bool {
	if o == nil || IsNil(o.NotificationFollowingsInCall.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFollowingsInCall.Get()
}

// GetNotificationFollowingsInCallOk returns a tuple with the NotificationFollowingsInCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFollowingsInCallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFollowingsInCall.Get(), o.NotificationFollowingsInCall.IsSet()
}

// HasNotificationFollowingsInCall returns a boolean if a field has been set.
func (o *Settings) HasNotificationFollowingsInCall() bool {
	if o != nil && o.NotificationFollowingsInCall.IsSet() {
		return true
	}

	return false
}

// SetNotificationFollowingsInCall gets a reference to the given NullableBool and assigns it to the NotificationFollowingsInCall field.
func (o *Settings) SetNotificationFollowingsInCall(v bool) {
	o.NotificationFollowingsInCall.Set(&v)
}
// SetNotificationFollowingsInCallNil sets the value for NotificationFollowingsInCall to be an explicit nil
func (o *Settings) SetNotificationFollowingsInCallNil() {
	o.NotificationFollowingsInCall.Set(nil)
}

// UnsetNotificationFollowingsInCall ensures that no value is present for NotificationFollowingsInCall, not even an explicit nil
func (o *Settings) UnsetNotificationFollowingsInCall() {
	o.NotificationFollowingsInCall.Unset()
}

// GetNotificationFootprint returns the NotificationFootprint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationFootprint() bool {
	if o == nil || IsNil(o.NotificationFootprint.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationFootprint.Get()
}

// GetNotificationFootprintOk returns a tuple with the NotificationFootprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationFootprintOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationFootprint.Get(), o.NotificationFootprint.IsSet()
}

// HasNotificationFootprint returns a boolean if a field has been set.
func (o *Settings) HasNotificationFootprint() bool {
	if o != nil && o.NotificationFootprint.IsSet() {
		return true
	}

	return false
}

// SetNotificationFootprint gets a reference to the given NullableBool and assigns it to the NotificationFootprint field.
func (o *Settings) SetNotificationFootprint(v bool) {
	o.NotificationFootprint.Set(&v)
}
// SetNotificationFootprintNil sets the value for NotificationFootprint to be an explicit nil
func (o *Settings) SetNotificationFootprintNil() {
	o.NotificationFootprint.Set(nil)
}

// UnsetNotificationFootprint ensures that no value is present for NotificationFootprint, not even an explicit nil
func (o *Settings) UnsetNotificationFootprint() {
	o.NotificationFootprint.Unset()
}

// GetNotificationGiftReceive returns the NotificationGiftReceive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGiftReceive() bool {
	if o == nil || IsNil(o.NotificationGiftReceive.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGiftReceive.Get()
}

// GetNotificationGiftReceiveOk returns a tuple with the NotificationGiftReceive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGiftReceiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGiftReceive.Get(), o.NotificationGiftReceive.IsSet()
}

// HasNotificationGiftReceive returns a boolean if a field has been set.
func (o *Settings) HasNotificationGiftReceive() bool {
	if o != nil && o.NotificationGiftReceive.IsSet() {
		return true
	}

	return false
}

// SetNotificationGiftReceive gets a reference to the given NullableBool and assigns it to the NotificationGiftReceive field.
func (o *Settings) SetNotificationGiftReceive(v bool) {
	o.NotificationGiftReceive.Set(&v)
}
// SetNotificationGiftReceiveNil sets the value for NotificationGiftReceive to be an explicit nil
func (o *Settings) SetNotificationGiftReceiveNil() {
	o.NotificationGiftReceive.Set(nil)
}

// UnsetNotificationGiftReceive ensures that no value is present for NotificationGiftReceive, not even an explicit nil
func (o *Settings) UnsetNotificationGiftReceive() {
	o.NotificationGiftReceive.Unset()
}

// GetNotificationGroupAccept returns the NotificationGroupAccept field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupAccept() bool {
	if o == nil || IsNil(o.NotificationGroupAccept.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupAccept.Get()
}

// GetNotificationGroupAcceptOk returns a tuple with the NotificationGroupAccept field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupAcceptOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupAccept.Get(), o.NotificationGroupAccept.IsSet()
}

// HasNotificationGroupAccept returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupAccept() bool {
	if o != nil && o.NotificationGroupAccept.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupAccept gets a reference to the given NullableBool and assigns it to the NotificationGroupAccept field.
func (o *Settings) SetNotificationGroupAccept(v bool) {
	o.NotificationGroupAccept.Set(&v)
}
// SetNotificationGroupAcceptNil sets the value for NotificationGroupAccept to be an explicit nil
func (o *Settings) SetNotificationGroupAcceptNil() {
	o.NotificationGroupAccept.Set(nil)
}

// UnsetNotificationGroupAccept ensures that no value is present for NotificationGroupAccept, not even an explicit nil
func (o *Settings) UnsetNotificationGroupAccept() {
	o.NotificationGroupAccept.Unset()
}

// GetNotificationGroupConferenceCall returns the NotificationGroupConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupConferenceCall() bool {
	if o == nil || IsNil(o.NotificationGroupConferenceCall.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupConferenceCall.Get()
}

// GetNotificationGroupConferenceCallOk returns a tuple with the NotificationGroupConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupConferenceCallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupConferenceCall.Get(), o.NotificationGroupConferenceCall.IsSet()
}

// HasNotificationGroupConferenceCall returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupConferenceCall() bool {
	if o != nil && o.NotificationGroupConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupConferenceCall gets a reference to the given NullableBool and assigns it to the NotificationGroupConferenceCall field.
func (o *Settings) SetNotificationGroupConferenceCall(v bool) {
	o.NotificationGroupConferenceCall.Set(&v)
}
// SetNotificationGroupConferenceCallNil sets the value for NotificationGroupConferenceCall to be an explicit nil
func (o *Settings) SetNotificationGroupConferenceCallNil() {
	o.NotificationGroupConferenceCall.Set(nil)
}

// UnsetNotificationGroupConferenceCall ensures that no value is present for NotificationGroupConferenceCall, not even an explicit nil
func (o *Settings) UnsetNotificationGroupConferenceCall() {
	o.NotificationGroupConferenceCall.Unset()
}

// GetNotificationGroupInvite returns the NotificationGroupInvite field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupInvite() bool {
	if o == nil || IsNil(o.NotificationGroupInvite.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupInvite.Get()
}

// GetNotificationGroupInviteOk returns a tuple with the NotificationGroupInvite field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupInviteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupInvite.Get(), o.NotificationGroupInvite.IsSet()
}

// HasNotificationGroupInvite returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupInvite() bool {
	if o != nil && o.NotificationGroupInvite.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupInvite gets a reference to the given NullableBool and assigns it to the NotificationGroupInvite field.
func (o *Settings) SetNotificationGroupInvite(v bool) {
	o.NotificationGroupInvite.Set(&v)
}
// SetNotificationGroupInviteNil sets the value for NotificationGroupInvite to be an explicit nil
func (o *Settings) SetNotificationGroupInviteNil() {
	o.NotificationGroupInvite.Set(nil)
}

// UnsetNotificationGroupInvite ensures that no value is present for NotificationGroupInvite, not even an explicit nil
func (o *Settings) UnsetNotificationGroupInvite() {
	o.NotificationGroupInvite.Unset()
}

// GetNotificationGroupJoin returns the NotificationGroupJoin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupJoin() bool {
	if o == nil || IsNil(o.NotificationGroupJoin.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupJoin.Get()
}

// GetNotificationGroupJoinOk returns a tuple with the NotificationGroupJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupJoinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupJoin.Get(), o.NotificationGroupJoin.IsSet()
}

// HasNotificationGroupJoin returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupJoin() bool {
	if o != nil && o.NotificationGroupJoin.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupJoin gets a reference to the given NullableBool and assigns it to the NotificationGroupJoin field.
func (o *Settings) SetNotificationGroupJoin(v bool) {
	o.NotificationGroupJoin.Set(&v)
}
// SetNotificationGroupJoinNil sets the value for NotificationGroupJoin to be an explicit nil
func (o *Settings) SetNotificationGroupJoinNil() {
	o.NotificationGroupJoin.Set(nil)
}

// UnsetNotificationGroupJoin ensures that no value is present for NotificationGroupJoin, not even an explicit nil
func (o *Settings) UnsetNotificationGroupJoin() {
	o.NotificationGroupJoin.Unset()
}

// GetNotificationGroupMessageTagAll returns the NotificationGroupMessageTagAll field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupMessageTagAll() bool {
	if o == nil || IsNil(o.NotificationGroupMessageTagAll.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupMessageTagAll.Get()
}

// GetNotificationGroupMessageTagAllOk returns a tuple with the NotificationGroupMessageTagAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupMessageTagAllOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupMessageTagAll.Get(), o.NotificationGroupMessageTagAll.IsSet()
}

// HasNotificationGroupMessageTagAll returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupMessageTagAll() bool {
	if o != nil && o.NotificationGroupMessageTagAll.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupMessageTagAll gets a reference to the given NullableBool and assigns it to the NotificationGroupMessageTagAll field.
func (o *Settings) SetNotificationGroupMessageTagAll(v bool) {
	o.NotificationGroupMessageTagAll.Set(&v)
}
// SetNotificationGroupMessageTagAllNil sets the value for NotificationGroupMessageTagAll to be an explicit nil
func (o *Settings) SetNotificationGroupMessageTagAllNil() {
	o.NotificationGroupMessageTagAll.Set(nil)
}

// UnsetNotificationGroupMessageTagAll ensures that no value is present for NotificationGroupMessageTagAll, not even an explicit nil
func (o *Settings) UnsetNotificationGroupMessageTagAll() {
	o.NotificationGroupMessageTagAll.Unset()
}

// GetNotificationGroupModerator returns the NotificationGroupModerator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupModerator() bool {
	if o == nil || IsNil(o.NotificationGroupModerator.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupModerator.Get()
}

// GetNotificationGroupModeratorOk returns a tuple with the NotificationGroupModerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupModeratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupModerator.Get(), o.NotificationGroupModerator.IsSet()
}

// HasNotificationGroupModerator returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupModerator() bool {
	if o != nil && o.NotificationGroupModerator.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupModerator gets a reference to the given NullableBool and assigns it to the NotificationGroupModerator field.
func (o *Settings) SetNotificationGroupModerator(v bool) {
	o.NotificationGroupModerator.Set(&v)
}
// SetNotificationGroupModeratorNil sets the value for NotificationGroupModerator to be an explicit nil
func (o *Settings) SetNotificationGroupModeratorNil() {
	o.NotificationGroupModerator.Set(nil)
}

// UnsetNotificationGroupModerator ensures that no value is present for NotificationGroupModerator, not even an explicit nil
func (o *Settings) UnsetNotificationGroupModerator() {
	o.NotificationGroupModerator.Unset()
}

// GetNotificationGroupPost returns the NotificationGroupPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupPost() bool {
	if o == nil || IsNil(o.NotificationGroupPost.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupPost.Get()
}

// GetNotificationGroupPostOk returns a tuple with the NotificationGroupPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupPost.Get(), o.NotificationGroupPost.IsSet()
}

// HasNotificationGroupPost returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupPost() bool {
	if o != nil && o.NotificationGroupPost.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupPost gets a reference to the given NullableBool and assigns it to the NotificationGroupPost field.
func (o *Settings) SetNotificationGroupPost(v bool) {
	o.NotificationGroupPost.Set(&v)
}
// SetNotificationGroupPostNil sets the value for NotificationGroupPost to be an explicit nil
func (o *Settings) SetNotificationGroupPostNil() {
	o.NotificationGroupPost.Set(nil)
}

// UnsetNotificationGroupPost ensures that no value is present for NotificationGroupPost, not even an explicit nil
func (o *Settings) UnsetNotificationGroupPost() {
	o.NotificationGroupPost.Unset()
}

// GetNotificationGroupRequest returns the NotificationGroupRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationGroupRequest() bool {
	if o == nil || IsNil(o.NotificationGroupRequest.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationGroupRequest.Get()
}

// GetNotificationGroupRequestOk returns a tuple with the NotificationGroupRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationGroupRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationGroupRequest.Get(), o.NotificationGroupRequest.IsSet()
}

// HasNotificationGroupRequest returns a boolean if a field has been set.
func (o *Settings) HasNotificationGroupRequest() bool {
	if o != nil && o.NotificationGroupRequest.IsSet() {
		return true
	}

	return false
}

// SetNotificationGroupRequest gets a reference to the given NullableBool and assigns it to the NotificationGroupRequest field.
func (o *Settings) SetNotificationGroupRequest(v bool) {
	o.NotificationGroupRequest.Set(&v)
}
// SetNotificationGroupRequestNil sets the value for NotificationGroupRequest to be an explicit nil
func (o *Settings) SetNotificationGroupRequestNil() {
	o.NotificationGroupRequest.Set(nil)
}

// UnsetNotificationGroupRequest ensures that no value is present for NotificationGroupRequest, not even an explicit nil
func (o *Settings) UnsetNotificationGroupRequest() {
	o.NotificationGroupRequest.Unset()
}

// GetNotificationHimaNow returns the NotificationHimaNow field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationHimaNow() bool {
	if o == nil || IsNil(o.NotificationHimaNow.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationHimaNow.Get()
}

// GetNotificationHimaNowOk returns a tuple with the NotificationHimaNow field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationHimaNowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationHimaNow.Get(), o.NotificationHimaNow.IsSet()
}

// HasNotificationHimaNow returns a boolean if a field has been set.
func (o *Settings) HasNotificationHimaNow() bool {
	if o != nil && o.NotificationHimaNow.IsSet() {
		return true
	}

	return false
}

// SetNotificationHimaNow gets a reference to the given NullableBool and assigns it to the NotificationHimaNow field.
func (o *Settings) SetNotificationHimaNow(v bool) {
	o.NotificationHimaNow.Set(&v)
}
// SetNotificationHimaNowNil sets the value for NotificationHimaNow to be an explicit nil
func (o *Settings) SetNotificationHimaNowNil() {
	o.NotificationHimaNow.Set(nil)
}

// UnsetNotificationHimaNow ensures that no value is present for NotificationHimaNow, not even an explicit nil
func (o *Settings) UnsetNotificationHimaNow() {
	o.NotificationHimaNow.Unset()
}

// GetNotificationInviteeCreateAccount returns the NotificationInviteeCreateAccount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationInviteeCreateAccount() bool {
	if o == nil || IsNil(o.NotificationInviteeCreateAccount.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationInviteeCreateAccount.Get()
}

// GetNotificationInviteeCreateAccountOk returns a tuple with the NotificationInviteeCreateAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationInviteeCreateAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationInviteeCreateAccount.Get(), o.NotificationInviteeCreateAccount.IsSet()
}

// HasNotificationInviteeCreateAccount returns a boolean if a field has been set.
func (o *Settings) HasNotificationInviteeCreateAccount() bool {
	if o != nil && o.NotificationInviteeCreateAccount.IsSet() {
		return true
	}

	return false
}

// SetNotificationInviteeCreateAccount gets a reference to the given NullableBool and assigns it to the NotificationInviteeCreateAccount field.
func (o *Settings) SetNotificationInviteeCreateAccount(v bool) {
	o.NotificationInviteeCreateAccount.Set(&v)
}
// SetNotificationInviteeCreateAccountNil sets the value for NotificationInviteeCreateAccount to be an explicit nil
func (o *Settings) SetNotificationInviteeCreateAccountNil() {
	o.NotificationInviteeCreateAccount.Set(nil)
}

// UnsetNotificationInviteeCreateAccount ensures that no value is present for NotificationInviteeCreateAccount, not even an explicit nil
func (o *Settings) UnsetNotificationInviteeCreateAccount() {
	o.NotificationInviteeCreateAccount.Unset()
}

// GetNotificationLatestNews returns the NotificationLatestNews field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationLatestNews() bool {
	if o == nil || IsNil(o.NotificationLatestNews.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationLatestNews.Get()
}

// GetNotificationLatestNewsOk returns a tuple with the NotificationLatestNews field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationLatestNewsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationLatestNews.Get(), o.NotificationLatestNews.IsSet()
}

// HasNotificationLatestNews returns a boolean if a field has been set.
func (o *Settings) HasNotificationLatestNews() bool {
	if o != nil && o.NotificationLatestNews.IsSet() {
		return true
	}

	return false
}

// SetNotificationLatestNews gets a reference to the given NullableBool and assigns it to the NotificationLatestNews field.
func (o *Settings) SetNotificationLatestNews(v bool) {
	o.NotificationLatestNews.Set(&v)
}
// SetNotificationLatestNewsNil sets the value for NotificationLatestNews to be an explicit nil
func (o *Settings) SetNotificationLatestNewsNil() {
	o.NotificationLatestNews.Set(nil)
}

// UnsetNotificationLatestNews ensures that no value is present for NotificationLatestNews, not even an explicit nil
func (o *Settings) UnsetNotificationLatestNews() {
	o.NotificationLatestNews.Unset()
}

// GetNotificationLike returns the NotificationLike field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationLike() bool {
	if o == nil || IsNil(o.NotificationLike.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationLike.Get()
}

// GetNotificationLikeOk returns a tuple with the NotificationLike field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationLikeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationLike.Get(), o.NotificationLike.IsSet()
}

// HasNotificationLike returns a boolean if a field has been set.
func (o *Settings) HasNotificationLike() bool {
	if o != nil && o.NotificationLike.IsSet() {
		return true
	}

	return false
}

// SetNotificationLike gets a reference to the given NullableBool and assigns it to the NotificationLike field.
func (o *Settings) SetNotificationLike(v bool) {
	o.NotificationLike.Set(&v)
}
// SetNotificationLikeNil sets the value for NotificationLike to be an explicit nil
func (o *Settings) SetNotificationLikeNil() {
	o.NotificationLike.Set(nil)
}

// UnsetNotificationLike ensures that no value is present for NotificationLike, not even an explicit nil
func (o *Settings) UnsetNotificationLike() {
	o.NotificationLike.Unset()
}

// GetNotificationMessageTag returns the NotificationMessageTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationMessageTag() bool {
	if o == nil || IsNil(o.NotificationMessageTag.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationMessageTag.Get()
}

// GetNotificationMessageTagOk returns a tuple with the NotificationMessageTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationMessageTagOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationMessageTag.Get(), o.NotificationMessageTag.IsSet()
}

// HasNotificationMessageTag returns a boolean if a field has been set.
func (o *Settings) HasNotificationMessageTag() bool {
	if o != nil && o.NotificationMessageTag.IsSet() {
		return true
	}

	return false
}

// SetNotificationMessageTag gets a reference to the given NullableBool and assigns it to the NotificationMessageTag field.
func (o *Settings) SetNotificationMessageTag(v bool) {
	o.NotificationMessageTag.Set(&v)
}
// SetNotificationMessageTagNil sets the value for NotificationMessageTag to be an explicit nil
func (o *Settings) SetNotificationMessageTagNil() {
	o.NotificationMessageTag.Set(nil)
}

// UnsetNotificationMessageTag ensures that no value is present for NotificationMessageTag, not even an explicit nil
func (o *Settings) UnsetNotificationMessageTag() {
	o.NotificationMessageTag.Unset()
}

// GetNotificationPopularPost returns the NotificationPopularPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationPopularPost() bool {
	if o == nil || IsNil(o.NotificationPopularPost.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationPopularPost.Get()
}

// GetNotificationPopularPostOk returns a tuple with the NotificationPopularPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationPopularPostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationPopularPost.Get(), o.NotificationPopularPost.IsSet()
}

// HasNotificationPopularPost returns a boolean if a field has been set.
func (o *Settings) HasNotificationPopularPost() bool {
	if o != nil && o.NotificationPopularPost.IsSet() {
		return true
	}

	return false
}

// SetNotificationPopularPost gets a reference to the given NullableBool and assigns it to the NotificationPopularPost field.
func (o *Settings) SetNotificationPopularPost(v bool) {
	o.NotificationPopularPost.Set(&v)
}
// SetNotificationPopularPostNil sets the value for NotificationPopularPost to be an explicit nil
func (o *Settings) SetNotificationPopularPostNil() {
	o.NotificationPopularPost.Set(nil)
}

// UnsetNotificationPopularPost ensures that no value is present for NotificationPopularPost, not even an explicit nil
func (o *Settings) UnsetNotificationPopularPost() {
	o.NotificationPopularPost.Unset()
}

// GetNotificationProfileScreenshot returns the NotificationProfileScreenshot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationProfileScreenshot() bool {
	if o == nil || IsNil(o.NotificationProfileScreenshot.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationProfileScreenshot.Get()
}

// GetNotificationProfileScreenshotOk returns a tuple with the NotificationProfileScreenshot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationProfileScreenshotOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationProfileScreenshot.Get(), o.NotificationProfileScreenshot.IsSet()
}

// HasNotificationProfileScreenshot returns a boolean if a field has been set.
func (o *Settings) HasNotificationProfileScreenshot() bool {
	if o != nil && o.NotificationProfileScreenshot.IsSet() {
		return true
	}

	return false
}

// SetNotificationProfileScreenshot gets a reference to the given NullableBool and assigns it to the NotificationProfileScreenshot field.
func (o *Settings) SetNotificationProfileScreenshot(v bool) {
	o.NotificationProfileScreenshot.Set(&v)
}
// SetNotificationProfileScreenshotNil sets the value for NotificationProfileScreenshot to be an explicit nil
func (o *Settings) SetNotificationProfileScreenshotNil() {
	o.NotificationProfileScreenshot.Set(nil)
}

// UnsetNotificationProfileScreenshot ensures that no value is present for NotificationProfileScreenshot, not even an explicit nil
func (o *Settings) UnsetNotificationProfileScreenshot() {
	o.NotificationProfileScreenshot.Unset()
}

// GetNotificationReply returns the NotificationReply field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationReply() bool {
	if o == nil || IsNil(o.NotificationReply.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationReply.Get()
}

// GetNotificationReplyOk returns a tuple with the NotificationReply field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationReplyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationReply.Get(), o.NotificationReply.IsSet()
}

// HasNotificationReply returns a boolean if a field has been set.
func (o *Settings) HasNotificationReply() bool {
	if o != nil && o.NotificationReply.IsSet() {
		return true
	}

	return false
}

// SetNotificationReply gets a reference to the given NullableBool and assigns it to the NotificationReply field.
func (o *Settings) SetNotificationReply(v bool) {
	o.NotificationReply.Set(&v)
}
// SetNotificationReplyNil sets the value for NotificationReply to be an explicit nil
func (o *Settings) SetNotificationReplyNil() {
	o.NotificationReply.Set(nil)
}

// UnsetNotificationReply ensures that no value is present for NotificationReply, not even an explicit nil
func (o *Settings) UnsetNotificationReply() {
	o.NotificationReply.Unset()
}

// GetNotificationRepost returns the NotificationRepost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationRepost() bool {
	if o == nil || IsNil(o.NotificationRepost.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationRepost.Get()
}

// GetNotificationRepostOk returns a tuple with the NotificationRepost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationRepostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationRepost.Get(), o.NotificationRepost.IsSet()
}

// HasNotificationRepost returns a boolean if a field has been set.
func (o *Settings) HasNotificationRepost() bool {
	if o != nil && o.NotificationRepost.IsSet() {
		return true
	}

	return false
}

// SetNotificationRepost gets a reference to the given NullableBool and assigns it to the NotificationRepost field.
func (o *Settings) SetNotificationRepost(v bool) {
	o.NotificationRepost.Set(&v)
}
// SetNotificationRepostNil sets the value for NotificationRepost to be an explicit nil
func (o *Settings) SetNotificationRepostNil() {
	o.NotificationRepost.Set(nil)
}

// UnsetNotificationRepost ensures that no value is present for NotificationRepost, not even an explicit nil
func (o *Settings) UnsetNotificationRepost() {
	o.NotificationRepost.Unset()
}

// GetNotificationReview returns the NotificationReview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationReview() bool {
	if o == nil || IsNil(o.NotificationReview.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationReview.Get()
}

// GetNotificationReviewOk returns a tuple with the NotificationReview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationReviewOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationReview.Get(), o.NotificationReview.IsSet()
}

// HasNotificationReview returns a boolean if a field has been set.
func (o *Settings) HasNotificationReview() bool {
	if o != nil && o.NotificationReview.IsSet() {
		return true
	}

	return false
}

// SetNotificationReview gets a reference to the given NullableBool and assigns it to the NotificationReview field.
func (o *Settings) SetNotificationReview(v bool) {
	o.NotificationReview.Set(&v)
}
// SetNotificationReviewNil sets the value for NotificationReview to be an explicit nil
func (o *Settings) SetNotificationReviewNil() {
	o.NotificationReview.Set(nil)
}

// UnsetNotificationReview ensures that no value is present for NotificationReview, not even an explicit nil
func (o *Settings) UnsetNotificationReview() {
	o.NotificationReview.Unset()
}

// GetNotificationSecurityWarning returns the NotificationSecurityWarning field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationSecurityWarning() bool {
	if o == nil || IsNil(o.NotificationSecurityWarning.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationSecurityWarning.Get()
}

// GetNotificationSecurityWarningOk returns a tuple with the NotificationSecurityWarning field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationSecurityWarningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationSecurityWarning.Get(), o.NotificationSecurityWarning.IsSet()
}

// HasNotificationSecurityWarning returns a boolean if a field has been set.
func (o *Settings) HasNotificationSecurityWarning() bool {
	if o != nil && o.NotificationSecurityWarning.IsSet() {
		return true
	}

	return false
}

// SetNotificationSecurityWarning gets a reference to the given NullableBool and assigns it to the NotificationSecurityWarning field.
func (o *Settings) SetNotificationSecurityWarning(v bool) {
	o.NotificationSecurityWarning.Set(&v)
}
// SetNotificationSecurityWarningNil sets the value for NotificationSecurityWarning to be an explicit nil
func (o *Settings) SetNotificationSecurityWarningNil() {
	o.NotificationSecurityWarning.Set(nil)
}

// UnsetNotificationSecurityWarning ensures that no value is present for NotificationSecurityWarning, not even an explicit nil
func (o *Settings) UnsetNotificationSecurityWarning() {
	o.NotificationSecurityWarning.Unset()
}

// GetNotificationTwitterFriend returns the NotificationTwitterFriend field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetNotificationTwitterFriend() bool {
	if o == nil || IsNil(o.NotificationTwitterFriend.Get()) {
		var ret bool
		return ret
	}
	return *o.NotificationTwitterFriend.Get()
}

// GetNotificationTwitterFriendOk returns a tuple with the NotificationTwitterFriend field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetNotificationTwitterFriendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationTwitterFriend.Get(), o.NotificationTwitterFriend.IsSet()
}

// HasNotificationTwitterFriend returns a boolean if a field has been set.
func (o *Settings) HasNotificationTwitterFriend() bool {
	if o != nil && o.NotificationTwitterFriend.IsSet() {
		return true
	}

	return false
}

// SetNotificationTwitterFriend gets a reference to the given NullableBool and assigns it to the NotificationTwitterFriend field.
func (o *Settings) SetNotificationTwitterFriend(v bool) {
	o.NotificationTwitterFriend.Set(&v)
}
// SetNotificationTwitterFriendNil sets the value for NotificationTwitterFriend to be an explicit nil
func (o *Settings) SetNotificationTwitterFriendNil() {
	o.NotificationTwitterFriend.Set(nil)
}

// UnsetNotificationTwitterFriend ensures that no value is present for NotificationTwitterFriend, not even an explicit nil
func (o *Settings) UnsetNotificationTwitterFriend() {
	o.NotificationTwitterFriend.Unset()
}

// GetPrivacyMode returns the PrivacyMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetPrivacyMode() bool {
	if o == nil || IsNil(o.PrivacyMode.Get()) {
		var ret bool
		return ret
	}
	return *o.PrivacyMode.Get()
}

// GetPrivacyModeOk returns a tuple with the PrivacyMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetPrivacyModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivacyMode.Get(), o.PrivacyMode.IsSet()
}

// HasPrivacyMode returns a boolean if a field has been set.
func (o *Settings) HasPrivacyMode() bool {
	if o != nil && o.PrivacyMode.IsSet() {
		return true
	}

	return false
}

// SetPrivacyMode gets a reference to the given NullableBool and assigns it to the PrivacyMode field.
func (o *Settings) SetPrivacyMode(v bool) {
	o.PrivacyMode.Set(&v)
}
// SetPrivacyModeNil sets the value for PrivacyMode to be an explicit nil
func (o *Settings) SetPrivacyModeNil() {
	o.PrivacyMode.Set(nil)
}

// UnsetPrivacyMode ensures that no value is present for PrivacyMode, not even an explicit nil
func (o *Settings) UnsetPrivacyMode() {
	o.PrivacyMode.Unset()
}

// GetPrivatePost returns the PrivatePost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetPrivatePost() bool {
	if o == nil || IsNil(o.PrivatePost.Get()) {
		var ret bool
		return ret
	}
	return *o.PrivatePost.Get()
}

// GetPrivatePostOk returns a tuple with the PrivatePost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetPrivatePostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivatePost.Get(), o.PrivatePost.IsSet()
}

// HasPrivatePost returns a boolean if a field has been set.
func (o *Settings) HasPrivatePost() bool {
	if o != nil && o.PrivatePost.IsSet() {
		return true
	}

	return false
}

// SetPrivatePost gets a reference to the given NullableBool and assigns it to the PrivatePost field.
func (o *Settings) SetPrivatePost(v bool) {
	o.PrivatePost.Set(&v)
}
// SetPrivatePostNil sets the value for PrivatePost to be an explicit nil
func (o *Settings) SetPrivatePostNil() {
	o.PrivatePost.Set(nil)
}

// UnsetPrivatePost ensures that no value is present for PrivatePost, not even an explicit nil
func (o *Settings) UnsetPrivatePost() {
	o.PrivatePost.Unset()
}

// GetPrivateUserTimeline returns the PrivateUserTimeline field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetPrivateUserTimeline() bool {
	if o == nil || IsNil(o.PrivateUserTimeline.Get()) {
		var ret bool
		return ret
	}
	return *o.PrivateUserTimeline.Get()
}

// GetPrivateUserTimelineOk returns a tuple with the PrivateUserTimeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetPrivateUserTimelineOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivateUserTimeline.Get(), o.PrivateUserTimeline.IsSet()
}

// HasPrivateUserTimeline returns a boolean if a field has been set.
func (o *Settings) HasPrivateUserTimeline() bool {
	if o != nil && o.PrivateUserTimeline.IsSet() {
		return true
	}

	return false
}

// SetPrivateUserTimeline gets a reference to the given NullableBool and assigns it to the PrivateUserTimeline field.
func (o *Settings) SetPrivateUserTimeline(v bool) {
	o.PrivateUserTimeline.Set(&v)
}
// SetPrivateUserTimelineNil sets the value for PrivateUserTimeline to be an explicit nil
func (o *Settings) SetPrivateUserTimelineNil() {
	o.PrivateUserTimeline.Set(nil)
}

// UnsetPrivateUserTimeline ensures that no value is present for PrivateUserTimeline, not even an explicit nil
func (o *Settings) UnsetPrivateUserTimeline() {
	o.PrivateUserTimeline.Unset()
}

// GetShowTotalGiftsReceivedCount returns the ShowTotalGiftsReceivedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetShowTotalGiftsReceivedCount() bool {
	if o == nil || IsNil(o.ShowTotalGiftsReceivedCount.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowTotalGiftsReceivedCount.Get()
}

// GetShowTotalGiftsReceivedCountOk returns a tuple with the ShowTotalGiftsReceivedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetShowTotalGiftsReceivedCountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowTotalGiftsReceivedCount.Get(), o.ShowTotalGiftsReceivedCount.IsSet()
}

// HasShowTotalGiftsReceivedCount returns a boolean if a field has been set.
func (o *Settings) HasShowTotalGiftsReceivedCount() bool {
	if o != nil && o.ShowTotalGiftsReceivedCount.IsSet() {
		return true
	}

	return false
}

// SetShowTotalGiftsReceivedCount gets a reference to the given NullableBool and assigns it to the ShowTotalGiftsReceivedCount field.
func (o *Settings) SetShowTotalGiftsReceivedCount(v bool) {
	o.ShowTotalGiftsReceivedCount.Set(&v)
}
// SetShowTotalGiftsReceivedCountNil sets the value for ShowTotalGiftsReceivedCount to be an explicit nil
func (o *Settings) SetShowTotalGiftsReceivedCountNil() {
	o.ShowTotalGiftsReceivedCount.Set(nil)
}

// UnsetShowTotalGiftsReceivedCount ensures that no value is present for ShowTotalGiftsReceivedCount, not even an explicit nil
func (o *Settings) UnsetShowTotalGiftsReceivedCount() {
	o.ShowTotalGiftsReceivedCount.Unset()
}

// GetVipInvisibleFootprintMode returns the VipInvisibleFootprintMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetVipInvisibleFootprintMode() bool {
	if o == nil || IsNil(o.VipInvisibleFootprintMode.Get()) {
		var ret bool
		return ret
	}
	return *o.VipInvisibleFootprintMode.Get()
}

// GetVipInvisibleFootprintModeOk returns a tuple with the VipInvisibleFootprintMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetVipInvisibleFootprintModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipInvisibleFootprintMode.Get(), o.VipInvisibleFootprintMode.IsSet()
}

// HasVipInvisibleFootprintMode returns a boolean if a field has been set.
func (o *Settings) HasVipInvisibleFootprintMode() bool {
	if o != nil && o.VipInvisibleFootprintMode.IsSet() {
		return true
	}

	return false
}

// SetVipInvisibleFootprintMode gets a reference to the given NullableBool and assigns it to the VipInvisibleFootprintMode field.
func (o *Settings) SetVipInvisibleFootprintMode(v bool) {
	o.VipInvisibleFootprintMode.Set(&v)
}
// SetVipInvisibleFootprintModeNil sets the value for VipInvisibleFootprintMode to be an explicit nil
func (o *Settings) SetVipInvisibleFootprintModeNil() {
	o.VipInvisibleFootprintMode.Set(nil)
}

// UnsetVipInvisibleFootprintMode ensures that no value is present for VipInvisibleFootprintMode, not even an explicit nil
func (o *Settings) UnsetVipInvisibleFootprintMode() {
	o.VipInvisibleFootprintMode.Unset()
}

// GetVisibleOnSnsFriendRecommendation returns the VisibleOnSnsFriendRecommendation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Settings) GetVisibleOnSnsFriendRecommendation() bool {
	if o == nil || IsNil(o.VisibleOnSnsFriendRecommendation.Get()) {
		var ret bool
		return ret
	}
	return *o.VisibleOnSnsFriendRecommendation.Get()
}

// GetVisibleOnSnsFriendRecommendationOk returns a tuple with the VisibleOnSnsFriendRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Settings) GetVisibleOnSnsFriendRecommendationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisibleOnSnsFriendRecommendation.Get(), o.VisibleOnSnsFriendRecommendation.IsSet()
}

// HasVisibleOnSnsFriendRecommendation returns a boolean if a field has been set.
func (o *Settings) HasVisibleOnSnsFriendRecommendation() bool {
	if o != nil && o.VisibleOnSnsFriendRecommendation.IsSet() {
		return true
	}

	return false
}

// SetVisibleOnSnsFriendRecommendation gets a reference to the given NullableBool and assigns it to the VisibleOnSnsFriendRecommendation field.
func (o *Settings) SetVisibleOnSnsFriendRecommendation(v bool) {
	o.VisibleOnSnsFriendRecommendation.Set(&v)
}
// SetVisibleOnSnsFriendRecommendationNil sets the value for VisibleOnSnsFriendRecommendation to be an explicit nil
func (o *Settings) SetVisibleOnSnsFriendRecommendationNil() {
	o.VisibleOnSnsFriendRecommendation.Set(nil)
}

// UnsetVisibleOnSnsFriendRecommendation ensures that no value is present for VisibleOnSnsFriendRecommendation, not even an explicit nil
func (o *Settings) UnsetVisibleOnSnsFriendRecommendation() {
	o.VisibleOnSnsFriendRecommendation.Unset()
}

func (o Settings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Settings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgeRestrictedOnReview.IsSet() {
		toSerialize["age_restricted_on_review"] = o.AgeRestrictedOnReview.Get()
	}
	if o.AllowReposts.IsSet() {
		toSerialize["allow_reposts"] = o.AllowReposts.Get()
	}
	if o.CautionUserChat.IsSet() {
		toSerialize["caution_user_chat"] = o.CautionUserChat.Get()
	}
	if o.FollowingOnlyCallInvite.IsSet() {
		toSerialize["following_only_call_invite"] = o.FollowingOnlyCallInvite.Get()
	}
	if o.FollowingOnlyGroupInvite.IsSet() {
		toSerialize["following_only_group_invite"] = o.FollowingOnlyGroupInvite.Get()
	}
	if o.FollowingRestrictedOnReview.IsSet() {
		toSerialize["following_restricted_on_review"] = o.FollowingRestrictedOnReview.Get()
	}
	if o.HideActiveCall.IsSet() {
		toSerialize["hide_active_call"] = o.HideActiveCall.Get()
	}
	if o.HideGiftsReceived.IsSet() {
		toSerialize["hide_gifts_received"] = o.HideGiftsReceived.Get()
	}
	if o.HideHotPost.IsSet() {
		toSerialize["hide_hot_post"] = o.HideHotPost.Get()
	}
	if o.HideOnInvitable.IsSet() {
		toSerialize["hide_on_invitable"] = o.HideOnInvitable.Get()
	}
	if o.HideOnlineStatus.IsSet() {
		toSerialize["hide_online_status"] = o.HideOnlineStatus.Get()
	}
	if o.HideReplyFollowingTimeline.IsSet() {
		toSerialize["hide_reply_following_timeline"] = o.HideReplyFollowingTimeline.Get()
	}
	if o.HideReplyPublicTimeline.IsSet() {
		toSerialize["hide_reply_public_timeline"] = o.HideReplyPublicTimeline.Get()
	}
	if o.HideVip.IsSet() {
		toSerialize["hide_vip"] = o.HideVip.Get()
	}
	if o.InvisibleOnUserSearch.IsSet() {
		toSerialize["invisible_on_user_search"] = o.InvisibleOnUserSearch.Get()
	}
	if o.NoReplyGroupTimeline.IsSet() {
		toSerialize["no_reply_group_timeline"] = o.NoReplyGroupTimeline.Get()
	}
	if o.NotificationBirthdayToFollowers.IsSet() {
		toSerialize["notification_birthday_to_followers"] = o.NotificationBirthdayToFollowers.Get()
	}
	if o.NotificationBulkCallInvite.IsSet() {
		toSerialize["notification_bulk_call_invite"] = o.NotificationBulkCallInvite.Get()
	}
	if o.NotificationCallInvite.IsSet() {
		toSerialize["notification_call_invite"] = o.NotificationCallInvite.Get()
	}
	if o.NotificationChat.IsSet() {
		toSerialize["notification_chat"] = o.NotificationChat.Get()
	}
	if o.NotificationChatDelete.IsSet() {
		toSerialize["notification_chat_delete"] = o.NotificationChatDelete.Get()
	}
	if o.NotificationContactFriend.IsSet() {
		toSerialize["notification_contact_friend"] = o.NotificationContactFriend.Get()
	}
	if o.NotificationDailyQuest.IsSet() {
		toSerialize["notification_daily_quest"] = o.NotificationDailyQuest.Get()
	}
	if o.NotificationDailySummary.IsSet() {
		toSerialize["notification_daily_summary"] = o.NotificationDailySummary.Get()
	}
	if o.NotificationFollow.IsSet() {
		toSerialize["notification_follow"] = o.NotificationFollow.Get()
	}
	if o.NotificationFollowAccept.IsSet() {
		toSerialize["notification_follow_accept"] = o.NotificationFollowAccept.Get()
	}
	if o.NotificationFollowRequest.IsSet() {
		toSerialize["notification_follow_request"] = o.NotificationFollowRequest.Get()
	}
	if o.NotificationFollowerConferenceCall.IsSet() {
		toSerialize["notification_follower_conference_call"] = o.NotificationFollowerConferenceCall.Get()
	}
	if o.NotificationFollowerCreateGroup.IsSet() {
		toSerialize["notification_follower_create_group"] = o.NotificationFollowerCreateGroup.Get()
	}
	if o.NotificationFollowingBirthdateOn.IsSet() {
		toSerialize["notification_following_birthdate_on"] = o.NotificationFollowingBirthdateOn.Get()
	}
	if o.NotificationFollowingPostAfterBreak.IsSet() {
		toSerialize["notification_following_post_after_break"] = o.NotificationFollowingPostAfterBreak.Get()
	}
	if o.NotificationFollowingsInCall.IsSet() {
		toSerialize["notification_followings_in_call"] = o.NotificationFollowingsInCall.Get()
	}
	if o.NotificationFootprint.IsSet() {
		toSerialize["notification_footprint"] = o.NotificationFootprint.Get()
	}
	if o.NotificationGiftReceive.IsSet() {
		toSerialize["notification_gift_receive"] = o.NotificationGiftReceive.Get()
	}
	if o.NotificationGroupAccept.IsSet() {
		toSerialize["notification_group_accept"] = o.NotificationGroupAccept.Get()
	}
	if o.NotificationGroupConferenceCall.IsSet() {
		toSerialize["notification_group_conference_call"] = o.NotificationGroupConferenceCall.Get()
	}
	if o.NotificationGroupInvite.IsSet() {
		toSerialize["notification_group_invite"] = o.NotificationGroupInvite.Get()
	}
	if o.NotificationGroupJoin.IsSet() {
		toSerialize["notification_group_join"] = o.NotificationGroupJoin.Get()
	}
	if o.NotificationGroupMessageTagAll.IsSet() {
		toSerialize["notification_group_message_tag_all"] = o.NotificationGroupMessageTagAll.Get()
	}
	if o.NotificationGroupModerator.IsSet() {
		toSerialize["notification_group_moderator"] = o.NotificationGroupModerator.Get()
	}
	if o.NotificationGroupPost.IsSet() {
		toSerialize["notification_group_post"] = o.NotificationGroupPost.Get()
	}
	if o.NotificationGroupRequest.IsSet() {
		toSerialize["notification_group_request"] = o.NotificationGroupRequest.Get()
	}
	if o.NotificationHimaNow.IsSet() {
		toSerialize["notification_hima_now"] = o.NotificationHimaNow.Get()
	}
	if o.NotificationInviteeCreateAccount.IsSet() {
		toSerialize["notification_invitee_create_account"] = o.NotificationInviteeCreateAccount.Get()
	}
	if o.NotificationLatestNews.IsSet() {
		toSerialize["notification_latest_news"] = o.NotificationLatestNews.Get()
	}
	if o.NotificationLike.IsSet() {
		toSerialize["notification_like"] = o.NotificationLike.Get()
	}
	if o.NotificationMessageTag.IsSet() {
		toSerialize["notification_message_tag"] = o.NotificationMessageTag.Get()
	}
	if o.NotificationPopularPost.IsSet() {
		toSerialize["notification_popular_post"] = o.NotificationPopularPost.Get()
	}
	if o.NotificationProfileScreenshot.IsSet() {
		toSerialize["notification_profile_screenshot"] = o.NotificationProfileScreenshot.Get()
	}
	if o.NotificationReply.IsSet() {
		toSerialize["notification_reply"] = o.NotificationReply.Get()
	}
	if o.NotificationRepost.IsSet() {
		toSerialize["notification_repost"] = o.NotificationRepost.Get()
	}
	if o.NotificationReview.IsSet() {
		toSerialize["notification_review"] = o.NotificationReview.Get()
	}
	if o.NotificationSecurityWarning.IsSet() {
		toSerialize["notification_security_warning"] = o.NotificationSecurityWarning.Get()
	}
	if o.NotificationTwitterFriend.IsSet() {
		toSerialize["notification_twitter_friend"] = o.NotificationTwitterFriend.Get()
	}
	if o.PrivacyMode.IsSet() {
		toSerialize["privacy_mode"] = o.PrivacyMode.Get()
	}
	if o.PrivatePost.IsSet() {
		toSerialize["private_post"] = o.PrivatePost.Get()
	}
	if o.PrivateUserTimeline.IsSet() {
		toSerialize["private_user_timeline"] = o.PrivateUserTimeline.Get()
	}
	if o.ShowTotalGiftsReceivedCount.IsSet() {
		toSerialize["show_total_gifts_received_count"] = o.ShowTotalGiftsReceivedCount.Get()
	}
	if o.VipInvisibleFootprintMode.IsSet() {
		toSerialize["vip_invisible_footprint_mode"] = o.VipInvisibleFootprintMode.Get()
	}
	if o.VisibleOnSnsFriendRecommendation.IsSet() {
		toSerialize["visible_on_sns_friend_recommendation"] = o.VisibleOnSnsFriendRecommendation.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Settings) UnmarshalJSON(data []byte) (err error) {
	varSettings := _Settings{}

	err = json.Unmarshal(data, &varSettings)

	if err != nil {
		return err
	}

	*o = Settings(varSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "age_restricted_on_review")
		delete(additionalProperties, "allow_reposts")
		delete(additionalProperties, "caution_user_chat")
		delete(additionalProperties, "following_only_call_invite")
		delete(additionalProperties, "following_only_group_invite")
		delete(additionalProperties, "following_restricted_on_review")
		delete(additionalProperties, "hide_active_call")
		delete(additionalProperties, "hide_gifts_received")
		delete(additionalProperties, "hide_hot_post")
		delete(additionalProperties, "hide_on_invitable")
		delete(additionalProperties, "hide_online_status")
		delete(additionalProperties, "hide_reply_following_timeline")
		delete(additionalProperties, "hide_reply_public_timeline")
		delete(additionalProperties, "hide_vip")
		delete(additionalProperties, "invisible_on_user_search")
		delete(additionalProperties, "no_reply_group_timeline")
		delete(additionalProperties, "notification_birthday_to_followers")
		delete(additionalProperties, "notification_bulk_call_invite")
		delete(additionalProperties, "notification_call_invite")
		delete(additionalProperties, "notification_chat")
		delete(additionalProperties, "notification_chat_delete")
		delete(additionalProperties, "notification_contact_friend")
		delete(additionalProperties, "notification_daily_quest")
		delete(additionalProperties, "notification_daily_summary")
		delete(additionalProperties, "notification_follow")
		delete(additionalProperties, "notification_follow_accept")
		delete(additionalProperties, "notification_follow_request")
		delete(additionalProperties, "notification_follower_conference_call")
		delete(additionalProperties, "notification_follower_create_group")
		delete(additionalProperties, "notification_following_birthdate_on")
		delete(additionalProperties, "notification_following_post_after_break")
		delete(additionalProperties, "notification_followings_in_call")
		delete(additionalProperties, "notification_footprint")
		delete(additionalProperties, "notification_gift_receive")
		delete(additionalProperties, "notification_group_accept")
		delete(additionalProperties, "notification_group_conference_call")
		delete(additionalProperties, "notification_group_invite")
		delete(additionalProperties, "notification_group_join")
		delete(additionalProperties, "notification_group_message_tag_all")
		delete(additionalProperties, "notification_group_moderator")
		delete(additionalProperties, "notification_group_post")
		delete(additionalProperties, "notification_group_request")
		delete(additionalProperties, "notification_hima_now")
		delete(additionalProperties, "notification_invitee_create_account")
		delete(additionalProperties, "notification_latest_news")
		delete(additionalProperties, "notification_like")
		delete(additionalProperties, "notification_message_tag")
		delete(additionalProperties, "notification_popular_post")
		delete(additionalProperties, "notification_profile_screenshot")
		delete(additionalProperties, "notification_reply")
		delete(additionalProperties, "notification_repost")
		delete(additionalProperties, "notification_review")
		delete(additionalProperties, "notification_security_warning")
		delete(additionalProperties, "notification_twitter_friend")
		delete(additionalProperties, "privacy_mode")
		delete(additionalProperties, "private_post")
		delete(additionalProperties, "private_user_timeline")
		delete(additionalProperties, "show_total_gifts_received_count")
		delete(additionalProperties, "vip_invisible_footprint_mode")
		delete(additionalProperties, "visible_on_sns_friend_recommendation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSettings struct {
	value *Settings
	isSet bool
}

func (v NullableSettings) Get() *Settings {
	return v.value
}

func (v *NullableSettings) Set(val *Settings) {
	v.value = val
	v.isSet = true
}

func (v NullableSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettings(val *Settings) *NullableSettings {
	return &NullableSettings{value: val, isSet: true}
}

func (v NullableSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


