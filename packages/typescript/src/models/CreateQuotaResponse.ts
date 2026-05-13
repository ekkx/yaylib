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
import type { CreateGroupQuota } from './CreateGroupQuota';
import {
    CreateGroupQuotaFromJSON,
    CreateGroupQuotaFromJSONTyped,
    CreateGroupQuotaToJSON,
    CreateGroupQuotaToJSONTyped,
} from './CreateGroupQuota';

/**
 * 
 * @export
 * @interface CreateQuotaResponse
 */
export interface CreateQuotaResponse {
    /**
     * 
     * @type {CreateGroupQuota}
     * @memberof CreateQuotaResponse
     */
    create?: CreateGroupQuota | null;
}

/**
 * Check if a given object implements the CreateQuotaResponse interface.
 */
export function instanceOfCreateQuotaResponse(value: object): value is CreateQuotaResponse {
    return true;
}

export function CreateQuotaResponseFromJSON(json: any): CreateQuotaResponse {
    return CreateQuotaResponseFromJSONTyped(json, false);
}

export function CreateQuotaResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): CreateQuotaResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'create': json['create'] == null ? undefined : CreateGroupQuotaFromJSON(json['create']),
    };
}

export function CreateQuotaResponseToJSON(json: any): CreateQuotaResponse {
    return CreateQuotaResponseToJSONTyped(json, false);
}

export function CreateQuotaResponseToJSONTyped(value?: CreateQuotaResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'create': CreateGroupQuotaToJSON(value['create']),
    };
}

