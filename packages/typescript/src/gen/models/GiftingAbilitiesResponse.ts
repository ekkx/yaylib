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
import type { RealmGiftingAbility } from './RealmGiftingAbility.js';
import {
    RealmGiftingAbilityFromJSON,
    RealmGiftingAbilityFromJSONTyped,
    RealmGiftingAbilityToJSON,
    RealmGiftingAbilityToJSONTyped,
} from './RealmGiftingAbility.js';

/**
 * 
 * @export
 * @interface GiftingAbilitiesResponse
 */
export interface GiftingAbilitiesResponse {
    /**
     * 
     * @type {Array<RealmGiftingAbility>}
     * @memberof GiftingAbilitiesResponse
     */
    giftingAbilities?: Array<RealmGiftingAbility> | null;
}

/**
 * Check if a given object implements the GiftingAbilitiesResponse interface.
 */
export function instanceOfGiftingAbilitiesResponse(value: object): value is GiftingAbilitiesResponse {
    return true;
}

export function GiftingAbilitiesResponseFromJSON(json: any): GiftingAbilitiesResponse {
    return GiftingAbilitiesResponseFromJSONTyped(json, false);
}

export function GiftingAbilitiesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GiftingAbilitiesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'giftingAbilities': json['gifting_abilities'] == null ? undefined : ((json['gifting_abilities'] as Array<any>).map(RealmGiftingAbilityFromJSON)),
    };
}

export function GiftingAbilitiesResponseToJSON(json: any): GiftingAbilitiesResponse {
    return GiftingAbilitiesResponseToJSONTyped(json, false);
}

export function GiftingAbilitiesResponseToJSONTyped(value?: GiftingAbilitiesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'gifting_abilities': value['giftingAbilities'] == null ? undefined : ((value['giftingAbilities'] as Array<any>).map(RealmGiftingAbilityToJSON)),
    };
}

