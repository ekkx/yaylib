
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RecentSearchesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecentSearchesResponse{}

// RecentSearchesResponse struct for RecentSearchesResponse
type RecentSearchesResponse struct {
	RecentSearches []RecentSearch `json:"recent_searches,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RecentSearchesResponse RecentSearchesResponse

// NewRecentSearchesResponse instantiates a new RecentSearchesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecentSearchesResponse() *RecentSearchesResponse {
	this := RecentSearchesResponse{}
	return &this
}

// NewRecentSearchesResponseWithDefaults instantiates a new RecentSearchesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecentSearchesResponseWithDefaults() *RecentSearchesResponse {
	this := RecentSearchesResponse{}
	return &this
}

// GetRecentSearches returns the RecentSearches field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RecentSearchesResponse) GetRecentSearches() []RecentSearch {
	if o == nil {
		var ret []RecentSearch
		return ret
	}
	return o.RecentSearches
}

// GetRecentSearchesOk returns a tuple with the RecentSearches field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RecentSearchesResponse) GetRecentSearchesOk() ([]RecentSearch, bool) {
	if o == nil || IsNil(o.RecentSearches) {
		return nil, false
	}
	return o.RecentSearches, true
}

// HasRecentSearches returns a boolean if a field has been set.
func (o *RecentSearchesResponse) HasRecentSearches() bool {
	if o != nil && !IsNil(o.RecentSearches) {
		return true
	}

	return false
}

// SetRecentSearches gets a reference to the given []RecentSearch and assigns it to the RecentSearches field.
func (o *RecentSearchesResponse) SetRecentSearches(v []RecentSearch) {
	o.RecentSearches = v
}

func (o RecentSearchesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RecentSearchesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.RecentSearches != nil {
		toSerialize["recent_searches"] = o.RecentSearches
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RecentSearchesResponse) UnmarshalJSON(data []byte) (err error) {
	varRecentSearchesResponse := _RecentSearchesResponse{}

	err = json.Unmarshal(data, &varRecentSearchesResponse)

	if err != nil {
		return err
	}

	*o = RecentSearchesResponse(varRecentSearchesResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "recent_searches")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRecentSearchesResponse struct {
	value *RecentSearchesResponse
	isSet bool
}

func (v NullableRecentSearchesResponse) Get() *RecentSearchesResponse {
	return v.value
}

func (v *NullableRecentSearchesResponse) Set(val *RecentSearchesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecentSearchesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecentSearchesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecentSearchesResponse(val *RecentSearchesResponse) *NullableRecentSearchesResponse {
	return &NullableRecentSearchesResponse{value: val, isSet: true}
}

func (v NullableRecentSearchesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecentSearchesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


