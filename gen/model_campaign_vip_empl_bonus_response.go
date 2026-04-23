
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignVipEmplBonusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignVipEmplBonusResponse{}

// CampaignVipEmplBonusResponse struct for CampaignVipEmplBonusResponse
type CampaignVipEmplBonusResponse struct {
	AppCampaign NullableCampaignVipEmplBonusDTO `json:"app_campaign,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignVipEmplBonusResponse CampaignVipEmplBonusResponse

// NewCampaignVipEmplBonusResponse instantiates a new CampaignVipEmplBonusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignVipEmplBonusResponse() *CampaignVipEmplBonusResponse {
	this := CampaignVipEmplBonusResponse{}
	return &this
}

// NewCampaignVipEmplBonusResponseWithDefaults instantiates a new CampaignVipEmplBonusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignVipEmplBonusResponseWithDefaults() *CampaignVipEmplBonusResponse {
	this := CampaignVipEmplBonusResponse{}
	return &this
}

// GetAppCampaign returns the AppCampaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignVipEmplBonusResponse) GetAppCampaign() CampaignVipEmplBonusDTO {
	if o == nil || IsNil(o.AppCampaign.Get()) {
		var ret CampaignVipEmplBonusDTO
		return ret
	}
	return *o.AppCampaign.Get()
}

// GetAppCampaignOk returns a tuple with the AppCampaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignVipEmplBonusResponse) GetAppCampaignOk() (*CampaignVipEmplBonusDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.AppCampaign.Get(), o.AppCampaign.IsSet()
}

// HasAppCampaign returns a boolean if a field has been set.
func (o *CampaignVipEmplBonusResponse) HasAppCampaign() bool {
	if o != nil && o.AppCampaign.IsSet() {
		return true
	}

	return false
}

// SetAppCampaign gets a reference to the given NullableCampaignVipEmplBonusDTO and assigns it to the AppCampaign field.
func (o *CampaignVipEmplBonusResponse) SetAppCampaign(v CampaignVipEmplBonusDTO) {
	o.AppCampaign.Set(&v)
}
// SetAppCampaignNil sets the value for AppCampaign to be an explicit nil
func (o *CampaignVipEmplBonusResponse) SetAppCampaignNil() {
	o.AppCampaign.Set(nil)
}

// UnsetAppCampaign ensures that no value is present for AppCampaign, not even an explicit nil
func (o *CampaignVipEmplBonusResponse) UnsetAppCampaign() {
	o.AppCampaign.Unset()
}

func (o CampaignVipEmplBonusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignVipEmplBonusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AppCampaign.IsSet() {
		toSerialize["app_campaign"] = o.AppCampaign.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignVipEmplBonusResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignVipEmplBonusResponse := _CampaignVipEmplBonusResponse{}

	err = json.Unmarshal(data, &varCampaignVipEmplBonusResponse)

	if err != nil {
		return err
	}

	*o = CampaignVipEmplBonusResponse(varCampaignVipEmplBonusResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app_campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignVipEmplBonusResponse struct {
	value *CampaignVipEmplBonusResponse
	isSet bool
}

func (v NullableCampaignVipEmplBonusResponse) Get() *CampaignVipEmplBonusResponse {
	return v.value
}

func (v *NullableCampaignVipEmplBonusResponse) Set(val *CampaignVipEmplBonusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignVipEmplBonusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignVipEmplBonusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignVipEmplBonusResponse(val *CampaignVipEmplBonusResponse) *NullableCampaignVipEmplBonusResponse {
	return &NullableCampaignVipEmplBonusResponse{value: val, isSet: true}
}

func (v NullableCampaignVipEmplBonusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignVipEmplBonusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


