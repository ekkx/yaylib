
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostContentState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostContentState{}

// PostContentState struct for PostContentState
type PostContentState struct {
	PlaceHolderId NullableInt32 `json:"place_holder_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostContentState PostContentState

// NewPostContentState instantiates a new PostContentState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostContentState() *PostContentState {
	this := PostContentState{}
	return &this
}

// NewPostContentStateWithDefaults instantiates a new PostContentState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostContentStateWithDefaults() *PostContentState {
	this := PostContentState{}
	return &this
}

// GetPlaceHolderId returns the PlaceHolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostContentState) GetPlaceHolderId() int32 {
	if o == nil || IsNil(o.PlaceHolderId.Get()) {
		var ret int32
		return ret
	}
	return *o.PlaceHolderId.Get()
}

// GetPlaceHolderIdOk returns a tuple with the PlaceHolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostContentState) GetPlaceHolderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlaceHolderId.Get(), o.PlaceHolderId.IsSet()
}

// HasPlaceHolderId returns a boolean if a field has been set.
func (o *PostContentState) HasPlaceHolderId() bool {
	if o != nil && o.PlaceHolderId.IsSet() {
		return true
	}

	return false
}

// SetPlaceHolderId gets a reference to the given NullableInt32 and assigns it to the PlaceHolderId field.
func (o *PostContentState) SetPlaceHolderId(v int32) {
	o.PlaceHolderId.Set(&v)
}
// SetPlaceHolderIdNil sets the value for PlaceHolderId to be an explicit nil
func (o *PostContentState) SetPlaceHolderIdNil() {
	o.PlaceHolderId.Set(nil)
}

// UnsetPlaceHolderId ensures that no value is present for PlaceHolderId, not even an explicit nil
func (o *PostContentState) UnsetPlaceHolderId() {
	o.PlaceHolderId.Unset()
}

func (o PostContentState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostContentState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PlaceHolderId.IsSet() {
		toSerialize["place_holder_id"] = o.PlaceHolderId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostContentState) UnmarshalJSON(data []byte) (err error) {
	varPostContentState := _PostContentState{}

	err = json.Unmarshal(data, &varPostContentState)

	if err != nil {
		return err
	}

	*o = PostContentState(varPostContentState)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "place_holder_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostContentState struct {
	value *PostContentState
	isSet bool
}

func (v NullablePostContentState) Get() *PostContentState {
	return v.value
}

func (v *NullablePostContentState) Set(val *PostContentState) {
	v.value = val
	v.isSet = true
}

func (v NullablePostContentState) IsSet() bool {
	return v.isSet
}

func (v *NullablePostContentState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostContentState(val *PostContentState) *NullablePostContentState {
	return &NullablePostContentState{value: val, isSet: true}
}

func (v NullablePostContentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostContentState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


