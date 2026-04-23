
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the AdditionalSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AdditionalSettingsResponse{}

// AdditionalSettingsResponse struct for AdditionalSettingsResponse
type AdditionalSettingsResponse struct {
	Settings NullableSettings `json:"settings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdditionalSettingsResponse AdditionalSettingsResponse

// NewAdditionalSettingsResponse instantiates a new AdditionalSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdditionalSettingsResponse() *AdditionalSettingsResponse {
	this := AdditionalSettingsResponse{}
	return &this
}

// NewAdditionalSettingsResponseWithDefaults instantiates a new AdditionalSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdditionalSettingsResponseWithDefaults() *AdditionalSettingsResponse {
	this := AdditionalSettingsResponse{}
	return &this
}

// GetSettings returns the Settings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdditionalSettingsResponse) GetSettings() Settings {
	if o == nil || IsNil(o.Settings.Get()) {
		var ret Settings
		return ret
	}
	return *o.Settings.Get()
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdditionalSettingsResponse) GetSettingsOk() (*Settings, bool) {
	if o == nil {
		return nil, false
	}
	return o.Settings.Get(), o.Settings.IsSet()
}

// HasSettings returns a boolean if a field has been set.
func (o *AdditionalSettingsResponse) HasSettings() bool {
	if o != nil && o.Settings.IsSet() {
		return true
	}

	return false
}

// SetSettings gets a reference to the given NullableSettings and assigns it to the Settings field.
func (o *AdditionalSettingsResponse) SetSettings(v Settings) {
	o.Settings.Set(&v)
}
// SetSettingsNil sets the value for Settings to be an explicit nil
func (o *AdditionalSettingsResponse) SetSettingsNil() {
	o.Settings.Set(nil)
}

// UnsetSettings ensures that no value is present for Settings, not even an explicit nil
func (o *AdditionalSettingsResponse) UnsetSettings() {
	o.Settings.Unset()
}

func (o AdditionalSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AdditionalSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Settings.IsSet() {
		toSerialize["settings"] = o.Settings.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AdditionalSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	varAdditionalSettingsResponse := _AdditionalSettingsResponse{}

	err = json.Unmarshal(data, &varAdditionalSettingsResponse)

	if err != nil {
		return err
	}

	*o = AdditionalSettingsResponse(varAdditionalSettingsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "settings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAdditionalSettingsResponse struct {
	value *AdditionalSettingsResponse
	isSet bool
}

func (v NullableAdditionalSettingsResponse) Get() *AdditionalSettingsResponse {
	return v.value
}

func (v *NullableAdditionalSettingsResponse) Set(val *AdditionalSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalSettingsResponse(val *AdditionalSettingsResponse) *NullableAdditionalSettingsResponse {
	return &NullableAdditionalSettingsResponse{value: val, isSet: true}
}

func (v NullableAdditionalSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


