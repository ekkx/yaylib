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
import type { EmplDTO } from './EmplDTO';
import {
    EmplDTOFromJSON,
    EmplDTOFromJSONTyped,
    EmplDTOToJSON,
    EmplDTOToJSONTyped,
} from './EmplDTO';
import type { PalDetailsDTO } from './PalDetailsDTO';
import {
    PalDetailsDTOFromJSON,
    PalDetailsDTOFromJSONTyped,
    PalDetailsDTOToJSON,
    PalDetailsDTOToJSONTyped,
} from './PalDetailsDTO';

/**
 * 
 * @export
 * @interface EmplActivityDTO
 */
export interface EmplActivityDTO {
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    amount?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    avatarFrameUserUuid?: string | null;
    /**
     * 
     * @type {number}
     * @memberof EmplActivityDTO
     */
    battleHistoryId?: number | null;
    /**
     * 
     * @type {number}
     * @memberof EmplActivityDTO
     */
    chainId?: number | null;
    /**
     * 
     * @type {PalDetailsDTO}
     * @memberof EmplActivityDTO
     */
    details?: PalDetailsDTO | null;
    /**
     * 
     * @type {EmplDTO}
     * @memberof EmplActivityDTO
     */
    empl?: EmplDTO | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    giftExchangeUuid?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    giftTransactionUuid?: string | null;
    /**
     * 
     * @type {number}
     * @memberof EmplActivityDTO
     */
    id?: number | null;
    /**
     * 
     * @type {number}
     * @memberof EmplActivityDTO
     */
    moneyAmount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof EmplActivityDTO
     */
    occurredAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof EmplActivityDTO
     */
    price?: number | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    raceId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    status?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    tokenAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    tokenSymbol?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    transactionCode?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplActivityDTO
     */
    type?: string | null;
}

/**
 * Check if a given object implements the EmplActivityDTO interface.
 */
export function instanceOfEmplActivityDTO(value: object): value is EmplActivityDTO {
    return true;
}

export function EmplActivityDTOFromJSON(json: any): EmplActivityDTO {
    return EmplActivityDTOFromJSONTyped(json, false);
}

export function EmplActivityDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplActivityDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'amount': json['amount'] == null ? undefined : json['amount'],
        'avatarFrameUserUuid': json['avatar_frame_user_uuid'] == null ? undefined : json['avatar_frame_user_uuid'],
        'battleHistoryId': json['battle_history_id'] == null ? undefined : json['battle_history_id'],
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'details': json['details'] == null ? undefined : PalDetailsDTOFromJSON(json['details']),
        'empl': json['empl'] == null ? undefined : EmplDTOFromJSON(json['empl']),
        'giftExchangeUuid': json['gift_exchange_uuid'] == null ? undefined : json['gift_exchange_uuid'],
        'giftTransactionUuid': json['gift_transaction_uuid'] == null ? undefined : json['gift_transaction_uuid'],
        'id': json['id'] == null ? undefined : json['id'],
        'moneyAmount': json['money_amount'] == null ? undefined : json['money_amount'],
        'occurredAt': json['occurred_at'] == null ? undefined : json['occurred_at'],
        'price': json['price'] == null ? undefined : json['price'],
        'raceId': json['race_id'] == null ? undefined : json['race_id'],
        'status': json['status'] == null ? undefined : json['status'],
        'tokenAddress': json['token_address'] == null ? undefined : json['token_address'],
        'tokenSymbol': json['token_symbol'] == null ? undefined : json['token_symbol'],
        'transactionCode': json['transaction_code'] == null ? undefined : json['transaction_code'],
        'type': json['type'] == null ? undefined : json['type'],
    };
}

export function EmplActivityDTOToJSON(json: any): EmplActivityDTO {
    return EmplActivityDTOToJSONTyped(json, false);
}

export function EmplActivityDTOToJSONTyped(value?: EmplActivityDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'amount': value['amount'],
        'avatar_frame_user_uuid': value['avatarFrameUserUuid'],
        'battle_history_id': value['battleHistoryId'],
        'chain_id': value['chainId'],
        'details': PalDetailsDTOToJSON(value['details']),
        'empl': EmplDTOToJSON(value['empl']),
        'gift_exchange_uuid': value['giftExchangeUuid'],
        'gift_transaction_uuid': value['giftTransactionUuid'],
        'id': value['id'],
        'money_amount': value['moneyAmount'],
        'occurred_at': value['occurredAt'],
        'price': value['price'],
        'race_id': value['raceId'],
        'status': value['status'],
        'token_address': value['tokenAddress'],
        'token_symbol': value['tokenSymbol'],
        'transaction_code': value['transactionCode'],
        'type': value['type'],
    };
}

