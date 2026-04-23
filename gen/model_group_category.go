
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupCategory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupCategory{}

// GroupCategory struct for GroupCategory
type GroupCategory struct {
	Icon NullableString `json:"icon,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Rank NullableInt32 `json:"rank,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupCategory GroupCategory

// NewGroupCategory instantiates a new GroupCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCategory() *GroupCategory {
	this := GroupCategory{}
	return &this
}

// NewGroupCategoryWithDefaults instantiates a new GroupCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCategoryWithDefaults() *GroupCategory {
	this := GroupCategory{}
	return &this
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCategory) GetIcon() string {
	if o == nil || IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCategory) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *GroupCategory) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *GroupCategory) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *GroupCategory) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *GroupCategory) UnsetIcon() {
	o.Icon.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCategory) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCategory) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GroupCategory) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *GroupCategory) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GroupCategory) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GroupCategory) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCategory) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *GroupCategory) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *GroupCategory) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *GroupCategory) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *GroupCategory) UnsetName() {
	o.Name.Unset()
}

// GetRank returns the Rank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCategory) GetRank() int32 {
	if o == nil || IsNil(o.Rank.Get()) {
		var ret int32
		return ret
	}
	return *o.Rank.Get()
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCategory) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rank.Get(), o.Rank.IsSet()
}

// HasRank returns a boolean if a field has been set.
func (o *GroupCategory) HasRank() bool {
	if o != nil && o.Rank.IsSet() {
		return true
	}

	return false
}

// SetRank gets a reference to the given NullableInt32 and assigns it to the Rank field.
func (o *GroupCategory) SetRank(v int32) {
	o.Rank.Set(&v)
}
// SetRankNil sets the value for Rank to be an explicit nil
func (o *GroupCategory) SetRankNil() {
	o.Rank.Set(nil)
}

// UnsetRank ensures that no value is present for Rank, not even an explicit nil
func (o *GroupCategory) UnsetRank() {
	o.Rank.Unset()
}

func (o GroupCategory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupCategory) ToMap() (map[string]interface{}, error) {
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
	if o.Rank.IsSet() {
		toSerialize["rank"] = o.Rank.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupCategory) UnmarshalJSON(data []byte) (err error) {
	varGroupCategory := _GroupCategory{}

	err = json.Unmarshal(data, &varGroupCategory)

	if err != nil {
		return err
	}

	*o = GroupCategory(varGroupCategory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "icon")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "rank")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupCategory struct {
	value *GroupCategory
	isSet bool
}

func (v NullableGroupCategory) Get() *GroupCategory {
	return v.value
}

func (v *NullableGroupCategory) Set(val *GroupCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCategory(val *GroupCategory) *NullableGroupCategory {
	return &NullableGroupCategory{value: val, isSet: true}
}

func (v NullableGroupCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


