
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalActivityLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalActivityLogs{}

// PalActivityLogs struct for PalActivityLogs
type PalActivityLogs struct {
	Result []PalActivityLog `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalActivityLogs PalActivityLogs

// NewPalActivityLogs instantiates a new PalActivityLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalActivityLogs() *PalActivityLogs {
	this := PalActivityLogs{}
	return &this
}

// NewPalActivityLogsWithDefaults instantiates a new PalActivityLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalActivityLogsWithDefaults() *PalActivityLogs {
	this := PalActivityLogs{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogs) GetResult() []PalActivityLog {
	if o == nil {
		var ret []PalActivityLog
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogs) GetResultOk() ([]PalActivityLog, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *PalActivityLogs) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []PalActivityLog and assigns it to the Result field.
func (o *PalActivityLogs) SetResult(v []PalActivityLog) {
	o.Result = v
}

func (o PalActivityLogs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalActivityLogs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalActivityLogs) UnmarshalJSON(data []byte) (err error) {
	varPalActivityLogs := _PalActivityLogs{}

	err = json.Unmarshal(data, &varPalActivityLogs)

	if err != nil {
		return err
	}

	*o = PalActivityLogs(varPalActivityLogs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalActivityLogs struct {
	value *PalActivityLogs
	isSet bool
}

func (v NullablePalActivityLogs) Get() *PalActivityLogs {
	return v.value
}

func (v *NullablePalActivityLogs) Set(val *PalActivityLogs) {
	v.value = val
	v.isSet = true
}

func (v NullablePalActivityLogs) IsSet() bool {
	return v.isSet
}

func (v *NullablePalActivityLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalActivityLogs(val *PalActivityLogs) *NullablePalActivityLogs {
	return &NullablePalActivityLogs{value: val, isSet: true}
}

func (v NullablePalActivityLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalActivityLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


