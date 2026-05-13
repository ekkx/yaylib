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
 * @interface GiftingAbility
 */
export interface GiftingAbility {
    /**
     * 
     * @type {boolean}
     * @memberof GiftingAbility
     */
    canReceive?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GiftingAbility
     */
    canSend?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof GiftingAbility
     */
    enabled?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof GiftingAbility
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the GiftingAbility interface.
 */
export function instanceOfGiftingAbility(value: object): value is GiftingAbility {
    return true;
}

export function GiftingAbilityFromJSON(json: any): GiftingAbility {
    return GiftingAbilityFromJSONTyped(json, false);
}

export function GiftingAbilityFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftingAbility {
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

export function GiftingAbilityToJSON(json: any): GiftingAbility {
    return GiftingAbilityToJSONTyped(json, false);
}

export function GiftingAbilityToJSONTyped(value?: GiftingAbility | null, ignoreDiscriminator: boolean = false): any {
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

