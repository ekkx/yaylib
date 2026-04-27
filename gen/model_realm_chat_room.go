
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmChatRoom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmChatRoom{}

// RealmChatRoom struct for RealmChatRoom
type RealmChatRoom struct {
	Background NullableString `json:"background,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsGroup NullableBool `json:"is_group,omitempty"`
	IsRequest NullableBool `json:"is_request,omitempty"`
	LastMessage NullableChatRoomLastMessage `json:"last_message,omitempty"`
	Members map[string]interface{} `json:"members,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Owner NullableRealmUser `json:"owner,omitempty"`
	UnreadCount NullableInt32 `json:"unread_count,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	UserSetting NullableUserSetting `json:"user_setting,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmChatRoom RealmChatRoom

// NewRealmChatRoom instantiates a new RealmChatRoom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmChatRoom() *RealmChatRoom {
	this := RealmChatRoom{}
	return &this
}

// NewRealmChatRoomWithDefaults instantiates a new RealmChatRoom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmChatRoomWithDefaults() *RealmChatRoom {
	this := RealmChatRoom{}
	return &this
}

// GetBackground returns the Background field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetBackground() string {
	if o == nil || IsNil(o.Background.Get()) {
		var ret string
		return ret
	}
	return *o.Background.Get()
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Background.Get(), o.Background.IsSet()
}

// HasBackground returns a boolean if a field has been set.
func (o *RealmChatRoom) HasBackground() bool {
	if o != nil && o.Background.IsSet() {
		return true
	}

	return false
}

// SetBackground gets a reference to the given NullableString and assigns it to the Background field.
func (o *RealmChatRoom) SetBackground(v string) {
	o.Background.Set(&v)
}
// SetBackgroundNil sets the value for Background to be an explicit nil
func (o *RealmChatRoom) SetBackgroundNil() {
	o.Background.Set(nil)
}

// UnsetBackground ensures that no value is present for Background, not even an explicit nil
func (o *RealmChatRoom) UnsetBackground() {
	o.Background.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RealmChatRoom) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RealmChatRoom) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RealmChatRoom) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RealmChatRoom) UnsetId() {
	o.Id.Unset()
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetIsGroup() bool {
	if o == nil || IsNil(o.IsGroup.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGroup.Get()
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetIsGroupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGroup.Get(), o.IsGroup.IsSet()
}

// HasIsGroup returns a boolean if a field has been set.
func (o *RealmChatRoom) HasIsGroup() bool {
	if o != nil && o.IsGroup.IsSet() {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given NullableBool and assigns it to the IsGroup field.
func (o *RealmChatRoom) SetIsGroup(v bool) {
	o.IsGroup.Set(&v)
}
// SetIsGroupNil sets the value for IsGroup to be an explicit nil
func (o *RealmChatRoom) SetIsGroupNil() {
	o.IsGroup.Set(nil)
}

// UnsetIsGroup ensures that no value is present for IsGroup, not even an explicit nil
func (o *RealmChatRoom) UnsetIsGroup() {
	o.IsGroup.Unset()
}

// GetIsRequest returns the IsRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetIsRequest() bool {
	if o == nil || IsNil(o.IsRequest.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequest.Get()
}

// GetIsRequestOk returns a tuple with the IsRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetIsRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequest.Get(), o.IsRequest.IsSet()
}

// HasIsRequest returns a boolean if a field has been set.
func (o *RealmChatRoom) HasIsRequest() bool {
	if o != nil && o.IsRequest.IsSet() {
		return true
	}

	return false
}

// SetIsRequest gets a reference to the given NullableBool and assigns it to the IsRequest field.
func (o *RealmChatRoom) SetIsRequest(v bool) {
	o.IsRequest.Set(&v)
}
// SetIsRequestNil sets the value for IsRequest to be an explicit nil
func (o *RealmChatRoom) SetIsRequestNil() {
	o.IsRequest.Set(nil)
}

// UnsetIsRequest ensures that no value is present for IsRequest, not even an explicit nil
func (o *RealmChatRoom) UnsetIsRequest() {
	o.IsRequest.Unset()
}

// GetLastMessage returns the LastMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetLastMessage() ChatRoomLastMessage {
	if o == nil || IsNil(o.LastMessage.Get()) {
		var ret ChatRoomLastMessage
		return ret
	}
	return *o.LastMessage.Get()
}

// GetLastMessageOk returns a tuple with the LastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetLastMessageOk() (*ChatRoomLastMessage, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMessage.Get(), o.LastMessage.IsSet()
}

// HasLastMessage returns a boolean if a field has been set.
func (o *RealmChatRoom) HasLastMessage() bool {
	if o != nil && o.LastMessage.IsSet() {
		return true
	}

	return false
}

// SetLastMessage gets a reference to the given NullableChatRoomLastMessage and assigns it to the LastMessage field.
func (o *RealmChatRoom) SetLastMessage(v ChatRoomLastMessage) {
	o.LastMessage.Set(&v)
}
// SetLastMessageNil sets the value for LastMessage to be an explicit nil
func (o *RealmChatRoom) SetLastMessageNil() {
	o.LastMessage.Set(nil)
}

// UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
func (o *RealmChatRoom) UnsetLastMessage() {
	o.LastMessage.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetMembers() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetMembersOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Members) {
		return map[string]interface{}{}, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *RealmChatRoom) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given map[string]interface{} and assigns it to the Members field.
func (o *RealmChatRoom) SetMembers(v map[string]interface{}) {
	o.Members = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *RealmChatRoom) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *RealmChatRoom) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *RealmChatRoom) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *RealmChatRoom) UnsetName() {
	o.Name.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetOwner() RealmUser {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetOwnerOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *RealmChatRoom) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableRealmUser and assigns it to the Owner field.
func (o *RealmChatRoom) SetOwner(v RealmUser) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *RealmChatRoom) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *RealmChatRoom) UnsetOwner() {
	o.Owner.Unset()
}

// GetUnreadCount returns the UnreadCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetUnreadCount() int32 {
	if o == nil || IsNil(o.UnreadCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnreadCount.Get()
}

// GetUnreadCountOk returns a tuple with the UnreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetUnreadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadCount.Get(), o.UnreadCount.IsSet()
}

// HasUnreadCount returns a boolean if a field has been set.
func (o *RealmChatRoom) HasUnreadCount() bool {
	if o != nil && o.UnreadCount.IsSet() {
		return true
	}

	return false
}

// SetUnreadCount gets a reference to the given NullableInt32 and assigns it to the UnreadCount field.
func (o *RealmChatRoom) SetUnreadCount(v int32) {
	o.UnreadCount.Set(&v)
}
// SetUnreadCountNil sets the value for UnreadCount to be an explicit nil
func (o *RealmChatRoom) SetUnreadCountNil() {
	o.UnreadCount.Set(nil)
}

// UnsetUnreadCount ensures that no value is present for UnreadCount, not even an explicit nil
func (o *RealmChatRoom) UnsetUnreadCount() {
	o.UnreadCount.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RealmChatRoom) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *RealmChatRoom) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *RealmChatRoom) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *RealmChatRoom) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUserSetting returns the UserSetting field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmChatRoom) GetUserSetting() UserSetting {
	if o == nil || IsNil(o.UserSetting.Get()) {
		var ret UserSetting
		return ret
	}
	return *o.UserSetting.Get()
}

// GetUserSettingOk returns a tuple with the UserSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmChatRoom) GetUserSettingOk() (*UserSetting, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserSetting.Get(), o.UserSetting.IsSet()
}

// HasUserSetting returns a boolean if a field has been set.
func (o *RealmChatRoom) HasUserSetting() bool {
	if o != nil && o.UserSetting.IsSet() {
		return true
	}

	return false
}

// SetUserSetting gets a reference to the given NullableUserSetting and assigns it to the UserSetting field.
func (o *RealmChatRoom) SetUserSetting(v UserSetting) {
	o.UserSetting.Set(&v)
}
// SetUserSettingNil sets the value for UserSetting to be an explicit nil
func (o *RealmChatRoom) SetUserSettingNil() {
	o.UserSetting.Set(nil)
}

// UnsetUserSetting ensures that no value is present for UserSetting, not even an explicit nil
func (o *RealmChatRoom) UnsetUserSetting() {
	o.UserSetting.Unset()
}

func (o RealmChatRoom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmChatRoom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Background.IsSet() {
		toSerialize["background"] = o.Background.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsGroup.IsSet() {
		toSerialize["is_group"] = o.IsGroup.Get()
	}
	if o.IsRequest.IsSet() {
		toSerialize["is_request"] = o.IsRequest.Get()
	}
	if o.LastMessage.IsSet() {
		toSerialize["last_message"] = o.LastMessage.Get()
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.UnreadCount.IsSet() {
		toSerialize["unread_count"] = o.UnreadCount.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.UserSetting.IsSet() {
		toSerialize["user_setting"] = o.UserSetting.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RealmChatRoom) UnmarshalJSON(data []byte) (err error) {
	varRealmChatRoom := _RealmChatRoom{}

	err = json.Unmarshal(data, &varRealmChatRoom)

	if err != nil {
		return err
	}

	*o = RealmChatRoom(varRealmChatRoom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "background")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_group")
		delete(additionalProperties, "is_request")
		delete(additionalProperties, "last_message")
		delete(additionalProperties, "members")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "unread_count")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "user_setting")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmChatRoom struct {
	value *RealmChatRoom
	isSet bool
}

func (v NullableRealmChatRoom) Get() *RealmChatRoom {
	return v.value
}

func (v *NullableRealmChatRoom) Set(val *RealmChatRoom) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmChatRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmChatRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmChatRoom(val *RealmChatRoom) *NullableRealmChatRoom {
	return &NullableRealmChatRoom{value: val, isSet: true}
}

func (v NullableRealmChatRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmChatRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


