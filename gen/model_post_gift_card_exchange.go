
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PostGiftCardExchange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostGiftCardExchange{}

// PostGiftCardExchange struct for PostGiftCardExchange
type PostGiftCardExchange struct {
	Email NullableString `json:"email,omitempty"`
	GiftCard NullableString `json:"gift_card,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostGiftCardExchange PostGiftCardExchange

// NewPostGiftCardExchange instantiates a new PostGiftCardExchange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostGiftCardExchange() *PostGiftCardExchange {
	this := PostGiftCardExchange{}
	return &this
}

// NewPostGiftCardExchangeWithDefaults instantiates a new PostGiftCardExchange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostGiftCardExchangeWithDefaults() *PostGiftCardExchange {
	this := PostGiftCardExchange{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostGiftCardExchange) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostGiftCardExchange) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PostGiftCardExchange) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *PostGiftCardExchange) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PostGiftCardExchange) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PostGiftCardExchange) UnsetEmail() {
	o.Email.Unset()
}

// GetGiftCard returns the GiftCard field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostGiftCardExchange) GetGiftCard() string {
	if o == nil || IsNil(o.GiftCard.Get()) {
		var ret string
		return ret
	}
	return *o.GiftCard.Get()
}

// GetGiftCardOk returns a tuple with the GiftCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostGiftCardExchange) GetGiftCardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GiftCard.Get(), o.GiftCard.IsSet()
}

// HasGiftCard returns a boolean if a field has been set.
func (o *PostGiftCardExchange) HasGiftCard() bool {
	if o != nil && o.GiftCard.IsSet() {
		return true
	}

	return false
}

// SetGiftCard gets a reference to the given NullableString and assigns it to the GiftCard field.
func (o *PostGiftCardExchange) SetGiftCard(v string) {
	o.GiftCard.Set(&v)
}
// SetGiftCardNil sets the value for GiftCard to be an explicit nil
func (o *PostGiftCardExchange) SetGiftCardNil() {
	o.GiftCard.Set(nil)
}

// UnsetGiftCard ensures that no value is present for GiftCard, not even an explicit nil
func (o *PostGiftCardExchange) UnsetGiftCard() {
	o.GiftCard.Unset()
}

func (o PostGiftCardExchange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostGiftCardExchange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.GiftCard.IsSet() {
		toSerialize["gift_card"] = o.GiftCard.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostGiftCardExchange) UnmarshalJSON(data []byte) (err error) {
	varPostGiftCardExchange := _PostGiftCardExchange{}

	err = json.Unmarshal(data, &varPostGiftCardExchange)

	if err != nil {
		return err
	}

	*o = PostGiftCardExchange(varPostGiftCardExchange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "email")
		delete(additionalProperties, "gift_card")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePostGiftCardExchange struct {
	value *PostGiftCardExchange
	isSet bool
}

func (v NullablePostGiftCardExchange) Get() *PostGiftCardExchange {
	return v.value
}

func (v *NullablePostGiftCardExchange) Set(val *PostGiftCardExchange) {
	v.value = val
	v.isSet = true
}

func (v NullablePostGiftCardExchange) IsSet() bool {
	return v.isSet
}

func (v *NullablePostGiftCardExchange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostGiftCardExchange(val *PostGiftCardExchange) *NullablePostGiftCardExchange {
	return &NullablePostGiftCardExchange{value: val, isSet: true}
}

func (v NullablePostGiftCardExchange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostGiftCardExchange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


