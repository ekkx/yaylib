
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupCategoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupCategoriesResponse{}

// GroupCategoriesResponse struct for GroupCategoriesResponse
type GroupCategoriesResponse struct {
	GroupCategories []GroupCategory `json:"group_categories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupCategoriesResponse GroupCategoriesResponse

// NewGroupCategoriesResponse instantiates a new GroupCategoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupCategoriesResponse() *GroupCategoriesResponse {
	this := GroupCategoriesResponse{}
	return &this
}

// NewGroupCategoriesResponseWithDefaults instantiates a new GroupCategoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupCategoriesResponseWithDefaults() *GroupCategoriesResponse {
	this := GroupCategoriesResponse{}
	return &this
}

// GetGroupCategories returns the GroupCategories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupCategoriesResponse) GetGroupCategories() []GroupCategory {
	if o == nil {
		var ret []GroupCategory
		return ret
	}
	return o.GroupCategories
}

// GetGroupCategoriesOk returns a tuple with the GroupCategories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupCategoriesResponse) GetGroupCategoriesOk() ([]GroupCategory, bool) {
	if o == nil || IsNil(o.GroupCategories) {
		return nil, false
	}
	return o.GroupCategories, true
}

// HasGroupCategories returns a boolean if a field has been set.
func (o *GroupCategoriesResponse) HasGroupCategories() bool {
	if o != nil && !IsNil(o.GroupCategories) {
		return true
	}

	return false
}

// SetGroupCategories gets a reference to the given []GroupCategory and assigns it to the GroupCategories field.
func (o *GroupCategoriesResponse) SetGroupCategories(v []GroupCategory) {
	o.GroupCategories = v
}

func (o GroupCategoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupCategoriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupCategories != nil {
		toSerialize["group_categories"] = o.GroupCategories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupCategoriesResponse := _GroupCategoriesResponse{}

	err = json.Unmarshal(data, &varGroupCategoriesResponse)

	if err != nil {
		return err
	}

	*o = GroupCategoriesResponse(varGroupCategoriesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "group_categories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupCategoriesResponse struct {
	value *GroupCategoriesResponse
	isSet bool
}

func (v NullableGroupCategoriesResponse) Get() *GroupCategoriesResponse {
	return v.value
}

func (v *NullableGroupCategoriesResponse) Set(val *GroupCategoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupCategoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupCategoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupCategoriesResponse(val *GroupCategoriesResponse) *NullableGroupCategoriesResponse {
	return &NullableGroupCategoriesResponse{value: val, isSet: true}
}

func (v NullableGroupCategoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupCategoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


