
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ModelInterest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInterest{}

// ModelInterest struct for ModelInterest
type ModelInterest struct {
	IconUrl NullableString `json:"icon_url,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	IsSelected NullableBool `json:"is_selected,omitempty"`
	Name NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ModelInterest ModelInterest

// NewModelInterest instantiates a new ModelInterest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInterest() *ModelInterest {
	this := ModelInterest{}
	return &this
}

// NewModelInterestWithDefaults instantiates a new ModelInterest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInterestWithDefaults() *ModelInterest {
	this := ModelInterest{}
	return &this
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInterest) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInterest) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *ModelInterest) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *ModelInterest) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *ModelInterest) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *ModelInterest) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInterest) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInterest) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *ModelInterest) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *ModelInterest) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *ModelInterest) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *ModelInterest) UnsetId() {
	o.Id.Unset()
}

// GetIsSelected returns the IsSelected field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInterest) GetIsSelected() bool {
	if o == nil || IsNil(o.IsSelected.Get()) {
		var ret bool
		return ret
	}
	return *o.IsSelected.Get()
}

// GetIsSelectedOk returns a tuple with the IsSelected field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInterest) GetIsSelectedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSelected.Get(), o.IsSelected.IsSet()
}

// HasIsSelected returns a boolean if a field has been set.
func (o *ModelInterest) HasIsSelected() bool {
	if o != nil && o.IsSelected.IsSet() {
		return true
	}

	return false
}

// SetIsSelected gets a reference to the given NullableBool and assigns it to the IsSelected field.
func (o *ModelInterest) SetIsSelected(v bool) {
	o.IsSelected.Set(&v)
}
// SetIsSelectedNil sets the value for IsSelected to be an explicit nil
func (o *ModelInterest) SetIsSelectedNil() {
	o.IsSelected.Set(nil)
}

// UnsetIsSelected ensures that no value is present for IsSelected, not even an explicit nil
func (o *ModelInterest) UnsetIsSelected() {
	o.IsSelected.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelInterest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelInterest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ModelInterest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ModelInterest) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ModelInterest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ModelInterest) UnsetName() {
	o.Name.Unset()
}

func (o ModelInterest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInterest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.IsSelected.IsSet() {
		toSerialize["is_selected"] = o.IsSelected.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ModelInterest) UnmarshalJSON(data []byte) (err error) {
	varModelInterest := _ModelInterest{}

	err = json.Unmarshal(data, &varModelInterest)

	if err != nil {
		return err
	}

	*o = ModelInterest(varModelInterest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon_url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_selected")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableModelInterest struct {
	value *ModelInterest
	isSet bool
}

func (v NullableModelInterest) Get() *ModelInterest {
	return v.value
}

func (v *NullableModelInterest) Set(val *ModelInterest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInterest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInterest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInterest(val *ModelInterest) *NullableModelInterest {
	return &NullableModelInterest{value: val, isSet: true}
}

func (v NullableModelInterest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInterest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


