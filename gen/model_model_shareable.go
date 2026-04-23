
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelShareable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelShareable{}

// ModelShareable struct for ModelShareable
type ModelShareable struct {
	Group NullableModelGroup `json:"group,omitempty"`
	Post NullableModelPost `json:"post,omitempty"`
	Thread NullableModelThreadInfo `json:"thread,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelShareable ModelShareable

// NewModelShareable instantiates a new ModelShareable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelShareable() *ModelShareable {
	this := ModelShareable{}
	return &this
}

// NewModelShareableWithDefaults instantiates a new ModelShareable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelShareableWithDefaults() *ModelShareable {
	this := ModelShareable{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelShareable) GetGroup() ModelGroup {
	if o == nil || IsNil(o.Group.Get()) {
		var ret ModelGroup
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelShareable) GetGroupOk() (*ModelGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *ModelShareable) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableModelGroup and assigns it to the Group field.
func (o *ModelShareable) SetGroup(v ModelGroup) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *ModelShareable) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *ModelShareable) UnsetGroup() {
	o.Group.Unset()
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelShareable) GetPost() ModelPost {
	if o == nil || IsNil(o.Post.Get()) {
		var ret ModelPost
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelShareable) GetPostOk() (*ModelPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *ModelShareable) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullableModelPost and assigns it to the Post field.
func (o *ModelShareable) SetPost(v ModelPost) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *ModelShareable) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *ModelShareable) UnsetPost() {
	o.Post.Unset()
}

// GetThread returns the Thread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelShareable) GetThread() ModelThreadInfo {
	if o == nil || IsNil(o.Thread.Get()) {
		var ret ModelThreadInfo
		return ret
	}
	return *o.Thread.Get()
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelShareable) GetThreadOk() (*ModelThreadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thread.Get(), o.Thread.IsSet()
}

// HasThread returns a boolean if a field has been set.
func (o *ModelShareable) HasThread() bool {
	if o != nil && o.Thread.IsSet() {
		return true
	}

	return false
}

// SetThread gets a reference to the given NullableModelThreadInfo and assigns it to the Thread field.
func (o *ModelShareable) SetThread(v ModelThreadInfo) {
	o.Thread.Set(&v)
}
// SetThreadNil sets the value for Thread to be an explicit nil
func (o *ModelShareable) SetThreadNil() {
	o.Thread.Set(nil)
}

// UnsetThread ensures that no value is present for Thread, not even an explicit nil
func (o *ModelShareable) UnsetThread() {
	o.Thread.Unset()
}

func (o ModelShareable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelShareable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	if o.Post.IsSet() {
		toSerialize["post"] = o.Post.Get()
	}
	if o.Thread.IsSet() {
		toSerialize["thread"] = o.Thread.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelShareable) UnmarshalJSON(data []byte) (err error) {
	varModelShareable := _ModelShareable{}

	err = json.Unmarshal(data, &varModelShareable)

	if err != nil {
		return err
	}

	*o = ModelShareable(varModelShareable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "post")
		delete(additionalProperties, "thread")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelShareable struct {
	value *ModelShareable
	isSet bool
}

func (v NullableModelShareable) Get() *ModelShareable {
	return v.value
}

func (v *NullableModelShareable) Set(val *ModelShareable) {
	v.value = val
	v.isSet = true
}

func (v NullableModelShareable) IsSet() bool {
	return v.isSet
}

func (v *NullableModelShareable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelShareable(val *ModelShareable) *NullableModelShareable {
	return &NullableModelShareable{value: val, isSet: true}
}

func (v NullableModelShareable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelShareable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


