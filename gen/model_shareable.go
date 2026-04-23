
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Shareable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Shareable{}

// Shareable struct for Shareable
type Shareable struct {
	Group NullableGroup `json:"group,omitempty"`
	Post NullablePost `json:"post,omitempty"`
	Thread NullableThreadInfo `json:"thread,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Shareable Shareable

// NewShareable instantiates a new Shareable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShareable() *Shareable {
	this := Shareable{}
	return &this
}

// NewShareableWithDefaults instantiates a new Shareable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareableWithDefaults() *Shareable {
	this := Shareable{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shareable) GetGroup() Group {
	if o == nil || IsNil(o.Group.Get()) {
		var ret Group
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shareable) GetGroupOk() (*Group, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *Shareable) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableGroup and assigns it to the Group field.
func (o *Shareable) SetGroup(v Group) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *Shareable) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *Shareable) UnsetGroup() {
	o.Group.Unset()
}

// GetPost returns the Post field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shareable) GetPost() Post {
	if o == nil || IsNil(o.Post.Get()) {
		var ret Post
		return ret
	}
	return *o.Post.Get()
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shareable) GetPostOk() (*Post, bool) {
	if o == nil {
		return nil, false
	}
	return o.Post.Get(), o.Post.IsSet()
}

// HasPost returns a boolean if a field has been set.
func (o *Shareable) HasPost() bool {
	if o != nil && o.Post.IsSet() {
		return true
	}

	return false
}

// SetPost gets a reference to the given NullablePost and assigns it to the Post field.
func (o *Shareable) SetPost(v Post) {
	o.Post.Set(&v)
}
// SetPostNil sets the value for Post to be an explicit nil
func (o *Shareable) SetPostNil() {
	o.Post.Set(nil)
}

// UnsetPost ensures that no value is present for Post, not even an explicit nil
func (o *Shareable) UnsetPost() {
	o.Post.Unset()
}

// GetThread returns the Thread field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Shareable) GetThread() ThreadInfo {
	if o == nil || IsNil(o.Thread.Get()) {
		var ret ThreadInfo
		return ret
	}
	return *o.Thread.Get()
}

// GetThreadOk returns a tuple with the Thread field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Shareable) GetThreadOk() (*ThreadInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Thread.Get(), o.Thread.IsSet()
}

// HasThread returns a boolean if a field has been set.
func (o *Shareable) HasThread() bool {
	if o != nil && o.Thread.IsSet() {
		return true
	}

	return false
}

// SetThread gets a reference to the given NullableThreadInfo and assigns it to the Thread field.
func (o *Shareable) SetThread(v ThreadInfo) {
	o.Thread.Set(&v)
}
// SetThreadNil sets the value for Thread to be an explicit nil
func (o *Shareable) SetThreadNil() {
	o.Thread.Set(nil)
}

// UnsetThread ensures that no value is present for Thread, not even an explicit nil
func (o *Shareable) UnsetThread() {
	o.Thread.Unset()
}

func (o Shareable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Shareable) ToMap() (map[string]interface{}, error) {
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

func (o *Shareable) UnmarshalJSON(data []byte) (err error) {
	varShareable := _Shareable{}

	err = json.Unmarshal(data, &varShareable)

	if err != nil {
		return err
	}

	*o = Shareable(varShareable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group")
		delete(additionalProperties, "post")
		delete(additionalProperties, "thread")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableShareable struct {
	value *Shareable
	isSet bool
}

func (v NullableShareable) Get() *Shareable {
	return v.value
}

func (v *NullableShareable) Set(val *Shareable) {
	v.value = val
	v.isSet = true
}

func (v NullableShareable) IsSet() bool {
	return v.isSet
}

func (v *NullableShareable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShareable(val *Shareable) *NullableShareable {
	return &NullableShareable{value: val, isSet: true}
}

func (v NullableShareable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShareable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


