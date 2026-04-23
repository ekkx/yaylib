
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PromotionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotionsResponse{}

// PromotionsResponse struct for PromotionsResponse
type PromotionsResponse struct {
	Promotions []Promotion `json:"promotions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PromotionsResponse PromotionsResponse

// NewPromotionsResponse instantiates a new PromotionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotionsResponse() *PromotionsResponse {
	this := PromotionsResponse{}
	return &this
}

// NewPromotionsResponseWithDefaults instantiates a new PromotionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionsResponseWithDefaults() *PromotionsResponse {
	this := PromotionsResponse{}
	return &this
}

// GetPromotions returns the Promotions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PromotionsResponse) GetPromotions() []Promotion {
	if o == nil {
		var ret []Promotion
		return ret
	}
	return o.Promotions
}

// GetPromotionsOk returns a tuple with the Promotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PromotionsResponse) GetPromotionsOk() ([]Promotion, bool) {
	if o == nil || IsNil(o.Promotions) {
		return nil, false
	}
	return o.Promotions, true
}

// HasPromotions returns a boolean if a field has been set.
func (o *PromotionsResponse) HasPromotions() bool {
	if o != nil && !IsNil(o.Promotions) {
		return true
	}

	return false
}

// SetPromotions gets a reference to the given []Promotion and assigns it to the Promotions field.
func (o *PromotionsResponse) SetPromotions(v []Promotion) {
	o.Promotions = v
}

func (o PromotionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Promotions != nil {
		toSerialize["promotions"] = o.Promotions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PromotionsResponse) UnmarshalJSON(data []byte) (err error) {
	varPromotionsResponse := _PromotionsResponse{}

	err = json.Unmarshal(data, &varPromotionsResponse)

	if err != nil {
		return err
	}

	*o = PromotionsResponse(varPromotionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "promotions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePromotionsResponse struct {
	value *PromotionsResponse
	isSet bool
}

func (v NullablePromotionsResponse) Get() *PromotionsResponse {
	return v.value
}

func (v *NullablePromotionsResponse) Set(val *PromotionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotionsResponse(val *PromotionsResponse) *NullablePromotionsResponse {
	return &NullablePromotionsResponse{value: val, isSet: true}
}

func (v NullablePromotionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


