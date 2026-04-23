
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ProgressDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgressDTO{}

// ProgressDTO struct for ProgressDTO
type ProgressDTO struct {
	ActionsCount NullableInt32 `json:"actions_count,omitempty"`
	CompletedAt NullableInt64 `json:"completed_at,omitempty"`
	TotalEarnedPoints NullableInt32 `json:"total_earned_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProgressDTO ProgressDTO

// NewProgressDTO instantiates a new ProgressDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgressDTO() *ProgressDTO {
	this := ProgressDTO{}
	return &this
}

// NewProgressDTOWithDefaults instantiates a new ProgressDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgressDTOWithDefaults() *ProgressDTO {
	this := ProgressDTO{}
	return &this
}

// GetActionsCount returns the ActionsCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressDTO) GetActionsCount() int32 {
	if o == nil || IsNil(o.ActionsCount.Get()) {
		var ret int32
		return ret
	}
	return *o.ActionsCount.Get()
}

// GetActionsCountOk returns a tuple with the ActionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressDTO) GetActionsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActionsCount.Get(), o.ActionsCount.IsSet()
}

// HasActionsCount returns a boolean if a field has been set.
func (o *ProgressDTO) HasActionsCount() bool {
	if o != nil && o.ActionsCount.IsSet() {
		return true
	}

	return false
}

// SetActionsCount gets a reference to the given NullableInt32 and assigns it to the ActionsCount field.
func (o *ProgressDTO) SetActionsCount(v int32) {
	o.ActionsCount.Set(&v)
}
// SetActionsCountNil sets the value for ActionsCount to be an explicit nil
func (o *ProgressDTO) SetActionsCountNil() {
	o.ActionsCount.Set(nil)
}

// UnsetActionsCount ensures that no value is present for ActionsCount, not even an explicit nil
func (o *ProgressDTO) UnsetActionsCount() {
	o.ActionsCount.Unset()
}

// GetCompletedAt returns the CompletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressDTO) GetCompletedAt() int64 {
	if o == nil || IsNil(o.CompletedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CompletedAt.Get()
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressDTO) GetCompletedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletedAt.Get(), o.CompletedAt.IsSet()
}

// HasCompletedAt returns a boolean if a field has been set.
func (o *ProgressDTO) HasCompletedAt() bool {
	if o != nil && o.CompletedAt.IsSet() {
		return true
	}

	return false
}

// SetCompletedAt gets a reference to the given NullableInt64 and assigns it to the CompletedAt field.
func (o *ProgressDTO) SetCompletedAt(v int64) {
	o.CompletedAt.Set(&v)
}
// SetCompletedAtNil sets the value for CompletedAt to be an explicit nil
func (o *ProgressDTO) SetCompletedAtNil() {
	o.CompletedAt.Set(nil)
}

// UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
func (o *ProgressDTO) UnsetCompletedAt() {
	o.CompletedAt.Unset()
}

// GetTotalEarnedPoints returns the TotalEarnedPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProgressDTO) GetTotalEarnedPoints() int32 {
	if o == nil || IsNil(o.TotalEarnedPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalEarnedPoints.Get()
}

// GetTotalEarnedPointsOk returns a tuple with the TotalEarnedPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProgressDTO) GetTotalEarnedPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalEarnedPoints.Get(), o.TotalEarnedPoints.IsSet()
}

// HasTotalEarnedPoints returns a boolean if a field has been set.
func (o *ProgressDTO) HasTotalEarnedPoints() bool {
	if o != nil && o.TotalEarnedPoints.IsSet() {
		return true
	}

	return false
}

// SetTotalEarnedPoints gets a reference to the given NullableInt32 and assigns it to the TotalEarnedPoints field.
func (o *ProgressDTO) SetTotalEarnedPoints(v int32) {
	o.TotalEarnedPoints.Set(&v)
}
// SetTotalEarnedPointsNil sets the value for TotalEarnedPoints to be an explicit nil
func (o *ProgressDTO) SetTotalEarnedPointsNil() {
	o.TotalEarnedPoints.Set(nil)
}

// UnsetTotalEarnedPoints ensures that no value is present for TotalEarnedPoints, not even an explicit nil
func (o *ProgressDTO) UnsetTotalEarnedPoints() {
	o.TotalEarnedPoints.Unset()
}

func (o ProgressDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgressDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ActionsCount.IsSet() {
		toSerialize["actions_count"] = o.ActionsCount.Get()
	}
	if o.CompletedAt.IsSet() {
		toSerialize["completed_at"] = o.CompletedAt.Get()
	}
	if o.TotalEarnedPoints.IsSet() {
		toSerialize["total_earned_points"] = o.TotalEarnedPoints.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProgressDTO) UnmarshalJSON(data []byte) (err error) {
	varProgressDTO := _ProgressDTO{}

	err = json.Unmarshal(data, &varProgressDTO)

	if err != nil {
		return err
	}

	*o = ProgressDTO(varProgressDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "actions_count")
		delete(additionalProperties, "completed_at")
		delete(additionalProperties, "total_earned_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableProgressDTO struct {
	value *ProgressDTO
	isSet bool
}

func (v NullableProgressDTO) Get() *ProgressDTO {
	return v.value
}

func (v *NullableProgressDTO) Set(val *ProgressDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableProgressDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableProgressDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgressDTO(val *ProgressDTO) *NullableProgressDTO {
	return &NullableProgressDTO{value: val, isSet: true}
}

func (v NullableProgressDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgressDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


