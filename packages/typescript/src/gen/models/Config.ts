/* tslint:disable */
/* eslint-disable */
/**
 * Yay! API
 * No description provided
 *
 * Schema version: 4.26.1
 * 
 *
 * NOTE: Code generated; DO NOT EDIT.
 * 
 * Do not edit the class manually.
 */

import { mapValues } from '../runtime.js';
/**
 * 
 * @export
 * @interface Config
 */
export interface Config {
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    adTesterUserIds?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    androidNewsVersion?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    dapps?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isChatWebsocketEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isDirectVipPurchaseEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isGiftFeaturesEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isIdCardAndPhoneVerificationCheckForReviewEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isIdCardCheckRequiredBlockerViewEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isMaintenanced?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isPhoneVerificationRequiredBlockerViewEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isSpeedTestEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isSslPinningEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    isStarEnabled?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    latestAndroidAppVersion?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    lineOfficialAccountId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    localizedAndroidNewsTitle?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    localizedAndroidNewsUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    localizedMaintenanceUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    localizedNewsTitle?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    localizedNewsUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    maintenanceFeatures?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    maxImageFrameCount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    minimumAndroidAppVersionRequired?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    minimumAndroidVersionSupported?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    newsVersion?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    nftInfos?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    promotionStickerPackIds?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    shouldAppendUserIdToNewsUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    supportEmailAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    tokenInfos?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    useRandomMessageRefreshRate?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    web3IsMaintenanced?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    web3LocalizedMaintenanceUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    web3MaintenanceFeatures?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    web3MaintenanceUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    xYayGlobalId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    xYayJpId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Config
     */
    xYayUpdateId?: string | null;
}

/**
 * Check if a given object implements the Config interface.
 */
export function instanceOfConfig(value: object): value is Config {
    return true;
}

export function ConfigFromJSON(json: any): Config {
    return ConfigFromJSONTyped(json, false);
}

export function ConfigFromJSONTyped(json: any, ignoreDiscriminator: boolean): Config {
    if (json == null) {
        return json;
    }
    return {
        
        'adTesterUserIds': json['ad_tester_user_ids'] == null ? undefined : json['ad_tester_user_ids'],
        'androidNewsVersion': json['android_news_version'] == null ? undefined : json['android_news_version'],
        'dapps': json['dapps'] == null ? undefined : json['dapps'],
        'isChatWebsocketEnabled': json['is_chat_websocket_enabled'] == null ? undefined : json['is_chat_websocket_enabled'],
        'isDirectVipPurchaseEnabled': json['is_direct_vip_purchase_enabled'] == null ? undefined : json['is_direct_vip_purchase_enabled'],
        'isGiftFeaturesEnabled': json['is_gift_features_enabled'] == null ? undefined : json['is_gift_features_enabled'],
        'isIdCardAndPhoneVerificationCheckForReviewEnabled': json['is_id_card_and_phone_verification_check_for_review_enabled'] == null ? undefined : json['is_id_card_and_phone_verification_check_for_review_enabled'],
        'isIdCardCheckRequiredBlockerViewEnabled': json['is_id_card_check_required_blocker_view_enabled'] == null ? undefined : json['is_id_card_check_required_blocker_view_enabled'],
        'isMaintenanced': json['is_maintenanced'] == null ? undefined : json['is_maintenanced'],
        'isPhoneVerificationRequiredBlockerViewEnabled': json['is_phone_verification_required_blocker_view_enabled'] == null ? undefined : json['is_phone_verification_required_blocker_view_enabled'],
        'isSpeedTestEnabled': json['is_speed_test_enabled'] == null ? undefined : json['is_speed_test_enabled'],
        'isSslPinningEnabled': json['is_ssl_pinning_enabled'] == null ? undefined : json['is_ssl_pinning_enabled'],
        'isStarEnabled': json['is_star_enabled'] == null ? undefined : json['is_star_enabled'],
        'latestAndroidAppVersion': json['latest_android_app_version'] == null ? undefined : json['latest_android_app_version'],
        'lineOfficialAccountId': json['line_official_account_id'] == null ? undefined : json['line_official_account_id'],
        'localizedAndroidNewsTitle': json['localized_android_news_title'] == null ? undefined : json['localized_android_news_title'],
        'localizedAndroidNewsUrl': json['localized_android_news_url'] == null ? undefined : json['localized_android_news_url'],
        'localizedMaintenanceUrl': json['localized_maintenance_url'] == null ? undefined : json['localized_maintenance_url'],
        'localizedNewsTitle': json['localized_news_title'] == null ? undefined : json['localized_news_title'],
        'localizedNewsUrl': json['localized_news_url'] == null ? undefined : json['localized_news_url'],
        'maintenanceFeatures': json['maintenance_features'] == null ? undefined : json['maintenance_features'],
        'maxImageFrameCount': json['max_image_frame_count'] == null ? undefined : json['max_image_frame_count'],
        'minimumAndroidAppVersionRequired': json['minimum_android_app_version_required'] == null ? undefined : json['minimum_android_app_version_required'],
        'minimumAndroidVersionSupported': json['minimum_android_version_supported'] == null ? undefined : json['minimum_android_version_supported'],
        'newsVersion': json['news_version'] == null ? undefined : json['news_version'],
        'nftInfos': json['nft_infos'] == null ? undefined : json['nft_infos'],
        'promotionStickerPackIds': json['promotion_sticker_pack_ids'] == null ? undefined : json['promotion_sticker_pack_ids'],
        'shouldAppendUserIdToNewsUrl': json['should_append_user_id_to_news_url'] == null ? undefined : json['should_append_user_id_to_news_url'],
        'supportEmailAddress': json['support_email_address'] == null ? undefined : json['support_email_address'],
        'tokenInfos': json['token_infos'] == null ? undefined : json['token_infos'],
        'useRandomMessageRefreshRate': json['use_random_message_refresh_rate'] == null ? undefined : json['use_random_message_refresh_rate'],
        'web3IsMaintenanced': json['web3_is_maintenanced'] == null ? undefined : json['web3_is_maintenanced'],
        'web3LocalizedMaintenanceUrl': json['web3_localized_maintenance_url'] == null ? undefined : json['web3_localized_maintenance_url'],
        'web3MaintenanceFeatures': json['web3_maintenance_features'] == null ? undefined : json['web3_maintenance_features'],
        'web3MaintenanceUrl': json['web3_maintenance_url'] == null ? undefined : json['web3_maintenance_url'],
        'xYayGlobalId': json['x_yay_global_id'] == null ? undefined : json['x_yay_global_id'],
        'xYayJpId': json['x_yay_jp_id'] == null ? undefined : json['x_yay_jp_id'],
        'xYayUpdateId': json['x_yay_update_id'] == null ? undefined : json['x_yay_update_id'],
    };
}

export function ConfigToJSON(json: any): Config {
    return ConfigToJSONTyped(json, false);
}

export function ConfigToJSONTyped(value?: Config | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'ad_tester_user_ids': value['adTesterUserIds'],
        'android_news_version': value['androidNewsVersion'],
        'dapps': value['dapps'],
        'is_chat_websocket_enabled': value['isChatWebsocketEnabled'],
        'is_direct_vip_purchase_enabled': value['isDirectVipPurchaseEnabled'],
        'is_gift_features_enabled': value['isGiftFeaturesEnabled'],
        'is_id_card_and_phone_verification_check_for_review_enabled': value['isIdCardAndPhoneVerificationCheckForReviewEnabled'],
        'is_id_card_check_required_blocker_view_enabled': value['isIdCardCheckRequiredBlockerViewEnabled'],
        'is_maintenanced': value['isMaintenanced'],
        'is_phone_verification_required_blocker_view_enabled': value['isPhoneVerificationRequiredBlockerViewEnabled'],
        'is_speed_test_enabled': value['isSpeedTestEnabled'],
        'is_ssl_pinning_enabled': value['isSslPinningEnabled'],
        'is_star_enabled': value['isStarEnabled'],
        'latest_android_app_version': value['latestAndroidAppVersion'],
        'line_official_account_id': value['lineOfficialAccountId'],
        'localized_android_news_title': value['localizedAndroidNewsTitle'],
        'localized_android_news_url': value['localizedAndroidNewsUrl'],
        'localized_maintenance_url': value['localizedMaintenanceUrl'],
        'localized_news_title': value['localizedNewsTitle'],
        'localized_news_url': value['localizedNewsUrl'],
        'maintenance_features': value['maintenanceFeatures'],
        'max_image_frame_count': value['maxImageFrameCount'],
        'minimum_android_app_version_required': value['minimumAndroidAppVersionRequired'],
        'minimum_android_version_supported': value['minimumAndroidVersionSupported'],
        'news_version': value['newsVersion'],
        'nft_infos': value['nftInfos'],
        'promotion_sticker_pack_ids': value['promotionStickerPackIds'],
        'should_append_user_id_to_news_url': value['shouldAppendUserIdToNewsUrl'],
        'support_email_address': value['supportEmailAddress'],
        'token_infos': value['tokenInfos'],
        'use_random_message_refresh_rate': value['useRandomMessageRefreshRate'],
        'web3_is_maintenanced': value['web3IsMaintenanced'],
        'web3_localized_maintenance_url': value['web3LocalizedMaintenanceUrl'],
        'web3_maintenance_features': value['web3MaintenanceFeatures'],
        'web3_maintenance_url': value['web3MaintenanceUrl'],
        'x_yay_global_id': value['xYayGlobalId'],
        'x_yay_jp_id': value['xYayJpId'],
        'x_yay_update_id': value['xYayUpdateId'],
    };
}

