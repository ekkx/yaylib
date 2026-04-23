
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftExchangeHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftExchangeHistory{}

// GiftExchangeHistory struct for GiftExchangeHistory
type GiftExchangeHistory struct {
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Id NullableInt32 `json:"id,omitempty"`
	Items []Item `json:"items,omitempty"`
	Quantity NullableInt32 `json:"quantity,omitempty"`
	RewardedAmount NullableInt64 `json:"rewarded_amount,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftExchangeHistory GiftExchangeHistory

// NewGiftExchangeHistory instantiates a new GiftExchangeHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftExchangeHistory() *GiftExchangeHistory {
	this := GiftExchangeHistory{}
	return &this
}

// NewGiftExchangeHistoryWithDefaults instantiates a new GiftExchangeHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftExchangeHistoryWithDefaults() *GiftExchangeHistory {
	this := GiftExchangeHistory{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *GiftExchangeHistory) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *GiftExchangeHistory) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *GiftExchangeHistory) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetId() int32 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int32
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt32 and assigns it to the Id field.
func (o *GiftExchangeHistory) SetId(v int32) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *GiftExchangeHistory) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *GiftExchangeHistory) UnsetId() {
	o.Id.Unset()
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetItems() []Item {
	if o == nil {
		var ret []Item
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetItemsOk() ([]Item, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Item and assigns it to the Items field.
func (o *GiftExchangeHistory) SetItems(v []Item) {
	o.Items = v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetQuantity() int32 {
	if o == nil || IsNil(o.Quantity.Get()) {
		var ret int32
		return ret
	}
	return *o.Quantity.Get()
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetQuantityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Quantity.Get(), o.Quantity.IsSet()
}

// HasQuantity returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasQuantity() bool {
	if o != nil && o.Quantity.IsSet() {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given NullableInt32 and assigns it to the Quantity field.
func (o *GiftExchangeHistory) SetQuantity(v int32) {
	o.Quantity.Set(&v)
}
// SetQuantityNil sets the value for Quantity to be an explicit nil
func (o *GiftExchangeHistory) SetQuantityNil() {
	o.Quantity.Set(nil)
}

// UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
func (o *GiftExchangeHistory) UnsetQuantity() {
	o.Quantity.Unset()
}

// GetRewardedAmount returns the RewardedAmount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetRewardedAmount() int64 {
	if o == nil || IsNil(o.RewardedAmount.Get()) {
		var ret int64
		return ret
	}
	return *o.RewardedAmount.Get()
}

// GetRewardedAmountOk returns a tuple with the RewardedAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetRewardedAmountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RewardedAmount.Get(), o.RewardedAmount.IsSet()
}

// HasRewardedAmount returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasRewardedAmount() bool {
	if o != nil && o.RewardedAmount.IsSet() {
		return true
	}

	return false
}

// SetRewardedAmount gets a reference to the given NullableInt64 and assigns it to the RewardedAmount field.
func (o *GiftExchangeHistory) SetRewardedAmount(v int64) {
	o.RewardedAmount.Set(&v)
}
// SetRewardedAmountNil sets the value for RewardedAmount to be an explicit nil
func (o *GiftExchangeHistory) SetRewardedAmountNil() {
	o.RewardedAmount.Set(nil)
}

// UnsetRewardedAmount ensures that no value is present for RewardedAmount, not even an explicit nil
func (o *GiftExchangeHistory) UnsetRewardedAmount() {
	o.RewardedAmount.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *GiftExchangeHistory) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *GiftExchangeHistory) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *GiftExchangeHistory) UnsetStatus() {
	o.Status.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftExchangeHistory) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftExchangeHistory) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *GiftExchangeHistory) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *GiftExchangeHistory) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *GiftExchangeHistory) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *GiftExchangeHistory) UnsetUuid() {
	o.Uuid.Unset()
}

func (o GiftExchangeHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftExchangeHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Quantity.IsSet() {
		toSerialize["quantity"] = o.Quantity.Get()
	}
	if o.RewardedAmount.IsSet() {
		toSerialize["rewarded_amount"] = o.RewardedAmount.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftExchangeHistory) UnmarshalJSON(data []byte) (err error) {
	varGiftExchangeHistory := _GiftExchangeHistory{}

	err = json.Unmarshal(data, &varGiftExchangeHistory)

	if err != nil {
		return err
	}

	*o = GiftExchangeHistory(varGiftExchangeHistory)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "items")
		delete(additionalProperties, "quantity")
		delete(additionalProperties, "rewarded_amount")
		delete(additionalProperties, "status")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftExchangeHistory struct {
	value *GiftExchangeHistory
	isSet bool
}

func (v NullableGiftExchangeHistory) Get() *GiftExchangeHistory {
	return v.value
}

func (v *NullableGiftExchangeHistory) Set(val *GiftExchangeHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftExchangeHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftExchangeHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftExchangeHistory(val *GiftExchangeHistory) *NullableGiftExchangeHistory {
	return &NullableGiftExchangeHistory{value: val, isSet: true}
}

func (v NullableGiftExchangeHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftExchangeHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


