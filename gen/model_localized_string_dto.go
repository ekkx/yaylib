
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the LocalizedStringDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalizedStringDTO{}

// LocalizedStringDTO struct for LocalizedStringDTO
type LocalizedStringDTO struct {
	En NullableString `json:"en,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Ja NullableString `json:"ja,omitempty"`
	Ko NullableString `json:"ko,omitempty"`
	Th NullableString `json:"th,omitempty"`
	ZhCn NullableString `json:"zh_cn,omitempty"`
	ZhHk NullableString `json:"zh_hk,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _LocalizedStringDTO LocalizedStringDTO

// NewLocalizedStringDTO instantiates a new LocalizedStringDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalizedStringDTO() *LocalizedStringDTO {
	this := LocalizedStringDTO{}
	return &this
}

// NewLocalizedStringDTOWithDefaults instantiates a new LocalizedStringDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalizedStringDTOWithDefaults() *LocalizedStringDTO {
	this := LocalizedStringDTO{}
	return &this
}

// GetEn returns the En field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetEn() string {
	if o == nil || IsNil(o.En.Get()) {
		var ret string
		return ret
	}
	return *o.En.Get()
}

// GetEnOk returns a tuple with the En field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetEnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.En.Get(), o.En.IsSet()
}

// HasEn returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasEn() bool {
	if o != nil && o.En.IsSet() {
		return true
	}

	return false
}

// SetEn gets a reference to the given NullableString and assigns it to the En field.
func (o *LocalizedStringDTO) SetEn(v string) {
	o.En.Set(&v)
}
// SetEnNil sets the value for En to be an explicit nil
func (o *LocalizedStringDTO) SetEnNil() {
	o.En.Set(nil)
}

// UnsetEn ensures that no value is present for En, not even an explicit nil
func (o *LocalizedStringDTO) UnsetEn() {
	o.En.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *LocalizedStringDTO) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *LocalizedStringDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *LocalizedStringDTO) UnsetId() {
	o.Id.Unset()
}

// GetJa returns the Ja field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetJa() string {
	if o == nil || IsNil(o.Ja.Get()) {
		var ret string
		return ret
	}
	return *o.Ja.Get()
}

// GetJaOk returns a tuple with the Ja field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetJaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ja.Get(), o.Ja.IsSet()
}

// HasJa returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasJa() bool {
	if o != nil && o.Ja.IsSet() {
		return true
	}

	return false
}

// SetJa gets a reference to the given NullableString and assigns it to the Ja field.
func (o *LocalizedStringDTO) SetJa(v string) {
	o.Ja.Set(&v)
}
// SetJaNil sets the value for Ja to be an explicit nil
func (o *LocalizedStringDTO) SetJaNil() {
	o.Ja.Set(nil)
}

// UnsetJa ensures that no value is present for Ja, not even an explicit nil
func (o *LocalizedStringDTO) UnsetJa() {
	o.Ja.Unset()
}

// GetKo returns the Ko field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetKo() string {
	if o == nil || IsNil(o.Ko.Get()) {
		var ret string
		return ret
	}
	return *o.Ko.Get()
}

// GetKoOk returns a tuple with the Ko field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetKoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ko.Get(), o.Ko.IsSet()
}

// HasKo returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasKo() bool {
	if o != nil && o.Ko.IsSet() {
		return true
	}

	return false
}

// SetKo gets a reference to the given NullableString and assigns it to the Ko field.
func (o *LocalizedStringDTO) SetKo(v string) {
	o.Ko.Set(&v)
}
// SetKoNil sets the value for Ko to be an explicit nil
func (o *LocalizedStringDTO) SetKoNil() {
	o.Ko.Set(nil)
}

// UnsetKo ensures that no value is present for Ko, not even an explicit nil
func (o *LocalizedStringDTO) UnsetKo() {
	o.Ko.Unset()
}

// GetTh returns the Th field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetTh() string {
	if o == nil || IsNil(o.Th.Get()) {
		var ret string
		return ret
	}
	return *o.Th.Get()
}

// GetThOk returns a tuple with the Th field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetThOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Th.Get(), o.Th.IsSet()
}

// HasTh returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasTh() bool {
	if o != nil && o.Th.IsSet() {
		return true
	}

	return false
}

// SetTh gets a reference to the given NullableString and assigns it to the Th field.
func (o *LocalizedStringDTO) SetTh(v string) {
	o.Th.Set(&v)
}
// SetThNil sets the value for Th to be an explicit nil
func (o *LocalizedStringDTO) SetThNil() {
	o.Th.Set(nil)
}

// UnsetTh ensures that no value is present for Th, not even an explicit nil
func (o *LocalizedStringDTO) UnsetTh() {
	o.Th.Unset()
}

// GetZhCn returns the ZhCn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetZhCn() string {
	if o == nil || IsNil(o.ZhCn.Get()) {
		var ret string
		return ret
	}
	return *o.ZhCn.Get()
}

// GetZhCnOk returns a tuple with the ZhCn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetZhCnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZhCn.Get(), o.ZhCn.IsSet()
}

// HasZhCn returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasZhCn() bool {
	if o != nil && o.ZhCn.IsSet() {
		return true
	}

	return false
}

// SetZhCn gets a reference to the given NullableString and assigns it to the ZhCn field.
func (o *LocalizedStringDTO) SetZhCn(v string) {
	o.ZhCn.Set(&v)
}
// SetZhCnNil sets the value for ZhCn to be an explicit nil
func (o *LocalizedStringDTO) SetZhCnNil() {
	o.ZhCn.Set(nil)
}

// UnsetZhCn ensures that no value is present for ZhCn, not even an explicit nil
func (o *LocalizedStringDTO) UnsetZhCn() {
	o.ZhCn.Unset()
}

// GetZhHk returns the ZhHk field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LocalizedStringDTO) GetZhHk() string {
	if o == nil || IsNil(o.ZhHk.Get()) {
		var ret string
		return ret
	}
	return *o.ZhHk.Get()
}

// GetZhHkOk returns a tuple with the ZhHk field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LocalizedStringDTO) GetZhHkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ZhHk.Get(), o.ZhHk.IsSet()
}

// HasZhHk returns a boolean if a field has been set.
func (o *LocalizedStringDTO) HasZhHk() bool {
	if o != nil && o.ZhHk.IsSet() {
		return true
	}

	return false
}

// SetZhHk gets a reference to the given NullableString and assigns it to the ZhHk field.
func (o *LocalizedStringDTO) SetZhHk(v string) {
	o.ZhHk.Set(&v)
}
// SetZhHkNil sets the value for ZhHk to be an explicit nil
func (o *LocalizedStringDTO) SetZhHkNil() {
	o.ZhHk.Set(nil)
}

// UnsetZhHk ensures that no value is present for ZhHk, not even an explicit nil
func (o *LocalizedStringDTO) UnsetZhHk() {
	o.ZhHk.Unset()
}

func (o LocalizedStringDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalizedStringDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.En.IsSet() {
		toSerialize["en"] = o.En.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Ja.IsSet() {
		toSerialize["ja"] = o.Ja.Get()
	}
	if o.Ko.IsSet() {
		toSerialize["ko"] = o.Ko.Get()
	}
	if o.Th.IsSet() {
		toSerialize["th"] = o.Th.Get()
	}
	if o.ZhCn.IsSet() {
		toSerialize["zh_cn"] = o.ZhCn.Get()
	}
	if o.ZhHk.IsSet() {
		toSerialize["zh_hk"] = o.ZhHk.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *LocalizedStringDTO) UnmarshalJSON(data []byte) (err error) {
	varLocalizedStringDTO := _LocalizedStringDTO{}

	err = json.Unmarshal(data, &varLocalizedStringDTO)

	if err != nil {
		return err
	}

	*o = LocalizedStringDTO(varLocalizedStringDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "en")
		delete(additionalProperties, "id")
		delete(additionalProperties, "ja")
		delete(additionalProperties, "ko")
		delete(additionalProperties, "th")
		delete(additionalProperties, "zh_cn")
		delete(additionalProperties, "zh_hk")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableLocalizedStringDTO struct {
	value *LocalizedStringDTO
	isSet bool
}

func (v NullableLocalizedStringDTO) Get() *LocalizedStringDTO {
	return v.value
}

func (v *NullableLocalizedStringDTO) Set(val *LocalizedStringDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalizedStringDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalizedStringDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalizedStringDTO(val *LocalizedStringDTO) *NullableLocalizedStringDTO {
	return &NullableLocalizedStringDTO{value: val, isSet: true}
}

func (v NullableLocalizedStringDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalizedStringDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


