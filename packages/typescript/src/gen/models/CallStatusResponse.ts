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
 * @interface CallStatusResponse
 */
export interface CallStatusResponse {
    /**
     * 
     * @type {boolean}
     * @memberof CallStatusResponse
     */
    phoneStatus?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof CallStatusResponse
     */
    roomUrl?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof CallStatusResponse
     */
    videoStatus?: boolean | null;
}

/**
 * Check if a given object implements the CallStatusResponse interface.
 */
export function instanceOfCallStatusResponse(value: object): value is CallStatusResponse {
    return true;
}

export function CallStatusResponseFromJSON(json: any): CallStatusResponse {
    return CallStatusResponseFromJSONTyped(json, false);
}

export function CallStatusResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CallStatusResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'phoneStatus': json['phone_status'] == null ? undefined : json['phone_status'],
        'roomUrl': json['room_url'] == null ? undefined : json['room_url'],
        'videoStatus': json['video_status'] == null ? undefined : json['video_status'],
    };
}

export function CallStatusResponseToJSON(json: any): CallStatusResponse {
    return CallStatusResponseToJSONTyped(json, false);
}

export function CallStatusResponseToJSONTyped(value?: CallStatusResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'phone_status': value['phoneStatus'],
        'room_url': value['roomUrl'],
        'video_status': value['videoStatus'],
    };
}

