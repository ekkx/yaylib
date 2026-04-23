
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignRankingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignRankingResponse{}

// CampaignRankingResponse struct for CampaignRankingResponse
type CampaignRankingResponse struct {
	MyRanking NullableInt32 `json:"my_ranking,omitempty"`
	NextCursor NullableString `json:"next_cursor,omitempty"`
	TopCampaignUsers []TopCampaignUserDTO `json:"top_campaign_users,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignRankingResponse CampaignRankingResponse

// NewCampaignRankingResponse instantiates a new CampaignRankingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignRankingResponse() *CampaignRankingResponse {
	this := CampaignRankingResponse{}
	return &this
}

// NewCampaignRankingResponseWithDefaults instantiates a new CampaignRankingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignRankingResponseWithDefaults() *CampaignRankingResponse {
	this := CampaignRankingResponse{}
	return &this
}

// GetMyRanking returns the MyRanking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignRankingResponse) GetMyRanking() int32 {
	if o == nil || IsNil(o.MyRanking.Get()) {
		var ret int32
		return ret
	}
	return *o.MyRanking.Get()
}

// GetMyRankingOk returns a tuple with the MyRanking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignRankingResponse) GetMyRankingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MyRanking.Get(), o.MyRanking.IsSet()
}

// HasMyRanking returns a boolean if a field has been set.
func (o *CampaignRankingResponse) HasMyRanking() bool {
	if o != nil && o.MyRanking.IsSet() {
		return true
	}

	return false
}

// SetMyRanking gets a reference to the given NullableInt32 and assigns it to the MyRanking field.
func (o *CampaignRankingResponse) SetMyRanking(v int32) {
	o.MyRanking.Set(&v)
}
// SetMyRankingNil sets the value for MyRanking to be an explicit nil
func (o *CampaignRankingResponse) SetMyRankingNil() {
	o.MyRanking.Set(nil)
}

// UnsetMyRanking ensures that no value is present for MyRanking, not even an explicit nil
func (o *CampaignRankingResponse) UnsetMyRanking() {
	o.MyRanking.Unset()
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignRankingResponse) GetNextCursor() string {
	if o == nil || IsNil(o.NextCursor.Get()) {
		var ret string
		return ret
	}
	return *o.NextCursor.Get()
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignRankingResponse) GetNextCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextCursor.Get(), o.NextCursor.IsSet()
}

// HasNextCursor returns a boolean if a field has been set.
func (o *CampaignRankingResponse) HasNextCursor() bool {
	if o != nil && o.NextCursor.IsSet() {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given NullableString and assigns it to the NextCursor field.
func (o *CampaignRankingResponse) SetNextCursor(v string) {
	o.NextCursor.Set(&v)
}
// SetNextCursorNil sets the value for NextCursor to be an explicit nil
func (o *CampaignRankingResponse) SetNextCursorNil() {
	o.NextCursor.Set(nil)
}

// UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
func (o *CampaignRankingResponse) UnsetNextCursor() {
	o.NextCursor.Unset()
}

// GetTopCampaignUsers returns the TopCampaignUsers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignRankingResponse) GetTopCampaignUsers() []TopCampaignUserDTO {
	if o == nil {
		var ret []TopCampaignUserDTO
		return ret
	}
	return o.TopCampaignUsers
}

// GetTopCampaignUsersOk returns a tuple with the TopCampaignUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignRankingResponse) GetTopCampaignUsersOk() ([]TopCampaignUserDTO, bool) {
	if o == nil || IsNil(o.TopCampaignUsers) {
		return nil, false
	}
	return o.TopCampaignUsers, true
}

// HasTopCampaignUsers returns a boolean if a field has been set.
func (o *CampaignRankingResponse) HasTopCampaignUsers() bool {
	if o != nil && !IsNil(o.TopCampaignUsers) {
		return true
	}

	return false
}

// SetTopCampaignUsers gets a reference to the given []TopCampaignUserDTO and assigns it to the TopCampaignUsers field.
func (o *CampaignRankingResponse) SetTopCampaignUsers(v []TopCampaignUserDTO) {
	o.TopCampaignUsers = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignRankingResponse) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignRankingResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *CampaignRankingResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *CampaignRankingResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *CampaignRankingResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *CampaignRankingResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

func (o CampaignRankingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignRankingResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MyRanking.IsSet() {
		toSerialize["my_ranking"] = o.MyRanking.Get()
	}
	if o.NextCursor.IsSet() {
		toSerialize["next_cursor"] = o.NextCursor.Get()
	}
	if o.TopCampaignUsers != nil {
		toSerialize["top_campaign_users"] = o.TopCampaignUsers
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignRankingResponse) UnmarshalJSON(data []byte) (err error) {
	varCampaignRankingResponse := _CampaignRankingResponse{}

	err = json.Unmarshal(data, &varCampaignRankingResponse)

	if err != nil {
		return err
	}

	*o = CampaignRankingResponse(varCampaignRankingResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "my_ranking")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "top_campaign_users")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignRankingResponse struct {
	value *CampaignRankingResponse
	isSet bool
}

func (v NullableCampaignRankingResponse) Get() *CampaignRankingResponse {
	return v.value
}

func (v *NullableCampaignRankingResponse) Set(val *CampaignRankingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignRankingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignRankingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignRankingResponse(val *CampaignRankingResponse) *NullableCampaignRankingResponse {
	return &NullableCampaignRankingResponse{value: val, isSet: true}
}

func (v NullableCampaignRankingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignRankingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


