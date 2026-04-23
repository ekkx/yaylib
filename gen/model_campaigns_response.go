
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignsResponse{}

// CampaignsResponse struct for CampaignsResponse
type CampaignsResponse struct {
	Campaigns []CampaignDTO `json:"campaigns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignsResponse CampaignsResponse

// NewCampaignsResponse instantiates a new CampaignsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignsResponse() *CampaignsResponse {
	this := CampaignsResponse{}
	return &this
}

// NewCampaignsResponseWithDefaults instantiates a new CampaignsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignsResponseWithDefaults() *CampaignsResponse {
	this := CampaignsResponse{}
	return &this
}

// GetCampaigns returns the Campaigns field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignsResponse) GetCampaigns() []CampaignDTO {
	if o == nil {
		var ret []CampaignDTO
		return ret
	}
	return o.Campaigns
}

// GetCampaignsOk returns a tuple with the Campaigns field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignsResponse) GetCampaignsOk() ([]CampaignDTO, bool) {
	if o == nil || IsNil(o.Campaigns) {
		return nil, false
	}
	return o.Campaigns, true
}

// HasCampaigns returns a boolean if a field has been set.
func (o *CampaignsResponse) HasCampaigns() bool {
	if o != nil && !IsNil(o.Campaigns) {
		return true
	}

	return false
}

// SetCampaigns gets a reference to the given []CampaignDTO and assigns it to the Campaigns field.
func (o *CampaignsResponse) SetCampaigns(v []CampaignDTO) {
	o.Campaigns = v
}

func (o CampaignsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Campaigns != nil {
		toSerialize["campaigns"] = o.Campaigns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignsResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignsResponse := _CampaignsResponse{}

	err = json.Unmarshal(data, &varCampaignsResponse)

	if err != nil {
		return err
	}

	*o = CampaignsResponse(varCampaignsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "campaigns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignsResponse struct {
	value *CampaignsResponse
	isSet bool
}

func (v NullableCampaignsResponse) Get() *CampaignsResponse {
	return v.value
}

func (v *NullableCampaignsResponse) Set(val *CampaignsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignsResponse(val *CampaignsResponse) *NullableCampaignsResponse {
	return &NullableCampaignsResponse{value: val, isSet: true}
}

func (v NullableCampaignsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


