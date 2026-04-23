
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalDTO{}

// PalDTO struct for PalDTO
type PalDTO struct {
	AgeInDays NullableInt32 `json:"age_in_days,omitempty"`
	Attributes NullableAttribute `json:"attributes,omitempty"`
	CurrentEmplEarned NullableString `json:"current_empl_earned,omitempty"`
	CurrentLevel NullableInt32 `json:"current_level,omitempty"`
	Description NullableString `json:"description,omitempty"`
	EmplEarnedLimit NullableString `json:"empl_earned_limit,omitempty"`
	FromLiquidityPool NullableBool `json:"from_liquidity_pool,omitempty"`
	HatchingAnimationUrl NullableString `json:"hatching_animation_url,omitempty"`
	Image NullableString `json:"image,omitempty"`
	InLockedParty NullableBool `json:"in_locked_party,omitempty"`
	IsAlive NullableBool `json:"is_alive,omitempty"`
	IsPending NullableBool `json:"is_pending,omitempty"`
	IsValidToEvolve NullableBool `json:"is_valid_to_evolve,omitempty"`
	IsValidToLevelUp NullableBool `json:"is_valid_to_level_up,omitempty"`
	MaxAgeInDays NullableInt32 `json:"max_age_in_days,omitempty"`
	MaxAttributes NullableMaxAttribute `json:"max_attributes,omitempty"`
	Name NullableString `json:"name,omitempty"`
	OriginalWalletAddress NullableString `json:"original_wallet_address,omitempty"`
	RaceStatus NullableString `json:"race_status,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalDTO PalDTO

// NewPalDTO instantiates a new PalDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalDTO() *PalDTO {
	this := PalDTO{}
	return &this
}

// NewPalDTOWithDefaults instantiates a new PalDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalDTOWithDefaults() *PalDTO {
	this := PalDTO{}
	return &this
}

// GetAgeInDays returns the AgeInDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetAgeInDays() int32 {
	if o == nil || IsNil(o.AgeInDays.Get()) {
		var ret int32
		return ret
	}
	return *o.AgeInDays.Get()
}

// GetAgeInDaysOk returns a tuple with the AgeInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetAgeInDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AgeInDays.Get(), o.AgeInDays.IsSet()
}

// HasAgeInDays returns a boolean if a field has been set.
func (o *PalDTO) HasAgeInDays() bool {
	if o != nil && o.AgeInDays.IsSet() {
		return true
	}

	return false
}

// SetAgeInDays gets a reference to the given NullableInt32 and assigns it to the AgeInDays field.
func (o *PalDTO) SetAgeInDays(v int32) {
	o.AgeInDays.Set(&v)
}
// SetAgeInDaysNil sets the value for AgeInDays to be an explicit nil
func (o *PalDTO) SetAgeInDaysNil() {
	o.AgeInDays.Set(nil)
}

// UnsetAgeInDays ensures that no value is present for AgeInDays, not even an explicit nil
func (o *PalDTO) UnsetAgeInDays() {
	o.AgeInDays.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetAttributes() Attribute {
	if o == nil || IsNil(o.Attributes.Get()) {
		var ret Attribute
		return ret
	}
	return *o.Attributes.Get()
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetAttributesOk() (*Attribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes.Get(), o.Attributes.IsSet()
}

// HasAttributes returns a boolean if a field has been set.
func (o *PalDTO) HasAttributes() bool {
	if o != nil && o.Attributes.IsSet() {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given NullableAttribute and assigns it to the Attributes field.
func (o *PalDTO) SetAttributes(v Attribute) {
	o.Attributes.Set(&v)
}
// SetAttributesNil sets the value for Attributes to be an explicit nil
func (o *PalDTO) SetAttributesNil() {
	o.Attributes.Set(nil)
}

// UnsetAttributes ensures that no value is present for Attributes, not even an explicit nil
func (o *PalDTO) UnsetAttributes() {
	o.Attributes.Unset()
}

// GetCurrentEmplEarned returns the CurrentEmplEarned field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetCurrentEmplEarned() string {
	if o == nil || IsNil(o.CurrentEmplEarned.Get()) {
		var ret string
		return ret
	}
	return *o.CurrentEmplEarned.Get()
}

// GetCurrentEmplEarnedOk returns a tuple with the CurrentEmplEarned field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetCurrentEmplEarnedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentEmplEarned.Get(), o.CurrentEmplEarned.IsSet()
}

// HasCurrentEmplEarned returns a boolean if a field has been set.
func (o *PalDTO) HasCurrentEmplEarned() bool {
	if o != nil && o.CurrentEmplEarned.IsSet() {
		return true
	}

	return false
}

// SetCurrentEmplEarned gets a reference to the given NullableString and assigns it to the CurrentEmplEarned field.
func (o *PalDTO) SetCurrentEmplEarned(v string) {
	o.CurrentEmplEarned.Set(&v)
}
// SetCurrentEmplEarnedNil sets the value for CurrentEmplEarned to be an explicit nil
func (o *PalDTO) SetCurrentEmplEarnedNil() {
	o.CurrentEmplEarned.Set(nil)
}

// UnsetCurrentEmplEarned ensures that no value is present for CurrentEmplEarned, not even an explicit nil
func (o *PalDTO) UnsetCurrentEmplEarned() {
	o.CurrentEmplEarned.Unset()
}

// GetCurrentLevel returns the CurrentLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetCurrentLevel() int32 {
	if o == nil || IsNil(o.CurrentLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.CurrentLevel.Get()
}

// GetCurrentLevelOk returns a tuple with the CurrentLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetCurrentLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentLevel.Get(), o.CurrentLevel.IsSet()
}

// HasCurrentLevel returns a boolean if a field has been set.
func (o *PalDTO) HasCurrentLevel() bool {
	if o != nil && o.CurrentLevel.IsSet() {
		return true
	}

	return false
}

// SetCurrentLevel gets a reference to the given NullableInt32 and assigns it to the CurrentLevel field.
func (o *PalDTO) SetCurrentLevel(v int32) {
	o.CurrentLevel.Set(&v)
}
// SetCurrentLevelNil sets the value for CurrentLevel to be an explicit nil
func (o *PalDTO) SetCurrentLevelNil() {
	o.CurrentLevel.Set(nil)
}

// UnsetCurrentLevel ensures that no value is present for CurrentLevel, not even an explicit nil
func (o *PalDTO) UnsetCurrentLevel() {
	o.CurrentLevel.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetDescription() string {
	if o == nil || IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *PalDTO) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *PalDTO) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *PalDTO) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *PalDTO) UnsetDescription() {
	o.Description.Unset()
}

// GetEmplEarnedLimit returns the EmplEarnedLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetEmplEarnedLimit() string {
	if o == nil || IsNil(o.EmplEarnedLimit.Get()) {
		var ret string
		return ret
	}
	return *o.EmplEarnedLimit.Get()
}

// GetEmplEarnedLimitOk returns a tuple with the EmplEarnedLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetEmplEarnedLimitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmplEarnedLimit.Get(), o.EmplEarnedLimit.IsSet()
}

// HasEmplEarnedLimit returns a boolean if a field has been set.
func (o *PalDTO) HasEmplEarnedLimit() bool {
	if o != nil && o.EmplEarnedLimit.IsSet() {
		return true
	}

	return false
}

// SetEmplEarnedLimit gets a reference to the given NullableString and assigns it to the EmplEarnedLimit field.
func (o *PalDTO) SetEmplEarnedLimit(v string) {
	o.EmplEarnedLimit.Set(&v)
}
// SetEmplEarnedLimitNil sets the value for EmplEarnedLimit to be an explicit nil
func (o *PalDTO) SetEmplEarnedLimitNil() {
	o.EmplEarnedLimit.Set(nil)
}

// UnsetEmplEarnedLimit ensures that no value is present for EmplEarnedLimit, not even an explicit nil
func (o *PalDTO) UnsetEmplEarnedLimit() {
	o.EmplEarnedLimit.Unset()
}

// GetFromLiquidityPool returns the FromLiquidityPool field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetFromLiquidityPool() bool {
	if o == nil || IsNil(o.FromLiquidityPool.Get()) {
		var ret bool
		return ret
	}
	return *o.FromLiquidityPool.Get()
}

// GetFromLiquidityPoolOk returns a tuple with the FromLiquidityPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetFromLiquidityPoolOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromLiquidityPool.Get(), o.FromLiquidityPool.IsSet()
}

// HasFromLiquidityPool returns a boolean if a field has been set.
func (o *PalDTO) HasFromLiquidityPool() bool {
	if o != nil && o.FromLiquidityPool.IsSet() {
		return true
	}

	return false
}

// SetFromLiquidityPool gets a reference to the given NullableBool and assigns it to the FromLiquidityPool field.
func (o *PalDTO) SetFromLiquidityPool(v bool) {
	o.FromLiquidityPool.Set(&v)
}
// SetFromLiquidityPoolNil sets the value for FromLiquidityPool to be an explicit nil
func (o *PalDTO) SetFromLiquidityPoolNil() {
	o.FromLiquidityPool.Set(nil)
}

// UnsetFromLiquidityPool ensures that no value is present for FromLiquidityPool, not even an explicit nil
func (o *PalDTO) UnsetFromLiquidityPool() {
	o.FromLiquidityPool.Unset()
}

// GetHatchingAnimationUrl returns the HatchingAnimationUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetHatchingAnimationUrl() string {
	if o == nil || IsNil(o.HatchingAnimationUrl.Get()) {
		var ret string
		return ret
	}
	return *o.HatchingAnimationUrl.Get()
}

// GetHatchingAnimationUrlOk returns a tuple with the HatchingAnimationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetHatchingAnimationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HatchingAnimationUrl.Get(), o.HatchingAnimationUrl.IsSet()
}

// HasHatchingAnimationUrl returns a boolean if a field has been set.
func (o *PalDTO) HasHatchingAnimationUrl() bool {
	if o != nil && o.HatchingAnimationUrl.IsSet() {
		return true
	}

	return false
}

// SetHatchingAnimationUrl gets a reference to the given NullableString and assigns it to the HatchingAnimationUrl field.
func (o *PalDTO) SetHatchingAnimationUrl(v string) {
	o.HatchingAnimationUrl.Set(&v)
}
// SetHatchingAnimationUrlNil sets the value for HatchingAnimationUrl to be an explicit nil
func (o *PalDTO) SetHatchingAnimationUrlNil() {
	o.HatchingAnimationUrl.Set(nil)
}

// UnsetHatchingAnimationUrl ensures that no value is present for HatchingAnimationUrl, not even an explicit nil
func (o *PalDTO) UnsetHatchingAnimationUrl() {
	o.HatchingAnimationUrl.Unset()
}

// GetImage returns the Image field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetImage() string {
	if o == nil || IsNil(o.Image.Get()) {
		var ret string
		return ret
	}
	return *o.Image.Get()
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Image.Get(), o.Image.IsSet()
}

// HasImage returns a boolean if a field has been set.
func (o *PalDTO) HasImage() bool {
	if o != nil && o.Image.IsSet() {
		return true
	}

	return false
}

// SetImage gets a reference to the given NullableString and assigns it to the Image field.
func (o *PalDTO) SetImage(v string) {
	o.Image.Set(&v)
}
// SetImageNil sets the value for Image to be an explicit nil
func (o *PalDTO) SetImageNil() {
	o.Image.Set(nil)
}

// UnsetImage ensures that no value is present for Image, not even an explicit nil
func (o *PalDTO) UnsetImage() {
	o.Image.Unset()
}

// GetInLockedParty returns the InLockedParty field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetInLockedParty() bool {
	if o == nil || IsNil(o.InLockedParty.Get()) {
		var ret bool
		return ret
	}
	return *o.InLockedParty.Get()
}

// GetInLockedPartyOk returns a tuple with the InLockedParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetInLockedPartyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.InLockedParty.Get(), o.InLockedParty.IsSet()
}

// HasInLockedParty returns a boolean if a field has been set.
func (o *PalDTO) HasInLockedParty() bool {
	if o != nil && o.InLockedParty.IsSet() {
		return true
	}

	return false
}

// SetInLockedParty gets a reference to the given NullableBool and assigns it to the InLockedParty field.
func (o *PalDTO) SetInLockedParty(v bool) {
	o.InLockedParty.Set(&v)
}
// SetInLockedPartyNil sets the value for InLockedParty to be an explicit nil
func (o *PalDTO) SetInLockedPartyNil() {
	o.InLockedParty.Set(nil)
}

// UnsetInLockedParty ensures that no value is present for InLockedParty, not even an explicit nil
func (o *PalDTO) UnsetInLockedParty() {
	o.InLockedParty.Unset()
}

// GetIsAlive returns the IsAlive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetIsAlive() bool {
	if o == nil || IsNil(o.IsAlive.Get()) {
		var ret bool
		return ret
	}
	return *o.IsAlive.Get()
}

// GetIsAliveOk returns a tuple with the IsAlive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetIsAliveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsAlive.Get(), o.IsAlive.IsSet()
}

// HasIsAlive returns a boolean if a field has been set.
func (o *PalDTO) HasIsAlive() bool {
	if o != nil && o.IsAlive.IsSet() {
		return true
	}

	return false
}

// SetIsAlive gets a reference to the given NullableBool and assigns it to the IsAlive field.
func (o *PalDTO) SetIsAlive(v bool) {
	o.IsAlive.Set(&v)
}
// SetIsAliveNil sets the value for IsAlive to be an explicit nil
func (o *PalDTO) SetIsAliveNil() {
	o.IsAlive.Set(nil)
}

// UnsetIsAlive ensures that no value is present for IsAlive, not even an explicit nil
func (o *PalDTO) UnsetIsAlive() {
	o.IsAlive.Unset()
}

// GetIsPending returns the IsPending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetIsPending() bool {
	if o == nil || IsNil(o.IsPending.Get()) {
		var ret bool
		return ret
	}
	return *o.IsPending.Get()
}

// GetIsPendingOk returns a tuple with the IsPending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetIsPendingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPending.Get(), o.IsPending.IsSet()
}

// HasIsPending returns a boolean if a field has been set.
func (o *PalDTO) HasIsPending() bool {
	if o != nil && o.IsPending.IsSet() {
		return true
	}

	return false
}

// SetIsPending gets a reference to the given NullableBool and assigns it to the IsPending field.
func (o *PalDTO) SetIsPending(v bool) {
	o.IsPending.Set(&v)
}
// SetIsPendingNil sets the value for IsPending to be an explicit nil
func (o *PalDTO) SetIsPendingNil() {
	o.IsPending.Set(nil)
}

// UnsetIsPending ensures that no value is present for IsPending, not even an explicit nil
func (o *PalDTO) UnsetIsPending() {
	o.IsPending.Unset()
}

// GetIsValidToEvolve returns the IsValidToEvolve field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetIsValidToEvolve() bool {
	if o == nil || IsNil(o.IsValidToEvolve.Get()) {
		var ret bool
		return ret
	}
	return *o.IsValidToEvolve.Get()
}

// GetIsValidToEvolveOk returns a tuple with the IsValidToEvolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetIsValidToEvolveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsValidToEvolve.Get(), o.IsValidToEvolve.IsSet()
}

// HasIsValidToEvolve returns a boolean if a field has been set.
func (o *PalDTO) HasIsValidToEvolve() bool {
	if o != nil && o.IsValidToEvolve.IsSet() {
		return true
	}

	return false
}

// SetIsValidToEvolve gets a reference to the given NullableBool and assigns it to the IsValidToEvolve field.
func (o *PalDTO) SetIsValidToEvolve(v bool) {
	o.IsValidToEvolve.Set(&v)
}
// SetIsValidToEvolveNil sets the value for IsValidToEvolve to be an explicit nil
func (o *PalDTO) SetIsValidToEvolveNil() {
	o.IsValidToEvolve.Set(nil)
}

// UnsetIsValidToEvolve ensures that no value is present for IsValidToEvolve, not even an explicit nil
func (o *PalDTO) UnsetIsValidToEvolve() {
	o.IsValidToEvolve.Unset()
}

// GetIsValidToLevelUp returns the IsValidToLevelUp field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetIsValidToLevelUp() bool {
	if o == nil || IsNil(o.IsValidToLevelUp.Get()) {
		var ret bool
		return ret
	}
	return *o.IsValidToLevelUp.Get()
}

// GetIsValidToLevelUpOk returns a tuple with the IsValidToLevelUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetIsValidToLevelUpOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsValidToLevelUp.Get(), o.IsValidToLevelUp.IsSet()
}

// HasIsValidToLevelUp returns a boolean if a field has been set.
func (o *PalDTO) HasIsValidToLevelUp() bool {
	if o != nil && o.IsValidToLevelUp.IsSet() {
		return true
	}

	return false
}

// SetIsValidToLevelUp gets a reference to the given NullableBool and assigns it to the IsValidToLevelUp field.
func (o *PalDTO) SetIsValidToLevelUp(v bool) {
	o.IsValidToLevelUp.Set(&v)
}
// SetIsValidToLevelUpNil sets the value for IsValidToLevelUp to be an explicit nil
func (o *PalDTO) SetIsValidToLevelUpNil() {
	o.IsValidToLevelUp.Set(nil)
}

// UnsetIsValidToLevelUp ensures that no value is present for IsValidToLevelUp, not even an explicit nil
func (o *PalDTO) UnsetIsValidToLevelUp() {
	o.IsValidToLevelUp.Unset()
}

// GetMaxAgeInDays returns the MaxAgeInDays field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetMaxAgeInDays() int32 {
	if o == nil || IsNil(o.MaxAgeInDays.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAgeInDays.Get()
}

// GetMaxAgeInDaysOk returns a tuple with the MaxAgeInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetMaxAgeInDaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAgeInDays.Get(), o.MaxAgeInDays.IsSet()
}

// HasMaxAgeInDays returns a boolean if a field has been set.
func (o *PalDTO) HasMaxAgeInDays() bool {
	if o != nil && o.MaxAgeInDays.IsSet() {
		return true
	}

	return false
}

// SetMaxAgeInDays gets a reference to the given NullableInt32 and assigns it to the MaxAgeInDays field.
func (o *PalDTO) SetMaxAgeInDays(v int32) {
	o.MaxAgeInDays.Set(&v)
}
// SetMaxAgeInDaysNil sets the value for MaxAgeInDays to be an explicit nil
func (o *PalDTO) SetMaxAgeInDaysNil() {
	o.MaxAgeInDays.Set(nil)
}

// UnsetMaxAgeInDays ensures that no value is present for MaxAgeInDays, not even an explicit nil
func (o *PalDTO) UnsetMaxAgeInDays() {
	o.MaxAgeInDays.Unset()
}

// GetMaxAttributes returns the MaxAttributes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetMaxAttributes() MaxAttribute {
	if o == nil || IsNil(o.MaxAttributes.Get()) {
		var ret MaxAttribute
		return ret
	}
	return *o.MaxAttributes.Get()
}

// GetMaxAttributesOk returns a tuple with the MaxAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetMaxAttributesOk() (*MaxAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAttributes.Get(), o.MaxAttributes.IsSet()
}

// HasMaxAttributes returns a boolean if a field has been set.
func (o *PalDTO) HasMaxAttributes() bool {
	if o != nil && o.MaxAttributes.IsSet() {
		return true
	}

	return false
}

// SetMaxAttributes gets a reference to the given NullableMaxAttribute and assigns it to the MaxAttributes field.
func (o *PalDTO) SetMaxAttributes(v MaxAttribute) {
	o.MaxAttributes.Set(&v)
}
// SetMaxAttributesNil sets the value for MaxAttributes to be an explicit nil
func (o *PalDTO) SetMaxAttributesNil() {
	o.MaxAttributes.Set(nil)
}

// UnsetMaxAttributes ensures that no value is present for MaxAttributes, not even an explicit nil
func (o *PalDTO) UnsetMaxAttributes() {
	o.MaxAttributes.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *PalDTO) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *PalDTO) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *PalDTO) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *PalDTO) UnsetName() {
	o.Name.Unset()
}

// GetOriginalWalletAddress returns the OriginalWalletAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetOriginalWalletAddress() string {
	if o == nil || IsNil(o.OriginalWalletAddress.Get()) {
		var ret string
		return ret
	}
	return *o.OriginalWalletAddress.Get()
}

// GetOriginalWalletAddressOk returns a tuple with the OriginalWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetOriginalWalletAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OriginalWalletAddress.Get(), o.OriginalWalletAddress.IsSet()
}

// HasOriginalWalletAddress returns a boolean if a field has been set.
func (o *PalDTO) HasOriginalWalletAddress() bool {
	if o != nil && o.OriginalWalletAddress.IsSet() {
		return true
	}

	return false
}

// SetOriginalWalletAddress gets a reference to the given NullableString and assigns it to the OriginalWalletAddress field.
func (o *PalDTO) SetOriginalWalletAddress(v string) {
	o.OriginalWalletAddress.Set(&v)
}
// SetOriginalWalletAddressNil sets the value for OriginalWalletAddress to be an explicit nil
func (o *PalDTO) SetOriginalWalletAddressNil() {
	o.OriginalWalletAddress.Set(nil)
}

// UnsetOriginalWalletAddress ensures that no value is present for OriginalWalletAddress, not even an explicit nil
func (o *PalDTO) UnsetOriginalWalletAddress() {
	o.OriginalWalletAddress.Unset()
}

// GetRaceStatus returns the RaceStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetRaceStatus() string {
	if o == nil || IsNil(o.RaceStatus.Get()) {
		var ret string
		return ret
	}
	return *o.RaceStatus.Get()
}

// GetRaceStatusOk returns a tuple with the RaceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetRaceStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceStatus.Get(), o.RaceStatus.IsSet()
}

// HasRaceStatus returns a boolean if a field has been set.
func (o *PalDTO) HasRaceStatus() bool {
	if o != nil && o.RaceStatus.IsSet() {
		return true
	}

	return false
}

// SetRaceStatus gets a reference to the given NullableString and assigns it to the RaceStatus field.
func (o *PalDTO) SetRaceStatus(v string) {
	o.RaceStatus.Set(&v)
}
// SetRaceStatusNil sets the value for RaceStatus to be an explicit nil
func (o *PalDTO) SetRaceStatusNil() {
	o.RaceStatus.Set(nil)
}

// UnsetRaceStatus ensures that no value is present for RaceStatus, not even an explicit nil
func (o *PalDTO) UnsetRaceStatus() {
	o.RaceStatus.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalDTO) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalDTO) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *PalDTO) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *PalDTO) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *PalDTO) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *PalDTO) UnsetTokenId() {
	o.TokenId.Unset()
}

func (o PalDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AgeInDays.IsSet() {
		toSerialize["age_in_days"] = o.AgeInDays.Get()
	}
	if o.Attributes.IsSet() {
		toSerialize["attributes"] = o.Attributes.Get()
	}
	if o.CurrentEmplEarned.IsSet() {
		toSerialize["current_empl_earned"] = o.CurrentEmplEarned.Get()
	}
	if o.CurrentLevel.IsSet() {
		toSerialize["current_level"] = o.CurrentLevel.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.EmplEarnedLimit.IsSet() {
		toSerialize["empl_earned_limit"] = o.EmplEarnedLimit.Get()
	}
	if o.FromLiquidityPool.IsSet() {
		toSerialize["from_liquidity_pool"] = o.FromLiquidityPool.Get()
	}
	if o.HatchingAnimationUrl.IsSet() {
		toSerialize["hatching_animation_url"] = o.HatchingAnimationUrl.Get()
	}
	if o.Image.IsSet() {
		toSerialize["image"] = o.Image.Get()
	}
	if o.InLockedParty.IsSet() {
		toSerialize["in_locked_party"] = o.InLockedParty.Get()
	}
	if o.IsAlive.IsSet() {
		toSerialize["is_alive"] = o.IsAlive.Get()
	}
	if o.IsPending.IsSet() {
		toSerialize["is_pending"] = o.IsPending.Get()
	}
	if o.IsValidToEvolve.IsSet() {
		toSerialize["is_valid_to_evolve"] = o.IsValidToEvolve.Get()
	}
	if o.IsValidToLevelUp.IsSet() {
		toSerialize["is_valid_to_level_up"] = o.IsValidToLevelUp.Get()
	}
	if o.MaxAgeInDays.IsSet() {
		toSerialize["max_age_in_days"] = o.MaxAgeInDays.Get()
	}
	if o.MaxAttributes.IsSet() {
		toSerialize["max_attributes"] = o.MaxAttributes.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.OriginalWalletAddress.IsSet() {
		toSerialize["original_wallet_address"] = o.OriginalWalletAddress.Get()
	}
	if o.RaceStatus.IsSet() {
		toSerialize["race_status"] = o.RaceStatus.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalDTO) UnmarshalJSON(data []byte) (err error) {
	varPalDTO := _PalDTO{}

	err = json.Unmarshal(data, &varPalDTO)

	if err != nil {
		return err
	}

	*o = PalDTO(varPalDTO)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "age_in_days")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "current_empl_earned")
		delete(additionalProperties, "current_level")
		delete(additionalProperties, "description")
		delete(additionalProperties, "empl_earned_limit")
		delete(additionalProperties, "from_liquidity_pool")
		delete(additionalProperties, "hatching_animation_url")
		delete(additionalProperties, "image")
		delete(additionalProperties, "in_locked_party")
		delete(additionalProperties, "is_alive")
		delete(additionalProperties, "is_pending")
		delete(additionalProperties, "is_valid_to_evolve")
		delete(additionalProperties, "is_valid_to_level_up")
		delete(additionalProperties, "max_age_in_days")
		delete(additionalProperties, "max_attributes")
		delete(additionalProperties, "name")
		delete(additionalProperties, "original_wallet_address")
		delete(additionalProperties, "race_status")
		delete(additionalProperties, "token_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalDTO struct {
	value *PalDTO
	isSet bool
}

func (v NullablePalDTO) Get() *PalDTO {
	return v.value
}

func (v *NullablePalDTO) Set(val *PalDTO) {
	v.value = val
	v.isSet = true
}

func (v NullablePalDTO) IsSet() bool {
	return v.isSet
}

func (v *NullablePalDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalDTO(val *PalDTO) *NullablePalDTO {
	return &NullablePalDTO{value: val, isSet: true}
}

func (v NullablePalDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


