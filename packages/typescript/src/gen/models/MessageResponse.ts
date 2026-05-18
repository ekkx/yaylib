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
import type { RealmConferenceCall } from './RealmConferenceCall.js';
import {
    RealmConferenceCallFromJSON,
    RealmConferenceCallFromJSONTyped,
    RealmConferenceCallToJSON,
    RealmConferenceCallToJSONTyped,
} from './RealmConferenceCall.js';

/**
 * 
 * @export
 * @interface MessageResponse
 */
export interface MessageResponse {
    /**
     * 
     * @type {RealmConferenceCall}
     * @memberof MessageResponse
     */
    conferenceCall?: RealmConferenceCall | null;
    /**
     * 
     * @type {number}
     * @memberof MessageResponse
     */
    id?: number | null;
}

/**
 * Check if a given object implements the MessageResponse interface.
 */
export function instanceOfMessageResponse(value: object): value is MessageResponse {
    return true;
}

export function MessageResponseFromJSON(json: any): MessageResponse {
    return MessageResponseFromJSONTyped(json, false);
}

export function MessageResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessageResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'conferenceCall': json['conference_call'] == null ? undefined : RealmConferenceCallFromJSON(json['conference_call']),
        'id': json['id'] == null ? undefined : json['id'],
    };
}

export function MessageResponseToJSON(json: any): MessageResponse {
    return MessageResponseToJSONTyped(json, false);
}

export function MessageResponseToJSONTyped(value?: MessageResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'conference_call': RealmConferenceCallToJSON(value['conferenceCall']),
        'id': value['id'],
    };
}

