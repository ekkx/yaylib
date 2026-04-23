
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Navigation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Navigation{}

// Navigation struct for Navigation
type Navigation struct {
	NavigationItemList []SectionNavigation `json:"navigation_item_list,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Navigation Navigation

// NewNavigation instantiates a new Navigation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNavigation() *Navigation {
	this := Navigation{}
	return &this
}

// NewNavigationWithDefaults instantiates a new Navigation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNavigationWithDefaults() *Navigation {
	this := Navigation{}
	return &this
}

// GetNavigationItemList returns the NavigationItemList field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Navigation) GetNavigationItemList() []SectionNavigation {
	if o == nil {
		var ret []SectionNavigation
		return ret
	}
	return o.NavigationItemList
}

// GetNavigationItemListOk returns a tuple with the NavigationItemList field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Navigation) GetNavigationItemListOk() ([]SectionNavigation, bool) {
	if o == nil || IsNil(o.NavigationItemList) {
		return nil, false
	}
	return o.NavigationItemList, true
}

// HasNavigationItemList returns a boolean if a field has been set.
func (o *Navigation) HasNavigationItemList() bool {
	if o != nil && !IsNil(o.NavigationItemList) {
		return true
	}

	return false
}

// SetNavigationItemList gets a reference to the given []SectionNavigation and assigns it to the NavigationItemList field.
func (o *Navigation) SetNavigationItemList(v []SectionNavigation) {
	o.NavigationItemList = v
}

func (o Navigation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Navigation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NavigationItemList != nil {
		toSerialize["navigation_item_list"] = o.NavigationItemList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Navigation) UnmarshalJSON(data []byte) (err error) {
	varNavigation := _Navigation{}

	err = json.Unmarshal(data, &varNavigation)

	if err != nil {
		return err
	}

	*o = Navigation(varNavigation)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "navigation_item_list")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNavigation struct {
	value *Navigation
	isSet bool
}

func (v NullableNavigation) Get() *Navigation {
	return v.value
}

func (v *NullableNavigation) Set(val *Navigation) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigation) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigation(val *Navigation) *NullableNavigation {
	return &NullableNavigation{value: val, isSet: true}
}

func (v NullableNavigation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


