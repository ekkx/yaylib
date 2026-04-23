
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostGift type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGift{}

// PostGift struct for PostGift
type PostGift struct {
	Count NullableInt32 `json:"count,omitempty"`
	Gift NullableGift `json:"gift,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostGift PostGift

// NewPostGift instantiates a new PostGift object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGift() *PostGift {
	this := PostGift{}
	return &this
}

// NewPostGiftWithDefaults instantiates a new PostGift object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGiftWithDefaults() *PostGift {
	this := PostGift{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostGift) GetCount() int32 {
	if o == nil || IsNil(o.Count.Get()) {
		var ret int32
		return ret
	}
	return *o.Count.Get()
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostGift) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Count.Get(), o.Count.IsSet()
}

// HasCount returns a boolean if a field has been set.
func (o *PostGift) HasCount() bool {
	if o != nil && o.Count.IsSet() {
		return true
	}

	return false
}

// SetCount gets a reference to the given NullableInt32 and assigns it to the Count field.
func (o *PostGift) SetCount(v int32) {
	o.Count.Set(&v)
}
// SetCountNil sets the value for Count to be an explicit nil
func (o *PostGift) SetCountNil() {
	o.Count.Set(nil)
}

// UnsetCount ensures that no value is present for Count, not even an explicit nil
func (o *PostGift) UnsetCount() {
	o.Count.Unset()
}

// GetGift returns the Gift field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostGift) GetGift() Gift {
	if o == nil || IsNil(o.Gift.Get()) {
		var ret Gift
		return ret
	}
	return *o.Gift.Get()
}

// GetGiftOk returns a tuple with the Gift field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostGift) GetGiftOk() (*Gift, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gift.Get(), o.Gift.IsSet()
}

// HasGift returns a boolean if a field has been set.
func (o *PostGift) HasGift() bool {
	if o != nil && o.Gift.IsSet() {
		return true
	}

	return false
}

// SetGift gets a reference to the given NullableGift and assigns it to the Gift field.
func (o *PostGift) SetGift(v Gift) {
	o.Gift.Set(&v)
}
// SetGiftNil sets the value for Gift to be an explicit nil
func (o *PostGift) SetGiftNil() {
	o.Gift.Set(nil)
}

// UnsetGift ensures that no value is present for Gift, not even an explicit nil
func (o *PostGift) UnsetGift() {
	o.Gift.Unset()
}

func (o PostGift) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGift) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Count.IsSet() {
		toSerialize["count"] = o.Count.Get()
	}
	if o.Gift.IsSet() {
		toSerialize["gift"] = o.Gift.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostGift) UnmarshalJSON(data []byte) (err error) {
	varPostGift := _PostGift{}

	err = json.Unmarshal(data, &varPostGift)

	if err != nil {
		return err
	}

	*o = PostGift(varPostGift)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "count")
		delete(additionalProperties, "gift")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostGift struct {
	value *PostGift
	isSet bool
}

func (v NullablePostGift) Get() *PostGift {
	return v.value
}

func (v *NullablePostGift) Set(val *PostGift) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGift) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGift) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGift(val *PostGift) *NullablePostGift {
	return &NullablePostGift{value: val, isSet: true}
}

func (v NullablePostGift) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGift) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


