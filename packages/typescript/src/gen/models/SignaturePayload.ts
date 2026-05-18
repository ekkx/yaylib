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
 * @interface SignaturePayload
 */
export interface SignaturePayload {
    /**
     * 
     * @type {string}
     * @memberof SignaturePayload
     */
    action?: string | null;
    /**
     * 
     * @type {number}
     * @memberof SignaturePayload
     */
    callId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof SignaturePayload
     */
    receiverUuid?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SignaturePayload
     */
    senderUuid?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SignaturePayload
     */
    signature?: string | null;
    /**
     * 
     * @type {number}
     * @memberof SignaturePayload
     */
    timestamp?: number | null;
}

/**
 * Check if a given object implements the SignaturePayload interface.
 */
export function instanceOfSignaturePayload(value: object): value is SignaturePayload {
    return true;
}

export function SignaturePayloadFromJSON(json: any): SignaturePayload {
    return SignaturePayloadFromJSONTyped(json, false);
}

export function SignaturePayloadFromJSONTyped(json: any, ignoreDiscriminator: boolean): SignaturePayload {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'] == null ? undefined : json['action'],
        'callId': json['call_id'] == null ? undefined : json['call_id'],
        'receiverUuid': json['receiver_uuid'] == null ? undefined : json['receiver_uuid'],
        'senderUuid': json['sender_uuid'] == null ? undefined : json['sender_uuid'],
        'signature': json['signature'] == null ? undefined : json['signature'],
        'timestamp': json['timestamp'] == null ? undefined : json['timestamp'],
    };
}

export function SignaturePayloadToJSON(json: any): SignaturePayload {
    return SignaturePayloadToJSONTyped(json, false);
}

export function SignaturePayloadToJSONTyped(value?: SignaturePayload | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'action': value['action'],
        'call_id': value['callId'],
        'receiver_uuid': value['receiverUuid'],
        'sender_uuid': value['senderUuid'],
        'signature': value['signature'],
        'timestamp': value['timestamp'],
    };
}

