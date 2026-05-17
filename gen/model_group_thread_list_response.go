
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the GroupThreadListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupThreadListResponse{}

// GroupThreadListResponse struct for GroupThreadListResponse
type GroupThreadListResponse struct {
	NextPageValue NullableCursor `json:"next_page_value,omitempty"`
	Threads []ThreadInfo `json:"threads,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GroupThreadListResponse GroupThreadListResponse

// NewGroupThreadListResponse instantiates a new GroupThreadListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupThreadListResponse() *GroupThreadListResponse {
	this := GroupThreadListResponse{}
	return &this
}

// NewGroupThreadListResponseWithDefaults instantiates a new GroupThreadListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupThreadListResponseWithDefaults() *GroupThreadListResponse {
	this := GroupThreadListResponse{}
	return &this
}

// GetNextPageValue returns the NextPageValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupThreadListResponse) GetNextPageValue() string {
	if o == nil || IsNil(o.NextPageValue.Get()) {
		var ret string
		return ret
	}
	return *o.NextPageValue.Get()
}

// GetNextPageValueOk returns a tuple with the NextPageValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupThreadListResponse) GetNextPageValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPageValue.Get(), o.NextPageValue.IsSet()
}

// HasNextPageValue returns a boolean if a field has been set.
func (o *GroupThreadListResponse) HasNextPageValue() bool {
	if o != nil && o.NextPageValue.IsSet() {
		return true
	}

	return false
}

// SetNextPageValue gets a reference to the given NullableString and assigns it to the NextPageValue field.
func (o *GroupThreadListResponse) SetNextPageValue(v string) {
	o.NextPageValue.Set(&v)
}
// SetNextPageValueNil sets the value for NextPageValue to be an explicit nil
func (o *GroupThreadListResponse) SetNextPageValueNil() {
	o.NextPageValue.Set(nil)
}

// UnsetNextPageValue ensures that no value is present for NextPageValue, not even an explicit nil
func (o *GroupThreadListResponse) UnsetNextPageValue() {
	o.NextPageValue.Unset()
}

// GetThreads returns the Threads field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GroupThreadListResponse) GetThreads() []ThreadInfo {
	if o == nil {
		var ret []ThreadInfo
		return ret
	}
	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GroupThreadListResponse) GetThreadsOk() ([]ThreadInfo, bool) {
	if o == nil || IsNil(o.Threads) {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *GroupThreadListResponse) HasThreads() bool {
	if o != nil && !IsNil(o.Threads) {
		return true
	}

	return false
}

// SetThreads gets a reference to the given []ThreadInfo and assigns it to the Threads field.
func (o *GroupThreadListResponse) SetThreads(v []ThreadInfo) {
	o.Threads = v
}

func (o GroupThreadListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupThreadListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NextPageValue.IsSet() {
		toSerialize["next_page_value"] = o.NextPageValue.Get()
	}
	if o.Threads != nil {
		toSerialize["threads"] = o.Threads
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GroupThreadListResponse) UnmarshalJSON(data []byte) (err error) {
	varGroupThreadListResponse := _GroupThreadListResponse{}

	err = json.Unmarshal(data, &varGroupThreadListResponse)

	if err != nil {
		return err
	}

	*o = GroupThreadListResponse(varGroupThreadListResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "next_page_value")
		delete(additionalProperties, "threads")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGroupThreadListResponse struct {
	value *GroupThreadListResponse
	isSet bool
}

func (v NullableGroupThreadListResponse) Get() *GroupThreadListResponse {
	return v.value
}

func (v *NullableGroupThreadListResponse) Set(val *GroupThreadListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupThreadListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupThreadListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupThreadListResponse(val *GroupThreadListResponse) *NullableGroupThreadListResponse {
	return &NullableGroupThreadListResponse{value: val, isSet: true}
}

func (v NullableGroupThreadListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupThreadListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


