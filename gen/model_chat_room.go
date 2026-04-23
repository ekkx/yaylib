
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ChatRoom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatRoom{}

// ChatRoom struct for ChatRoom
type ChatRoom struct {
	Background NullableString `json:"background,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsGroup NullableBool `json:"is_group,omitempty"`
	IsNotificationOn NullableBool `json:"is_notification_on,omitempty"`
	IsPinned NullableBool `json:"is_pinned,omitempty"`
	IsRequest NullableBool `json:"is_request,omitempty"`
	LastMessage NullableMessage `json:"last_message,omitempty"`
	Members []User `json:"members,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Owner NullableUser `json:"owner,omitempty"`
	UnreadCount NullableInt32 `json:"unread_count,omitempty"`
	UpdatedAtMillis NullableInt64 `json:"updated_at_millis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChatRoom ChatRoom

// NewChatRoom instantiates a new ChatRoom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatRoom() *ChatRoom {
	this := ChatRoom{}
	return &this
}

// NewChatRoomWithDefaults instantiates a new ChatRoom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatRoomWithDefaults() *ChatRoom {
	this := ChatRoom{}
	return &this
}

// GetBackground returns the Background field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetBackground() string {
	if o == nil || IsNil(o.Background.Get()) {
		var ret string
		return ret
	}
	return *o.Background.Get()
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetBackgroundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Background.Get(), o.Background.IsSet()
}

// HasBackground returns a boolean if a field has been set.
func (o *ChatRoom) HasBackground() bool {
	if o != nil && o.Background.IsSet() {
		return true
	}

	return false
}

// SetBackground gets a reference to the given NullableString and assigns it to the Background field.
func (o *ChatRoom) SetBackground(v string) {
	o.Background.Set(&v)
}
// SetBackgroundNil sets the value for Background to be an explicit nil
func (o *ChatRoom) SetBackgroundNil() {
	o.Background.Set(nil)
}

// UnsetBackground ensures that no value is present for Background, not even an explicit nil
func (o *ChatRoom) UnsetBackground() {
	o.Background.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ChatRoom) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ChatRoom) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ChatRoom) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ChatRoom) UnsetId() {
	o.Id.Unset()
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetIsGroup() bool {
	if o == nil || IsNil(o.IsGroup.Get()) {
		var ret bool
		return ret
	}
	return *o.IsGroup.Get()
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetIsGroupOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGroup.Get(), o.IsGroup.IsSet()
}

// HasIsGroup returns a boolean if a field has been set.
func (o *ChatRoom) HasIsGroup() bool {
	if o != nil && o.IsGroup.IsSet() {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given NullableBool and assigns it to the IsGroup field.
func (o *ChatRoom) SetIsGroup(v bool) {
	o.IsGroup.Set(&v)
}
// SetIsGroupNil sets the value for IsGroup to be an explicit nil
func (o *ChatRoom) SetIsGroupNil() {
	o.IsGroup.Set(nil)
}

// UnsetIsGroup ensures that no value is present for IsGroup, not even an explicit nil
func (o *ChatRoom) UnsetIsGroup() {
	o.IsGroup.Unset()
}

// GetIsNotificationOn returns the IsNotificationOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetIsNotificationOn() bool {
	if o == nil || IsNil(o.IsNotificationOn.Get()) {
		var ret bool
		return ret
	}
	return *o.IsNotificationOn.Get()
}

// GetIsNotificationOnOk returns a tuple with the IsNotificationOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetIsNotificationOnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsNotificationOn.Get(), o.IsNotificationOn.IsSet()
}

// HasIsNotificationOn returns a boolean if a field has been set.
func (o *ChatRoom) HasIsNotificationOn() bool {
	if o != nil && o.IsNotificationOn.IsSet() {
		return true
	}

	return false
}

// SetIsNotificationOn gets a reference to the given NullableBool and assigns it to the IsNotificationOn field.
func (o *ChatRoom) SetIsNotificationOn(v bool) {
	o.IsNotificationOn.Set(&v)
}
// SetIsNotificationOnNil sets the value for IsNotificationOn to be an explicit nil
func (o *ChatRoom) SetIsNotificationOnNil() {
	o.IsNotificationOn.Set(nil)
}

// UnsetIsNotificationOn ensures that no value is present for IsNotificationOn, not even an explicit nil
func (o *ChatRoom) UnsetIsNotificationOn() {
	o.IsNotificationOn.Unset()
}

// GetIsPinned returns the IsPinned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetIsPinned() bool {
	if o == nil || IsNil(o.IsPinned.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPinned.Get()
}

// GetIsPinnedOk returns a tuple with the IsPinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetIsPinnedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPinned.Get(), o.IsPinned.IsSet()
}

// HasIsPinned returns a boolean if a field has been set.
func (o *ChatRoom) HasIsPinned() bool {
	if o != nil && o.IsPinned.IsSet() {
		return true
	}

	return false
}

// SetIsPinned gets a reference to the given NullableBool and assigns it to the IsPinned field.
func (o *ChatRoom) SetIsPinned(v bool) {
	o.IsPinned.Set(&v)
}
// SetIsPinnedNil sets the value for IsPinned to be an explicit nil
func (o *ChatRoom) SetIsPinnedNil() {
	o.IsPinned.Set(nil)
}

// UnsetIsPinned ensures that no value is present for IsPinned, not even an explicit nil
func (o *ChatRoom) UnsetIsPinned() {
	o.IsPinned.Unset()
}

// GetIsRequest returns the IsRequest field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetIsRequest() bool {
	if o == nil || IsNil(o.IsRequest.Get()) {
		var ret bool
		return ret
	}
	return *o.IsRequest.Get()
}

// GetIsRequestOk returns a tuple with the IsRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetIsRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsRequest.Get(), o.IsRequest.IsSet()
}

// HasIsRequest returns a boolean if a field has been set.
func (o *ChatRoom) HasIsRequest() bool {
	if o != nil && o.IsRequest.IsSet() {
		return true
	}

	return false
}

// SetIsRequest gets a reference to the given NullableBool and assigns it to the IsRequest field.
func (o *ChatRoom) SetIsRequest(v bool) {
	o.IsRequest.Set(&v)
}
// SetIsRequestNil sets the value for IsRequest to be an explicit nil
func (o *ChatRoom) SetIsRequestNil() {
	o.IsRequest.Set(nil)
}

// UnsetIsRequest ensures that no value is present for IsRequest, not even an explicit nil
func (o *ChatRoom) UnsetIsRequest() {
	o.IsRequest.Unset()
}

// GetLastMessage returns the LastMessage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetLastMessage() Message {
	if o == nil || IsNil(o.LastMessage.Get()) {
		var ret Message
		return ret
	}
	return *o.LastMessage.Get()
}

// GetLastMessageOk returns a tuple with the LastMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetLastMessageOk() (*Message, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastMessage.Get(), o.LastMessage.IsSet()
}

// HasLastMessage returns a boolean if a field has been set.
func (o *ChatRoom) HasLastMessage() bool {
	if o != nil && o.LastMessage.IsSet() {
		return true
	}

	return false
}

// SetLastMessage gets a reference to the given NullableMessage and assigns it to the LastMessage field.
func (o *ChatRoom) SetLastMessage(v Message) {
	o.LastMessage.Set(&v)
}
// SetLastMessageNil sets the value for LastMessage to be an explicit nil
func (o *ChatRoom) SetLastMessageNil() {
	o.LastMessage.Set(nil)
}

// UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
func (o *ChatRoom) UnsetLastMessage() {
	o.LastMessage.Unset()
}

// GetMembers returns the Members field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetMembers() []User {
	if o == nil {
		var ret []User
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetMembersOk() ([]User, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *ChatRoom) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []User and assigns it to the Members field.
func (o *ChatRoom) SetMembers(v []User) {
	o.Members = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ChatRoom) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ChatRoom) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ChatRoom) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ChatRoom) UnsetName() {
	o.Name.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetOwner() User {
	if o == nil || IsNil(o.Owner.Get()) {
		var ret User
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetOwnerOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *ChatRoom) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableUser and assigns it to the Owner field.
func (o *ChatRoom) SetOwner(v User) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *ChatRoom) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *ChatRoom) UnsetOwner() {
	o.Owner.Unset()
}

// GetUnreadCount returns the UnreadCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetUnreadCount() int32 {
	if o == nil || IsNil(o.UnreadCount.Get()) {
		var ret int32
		return ret
	}
	return *o.UnreadCount.Get()
}

// GetUnreadCountOk returns a tuple with the UnreadCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetUnreadCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.UnreadCount.Get(), o.UnreadCount.IsSet()
}

// HasUnreadCount returns a boolean if a field has been set.
func (o *ChatRoom) HasUnreadCount() bool {
	if o != nil && o.UnreadCount.IsSet() {
		return true
	}

	return false
}

// SetUnreadCount gets a reference to the given NullableInt32 and assigns it to the UnreadCount field.
func (o *ChatRoom) SetUnreadCount(v int32) {
	o.UnreadCount.Set(&v)
}
// SetUnreadCountNil sets the value for UnreadCount to be an explicit nil
func (o *ChatRoom) SetUnreadCountNil() {
	o.UnreadCount.Set(nil)
}

// UnsetUnreadCount ensures that no value is present for UnreadCount, not even an explicit nil
func (o *ChatRoom) UnsetUnreadCount() {
	o.UnreadCount.Unset()
}

// GetUpdatedAtMillis returns the UpdatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChatRoom) GetUpdatedAtMillis() int64 {
	if o == nil || IsNil(o.UpdatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAtMillis.Get()
}

// GetUpdatedAtMillisOk returns a tuple with the UpdatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChatRoom) GetUpdatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAtMillis.Get(), o.UpdatedAtMillis.IsSet()
}

// HasUpdatedAtMillis returns a boolean if a field has been set.
func (o *ChatRoom) HasUpdatedAtMillis() bool {
	if o != nil && o.UpdatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtMillis gets a reference to the given NullableInt64 and assigns it to the UpdatedAtMillis field.
func (o *ChatRoom) SetUpdatedAtMillis(v int64) {
	o.UpdatedAtMillis.Set(&v)
}
// SetUpdatedAtMillisNil sets the value for UpdatedAtMillis to be an explicit nil
func (o *ChatRoom) SetUpdatedAtMillisNil() {
	o.UpdatedAtMillis.Set(nil)
}

// UnsetUpdatedAtMillis ensures that no value is present for UpdatedAtMillis, not even an explicit nil
func (o *ChatRoom) UnsetUpdatedAtMillis() {
	o.UpdatedAtMillis.Unset()
}

func (o ChatRoom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatRoom) ToMap() (map[string]interface{}, error) {
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
	if o.IsNotificationOn.IsSet() {
		toSerialize["is_notification_on"] = o.IsNotificationOn.Get()
	}
	if o.IsPinned.IsSet() {
		toSerialize["is_pinned"] = o.IsPinned.Get()
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
	if o.UpdatedAtMillis.IsSet() {
		toSerialize["updated_at_millis"] = o.UpdatedAtMillis.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ChatRoom) UnmarshalJSON(data []byte) (err error) {
	varChatRoom := _ChatRoom{}

	err = json.Unmarshal(data, &varChatRoom)

	if err != nil {
		return err
	}

	*o = ChatRoom(varChatRoom)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "background")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_group")
		delete(additionalProperties, "is_notification_on")
		delete(additionalProperties, "is_pinned")
		delete(additionalProperties, "is_request")
		delete(additionalProperties, "last_message")
		delete(additionalProperties, "members")
		delete(additionalProperties, "name")
		delete(additionalProperties, "owner")
		delete(additionalProperties, "unread_count")
		delete(additionalProperties, "updated_at_millis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChatRoom struct {
	value *ChatRoom
	isSet bool
}

func (v NullableChatRoom) Get() *ChatRoom {
	return v.value
}

func (v *NullableChatRoom) Set(val *ChatRoom) {
	v.value = val
	v.isSet = true
}

func (v NullableChatRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableChatRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatRoom(val *ChatRoom) *NullableChatRoom {
	return &NullableChatRoom{value: val, isSet: true}
}

func (v NullableChatRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


