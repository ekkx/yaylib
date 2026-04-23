
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CommonErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonErrorResponse{}

// CommonErrorResponse struct for CommonErrorResponse
type CommonErrorResponse struct {
	BanUntil NullableInt64 `json:"ban_until,omitempty"`
	ErrorCode NullableInt32 `json:"error_code,omitempty"`
	Message NullableString `json:"message,omitempty"`
	RemainingQuota NullableInt64 `json:"remaining_quota,omitempty"`
	RetryIn NullableInt64 `json:"retry_in,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CommonErrorResponse CommonErrorResponse

// NewCommonErrorResponse instantiates a new CommonErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonErrorResponse() *CommonErrorResponse {
	this := CommonErrorResponse{}
	return &this
}

// NewCommonErrorResponseWithDefaults instantiates a new CommonErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonErrorResponseWithDefaults() *CommonErrorResponse {
	this := CommonErrorResponse{}
	return &this
}

// GetBanUntil returns the BanUntil field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonErrorResponse) GetBanUntil() int64 {
	if o == nil || IsNil(o.BanUntil.Get()) {
		var ret int64
		return ret
	}
	return *o.BanUntil.Get()
}

// GetBanUntilOk returns a tuple with the BanUntil field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonErrorResponse) GetBanUntilOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BanUntil.Get(), o.BanUntil.IsSet()
}

// HasBanUntil returns a boolean if a field has been set.
func (o *CommonErrorResponse) HasBanUntil() bool {
	if o != nil && o.BanUntil.IsSet() {
		return true
	}

	return false
}

// SetBanUntil gets a reference to the given NullableInt64 and assigns it to the BanUntil field.
func (o *CommonErrorResponse) SetBanUntil(v int64) {
	o.BanUntil.Set(&v)
}
// SetBanUntilNil sets the value for BanUntil to be an explicit nil
func (o *CommonErrorResponse) SetBanUntilNil() {
	o.BanUntil.Set(nil)
}

// UnsetBanUntil ensures that no value is present for BanUntil, not even an explicit nil
func (o *CommonErrorResponse) UnsetBanUntil() {
	o.BanUntil.Unset()
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonErrorResponse) GetErrorCode() int32 {
	if o == nil || IsNil(o.ErrorCode.Get()) {
		var ret int32
		return ret
	}
	return *o.ErrorCode.Get()
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonErrorResponse) GetErrorCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ErrorCode.Get(), o.ErrorCode.IsSet()
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CommonErrorResponse) HasErrorCode() bool {
	if o != nil && o.ErrorCode.IsSet() {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given NullableInt32 and assigns it to the ErrorCode field.
func (o *CommonErrorResponse) SetErrorCode(v int32) {
	o.ErrorCode.Set(&v)
}
// SetErrorCodeNil sets the value for ErrorCode to be an explicit nil
func (o *CommonErrorResponse) SetErrorCodeNil() {
	o.ErrorCode.Set(nil)
}

// UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
func (o *CommonErrorResponse) UnsetErrorCode() {
	o.ErrorCode.Unset()
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonErrorResponse) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *CommonErrorResponse) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *CommonErrorResponse) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *CommonErrorResponse) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *CommonErrorResponse) UnsetMessage() {
	o.Message.Unset()
}

// GetRemainingQuota returns the RemainingQuota field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonErrorResponse) GetRemainingQuota() int64 {
	if o == nil || IsNil(o.RemainingQuota.Get()) {
		var ret int64
		return ret
	}
	return *o.RemainingQuota.Get()
}

// GetRemainingQuotaOk returns a tuple with the RemainingQuota field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonErrorResponse) GetRemainingQuotaOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemainingQuota.Get(), o.RemainingQuota.IsSet()
}

// HasRemainingQuota returns a boolean if a field has been set.
func (o *CommonErrorResponse) HasRemainingQuota() bool {
	if o != nil && o.RemainingQuota.IsSet() {
		return true
	}

	return false
}

// SetRemainingQuota gets a reference to the given NullableInt64 and assigns it to the RemainingQuota field.
func (o *CommonErrorResponse) SetRemainingQuota(v int64) {
	o.RemainingQuota.Set(&v)
}
// SetRemainingQuotaNil sets the value for RemainingQuota to be an explicit nil
func (o *CommonErrorResponse) SetRemainingQuotaNil() {
	o.RemainingQuota.Set(nil)
}

// UnsetRemainingQuota ensures that no value is present for RemainingQuota, not even an explicit nil
func (o *CommonErrorResponse) UnsetRemainingQuota() {
	o.RemainingQuota.Unset()
}

// GetRetryIn returns the RetryIn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommonErrorResponse) GetRetryIn() int64 {
	if o == nil || IsNil(o.RetryIn.Get()) {
		var ret int64
		return ret
	}
	return *o.RetryIn.Get()
}

// GetRetryInOk returns a tuple with the RetryIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommonErrorResponse) GetRetryInOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.RetryIn.Get(), o.RetryIn.IsSet()
}

// HasRetryIn returns a boolean if a field has been set.
func (o *CommonErrorResponse) HasRetryIn() bool {
	if o != nil && o.RetryIn.IsSet() {
		return true
	}

	return false
}

// SetRetryIn gets a reference to the given NullableInt64 and assigns it to the RetryIn field.
func (o *CommonErrorResponse) SetRetryIn(v int64) {
	o.RetryIn.Set(&v)
}
// SetRetryInNil sets the value for RetryIn to be an explicit nil
func (o *CommonErrorResponse) SetRetryInNil() {
	o.RetryIn.Set(nil)
}

// UnsetRetryIn ensures that no value is present for RetryIn, not even an explicit nil
func (o *CommonErrorResponse) UnsetRetryIn() {
	o.RetryIn.Unset()
}

func (o CommonErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BanUntil.IsSet() {
		toSerialize["ban_until"] = o.BanUntil.Get()
	}
	if o.ErrorCode.IsSet() {
		toSerialize["error_code"] = o.ErrorCode.Get()
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if o.RemainingQuota.IsSet() {
		toSerialize["remaining_quota"] = o.RemainingQuota.Get()
	}
	if o.RetryIn.IsSet() {
		toSerialize["retry_in"] = o.RetryIn.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CommonErrorResponse) UnmarshalJSON(data []byte) (err error) {
	varCommonErrorResponse := _CommonErrorResponse{}

	err = json.Unmarshal(data, &varCommonErrorResponse)

	if err != nil {
		return err
	}

	*o = CommonErrorResponse(varCommonErrorResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ban_until")
		delete(additionalProperties, "error_code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "remaining_quota")
		delete(additionalProperties, "retry_in")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCommonErrorResponse struct {
	value *CommonErrorResponse
	isSet bool
}

func (v NullableCommonErrorResponse) Get() *CommonErrorResponse {
	return v.value
}

func (v *NullableCommonErrorResponse) Set(val *CommonErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonErrorResponse(val *CommonErrorResponse) *NullableCommonErrorResponse {
	return &NullableCommonErrorResponse{value: val, isSet: true}
}

func (v NullableCommonErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


