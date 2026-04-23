
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the FirebaseExperiment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirebaseExperiment{}

// FirebaseExperiment struct for FirebaseExperiment
type FirebaseExperiment struct {
	ExperimentNumber NullableInt32 `json:"experiment_number,omitempty"`
	Number NullableInt32 `json:"number,omitempty"`
	VariantNumber NullableInt32 `json:"variant_number,omitempty"`
	Visible NullableBool `json:"visible,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirebaseExperiment FirebaseExperiment

// NewFirebaseExperiment instantiates a new FirebaseExperiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirebaseExperiment() *FirebaseExperiment {
	this := FirebaseExperiment{}
	return &this
}

// NewFirebaseExperimentWithDefaults instantiates a new FirebaseExperiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirebaseExperimentWithDefaults() *FirebaseExperiment {
	this := FirebaseExperiment{}
	return &this
}

// GetExperimentNumber returns the ExperimentNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseExperiment) GetExperimentNumber() int32 {
	if o == nil || IsNil(o.ExperimentNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ExperimentNumber.Get()
}

// GetExperimentNumberOk returns a tuple with the ExperimentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseExperiment) GetExperimentNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentNumber.Get(), o.ExperimentNumber.IsSet()
}

// HasExperimentNumber returns a boolean if a field has been set.
func (o *FirebaseExperiment) HasExperimentNumber() bool {
	if o != nil && o.ExperimentNumber.IsSet() {
		return true
	}

	return false
}

// SetExperimentNumber gets a reference to the given NullableInt32 and assigns it to the ExperimentNumber field.
func (o *FirebaseExperiment) SetExperimentNumber(v int32) {
	o.ExperimentNumber.Set(&v)
}
// SetExperimentNumberNil sets the value for ExperimentNumber to be an explicit nil
func (o *FirebaseExperiment) SetExperimentNumberNil() {
	o.ExperimentNumber.Set(nil)
}

// UnsetExperimentNumber ensures that no value is present for ExperimentNumber, not even an explicit nil
func (o *FirebaseExperiment) UnsetExperimentNumber() {
	o.ExperimentNumber.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseExperiment) GetNumber() int32 {
	if o == nil || IsNil(o.Number.Get()) {
		var ret int32
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseExperiment) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *FirebaseExperiment) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableInt32 and assigns it to the Number field.
func (o *FirebaseExperiment) SetNumber(v int32) {
	o.Number.Set(&v)
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *FirebaseExperiment) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *FirebaseExperiment) UnsetNumber() {
	o.Number.Unset()
}

// GetVariantNumber returns the VariantNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseExperiment) GetVariantNumber() int32 {
	if o == nil || IsNil(o.VariantNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.VariantNumber.Get()
}

// GetVariantNumberOk returns a tuple with the VariantNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseExperiment) GetVariantNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantNumber.Get(), o.VariantNumber.IsSet()
}

// HasVariantNumber returns a boolean if a field has been set.
func (o *FirebaseExperiment) HasVariantNumber() bool {
	if o != nil && o.VariantNumber.IsSet() {
		return true
	}

	return false
}

// SetVariantNumber gets a reference to the given NullableInt32 and assigns it to the VariantNumber field.
func (o *FirebaseExperiment) SetVariantNumber(v int32) {
	o.VariantNumber.Set(&v)
}
// SetVariantNumberNil sets the value for VariantNumber to be an explicit nil
func (o *FirebaseExperiment) SetVariantNumberNil() {
	o.VariantNumber.Set(nil)
}

// UnsetVariantNumber ensures that no value is present for VariantNumber, not even an explicit nil
func (o *FirebaseExperiment) UnsetVariantNumber() {
	o.VariantNumber.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FirebaseExperiment) GetVisible() bool {
	if o == nil || IsNil(o.Visible.Get()) {
		var ret bool
		return ret
	}
	return *o.Visible.Get()
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FirebaseExperiment) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visible.Get(), o.Visible.IsSet()
}

// HasVisible returns a boolean if a field has been set.
func (o *FirebaseExperiment) HasVisible() bool {
	if o != nil && o.Visible.IsSet() {
		return true
	}

	return false
}

// SetVisible gets a reference to the given NullableBool and assigns it to the Visible field.
func (o *FirebaseExperiment) SetVisible(v bool) {
	o.Visible.Set(&v)
}
// SetVisibleNil sets the value for Visible to be an explicit nil
func (o *FirebaseExperiment) SetVisibleNil() {
	o.Visible.Set(nil)
}

// UnsetVisible ensures that no value is present for Visible, not even an explicit nil
func (o *FirebaseExperiment) UnsetVisible() {
	o.Visible.Unset()
}

func (o FirebaseExperiment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirebaseExperiment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ExperimentNumber.IsSet() {
		toSerialize["experiment_number"] = o.ExperimentNumber.Get()
	}
	if o.Number.IsSet() {
		toSerialize["number"] = o.Number.Get()
	}
	if o.VariantNumber.IsSet() {
		toSerialize["variant_number"] = o.VariantNumber.Get()
	}
	if o.Visible.IsSet() {
		toSerialize["visible"] = o.Visible.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FirebaseExperiment) UnmarshalJSON(data []byte) (err error) {
	varFirebaseExperiment := _FirebaseExperiment{}

	err = json.Unmarshal(data, &varFirebaseExperiment)

	if err != nil {
		return err
	}

	*o = FirebaseExperiment(varFirebaseExperiment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "experiment_number")
		delete(additionalProperties, "number")
		delete(additionalProperties, "variant_number")
		delete(additionalProperties, "visible")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirebaseExperiment struct {
	value *FirebaseExperiment
	isSet bool
}

func (v NullableFirebaseExperiment) Get() *FirebaseExperiment {
	return v.value
}

func (v *NullableFirebaseExperiment) Set(val *FirebaseExperiment) {
	v.value = val
	v.isSet = true
}

func (v NullableFirebaseExperiment) IsSet() bool {
	return v.isSet
}

func (v *NullableFirebaseExperiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirebaseExperiment(val *FirebaseExperiment) *NullableFirebaseExperiment {
	return &NullableFirebaseExperiment{value: val, isSet: true}
}

func (v NullableFirebaseExperiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirebaseExperiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


