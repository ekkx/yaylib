
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignPointHistoryResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignPointHistoryResponse{}

// CampaignPointHistoryResponse struct for CampaignPointHistoryResponse
type CampaignPointHistoryResponse struct {
	PointHistories []CampaignPointHistoryDTO `json:"point_histories,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignPointHistoryResponse CampaignPointHistoryResponse

// NewCampaignPointHistoryResponse instantiates a new CampaignPointHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignPointHistoryResponse() *CampaignPointHistoryResponse {
	this := CampaignPointHistoryResponse{}
	return &this
}

// NewCampaignPointHistoryResponseWithDefaults instantiates a new CampaignPointHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignPointHistoryResponseWithDefaults() *CampaignPointHistoryResponse {
	this := CampaignPointHistoryResponse{}
	return &this
}

// GetPointHistories returns the PointHistories field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignPointHistoryResponse) GetPointHistories() []CampaignPointHistoryDTO {
	if o == nil {
		var ret []CampaignPointHistoryDTO
		return ret
	}
	return o.PointHistories
}

// GetPointHistoriesOk returns a tuple with the PointHistories field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignPointHistoryResponse) GetPointHistoriesOk() ([]CampaignPointHistoryDTO, bool) {
	if o == nil || IsNil(o.PointHistories) {
		return nil, false
	}
	return o.PointHistories, true
}

// HasPointHistories returns a boolean if a field has been set.
func (o *CampaignPointHistoryResponse) HasPointHistories() bool {
	if o != nil && !IsNil(o.PointHistories) {
		return true
	}

	return false
}

// SetPointHistories gets a reference to the given []CampaignPointHistoryDTO and assigns it to the PointHistories field.
func (o *CampaignPointHistoryResponse) SetPointHistories(v []CampaignPointHistoryDTO) {
	o.PointHistories = v
}

func (o CampaignPointHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignPointHistoryResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PointHistories != nil {
		toSerialize["point_histories"] = o.PointHistories
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignPointHistoryResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignPointHistoryResponse := _CampaignPointHistoryResponse{}

	err = json.Unmarshal(data, &varCampaignPointHistoryResponse)

	if err != nil {
		return err
	}

	*o = CampaignPointHistoryResponse(varCampaignPointHistoryResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "point_histories")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignPointHistoryResponse struct {
	value *CampaignPointHistoryResponse
	isSet bool
}

func (v NullableCampaignPointHistoryResponse) Get() *CampaignPointHistoryResponse {
	return v.value
}

func (v *NullableCampaignPointHistoryResponse) Set(val *CampaignPointHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignPointHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignPointHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignPointHistoryResponse(val *CampaignPointHistoryResponse) *NullableCampaignPointHistoryResponse {
	return &NullableCampaignPointHistoryResponse{value: val, isSet: true}
}

func (v NullableCampaignPointHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignPointHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


