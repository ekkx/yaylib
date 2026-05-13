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
import type { Web3WalletPreSignApprove } from './Web3WalletPreSignApprove';
import {
    Web3WalletPreSignApproveFromJSON,
    Web3WalletPreSignApproveFromJSONTyped,
    Web3WalletPreSignApproveToJSON,
    Web3WalletPreSignApproveToJSONTyped,
} from './Web3WalletPreSignApprove';

/**
 * 
 * @export
 * @interface PreSignApprove
 */
export interface PreSignApprove {
    /**
     * 
     * @type {Web3WalletPreSignApprove}
     * @memberof PreSignApprove
     */
    result?: Web3WalletPreSignApprove | null;
}

/**
 * Check if a given object implements the PreSignApprove interface.
 */
export function instanceOfPreSignApprove(value: object): value is PreSignApprove {
    return true;
}

export function PreSignApproveFromJSON(json: any): PreSignApprove {
    return PreSignApproveFromJSONTyped(json, false);
}

export function PreSignApproveFromJSONTyped(json: any, ignoreDiscriminator: boolean): PreSignApprove {
    if (json == null) {
        return json;
    }
    return {
        
        'result': json['result'] == null ? undefined : Web3WalletPreSignApproveFromJSON(json['result']),
    };
}

export function PreSignApproveToJSON(json: any): PreSignApprove {
    return PreSignApproveToJSONTyped(json, false);
}

export function PreSignApproveToJSONTyped(value?: PreSignApprove | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'result': Web3WalletPreSignApproveToJSON(value['result']),
    };
}

