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

/**
 * 
 * @export
 * @interface GiftSendersResponse
 */
export interface GiftSendersResponse {
    /**
     * 
     * @type {Array<RealmUser>}
     * @memberof GiftSendersResponse
     */
    senders?: Array<RealmUser> | null;
    /**
     * 
     * @type {number}
     * @memberof GiftSendersResponse
     */
    totalSendersCount?: number | null;
}

/**
 * Check if a given object implements the GiftSendersResponse interface.
 */
export function instanceOfGiftSendersResponse(value: object): value is GiftSendersResponse {
    return true;
}

export function GiftSendersResponseFromJSON(json: any): GiftSendersResponse {
    return GiftSendersResponseFromJSONTyped(json, false);
}

export function GiftSendersResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftSendersResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'senders': json['senders'] == null ? undefined : ((json['senders'] as Array<any>).map(RealmUserFromJSON)),
        'totalSendersCount': json['total_senders_count'] == null ? undefined : json['total_senders_count'],
    };
}

export function GiftSendersResponseToJSON(json: any): GiftSendersResponse {
    return GiftSendersResponseToJSONTyped(json, false);
}

export function GiftSendersResponseToJSONTyped(value?: GiftSendersResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'senders': value['senders'] == null ? undefined : ((value['senders'] as Array<any>).map(RealmUserToJSON)),
        'total_senders_count': value['totalSendersCount'],
    };
}

