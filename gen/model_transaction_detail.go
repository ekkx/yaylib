
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TransactionDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransactionDetail{}

// TransactionDetail struct for TransactionDetail
type TransactionDetail struct {
	Result NullableResult `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TransactionDetail TransactionDetail

// NewTransactionDetail instantiates a new TransactionDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionDetail() *TransactionDetail {
	this := TransactionDetail{}
	return &this
}

// NewTransactionDetailWithDefaults instantiates a new TransactionDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionDetailWithDefaults() *TransactionDetail {
	this := TransactionDetail{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionDetail) GetResult() Result {
	if o == nil || IsNil(o.Result.Get()) {
		var ret Result
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionDetail) GetResultOk() (*Result, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *TransactionDetail) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableResult and assigns it to the Result field.
func (o *TransactionDetail) SetResult(v Result) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *TransactionDetail) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *TransactionDetail) UnsetResult() {
	o.Result.Unset()
}

func (o TransactionDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransactionDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TransactionDetail) UnmarshalJSON(data []byte) (err error) {
	varTransactionDetail := _TransactionDetail{}

	err = json.Unmarshal(data, &varTransactionDetail)

	if err != nil {
		return err
	}

	*o = TransactionDetail(varTransactionDetail)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionDetail struct {
	value *TransactionDetail
	isSet bool
}

func (v NullableTransactionDetail) Get() *TransactionDetail {
	return v.value
}

func (v *NullableTransactionDetail) Set(val *TransactionDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionDetail(val *TransactionDetail) *NullableTransactionDetail {
	return &NullableTransactionDetail{value: val, isSet: true}
}

func (v NullableTransactionDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


