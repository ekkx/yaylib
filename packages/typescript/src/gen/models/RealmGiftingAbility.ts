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
 * @interface RealmGiftingAbility
 */
export interface RealmGiftingAbility {
    /**
     * 
     * @type {boolean}
     * @memberof RealmGiftingAbility
     */
    canReceive?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmGiftingAbility
     */
    canSend?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof RealmGiftingAbility
     */
    enabled?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof RealmGiftingAbility
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the RealmGiftingAbility interface.
 */
export function instanceOfRealmGiftingAbility(value: object): value is RealmGiftingAbility {
    return true;
}

export function RealmGiftingAbilityFromJSON(json: any): RealmGiftingAbility {
    return RealmGiftingAbilityFromJSONTyped(json, false);
}

export function RealmGiftingAbilityFromJSONTyped(json: any, ignoreDiscriminator: boolean): RealmGiftingAbility {
    if (json == null) {
        return json;
    }
    return {
        
        'canReceive': json['can_receive'] == null ? undefined : json['can_receive'],
        'canSend': json['can_send'] == null ? undefined : json['can_send'],
        'enabled': json['enabled'] == null ? undefined : json['enabled'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function RealmGiftingAbilityToJSON(json: any): RealmGiftingAbility {
    return RealmGiftingAbilityToJSONTyped(json, false);
}

export function RealmGiftingAbilityToJSONTyped(value?: RealmGiftingAbility | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'can_receive': value['canReceive'],
        'can_send': value['canSend'],
        'enabled': value['enabled'],
        'user_id': value['userId'],
    };
}

