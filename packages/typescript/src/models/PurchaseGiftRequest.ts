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
 * @interface PurchaseGiftRequest
 */
export interface PurchaseGiftRequest {
    /**
     * 
     * @type {{ [key: string]: number; }}
     * @memberof PurchaseGiftRequest
     */
    gifts?: { [key: string]: number; };
    /**
     * 
     * @type {boolean}
     * @memberof PurchaseGiftRequest
     */
    isAnonymously?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof PurchaseGiftRequest
     */
    receiverId?: number | null;
}

/**
 * Check if a given object implements the PurchaseGiftRequest interface.
 */
export function instanceOfPurchaseGiftRequest(value: object): value is PurchaseGiftRequest {
    return true;
}

export function PurchaseGiftRequestFromJSON(json: any): PurchaseGiftRequest {
    return PurchaseGiftRequestFromJSONTyped(json, false);
}

export function PurchaseGiftRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): PurchaseGiftRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'gifts': json['gifts'] == null ? undefined : json['gifts'],
        'isAnonymously': json['is_anonymously'] == null ? undefined : json['is_anonymously'],
        'receiverId': json['receiver_id'] == null ? undefined : json['receiver_id'],
    };
}

export function PurchaseGiftRequestToJSON(json: any): PurchaseGiftRequest {
    return PurchaseGiftRequestToJSONTyped(json, false);
}

export function PurchaseGiftRequestToJSONTyped(value?: PurchaseGiftRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifts': value['gifts'],
        'is_anonymously': value['isAnonymously'],
        'receiver_id': value['receiverId'],
    };
}

