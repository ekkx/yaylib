
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Group type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Group{}

// Group struct for Group
type Group struct {
	AllowMembersToPostImageAndVideo NullableBool `json:"allow_members_to_post_image_and_video,omitempty"`
	AllowMembersToPostUrl NullableBool `json:"allow_members_to_post_url,omitempty"`
	AllowOwnershipTransfer NullableBool `json:"allow_ownership_transfer,omitempty"`
	AllowThreadCreationBy NullableString `json:"allow_thread_creation_by,omitempty"`
	CallTimelineDisplay NullableBool `json:"call_timeline_display,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	CoverImageThumbnail NullableString `json:"cover_image_thumbnail,omitempty"`
	Description NullableString `json:"description,omitempty"`
	Gender NullableInt32 `json:"gender,omitempty"`
	GenerationGroupsLimit NullableInt32 `json:"generation_groups_limit,omitempty"`
	GroupCategoryId NullableInt64 `json:"group_category_id,omitempty"`
	GroupIcon NullableString `json:"group_icon,omitempty"`
	GroupIconThumbnail NullableString `json:"group_icon_thumbnail,omitempty"`
	GroupsUsersCount NullableInt64 `json:"groups_users_count,omitempty"`
	Guidelines NullableString `json:"guidelines,omitempty"`
	HideConferenceCall NullableBool `json:"hide_conference_call,omitempty"`
	HideFromGameEight NullableBool `json:"hide_from_game_eight,omitempty"`
	HideReportedPosts NullableBool `json:"hide_reported_posts,omitempty"`
	HighlightedCount NullableInt64 `json:"highlighted_count,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	InvitedToJoin NullableBool `json:"invited_to_join,omitempty"`
	IsJoined NullableBool `json:"is_joined,omitempty"`
	IsPending NullableBool `json:"is_pending,omitempty"`
	IsPrivate NullableBool `json:"is_private,omitempty"`
	IsRelated NullableBool `json:"is_related,omitempty"`
	JoinedCommunityCampaign NullableBool `json:"joined_community_campaign,omitempty"`
	ModeratorIds []int64 `json:"moderator_ids,omitempty"`
	OnlyMobileVerified NullableBool `json:"only_mobile_verified,omitempty"`
	OnlyVerifiedAge NullableBool `json:"only_verified_age,omitempty"`
	Owner NullableRealmUser `json:"owner,omitempty"`
	PendingCount NullableInt64 `json:"pending_count,omitempty"`
	PendingDeputizeIds []int64 `json:"pending_deputize_ids,omitempty"`
	PendingTransferId NullableInt64 `json:"pending_transfer_id,omitempty"`
	PostsCount NullableInt64 `json:"posts_count,omitempty"`
	RelatedCount NullableInt64 `json:"related_count,omitempty"`
	SafeMode NullableBool `json:"safe_mode,omitempty"`
	Secret NullableBool `json:"secret,omitempty"`
	Seizable NullableBool `json:"seizable,omitempty"`
	SeizableBefore NullableInt64 `json:"seizable_before,omitempty"`
	SubCategoryId NullableInt64 `json:"sub_category_id,omitempty"`
	ThreadsCount NullableInt64 `json:"threads_count,omitempty"`
	Title NullableString `json:"title,omitempty"`
	Topic NullableString `json:"topic,omitempty"`
	UnreadCounts NullableInt64 `json:"unread_counts,omitempty"`
	UnreadThreadsCount NullableInt64 `json:"unread_threads_count,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	UserId NullableInt64 `json:"user_id,omitempty"`
	ViewsCount NullableInt64 `json:"views_count,omitempty"`
	WalkthroughRequested NullableBool `json:"walkthrough_requested,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Group Group

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup() *Group {
	this := Group{}
	return &this
}

// NewGroupWithDefaults instantiates a new Group object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupWithDefaults() *Group {
	this := Group{}
	return &this
}

// GetAllowMembersToPostImageAndVideo returns the AllowMembersToPostImageAndVideo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetAllowMembersToPostImageAndVideo() bool {
	if o == nil || IsNil(o.AllowMembersToPostImageAndVideo.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowMembersToPostImageAndVideo.Get()
}

// GetAllowMembersToPostImageAndVideoOk returns a tuple with the AllowMembersToPostImageAndVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetAllowMembersToPostImageAndVideoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowMembersToPostImageAndVideo.Get(), o.AllowMembersToPostImageAndVideo.IsSet()
}

// HasAllowMembersToPostImageAndVideo returns a boolean if a field has been set.
func (o *Group) HasAllowMembersToPostImageAndVideo() bool {
	if o != nil && o.AllowMembersToPostImageAndVideo.IsSet() {
		return true
	}

	return false
}

// SetAllowMembersToPostImageAndVideo gets a reference to the given NullableBool and assigns it to the AllowMembersToPostImageAndVideo field.
func (o *Group) SetAllowMembersToPostImageAndVideo(v bool) {
	o.AllowMembersToPostImageAndVideo.Set(&v)
}
// SetAllowMembersToPostImageAndVideoNil sets the value for AllowMembersToPostImageAndVideo to be an explicit nil
func (o *Group) SetAllowMembersToPostImageAndVideoNil() {
	o.AllowMembersToPostImageAndVideo.Set(nil)
}

// UnsetAllowMembersToPostImageAndVideo ensures that no value is present for AllowMembersToPostImageAndVideo, not even an explicit nil
func (o *Group) UnsetAllowMembersToPostImageAndVideo() {
	o.AllowMembersToPostImageAndVideo.Unset()
}

// GetAllowMembersToPostUrl returns the AllowMembersToPostUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetAllowMembersToPostUrl() bool {
	if o == nil || IsNil(o.AllowMembersToPostUrl.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowMembersToPostUrl.Get()
}

// GetAllowMembersToPostUrlOk returns a tuple with the AllowMembersToPostUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetAllowMembersToPostUrlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowMembersToPostUrl.Get(), o.AllowMembersToPostUrl.IsSet()
}

// HasAllowMembersToPostUrl returns a boolean if a field has been set.
func (o *Group) HasAllowMembersToPostUrl() bool {
	if o != nil && o.AllowMembersToPostUrl.IsSet() {
		return true
	}

	return false
}

// SetAllowMembersToPostUrl gets a reference to the given NullableBool and assigns it to the AllowMembersToPostUrl field.
func (o *Group) SetAllowMembersToPostUrl(v bool) {
	o.AllowMembersToPostUrl.Set(&v)
}
// SetAllowMembersToPostUrlNil sets the value for AllowMembersToPostUrl to be an explicit nil
func (o *Group) SetAllowMembersToPostUrlNil() {
	o.AllowMembersToPostUrl.Set(nil)
}

// UnsetAllowMembersToPostUrl ensures that no value is present for AllowMembersToPostUrl, not even an explicit nil
func (o *Group) UnsetAllowMembersToPostUrl() {
	o.AllowMembersToPostUrl.Unset()
}

// GetAllowOwnershipTransfer returns the AllowOwnershipTransfer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetAllowOwnershipTransfer() bool {
	if o == nil || IsNil(o.AllowOwnershipTransfer.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowOwnershipTransfer.Get()
}

// GetAllowOwnershipTransferOk returns a tuple with the AllowOwnershipTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetAllowOwnershipTransferOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowOwnershipTransfer.Get(), o.AllowOwnershipTransfer.IsSet()
}

// HasAllowOwnershipTransfer returns a boolean if a field has been set.
func (o *Group) HasAllowOwnershipTransfer() bool {
	if o != nil && o.AllowOwnershipTransfer.IsSet() {
		return true
	}

	return false
}

// SetAllowOwnershipTransfer gets a reference to the given NullableBool and assigns it to the AllowOwnershipTransfer field.
func (o *Group) SetAllowOwnershipTransfer(v bool) {
	o.AllowOwnershipTransfer.Set(&v)
}
// SetAllowOwnershipTransferNil sets the value for AllowOwnershipTransfer to be an explicit nil
func (o *Group) SetAllowOwnershipTransferNil() {
	o.AllowOwnershipTransfer.Set(nil)
}

// UnsetAllowOwnershipTransfer ensures that no value is present for AllowOwnershipTransfer, not even an explicit nil
func (o *Group) UnsetAllowOwnershipTransfer() {
	o.AllowOwnershipTransfer.Unset()
}

// GetAllowThreadCreationBy returns the AllowThreadCreationBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetAllowThreadCreationBy() string {
	if o == nil || IsNil(o.AllowThreadCreationBy.Get()) {
		var ret string
		return ret
	}
	return *o.AllowThreadCreationBy.Get()
}

// GetAllowThreadCreationByOk returns a tuple with the AllowThreadCreationBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetAllowThreadCreationByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowThreadCreationBy.Get(), o.AllowThreadCreationBy.IsSet()
}

// HasAllowThreadCreationBy returns a boolean if a field has been set.
func (o *Group) HasAllowThreadCreationBy() bool {
	if o != nil && o.AllowThreadCreationBy.IsSet() {
		return true
	}

	return false
}

// SetAllowThreadCreationBy gets a reference to the given NullableString and assigns it to the AllowThreadCreationBy field.
func (o *Group) SetAllowThreadCreationBy(v string) {
	o.AllowThreadCreationBy.Set(&v)
}
// SetAllowThreadCreationByNil sets the value for AllowThreadCreationBy to be an explicit nil
func (o *Group) SetAllowThreadCreationByNil() {
	o.AllowThreadCreationBy.Set(nil)
}

// UnsetAllowThreadCreationBy ensures that no value is present for AllowThreadCreationBy, not even an explicit nil
func (o *Group) UnsetAllowThreadCreationBy() {
	o.AllowThreadCreationBy.Unset()
}

// GetCallTimelineDisplay returns the CallTimelineDisplay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetCallTimelineDisplay() bool {
	if o == nil || IsNil(o.CallTimelineDisplay.Get()) {
		var ret bool
		return ret
	}
	return *o.CallTimelineDisplay.Get()
}

// GetCallTimelineDisplayOk returns a tuple with the CallTimelineDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetCallTimelineDisplayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.CallTimelineDisplay.Get(), o.CallTimelineDisplay.IsSet()
}

// HasCallTimelineDisplay returns a boolean if a field has been set.
func (o *Group) HasCallTimelineDisplay() bool {
	if o != nil && o.CallTimelineDisplay.IsSet() {
		return true
	}

	return false
}

// SetCallTimelineDisplay gets a reference to the given NullableBool and assigns it to the CallTimelineDisplay field.
func (o *Group) SetCallTimelineDisplay(v bool) {
	o.CallTimelineDisplay.Set(&v)
}
// SetCallTimelineDisplayNil sets the value for CallTimelineDisplay to be an explicit nil
func (o *Group) SetCallTimelineDisplayNil() {
	o.CallTimelineDisplay.Set(nil)
}

// UnsetCallTimelineDisplay ensures that no value is present for CallTimelineDisplay, not even an explicit nil
func (o *Group) UnsetCallTimelineDisplay() {
	o.CallTimelineDisplay.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetCoverImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *Group) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *Group) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}
// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *Group) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *Group) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetCoverImageThumbnail returns the CoverImageThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetCoverImageThumbnail() string {
	if o == nil || IsNil(o.CoverImageThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImageThumbnail.Get()
}

// GetCoverImageThumbnailOk returns a tuple with the CoverImageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetCoverImageThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImageThumbnail.Get(), o.CoverImageThumbnail.IsSet()
}

// HasCoverImageThumbnail returns a boolean if a field has been set.
func (o *Group) HasCoverImageThumbnail() bool {
	if o != nil && o.CoverImageThumbnail.IsSet() {
		return true
	}

	return false
}

// SetCoverImageThumbnail gets a reference to the given NullableString and assigns it to the CoverImageThumbnail field.
func (o *Group) SetCoverImageThumbnail(v string) {
	o.CoverImageThumbnail.Set(&v)
}
// SetCoverImageThumbnailNil sets the value for CoverImageThumbnail to be an explicit nil
func (o *Group) SetCoverImageThumbnailNil() {
	o.CoverImageThumbnail.Set(nil)
}

// UnsetCoverImageThumbnail ensures that no value is present for CoverImageThumbnail, not even an explicit nil
func (o *Group) UnsetCoverImageThumbnail() {
	o.CoverImageThumbnail.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Group) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Group) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Group) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Group) UnsetDescription() {
	o.Description.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGender() int32 {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret int32
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGenderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *Group) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableInt32 and assigns it to the Gender field.
func (o *Group) SetGender(v int32) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *Group) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *Group) UnsetGender() {
	o.Gender.Unset()
}

// GetGenerationGroupsLimit returns the GenerationGroupsLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGenerationGroupsLimit() int32 {
	if o == nil || IsNil(o.GenerationGroupsLimit.Get()) {
		var ret int32
		return ret
	}
	return *o.GenerationGroupsLimit.Get()
}

// GetGenerationGroupsLimitOk returns a tuple with the GenerationGroupsLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGenerationGroupsLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.GenerationGroupsLimit.Get(), o.GenerationGroupsLimit.IsSet()
}

// HasGenerationGroupsLimit returns a boolean if a field has been set.
func (o *Group) HasGenerationGroupsLimit() bool {
	if o != nil && o.GenerationGroupsLimit.IsSet() {
		return true
	}

	return false
}

// SetGenerationGroupsLimit gets a reference to the given NullableInt32 and assigns it to the GenerationGroupsLimit field.
func (o *Group) SetGenerationGroupsLimit(v int32) {
	o.GenerationGroupsLimit.Set(&v)
}
// SetGenerationGroupsLimitNil sets the value for GenerationGroupsLimit to be an explicit nil
func (o *Group) SetGenerationGroupsLimitNil() {
	o.GenerationGroupsLimit.Set(nil)
}

// UnsetGenerationGroupsLimit ensures that no value is present for GenerationGroupsLimit, not even an explicit nil
func (o *Group) UnsetGenerationGroupsLimit() {
	o.GenerationGroupsLimit.Unset()
}

// GetGroupCategoryId returns the GroupCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGroupCategoryId() int64 {
	if o == nil || IsNil(o.GroupCategoryId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupCategoryId.Get()
}

// GetGroupCategoryIdOk returns a tuple with the GroupCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGroupCategoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupCategoryId.Get(), o.GroupCategoryId.IsSet()
}

// HasGroupCategoryId returns a boolean if a field has been set.
func (o *Group) HasGroupCategoryId() bool {
	if o != nil && o.GroupCategoryId.IsSet() {
		return true
	}

	return false
}

// SetGroupCategoryId gets a reference to the given NullableInt64 and assigns it to the GroupCategoryId field.
func (o *Group) SetGroupCategoryId(v int64) {
	o.GroupCategoryId.Set(&v)
}
// SetGroupCategoryIdNil sets the value for GroupCategoryId to be an explicit nil
func (o *Group) SetGroupCategoryIdNil() {
	o.GroupCategoryId.Set(nil)
}

// UnsetGroupCategoryId ensures that no value is present for GroupCategoryId, not even an explicit nil
func (o *Group) UnsetGroupCategoryId() {
	o.GroupCategoryId.Unset()
}

// GetGroupIcon returns the GroupIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGroupIcon() string {
	if o == nil || IsNil(o.GroupIcon.Get()) {
		var ret string
		return ret
	}
	return *o.GroupIcon.Get()
}

// GetGroupIconOk returns a tuple with the GroupIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGroupIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIcon.Get(), o.GroupIcon.IsSet()
}

// HasGroupIcon returns a boolean if a field has been set.
func (o *Group) HasGroupIcon() bool {
	if o != nil && o.GroupIcon.IsSet() {
		return true
	}

	return false
}

// SetGroupIcon gets a reference to the given NullableString and assigns it to the GroupIcon field.
func (o *Group) SetGroupIcon(v string) {
	o.GroupIcon.Set(&v)
}
// SetGroupIconNil sets the value for GroupIcon to be an explicit nil
func (o *Group) SetGroupIconNil() {
	o.GroupIcon.Set(nil)
}

// UnsetGroupIcon ensures that no value is present for GroupIcon, not even an explicit nil
func (o *Group) UnsetGroupIcon() {
	o.GroupIcon.Unset()
}

// GetGroupIconThumbnail returns the GroupIconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGroupIconThumbnail() string {
	if o == nil || IsNil(o.GroupIconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.GroupIconThumbnail.Get()
}

// GetGroupIconThumbnailOk returns a tuple with the GroupIconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGroupIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIconThumbnail.Get(), o.GroupIconThumbnail.IsSet()
}

// HasGroupIconThumbnail returns a boolean if a field has been set.
func (o *Group) HasGroupIconThumbnail() bool {
	if o != nil && o.GroupIconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetGroupIconThumbnail gets a reference to the given NullableString and assigns it to the GroupIconThumbnail field.
func (o *Group) SetGroupIconThumbnail(v string) {
	o.GroupIconThumbnail.Set(&v)
}
// SetGroupIconThumbnailNil sets the value for GroupIconThumbnail to be an explicit nil
func (o *Group) SetGroupIconThumbnailNil() {
	o.GroupIconThumbnail.Set(nil)
}

// UnsetGroupIconThumbnail ensures that no value is present for GroupIconThumbnail, not even an explicit nil
func (o *Group) UnsetGroupIconThumbnail() {
	o.GroupIconThumbnail.Unset()
}

// GetGroupsUsersCount returns the GroupsUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGroupsUsersCount() int64 {
	if o == nil || IsNil(o.GroupsUsersCount.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupsUsersCount.Get()
}

// GetGroupsUsersCountOk returns a tuple with the GroupsUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGroupsUsersCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsUsersCount.Get(), o.GroupsUsersCount.IsSet()
}

// HasGroupsUsersCount returns a boolean if a field has been set.
func (o *Group) HasGroupsUsersCount() bool {
	if o != nil && o.GroupsUsersCount.IsSet() {
		return true
	}

	return false
}

// SetGroupsUsersCount gets a reference to the given NullableInt64 and assigns it to the GroupsUsersCount field.
func (o *Group) SetGroupsUsersCount(v int64) {
	o.GroupsUsersCount.Set(&v)
}
// SetGroupsUsersCountNil sets the value for GroupsUsersCount to be an explicit nil
func (o *Group) SetGroupsUsersCountNil() {
	o.GroupsUsersCount.Set(nil)
}

// UnsetGroupsUsersCount ensures that no value is present for GroupsUsersCount, not even an explicit nil
func (o *Group) UnsetGroupsUsersCount() {
	o.GroupsUsersCount.Unset()
}

// GetGuidelines returns the Guidelines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetGuidelines() string {
	if o == nil || IsNil(o.Guidelines.Get()) {
		var ret string
		return ret
	}
	return *o.Guidelines.Get()
}

// GetGuidelinesOk returns a tuple with the Guidelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetGuidelinesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guidelines.Get(), o.Guidelines.IsSet()
}

// HasGuidelines returns a boolean if a field has been set.
func (o *Group) HasGuidelines() bool {
	if o != nil && o.Guidelines.IsSet() {
		return true
	}

	return false
}

// SetGuidelines gets a reference to the given NullableString and assigns it to the Guidelines field.
func (o *Group) SetGuidelines(v string) {
	o.Guidelines.Set(&v)
}
// SetGuidelinesNil sets the value for Guidelines to be an explicit nil
func (o *Group) SetGuidelinesNil() {
	o.Guidelines.Set(nil)
}

// UnsetGuidelines ensures that no value is present for Guidelines, not even an explicit nil
func (o *Group) UnsetGuidelines() {
	o.Guidelines.Unset()
}

// GetHideConferenceCall returns the HideConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetHideConferenceCall() bool {
	if o == nil || IsNil(o.HideConferenceCall.Get()) {
		var ret bool
		return ret
	}
	return *o.HideConferenceCall.Get()
}

// GetHideConferenceCallOk returns a tuple with the HideConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetHideConferenceCallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideConferenceCall.Get(), o.HideConferenceCall.IsSet()
}

// HasHideConferenceCall returns a boolean if a field has been set.
func (o *Group) HasHideConferenceCall() bool {
	if o != nil && o.HideConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetHideConferenceCall gets a reference to the given NullableBool and assigns it to the HideConferenceCall field.
func (o *Group) SetHideConferenceCall(v bool) {
	o.HideConferenceCall.Set(&v)
}
// SetHideConferenceCallNil sets the value for HideConferenceCall to be an explicit nil
func (o *Group) SetHideConferenceCallNil() {
	o.HideConferenceCall.Set(nil)
}

// UnsetHideConferenceCall ensures that no value is present for HideConferenceCall, not even an explicit nil
func (o *Group) UnsetHideConferenceCall() {
	o.HideConferenceCall.Unset()
}

// GetHideFromGameEight returns the HideFromGameEight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetHideFromGameEight() bool {
	if o == nil || IsNil(o.HideFromGameEight.Get()) {
		var ret bool
		return ret
	}
	return *o.HideFromGameEight.Get()
}

// GetHideFromGameEightOk returns a tuple with the HideFromGameEight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetHideFromGameEightOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideFromGameEight.Get(), o.HideFromGameEight.IsSet()
}

// HasHideFromGameEight returns a boolean if a field has been set.
func (o *Group) HasHideFromGameEight() bool {
	if o != nil && o.HideFromGameEight.IsSet() {
		return true
	}

	return false
}

// SetHideFromGameEight gets a reference to the given NullableBool and assigns it to the HideFromGameEight field.
func (o *Group) SetHideFromGameEight(v bool) {
	o.HideFromGameEight.Set(&v)
}
// SetHideFromGameEightNil sets the value for HideFromGameEight to be an explicit nil
func (o *Group) SetHideFromGameEightNil() {
	o.HideFromGameEight.Set(nil)
}

// UnsetHideFromGameEight ensures that no value is present for HideFromGameEight, not even an explicit nil
func (o *Group) UnsetHideFromGameEight() {
	o.HideFromGameEight.Unset()
}

// GetHideReportedPosts returns the HideReportedPosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetHideReportedPosts() bool {
	if o == nil || IsNil(o.HideReportedPosts.Get()) {
		var ret bool
		return ret
	}
	return *o.HideReportedPosts.Get()
}

// GetHideReportedPostsOk returns a tuple with the HideReportedPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetHideReportedPostsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HideReportedPosts.Get(), o.HideReportedPosts.IsSet()
}

// HasHideReportedPosts returns a boolean if a field has been set.
func (o *Group) HasHideReportedPosts() bool {
	if o != nil && o.HideReportedPosts.IsSet() {
		return true
	}

	return false
}

// SetHideReportedPosts gets a reference to the given NullableBool and assigns it to the HideReportedPosts field.
func (o *Group) SetHideReportedPosts(v bool) {
	o.HideReportedPosts.Set(&v)
}
// SetHideReportedPostsNil sets the value for HideReportedPosts to be an explicit nil
func (o *Group) SetHideReportedPostsNil() {
	o.HideReportedPosts.Set(nil)
}

// UnsetHideReportedPosts ensures that no value is present for HideReportedPosts, not even an explicit nil
func (o *Group) UnsetHideReportedPosts() {
	o.HideReportedPosts.Unset()
}

// GetHighlightedCount returns the HighlightedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetHighlightedCount() int64 {
	if o == nil || IsNil(o.HighlightedCount.Get()) {
		var ret int64
		return ret
	}
	return *o.HighlightedCount.Get()
}

// GetHighlightedCountOk returns a tuple with the HighlightedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetHighlightedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighlightedCount.Get(), o.HighlightedCount.IsSet()
}

// HasHighlightedCount returns a boolean if a field has been set.
func (o *Group) HasHighlightedCount() bool {
	if o != nil && o.HighlightedCount.IsSet() {
		return true
	}

	return false
}

// SetHighlightedCount gets a reference to the given NullableInt64 and assigns it to the HighlightedCount field.
func (o *Group) SetHighlightedCount(v int64) {
	o.HighlightedCount.Set(&v)
}
// SetHighlightedCountNil sets the value for HighlightedCount to be an explicit nil
func (o *Group) SetHighlightedCountNil() {
	o.HighlightedCount.Set(nil)
}

// UnsetHighlightedCount ensures that no value is present for HighlightedCount, not even an explicit nil
func (o *Group) UnsetHighlightedCount() {
	o.HighlightedCount.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Group) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Group) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Group) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Group) UnsetId() {
	o.Id.Unset()
}

// GetInvitedToJoin returns the InvitedToJoin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetInvitedToJoin() bool {
	if o == nil || IsNil(o.InvitedToJoin.Get()) {
		var ret bool
		return ret
	}
	return *o.InvitedToJoin.Get()
}

// GetInvitedToJoinOk returns a tuple with the InvitedToJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetInvitedToJoinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvitedToJoin.Get(), o.InvitedToJoin.IsSet()
}

// HasInvitedToJoin returns a boolean if a field has been set.
func (o *Group) HasInvitedToJoin() bool {
	if o != nil && o.InvitedToJoin.IsSet() {
		return true
	}

	return false
}

// SetInvitedToJoin gets a reference to the given NullableBool and assigns it to the InvitedToJoin field.
func (o *Group) SetInvitedToJoin(v bool) {
	o.InvitedToJoin.Set(&v)
}
// SetInvitedToJoinNil sets the value for InvitedToJoin to be an explicit nil
func (o *Group) SetInvitedToJoinNil() {
	o.InvitedToJoin.Set(nil)
}

// UnsetInvitedToJoin ensures that no value is present for InvitedToJoin, not even an explicit nil
func (o *Group) UnsetInvitedToJoin() {
	o.InvitedToJoin.Unset()
}

// GetIsJoined returns the IsJoined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetIsJoined() bool {
	if o == nil || IsNil(o.IsJoined.Get()) {
		var ret bool
		return ret
	}
	return *o.IsJoined.Get()
}

// GetIsJoinedOk returns a tuple with the IsJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIsJoinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsJoined.Get(), o.IsJoined.IsSet()
}

// HasIsJoined returns a boolean if a field has been set.
func (o *Group) HasIsJoined() bool {
	if o != nil && o.IsJoined.IsSet() {
		return true
	}

	return false
}

// SetIsJoined gets a reference to the given NullableBool and assigns it to the IsJoined field.
func (o *Group) SetIsJoined(v bool) {
	o.IsJoined.Set(&v)
}
// SetIsJoinedNil sets the value for IsJoined to be an explicit nil
func (o *Group) SetIsJoinedNil() {
	o.IsJoined.Set(nil)
}

// UnsetIsJoined ensures that no value is present for IsJoined, not even an explicit nil
func (o *Group) UnsetIsJoined() {
	o.IsJoined.Unset()
}

// GetIsPending returns the IsPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetIsPending() bool {
	if o == nil || IsNil(o.IsPending.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPending.Get()
}

// GetIsPendingOk returns a tuple with the IsPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIsPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPending.Get(), o.IsPending.IsSet()
}

// HasIsPending returns a boolean if a field has been set.
func (o *Group) HasIsPending() bool {
	if o != nil && o.IsPending.IsSet() {
		return true
	}

	return false
}

// SetIsPending gets a reference to the given NullableBool and assigns it to the IsPending field.
func (o *Group) SetIsPending(v bool) {
	o.IsPending.Set(&v)
}
// SetIsPendingNil sets the value for IsPending to be an explicit nil
func (o *Group) SetIsPendingNil() {
	o.IsPending.Set(nil)
}

// UnsetIsPending ensures that no value is present for IsPending, not even an explicit nil
func (o *Group) UnsetIsPending() {
	o.IsPending.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *Group) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *Group) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *Group) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *Group) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetIsRelated returns the IsRelated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetIsRelated() bool {
	if o == nil || IsNil(o.IsRelated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRelated.Get()
}

// GetIsRelatedOk returns a tuple with the IsRelated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetIsRelatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRelated.Get(), o.IsRelated.IsSet()
}

// HasIsRelated returns a boolean if a field has been set.
func (o *Group) HasIsRelated() bool {
	if o != nil && o.IsRelated.IsSet() {
		return true
	}

	return false
}

// SetIsRelated gets a reference to the given NullableBool and assigns it to the IsRelated field.
func (o *Group) SetIsRelated(v bool) {
	o.IsRelated.Set(&v)
}
// SetIsRelatedNil sets the value for IsRelated to be an explicit nil
func (o *Group) SetIsRelatedNil() {
	o.IsRelated.Set(nil)
}

// UnsetIsRelated ensures that no value is present for IsRelated, not even an explicit nil
func (o *Group) UnsetIsRelated() {
	o.IsRelated.Unset()
}

// GetJoinedCommunityCampaign returns the JoinedCommunityCampaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetJoinedCommunityCampaign() bool {
	if o == nil || IsNil(o.JoinedCommunityCampaign.Get()) {
		var ret bool
		return ret
	}
	return *o.JoinedCommunityCampaign.Get()
}

// GetJoinedCommunityCampaignOk returns a tuple with the JoinedCommunityCampaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetJoinedCommunityCampaignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinedCommunityCampaign.Get(), o.JoinedCommunityCampaign.IsSet()
}

// HasJoinedCommunityCampaign returns a boolean if a field has been set.
func (o *Group) HasJoinedCommunityCampaign() bool {
	if o != nil && o.JoinedCommunityCampaign.IsSet() {
		return true
	}

	return false
}

// SetJoinedCommunityCampaign gets a reference to the given NullableBool and assigns it to the JoinedCommunityCampaign field.
func (o *Group) SetJoinedCommunityCampaign(v bool) {
	o.JoinedCommunityCampaign.Set(&v)
}
// SetJoinedCommunityCampaignNil sets the value for JoinedCommunityCampaign to be an explicit nil
func (o *Group) SetJoinedCommunityCampaignNil() {
	o.JoinedCommunityCampaign.Set(nil)
}

// UnsetJoinedCommunityCampaign ensures that no value is present for JoinedCommunityCampaign, not even an explicit nil
func (o *Group) UnsetJoinedCommunityCampaign() {
	o.JoinedCommunityCampaign.Unset()
}

// GetModeratorIds returns the ModeratorIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetModeratorIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ModeratorIds
}

// GetModeratorIdsOk returns a tuple with the ModeratorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetModeratorIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ModeratorIds) {
		return nil, false
	}
	return o.ModeratorIds, true
}

// HasModeratorIds returns a boolean if a field has been set.
func (o *Group) HasModeratorIds() bool {
	if o != nil && !IsNil(o.ModeratorIds) {
		return true
	}

	return false
}

// SetModeratorIds gets a reference to the given []int64 and assigns it to the ModeratorIds field.
func (o *Group) SetModeratorIds(v []int64) {
	o.ModeratorIds = v
}

// GetOnlyMobileVerified returns the OnlyMobileVerified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetOnlyMobileVerified() bool {
	if o == nil || IsNil(o.OnlyMobileVerified.Get()) {
		var ret bool
		return ret
	}
	return *o.OnlyMobileVerified.Get()
}

// GetOnlyMobileVerifiedOk returns a tuple with the OnlyMobileVerified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetOnlyMobileVerifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlyMobileVerified.Get(), o.OnlyMobileVerified.IsSet()
}

// HasOnlyMobileVerified returns a boolean if a field has been set.
func (o *Group) HasOnlyMobileVerified() bool {
	if o != nil && o.OnlyMobileVerified.IsSet() {
		return true
	}

	return false
}

// SetOnlyMobileVerified gets a reference to the given NullableBool and assigns it to the OnlyMobileVerified field.
func (o *Group) SetOnlyMobileVerified(v bool) {
	o.OnlyMobileVerified.Set(&v)
}
// SetOnlyMobileVerifiedNil sets the value for OnlyMobileVerified to be an explicit nil
func (o *Group) SetOnlyMobileVerifiedNil() {
	o.OnlyMobileVerified.Set(nil)
}

// UnsetOnlyMobileVerified ensures that no value is present for OnlyMobileVerified, not even an explicit nil
func (o *Group) UnsetOnlyMobileVerified() {
	o.OnlyMobileVerified.Unset()
}

// GetOnlyVerifiedAge returns the OnlyVerifiedAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetOnlyVerifiedAge() bool {
	if o == nil || IsNil(o.OnlyVerifiedAge.Get()) {
		var ret bool
		return ret
	}
	return *o.OnlyVerifiedAge.Get()
}

// GetOnlyVerifiedAgeOk returns a tuple with the OnlyVerifiedAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetOnlyVerifiedAgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlyVerifiedAge.Get(), o.OnlyVerifiedAge.IsSet()
}

// HasOnlyVerifiedAge returns a boolean if a field has been set.
func (o *Group) HasOnlyVerifiedAge() bool {
	if o != nil && o.OnlyVerifiedAge.IsSet() {
		return true
	}

	return false
}

// SetOnlyVerifiedAge gets a reference to the given NullableBool and assigns it to the OnlyVerifiedAge field.
func (o *Group) SetOnlyVerifiedAge(v bool) {
	o.OnlyVerifiedAge.Set(&v)
}
// SetOnlyVerifiedAgeNil sets the value for OnlyVerifiedAge to be an explicit nil
func (o *Group) SetOnlyVerifiedAgeNil() {
	o.OnlyVerifiedAge.Set(nil)
}

// UnsetOnlyVerifiedAge ensures that no value is present for OnlyVerifiedAge, not even an explicit nil
func (o *Group) UnsetOnlyVerifiedAge() {
	o.OnlyVerifiedAge.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetOwner() RealmUser {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetOwnerOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *Group) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableRealmUser and assigns it to the Owner field.
func (o *Group) SetOwner(v RealmUser) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *Group) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *Group) UnsetOwner() {
	o.Owner.Unset()
}

// GetPendingCount returns the PendingCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetPendingCount() int64 {
	if o == nil || IsNil(o.PendingCount.Get()) {
		var ret int64
		return ret
	}
	return *o.PendingCount.Get()
}

// GetPendingCountOk returns a tuple with the PendingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetPendingCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingCount.Get(), o.PendingCount.IsSet()
}

// HasPendingCount returns a boolean if a field has been set.
func (o *Group) HasPendingCount() bool {
	if o != nil && o.PendingCount.IsSet() {
		return true
	}

	return false
}

// SetPendingCount gets a reference to the given NullableInt64 and assigns it to the PendingCount field.
func (o *Group) SetPendingCount(v int64) {
	o.PendingCount.Set(&v)
}
// SetPendingCountNil sets the value for PendingCount to be an explicit nil
func (o *Group) SetPendingCountNil() {
	o.PendingCount.Set(nil)
}

// UnsetPendingCount ensures that no value is present for PendingCount, not even an explicit nil
func (o *Group) UnsetPendingCount() {
	o.PendingCount.Unset()
}

// GetPendingDeputizeIds returns the PendingDeputizeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetPendingDeputizeIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.PendingDeputizeIds
}

// GetPendingDeputizeIdsOk returns a tuple with the PendingDeputizeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetPendingDeputizeIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PendingDeputizeIds) {
		return nil, false
	}
	return o.PendingDeputizeIds, true
}

// HasPendingDeputizeIds returns a boolean if a field has been set.
func (o *Group) HasPendingDeputizeIds() bool {
	if o != nil && !IsNil(o.PendingDeputizeIds) {
		return true
	}

	return false
}

// SetPendingDeputizeIds gets a reference to the given []int64 and assigns it to the PendingDeputizeIds field.
func (o *Group) SetPendingDeputizeIds(v []int64) {
	o.PendingDeputizeIds = v
}

// GetPendingTransferId returns the PendingTransferId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetPendingTransferId() int64 {
	if o == nil || IsNil(o.PendingTransferId.Get()) {
		var ret int64
		return ret
	}
	return *o.PendingTransferId.Get()
}

// GetPendingTransferIdOk returns a tuple with the PendingTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetPendingTransferIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingTransferId.Get(), o.PendingTransferId.IsSet()
}

// HasPendingTransferId returns a boolean if a field has been set.
func (o *Group) HasPendingTransferId() bool {
	if o != nil && o.PendingTransferId.IsSet() {
		return true
	}

	return false
}

// SetPendingTransferId gets a reference to the given NullableInt64 and assigns it to the PendingTransferId field.
func (o *Group) SetPendingTransferId(v int64) {
	o.PendingTransferId.Set(&v)
}
// SetPendingTransferIdNil sets the value for PendingTransferId to be an explicit nil
func (o *Group) SetPendingTransferIdNil() {
	o.PendingTransferId.Set(nil)
}

// UnsetPendingTransferId ensures that no value is present for PendingTransferId, not even an explicit nil
func (o *Group) UnsetPendingTransferId() {
	o.PendingTransferId.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetPostsCount() int64 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetPostsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *Group) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt64 and assigns it to the PostsCount field.
func (o *Group) SetPostsCount(v int64) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *Group) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *Group) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetRelatedCount returns the RelatedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetRelatedCount() int64 {
	if o == nil || IsNil(o.RelatedCount.Get()) {
		var ret int64
		return ret
	}
	return *o.RelatedCount.Get()
}

// GetRelatedCountOk returns a tuple with the RelatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetRelatedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelatedCount.Get(), o.RelatedCount.IsSet()
}

// HasRelatedCount returns a boolean if a field has been set.
func (o *Group) HasRelatedCount() bool {
	if o != nil && o.RelatedCount.IsSet() {
		return true
	}

	return false
}

// SetRelatedCount gets a reference to the given NullableInt64 and assigns it to the RelatedCount field.
func (o *Group) SetRelatedCount(v int64) {
	o.RelatedCount.Set(&v)
}
// SetRelatedCountNil sets the value for RelatedCount to be an explicit nil
func (o *Group) SetRelatedCountNil() {
	o.RelatedCount.Set(nil)
}

// UnsetRelatedCount ensures that no value is present for RelatedCount, not even an explicit nil
func (o *Group) UnsetRelatedCount() {
	o.RelatedCount.Unset()
}

// GetSafeMode returns the SafeMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetSafeMode() bool {
	if o == nil || IsNil(o.SafeMode.Get()) {
		var ret bool
		return ret
	}
	return *o.SafeMode.Get()
}

// GetSafeModeOk returns a tuple with the SafeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetSafeModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.SafeMode.Get(), o.SafeMode.IsSet()
}

// HasSafeMode returns a boolean if a field has been set.
func (o *Group) HasSafeMode() bool {
	if o != nil && o.SafeMode.IsSet() {
		return true
	}

	return false
}

// SetSafeMode gets a reference to the given NullableBool and assigns it to the SafeMode field.
func (o *Group) SetSafeMode(v bool) {
	o.SafeMode.Set(&v)
}
// SetSafeModeNil sets the value for SafeMode to be an explicit nil
func (o *Group) SetSafeModeNil() {
	o.SafeMode.Set(nil)
}

// UnsetSafeMode ensures that no value is present for SafeMode, not even an explicit nil
func (o *Group) UnsetSafeMode() {
	o.SafeMode.Unset()
}

// GetSecret returns the Secret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetSecret() bool {
	if o == nil || IsNil(o.Secret.Get()) {
		var ret bool
		return ret
	}
	return *o.Secret.Get()
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Secret.Get(), o.Secret.IsSet()
}

// HasSecret returns a boolean if a field has been set.
func (o *Group) HasSecret() bool {
	if o != nil && o.Secret.IsSet() {
		return true
	}

	return false
}

// SetSecret gets a reference to the given NullableBool and assigns it to the Secret field.
func (o *Group) SetSecret(v bool) {
	o.Secret.Set(&v)
}
// SetSecretNil sets the value for Secret to be an explicit nil
func (o *Group) SetSecretNil() {
	o.Secret.Set(nil)
}

// UnsetSecret ensures that no value is present for Secret, not even an explicit nil
func (o *Group) UnsetSecret() {
	o.Secret.Unset()
}

// GetSeizable returns the Seizable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetSeizable() bool {
	if o == nil || IsNil(o.Seizable.Get()) {
		var ret bool
		return ret
	}
	return *o.Seizable.Get()
}

// GetSeizableOk returns a tuple with the Seizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetSeizableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Seizable.Get(), o.Seizable.IsSet()
}

// HasSeizable returns a boolean if a field has been set.
func (o *Group) HasSeizable() bool {
	if o != nil && o.Seizable.IsSet() {
		return true
	}

	return false
}

// SetSeizable gets a reference to the given NullableBool and assigns it to the Seizable field.
func (o *Group) SetSeizable(v bool) {
	o.Seizable.Set(&v)
}
// SetSeizableNil sets the value for Seizable to be an explicit nil
func (o *Group) SetSeizableNil() {
	o.Seizable.Set(nil)
}

// UnsetSeizable ensures that no value is present for Seizable, not even an explicit nil
func (o *Group) UnsetSeizable() {
	o.Seizable.Unset()
}

// GetSeizableBefore returns the SeizableBefore field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetSeizableBefore() int64 {
	if o == nil || IsNil(o.SeizableBefore.Get()) {
		var ret int64
		return ret
	}
	return *o.SeizableBefore.Get()
}

// GetSeizableBeforeOk returns a tuple with the SeizableBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetSeizableBeforeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeizableBefore.Get(), o.SeizableBefore.IsSet()
}

// HasSeizableBefore returns a boolean if a field has been set.
func (o *Group) HasSeizableBefore() bool {
	if o != nil && o.SeizableBefore.IsSet() {
		return true
	}

	return false
}

// SetSeizableBefore gets a reference to the given NullableInt64 and assigns it to the SeizableBefore field.
func (o *Group) SetSeizableBefore(v int64) {
	o.SeizableBefore.Set(&v)
}
// SetSeizableBeforeNil sets the value for SeizableBefore to be an explicit nil
func (o *Group) SetSeizableBeforeNil() {
	o.SeizableBefore.Set(nil)
}

// UnsetSeizableBefore ensures that no value is present for SeizableBefore, not even an explicit nil
func (o *Group) UnsetSeizableBefore() {
	o.SeizableBefore.Unset()
}

// GetSubCategoryId returns the SubCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetSubCategoryId() int64 {
	if o == nil || IsNil(o.SubCategoryId.Get()) {
		var ret int64
		return ret
	}
	return *o.SubCategoryId.Get()
}

// GetSubCategoryIdOk returns a tuple with the SubCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetSubCategoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubCategoryId.Get(), o.SubCategoryId.IsSet()
}

// HasSubCategoryId returns a boolean if a field has been set.
func (o *Group) HasSubCategoryId() bool {
	if o != nil && o.SubCategoryId.IsSet() {
		return true
	}

	return false
}

// SetSubCategoryId gets a reference to the given NullableInt64 and assigns it to the SubCategoryId field.
func (o *Group) SetSubCategoryId(v int64) {
	o.SubCategoryId.Set(&v)
}
// SetSubCategoryIdNil sets the value for SubCategoryId to be an explicit nil
func (o *Group) SetSubCategoryIdNil() {
	o.SubCategoryId.Set(nil)
}

// UnsetSubCategoryId ensures that no value is present for SubCategoryId, not even an explicit nil
func (o *Group) UnsetSubCategoryId() {
	o.SubCategoryId.Unset()
}

// GetThreadsCount returns the ThreadsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetThreadsCount() int64 {
	if o == nil || IsNil(o.ThreadsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.ThreadsCount.Get()
}

// GetThreadsCountOk returns a tuple with the ThreadsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetThreadsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadsCount.Get(), o.ThreadsCount.IsSet()
}

// HasThreadsCount returns a boolean if a field has been set.
func (o *Group) HasThreadsCount() bool {
	if o != nil && o.ThreadsCount.IsSet() {
		return true
	}

	return false
}

// SetThreadsCount gets a reference to the given NullableInt64 and assigns it to the ThreadsCount field.
func (o *Group) SetThreadsCount(v int64) {
	o.ThreadsCount.Set(&v)
}
// SetThreadsCountNil sets the value for ThreadsCount to be an explicit nil
func (o *Group) SetThreadsCountNil() {
	o.ThreadsCount.Set(nil)
}

// UnsetThreadsCount ensures that no value is present for ThreadsCount, not even an explicit nil
func (o *Group) UnsetThreadsCount() {
	o.ThreadsCount.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Group) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Group) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Group) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Group) UnsetTitle() {
	o.Title.Unset()
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *Group) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *Group) SetTopic(v string) {
	o.Topic.Set(&v)
}
// SetTopicNil sets the value for Topic to be an explicit nil
func (o *Group) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *Group) UnsetTopic() {
	o.Topic.Unset()
}

// GetUnreadCounts returns the UnreadCounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetUnreadCounts() int64 {
	if o == nil || IsNil(o.UnreadCounts.Get()) {
		var ret int64
		return ret
	}
	return *o.UnreadCounts.Get()
}

// GetUnreadCountsOk returns a tuple with the UnreadCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetUnreadCountsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadCounts.Get(), o.UnreadCounts.IsSet()
}

// HasUnreadCounts returns a boolean if a field has been set.
func (o *Group) HasUnreadCounts() bool {
	if o != nil && o.UnreadCounts.IsSet() {
		return true
	}

	return false
}

// SetUnreadCounts gets a reference to the given NullableInt64 and assigns it to the UnreadCounts field.
func (o *Group) SetUnreadCounts(v int64) {
	o.UnreadCounts.Set(&v)
}
// SetUnreadCountsNil sets the value for UnreadCounts to be an explicit nil
func (o *Group) SetUnreadCountsNil() {
	o.UnreadCounts.Set(nil)
}

// UnsetUnreadCounts ensures that no value is present for UnreadCounts, not even an explicit nil
func (o *Group) UnsetUnreadCounts() {
	o.UnreadCounts.Unset()
}

// GetUnreadThreadsCount returns the UnreadThreadsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetUnreadThreadsCount() int64 {
	if o == nil || IsNil(o.UnreadThreadsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.UnreadThreadsCount.Get()
}

// GetUnreadThreadsCountOk returns a tuple with the UnreadThreadsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetUnreadThreadsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadThreadsCount.Get(), o.UnreadThreadsCount.IsSet()
}

// HasUnreadThreadsCount returns a boolean if a field has been set.
func (o *Group) HasUnreadThreadsCount() bool {
	if o != nil && o.UnreadThreadsCount.IsSet() {
		return true
	}

	return false
}

// SetUnreadThreadsCount gets a reference to the given NullableInt64 and assigns it to the UnreadThreadsCount field.
func (o *Group) SetUnreadThreadsCount(v int64) {
	o.UnreadThreadsCount.Set(&v)
}
// SetUnreadThreadsCountNil sets the value for UnreadThreadsCount to be an explicit nil
func (o *Group) SetUnreadThreadsCountNil() {
	o.UnreadThreadsCount.Set(nil)
}

// UnsetUnreadThreadsCount ensures that no value is present for UnreadThreadsCount, not even an explicit nil
func (o *Group) UnsetUnreadThreadsCount() {
	o.UnreadThreadsCount.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Group) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *Group) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *Group) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *Group) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetUserId() int64 {
	if o == nil || IsNil(o.UserId.Get()) {
		var ret int64
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetUserIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *Group) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableInt64 and assigns it to the UserId field.
func (o *Group) SetUserId(v int64) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *Group) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *Group) UnsetUserId() {
	o.UserId.Unset()
}

// GetViewsCount returns the ViewsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetViewsCount() int64 {
	if o == nil || IsNil(o.ViewsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.ViewsCount.Get()
}

// GetViewsCountOk returns a tuple with the ViewsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetViewsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewsCount.Get(), o.ViewsCount.IsSet()
}

// HasViewsCount returns a boolean if a field has been set.
func (o *Group) HasViewsCount() bool {
	if o != nil && o.ViewsCount.IsSet() {
		return true
	}

	return false
}

// SetViewsCount gets a reference to the given NullableInt64 and assigns it to the ViewsCount field.
func (o *Group) SetViewsCount(v int64) {
	o.ViewsCount.Set(&v)
}
// SetViewsCountNil sets the value for ViewsCount to be an explicit nil
func (o *Group) SetViewsCountNil() {
	o.ViewsCount.Set(nil)
}

// UnsetViewsCount ensures that no value is present for ViewsCount, not even an explicit nil
func (o *Group) UnsetViewsCount() {
	o.ViewsCount.Unset()
}

// GetWalkthroughRequested returns the WalkthroughRequested field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Group) GetWalkthroughRequested() bool {
	if o == nil || IsNil(o.WalkthroughRequested.Get()) {
		var ret bool
		return ret
	}
	return *o.WalkthroughRequested.Get()
}

// GetWalkthroughRequestedOk returns a tuple with the WalkthroughRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Group) GetWalkthroughRequestedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalkthroughRequested.Get(), o.WalkthroughRequested.IsSet()
}

// HasWalkthroughRequested returns a boolean if a field has been set.
func (o *Group) HasWalkthroughRequested() bool {
	if o != nil && o.WalkthroughRequested.IsSet() {
		return true
	}

	return false
}

// SetWalkthroughRequested gets a reference to the given NullableBool and assigns it to the WalkthroughRequested field.
func (o *Group) SetWalkthroughRequested(v bool) {
	o.WalkthroughRequested.Set(&v)
}
// SetWalkthroughRequestedNil sets the value for WalkthroughRequested to be an explicit nil
func (o *Group) SetWalkthroughRequestedNil() {
	o.WalkthroughRequested.Set(nil)
}

// UnsetWalkthroughRequested ensures that no value is present for WalkthroughRequested, not even an explicit nil
func (o *Group) UnsetWalkthroughRequested() {
	o.WalkthroughRequested.Unset()
}

func (o Group) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Group) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowMembersToPostImageAndVideo.IsSet() {
		toSerialize["allow_members_to_post_image_and_video"] = o.AllowMembersToPostImageAndVideo.Get()
	}
	if o.AllowMembersToPostUrl.IsSet() {
		toSerialize["allow_members_to_post_url"] = o.AllowMembersToPostUrl.Get()
	}
	if o.AllowOwnershipTransfer.IsSet() {
		toSerialize["allow_ownership_transfer"] = o.AllowOwnershipTransfer.Get()
	}
	if o.AllowThreadCreationBy.IsSet() {
		toSerialize["allow_thread_creation_by"] = o.AllowThreadCreationBy.Get()
	}
	if o.CallTimelineDisplay.IsSet() {
		toSerialize["call_timeline_display"] = o.CallTimelineDisplay.Get()
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if o.CoverImageThumbnail.IsSet() {
		toSerialize["cover_image_thumbnail"] = o.CoverImageThumbnail.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.GenerationGroupsLimit.IsSet() {
		toSerialize["generation_groups_limit"] = o.GenerationGroupsLimit.Get()
	}
	if o.GroupCategoryId.IsSet() {
		toSerialize["group_category_id"] = o.GroupCategoryId.Get()
	}
	if o.GroupIcon.IsSet() {
		toSerialize["group_icon"] = o.GroupIcon.Get()
	}
	if o.GroupIconThumbnail.IsSet() {
		toSerialize["group_icon_thumbnail"] = o.GroupIconThumbnail.Get()
	}
	if o.GroupsUsersCount.IsSet() {
		toSerialize["groups_users_count"] = o.GroupsUsersCount.Get()
	}
	if o.Guidelines.IsSet() {
		toSerialize["guidelines"] = o.Guidelines.Get()
	}
	if o.HideConferenceCall.IsSet() {
		toSerialize["hide_conference_call"] = o.HideConferenceCall.Get()
	}
	if o.HideFromGameEight.IsSet() {
		toSerialize["hide_from_game_eight"] = o.HideFromGameEight.Get()
	}
	if o.HideReportedPosts.IsSet() {
		toSerialize["hide_reported_posts"] = o.HideReportedPosts.Get()
	}
	if o.HighlightedCount.IsSet() {
		toSerialize["highlighted_count"] = o.HighlightedCount.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InvitedToJoin.IsSet() {
		toSerialize["invited_to_join"] = o.InvitedToJoin.Get()
	}
	if o.IsJoined.IsSet() {
		toSerialize["is_joined"] = o.IsJoined.Get()
	}
	if o.IsPending.IsSet() {
		toSerialize["is_pending"] = o.IsPending.Get()
	}
	if o.IsPrivate.IsSet() {
		toSerialize["is_private"] = o.IsPrivate.Get()
	}
	if o.IsRelated.IsSet() {
		toSerialize["is_related"] = o.IsRelated.Get()
	}
	if o.JoinedCommunityCampaign.IsSet() {
		toSerialize["joined_community_campaign"] = o.JoinedCommunityCampaign.Get()
	}
	if o.ModeratorIds != nil {
		toSerialize["moderator_ids"] = o.ModeratorIds
	}
	if o.OnlyMobileVerified.IsSet() {
		toSerialize["only_mobile_verified"] = o.OnlyMobileVerified.Get()
	}
	if o.OnlyVerifiedAge.IsSet() {
		toSerialize["only_verified_age"] = o.OnlyVerifiedAge.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.PendingCount.IsSet() {
		toSerialize["pending_count"] = o.PendingCount.Get()
	}
	if o.PendingDeputizeIds != nil {
		toSerialize["pending_deputize_ids"] = o.PendingDeputizeIds
	}
	if o.PendingTransferId.IsSet() {
		toSerialize["pending_transfer_id"] = o.PendingTransferId.Get()
	}
	if o.PostsCount.IsSet() {
		toSerialize["posts_count"] = o.PostsCount.Get()
	}
	if o.RelatedCount.IsSet() {
		toSerialize["related_count"] = o.RelatedCount.Get()
	}
	if o.SafeMode.IsSet() {
		toSerialize["safe_mode"] = o.SafeMode.Get()
	}
	if o.Secret.IsSet() {
		toSerialize["secret"] = o.Secret.Get()
	}
	if o.Seizable.IsSet() {
		toSerialize["seizable"] = o.Seizable.Get()
	}
	if o.SeizableBefore.IsSet() {
		toSerialize["seizable_before"] = o.SeizableBefore.Get()
	}
	if o.SubCategoryId.IsSet() {
		toSerialize["sub_category_id"] = o.SubCategoryId.Get()
	}
	if o.ThreadsCount.IsSet() {
		toSerialize["threads_count"] = o.ThreadsCount.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	if o.UnreadCounts.IsSet() {
		toSerialize["unread_counts"] = o.UnreadCounts.Get()
	}
	if o.UnreadThreadsCount.IsSet() {
		toSerialize["unread_threads_count"] = o.UnreadThreadsCount.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["user_id"] = o.UserId.Get()
	}
	if o.ViewsCount.IsSet() {
		toSerialize["views_count"] = o.ViewsCount.Get()
	}
	if o.WalkthroughRequested.IsSet() {
		toSerialize["walkthrough_requested"] = o.WalkthroughRequested.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Group) UnmarshalJSON(data []byte) (err error) {
	varGroup := _Group{}

	err = json.Unmarshal(data, &varGroup)

	if err != nil {
		return err
	}

	*o = Group(varGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_members_to_post_image_and_video")
		delete(additionalProperties, "allow_members_to_post_url")
		delete(additionalProperties, "allow_ownership_transfer")
		delete(additionalProperties, "allow_thread_creation_by")
		delete(additionalProperties, "call_timeline_display")
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "cover_image_thumbnail")
		delete(additionalProperties, "description")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "generation_groups_limit")
		delete(additionalProperties, "group_category_id")
		delete(additionalProperties, "group_icon")
		delete(additionalProperties, "group_icon_thumbnail")
		delete(additionalProperties, "groups_users_count")
		delete(additionalProperties, "guidelines")
		delete(additionalProperties, "hide_conference_call")
		delete(additionalProperties, "hide_from_game_eight")
		delete(additionalProperties, "hide_reported_posts")
		delete(additionalProperties, "highlighted_count")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invited_to_join")
		delete(additionalProperties, "is_joined")
		delete(additionalProperties, "is_pending")
		delete(additionalProperties, "is_private")
		delete(additionalProperties, "is_related")
		delete(additionalProperties, "joined_community_campaign")
		delete(additionalProperties, "moderator_ids")
		delete(additionalProperties, "only_mobile_verified")
		delete(additionalProperties, "only_verified_age")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "pending_count")
		delete(additionalProperties, "pending_deputize_ids")
		delete(additionalProperties, "pending_transfer_id")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "related_count")
		delete(additionalProperties, "safe_mode")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "seizable")
		delete(additionalProperties, "seizable_before")
		delete(additionalProperties, "sub_category_id")
		delete(additionalProperties, "threads_count")
		delete(additionalProperties, "title")
		delete(additionalProperties, "topic")
		delete(additionalProperties, "unread_counts")
		delete(additionalProperties, "unread_threads_count")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "views_count")
		delete(additionalProperties, "walkthrough_requested")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroup struct {
	value *Group
	isSet bool
}

func (v NullableGroup) Get() *Group {
	return v.value
}

func (v *NullableGroup) Set(val *Group) {
	v.value = val
	v.isSet = true
}

func (v NullableGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroup(val *Group) *NullableGroup {
	return &NullableGroup{value: val, isSet: true}
}

func (v NullableGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


