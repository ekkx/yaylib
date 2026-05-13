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
 * @interface SignUpSnsInfoRequest
 */
export interface SignUpSnsInfoRequest {
    /**
     * 
     * @type {string}
     * @memberof SignUpSnsInfoRequest
     */
    type?: string | null;
    /**
     * 
     * @type {string}
     * @memberof SignUpSnsInfoRequest
     */
    uid?: string | null;
}

/**
 * Check if a given object implements the SignUpSnsInfoRequest interface.
 */
export function instanceOfSignUpSnsInfoRequest(value: object): value is SignUpSnsInfoRequest {
    return true;
}

export function SignUpSnsInfoRequestFromJSON(json: any): SignUpSnsInfoRequest {
    return SignUpSnsInfoRequestFromJSONTyped(json, false);
}

export function SignUpSnsInfoRequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): SignUpSnsInfoRequest {
    if (json == null) {
        return json;
    }
    return {
        
        'type': json['type'] == null ? undefined : json['type'],
        'uid': json['uid'] == null ? undefined : json['uid'],
    };
}

export function SignUpSnsInfoRequestToJSON(json: any): SignUpSnsInfoRequest {
    return SignUpSnsInfoRequestToJSONTyped(json, false);
}

export function SignUpSnsInfoRequestToJSONTyped(value?: SignUpSnsInfoRequest | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'type': value['type'],
        'uid': value['uid'],
    };
}

