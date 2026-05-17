
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupsRelatedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsRelatedResponse{}

// GroupsRelatedResponse struct for GroupsRelatedResponse
type GroupsRelatedResponse struct {
	Groups []Group `json:"groups,omitempty"`
	NextPageValue NullableCursor `json:"next_page_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupsRelatedResponse GroupsRelatedResponse

// NewGroupsRelatedResponse instantiates a new GroupsRelatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsRelatedResponse() *GroupsRelatedResponse {
	this := GroupsRelatedResponse{}
	return &this
}

// NewGroupsRelatedResponseWithDefaults instantiates a new GroupsRelatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsRelatedResponseWithDefaults() *GroupsRelatedResponse {
	this := GroupsRelatedResponse{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsRelatedResponse) GetGroups() []Group {
	if o == nil {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsRelatedResponse) GetGroupsOk() ([]Group, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *GroupsRelatedResponse) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *GroupsRelatedResponse) SetGroups(v []Group) {
	o.Groups = v
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupsRelatedResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupsRelatedResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GroupsRelatedResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *GroupsRelatedResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GroupsRelatedResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GroupsRelatedResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

func (o GroupsRelatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsRelatedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupsRelatedResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupsRelatedResponse := _GroupsRelatedResponse{}

	err = json.Unmarshal(data, &varGroupsRelatedResponse)

	if err != nil {
		return err
	}

	*o = GroupsRelatedResponse(varGroupsRelatedResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "groups")
		delete(additionalProperties, "next_page_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupsRelatedResponse struct {
	value *GroupsRelatedResponse
	isSet bool
}

func (v NullableGroupsRelatedResponse) Get() *GroupsRelatedResponse {
	return v.value
}

func (v *NullableGroupsRelatedResponse) Set(val *GroupsRelatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsRelatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsRelatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsRelatedResponse(val *GroupsRelatedResponse) *NullableGroupsRelatedResponse {
	return &NullableGroupsRelatedResponse{value: val, isSet: true}
}

func (v NullableGroupsRelatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsRelatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


