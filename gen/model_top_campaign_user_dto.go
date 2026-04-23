
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the TopCampaignUserDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TopCampaignUserDTO{}

// TopCampaignUserDTO struct for TopCampaignUserDTO
type TopCampaignUserDTO struct {
	Points NullableInt32 `json:"points,omitempty"`
	Ranking NullableInt32 `json:"ranking,omitempty"`
	User NullableUserDTO `json:"user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TopCampaignUserDTO TopCampaignUserDTO

// NewTopCampaignUserDTO instantiates a new TopCampaignUserDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopCampaignUserDTO() *TopCampaignUserDTO {
	this := TopCampaignUserDTO{}
	return &this
}

// NewTopCampaignUserDTOWithDefaults instantiates a new TopCampaignUserDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopCampaignUserDTOWithDefaults() *TopCampaignUserDTO {
	this := TopCampaignUserDTO{}
	return &this
}

// GetPoints returns the Points field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopCampaignUserDTO) GetPoints() int32 {
	if o == nil || IsNil(o.Points.Get()) {
		var ret int32
		return ret
	}
	return *o.Points.Get()
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopCampaignUserDTO) GetPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Points.Get(), o.Points.IsSet()
}

// HasPoints returns a boolean if a field has been set.
func (o *TopCampaignUserDTO) HasPoints() bool {
	if o != nil && o.Points.IsSet() {
		return true
	}

	return false
}

// SetPoints gets a reference to the given NullableInt32 and assigns it to the Points field.
func (o *TopCampaignUserDTO) SetPoints(v int32) {
	o.Points.Set(&v)
}
// SetPointsNil sets the value for Points to be an explicit nil
func (o *TopCampaignUserDTO) SetPointsNil() {
	o.Points.Set(nil)
}

// UnsetPoints ensures that no value is present for Points, not even an explicit nil
func (o *TopCampaignUserDTO) UnsetPoints() {
	o.Points.Unset()
}

// GetRanking returns the Ranking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopCampaignUserDTO) GetRanking() int32 {
	if o == nil || IsNil(o.Ranking.Get()) {
		var ret int32
		return ret
	}
	return *o.Ranking.Get()
}

// GetRankingOk returns a tuple with the Ranking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopCampaignUserDTO) GetRankingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ranking.Get(), o.Ranking.IsSet()
}

// HasRanking returns a boolean if a field has been set.
func (o *TopCampaignUserDTO) HasRanking() bool {
	if o != nil && o.Ranking.IsSet() {
		return true
	}

	return false
}

// SetRanking gets a reference to the given NullableInt32 and assigns it to the Ranking field.
func (o *TopCampaignUserDTO) SetRanking(v int32) {
	o.Ranking.Set(&v)
}
// SetRankingNil sets the value for Ranking to be an explicit nil
func (o *TopCampaignUserDTO) SetRankingNil() {
	o.Ranking.Set(nil)
}

// UnsetRanking ensures that no value is present for Ranking, not even an explicit nil
func (o *TopCampaignUserDTO) UnsetRanking() {
	o.Ranking.Unset()
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TopCampaignUserDTO) GetUser() UserDTO {
	if o == nil || IsNil(o.User.Get()) {
		var ret UserDTO
		return ret
	}
	return *o.User.Get()
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TopCampaignUserDTO) GetUserOk() (*UserDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.User.Get(), o.User.IsSet()
}

// HasUser returns a boolean if a field has been set.
func (o *TopCampaignUserDTO) HasUser() bool {
	if o != nil && o.User.IsSet() {
		return true
	}

	return false
}

// SetUser gets a reference to the given NullableUserDTO and assigns it to the User field.
func (o *TopCampaignUserDTO) SetUser(v UserDTO) {
	o.User.Set(&v)
}
// SetUserNil sets the value for User to be an explicit nil
func (o *TopCampaignUserDTO) SetUserNil() {
	o.User.Set(nil)
}

// UnsetUser ensures that no value is present for User, not even an explicit nil
func (o *TopCampaignUserDTO) UnsetUser() {
	o.User.Unset()
}

func (o TopCampaignUserDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopCampaignUserDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Points.IsSet() {
		toSerialize["points"] = o.Points.Get()
	}
	if o.Ranking.IsSet() {
		toSerialize["ranking"] = o.Ranking.Get()
	}
	if o.User.IsSet() {
		toSerialize["user"] = o.User.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *TopCampaignUserDTO) UnmarshalJSON(data []byte) (err error) {
	varTopCampaignUserDTO := _TopCampaignUserDTO{}

	err = json.Unmarshal(data, &varTopCampaignUserDTO)

	if err != nil {
		return err
	}

	*o = TopCampaignUserDTO(varTopCampaignUserDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "points")
		delete(additionalProperties, "ranking")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTopCampaignUserDTO struct {
	value *TopCampaignUserDTO
	isSet bool
}

func (v NullableTopCampaignUserDTO) Get() *TopCampaignUserDTO {
	return v.value
}

func (v *NullableTopCampaignUserDTO) Set(val *TopCampaignUserDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableTopCampaignUserDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableTopCampaignUserDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopCampaignUserDTO(val *TopCampaignUserDTO) *NullableTopCampaignUserDTO {
	return &NullableTopCampaignUserDTO{value: val, isSet: true}
}

func (v NullableTopCampaignUserDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopCampaignUserDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


