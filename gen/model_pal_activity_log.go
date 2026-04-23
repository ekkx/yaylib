
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalActivityLog type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalActivityLog{}

// PalActivityLog struct for PalActivityLog
type PalActivityLog struct {
	CreatedAtSeconds NullableInt64 `json:"created_at_seconds,omitempty"`
	Details NullablePalActivityLogDetails `json:"details,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	// Unresolved Java type: jp.nanameue.yay.api.PalActivityType
	Type map[string]interface{} `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalActivityLog PalActivityLog

// NewPalActivityLog instantiates a new PalActivityLog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalActivityLog() *PalActivityLog {
	this := PalActivityLog{}
	return &this
}

// NewPalActivityLogWithDefaults instantiates a new PalActivityLog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalActivityLogWithDefaults() *PalActivityLog {
	this := PalActivityLog{}
	return &this
}

// GetCreatedAtSeconds returns the CreatedAtSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLog) GetCreatedAtSeconds() int64 {
	if o == nil || IsNil(o.CreatedAtSeconds.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAtSeconds.Get()
}

// GetCreatedAtSecondsOk returns a tuple with the CreatedAtSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLog) GetCreatedAtSecondsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAtSeconds.Get(), o.CreatedAtSeconds.IsSet()
}

// HasCreatedAtSeconds returns a boolean if a field has been set.
func (o *PalActivityLog) HasCreatedAtSeconds() bool {
	if o != nil && o.CreatedAtSeconds.IsSet() {
		return true
	}

	return false
}

// SetCreatedAtSeconds gets a reference to the given NullableInt64 and assigns it to the CreatedAtSeconds field.
func (o *PalActivityLog) SetCreatedAtSeconds(v int64) {
	o.CreatedAtSeconds.Set(&v)
}
// SetCreatedAtSecondsNil sets the value for CreatedAtSeconds to be an explicit nil
func (o *PalActivityLog) SetCreatedAtSecondsNil() {
	o.CreatedAtSeconds.Set(nil)
}

// UnsetCreatedAtSeconds ensures that no value is present for CreatedAtSeconds, not even an explicit nil
func (o *PalActivityLog) UnsetCreatedAtSeconds() {
	o.CreatedAtSeconds.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLog) GetDetails() PalActivityLogDetails {
	if o == nil || IsNil(o.Details.Get()) {
		var ret PalActivityLogDetails
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLog) GetDetailsOk() (*PalActivityLogDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *PalActivityLog) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullablePalActivityLogDetails and assigns it to the Details field.
func (o *PalActivityLog) SetDetails(v PalActivityLogDetails) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *PalActivityLog) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *PalActivityLog) UnsetDetails() {
	o.Details.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLog) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLog) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *PalActivityLog) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *PalActivityLog) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *PalActivityLog) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *PalActivityLog) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLog) GetType() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLog) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PalActivityLog) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *PalActivityLog) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o PalActivityLog) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalActivityLog) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAtSeconds.IsSet() {
		toSerialize["created_at_seconds"] = o.CreatedAtSeconds.Get()
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalActivityLog) UnmarshalJSON(data []byte) (err error) {
	varPalActivityLog := _PalActivityLog{}

	err = json.Unmarshal(data, &varPalActivityLog)

	if err != nil {
		return err
	}

	*o = PalActivityLog(varPalActivityLog)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at_seconds")
		delete(additionalProperties, "details")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalActivityLog struct {
	value *PalActivityLog
	isSet bool
}

func (v NullablePalActivityLog) Get() *PalActivityLog {
	return v.value
}

func (v *NullablePalActivityLog) Set(val *PalActivityLog) {
	v.value = val
	v.isSet = true
}

func (v NullablePalActivityLog) IsSet() bool {
	return v.isSet
}

func (v *NullablePalActivityLog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalActivityLog(val *PalActivityLog) *NullablePalActivityLog {
	return &NullablePalActivityLog{value: val, isSet: true}
}

func (v NullablePalActivityLog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalActivityLog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


