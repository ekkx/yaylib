
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupUser{}

// GroupUser struct for GroupUser
type GroupUser struct {
	IsModerator NullableBool `json:"is_moderator,omitempty"`
	IsMute NullableBool `json:"is_mute,omitempty"`
	PendingDeputize NullableBool `json:"pending_deputize,omitempty"`
	PendingTransfer NullableBool `json:"pending_transfer,omitempty"`
	Title NullableString `json:"title,omitempty"`
	User NullableRealmUser `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupUser GroupUser

// NewGroupUser instantiates a new GroupUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUser() *GroupUser {
	this := GroupUser{}
	return &this
}

// NewGroupUserWithDefaults instantiates a new GroupUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUserWithDefaults() *GroupUser {
	this := GroupUser{}
	return &this
}

// GetIsModerator returns the IsModerator field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUser) GetIsModerator() bool {
	if o == nil || IsNil(o.IsModerator.Get()) {
		var ret bool
		return ret
	}
	return *o.IsModerator.Get()
}

// GetIsModeratorOk returns a tuple with the IsModerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetIsModeratorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsModerator.Get(), o.IsModerator.IsSet()
}

// HasIsModerator returns a boolean if a field has been set.
func (o *GroupUser) HasIsModerator() bool {
	if o != nil && o.IsModerator.IsSet() {
		return true
	}

	return false
}

// SetIsModerator gets a reference to the given NullableBool and assigns it to the IsModerator field.
func (o *GroupUser) SetIsModerator(v bool) {
	o.IsModerator.Set(&v)
}
// SetIsModeratorNil sets the value for IsModerator to be an explicit nil
func (o *GroupUser) SetIsModeratorNil() {
	o.IsModerator.Set(nil)
}

// UnsetIsModerator ensures that no value is present for IsModerator, not even an explicit nil
func (o *GroupUser) UnsetIsModerator() {
	o.IsModerator.Unset()
}

// GetIsMute returns the IsMute field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUser) GetIsMute() bool {
	if o == nil || IsNil(o.IsMute.Get()) {
		var ret bool
		return ret
	}
	return *o.IsMute.Get()
}

// GetIsMuteOk returns a tuple with the IsMute field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetIsMuteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMute.Get(), o.IsMute.IsSet()
}

// HasIsMute returns a boolean if a field has been set.
func (o *GroupUser) HasIsMute() bool {
	if o != nil && o.IsMute.IsSet() {
		return true
	}

	return false
}

// SetIsMute gets a reference to the given NullableBool and assigns it to the IsMute field.
func (o *GroupUser) SetIsMute(v bool) {
	o.IsMute.Set(&v)
}
// SetIsMuteNil sets the value for IsMute to be an explicit nil
func (o *GroupUser) SetIsMuteNil() {
	o.IsMute.Set(nil)
}

// UnsetIsMute ensures that no value is present for IsMute, not even an explicit nil
func (o *GroupUser) UnsetIsMute() {
	o.IsMute.Unset()
}

// GetPendingDeputize returns the PendingDeputize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUser) GetPendingDeputize() bool {
	if o == nil || IsNil(o.PendingDeputize.Get()) {
		var ret bool
		return ret
	}
	return *o.PendingDeputize.Get()
}

// GetPendingDeputizeOk returns a tuple with the PendingDeputize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetPendingDeputizeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingDeputize.Get(), o.PendingDeputize.IsSet()
}

// HasPendingDeputize returns a boolean if a field has been set.
func (o *GroupUser) HasPendingDeputize() bool {
	if o != nil && o.PendingDeputize.IsSet() {
		return true
	}

	return false
}

// SetPendingDeputize gets a reference to the given NullableBool and assigns it to the PendingDeputize field.
func (o *GroupUser) SetPendingDeputize(v bool) {
	o.PendingDeputize.Set(&v)
}
// SetPendingDeputizeNil sets the value for PendingDeputize to be an explicit nil
func (o *GroupUser) SetPendingDeputizeNil() {
	o.PendingDeputize.Set(nil)
}

// UnsetPendingDeputize ensures that no value is present for PendingDeputize, not even an explicit nil
func (o *GroupUser) UnsetPendingDeputize() {
	o.PendingDeputize.Unset()
}

// GetPendingTransfer returns the PendingTransfer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUser) GetPendingTransfer() bool {
	if o == nil || IsNil(o.PendingTransfer.Get()) {
		var ret bool
		return ret
	}
	return *o.PendingTransfer.Get()
}

// GetPendingTransferOk returns a tuple with the PendingTransfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetPendingTransferOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PendingTransfer.Get(), o.PendingTransfer.IsSet()
}

// HasPendingTransfer returns a boolean if a field has been set.
func (o *GroupUser) HasPendingTransfer() bool {
	if o != nil && o.PendingTransfer.IsSet() {
		return true
	}

	return false
}

// SetPendingTransfer gets a reference to the given NullableBool and assigns it to the PendingTransfer field.
func (o *GroupUser) SetPendingTransfer(v bool) {
	o.PendingTransfer.Set(&v)
}
// SetPendingTransferNil sets the value for PendingTransfer to be an explicit nil
func (o *GroupUser) SetPendingTransferNil() {
	o.PendingTransfer.Set(nil)
}

// UnsetPendingTransfer ensures that no value is present for PendingTransfer, not even an explicit nil
func (o *GroupUser) UnsetPendingTransfer() {
	o.PendingTransfer.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUser) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *GroupUser) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *GroupUser) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *GroupUser) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *GroupUser) UnsetTitle() {
	o.Title.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupUser) GetUser() RealmUser {
	if o == nil || IsNil(o.User.Get()) {
		var ret RealmUser
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupUser) GetUserOk() (*RealmUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *GroupUser) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableRealmUser and assigns it to the User field.
func (o *GroupUser) SetUser(v RealmUser) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *GroupUser) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *GroupUser) UnsetUser() {
	o.User.Unset()
}

func (o GroupUser) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsModerator.IsSet() {
		toSerialize["is_moderator"] = o.IsModerator.Get()
	}
	if o.IsMute.IsSet() {
		toSerialize["is_mute"] = o.IsMute.Get()
	}
	if o.PendingDeputize.IsSet() {
		toSerialize["pending_deputize"] = o.PendingDeputize.Get()
	}
	if o.PendingTransfer.IsSet() {
		toSerialize["pending_transfer"] = o.PendingTransfer.Get()
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupUser) UnmarshalJSON(data []byte) (err error) {
	varGroupUser := _GroupUser{}

	err = json.Unmarshal(data, &varGroupUser)

	if err != nil {
		return err
	}

	*o = GroupUser(varGroupUser)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_moderator")
		delete(additionalProperties, "is_mute")
		delete(additionalProperties, "pending_deputize")
		delete(additionalProperties, "pending_transfer")
		delete(additionalProperties, "title")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupUser struct {
	value *GroupUser
	isSet bool
}

func (v NullableGroupUser) Get() *GroupUser {
	return v.value
}

func (v *NullableGroupUser) Set(val *GroupUser) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUser) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUser(val *GroupUser) *NullableGroupUser {
	return &NullableGroupUser{value: val, isSet: true}
}

func (v NullableGroupUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


