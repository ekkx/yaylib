
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RegisterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterRequest{}

// RegisterRequest struct for RegisterRequest
type RegisterRequest struct {
	SignedDataType NullableString `json:"signed_data_type,omitempty"`
	WalletAddress NullableString `json:"wallet_address,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegisterRequest RegisterRequest

// NewRegisterRequest instantiates a new RegisterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterRequest() *RegisterRequest {
	this := RegisterRequest{}
	return &this
}

// NewRegisterRequestWithDefaults instantiates a new RegisterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterRequestWithDefaults() *RegisterRequest {
	this := RegisterRequest{}
	return &this
}

// GetSignedDataType returns the SignedDataType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterRequest) GetSignedDataType() string {
	if o == nil || IsNil(o.SignedDataType.Get()) {
		var ret string
		return ret
	}
	return *o.SignedDataType.Get()
}

// GetSignedDataTypeOk returns a tuple with the SignedDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterRequest) GetSignedDataTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignedDataType.Get(), o.SignedDataType.IsSet()
}

// HasSignedDataType returns a boolean if a field has been set.
func (o *RegisterRequest) HasSignedDataType() bool {
	if o != nil && o.SignedDataType.IsSet() {
		return true
	}

	return false
}

// SetSignedDataType gets a reference to the given NullableString and assigns it to the SignedDataType field.
func (o *RegisterRequest) SetSignedDataType(v string) {
	o.SignedDataType.Set(&v)
}
// SetSignedDataTypeNil sets the value for SignedDataType to be an explicit nil
func (o *RegisterRequest) SetSignedDataTypeNil() {
	o.SignedDataType.Set(nil)
}

// UnsetSignedDataType ensures that no value is present for SignedDataType, not even an explicit nil
func (o *RegisterRequest) UnsetSignedDataType() {
	o.SignedDataType.Unset()
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterRequest) GetWalletAddress() string {
	if o == nil || IsNil(o.WalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.WalletAddress.Get()
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterRequest) GetWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WalletAddress.Get(), o.WalletAddress.IsSet()
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *RegisterRequest) HasWalletAddress() bool {
	if o != nil && o.WalletAddress.IsSet() {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given NullableString and assigns it to the WalletAddress field.
func (o *RegisterRequest) SetWalletAddress(v string) {
	o.WalletAddress.Set(&v)
}
// SetWalletAddressNil sets the value for WalletAddress to be an explicit nil
func (o *RegisterRequest) SetWalletAddressNil() {
	o.WalletAddress.Set(nil)
}

// UnsetWalletAddress ensures that no value is present for WalletAddress, not even an explicit nil
func (o *RegisterRequest) UnsetWalletAddress() {
	o.WalletAddress.Unset()
}

func (o RegisterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.SignedDataType.IsSet() {
		toSerialize["signed_data_type"] = o.SignedDataType.Get()
	}
	if o.WalletAddress.IsSet() {
		toSerialize["wallet_address"] = o.WalletAddress.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegisterRequest) UnmarshalJSON(data []byte) (err error) {
	varRegisterRequest := _RegisterRequest{}

	err = json.Unmarshal(data, &varRegisterRequest)

	if err != nil {
		return err
	}

	*o = RegisterRequest(varRegisterRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "signed_data_type")
		delete(additionalProperties, "wallet_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegisterRequest struct {
	value *RegisterRequest
	isSet bool
}

func (v NullableRegisterRequest) Get() *RegisterRequest {
	return v.value
}

func (v *NullableRegisterRequest) Set(val *RegisterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterRequest(val *RegisterRequest) *NullableRegisterRequest {
	return &NullableRegisterRequest{value: val, isSet: true}
}

func (v NullableRegisterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


