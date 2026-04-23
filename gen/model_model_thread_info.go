
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelThreadInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelThreadInfo{}

// ModelThreadInfo struct for ModelThreadInfo
type ModelThreadInfo struct {
	CreatedAtSeconds NullableInt64 `json:"created_at_seconds,omitempty"`
	HasNewUpdates NullableBool `json:"has_new_updates,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsJoined NullableBool `json:"is_joined,omitempty"`
	LastPost NullableModelPost `json:"last_post,omitempty"`
	Owner NullableUser `json:"owner,omitempty"`
	PostsCount NullableInt32 `json:"posts_count,omitempty"`
	ThreadIcon NullableString `json:"thread_icon,omitempty"`
	Title NullableString `json:"title,omitempty"`
	UnreadCount NullableInt32 `json:"unread_count,omitempty"`
	UpdatedAtSeconds NullableInt64 `json:"updated_at_seconds,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelThreadInfo ModelThreadInfo

// NewModelThreadInfo instantiates a new ModelThreadInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelThreadInfo() *ModelThreadInfo {
	this := ModelThreadInfo{}
	return &this
}

// NewModelThreadInfoWithDefaults instantiates a new ModelThreadInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelThreadInfoWithDefaults() *ModelThreadInfo {
	this := ModelThreadInfo{}
	return &this
}

// GetCreatedAtSeconds returns the CreatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetCreatedAtSeconds() int64 {
	if o == nil || IsNil(o.CreatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtSeconds.Get()
}

// GetCreatedAtSecondsOk returns a tuple with the CreatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetCreatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtSeconds.Get(), o.CreatedAtSeconds.IsSet()
}

// HasCreatedAtSeconds returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasCreatedAtSeconds() bool {
	if o != nil && o.CreatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the CreatedAtSeconds field.
func (o *ModelThreadInfo) SetCreatedAtSeconds(v int64) {
	o.CreatedAtSeconds.Set(&v)
}
// SetCreatedAtSecondsNil sets the value for CreatedAtSeconds to be an explicit nil
func (o *ModelThreadInfo) SetCreatedAtSecondsNil() {
	o.CreatedAtSeconds.Set(nil)
}

// UnsetCreatedAtSeconds ensures that no value is present for CreatedAtSeconds, not even an explicit nil
func (o *ModelThreadInfo) UnsetCreatedAtSeconds() {
	o.CreatedAtSeconds.Unset()
}

// GetHasNewUpdates returns the HasNewUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetHasNewUpdates() bool {
	if o == nil || IsNil(o.HasNewUpdates.Get()) {
		var ret bool
		return ret
	}
	return *o.HasNewUpdates.Get()
}

// GetHasNewUpdatesOk returns a tuple with the HasNewUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetHasNewUpdatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HasNewUpdates.Get(), o.HasNewUpdates.IsSet()
}

// HasHasNewUpdates returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasHasNewUpdates() bool {
	if o != nil && o.HasNewUpdates.IsSet() {
		return true
	}

	return false
}

// SetHasNewUpdates gets a reference to the given NullableBool and assigns it to the HasNewUpdates field.
func (o *ModelThreadInfo) SetHasNewUpdates(v bool) {
	o.HasNewUpdates.Set(&v)
}
// SetHasNewUpdatesNil sets the value for HasNewUpdates to be an explicit nil
func (o *ModelThreadInfo) SetHasNewUpdatesNil() {
	o.HasNewUpdates.Set(nil)
}

// UnsetHasNewUpdates ensures that no value is present for HasNewUpdates, not even an explicit nil
func (o *ModelThreadInfo) UnsetHasNewUpdates() {
	o.HasNewUpdates.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelThreadInfo) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelThreadInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelThreadInfo) UnsetId() {
	o.Id.Unset()
}

// GetIsJoined returns the IsJoined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetIsJoined() bool {
	if o == nil || IsNil(o.IsJoined.Get()) {
		var ret bool
		return ret
	}
	return *o.IsJoined.Get()
}

// GetIsJoinedOk returns a tuple with the IsJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetIsJoinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsJoined.Get(), o.IsJoined.IsSet()
}

// HasIsJoined returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasIsJoined() bool {
	if o != nil && o.IsJoined.IsSet() {
		return true
	}

	return false
}

// SetIsJoined gets a reference to the given NullableBool and assigns it to the IsJoined field.
func (o *ModelThreadInfo) SetIsJoined(v bool) {
	o.IsJoined.Set(&v)
}
// SetIsJoinedNil sets the value for IsJoined to be an explicit nil
func (o *ModelThreadInfo) SetIsJoinedNil() {
	o.IsJoined.Set(nil)
}

// UnsetIsJoined ensures that no value is present for IsJoined, not even an explicit nil
func (o *ModelThreadInfo) UnsetIsJoined() {
	o.IsJoined.Unset()
}

// GetLastPost returns the LastPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetLastPost() ModelPost {
	if o == nil || IsNil(o.LastPost.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.LastPost.Get()
}

// GetLastPostOk returns a tuple with the LastPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetLastPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPost.Get(), o.LastPost.IsSet()
}

// HasLastPost returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasLastPost() bool {
	if o != nil && o.LastPost.IsSet() {
		return true
	}

	return false
}

// SetLastPost gets a reference to the given NullableModelPost and assigns it to the LastPost field.
func (o *ModelThreadInfo) SetLastPost(v ModelPost) {
	o.LastPost.Set(&v)
}
// SetLastPostNil sets the value for LastPost to be an explicit nil
func (o *ModelThreadInfo) SetLastPostNil() {
	o.LastPost.Set(nil)
}

// UnsetLastPost ensures that no value is present for LastPost, not even an explicit nil
func (o *ModelThreadInfo) UnsetLastPost() {
	o.LastPost.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetOwner() User {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret User
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetOwnerOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableUser and assigns it to the Owner field.
func (o *ModelThreadInfo) SetOwner(v User) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ModelThreadInfo) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ModelThreadInfo) UnsetOwner() {
	o.Owner.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetPostsCount() int32 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetPostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt32 and assigns it to the PostsCount field.
func (o *ModelThreadInfo) SetPostsCount(v int32) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *ModelThreadInfo) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *ModelThreadInfo) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetThreadIcon returns the ThreadIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetThreadIcon() string {
	if o == nil || IsNil(o.ThreadIcon.Get()) {
		var ret string
		return ret
	}
	return *o.ThreadIcon.Get()
}

// GetThreadIconOk returns a tuple with the ThreadIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetThreadIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadIcon.Get(), o.ThreadIcon.IsSet()
}

// HasThreadIcon returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasThreadIcon() bool {
	if o != nil && o.ThreadIcon.IsSet() {
		return true
	}

	return false
}

// SetThreadIcon gets a reference to the given NullableString and assigns it to the ThreadIcon field.
func (o *ModelThreadInfo) SetThreadIcon(v string) {
	o.ThreadIcon.Set(&v)
}
// SetThreadIconNil sets the value for ThreadIcon to be an explicit nil
func (o *ModelThreadInfo) SetThreadIconNil() {
	o.ThreadIcon.Set(nil)
}

// UnsetThreadIcon ensures that no value is present for ThreadIcon, not even an explicit nil
func (o *ModelThreadInfo) UnsetThreadIcon() {
	o.ThreadIcon.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ModelThreadInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ModelThreadInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ModelThreadInfo) UnsetTitle() {
	o.Title.Unset()
}

// GetUnreadCount returns the UnreadCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetUnreadCount() int32 {
	if o == nil || IsNil(o.UnreadCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnreadCount.Get()
}

// GetUnreadCountOk returns a tuple with the UnreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetUnreadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadCount.Get(), o.UnreadCount.IsSet()
}

// HasUnreadCount returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasUnreadCount() bool {
	if o != nil && o.UnreadCount.IsSet() {
		return true
	}

	return false
}

// SetUnreadCount gets a reference to the given NullableInt32 and assigns it to the UnreadCount field.
func (o *ModelThreadInfo) SetUnreadCount(v int32) {
	o.UnreadCount.Set(&v)
}
// SetUnreadCountNil sets the value for UnreadCount to be an explicit nil
func (o *ModelThreadInfo) SetUnreadCountNil() {
	o.UnreadCount.Set(nil)
}

// UnsetUnreadCount ensures that no value is present for UnreadCount, not even an explicit nil
func (o *ModelThreadInfo) UnsetUnreadCount() {
	o.UnreadCount.Unset()
}

// GetUpdatedAtSeconds returns the UpdatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelThreadInfo) GetUpdatedAtSeconds() int64 {
	if o == nil || IsNil(o.UpdatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAtSeconds.Get()
}

// GetUpdatedAtSecondsOk returns a tuple with the UpdatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelThreadInfo) GetUpdatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAtSeconds.Get(), o.UpdatedAtSeconds.IsSet()
}

// HasUpdatedAtSeconds returns a boolean if a field has been set.
func (o *ModelThreadInfo) HasUpdatedAtSeconds() bool {
	if o != nil && o.UpdatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the UpdatedAtSeconds field.
func (o *ModelThreadInfo) SetUpdatedAtSeconds(v int64) {
	o.UpdatedAtSeconds.Set(&v)
}
// SetUpdatedAtSecondsNil sets the value for UpdatedAtSeconds to be an explicit nil
func (o *ModelThreadInfo) SetUpdatedAtSecondsNil() {
	o.UpdatedAtSeconds.Set(nil)
}

// UnsetUpdatedAtSeconds ensures that no value is present for UpdatedAtSeconds, not even an explicit nil
func (o *ModelThreadInfo) UnsetUpdatedAtSeconds() {
	o.UpdatedAtSeconds.Unset()
}

func (o ModelThreadInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelThreadInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAtSeconds.IsSet() {
		toSerialize["created_at_seconds"] = o.CreatedAtSeconds.Get()
	}
	if o.HasNewUpdates.IsSet() {
		toSerialize["has_new_updates"] = o.HasNewUpdates.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsJoined.IsSet() {
		toSerialize["is_joined"] = o.IsJoined.Get()
	}
	if o.LastPost.IsSet() {
		toSerialize["last_post"] = o.LastPost.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.PostsCount.IsSet() {
		toSerialize["posts_count"] = o.PostsCount.Get()
	}
	if o.ThreadIcon.IsSet() {
		toSerialize["thread_icon"] = o.ThreadIcon.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.UnreadCount.IsSet() {
		toSerialize["unread_count"] = o.UnreadCount.Get()
	}
	if o.UpdatedAtSeconds.IsSet() {
		toSerialize["updated_at_seconds"] = o.UpdatedAtSeconds.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelThreadInfo) UnmarshalJSON(data []byte) (err error) {
	varModelThreadInfo := _ModelThreadInfo{}

	err = json.Unmarshal(data, &varModelThreadInfo)

	if err != nil {
		return err
	}

	*o = ModelThreadInfo(varModelThreadInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at_seconds")
		delete(additionalProperties, "has_new_updates")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_joined")
		delete(additionalProperties, "last_post")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "thread_icon")
		delete(additionalProperties, "title")
		delete(additionalProperties, "unread_count")
		delete(additionalProperties, "updated_at_seconds")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelThreadInfo struct {
	value *ModelThreadInfo
	isSet bool
}

func (v NullableModelThreadInfo) Get() *ModelThreadInfo {
	return v.value
}

func (v *NullableModelThreadInfo) Set(val *ModelThreadInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelThreadInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelThreadInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelThreadInfo(val *ModelThreadInfo) *NullableModelThreadInfo {
	return &NullableModelThreadInfo{value: val, isSet: true}
}

func (v NullableModelThreadInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelThreadInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


