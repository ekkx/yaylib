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
 * @interface ContactStatus
 */
export interface ContactStatus {
    /**
     * 
     * @type {string}
     * @memberof ContactStatus
     */
    status?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ContactStatus
     */
    userId?: number | null;
}

/**
 * Check if a given object implements the ContactStatus interface.
 */
export function instanceOfContactStatus(value: object): value is ContactStatus {
    return true;
}

export function ContactStatusFromJSON(json: any): ContactStatus {
    return ContactStatusFromJSONTyped(json, false);
}

export function ContactStatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): ContactStatus {
    if (json == null) {
        return json;
    }
    return {
        
        'status': json['status'] == null ? undefined : json['status'],
        'userId': json['user_id'] == null ? undefined : json['user_id'],
    };
}

export function ContactStatusToJSON(json: any): ContactStatus {
    return ContactStatusToJSONTyped(json, false);
}

export function ContactStatusToJSONTyped(value?: ContactStatus | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'status': value['status'],
        'user_id': value['userId'],
    };
}

