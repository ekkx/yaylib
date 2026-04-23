
// Code generated from the Yay! API schema; DO NOT EDIT.

package gen

import (
	"encoding/json"
)

// checks if the Config type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Config{}

// Config struct for Config
type Config struct {
	AdTesterUserIds NullableString `json:"ad_tester_user_ids,omitempty"`
	AndroidNewsVersion NullableString `json:"android_news_version,omitempty"`
	Dapps NullableString `json:"dapps,omitempty"`
	IsChatWebsocketEnabled NullableString `json:"is_chat_websocket_enabled,omitempty"`
	IsDirectVipPurchaseEnabled NullableString `json:"is_direct_vip_purchase_enabled,omitempty"`
	IsGiftFeaturesEnabled NullableString `json:"is_gift_features_enabled,omitempty"`
	IsIdCardAndPhoneVerificationCheckForReviewEnabled NullableString `json:"is_id_card_and_phone_verification_check_for_review_enabled,omitempty"`
	IsIdCardCheckRequiredBlockerViewEnabled NullableString `json:"is_id_card_check_required_blocker_view_enabled,omitempty"`
	IsMaintenanced NullableString `json:"is_maintenanced,omitempty"`
	IsPhoneVerificationRequiredBlockerViewEnabled NullableString `json:"is_phone_verification_required_blocker_view_enabled,omitempty"`
	IsSpeedTestEnabled NullableString `json:"is_speed_test_enabled,omitempty"`
	IsSslPinningEnabled NullableString `json:"is_ssl_pinning_enabled,omitempty"`
	IsStarEnabled NullableString `json:"is_star_enabled,omitempty"`
	LatestAndroidAppVersion NullableString `json:"latest_android_app_version,omitempty"`
	LineOfficialAccountId NullableString `json:"line_official_account_id,omitempty"`
	LocalizedAndroidNewsTitle NullableString `json:"localized_android_news_title,omitempty"`
	LocalizedAndroidNewsUrl NullableString `json:"localized_android_news_url,omitempty"`
	LocalizedMaintenanceUrl NullableString `json:"localized_maintenance_url,omitempty"`
	LocalizedNewsTitle NullableString `json:"localized_news_title,omitempty"`
	LocalizedNewsUrl NullableString `json:"localized_news_url,omitempty"`
	MaintenanceFeatures NullableString `json:"maintenance_features,omitempty"`
	MaxImageFrameCount NullableString `json:"max_image_frame_count,omitempty"`
	MinimumAndroidAppVersionRequired NullableString `json:"minimum_android_app_version_required,omitempty"`
	MinimumAndroidVersionSupported NullableString `json:"minimum_android_version_supported,omitempty"`
	NewsVersion NullableString `json:"news_version,omitempty"`
	NftInfos NullableString `json:"nft_infos,omitempty"`
	PromotionStickerPackIds NullableString `json:"promotion_sticker_pack_ids,omitempty"`
	ShouldAppendUserIdToNewsUrl NullableString `json:"should_append_user_id_to_news_url,omitempty"`
	SupportEmailAddress NullableString `json:"support_email_address,omitempty"`
	TokenInfos NullableString `json:"token_infos,omitempty"`
	UseRandomMessageRefreshRate NullableString `json:"use_random_message_refresh_rate,omitempty"`
	Web3IsMaintenanced NullableString `json:"web3_is_maintenanced,omitempty"`
	Web3LocalizedMaintenanceUrl NullableString `json:"web3_localized_maintenance_url,omitempty"`
	Web3MaintenanceFeatures NullableString `json:"web3_maintenance_features,omitempty"`
	Web3MaintenanceUrl NullableString `json:"web3_maintenance_url,omitempty"`
	XYayGlobalId NullableString `json:"x_yay_global_id,omitempty"`
	XYayJpId NullableString `json:"x_yay_jp_id,omitempty"`
	XYayUpdateId NullableString `json:"x_yay_update_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Config Config

// NewConfig instantiates a new Config object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfig() *Config {
	this := Config{}
	return &this
}

// NewConfigWithDefaults instantiates a new Config object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigWithDefaults() *Config {
	this := Config{}
	return &this
}

// GetAdTesterUserIds returns the AdTesterUserIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetAdTesterUserIds() string {
	if o == nil || IsNil(o.AdTesterUserIds.Get()) {
		var ret string
		return ret
	}
	return *o.AdTesterUserIds.Get()
}

// GetAdTesterUserIdsOk returns a tuple with the AdTesterUserIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetAdTesterUserIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AdTesterUserIds.Get(), o.AdTesterUserIds.IsSet()
}

// HasAdTesterUserIds returns a boolean if a field has been set.
func (o *Config) HasAdTesterUserIds() bool {
	if o != nil && o.AdTesterUserIds.IsSet() {
		return true
	}

	return false
}

// SetAdTesterUserIds gets a reference to the given NullableString and assigns it to the AdTesterUserIds field.
func (o *Config) SetAdTesterUserIds(v string) {
	o.AdTesterUserIds.Set(&v)
}
// SetAdTesterUserIdsNil sets the value for AdTesterUserIds to be an explicit nil
func (o *Config) SetAdTesterUserIdsNil() {
	o.AdTesterUserIds.Set(nil)
}

// UnsetAdTesterUserIds ensures that no value is present for AdTesterUserIds, not even an explicit nil
func (o *Config) UnsetAdTesterUserIds() {
	o.AdTesterUserIds.Unset()
}

// GetAndroidNewsVersion returns the AndroidNewsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetAndroidNewsVersion() string {
	if o == nil || IsNil(o.AndroidNewsVersion.Get()) {
		var ret string
		return ret
	}
	return *o.AndroidNewsVersion.Get()
}

// GetAndroidNewsVersionOk returns a tuple with the AndroidNewsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetAndroidNewsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AndroidNewsVersion.Get(), o.AndroidNewsVersion.IsSet()
}

// HasAndroidNewsVersion returns a boolean if a field has been set.
func (o *Config) HasAndroidNewsVersion() bool {
	if o != nil && o.AndroidNewsVersion.IsSet() {
		return true
	}

	return false
}

// SetAndroidNewsVersion gets a reference to the given NullableString and assigns it to the AndroidNewsVersion field.
func (o *Config) SetAndroidNewsVersion(v string) {
	o.AndroidNewsVersion.Set(&v)
}
// SetAndroidNewsVersionNil sets the value for AndroidNewsVersion to be an explicit nil
func (o *Config) SetAndroidNewsVersionNil() {
	o.AndroidNewsVersion.Set(nil)
}

// UnsetAndroidNewsVersion ensures that no value is present for AndroidNewsVersion, not even an explicit nil
func (o *Config) UnsetAndroidNewsVersion() {
	o.AndroidNewsVersion.Unset()
}

// GetDapps returns the Dapps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetDapps() string {
	if o == nil || IsNil(o.Dapps.Get()) {
		var ret string
		return ret
	}
	return *o.Dapps.Get()
}

// GetDappsOk returns a tuple with the Dapps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetDappsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dapps.Get(), o.Dapps.IsSet()
}

// HasDapps returns a boolean if a field has been set.
func (o *Config) HasDapps() bool {
	if o != nil && o.Dapps.IsSet() {
		return true
	}

	return false
}

// SetDapps gets a reference to the given NullableString and assigns it to the Dapps field.
func (o *Config) SetDapps(v string) {
	o.Dapps.Set(&v)
}
// SetDappsNil sets the value for Dapps to be an explicit nil
func (o *Config) SetDappsNil() {
	o.Dapps.Set(nil)
}

// UnsetDapps ensures that no value is present for Dapps, not even an explicit nil
func (o *Config) UnsetDapps() {
	o.Dapps.Unset()
}

// GetIsChatWebsocketEnabled returns the IsChatWebsocketEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsChatWebsocketEnabled() string {
	if o == nil || IsNil(o.IsChatWebsocketEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsChatWebsocketEnabled.Get()
}

// GetIsChatWebsocketEnabledOk returns a tuple with the IsChatWebsocketEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsChatWebsocketEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsChatWebsocketEnabled.Get(), o.IsChatWebsocketEnabled.IsSet()
}

// HasIsChatWebsocketEnabled returns a boolean if a field has been set.
func (o *Config) HasIsChatWebsocketEnabled() bool {
	if o != nil && o.IsChatWebsocketEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsChatWebsocketEnabled gets a reference to the given NullableString and assigns it to the IsChatWebsocketEnabled field.
func (o *Config) SetIsChatWebsocketEnabled(v string) {
	o.IsChatWebsocketEnabled.Set(&v)
}
// SetIsChatWebsocketEnabledNil sets the value for IsChatWebsocketEnabled to be an explicit nil
func (o *Config) SetIsChatWebsocketEnabledNil() {
	o.IsChatWebsocketEnabled.Set(nil)
}

// UnsetIsChatWebsocketEnabled ensures that no value is present for IsChatWebsocketEnabled, not even an explicit nil
func (o *Config) UnsetIsChatWebsocketEnabled() {
	o.IsChatWebsocketEnabled.Unset()
}

// GetIsDirectVipPurchaseEnabled returns the IsDirectVipPurchaseEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsDirectVipPurchaseEnabled() string {
	if o == nil || IsNil(o.IsDirectVipPurchaseEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsDirectVipPurchaseEnabled.Get()
}

// GetIsDirectVipPurchaseEnabledOk returns a tuple with the IsDirectVipPurchaseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsDirectVipPurchaseEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsDirectVipPurchaseEnabled.Get(), o.IsDirectVipPurchaseEnabled.IsSet()
}

// HasIsDirectVipPurchaseEnabled returns a boolean if a field has been set.
func (o *Config) HasIsDirectVipPurchaseEnabled() bool {
	if o != nil && o.IsDirectVipPurchaseEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsDirectVipPurchaseEnabled gets a reference to the given NullableString and assigns it to the IsDirectVipPurchaseEnabled field.
func (o *Config) SetIsDirectVipPurchaseEnabled(v string) {
	o.IsDirectVipPurchaseEnabled.Set(&v)
}
// SetIsDirectVipPurchaseEnabledNil sets the value for IsDirectVipPurchaseEnabled to be an explicit nil
func (o *Config) SetIsDirectVipPurchaseEnabledNil() {
	o.IsDirectVipPurchaseEnabled.Set(nil)
}

// UnsetIsDirectVipPurchaseEnabled ensures that no value is present for IsDirectVipPurchaseEnabled, not even an explicit nil
func (o *Config) UnsetIsDirectVipPurchaseEnabled() {
	o.IsDirectVipPurchaseEnabled.Unset()
}

// GetIsGiftFeaturesEnabled returns the IsGiftFeaturesEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsGiftFeaturesEnabled() string {
	if o == nil || IsNil(o.IsGiftFeaturesEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsGiftFeaturesEnabled.Get()
}

// GetIsGiftFeaturesEnabledOk returns a tuple with the IsGiftFeaturesEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsGiftFeaturesEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsGiftFeaturesEnabled.Get(), o.IsGiftFeaturesEnabled.IsSet()
}

// HasIsGiftFeaturesEnabled returns a boolean if a field has been set.
func (o *Config) HasIsGiftFeaturesEnabled() bool {
	if o != nil && o.IsGiftFeaturesEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsGiftFeaturesEnabled gets a reference to the given NullableString and assigns it to the IsGiftFeaturesEnabled field.
func (o *Config) SetIsGiftFeaturesEnabled(v string) {
	o.IsGiftFeaturesEnabled.Set(&v)
}
// SetIsGiftFeaturesEnabledNil sets the value for IsGiftFeaturesEnabled to be an explicit nil
func (o *Config) SetIsGiftFeaturesEnabledNil() {
	o.IsGiftFeaturesEnabled.Set(nil)
}

// UnsetIsGiftFeaturesEnabled ensures that no value is present for IsGiftFeaturesEnabled, not even an explicit nil
func (o *Config) UnsetIsGiftFeaturesEnabled() {
	o.IsGiftFeaturesEnabled.Unset()
}

// GetIsIdCardAndPhoneVerificationCheckForReviewEnabled returns the IsIdCardAndPhoneVerificationCheckForReviewEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsIdCardAndPhoneVerificationCheckForReviewEnabled() string {
	if o == nil || IsNil(o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Get()
}

// GetIsIdCardAndPhoneVerificationCheckForReviewEnabledOk returns a tuple with the IsIdCardAndPhoneVerificationCheckForReviewEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsIdCardAndPhoneVerificationCheckForReviewEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Get(), o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.IsSet()
}

// HasIsIdCardAndPhoneVerificationCheckForReviewEnabled returns a boolean if a field has been set.
func (o *Config) HasIsIdCardAndPhoneVerificationCheckForReviewEnabled() bool {
	if o != nil && o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsIdCardAndPhoneVerificationCheckForReviewEnabled gets a reference to the given NullableString and assigns it to the IsIdCardAndPhoneVerificationCheckForReviewEnabled field.
func (o *Config) SetIsIdCardAndPhoneVerificationCheckForReviewEnabled(v string) {
	o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Set(&v)
}
// SetIsIdCardAndPhoneVerificationCheckForReviewEnabledNil sets the value for IsIdCardAndPhoneVerificationCheckForReviewEnabled to be an explicit nil
func (o *Config) SetIsIdCardAndPhoneVerificationCheckForReviewEnabledNil() {
	o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Set(nil)
}

// UnsetIsIdCardAndPhoneVerificationCheckForReviewEnabled ensures that no value is present for IsIdCardAndPhoneVerificationCheckForReviewEnabled, not even an explicit nil
func (o *Config) UnsetIsIdCardAndPhoneVerificationCheckForReviewEnabled() {
	o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Unset()
}

// GetIsIdCardCheckRequiredBlockerViewEnabled returns the IsIdCardCheckRequiredBlockerViewEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsIdCardCheckRequiredBlockerViewEnabled() string {
	if o == nil || IsNil(o.IsIdCardCheckRequiredBlockerViewEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsIdCardCheckRequiredBlockerViewEnabled.Get()
}

// GetIsIdCardCheckRequiredBlockerViewEnabledOk returns a tuple with the IsIdCardCheckRequiredBlockerViewEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsIdCardCheckRequiredBlockerViewEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsIdCardCheckRequiredBlockerViewEnabled.Get(), o.IsIdCardCheckRequiredBlockerViewEnabled.IsSet()
}

// HasIsIdCardCheckRequiredBlockerViewEnabled returns a boolean if a field has been set.
func (o *Config) HasIsIdCardCheckRequiredBlockerViewEnabled() bool {
	if o != nil && o.IsIdCardCheckRequiredBlockerViewEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsIdCardCheckRequiredBlockerViewEnabled gets a reference to the given NullableString and assigns it to the IsIdCardCheckRequiredBlockerViewEnabled field.
func (o *Config) SetIsIdCardCheckRequiredBlockerViewEnabled(v string) {
	o.IsIdCardCheckRequiredBlockerViewEnabled.Set(&v)
}
// SetIsIdCardCheckRequiredBlockerViewEnabledNil sets the value for IsIdCardCheckRequiredBlockerViewEnabled to be an explicit nil
func (o *Config) SetIsIdCardCheckRequiredBlockerViewEnabledNil() {
	o.IsIdCardCheckRequiredBlockerViewEnabled.Set(nil)
}

// UnsetIsIdCardCheckRequiredBlockerViewEnabled ensures that no value is present for IsIdCardCheckRequiredBlockerViewEnabled, not even an explicit nil
func (o *Config) UnsetIsIdCardCheckRequiredBlockerViewEnabled() {
	o.IsIdCardCheckRequiredBlockerViewEnabled.Unset()
}

// GetIsMaintenanced returns the IsMaintenanced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsMaintenanced() string {
	if o == nil || IsNil(o.IsMaintenanced.Get()) {
		var ret string
		return ret
	}
	return *o.IsMaintenanced.Get()
}

// GetIsMaintenancedOk returns a tuple with the IsMaintenanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsMaintenancedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsMaintenanced.Get(), o.IsMaintenanced.IsSet()
}

// HasIsMaintenanced returns a boolean if a field has been set.
func (o *Config) HasIsMaintenanced() bool {
	if o != nil && o.IsMaintenanced.IsSet() {
		return true
	}

	return false
}

// SetIsMaintenanced gets a reference to the given NullableString and assigns it to the IsMaintenanced field.
func (o *Config) SetIsMaintenanced(v string) {
	o.IsMaintenanced.Set(&v)
}
// SetIsMaintenancedNil sets the value for IsMaintenanced to be an explicit nil
func (o *Config) SetIsMaintenancedNil() {
	o.IsMaintenanced.Set(nil)
}

// UnsetIsMaintenanced ensures that no value is present for IsMaintenanced, not even an explicit nil
func (o *Config) UnsetIsMaintenanced() {
	o.IsMaintenanced.Unset()
}

// GetIsPhoneVerificationRequiredBlockerViewEnabled returns the IsPhoneVerificationRequiredBlockerViewEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsPhoneVerificationRequiredBlockerViewEnabled() string {
	if o == nil || IsNil(o.IsPhoneVerificationRequiredBlockerViewEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsPhoneVerificationRequiredBlockerViewEnabled.Get()
}

// GetIsPhoneVerificationRequiredBlockerViewEnabledOk returns a tuple with the IsPhoneVerificationRequiredBlockerViewEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsPhoneVerificationRequiredBlockerViewEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsPhoneVerificationRequiredBlockerViewEnabled.Get(), o.IsPhoneVerificationRequiredBlockerViewEnabled.IsSet()
}

// HasIsPhoneVerificationRequiredBlockerViewEnabled returns a boolean if a field has been set.
func (o *Config) HasIsPhoneVerificationRequiredBlockerViewEnabled() bool {
	if o != nil && o.IsPhoneVerificationRequiredBlockerViewEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsPhoneVerificationRequiredBlockerViewEnabled gets a reference to the given NullableString and assigns it to the IsPhoneVerificationRequiredBlockerViewEnabled field.
func (o *Config) SetIsPhoneVerificationRequiredBlockerViewEnabled(v string) {
	o.IsPhoneVerificationRequiredBlockerViewEnabled.Set(&v)
}
// SetIsPhoneVerificationRequiredBlockerViewEnabledNil sets the value for IsPhoneVerificationRequiredBlockerViewEnabled to be an explicit nil
func (o *Config) SetIsPhoneVerificationRequiredBlockerViewEnabledNil() {
	o.IsPhoneVerificationRequiredBlockerViewEnabled.Set(nil)
}

// UnsetIsPhoneVerificationRequiredBlockerViewEnabled ensures that no value is present for IsPhoneVerificationRequiredBlockerViewEnabled, not even an explicit nil
func (o *Config) UnsetIsPhoneVerificationRequiredBlockerViewEnabled() {
	o.IsPhoneVerificationRequiredBlockerViewEnabled.Unset()
}

// GetIsSpeedTestEnabled returns the IsSpeedTestEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsSpeedTestEnabled() string {
	if o == nil || IsNil(o.IsSpeedTestEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsSpeedTestEnabled.Get()
}

// GetIsSpeedTestEnabledOk returns a tuple with the IsSpeedTestEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsSpeedTestEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSpeedTestEnabled.Get(), o.IsSpeedTestEnabled.IsSet()
}

// HasIsSpeedTestEnabled returns a boolean if a field has been set.
func (o *Config) HasIsSpeedTestEnabled() bool {
	if o != nil && o.IsSpeedTestEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsSpeedTestEnabled gets a reference to the given NullableString and assigns it to the IsSpeedTestEnabled field.
func (o *Config) SetIsSpeedTestEnabled(v string) {
	o.IsSpeedTestEnabled.Set(&v)
}
// SetIsSpeedTestEnabledNil sets the value for IsSpeedTestEnabled to be an explicit nil
func (o *Config) SetIsSpeedTestEnabledNil() {
	o.IsSpeedTestEnabled.Set(nil)
}

// UnsetIsSpeedTestEnabled ensures that no value is present for IsSpeedTestEnabled, not even an explicit nil
func (o *Config) UnsetIsSpeedTestEnabled() {
	o.IsSpeedTestEnabled.Unset()
}

// GetIsSslPinningEnabled returns the IsSslPinningEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsSslPinningEnabled() string {
	if o == nil || IsNil(o.IsSslPinningEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsSslPinningEnabled.Get()
}

// GetIsSslPinningEnabledOk returns a tuple with the IsSslPinningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsSslPinningEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsSslPinningEnabled.Get(), o.IsSslPinningEnabled.IsSet()
}

// HasIsSslPinningEnabled returns a boolean if a field has been set.
func (o *Config) HasIsSslPinningEnabled() bool {
	if o != nil && o.IsSslPinningEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsSslPinningEnabled gets a reference to the given NullableString and assigns it to the IsSslPinningEnabled field.
func (o *Config) SetIsSslPinningEnabled(v string) {
	o.IsSslPinningEnabled.Set(&v)
}
// SetIsSslPinningEnabledNil sets the value for IsSslPinningEnabled to be an explicit nil
func (o *Config) SetIsSslPinningEnabledNil() {
	o.IsSslPinningEnabled.Set(nil)
}

// UnsetIsSslPinningEnabled ensures that no value is present for IsSslPinningEnabled, not even an explicit nil
func (o *Config) UnsetIsSslPinningEnabled() {
	o.IsSslPinningEnabled.Unset()
}

// GetIsStarEnabled returns the IsStarEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetIsStarEnabled() string {
	if o == nil || IsNil(o.IsStarEnabled.Get()) {
		var ret string
		return ret
	}
	return *o.IsStarEnabled.Get()
}

// GetIsStarEnabledOk returns a tuple with the IsStarEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetIsStarEnabledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IsStarEnabled.Get(), o.IsStarEnabled.IsSet()
}

// HasIsStarEnabled returns a boolean if a field has been set.
func (o *Config) HasIsStarEnabled() bool {
	if o != nil && o.IsStarEnabled.IsSet() {
		return true
	}

	return false
}

// SetIsStarEnabled gets a reference to the given NullableString and assigns it to the IsStarEnabled field.
func (o *Config) SetIsStarEnabled(v string) {
	o.IsStarEnabled.Set(&v)
}
// SetIsStarEnabledNil sets the value for IsStarEnabled to be an explicit nil
func (o *Config) SetIsStarEnabledNil() {
	o.IsStarEnabled.Set(nil)
}

// UnsetIsStarEnabled ensures that no value is present for IsStarEnabled, not even an explicit nil
func (o *Config) UnsetIsStarEnabled() {
	o.IsStarEnabled.Unset()
}

// GetLatestAndroidAppVersion returns the LatestAndroidAppVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLatestAndroidAppVersion() string {
	if o == nil || IsNil(o.LatestAndroidAppVersion.Get()) {
		var ret string
		return ret
	}
	return *o.LatestAndroidAppVersion.Get()
}

// GetLatestAndroidAppVersionOk returns a tuple with the LatestAndroidAppVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLatestAndroidAppVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LatestAndroidAppVersion.Get(), o.LatestAndroidAppVersion.IsSet()
}

// HasLatestAndroidAppVersion returns a boolean if a field has been set.
func (o *Config) HasLatestAndroidAppVersion() bool {
	if o != nil && o.LatestAndroidAppVersion.IsSet() {
		return true
	}

	return false
}

// SetLatestAndroidAppVersion gets a reference to the given NullableString and assigns it to the LatestAndroidAppVersion field.
func (o *Config) SetLatestAndroidAppVersion(v string) {
	o.LatestAndroidAppVersion.Set(&v)
}
// SetLatestAndroidAppVersionNil sets the value for LatestAndroidAppVersion to be an explicit nil
func (o *Config) SetLatestAndroidAppVersionNil() {
	o.LatestAndroidAppVersion.Set(nil)
}

// UnsetLatestAndroidAppVersion ensures that no value is present for LatestAndroidAppVersion, not even an explicit nil
func (o *Config) UnsetLatestAndroidAppVersion() {
	o.LatestAndroidAppVersion.Unset()
}

// GetLineOfficialAccountId returns the LineOfficialAccountId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLineOfficialAccountId() string {
	if o == nil || IsNil(o.LineOfficialAccountId.Get()) {
		var ret string
		return ret
	}
	return *o.LineOfficialAccountId.Get()
}

// GetLineOfficialAccountIdOk returns a tuple with the LineOfficialAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLineOfficialAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineOfficialAccountId.Get(), o.LineOfficialAccountId.IsSet()
}

// HasLineOfficialAccountId returns a boolean if a field has been set.
func (o *Config) HasLineOfficialAccountId() bool {
	if o != nil && o.LineOfficialAccountId.IsSet() {
		return true
	}

	return false
}

// SetLineOfficialAccountId gets a reference to the given NullableString and assigns it to the LineOfficialAccountId field.
func (o *Config) SetLineOfficialAccountId(v string) {
	o.LineOfficialAccountId.Set(&v)
}
// SetLineOfficialAccountIdNil sets the value for LineOfficialAccountId to be an explicit nil
func (o *Config) SetLineOfficialAccountIdNil() {
	o.LineOfficialAccountId.Set(nil)
}

// UnsetLineOfficialAccountId ensures that no value is present for LineOfficialAccountId, not even an explicit nil
func (o *Config) UnsetLineOfficialAccountId() {
	o.LineOfficialAccountId.Unset()
}

// GetLocalizedAndroidNewsTitle returns the LocalizedAndroidNewsTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLocalizedAndroidNewsTitle() string {
	if o == nil || IsNil(o.LocalizedAndroidNewsTitle.Get()) {
		var ret string
		return ret
	}
	return *o.LocalizedAndroidNewsTitle.Get()
}

// GetLocalizedAndroidNewsTitleOk returns a tuple with the LocalizedAndroidNewsTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLocalizedAndroidNewsTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedAndroidNewsTitle.Get(), o.LocalizedAndroidNewsTitle.IsSet()
}

// HasLocalizedAndroidNewsTitle returns a boolean if a field has been set.
func (o *Config) HasLocalizedAndroidNewsTitle() bool {
	if o != nil && o.LocalizedAndroidNewsTitle.IsSet() {
		return true
	}

	return false
}

// SetLocalizedAndroidNewsTitle gets a reference to the given NullableString and assigns it to the LocalizedAndroidNewsTitle field.
func (o *Config) SetLocalizedAndroidNewsTitle(v string) {
	o.LocalizedAndroidNewsTitle.Set(&v)
}
// SetLocalizedAndroidNewsTitleNil sets the value for LocalizedAndroidNewsTitle to be an explicit nil
func (o *Config) SetLocalizedAndroidNewsTitleNil() {
	o.LocalizedAndroidNewsTitle.Set(nil)
}

// UnsetLocalizedAndroidNewsTitle ensures that no value is present for LocalizedAndroidNewsTitle, not even an explicit nil
func (o *Config) UnsetLocalizedAndroidNewsTitle() {
	o.LocalizedAndroidNewsTitle.Unset()
}

// GetLocalizedAndroidNewsUrl returns the LocalizedAndroidNewsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLocalizedAndroidNewsUrl() string {
	if o == nil || IsNil(o.LocalizedAndroidNewsUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LocalizedAndroidNewsUrl.Get()
}

// GetLocalizedAndroidNewsUrlOk returns a tuple with the LocalizedAndroidNewsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLocalizedAndroidNewsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedAndroidNewsUrl.Get(), o.LocalizedAndroidNewsUrl.IsSet()
}

// HasLocalizedAndroidNewsUrl returns a boolean if a field has been set.
func (o *Config) HasLocalizedAndroidNewsUrl() bool {
	if o != nil && o.LocalizedAndroidNewsUrl.IsSet() {
		return true
	}

	return false
}

// SetLocalizedAndroidNewsUrl gets a reference to the given NullableString and assigns it to the LocalizedAndroidNewsUrl field.
func (o *Config) SetLocalizedAndroidNewsUrl(v string) {
	o.LocalizedAndroidNewsUrl.Set(&v)
}
// SetLocalizedAndroidNewsUrlNil sets the value for LocalizedAndroidNewsUrl to be an explicit nil
func (o *Config) SetLocalizedAndroidNewsUrlNil() {
	o.LocalizedAndroidNewsUrl.Set(nil)
}

// UnsetLocalizedAndroidNewsUrl ensures that no value is present for LocalizedAndroidNewsUrl, not even an explicit nil
func (o *Config) UnsetLocalizedAndroidNewsUrl() {
	o.LocalizedAndroidNewsUrl.Unset()
}

// GetLocalizedMaintenanceUrl returns the LocalizedMaintenanceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLocalizedMaintenanceUrl() string {
	if o == nil || IsNil(o.LocalizedMaintenanceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LocalizedMaintenanceUrl.Get()
}

// GetLocalizedMaintenanceUrlOk returns a tuple with the LocalizedMaintenanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLocalizedMaintenanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedMaintenanceUrl.Get(), o.LocalizedMaintenanceUrl.IsSet()
}

// HasLocalizedMaintenanceUrl returns a boolean if a field has been set.
func (o *Config) HasLocalizedMaintenanceUrl() bool {
	if o != nil && o.LocalizedMaintenanceUrl.IsSet() {
		return true
	}

	return false
}

// SetLocalizedMaintenanceUrl gets a reference to the given NullableString and assigns it to the LocalizedMaintenanceUrl field.
func (o *Config) SetLocalizedMaintenanceUrl(v string) {
	o.LocalizedMaintenanceUrl.Set(&v)
}
// SetLocalizedMaintenanceUrlNil sets the value for LocalizedMaintenanceUrl to be an explicit nil
func (o *Config) SetLocalizedMaintenanceUrlNil() {
	o.LocalizedMaintenanceUrl.Set(nil)
}

// UnsetLocalizedMaintenanceUrl ensures that no value is present for LocalizedMaintenanceUrl, not even an explicit nil
func (o *Config) UnsetLocalizedMaintenanceUrl() {
	o.LocalizedMaintenanceUrl.Unset()
}

// GetLocalizedNewsTitle returns the LocalizedNewsTitle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLocalizedNewsTitle() string {
	if o == nil || IsNil(o.LocalizedNewsTitle.Get()) {
		var ret string
		return ret
	}
	return *o.LocalizedNewsTitle.Get()
}

// GetLocalizedNewsTitleOk returns a tuple with the LocalizedNewsTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLocalizedNewsTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedNewsTitle.Get(), o.LocalizedNewsTitle.IsSet()
}

// HasLocalizedNewsTitle returns a boolean if a field has been set.
func (o *Config) HasLocalizedNewsTitle() bool {
	if o != nil && o.LocalizedNewsTitle.IsSet() {
		return true
	}

	return false
}

// SetLocalizedNewsTitle gets a reference to the given NullableString and assigns it to the LocalizedNewsTitle field.
func (o *Config) SetLocalizedNewsTitle(v string) {
	o.LocalizedNewsTitle.Set(&v)
}
// SetLocalizedNewsTitleNil sets the value for LocalizedNewsTitle to be an explicit nil
func (o *Config) SetLocalizedNewsTitleNil() {
	o.LocalizedNewsTitle.Set(nil)
}

// UnsetLocalizedNewsTitle ensures that no value is present for LocalizedNewsTitle, not even an explicit nil
func (o *Config) UnsetLocalizedNewsTitle() {
	o.LocalizedNewsTitle.Unset()
}

// GetLocalizedNewsUrl returns the LocalizedNewsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetLocalizedNewsUrl() string {
	if o == nil || IsNil(o.LocalizedNewsUrl.Get()) {
		var ret string
		return ret
	}
	return *o.LocalizedNewsUrl.Get()
}

// GetLocalizedNewsUrlOk returns a tuple with the LocalizedNewsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetLocalizedNewsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocalizedNewsUrl.Get(), o.LocalizedNewsUrl.IsSet()
}

// HasLocalizedNewsUrl returns a boolean if a field has been set.
func (o *Config) HasLocalizedNewsUrl() bool {
	if o != nil && o.LocalizedNewsUrl.IsSet() {
		return true
	}

	return false
}

// SetLocalizedNewsUrl gets a reference to the given NullableString and assigns it to the LocalizedNewsUrl field.
func (o *Config) SetLocalizedNewsUrl(v string) {
	o.LocalizedNewsUrl.Set(&v)
}
// SetLocalizedNewsUrlNil sets the value for LocalizedNewsUrl to be an explicit nil
func (o *Config) SetLocalizedNewsUrlNil() {
	o.LocalizedNewsUrl.Set(nil)
}

// UnsetLocalizedNewsUrl ensures that no value is present for LocalizedNewsUrl, not even an explicit nil
func (o *Config) UnsetLocalizedNewsUrl() {
	o.LocalizedNewsUrl.Unset()
}

// GetMaintenanceFeatures returns the MaintenanceFeatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetMaintenanceFeatures() string {
	if o == nil || IsNil(o.MaintenanceFeatures.Get()) {
		var ret string
		return ret
	}
	return *o.MaintenanceFeatures.Get()
}

// GetMaintenanceFeaturesOk returns a tuple with the MaintenanceFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetMaintenanceFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaintenanceFeatures.Get(), o.MaintenanceFeatures.IsSet()
}

// HasMaintenanceFeatures returns a boolean if a field has been set.
func (o *Config) HasMaintenanceFeatures() bool {
	if o != nil && o.MaintenanceFeatures.IsSet() {
		return true
	}

	return false
}

// SetMaintenanceFeatures gets a reference to the given NullableString and assigns it to the MaintenanceFeatures field.
func (o *Config) SetMaintenanceFeatures(v string) {
	o.MaintenanceFeatures.Set(&v)
}
// SetMaintenanceFeaturesNil sets the value for MaintenanceFeatures to be an explicit nil
func (o *Config) SetMaintenanceFeaturesNil() {
	o.MaintenanceFeatures.Set(nil)
}

// UnsetMaintenanceFeatures ensures that no value is present for MaintenanceFeatures, not even an explicit nil
func (o *Config) UnsetMaintenanceFeatures() {
	o.MaintenanceFeatures.Unset()
}

// GetMaxImageFrameCount returns the MaxImageFrameCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetMaxImageFrameCount() string {
	if o == nil || IsNil(o.MaxImageFrameCount.Get()) {
		var ret string
		return ret
	}
	return *o.MaxImageFrameCount.Get()
}

// GetMaxImageFrameCountOk returns a tuple with the MaxImageFrameCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetMaxImageFrameCountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxImageFrameCount.Get(), o.MaxImageFrameCount.IsSet()
}

// HasMaxImageFrameCount returns a boolean if a field has been set.
func (o *Config) HasMaxImageFrameCount() bool {
	if o != nil && o.MaxImageFrameCount.IsSet() {
		return true
	}

	return false
}

// SetMaxImageFrameCount gets a reference to the given NullableString and assigns it to the MaxImageFrameCount field.
func (o *Config) SetMaxImageFrameCount(v string) {
	o.MaxImageFrameCount.Set(&v)
}
// SetMaxImageFrameCountNil sets the value for MaxImageFrameCount to be an explicit nil
func (o *Config) SetMaxImageFrameCountNil() {
	o.MaxImageFrameCount.Set(nil)
}

// UnsetMaxImageFrameCount ensures that no value is present for MaxImageFrameCount, not even an explicit nil
func (o *Config) UnsetMaxImageFrameCount() {
	o.MaxImageFrameCount.Unset()
}

// GetMinimumAndroidAppVersionRequired returns the MinimumAndroidAppVersionRequired field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetMinimumAndroidAppVersionRequired() string {
	if o == nil || IsNil(o.MinimumAndroidAppVersionRequired.Get()) {
		var ret string
		return ret
	}
	return *o.MinimumAndroidAppVersionRequired.Get()
}

// GetMinimumAndroidAppVersionRequiredOk returns a tuple with the MinimumAndroidAppVersionRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetMinimumAndroidAppVersionRequiredOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumAndroidAppVersionRequired.Get(), o.MinimumAndroidAppVersionRequired.IsSet()
}

// HasMinimumAndroidAppVersionRequired returns a boolean if a field has been set.
func (o *Config) HasMinimumAndroidAppVersionRequired() bool {
	if o != nil && o.MinimumAndroidAppVersionRequired.IsSet() {
		return true
	}

	return false
}

// SetMinimumAndroidAppVersionRequired gets a reference to the given NullableString and assigns it to the MinimumAndroidAppVersionRequired field.
func (o *Config) SetMinimumAndroidAppVersionRequired(v string) {
	o.MinimumAndroidAppVersionRequired.Set(&v)
}
// SetMinimumAndroidAppVersionRequiredNil sets the value for MinimumAndroidAppVersionRequired to be an explicit nil
func (o *Config) SetMinimumAndroidAppVersionRequiredNil() {
	o.MinimumAndroidAppVersionRequired.Set(nil)
}

// UnsetMinimumAndroidAppVersionRequired ensures that no value is present for MinimumAndroidAppVersionRequired, not even an explicit nil
func (o *Config) UnsetMinimumAndroidAppVersionRequired() {
	o.MinimumAndroidAppVersionRequired.Unset()
}

// GetMinimumAndroidVersionSupported returns the MinimumAndroidVersionSupported field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetMinimumAndroidVersionSupported() string {
	if o == nil || IsNil(o.MinimumAndroidVersionSupported.Get()) {
		var ret string
		return ret
	}
	return *o.MinimumAndroidVersionSupported.Get()
}

// GetMinimumAndroidVersionSupportedOk returns a tuple with the MinimumAndroidVersionSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetMinimumAndroidVersionSupportedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MinimumAndroidVersionSupported.Get(), o.MinimumAndroidVersionSupported.IsSet()
}

// HasMinimumAndroidVersionSupported returns a boolean if a field has been set.
func (o *Config) HasMinimumAndroidVersionSupported() bool {
	if o != nil && o.MinimumAndroidVersionSupported.IsSet() {
		return true
	}

	return false
}

// SetMinimumAndroidVersionSupported gets a reference to the given NullableString and assigns it to the MinimumAndroidVersionSupported field.
func (o *Config) SetMinimumAndroidVersionSupported(v string) {
	o.MinimumAndroidVersionSupported.Set(&v)
}
// SetMinimumAndroidVersionSupportedNil sets the value for MinimumAndroidVersionSupported to be an explicit nil
func (o *Config) SetMinimumAndroidVersionSupportedNil() {
	o.MinimumAndroidVersionSupported.Set(nil)
}

// UnsetMinimumAndroidVersionSupported ensures that no value is present for MinimumAndroidVersionSupported, not even an explicit nil
func (o *Config) UnsetMinimumAndroidVersionSupported() {
	o.MinimumAndroidVersionSupported.Unset()
}

// GetNewsVersion returns the NewsVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetNewsVersion() string {
	if o == nil || IsNil(o.NewsVersion.Get()) {
		var ret string
		return ret
	}
	return *o.NewsVersion.Get()
}

// GetNewsVersionOk returns a tuple with the NewsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetNewsVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NewsVersion.Get(), o.NewsVersion.IsSet()
}

// HasNewsVersion returns a boolean if a field has been set.
func (o *Config) HasNewsVersion() bool {
	if o != nil && o.NewsVersion.IsSet() {
		return true
	}

	return false
}

// SetNewsVersion gets a reference to the given NullableString and assigns it to the NewsVersion field.
func (o *Config) SetNewsVersion(v string) {
	o.NewsVersion.Set(&v)
}
// SetNewsVersionNil sets the value for NewsVersion to be an explicit nil
func (o *Config) SetNewsVersionNil() {
	o.NewsVersion.Set(nil)
}

// UnsetNewsVersion ensures that no value is present for NewsVersion, not even an explicit nil
func (o *Config) UnsetNewsVersion() {
	o.NewsVersion.Unset()
}

// GetNftInfos returns the NftInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetNftInfos() string {
	if o == nil || IsNil(o.NftInfos.Get()) {
		var ret string
		return ret
	}
	return *o.NftInfos.Get()
}

// GetNftInfosOk returns a tuple with the NftInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetNftInfosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NftInfos.Get(), o.NftInfos.IsSet()
}

// HasNftInfos returns a boolean if a field has been set.
func (o *Config) HasNftInfos() bool {
	if o != nil && o.NftInfos.IsSet() {
		return true
	}

	return false
}

// SetNftInfos gets a reference to the given NullableString and assigns it to the NftInfos field.
func (o *Config) SetNftInfos(v string) {
	o.NftInfos.Set(&v)
}
// SetNftInfosNil sets the value for NftInfos to be an explicit nil
func (o *Config) SetNftInfosNil() {
	o.NftInfos.Set(nil)
}

// UnsetNftInfos ensures that no value is present for NftInfos, not even an explicit nil
func (o *Config) UnsetNftInfos() {
	o.NftInfos.Unset()
}

// GetPromotionStickerPackIds returns the PromotionStickerPackIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetPromotionStickerPackIds() string {
	if o == nil || IsNil(o.PromotionStickerPackIds.Get()) {
		var ret string
		return ret
	}
	return *o.PromotionStickerPackIds.Get()
}

// GetPromotionStickerPackIdsOk returns a tuple with the PromotionStickerPackIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetPromotionStickerPackIdsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PromotionStickerPackIds.Get(), o.PromotionStickerPackIds.IsSet()
}

// HasPromotionStickerPackIds returns a boolean if a field has been set.
func (o *Config) HasPromotionStickerPackIds() bool {
	if o != nil && o.PromotionStickerPackIds.IsSet() {
		return true
	}

	return false
}

// SetPromotionStickerPackIds gets a reference to the given NullableString and assigns it to the PromotionStickerPackIds field.
func (o *Config) SetPromotionStickerPackIds(v string) {
	o.PromotionStickerPackIds.Set(&v)
}
// SetPromotionStickerPackIdsNil sets the value for PromotionStickerPackIds to be an explicit nil
func (o *Config) SetPromotionStickerPackIdsNil() {
	o.PromotionStickerPackIds.Set(nil)
}

// UnsetPromotionStickerPackIds ensures that no value is present for PromotionStickerPackIds, not even an explicit nil
func (o *Config) UnsetPromotionStickerPackIds() {
	o.PromotionStickerPackIds.Unset()
}

// GetShouldAppendUserIdToNewsUrl returns the ShouldAppendUserIdToNewsUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetShouldAppendUserIdToNewsUrl() string {
	if o == nil || IsNil(o.ShouldAppendUserIdToNewsUrl.Get()) {
		var ret string
		return ret
	}
	return *o.ShouldAppendUserIdToNewsUrl.Get()
}

// GetShouldAppendUserIdToNewsUrlOk returns a tuple with the ShouldAppendUserIdToNewsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetShouldAppendUserIdToNewsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShouldAppendUserIdToNewsUrl.Get(), o.ShouldAppendUserIdToNewsUrl.IsSet()
}

// HasShouldAppendUserIdToNewsUrl returns a boolean if a field has been set.
func (o *Config) HasShouldAppendUserIdToNewsUrl() bool {
	if o != nil && o.ShouldAppendUserIdToNewsUrl.IsSet() {
		return true
	}

	return false
}

// SetShouldAppendUserIdToNewsUrl gets a reference to the given NullableString and assigns it to the ShouldAppendUserIdToNewsUrl field.
func (o *Config) SetShouldAppendUserIdToNewsUrl(v string) {
	o.ShouldAppendUserIdToNewsUrl.Set(&v)
}
// SetShouldAppendUserIdToNewsUrlNil sets the value for ShouldAppendUserIdToNewsUrl to be an explicit nil
func (o *Config) SetShouldAppendUserIdToNewsUrlNil() {
	o.ShouldAppendUserIdToNewsUrl.Set(nil)
}

// UnsetShouldAppendUserIdToNewsUrl ensures that no value is present for ShouldAppendUserIdToNewsUrl, not even an explicit nil
func (o *Config) UnsetShouldAppendUserIdToNewsUrl() {
	o.ShouldAppendUserIdToNewsUrl.Unset()
}

// GetSupportEmailAddress returns the SupportEmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetSupportEmailAddress() string {
	if o == nil || IsNil(o.SupportEmailAddress.Get()) {
		var ret string
		return ret
	}
	return *o.SupportEmailAddress.Get()
}

// GetSupportEmailAddressOk returns a tuple with the SupportEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetSupportEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportEmailAddress.Get(), o.SupportEmailAddress.IsSet()
}

// HasSupportEmailAddress returns a boolean if a field has been set.
func (o *Config) HasSupportEmailAddress() bool {
	if o != nil && o.SupportEmailAddress.IsSet() {
		return true
	}

	return false
}

// SetSupportEmailAddress gets a reference to the given NullableString and assigns it to the SupportEmailAddress field.
func (o *Config) SetSupportEmailAddress(v string) {
	o.SupportEmailAddress.Set(&v)
}
// SetSupportEmailAddressNil sets the value for SupportEmailAddress to be an explicit nil
func (o *Config) SetSupportEmailAddressNil() {
	o.SupportEmailAddress.Set(nil)
}

// UnsetSupportEmailAddress ensures that no value is present for SupportEmailAddress, not even an explicit nil
func (o *Config) UnsetSupportEmailAddress() {
	o.SupportEmailAddress.Unset()
}

// GetTokenInfos returns the TokenInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetTokenInfos() string {
	if o == nil || IsNil(o.TokenInfos.Get()) {
		var ret string
		return ret
	}
	return *o.TokenInfos.Get()
}

// GetTokenInfosOk returns a tuple with the TokenInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetTokenInfosOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenInfos.Get(), o.TokenInfos.IsSet()
}

// HasTokenInfos returns a boolean if a field has been set.
func (o *Config) HasTokenInfos() bool {
	if o != nil && o.TokenInfos.IsSet() {
		return true
	}

	return false
}

// SetTokenInfos gets a reference to the given NullableString and assigns it to the TokenInfos field.
func (o *Config) SetTokenInfos(v string) {
	o.TokenInfos.Set(&v)
}
// SetTokenInfosNil sets the value for TokenInfos to be an explicit nil
func (o *Config) SetTokenInfosNil() {
	o.TokenInfos.Set(nil)
}

// UnsetTokenInfos ensures that no value is present for TokenInfos, not even an explicit nil
func (o *Config) UnsetTokenInfos() {
	o.TokenInfos.Unset()
}

// GetUseRandomMessageRefreshRate returns the UseRandomMessageRefreshRate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetUseRandomMessageRefreshRate() string {
	if o == nil || IsNil(o.UseRandomMessageRefreshRate.Get()) {
		var ret string
		return ret
	}
	return *o.UseRandomMessageRefreshRate.Get()
}

// GetUseRandomMessageRefreshRateOk returns a tuple with the UseRandomMessageRefreshRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetUseRandomMessageRefreshRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UseRandomMessageRefreshRate.Get(), o.UseRandomMessageRefreshRate.IsSet()
}

// HasUseRandomMessageRefreshRate returns a boolean if a field has been set.
func (o *Config) HasUseRandomMessageRefreshRate() bool {
	if o != nil && o.UseRandomMessageRefreshRate.IsSet() {
		return true
	}

	return false
}

// SetUseRandomMessageRefreshRate gets a reference to the given NullableString and assigns it to the UseRandomMessageRefreshRate field.
func (o *Config) SetUseRandomMessageRefreshRate(v string) {
	o.UseRandomMessageRefreshRate.Set(&v)
}
// SetUseRandomMessageRefreshRateNil sets the value for UseRandomMessageRefreshRate to be an explicit nil
func (o *Config) SetUseRandomMessageRefreshRateNil() {
	o.UseRandomMessageRefreshRate.Set(nil)
}

// UnsetUseRandomMessageRefreshRate ensures that no value is present for UseRandomMessageRefreshRate, not even an explicit nil
func (o *Config) UnsetUseRandomMessageRefreshRate() {
	o.UseRandomMessageRefreshRate.Unset()
}

// GetWeb3IsMaintenanced returns the Web3IsMaintenanced field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetWeb3IsMaintenanced() string {
	if o == nil || IsNil(o.Web3IsMaintenanced.Get()) {
		var ret string
		return ret
	}
	return *o.Web3IsMaintenanced.Get()
}

// GetWeb3IsMaintenancedOk returns a tuple with the Web3IsMaintenanced field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetWeb3IsMaintenancedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Web3IsMaintenanced.Get(), o.Web3IsMaintenanced.IsSet()
}

// HasWeb3IsMaintenanced returns a boolean if a field has been set.
func (o *Config) HasWeb3IsMaintenanced() bool {
	if o != nil && o.Web3IsMaintenanced.IsSet() {
		return true
	}

	return false
}

// SetWeb3IsMaintenanced gets a reference to the given NullableString and assigns it to the Web3IsMaintenanced field.
func (o *Config) SetWeb3IsMaintenanced(v string) {
	o.Web3IsMaintenanced.Set(&v)
}
// SetWeb3IsMaintenancedNil sets the value for Web3IsMaintenanced to be an explicit nil
func (o *Config) SetWeb3IsMaintenancedNil() {
	o.Web3IsMaintenanced.Set(nil)
}

// UnsetWeb3IsMaintenanced ensures that no value is present for Web3IsMaintenanced, not even an explicit nil
func (o *Config) UnsetWeb3IsMaintenanced() {
	o.Web3IsMaintenanced.Unset()
}

// GetWeb3LocalizedMaintenanceUrl returns the Web3LocalizedMaintenanceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetWeb3LocalizedMaintenanceUrl() string {
	if o == nil || IsNil(o.Web3LocalizedMaintenanceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.Web3LocalizedMaintenanceUrl.Get()
}

// GetWeb3LocalizedMaintenanceUrlOk returns a tuple with the Web3LocalizedMaintenanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetWeb3LocalizedMaintenanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Web3LocalizedMaintenanceUrl.Get(), o.Web3LocalizedMaintenanceUrl.IsSet()
}

// HasWeb3LocalizedMaintenanceUrl returns a boolean if a field has been set.
func (o *Config) HasWeb3LocalizedMaintenanceUrl() bool {
	if o != nil && o.Web3LocalizedMaintenanceUrl.IsSet() {
		return true
	}

	return false
}

// SetWeb3LocalizedMaintenanceUrl gets a reference to the given NullableString and assigns it to the Web3LocalizedMaintenanceUrl field.
func (o *Config) SetWeb3LocalizedMaintenanceUrl(v string) {
	o.Web3LocalizedMaintenanceUrl.Set(&v)
}
// SetWeb3LocalizedMaintenanceUrlNil sets the value for Web3LocalizedMaintenanceUrl to be an explicit nil
func (o *Config) SetWeb3LocalizedMaintenanceUrlNil() {
	o.Web3LocalizedMaintenanceUrl.Set(nil)
}

// UnsetWeb3LocalizedMaintenanceUrl ensures that no value is present for Web3LocalizedMaintenanceUrl, not even an explicit nil
func (o *Config) UnsetWeb3LocalizedMaintenanceUrl() {
	o.Web3LocalizedMaintenanceUrl.Unset()
}

// GetWeb3MaintenanceFeatures returns the Web3MaintenanceFeatures field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetWeb3MaintenanceFeatures() string {
	if o == nil || IsNil(o.Web3MaintenanceFeatures.Get()) {
		var ret string
		return ret
	}
	return *o.Web3MaintenanceFeatures.Get()
}

// GetWeb3MaintenanceFeaturesOk returns a tuple with the Web3MaintenanceFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetWeb3MaintenanceFeaturesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Web3MaintenanceFeatures.Get(), o.Web3MaintenanceFeatures.IsSet()
}

// HasWeb3MaintenanceFeatures returns a boolean if a field has been set.
func (o *Config) HasWeb3MaintenanceFeatures() bool {
	if o != nil && o.Web3MaintenanceFeatures.IsSet() {
		return true
	}

	return false
}

// SetWeb3MaintenanceFeatures gets a reference to the given NullableString and assigns it to the Web3MaintenanceFeatures field.
func (o *Config) SetWeb3MaintenanceFeatures(v string) {
	o.Web3MaintenanceFeatures.Set(&v)
}
// SetWeb3MaintenanceFeaturesNil sets the value for Web3MaintenanceFeatures to be an explicit nil
func (o *Config) SetWeb3MaintenanceFeaturesNil() {
	o.Web3MaintenanceFeatures.Set(nil)
}

// UnsetWeb3MaintenanceFeatures ensures that no value is present for Web3MaintenanceFeatures, not even an explicit nil
func (o *Config) UnsetWeb3MaintenanceFeatures() {
	o.Web3MaintenanceFeatures.Unset()
}

// GetWeb3MaintenanceUrl returns the Web3MaintenanceUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetWeb3MaintenanceUrl() string {
	if o == nil || IsNil(o.Web3MaintenanceUrl.Get()) {
		var ret string
		return ret
	}
	return *o.Web3MaintenanceUrl.Get()
}

// GetWeb3MaintenanceUrlOk returns a tuple with the Web3MaintenanceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetWeb3MaintenanceUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Web3MaintenanceUrl.Get(), o.Web3MaintenanceUrl.IsSet()
}

// HasWeb3MaintenanceUrl returns a boolean if a field has been set.
func (o *Config) HasWeb3MaintenanceUrl() bool {
	if o != nil && o.Web3MaintenanceUrl.IsSet() {
		return true
	}

	return false
}

// SetWeb3MaintenanceUrl gets a reference to the given NullableString and assigns it to the Web3MaintenanceUrl field.
func (o *Config) SetWeb3MaintenanceUrl(v string) {
	o.Web3MaintenanceUrl.Set(&v)
}
// SetWeb3MaintenanceUrlNil sets the value for Web3MaintenanceUrl to be an explicit nil
func (o *Config) SetWeb3MaintenanceUrlNil() {
	o.Web3MaintenanceUrl.Set(nil)
}

// UnsetWeb3MaintenanceUrl ensures that no value is present for Web3MaintenanceUrl, not even an explicit nil
func (o *Config) UnsetWeb3MaintenanceUrl() {
	o.Web3MaintenanceUrl.Unset()
}

// GetXYayGlobalId returns the XYayGlobalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetXYayGlobalId() string {
	if o == nil || IsNil(o.XYayGlobalId.Get()) {
		var ret string
		return ret
	}
	return *o.XYayGlobalId.Get()
}

// GetXYayGlobalIdOk returns a tuple with the XYayGlobalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetXYayGlobalIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XYayGlobalId.Get(), o.XYayGlobalId.IsSet()
}

// HasXYayGlobalId returns a boolean if a field has been set.
func (o *Config) HasXYayGlobalId() bool {
	if o != nil && o.XYayGlobalId.IsSet() {
		return true
	}

	return false
}

// SetXYayGlobalId gets a reference to the given NullableString and assigns it to the XYayGlobalId field.
func (o *Config) SetXYayGlobalId(v string) {
	o.XYayGlobalId.Set(&v)
}
// SetXYayGlobalIdNil sets the value for XYayGlobalId to be an explicit nil
func (o *Config) SetXYayGlobalIdNil() {
	o.XYayGlobalId.Set(nil)
}

// UnsetXYayGlobalId ensures that no value is present for XYayGlobalId, not even an explicit nil
func (o *Config) UnsetXYayGlobalId() {
	o.XYayGlobalId.Unset()
}

// GetXYayJpId returns the XYayJpId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetXYayJpId() string {
	if o == nil || IsNil(o.XYayJpId.Get()) {
		var ret string
		return ret
	}
	return *o.XYayJpId.Get()
}

// GetXYayJpIdOk returns a tuple with the XYayJpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetXYayJpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XYayJpId.Get(), o.XYayJpId.IsSet()
}

// HasXYayJpId returns a boolean if a field has been set.
func (o *Config) HasXYayJpId() bool {
	if o != nil && o.XYayJpId.IsSet() {
		return true
	}

	return false
}

// SetXYayJpId gets a reference to the given NullableString and assigns it to the XYayJpId field.
func (o *Config) SetXYayJpId(v string) {
	o.XYayJpId.Set(&v)
}
// SetXYayJpIdNil sets the value for XYayJpId to be an explicit nil
func (o *Config) SetXYayJpIdNil() {
	o.XYayJpId.Set(nil)
}

// UnsetXYayJpId ensures that no value is present for XYayJpId, not even an explicit nil
func (o *Config) UnsetXYayJpId() {
	o.XYayJpId.Unset()
}

// GetXYayUpdateId returns the XYayUpdateId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Config) GetXYayUpdateId() string {
	if o == nil || IsNil(o.XYayUpdateId.Get()) {
		var ret string
		return ret
	}
	return *o.XYayUpdateId.Get()
}

// GetXYayUpdateIdOk returns a tuple with the XYayUpdateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Config) GetXYayUpdateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.XYayUpdateId.Get(), o.XYayUpdateId.IsSet()
}

// HasXYayUpdateId returns a boolean if a field has been set.
func (o *Config) HasXYayUpdateId() bool {
	if o != nil && o.XYayUpdateId.IsSet() {
		return true
	}

	return false
}

// SetXYayUpdateId gets a reference to the given NullableString and assigns it to the XYayUpdateId field.
func (o *Config) SetXYayUpdateId(v string) {
	o.XYayUpdateId.Set(&v)
}
// SetXYayUpdateIdNil sets the value for XYayUpdateId to be an explicit nil
func (o *Config) SetXYayUpdateIdNil() {
	o.XYayUpdateId.Set(nil)
}

// UnsetXYayUpdateId ensures that no value is present for XYayUpdateId, not even an explicit nil
func (o *Config) UnsetXYayUpdateId() {
	o.XYayUpdateId.Unset()
}

func (o Config) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Config) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AdTesterUserIds.IsSet() {
		toSerialize["ad_tester_user_ids"] = o.AdTesterUserIds.Get()
	}
	if o.AndroidNewsVersion.IsSet() {
		toSerialize["android_news_version"] = o.AndroidNewsVersion.Get()
	}
	if o.Dapps.IsSet() {
		toSerialize["dapps"] = o.Dapps.Get()
	}
	if o.IsChatWebsocketEnabled.IsSet() {
		toSerialize["is_chat_websocket_enabled"] = o.IsChatWebsocketEnabled.Get()
	}
	if o.IsDirectVipPurchaseEnabled.IsSet() {
		toSerialize["is_direct_vip_purchase_enabled"] = o.IsDirectVipPurchaseEnabled.Get()
	}
	if o.IsGiftFeaturesEnabled.IsSet() {
		toSerialize["is_gift_features_enabled"] = o.IsGiftFeaturesEnabled.Get()
	}
	if o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.IsSet() {
		toSerialize["is_id_card_and_phone_verification_check_for_review_enabled"] = o.IsIdCardAndPhoneVerificationCheckForReviewEnabled.Get()
	}
	if o.IsIdCardCheckRequiredBlockerViewEnabled.IsSet() {
		toSerialize["is_id_card_check_required_blocker_view_enabled"] = o.IsIdCardCheckRequiredBlockerViewEnabled.Get()
	}
	if o.IsMaintenanced.IsSet() {
		toSerialize["is_maintenanced"] = o.IsMaintenanced.Get()
	}
	if o.IsPhoneVerificationRequiredBlockerViewEnabled.IsSet() {
		toSerialize["is_phone_verification_required_blocker_view_enabled"] = o.IsPhoneVerificationRequiredBlockerViewEnabled.Get()
	}
	if o.IsSpeedTestEnabled.IsSet() {
		toSerialize["is_speed_test_enabled"] = o.IsSpeedTestEnabled.Get()
	}
	if o.IsSslPinningEnabled.IsSet() {
		toSerialize["is_ssl_pinning_enabled"] = o.IsSslPinningEnabled.Get()
	}
	if o.IsStarEnabled.IsSet() {
		toSerialize["is_star_enabled"] = o.IsStarEnabled.Get()
	}
	if o.LatestAndroidAppVersion.IsSet() {
		toSerialize["latest_android_app_version"] = o.LatestAndroidAppVersion.Get()
	}
	if o.LineOfficialAccountId.IsSet() {
		toSerialize["line_official_account_id"] = o.LineOfficialAccountId.Get()
	}
	if o.LocalizedAndroidNewsTitle.IsSet() {
		toSerialize["localized_android_news_title"] = o.LocalizedAndroidNewsTitle.Get()
	}
	if o.LocalizedAndroidNewsUrl.IsSet() {
		toSerialize["localized_android_news_url"] = o.LocalizedAndroidNewsUrl.Get()
	}
	if o.LocalizedMaintenanceUrl.IsSet() {
		toSerialize["localized_maintenance_url"] = o.LocalizedMaintenanceUrl.Get()
	}
	if o.LocalizedNewsTitle.IsSet() {
		toSerialize["localized_news_title"] = o.LocalizedNewsTitle.Get()
	}
	if o.LocalizedNewsUrl.IsSet() {
		toSerialize["localized_news_url"] = o.LocalizedNewsUrl.Get()
	}
	if o.MaintenanceFeatures.IsSet() {
		toSerialize["maintenance_features"] = o.MaintenanceFeatures.Get()
	}
	if o.MaxImageFrameCount.IsSet() {
		toSerialize["max_image_frame_count"] = o.MaxImageFrameCount.Get()
	}
	if o.MinimumAndroidAppVersionRequired.IsSet() {
		toSerialize["minimum_android_app_version_required"] = o.MinimumAndroidAppVersionRequired.Get()
	}
	if o.MinimumAndroidVersionSupported.IsSet() {
		toSerialize["minimum_android_version_supported"] = o.MinimumAndroidVersionSupported.Get()
	}
	if o.NewsVersion.IsSet() {
		toSerialize["news_version"] = o.NewsVersion.Get()
	}
	if o.NftInfos.IsSet() {
		toSerialize["nft_infos"] = o.NftInfos.Get()
	}
	if o.PromotionStickerPackIds.IsSet() {
		toSerialize["promotion_sticker_pack_ids"] = o.PromotionStickerPackIds.Get()
	}
	if o.ShouldAppendUserIdToNewsUrl.IsSet() {
		toSerialize["should_append_user_id_to_news_url"] = o.ShouldAppendUserIdToNewsUrl.Get()
	}
	if o.SupportEmailAddress.IsSet() {
		toSerialize["support_email_address"] = o.SupportEmailAddress.Get()
	}
	if o.TokenInfos.IsSet() {
		toSerialize["token_infos"] = o.TokenInfos.Get()
	}
	if o.UseRandomMessageRefreshRate.IsSet() {
		toSerialize["use_random_message_refresh_rate"] = o.UseRandomMessageRefreshRate.Get()
	}
	if o.Web3IsMaintenanced.IsSet() {
		toSerialize["web3_is_maintenanced"] = o.Web3IsMaintenanced.Get()
	}
	if o.Web3LocalizedMaintenanceUrl.IsSet() {
		toSerialize["web3_localized_maintenance_url"] = o.Web3LocalizedMaintenanceUrl.Get()
	}
	if o.Web3MaintenanceFeatures.IsSet() {
		toSerialize["web3_maintenance_features"] = o.Web3MaintenanceFeatures.Get()
	}
	if o.Web3MaintenanceUrl.IsSet() {
		toSerialize["web3_maintenance_url"] = o.Web3MaintenanceUrl.Get()
	}
	if o.XYayGlobalId.IsSet() {
		toSerialize["x_yay_global_id"] = o.XYayGlobalId.Get()
	}
	if o.XYayJpId.IsSet() {
		toSerialize["x_yay_jp_id"] = o.XYayJpId.Get()
	}
	if o.XYayUpdateId.IsSet() {
		toSerialize["x_yay_update_id"] = o.XYayUpdateId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Config) UnmarshalJSON(data []byte) (err error) {
	varConfig := _Config{}

	err = json.Unmarshal(data, &varConfig)

	if err != nil {
		return err
	}

	*o = Config(varConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ad_tester_user_ids")
		delete(additionalProperties, "android_news_version")
		delete(additionalProperties, "dapps")
		delete(additionalProperties, "is_chat_websocket_enabled")
		delete(additionalProperties, "is_direct_vip_purchase_enabled")
		delete(additionalProperties, "is_gift_features_enabled")
		delete(additionalProperties, "is_id_card_and_phone_verification_check_for_review_enabled")
		delete(additionalProperties, "is_id_card_check_required_blocker_view_enabled")
		delete(additionalProperties, "is_maintenanced")
		delete(additionalProperties, "is_phone_verification_required_blocker_view_enabled")
		delete(additionalProperties, "is_speed_test_enabled")
		delete(additionalProperties, "is_ssl_pinning_enabled")
		delete(additionalProperties, "is_star_enabled")
		delete(additionalProperties, "latest_android_app_version")
		delete(additionalProperties, "line_official_account_id")
		delete(additionalProperties, "localized_android_news_title")
		delete(additionalProperties, "localized_android_news_url")
		delete(additionalProperties, "localized_maintenance_url")
		delete(additionalProperties, "localized_news_title")
		delete(additionalProperties, "localized_news_url")
		delete(additionalProperties, "maintenance_features")
		delete(additionalProperties, "max_image_frame_count")
		delete(additionalProperties, "minimum_android_app_version_required")
		delete(additionalProperties, "minimum_android_version_supported")
		delete(additionalProperties, "news_version")
		delete(additionalProperties, "nft_infos")
		delete(additionalProperties, "promotion_sticker_pack_ids")
		delete(additionalProperties, "should_append_user_id_to_news_url")
		delete(additionalProperties, "support_email_address")
		delete(additionalProperties, "token_infos")
		delete(additionalProperties, "use_random_message_refresh_rate")
		delete(additionalProperties, "web3_is_maintenanced")
		delete(additionalProperties, "web3_localized_maintenance_url")
		delete(additionalProperties, "web3_maintenance_features")
		delete(additionalProperties, "web3_maintenance_url")
		delete(additionalProperties, "x_yay_global_id")
		delete(additionalProperties, "x_yay_jp_id")
		delete(additionalProperties, "x_yay_update_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConfig struct {
	value *Config
	isSet bool
}

func (v NullableConfig) Get() *Config {
	return v.value
}

func (v *NullableConfig) Set(val *Config) {
	v.value = val
	v.isSet = true
}

func (v NullableConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfig(val *Config) *NullableConfig {
	return &NullableConfig{value: val, isSet: true}
}

func (v NullableConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


