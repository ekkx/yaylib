
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the BookmarkPostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BookmarkPostResponse{}

// BookmarkPostResponse struct for BookmarkPostResponse
type BookmarkPostResponse struct {
	IsBookmarked NullableBool `json:"is_bookmarked,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BookmarkPostResponse BookmarkPostResponse

// NewBookmarkPostResponse instantiates a new BookmarkPostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBookmarkPostResponse() *BookmarkPostResponse {
	this := BookmarkPostResponse{}
	return &this
}

// NewBookmarkPostResponseWithDefaults instantiates a new BookmarkPostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBookmarkPostResponseWithDefaults() *BookmarkPostResponse {
	this := BookmarkPostResponse{}
	return &this
}

// GetIsBookmarked returns the IsBookmarked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BookmarkPostResponse) GetIsBookmarked() bool {
	if o == nil || IsNil(o.IsBookmarked.Get()) {
		var ret bool
		return ret
	}
	return *o.IsBookmarked.Get()
}

// GetIsBookmarkedOk returns a tuple with the IsBookmarked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BookmarkPostResponse) GetIsBookmarkedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsBookmarked.Get(), o.IsBookmarked.IsSet()
}

// HasIsBookmarked returns a boolean if a field has been set.
func (o *BookmarkPostResponse) HasIsBookmarked() bool {
	if o != nil && o.IsBookmarked.IsSet() {
		return true
	}

	return false
}

// SetIsBookmarked gets a reference to the given NullableBool and assigns it to the IsBookmarked field.
func (o *BookmarkPostResponse) SetIsBookmarked(v bool) {
	o.IsBookmarked.Set(&v)
}
// SetIsBookmarkedNil sets the value for IsBookmarked to be an explicit nil
func (o *BookmarkPostResponse) SetIsBookmarkedNil() {
	o.IsBookmarked.Set(nil)
}

// UnsetIsBookmarked ensures that no value is present for IsBookmarked, not even an explicit nil
func (o *BookmarkPostResponse) UnsetIsBookmarked() {
	o.IsBookmarked.Unset()
}

func (o BookmarkPostResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BookmarkPostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsBookmarked.IsSet() {
		toSerialize["is_bookmarked"] = o.IsBookmarked.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BookmarkPostResponse) UnmarshalJSON(data []byte) (err error) {
	varBookmarkPostResponse := _BookmarkPostResponse{}

	err = json.Unmarshal(data, &varBookmarkPostResponse)

	if err != nil {
		return err
	}

	*o = BookmarkPostResponse(varBookmarkPostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_bookmarked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBookmarkPostResponse struct {
	value *BookmarkPostResponse
	isSet bool
}

func (v NullableBookmarkPostResponse) Get() *BookmarkPostResponse {
	return v.value
}

func (v *NullableBookmarkPostResponse) Set(val *BookmarkPostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBookmarkPostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBookmarkPostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBookmarkPostResponse(val *BookmarkPostResponse) *NullableBookmarkPostResponse {
	return &NullableBookmarkPostResponse{value: val, isSet: true}
}

func (v NullableBookmarkPostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBookmarkPostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


