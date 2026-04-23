
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PlatformDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlatformDetails{}

// PlatformDetails struct for PlatformDetails
type PlatformDetails struct {
	AffiliateUrl NullableString `json:"affiliate_url,omitempty"`
	PackageId NullableString `json:"package_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PlatformDetails PlatformDetails

// NewPlatformDetails instantiates a new PlatformDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformDetails() *PlatformDetails {
	this := PlatformDetails{}
	return &this
}

// NewPlatformDetailsWithDefaults instantiates a new PlatformDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformDetailsWithDefaults() *PlatformDetails {
	this := PlatformDetails{}
	return &this
}

// GetAffiliateUrl returns the AffiliateUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDetails) GetAffiliateUrl() string {
	if o == nil || IsNil(o.AffiliateUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AffiliateUrl.Get()
}

// GetAffiliateUrlOk returns a tuple with the AffiliateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDetails) GetAffiliateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AffiliateUrl.Get(), o.AffiliateUrl.IsSet()
}

// HasAffiliateUrl returns a boolean if a field has been set.
func (o *PlatformDetails) HasAffiliateUrl() bool {
	if o != nil && o.AffiliateUrl.IsSet() {
		return true
	}

	return false
}

// SetAffiliateUrl gets a reference to the given NullableString and assigns it to the AffiliateUrl field.
func (o *PlatformDetails) SetAffiliateUrl(v string) {
	o.AffiliateUrl.Set(&v)
}
// SetAffiliateUrlNil sets the value for AffiliateUrl to be an explicit nil
func (o *PlatformDetails) SetAffiliateUrlNil() {
	o.AffiliateUrl.Set(nil)
}

// UnsetAffiliateUrl ensures that no value is present for AffiliateUrl, not even an explicit nil
func (o *PlatformDetails) UnsetAffiliateUrl() {
	o.AffiliateUrl.Unset()
}

// GetPackageId returns the PackageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlatformDetails) GetPackageId() string {
	if o == nil || IsNil(o.PackageId.Get()) {
		var ret string
		return ret
	}
	return *o.PackageId.Get()
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlatformDetails) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageId.Get(), o.PackageId.IsSet()
}

// HasPackageId returns a boolean if a field has been set.
func (o *PlatformDetails) HasPackageId() bool {
	if o != nil && o.PackageId.IsSet() {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given NullableString and assigns it to the PackageId field.
func (o *PlatformDetails) SetPackageId(v string) {
	o.PackageId.Set(&v)
}
// SetPackageIdNil sets the value for PackageId to be an explicit nil
func (o *PlatformDetails) SetPackageIdNil() {
	o.PackageId.Set(nil)
}

// UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
func (o *PlatformDetails) UnsetPackageId() {
	o.PackageId.Unset()
}

func (o PlatformDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlatformDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AffiliateUrl.IsSet() {
		toSerialize["affiliate_url"] = o.AffiliateUrl.Get()
	}
	if o.PackageId.IsSet() {
		toSerialize["package_id"] = o.PackageId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PlatformDetails) UnmarshalJSON(data []byte) (err error) {
	varPlatformDetails := _PlatformDetails{}

	err = json.Unmarshal(data, &varPlatformDetails)

	if err != nil {
		return err
	}

	*o = PlatformDetails(varPlatformDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "affiliate_url")
		delete(additionalProperties, "package_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePlatformDetails struct {
	value *PlatformDetails
	isSet bool
}

func (v NullablePlatformDetails) Get() *PlatformDetails {
	return v.value
}

func (v *NullablePlatformDetails) Set(val *PlatformDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePlatformDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePlatformDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlatformDetails(val *PlatformDetails) *NullablePlatformDetails {
	return &NullablePlatformDetails{value: val, isSet: true}
}

func (v NullablePlatformDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlatformDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


