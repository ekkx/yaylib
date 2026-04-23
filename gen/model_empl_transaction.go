
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplTransaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplTransaction{}

// EmplTransaction struct for EmplTransaction
type EmplTransaction struct {
	Amount NullableString `json:"amount,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	OccurredAt NullableInt64 `json:"occurred_at,omitempty"`
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplTransaction EmplTransaction

// NewEmplTransaction instantiates a new EmplTransaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplTransaction() *EmplTransaction {
	this := EmplTransaction{}
	return &this
}

// NewEmplTransactionWithDefaults instantiates a new EmplTransaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplTransactionWithDefaults() *EmplTransaction {
	this := EmplTransaction{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTransaction) GetAmount() string {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret string
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTransaction) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *EmplTransaction) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableString and assigns it to the Amount field.
func (o *EmplTransaction) SetAmount(v string) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *EmplTransaction) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *EmplTransaction) UnsetAmount() {
	o.Amount.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTransaction) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTransaction) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *EmplTransaction) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *EmplTransaction) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *EmplTransaction) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *EmplTransaction) UnsetId() {
	o.Id.Unset()
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTransaction) GetOccurredAt() int64 {
	if o == nil || IsNil(o.OccurredAt.Get()) {
		var ret int64
		return ret
	}
	return *o.OccurredAt.Get()
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTransaction) GetOccurredAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.OccurredAt.Get(), o.OccurredAt.IsSet()
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *EmplTransaction) HasOccurredAt() bool {
	if o != nil && o.OccurredAt.IsSet() {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given NullableInt64 and assigns it to the OccurredAt field.
func (o *EmplTransaction) SetOccurredAt(v int64) {
	o.OccurredAt.Set(&v)
}
// SetOccurredAtNil sets the value for OccurredAt to be an explicit nil
func (o *EmplTransaction) SetOccurredAtNil() {
	o.OccurredAt.Set(nil)
}

// UnsetOccurredAt ensures that no value is present for OccurredAt, not even an explicit nil
func (o *EmplTransaction) UnsetOccurredAt() {
	o.OccurredAt.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplTransaction) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplTransaction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *EmplTransaction) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *EmplTransaction) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *EmplTransaction) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *EmplTransaction) UnsetType() {
	o.Type.Unset()
}

func (o EmplTransaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplTransaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.OccurredAt.IsSet() {
		toSerialize["occurred_at"] = o.OccurredAt.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplTransaction) UnmarshalJSON(data []byte) (err error) {
	varEmplTransaction := _EmplTransaction{}

	err = json.Unmarshal(data, &varEmplTransaction)

	if err != nil {
		return err
	}

	*o = EmplTransaction(varEmplTransaction)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "amount")
		delete(additionalProperties, "id")
		delete(additionalProperties, "occurred_at")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplTransaction struct {
	value *EmplTransaction
	isSet bool
}

func (v NullableEmplTransaction) Get() *EmplTransaction {
	return v.value
}

func (v *NullableEmplTransaction) Set(val *EmplTransaction) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplTransaction) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplTransaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplTransaction(val *EmplTransaction) *NullableEmplTransaction {
	return &NullableEmplTransaction{value: val, isSet: true}
}

func (v NullableEmplTransaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplTransaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


