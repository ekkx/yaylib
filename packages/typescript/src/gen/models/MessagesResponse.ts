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
import type { RealmMessage } from './RealmMessage';
import {
    RealmMessageFromJSON,
    RealmMessageFromJSONTyped,
    RealmMessageToJSON,
    RealmMessageToJSONTyped,
} from './RealmMessage';

/**
 * 
 * @export
 * @interface MessagesResponse
 */
export interface MessagesResponse {
    /**
     * 
     * @type {Array<RealmMessage>}
     * @memberof MessagesResponse
     */
    messages?: Array<RealmMessage> | null;
}

/**
 * Check if a given object implements the MessagesResponse interface.
 */
export function instanceOfMessagesResponse(value: object): value is MessagesResponse {
    return true;
}

export function MessagesResponseFromJSON(json: any): MessagesResponse {
    return MessagesResponseFromJSONTyped(json, false);
}

export function MessagesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessagesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'messages': json['messages'] == null ? undefined : ((json['messages'] as Array<any>).map(RealmMessageFromJSON)),
    };
}

export function MessagesResponseToJSON(json: any): MessagesResponse {
    return MessagesResponseToJSONTyped(json, false);
}

export function MessagesResponseToJSONTyped(value?: MessagesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'messages': value['messages'] == null ? undefined : ((value['messages'] as Array<any>).map(RealmMessageToJSON)),
    };
}

