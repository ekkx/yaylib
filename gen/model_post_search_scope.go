
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostSearchScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSearchScope{}

// PostSearchScope struct for PostSearchScope
type PostSearchScope struct {
	ApiValue NullableString `json:"api_value,omitempty"`
	TitleId NullableInt32 `json:"title_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostSearchScope PostSearchScope

// NewPostSearchScope instantiates a new PostSearchScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSearchScope() *PostSearchScope {
	this := PostSearchScope{}
	return &this
}

// NewPostSearchScopeWithDefaults instantiates a new PostSearchScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSearchScopeWithDefaults() *PostSearchScope {
	this := PostSearchScope{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchScope) GetApiValue() string {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchScope) GetApiValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *PostSearchScope) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableString and assigns it to the ApiValue field.
func (o *PostSearchScope) SetApiValue(v string) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *PostSearchScope) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *PostSearchScope) UnsetApiValue() {
	o.ApiValue.Unset()
}

// GetTitleId returns the TitleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchScope) GetTitleId() int32 {
	if o == nil || IsNil(o.TitleId.Get()) {
		var ret int32
		return ret
	}
	return *o.TitleId.Get()
}

// GetTitleIdOk returns a tuple with the TitleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchScope) GetTitleIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleId.Get(), o.TitleId.IsSet()
}

// HasTitleId returns a boolean if a field has been set.
func (o *PostSearchScope) HasTitleId() bool {
	if o != nil && o.TitleId.IsSet() {
		return true
	}

	return false
}

// SetTitleId gets a reference to the given NullableInt32 and assigns it to the TitleId field.
func (o *PostSearchScope) SetTitleId(v int32) {
	o.TitleId.Set(&v)
}
// SetTitleIdNil sets the value for TitleId to be an explicit nil
func (o *PostSearchScope) SetTitleIdNil() {
	o.TitleId.Set(nil)
}

// UnsetTitleId ensures that no value is present for TitleId, not even an explicit nil
func (o *PostSearchScope) UnsetTitleId() {
	o.TitleId.Unset()
}

func (o PostSearchScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSearchScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}
	if o.TitleId.IsSet() {
		toSerialize["title_id"] = o.TitleId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostSearchScope) UnmarshalJSON(data []byte) (err error) {
	varPostSearchScope := _PostSearchScope{}

	err = json.Unmarshal(data, &varPostSearchScope)

	if err != nil {
		return err
	}

	*o = PostSearchScope(varPostSearchScope)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		delete(additionalProperties, "title_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostSearchScope struct {
	value *PostSearchScope
	isSet bool
}

func (v NullablePostSearchScope) Get() *PostSearchScope {
	return v.value
}

func (v *NullablePostSearchScope) Set(val *PostSearchScope) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSearchScope) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSearchScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSearchScope(val *PostSearchScope) *NullablePostSearchScope {
	return &NullablePostSearchScope{value: val, isSet: true}
}

func (v NullablePostSearchScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSearchScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


