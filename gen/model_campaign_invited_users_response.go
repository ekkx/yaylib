
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignInvitedUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignInvitedUsersResponse{}

// CampaignInvitedUsersResponse struct for CampaignInvitedUsersResponse
type CampaignInvitedUsersResponse struct {
	InvitedUsers []InvitedUserDTO `json:"invited_users,omitempty"`
	TotalPoints NullableInt32 `json:"total_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignInvitedUsersResponse CampaignInvitedUsersResponse

// NewCampaignInvitedUsersResponse instantiates a new CampaignInvitedUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignInvitedUsersResponse() *CampaignInvitedUsersResponse {
	this := CampaignInvitedUsersResponse{}
	return &this
}

// NewCampaignInvitedUsersResponseWithDefaults instantiates a new CampaignInvitedUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignInvitedUsersResponseWithDefaults() *CampaignInvitedUsersResponse {
	this := CampaignInvitedUsersResponse{}
	return &this
}

// GetInvitedUsers returns the InvitedUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignInvitedUsersResponse) GetInvitedUsers() []InvitedUserDTO {
	if o == nil {
		var ret []InvitedUserDTO
		return ret
	}
	return o.InvitedUsers
}

// GetInvitedUsersOk returns a tuple with the InvitedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignInvitedUsersResponse) GetInvitedUsersOk() ([]InvitedUserDTO, bool) {
	if o == nil || IsNil(o.InvitedUsers) {
		return nil, false
	}
	return o.InvitedUsers, true
}

// HasInvitedUsers returns a boolean if a field has been set.
func (o *CampaignInvitedUsersResponse) HasInvitedUsers() bool {
	if o != nil && !IsNil(o.InvitedUsers) {
		return true
	}

	return false
}

// SetInvitedUsers gets a reference to the given []InvitedUserDTO and assigns it to the InvitedUsers field.
func (o *CampaignInvitedUsersResponse) SetInvitedUsers(v []InvitedUserDTO) {
	o.InvitedUsers = v
}

// GetTotalPoints returns the TotalPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignInvitedUsersResponse) GetTotalPoints() int32 {
	if o == nil || IsNil(o.TotalPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalPoints.Get()
}

// GetTotalPointsOk returns a tuple with the TotalPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignInvitedUsersResponse) GetTotalPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPoints.Get(), o.TotalPoints.IsSet()
}

// HasTotalPoints returns a boolean if a field has been set.
func (o *CampaignInvitedUsersResponse) HasTotalPoints() bool {
	if o != nil && o.TotalPoints.IsSet() {
		return true
	}

	return false
}

// SetTotalPoints gets a reference to the given NullableInt32 and assigns it to the TotalPoints field.
func (o *CampaignInvitedUsersResponse) SetTotalPoints(v int32) {
	o.TotalPoints.Set(&v)
}
// SetTotalPointsNil sets the value for TotalPoints to be an explicit nil
func (o *CampaignInvitedUsersResponse) SetTotalPointsNil() {
	o.TotalPoints.Set(nil)
}

// UnsetTotalPoints ensures that no value is present for TotalPoints, not even an explicit nil
func (o *CampaignInvitedUsersResponse) UnsetTotalPoints() {
	o.TotalPoints.Unset()
}

func (o CampaignInvitedUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignInvitedUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InvitedUsers != nil {
		toSerialize["invited_users"] = o.InvitedUsers
	}
	if o.TotalPoints.IsSet() {
		toSerialize["total_points"] = o.TotalPoints.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignInvitedUsersResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignInvitedUsersResponse := _CampaignInvitedUsersResponse{}

	err = json.Unmarshal(data, &varCampaignInvitedUsersResponse)

	if err != nil {
		return err
	}

	*o = CampaignInvitedUsersResponse(varCampaignInvitedUsersResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "invited_users")
		delete(additionalProperties, "total_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignInvitedUsersResponse struct {
	value *CampaignInvitedUsersResponse
	isSet bool
}

func (v NullableCampaignInvitedUsersResponse) Get() *CampaignInvitedUsersResponse {
	return v.value
}

func (v *NullableCampaignInvitedUsersResponse) Set(val *CampaignInvitedUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignInvitedUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignInvitedUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignInvitedUsersResponse(val *CampaignInvitedUsersResponse) *NullableCampaignInvitedUsersResponse {
	return &NullableCampaignInvitedUsersResponse{value: val, isSet: true}
}

func (v NullableCampaignInvitedUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignInvitedUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


