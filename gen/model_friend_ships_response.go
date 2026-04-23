
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FriendShipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FriendShipsResponse{}

// FriendShipsResponse struct for FriendShipsResponse
type FriendShipsResponse struct {
	Relationship NullableRelationship `json:"relationship,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FriendShipsResponse FriendShipsResponse

// NewFriendShipsResponse instantiates a new FriendShipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFriendShipsResponse() *FriendShipsResponse {
	this := FriendShipsResponse{}
	return &this
}

// NewFriendShipsResponseWithDefaults instantiates a new FriendShipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFriendShipsResponseWithDefaults() *FriendShipsResponse {
	this := FriendShipsResponse{}
	return &this
}

// GetRelationship returns the Relationship field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FriendShipsResponse) GetRelationship() Relationship {
	if o == nil || IsNil(o.Relationship.Get()) {
		var ret Relationship
		return ret
	}
	return *o.Relationship.Get()
}

// GetRelationshipOk returns a tuple with the Relationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FriendShipsResponse) GetRelationshipOk() (*Relationship, bool) {
	if o == nil {
		return nil, false
	}
	return o.Relationship.Get(), o.Relationship.IsSet()
}

// HasRelationship returns a boolean if a field has been set.
func (o *FriendShipsResponse) HasRelationship() bool {
	if o != nil && o.Relationship.IsSet() {
		return true
	}

	return false
}

// SetRelationship gets a reference to the given NullableRelationship and assigns it to the Relationship field.
func (o *FriendShipsResponse) SetRelationship(v Relationship) {
	o.Relationship.Set(&v)
}
// SetRelationshipNil sets the value for Relationship to be an explicit nil
func (o *FriendShipsResponse) SetRelationshipNil() {
	o.Relationship.Set(nil)
}

// UnsetRelationship ensures that no value is present for Relationship, not even an explicit nil
func (o *FriendShipsResponse) UnsetRelationship() {
	o.Relationship.Unset()
}

func (o FriendShipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FriendShipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Relationship.IsSet() {
		toSerialize["relationship"] = o.Relationship.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FriendShipsResponse) UnmarshalJSON(data []byte) (err error) {
	varFriendShipsResponse := _FriendShipsResponse{}

	err = json.Unmarshal(data, &varFriendShipsResponse)

	if err != nil {
		return err
	}

	*o = FriendShipsResponse(varFriendShipsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "relationship")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFriendShipsResponse struct {
	value *FriendShipsResponse
	isSet bool
}

func (v NullableFriendShipsResponse) Get() *FriendShipsResponse {
	return v.value
}

func (v *NullableFriendShipsResponse) Set(val *FriendShipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFriendShipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFriendShipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFriendShipsResponse(val *FriendShipsResponse) *NullableFriendShipsResponse {
	return &NullableFriendShipsResponse{value: val, isSet: true}
}

func (v NullableFriendShipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFriendShipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


