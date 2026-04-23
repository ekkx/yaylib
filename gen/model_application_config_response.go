
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the ApplicationConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationConfigResponse{}

// ApplicationConfigResponse struct for ApplicationConfigResponse
type ApplicationConfigResponse struct {
	App NullableApplication `json:"app,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplicationConfigResponse ApplicationConfigResponse

// NewApplicationConfigResponse instantiates a new ApplicationConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationConfigResponse() *ApplicationConfigResponse {
	this := ApplicationConfigResponse{}
	return &this
}

// NewApplicationConfigResponseWithDefaults instantiates a new ApplicationConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationConfigResponseWithDefaults() *ApplicationConfigResponse {
	this := ApplicationConfigResponse{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApplicationConfigResponse) GetApp() Application {
	if o == nil || IsNil(o.App.Get()) {
		var ret Application
		return ret
	}
	return *o.App.Get()
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationConfigResponse) GetAppOk() (*Application, bool) {
	if o == nil {
		return nil, false
	}
	return o.App.Get(), o.App.IsSet()
}

// HasApp returns a boolean if a field has been set.
func (o *ApplicationConfigResponse) HasApp() bool {
	if o != nil && o.App.IsSet() {
		return true
	}

	return false
}

// SetApp gets a reference to the given NullableApplication and assigns it to the App field.
func (o *ApplicationConfigResponse) SetApp(v Application) {
	o.App.Set(&v)
}
// SetAppNil sets the value for App to be an explicit nil
func (o *ApplicationConfigResponse) SetAppNil() {
	o.App.Set(nil)
}

// UnsetApp ensures that no value is present for App, not even an explicit nil
func (o *ApplicationConfigResponse) UnsetApp() {
	o.App.Unset()
}

func (o ApplicationConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.App.IsSet() {
		toSerialize["app"] = o.App.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ApplicationConfigResponse) UnmarshalJSON(data []byte) (err error) {
	varApplicationConfigResponse := _ApplicationConfigResponse{}

	err = json.Unmarshal(data, &varApplicationConfigResponse)

	if err != nil {
		return err
	}

	*o = ApplicationConfigResponse(varApplicationConfigResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "app")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplicationConfigResponse struct {
	value *ApplicationConfigResponse
	isSet bool
}

func (v NullableApplicationConfigResponse) Get() *ApplicationConfigResponse {
	return v.value
}

func (v *NullableApplicationConfigResponse) Set(val *ApplicationConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationConfigResponse(val *ApplicationConfigResponse) *NullableApplicationConfigResponse {
	return &NullableApplicationConfigResponse{value: val, isSet: true}
}

func (v NullableApplicationConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


