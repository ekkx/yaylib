
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FriendIdsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FriendIdsResponse{}

// FriendIdsResponse struct for FriendIdsResponse
type FriendIdsResponse struct {
	Ids []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FriendIdsResponse FriendIdsResponse

// NewFriendIdsResponse instantiates a new FriendIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFriendIdsResponse() *FriendIdsResponse {
	this := FriendIdsResponse{}
	return &this
}

// NewFriendIdsResponseWithDefaults instantiates a new FriendIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFriendIdsResponseWithDefaults() *FriendIdsResponse {
	this := FriendIdsResponse{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendIdsResponse) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendIdsResponse) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *FriendIdsResponse) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *FriendIdsResponse) SetIds(v []string) {
	o.Ids = v
}

func (o FriendIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FriendIdsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FriendIdsResponse) UnmarshalJSON(data []byte) (err error) {
	varFriendIdsResponse := _FriendIdsResponse{}

	err = json.Unmarshal(data, &varFriendIdsResponse)

	if err != nil {
		return err
	}

	*o = FriendIdsResponse(varFriendIdsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFriendIdsResponse struct {
	value *FriendIdsResponse
	isSet bool
}

func (v NullableFriendIdsResponse) Get() *FriendIdsResponse {
	return v.value
}

func (v *NullableFriendIdsResponse) Set(val *FriendIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFriendIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFriendIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFriendIdsResponse(val *FriendIdsResponse) *NullableFriendIdsResponse {
	return &NullableFriendIdsResponse{value: val, isSet: true}
}

func (v NullableFriendIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFriendIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


