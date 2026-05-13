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
import type { RequestMethod } from './RequestMethod';
import {
    RequestMethodFromJSON,
    RequestMethodFromJSONTyped,
    RequestMethodToJSON,
    RequestMethodToJSONTyped,
} from './RequestMethod';

/**
 * 
 * @export
 * @interface Request
 */
export interface Request {
    /**
     * 
     * @type {RequestMethod}
     * @memberof Request
     */
    method?: RequestMethod | null;
    /**
     * 
     * @type {object}
     * @memberof Request
     */
    params?: object | null;
}

/**
 * Check if a given object implements the Request interface.
 */
export function instanceOfRequest(value: object): value is Request {
    return true;
}

export function RequestFromJSON(json: any): Request {
    return RequestFromJSONTyped(json, false);
}

export function RequestFromJSONTyped(json: any, ignoreDiscriminator: boolean): Request {
    if (json == null) {
        return json;
    }
    return {
        
        'method': json['method'] == null ? undefined : RequestMethodFromJSON(json['method']),
        'params': json['params'] == null ? undefined : json['params'],
    };
}

export function RequestToJSON(json: any): Request {
    return RequestToJSONTyped(json, false);
}

export function RequestToJSONTyped(value?: Request | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'method': RequestMethodToJSON(value['method']),
        'params': value['params'],
    };
}

