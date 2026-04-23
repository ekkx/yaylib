
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RefreshCounterRequestsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshCounterRequestsResponse{}

// RefreshCounterRequestsResponse struct for RefreshCounterRequestsResponse
type RefreshCounterRequestsResponse struct {
	ResetCounterRequests []RefreshCounterRequest `json:"reset_counter_requests,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RefreshCounterRequestsResponse RefreshCounterRequestsResponse

// NewRefreshCounterRequestsResponse instantiates a new RefreshCounterRequestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshCounterRequestsResponse() *RefreshCounterRequestsResponse {
	this := RefreshCounterRequestsResponse{}
	return &this
}

// NewRefreshCounterRequestsResponseWithDefaults instantiates a new RefreshCounterRequestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshCounterRequestsResponseWithDefaults() *RefreshCounterRequestsResponse {
	this := RefreshCounterRequestsResponse{}
	return &this
}

// GetResetCounterRequests returns the ResetCounterRequests field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RefreshCounterRequestsResponse) GetResetCounterRequests() []RefreshCounterRequest {
	if o == nil {
		var ret []RefreshCounterRequest
		return ret
	}
	return o.ResetCounterRequests
}

// GetResetCounterRequestsOk returns a tuple with the ResetCounterRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RefreshCounterRequestsResponse) GetResetCounterRequestsOk() ([]RefreshCounterRequest, bool) {
	if o == nil || IsNil(o.ResetCounterRequests) {
		return nil, false
	}
	return o.ResetCounterRequests, true
}

// HasResetCounterRequests returns a boolean if a field has been set.
func (o *RefreshCounterRequestsResponse) HasResetCounterRequests() bool {
	if o != nil && !IsNil(o.ResetCounterRequests) {
		return true
	}

	return false
}

// SetResetCounterRequests gets a reference to the given []RefreshCounterRequest and assigns it to the ResetCounterRequests field.
func (o *RefreshCounterRequestsResponse) SetResetCounterRequests(v []RefreshCounterRequest) {
	o.ResetCounterRequests = v
}

func (o RefreshCounterRequestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshCounterRequestsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ResetCounterRequests != nil {
		toSerialize["reset_counter_requests"] = o.ResetCounterRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RefreshCounterRequestsResponse) UnmarshalJSON(data []byte) (err error) {
	varRefreshCounterRequestsResponse := _RefreshCounterRequestsResponse{}

	err = json.Unmarshal(data, &varRefreshCounterRequestsResponse)

	if err != nil {
		return err
	}

	*o = RefreshCounterRequestsResponse(varRefreshCounterRequestsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "reset_counter_requests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRefreshCounterRequestsResponse struct {
	value *RefreshCounterRequestsResponse
	isSet bool
}

func (v NullableRefreshCounterRequestsResponse) Get() *RefreshCounterRequestsResponse {
	return v.value
}

func (v *NullableRefreshCounterRequestsResponse) Set(val *RefreshCounterRequestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshCounterRequestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshCounterRequestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshCounterRequestsResponse(val *RefreshCounterRequestsResponse) *NullableRefreshCounterRequestsResponse {
	return &NullableRefreshCounterRequestsResponse{value: val, isSet: true}
}

func (v NullableRefreshCounterRequestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshCounterRequestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


