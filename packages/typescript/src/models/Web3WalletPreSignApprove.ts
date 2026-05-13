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
import type { TXNData } from './TXNData';
import {
    TXNDataFromJSON,
    TXNDataFromJSONTyped,
    TXNDataToJSON,
    TXNDataToJSONTyped,
} from './TXNData';
import type { Web3WalletPreSignApproveBCNetworkData } from './Web3WalletPreSignApproveBCNetworkData';
import {
    Web3WalletPreSignApproveBCNetworkDataFromJSON,
    Web3WalletPreSignApproveBCNetworkDataFromJSONTyped,
    Web3WalletPreSignApproveBCNetworkDataToJSON,
    Web3WalletPreSignApproveBCNetworkDataToJSONTyped,
} from './Web3WalletPreSignApproveBCNetworkData';

/**
 * 
 * @export
 * @interface Web3WalletPreSignApprove
 */
export interface Web3WalletPreSignApprove {
    /**
     * 
     * @type {Web3WalletPreSignApproveBCNetworkData}
     * @memberof Web3WalletPreSignApprove
     */
    blockchainNetworkData?: Web3WalletPreSignApproveBCNetworkData | null;
    /**
     * 
     * @type {TXNData}
     * @memberof Web3WalletPreSignApprove
     */
    transactionData?: TXNData | null;
}

/**
 * Check if a given object implements the Web3WalletPreSignApprove interface.
 */
export function instanceOfWeb3WalletPreSignApprove(value: object): value is Web3WalletPreSignApprove {
    return true;
}

export function Web3WalletPreSignApproveFromJSON(json: any): Web3WalletPreSignApprove {
    return Web3WalletPreSignApproveFromJSONTyped(json, false);
}

export function Web3WalletPreSignApproveFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletPreSignApprove {
    if (json == null) {
        return json;
    }
    return {
        
        'blockchainNetworkData': json['blockchain_network_data'] == null ? undefined : Web3WalletPreSignApproveBCNetworkDataFromJSON(json['blockchain_network_data']),
        'transactionData': json['transaction_data'] == null ? undefined : TXNDataFromJSON(json['transaction_data']),
    };
}

export function Web3WalletPreSignApproveToJSON(json: any): Web3WalletPreSignApprove {
    return Web3WalletPreSignApproveToJSONTyped(json, false);
}

export function Web3WalletPreSignApproveToJSONTyped(value?: Web3WalletPreSignApprove | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'blockchain_network_data': Web3WalletPreSignApproveBCNetworkDataToJSON(value['blockchainNetworkData']),
        'transaction_data': TXNDataToJSON(value['transactionData']),
    };
}

