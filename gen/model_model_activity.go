
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelActivity{}

// ModelActivity struct for ModelActivity
type ModelActivity struct {
	BirthdayUsers []User `json:"birthday_users,omitempty"`
	BirthdayUsersCount NullableInt32 `json:"birthday_users_count,omitempty"`
	Body NullableString `json:"body,omitempty"`
	ContentPreview NullableString `json:"content_preview,omitempty"`
	CreatedAtMillis NullableInt64 `json:"created_at_millis,omitempty"`
	EmplAmount NullableInt32 `json:"empl_amount,omitempty"`
	Followers []User `json:"followers,omitempty"`
	FollowersCount NullableInt32 `json:"followers_count,omitempty"`
	FromPost NullableModelPost `json:"from_post,omitempty"`
	FromPostIds []int64 `json:"from_post_ids,omitempty"`
	GiftItem NullableGift `json:"gift_item,omitempty"`
	Group NullableModelGroup `json:"group,omitempty"`
	GroupTopic NullableString `json:"group_topic,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsAnonymous NullableBool `json:"is_anonymous,omitempty"`
	IsBulkInvitation NullableBool `json:"is_bulk_invitation,omitempty"`
	Title NullableString `json:"title,omitempty"`
	ToPost NullableModelPost `json:"to_post,omitempty"`
	// Unresolved Java type: jp.nanameue.yay.data.model.Activity.Companion.ActivityType
	Type map[string]interface{} `json:"type,omitempty"`
	Url NullableString `json:"url,omitempty"`
	User NullableUser `json:"user,omitempty"`
	VipReward NullableInt32 `json:"vip_reward,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelActivity ModelActivity

// NewModelActivity instantiates a new ModelActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelActivity() *ModelActivity {
	this := ModelActivity{}
	return &this
}

// NewModelActivityWithDefaults instantiates a new ModelActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelActivityWithDefaults() *ModelActivity {
	this := ModelActivity{}
	return &this
}

// GetBirthdayUsers returns the BirthdayUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetBirthdayUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.BirthdayUsers
}

// GetBirthdayUsersOk returns a tuple with the BirthdayUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetBirthdayUsersOk() ([]User, bool) {
	if o == nil || IsNil(o.BirthdayUsers) {
		return nil, false
	}
	return o.BirthdayUsers, true
}

// HasBirthdayUsers returns a boolean if a field has been set.
func (o *ModelActivity) HasBirthdayUsers() bool {
	if o != nil && !IsNil(o.BirthdayUsers) {
		return true
	}

	return false
}

// SetBirthdayUsers gets a reference to the given []User and assigns it to the BirthdayUsers field.
func (o *ModelActivity) SetBirthdayUsers(v []User) {
	o.BirthdayUsers = v
}

// GetBirthdayUsersCount returns the BirthdayUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetBirthdayUsersCount() int32 {
	if o == nil || IsNil(o.BirthdayUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.BirthdayUsersCount.Get()
}

// GetBirthdayUsersCountOk returns a tuple with the BirthdayUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetBirthdayUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthdayUsersCount.Get(), o.BirthdayUsersCount.IsSet()
}

// HasBirthdayUsersCount returns a boolean if a field has been set.
func (o *ModelActivity) HasBirthdayUsersCount() bool {
	if o != nil && o.BirthdayUsersCount.IsSet() {
		return true
	}

	return false
}

// SetBirthdayUsersCount gets a reference to the given NullableInt32 and assigns it to the BirthdayUsersCount field.
func (o *ModelActivity) SetBirthdayUsersCount(v int32) {
	o.BirthdayUsersCount.Set(&v)
}
// SetBirthdayUsersCountNil sets the value for BirthdayUsersCount to be an explicit nil
func (o *ModelActivity) SetBirthdayUsersCountNil() {
	o.BirthdayUsersCount.Set(nil)
}

// UnsetBirthdayUsersCount ensures that no value is present for BirthdayUsersCount, not even an explicit nil
func (o *ModelActivity) UnsetBirthdayUsersCount() {
	o.BirthdayUsersCount.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *ModelActivity) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *ModelActivity) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *ModelActivity) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *ModelActivity) UnsetBody() {
	o.Body.Unset()
}

// GetContentPreview returns the ContentPreview field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetContentPreview() string {
	if o == nil || IsNil(o.ContentPreview.Get()) {
		var ret string
		return ret
	}
	return *o.ContentPreview.Get()
}

// GetContentPreviewOk returns a tuple with the ContentPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetContentPreviewOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentPreview.Get(), o.ContentPreview.IsSet()
}

// HasContentPreview returns a boolean if a field has been set.
func (o *ModelActivity) HasContentPreview() bool {
	if o != nil && o.ContentPreview.IsSet() {
		return true
	}

	return false
}

// SetContentPreview gets a reference to the given NullableString and assigns it to the ContentPreview field.
func (o *ModelActivity) SetContentPreview(v string) {
	o.ContentPreview.Set(&v)
}
// SetContentPreviewNil sets the value for ContentPreview to be an explicit nil
func (o *ModelActivity) SetContentPreviewNil() {
	o.ContentPreview.Set(nil)
}

// UnsetContentPreview ensures that no value is present for ContentPreview, not even an explicit nil
func (o *ModelActivity) UnsetContentPreview() {
	o.ContentPreview.Unset()
}

// GetCreatedAtMillis returns the CreatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetCreatedAtMillis() int64 {
	if o == nil || IsNil(o.CreatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtMillis.Get()
}

// GetCreatedAtMillisOk returns a tuple with the CreatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetCreatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtMillis.Get(), o.CreatedAtMillis.IsSet()
}

// HasCreatedAtMillis returns a boolean if a field has been set.
func (o *ModelActivity) HasCreatedAtMillis() bool {
	if o != nil && o.CreatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtMillis gets a reference to the given NullableInt64 and assigns it to the CreatedAtMillis field.
func (o *ModelActivity) SetCreatedAtMillis(v int64) {
	o.CreatedAtMillis.Set(&v)
}
// SetCreatedAtMillisNil sets the value for CreatedAtMillis to be an explicit nil
func (o *ModelActivity) SetCreatedAtMillisNil() {
	o.CreatedAtMillis.Set(nil)
}

// UnsetCreatedAtMillis ensures that no value is present for CreatedAtMillis, not even an explicit nil
func (o *ModelActivity) UnsetCreatedAtMillis() {
	o.CreatedAtMillis.Unset()
}

// GetEmplAmount returns the EmplAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetEmplAmount() int32 {
	if o == nil || IsNil(o.EmplAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.EmplAmount.Get()
}

// GetEmplAmountOk returns a tuple with the EmplAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetEmplAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplAmount.Get(), o.EmplAmount.IsSet()
}

// HasEmplAmount returns a boolean if a field has been set.
func (o *ModelActivity) HasEmplAmount() bool {
	if o != nil && o.EmplAmount.IsSet() {
		return true
	}

	return false
}

// SetEmplAmount gets a reference to the given NullableInt32 and assigns it to the EmplAmount field.
func (o *ModelActivity) SetEmplAmount(v int32) {
	o.EmplAmount.Set(&v)
}
// SetEmplAmountNil sets the value for EmplAmount to be an explicit nil
func (o *ModelActivity) SetEmplAmountNil() {
	o.EmplAmount.Set(nil)
}

// UnsetEmplAmount ensures that no value is present for EmplAmount, not even an explicit nil
func (o *ModelActivity) UnsetEmplAmount() {
	o.EmplAmount.Unset()
}

// GetFollowers returns the Followers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetFollowers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetFollowersOk() ([]User, bool) {
	if o == nil || IsNil(o.Followers) {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *ModelActivity) HasFollowers() bool {
	if o != nil && !IsNil(o.Followers) {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given []User and assigns it to the Followers field.
func (o *ModelActivity) SetFollowers(v []User) {
	o.Followers = v
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowersCount.Get()
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetFollowersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowersCount.Get(), o.FollowersCount.IsSet()
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *ModelActivity) HasFollowersCount() bool {
	if o != nil && o.FollowersCount.IsSet() {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given NullableInt32 and assigns it to the FollowersCount field.
func (o *ModelActivity) SetFollowersCount(v int32) {
	o.FollowersCount.Set(&v)
}
// SetFollowersCountNil sets the value for FollowersCount to be an explicit nil
func (o *ModelActivity) SetFollowersCountNil() {
	o.FollowersCount.Set(nil)
}

// UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
func (o *ModelActivity) UnsetFollowersCount() {
	o.FollowersCount.Unset()
}

// GetFromPost returns the FromPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetFromPost() ModelPost {
	if o == nil || IsNil(o.FromPost.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.FromPost.Get()
}

// GetFromPostOk returns a tuple with the FromPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetFromPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromPost.Get(), o.FromPost.IsSet()
}

// HasFromPost returns a boolean if a field has been set.
func (o *ModelActivity) HasFromPost() bool {
	if o != nil && o.FromPost.IsSet() {
		return true
	}

	return false
}

// SetFromPost gets a reference to the given NullableModelPost and assigns it to the FromPost field.
func (o *ModelActivity) SetFromPost(v ModelPost) {
	o.FromPost.Set(&v)
}
// SetFromPostNil sets the value for FromPost to be an explicit nil
func (o *ModelActivity) SetFromPostNil() {
	o.FromPost.Set(nil)
}

// UnsetFromPost ensures that no value is present for FromPost, not even an explicit nil
func (o *ModelActivity) UnsetFromPost() {
	o.FromPost.Unset()
}

// GetFromPostIds returns the FromPostIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetFromPostIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.FromPostIds
}

// GetFromPostIdsOk returns a tuple with the FromPostIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetFromPostIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.FromPostIds) {
		return nil, false
	}
	return o.FromPostIds, true
}

// HasFromPostIds returns a boolean if a field has been set.
func (o *ModelActivity) HasFromPostIds() bool {
	if o != nil && !IsNil(o.FromPostIds) {
		return true
	}

	return false
}

// SetFromPostIds gets a reference to the given []int64 and assigns it to the FromPostIds field.
func (o *ModelActivity) SetFromPostIds(v []int64) {
	o.FromPostIds = v
}

// GetGiftItem returns the GiftItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetGiftItem() Gift {
	if o == nil || IsNil(o.GiftItem.Get()) {
		var ret Gift
		return ret
	}
	return *o.GiftItem.Get()
}

// GetGiftItemOk returns a tuple with the GiftItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetGiftItemOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftItem.Get(), o.GiftItem.IsSet()
}

// HasGiftItem returns a boolean if a field has been set.
func (o *ModelActivity) HasGiftItem() bool {
	if o != nil && o.GiftItem.IsSet() {
		return true
	}

	return false
}

// SetGiftItem gets a reference to the given NullableGift and assigns it to the GiftItem field.
func (o *ModelActivity) SetGiftItem(v Gift) {
	o.GiftItem.Set(&v)
}
// SetGiftItemNil sets the value for GiftItem to be an explicit nil
func (o *ModelActivity) SetGiftItemNil() {
	o.GiftItem.Set(nil)
}

// UnsetGiftItem ensures that no value is present for GiftItem, not even an explicit nil
func (o *ModelActivity) UnsetGiftItem() {
	o.GiftItem.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetGroup() ModelGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret ModelGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetGroupOk() (*ModelGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ModelActivity) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableModelGroup and assigns it to the Group field.
func (o *ModelActivity) SetGroup(v ModelGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *ModelActivity) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *ModelActivity) UnsetGroup() {
	o.Group.Unset()
}

// GetGroupTopic returns the GroupTopic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetGroupTopic() string {
	if o == nil || IsNil(o.GroupTopic.Get()) {
		var ret string
		return ret
	}
	return *o.GroupTopic.Get()
}

// GetGroupTopicOk returns a tuple with the GroupTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetGroupTopicOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupTopic.Get(), o.GroupTopic.IsSet()
}

// HasGroupTopic returns a boolean if a field has been set.
func (o *ModelActivity) HasGroupTopic() bool {
	if o != nil && o.GroupTopic.IsSet() {
		return true
	}

	return false
}

// SetGroupTopic gets a reference to the given NullableString and assigns it to the GroupTopic field.
func (o *ModelActivity) SetGroupTopic(v string) {
	o.GroupTopic.Set(&v)
}
// SetGroupTopicNil sets the value for GroupTopic to be an explicit nil
func (o *ModelActivity) SetGroupTopicNil() {
	o.GroupTopic.Set(nil)
}

// UnsetGroupTopic ensures that no value is present for GroupTopic, not even an explicit nil
func (o *ModelActivity) UnsetGroupTopic() {
	o.GroupTopic.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelActivity) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelActivity) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelActivity) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelActivity) UnsetId() {
	o.Id.Unset()
}

// GetIsAnonymous returns the IsAnonymous field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetIsAnonymous() bool {
	if o == nil || IsNil(o.IsAnonymous.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAnonymous.Get()
}

// GetIsAnonymousOk returns a tuple with the IsAnonymous field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetIsAnonymousOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAnonymous.Get(), o.IsAnonymous.IsSet()
}

// HasIsAnonymous returns a boolean if a field has been set.
func (o *ModelActivity) HasIsAnonymous() bool {
	if o != nil && o.IsAnonymous.IsSet() {
		return true
	}

	return false
}

// SetIsAnonymous gets a reference to the given NullableBool and assigns it to the IsAnonymous field.
func (o *ModelActivity) SetIsAnonymous(v bool) {
	o.IsAnonymous.Set(&v)
}
// SetIsAnonymousNil sets the value for IsAnonymous to be an explicit nil
func (o *ModelActivity) SetIsAnonymousNil() {
	o.IsAnonymous.Set(nil)
}

// UnsetIsAnonymous ensures that no value is present for IsAnonymous, not even an explicit nil
func (o *ModelActivity) UnsetIsAnonymous() {
	o.IsAnonymous.Unset()
}

// GetIsBulkInvitation returns the IsBulkInvitation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetIsBulkInvitation() bool {
	if o == nil || IsNil(o.IsBulkInvitation.Get()) {
		var ret bool
		return ret
	}
	return *o.IsBulkInvitation.Get()
}

// GetIsBulkInvitationOk returns a tuple with the IsBulkInvitation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetIsBulkInvitationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBulkInvitation.Get(), o.IsBulkInvitation.IsSet()
}

// HasIsBulkInvitation returns a boolean if a field has been set.
func (o *ModelActivity) HasIsBulkInvitation() bool {
	if o != nil && o.IsBulkInvitation.IsSet() {
		return true
	}

	return false
}

// SetIsBulkInvitation gets a reference to the given NullableBool and assigns it to the IsBulkInvitation field.
func (o *ModelActivity) SetIsBulkInvitation(v bool) {
	o.IsBulkInvitation.Set(&v)
}
// SetIsBulkInvitationNil sets the value for IsBulkInvitation to be an explicit nil
func (o *ModelActivity) SetIsBulkInvitationNil() {
	o.IsBulkInvitation.Set(nil)
}

// UnsetIsBulkInvitation ensures that no value is present for IsBulkInvitation, not even an explicit nil
func (o *ModelActivity) UnsetIsBulkInvitation() {
	o.IsBulkInvitation.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ModelActivity) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ModelActivity) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ModelActivity) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ModelActivity) UnsetTitle() {
	o.Title.Unset()
}

// GetToPost returns the ToPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetToPost() ModelPost {
	if o == nil || IsNil(o.ToPost.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.ToPost.Get()
}

// GetToPostOk returns a tuple with the ToPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetToPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToPost.Get(), o.ToPost.IsSet()
}

// HasToPost returns a boolean if a field has been set.
func (o *ModelActivity) HasToPost() bool {
	if o != nil && o.ToPost.IsSet() {
		return true
	}

	return false
}

// SetToPost gets a reference to the given NullableModelPost and assigns it to the ToPost field.
func (o *ModelActivity) SetToPost(v ModelPost) {
	o.ToPost.Set(&v)
}
// SetToPostNil sets the value for ToPost to be an explicit nil
func (o *ModelActivity) SetToPostNil() {
	o.ToPost.Set(nil)
}

// UnsetToPost ensures that no value is present for ToPost, not even an explicit nil
func (o *ModelActivity) UnsetToPost() {
	o.ToPost.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelActivity) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *ModelActivity) SetType(v map[string]interface{}) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *ModelActivity) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *ModelActivity) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *ModelActivity) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *ModelActivity) UnsetUrl() {
	o.Url.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetUser() User {
	if o == nil || IsNil(o.User.Get()) {
		var ret User
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *ModelActivity) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUser and assigns it to the User field.
func (o *ModelActivity) SetUser(v User) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *ModelActivity) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *ModelActivity) UnsetUser() {
	o.User.Unset()
}

// GetVipReward returns the VipReward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelActivity) GetVipReward() int32 {
	if o == nil || IsNil(o.VipReward.Get()) {
		var ret int32
		return ret
	}
	return *o.VipReward.Get()
}

// GetVipRewardOk returns a tuple with the VipReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelActivity) GetVipRewardOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipReward.Get(), o.VipReward.IsSet()
}

// HasVipReward returns a boolean if a field has been set.
func (o *ModelActivity) HasVipReward() bool {
	if o != nil && o.VipReward.IsSet() {
		return true
	}

	return false
}

// SetVipReward gets a reference to the given NullableInt32 and assigns it to the VipReward field.
func (o *ModelActivity) SetVipReward(v int32) {
	o.VipReward.Set(&v)
}
// SetVipRewardNil sets the value for VipReward to be an explicit nil
func (o *ModelActivity) SetVipRewardNil() {
	o.VipReward.Set(nil)
}

// UnsetVipReward ensures that no value is present for VipReward, not even an explicit nil
func (o *ModelActivity) UnsetVipReward() {
	o.VipReward.Unset()
}

func (o ModelActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BirthdayUsers != nil {
		toSerialize["birthday_users"] = o.BirthdayUsers
	}
	if o.BirthdayUsersCount.IsSet() {
		toSerialize["birthday_users_count"] = o.BirthdayUsersCount.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.ContentPreview.IsSet() {
		toSerialize["content_preview"] = o.ContentPreview.Get()
	}
	if o.CreatedAtMillis.IsSet() {
		toSerialize["created_at_millis"] = o.CreatedAtMillis.Get()
	}
	if o.EmplAmount.IsSet() {
		toSerialize["empl_amount"] = o.EmplAmount.Get()
	}
	if o.Followers != nil {
		toSerialize["followers"] = o.Followers
	}
	if o.FollowersCount.IsSet() {
		toSerialize["followers_count"] = o.FollowersCount.Get()
	}
	if o.FromPost.IsSet() {
		toSerialize["from_post"] = o.FromPost.Get()
	}
	if o.FromPostIds != nil {
		toSerialize["from_post_ids"] = o.FromPostIds
	}
	if o.GiftItem.IsSet() {
		toSerialize["gift_item"] = o.GiftItem.Get()
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.GroupTopic.IsSet() {
		toSerialize["group_topic"] = o.GroupTopic.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsAnonymous.IsSet() {
		toSerialize["is_anonymous"] = o.IsAnonymous.Get()
	}
	if o.IsBulkInvitation.IsSet() {
		toSerialize["is_bulk_invitation"] = o.IsBulkInvitation.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.ToPost.IsSet() {
		toSerialize["to_post"] = o.ToPost.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}
	if o.VipReward.IsSet() {
		toSerialize["vip_reward"] = o.VipReward.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelActivity) UnmarshalJSON(data []byte) (err error) {
	varModelActivity := _ModelActivity{}

	err = json.Unmarshal(data, &varModelActivity)

	if err != nil {
		return err
	}

	*o = ModelActivity(varModelActivity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "birthday_users")
		delete(additionalProperties, "birthday_users_count")
		delete(additionalProperties, "body")
		delete(additionalProperties, "content_preview")
		delete(additionalProperties, "created_at_millis")
		delete(additionalProperties, "empl_amount")
		delete(additionalProperties, "followers")
		delete(additionalProperties, "followers_count")
		delete(additionalProperties, "from_post")
		delete(additionalProperties, "from_post_ids")
		delete(additionalProperties, "gift_item")
		delete(additionalProperties, "group")
		delete(additionalProperties, "group_topic")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_anonymous")
		delete(additionalProperties, "is_bulk_invitation")
		delete(additionalProperties, "title")
		delete(additionalProperties, "to_post")
		delete(additionalProperties, "type")
		delete(additionalProperties, "url")
		delete(additionalProperties, "user")
		delete(additionalProperties, "vip_reward")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelActivity struct {
	value *ModelActivity
	isSet bool
}

func (v NullableModelActivity) Get() *ModelActivity {
	return v.value
}

func (v *NullableModelActivity) Set(val *ModelActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableModelActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableModelActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelActivity(val *ModelActivity) *NullableModelActivity {
	return &NullableModelActivity{value: val, isSet: true}
}

func (v NullableModelActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


