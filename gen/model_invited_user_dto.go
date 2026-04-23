
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the InvitedUserDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InvitedUserDTO{}

// InvitedUserDTO struct for InvitedUserDTO
type InvitedUserDTO struct {
	Referred NullableInvitedUserDetailsDTO `json:"referred,omitempty"`
	TotalEarnedPoints NullableInt32 `json:"total_earned_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InvitedUserDTO InvitedUserDTO

// NewInvitedUserDTO instantiates a new InvitedUserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitedUserDTO() *InvitedUserDTO {
	this := InvitedUserDTO{}
	return &this
}

// NewInvitedUserDTOWithDefaults instantiates a new InvitedUserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitedUserDTOWithDefaults() *InvitedUserDTO {
	this := InvitedUserDTO{}
	return &this
}

// GetReferred returns the Referred field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitedUserDTO) GetReferred() InvitedUserDetailsDTO {
	if o == nil || IsNil(o.Referred.Get()) {
		var ret InvitedUserDetailsDTO
		return ret
	}
	return *o.Referred.Get()
}

// GetReferredOk returns a tuple with the Referred field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitedUserDTO) GetReferredOk() (*InvitedUserDetailsDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Referred.Get(), o.Referred.IsSet()
}

// HasReferred returns a boolean if a field has been set.
func (o *InvitedUserDTO) HasReferred() bool {
	if o != nil && o.Referred.IsSet() {
		return true
	}

	return false
}

// SetReferred gets a reference to the given NullableInvitedUserDetailsDTO and assigns it to the Referred field.
func (o *InvitedUserDTO) SetReferred(v InvitedUserDetailsDTO) {
	o.Referred.Set(&v)
}
// SetReferredNil sets the value for Referred to be an explicit nil
func (o *InvitedUserDTO) SetReferredNil() {
	o.Referred.Set(nil)
}

// UnsetReferred ensures that no value is present for Referred, not even an explicit nil
func (o *InvitedUserDTO) UnsetReferred() {
	o.Referred.Unset()
}

// GetTotalEarnedPoints returns the TotalEarnedPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvitedUserDTO) GetTotalEarnedPoints() int32 {
	if o == nil || IsNil(o.TotalEarnedPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalEarnedPoints.Get()
}

// GetTotalEarnedPointsOk returns a tuple with the TotalEarnedPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvitedUserDTO) GetTotalEarnedPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalEarnedPoints.Get(), o.TotalEarnedPoints.IsSet()
}

// HasTotalEarnedPoints returns a boolean if a field has been set.
func (o *InvitedUserDTO) HasTotalEarnedPoints() bool {
	if o != nil && o.TotalEarnedPoints.IsSet() {
		return true
	}

	return false
}

// SetTotalEarnedPoints gets a reference to the given NullableInt32 and assigns it to the TotalEarnedPoints field.
func (o *InvitedUserDTO) SetTotalEarnedPoints(v int32) {
	o.TotalEarnedPoints.Set(&v)
}
// SetTotalEarnedPointsNil sets the value for TotalEarnedPoints to be an explicit nil
func (o *InvitedUserDTO) SetTotalEarnedPointsNil() {
	o.TotalEarnedPoints.Set(nil)
}

// UnsetTotalEarnedPoints ensures that no value is present for TotalEarnedPoints, not even an explicit nil
func (o *InvitedUserDTO) UnsetTotalEarnedPoints() {
	o.TotalEarnedPoints.Unset()
}

func (o InvitedUserDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InvitedUserDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Referred.IsSet() {
		toSerialize["referred"] = o.Referred.Get()
	}
	if o.TotalEarnedPoints.IsSet() {
		toSerialize["total_earned_points"] = o.TotalEarnedPoints.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InvitedUserDTO) UnmarshalJSON(data []byte) (err error) {
	varInvitedUserDTO := _InvitedUserDTO{}

	err = json.Unmarshal(data, &varInvitedUserDTO)

	if err != nil {
		return err
	}

	*o = InvitedUserDTO(varInvitedUserDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "referred")
		delete(additionalProperties, "total_earned_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvitedUserDTO struct {
	value *InvitedUserDTO
	isSet bool
}

func (v NullableInvitedUserDTO) Get() *InvitedUserDTO {
	return v.value
}

func (v *NullableInvitedUserDTO) Set(val *InvitedUserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableInvitedUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableInvitedUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvitedUserDTO(val *InvitedUserDTO) *NullableInvitedUserDTO {
	return &NullableInvitedUserDTO{value: val, isSet: true}
}

func (v NullableInvitedUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvitedUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


