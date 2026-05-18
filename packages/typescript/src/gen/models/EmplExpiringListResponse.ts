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
import type { EmplExpiringResponse } from './EmplExpiringResponse.js';
import {
    EmplExpiringResponseFromJSON,
    EmplExpiringResponseFromJSONTyped,
    EmplExpiringResponseToJSON,
    EmplExpiringResponseToJSONTyped,
} from './EmplExpiringResponse.js';

/**
 * 
 * @export
 * @interface EmplExpiringListResponse
 */
export interface EmplExpiringListResponse {
    /**
     * 
     * @type {Array<EmplExpiringResponse>}
     * @memberof EmplExpiringListResponse
     */
    regularExpiring?: Array<EmplExpiringResponse> | null;
    /**
     * 
     * @type {Array<EmplExpiringResponse>}
     * @memberof EmplExpiringListResponse
     */
    rewardedExpiring?: Array<EmplExpiringResponse> | null;
}

/**
 * Check if a given object implements the EmplExpiringListResponse interface.
 */
export function instanceOfEmplExpiringListResponse(value: object): value is EmplExpiringListResponse {
    return true;
}

export function EmplExpiringListResponseFromJSON(json: any): EmplExpiringListResponse {
    return EmplExpiringListResponseFromJSONTyped(json, false);
}

export function EmplExpiringListResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): EmplExpiringListResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'regularExpiring': json['regular_expiring'] == null ? undefined : ((json['regular_expiring'] as Array<any>).map(EmplExpiringResponseFromJSON)),
        'rewardedExpiring': json['rewarded_expiring'] == null ? undefined : ((json['rewarded_expiring'] as Array<any>).map(EmplExpiringResponseFromJSON)),
    };
}

export function EmplExpiringListResponseToJSON(json: any): EmplExpiringListResponse {
    return EmplExpiringListResponseToJSONTyped(json, false);
}

export function EmplExpiringListResponseToJSONTyped(value?: EmplExpiringListResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'regular_expiring': value['regularExpiring'] == null ? undefined : ((value['regularExpiring'] as Array<any>).map(EmplExpiringResponseToJSON)),
        'rewarded_expiring': value['rewardedExpiring'] == null ? undefined : ((value['rewardedExpiring'] as Array<any>).map(EmplExpiringResponseToJSON)),
    };
}

