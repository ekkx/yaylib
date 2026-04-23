
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Interest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Interest{}

// Interest struct for Interest
type Interest struct {
	Icon NullableString `json:"icon,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Selected NullableBool `json:"selected,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Interest Interest

// NewInterest instantiates a new Interest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterest() *Interest {
	this := Interest{}
	return &this
}

// NewInterestWithDefaults instantiates a new Interest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterestWithDefaults() *Interest {
	this := Interest{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interest) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interest) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *Interest) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *Interest) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *Interest) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *Interest) UnsetIcon() {
	o.Icon.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interest) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *Interest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *Interest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *Interest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *Interest) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Interest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Interest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Interest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Interest) UnsetName() {
	o.Name.Unset()
}

// GetSelected returns the Selected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Interest) GetSelected() bool {
	if o == nil || IsNil(o.Selected.Get()) {
		var ret bool
		return ret
	}
	return *o.Selected.Get()
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Interest) GetSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Selected.Get(), o.Selected.IsSet()
}

// HasSelected returns a boolean if a field has been set.
func (o *Interest) HasSelected() bool {
	if o != nil && o.Selected.IsSet() {
		return true
	}

	return false
}

// SetSelected gets a reference to the given NullableBool and assigns it to the Selected field.
func (o *Interest) SetSelected(v bool) {
	o.Selected.Set(&v)
}
// SetSelectedNil sets the value for Selected to be an explicit nil
func (o *Interest) SetSelectedNil() {
	o.Selected.Set(nil)
}

// UnsetSelected ensures that no value is present for Selected, not even an explicit nil
func (o *Interest) UnsetSelected() {
	o.Selected.Unset()
}

func (o Interest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Interest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Selected.IsSet() {
		toSerialize["selected"] = o.Selected.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Interest) UnmarshalJSON(data []byte) (err error) {
	varInterest := _Interest{}

	err = json.Unmarshal(data, &varInterest)

	if err != nil {
		return err
	}

	*o = Interest(varInterest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "selected")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInterest struct {
	value *Interest
	isSet bool
}

func (v NullableInterest) Get() *Interest {
	return v.value
}

func (v *NullableInterest) Set(val *Interest) {
	v.value = val
	v.isSet = true
}

func (v NullableInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterest(val *Interest) *NullableInterest {
	return &NullableInterest{value: val, isSet: true}
}

func (v NullableInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


