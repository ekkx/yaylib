
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Web3WalletBlockStoreEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Web3WalletBlockStoreEntry{}

// Web3WalletBlockStoreEntry struct for Web3WalletBlockStoreEntry
type Web3WalletBlockStoreEntry struct {
	Items []Web3WalletBlockStoreItem `json:"items,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Web3WalletBlockStoreEntry Web3WalletBlockStoreEntry

// NewWeb3WalletBlockStoreEntry instantiates a new Web3WalletBlockStoreEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWeb3WalletBlockStoreEntry() *Web3WalletBlockStoreEntry {
	this := Web3WalletBlockStoreEntry{}
	return &this
}

// NewWeb3WalletBlockStoreEntryWithDefaults instantiates a new Web3WalletBlockStoreEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWeb3WalletBlockStoreEntryWithDefaults() *Web3WalletBlockStoreEntry {
	this := Web3WalletBlockStoreEntry{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Web3WalletBlockStoreEntry) GetItems() []Web3WalletBlockStoreItem {
	if o == nil {
		var ret []Web3WalletBlockStoreItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Web3WalletBlockStoreEntry) GetItemsOk() ([]Web3WalletBlockStoreItem, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Web3WalletBlockStoreEntry) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Web3WalletBlockStoreItem and assigns it to the Items field.
func (o *Web3WalletBlockStoreEntry) SetItems(v []Web3WalletBlockStoreItem) {
	o.Items = v
}

func (o Web3WalletBlockStoreEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Web3WalletBlockStoreEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Web3WalletBlockStoreEntry) UnmarshalJSON(data []byte) (err error) {
	varWeb3WalletBlockStoreEntry := _Web3WalletBlockStoreEntry{}

	err = json.Unmarshal(data, &varWeb3WalletBlockStoreEntry)

	if err != nil {
		return err
	}

	*o = Web3WalletBlockStoreEntry(varWeb3WalletBlockStoreEntry)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWeb3WalletBlockStoreEntry struct {
	value *Web3WalletBlockStoreEntry
	isSet bool
}

func (v NullableWeb3WalletBlockStoreEntry) Get() *Web3WalletBlockStoreEntry {
	return v.value
}

func (v *NullableWeb3WalletBlockStoreEntry) Set(val *Web3WalletBlockStoreEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableWeb3WalletBlockStoreEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableWeb3WalletBlockStoreEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWeb3WalletBlockStoreEntry(val *Web3WalletBlockStoreEntry) *NullableWeb3WalletBlockStoreEntry {
	return &NullableWeb3WalletBlockStoreEntry{value: val, isSet: true}
}

func (v NullableWeb3WalletBlockStoreEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWeb3WalletBlockStoreEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


