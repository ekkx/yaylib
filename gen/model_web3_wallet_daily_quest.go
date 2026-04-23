
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletDailyQuest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletDailyQuest{}

// Web3WalletDailyQuest struct for Web3WalletDailyQuest
type Web3WalletDailyQuest struct {
	IsCalculated NullableBool `json:"is_calculated,omitempty"`
	PalCount NullableInt32 `json:"pal_count,omitempty"`
	SpentDurationSeconds NullableInt32 `json:"spent_duration_seconds,omitempty"`
	TotalDurationNeededSeconds NullableInt32 `json:"total_duration_needed_seconds,omitempty"`
	UpdatedAtMillis NullableInt64 `json:"updated_at_millis,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletDailyQuest Web3WalletDailyQuest

// NewWeb3WalletDailyQuest instantiates a new Web3WalletDailyQuest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletDailyQuest() *Web3WalletDailyQuest {
	this := Web3WalletDailyQuest{}
	return &this
}

// NewWeb3WalletDailyQuestWithDefaults instantiates a new Web3WalletDailyQuest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletDailyQuestWithDefaults() *Web3WalletDailyQuest {
	this := Web3WalletDailyQuest{}
	return &this
}

// GetIsCalculated returns the IsCalculated field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletDailyQuest) GetIsCalculated() bool {
	if o == nil || IsNil(o.IsCalculated.Get()) {
		var ret bool
		return ret
	}
	return *o.IsCalculated.Get()
}

// GetIsCalculatedOk returns a tuple with the IsCalculated field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletDailyQuest) GetIsCalculatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsCalculated.Get(), o.IsCalculated.IsSet()
}

// HasIsCalculated returns a boolean if a field has been set.
func (o *Web3WalletDailyQuest) HasIsCalculated() bool {
	if o != nil && o.IsCalculated.IsSet() {
		return true
	}

	return false
}

// SetIsCalculated gets a reference to the given NullableBool and assigns it to the IsCalculated field.
func (o *Web3WalletDailyQuest) SetIsCalculated(v bool) {
	o.IsCalculated.Set(&v)
}
// SetIsCalculatedNil sets the value for IsCalculated to be an explicit nil
func (o *Web3WalletDailyQuest) SetIsCalculatedNil() {
	o.IsCalculated.Set(nil)
}

// UnsetIsCalculated ensures that no value is present for IsCalculated, not even an explicit nil
func (o *Web3WalletDailyQuest) UnsetIsCalculated() {
	o.IsCalculated.Unset()
}

// GetPalCount returns the PalCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletDailyQuest) GetPalCount() int32 {
	if o == nil || IsNil(o.PalCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PalCount.Get()
}

// GetPalCountOk returns a tuple with the PalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletDailyQuest) GetPalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalCount.Get(), o.PalCount.IsSet()
}

// HasPalCount returns a boolean if a field has been set.
func (o *Web3WalletDailyQuest) HasPalCount() bool {
	if o != nil && o.PalCount.IsSet() {
		return true
	}

	return false
}

// SetPalCount gets a reference to the given NullableInt32 and assigns it to the PalCount field.
func (o *Web3WalletDailyQuest) SetPalCount(v int32) {
	o.PalCount.Set(&v)
}
// SetPalCountNil sets the value for PalCount to be an explicit nil
func (o *Web3WalletDailyQuest) SetPalCountNil() {
	o.PalCount.Set(nil)
}

// UnsetPalCount ensures that no value is present for PalCount, not even an explicit nil
func (o *Web3WalletDailyQuest) UnsetPalCount() {
	o.PalCount.Unset()
}

// GetSpentDurationSeconds returns the SpentDurationSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletDailyQuest) GetSpentDurationSeconds() int32 {
	if o == nil || IsNil(o.SpentDurationSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.SpentDurationSeconds.Get()
}

// GetSpentDurationSecondsOk returns a tuple with the SpentDurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletDailyQuest) GetSpentDurationSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpentDurationSeconds.Get(), o.SpentDurationSeconds.IsSet()
}

// HasSpentDurationSeconds returns a boolean if a field has been set.
func (o *Web3WalletDailyQuest) HasSpentDurationSeconds() bool {
	if o != nil && o.SpentDurationSeconds.IsSet() {
		return true
	}

	return false
}

// SetSpentDurationSeconds gets a reference to the given NullableInt32 and assigns it to the SpentDurationSeconds field.
func (o *Web3WalletDailyQuest) SetSpentDurationSeconds(v int32) {
	o.SpentDurationSeconds.Set(&v)
}
// SetSpentDurationSecondsNil sets the value for SpentDurationSeconds to be an explicit nil
func (o *Web3WalletDailyQuest) SetSpentDurationSecondsNil() {
	o.SpentDurationSeconds.Set(nil)
}

// UnsetSpentDurationSeconds ensures that no value is present for SpentDurationSeconds, not even an explicit nil
func (o *Web3WalletDailyQuest) UnsetSpentDurationSeconds() {
	o.SpentDurationSeconds.Unset()
}

// GetTotalDurationNeededSeconds returns the TotalDurationNeededSeconds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletDailyQuest) GetTotalDurationNeededSeconds() int32 {
	if o == nil || IsNil(o.TotalDurationNeededSeconds.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalDurationNeededSeconds.Get()
}

// GetTotalDurationNeededSecondsOk returns a tuple with the TotalDurationNeededSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletDailyQuest) GetTotalDurationNeededSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalDurationNeededSeconds.Get(), o.TotalDurationNeededSeconds.IsSet()
}

// HasTotalDurationNeededSeconds returns a boolean if a field has been set.
func (o *Web3WalletDailyQuest) HasTotalDurationNeededSeconds() bool {
	if o != nil && o.TotalDurationNeededSeconds.IsSet() {
		return true
	}

	return false
}

// SetTotalDurationNeededSeconds gets a reference to the given NullableInt32 and assigns it to the TotalDurationNeededSeconds field.
func (o *Web3WalletDailyQuest) SetTotalDurationNeededSeconds(v int32) {
	o.TotalDurationNeededSeconds.Set(&v)
}
// SetTotalDurationNeededSecondsNil sets the value for TotalDurationNeededSeconds to be an explicit nil
func (o *Web3WalletDailyQuest) SetTotalDurationNeededSecondsNil() {
	o.TotalDurationNeededSeconds.Set(nil)
}

// UnsetTotalDurationNeededSeconds ensures that no value is present for TotalDurationNeededSeconds, not even an explicit nil
func (o *Web3WalletDailyQuest) UnsetTotalDurationNeededSeconds() {
	o.TotalDurationNeededSeconds.Unset()
}

// GetUpdatedAtMillis returns the UpdatedAtMillis field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletDailyQuest) GetUpdatedAtMillis() int64 {
	if o == nil || IsNil(o.UpdatedAtMillis.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAtMillis.Get()
}

// GetUpdatedAtMillisOk returns a tuple with the UpdatedAtMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletDailyQuest) GetUpdatedAtMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAtMillis.Get(), o.UpdatedAtMillis.IsSet()
}

// HasUpdatedAtMillis returns a boolean if a field has been set.
func (o *Web3WalletDailyQuest) HasUpdatedAtMillis() bool {
	if o != nil && o.UpdatedAtMillis.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAtMillis gets a reference to the given NullableInt64 and assigns it to the UpdatedAtMillis field.
func (o *Web3WalletDailyQuest) SetUpdatedAtMillis(v int64) {
	o.UpdatedAtMillis.Set(&v)
}
// SetUpdatedAtMillisNil sets the value for UpdatedAtMillis to be an explicit nil
func (o *Web3WalletDailyQuest) SetUpdatedAtMillisNil() {
	o.UpdatedAtMillis.Set(nil)
}

// UnsetUpdatedAtMillis ensures that no value is present for UpdatedAtMillis, not even an explicit nil
func (o *Web3WalletDailyQuest) UnsetUpdatedAtMillis() {
	o.UpdatedAtMillis.Unset()
}

func (o Web3WalletDailyQuest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletDailyQuest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IsCalculated.IsSet() {
		toSerialize["is_calculated"] = o.IsCalculated.Get()
	}
	if o.PalCount.IsSet() {
		toSerialize["pal_count"] = o.PalCount.Get()
	}
	if o.SpentDurationSeconds.IsSet() {
		toSerialize["spent_duration_seconds"] = o.SpentDurationSeconds.Get()
	}
	if o.TotalDurationNeededSeconds.IsSet() {
		toSerialize["total_duration_needed_seconds"] = o.TotalDurationNeededSeconds.Get()
	}
	if o.UpdatedAtMillis.IsSet() {
		toSerialize["updated_at_millis"] = o.UpdatedAtMillis.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletDailyQuest) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletDailyQuest := _Web3WalletDailyQuest{}

	err = json.Unmarshal(data, &varWeb3WalletDailyQuest)

	if err != nil {
		return err
	}

	*o = Web3WalletDailyQuest(varWeb3WalletDailyQuest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "is_calculated")
		delete(additionalProperties, "pal_count")
		delete(additionalProperties, "spent_duration_seconds")
		delete(additionalProperties, "total_duration_needed_seconds")
		delete(additionalProperties, "updated_at_millis")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletDailyQuest struct {
	value *Web3WalletDailyQuest
	isSet bool
}

func (v NullableWeb3WalletDailyQuest) Get() *Web3WalletDailyQuest {
	return v.value
}

func (v *NullableWeb3WalletDailyQuest) Set(val *Web3WalletDailyQuest) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletDailyQuest) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletDailyQuest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletDailyQuest(val *Web3WalletDailyQuest) *NullableWeb3WalletDailyQuest {
	return &NullableWeb3WalletDailyQuest{value: val, isSet: true}
}

func (v NullableWeb3WalletDailyQuest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletDailyQuest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


