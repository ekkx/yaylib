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
 * @interface MaxAttribute
 */
export interface MaxAttribute {
    /**
     * 
     * @type {number}
     * @memberof MaxAttribute
     */
    luck?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MaxAttribute
     */
    power?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MaxAttribute
     */
    speed?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MaxAttribute
     */
    stamina?: number | null;
    /**
     * 
     * @type {number}
     * @memberof MaxAttribute
     */
    technique?: number | null;
}

/**
 * Check if a given object implements the MaxAttribute interface.
 */
export function instanceOfMaxAttribute(value: object): value is MaxAttribute {
    return true;
}

export function MaxAttributeFromJSON(json: any): MaxAttribute {
    return MaxAttributeFromJSONTyped(json, false);
}

export function MaxAttributeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MaxAttribute {
    if (json == null) {
        return json;
    }
    return {
        
        'luck': json['luck'] == null ? undefined : json['luck'],
        'power': json['power'] == null ? undefined : json['power'],
        'speed': json['speed'] == null ? undefined : json['speed'],
        'stamina': json['stamina'] == null ? undefined : json['stamina'],
        'technique': json['technique'] == null ? undefined : json['technique'],
    };
}

export function MaxAttributeToJSON(json: any): MaxAttribute {
    return MaxAttributeToJSONTyped(json, false);
}

export function MaxAttributeToJSONTyped(value?: MaxAttribute | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'luck': value['luck'],
        'power': value['power'],
        'speed': value['speed'],
        'stamina': value['stamina'],
        'technique': value['technique'],
    };
}

