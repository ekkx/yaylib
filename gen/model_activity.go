
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Activity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Activity{}

// Activity struct for Activity
type Activity struct {
	BirthdayUsers []RealmUser `json:"birthday_users,omitempty"`
	BirthdayUsersCount NullableInt32 `json:"birthday_users_count,omitempty"`
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	EmplAmount NullableInt32 `json:"empl_amount,omitempty"`
	Followers []RealmUser `json:"followers,omitempty"`
	FollowersCount NullableInt32 `json:"followers_count,omitempty"`
	FromPost NullablePost `json:"from_post,omitempty"`
	FromPostIds []int64 `json:"from_post_ids,omitempty"`
	GiftItem NullableRealmGift `json:"gift_item,omitempty"`
	Group NullableGroup `json:"group,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Metadata NullableMetadata `json:"metadata,omitempty"`
	ToPost NullablePost `json:"to_post,omitempty"`
	Type NullableString `json:"type,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	VipReward NullableInt32 `json:"vip_reward,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Activity Activity

// NewActivity instantiates a new Activity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivity() *Activity {
	this := Activity{}
	return &this
}

// NewActivityWithDefaults instantiates a new Activity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityWithDefaults() *Activity {
	this := Activity{}
	return &this
}

// GetBirthdayUsers returns the BirthdayUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetBirthdayUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.BirthdayUsers
}

// GetBirthdayUsersOk returns a tuple with the BirthdayUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetBirthdayUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.BirthdayUsers) {
		return nil, false
	}
	return o.BirthdayUsers, true
}

// HasBirthdayUsers returns a boolean if a field has been set.
func (o *Activity) HasBirthdayUsers() bool {
	if o != nil && !IsNil(o.BirthdayUsers) {
		return true
	}

	return false
}

// SetBirthdayUsers gets a reference to the given []RealmUser and assigns it to the BirthdayUsers field.
func (o *Activity) SetBirthdayUsers(v []RealmUser) {
	o.BirthdayUsers = v
}

// GetBirthdayUsersCount returns the BirthdayUsersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetBirthdayUsersCount() int32 {
	if o == nil || IsNil(o.BirthdayUsersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.BirthdayUsersCount.Get()
}

// GetBirthdayUsersCountOk returns a tuple with the BirthdayUsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetBirthdayUsersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BirthdayUsersCount.Get(), o.BirthdayUsersCount.IsSet()
}

// HasBirthdayUsersCount returns a boolean if a field has been set.
func (o *Activity) HasBirthdayUsersCount() bool {
	if o != nil && o.BirthdayUsersCount.IsSet() {
		return true
	}

	return false
}

// SetBirthdayUsersCount gets a reference to the given NullableInt32 and assigns it to the BirthdayUsersCount field.
func (o *Activity) SetBirthdayUsersCount(v int32) {
	o.BirthdayUsersCount.Set(&v)
}
// SetBirthdayUsersCountNil sets the value for BirthdayUsersCount to be an explicit nil
func (o *Activity) SetBirthdayUsersCountNil() {
	o.BirthdayUsersCount.Set(nil)
}

// UnsetBirthdayUsersCount ensures that no value is present for BirthdayUsersCount, not even an explicit nil
func (o *Activity) UnsetBirthdayUsersCount() {
	o.BirthdayUsersCount.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Activity) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *Activity) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *Activity) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *Activity) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetEmplAmount returns the EmplAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetEmplAmount() int32 {
	if o == nil || IsNil(o.EmplAmount.Get()) {
		var ret int32
		return ret
	}
	return *o.EmplAmount.Get()
}

// GetEmplAmountOk returns a tuple with the EmplAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetEmplAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplAmount.Get(), o.EmplAmount.IsSet()
}

// HasEmplAmount returns a boolean if a field has been set.
func (o *Activity) HasEmplAmount() bool {
	if o != nil && o.EmplAmount.IsSet() {
		return true
	}

	return false
}

// SetEmplAmount gets a reference to the given NullableInt32 and assigns it to the EmplAmount field.
func (o *Activity) SetEmplAmount(v int32) {
	o.EmplAmount.Set(&v)
}
// SetEmplAmountNil sets the value for EmplAmount to be an explicit nil
func (o *Activity) SetEmplAmountNil() {
	o.EmplAmount.Set(nil)
}

// UnsetEmplAmount ensures that no value is present for EmplAmount, not even an explicit nil
func (o *Activity) UnsetEmplAmount() {
	o.EmplAmount.Unset()
}

// GetFollowers returns the Followers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetFollowers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Followers
}

// GetFollowersOk returns a tuple with the Followers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetFollowersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Followers) {
		return nil, false
	}
	return o.Followers, true
}

// HasFollowers returns a boolean if a field has been set.
func (o *Activity) HasFollowers() bool {
	if o != nil && !IsNil(o.Followers) {
		return true
	}

	return false
}

// SetFollowers gets a reference to the given []RealmUser and assigns it to the Followers field.
func (o *Activity) SetFollowers(v []RealmUser) {
	o.Followers = v
}

// GetFollowersCount returns the FollowersCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetFollowersCount() int32 {
	if o == nil || IsNil(o.FollowersCount.Get()) {
		var ret int32
		return ret
	}
	return *o.FollowersCount.Get()
}

// GetFollowersCountOk returns a tuple with the FollowersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetFollowersCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FollowersCount.Get(), o.FollowersCount.IsSet()
}

// HasFollowersCount returns a boolean if a field has been set.
func (o *Activity) HasFollowersCount() bool {
	if o != nil && o.FollowersCount.IsSet() {
		return true
	}

	return false
}

// SetFollowersCount gets a reference to the given NullableInt32 and assigns it to the FollowersCount field.
func (o *Activity) SetFollowersCount(v int32) {
	o.FollowersCount.Set(&v)
}
// SetFollowersCountNil sets the value for FollowersCount to be an explicit nil
func (o *Activity) SetFollowersCountNil() {
	o.FollowersCount.Set(nil)
}

// UnsetFollowersCount ensures that no value is present for FollowersCount, not even an explicit nil
func (o *Activity) UnsetFollowersCount() {
	o.FollowersCount.Unset()
}

// GetFromPost returns the FromPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetFromPost() Post {
	if o == nil || IsNil(o.FromPost.Get()) {
		var ret Post
		return ret
	}
	return *o.FromPost.Get()
}

// GetFromPostOk returns a tuple with the FromPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetFromPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromPost.Get(), o.FromPost.IsSet()
}

// HasFromPost returns a boolean if a field has been set.
func (o *Activity) HasFromPost() bool {
	if o != nil && o.FromPost.IsSet() {
		return true
	}

	return false
}

// SetFromPost gets a reference to the given NullablePost and assigns it to the FromPost field.
func (o *Activity) SetFromPost(v Post) {
	o.FromPost.Set(&v)
}
// SetFromPostNil sets the value for FromPost to be an explicit nil
func (o *Activity) SetFromPostNil() {
	o.FromPost.Set(nil)
}

// UnsetFromPost ensures that no value is present for FromPost, not even an explicit nil
func (o *Activity) UnsetFromPost() {
	o.FromPost.Unset()
}

// GetFromPostIds returns the FromPostIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetFromPostIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.FromPostIds
}

// GetFromPostIdsOk returns a tuple with the FromPostIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetFromPostIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.FromPostIds) {
		return nil, false
	}
	return o.FromPostIds, true
}

// HasFromPostIds returns a boolean if a field has been set.
func (o *Activity) HasFromPostIds() bool {
	if o != nil && !IsNil(o.FromPostIds) {
		return true
	}

	return false
}

// SetFromPostIds gets a reference to the given []int64 and assigns it to the FromPostIds field.
func (o *Activity) SetFromPostIds(v []int64) {
	o.FromPostIds = v
}

// GetGiftItem returns the GiftItem field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetGiftItem() RealmGift {
	if o == nil || IsNil(o.GiftItem.Get()) {
		var ret RealmGift
		return ret
	}
	return *o.GiftItem.Get()
}

// GetGiftItemOk returns a tuple with the GiftItem field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetGiftItemOk() (*RealmGift, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftItem.Get(), o.GiftItem.IsSet()
}

// HasGiftItem returns a boolean if a field has been set.
func (o *Activity) HasGiftItem() bool {
	if o != nil && o.GiftItem.IsSet() {
		return true
	}

	return false
}

// SetGiftItem gets a reference to the given NullableRealmGift and assigns it to the GiftItem field.
func (o *Activity) SetGiftItem(v RealmGift) {
	o.GiftItem.Set(&v)
}
// SetGiftItemNil sets the value for GiftItem to be an explicit nil
func (o *Activity) SetGiftItemNil() {
	o.GiftItem.Set(nil)
}

// UnsetGiftItem ensures that no value is present for GiftItem, not even an explicit nil
func (o *Activity) UnsetGiftItem() {
	o.GiftItem.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetGroup() Group {
	if o == nil || IsNil(o.Group.Get()) {
		var ret Group
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Activity) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableGroup and assigns it to the Group field.
func (o *Activity) SetGroup(v Group) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Activity) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Activity) UnsetGroup() {
	o.Group.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Activity) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Activity) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Activity) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Activity) UnsetId() {
	o.Id.Unset()
}

// GetMetadata returns the Metadata field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetMetadata() Metadata {
	if o == nil || IsNil(o.Metadata.Get()) {
		var ret Metadata
		return ret
	}
	return *o.Metadata.Get()
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetMetadataOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata.Get(), o.Metadata.IsSet()
}

// HasMetadata returns a boolean if a field has been set.
func (o *Activity) HasMetadata() bool {
	if o != nil && o.Metadata.IsSet() {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given NullableMetadata and assigns it to the Metadata field.
func (o *Activity) SetMetadata(v Metadata) {
	o.Metadata.Set(&v)
}
// SetMetadataNil sets the value for Metadata to be an explicit nil
func (o *Activity) SetMetadataNil() {
	o.Metadata.Set(nil)
}

// UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
func (o *Activity) UnsetMetadata() {
	o.Metadata.Unset()
}

// GetToPost returns the ToPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetToPost() Post {
	if o == nil || IsNil(o.ToPost.Get()) {
		var ret Post
		return ret
	}
	return *o.ToPost.Get()
}

// GetToPostOk returns a tuple with the ToPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetToPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToPost.Get(), o.ToPost.IsSet()
}

// HasToPost returns a boolean if a field has been set.
func (o *Activity) HasToPost() bool {
	if o != nil && o.ToPost.IsSet() {
		return true
	}

	return false
}

// SetToPost gets a reference to the given NullablePost and assigns it to the ToPost field.
func (o *Activity) SetToPost(v Post) {
	o.ToPost.Set(&v)
}
// SetToPostNil sets the value for ToPost to be an explicit nil
func (o *Activity) SetToPostNil() {
	o.ToPost.Set(nil)
}

// UnsetToPost ensures that no value is present for ToPost, not even an explicit nil
func (o *Activity) UnsetToPost() {
	o.ToPost.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Activity) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Activity) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Activity) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Activity) UnsetType() {
	o.Type.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *Activity) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *Activity) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *Activity) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *Activity) UnsetUser() {
	o.User.Unset()
}

// GetVipReward returns the VipReward field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Activity) GetVipReward() int32 {
	if o == nil || IsNil(o.VipReward.Get()) {
		var ret int32
		return ret
	}
	return *o.VipReward.Get()
}

// GetVipRewardOk returns a tuple with the VipReward field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Activity) GetVipRewardOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VipReward.Get(), o.VipReward.IsSet()
}

// HasVipReward returns a boolean if a field has been set.
func (o *Activity) HasVipReward() bool {
	if o != nil && o.VipReward.IsSet() {
		return true
	}

	return false
}

// SetVipReward gets a reference to the given NullableInt32 and assigns it to the VipReward field.
func (o *Activity) SetVipReward(v int32) {
	o.VipReward.Set(&v)
}
// SetVipRewardNil sets the value for VipReward to be an explicit nil
func (o *Activity) SetVipRewardNil() {
	o.VipReward.Set(nil)
}

// UnsetVipReward ensures that no value is present for VipReward, not even an explicit nil
func (o *Activity) UnsetVipReward() {
	o.VipReward.Unset()
}

func (o Activity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Activity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BirthdayUsers != nil {
		toSerialize["birthday_users"] = o.BirthdayUsers
	}
	if o.BirthdayUsersCount.IsSet() {
		toSerialize["birthday_users_count"] = o.BirthdayUsersCount.Get()
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
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
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Metadata.IsSet() {
		toSerialize["metadata"] = o.Metadata.Get()
	}
	if o.ToPost.IsSet() {
		toSerialize["to_post"] = o.ToPost.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
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

func (o *Activity) UnmarshalJSON(data []byte) (err error) {
	varActivity := _Activity{}

	err = json.Unmarshal(data, &varActivity)

	if err != nil {
		return err
	}

	*o = Activity(varActivity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "birthday_users")
		delete(additionalProperties, "birthday_users_count")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "empl_amount")
		delete(additionalProperties, "followers")
		delete(additionalProperties, "followers_count")
		delete(additionalProperties, "from_post")
		delete(additionalProperties, "from_post_ids")
		delete(additionalProperties, "gift_item")
		delete(additionalProperties, "group")
		delete(additionalProperties, "id")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "to_post")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user")
		delete(additionalProperties, "vip_reward")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActivity struct {
	value *Activity
	isSet bool
}

func (v NullableActivity) Get() *Activity {
	return v.value
}

func (v *NullableActivity) Set(val *Activity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivity(val *Activity) *NullableActivity {
	return &NullableActivity{value: val, isSet: true}
}

func (v NullableActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


