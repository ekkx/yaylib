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
 * @interface Log
 */
export interface Log {
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    address?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    blockHash?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    blockNumber?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    data?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    logIndex?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Log
     */
    removed?: boolean | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof Log
     */
    topics?: Array<string> | null;
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    transactionHash?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Log
     */
    transactionIndex?: string | null;
}

/**
 * Check if a given object implements the Log interface.
 */
export function instanceOfLog(value: object): value is Log {
    return true;
}

export function LogFromJSON(json: any): Log {
    return LogFromJSONTyped(json, false);
}

export function LogFromJSONTyped(json: any, ignoreDiscriminator: boolean): Log {
    if (json == null) {
        return json;
    }
    return {
        
        'address': json['address'] == null ? undefined : json['address'],
        'blockHash': json['block_hash'] == null ? undefined : json['block_hash'],
        'blockNumber': json['block_number'] == null ? undefined : json['block_number'],
        'data': json['data'] == null ? undefined : json['data'],
        'logIndex': json['log_index'] == null ? undefined : json['log_index'],
        'removed': json['removed'] == null ? undefined : json['removed'],
        'topics': json['topics'] == null ? undefined : json['topics'],
        'transactionHash': json['transaction_hash'] == null ? undefined : json['transaction_hash'],
        'transactionIndex': json['transaction_index'] == null ? undefined : json['transaction_index'],
    };
}

export function LogToJSON(json: any): Log {
    return LogToJSONTyped(json, false);
}

export function LogToJSONTyped(value?: Log | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'address': value['address'],
        'block_hash': value['blockHash'],
        'block_number': value['blockNumber'],
        'data': value['data'],
        'log_index': value['logIndex'],
        'removed': value['removed'],
        'topics': value['topics'],
        'transaction_hash': value['transactionHash'],
        'transaction_index': value['transactionIndex'],
    };
}

