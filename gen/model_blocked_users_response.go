
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BlockedUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockedUsersResponse{}

// BlockedUsersResponse struct for BlockedUsersResponse
type BlockedUsersResponse struct {
	BlockedCount NullableInt32 `json:"blocked_count,omitempty"`
	LastId NullableInt64 `json:"last_id,omitempty"`
	Users []RealmUser `json:"users,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockedUsersResponse BlockedUsersResponse

// NewBlockedUsersResponse instantiates a new BlockedUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockedUsersResponse() *BlockedUsersResponse {
	this := BlockedUsersResponse{}
	return &this
}

// NewBlockedUsersResponseWithDefaults instantiates a new BlockedUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockedUsersResponseWithDefaults() *BlockedUsersResponse {
	this := BlockedUsersResponse{}
	return &this
}

// GetBlockedCount returns the BlockedCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockedUsersResponse) GetBlockedCount() int32 {
	if o == nil || IsNil(o.BlockedCount.Get()) {
		var ret int32
		return ret
	}
	return *o.BlockedCount.Get()
}

// GetBlockedCountOk returns a tuple with the BlockedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockedUsersResponse) GetBlockedCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.BlockedCount.Get(), o.BlockedCount.IsSet()
}

// HasBlockedCount returns a boolean if a field has been set.
func (o *BlockedUsersResponse) HasBlockedCount() bool {
	if o != nil && o.BlockedCount.IsSet() {
		return true
	}

	return false
}

// SetBlockedCount gets a reference to the given NullableInt32 and assigns it to the BlockedCount field.
func (o *BlockedUsersResponse) SetBlockedCount(v int32) {
	o.BlockedCount.Set(&v)
}
// SetBlockedCountNil sets the value for BlockedCount to be an explicit nil
func (o *BlockedUsersResponse) SetBlockedCountNil() {
	o.BlockedCount.Set(nil)
}

// UnsetBlockedCount ensures that no value is present for BlockedCount, not even an explicit nil
func (o *BlockedUsersResponse) UnsetBlockedCount() {
	o.BlockedCount.Unset()
}

// GetLastId returns the LastId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockedUsersResponse) GetLastId() int64 {
	if o == nil || IsNil(o.LastId.Get()) {
		var ret int64
		return ret
	}
	return *o.LastId.Get()
}

// GetLastIdOk returns a tuple with the LastId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockedUsersResponse) GetLastIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastId.Get(), o.LastId.IsSet()
}

// HasLastId returns a boolean if a field has been set.
func (o *BlockedUsersResponse) HasLastId() bool {
	if o != nil && o.LastId.IsSet() {
		return true
	}

	return false
}

// SetLastId gets a reference to the given NullableInt64 and assigns it to the LastId field.
func (o *BlockedUsersResponse) SetLastId(v int64) {
	o.LastId.Set(&v)
}
// SetLastIdNil sets the value for LastId to be an explicit nil
func (o *BlockedUsersResponse) SetLastIdNil() {
	o.LastId.Set(nil)
}

// UnsetLastId ensures that no value is present for LastId, not even an explicit nil
func (o *BlockedUsersResponse) UnsetLastId() {
	o.LastId.Unset()
}

// GetUsers returns the Users field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockedUsersResponse) GetUsers() []RealmUser {
	if o == nil {
		var ret []RealmUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockedUsersResponse) GetUsersOk() ([]RealmUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BlockedUsersResponse) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []RealmUser and assigns it to the Users field.
func (o *BlockedUsersResponse) SetUsers(v []RealmUser) {
	o.Users = v
}

func (o BlockedUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockedUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockedCount.IsSet() {
		toSerialize["blocked_count"] = o.BlockedCount.Get()
	}
	if o.LastId.IsSet() {
		toSerialize["last_id"] = o.LastId.Get()
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockedUsersResponse) UnmarshalJSON(data []byte) (err error) {
	varBlockedUsersResponse := _BlockedUsersResponse{}

	err = json.Unmarshal(data, &varBlockedUsersResponse)

	if err != nil {
		return err
	}

	*o = BlockedUsersResponse(varBlockedUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "blocked_count")
		delete(additionalProperties, "last_id")
		delete(additionalProperties, "users")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockedUsersResponse struct {
	value *BlockedUsersResponse
	isSet bool
}

func (v NullableBlockedUsersResponse) Get() *BlockedUsersResponse {
	return v.value
}

func (v *NullableBlockedUsersResponse) Set(val *BlockedUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockedUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockedUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockedUsersResponse(val *BlockedUsersResponse) *NullableBlockedUsersResponse {
	return &NullableBlockedUsersResponse{value: val, isSet: true}
}

func (v NullableBlockedUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockedUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


