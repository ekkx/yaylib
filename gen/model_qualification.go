
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Qualification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Qualification{}

// Qualification struct for Qualification
type Qualification struct {
	Qualified NullableBool `json:"qualified,omitempty"`
	Type map[string]interface{} `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Qualification Qualification

// NewQualification instantiates a new Qualification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQualification() *Qualification {
	this := Qualification{}
	return &this
}

// NewQualificationWithDefaults instantiates a new Qualification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQualificationWithDefaults() *Qualification {
	this := Qualification{}
	return &this
}

// GetQualified returns the Qualified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Qualification) GetQualified() bool {
	if o == nil || IsNil(o.Qualified.Get()) {
		var ret bool
		return ret
	}
	return *o.Qualified.Get()
}

// GetQualifiedOk returns a tuple with the Qualified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Qualification) GetQualifiedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Qualified.Get(), o.Qualified.IsSet()
}

// HasQualified returns a boolean if a field has been set.
func (o *Qualification) HasQualified() bool {
	if o != nil && o.Qualified.IsSet() {
		return true
	}

	return false
}

// SetQualified gets a reference to the given NullableBool and assigns it to the Qualified field.
func (o *Qualification) SetQualified(v bool) {
	o.Qualified.Set(&v)
}
// SetQualifiedNil sets the value for Qualified to be an explicit nil
func (o *Qualification) SetQualifiedNil() {
	o.Qualified.Set(nil)
}

// UnsetQualified ensures that no value is present for Qualified, not even an explicit nil
func (o *Qualification) UnsetQualified() {
	o.Qualified.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Qualification) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Qualification) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Qualification) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *Qualification) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o Qualification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Qualification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Qualified.IsSet() {
		toSerialize["qualified"] = o.Qualified.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Qualification) UnmarshalJSON(data []byte) (err error) {
	varQualification := _Qualification{}

	err = json.Unmarshal(data, &varQualification)

	if err != nil {
		return err
	}

	*o = Qualification(varQualification)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "qualified")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableQualification struct {
	value *Qualification
	isSet bool
}

func (v NullableQualification) Get() *Qualification {
	return v.value
}

func (v *NullableQualification) Set(val *Qualification) {
	v.value = val
	v.isSet = true
}

func (v NullableQualification) IsSet() bool {
	return v.isSet
}

func (v *NullableQualification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQualification(val *Qualification) *NullableQualification {
	return &NullableQualification{value: val, isSet: true}
}

func (v NullableQualification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQualification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


