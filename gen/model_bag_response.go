
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BagResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BagResponse{}

// BagResponse struct for BagResponse
type BagResponse struct {
	Result NullableBagDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BagResponse BagResponse

// NewBagResponse instantiates a new BagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBagResponse() *BagResponse {
	this := BagResponse{}
	return &this
}

// NewBagResponseWithDefaults instantiates a new BagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBagResponseWithDefaults() *BagResponse {
	this := BagResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BagResponse) GetResult() BagDTO {
	if o == nil || IsNil(o.Result.Get()) {
		var ret BagDTO
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BagResponse) GetResultOk() (*BagDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *BagResponse) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableBagDTO and assigns it to the Result field.
func (o *BagResponse) SetResult(v BagDTO) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *BagResponse) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *BagResponse) UnsetResult() {
	o.Result.Unset()
}

func (o BagResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BagResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BagResponse) UnmarshalJSON(data []byte) (err error) {
	varBagResponse := _BagResponse{}

	err = json.Unmarshal(data, &varBagResponse)

	if err != nil {
		return err
	}

	*o = BagResponse(varBagResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBagResponse struct {
	value *BagResponse
	isSet bool
}

func (v NullableBagResponse) Get() *BagResponse {
	return v.value
}

func (v *NullableBagResponse) Set(val *BagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBagResponse(val *BagResponse) *NullableBagResponse {
	return &NullableBagResponse{value: val, isSet: true}
}

func (v NullableBagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


