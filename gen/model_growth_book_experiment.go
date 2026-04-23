
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GrowthBookExperiment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrowthBookExperiment{}

// GrowthBookExperiment struct for GrowthBookExperiment
type GrowthBookExperiment struct {
	ExperimentNumber NullableInt32 `json:"experiment_number,omitempty"`
	Number NullableInt32 `json:"number,omitempty"`
	VariantNumber NullableInt32 `json:"variant_number,omitempty"`
	Visible NullableBool `json:"visible,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrowthBookExperiment GrowthBookExperiment

// NewGrowthBookExperiment instantiates a new GrowthBookExperiment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrowthBookExperiment() *GrowthBookExperiment {
	this := GrowthBookExperiment{}
	return &this
}

// NewGrowthBookExperimentWithDefaults instantiates a new GrowthBookExperiment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrowthBookExperimentWithDefaults() *GrowthBookExperiment {
	this := GrowthBookExperiment{}
	return &this
}

// GetExperimentNumber returns the ExperimentNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrowthBookExperiment) GetExperimentNumber() int32 {
	if o == nil || IsNil(o.ExperimentNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.ExperimentNumber.Get()
}

// GetExperimentNumberOk returns a tuple with the ExperimentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrowthBookExperiment) GetExperimentNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExperimentNumber.Get(), o.ExperimentNumber.IsSet()
}

// HasExperimentNumber returns a boolean if a field has been set.
func (o *GrowthBookExperiment) HasExperimentNumber() bool {
	if o != nil && o.ExperimentNumber.IsSet() {
		return true
	}

	return false
}

// SetExperimentNumber gets a reference to the given NullableInt32 and assigns it to the ExperimentNumber field.
func (o *GrowthBookExperiment) SetExperimentNumber(v int32) {
	o.ExperimentNumber.Set(&v)
}
// SetExperimentNumberNil sets the value for ExperimentNumber to be an explicit nil
func (o *GrowthBookExperiment) SetExperimentNumberNil() {
	o.ExperimentNumber.Set(nil)
}

// UnsetExperimentNumber ensures that no value is present for ExperimentNumber, not even an explicit nil
func (o *GrowthBookExperiment) UnsetExperimentNumber() {
	o.ExperimentNumber.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrowthBookExperiment) GetNumber() int32 {
	if o == nil || IsNil(o.Number.Get()) {
		var ret int32
		return ret
	}
	return *o.Number.Get()
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrowthBookExperiment) GetNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Number.Get(), o.Number.IsSet()
}

// HasNumber returns a boolean if a field has been set.
func (o *GrowthBookExperiment) HasNumber() bool {
	if o != nil && o.Number.IsSet() {
		return true
	}

	return false
}

// SetNumber gets a reference to the given NullableInt32 and assigns it to the Number field.
func (o *GrowthBookExperiment) SetNumber(v int32) {
	o.Number.Set(&v)
}
// SetNumberNil sets the value for Number to be an explicit nil
func (o *GrowthBookExperiment) SetNumberNil() {
	o.Number.Set(nil)
}

// UnsetNumber ensures that no value is present for Number, not even an explicit nil
func (o *GrowthBookExperiment) UnsetNumber() {
	o.Number.Unset()
}

// GetVariantNumber returns the VariantNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrowthBookExperiment) GetVariantNumber() int32 {
	if o == nil || IsNil(o.VariantNumber.Get()) {
		var ret int32
		return ret
	}
	return *o.VariantNumber.Get()
}

// GetVariantNumberOk returns a tuple with the VariantNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrowthBookExperiment) GetVariantNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.VariantNumber.Get(), o.VariantNumber.IsSet()
}

// HasVariantNumber returns a boolean if a field has been set.
func (o *GrowthBookExperiment) HasVariantNumber() bool {
	if o != nil && o.VariantNumber.IsSet() {
		return true
	}

	return false
}

// SetVariantNumber gets a reference to the given NullableInt32 and assigns it to the VariantNumber field.
func (o *GrowthBookExperiment) SetVariantNumber(v int32) {
	o.VariantNumber.Set(&v)
}
// SetVariantNumberNil sets the value for VariantNumber to be an explicit nil
func (o *GrowthBookExperiment) SetVariantNumberNil() {
	o.VariantNumber.Set(nil)
}

// UnsetVariantNumber ensures that no value is present for VariantNumber, not even an explicit nil
func (o *GrowthBookExperiment) UnsetVariantNumber() {
	o.VariantNumber.Unset()
}

// GetVisible returns the Visible field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GrowthBookExperiment) GetVisible() bool {
	if o == nil || IsNil(o.Visible.Get()) {
		var ret bool
		return ret
	}
	return *o.Visible.Get()
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GrowthBookExperiment) GetVisibleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Visible.Get(), o.Visible.IsSet()
}

// HasVisible returns a boolean if a field has been set.
func (o *GrowthBookExperiment) HasVisible() bool {
	if o != nil && o.Visible.IsSet() {
		return true
	}

	return false
}

// SetVisible gets a reference to the given NullableBool and assigns it to the Visible field.
func (o *GrowthBookExperiment) SetVisible(v bool) {
	o.Visible.Set(&v)
}
// SetVisibleNil sets the value for Visible to be an explicit nil
func (o *GrowthBookExperiment) SetVisibleNil() {
	o.Visible.Set(nil)
}

// UnsetVisible ensures that no value is present for Visible, not even an explicit nil
func (o *GrowthBookExperiment) UnsetVisible() {
	o.Visible.Unset()
}

func (o GrowthBookExperiment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrowthBookExperiment) ToMap() (map[string]interface{}, error) {
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

func (o *GrowthBookExperiment) UnmarshalJSON(data []byte) (err error) {
	varGrowthBookExperiment := _GrowthBookExperiment{}

	err = json.Unmarshal(data, &varGrowthBookExperiment)

	if err != nil {
		return err
	}

	*o = GrowthBookExperiment(varGrowthBookExperiment)

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

type NullableGrowthBookExperiment struct {
	value *GrowthBookExperiment
	isSet bool
}

func (v NullableGrowthBookExperiment) Get() *GrowthBookExperiment {
	return v.value
}

func (v *NullableGrowthBookExperiment) Set(val *GrowthBookExperiment) {
	v.value = val
	v.isSet = true
}

func (v NullableGrowthBookExperiment) IsSet() bool {
	return v.isSet
}

func (v *NullableGrowthBookExperiment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrowthBookExperiment(val *GrowthBookExperiment) *NullableGrowthBookExperiment {
	return &NullableGrowthBookExperiment{value: val, isSet: true}
}

func (v NullableGrowthBookExperiment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrowthBookExperiment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


