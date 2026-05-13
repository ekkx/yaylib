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
import type { RealmGift } from './RealmGift';
import {
    RealmGiftFromJSON,
    RealmGiftFromJSONTyped,
    RealmGiftToJSON,
    RealmGiftToJSONTyped,
} from './RealmGift';

/**
 * 
 * @export
 * @interface Requirement
 */
export interface Requirement {
    /**
     * 
     * @type {RealmGift}
     * @memberof Requirement
     */
    gift?: RealmGift | null;
    /**
     * 
     * @type {number}
     * @memberof Requirement
     */
    receivedCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Requirement
     */
    requiresAmount?: number | null;
}

/**
 * Check if a given object implements the Requirement interface.
 */
export function instanceOfRequirement(value: object): value is Requirement {
    return true;
}

export function RequirementFromJSON(json: any): Requirement {
    return RequirementFromJSONTyped(json, false);
}

export function RequirementFromJSONTyped(json: any, ignoreDiscriminator: boolean): Requirement {
    if (json == null) {
        return json;
    }
    return {
        
        'gift': json['gift'] == null ? undefined : RealmGiftFromJSON(json['gift']),
        'receivedCount': json['received_count'] == null ? undefined : json['received_count'],
        'requiresAmount': json['requires_amount'] == null ? undefined : json['requires_amount'],
    };
}

export function RequirementToJSON(json: any): Requirement {
    return RequirementToJSONTyped(json, false);
}

export function RequirementToJSONTyped(value?: Requirement | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gift': RealmGiftToJSON(value['gift']),
        'received_count': value['receivedCount'],
        'requires_amount': value['requiresAmount'],
    };
}

