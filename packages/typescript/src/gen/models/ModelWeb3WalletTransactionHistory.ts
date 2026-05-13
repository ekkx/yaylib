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
import type { Web3WalletTransactionHistoryAssetInfo } from './Web3WalletTransactionHistoryAssetInfo';
import {
    Web3WalletTransactionHistoryAssetInfoFromJSON,
    Web3WalletTransactionHistoryAssetInfoFromJSONTyped,
    Web3WalletTransactionHistoryAssetInfoToJSON,
    Web3WalletTransactionHistoryAssetInfoToJSONTyped,
} from './Web3WalletTransactionHistoryAssetInfo';
import type { AssetType } from './AssetType';
import {
    AssetTypeFromJSON,
    AssetTypeFromJSONTyped,
    AssetTypeToJSON,
    AssetTypeToJSONTyped,
} from './AssetType';
import type { EmplDetails } from './EmplDetails';
import {
    EmplDetailsFromJSON,
    EmplDetailsFromJSONTyped,
    EmplDetailsToJSON,
    EmplDetailsToJSONTyped,
} from './EmplDetails';
import type { GiftTransactionDetail } from './GiftTransactionDetail';
import {
    GiftTransactionDetailFromJSON,
    GiftTransactionDetailFromJSONTyped,
    GiftTransactionDetailToJSON,
    GiftTransactionDetailToJSONTyped,
} from './GiftTransactionDetail';

/**
 * 
 * @export
 * @interface ModelWeb3WalletTransactionHistory
 */
export interface ModelWeb3WalletTransactionHistory {
    /**
     * 
     * @type {Web3WalletTransactionHistoryAssetInfo}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    assetInfo?: Web3WalletTransactionHistoryAssetInfo | null;
    /**
     * 
     * @type {AssetType}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    assetType?: AssetType | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    createdAtMillis?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    currency?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    dotMoneyAmount?: number | null;
    /**
     * 
     * @type {EmplDetails}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    emplDetails?: EmplDetails | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    fromAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    giftExchangeUuid?: string | null;
    /**
     * 
     * @type {GiftTransactionDetail}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    giftTransactionDetail?: GiftTransactionDetail | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    giftTransactionUuid?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    isPendingTransaction?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    nftTokenId?: string | null;
    /**
     * 
     * @type {object}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    palTransactionDetail?: object | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    price?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    raceId?: string | null;
    /**
     * 
     * @type {Array<number>}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    raceRanks?: Array<number> | null;
    /**
     * 
     * @type {object}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    status?: object | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    toAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    tokenSymbol?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    transactionCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    transactionHashHex?: string | null;
    /**
     * 
     * @type {object}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    transactionHistoryType?: object | null;
    /**
     * 
     * @type {object}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    transactionType?: object | null;
    /**
     * 
     * @type {number}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    value?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    valueSymbol?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWeb3WalletTransactionHistory
     */
    valueText?: string | null;
}

/**
 * Check if a given object implements the ModelWeb3WalletTransactionHistory interface.
 */
export function instanceOfModelWeb3WalletTransactionHistory(value: object): value is ModelWeb3WalletTransactionHistory {
    return true;
}

export function ModelWeb3WalletTransactionHistoryFromJSON(json: any): ModelWeb3WalletTransactionHistory {
    return ModelWeb3WalletTransactionHistoryFromJSONTyped(json, false);
}

export function ModelWeb3WalletTransactionHistoryFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelWeb3WalletTransactionHistory {
    if (json == null) {
        return json;
    }
    return {
        
        'assetInfo': json['asset_info'] == null ? undefined : Web3WalletTransactionHistoryAssetInfoFromJSON(json['asset_info']),
        'assetType': json['asset_type'] == null ? undefined : AssetTypeFromJSON(json['asset_type']),
        'createdAtMillis': json['created_at_millis'] == null ? undefined : json['created_at_millis'],
        'currency': json['currency'] == null ? undefined : json['currency'],
        'dotMoneyAmount': json['dot_money_amount'] == null ? undefined : json['dot_money_amount'],
        'emplDetails': json['empl_details'] == null ? undefined : EmplDetailsFromJSON(json['empl_details']),
        'fromAddress': json['from_address'] == null ? undefined : json['from_address'],
        'giftExchangeUuid': json['gift_exchange_uuid'] == null ? undefined : json['gift_exchange_uuid'],
        'giftTransactionDetail': json['gift_transaction_detail'] == null ? undefined : GiftTransactionDetailFromJSON(json['gift_transaction_detail']),
        'giftTransactionUuid': json['gift_transaction_uuid'] == null ? undefined : json['gift_transaction_uuid'],
        'isPendingTransaction': json['is_pending_transaction'] == null ? undefined : json['is_pending_transaction'],
        'nftTokenId': json['nft_token_id'] == null ? undefined : json['nft_token_id'],
        'palTransactionDetail': json['pal_transaction_detail'] == null ? undefined : json['pal_transaction_detail'],
        'price': json['price'] == null ? undefined : json['price'],
        'raceId': json['race_id'] == null ? undefined : json['race_id'],
        'raceRanks': json['race_ranks'] == null ? undefined : json['race_ranks'],
        'status': json['status'] == null ? undefined : json['status'],
        'toAddress': json['to_address'] == null ? undefined : json['to_address'],
        'tokenSymbol': json['token_symbol'] == null ? undefined : json['token_symbol'],
        'transactionCode': json['transaction_code'] == null ? undefined : json['transaction_code'],
        'transactionHashHex': json['transaction_hash_hex'] == null ? undefined : json['transaction_hash_hex'],
        'transactionHistoryType': json['transaction_history_type'] == null ? undefined : json['transaction_history_type'],
        'transactionType': json['transaction_type'] == null ? undefined : json['transaction_type'],
        'value': json['value'] == null ? undefined : json['value'],
        'valueSymbol': json['value_symbol'] == null ? undefined : json['value_symbol'],
        'valueText': json['value_text'] == null ? undefined : json['value_text'],
    };
}

export function ModelWeb3WalletTransactionHistoryToJSON(json: any): ModelWeb3WalletTransactionHistory {
    return ModelWeb3WalletTransactionHistoryToJSONTyped(json, false);
}

export function ModelWeb3WalletTransactionHistoryToJSONTyped(value?: ModelWeb3WalletTransactionHistory | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'asset_info': Web3WalletTransactionHistoryAssetInfoToJSON(value['assetInfo']),
        'asset_type': AssetTypeToJSON(value['assetType']),
        'created_at_millis': value['createdAtMillis'],
        'currency': value['currency'],
        'dot_money_amount': value['dotMoneyAmount'],
        'empl_details': EmplDetailsToJSON(value['emplDetails']),
        'from_address': value['fromAddress'],
        'gift_exchange_uuid': value['giftExchangeUuid'],
        'gift_transaction_detail': GiftTransactionDetailToJSON(value['giftTransactionDetail']),
        'gift_transaction_uuid': value['giftTransactionUuid'],
        'is_pending_transaction': value['isPendingTransaction'],
        'nft_token_id': value['nftTokenId'],
        'pal_transaction_detail': value['palTransactionDetail'],
        'price': value['price'],
        'race_id': value['raceId'],
        'race_ranks': value['raceRanks'],
        'status': value['status'],
        'to_address': value['toAddress'],
        'token_symbol': value['tokenSymbol'],
        'transaction_code': value['transactionCode'],
        'transaction_hash_hex': value['transactionHashHex'],
        'transaction_history_type': value['transactionHistoryType'],
        'transaction_type': value['transactionType'],
        'value': value['value'],
        'value_symbol': value['valueSymbol'],
        'value_text': value['valueText'],
    };
}

