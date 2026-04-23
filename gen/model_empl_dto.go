
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplDTO{}

// EmplDTO struct for EmplDTO
type EmplDTO struct {
	Campaign NullableString `json:"campaign,omitempty"`
	Purchased NullableString `json:"purchased,omitempty"`
	Regular NullableString `json:"regular,omitempty"`
	Rewarded NullableString `json:"rewarded,omitempty"`
	Total NullableString `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplDTO EmplDTO

// NewEmplDTO instantiates a new EmplDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplDTO() *EmplDTO {
	this := EmplDTO{}
	return &this
}

// NewEmplDTOWithDefaults instantiates a new EmplDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplDTOWithDefaults() *EmplDTO {
	this := EmplDTO{}
	return &this
}

// GetCampaign returns the Campaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDTO) GetCampaign() string {
	if o == nil || IsNil(o.Campaign.Get()) {
		var ret string
		return ret
	}
	return *o.Campaign.Get()
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDTO) GetCampaignOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaign.Get(), o.Campaign.IsSet()
}

// HasCampaign returns a boolean if a field has been set.
func (o *EmplDTO) HasCampaign() bool {
	if o != nil && o.Campaign.IsSet() {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given NullableString and assigns it to the Campaign field.
func (o *EmplDTO) SetCampaign(v string) {
	o.Campaign.Set(&v)
}
// SetCampaignNil sets the value for Campaign to be an explicit nil
func (o *EmplDTO) SetCampaignNil() {
	o.Campaign.Set(nil)
}

// UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
func (o *EmplDTO) UnsetCampaign() {
	o.Campaign.Unset()
}

// GetPurchased returns the Purchased field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDTO) GetPurchased() string {
	if o == nil || IsNil(o.Purchased.Get()) {
		var ret string
		return ret
	}
	return *o.Purchased.Get()
}

// GetPurchasedOk returns a tuple with the Purchased field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDTO) GetPurchasedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Purchased.Get(), o.Purchased.IsSet()
}

// HasPurchased returns a boolean if a field has been set.
func (o *EmplDTO) HasPurchased() bool {
	if o != nil && o.Purchased.IsSet() {
		return true
	}

	return false
}

// SetPurchased gets a reference to the given NullableString and assigns it to the Purchased field.
func (o *EmplDTO) SetPurchased(v string) {
	o.Purchased.Set(&v)
}
// SetPurchasedNil sets the value for Purchased to be an explicit nil
func (o *EmplDTO) SetPurchasedNil() {
	o.Purchased.Set(nil)
}

// UnsetPurchased ensures that no value is present for Purchased, not even an explicit nil
func (o *EmplDTO) UnsetPurchased() {
	o.Purchased.Unset()
}

// GetRegular returns the Regular field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDTO) GetRegular() string {
	if o == nil || IsNil(o.Regular.Get()) {
		var ret string
		return ret
	}
	return *o.Regular.Get()
}

// GetRegularOk returns a tuple with the Regular field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDTO) GetRegularOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Regular.Get(), o.Regular.IsSet()
}

// HasRegular returns a boolean if a field has been set.
func (o *EmplDTO) HasRegular() bool {
	if o != nil && o.Regular.IsSet() {
		return true
	}

	return false
}

// SetRegular gets a reference to the given NullableString and assigns it to the Regular field.
func (o *EmplDTO) SetRegular(v string) {
	o.Regular.Set(&v)
}
// SetRegularNil sets the value for Regular to be an explicit nil
func (o *EmplDTO) SetRegularNil() {
	o.Regular.Set(nil)
}

// UnsetRegular ensures that no value is present for Regular, not even an explicit nil
func (o *EmplDTO) UnsetRegular() {
	o.Regular.Unset()
}

// GetRewarded returns the Rewarded field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDTO) GetRewarded() string {
	if o == nil || IsNil(o.Rewarded.Get()) {
		var ret string
		return ret
	}
	return *o.Rewarded.Get()
}

// GetRewardedOk returns a tuple with the Rewarded field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDTO) GetRewardedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rewarded.Get(), o.Rewarded.IsSet()
}

// HasRewarded returns a boolean if a field has been set.
func (o *EmplDTO) HasRewarded() bool {
	if o != nil && o.Rewarded.IsSet() {
		return true
	}

	return false
}

// SetRewarded gets a reference to the given NullableString and assigns it to the Rewarded field.
func (o *EmplDTO) SetRewarded(v string) {
	o.Rewarded.Set(&v)
}
// SetRewardedNil sets the value for Rewarded to be an explicit nil
func (o *EmplDTO) SetRewardedNil() {
	o.Rewarded.Set(nil)
}

// UnsetRewarded ensures that no value is present for Rewarded, not even an explicit nil
func (o *EmplDTO) UnsetRewarded() {
	o.Rewarded.Unset()
}

// GetTotal returns the Total field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplDTO) GetTotal() string {
	if o == nil || IsNil(o.Total.Get()) {
		var ret string
		return ret
	}
	return *o.Total.Get()
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplDTO) GetTotalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Total.Get(), o.Total.IsSet()
}

// HasTotal returns a boolean if a field has been set.
func (o *EmplDTO) HasTotal() bool {
	if o != nil && o.Total.IsSet() {
		return true
	}

	return false
}

// SetTotal gets a reference to the given NullableString and assigns it to the Total field.
func (o *EmplDTO) SetTotal(v string) {
	o.Total.Set(&v)
}
// SetTotalNil sets the value for Total to be an explicit nil
func (o *EmplDTO) SetTotalNil() {
	o.Total.Set(nil)
}

// UnsetTotal ensures that no value is present for Total, not even an explicit nil
func (o *EmplDTO) UnsetTotal() {
	o.Total.Unset()
}

func (o EmplDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Campaign.IsSet() {
		toSerialize["campaign"] = o.Campaign.Get()
	}
	if o.Purchased.IsSet() {
		toSerialize["purchased"] = o.Purchased.Get()
	}
	if o.Regular.IsSet() {
		toSerialize["regular"] = o.Regular.Get()
	}
	if o.Rewarded.IsSet() {
		toSerialize["rewarded"] = o.Rewarded.Get()
	}
	if o.Total.IsSet() {
		toSerialize["total"] = o.Total.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplDTO) UnmarshalJSON(data []byte) (err error) {
	varEmplDTO := _EmplDTO{}

	err = json.Unmarshal(data, &varEmplDTO)

	if err != nil {
		return err
	}

	*o = EmplDTO(varEmplDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		delete(additionalProperties, "purchased")
		delete(additionalProperties, "regular")
		delete(additionalProperties, "rewarded")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplDTO struct {
	value *EmplDTO
	isSet bool
}

func (v NullableEmplDTO) Get() *EmplDTO {
	return v.value
}

func (v *NullableEmplDTO) Set(val *EmplDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplDTO(val *EmplDTO) *NullableEmplDTO {
	return &NullableEmplDTO{value: val, isSet: true}
}

func (v NullableEmplDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


