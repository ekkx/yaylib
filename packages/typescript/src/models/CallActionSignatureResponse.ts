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
import type { SignaturePayload } from './SignaturePayload';
import {
    SignaturePayloadFromJSON,
    SignaturePayloadFromJSONTyped,
    SignaturePayloadToJSON,
    SignaturePayloadToJSONTyped,
} from './SignaturePayload';

/**
 * 
 * @export
 * @interface CallActionSignatureResponse
 */
export interface CallActionSignatureResponse {
    /**
     * 
     * @type {SignaturePayload}
     * @memberof CallActionSignatureResponse
     */
    signaturePayload?: SignaturePayload | null;
}

/**
 * Check if a given object implements the CallActionSignatureResponse interface.
 */
export function instanceOfCallActionSignatureResponse(value: object): value is CallActionSignatureResponse {
    return true;
}

export function CallActionSignatureResponseFromJSON(json: any): CallActionSignatureResponse {
    return CallActionSignatureResponseFromJSONTyped(json, false);
}

export function CallActionSignatureResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CallActionSignatureResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'signaturePayload': json['signature_payload'] == null ? undefined : SignaturePayloadFromJSON(json['signature_payload']),
    };
}

export function CallActionSignatureResponseToJSON(json: any): CallActionSignatureResponse {
    return CallActionSignatureResponseToJSONTyped(json, false);
}

export function CallActionSignatureResponseToJSONTyped(value?: CallActionSignatureResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'signature_payload': SignaturePayloadToJSON(value['signaturePayload']),
    };
}

