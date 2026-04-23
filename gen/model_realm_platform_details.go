
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the RealmPlatformDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RealmPlatformDetails{}

// RealmPlatformDetails struct for RealmPlatformDetails
type RealmPlatformDetails struct {
	AffiliateUrl NullableString `json:"affiliate_url,omitempty"`
	PackageId NullableString `json:"package_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RealmPlatformDetails RealmPlatformDetails

// NewRealmPlatformDetails instantiates a new RealmPlatformDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealmPlatformDetails() *RealmPlatformDetails {
	this := RealmPlatformDetails{}
	return &this
}

// NewRealmPlatformDetailsWithDefaults instantiates a new RealmPlatformDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealmPlatformDetailsWithDefaults() *RealmPlatformDetails {
	this := RealmPlatformDetails{}
	return &this
}

// GetAffiliateUrl returns the AffiliateUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmPlatformDetails) GetAffiliateUrl() string {
	if o == nil || IsNil(o.AffiliateUrl.Get()) {
		var ret string
		return ret
	}
	return *o.AffiliateUrl.Get()
}

// GetAffiliateUrlOk returns a tuple with the AffiliateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmPlatformDetails) GetAffiliateUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AffiliateUrl.Get(), o.AffiliateUrl.IsSet()
}

// HasAffiliateUrl returns a boolean if a field has been set.
func (o *RealmPlatformDetails) HasAffiliateUrl() bool {
	if o != nil && o.AffiliateUrl.IsSet() {
		return true
	}

	return false
}

// SetAffiliateUrl gets a reference to the given NullableString and assigns it to the AffiliateUrl field.
func (o *RealmPlatformDetails) SetAffiliateUrl(v string) {
	o.AffiliateUrl.Set(&v)
}
// SetAffiliateUrlNil sets the value for AffiliateUrl to be an explicit nil
func (o *RealmPlatformDetails) SetAffiliateUrlNil() {
	o.AffiliateUrl.Set(nil)
}

// UnsetAffiliateUrl ensures that no value is present for AffiliateUrl, not even an explicit nil
func (o *RealmPlatformDetails) UnsetAffiliateUrl() {
	o.AffiliateUrl.Unset()
}

// GetPackageId returns the PackageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RealmPlatformDetails) GetPackageId() string {
	if o == nil || IsNil(o.PackageId.Get()) {
		var ret string
		return ret
	}
	return *o.PackageId.Get()
}

// GetPackageIdOk returns a tuple with the PackageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RealmPlatformDetails) GetPackageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PackageId.Get(), o.PackageId.IsSet()
}

// HasPackageId returns a boolean if a field has been set.
func (o *RealmPlatformDetails) HasPackageId() bool {
	if o != nil && o.PackageId.IsSet() {
		return true
	}

	return false
}

// SetPackageId gets a reference to the given NullableString and assigns it to the PackageId field.
func (o *RealmPlatformDetails) SetPackageId(v string) {
	o.PackageId.Set(&v)
}
// SetPackageIdNil sets the value for PackageId to be an explicit nil
func (o *RealmPlatformDetails) SetPackageIdNil() {
	o.PackageId.Set(nil)
}

// UnsetPackageId ensures that no value is present for PackageId, not even an explicit nil
func (o *RealmPlatformDetails) UnsetPackageId() {
	o.PackageId.Unset()
}

func (o RealmPlatformDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealmPlatformDetails) ToMap() (map[string]interface{}, error) {
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

func (o *RealmPlatformDetails) UnmarshalJSON(data []byte) (err error) {
	varRealmPlatformDetails := _RealmPlatformDetails{}

	err = json.Unmarshal(data, &varRealmPlatformDetails)

	if err != nil {
		return err
	}

	*o = RealmPlatformDetails(varRealmPlatformDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "affiliate_url")
		delete(additionalProperties, "package_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRealmPlatformDetails struct {
	value *RealmPlatformDetails
	isSet bool
}

func (v NullableRealmPlatformDetails) Get() *RealmPlatformDetails {
	return v.value
}

func (v *NullableRealmPlatformDetails) Set(val *RealmPlatformDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableRealmPlatformDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableRealmPlatformDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealmPlatformDetails(val *RealmPlatformDetails) *NullableRealmPlatformDetails {
	return &NullableRealmPlatformDetails{value: val, isSet: true}
}

func (v NullableRealmPlatformDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealmPlatformDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


