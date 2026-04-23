
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PendingEmplActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PendingEmplActivity{}

// PendingEmplActivity struct for PendingEmplActivity
type PendingEmplActivity struct {
	// Unresolved Java type: jp.nanameue.yay.shared.data.model.empl.EmplActivity
	Activity map[string]interface{} `json:"activity,omitempty"`
	Details NullablePendingEmplActivityDetails `json:"details,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PendingEmplActivity PendingEmplActivity

// NewPendingEmplActivity instantiates a new PendingEmplActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPendingEmplActivity() *PendingEmplActivity {
	this := PendingEmplActivity{}
	return &this
}

// NewPendingEmplActivityWithDefaults instantiates a new PendingEmplActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPendingEmplActivityWithDefaults() *PendingEmplActivity {
	this := PendingEmplActivity{}
	return &this
}

// GetActivity returns the Activity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingEmplActivity) GetActivity() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Activity
}

// GetActivityOk returns a tuple with the Activity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingEmplActivity) GetActivityOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Activity) {
		return map[string]interface{}{}, false
	}
	return o.Activity, true
}

// HasActivity returns a boolean if a field has been set.
func (o *PendingEmplActivity) HasActivity() bool {
	if o != nil && !IsNil(o.Activity) {
		return true
	}

	return false
}

// SetActivity gets a reference to the given map[string]interface{} and assigns it to the Activity field.
func (o *PendingEmplActivity) SetActivity(v map[string]interface{}) {
	o.Activity = v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PendingEmplActivity) GetDetails() PendingEmplActivityDetails {
	if o == nil || IsNil(o.Details.Get()) {
		var ret PendingEmplActivityDetails
		return ret
	}
	return *o.Details.Get()
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PendingEmplActivity) GetDetailsOk() (*PendingEmplActivityDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Details.Get(), o.Details.IsSet()
}

// HasDetails returns a boolean if a field has been set.
func (o *PendingEmplActivity) HasDetails() bool {
	if o != nil && o.Details.IsSet() {
		return true
	}

	return false
}

// SetDetails gets a reference to the given NullablePendingEmplActivityDetails and assigns it to the Details field.
func (o *PendingEmplActivity) SetDetails(v PendingEmplActivityDetails) {
	o.Details.Set(&v)
}
// SetDetailsNil sets the value for Details to be an explicit nil
func (o *PendingEmplActivity) SetDetailsNil() {
	o.Details.Set(nil)
}

// UnsetDetails ensures that no value is present for Details, not even an explicit nil
func (o *PendingEmplActivity) UnsetDetails() {
	o.Details.Unset()
}

func (o PendingEmplActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PendingEmplActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Activity != nil {
		toSerialize["activity"] = o.Activity
	}
	if o.Details.IsSet() {
		toSerialize["details"] = o.Details.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PendingEmplActivity) UnmarshalJSON(data []byte) (err error) {
	varPendingEmplActivity := _PendingEmplActivity{}

	err = json.Unmarshal(data, &varPendingEmplActivity)

	if err != nil {
		return err
	}

	*o = PendingEmplActivity(varPendingEmplActivity)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activity")
		delete(additionalProperties, "details")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePendingEmplActivity struct {
	value *PendingEmplActivity
	isSet bool
}

func (v NullablePendingEmplActivity) Get() *PendingEmplActivity {
	return v.value
}

func (v *NullablePendingEmplActivity) Set(val *PendingEmplActivity) {
	v.value = val
	v.isSet = true
}

func (v NullablePendingEmplActivity) IsSet() bool {
	return v.isSet
}

func (v *NullablePendingEmplActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePendingEmplActivity(val *PendingEmplActivity) *NullablePendingEmplActivity {
	return &NullablePendingEmplActivity{value: val, isSet: true}
}

func (v NullablePendingEmplActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePendingEmplActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


