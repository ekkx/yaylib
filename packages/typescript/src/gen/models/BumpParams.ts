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
 * @interface BumpParams
 */
export interface BumpParams {
    /**
     * 
     * @type {number}
     * @memberof BumpParams
     */
    participantLimit?: number | null;
}

/**
 * Check if a given object implements the BumpParams interface.
 */
export function instanceOfBumpParams(value: object): value is BumpParams {
    return true;
}

export function BumpParamsFromJSON(json: any): BumpParams {
    return BumpParamsFromJSONTyped(json, false);
}

export function BumpParamsFromJSONTyped(json: any, ignoreDiscriminator: boolean): BumpParams {
    if (json == null) {
        return json;
    }
    return {
        
        'participantLimit': json['participant_limit'] == null ? undefined : json['participant_limit'],
    };
}

export function BumpParamsToJSON(json: any): BumpParams {
    return BumpParamsToJSONTyped(json, false);
}

export function BumpParamsToJSONTyped(value?: BumpParams | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'participant_limit': value['participantLimit'],
    };
}

