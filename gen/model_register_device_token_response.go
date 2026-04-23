
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RegisterDeviceTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterDeviceTokenResponse{}

// RegisterDeviceTokenResponse struct for RegisterDeviceTokenResponse
type RegisterDeviceTokenResponse struct {
	CreatedAt NullableInt64 `json:"created_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	ServerDeviceId NullableString `json:"server_device_id,omitempty"`
	UpdatedAt NullableInt64 `json:"updated_at,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RegisterDeviceTokenResponse RegisterDeviceTokenResponse

// NewRegisterDeviceTokenResponse instantiates a new RegisterDeviceTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterDeviceTokenResponse() *RegisterDeviceTokenResponse {
	this := RegisterDeviceTokenResponse{}
	return &this
}

// NewRegisterDeviceTokenResponseWithDefaults instantiates a new RegisterDeviceTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterDeviceTokenResponseWithDefaults() *RegisterDeviceTokenResponse {
	this := RegisterDeviceTokenResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterDeviceTokenResponse) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterDeviceTokenResponse) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RegisterDeviceTokenResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableInt64 and assigns it to the CreatedAt field.
func (o *RegisterDeviceTokenResponse) SetCreatedAt(v int64) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *RegisterDeviceTokenResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *RegisterDeviceTokenResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterDeviceTokenResponse) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterDeviceTokenResponse) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *RegisterDeviceTokenResponse) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *RegisterDeviceTokenResponse) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *RegisterDeviceTokenResponse) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *RegisterDeviceTokenResponse) UnsetId() {
	o.Id.Unset()
}

// GetServerDeviceId returns the ServerDeviceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterDeviceTokenResponse) GetServerDeviceId() string {
	if o == nil || IsNil(o.ServerDeviceId.Get()) {
		var ret string
		return ret
	}
	return *o.ServerDeviceId.Get()
}

// GetServerDeviceIdOk returns a tuple with the ServerDeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterDeviceTokenResponse) GetServerDeviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerDeviceId.Get(), o.ServerDeviceId.IsSet()
}

// HasServerDeviceId returns a boolean if a field has been set.
func (o *RegisterDeviceTokenResponse) HasServerDeviceId() bool {
	if o != nil && o.ServerDeviceId.IsSet() {
		return true
	}

	return false
}

// SetServerDeviceId gets a reference to the given NullableString and assigns it to the ServerDeviceId field.
func (o *RegisterDeviceTokenResponse) SetServerDeviceId(v string) {
	o.ServerDeviceId.Set(&v)
}
// SetServerDeviceIdNil sets the value for ServerDeviceId to be an explicit nil
func (o *RegisterDeviceTokenResponse) SetServerDeviceIdNil() {
	o.ServerDeviceId.Set(nil)
}

// UnsetServerDeviceId ensures that no value is present for ServerDeviceId, not even an explicit nil
func (o *RegisterDeviceTokenResponse) UnsetServerDeviceId() {
	o.ServerDeviceId.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterDeviceTokenResponse) GetUpdatedAt() int64 {
	if o == nil || IsNil(o.UpdatedAt.Get()) {
		var ret int64
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterDeviceTokenResponse) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RegisterDeviceTokenResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableInt64 and assigns it to the UpdatedAt field.
func (o *RegisterDeviceTokenResponse) SetUpdatedAt(v int64) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *RegisterDeviceTokenResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *RegisterDeviceTokenResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RegisterDeviceTokenResponse) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RegisterDeviceTokenResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *RegisterDeviceTokenResponse) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *RegisterDeviceTokenResponse) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *RegisterDeviceTokenResponse) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *RegisterDeviceTokenResponse) UnsetUuid() {
	o.Uuid.Unset()
}

func (o RegisterDeviceTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterDeviceTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ServerDeviceId.IsSet() {
		toSerialize["server_device_id"] = o.ServerDeviceId.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RegisterDeviceTokenResponse) UnmarshalJSON(data []byte) (err error) {
	varRegisterDeviceTokenResponse := _RegisterDeviceTokenResponse{}

	err = json.Unmarshal(data, &varRegisterDeviceTokenResponse)

	if err != nil {
		return err
	}

	*o = RegisterDeviceTokenResponse(varRegisterDeviceTokenResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "server_device_id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRegisterDeviceTokenResponse struct {
	value *RegisterDeviceTokenResponse
	isSet bool
}

func (v NullableRegisterDeviceTokenResponse) Get() *RegisterDeviceTokenResponse {
	return v.value
}

func (v *NullableRegisterDeviceTokenResponse) Set(val *RegisterDeviceTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterDeviceTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterDeviceTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterDeviceTokenResponse(val *RegisterDeviceTokenResponse) *NullableRegisterDeviceTokenResponse {
	return &NullableRegisterDeviceTokenResponse{value: val, isSet: true}
}

func (v NullableRegisterDeviceTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterDeviceTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


