
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FootprintsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FootprintsResponse{}

// FootprintsResponse struct for FootprintsResponse
type FootprintsResponse struct {
	Footprints []FootprintDTO `json:"footprints,omitempty"`
	NextPageValue NullableString `json:"next_page_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FootprintsResponse FootprintsResponse

// NewFootprintsResponse instantiates a new FootprintsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFootprintsResponse() *FootprintsResponse {
	this := FootprintsResponse{}
	return &this
}

// NewFootprintsResponseWithDefaults instantiates a new FootprintsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFootprintsResponseWithDefaults() *FootprintsResponse {
	this := FootprintsResponse{}
	return &this
}

// GetFootprints returns the Footprints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FootprintsResponse) GetFootprints() []FootprintDTO {
	if o == nil {
		var ret []FootprintDTO
		return ret
	}
	return o.Footprints
}

// GetFootprintsOk returns a tuple with the Footprints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FootprintsResponse) GetFootprintsOk() ([]FootprintDTO, bool) {
	if o == nil || IsNil(o.Footprints) {
		return nil, false
	}
	return o.Footprints, true
}

// HasFootprints returns a boolean if a field has been set.
func (o *FootprintsResponse) HasFootprints() bool {
	if o != nil && !IsNil(o.Footprints) {
		return true
	}

	return false
}

// SetFootprints gets a reference to the given []FootprintDTO and assigns it to the Footprints field.
func (o *FootprintsResponse) SetFootprints(v []FootprintDTO) {
	o.Footprints = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FootprintsResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FootprintsResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *FootprintsResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *FootprintsResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *FootprintsResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *FootprintsResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

func (o FootprintsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FootprintsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Footprints != nil {
		toSerialize["footprints"] = o.Footprints
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FootprintsResponse) UnmarshalJSON(data []byte) (err error) {
	varFootprintsResponse := _FootprintsResponse{}

	err = json.Unmarshal(data, &varFootprintsResponse)

	if err != nil {
		return err
	}

	*o = FootprintsResponse(varFootprintsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "footprints")
		delete(additionalProperties, "next_page_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFootprintsResponse struct {
	value *FootprintsResponse
	isSet bool
}

func (v NullableFootprintsResponse) Get() *FootprintsResponse {
	return v.value
}

func (v *NullableFootprintsResponse) Set(val *FootprintsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFootprintsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFootprintsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFootprintsResponse(val *FootprintsResponse) *NullableFootprintsResponse {
	return &NullableFootprintsResponse{value: val, isSet: true}
}

func (v NullableFootprintsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFootprintsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


