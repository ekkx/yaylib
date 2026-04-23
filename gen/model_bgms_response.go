
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BgmsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgmsResponse{}

// BgmsResponse struct for BgmsResponse
type BgmsResponse struct {
	Bgm []Bgm `json:"bgm,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgmsResponse BgmsResponse

// NewBgmsResponse instantiates a new BgmsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgmsResponse() *BgmsResponse {
	this := BgmsResponse{}
	return &this
}

// NewBgmsResponseWithDefaults instantiates a new BgmsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgmsResponseWithDefaults() *BgmsResponse {
	this := BgmsResponse{}
	return &this
}

// GetBgm returns the Bgm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BgmsResponse) GetBgm() []Bgm {
	if o == nil {
		var ret []Bgm
		return ret
	}
	return o.Bgm
}

// GetBgmOk returns a tuple with the Bgm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BgmsResponse) GetBgmOk() ([]Bgm, bool) {
	if o == nil || IsNil(o.Bgm) {
		return nil, false
	}
	return o.Bgm, true
}

// HasBgm returns a boolean if a field has been set.
func (o *BgmsResponse) HasBgm() bool {
	if o != nil && !IsNil(o.Bgm) {
		return true
	}

	return false
}

// SetBgm gets a reference to the given []Bgm and assigns it to the Bgm field.
func (o *BgmsResponse) SetBgm(v []Bgm) {
	o.Bgm = v
}

func (o BgmsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgmsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Bgm != nil {
		toSerialize["bgm"] = o.Bgm
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgmsResponse) UnmarshalJSON(data []byte) (err error) {
	varBgmsResponse := _BgmsResponse{}

	err = json.Unmarshal(data, &varBgmsResponse)

	if err != nil {
		return err
	}

	*o = BgmsResponse(varBgmsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "bgm")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgmsResponse struct {
	value *BgmsResponse
	isSet bool
}

func (v NullableBgmsResponse) Get() *BgmsResponse {
	return v.value
}

func (v *NullableBgmsResponse) Set(val *BgmsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBgmsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBgmsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgmsResponse(val *BgmsResponse) *NullableBgmsResponse {
	return &NullableBgmsResponse{value: val, isSet: true}
}

func (v NullableBgmsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgmsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


