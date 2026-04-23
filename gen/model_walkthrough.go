
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Walkthrough type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Walkthrough{}

// Walkthrough struct for Walkthrough
type Walkthrough struct {
	Title NullableString `json:"title,omitempty"`
	Url NullableString `json:"url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Walkthrough Walkthrough

// NewWalkthrough instantiates a new Walkthrough object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalkthrough() *Walkthrough {
	this := Walkthrough{}
	return &this
}

// NewWalkthroughWithDefaults instantiates a new Walkthrough object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalkthroughWithDefaults() *Walkthrough {
	this := Walkthrough{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Walkthrough) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Walkthrough) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *Walkthrough) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *Walkthrough) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *Walkthrough) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *Walkthrough) UnsetTitle() {
	o.Title.Unset()
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Walkthrough) GetUrl() string {
	if o == nil || IsNil(o.Url.Get()) {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Walkthrough) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *Walkthrough) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *Walkthrough) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *Walkthrough) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *Walkthrough) UnsetUrl() {
	o.Url.Unset()
}

func (o Walkthrough) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Walkthrough) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Walkthrough) UnmarshalJSON(data []byte) (err error) {
	varWalkthrough := _Walkthrough{}

	err = json.Unmarshal(data, &varWalkthrough)

	if err != nil {
		return err
	}

	*o = Walkthrough(varWalkthrough)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWalkthrough struct {
	value *Walkthrough
	isSet bool
}

func (v NullableWalkthrough) Get() *Walkthrough {
	return v.value
}

func (v *NullableWalkthrough) Set(val *Walkthrough) {
	v.value = val
	v.isSet = true
}

func (v NullableWalkthrough) IsSet() bool {
	return v.isSet
}

func (v *NullableWalkthrough) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalkthrough(val *Walkthrough) *NullableWalkthrough {
	return &NullableWalkthrough{value: val, isSet: true}
}

func (v NullableWalkthrough) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalkthrough) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


