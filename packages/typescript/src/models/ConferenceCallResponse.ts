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
import type { RealmConferenceCall } from './RealmConferenceCall';
import {
    RealmConferenceCallFromJSON,
    RealmConferenceCallFromJSONTyped,
    RealmConferenceCallToJSON,
    RealmConferenceCallToJSONTyped,
} from './RealmConferenceCall';

/**
 * 
 * @export
 * @interface ConferenceCallResponse
 */
export interface ConferenceCallResponse {
    /**
     * 
     * @type {RealmConferenceCall}
     * @memberof ConferenceCallResponse
     */
    conferenceCall?: RealmConferenceCall | null;
    /**
     * 
     * @type {string}
     * @memberof ConferenceCallResponse
     */
    conferenceCallUserUuid?: string | null;
}

/**
 * Check if a given object implements the ConferenceCallResponse interface.
 */
export function instanceOfConferenceCallResponse(value: object): value is ConferenceCallResponse {
    return true;
}

export function ConferenceCallResponseFromJSON(json: any): ConferenceCallResponse {
    return ConferenceCallResponseFromJSONTyped(json, false);
}

export function ConferenceCallResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ConferenceCallResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'conferenceCall': json['conference_call'] == null ? undefined : RealmConferenceCallFromJSON(json['conference_call']),
        'conferenceCallUserUuid': json['conference_call_user_uuid'] == null ? undefined : json['conference_call_user_uuid'],
    };
}

export function ConferenceCallResponseToJSON(json: any): ConferenceCallResponse {
    return ConferenceCallResponseToJSONTyped(json, false);
}

export function ConferenceCallResponseToJSONTyped(value?: ConferenceCallResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'conference_call': RealmConferenceCallToJSON(value['conferenceCall']),
        'conference_call_user_uuid': value['conferenceCallUserUuid'],
    };
}

