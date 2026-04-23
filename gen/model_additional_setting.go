
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AdditionalSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalSetting{}

// AdditionalSetting struct for AdditionalSetting
type AdditionalSetting struct {
	ApiId NullableString `json:"api_id,omitempty"`
	// Unresolved Java type: kotlin.jvm.functions.Function1
	IsOn map[string]interface{} `json:"is_on,omitempty"`
	// Unresolved Java type: kotlin.jvm.functions.Function1
	IsOnFromUser map[string]interface{} `json:"is_on_from_user,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdditionalSetting AdditionalSetting

// NewAdditionalSetting instantiates a new AdditionalSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSetting() *AdditionalSetting {
	this := AdditionalSetting{}
	return &this
}

// NewAdditionalSettingWithDefaults instantiates a new AdditionalSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSettingWithDefaults() *AdditionalSetting {
	this := AdditionalSetting{}
	return &this
}

// GetApiId returns the ApiId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionalSetting) GetApiId() string {
	if o == nil || IsNil(o.ApiId.Get()) {
		var ret string
		return ret
	}
	return *o.ApiId.Get()
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionalSetting) GetApiIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiId.Get(), o.ApiId.IsSet()
}

// HasApiId returns a boolean if a field has been set.
func (o *AdditionalSetting) HasApiId() bool {
	if o != nil && o.ApiId.IsSet() {
		return true
	}

	return false
}

// SetApiId gets a reference to the given NullableString and assigns it to the ApiId field.
func (o *AdditionalSetting) SetApiId(v string) {
	o.ApiId.Set(&v)
}
// SetApiIdNil sets the value for ApiId to be an explicit nil
func (o *AdditionalSetting) SetApiIdNil() {
	o.ApiId.Set(nil)
}

// UnsetApiId ensures that no value is present for ApiId, not even an explicit nil
func (o *AdditionalSetting) UnsetApiId() {
	o.ApiId.Unset()
}

// GetIsOn returns the IsOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionalSetting) GetIsOn() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IsOn
}

// GetIsOnOk returns a tuple with the IsOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionalSetting) GetIsOnOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IsOn) {
		return map[string]interface{}{}, false
	}
	return o.IsOn, true
}

// HasIsOn returns a boolean if a field has been set.
func (o *AdditionalSetting) HasIsOn() bool {
	if o != nil && !IsNil(o.IsOn) {
		return true
	}

	return false
}

// SetIsOn gets a reference to the given map[string]interface{} and assigns it to the IsOn field.
func (o *AdditionalSetting) SetIsOn(v map[string]interface{}) {
	o.IsOn = v
}

// GetIsOnFromUser returns the IsOnFromUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionalSetting) GetIsOnFromUser() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.IsOnFromUser
}

// GetIsOnFromUserOk returns a tuple with the IsOnFromUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionalSetting) GetIsOnFromUserOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.IsOnFromUser) {
		return map[string]interface{}{}, false
	}
	return o.IsOnFromUser, true
}

// HasIsOnFromUser returns a boolean if a field has been set.
func (o *AdditionalSetting) HasIsOnFromUser() bool {
	if o != nil && !IsNil(o.IsOnFromUser) {
		return true
	}

	return false
}

// SetIsOnFromUser gets a reference to the given map[string]interface{} and assigns it to the IsOnFromUser field.
func (o *AdditionalSetting) SetIsOnFromUser(v map[string]interface{}) {
	o.IsOnFromUser = v
}

func (o AdditionalSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiId.IsSet() {
		toSerialize["api_id"] = o.ApiId.Get()
	}
	if o.IsOn != nil {
		toSerialize["is_on"] = o.IsOn
	}
	if o.IsOnFromUser != nil {
		toSerialize["is_on_from_user"] = o.IsOnFromUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdditionalSetting) UnmarshalJSON(data []byte) (err error) {
	varAdditionalSetting := _AdditionalSetting{}

	err = json.Unmarshal(data, &varAdditionalSetting)

	if err != nil {
		return err
	}

	*o = AdditionalSetting(varAdditionalSetting)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "api_id")
		delete(additionalProperties, "is_on")
		delete(additionalProperties, "is_on_from_user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdditionalSetting struct {
	value *AdditionalSetting
	isSet bool
}

func (v NullableAdditionalSetting) Get() *AdditionalSetting {
	return v.value
}

func (v *NullableAdditionalSetting) Set(val *AdditionalSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSetting(val *AdditionalSetting) *NullableAdditionalSetting {
	return &NullableAdditionalSetting{value: val, isSet: true}
}

func (v NullableAdditionalSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


