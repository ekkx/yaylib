
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Balance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Balance{}

// Balance struct for Balance
type Balance struct {
	Egg NullableString `json:"egg,omitempty"`
	Empl NullableString `json:"empl,omitempty"`
	L1Native NullableString `json:"l1_native,omitempty"`
	L1Yay NullableString `json:"l1_yay,omitempty"`
	Native NullableString `json:"native,omitempty"`
	Pal NullableString `json:"pal,omitempty"`
	Yay NullableString `json:"yay,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Balance Balance

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance() *Balance {
	this := Balance{}
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetEgg returns the Egg field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetEgg() string {
	if o == nil || IsNil(o.Egg.Get()) {
		var ret string
		return ret
	}
	return *o.Egg.Get()
}

// GetEggOk returns a tuple with the Egg field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetEggOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Egg.Get(), o.Egg.IsSet()
}

// HasEgg returns a boolean if a field has been set.
func (o *Balance) HasEgg() bool {
	if o != nil && o.Egg.IsSet() {
		return true
	}

	return false
}

// SetEgg gets a reference to the given NullableString and assigns it to the Egg field.
func (o *Balance) SetEgg(v string) {
	o.Egg.Set(&v)
}
// SetEggNil sets the value for Egg to be an explicit nil
func (o *Balance) SetEggNil() {
	o.Egg.Set(nil)
}

// UnsetEgg ensures that no value is present for Egg, not even an explicit nil
func (o *Balance) UnsetEgg() {
	o.Egg.Unset()
}

// GetEmpl returns the Empl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetEmpl() string {
	if o == nil || IsNil(o.Empl.Get()) {
		var ret string
		return ret
	}
	return *o.Empl.Get()
}

// GetEmplOk returns a tuple with the Empl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetEmplOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Empl.Get(), o.Empl.IsSet()
}

// HasEmpl returns a boolean if a field has been set.
func (o *Balance) HasEmpl() bool {
	if o != nil && o.Empl.IsSet() {
		return true
	}

	return false
}

// SetEmpl gets a reference to the given NullableString and assigns it to the Empl field.
func (o *Balance) SetEmpl(v string) {
	o.Empl.Set(&v)
}
// SetEmplNil sets the value for Empl to be an explicit nil
func (o *Balance) SetEmplNil() {
	o.Empl.Set(nil)
}

// UnsetEmpl ensures that no value is present for Empl, not even an explicit nil
func (o *Balance) UnsetEmpl() {
	o.Empl.Unset()
}

// GetL1Native returns the L1Native field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetL1Native() string {
	if o == nil || IsNil(o.L1Native.Get()) {
		var ret string
		return ret
	}
	return *o.L1Native.Get()
}

// GetL1NativeOk returns a tuple with the L1Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetL1NativeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1Native.Get(), o.L1Native.IsSet()
}

// HasL1Native returns a boolean if a field has been set.
func (o *Balance) HasL1Native() bool {
	if o != nil && o.L1Native.IsSet() {
		return true
	}

	return false
}

// SetL1Native gets a reference to the given NullableString and assigns it to the L1Native field.
func (o *Balance) SetL1Native(v string) {
	o.L1Native.Set(&v)
}
// SetL1NativeNil sets the value for L1Native to be an explicit nil
func (o *Balance) SetL1NativeNil() {
	o.L1Native.Set(nil)
}

// UnsetL1Native ensures that no value is present for L1Native, not even an explicit nil
func (o *Balance) UnsetL1Native() {
	o.L1Native.Unset()
}

// GetL1Yay returns the L1Yay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetL1Yay() string {
	if o == nil || IsNil(o.L1Yay.Get()) {
		var ret string
		return ret
	}
	return *o.L1Yay.Get()
}

// GetL1YayOk returns a tuple with the L1Yay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetL1YayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.L1Yay.Get(), o.L1Yay.IsSet()
}

// HasL1Yay returns a boolean if a field has been set.
func (o *Balance) HasL1Yay() bool {
	if o != nil && o.L1Yay.IsSet() {
		return true
	}

	return false
}

// SetL1Yay gets a reference to the given NullableString and assigns it to the L1Yay field.
func (o *Balance) SetL1Yay(v string) {
	o.L1Yay.Set(&v)
}
// SetL1YayNil sets the value for L1Yay to be an explicit nil
func (o *Balance) SetL1YayNil() {
	o.L1Yay.Set(nil)
}

// UnsetL1Yay ensures that no value is present for L1Yay, not even an explicit nil
func (o *Balance) UnsetL1Yay() {
	o.L1Yay.Unset()
}

// GetNative returns the Native field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetNative() string {
	if o == nil || IsNil(o.Native.Get()) {
		var ret string
		return ret
	}
	return *o.Native.Get()
}

// GetNativeOk returns a tuple with the Native field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetNativeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Native.Get(), o.Native.IsSet()
}

// HasNative returns a boolean if a field has been set.
func (o *Balance) HasNative() bool {
	if o != nil && o.Native.IsSet() {
		return true
	}

	return false
}

// SetNative gets a reference to the given NullableString and assigns it to the Native field.
func (o *Balance) SetNative(v string) {
	o.Native.Set(&v)
}
// SetNativeNil sets the value for Native to be an explicit nil
func (o *Balance) SetNativeNil() {
	o.Native.Set(nil)
}

// UnsetNative ensures that no value is present for Native, not even an explicit nil
func (o *Balance) UnsetNative() {
	o.Native.Unset()
}

// GetPal returns the Pal field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetPal() string {
	if o == nil || IsNil(o.Pal.Get()) {
		var ret string
		return ret
	}
	return *o.Pal.Get()
}

// GetPalOk returns a tuple with the Pal field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetPalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Pal.Get(), o.Pal.IsSet()
}

// HasPal returns a boolean if a field has been set.
func (o *Balance) HasPal() bool {
	if o != nil && o.Pal.IsSet() {
		return true
	}

	return false
}

// SetPal gets a reference to the given NullableString and assigns it to the Pal field.
func (o *Balance) SetPal(v string) {
	o.Pal.Set(&v)
}
// SetPalNil sets the value for Pal to be an explicit nil
func (o *Balance) SetPalNil() {
	o.Pal.Set(nil)
}

// UnsetPal ensures that no value is present for Pal, not even an explicit nil
func (o *Balance) UnsetPal() {
	o.Pal.Unset()
}

// GetYay returns the Yay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Balance) GetYay() string {
	if o == nil || IsNil(o.Yay.Get()) {
		var ret string
		return ret
	}
	return *o.Yay.Get()
}

// GetYayOk returns a tuple with the Yay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Balance) GetYayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Yay.Get(), o.Yay.IsSet()
}

// HasYay returns a boolean if a field has been set.
func (o *Balance) HasYay() bool {
	if o != nil && o.Yay.IsSet() {
		return true
	}

	return false
}

// SetYay gets a reference to the given NullableString and assigns it to the Yay field.
func (o *Balance) SetYay(v string) {
	o.Yay.Set(&v)
}
// SetYayNil sets the value for Yay to be an explicit nil
func (o *Balance) SetYayNil() {
	o.Yay.Set(nil)
}

// UnsetYay ensures that no value is present for Yay, not even an explicit nil
func (o *Balance) UnsetYay() {
	o.Yay.Unset()
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Balance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Egg.IsSet() {
		toSerialize["egg"] = o.Egg.Get()
	}
	if o.Empl.IsSet() {
		toSerialize["empl"] = o.Empl.Get()
	}
	if o.L1Native.IsSet() {
		toSerialize["l1_native"] = o.L1Native.Get()
	}
	if o.L1Yay.IsSet() {
		toSerialize["l1_yay"] = o.L1Yay.Get()
	}
	if o.Native.IsSet() {
		toSerialize["native"] = o.Native.Get()
	}
	if o.Pal.IsSet() {
		toSerialize["pal"] = o.Pal.Get()
	}
	if o.Yay.IsSet() {
		toSerialize["yay"] = o.Yay.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Balance) UnmarshalJSON(data []byte) (err error) {
	varBalance := _Balance{}

	err = json.Unmarshal(data, &varBalance)

	if err != nil {
		return err
	}

	*o = Balance(varBalance)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "egg")
		delete(additionalProperties, "empl")
		delete(additionalProperties, "l1_native")
		delete(additionalProperties, "l1_yay")
		delete(additionalProperties, "native")
		delete(additionalProperties, "pal")
		delete(additionalProperties, "yay")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


