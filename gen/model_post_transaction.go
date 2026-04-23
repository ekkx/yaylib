
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostTransaction{}

// PostTransaction struct for PostTransaction
type PostTransaction struct {
	ReceiptJsonStr NullableString `json:"receipt_json_str,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostTransaction PostTransaction

// NewPostTransaction instantiates a new PostTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostTransaction() *PostTransaction {
	this := PostTransaction{}
	return &this
}

// NewPostTransactionWithDefaults instantiates a new PostTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostTransactionWithDefaults() *PostTransaction {
	this := PostTransaction{}
	return &this
}

// GetReceiptJsonStr returns the ReceiptJsonStr field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostTransaction) GetReceiptJsonStr() string {
	if o == nil || IsNil(o.ReceiptJsonStr.Get()) {
		var ret string
		return ret
	}
	return *o.ReceiptJsonStr.Get()
}

// GetReceiptJsonStrOk returns a tuple with the ReceiptJsonStr field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostTransaction) GetReceiptJsonStrOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReceiptJsonStr.Get(), o.ReceiptJsonStr.IsSet()
}

// HasReceiptJsonStr returns a boolean if a field has been set.
func (o *PostTransaction) HasReceiptJsonStr() bool {
	if o != nil && o.ReceiptJsonStr.IsSet() {
		return true
	}

	return false
}

// SetReceiptJsonStr gets a reference to the given NullableString and assigns it to the ReceiptJsonStr field.
func (o *PostTransaction) SetReceiptJsonStr(v string) {
	o.ReceiptJsonStr.Set(&v)
}
// SetReceiptJsonStrNil sets the value for ReceiptJsonStr to be an explicit nil
func (o *PostTransaction) SetReceiptJsonStrNil() {
	o.ReceiptJsonStr.Set(nil)
}

// UnsetReceiptJsonStr ensures that no value is present for ReceiptJsonStr, not even an explicit nil
func (o *PostTransaction) UnsetReceiptJsonStr() {
	o.ReceiptJsonStr.Unset()
}

func (o PostTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ReceiptJsonStr.IsSet() {
		toSerialize["receipt_json_str"] = o.ReceiptJsonStr.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostTransaction) UnmarshalJSON(data []byte) (err error) {
	varPostTransaction := _PostTransaction{}

	err = json.Unmarshal(data, &varPostTransaction)

	if err != nil {
		return err
	}

	*o = PostTransaction(varPostTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "receipt_json_str")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostTransaction struct {
	value *PostTransaction
	isSet bool
}

func (v NullablePostTransaction) Get() *PostTransaction {
	return v.value
}

func (v *NullablePostTransaction) Set(val *PostTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullablePostTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullablePostTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostTransaction(val *PostTransaction) *NullablePostTransaction {
	return &NullablePostTransaction{value: val, isSet: true}
}

func (v NullablePostTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


