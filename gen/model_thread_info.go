
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ThreadInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThreadInfo{}

// ThreadInfo struct for ThreadInfo
type ThreadInfo struct {
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsJoined NullableBool `json:"is_joined,omitempty"`
	LastPost NullablePost `json:"last_post,omitempty"`
	NewUpdates NullableBool `json:"new_updates,omitempty"`
	Owner NullableRealmUser `json:"owner,omitempty"`
	PostsCount NullableInt32 `json:"posts_count,omitempty"`
	ThreadIcon NullableString `json:"thread_icon,omitempty"`
	Title NullableString `json:"title,omitempty"`
	UnreadCount NullableInt32 `json:"unread_count,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ThreadInfo ThreadInfo

// NewThreadInfo instantiates a new ThreadInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThreadInfo() *ThreadInfo {
	this := ThreadInfo{}
	return &this
}

// NewThreadInfoWithDefaults instantiates a new ThreadInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThreadInfoWithDefaults() *ThreadInfo {
	this := ThreadInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ThreadInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *ThreadInfo) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *ThreadInfo) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *ThreadInfo) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ThreadInfo) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ThreadInfo) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ThreadInfo) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ThreadInfo) UnsetId() {
	o.Id.Unset()
}

// GetIsJoined returns the IsJoined field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetIsJoined() bool {
	if o == nil || IsNil(o.IsJoined.Get()) {
		var ret bool
		return ret
	}
	return *o.IsJoined.Get()
}

// GetIsJoinedOk returns a tuple with the IsJoined field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetIsJoinedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsJoined.Get(), o.IsJoined.IsSet()
}

// HasIsJoined returns a boolean if a field has been set.
func (o *ThreadInfo) HasIsJoined() bool {
	if o != nil && o.IsJoined.IsSet() {
		return true
	}

	return false
}

// SetIsJoined gets a reference to the given NullableBool and assigns it to the IsJoined field.
func (o *ThreadInfo) SetIsJoined(v bool) {
	o.IsJoined.Set(&v)
}
// SetIsJoinedNil sets the value for IsJoined to be an explicit nil
func (o *ThreadInfo) SetIsJoinedNil() {
	o.IsJoined.Set(nil)
}

// UnsetIsJoined ensures that no value is present for IsJoined, not even an explicit nil
func (o *ThreadInfo) UnsetIsJoined() {
	o.IsJoined.Unset()
}

// GetLastPost returns the LastPost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetLastPost() Post {
	if o == nil || IsNil(o.LastPost.Get()) {
		var ret Post
		return ret
	}
	return *o.LastPost.Get()
}

// GetLastPostOk returns a tuple with the LastPost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetLastPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastPost.Get(), o.LastPost.IsSet()
}

// HasLastPost returns a boolean if a field has been set.
func (o *ThreadInfo) HasLastPost() bool {
	if o != nil && o.LastPost.IsSet() {
		return true
	}

	return false
}

// SetLastPost gets a reference to the given NullablePost and assigns it to the LastPost field.
func (o *ThreadInfo) SetLastPost(v Post) {
	o.LastPost.Set(&v)
}
// SetLastPostNil sets the value for LastPost to be an explicit nil
func (o *ThreadInfo) SetLastPostNil() {
	o.LastPost.Set(nil)
}

// UnsetLastPost ensures that no value is present for LastPost, not even an explicit nil
func (o *ThreadInfo) UnsetLastPost() {
	o.LastPost.Unset()
}

// GetNewUpdates returns the NewUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetNewUpdates() bool {
	if o == nil || IsNil(o.NewUpdates.Get()) {
		var ret bool
		return ret
	}
	return *o.NewUpdates.Get()
}

// GetNewUpdatesOk returns a tuple with the NewUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetNewUpdatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewUpdates.Get(), o.NewUpdates.IsSet()
}

// HasNewUpdates returns a boolean if a field has been set.
func (o *ThreadInfo) HasNewUpdates() bool {
	if o != nil && o.NewUpdates.IsSet() {
		return true
	}

	return false
}

// SetNewUpdates gets a reference to the given NullableBool and assigns it to the NewUpdates field.
func (o *ThreadInfo) SetNewUpdates(v bool) {
	o.NewUpdates.Set(&v)
}
// SetNewUpdatesNil sets the value for NewUpdates to be an explicit nil
func (o *ThreadInfo) SetNewUpdatesNil() {
	o.NewUpdates.Set(nil)
}

// UnsetNewUpdates ensures that no value is present for NewUpdates, not even an explicit nil
func (o *ThreadInfo) UnsetNewUpdates() {
	o.NewUpdates.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetOwner() RealmUser {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetOwnerOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ThreadInfo) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableRealmUser and assigns it to the Owner field.
func (o *ThreadInfo) SetOwner(v RealmUser) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ThreadInfo) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ThreadInfo) UnsetOwner() {
	o.Owner.Unset()
}

// GetPostsCount returns the PostsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetPostsCount() int32 {
	if o == nil || IsNil(o.PostsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PostsCount.Get()
}

// GetPostsCountOk returns a tuple with the PostsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetPostsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostsCount.Get(), o.PostsCount.IsSet()
}

// HasPostsCount returns a boolean if a field has been set.
func (o *ThreadInfo) HasPostsCount() bool {
	if o != nil && o.PostsCount.IsSet() {
		return true
	}

	return false
}

// SetPostsCount gets a reference to the given NullableInt32 and assigns it to the PostsCount field.
func (o *ThreadInfo) SetPostsCount(v int32) {
	o.PostsCount.Set(&v)
}
// SetPostsCountNil sets the value for PostsCount to be an explicit nil
func (o *ThreadInfo) SetPostsCountNil() {
	o.PostsCount.Set(nil)
}

// UnsetPostsCount ensures that no value is present for PostsCount, not even an explicit nil
func (o *ThreadInfo) UnsetPostsCount() {
	o.PostsCount.Unset()
}

// GetThreadIcon returns the ThreadIcon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetThreadIcon() string {
	if o == nil || IsNil(o.ThreadIcon.Get()) {
		var ret string
		return ret
	}
	return *o.ThreadIcon.Get()
}

// GetThreadIconOk returns a tuple with the ThreadIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetThreadIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreadIcon.Get(), o.ThreadIcon.IsSet()
}

// HasThreadIcon returns a boolean if a field has been set.
func (o *ThreadInfo) HasThreadIcon() bool {
	if o != nil && o.ThreadIcon.IsSet() {
		return true
	}

	return false
}

// SetThreadIcon gets a reference to the given NullableString and assigns it to the ThreadIcon field.
func (o *ThreadInfo) SetThreadIcon(v string) {
	o.ThreadIcon.Set(&v)
}
// SetThreadIconNil sets the value for ThreadIcon to be an explicit nil
func (o *ThreadInfo) SetThreadIconNil() {
	o.ThreadIcon.Set(nil)
}

// UnsetThreadIcon ensures that no value is present for ThreadIcon, not even an explicit nil
func (o *ThreadInfo) UnsetThreadIcon() {
	o.ThreadIcon.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *ThreadInfo) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *ThreadInfo) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *ThreadInfo) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *ThreadInfo) UnsetTitle() {
	o.Title.Unset()
}

// GetUnreadCount returns the UnreadCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetUnreadCount() int32 {
	if o == nil || IsNil(o.UnreadCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnreadCount.Get()
}

// GetUnreadCountOk returns a tuple with the UnreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetUnreadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadCount.Get(), o.UnreadCount.IsSet()
}

// HasUnreadCount returns a boolean if a field has been set.
func (o *ThreadInfo) HasUnreadCount() bool {
	if o != nil && o.UnreadCount.IsSet() {
		return true
	}

	return false
}

// SetUnreadCount gets a reference to the given NullableInt32 and assigns it to the UnreadCount field.
func (o *ThreadInfo) SetUnreadCount(v int32) {
	o.UnreadCount.Set(&v)
}
// SetUnreadCountNil sets the value for UnreadCount to be an explicit nil
func (o *ThreadInfo) SetUnreadCountNil() {
	o.UnreadCount.Set(nil)
}

// UnsetUnreadCount ensures that no value is present for UnreadCount, not even an explicit nil
func (o *ThreadInfo) UnsetUnreadCount() {
	o.UnreadCount.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ThreadInfo) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ThreadInfo) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ThreadInfo) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *ThreadInfo) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *ThreadInfo) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *ThreadInfo) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o ThreadInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ThreadInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
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
	if o.NewUpdates.IsSet() {
		toSerialize["new_updates"] = o.NewUpdates.Get()
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
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ThreadInfo) UnmarshalJSON(data []byte) (err error) {
	varThreadInfo := _ThreadInfo{}

	err = json.Unmarshal(data, &varThreadInfo)

	if err != nil {
		return err
	}

	*o = ThreadInfo(varThreadInfo)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_joined")
		delete(additionalProperties, "last_post")
		delete(additionalProperties, "new_updates")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "posts_count")
		delete(additionalProperties, "thread_icon")
		delete(additionalProperties, "title")
		delete(additionalProperties, "unread_count")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableThreadInfo struct {
	value *ThreadInfo
	isSet bool
}

func (v NullableThreadInfo) Get() *ThreadInfo {
	return v.value
}

func (v *NullableThreadInfo) Set(val *ThreadInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableThreadInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableThreadInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThreadInfo(val *ThreadInfo) *NullableThreadInfo {
	return &NullableThreadInfo{value: val, isSet: true}
}

func (v NullableThreadInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThreadInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


