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
import type { EmplTokenExchangeDetails } from './EmplTokenExchangeDetails';
import {
    EmplTokenExchangeDetailsFromJSON,
    EmplTokenExchangeDetailsFromJSONTyped,
    EmplTokenExchangeDetailsToJSON,
    EmplTokenExchangeDetailsToJSONTyped,
} from './EmplTokenExchangeDetails';
import type { AssetInfo } from './AssetInfo';
import {
    AssetInfoFromJSON,
    AssetInfoFromJSONTyped,
    AssetInfoToJSON,
    AssetInfoToJSONTyped,
} from './AssetInfo';

/**
 * 
 * @export
 * @interface Web3WalletTransactionHistory
 */
export interface Web3WalletTransactionHistory {
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    amount?: string | null;
    /**
     * 
     * @type {AssetInfo}
     * @memberof Web3WalletTransactionHistory
     */
    assetInfo?: AssetInfo | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletTransactionHistory
     */
    createdAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    currency?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    description?: string | null;
    /**
     * 
     * @type {EmplTokenExchangeDetails}
     * @memberof Web3WalletTransactionHistory
     */
    details?: EmplTokenExchangeDetails | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletTransactionHistory
     */
    finalStatusUpdatedAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    fromAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    giftTransactionUuid?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    hash?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    id?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    toAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    tokenContract?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    tokenId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3WalletTransactionHistory
     */
    transactionCode?: string | null;
    /**
     * 
     * @type {object}
     * @memberof Web3WalletTransactionHistory
     */
    type?: object | null;
    /**
     * 
     * @type {number}
     * @memberof Web3WalletTransactionHistory
     */
    updatedAt?: number | null;
}

/**
 * Check if a given object implements the Web3WalletTransactionHistory interface.
 */
export function instanceOfWeb3WalletTransactionHistory(value: object): value is Web3WalletTransactionHistory {
    return true;
}

export function Web3WalletTransactionHistoryFromJSON(json: any): Web3WalletTransactionHistory {
    return Web3WalletTransactionHistoryFromJSONTyped(json, false);
}

export function Web3WalletTransactionHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletTransactionHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'assetInfo': json['asset_info'] == null ? undefined : AssetInfoFromJSON(json['asset_info']),
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'currency': json['currency'] == null ? undefined : json['currency'],
        'description': json['description'] == null ? undefined : json['description'],
        'details': json['details'] == null ? undefined : EmplTokenExchangeDetailsFromJSON(json['details']),
        'finalStatusUpdatedAt': json['final_status_updated_at'] == null ? undefined : json['final_status_updated_at'],
        'fromAddress': json['from_address'] == null ? undefined : json['from_address'],
        'giftTransactionUuid': json['gift_transaction_uuid'] == null ? undefined : json['gift_transaction_uuid'],
        'hash': json['hash'] == null ? undefined : json['hash'],
        'id': json['id'] == null ? undefined : json['id'],
        'status': json['status'] == null ? undefined : json['status'],
        'toAddress': json['to_address'] == null ? undefined : json['to_address'],
        'tokenContract': json['token_contract'] == null ? undefined : json['token_contract'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
        'transactionCode': json['transaction_code'] == null ? undefined : json['transaction_code'],
        'type': json['type'] == null ? undefined : json['type'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
    };
}

export function Web3WalletTransactionHistoryToJSON(json: any): Web3WalletTransactionHistory {
    return Web3WalletTransactionHistoryToJSONTyped(json, false);
}

export function Web3WalletTransactionHistoryToJSONTyped(value?: Web3WalletTransactionHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'asset_info': AssetInfoToJSON(value['assetInfo']),
        'created_at': value['createdAt'],
        'currency': value['currency'],
        'description': value['description'],
        'details': EmplTokenExchangeDetailsToJSON(value['details']),
        'final_status_updated_at': value['finalStatusUpdatedAt'],
        'from_address': value['fromAddress'],
        'gift_transaction_uuid': value['giftTransactionUuid'],
        'hash': value['hash'],
        'id': value['id'],
        'status': value['status'],
        'to_address': value['toAddress'],
        'token_contract': value['tokenContract'],
        'token_id': value['tokenId'],
        'transaction_code': value['transactionCode'],
        'type': value['type'],
        'updated_at': value['updatedAt'],
    };
}

