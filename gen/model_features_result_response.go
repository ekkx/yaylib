
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FeaturesResultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeaturesResultResponse{}

// FeaturesResultResponse struct for FeaturesResultResponse
type FeaturesResultResponse struct {
	Features NullableFeaturesDTO `json:"features,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FeaturesResultResponse FeaturesResultResponse

// NewFeaturesResultResponse instantiates a new FeaturesResultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeaturesResultResponse() *FeaturesResultResponse {
	this := FeaturesResultResponse{}
	return &this
}

// NewFeaturesResultResponseWithDefaults instantiates a new FeaturesResultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeaturesResultResponseWithDefaults() *FeaturesResultResponse {
	this := FeaturesResultResponse{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FeaturesResultResponse) GetFeatures() FeaturesDTO {
	if o == nil || IsNil(o.Features.Get()) {
		var ret FeaturesDTO
		return ret
	}
	return *o.Features.Get()
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FeaturesResultResponse) GetFeaturesOk() (*FeaturesDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features.Get(), o.Features.IsSet()
}

// HasFeatures returns a boolean if a field has been set.
func (o *FeaturesResultResponse) HasFeatures() bool {
	if o != nil && o.Features.IsSet() {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given NullableFeaturesDTO and assigns it to the Features field.
func (o *FeaturesResultResponse) SetFeatures(v FeaturesDTO) {
	o.Features.Set(&v)
}
// SetFeaturesNil sets the value for Features to be an explicit nil
func (o *FeaturesResultResponse) SetFeaturesNil() {
	o.Features.Set(nil)
}

// UnsetFeatures ensures that no value is present for Features, not even an explicit nil
func (o *FeaturesResultResponse) UnsetFeatures() {
	o.Features.Unset()
}

func (o FeaturesResultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeaturesResultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Features.IsSet() {
		toSerialize["features"] = o.Features.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FeaturesResultResponse) UnmarshalJSON(data []byte) (err error) {
	varFeaturesResultResponse := _FeaturesResultResponse{}

	err = json.Unmarshal(data, &varFeaturesResultResponse)

	if err != nil {
		return err
	}

	*o = FeaturesResultResponse(varFeaturesResultResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "features")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFeaturesResultResponse struct {
	value *FeaturesResultResponse
	isSet bool
}

func (v NullableFeaturesResultResponse) Get() *FeaturesResultResponse {
	return v.value
}

func (v *NullableFeaturesResultResponse) Set(val *FeaturesResultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFeaturesResultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFeaturesResultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeaturesResultResponse(val *FeaturesResultResponse) *NullableFeaturesResultResponse {
	return &NullableFeaturesResultResponse{value: val, isSet: true}
}

func (v NullableFeaturesResultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeaturesResultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


