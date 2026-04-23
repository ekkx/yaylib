
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CreateQuotaResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateQuotaResponse{}

// CreateQuotaResponse struct for CreateQuotaResponse
type CreateQuotaResponse struct {
	Create NullableCreateGroupQuota `json:"create,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateQuotaResponse CreateQuotaResponse

// NewCreateQuotaResponse instantiates a new CreateQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateQuotaResponse() *CreateQuotaResponse {
	this := CreateQuotaResponse{}
	return &this
}

// NewCreateQuotaResponseWithDefaults instantiates a new CreateQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateQuotaResponseWithDefaults() *CreateQuotaResponse {
	this := CreateQuotaResponse{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateQuotaResponse) GetCreate() CreateGroupQuota {
	if o == nil || IsNil(o.Create.Get()) {
		var ret CreateGroupQuota
		return ret
	}
	return *o.Create.Get()
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateQuotaResponse) GetCreateOk() (*CreateGroupQuota, bool) {
	if o == nil {
		return nil, false
	}
	return o.Create.Get(), o.Create.IsSet()
}

// HasCreate returns a boolean if a field has been set.
func (o *CreateQuotaResponse) HasCreate() bool {
	if o != nil && o.Create.IsSet() {
		return true
	}

	return false
}

// SetCreate gets a reference to the given NullableCreateGroupQuota and assigns it to the Create field.
func (o *CreateQuotaResponse) SetCreate(v CreateGroupQuota) {
	o.Create.Set(&v)
}
// SetCreateNil sets the value for Create to be an explicit nil
func (o *CreateQuotaResponse) SetCreateNil() {
	o.Create.Set(nil)
}

// UnsetCreate ensures that no value is present for Create, not even an explicit nil
func (o *CreateQuotaResponse) UnsetCreate() {
	o.Create.Unset()
}

func (o CreateQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateQuotaResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Create.IsSet() {
		toSerialize["create"] = o.Create.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateQuotaResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateQuotaResponse := _CreateQuotaResponse{}

	err = json.Unmarshal(data, &varCreateQuotaResponse)

	if err != nil {
		return err
	}

	*o = CreateQuotaResponse(varCreateQuotaResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "create")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateQuotaResponse struct {
	value *CreateQuotaResponse
	isSet bool
}

func (v NullableCreateQuotaResponse) Get() *CreateQuotaResponse {
	return v.value
}

func (v *NullableCreateQuotaResponse) Set(val *CreateQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateQuotaResponse(val *CreateQuotaResponse) *NullableCreateQuotaResponse {
	return &NullableCreateQuotaResponse{value: val, isSet: true}
}

func (v NullableCreateQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


