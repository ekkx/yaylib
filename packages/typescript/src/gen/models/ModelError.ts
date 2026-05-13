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
import type { ErrorType } from './ErrorType';
import {
    ErrorTypeFromJSON,
    ErrorTypeFromJSONTyped,
    ErrorTypeToJSON,
    ErrorTypeToJSONTyped,
} from './ErrorType';

/**
 * 
 * @export
 * @interface ModelError
 */
export interface ModelError {
    /**
     * 
     * @type {object}
     * @memberof ModelError
     */
    action?: object | null;
    /**
     * 
     * @type {number}
     * @memberof ModelError
     */
    originalCode?: number | null;
    /**
     * 
     * @type {object}
     * @memberof ModelError
     */
    throwable?: object | null;
    /**
     * 
     * @type {ErrorType}
     * @memberof ModelError
     */
    type?: ErrorType | null;
}

/**
 * Check if a given object implements the ModelError interface.
 */
export function instanceOfModelError(value: object): value is ModelError {
    return true;
}

export function ModelErrorFromJSON(json: any): ModelError {
    return ModelErrorFromJSONTyped(json, false);
}

export function ModelErrorFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelError {
    if (json == null) {
        return json;
    }
    return {
        
        'action': json['action'] == null ? undefined : json['action'],
        'originalCode': json['original_code'] == null ? undefined : json['original_code'],
        'throwable': json['throwable'] == null ? undefined : json['throwable'],
        'type': json['type'] == null ? undefined : ErrorTypeFromJSON(json['type']),
    };
}

export function ModelErrorToJSON(json: any): ModelError {
    return ModelErrorToJSONTyped(json, false);
}

export function ModelErrorToJSONTyped(value?: ModelError | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'action': value['action'],
        'original_code': value['originalCode'],
        'throwable': value['throwable'],
        'type': ErrorTypeToJSON(value['type']),
    };
}

