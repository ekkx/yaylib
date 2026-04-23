
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GifsDataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GifsDataResponse{}

// GifsDataResponse struct for GifsDataResponse
type GifsDataResponse struct {
	GifCategories []GifImageCategory `json:"gif_categories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GifsDataResponse GifsDataResponse

// NewGifsDataResponse instantiates a new GifsDataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGifsDataResponse() *GifsDataResponse {
	this := GifsDataResponse{}
	return &this
}

// NewGifsDataResponseWithDefaults instantiates a new GifsDataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGifsDataResponseWithDefaults() *GifsDataResponse {
	this := GifsDataResponse{}
	return &this
}

// GetGifCategories returns the GifCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GifsDataResponse) GetGifCategories() []GifImageCategory {
	if o == nil {
		var ret []GifImageCategory
		return ret
	}
	return o.GifCategories
}

// GetGifCategoriesOk returns a tuple with the GifCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GifsDataResponse) GetGifCategoriesOk() ([]GifImageCategory, bool) {
	if o == nil || IsNil(o.GifCategories) {
		return nil, false
	}
	return o.GifCategories, true
}

// HasGifCategories returns a boolean if a field has been set.
func (o *GifsDataResponse) HasGifCategories() bool {
	if o != nil && !IsNil(o.GifCategories) {
		return true
	}

	return false
}

// SetGifCategories gets a reference to the given []GifImageCategory and assigns it to the GifCategories field.
func (o *GifsDataResponse) SetGifCategories(v []GifImageCategory) {
	o.GifCategories = v
}

func (o GifsDataResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GifsDataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GifCategories != nil {
		toSerialize["gif_categories"] = o.GifCategories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GifsDataResponse) UnmarshalJSON(data []byte) (err error) {
	varGifsDataResponse := _GifsDataResponse{}

	err = json.Unmarshal(data, &varGifsDataResponse)

	if err != nil {
		return err
	}

	*o = GifsDataResponse(varGifsDataResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gif_categories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGifsDataResponse struct {
	value *GifsDataResponse
	isSet bool
}

func (v NullableGifsDataResponse) Get() *GifsDataResponse {
	return v.value
}

func (v *NullableGifsDataResponse) Set(val *GifsDataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGifsDataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGifsDataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGifsDataResponse(val *GifsDataResponse) *NullableGifsDataResponse {
	return &NullableGifsDataResponse{value: val, isSet: true}
}

func (v NullableGifsDataResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGifsDataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


