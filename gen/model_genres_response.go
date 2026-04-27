
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GenresResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenresResponse{}

// GenresResponse struct for GenresResponse
type GenresResponse struct {
	Genres []RealmGenre `json:"genres,omitempty"`
	NextPageValue NullableInt64 `json:"next_page_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GenresResponse GenresResponse

// NewGenresResponse instantiates a new GenresResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenresResponse() *GenresResponse {
	this := GenresResponse{}
	return &this
}

// NewGenresResponseWithDefaults instantiates a new GenresResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenresResponseWithDefaults() *GenresResponse {
	this := GenresResponse{}
	return &this
}

// GetGenres returns the Genres field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenresResponse) GetGenres() []RealmGenre {
	if o == nil {
		var ret []RealmGenre
		return ret
	}
	return o.Genres
}

// GetGenresOk returns a tuple with the Genres field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenresResponse) GetGenresOk() ([]RealmGenre, bool) {
	if o == nil || IsNil(o.Genres) {
		return nil, false
	}
	return o.Genres, true
}

// HasGenres returns a boolean if a field has been set.
func (o *GenresResponse) HasGenres() bool {
	if o != nil && !IsNil(o.Genres) {
		return true
	}

	return false
}

// SetGenres gets a reference to the given []RealmGenre and assigns it to the Genres field.
func (o *GenresResponse) SetGenres(v []RealmGenre) {
	o.Genres = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GenresResponse) GetNextPageValue() int64 {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret int64
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GenresResponse) GetNextPageValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GenresResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableInt64 and assigns it to the NextPageValue field.
func (o *GenresResponse) SetNextPageValue(v int64) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GenresResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GenresResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

func (o GenresResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenresResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Genres != nil {
		toSerialize["genres"] = o.Genres
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GenresResponse) UnmarshalJSON(data []byte) (err error) {
	varGenresResponse := _GenresResponse{}

	err = json.Unmarshal(data, &varGenresResponse)

	if err != nil {
		return err
	}

	*o = GenresResponse(varGenresResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "genres")
		delete(additionalProperties, "next_page_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGenresResponse struct {
	value *GenresResponse
	isSet bool
}

func (v NullableGenresResponse) Get() *GenresResponse {
	return v.value
}

func (v *NullableGenresResponse) Set(val *GenresResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenresResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenresResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenresResponse(val *GenresResponse) *NullableGenresResponse {
	return &NullableGenresResponse{value: val, isSet: true}
}

func (v NullableGenresResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenresResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


