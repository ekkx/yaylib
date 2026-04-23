
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignResponse{}

// CampaignResponse struct for CampaignResponse
type CampaignResponse struct {
	Campaign NullableCampaignDTO `json:"campaign,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignResponse CampaignResponse

// NewCampaignResponse instantiates a new CampaignResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignResponse() *CampaignResponse {
	this := CampaignResponse{}
	return &this
}

// NewCampaignResponseWithDefaults instantiates a new CampaignResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignResponseWithDefaults() *CampaignResponse {
	this := CampaignResponse{}
	return &this
}

// GetCampaign returns the Campaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignResponse) GetCampaign() CampaignDTO {
	if o == nil || IsNil(o.Campaign.Get()) {
		var ret CampaignDTO
		return ret
	}
	return *o.Campaign.Get()
}

// GetCampaignOk returns a tuple with the Campaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignResponse) GetCampaignOk() (*CampaignDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.Campaign.Get(), o.Campaign.IsSet()
}

// HasCampaign returns a boolean if a field has been set.
func (o *CampaignResponse) HasCampaign() bool {
	if o != nil && o.Campaign.IsSet() {
		return true
	}

	return false
}

// SetCampaign gets a reference to the given NullableCampaignDTO and assigns it to the Campaign field.
func (o *CampaignResponse) SetCampaign(v CampaignDTO) {
	o.Campaign.Set(&v)
}
// SetCampaignNil sets the value for Campaign to be an explicit nil
func (o *CampaignResponse) SetCampaignNil() {
	o.Campaign.Set(nil)
}

// UnsetCampaign ensures that no value is present for Campaign, not even an explicit nil
func (o *CampaignResponse) UnsetCampaign() {
	o.Campaign.Unset()
}

func (o CampaignResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Campaign.IsSet() {
		toSerialize["campaign"] = o.Campaign.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignResponse := _CampaignResponse{}

	err = json.Unmarshal(data, &varCampaignResponse)

	if err != nil {
		return err
	}

	*o = CampaignResponse(varCampaignResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignResponse struct {
	value *CampaignResponse
	isSet bool
}

func (v NullableCampaignResponse) Get() *CampaignResponse {
	return v.value
}

func (v *NullableCampaignResponse) Set(val *CampaignResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignResponse(val *CampaignResponse) *NullableCampaignResponse {
	return &NullableCampaignResponse{value: val, isSet: true}
}

func (v NullableCampaignResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


