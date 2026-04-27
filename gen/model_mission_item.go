
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the MissionItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MissionItem{}

// MissionItem struct for MissionItem
type MissionItem struct {
	Mission map[string]interface{} `json:"mission,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MissionItem MissionItem

// NewMissionItem instantiates a new MissionItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMissionItem() *MissionItem {
	this := MissionItem{}
	return &this
}

// NewMissionItemWithDefaults instantiates a new MissionItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMissionItemWithDefaults() *MissionItem {
	this := MissionItem{}
	return &this
}

// GetMission returns the Mission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MissionItem) GetMission() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Mission
}

// GetMissionOk returns a tuple with the Mission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MissionItem) GetMissionOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Mission) {
		return map[string]interface{}{}, false
	}
	return o.Mission, true
}

// HasMission returns a boolean if a field has been set.
func (o *MissionItem) HasMission() bool {
	if o != nil && !IsNil(o.Mission) {
		return true
	}

	return false
}

// SetMission gets a reference to the given map[string]interface{} and assigns it to the Mission field.
func (o *MissionItem) SetMission(v map[string]interface{}) {
	o.Mission = v
}

func (o MissionItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MissionItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Mission != nil {
		toSerialize["mission"] = o.Mission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *MissionItem) UnmarshalJSON(data []byte) (err error) {
	varMissionItem := _MissionItem{}

	err = json.Unmarshal(data, &varMissionItem)

	if err != nil {
		return err
	}

	*o = MissionItem(varMissionItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "mission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMissionItem struct {
	value *MissionItem
	isSet bool
}

func (v NullableMissionItem) Get() *MissionItem {
	return v.value
}

func (v *NullableMissionItem) Set(val *MissionItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMissionItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMissionItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMissionItem(val *MissionItem) *NullableMissionItem {
	return &NullableMissionItem{value: val, isSet: true}
}

func (v NullableMissionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMissionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


