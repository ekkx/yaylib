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
import type { AssetInfoDTO } from './AssetInfoDTO.js';
import {
    AssetInfoDTOFromJSON,
    AssetInfoDTOFromJSONTyped,
    AssetInfoDTOToJSON,
    AssetInfoDTOToJSONTyped,
} from './AssetInfoDTO.js';

/**
 * 
 * @export
 * @interface OnChainTransactionDTO
 */
export interface OnChainTransactionDTO {
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    amount?: string | null;
    /**
     * 
     * @type {AssetInfoDTO}
     * @memberof OnChainTransactionDTO
     */
    assetInfo?: AssetInfoDTO | null;
    /**
     * 
     * @type {number}
     * @memberof OnChainTransactionDTO
     */
    createdAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    currency?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    description?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    fromAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    hash?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    toAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    tokenContract?: string | null;
    /**
     * 
     * @type {string}
     * @memberof OnChainTransactionDTO
     */
    tokenType?: string | null;
}

/**
 * Check if a given object implements the OnChainTransactionDTO interface.
 */
export function instanceOfOnChainTransactionDTO(value: object): value is OnChainTransactionDTO {
    return true;
}

export function OnChainTransactionDTOFromJSON(json: any): OnChainTransactionDTO {
    return OnChainTransactionDTOFromJSONTyped(json, false);
}

export function OnChainTransactionDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): OnChainTransactionDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'assetInfo': json['asset_info'] == null ? undefined : AssetInfoDTOFromJSON(json['asset_info']),
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'currency': json['currency'] == null ? undefined : json['currency'],
        'description': json['description'] == null ? undefined : json['description'],
        'fromAddress': json['from_address'] == null ? undefined : json['from_address'],
        'hash': json['hash'] == null ? undefined : json['hash'],
        'status': json['status'] == null ? undefined : json['status'],
        'toAddress': json['to_address'] == null ? undefined : json['to_address'],
        'tokenContract': json['token_contract'] == null ? undefined : json['token_contract'],
        'tokenType': json['token_type'] == null ? undefined : json['token_type'],
    };
}

export function OnChainTransactionDTOToJSON(json: any): OnChainTransactionDTO {
    return OnChainTransactionDTOToJSONTyped(json, false);
}

export function OnChainTransactionDTOToJSONTyped(value?: OnChainTransactionDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'asset_info': AssetInfoDTOToJSON(value['assetInfo']),
        'created_at': value['createdAt'],
        'currency': value['currency'],
        'description': value['description'],
        'from_address': value['fromAddress'],
        'hash': value['hash'],
        'status': value['status'],
        'to_address': value['toAddress'],
        'token_contract': value['tokenContract'],
        'token_type': value['tokenType'],
    };
}

