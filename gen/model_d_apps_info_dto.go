
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DAppsInfoDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DAppsInfoDTO{}

// DAppsInfoDTO struct for DAppsInfoDTO
type DAppsInfoDTO struct {
	Dapps []DAppDTO `json:"dapps,omitempty"`
	EducationalLinks []DAppDTO `json:"educational_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DAppsInfoDTO DAppsInfoDTO

// NewDAppsInfoDTO instantiates a new DAppsInfoDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDAppsInfoDTO() *DAppsInfoDTO {
	this := DAppsInfoDTO{}
	return &this
}

// NewDAppsInfoDTOWithDefaults instantiates a new DAppsInfoDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDAppsInfoDTOWithDefaults() *DAppsInfoDTO {
	this := DAppsInfoDTO{}
	return &this
}

// GetDapps returns the Dapps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppsInfoDTO) GetDapps() []DAppDTO {
	if o == nil {
		var ret []DAppDTO
		return ret
	}
	return o.Dapps
}

// GetDappsOk returns a tuple with the Dapps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppsInfoDTO) GetDappsOk() ([]DAppDTO, bool) {
	if o == nil || IsNil(o.Dapps) {
		return nil, false
	}
	return o.Dapps, true
}

// HasDapps returns a boolean if a field has been set.
func (o *DAppsInfoDTO) HasDapps() bool {
	if o != nil && !IsNil(o.Dapps) {
		return true
	}

	return false
}

// SetDapps gets a reference to the given []DAppDTO and assigns it to the Dapps field.
func (o *DAppsInfoDTO) SetDapps(v []DAppDTO) {
	o.Dapps = v
}

// GetEducationalLinks returns the EducationalLinks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppsInfoDTO) GetEducationalLinks() []DAppDTO {
	if o == nil {
		var ret []DAppDTO
		return ret
	}
	return o.EducationalLinks
}

// GetEducationalLinksOk returns a tuple with the EducationalLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppsInfoDTO) GetEducationalLinksOk() ([]DAppDTO, bool) {
	if o == nil || IsNil(o.EducationalLinks) {
		return nil, false
	}
	return o.EducationalLinks, true
}

// HasEducationalLinks returns a boolean if a field has been set.
func (o *DAppsInfoDTO) HasEducationalLinks() bool {
	if o != nil && !IsNil(o.EducationalLinks) {
		return true
	}

	return false
}

// SetEducationalLinks gets a reference to the given []DAppDTO and assigns it to the EducationalLinks field.
func (o *DAppsInfoDTO) SetEducationalLinks(v []DAppDTO) {
	o.EducationalLinks = v
}

func (o DAppsInfoDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DAppsInfoDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Dapps != nil {
		toSerialize["dapps"] = o.Dapps
	}
	if o.EducationalLinks != nil {
		toSerialize["educational_links"] = o.EducationalLinks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DAppsInfoDTO) UnmarshalJSON(data []byte) (err error) {
	varDAppsInfoDTO := _DAppsInfoDTO{}

	err = json.Unmarshal(data, &varDAppsInfoDTO)

	if err != nil {
		return err
	}

	*o = DAppsInfoDTO(varDAppsInfoDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "dapps")
		delete(additionalProperties, "educational_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDAppsInfoDTO struct {
	value *DAppsInfoDTO
	isSet bool
}

func (v NullableDAppsInfoDTO) Get() *DAppsInfoDTO {
	return v.value
}

func (v *NullableDAppsInfoDTO) Set(val *DAppsInfoDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDAppsInfoDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDAppsInfoDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDAppsInfoDTO(val *DAppsInfoDTO) *NullableDAppsInfoDTO {
	return &NullableDAppsInfoDTO{value: val, isSet: true}
}

func (v NullableDAppsInfoDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDAppsInfoDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


