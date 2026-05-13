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
 * @interface ConferenceCallBumpParams
 */
export interface ConferenceCallBumpParams {
    /**
     * 
     * @type {number}
     * @memberof ConferenceCallBumpParams
     */
    participantLimit?: number | null;
}

/**
 * Check if a given object implements the ConferenceCallBumpParams interface.
 */
export function instanceOfConferenceCallBumpParams(value: object): value is ConferenceCallBumpParams {
    return true;
}

export function ConferenceCallBumpParamsFromJSON(json: any): ConferenceCallBumpParams {
    return ConferenceCallBumpParamsFromJSONTyped(json, false);
}

export function ConferenceCallBumpParamsFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConferenceCallBumpParams {
    if (json == null) {
        return json;
    }
    return {
        
        'participantLimit': json['participant_limit'] == null ? undefined : json['participant_limit'],
    };
}

export function ConferenceCallBumpParamsToJSON(json: any): ConferenceCallBumpParams {
    return ConferenceCallBumpParamsToJSONTyped(json, false);
}

export function ConferenceCallBumpParamsToJSONTyped(value?: ConferenceCallBumpParams | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'participant_limit': value['participantLimit'],
    };
}

