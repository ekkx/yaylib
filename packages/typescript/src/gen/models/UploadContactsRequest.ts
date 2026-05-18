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
import type { Contact } from './Contact.js';
import {
    ContactFromJSON,
    ContactFromJSONTyped,
    ContactToJSON,
    ContactToJSONTyped,
} from './Contact.js';

/**
 * 
 * @export
 * @interface UploadContactsRequest
 */
export interface UploadContactsRequest {
    /**
     * 
     * @type {Array<Contact>}
     * @memberof UploadContactsRequest
     */
    contacts?: Array<Contact> | null;
}

/**
 * Check if a given object implements the UploadContactsRequest interface.
 */
export function instanceOfUploadContactsRequest(value: object): value is UploadContactsRequest {
    return true;
}

export function UploadContactsRequestFromJSON(json: any): UploadContactsRequest {
    return UploadContactsRequestFromJSONTyped(json, false);
}

export function UploadContactsRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): UploadContactsRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'contacts': json['contacts'] == null ? undefined : ((json['contacts'] as Array<any>).map(ContactFromJSON)),
    };
}

export function UploadContactsRequestToJSON(json: any): UploadContactsRequest {
    return UploadContactsRequestToJSONTyped(json, false);
}

export function UploadContactsRequestToJSONTyped(value?: UploadContactsRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'contacts': value['contacts'] == null ? undefined : ((value['contacts'] as Array<any>).map(ContactToJSON)),
    };
}

