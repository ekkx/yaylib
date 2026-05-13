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
 * @interface Contact
 */
export interface Contact {
    /**
     * 
     * @type {string}
     * @memberof Contact
     */
    email?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Contact
     */
    mobileNumber?: string | null;
}

/**
 * Check if a given object implements the Contact interface.
 */
export function instanceOfContact(value: object): value is Contact {
    return true;
}

export function ContactFromJSON(json: any): Contact {
    return ContactFromJSONTyped(json, false);
}

export function ContactFromJSONTyped(json: any, ignoreDiscriminator: boolean): Contact {
    if (json == null) {
        return json;
    }
    return {
        
        'email': json['email'] == null ? undefined : json['email'],
        'mobileNumber': json['mobile_number'] == null ? undefined : json['mobile_number'],
    };
}

export function ContactToJSON(json: any): Contact {
    return ContactToJSONTyped(json, false);
}

export function ContactToJSONTyped(value?: Contact | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'email': value['email'],
        'mobile_number': value['mobileNumber'],
    };
}

