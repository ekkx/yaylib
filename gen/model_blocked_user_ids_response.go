
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BlockedUserIdsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlockedUserIdsResponse{}

// BlockedUserIdsResponse struct for BlockedUserIdsResponse
type BlockedUserIdsResponse struct {
	BlockIds []int64 `json:"block_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BlockedUserIdsResponse BlockedUserIdsResponse

// NewBlockedUserIdsResponse instantiates a new BlockedUserIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockedUserIdsResponse() *BlockedUserIdsResponse {
	this := BlockedUserIdsResponse{}
	return &this
}

// NewBlockedUserIdsResponseWithDefaults instantiates a new BlockedUserIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockedUserIdsResponseWithDefaults() *BlockedUserIdsResponse {
	this := BlockedUserIdsResponse{}
	return &this
}

// GetBlockIds returns the BlockIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockedUserIdsResponse) GetBlockIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.BlockIds
}

// GetBlockIdsOk returns a tuple with the BlockIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockedUserIdsResponse) GetBlockIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.BlockIds) {
		return nil, false
	}
	return o.BlockIds, true
}

// HasBlockIds returns a boolean if a field has been set.
func (o *BlockedUserIdsResponse) HasBlockIds() bool {
	if o != nil && !IsNil(o.BlockIds) {
		return true
	}

	return false
}

// SetBlockIds gets a reference to the given []int64 and assigns it to the BlockIds field.
func (o *BlockedUserIdsResponse) SetBlockIds(v []int64) {
	o.BlockIds = v
}

func (o BlockedUserIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlockedUserIdsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BlockIds != nil {
		toSerialize["block_ids"] = o.BlockIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BlockedUserIdsResponse) UnmarshalJSON(data []byte) (err error) {
	varBlockedUserIdsResponse := _BlockedUserIdsResponse{}

	err = json.Unmarshal(data, &varBlockedUserIdsResponse)

	if err != nil {
		return err
	}

	*o = BlockedUserIdsResponse(varBlockedUserIdsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "block_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBlockedUserIdsResponse struct {
	value *BlockedUserIdsResponse
	isSet bool
}

func (v NullableBlockedUserIdsResponse) Get() *BlockedUserIdsResponse {
	return v.value
}

func (v *NullableBlockedUserIdsResponse) Set(val *BlockedUserIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockedUserIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockedUserIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockedUserIdsResponse(val *BlockedUserIdsResponse) *NullableBlockedUserIdsResponse {
	return &NullableBlockedUserIdsResponse{value: val, isSet: true}
}

func (v NullableBlockedUserIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockedUserIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


