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
 * @interface Cooldown
 */
export interface Cooldown {
    /**
     * 
     * @type {boolean}
     * @memberof Cooldown
     */
    active?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Cooldown
     */
    retryAfter?: number | null;
}

/**
 * Check if a given object implements the Cooldown interface.
 */
export function instanceOfCooldown(value: object): value is Cooldown {
    return true;
}

export function CooldownFromJSON(json: any): Cooldown {
    return CooldownFromJSONTyped(json, false);
}

export function CooldownFromJSONTyped(json: any, ignoreDiscriminator: boolean): Cooldown {
    if (json == null) {
        return json;
    }
    return {
        
        'active': json['active'] == null ? undefined : json['active'],
        'retryAfter': json['retry_after'] == null ? undefined : json['retry_after'],
    };
}

export function CooldownToJSON(json: any): Cooldown {
    return CooldownToJSONTyped(json, false);
}

export function CooldownToJSONTyped(value?: Cooldown | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'active': value['active'],
        'retry_after': value['retryAfter'],
    };
}

