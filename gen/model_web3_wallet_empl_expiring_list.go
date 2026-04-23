
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletEmplExpiringList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletEmplExpiringList{}

// Web3WalletEmplExpiringList struct for Web3WalletEmplExpiringList
type Web3WalletEmplExpiringList struct {
	RegularExpiring []Web3WalletEmplExpiring `json:"regular_expiring,omitempty"`
	RewardedExpiring []Web3WalletEmplExpiring `json:"rewarded_expiring,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletEmplExpiringList Web3WalletEmplExpiringList

// NewWeb3WalletEmplExpiringList instantiates a new Web3WalletEmplExpiringList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletEmplExpiringList() *Web3WalletEmplExpiringList {
	this := Web3WalletEmplExpiringList{}
	return &this
}

// NewWeb3WalletEmplExpiringListWithDefaults instantiates a new Web3WalletEmplExpiringList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletEmplExpiringListWithDefaults() *Web3WalletEmplExpiringList {
	this := Web3WalletEmplExpiringList{}
	return &this
}

// GetRegularExpiring returns the RegularExpiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletEmplExpiringList) GetRegularExpiring() []Web3WalletEmplExpiring {
	if o == nil {
		var ret []Web3WalletEmplExpiring
		return ret
	}
	return o.RegularExpiring
}

// GetRegularExpiringOk returns a tuple with the RegularExpiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletEmplExpiringList) GetRegularExpiringOk() ([]Web3WalletEmplExpiring, bool) {
	if o == nil || IsNil(o.RegularExpiring) {
		return nil, false
	}
	return o.RegularExpiring, true
}

// HasRegularExpiring returns a boolean if a field has been set.
func (o *Web3WalletEmplExpiringList) HasRegularExpiring() bool {
	if o != nil && !IsNil(o.RegularExpiring) {
		return true
	}

	return false
}

// SetRegularExpiring gets a reference to the given []Web3WalletEmplExpiring and assigns it to the RegularExpiring field.
func (o *Web3WalletEmplExpiringList) SetRegularExpiring(v []Web3WalletEmplExpiring) {
	o.RegularExpiring = v
}

// GetRewardedExpiring returns the RewardedExpiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletEmplExpiringList) GetRewardedExpiring() []Web3WalletEmplExpiring {
	if o == nil {
		var ret []Web3WalletEmplExpiring
		return ret
	}
	return o.RewardedExpiring
}

// GetRewardedExpiringOk returns a tuple with the RewardedExpiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletEmplExpiringList) GetRewardedExpiringOk() ([]Web3WalletEmplExpiring, bool) {
	if o == nil || IsNil(o.RewardedExpiring) {
		return nil, false
	}
	return o.RewardedExpiring, true
}

// HasRewardedExpiring returns a boolean if a field has been set.
func (o *Web3WalletEmplExpiringList) HasRewardedExpiring() bool {
	if o != nil && !IsNil(o.RewardedExpiring) {
		return true
	}

	return false
}

// SetRewardedExpiring gets a reference to the given []Web3WalletEmplExpiring and assigns it to the RewardedExpiring field.
func (o *Web3WalletEmplExpiringList) SetRewardedExpiring(v []Web3WalletEmplExpiring) {
	o.RewardedExpiring = v
}

func (o Web3WalletEmplExpiringList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletEmplExpiringList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RegularExpiring != nil {
		toSerialize["regular_expiring"] = o.RegularExpiring
	}
	if o.RewardedExpiring != nil {
		toSerialize["rewarded_expiring"] = o.RewardedExpiring
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletEmplExpiringList) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletEmplExpiringList := _Web3WalletEmplExpiringList{}

	err = json.Unmarshal(data, &varWeb3WalletEmplExpiringList)

	if err != nil {
		return err
	}

	*o = Web3WalletEmplExpiringList(varWeb3WalletEmplExpiringList)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "regular_expiring")
		delete(additionalProperties, "rewarded_expiring")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletEmplExpiringList struct {
	value *Web3WalletEmplExpiringList
	isSet bool
}

func (v NullableWeb3WalletEmplExpiringList) Get() *Web3WalletEmplExpiringList {
	return v.value
}

func (v *NullableWeb3WalletEmplExpiringList) Set(val *Web3WalletEmplExpiringList) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletEmplExpiringList) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletEmplExpiringList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletEmplExpiringList(val *Web3WalletEmplExpiringList) *NullableWeb3WalletEmplExpiringList {
	return &NullableWeb3WalletEmplExpiringList{value: val, isSet: true}
}

func (v NullableWeb3WalletEmplExpiringList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletEmplExpiringList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


