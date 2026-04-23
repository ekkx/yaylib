
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the YayPoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &YayPoints{}

// YayPoints struct for YayPoints
type YayPoints struct {
	Total NullableInt64 `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _YayPoints YayPoints

// NewYayPoints instantiates a new YayPoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewYayPoints() *YayPoints {
	this := YayPoints{}
	return &this
}

// NewYayPointsWithDefaults instantiates a new YayPoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewYayPointsWithDefaults() *YayPoints {
	this := YayPoints{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *YayPoints) GetTotal() int64 {
	if o == nil || IsNil(o.Total.Get()) {
		var ret int64
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *YayPoints) GetTotalOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *YayPoints) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableInt64 and assigns it to the Total field.
func (o *YayPoints) SetTotal(v int64) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *YayPoints) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *YayPoints) UnsetTotal() {
	o.Total.Unset()
}

func (o YayPoints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o YayPoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *YayPoints) UnmarshalJSON(data []byte) (err error) {
	varYayPoints := _YayPoints{}

	err = json.Unmarshal(data, &varYayPoints)

	if err != nil {
		return err
	}

	*o = YayPoints(varYayPoints)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableYayPoints struct {
	value *YayPoints
	isSet bool
}

func (v NullableYayPoints) Get() *YayPoints {
	return v.value
}

func (v *NullableYayPoints) Set(val *YayPoints) {
	v.value = val
	v.isSet = true
}

func (v NullableYayPoints) IsSet() bool {
	return v.isSet
}

func (v *NullableYayPoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableYayPoints(val *YayPoints) *NullableYayPoints {
	return &NullableYayPoints{value: val, isSet: true}
}

func (v NullableYayPoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableYayPoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


