
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LikePostsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LikePostsResponse{}

// LikePostsResponse struct for LikePostsResponse
type LikePostsResponse struct {
	LikeIds []int64 `json:"like_ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LikePostsResponse LikePostsResponse

// NewLikePostsResponse instantiates a new LikePostsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLikePostsResponse() *LikePostsResponse {
	this := LikePostsResponse{}
	return &this
}

// NewLikePostsResponseWithDefaults instantiates a new LikePostsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLikePostsResponseWithDefaults() *LikePostsResponse {
	this := LikePostsResponse{}
	return &this
}

// GetLikeIds returns the LikeIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LikePostsResponse) GetLikeIds() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}
	return o.LikeIds
}

// GetLikeIdsOk returns a tuple with the LikeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LikePostsResponse) GetLikeIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.LikeIds) {
		return nil, false
	}
	return o.LikeIds, true
}

// HasLikeIds returns a boolean if a field has been set.
func (o *LikePostsResponse) HasLikeIds() bool {
	if o != nil && !IsNil(o.LikeIds) {
		return true
	}

	return false
}

// SetLikeIds gets a reference to the given []int64 and assigns it to the LikeIds field.
func (o *LikePostsResponse) SetLikeIds(v []int64) {
	o.LikeIds = v
}

func (o LikePostsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LikePostsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LikeIds != nil {
		toSerialize["like_ids"] = o.LikeIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LikePostsResponse) UnmarshalJSON(data []byte) (err error) {
	varLikePostsResponse := _LikePostsResponse{}

	err = json.Unmarshal(data, &varLikePostsResponse)

	if err != nil {
		return err
	}

	*o = LikePostsResponse(varLikePostsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "like_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLikePostsResponse struct {
	value *LikePostsResponse
	isSet bool
}

func (v NullableLikePostsResponse) Get() *LikePostsResponse {
	return v.value
}

func (v *NullableLikePostsResponse) Set(val *LikePostsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLikePostsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLikePostsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLikePostsResponse(val *LikePostsResponse) *NullableLikePostsResponse {
	return &NullableLikePostsResponse{value: val, isSet: true}
}

func (v NullableLikePostsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLikePostsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


