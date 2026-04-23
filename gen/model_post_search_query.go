
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostSearchQuery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSearchQuery{}

// PostSearchQuery struct for PostSearchQuery
type PostSearchQuery struct {
	Keyword NullableString `json:"keyword,omitempty"`
	OnlyMedia NullableBool `json:"only_media,omitempty"`
	Scope NullablePostSearchScope `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostSearchQuery PostSearchQuery

// NewPostSearchQuery instantiates a new PostSearchQuery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSearchQuery() *PostSearchQuery {
	this := PostSearchQuery{}
	return &this
}

// NewPostSearchQueryWithDefaults instantiates a new PostSearchQuery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSearchQueryWithDefaults() *PostSearchQuery {
	this := PostSearchQuery{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchQuery) GetKeyword() string {
	if o == nil || IsNil(o.Keyword.Get()) {
		var ret string
		return ret
	}
	return *o.Keyword.Get()
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchQuery) GetKeywordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Keyword.Get(), o.Keyword.IsSet()
}

// HasKeyword returns a boolean if a field has been set.
func (o *PostSearchQuery) HasKeyword() bool {
	if o != nil && o.Keyword.IsSet() {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given NullableString and assigns it to the Keyword field.
func (o *PostSearchQuery) SetKeyword(v string) {
	o.Keyword.Set(&v)
}
// SetKeywordNil sets the value for Keyword to be an explicit nil
func (o *PostSearchQuery) SetKeywordNil() {
	o.Keyword.Set(nil)
}

// UnsetKeyword ensures that no value is present for Keyword, not even an explicit nil
func (o *PostSearchQuery) UnsetKeyword() {
	o.Keyword.Unset()
}

// GetOnlyMedia returns the OnlyMedia field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchQuery) GetOnlyMedia() bool {
	if o == nil || IsNil(o.OnlyMedia.Get()) {
		var ret bool
		return ret
	}
	return *o.OnlyMedia.Get()
}

// GetOnlyMediaOk returns a tuple with the OnlyMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchQuery) GetOnlyMediaOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.OnlyMedia.Get(), o.OnlyMedia.IsSet()
}

// HasOnlyMedia returns a boolean if a field has been set.
func (o *PostSearchQuery) HasOnlyMedia() bool {
	if o != nil && o.OnlyMedia.IsSet() {
		return true
	}

	return false
}

// SetOnlyMedia gets a reference to the given NullableBool and assigns it to the OnlyMedia field.
func (o *PostSearchQuery) SetOnlyMedia(v bool) {
	o.OnlyMedia.Set(&v)
}
// SetOnlyMediaNil sets the value for OnlyMedia to be an explicit nil
func (o *PostSearchQuery) SetOnlyMediaNil() {
	o.OnlyMedia.Set(nil)
}

// UnsetOnlyMedia ensures that no value is present for OnlyMedia, not even an explicit nil
func (o *PostSearchQuery) UnsetOnlyMedia() {
	o.OnlyMedia.Unset()
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostSearchQuery) GetScope() PostSearchScope {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret PostSearchScope
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostSearchQuery) GetScopeOk() (*PostSearchScope, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *PostSearchQuery) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullablePostSearchScope and assigns it to the Scope field.
func (o *PostSearchQuery) SetScope(v PostSearchScope) {
	o.Scope.Set(&v)
}
// SetScopeNil sets the value for Scope to be an explicit nil
func (o *PostSearchQuery) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *PostSearchQuery) UnsetScope() {
	o.Scope.Unset()
}

func (o PostSearchQuery) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSearchQuery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Keyword.IsSet() {
		toSerialize["keyword"] = o.Keyword.Get()
	}
	if o.OnlyMedia.IsSet() {
		toSerialize["only_media"] = o.OnlyMedia.Get()
	}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostSearchQuery) UnmarshalJSON(data []byte) (err error) {
	varPostSearchQuery := _PostSearchQuery{}

	err = json.Unmarshal(data, &varPostSearchQuery)

	if err != nil {
		return err
	}

	*o = PostSearchQuery(varPostSearchQuery)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "keyword")
		delete(additionalProperties, "only_media")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostSearchQuery struct {
	value *PostSearchQuery
	isSet bool
}

func (v NullablePostSearchQuery) Get() *PostSearchQuery {
	return v.value
}

func (v *NullablePostSearchQuery) Set(val *PostSearchQuery) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSearchQuery) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSearchQuery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSearchQuery(val *PostSearchQuery) *NullablePostSearchQuery {
	return &NullablePostSearchQuery{value: val, isSet: true}
}

func (v NullablePostSearchQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSearchQuery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


