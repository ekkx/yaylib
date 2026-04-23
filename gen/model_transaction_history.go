
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TransactionHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionHistory{}

// TransactionHistory struct for TransactionHistory
type TransactionHistory struct {
	Cursor NullableString `json:"cursor,omitempty"`
	Result []Web3WalletTransactionHistory `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionHistory TransactionHistory

// NewTransactionHistory instantiates a new TransactionHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionHistory() *TransactionHistory {
	this := TransactionHistory{}
	return &this
}

// NewTransactionHistoryWithDefaults instantiates a new TransactionHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionHistoryWithDefaults() *TransactionHistory {
	this := TransactionHistory{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionHistory) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionHistory) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *TransactionHistory) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *TransactionHistory) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *TransactionHistory) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *TransactionHistory) UnsetCursor() {
	o.Cursor.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionHistory) GetResult() []Web3WalletTransactionHistory {
	if o == nil {
		var ret []Web3WalletTransactionHistory
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionHistory) GetResultOk() ([]Web3WalletTransactionHistory, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TransactionHistory) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []Web3WalletTransactionHistory and assigns it to the Result field.
func (o *TransactionHistory) SetResult(v []Web3WalletTransactionHistory) {
	o.Result = v
}

func (o TransactionHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionHistory) UnmarshalJSON(data []byte) (err error) {
	varTransactionHistory := _TransactionHistory{}

	err = json.Unmarshal(data, &varTransactionHistory)

	if err != nil {
		return err
	}

	*o = TransactionHistory(varTransactionHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionHistory struct {
	value *TransactionHistory
	isSet bool
}

func (v NullableTransactionHistory) Get() *TransactionHistory {
	return v.value
}

func (v *NullableTransactionHistory) Set(val *TransactionHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionHistory(val *TransactionHistory) *NullableTransactionHistory {
	return &NullableTransactionHistory{value: val, isSet: true}
}

func (v NullableTransactionHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


