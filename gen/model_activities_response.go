
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ActivitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivitiesResponse{}

// ActivitiesResponse struct for ActivitiesResponse
type ActivitiesResponse struct {
	Activities []Activity `json:"activities,omitempty"`
	LastTimestamp NullableInt64 `json:"last_timestamp,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ActivitiesResponse ActivitiesResponse

// NewActivitiesResponse instantiates a new ActivitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivitiesResponse() *ActivitiesResponse {
	this := ActivitiesResponse{}
	return &this
}

// NewActivitiesResponseWithDefaults instantiates a new ActivitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivitiesResponseWithDefaults() *ActivitiesResponse {
	this := ActivitiesResponse{}
	return &this
}

// GetActivities returns the Activities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivitiesResponse) GetActivities() []Activity {
	if o == nil {
		var ret []Activity
		return ret
	}
	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivitiesResponse) GetActivitiesOk() ([]Activity, bool) {
	if o == nil || IsNil(o.Activities) {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *ActivitiesResponse) HasActivities() bool {
	if o != nil && !IsNil(o.Activities) {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []Activity and assigns it to the Activities field.
func (o *ActivitiesResponse) SetActivities(v []Activity) {
	o.Activities = v
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ActivitiesResponse) GetLastTimestamp() int64 {
	if o == nil || IsNil(o.LastTimestamp.Get()) {
		var ret int64
		return ret
	}
	return *o.LastTimestamp.Get()
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivitiesResponse) GetLastTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTimestamp.Get(), o.LastTimestamp.IsSet()
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *ActivitiesResponse) HasLastTimestamp() bool {
	if o != nil && o.LastTimestamp.IsSet() {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given NullableInt64 and assigns it to the LastTimestamp field.
func (o *ActivitiesResponse) SetLastTimestamp(v int64) {
	o.LastTimestamp.Set(&v)
}
// SetLastTimestampNil sets the value for LastTimestamp to be an explicit nil
func (o *ActivitiesResponse) SetLastTimestampNil() {
	o.LastTimestamp.Set(nil)
}

// UnsetLastTimestamp ensures that no value is present for LastTimestamp, not even an explicit nil
func (o *ActivitiesResponse) UnsetLastTimestamp() {
	o.LastTimestamp.Unset()
}

func (o ActivitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Activities != nil {
		toSerialize["activities"] = o.Activities
	}
	if o.LastTimestamp.IsSet() {
		toSerialize["last_timestamp"] = o.LastTimestamp.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ActivitiesResponse) UnmarshalJSON(data []byte) (err error) {
	varActivitiesResponse := _ActivitiesResponse{}

	err = json.Unmarshal(data, &varActivitiesResponse)

	if err != nil {
		return err
	}

	*o = ActivitiesResponse(varActivitiesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "activities")
		delete(additionalProperties, "last_timestamp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableActivitiesResponse struct {
	value *ActivitiesResponse
	isSet bool
}

func (v NullableActivitiesResponse) Get() *ActivitiesResponse {
	return v.value
}

func (v *NullableActivitiesResponse) Set(val *ActivitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableActivitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableActivitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivitiesResponse(val *ActivitiesResponse) *NullableActivitiesResponse {
	return &NullableActivitiesResponse{value: val, isSet: true}
}

func (v NullableActivitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


