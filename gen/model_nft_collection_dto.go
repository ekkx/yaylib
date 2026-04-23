
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the NftCollectionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NftCollectionDTO{}

// NftCollectionDTO struct for NftCollectionDTO
type NftCollectionDTO struct {
	ContractAddress NullableString `json:"contract_address,omitempty"`
	Image NullableString `json:"image,omitempty"`
	MediaType NullableString `json:"media_type,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Priority NullableInt32 `json:"priority,omitempty"`
	Status NullableString `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NftCollectionDTO NftCollectionDTO

// NewNftCollectionDTO instantiates a new NftCollectionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftCollectionDTO() *NftCollectionDTO {
	this := NftCollectionDTO{}
	return &this
}

// NewNftCollectionDTOWithDefaults instantiates a new NftCollectionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftCollectionDTOWithDefaults() *NftCollectionDTO {
	this := NftCollectionDTO{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftCollectionDTO) GetContractAddress() string {
	if o == nil || IsNil(o.ContractAddress.Get()) {
		var ret string
		return ret
	}
	return *o.ContractAddress.Get()
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftCollectionDTO) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractAddress.Get(), o.ContractAddress.IsSet()
}

// HasContractAddress returns a boolean if a field has been set.
func (o *NftCollectionDTO) HasContractAddress() bool {
	if o != nil && o.ContractAddress.IsSet() {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given NullableString and assigns it to the ContractAddress field.
func (o *NftCollectionDTO) SetContractAddress(v string) {
	o.ContractAddress.Set(&v)
}
// SetContractAddressNil sets the value for ContractAddress to be an explicit nil
func (o *NftCollectionDTO) SetContractAddressNil() {
	o.ContractAddress.Set(nil)
}

// UnsetContractAddress ensures that no value is present for ContractAddress, not even an explicit nil
func (o *NftCollectionDTO) UnsetContractAddress() {
	o.ContractAddress.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftCollectionDTO) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftCollectionDTO) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *NftCollectionDTO) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *NftCollectionDTO) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *NftCollectionDTO) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *NftCollectionDTO) UnsetImage() {
	o.Image.Unset()
}

// GetMediaType returns the MediaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftCollectionDTO) GetMediaType() string {
	if o == nil || IsNil(o.MediaType.Get()) {
		var ret string
		return ret
	}
	return *o.MediaType.Get()
}

// GetMediaTypeOk returns a tuple with the MediaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftCollectionDTO) GetMediaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MediaType.Get(), o.MediaType.IsSet()
}

// HasMediaType returns a boolean if a field has been set.
func (o *NftCollectionDTO) HasMediaType() bool {
	if o != nil && o.MediaType.IsSet() {
		return true
	}

	return false
}

// SetMediaType gets a reference to the given NullableString and assigns it to the MediaType field.
func (o *NftCollectionDTO) SetMediaType(v string) {
	o.MediaType.Set(&v)
}
// SetMediaTypeNil sets the value for MediaType to be an explicit nil
func (o *NftCollectionDTO) SetMediaTypeNil() {
	o.MediaType.Set(nil)
}

// UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
func (o *NftCollectionDTO) UnsetMediaType() {
	o.MediaType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftCollectionDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftCollectionDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NftCollectionDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NftCollectionDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *NftCollectionDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NftCollectionDTO) UnsetName() {
	o.Name.Unset()
}

// GetPriority returns the Priority field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftCollectionDTO) GetPriority() int32 {
	if o == nil || IsNil(o.Priority.Get()) {
		var ret int32
		return ret
	}
	return *o.Priority.Get()
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftCollectionDTO) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Priority.Get(), o.Priority.IsSet()
}

// HasPriority returns a boolean if a field has been set.
func (o *NftCollectionDTO) HasPriority() bool {
	if o != nil && o.Priority.IsSet() {
		return true
	}

	return false
}

// SetPriority gets a reference to the given NullableInt32 and assigns it to the Priority field.
func (o *NftCollectionDTO) SetPriority(v int32) {
	o.Priority.Set(&v)
}
// SetPriorityNil sets the value for Priority to be an explicit nil
func (o *NftCollectionDTO) SetPriorityNil() {
	o.Priority.Set(nil)
}

// UnsetPriority ensures that no value is present for Priority, not even an explicit nil
func (o *NftCollectionDTO) UnsetPriority() {
	o.Priority.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NftCollectionDTO) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NftCollectionDTO) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *NftCollectionDTO) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *NftCollectionDTO) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *NftCollectionDTO) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *NftCollectionDTO) UnsetStatus() {
	o.Status.Unset()
}

func (o NftCollectionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NftCollectionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractAddress.IsSet() {
		toSerialize["contract_address"] = o.ContractAddress.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.MediaType.IsSet() {
		toSerialize["media_type"] = o.MediaType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Priority.IsSet() {
		toSerialize["priority"] = o.Priority.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NftCollectionDTO) UnmarshalJSON(data []byte) (err error) {
	varNftCollectionDTO := _NftCollectionDTO{}

	err = json.Unmarshal(data, &varNftCollectionDTO)

	if err != nil {
		return err
	}

	*o = NftCollectionDTO(varNftCollectionDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_address")
		delete(additionalProperties, "image")
		delete(additionalProperties, "media_type")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNftCollectionDTO struct {
	value *NftCollectionDTO
	isSet bool
}

func (v NullableNftCollectionDTO) Get() *NftCollectionDTO {
	return v.value
}

func (v *NullableNftCollectionDTO) Set(val *NftCollectionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableNftCollectionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableNftCollectionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftCollectionDTO(val *NftCollectionDTO) *NullableNftCollectionDTO {
	return &NullableNftCollectionDTO{value: val, isSet: true}
}

func (v NullableNftCollectionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftCollectionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


