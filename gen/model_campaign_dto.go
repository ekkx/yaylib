
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the CampaignDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignDTO{}

// CampaignDTO struct for CampaignDTO
type CampaignDTO struct {
	BannerImage NullableString `json:"banner_image,omitempty"`
	CoverImage NullableString `json:"cover_image,omitempty"`
	Description NullableString `json:"description,omitempty"`
	DescriptionUrl NullableString `json:"description_url,omitempty"`
	EndAt NullableInt64 `json:"end_at,omitempty"`
	Id NullableInt64 `json:"id,omitempty"`
	MultiplierDetailUrl NullableString `json:"multiplier_detail_url,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Ranking NullableInt32 `json:"ranking,omitempty"`
	StartAt NullableInt64 `json:"start_at,omitempty"`
	TotalPoints NullableInt32 `json:"total_points,omitempty"`
	UserCampaign NullableUserCampaign `json:"user_campaign,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CampaignDTO CampaignDTO

// NewCampaignDTO instantiates a new CampaignDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignDTO() *CampaignDTO {
	this := CampaignDTO{}
	return &this
}

// NewCampaignDTOWithDefaults instantiates a new CampaignDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignDTOWithDefaults() *CampaignDTO {
	this := CampaignDTO{}
	return &this
}

// GetBannerImage returns the BannerImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetBannerImage() string {
	if o == nil || IsNil(o.BannerImage.Get()) {
		var ret string
		return ret
	}
	return *o.BannerImage.Get()
}

// GetBannerImageOk returns a tuple with the BannerImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetBannerImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BannerImage.Get(), o.BannerImage.IsSet()
}

// HasBannerImage returns a boolean if a field has been set.
func (o *CampaignDTO) HasBannerImage() bool {
	if o != nil && o.BannerImage.IsSet() {
		return true
	}

	return false
}

// SetBannerImage gets a reference to the given NullableString and assigns it to the BannerImage field.
func (o *CampaignDTO) SetBannerImage(v string) {
	o.BannerImage.Set(&v)
}
// SetBannerImageNil sets the value for BannerImage to be an explicit nil
func (o *CampaignDTO) SetBannerImageNil() {
	o.BannerImage.Set(nil)
}

// UnsetBannerImage ensures that no value is present for BannerImage, not even an explicit nil
func (o *CampaignDTO) UnsetBannerImage() {
	o.BannerImage.Unset()
}

// GetCoverImage returns the CoverImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetCoverImage() string {
	if o == nil || IsNil(o.CoverImage.Get()) {
		var ret string
		return ret
	}
	return *o.CoverImage.Get()
}

// GetCoverImageOk returns a tuple with the CoverImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetCoverImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CoverImage.Get(), o.CoverImage.IsSet()
}

// HasCoverImage returns a boolean if a field has been set.
func (o *CampaignDTO) HasCoverImage() bool {
	if o != nil && o.CoverImage.IsSet() {
		return true
	}

	return false
}

// SetCoverImage gets a reference to the given NullableString and assigns it to the CoverImage field.
func (o *CampaignDTO) SetCoverImage(v string) {
	o.CoverImage.Set(&v)
}
// SetCoverImageNil sets the value for CoverImage to be an explicit nil
func (o *CampaignDTO) SetCoverImageNil() {
	o.CoverImage.Set(nil)
}

// UnsetCoverImage ensures that no value is present for CoverImage, not even an explicit nil
func (o *CampaignDTO) UnsetCoverImage() {
	o.CoverImage.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *CampaignDTO) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *CampaignDTO) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *CampaignDTO) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *CampaignDTO) UnsetDescription() {
	o.Description.Unset()
}

// GetDescriptionUrl returns the DescriptionUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetDescriptionUrl() string {
	if o == nil || IsNil(o.DescriptionUrl.Get()) {
		var ret string
		return ret
	}
	return *o.DescriptionUrl.Get()
}

// GetDescriptionUrlOk returns a tuple with the DescriptionUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetDescriptionUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescriptionUrl.Get(), o.DescriptionUrl.IsSet()
}

// HasDescriptionUrl returns a boolean if a field has been set.
func (o *CampaignDTO) HasDescriptionUrl() bool {
	if o != nil && o.DescriptionUrl.IsSet() {
		return true
	}

	return false
}

// SetDescriptionUrl gets a reference to the given NullableString and assigns it to the DescriptionUrl field.
func (o *CampaignDTO) SetDescriptionUrl(v string) {
	o.DescriptionUrl.Set(&v)
}
// SetDescriptionUrlNil sets the value for DescriptionUrl to be an explicit nil
func (o *CampaignDTO) SetDescriptionUrlNil() {
	o.DescriptionUrl.Set(nil)
}

// UnsetDescriptionUrl ensures that no value is present for DescriptionUrl, not even an explicit nil
func (o *CampaignDTO) UnsetDescriptionUrl() {
	o.DescriptionUrl.Unset()
}

// GetEndAt returns the EndAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetEndAt() int64 {
	if o == nil || IsNil(o.EndAt.Get()) {
		var ret int64
		return ret
	}
	return *o.EndAt.Get()
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetEndAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndAt.Get(), o.EndAt.IsSet()
}

// HasEndAt returns a boolean if a field has been set.
func (o *CampaignDTO) HasEndAt() bool {
	if o != nil && o.EndAt.IsSet() {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given NullableInt64 and assigns it to the EndAt field.
func (o *CampaignDTO) SetEndAt(v int64) {
	o.EndAt.Set(&v)
}
// SetEndAtNil sets the value for EndAt to be an explicit nil
func (o *CampaignDTO) SetEndAtNil() {
	o.EndAt.Set(nil)
}

// UnsetEndAt ensures that no value is present for EndAt, not even an explicit nil
func (o *CampaignDTO) UnsetEndAt() {
	o.EndAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetId() int64 {
	if o == nil || IsNil(o.Id.Get()) {
		var ret int64
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *CampaignDTO) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableInt64 and assigns it to the Id field.
func (o *CampaignDTO) SetId(v int64) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *CampaignDTO) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *CampaignDTO) UnsetId() {
	o.Id.Unset()
}

// GetMultiplierDetailUrl returns the MultiplierDetailUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetMultiplierDetailUrl() string {
	if o == nil || IsNil(o.MultiplierDetailUrl.Get()) {
		var ret string
		return ret
	}
	return *o.MultiplierDetailUrl.Get()
}

// GetMultiplierDetailUrlOk returns a tuple with the MultiplierDetailUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetMultiplierDetailUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MultiplierDetailUrl.Get(), o.MultiplierDetailUrl.IsSet()
}

// HasMultiplierDetailUrl returns a boolean if a field has been set.
func (o *CampaignDTO) HasMultiplierDetailUrl() bool {
	if o != nil && o.MultiplierDetailUrl.IsSet() {
		return true
	}

	return false
}

// SetMultiplierDetailUrl gets a reference to the given NullableString and assigns it to the MultiplierDetailUrl field.
func (o *CampaignDTO) SetMultiplierDetailUrl(v string) {
	o.MultiplierDetailUrl.Set(&v)
}
// SetMultiplierDetailUrlNil sets the value for MultiplierDetailUrl to be an explicit nil
func (o *CampaignDTO) SetMultiplierDetailUrlNil() {
	o.MultiplierDetailUrl.Set(nil)
}

// UnsetMultiplierDetailUrl ensures that no value is present for MultiplierDetailUrl, not even an explicit nil
func (o *CampaignDTO) UnsetMultiplierDetailUrl() {
	o.MultiplierDetailUrl.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *CampaignDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *CampaignDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *CampaignDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *CampaignDTO) UnsetName() {
	o.Name.Unset()
}

// GetRanking returns the Ranking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetRanking() int32 {
	if o == nil || IsNil(o.Ranking.Get()) {
		var ret int32
		return ret
	}
	return *o.Ranking.Get()
}

// GetRankingOk returns a tuple with the Ranking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetRankingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ranking.Get(), o.Ranking.IsSet()
}

// HasRanking returns a boolean if a field has been set.
func (o *CampaignDTO) HasRanking() bool {
	if o != nil && o.Ranking.IsSet() {
		return true
	}

	return false
}

// SetRanking gets a reference to the given NullableInt32 and assigns it to the Ranking field.
func (o *CampaignDTO) SetRanking(v int32) {
	o.Ranking.Set(&v)
}
// SetRankingNil sets the value for Ranking to be an explicit nil
func (o *CampaignDTO) SetRankingNil() {
	o.Ranking.Set(nil)
}

// UnsetRanking ensures that no value is present for Ranking, not even an explicit nil
func (o *CampaignDTO) UnsetRanking() {
	o.Ranking.Unset()
}

// GetStartAt returns the StartAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetStartAt() int64 {
	if o == nil || IsNil(o.StartAt.Get()) {
		var ret int64
		return ret
	}
	return *o.StartAt.Get()
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetStartAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartAt.Get(), o.StartAt.IsSet()
}

// HasStartAt returns a boolean if a field has been set.
func (o *CampaignDTO) HasStartAt() bool {
	if o != nil && o.StartAt.IsSet() {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given NullableInt64 and assigns it to the StartAt field.
func (o *CampaignDTO) SetStartAt(v int64) {
	o.StartAt.Set(&v)
}
// SetStartAtNil sets the value for StartAt to be an explicit nil
func (o *CampaignDTO) SetStartAtNil() {
	o.StartAt.Set(nil)
}

// UnsetStartAt ensures that no value is present for StartAt, not even an explicit nil
func (o *CampaignDTO) UnsetStartAt() {
	o.StartAt.Unset()
}

// GetTotalPoints returns the TotalPoints field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetTotalPoints() int32 {
	if o == nil || IsNil(o.TotalPoints.Get()) {
		var ret int32
		return ret
	}
	return *o.TotalPoints.Get()
}

// GetTotalPointsOk returns a tuple with the TotalPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetTotalPointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TotalPoints.Get(), o.TotalPoints.IsSet()
}

// HasTotalPoints returns a boolean if a field has been set.
func (o *CampaignDTO) HasTotalPoints() bool {
	if o != nil && o.TotalPoints.IsSet() {
		return true
	}

	return false
}

// SetTotalPoints gets a reference to the given NullableInt32 and assigns it to the TotalPoints field.
func (o *CampaignDTO) SetTotalPoints(v int32) {
	o.TotalPoints.Set(&v)
}
// SetTotalPointsNil sets the value for TotalPoints to be an explicit nil
func (o *CampaignDTO) SetTotalPointsNil() {
	o.TotalPoints.Set(nil)
}

// UnsetTotalPoints ensures that no value is present for TotalPoints, not even an explicit nil
func (o *CampaignDTO) UnsetTotalPoints() {
	o.TotalPoints.Unset()
}

// GetUserCampaign returns the UserCampaign field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CampaignDTO) GetUserCampaign() UserCampaign {
	if o == nil || IsNil(o.UserCampaign.Get()) {
		var ret UserCampaign
		return ret
	}
	return *o.UserCampaign.Get()
}

// GetUserCampaignOk returns a tuple with the UserCampaign field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CampaignDTO) GetUserCampaignOk() (*UserCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserCampaign.Get(), o.UserCampaign.IsSet()
}

// HasUserCampaign returns a boolean if a field has been set.
func (o *CampaignDTO) HasUserCampaign() bool {
	if o != nil && o.UserCampaign.IsSet() {
		return true
	}

	return false
}

// SetUserCampaign gets a reference to the given NullableUserCampaign and assigns it to the UserCampaign field.
func (o *CampaignDTO) SetUserCampaign(v UserCampaign) {
	o.UserCampaign.Set(&v)
}
// SetUserCampaignNil sets the value for UserCampaign to be an explicit nil
func (o *CampaignDTO) SetUserCampaignNil() {
	o.UserCampaign.Set(nil)
}

// UnsetUserCampaign ensures that no value is present for UserCampaign, not even an explicit nil
func (o *CampaignDTO) UnsetUserCampaign() {
	o.UserCampaign.Unset()
}

func (o CampaignDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CampaignDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BannerImage.IsSet() {
		toSerialize["banner_image"] = o.BannerImage.Get()
	}
	if o.CoverImage.IsSet() {
		toSerialize["cover_image"] = o.CoverImage.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DescriptionUrl.IsSet() {
		toSerialize["description_url"] = o.DescriptionUrl.Get()
	}
	if o.EndAt.IsSet() {
		toSerialize["end_at"] = o.EndAt.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.MultiplierDetailUrl.IsSet() {
		toSerialize["multiplier_detail_url"] = o.MultiplierDetailUrl.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Ranking.IsSet() {
		toSerialize["ranking"] = o.Ranking.Get()
	}
	if o.StartAt.IsSet() {
		toSerialize["start_at"] = o.StartAt.Get()
	}
	if o.TotalPoints.IsSet() {
		toSerialize["total_points"] = o.TotalPoints.Get()
	}
	if o.UserCampaign.IsSet() {
		toSerialize["user_campaign"] = o.UserCampaign.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CampaignDTO) UnmarshalJSON(data []byte) (err error) {
	varCampaignDTO := _CampaignDTO{}

	err = json.Unmarshal(data, &varCampaignDTO)

	if err != nil {
		return err
	}

	*o = CampaignDTO(varCampaignDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "banner_image")
		delete(additionalProperties, "cover_image")
		delete(additionalProperties, "description")
		delete(additionalProperties, "description_url")
		delete(additionalProperties, "end_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "multiplier_detail_url")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ranking")
		delete(additionalProperties, "start_at")
		delete(additionalProperties, "total_points")
		delete(additionalProperties, "user_campaign")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCampaignDTO struct {
	value *CampaignDTO
	isSet bool
}

func (v NullableCampaignDTO) Get() *CampaignDTO {
	return v.value
}

func (v *NullableCampaignDTO) Set(val *CampaignDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignDTO(val *CampaignDTO) *NullableCampaignDTO {
	return &NullableCampaignDTO{value: val, isSet: true}
}

func (v NullableCampaignDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


