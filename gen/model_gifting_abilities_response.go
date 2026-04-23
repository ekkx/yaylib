
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GiftingAbilitiesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GiftingAbilitiesResponse{}

// GiftingAbilitiesResponse struct for GiftingAbilitiesResponse
type GiftingAbilitiesResponse struct {
	GiftingAbilities []RealmGiftingAbility `json:"gifting_abilities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GiftingAbilitiesResponse GiftingAbilitiesResponse

// NewGiftingAbilitiesResponse instantiates a new GiftingAbilitiesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGiftingAbilitiesResponse() *GiftingAbilitiesResponse {
	this := GiftingAbilitiesResponse{}
	return &this
}

// NewGiftingAbilitiesResponseWithDefaults instantiates a new GiftingAbilitiesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGiftingAbilitiesResponseWithDefaults() *GiftingAbilitiesResponse {
	this := GiftingAbilitiesResponse{}
	return &this
}

// GetGiftingAbilities returns the GiftingAbilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GiftingAbilitiesResponse) GetGiftingAbilities() []RealmGiftingAbility {
	if o == nil {
		var ret []RealmGiftingAbility
		return ret
	}
	return o.GiftingAbilities
}

// GetGiftingAbilitiesOk returns a tuple with the GiftingAbilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GiftingAbilitiesResponse) GetGiftingAbilitiesOk() ([]RealmGiftingAbility, bool) {
	if o == nil || IsNil(o.GiftingAbilities) {
		return nil, false
	}
	return o.GiftingAbilities, true
}

// HasGiftingAbilities returns a boolean if a field has been set.
func (o *GiftingAbilitiesResponse) HasGiftingAbilities() bool {
	if o != nil && !IsNil(o.GiftingAbilities) {
		return true
	}

	return false
}

// SetGiftingAbilities gets a reference to the given []RealmGiftingAbility and assigns it to the GiftingAbilities field.
func (o *GiftingAbilitiesResponse) SetGiftingAbilities(v []RealmGiftingAbility) {
	o.GiftingAbilities = v
}

func (o GiftingAbilitiesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GiftingAbilitiesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GiftingAbilities != nil {
		toSerialize["gifting_abilities"] = o.GiftingAbilities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GiftingAbilitiesResponse) UnmarshalJSON(data []byte) (err error) {
	varGiftingAbilitiesResponse := _GiftingAbilitiesResponse{}

	err = json.Unmarshal(data, &varGiftingAbilitiesResponse)

	if err != nil {
		return err
	}

	*o = GiftingAbilitiesResponse(varGiftingAbilitiesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "gifting_abilities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGiftingAbilitiesResponse struct {
	value *GiftingAbilitiesResponse
	isSet bool
}

func (v NullableGiftingAbilitiesResponse) Get() *GiftingAbilitiesResponse {
	return v.value
}

func (v *NullableGiftingAbilitiesResponse) Set(val *GiftingAbilitiesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGiftingAbilitiesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGiftingAbilitiesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGiftingAbilitiesResponse(val *GiftingAbilitiesResponse) *NullableGiftingAbilitiesResponse {
	return &NullableGiftingAbilitiesResponse{value: val, isSet: true}
}

func (v NullableGiftingAbilitiesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGiftingAbilitiesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


