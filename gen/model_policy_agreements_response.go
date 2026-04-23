
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PolicyAgreementsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAgreementsResponse{}

// PolicyAgreementsResponse struct for PolicyAgreementsResponse
type PolicyAgreementsResponse struct {
	LatestDotmoneyAgreed NullableBool `json:"latest_dotmoney_agreed,omitempty"`
	LatestEmpl2Agreed NullableBool `json:"latest_empl2_agreed,omitempty"`
	LatestPrivacyPolicyAgreed NullableBool `json:"latest_privacy_policy_agreed,omitempty"`
	LatestTermsOfUseAgreed NullableBool `json:"latest_terms_of_use_agreed,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PolicyAgreementsResponse PolicyAgreementsResponse

// NewPolicyAgreementsResponse instantiates a new PolicyAgreementsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAgreementsResponse() *PolicyAgreementsResponse {
	this := PolicyAgreementsResponse{}
	return &this
}

// NewPolicyAgreementsResponseWithDefaults instantiates a new PolicyAgreementsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAgreementsResponseWithDefaults() *PolicyAgreementsResponse {
	this := PolicyAgreementsResponse{}
	return &this
}

// GetLatestDotmoneyAgreed returns the LatestDotmoneyAgreed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAgreementsResponse) GetLatestDotmoneyAgreed() bool {
	if o == nil || IsNil(o.LatestDotmoneyAgreed.Get()) {
		var ret bool
		return ret
	}
	return *o.LatestDotmoneyAgreed.Get()
}

// GetLatestDotmoneyAgreedOk returns a tuple with the LatestDotmoneyAgreed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAgreementsResponse) GetLatestDotmoneyAgreedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestDotmoneyAgreed.Get(), o.LatestDotmoneyAgreed.IsSet()
}

// HasLatestDotmoneyAgreed returns a boolean if a field has been set.
func (o *PolicyAgreementsResponse) HasLatestDotmoneyAgreed() bool {
	if o != nil && o.LatestDotmoneyAgreed.IsSet() {
		return true
	}

	return false
}

// SetLatestDotmoneyAgreed gets a reference to the given NullableBool and assigns it to the LatestDotmoneyAgreed field.
func (o *PolicyAgreementsResponse) SetLatestDotmoneyAgreed(v bool) {
	o.LatestDotmoneyAgreed.Set(&v)
}
// SetLatestDotmoneyAgreedNil sets the value for LatestDotmoneyAgreed to be an explicit nil
func (o *PolicyAgreementsResponse) SetLatestDotmoneyAgreedNil() {
	o.LatestDotmoneyAgreed.Set(nil)
}

// UnsetLatestDotmoneyAgreed ensures that no value is present for LatestDotmoneyAgreed, not even an explicit nil
func (o *PolicyAgreementsResponse) UnsetLatestDotmoneyAgreed() {
	o.LatestDotmoneyAgreed.Unset()
}

// GetLatestEmpl2Agreed returns the LatestEmpl2Agreed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAgreementsResponse) GetLatestEmpl2Agreed() bool {
	if o == nil || IsNil(o.LatestEmpl2Agreed.Get()) {
		var ret bool
		return ret
	}
	return *o.LatestEmpl2Agreed.Get()
}

// GetLatestEmpl2AgreedOk returns a tuple with the LatestEmpl2Agreed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAgreementsResponse) GetLatestEmpl2AgreedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestEmpl2Agreed.Get(), o.LatestEmpl2Agreed.IsSet()
}

// HasLatestEmpl2Agreed returns a boolean if a field has been set.
func (o *PolicyAgreementsResponse) HasLatestEmpl2Agreed() bool {
	if o != nil && o.LatestEmpl2Agreed.IsSet() {
		return true
	}

	return false
}

// SetLatestEmpl2Agreed gets a reference to the given NullableBool and assigns it to the LatestEmpl2Agreed field.
func (o *PolicyAgreementsResponse) SetLatestEmpl2Agreed(v bool) {
	o.LatestEmpl2Agreed.Set(&v)
}
// SetLatestEmpl2AgreedNil sets the value for LatestEmpl2Agreed to be an explicit nil
func (o *PolicyAgreementsResponse) SetLatestEmpl2AgreedNil() {
	o.LatestEmpl2Agreed.Set(nil)
}

// UnsetLatestEmpl2Agreed ensures that no value is present for LatestEmpl2Agreed, not even an explicit nil
func (o *PolicyAgreementsResponse) UnsetLatestEmpl2Agreed() {
	o.LatestEmpl2Agreed.Unset()
}

// GetLatestPrivacyPolicyAgreed returns the LatestPrivacyPolicyAgreed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAgreementsResponse) GetLatestPrivacyPolicyAgreed() bool {
	if o == nil || IsNil(o.LatestPrivacyPolicyAgreed.Get()) {
		var ret bool
		return ret
	}
	return *o.LatestPrivacyPolicyAgreed.Get()
}

// GetLatestPrivacyPolicyAgreedOk returns a tuple with the LatestPrivacyPolicyAgreed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAgreementsResponse) GetLatestPrivacyPolicyAgreedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestPrivacyPolicyAgreed.Get(), o.LatestPrivacyPolicyAgreed.IsSet()
}

// HasLatestPrivacyPolicyAgreed returns a boolean if a field has been set.
func (o *PolicyAgreementsResponse) HasLatestPrivacyPolicyAgreed() bool {
	if o != nil && o.LatestPrivacyPolicyAgreed.IsSet() {
		return true
	}

	return false
}

// SetLatestPrivacyPolicyAgreed gets a reference to the given NullableBool and assigns it to the LatestPrivacyPolicyAgreed field.
func (o *PolicyAgreementsResponse) SetLatestPrivacyPolicyAgreed(v bool) {
	o.LatestPrivacyPolicyAgreed.Set(&v)
}
// SetLatestPrivacyPolicyAgreedNil sets the value for LatestPrivacyPolicyAgreed to be an explicit nil
func (o *PolicyAgreementsResponse) SetLatestPrivacyPolicyAgreedNil() {
	o.LatestPrivacyPolicyAgreed.Set(nil)
}

// UnsetLatestPrivacyPolicyAgreed ensures that no value is present for LatestPrivacyPolicyAgreed, not even an explicit nil
func (o *PolicyAgreementsResponse) UnsetLatestPrivacyPolicyAgreed() {
	o.LatestPrivacyPolicyAgreed.Unset()
}

// GetLatestTermsOfUseAgreed returns the LatestTermsOfUseAgreed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAgreementsResponse) GetLatestTermsOfUseAgreed() bool {
	if o == nil || IsNil(o.LatestTermsOfUseAgreed.Get()) {
		var ret bool
		return ret
	}
	return *o.LatestTermsOfUseAgreed.Get()
}

// GetLatestTermsOfUseAgreedOk returns a tuple with the LatestTermsOfUseAgreed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAgreementsResponse) GetLatestTermsOfUseAgreedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestTermsOfUseAgreed.Get(), o.LatestTermsOfUseAgreed.IsSet()
}

// HasLatestTermsOfUseAgreed returns a boolean if a field has been set.
func (o *PolicyAgreementsResponse) HasLatestTermsOfUseAgreed() bool {
	if o != nil && o.LatestTermsOfUseAgreed.IsSet() {
		return true
	}

	return false
}

// SetLatestTermsOfUseAgreed gets a reference to the given NullableBool and assigns it to the LatestTermsOfUseAgreed field.
func (o *PolicyAgreementsResponse) SetLatestTermsOfUseAgreed(v bool) {
	o.LatestTermsOfUseAgreed.Set(&v)
}
// SetLatestTermsOfUseAgreedNil sets the value for LatestTermsOfUseAgreed to be an explicit nil
func (o *PolicyAgreementsResponse) SetLatestTermsOfUseAgreedNil() {
	o.LatestTermsOfUseAgreed.Set(nil)
}

// UnsetLatestTermsOfUseAgreed ensures that no value is present for LatestTermsOfUseAgreed, not even an explicit nil
func (o *PolicyAgreementsResponse) UnsetLatestTermsOfUseAgreed() {
	o.LatestTermsOfUseAgreed.Unset()
}

func (o PolicyAgreementsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAgreementsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.LatestDotmoneyAgreed.IsSet() {
		toSerialize["latest_dotmoney_agreed"] = o.LatestDotmoneyAgreed.Get()
	}
	if o.LatestEmpl2Agreed.IsSet() {
		toSerialize["latest_empl2_agreed"] = o.LatestEmpl2Agreed.Get()
	}
	if o.LatestPrivacyPolicyAgreed.IsSet() {
		toSerialize["latest_privacy_policy_agreed"] = o.LatestPrivacyPolicyAgreed.Get()
	}
	if o.LatestTermsOfUseAgreed.IsSet() {
		toSerialize["latest_terms_of_use_agreed"] = o.LatestTermsOfUseAgreed.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PolicyAgreementsResponse) UnmarshalJSON(data []byte) (err error) {
	varPolicyAgreementsResponse := _PolicyAgreementsResponse{}

	err = json.Unmarshal(data, &varPolicyAgreementsResponse)

	if err != nil {
		return err
	}

	*o = PolicyAgreementsResponse(varPolicyAgreementsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "latest_dotmoney_agreed")
		delete(additionalProperties, "latest_empl2_agreed")
		delete(additionalProperties, "latest_privacy_policy_agreed")
		delete(additionalProperties, "latest_terms_of_use_agreed")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePolicyAgreementsResponse struct {
	value *PolicyAgreementsResponse
	isSet bool
}

func (v NullablePolicyAgreementsResponse) Get() *PolicyAgreementsResponse {
	return v.value
}

func (v *NullablePolicyAgreementsResponse) Set(val *PolicyAgreementsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAgreementsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAgreementsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAgreementsResponse(val *PolicyAgreementsResponse) *NullablePolicyAgreementsResponse {
	return &NullablePolicyAgreementsResponse{value: val, isSet: true}
}

func (v NullablePolicyAgreementsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAgreementsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


