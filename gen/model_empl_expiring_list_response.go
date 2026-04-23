
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the EmplExpiringListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmplExpiringListResponse{}

// EmplExpiringListResponse struct for EmplExpiringListResponse
type EmplExpiringListResponse struct {
	RegularExpiring []EmplExpiringResponse `json:"regular_expiring,omitempty"`
	RewardedExpiring []EmplExpiringResponse `json:"rewarded_expiring,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EmplExpiringListResponse EmplExpiringListResponse

// NewEmplExpiringListResponse instantiates a new EmplExpiringListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmplExpiringListResponse() *EmplExpiringListResponse {
	this := EmplExpiringListResponse{}
	return &this
}

// NewEmplExpiringListResponseWithDefaults instantiates a new EmplExpiringListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmplExpiringListResponseWithDefaults() *EmplExpiringListResponse {
	this := EmplExpiringListResponse{}
	return &this
}

// GetRegularExpiring returns the RegularExpiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplExpiringListResponse) GetRegularExpiring() []EmplExpiringResponse {
	if o == nil {
		var ret []EmplExpiringResponse
		return ret
	}
	return o.RegularExpiring
}

// GetRegularExpiringOk returns a tuple with the RegularExpiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplExpiringListResponse) GetRegularExpiringOk() ([]EmplExpiringResponse, bool) {
	if o == nil || IsNil(o.RegularExpiring) {
		return nil, false
	}
	return o.RegularExpiring, true
}

// HasRegularExpiring returns a boolean if a field has been set.
func (o *EmplExpiringListResponse) HasRegularExpiring() bool {
	if o != nil && !IsNil(o.RegularExpiring) {
		return true
	}

	return false
}

// SetRegularExpiring gets a reference to the given []EmplExpiringResponse and assigns it to the RegularExpiring field.
func (o *EmplExpiringListResponse) SetRegularExpiring(v []EmplExpiringResponse) {
	o.RegularExpiring = v
}

// GetRewardedExpiring returns the RewardedExpiring field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EmplExpiringListResponse) GetRewardedExpiring() []EmplExpiringResponse {
	if o == nil {
		var ret []EmplExpiringResponse
		return ret
	}
	return o.RewardedExpiring
}

// GetRewardedExpiringOk returns a tuple with the RewardedExpiring field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EmplExpiringListResponse) GetRewardedExpiringOk() ([]EmplExpiringResponse, bool) {
	if o == nil || IsNil(o.RewardedExpiring) {
		return nil, false
	}
	return o.RewardedExpiring, true
}

// HasRewardedExpiring returns a boolean if a field has been set.
func (o *EmplExpiringListResponse) HasRewardedExpiring() bool {
	if o != nil && !IsNil(o.RewardedExpiring) {
		return true
	}

	return false
}

// SetRewardedExpiring gets a reference to the given []EmplExpiringResponse and assigns it to the RewardedExpiring field.
func (o *EmplExpiringListResponse) SetRewardedExpiring(v []EmplExpiringResponse) {
	o.RewardedExpiring = v
}

func (o EmplExpiringListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmplExpiringListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RegularExpiring != nil {
		toSerialize["regular_expiring"] = o.RegularExpiring
	}
	if o.RewardedExpiring != nil {
		toSerialize["rewarded_expiring"] = o.RewardedExpiring
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EmplExpiringListResponse) UnmarshalJSON(data []byte) (err error) {
	varEmplExpiringListResponse := _EmplExpiringListResponse{}

	err = json.Unmarshal(data, &varEmplExpiringListResponse)

	if err != nil {
		return err
	}

	*o = EmplExpiringListResponse(varEmplExpiringListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "regular_expiring")
		delete(additionalProperties, "rewarded_expiring")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEmplExpiringListResponse struct {
	value *EmplExpiringListResponse
	isSet bool
}

func (v NullableEmplExpiringListResponse) Get() *EmplExpiringListResponse {
	return v.value
}

func (v *NullableEmplExpiringListResponse) Set(val *EmplExpiringListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmplExpiringListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmplExpiringListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmplExpiringListResponse(val *EmplExpiringListResponse) *NullableEmplExpiringListResponse {
	return &NullableEmplExpiringListResponse{value: val, isSet: true}
}

func (v NullableEmplExpiringListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmplExpiringListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


