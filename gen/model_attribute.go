
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Attribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Attribute{}

// Attribute struct for Attribute
type Attribute struct {
	Birthday NullableString `json:"birthday,omitempty"`
	Body NullableString `json:"body,omitempty"`
	Climbing NullableString `json:"climbing,omitempty"`
	EligibleLeague NullableInt32 `json:"eligible_league,omitempty"`
	Fatigue NullableFloat64 `json:"fatigue,omitempty"`
	Generation NullableInt32 `json:"generation,omitempty"`
	Grade NullableString `json:"grade,omitempty"`
	Head NullableString `json:"head,omitempty"`
	Item NullableString `json:"item,omitempty"`
	Level NullableInt32 `json:"level,omitempty"`
	Luck NullableInt32 `json:"luck,omitempty"`
	Momentum NullableInt32 `json:"momentum,omitempty"`
	Name NullableString `json:"name,omitempty"`
	Power NullableInt32 `json:"power,omitempty"`
	Racingstyle NullableString `json:"racingstyle,omitempty"`
	Rank NullableString `json:"rank,omitempty"`
	Running NullableString `json:"running,omitempty"`
	Skin NullableString `json:"skin,omitempty"`
	Speed NullableInt32 `json:"speed,omitempty"`
	Stamina NullableInt32 `json:"stamina,omitempty"`
	Status NullableString `json:"status,omitempty"`
	Swimming NullableString `json:"swimming,omitempty"`
	Technique NullableInt32 `json:"technique,omitempty"`
	Tribe NullableString `json:"tribe,omitempty"`
	Type NullableString `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Attribute Attribute

// NewAttribute instantiates a new Attribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttribute() *Attribute {
	this := Attribute{}
	return &this
}

// NewAttributeWithDefaults instantiates a new Attribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeWithDefaults() *Attribute {
	this := Attribute{}
	return &this
}

// GetBirthday returns the Birthday field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetBirthday() string {
	if o == nil || IsNil(o.Birthday.Get()) {
		var ret string
		return ret
	}
	return *o.Birthday.Get()
}

// GetBirthdayOk returns a tuple with the Birthday field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetBirthdayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Birthday.Get(), o.Birthday.IsSet()
}

// HasBirthday returns a boolean if a field has been set.
func (o *Attribute) HasBirthday() bool {
	if o != nil && o.Birthday.IsSet() {
		return true
	}

	return false
}

// SetBirthday gets a reference to the given NullableString and assigns it to the Birthday field.
func (o *Attribute) SetBirthday(v string) {
	o.Birthday.Set(&v)
}
// SetBirthdayNil sets the value for Birthday to be an explicit nil
func (o *Attribute) SetBirthdayNil() {
	o.Birthday.Set(nil)
}

// UnsetBirthday ensures that no value is present for Birthday, not even an explicit nil
func (o *Attribute) UnsetBirthday() {
	o.Birthday.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *Attribute) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *Attribute) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *Attribute) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *Attribute) UnsetBody() {
	o.Body.Unset()
}

// GetClimbing returns the Climbing field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetClimbing() string {
	if o == nil || IsNil(o.Climbing.Get()) {
		var ret string
		return ret
	}
	return *o.Climbing.Get()
}

// GetClimbingOk returns a tuple with the Climbing field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetClimbingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Climbing.Get(), o.Climbing.IsSet()
}

// HasClimbing returns a boolean if a field has been set.
func (o *Attribute) HasClimbing() bool {
	if o != nil && o.Climbing.IsSet() {
		return true
	}

	return false
}

// SetClimbing gets a reference to the given NullableString and assigns it to the Climbing field.
func (o *Attribute) SetClimbing(v string) {
	o.Climbing.Set(&v)
}
// SetClimbingNil sets the value for Climbing to be an explicit nil
func (o *Attribute) SetClimbingNil() {
	o.Climbing.Set(nil)
}

// UnsetClimbing ensures that no value is present for Climbing, not even an explicit nil
func (o *Attribute) UnsetClimbing() {
	o.Climbing.Unset()
}

// GetEligibleLeague returns the EligibleLeague field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetEligibleLeague() int32 {
	if o == nil || IsNil(o.EligibleLeague.Get()) {
		var ret int32
		return ret
	}
	return *o.EligibleLeague.Get()
}

// GetEligibleLeagueOk returns a tuple with the EligibleLeague field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetEligibleLeagueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EligibleLeague.Get(), o.EligibleLeague.IsSet()
}

// HasEligibleLeague returns a boolean if a field has been set.
func (o *Attribute) HasEligibleLeague() bool {
	if o != nil && o.EligibleLeague.IsSet() {
		return true
	}

	return false
}

// SetEligibleLeague gets a reference to the given NullableInt32 and assigns it to the EligibleLeague field.
func (o *Attribute) SetEligibleLeague(v int32) {
	o.EligibleLeague.Set(&v)
}
// SetEligibleLeagueNil sets the value for EligibleLeague to be an explicit nil
func (o *Attribute) SetEligibleLeagueNil() {
	o.EligibleLeague.Set(nil)
}

// UnsetEligibleLeague ensures that no value is present for EligibleLeague, not even an explicit nil
func (o *Attribute) UnsetEligibleLeague() {
	o.EligibleLeague.Unset()
}

// GetFatigue returns the Fatigue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetFatigue() float64 {
	if o == nil || IsNil(o.Fatigue.Get()) {
		var ret float64
		return ret
	}
	return *o.Fatigue.Get()
}

// GetFatigueOk returns a tuple with the Fatigue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetFatigueOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fatigue.Get(), o.Fatigue.IsSet()
}

// HasFatigue returns a boolean if a field has been set.
func (o *Attribute) HasFatigue() bool {
	if o != nil && o.Fatigue.IsSet() {
		return true
	}

	return false
}

// SetFatigue gets a reference to the given NullableFloat64 and assigns it to the Fatigue field.
func (o *Attribute) SetFatigue(v float64) {
	o.Fatigue.Set(&v)
}
// SetFatigueNil sets the value for Fatigue to be an explicit nil
func (o *Attribute) SetFatigueNil() {
	o.Fatigue.Set(nil)
}

// UnsetFatigue ensures that no value is present for Fatigue, not even an explicit nil
func (o *Attribute) UnsetFatigue() {
	o.Fatigue.Unset()
}

// GetGeneration returns the Generation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetGeneration() int32 {
	if o == nil || IsNil(o.Generation.Get()) {
		var ret int32
		return ret
	}
	return *o.Generation.Get()
}

// GetGenerationOk returns a tuple with the Generation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetGenerationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Generation.Get(), o.Generation.IsSet()
}

// HasGeneration returns a boolean if a field has been set.
func (o *Attribute) HasGeneration() bool {
	if o != nil && o.Generation.IsSet() {
		return true
	}

	return false
}

// SetGeneration gets a reference to the given NullableInt32 and assigns it to the Generation field.
func (o *Attribute) SetGeneration(v int32) {
	o.Generation.Set(&v)
}
// SetGenerationNil sets the value for Generation to be an explicit nil
func (o *Attribute) SetGenerationNil() {
	o.Generation.Set(nil)
}

// UnsetGeneration ensures that no value is present for Generation, not even an explicit nil
func (o *Attribute) UnsetGeneration() {
	o.Generation.Unset()
}

// GetGrade returns the Grade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetGrade() string {
	if o == nil || IsNil(o.Grade.Get()) {
		var ret string
		return ret
	}
	return *o.Grade.Get()
}

// GetGradeOk returns a tuple with the Grade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetGradeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Grade.Get(), o.Grade.IsSet()
}

// HasGrade returns a boolean if a field has been set.
func (o *Attribute) HasGrade() bool {
	if o != nil && o.Grade.IsSet() {
		return true
	}

	return false
}

// SetGrade gets a reference to the given NullableString and assigns it to the Grade field.
func (o *Attribute) SetGrade(v string) {
	o.Grade.Set(&v)
}
// SetGradeNil sets the value for Grade to be an explicit nil
func (o *Attribute) SetGradeNil() {
	o.Grade.Set(nil)
}

// UnsetGrade ensures that no value is present for Grade, not even an explicit nil
func (o *Attribute) UnsetGrade() {
	o.Grade.Unset()
}

// GetHead returns the Head field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetHead() string {
	if o == nil || IsNil(o.Head.Get()) {
		var ret string
		return ret
	}
	return *o.Head.Get()
}

// GetHeadOk returns a tuple with the Head field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetHeadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Head.Get(), o.Head.IsSet()
}

// HasHead returns a boolean if a field has been set.
func (o *Attribute) HasHead() bool {
	if o != nil && o.Head.IsSet() {
		return true
	}

	return false
}

// SetHead gets a reference to the given NullableString and assigns it to the Head field.
func (o *Attribute) SetHead(v string) {
	o.Head.Set(&v)
}
// SetHeadNil sets the value for Head to be an explicit nil
func (o *Attribute) SetHeadNil() {
	o.Head.Set(nil)
}

// UnsetHead ensures that no value is present for Head, not even an explicit nil
func (o *Attribute) UnsetHead() {
	o.Head.Unset()
}

// GetItem returns the Item field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetItem() string {
	if o == nil || IsNil(o.Item.Get()) {
		var ret string
		return ret
	}
	return *o.Item.Get()
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetItemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Item.Get(), o.Item.IsSet()
}

// HasItem returns a boolean if a field has been set.
func (o *Attribute) HasItem() bool {
	if o != nil && o.Item.IsSet() {
		return true
	}

	return false
}

// SetItem gets a reference to the given NullableString and assigns it to the Item field.
func (o *Attribute) SetItem(v string) {
	o.Item.Set(&v)
}
// SetItemNil sets the value for Item to be an explicit nil
func (o *Attribute) SetItemNil() {
	o.Item.Set(nil)
}

// UnsetItem ensures that no value is present for Item, not even an explicit nil
func (o *Attribute) UnsetItem() {
	o.Item.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *Attribute) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *Attribute) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *Attribute) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *Attribute) UnsetLevel() {
	o.Level.Unset()
}

// GetLuck returns the Luck field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetLuck() int32 {
	if o == nil || IsNil(o.Luck.Get()) {
		var ret int32
		return ret
	}
	return *o.Luck.Get()
}

// GetLuckOk returns a tuple with the Luck field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetLuckOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Luck.Get(), o.Luck.IsSet()
}

// HasLuck returns a boolean if a field has been set.
func (o *Attribute) HasLuck() bool {
	if o != nil && o.Luck.IsSet() {
		return true
	}

	return false
}

// SetLuck gets a reference to the given NullableInt32 and assigns it to the Luck field.
func (o *Attribute) SetLuck(v int32) {
	o.Luck.Set(&v)
}
// SetLuckNil sets the value for Luck to be an explicit nil
func (o *Attribute) SetLuckNil() {
	o.Luck.Set(nil)
}

// UnsetLuck ensures that no value is present for Luck, not even an explicit nil
func (o *Attribute) UnsetLuck() {
	o.Luck.Unset()
}

// GetMomentum returns the Momentum field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetMomentum() int32 {
	if o == nil || IsNil(o.Momentum.Get()) {
		var ret int32
		return ret
	}
	return *o.Momentum.Get()
}

// GetMomentumOk returns a tuple with the Momentum field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetMomentumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Momentum.Get(), o.Momentum.IsSet()
}

// HasMomentum returns a boolean if a field has been set.
func (o *Attribute) HasMomentum() bool {
	if o != nil && o.Momentum.IsSet() {
		return true
	}

	return false
}

// SetMomentum gets a reference to the given NullableInt32 and assigns it to the Momentum field.
func (o *Attribute) SetMomentum(v int32) {
	o.Momentum.Set(&v)
}
// SetMomentumNil sets the value for Momentum to be an explicit nil
func (o *Attribute) SetMomentumNil() {
	o.Momentum.Set(nil)
}

// UnsetMomentum ensures that no value is present for Momentum, not even an explicit nil
func (o *Attribute) UnsetMomentum() {
	o.Momentum.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Attribute) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Attribute) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *Attribute) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Attribute) UnsetName() {
	o.Name.Unset()
}

// GetPower returns the Power field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetPower() int32 {
	if o == nil || IsNil(o.Power.Get()) {
		var ret int32
		return ret
	}
	return *o.Power.Get()
}

// GetPowerOk returns a tuple with the Power field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetPowerOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Power.Get(), o.Power.IsSet()
}

// HasPower returns a boolean if a field has been set.
func (o *Attribute) HasPower() bool {
	if o != nil && o.Power.IsSet() {
		return true
	}

	return false
}

// SetPower gets a reference to the given NullableInt32 and assigns it to the Power field.
func (o *Attribute) SetPower(v int32) {
	o.Power.Set(&v)
}
// SetPowerNil sets the value for Power to be an explicit nil
func (o *Attribute) SetPowerNil() {
	o.Power.Set(nil)
}

// UnsetPower ensures that no value is present for Power, not even an explicit nil
func (o *Attribute) UnsetPower() {
	o.Power.Unset()
}

// GetRacingstyle returns the Racingstyle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetRacingstyle() string {
	if o == nil || IsNil(o.Racingstyle.Get()) {
		var ret string
		return ret
	}
	return *o.Racingstyle.Get()
}

// GetRacingstyleOk returns a tuple with the Racingstyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetRacingstyleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Racingstyle.Get(), o.Racingstyle.IsSet()
}

// HasRacingstyle returns a boolean if a field has been set.
func (o *Attribute) HasRacingstyle() bool {
	if o != nil && o.Racingstyle.IsSet() {
		return true
	}

	return false
}

// SetRacingstyle gets a reference to the given NullableString and assigns it to the Racingstyle field.
func (o *Attribute) SetRacingstyle(v string) {
	o.Racingstyle.Set(&v)
}
// SetRacingstyleNil sets the value for Racingstyle to be an explicit nil
func (o *Attribute) SetRacingstyleNil() {
	o.Racingstyle.Set(nil)
}

// UnsetRacingstyle ensures that no value is present for Racingstyle, not even an explicit nil
func (o *Attribute) UnsetRacingstyle() {
	o.Racingstyle.Unset()
}

// GetRank returns the Rank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetRank() string {
	if o == nil || IsNil(o.Rank.Get()) {
		var ret string
		return ret
	}
	return *o.Rank.Get()
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetRankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rank.Get(), o.Rank.IsSet()
}

// HasRank returns a boolean if a field has been set.
func (o *Attribute) HasRank() bool {
	if o != nil && o.Rank.IsSet() {
		return true
	}

	return false
}

// SetRank gets a reference to the given NullableString and assigns it to the Rank field.
func (o *Attribute) SetRank(v string) {
	o.Rank.Set(&v)
}
// SetRankNil sets the value for Rank to be an explicit nil
func (o *Attribute) SetRankNil() {
	o.Rank.Set(nil)
}

// UnsetRank ensures that no value is present for Rank, not even an explicit nil
func (o *Attribute) UnsetRank() {
	o.Rank.Unset()
}

// GetRunning returns the Running field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetRunning() string {
	if o == nil || IsNil(o.Running.Get()) {
		var ret string
		return ret
	}
	return *o.Running.Get()
}

// GetRunningOk returns a tuple with the Running field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetRunningOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Running.Get(), o.Running.IsSet()
}

// HasRunning returns a boolean if a field has been set.
func (o *Attribute) HasRunning() bool {
	if o != nil && o.Running.IsSet() {
		return true
	}

	return false
}

// SetRunning gets a reference to the given NullableString and assigns it to the Running field.
func (o *Attribute) SetRunning(v string) {
	o.Running.Set(&v)
}
// SetRunningNil sets the value for Running to be an explicit nil
func (o *Attribute) SetRunningNil() {
	o.Running.Set(nil)
}

// UnsetRunning ensures that no value is present for Running, not even an explicit nil
func (o *Attribute) UnsetRunning() {
	o.Running.Unset()
}

// GetSkin returns the Skin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetSkin() string {
	if o == nil || IsNil(o.Skin.Get()) {
		var ret string
		return ret
	}
	return *o.Skin.Get()
}

// GetSkinOk returns a tuple with the Skin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetSkinOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Skin.Get(), o.Skin.IsSet()
}

// HasSkin returns a boolean if a field has been set.
func (o *Attribute) HasSkin() bool {
	if o != nil && o.Skin.IsSet() {
		return true
	}

	return false
}

// SetSkin gets a reference to the given NullableString and assigns it to the Skin field.
func (o *Attribute) SetSkin(v string) {
	o.Skin.Set(&v)
}
// SetSkinNil sets the value for Skin to be an explicit nil
func (o *Attribute) SetSkinNil() {
	o.Skin.Set(nil)
}

// UnsetSkin ensures that no value is present for Skin, not even an explicit nil
func (o *Attribute) UnsetSkin() {
	o.Skin.Unset()
}

// GetSpeed returns the Speed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetSpeed() int32 {
	if o == nil || IsNil(o.Speed.Get()) {
		var ret int32
		return ret
	}
	return *o.Speed.Get()
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetSpeedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Speed.Get(), o.Speed.IsSet()
}

// HasSpeed returns a boolean if a field has been set.
func (o *Attribute) HasSpeed() bool {
	if o != nil && o.Speed.IsSet() {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given NullableInt32 and assigns it to the Speed field.
func (o *Attribute) SetSpeed(v int32) {
	o.Speed.Set(&v)
}
// SetSpeedNil sets the value for Speed to be an explicit nil
func (o *Attribute) SetSpeedNil() {
	o.Speed.Set(nil)
}

// UnsetSpeed ensures that no value is present for Speed, not even an explicit nil
func (o *Attribute) UnsetSpeed() {
	o.Speed.Unset()
}

// GetStamina returns the Stamina field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetStamina() int32 {
	if o == nil || IsNil(o.Stamina.Get()) {
		var ret int32
		return ret
	}
	return *o.Stamina.Get()
}

// GetStaminaOk returns a tuple with the Stamina field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetStaminaOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Stamina.Get(), o.Stamina.IsSet()
}

// HasStamina returns a boolean if a field has been set.
func (o *Attribute) HasStamina() bool {
	if o != nil && o.Stamina.IsSet() {
		return true
	}

	return false
}

// SetStamina gets a reference to the given NullableInt32 and assigns it to the Stamina field.
func (o *Attribute) SetStamina(v int32) {
	o.Stamina.Set(&v)
}
// SetStaminaNil sets the value for Stamina to be an explicit nil
func (o *Attribute) SetStaminaNil() {
	o.Stamina.Set(nil)
}

// UnsetStamina ensures that no value is present for Stamina, not even an explicit nil
func (o *Attribute) UnsetStamina() {
	o.Stamina.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *Attribute) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *Attribute) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *Attribute) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *Attribute) UnsetStatus() {
	o.Status.Unset()
}

// GetSwimming returns the Swimming field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetSwimming() string {
	if o == nil || IsNil(o.Swimming.Get()) {
		var ret string
		return ret
	}
	return *o.Swimming.Get()
}

// GetSwimmingOk returns a tuple with the Swimming field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetSwimmingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Swimming.Get(), o.Swimming.IsSet()
}

// HasSwimming returns a boolean if a field has been set.
func (o *Attribute) HasSwimming() bool {
	if o != nil && o.Swimming.IsSet() {
		return true
	}

	return false
}

// SetSwimming gets a reference to the given NullableString and assigns it to the Swimming field.
func (o *Attribute) SetSwimming(v string) {
	o.Swimming.Set(&v)
}
// SetSwimmingNil sets the value for Swimming to be an explicit nil
func (o *Attribute) SetSwimmingNil() {
	o.Swimming.Set(nil)
}

// UnsetSwimming ensures that no value is present for Swimming, not even an explicit nil
func (o *Attribute) UnsetSwimming() {
	o.Swimming.Unset()
}

// GetTechnique returns the Technique field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetTechnique() int32 {
	if o == nil || IsNil(o.Technique.Get()) {
		var ret int32
		return ret
	}
	return *o.Technique.Get()
}

// GetTechniqueOk returns a tuple with the Technique field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetTechniqueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Technique.Get(), o.Technique.IsSet()
}

// HasTechnique returns a boolean if a field has been set.
func (o *Attribute) HasTechnique() bool {
	if o != nil && o.Technique.IsSet() {
		return true
	}

	return false
}

// SetTechnique gets a reference to the given NullableInt32 and assigns it to the Technique field.
func (o *Attribute) SetTechnique(v int32) {
	o.Technique.Set(&v)
}
// SetTechniqueNil sets the value for Technique to be an explicit nil
func (o *Attribute) SetTechniqueNil() {
	o.Technique.Set(nil)
}

// UnsetTechnique ensures that no value is present for Technique, not even an explicit nil
func (o *Attribute) UnsetTechnique() {
	o.Technique.Unset()
}

// GetTribe returns the Tribe field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetTribe() string {
	if o == nil || IsNil(o.Tribe.Get()) {
		var ret string
		return ret
	}
	return *o.Tribe.Get()
}

// GetTribeOk returns a tuple with the Tribe field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetTribeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tribe.Get(), o.Tribe.IsSet()
}

// HasTribe returns a boolean if a field has been set.
func (o *Attribute) HasTribe() bool {
	if o != nil && o.Tribe.IsSet() {
		return true
	}

	return false
}

// SetTribe gets a reference to the given NullableString and assigns it to the Tribe field.
func (o *Attribute) SetTribe(v string) {
	o.Tribe.Set(&v)
}
// SetTribeNil sets the value for Tribe to be an explicit nil
func (o *Attribute) SetTribeNil() {
	o.Tribe.Set(nil)
}

// UnsetTribe ensures that no value is present for Tribe, not even an explicit nil
func (o *Attribute) UnsetTribe() {
	o.Tribe.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Attribute) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Attribute) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Attribute) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Attribute) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Attribute) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Attribute) UnsetType() {
	o.Type.Unset()
}

func (o Attribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Attribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Birthday.IsSet() {
		toSerialize["birthday"] = o.Birthday.Get()
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if o.Climbing.IsSet() {
		toSerialize["climbing"] = o.Climbing.Get()
	}
	if o.EligibleLeague.IsSet() {
		toSerialize["eligible_league"] = o.EligibleLeague.Get()
	}
	if o.Fatigue.IsSet() {
		toSerialize["fatigue"] = o.Fatigue.Get()
	}
	if o.Generation.IsSet() {
		toSerialize["generation"] = o.Generation.Get()
	}
	if o.Grade.IsSet() {
		toSerialize["grade"] = o.Grade.Get()
	}
	if o.Head.IsSet() {
		toSerialize["head"] = o.Head.Get()
	}
	if o.Item.IsSet() {
		toSerialize["item"] = o.Item.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.Luck.IsSet() {
		toSerialize["luck"] = o.Luck.Get()
	}
	if o.Momentum.IsSet() {
		toSerialize["momentum"] = o.Momentum.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Power.IsSet() {
		toSerialize["power"] = o.Power.Get()
	}
	if o.Racingstyle.IsSet() {
		toSerialize["racingstyle"] = o.Racingstyle.Get()
	}
	if o.Rank.IsSet() {
		toSerialize["rank"] = o.Rank.Get()
	}
	if o.Running.IsSet() {
		toSerialize["running"] = o.Running.Get()
	}
	if o.Skin.IsSet() {
		toSerialize["skin"] = o.Skin.Get()
	}
	if o.Speed.IsSet() {
		toSerialize["speed"] = o.Speed.Get()
	}
	if o.Stamina.IsSet() {
		toSerialize["stamina"] = o.Stamina.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Swimming.IsSet() {
		toSerialize["swimming"] = o.Swimming.Get()
	}
	if o.Technique.IsSet() {
		toSerialize["technique"] = o.Technique.Get()
	}
	if o.Tribe.IsSet() {
		toSerialize["tribe"] = o.Tribe.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Attribute) UnmarshalJSON(data []byte) (err error) {
	varAttribute := _Attribute{}

	err = json.Unmarshal(data, &varAttribute)

	if err != nil {
		return err
	}

	*o = Attribute(varAttribute)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "birthday")
		delete(additionalProperties, "body")
		delete(additionalProperties, "climbing")
		delete(additionalProperties, "eligible_league")
		delete(additionalProperties, "fatigue")
		delete(additionalProperties, "generation")
		delete(additionalProperties, "grade")
		delete(additionalProperties, "head")
		delete(additionalProperties, "item")
		delete(additionalProperties, "level")
		delete(additionalProperties, "luck")
		delete(additionalProperties, "momentum")
		delete(additionalProperties, "name")
		delete(additionalProperties, "power")
		delete(additionalProperties, "racingstyle")
		delete(additionalProperties, "rank")
		delete(additionalProperties, "running")
		delete(additionalProperties, "skin")
		delete(additionalProperties, "speed")
		delete(additionalProperties, "stamina")
		delete(additionalProperties, "status")
		delete(additionalProperties, "swimming")
		delete(additionalProperties, "technique")
		delete(additionalProperties, "tribe")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttribute struct {
	value *Attribute
	isSet bool
}

func (v NullableAttribute) Get() *Attribute {
	return v.value
}

func (v *NullableAttribute) Set(val *Attribute) {
	v.value = val
	v.isSet = true
}

func (v NullableAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttribute(val *Attribute) *NullableAttribute {
	return &NullableAttribute{value: val, isSet: true}
}

func (v NullableAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


