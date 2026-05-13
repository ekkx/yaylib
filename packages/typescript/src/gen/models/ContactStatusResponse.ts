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
import type { ContactStatus } from './ContactStatus';
import {
    ContactStatusFromJSON,
    ContactStatusFromJSONTyped,
    ContactStatusToJSON,
    ContactStatusToJSONTyped,
} from './ContactStatus';

/**
 * 
 * @export
 * @interface ContactStatusResponse
 */
export interface ContactStatusResponse {
    /**
     * 
     * @type {{ [key: string]: ContactStatus; }}
     * @memberof ContactStatusResponse
     */
    contacts?: { [key: string]: ContactStatus; };
}

/**
 * Check if a given object implements the ContactStatusResponse interface.
 */
export function instanceOfContactStatusResponse(value: object): value is ContactStatusResponse {
    return true;
}

export function ContactStatusResponseFromJSON(json: any): ContactStatusResponse {
    return ContactStatusResponseFromJSONTyped(json, false);
}

export function ContactStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContactStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'contacts': json['contacts'] == null ? undefined : (mapValues(json['contacts'], ContactStatusFromJSON)),
    };
}

export function ContactStatusResponseToJSON(json: any): ContactStatusResponse {
    return ContactStatusResponseToJSONTyped(json, false);
}

export function ContactStatusResponseToJSONTyped(value?: ContactStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contacts': value['contacts'] == null ? undefined : (mapValues(value['contacts'], ContactStatusToJSON)),
    };
}

