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
import type { Web3WalletBlockStoreItem } from './Web3WalletBlockStoreItem.js';
import {
    Web3WalletBlockStoreItemFromJSON,
    Web3WalletBlockStoreItemFromJSONTyped,
    Web3WalletBlockStoreItemToJSON,
    Web3WalletBlockStoreItemToJSONTyped,
} from './Web3WalletBlockStoreItem.js';

/**
 * 
 * @export
 * @interface Web3WalletBlockStoreEntry
 */
export interface Web3WalletBlockStoreEntry {
    /**
     * 
     * @type {Array<Web3WalletBlockStoreItem>}
     * @memberof Web3WalletBlockStoreEntry
     */
    items?: Array<Web3WalletBlockStoreItem> | null;
}

/**
 * Check if a given object implements the Web3WalletBlockStoreEntry interface.
 */
export function instanceOfWeb3WalletBlockStoreEntry(value: object): value is Web3WalletBlockStoreEntry {
    return true;
}

export function Web3WalletBlockStoreEntryFromJSON(json: any): Web3WalletBlockStoreEntry {
    return Web3WalletBlockStoreEntryFromJSONTyped(json, false);
}

export function Web3WalletBlockStoreEntryFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletBlockStoreEntry {
    if (json == null) {
        return json;
    }
    return {
        
        'items': json['items'] == null ? undefined : ((json['items'] as Array<any>).map(Web3WalletBlockStoreItemFromJSON)),
    };
}

export function Web3WalletBlockStoreEntryToJSON(json: any): Web3WalletBlockStoreEntry {
    return Web3WalletBlockStoreEntryToJSONTyped(json, false);
}

export function Web3WalletBlockStoreEntryToJSONTyped(value?: Web3WalletBlockStoreEntry | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'items': value['items'] == null ? undefined : ((value['items'] as Array<any>).map(Web3WalletBlockStoreItemToJSON)),
    };
}

