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
import type { BCNetworkData } from './BCNetworkData';
import {
    BCNetworkDataFromJSON,
    BCNetworkDataFromJSONTyped,
    BCNetworkDataToJSON,
    BCNetworkDataToJSONTyped,
} from './BCNetworkData';
import type { TXNLogData } from './TXNLogData';
import {
    TXNLogDataFromJSON,
    TXNLogDataFromJSONTyped,
    TXNLogDataToJSON,
    TXNLogDataToJSONTyped,
} from './TXNLogData';

/**
 * 
 * @export
 * @interface Web3WalletPreSign
 */
export interface Web3WalletPreSign {
    /**
     * 
     * @type {BCNetworkData}
     * @memberof Web3WalletPreSign
     */
    blockchainNetworkData?: BCNetworkData | null;
    /**
     * 
     * @type {TXNLogData}
     * @memberof Web3WalletPreSign
     */
    transactionAndLogData?: TXNLogData | null;
}

/**
 * Check if a given object implements the Web3WalletPreSign interface.
 */
export function instanceOfWeb3WalletPreSign(value: object): value is Web3WalletPreSign {
    return true;
}

export function Web3WalletPreSignFromJSON(json: any): Web3WalletPreSign {
    return Web3WalletPreSignFromJSONTyped(json, false);
}

export function Web3WalletPreSignFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3WalletPreSign {
    if (json == null) {
        return json;
    }
    return {
        
        'blockchainNetworkData': json['blockchain_network_data'] == null ? undefined : BCNetworkDataFromJSON(json['blockchain_network_data']),
        'transactionAndLogData': json['transaction_and_log_data'] == null ? undefined : TXNLogDataFromJSON(json['transaction_and_log_data']),
    };
}

export function Web3WalletPreSignToJSON(json: any): Web3WalletPreSign {
    return Web3WalletPreSignToJSONTyped(json, false);
}

export function Web3WalletPreSignToJSONTyped(value?: Web3WalletPreSign | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'blockchain_network_data': BCNetworkDataToJSON(value['blockchainNetworkData']),
        'transaction_and_log_data': TXNLogDataToJSON(value['transactionAndLogData']),
    };
}

