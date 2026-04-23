
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the OnChainTransactionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnChainTransactionsResponse{}

// OnChainTransactionsResponse struct for OnChainTransactionsResponse
type OnChainTransactionsResponse struct {
	Cursor NullableString `json:"cursor,omitempty"`
	Result []OnChainTransactionDTO `json:"result,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OnChainTransactionsResponse OnChainTransactionsResponse

// NewOnChainTransactionsResponse instantiates a new OnChainTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnChainTransactionsResponse() *OnChainTransactionsResponse {
	this := OnChainTransactionsResponse{}
	return &this
}

// NewOnChainTransactionsResponseWithDefaults instantiates a new OnChainTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnChainTransactionsResponseWithDefaults() *OnChainTransactionsResponse {
	this := OnChainTransactionsResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionsResponse) GetCursor() string {
	if o == nil || IsNil(o.Cursor.Get()) {
		var ret string
		return ret
	}
	return *o.Cursor.Get()
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionsResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cursor.Get(), o.Cursor.IsSet()
}

// HasCursor returns a boolean if a field has been set.
func (o *OnChainTransactionsResponse) HasCursor() bool {
	if o != nil && o.Cursor.IsSet() {
		return true
	}

	return false
}

// SetCursor gets a reference to the given NullableString and assigns it to the Cursor field.
func (o *OnChainTransactionsResponse) SetCursor(v string) {
	o.Cursor.Set(&v)
}
// SetCursorNil sets the value for Cursor to be an explicit nil
func (o *OnChainTransactionsResponse) SetCursorNil() {
	o.Cursor.Set(nil)
}

// UnsetCursor ensures that no value is present for Cursor, not even an explicit nil
func (o *OnChainTransactionsResponse) UnsetCursor() {
	o.Cursor.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OnChainTransactionsResponse) GetResult() []OnChainTransactionDTO {
	if o == nil {
		var ret []OnChainTransactionDTO
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OnChainTransactionsResponse) GetResultOk() ([]OnChainTransactionDTO, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *OnChainTransactionsResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given []OnChainTransactionDTO and assigns it to the Result field.
func (o *OnChainTransactionsResponse) SetResult(v []OnChainTransactionDTO) {
	o.Result = v
}

func (o OnChainTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnChainTransactionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor.IsSet() {
		toSerialize["cursor"] = o.Cursor.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OnChainTransactionsResponse) UnmarshalJSON(data []byte) (err error) {
	varOnChainTransactionsResponse := _OnChainTransactionsResponse{}

	err = json.Unmarshal(data, &varOnChainTransactionsResponse)

	if err != nil {
		return err
	}

	*o = OnChainTransactionsResponse(varOnChainTransactionsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "cursor")
		delete(additionalProperties, "result")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOnChainTransactionsResponse struct {
	value *OnChainTransactionsResponse
	isSet bool
}

func (v NullableOnChainTransactionsResponse) Get() *OnChainTransactionsResponse {
	return v.value
}

func (v *NullableOnChainTransactionsResponse) Set(val *OnChainTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOnChainTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOnChainTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnChainTransactionsResponse(val *OnChainTransactionsResponse) *NullableOnChainTransactionsResponse {
	return &NullableOnChainTransactionsResponse{value: val, isSet: true}
}

func (v NullableOnChainTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnChainTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


