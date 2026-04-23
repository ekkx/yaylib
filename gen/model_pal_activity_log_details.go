
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the PalActivityLogDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PalActivityLogDetails{}

// PalActivityLogDetails struct for PalActivityLogDetails
type PalActivityLogDetails struct {
	Attacking NullableInt32 `json:"attacking,omitempty"`
	Cost NullableFloat64 `json:"cost,omitempty"`
	GachaType NullableString `json:"gacha_type,omitempty"`
	ItemStatus NullableString `json:"item_status,omitempty"`
	ItemSubtype NullableString `json:"item_subtype,omitempty"`
	ItemType NullableString `json:"item_type,omitempty"`
	Level NullableInt32 `json:"level,omitempty"`
	NftType NullableString `json:"nft_type,omitempty"`
	Outcome NullableInt64 `json:"outcome,omitempty"`
	PalGrade NullableString `json:"pal_grade,omitempty"`
	PalImage NullableString `json:"pal_image,omitempty"`
	PalName NullableString `json:"pal_name,omitempty"`
	Places []int32 `json:"places,omitempty"`
	RaceId NullableString `json:"race_id,omitempty"`
	Rank NullableString `json:"rank,omitempty"`
	Result NullableString `json:"result,omitempty"`
	TokenId NullableString `json:"token_id,omitempty"`
	TransactionCode NullableString `json:"transaction_code,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PalActivityLogDetails PalActivityLogDetails

// NewPalActivityLogDetails instantiates a new PalActivityLogDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPalActivityLogDetails() *PalActivityLogDetails {
	this := PalActivityLogDetails{}
	return &this
}

// NewPalActivityLogDetailsWithDefaults instantiates a new PalActivityLogDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPalActivityLogDetailsWithDefaults() *PalActivityLogDetails {
	this := PalActivityLogDetails{}
	return &this
}

// GetAttacking returns the Attacking field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetAttacking() int32 {
	if o == nil || IsNil(o.Attacking.Get()) {
		var ret int32
		return ret
	}
	return *o.Attacking.Get()
}

// GetAttackingOk returns a tuple with the Attacking field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetAttackingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attacking.Get(), o.Attacking.IsSet()
}

// HasAttacking returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasAttacking() bool {
	if o != nil && o.Attacking.IsSet() {
		return true
	}

	return false
}

// SetAttacking gets a reference to the given NullableInt32 and assigns it to the Attacking field.
func (o *PalActivityLogDetails) SetAttacking(v int32) {
	o.Attacking.Set(&v)
}
// SetAttackingNil sets the value for Attacking to be an explicit nil
func (o *PalActivityLogDetails) SetAttackingNil() {
	o.Attacking.Set(nil)
}

// UnsetAttacking ensures that no value is present for Attacking, not even an explicit nil
func (o *PalActivityLogDetails) UnsetAttacking() {
	o.Attacking.Unset()
}

// GetCost returns the Cost field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetCost() float64 {
	if o == nil || IsNil(o.Cost.Get()) {
		var ret float64
		return ret
	}
	return *o.Cost.Get()
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetCostOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cost.Get(), o.Cost.IsSet()
}

// HasCost returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasCost() bool {
	if o != nil && o.Cost.IsSet() {
		return true
	}

	return false
}

// SetCost gets a reference to the given NullableFloat64 and assigns it to the Cost field.
func (o *PalActivityLogDetails) SetCost(v float64) {
	o.Cost.Set(&v)
}
// SetCostNil sets the value for Cost to be an explicit nil
func (o *PalActivityLogDetails) SetCostNil() {
	o.Cost.Set(nil)
}

// UnsetCost ensures that no value is present for Cost, not even an explicit nil
func (o *PalActivityLogDetails) UnsetCost() {
	o.Cost.Unset()
}

// GetGachaType returns the GachaType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetGachaType() string {
	if o == nil || IsNil(o.GachaType.Get()) {
		var ret string
		return ret
	}
	return *o.GachaType.Get()
}

// GetGachaTypeOk returns a tuple with the GachaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetGachaTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GachaType.Get(), o.GachaType.IsSet()
}

// HasGachaType returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasGachaType() bool {
	if o != nil && o.GachaType.IsSet() {
		return true
	}

	return false
}

// SetGachaType gets a reference to the given NullableString and assigns it to the GachaType field.
func (o *PalActivityLogDetails) SetGachaType(v string) {
	o.GachaType.Set(&v)
}
// SetGachaTypeNil sets the value for GachaType to be an explicit nil
func (o *PalActivityLogDetails) SetGachaTypeNil() {
	o.GachaType.Set(nil)
}

// UnsetGachaType ensures that no value is present for GachaType, not even an explicit nil
func (o *PalActivityLogDetails) UnsetGachaType() {
	o.GachaType.Unset()
}

// GetItemStatus returns the ItemStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetItemStatus() string {
	if o == nil || IsNil(o.ItemStatus.Get()) {
		var ret string
		return ret
	}
	return *o.ItemStatus.Get()
}

// GetItemStatusOk returns a tuple with the ItemStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetItemStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemStatus.Get(), o.ItemStatus.IsSet()
}

// HasItemStatus returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasItemStatus() bool {
	if o != nil && o.ItemStatus.IsSet() {
		return true
	}

	return false
}

// SetItemStatus gets a reference to the given NullableString and assigns it to the ItemStatus field.
func (o *PalActivityLogDetails) SetItemStatus(v string) {
	o.ItemStatus.Set(&v)
}
// SetItemStatusNil sets the value for ItemStatus to be an explicit nil
func (o *PalActivityLogDetails) SetItemStatusNil() {
	o.ItemStatus.Set(nil)
}

// UnsetItemStatus ensures that no value is present for ItemStatus, not even an explicit nil
func (o *PalActivityLogDetails) UnsetItemStatus() {
	o.ItemStatus.Unset()
}

// GetItemSubtype returns the ItemSubtype field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetItemSubtype() string {
	if o == nil || IsNil(o.ItemSubtype.Get()) {
		var ret string
		return ret
	}
	return *o.ItemSubtype.Get()
}

// GetItemSubtypeOk returns a tuple with the ItemSubtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetItemSubtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemSubtype.Get(), o.ItemSubtype.IsSet()
}

// HasItemSubtype returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasItemSubtype() bool {
	if o != nil && o.ItemSubtype.IsSet() {
		return true
	}

	return false
}

// SetItemSubtype gets a reference to the given NullableString and assigns it to the ItemSubtype field.
func (o *PalActivityLogDetails) SetItemSubtype(v string) {
	o.ItemSubtype.Set(&v)
}
// SetItemSubtypeNil sets the value for ItemSubtype to be an explicit nil
func (o *PalActivityLogDetails) SetItemSubtypeNil() {
	o.ItemSubtype.Set(nil)
}

// UnsetItemSubtype ensures that no value is present for ItemSubtype, not even an explicit nil
func (o *PalActivityLogDetails) UnsetItemSubtype() {
	o.ItemSubtype.Unset()
}

// GetItemType returns the ItemType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetItemType() string {
	if o == nil || IsNil(o.ItemType.Get()) {
		var ret string
		return ret
	}
	return *o.ItemType.Get()
}

// GetItemTypeOk returns a tuple with the ItemType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetItemTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ItemType.Get(), o.ItemType.IsSet()
}

// HasItemType returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasItemType() bool {
	if o != nil && o.ItemType.IsSet() {
		return true
	}

	return false
}

// SetItemType gets a reference to the given NullableString and assigns it to the ItemType field.
func (o *PalActivityLogDetails) SetItemType(v string) {
	o.ItemType.Set(&v)
}
// SetItemTypeNil sets the value for ItemType to be an explicit nil
func (o *PalActivityLogDetails) SetItemTypeNil() {
	o.ItemType.Set(nil)
}

// UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
func (o *PalActivityLogDetails) UnsetItemType() {
	o.ItemType.Unset()
}

// GetLevel returns the Level field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetLevel() int32 {
	if o == nil || IsNil(o.Level.Get()) {
		var ret int32
		return ret
	}
	return *o.Level.Get()
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Level.Get(), o.Level.IsSet()
}

// HasLevel returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasLevel() bool {
	if o != nil && o.Level.IsSet() {
		return true
	}

	return false
}

// SetLevel gets a reference to the given NullableInt32 and assigns it to the Level field.
func (o *PalActivityLogDetails) SetLevel(v int32) {
	o.Level.Set(&v)
}
// SetLevelNil sets the value for Level to be an explicit nil
func (o *PalActivityLogDetails) SetLevelNil() {
	o.Level.Set(nil)
}

// UnsetLevel ensures that no value is present for Level, not even an explicit nil
func (o *PalActivityLogDetails) UnsetLevel() {
	o.Level.Unset()
}

// GetNftType returns the NftType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetNftType() string {
	if o == nil || IsNil(o.NftType.Get()) {
		var ret string
		return ret
	}
	return *o.NftType.Get()
}

// GetNftTypeOk returns a tuple with the NftType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetNftTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NftType.Get(), o.NftType.IsSet()
}

// HasNftType returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasNftType() bool {
	if o != nil && o.NftType.IsSet() {
		return true
	}

	return false
}

// SetNftType gets a reference to the given NullableString and assigns it to the NftType field.
func (o *PalActivityLogDetails) SetNftType(v string) {
	o.NftType.Set(&v)
}
// SetNftTypeNil sets the value for NftType to be an explicit nil
func (o *PalActivityLogDetails) SetNftTypeNil() {
	o.NftType.Set(nil)
}

// UnsetNftType ensures that no value is present for NftType, not even an explicit nil
func (o *PalActivityLogDetails) UnsetNftType() {
	o.NftType.Unset()
}

// GetOutcome returns the Outcome field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetOutcome() int64 {
	if o == nil || IsNil(o.Outcome.Get()) {
		var ret int64
		return ret
	}
	return *o.Outcome.Get()
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetOutcomeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Outcome.Get(), o.Outcome.IsSet()
}

// HasOutcome returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasOutcome() bool {
	if o != nil && o.Outcome.IsSet() {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given NullableInt64 and assigns it to the Outcome field.
func (o *PalActivityLogDetails) SetOutcome(v int64) {
	o.Outcome.Set(&v)
}
// SetOutcomeNil sets the value for Outcome to be an explicit nil
func (o *PalActivityLogDetails) SetOutcomeNil() {
	o.Outcome.Set(nil)
}

// UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
func (o *PalActivityLogDetails) UnsetOutcome() {
	o.Outcome.Unset()
}

// GetPalGrade returns the PalGrade field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetPalGrade() string {
	if o == nil || IsNil(o.PalGrade.Get()) {
		var ret string
		return ret
	}
	return *o.PalGrade.Get()
}

// GetPalGradeOk returns a tuple with the PalGrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetPalGradeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalGrade.Get(), o.PalGrade.IsSet()
}

// HasPalGrade returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasPalGrade() bool {
	if o != nil && o.PalGrade.IsSet() {
		return true
	}

	return false
}

// SetPalGrade gets a reference to the given NullableString and assigns it to the PalGrade field.
func (o *PalActivityLogDetails) SetPalGrade(v string) {
	o.PalGrade.Set(&v)
}
// SetPalGradeNil sets the value for PalGrade to be an explicit nil
func (o *PalActivityLogDetails) SetPalGradeNil() {
	o.PalGrade.Set(nil)
}

// UnsetPalGrade ensures that no value is present for PalGrade, not even an explicit nil
func (o *PalActivityLogDetails) UnsetPalGrade() {
	o.PalGrade.Unset()
}

// GetPalImage returns the PalImage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetPalImage() string {
	if o == nil || IsNil(o.PalImage.Get()) {
		var ret string
		return ret
	}
	return *o.PalImage.Get()
}

// GetPalImageOk returns a tuple with the PalImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetPalImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalImage.Get(), o.PalImage.IsSet()
}

// HasPalImage returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasPalImage() bool {
	if o != nil && o.PalImage.IsSet() {
		return true
	}

	return false
}

// SetPalImage gets a reference to the given NullableString and assigns it to the PalImage field.
func (o *PalActivityLogDetails) SetPalImage(v string) {
	o.PalImage.Set(&v)
}
// SetPalImageNil sets the value for PalImage to be an explicit nil
func (o *PalActivityLogDetails) SetPalImageNil() {
	o.PalImage.Set(nil)
}

// UnsetPalImage ensures that no value is present for PalImage, not even an explicit nil
func (o *PalActivityLogDetails) UnsetPalImage() {
	o.PalImage.Unset()
}

// GetPalName returns the PalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetPalName() string {
	if o == nil || IsNil(o.PalName.Get()) {
		var ret string
		return ret
	}
	return *o.PalName.Get()
}

// GetPalNameOk returns a tuple with the PalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetPalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PalName.Get(), o.PalName.IsSet()
}

// HasPalName returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasPalName() bool {
	if o != nil && o.PalName.IsSet() {
		return true
	}

	return false
}

// SetPalName gets a reference to the given NullableString and assigns it to the PalName field.
func (o *PalActivityLogDetails) SetPalName(v string) {
	o.PalName.Set(&v)
}
// SetPalNameNil sets the value for PalName to be an explicit nil
func (o *PalActivityLogDetails) SetPalNameNil() {
	o.PalName.Set(nil)
}

// UnsetPalName ensures that no value is present for PalName, not even an explicit nil
func (o *PalActivityLogDetails) UnsetPalName() {
	o.PalName.Unset()
}

// GetPlaces returns the Places field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetPlaces() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.Places
}

// GetPlacesOk returns a tuple with the Places field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetPlacesOk() ([]int32, bool) {
	if o == nil || IsNil(o.Places) {
		return nil, false
	}
	return o.Places, true
}

// HasPlaces returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasPlaces() bool {
	if o != nil && !IsNil(o.Places) {
		return true
	}

	return false
}

// SetPlaces gets a reference to the given []int32 and assigns it to the Places field.
func (o *PalActivityLogDetails) SetPlaces(v []int32) {
	o.Places = v
}

// GetRaceId returns the RaceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetRaceId() string {
	if o == nil || IsNil(o.RaceId.Get()) {
		var ret string
		return ret
	}
	return *o.RaceId.Get()
}

// GetRaceIdOk returns a tuple with the RaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetRaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RaceId.Get(), o.RaceId.IsSet()
}

// HasRaceId returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasRaceId() bool {
	if o != nil && o.RaceId.IsSet() {
		return true
	}

	return false
}

// SetRaceId gets a reference to the given NullableString and assigns it to the RaceId field.
func (o *PalActivityLogDetails) SetRaceId(v string) {
	o.RaceId.Set(&v)
}
// SetRaceIdNil sets the value for RaceId to be an explicit nil
func (o *PalActivityLogDetails) SetRaceIdNil() {
	o.RaceId.Set(nil)
}

// UnsetRaceId ensures that no value is present for RaceId, not even an explicit nil
func (o *PalActivityLogDetails) UnsetRaceId() {
	o.RaceId.Unset()
}

// GetRank returns the Rank field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetRank() string {
	if o == nil || IsNil(o.Rank.Get()) {
		var ret string
		return ret
	}
	return *o.Rank.Get()
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetRankOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rank.Get(), o.Rank.IsSet()
}

// HasRank returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasRank() bool {
	if o != nil && o.Rank.IsSet() {
		return true
	}

	return false
}

// SetRank gets a reference to the given NullableString and assigns it to the Rank field.
func (o *PalActivityLogDetails) SetRank(v string) {
	o.Rank.Set(&v)
}
// SetRankNil sets the value for Rank to be an explicit nil
func (o *PalActivityLogDetails) SetRankNil() {
	o.Rank.Set(nil)
}

// UnsetRank ensures that no value is present for Rank, not even an explicit nil
func (o *PalActivityLogDetails) UnsetRank() {
	o.Rank.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetResult() string {
	if o == nil || IsNil(o.Result.Get()) {
		var ret string
		return ret
	}
	return *o.Result.Get()
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result.Get(), o.Result.IsSet()
}

// HasResult returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasResult() bool {
	if o != nil && o.Result.IsSet() {
		return true
	}

	return false
}

// SetResult gets a reference to the given NullableString and assigns it to the Result field.
func (o *PalActivityLogDetails) SetResult(v string) {
	o.Result.Set(&v)
}
// SetResultNil sets the value for Result to be an explicit nil
func (o *PalActivityLogDetails) SetResultNil() {
	o.Result.Set(nil)
}

// UnsetResult ensures that no value is present for Result, not even an explicit nil
func (o *PalActivityLogDetails) UnsetResult() {
	o.Result.Unset()
}

// GetTokenId returns the TokenId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetTokenId() string {
	if o == nil || IsNil(o.TokenId.Get()) {
		var ret string
		return ret
	}
	return *o.TokenId.Get()
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenId.Get(), o.TokenId.IsSet()
}

// HasTokenId returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasTokenId() bool {
	if o != nil && o.TokenId.IsSet() {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given NullableString and assigns it to the TokenId field.
func (o *PalActivityLogDetails) SetTokenId(v string) {
	o.TokenId.Set(&v)
}
// SetTokenIdNil sets the value for TokenId to be an explicit nil
func (o *PalActivityLogDetails) SetTokenIdNil() {
	o.TokenId.Set(nil)
}

// UnsetTokenId ensures that no value is present for TokenId, not even an explicit nil
func (o *PalActivityLogDetails) UnsetTokenId() {
	o.TokenId.Unset()
}

// GetTransactionCode returns the TransactionCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PalActivityLogDetails) GetTransactionCode() string {
	if o == nil || IsNil(o.TransactionCode.Get()) {
		var ret string
		return ret
	}
	return *o.TransactionCode.Get()
}

// GetTransactionCodeOk returns a tuple with the TransactionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PalActivityLogDetails) GetTransactionCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransactionCode.Get(), o.TransactionCode.IsSet()
}

// HasTransactionCode returns a boolean if a field has been set.
func (o *PalActivityLogDetails) HasTransactionCode() bool {
	if o != nil && o.TransactionCode.IsSet() {
		return true
	}

	return false
}

// SetTransactionCode gets a reference to the given NullableString and assigns it to the TransactionCode field.
func (o *PalActivityLogDetails) SetTransactionCode(v string) {
	o.TransactionCode.Set(&v)
}
// SetTransactionCodeNil sets the value for TransactionCode to be an explicit nil
func (o *PalActivityLogDetails) SetTransactionCodeNil() {
	o.TransactionCode.Set(nil)
}

// UnsetTransactionCode ensures that no value is present for TransactionCode, not even an explicit nil
func (o *PalActivityLogDetails) UnsetTransactionCode() {
	o.TransactionCode.Unset()
}

func (o PalActivityLogDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PalActivityLogDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Attacking.IsSet() {
		toSerialize["attacking"] = o.Attacking.Get()
	}
	if o.Cost.IsSet() {
		toSerialize["cost"] = o.Cost.Get()
	}
	if o.GachaType.IsSet() {
		toSerialize["gacha_type"] = o.GachaType.Get()
	}
	if o.ItemStatus.IsSet() {
		toSerialize["item_status"] = o.ItemStatus.Get()
	}
	if o.ItemSubtype.IsSet() {
		toSerialize["item_subtype"] = o.ItemSubtype.Get()
	}
	if o.ItemType.IsSet() {
		toSerialize["item_type"] = o.ItemType.Get()
	}
	if o.Level.IsSet() {
		toSerialize["level"] = o.Level.Get()
	}
	if o.NftType.IsSet() {
		toSerialize["nft_type"] = o.NftType.Get()
	}
	if o.Outcome.IsSet() {
		toSerialize["outcome"] = o.Outcome.Get()
	}
	if o.PalGrade.IsSet() {
		toSerialize["pal_grade"] = o.PalGrade.Get()
	}
	if o.PalImage.IsSet() {
		toSerialize["pal_image"] = o.PalImage.Get()
	}
	if o.PalName.IsSet() {
		toSerialize["pal_name"] = o.PalName.Get()
	}
	if o.Places != nil {
		toSerialize["places"] = o.Places
	}
	if o.RaceId.IsSet() {
		toSerialize["race_id"] = o.RaceId.Get()
	}
	if o.Rank.IsSet() {
		toSerialize["rank"] = o.Rank.Get()
	}
	if o.Result.IsSet() {
		toSerialize["result"] = o.Result.Get()
	}
	if o.TokenId.IsSet() {
		toSerialize["token_id"] = o.TokenId.Get()
	}
	if o.TransactionCode.IsSet() {
		toSerialize["transaction_code"] = o.TransactionCode.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PalActivityLogDetails) UnmarshalJSON(data []byte) (err error) {
	varPalActivityLogDetails := _PalActivityLogDetails{}

	err = json.Unmarshal(data, &varPalActivityLogDetails)

	if err != nil {
		return err
	}

	*o = PalActivityLogDetails(varPalActivityLogDetails)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "attacking")
		delete(additionalProperties, "cost")
		delete(additionalProperties, "gacha_type")
		delete(additionalProperties, "item_status")
		delete(additionalProperties, "item_subtype")
		delete(additionalProperties, "item_type")
		delete(additionalProperties, "level")
		delete(additionalProperties, "nft_type")
		delete(additionalProperties, "outcome")
		delete(additionalProperties, "pal_grade")
		delete(additionalProperties, "pal_image")
		delete(additionalProperties, "pal_name")
		delete(additionalProperties, "places")
		delete(additionalProperties, "race_id")
		delete(additionalProperties, "rank")
		delete(additionalProperties, "result")
		delete(additionalProperties, "token_id")
		delete(additionalProperties, "transaction_code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePalActivityLogDetails struct {
	value *PalActivityLogDetails
	isSet bool
}

func (v NullablePalActivityLogDetails) Get() *PalActivityLogDetails {
	return v.value
}

func (v *NullablePalActivityLogDetails) Set(val *PalActivityLogDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePalActivityLogDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePalActivityLogDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePalActivityLogDetails(val *PalActivityLogDetails) *NullablePalActivityLogDetails {
	return &NullablePalActivityLogDetails{value: val, isSet: true}
}

func (v NullablePalActivityLogDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePalActivityLogDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


