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
 * @interface RegisterDeviceTokenResponse
 */
export interface RegisterDeviceTokenResponse {
    /**
     * 
     * @type {number}
     * @memberof RegisterDeviceTokenResponse
     */
    createdAt?: number | null;
    /**
     * 
     * @type {number}
     * @memberof RegisterDeviceTokenResponse
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RegisterDeviceTokenResponse
     */
    serverDeviceId?: string | null;
    /**
     * 
     * @type {number}
     * @memberof RegisterDeviceTokenResponse
     */
    updatedAt?: number | null;
    /**
     * 
     * @type {string}
     * @memberof RegisterDeviceTokenResponse
     */
    uuid?: string | null;
}

/**
 * Check if a given object implements the RegisterDeviceTokenResponse interface.
 */
export function instanceOfRegisterDeviceTokenResponse(value: object): value is RegisterDeviceTokenResponse {
    return true;
}

export function RegisterDeviceTokenResponseFromJSON(json: any): RegisterDeviceTokenResponse {
    return RegisterDeviceTokenResponseFromJSONTyped(json, false);
}

export function RegisterDeviceTokenResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): RegisterDeviceTokenResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'createdAt': json['created_at'] == null ? undefined : json['created_at'],
        'id': json['id'] == null ? undefined : json['id'],
        'serverDeviceId': json['server_device_id'] == null ? undefined : json['server_device_id'],
        'updatedAt': json['updated_at'] == null ? undefined : json['updated_at'],
        'uuid': json['uuid'] == null ? undefined : json['uuid'],
    };
}

export function RegisterDeviceTokenResponseToJSON(json: any): RegisterDeviceTokenResponse {
    return RegisterDeviceTokenResponseToJSONTyped(json, false);
}

export function RegisterDeviceTokenResponseToJSONTyped(value?: RegisterDeviceTokenResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'created_at': value['createdAt'],
        'id': value['id'],
        'server_device_id': value['serverDeviceId'],
        'updated_at': value['updatedAt'],
        'uuid': value['uuid'],
    };
}

