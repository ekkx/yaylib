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
 * @interface EmplDTO
 */
export interface EmplDTO {
    /**
     * 
     * @type {string}
     * @memberof EmplDTO
     */
    campaign?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplDTO
     */
    purchased?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplDTO
     */
    regular?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplDTO
     */
    rewarded?: string | null;
    /**
     * 
     * @type {string}
     * @memberof EmplDTO
     */
    total?: string | null;
}

/**
 * Check if a given object implements the EmplDTO interface.
 */
export function instanceOfEmplDTO(value: object): value is EmplDTO {
    return true;
}

export function EmplDTOFromJSON(json: any): EmplDTO {
    return EmplDTOFromJSONTyped(json, false);
}

export function EmplDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'campaign': json['campaign'] == null ? undefined : json['campaign'],
        'purchased': json['purchased'] == null ? undefined : json['purchased'],
        'regular': json['regular'] == null ? undefined : json['regular'],
        'rewarded': json['rewarded'] == null ? undefined : json['rewarded'],
        'total': json['total'] == null ? undefined : json['total'],
    };
}

export function EmplDTOToJSON(json: any): EmplDTO {
    return EmplDTOToJSONTyped(json, false);
}

export function EmplDTOToJSONTyped(value?: EmplDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'campaign': value['campaign'],
        'purchased': value['purchased'],
        'regular': value['regular'],
        'rewarded': value['rewarded'],
        'total': value['total'],
    };
}

