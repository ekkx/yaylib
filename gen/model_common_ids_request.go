
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CommonIdsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonIdsRequest{}

// CommonIdsRequest struct for CommonIdsRequest
type CommonIdsRequest struct {
	Ids []int64 `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonIdsRequest CommonIdsRequest

// NewCommonIdsRequest instantiates a new CommonIdsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonIdsRequest() *CommonIdsRequest {
	this := CommonIdsRequest{}
	return &this
}

// NewCommonIdsRequestWithDefaults instantiates a new CommonIdsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonIdsRequestWithDefaults() *CommonIdsRequest {
	this := CommonIdsRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonIdsRequest) GetIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonIdsRequest) GetIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *CommonIdsRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []int64 and assigns it to the Ids field.
func (o *CommonIdsRequest) SetIds(v []int64) {
	o.Ids = v
}

func (o CommonIdsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonIdsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonIdsRequest) UnmarshalJSON(data []byte) (err error) {
	varCommonIdsRequest := _CommonIdsRequest{}

	err = json.Unmarshal(data, &varCommonIdsRequest)

	if err != nil {
		return err
	}

	*o = CommonIdsRequest(varCommonIdsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonIdsRequest struct {
	value *CommonIdsRequest
	isSet bool
}

func (v NullableCommonIdsRequest) Get() *CommonIdsRequest {
	return v.value
}

func (v *NullableCommonIdsRequest) Set(val *CommonIdsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonIdsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonIdsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonIdsRequest(val *CommonIdsRequest) *NullableCommonIdsRequest {
	return &NullableCommonIdsRequest{value: val, isSet: true}
}

func (v NullableCommonIdsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonIdsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


