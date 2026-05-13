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

import { mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Web3WalletTransactionHistoryAssetInfo
 */
export interface Web3WalletTransactionHistoryAssetInfo {
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistoryAssetInfo
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistoryAssetInfo
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistoryAssetInfo
     */
    tokenId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistoryAssetInfo
     */
    videoUrl?: string | null;
}

/**
 * Check if a given object implements the Web3WalletTransactionHistoryAssetInfo interface.
 */
export function instanceOfWeb3WalletTransactionHistoryAssetInfo(value: object): value is Web3WalletTransactionHistoryAssetInfo {
    return true;
}

export function Web3WalletTransactionHistoryAssetInfoFromJSON(json: any): Web3WalletTransactionHistoryAssetInfo {
    return Web3WalletTransactionHistoryAssetInfoFromJSONTyped(json, false);
}

export function Web3WalletTransactionHistoryAssetInfoFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletTransactionHistoryAssetInfo {
    if (json == null) {
        return json;
    }
    return {
        
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'name': json['name'] == null ? undefined : json['name'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
        'videoUrl': json['video_url'] == null ? undefined : json['video_url'],
    };
}

export function Web3WalletTransactionHistoryAssetInfoToJSON(json: any): Web3WalletTransactionHistoryAssetInfo {
    return Web3WalletTransactionHistoryAssetInfoToJSONTyped(json, false);
}

export function Web3WalletTransactionHistoryAssetInfoToJSONTyped(value?: Web3WalletTransactionHistoryAssetInfo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'image_url': value['imageUrl'],
        'name': value['name'],
        'token_id': value['tokenId'],
        'video_url': value['videoUrl'],
    };
}

