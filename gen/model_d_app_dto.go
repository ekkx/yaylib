
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the DAppDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DAppDTO{}

// DAppDTO struct for DAppDTO
type DAppDTO struct {
	DescriptionLocalized NullableLocalizedStringDTO `json:"description_localized,omitempty"`
	Enabled NullableBool `json:"enabled,omitempty"`
	IconUrl NullableString `json:"icon_url,omitempty"`
	Id NullableString `json:"id,omitempty"`
	Order NullableInt32 `json:"order,omitempty"`
	SearchKeywords []string `json:"search_keywords,omitempty"`
	ShowInRecommendDApps NullableBool `json:"show_in_recommend_d_apps,omitempty"`
	TitleLocalized NullableLocalizedStringDTO `json:"title_localized,omitempty"`
	UrlLocalized NullableLocalizedStringDTO `json:"url_localized,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DAppDTO DAppDTO

// NewDAppDTO instantiates a new DAppDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDAppDTO() *DAppDTO {
	this := DAppDTO{}
	return &this
}

// NewDAppDTOWithDefaults instantiates a new DAppDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDAppDTOWithDefaults() *DAppDTO {
	this := DAppDTO{}
	return &this
}

// GetDescriptionLocalized returns the DescriptionLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetDescriptionLocalized() LocalizedStringDTO {
	if o == nil || IsNil(o.DescriptionLocalized.Get()) {
		var ret LocalizedStringDTO
		return ret
	}
	return *o.DescriptionLocalized.Get()
}

// GetDescriptionLocalizedOk returns a tuple with the DescriptionLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetDescriptionLocalizedOk() (*LocalizedStringDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescriptionLocalized.Get(), o.DescriptionLocalized.IsSet()
}

// HasDescriptionLocalized returns a boolean if a field has been set.
func (o *DAppDTO) HasDescriptionLocalized() bool {
	if o != nil && o.DescriptionLocalized.IsSet() {
		return true
	}

	return false
}

// SetDescriptionLocalized gets a reference to the given NullableLocalizedStringDTO and assigns it to the DescriptionLocalized field.
func (o *DAppDTO) SetDescriptionLocalized(v LocalizedStringDTO) {
	o.DescriptionLocalized.Set(&v)
}
// SetDescriptionLocalizedNil sets the value for DescriptionLocalized to be an explicit nil
func (o *DAppDTO) SetDescriptionLocalizedNil() {
	o.DescriptionLocalized.Set(nil)
}

// UnsetDescriptionLocalized ensures that no value is present for DescriptionLocalized, not even an explicit nil
func (o *DAppDTO) UnsetDescriptionLocalized() {
	o.DescriptionLocalized.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DAppDTO) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DAppDTO) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DAppDTO) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DAppDTO) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetIconUrl returns the IconUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetIconUrl() string {
	if o == nil || IsNil(o.IconUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IconUrl.Get()
}

// GetIconUrlOk returns a tuple with the IconUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetIconUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IconUrl.Get(), o.IconUrl.IsSet()
}

// HasIconUrl returns a boolean if a field has been set.
func (o *DAppDTO) HasIconUrl() bool {
	if o != nil && o.IconUrl.IsSet() {
		return true
	}

	return false
}

// SetIconUrl gets a reference to the given NullableString and assigns it to the IconUrl field.
func (o *DAppDTO) SetIconUrl(v string) {
	o.IconUrl.Set(&v)
}
// SetIconUrlNil sets the value for IconUrl to be an explicit nil
func (o *DAppDTO) SetIconUrlNil() {
	o.IconUrl.Set(nil)
}

// UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
func (o *DAppDTO) UnsetIconUrl() {
	o.IconUrl.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetId() string {
	if o == nil || IsNil(o.Id.Get()) {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *DAppDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *DAppDTO) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *DAppDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *DAppDTO) UnsetId() {
	o.Id.Unset()
}

// GetOrder returns the Order field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetOrder() int32 {
	if o == nil || IsNil(o.Order.Get()) {
		var ret int32
		return ret
	}
	return *o.Order.Get()
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Order.Get(), o.Order.IsSet()
}

// HasOrder returns a boolean if a field has been set.
func (o *DAppDTO) HasOrder() bool {
	if o != nil && o.Order.IsSet() {
		return true
	}

	return false
}

// SetOrder gets a reference to the given NullableInt32 and assigns it to the Order field.
func (o *DAppDTO) SetOrder(v int32) {
	o.Order.Set(&v)
}
// SetOrderNil sets the value for Order to be an explicit nil
func (o *DAppDTO) SetOrderNil() {
	o.Order.Set(nil)
}

// UnsetOrder ensures that no value is present for Order, not even an explicit nil
func (o *DAppDTO) UnsetOrder() {
	o.Order.Unset()
}

// GetSearchKeywords returns the SearchKeywords field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetSearchKeywords() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.SearchKeywords
}

// GetSearchKeywordsOk returns a tuple with the SearchKeywords field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetSearchKeywordsOk() ([]string, bool) {
	if o == nil || IsNil(o.SearchKeywords) {
		return nil, false
	}
	return o.SearchKeywords, true
}

// HasSearchKeywords returns a boolean if a field has been set.
func (o *DAppDTO) HasSearchKeywords() bool {
	if o != nil && !IsNil(o.SearchKeywords) {
		return true
	}

	return false
}

// SetSearchKeywords gets a reference to the given []string and assigns it to the SearchKeywords field.
func (o *DAppDTO) SetSearchKeywords(v []string) {
	o.SearchKeywords = v
}

// GetShowInRecommendDApps returns the ShowInRecommendDApps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetShowInRecommendDApps() bool {
	if o == nil || IsNil(o.ShowInRecommendDApps.Get()) {
		var ret bool
		return ret
	}
	return *o.ShowInRecommendDApps.Get()
}

// GetShowInRecommendDAppsOk returns a tuple with the ShowInRecommendDApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetShowInRecommendDAppsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShowInRecommendDApps.Get(), o.ShowInRecommendDApps.IsSet()
}

// HasShowInRecommendDApps returns a boolean if a field has been set.
func (o *DAppDTO) HasShowInRecommendDApps() bool {
	if o != nil && o.ShowInRecommendDApps.IsSet() {
		return true
	}

	return false
}

// SetShowInRecommendDApps gets a reference to the given NullableBool and assigns it to the ShowInRecommendDApps field.
func (o *DAppDTO) SetShowInRecommendDApps(v bool) {
	o.ShowInRecommendDApps.Set(&v)
}
// SetShowInRecommendDAppsNil sets the value for ShowInRecommendDApps to be an explicit nil
func (o *DAppDTO) SetShowInRecommendDAppsNil() {
	o.ShowInRecommendDApps.Set(nil)
}

// UnsetShowInRecommendDApps ensures that no value is present for ShowInRecommendDApps, not even an explicit nil
func (o *DAppDTO) UnsetShowInRecommendDApps() {
	o.ShowInRecommendDApps.Unset()
}

// GetTitleLocalized returns the TitleLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetTitleLocalized() LocalizedStringDTO {
	if o == nil || IsNil(o.TitleLocalized.Get()) {
		var ret LocalizedStringDTO
		return ret
	}
	return *o.TitleLocalized.Get()
}

// GetTitleLocalizedOk returns a tuple with the TitleLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetTitleLocalizedOk() (*LocalizedStringDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.TitleLocalized.Get(), o.TitleLocalized.IsSet()
}

// HasTitleLocalized returns a boolean if a field has been set.
func (o *DAppDTO) HasTitleLocalized() bool {
	if o != nil && o.TitleLocalized.IsSet() {
		return true
	}

	return false
}

// SetTitleLocalized gets a reference to the given NullableLocalizedStringDTO and assigns it to the TitleLocalized field.
func (o *DAppDTO) SetTitleLocalized(v LocalizedStringDTO) {
	o.TitleLocalized.Set(&v)
}
// SetTitleLocalizedNil sets the value for TitleLocalized to be an explicit nil
func (o *DAppDTO) SetTitleLocalizedNil() {
	o.TitleLocalized.Set(nil)
}

// UnsetTitleLocalized ensures that no value is present for TitleLocalized, not even an explicit nil
func (o *DAppDTO) UnsetTitleLocalized() {
	o.TitleLocalized.Unset()
}

// GetUrlLocalized returns the UrlLocalized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DAppDTO) GetUrlLocalized() LocalizedStringDTO {
	if o == nil || IsNil(o.UrlLocalized.Get()) {
		var ret LocalizedStringDTO
		return ret
	}
	return *o.UrlLocalized.Get()
}

// GetUrlLocalizedOk returns a tuple with the UrlLocalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DAppDTO) GetUrlLocalizedOk() (*LocalizedStringDTO, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlLocalized.Get(), o.UrlLocalized.IsSet()
}

// HasUrlLocalized returns a boolean if a field has been set.
func (o *DAppDTO) HasUrlLocalized() bool {
	if o != nil && o.UrlLocalized.IsSet() {
		return true
	}

	return false
}

// SetUrlLocalized gets a reference to the given NullableLocalizedStringDTO and assigns it to the UrlLocalized field.
func (o *DAppDTO) SetUrlLocalized(v LocalizedStringDTO) {
	o.UrlLocalized.Set(&v)
}
// SetUrlLocalizedNil sets the value for UrlLocalized to be an explicit nil
func (o *DAppDTO) SetUrlLocalizedNil() {
	o.UrlLocalized.Set(nil)
}

// UnsetUrlLocalized ensures that no value is present for UrlLocalized, not even an explicit nil
func (o *DAppDTO) UnsetUrlLocalized() {
	o.UrlLocalized.Unset()
}

func (o DAppDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DAppDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DescriptionLocalized.IsSet() {
		toSerialize["description_localized"] = o.DescriptionLocalized.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if o.IconUrl.IsSet() {
		toSerialize["icon_url"] = o.IconUrl.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Order.IsSet() {
		toSerialize["order"] = o.Order.Get()
	}
	if o.SearchKeywords != nil {
		toSerialize["search_keywords"] = o.SearchKeywords
	}
	if o.ShowInRecommendDApps.IsSet() {
		toSerialize["show_in_recommend_d_apps"] = o.ShowInRecommendDApps.Get()
	}
	if o.TitleLocalized.IsSet() {
		toSerialize["title_localized"] = o.TitleLocalized.Get()
	}
	if o.UrlLocalized.IsSet() {
		toSerialize["url_localized"] = o.UrlLocalized.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DAppDTO) UnmarshalJSON(data []byte) (err error) {
	varDAppDTO := _DAppDTO{}

	err = json.Unmarshal(data, &varDAppDTO)

	if err != nil {
		return err
	}

	*o = DAppDTO(varDAppDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description_localized")
		delete(additionalProperties, "enabled")
		delete(additionalProperties, "icon_url")
		delete(additionalProperties, "id")
		delete(additionalProperties, "order")
		delete(additionalProperties, "search_keywords")
		delete(additionalProperties, "show_in_recommend_d_apps")
		delete(additionalProperties, "title_localized")
		delete(additionalProperties, "url_localized")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDAppDTO struct {
	value *DAppDTO
	isSet bool
}

func (v NullableDAppDTO) Get() *DAppDTO {
	return v.value
}

func (v *NullableDAppDTO) Set(val *DAppDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableDAppDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableDAppDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDAppDTO(val *DAppDTO) *NullableDAppDTO {
	return &NullableDAppDTO{value: val, isSet: true}
}

func (v NullableDAppDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDAppDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


