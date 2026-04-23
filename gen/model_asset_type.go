
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AssetType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetType{}

// AssetType struct for AssetType
type AssetType struct {
	ApiValue NullableString `json:"api_value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetType AssetType

// NewAssetType instantiates a new AssetType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetType() *AssetType {
	this := AssetType{}
	return &this
}

// NewAssetTypeWithDefaults instantiates a new AssetType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetTypeWithDefaults() *AssetType {
	this := AssetType{}
	return &this
}

// GetApiValue returns the ApiValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetType) GetApiValue() string {
	if o == nil || IsNil(o.ApiValue.Get()) {
		var ret string
		return ret
	}
	return *o.ApiValue.Get()
}

// GetApiValueOk returns a tuple with the ApiValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetType) GetApiValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiValue.Get(), o.ApiValue.IsSet()
}

// HasApiValue returns a boolean if a field has been set.
func (o *AssetType) HasApiValue() bool {
	if o != nil && o.ApiValue.IsSet() {
		return true
	}

	return false
}

// SetApiValue gets a reference to the given NullableString and assigns it to the ApiValue field.
func (o *AssetType) SetApiValue(v string) {
	o.ApiValue.Set(&v)
}
// SetApiValueNil sets the value for ApiValue to be an explicit nil
func (o *AssetType) SetApiValueNil() {
	o.ApiValue.Set(nil)
}

// UnsetApiValue ensures that no value is present for ApiValue, not even an explicit nil
func (o *AssetType) UnsetApiValue() {
	o.ApiValue.Unset()
}

func (o AssetType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiValue.IsSet() {
		toSerialize["api_value"] = o.ApiValue.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetType) UnmarshalJSON(data []byte) (err error) {
	varAssetType := _AssetType{}

	err = json.Unmarshal(data, &varAssetType)

	if err != nil {
		return err
	}

	*o = AssetType(varAssetType)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetType struct {
	value *AssetType
	isSet bool
}

func (v NullableAssetType) Get() *AssetType {
	return v.value
}

func (v *NullableAssetType) Set(val *AssetType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetType(val *AssetType) *NullableAssetType {
	return &NullableAssetType{value: val, isSet: true}
}

func (v NullableAssetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


