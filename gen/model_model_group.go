
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGroup{}

// ModelGroup struct for ModelGroup
type ModelGroup struct {
	AllowMembersToPostMedia NullableBool `json:"allow_members_to_post_media,omitempty"`
	AllowMembersToPostUrl NullableBool `json:"allow_members_to_post_url,omitempty"`
	AllowOwnershipTransfer NullableBool `json:"allow_ownership_transfer,omitempty"`
	AllowThreadCreationBy NullableGroupRole `json:"allow_thread_creation_by,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	CoverImageThumbnail NullableString `json:"cover_image_thumbnail,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DisplayIconThumbnail NullableString `json:"display_icon_thumbnail,omitempty"`
	Gender NullableGender `json:"gender,omitempty"`
	Generation NullableGeneration `json:"generation,omitempty"`
	GroupCategoryId NullableInt64 `json:"group_category_id,omitempty"`
	GroupsUsersCount NullableInt64 `json:"groups_users_count,omitempty"`
	Guidelines NullableString `json:"guidelines,omitempty"`
	HighlightedCount NullableInt64 `json:"highlighted_count,omitempty"`
	IconImage NullableString `json:"icon_image,omitempty"`
	IconImageThumbnail NullableString `json:"icon_image_thumbnail,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	InvitedToJoin NullableBool `json:"invited_to_join,omitempty"`
	IsCallTimelineDisplay NullableBool `json:"is_call_timeline_display,omitempty"`
	IsHideConferenceCall NullableBool `json:"is_hide_conference_call,omitempty"`
	IsHideFromGameEight NullableBool `json:"is_hide_from_game_eight,omitempty"`
	IsHideReportedPosts NullableBool `json:"is_hide_reported_posts,omitempty"`
	IsJoined NullableBool `json:"is_joined,omitempty"`
	IsOnlyVerifiedAge NullableBool `json:"is_only_verified_age,omitempty"`
	IsOnlyVerifiedPhone NullableBool `json:"is_only_verified_phone,omitempty"`
	IsPending NullableBool `json:"is_pending,omitempty"`
	IsPinned NullableBool `json:"is_pinned,omitempty"`
	IsPrivate NullableBool `json:"is_private,omitempty"`
	IsSafeMode NullableBool `json:"is_safe_mode,omitempty"`
	IsSecret NullableBool `json:"is_secret,omitempty"`
	IsSeizable NullableBool `json:"is_seizable,omitempty"`
	IsWalkthroughRequested NullableBool `json:"is_walkthrough_requested,omitempty"`
	JoinStatus map[string]interface{} `json:"join_status,omitempty"`
	JoinedCommunityCampaign NullableBool `json:"joined_community_campaign,omitempty"`
	ModeratorIds []int64 `json:"moderator_ids,omitempty"`
	Owner NullableUser `json:"owner,omitempty"`
	PendingCount NullableInt64 `json:"pending_count,omitempty"`
	PendingDeputizeIds []int64 `json:"pending_deputize_ids,omitempty"`
	PendingTransferId NullableInt64 `json:"pending_transfer_id,omitempty"`
	PostsCount NullableInt64 `json:"posts_count,omitempty"`
	RelatedCount NullableInt64 `json:"related_count,omitempty"`
	SeizableBeforeMillis NullableInt64 `json:"seizable_before_millis,omitempty"`
	SubCategoryId NullableInt64 `json:"sub_category_id,omitempty"`
	ThreadsCount NullableInt64 `json:"threads_count,omitempty"`
	Title map[string]interface{} `json:"title,omitempty"`
	Topic NullableString `json:"topic,omitempty"`
	UnreadCounts NullableInt64 `json:"unread_counts,omitempty"`
	UnreadThreadsCount NullableInt64 `json:"unread_threads_count,omitempty"`
	UpdatedAtSeconds NullableInt64 `json:"updated_at_seconds,omitempty"`
	ViewsCount NullableInt64 `json:"views_count,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelGroup ModelGroup

// NewModelGroup instantiates a new ModelGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGroup() *ModelGroup {
	this := ModelGroup{}
	return &this
}

// NewModelGroupWithDefaults instantiates a new ModelGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGroupWithDefaults() *ModelGroup {
	this := ModelGroup{}
	return &this
}

// GetAllowMembersToPostMedia returns the AllowMembersToPostMedia field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetAllowMembersToPostMedia() bool {
	if o == nil || IsNil(o.AllowMembersToPostMedia.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowMembersToPostMedia.Get()
}

// GetAllowMembersToPostMediaOk returns a tuple with the AllowMembersToPostMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetAllowMembersToPostMediaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowMembersToPostMedia.Get(), o.AllowMembersToPostMedia.IsSet()
}

// HasAllowMembersToPostMedia returns a boolean if a field has been set.
func (o *ModelGroup) HasAllowMembersToPostMedia() bool {
	if o != nil && o.AllowMembersToPostMedia.IsSet() {
		return true
	}

	return false
}

// SetAllowMembersToPostMedia gets a reference to the given NullableBool and assigns it to the AllowMembersToPostMedia field.
func (o *ModelGroup) SetAllowMembersToPostMedia(v bool) {
	o.AllowMembersToPostMedia.Set(&v)
}
// SetAllowMembersToPostMediaNil sets the value for AllowMembersToPostMedia to be an explicit nil
func (o *ModelGroup) SetAllowMembersToPostMediaNil() {
	o.AllowMembersToPostMedia.Set(nil)
}

// UnsetAllowMembersToPostMedia ensures that no value is present for AllowMembersToPostMedia, not even an explicit nil
func (o *ModelGroup) UnsetAllowMembersToPostMedia() {
	o.AllowMembersToPostMedia.Unset()
}

// GetAllowMembersToPostUrl returns the AllowMembersToPostUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetAllowMembersToPostUrl() bool {
	if o == nil || IsNil(o.AllowMembersToPostUrl.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowMembersToPostUrl.Get()
}

// GetAllowMembersToPostUrlOk returns a tuple with the AllowMembersToPostUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetAllowMembersToPostUrlOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowMembersToPostUrl.Get(), o.AllowMembersToPostUrl.IsSet()
}

// HasAllowMembersToPostUrl returns a boolean if a field has been set.
func (o *ModelGroup) HasAllowMembersToPostUrl() bool {
	if o != nil && o.AllowMembersToPostUrl.IsSet() {
		return true
	}

	return false
}

// SetAllowMembersToPostUrl gets a reference to the given NullableBool and assigns it to the AllowMembersToPostUrl field.
func (o *ModelGroup) SetAllowMembersToPostUrl(v bool) {
	o.AllowMembersToPostUrl.Set(&v)
}
// SetAllowMembersToPostUrlNil sets the value for AllowMembersToPostUrl to be an explicit nil
func (o *ModelGroup) SetAllowMembersToPostUrlNil() {
	o.AllowMembersToPostUrl.Set(nil)
}

// UnsetAllowMembersToPostUrl ensures that no value is present for AllowMembersToPostUrl, not even an explicit nil
func (o *ModelGroup) UnsetAllowMembersToPostUrl() {
	o.AllowMembersToPostUrl.Unset()
}

// GetAllowOwnershipTransfer returns the AllowOwnershipTransfer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetAllowOwnershipTransfer() bool {
	if o == nil || IsNil(o.AllowOwnershipTransfer.Get()) {
		var ret bool
		return ret
	}
	return *o.AllowOwnershipTransfer.Get()
}

// GetAllowOwnershipTransferOk returns a tuple with the AllowOwnershipTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetAllowOwnershipTransferOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowOwnershipTransfer.Get(), o.AllowOwnershipTransfer.IsSet()
}

// HasAllowOwnershipTransfer returns a boolean if a field has been set.
func (o *ModelGroup) HasAllowOwnershipTransfer() bool {
	if o != nil && o.AllowOwnershipTransfer.IsSet() {
		return true
	}

	return false
}

// SetAllowOwnershipTransfer gets a reference to the given NullableBool and assigns it to the AllowOwnershipTransfer field.
func (o *ModelGroup) SetAllowOwnershipTransfer(v bool) {
	o.AllowOwnershipTransfer.Set(&v)
}
// SetAllowOwnershipTransferNil sets the value for AllowOwnershipTransfer to be an explicit nil
func (o *ModelGroup) SetAllowOwnershipTransferNil() {
	o.AllowOwnershipTransfer.Set(nil)
}

// UnsetAllowOwnershipTransfer ensures that no value is present for AllowOwnershipTransfer, not even an explicit nil
func (o *ModelGroup) UnsetAllowOwnershipTransfer() {
	o.AllowOwnershipTransfer.Unset()
}

// GetAllowThreadCreationBy returns the AllowThreadCreationBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetAllowThreadCreationBy() GroupRole {
	if o == nil || IsNil(o.AllowThreadCreationBy.Get()) {
		var ret GroupRole
		return ret
	}
	return *o.AllowThreadCreationBy.Get()
}

// GetAllowThreadCreationByOk returns a tuple with the AllowThreadCreationBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetAllowThreadCreationByOk() (*GroupRole, bool) {
	if o == nil {
		return nil, false
	}
	return o.AllowThreadCreationBy.Get(), o.AllowThreadCreationBy.IsSet()
}

// HasAllowThreadCreationBy returns a boolean if a field has been set.
func (o *ModelGroup) HasAllowThreadCreationBy() bool {
	if o != nil && o.AllowThreadCreationBy.IsSet() {
		return true
	}

	return false
}

// SetAllowThreadCreationBy gets a reference to the given NullableGroupRole and assigns it to the AllowThreadCreationBy field.
func (o *ModelGroup) SetAllowThreadCreationBy(v GroupRole) {
	o.AllowThreadCreationBy.Set(&v)
}
// SetAllowThreadCreationByNil sets the value for AllowThreadCreationBy to be an explicit nil
func (o *ModelGroup) SetAllowThreadCreationByNil() {
	o.AllowThreadCreationBy.Set(nil)
}

// UnsetAllowThreadCreationBy ensures that no value is present for AllowThreadCreationBy, not even an explicit nil
func (o *ModelGroup) UnsetAllowThreadCreationBy() {
	o.AllowThreadCreationBy.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetCoverImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *ModelGroup) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *ModelGroup) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}
// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *ModelGroup) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *ModelGroup) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetCoverImageThumbnail returns the CoverImageThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetCoverImageThumbnail() string {
	if o == nil || IsNil(o.CoverImageThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImageThumbnail.Get()
}

// GetCoverImageThumbnailOk returns a tuple with the CoverImageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetCoverImageThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImageThumbnail.Get(), o.CoverImageThumbnail.IsSet()
}

// HasCoverImageThumbnail returns a boolean if a field has been set.
func (o *ModelGroup) HasCoverImageThumbnail() bool {
	if o != nil && o.CoverImageThumbnail.IsSet() {
		return true
	}

	return false
}

// SetCoverImageThumbnail gets a reference to the given NullableString and assigns it to the CoverImageThumbnail field.
func (o *ModelGroup) SetCoverImageThumbnail(v string) {
	o.CoverImageThumbnail.Set(&v)
}
// SetCoverImageThumbnailNil sets the value for CoverImageThumbnail to be an explicit nil
func (o *ModelGroup) SetCoverImageThumbnailNil() {
	o.CoverImageThumbnail.Set(nil)
}

// UnsetCoverImageThumbnail ensures that no value is present for CoverImageThumbnail, not even an explicit nil
func (o *ModelGroup) UnsetCoverImageThumbnail() {
	o.CoverImageThumbnail.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelGroup) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *ModelGroup) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *ModelGroup) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *ModelGroup) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayIconThumbnail returns the DisplayIconThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetDisplayIconThumbnail() string {
	if o == nil || IsNil(o.DisplayIconThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.DisplayIconThumbnail.Get()
}

// GetDisplayIconThumbnailOk returns a tuple with the DisplayIconThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetDisplayIconThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisplayIconThumbnail.Get(), o.DisplayIconThumbnail.IsSet()
}

// HasDisplayIconThumbnail returns a boolean if a field has been set.
func (o *ModelGroup) HasDisplayIconThumbnail() bool {
	if o != nil && o.DisplayIconThumbnail.IsSet() {
		return true
	}

	return false
}

// SetDisplayIconThumbnail gets a reference to the given NullableString and assigns it to the DisplayIconThumbnail field.
func (o *ModelGroup) SetDisplayIconThumbnail(v string) {
	o.DisplayIconThumbnail.Set(&v)
}
// SetDisplayIconThumbnailNil sets the value for DisplayIconThumbnail to be an explicit nil
func (o *ModelGroup) SetDisplayIconThumbnailNil() {
	o.DisplayIconThumbnail.Set(nil)
}

// UnsetDisplayIconThumbnail ensures that no value is present for DisplayIconThumbnail, not even an explicit nil
func (o *ModelGroup) UnsetDisplayIconThumbnail() {
	o.DisplayIconThumbnail.Unset()
}

// GetGender returns the Gender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetGender() Gender {
	if o == nil || IsNil(o.Gender.Get()) {
		var ret Gender
		return ret
	}
	return *o.Gender.Get()
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetGenderOk() (*Gender, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gender.Get(), o.Gender.IsSet()
}

// HasGender returns a boolean if a field has been set.
func (o *ModelGroup) HasGender() bool {
	if o != nil && o.Gender.IsSet() {
		return true
	}

	return false
}

// SetGender gets a reference to the given NullableGender and assigns it to the Gender field.
func (o *ModelGroup) SetGender(v Gender) {
	o.Gender.Set(&v)
}
// SetGenderNil sets the value for Gender to be an explicit nil
func (o *ModelGroup) SetGenderNil() {
	o.Gender.Set(nil)
}

// UnsetGender ensures that no value is present for Gender, not even an explicit nil
func (o *ModelGroup) UnsetGender() {
	o.Gender.Unset()
}

// GetGeneration returns the Generation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetGeneration() Generation {
	if o == nil || IsNil(o.Generation.Get()) {
		var ret Generation
		return ret
	}
	return *o.Generation.Get()
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetGenerationOk() (*Generation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Generation.Get(), o.Generation.IsSet()
}

// HasGeneration returns a boolean if a field has been set.
func (o *ModelGroup) HasGeneration() bool {
	if o != nil && o.Generation.IsSet() {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given NullableGeneration and assigns it to the Generation field.
func (o *ModelGroup) SetGeneration(v Generation) {
	o.Generation.Set(&v)
}
// SetGenerationNil sets the value for Generation to be an explicit nil
func (o *ModelGroup) SetGenerationNil() {
	o.Generation.Set(nil)
}

// UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
func (o *ModelGroup) UnsetGeneration() {
	o.Generation.Unset()
}

// GetGroupCategoryId returns the GroupCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetGroupCategoryId() int64 {
	if o == nil || IsNil(o.GroupCategoryId.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupCategoryId.Get()
}

// GetGroupCategoryIdOk returns a tuple with the GroupCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetGroupCategoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupCategoryId.Get(), o.GroupCategoryId.IsSet()
}

// HasGroupCategoryId returns a boolean if a field has been set.
func (o *ModelGroup) HasGroupCategoryId() bool {
	if o != nil && o.GroupCategoryId.IsSet() {
		return true
	}

	return false
}

// SetGroupCategoryId gets a reference to the given NullableInt64 and assigns it to the GroupCategoryId field.
func (o *ModelGroup) SetGroupCategoryId(v int64) {
	o.GroupCategoryId.Set(&v)
}
// SetGroupCategoryIdNil sets the value for GroupCategoryId to be an explicit nil
func (o *ModelGroup) SetGroupCategoryIdNil() {
	o.GroupCategoryId.Set(nil)
}

// UnsetGroupCategoryId ensures that no value is present for GroupCategoryId, not even an explicit nil
func (o *ModelGroup) UnsetGroupCategoryId() {
	o.GroupCategoryId.Unset()
}

// GetGroupsUsersCount returns the GroupsUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetGroupsUsersCount() int64 {
	if o == nil || IsNil(o.GroupsUsersCount.Get()) {
		var ret int64
		return ret
	}
	return *o.GroupsUsersCount.Get()
}

// GetGroupsUsersCountOk returns a tuple with the GroupsUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetGroupsUsersCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupsUsersCount.Get(), o.GroupsUsersCount.IsSet()
}

// HasGroupsUsersCount returns a boolean if a field has been set.
func (o *ModelGroup) HasGroupsUsersCount() bool {
	if o != nil && o.GroupsUsersCount.IsSet() {
		return true
	}

	return false
}

// SetGroupsUsersCount gets a reference to the given NullableInt64 and assigns it to the GroupsUsersCount field.
func (o *ModelGroup) SetGroupsUsersCount(v int64) {
	o.GroupsUsersCount.Set(&v)
}
// SetGroupsUsersCountNil sets the value for GroupsUsersCount to be an explicit nil
func (o *ModelGroup) SetGroupsUsersCountNil() {
	o.GroupsUsersCount.Set(nil)
}

// UnsetGroupsUsersCount ensures that no value is present for GroupsUsersCount, not even an explicit nil
func (o *ModelGroup) UnsetGroupsUsersCount() {
	o.GroupsUsersCount.Unset()
}

// GetGuidelines returns the Guidelines field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetGuidelines() string {
	if o == nil || IsNil(o.Guidelines.Get()) {
		var ret string
		return ret
	}
	return *o.Guidelines.Get()
}

// GetGuidelinesOk returns a tuple with the Guidelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetGuidelinesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Guidelines.Get(), o.Guidelines.IsSet()
}

// HasGuidelines returns a boolean if a field has been set.
func (o *ModelGroup) HasGuidelines() bool {
	if o != nil && o.Guidelines.IsSet() {
		return true
	}

	return false
}

// SetGuidelines gets a reference to the given NullableString and assigns it to the Guidelines field.
func (o *ModelGroup) SetGuidelines(v string) {
	o.Guidelines.Set(&v)
}
// SetGuidelinesNil sets the value for Guidelines to be an explicit nil
func (o *ModelGroup) SetGuidelinesNil() {
	o.Guidelines.Set(nil)
}

// UnsetGuidelines ensures that no value is present for Guidelines, not even an explicit nil
func (o *ModelGroup) UnsetGuidelines() {
	o.Guidelines.Unset()
}

// GetHighlightedCount returns the HighlightedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetHighlightedCount() int64 {
	if o == nil || IsNil(o.HighlightedCount.Get()) {
		var ret int64
		return ret
	}
	return *o.HighlightedCount.Get()
}

// GetHighlightedCountOk returns a tuple with the HighlightedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetHighlightedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighlightedCount.Get(), o.HighlightedCount.IsSet()
}

// HasHighlightedCount returns a boolean if a field has been set.
func (o *ModelGroup) HasHighlightedCount() bool {
	if o != nil && o.HighlightedCount.IsSet() {
		return true
	}

	return false
}

// SetHighlightedCount gets a reference to the given NullableInt64 and assigns it to the HighlightedCount field.
func (o *ModelGroup) SetHighlightedCount(v int64) {
	o.HighlightedCount.Set(&v)
}
// SetHighlightedCountNil sets the value for HighlightedCount to be an explicit nil
func (o *ModelGroup) SetHighlightedCountNil() {
	o.HighlightedCount.Set(nil)
}

// UnsetHighlightedCount ensures that no value is present for HighlightedCount, not even an explicit nil
func (o *ModelGroup) UnsetHighlightedCount() {
	o.HighlightedCount.Unset()
}

// GetIconImage returns the IconImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIconImage() string {
	if o == nil || IsNil(o.IconImage.Get()) {
		var ret string
		return ret
	}
	return *o.IconImage.Get()
}

// GetIconImageOk returns a tuple with the IconImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIconImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconImage.Get(), o.IconImage.IsSet()
}

// HasIconImage returns a boolean if a field has been set.
func (o *ModelGroup) HasIconImage() bool {
	if o != nil && o.IconImage.IsSet() {
		return true
	}

	return false
}

// SetIconImage gets a reference to the given NullableString and assigns it to the IconImage field.
func (o *ModelGroup) SetIconImage(v string) {
	o.IconImage.Set(&v)
}
// SetIconImageNil sets the value for IconImage to be an explicit nil
func (o *ModelGroup) SetIconImageNil() {
	o.IconImage.Set(nil)
}

// UnsetIconImage ensures that no value is present for IconImage, not even an explicit nil
func (o *ModelGroup) UnsetIconImage() {
	o.IconImage.Unset()
}

// GetIconImageThumbnail returns the IconImageThumbnail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIconImageThumbnail() string {
	if o == nil || IsNil(o.IconImageThumbnail.Get()) {
		var ret string
		return ret
	}
	return *o.IconImageThumbnail.Get()
}

// GetIconImageThumbnailOk returns a tuple with the IconImageThumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIconImageThumbnailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconImageThumbnail.Get(), o.IconImageThumbnail.IsSet()
}

// HasIconImageThumbnail returns a boolean if a field has been set.
func (o *ModelGroup) HasIconImageThumbnail() bool {
	if o != nil && o.IconImageThumbnail.IsSet() {
		return true
	}

	return false
}

// SetIconImageThumbnail gets a reference to the given NullableString and assigns it to the IconImageThumbnail field.
func (o *ModelGroup) SetIconImageThumbnail(v string) {
	o.IconImageThumbnail.Set(&v)
}
// SetIconImageThumbnailNil sets the value for IconImageThumbnail to be an explicit nil
func (o *ModelGroup) SetIconImageThumbnailNil() {
	o.IconImageThumbnail.Set(nil)
}

// UnsetIconImageThumbnail ensures that no value is present for IconImageThumbnail, not even an explicit nil
func (o *ModelGroup) UnsetIconImageThumbnail() {
	o.IconImageThumbnail.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelGroup) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelGroup) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelGroup) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelGroup) UnsetId() {
	o.Id.Unset()
}

// GetInvitedToJoin returns the InvitedToJoin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetInvitedToJoin() bool {
	if o == nil || IsNil(o.InvitedToJoin.Get()) {
		var ret bool
		return ret
	}
	return *o.InvitedToJoin.Get()
}

// GetInvitedToJoinOk returns a tuple with the InvitedToJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetInvitedToJoinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InvitedToJoin.Get(), o.InvitedToJoin.IsSet()
}

// HasInvitedToJoin returns a boolean if a field has been set.
func (o *ModelGroup) HasInvitedToJoin() bool {
	if o != nil && o.InvitedToJoin.IsSet() {
		return true
	}

	return false
}

// SetInvitedToJoin gets a reference to the given NullableBool and assigns it to the InvitedToJoin field.
func (o *ModelGroup) SetInvitedToJoin(v bool) {
	o.InvitedToJoin.Set(&v)
}
// SetInvitedToJoinNil sets the value for InvitedToJoin to be an explicit nil
func (o *ModelGroup) SetInvitedToJoinNil() {
	o.InvitedToJoin.Set(nil)
}

// UnsetInvitedToJoin ensures that no value is present for InvitedToJoin, not even an explicit nil
func (o *ModelGroup) UnsetInvitedToJoin() {
	o.InvitedToJoin.Unset()
}

// GetIsCallTimelineDisplay returns the IsCallTimelineDisplay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsCallTimelineDisplay() bool {
	if o == nil || IsNil(o.IsCallTimelineDisplay.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCallTimelineDisplay.Get()
}

// GetIsCallTimelineDisplayOk returns a tuple with the IsCallTimelineDisplay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsCallTimelineDisplayOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCallTimelineDisplay.Get(), o.IsCallTimelineDisplay.IsSet()
}

// HasIsCallTimelineDisplay returns a boolean if a field has been set.
func (o *ModelGroup) HasIsCallTimelineDisplay() bool {
	if o != nil && o.IsCallTimelineDisplay.IsSet() {
		return true
	}

	return false
}

// SetIsCallTimelineDisplay gets a reference to the given NullableBool and assigns it to the IsCallTimelineDisplay field.
func (o *ModelGroup) SetIsCallTimelineDisplay(v bool) {
	o.IsCallTimelineDisplay.Set(&v)
}
// SetIsCallTimelineDisplayNil sets the value for IsCallTimelineDisplay to be an explicit nil
func (o *ModelGroup) SetIsCallTimelineDisplayNil() {
	o.IsCallTimelineDisplay.Set(nil)
}

// UnsetIsCallTimelineDisplay ensures that no value is present for IsCallTimelineDisplay, not even an explicit nil
func (o *ModelGroup) UnsetIsCallTimelineDisplay() {
	o.IsCallTimelineDisplay.Unset()
}

// GetIsHideConferenceCall returns the IsHideConferenceCall field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsHideConferenceCall() bool {
	if o == nil || IsNil(o.IsHideConferenceCall.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHideConferenceCall.Get()
}

// GetIsHideConferenceCallOk returns a tuple with the IsHideConferenceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsHideConferenceCallOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHideConferenceCall.Get(), o.IsHideConferenceCall.IsSet()
}

// HasIsHideConferenceCall returns a boolean if a field has been set.
func (o *ModelGroup) HasIsHideConferenceCall() bool {
	if o != nil && o.IsHideConferenceCall.IsSet() {
		return true
	}

	return false
}

// SetIsHideConferenceCall gets a reference to the given NullableBool and assigns it to the IsHideConferenceCall field.
func (o *ModelGroup) SetIsHideConferenceCall(v bool) {
	o.IsHideConferenceCall.Set(&v)
}
// SetIsHideConferenceCallNil sets the value for IsHideConferenceCall to be an explicit nil
func (o *ModelGroup) SetIsHideConferenceCallNil() {
	o.IsHideConferenceCall.Set(nil)
}

// UnsetIsHideConferenceCall ensures that no value is present for IsHideConferenceCall, not even an explicit nil
func (o *ModelGroup) UnsetIsHideConferenceCall() {
	o.IsHideConferenceCall.Unset()
}

// GetIsHideFromGameEight returns the IsHideFromGameEight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsHideFromGameEight() bool {
	if o == nil || IsNil(o.IsHideFromGameEight.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHideFromGameEight.Get()
}

// GetIsHideFromGameEightOk returns a tuple with the IsHideFromGameEight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsHideFromGameEightOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHideFromGameEight.Get(), o.IsHideFromGameEight.IsSet()
}

// HasIsHideFromGameEight returns a boolean if a field has been set.
func (o *ModelGroup) HasIsHideFromGameEight() bool {
	if o != nil && o.IsHideFromGameEight.IsSet() {
		return true
	}

	return false
}

// SetIsHideFromGameEight gets a reference to the given NullableBool and assigns it to the IsHideFromGameEight field.
func (o *ModelGroup) SetIsHideFromGameEight(v bool) {
	o.IsHideFromGameEight.Set(&v)
}
// SetIsHideFromGameEightNil sets the value for IsHideFromGameEight to be an explicit nil
func (o *ModelGroup) SetIsHideFromGameEightNil() {
	o.IsHideFromGameEight.Set(nil)
}

// UnsetIsHideFromGameEight ensures that no value is present for IsHideFromGameEight, not even an explicit nil
func (o *ModelGroup) UnsetIsHideFromGameEight() {
	o.IsHideFromGameEight.Unset()
}

// GetIsHideReportedPosts returns the IsHideReportedPosts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsHideReportedPosts() bool {
	if o == nil || IsNil(o.IsHideReportedPosts.Get()) {
		var ret bool
		return ret
	}
	return *o.IsHideReportedPosts.Get()
}

// GetIsHideReportedPostsOk returns a tuple with the IsHideReportedPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsHideReportedPostsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsHideReportedPosts.Get(), o.IsHideReportedPosts.IsSet()
}

// HasIsHideReportedPosts returns a boolean if a field has been set.
func (o *ModelGroup) HasIsHideReportedPosts() bool {
	if o != nil && o.IsHideReportedPosts.IsSet() {
		return true
	}

	return false
}

// SetIsHideReportedPosts gets a reference to the given NullableBool and assigns it to the IsHideReportedPosts field.
func (o *ModelGroup) SetIsHideReportedPosts(v bool) {
	o.IsHideReportedPosts.Set(&v)
}
// SetIsHideReportedPostsNil sets the value for IsHideReportedPosts to be an explicit nil
func (o *ModelGroup) SetIsHideReportedPostsNil() {
	o.IsHideReportedPosts.Set(nil)
}

// UnsetIsHideReportedPosts ensures that no value is present for IsHideReportedPosts, not even an explicit nil
func (o *ModelGroup) UnsetIsHideReportedPosts() {
	o.IsHideReportedPosts.Unset()
}

// GetIsJoined returns the IsJoined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsJoined() bool {
	if o == nil || IsNil(o.IsJoined.Get()) {
		var ret bool
		return ret
	}
	return *o.IsJoined.Get()
}

// GetIsJoinedOk returns a tuple with the IsJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsJoinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsJoined.Get(), o.IsJoined.IsSet()
}

// HasIsJoined returns a boolean if a field has been set.
func (o *ModelGroup) HasIsJoined() bool {
	if o != nil && o.IsJoined.IsSet() {
		return true
	}

	return false
}

// SetIsJoined gets a reference to the given NullableBool and assigns it to the IsJoined field.
func (o *ModelGroup) SetIsJoined(v bool) {
	o.IsJoined.Set(&v)
}
// SetIsJoinedNil sets the value for IsJoined to be an explicit nil
func (o *ModelGroup) SetIsJoinedNil() {
	o.IsJoined.Set(nil)
}

// UnsetIsJoined ensures that no value is present for IsJoined, not even an explicit nil
func (o *ModelGroup) UnsetIsJoined() {
	o.IsJoined.Unset()
}

// GetIsOnlyVerifiedAge returns the IsOnlyVerifiedAge field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsOnlyVerifiedAge() bool {
	if o == nil || IsNil(o.IsOnlyVerifiedAge.Get()) {
		var ret bool
		return ret
	}
	return *o.IsOnlyVerifiedAge.Get()
}

// GetIsOnlyVerifiedAgeOk returns a tuple with the IsOnlyVerifiedAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsOnlyVerifiedAgeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsOnlyVerifiedAge.Get(), o.IsOnlyVerifiedAge.IsSet()
}

// HasIsOnlyVerifiedAge returns a boolean if a field has been set.
func (o *ModelGroup) HasIsOnlyVerifiedAge() bool {
	if o != nil && o.IsOnlyVerifiedAge.IsSet() {
		return true
	}

	return false
}

// SetIsOnlyVerifiedAge gets a reference to the given NullableBool and assigns it to the IsOnlyVerifiedAge field.
func (o *ModelGroup) SetIsOnlyVerifiedAge(v bool) {
	o.IsOnlyVerifiedAge.Set(&v)
}
// SetIsOnlyVerifiedAgeNil sets the value for IsOnlyVerifiedAge to be an explicit nil
func (o *ModelGroup) SetIsOnlyVerifiedAgeNil() {
	o.IsOnlyVerifiedAge.Set(nil)
}

// UnsetIsOnlyVerifiedAge ensures that no value is present for IsOnlyVerifiedAge, not even an explicit nil
func (o *ModelGroup) UnsetIsOnlyVerifiedAge() {
	o.IsOnlyVerifiedAge.Unset()
}

// GetIsOnlyVerifiedPhone returns the IsOnlyVerifiedPhone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsOnlyVerifiedPhone() bool {
	if o == nil || IsNil(o.IsOnlyVerifiedPhone.Get()) {
		var ret bool
		return ret
	}
	return *o.IsOnlyVerifiedPhone.Get()
}

// GetIsOnlyVerifiedPhoneOk returns a tuple with the IsOnlyVerifiedPhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsOnlyVerifiedPhoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsOnlyVerifiedPhone.Get(), o.IsOnlyVerifiedPhone.IsSet()
}

// HasIsOnlyVerifiedPhone returns a boolean if a field has been set.
func (o *ModelGroup) HasIsOnlyVerifiedPhone() bool {
	if o != nil && o.IsOnlyVerifiedPhone.IsSet() {
		return true
	}

	return false
}

// SetIsOnlyVerifiedPhone gets a reference to the given NullableBool and assigns it to the IsOnlyVerifiedPhone field.
func (o *ModelGroup) SetIsOnlyVerifiedPhone(v bool) {
	o.IsOnlyVerifiedPhone.Set(&v)
}
// SetIsOnlyVerifiedPhoneNil sets the value for IsOnlyVerifiedPhone to be an explicit nil
func (o *ModelGroup) SetIsOnlyVerifiedPhoneNil() {
	o.IsOnlyVerifiedPhone.Set(nil)
}

// UnsetIsOnlyVerifiedPhone ensures that no value is present for IsOnlyVerifiedPhone, not even an explicit nil
func (o *ModelGroup) UnsetIsOnlyVerifiedPhone() {
	o.IsOnlyVerifiedPhone.Unset()
}

// GetIsPending returns the IsPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsPending() bool {
	if o == nil || IsNil(o.IsPending.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPending.Get()
}

// GetIsPendingOk returns a tuple with the IsPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPending.Get(), o.IsPending.IsSet()
}

// HasIsPending returns a boolean if a field has been set.
func (o *ModelGroup) HasIsPending() bool {
	if o != nil && o.IsPending.IsSet() {
		return true
	}

	return false
}

// SetIsPending gets a reference to the given NullableBool and assigns it to the IsPending field.
func (o *ModelGroup) SetIsPending(v bool) {
	o.IsPending.Set(&v)
}
// SetIsPendingNil sets the value for IsPending to be an explicit nil
func (o *ModelGroup) SetIsPendingNil() {
	o.IsPending.Set(nil)
}

// UnsetIsPending ensures that no value is present for IsPending, not even an explicit nil
func (o *ModelGroup) UnsetIsPending() {
	o.IsPending.Unset()
}

// GetIsPinned returns the IsPinned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsPinned() bool {
	if o == nil || IsNil(o.IsPinned.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPinned.Get()
}

// GetIsPinnedOk returns a tuple with the IsPinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPinned.Get(), o.IsPinned.IsSet()
}

// HasIsPinned returns a boolean if a field has been set.
func (o *ModelGroup) HasIsPinned() bool {
	if o != nil && o.IsPinned.IsSet() {
		return true
	}

	return false
}

// SetIsPinned gets a reference to the given NullableBool and assigns it to the IsPinned field.
func (o *ModelGroup) SetIsPinned(v bool) {
	o.IsPinned.Set(&v)
}
// SetIsPinnedNil sets the value for IsPinned to be an explicit nil
func (o *ModelGroup) SetIsPinnedNil() {
	o.IsPinned.Set(nil)
}

// UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
func (o *ModelGroup) UnsetIsPinned() {
	o.IsPinned.Unset()
}

// GetIsPrivate returns the IsPrivate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsPrivate() bool {
	if o == nil || IsNil(o.IsPrivate.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPrivate.Get()
}

// GetIsPrivateOk returns a tuple with the IsPrivate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsPrivateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPrivate.Get(), o.IsPrivate.IsSet()
}

// HasIsPrivate returns a boolean if a field has been set.
func (o *ModelGroup) HasIsPrivate() bool {
	if o != nil && o.IsPrivate.IsSet() {
		return true
	}

	return false
}

// SetIsPrivate gets a reference to the given NullableBool and assigns it to the IsPrivate field.
func (o *ModelGroup) SetIsPrivate(v bool) {
	o.IsPrivate.Set(&v)
}
// SetIsPrivateNil sets the value for IsPrivate to be an explicit nil
func (o *ModelGroup) SetIsPrivateNil() {
	o.IsPrivate.Set(nil)
}

// UnsetIsPrivate ensures that no value is present for IsPrivate, not even an explicit nil
func (o *ModelGroup) UnsetIsPrivate() {
	o.IsPrivate.Unset()
}

// GetIsSafeMode returns the IsSafeMode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsSafeMode() bool {
	if o == nil || IsNil(o.IsSafeMode.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSafeMode.Get()
}

// GetIsSafeModeOk returns a tuple with the IsSafeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsSafeModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSafeMode.Get(), o.IsSafeMode.IsSet()
}

// HasIsSafeMode returns a boolean if a field has been set.
func (o *ModelGroup) HasIsSafeMode() bool {
	if o != nil && o.IsSafeMode.IsSet() {
		return true
	}

	return false
}

// SetIsSafeMode gets a reference to the given NullableBool and assigns it to the IsSafeMode field.
func (o *ModelGroup) SetIsSafeMode(v bool) {
	o.IsSafeMode.Set(&v)
}
// SetIsSafeModeNil sets the value for IsSafeMode to be an explicit nil
func (o *ModelGroup) SetIsSafeModeNil() {
	o.IsSafeMode.Set(nil)
}

// UnsetIsSafeMode ensures that no value is present for IsSafeMode, not even an explicit nil
func (o *ModelGroup) UnsetIsSafeMode() {
	o.IsSafeMode.Unset()
}

// GetIsSecret returns the IsSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsSecret() bool {
	if o == nil || IsNil(o.IsSecret.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSecret.Get()
}

// GetIsSecretOk returns a tuple with the IsSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsSecretOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSecret.Get(), o.IsSecret.IsSet()
}

// HasIsSecret returns a boolean if a field has been set.
func (o *ModelGroup) HasIsSecret() bool {
	if o != nil && o.IsSecret.IsSet() {
		return true
	}

	return false
}

// SetIsSecret gets a reference to the given NullableBool and assigns it to the IsSecret field.
func (o *ModelGroup) SetIsSecret(v bool) {
	o.IsSecret.Set(&v)
}
// SetIsSecretNil sets the value for IsSecret to be an explicit nil
func (o *ModelGroup) SetIsSecretNil() {
	o.IsSecret.Set(nil)
}

// UnsetIsSecret ensures that no value is present for IsSecret, not even an explicit nil
func (o *ModelGroup) UnsetIsSecret() {
	o.IsSecret.Unset()
}

// GetIsSeizable returns the IsSeizable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsSeizable() bool {
	if o == nil || IsNil(o.IsSeizable.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSeizable.Get()
}

// GetIsSeizableOk returns a tuple with the IsSeizable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsSeizableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSeizable.Get(), o.IsSeizable.IsSet()
}

// HasIsSeizable returns a boolean if a field has been set.
func (o *ModelGroup) HasIsSeizable() bool {
	if o != nil && o.IsSeizable.IsSet() {
		return true
	}

	return false
}

// SetIsSeizable gets a reference to the given NullableBool and assigns it to the IsSeizable field.
func (o *ModelGroup) SetIsSeizable(v bool) {
	o.IsSeizable.Set(&v)
}
// SetIsSeizableNil sets the value for IsSeizable to be an explicit nil
func (o *ModelGroup) SetIsSeizableNil() {
	o.IsSeizable.Set(nil)
}

// UnsetIsSeizable ensures that no value is present for IsSeizable, not even an explicit nil
func (o *ModelGroup) UnsetIsSeizable() {
	o.IsSeizable.Unset()
}

// GetIsWalkthroughRequested returns the IsWalkthroughRequested field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetIsWalkthroughRequested() bool {
	if o == nil || IsNil(o.IsWalkthroughRequested.Get()) {
		var ret bool
		return ret
	}
	return *o.IsWalkthroughRequested.Get()
}

// GetIsWalkthroughRequestedOk returns a tuple with the IsWalkthroughRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetIsWalkthroughRequestedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsWalkthroughRequested.Get(), o.IsWalkthroughRequested.IsSet()
}

// HasIsWalkthroughRequested returns a boolean if a field has been set.
func (o *ModelGroup) HasIsWalkthroughRequested() bool {
	if o != nil && o.IsWalkthroughRequested.IsSet() {
		return true
	}

	return false
}

// SetIsWalkthroughRequested gets a reference to the given NullableBool and assigns it to the IsWalkthroughRequested field.
func (o *ModelGroup) SetIsWalkthroughRequested(v bool) {
	o.IsWalkthroughRequested.Set(&v)
}
// SetIsWalkthroughRequestedNil sets the value for IsWalkthroughRequested to be an explicit nil
func (o *ModelGroup) SetIsWalkthroughRequestedNil() {
	o.IsWalkthroughRequested.Set(nil)
}

// UnsetIsWalkthroughRequested ensures that no value is present for IsWalkthroughRequested, not even an explicit nil
func (o *ModelGroup) UnsetIsWalkthroughRequested() {
	o.IsWalkthroughRequested.Unset()
}

// GetJoinStatus returns the JoinStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetJoinStatus() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.JoinStatus
}

// GetJoinStatusOk returns a tuple with the JoinStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetJoinStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.JoinStatus) {
		return map[string]interface{}{}, false
	}
	return o.JoinStatus, true
}

// HasJoinStatus returns a boolean if a field has been set.
func (o *ModelGroup) HasJoinStatus() bool {
	if o != nil && !IsNil(o.JoinStatus) {
		return true
	}

	return false
}

// SetJoinStatus gets a reference to the given map[string]interface{} and assigns it to the JoinStatus field.
func (o *ModelGroup) SetJoinStatus(v map[string]interface{}) {
	o.JoinStatus = v
}

// GetJoinedCommunityCampaign returns the JoinedCommunityCampaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetJoinedCommunityCampaign() bool {
	if o == nil || IsNil(o.JoinedCommunityCampaign.Get()) {
		var ret bool
		return ret
	}
	return *o.JoinedCommunityCampaign.Get()
}

// GetJoinedCommunityCampaignOk returns a tuple with the JoinedCommunityCampaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetJoinedCommunityCampaignOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.JoinedCommunityCampaign.Get(), o.JoinedCommunityCampaign.IsSet()
}

// HasJoinedCommunityCampaign returns a boolean if a field has been set.
func (o *ModelGroup) HasJoinedCommunityCampaign() bool {
	if o != nil && o.JoinedCommunityCampaign.IsSet() {
		return true
	}

	return false
}

// SetJoinedCommunityCampaign gets a reference to the given NullableBool and assigns it to the JoinedCommunityCampaign field.
func (o *ModelGroup) SetJoinedCommunityCampaign(v bool) {
	o.JoinedCommunityCampaign.Set(&v)
}
// SetJoinedCommunityCampaignNil sets the value for JoinedCommunityCampaign to be an explicit nil
func (o *ModelGroup) SetJoinedCommunityCampaignNil() {
	o.JoinedCommunityCampaign.Set(nil)
}

// UnsetJoinedCommunityCampaign ensures that no value is present for JoinedCommunityCampaign, not even an explicit nil
func (o *ModelGroup) UnsetJoinedCommunityCampaign() {
	o.JoinedCommunityCampaign.Unset()
}

// GetModeratorIds returns the ModeratorIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetModeratorIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.ModeratorIds
}

// GetModeratorIdsOk returns a tuple with the ModeratorIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetModeratorIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.ModeratorIds) {
		return nil, false
	}
	return o.ModeratorIds, true
}

// HasModeratorIds returns a boolean if a field has been set.
func (o *ModelGroup) HasModeratorIds() bool {
	if o != nil && !IsNil(o.ModeratorIds) {
		return true
	}

	return false
}

// SetModeratorIds gets a reference to the given []int64 and assigns it to the ModeratorIds field.
func (o *ModelGroup) SetModeratorIds(v []int64) {
	o.ModeratorIds = v
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetOwner() User {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret User
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetOwnerOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelGroup) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableUser and assigns it to the Owner field.
func (o *ModelGroup) SetOwner(v User) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ModelGroup) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ModelGroup) UnsetOwner() {
	o.Owner.Unset()
}

// GetPendingCount returns the PendingCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetPendingCount() int64 {
	if o == nil || IsNil(o.PendingCount.Get()) {
		var ret int64
		return ret
	}
	return *o.PendingCount.Get()
}

// GetPendingCountOk returns a tuple with the PendingCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetPendingCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingCount.Get(), o.PendingCount.IsSet()
}

// HasPendingCount returns a boolean if a field has been set.
func (o *ModelGroup) HasPendingCount() bool {
	if o != nil && o.PendingCount.IsSet() {
		return true
	}

	return false
}

// SetPendingCount gets a reference to the given NullableInt64 and assigns it to the PendingCount field.
func (o *ModelGroup) SetPendingCount(v int64) {
	o.PendingCount.Set(&v)
}
// SetPendingCountNil sets the value for PendingCount to be an explicit nil
func (o *ModelGroup) SetPendingCountNil() {
	o.PendingCount.Set(nil)
}

// UnsetPendingCount ensures that no value is present for PendingCount, not even an explicit nil
func (o *ModelGroup) UnsetPendingCount() {
	o.PendingCount.Unset()
}

// GetPendingDeputizeIds returns the PendingDeputizeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetPendingDeputizeIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.PendingDeputizeIds
}

// GetPendingDeputizeIdsOk returns a tuple with the PendingDeputizeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetPendingDeputizeIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.PendingDeputizeIds) {
		return nil, false
	}
	return o.PendingDeputizeIds, true
}

// HasPendingDeputizeIds returns a boolean if a field has been set.
func (o *ModelGroup) HasPendingDeputizeIds() bool {
	if o != nil && !IsNil(o.PendingDeputizeIds) {
		return true
	}

	return false
}

// SetPendingDeputizeIds gets a reference to the given []int64 and assigns it to the PendingDeputizeIds field.
func (o *ModelGroup) SetPendingDeputizeIds(v []int64) {
	o.PendingDeputizeIds = v
}

// GetPendingTransferId returns the PendingTransferId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetPendingTransferId() int64 {
	if o == nil || IsNil(o.PendingTransferId.Get()) {
		var ret int64
		return ret
	}
	return *o.PendingTransferId.Get()
}

// GetPendingTransferIdOk returns a tuple with the PendingTransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetPendingTransferIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingTransferId.Get(), o.PendingTransferId.IsSet()
}

// HasPendingTransferId returns a boolean if a field has been set.
func (o *ModelGroup) HasPendingTransferId() bool {
	if o != nil && o.PendingTransferId.IsSet() {
		return true
	}

	return false
}

// SetPendingTransferId gets a reference to the given NullableInt64 and assigns it to the PendingTransferId field.
func (o *ModelGroup) SetPendingTransferId(v int64) {
	o.PendingTransferId.Set(&v)
}
// SetPendingTransferIdNil sets the value for PendingTransferId to be an explicit nil
func (o *ModelGroup) SetPendingTransferIdNil() {
	o.PendingTransferId.Set(nil)
}

// UnsetPendingTransferId ensures that no value is present for PendingTransferId, not even an explicit nil
func (o *ModelGroup) UnsetPendingTransferId() {
	o.PendingTransferId.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetPostsCount() int64 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetPostsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *ModelGroup) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt64 and assigns it to the PostsCount field.
func (o *ModelGroup) SetPostsCount(v int64) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *ModelGroup) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *ModelGroup) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetRelatedCount returns the RelatedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetRelatedCount() int64 {
	if o == nil || IsNil(o.RelatedCount.Get()) {
		var ret int64
		return ret
	}
	return *o.RelatedCount.Get()
}

// GetRelatedCountOk returns a tuple with the RelatedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetRelatedCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelatedCount.Get(), o.RelatedCount.IsSet()
}

// HasRelatedCount returns a boolean if a field has been set.
func (o *ModelGroup) HasRelatedCount() bool {
	if o != nil && o.RelatedCount.IsSet() {
		return true
	}

	return false
}

// SetRelatedCount gets a reference to the given NullableInt64 and assigns it to the RelatedCount field.
func (o *ModelGroup) SetRelatedCount(v int64) {
	o.RelatedCount.Set(&v)
}
// SetRelatedCountNil sets the value for RelatedCount to be an explicit nil
func (o *ModelGroup) SetRelatedCountNil() {
	o.RelatedCount.Set(nil)
}

// UnsetRelatedCount ensures that no value is present for RelatedCount, not even an explicit nil
func (o *ModelGroup) UnsetRelatedCount() {
	o.RelatedCount.Unset()
}

// GetSeizableBeforeMillis returns the SeizableBeforeMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetSeizableBeforeMillis() int64 {
	if o == nil || IsNil(o.SeizableBeforeMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.SeizableBeforeMillis.Get()
}

// GetSeizableBeforeMillisOk returns a tuple with the SeizableBeforeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetSeizableBeforeMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SeizableBeforeMillis.Get(), o.SeizableBeforeMillis.IsSet()
}

// HasSeizableBeforeMillis returns a boolean if a field has been set.
func (o *ModelGroup) HasSeizableBeforeMillis() bool {
	if o != nil && o.SeizableBeforeMillis.IsSet() {
		return true
	}

	return false
}

// SetSeizableBeforeMillis gets a reference to the given NullableInt64 and assigns it to the SeizableBeforeMillis field.
func (o *ModelGroup) SetSeizableBeforeMillis(v int64) {
	o.SeizableBeforeMillis.Set(&v)
}
// SetSeizableBeforeMillisNil sets the value for SeizableBeforeMillis to be an explicit nil
func (o *ModelGroup) SetSeizableBeforeMillisNil() {
	o.SeizableBeforeMillis.Set(nil)
}

// UnsetSeizableBeforeMillis ensures that no value is present for SeizableBeforeMillis, not even an explicit nil
func (o *ModelGroup) UnsetSeizableBeforeMillis() {
	o.SeizableBeforeMillis.Unset()
}

// GetSubCategoryId returns the SubCategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetSubCategoryId() int64 {
	if o == nil || IsNil(o.SubCategoryId.Get()) {
		var ret int64
		return ret
	}
	return *o.SubCategoryId.Get()
}

// GetSubCategoryIdOk returns a tuple with the SubCategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetSubCategoryIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubCategoryId.Get(), o.SubCategoryId.IsSet()
}

// HasSubCategoryId returns a boolean if a field has been set.
func (o *ModelGroup) HasSubCategoryId() bool {
	if o != nil && o.SubCategoryId.IsSet() {
		return true
	}

	return false
}

// SetSubCategoryId gets a reference to the given NullableInt64 and assigns it to the SubCategoryId field.
func (o *ModelGroup) SetSubCategoryId(v int64) {
	o.SubCategoryId.Set(&v)
}
// SetSubCategoryIdNil sets the value for SubCategoryId to be an explicit nil
func (o *ModelGroup) SetSubCategoryIdNil() {
	o.SubCategoryId.Set(nil)
}

// UnsetSubCategoryId ensures that no value is present for SubCategoryId, not even an explicit nil
func (o *ModelGroup) UnsetSubCategoryId() {
	o.SubCategoryId.Unset()
}

// GetThreadsCount returns the ThreadsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetThreadsCount() int64 {
	if o == nil || IsNil(o.ThreadsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.ThreadsCount.Get()
}

// GetThreadsCountOk returns a tuple with the ThreadsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetThreadsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadsCount.Get(), o.ThreadsCount.IsSet()
}

// HasThreadsCount returns a boolean if a field has been set.
func (o *ModelGroup) HasThreadsCount() bool {
	if o != nil && o.ThreadsCount.IsSet() {
		return true
	}

	return false
}

// SetThreadsCount gets a reference to the given NullableInt64 and assigns it to the ThreadsCount field.
func (o *ModelGroup) SetThreadsCount(v int64) {
	o.ThreadsCount.Set(&v)
}
// SetThreadsCountNil sets the value for ThreadsCount to be an explicit nil
func (o *ModelGroup) SetThreadsCountNil() {
	o.ThreadsCount.Set(nil)
}

// UnsetThreadsCount ensures that no value is present for ThreadsCount, not even an explicit nil
func (o *ModelGroup) UnsetThreadsCount() {
	o.ThreadsCount.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetTitle() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetTitleOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Title) {
		return map[string]interface{}{}, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModelGroup) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given map[string]interface{} and assigns it to the Title field.
func (o *ModelGroup) SetTitle(v map[string]interface{}) {
	o.Title = v
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetTopic() string {
	if o == nil || IsNil(o.Topic.Get()) {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *ModelGroup) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *ModelGroup) SetTopic(v string) {
	o.Topic.Set(&v)
}
// SetTopicNil sets the value for Topic to be an explicit nil
func (o *ModelGroup) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *ModelGroup) UnsetTopic() {
	o.Topic.Unset()
}

// GetUnreadCounts returns the UnreadCounts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetUnreadCounts() int64 {
	if o == nil || IsNil(o.UnreadCounts.Get()) {
		var ret int64
		return ret
	}
	return *o.UnreadCounts.Get()
}

// GetUnreadCountsOk returns a tuple with the UnreadCounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetUnreadCountsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadCounts.Get(), o.UnreadCounts.IsSet()
}

// HasUnreadCounts returns a boolean if a field has been set.
func (o *ModelGroup) HasUnreadCounts() bool {
	if o != nil && o.UnreadCounts.IsSet() {
		return true
	}

	return false
}

// SetUnreadCounts gets a reference to the given NullableInt64 and assigns it to the UnreadCounts field.
func (o *ModelGroup) SetUnreadCounts(v int64) {
	o.UnreadCounts.Set(&v)
}
// SetUnreadCountsNil sets the value for UnreadCounts to be an explicit nil
func (o *ModelGroup) SetUnreadCountsNil() {
	o.UnreadCounts.Set(nil)
}

// UnsetUnreadCounts ensures that no value is present for UnreadCounts, not even an explicit nil
func (o *ModelGroup) UnsetUnreadCounts() {
	o.UnreadCounts.Unset()
}

// GetUnreadThreadsCount returns the UnreadThreadsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetUnreadThreadsCount() int64 {
	if o == nil || IsNil(o.UnreadThreadsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.UnreadThreadsCount.Get()
}

// GetUnreadThreadsCountOk returns a tuple with the UnreadThreadsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetUnreadThreadsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadThreadsCount.Get(), o.UnreadThreadsCount.IsSet()
}

// HasUnreadThreadsCount returns a boolean if a field has been set.
func (o *ModelGroup) HasUnreadThreadsCount() bool {
	if o != nil && o.UnreadThreadsCount.IsSet() {
		return true
	}

	return false
}

// SetUnreadThreadsCount gets a reference to the given NullableInt64 and assigns it to the UnreadThreadsCount field.
func (o *ModelGroup) SetUnreadThreadsCount(v int64) {
	o.UnreadThreadsCount.Set(&v)
}
// SetUnreadThreadsCountNil sets the value for UnreadThreadsCount to be an explicit nil
func (o *ModelGroup) SetUnreadThreadsCountNil() {
	o.UnreadThreadsCount.Set(nil)
}

// UnsetUnreadThreadsCount ensures that no value is present for UnreadThreadsCount, not even an explicit nil
func (o *ModelGroup) UnsetUnreadThreadsCount() {
	o.UnreadThreadsCount.Unset()
}

// GetUpdatedAtSeconds returns the UpdatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetUpdatedAtSeconds() int64 {
	if o == nil || IsNil(o.UpdatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAtSeconds.Get()
}

// GetUpdatedAtSecondsOk returns a tuple with the UpdatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetUpdatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAtSeconds.Get(), o.UpdatedAtSeconds.IsSet()
}

// HasUpdatedAtSeconds returns a boolean if a field has been set.
func (o *ModelGroup) HasUpdatedAtSeconds() bool {
	if o != nil && o.UpdatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the UpdatedAtSeconds field.
func (o *ModelGroup) SetUpdatedAtSeconds(v int64) {
	o.UpdatedAtSeconds.Set(&v)
}
// SetUpdatedAtSecondsNil sets the value for UpdatedAtSeconds to be an explicit nil
func (o *ModelGroup) SetUpdatedAtSecondsNil() {
	o.UpdatedAtSeconds.Set(nil)
}

// UnsetUpdatedAtSeconds ensures that no value is present for UpdatedAtSeconds, not even an explicit nil
func (o *ModelGroup) UnsetUpdatedAtSeconds() {
	o.UpdatedAtSeconds.Unset()
}

// GetViewsCount returns the ViewsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelGroup) GetViewsCount() int64 {
	if o == nil || IsNil(o.ViewsCount.Get()) {
		var ret int64
		return ret
	}
	return *o.ViewsCount.Get()
}

// GetViewsCountOk returns a tuple with the ViewsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelGroup) GetViewsCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ViewsCount.Get(), o.ViewsCount.IsSet()
}

// HasViewsCount returns a boolean if a field has been set.
func (o *ModelGroup) HasViewsCount() bool {
	if o != nil && o.ViewsCount.IsSet() {
		return true
	}

	return false
}

// SetViewsCount gets a reference to the given NullableInt64 and assigns it to the ViewsCount field.
func (o *ModelGroup) SetViewsCount(v int64) {
	o.ViewsCount.Set(&v)
}
// SetViewsCountNil sets the value for ViewsCount to be an explicit nil
func (o *ModelGroup) SetViewsCountNil() {
	o.ViewsCount.Set(nil)
}

// UnsetViewsCount ensures that no value is present for ViewsCount, not even an explicit nil
func (o *ModelGroup) UnsetViewsCount() {
	o.ViewsCount.Unset()
}

func (o ModelGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowMembersToPostMedia.IsSet() {
		toSerialize["allow_members_to_post_media"] = o.AllowMembersToPostMedia.Get()
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
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if o.CoverImageThumbnail.IsSet() {
		toSerialize["cover_image_thumbnail"] = o.CoverImageThumbnail.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayIconThumbnail.IsSet() {
		toSerialize["display_icon_thumbnail"] = o.DisplayIconThumbnail.Get()
	}
	if o.Gender.IsSet() {
		toSerialize["gender"] = o.Gender.Get()
	}
	if o.Generation.IsSet() {
		toSerialize["generation"] = o.Generation.Get()
	}
	if o.GroupCategoryId.IsSet() {
		toSerialize["group_category_id"] = o.GroupCategoryId.Get()
	}
	if o.GroupsUsersCount.IsSet() {
		toSerialize["groups_users_count"] = o.GroupsUsersCount.Get()
	}
	if o.Guidelines.IsSet() {
		toSerialize["guidelines"] = o.Guidelines.Get()
	}
	if o.HighlightedCount.IsSet() {
		toSerialize["highlighted_count"] = o.HighlightedCount.Get()
	}
	if o.IconImage.IsSet() {
		toSerialize["icon_image"] = o.IconImage.Get()
	}
	if o.IconImageThumbnail.IsSet() {
		toSerialize["icon_image_thumbnail"] = o.IconImageThumbnail.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.InvitedToJoin.IsSet() {
		toSerialize["invited_to_join"] = o.InvitedToJoin.Get()
	}
	if o.IsCallTimelineDisplay.IsSet() {
		toSerialize["is_call_timeline_display"] = o.IsCallTimelineDisplay.Get()
	}
	if o.IsHideConferenceCall.IsSet() {
		toSerialize["is_hide_conference_call"] = o.IsHideConferenceCall.Get()
	}
	if o.IsHideFromGameEight.IsSet() {
		toSerialize["is_hide_from_game_eight"] = o.IsHideFromGameEight.Get()
	}
	if o.IsHideReportedPosts.IsSet() {
		toSerialize["is_hide_reported_posts"] = o.IsHideReportedPosts.Get()
	}
	if o.IsJoined.IsSet() {
		toSerialize["is_joined"] = o.IsJoined.Get()
	}
	if o.IsOnlyVerifiedAge.IsSet() {
		toSerialize["is_only_verified_age"] = o.IsOnlyVerifiedAge.Get()
	}
	if o.IsOnlyVerifiedPhone.IsSet() {
		toSerialize["is_only_verified_phone"] = o.IsOnlyVerifiedPhone.Get()
	}
	if o.IsPending.IsSet() {
		toSerialize["is_pending"] = o.IsPending.Get()
	}
	if o.IsPinned.IsSet() {
		toSerialize["is_pinned"] = o.IsPinned.Get()
	}
	if o.IsPrivate.IsSet() {
		toSerialize["is_private"] = o.IsPrivate.Get()
	}
	if o.IsSafeMode.IsSet() {
		toSerialize["is_safe_mode"] = o.IsSafeMode.Get()
	}
	if o.IsSecret.IsSet() {
		toSerialize["is_secret"] = o.IsSecret.Get()
	}
	if o.IsSeizable.IsSet() {
		toSerialize["is_seizable"] = o.IsSeizable.Get()
	}
	if o.IsWalkthroughRequested.IsSet() {
		toSerialize["is_walkthrough_requested"] = o.IsWalkthroughRequested.Get()
	}
	if o.JoinStatus != nil {
		toSerialize["join_status"] = o.JoinStatus
	}
	if o.JoinedCommunityCampaign.IsSet() {
		toSerialize["joined_community_campaign"] = o.JoinedCommunityCampaign.Get()
	}
	if o.ModeratorIds != nil {
		toSerialize["moderator_ids"] = o.ModeratorIds
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
	if o.SeizableBeforeMillis.IsSet() {
		toSerialize["seizable_before_millis"] = o.SeizableBeforeMillis.Get()
	}
	if o.SubCategoryId.IsSet() {
		toSerialize["sub_category_id"] = o.SubCategoryId.Get()
	}
	if o.ThreadsCount.IsSet() {
		toSerialize["threads_count"] = o.ThreadsCount.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
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
	if o.UpdatedAtSeconds.IsSet() {
		toSerialize["updated_at_seconds"] = o.UpdatedAtSeconds.Get()
	}
	if o.ViewsCount.IsSet() {
		toSerialize["views_count"] = o.ViewsCount.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelGroup) UnmarshalJSON(data []byte) (err error) {
	varModelGroup := _ModelGroup{}

	err = json.Unmarshal(data, &varModelGroup)

	if err != nil {
		return err
	}

	*o = ModelGroup(varModelGroup)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_members_to_post_media")
		delete(additionalProperties, "allow_members_to_post_url")
		delete(additionalProperties, "allow_ownership_transfer")
		delete(additionalProperties, "allow_thread_creation_by")
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "cover_image_thumbnail")
		delete(additionalProperties, "description")
		delete(additionalProperties, "display_icon_thumbnail")
		delete(additionalProperties, "gender")
		delete(additionalProperties, "generation")
		delete(additionalProperties, "group_category_id")
		delete(additionalProperties, "groups_users_count")
		delete(additionalProperties, "guidelines")
		delete(additionalProperties, "highlighted_count")
		delete(additionalProperties, "icon_image")
		delete(additionalProperties, "icon_image_thumbnail")
		delete(additionalProperties, "id")
		delete(additionalProperties, "invited_to_join")
		delete(additionalProperties, "is_call_timeline_display")
		delete(additionalProperties, "is_hide_conference_call")
		delete(additionalProperties, "is_hide_from_game_eight")
		delete(additionalProperties, "is_hide_reported_posts")
		delete(additionalProperties, "is_joined")
		delete(additionalProperties, "is_only_verified_age")
		delete(additionalProperties, "is_only_verified_phone")
		delete(additionalProperties, "is_pending")
		delete(additionalProperties, "is_pinned")
		delete(additionalProperties, "is_private")
		delete(additionalProperties, "is_safe_mode")
		delete(additionalProperties, "is_secret")
		delete(additionalProperties, "is_seizable")
		delete(additionalProperties, "is_walkthrough_requested")
		delete(additionalProperties, "join_status")
		delete(additionalProperties, "joined_community_campaign")
		delete(additionalProperties, "moderator_ids")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "pending_count")
		delete(additionalProperties, "pending_deputize_ids")
		delete(additionalProperties, "pending_transfer_id")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "related_count")
		delete(additionalProperties, "seizable_before_millis")
		delete(additionalProperties, "sub_category_id")
		delete(additionalProperties, "threads_count")
		delete(additionalProperties, "title")
		delete(additionalProperties, "topic")
		delete(additionalProperties, "unread_counts")
		delete(additionalProperties, "unread_threads_count")
		delete(additionalProperties, "updated_at_seconds")
		delete(additionalProperties, "views_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelGroup struct {
	value *ModelGroup
	isSet bool
}

func (v NullableModelGroup) Get() *ModelGroup {
	return v.value
}

func (v *NullableModelGroup) Set(val *ModelGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGroup(val *ModelGroup) *NullableModelGroup {
	return &NullableModelGroup{value: val, isSet: true}
}

func (v NullableModelGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


