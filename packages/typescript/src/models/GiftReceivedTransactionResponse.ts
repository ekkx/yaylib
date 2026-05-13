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
import type { RealmUser } from './RealmUser';
import {
    RealmUserFromJSON,
    RealmUserFromJSONTyped,
    RealmUserToJSON,
    RealmUserToJSONTyped,
} from './RealmUser';
import type { TransactionGiftReceived } from './TransactionGiftReceived';
import {
    TransactionGiftReceivedFromJSON,
    TransactionGiftReceivedFromJSONTyped,
    TransactionGiftReceivedToJSON,
    TransactionGiftReceivedToJSONTyped,
} from './TransactionGiftReceived';

/**
 * 
 * @export
 * @interface GiftReceivedTransactionResponse
 */
export interface GiftReceivedTransactionResponse {
    /**
     * 
     * @type {number}
     * @memberof GiftReceivedTransactionResponse
     */
    createdAt?: number | null;
    /**
     * 
     * @type {Array<TransactionGiftReceived>}
     * @memberof GiftReceivedTransactionResponse
     */
    gifts?: Array<TransactionGiftReceived> | null;
    /**
     * 
     * @type {RealmUser}
     * @memberof GiftReceivedTransactionResponse
     */
    receiver?: RealmUser | null;
}

/**
 * Check if a given object implements the GiftReceivedTransactionResponse interface.
 */
export function instanceOfGiftReceivedTransactionResponse(value: object): value is GiftReceivedTransactionResponse {
    return true;
}

export function GiftReceivedTransactionResponseFromJSON(json: any): GiftReceivedTransactionResponse {
    return GiftReceivedTransactionResponseFromJSONTyped(json, false);
}

export function GiftReceivedTransactionResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftReceivedTransactionResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'gifts': json['gifts'] == null ? undefined : ((json['gifts'] as Array<any>).map(TransactionGiftReceivedFromJSON)),
        'receiver': json['receiver'] == null ? undefined : RealmUserFromJSON(json['receiver']),
    };
}

export function GiftReceivedTransactionResponseToJSON(json: any): GiftReceivedTransactionResponse {
    return GiftReceivedTransactionResponseToJSONTyped(json, false);
}

export function GiftReceivedTransactionResponseToJSONTyped(value?: GiftReceivedTransactionResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at': value['createdAt'],
        'gifts': value['gifts'] == null ? undefined : ((value['gifts'] as Array<any>).map(TransactionGiftReceivedToJSON)),
        'receiver': RealmUserToJSON(value['receiver']),
    };
}

