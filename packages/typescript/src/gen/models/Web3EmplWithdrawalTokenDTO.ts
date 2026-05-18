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
 * @interface Web3EmplWithdrawalTokenDTO
 */
export interface Web3EmplWithdrawalTokenDTO {
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    chainId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    id?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    tokenAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    tokenName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    tokenSymbol?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3EmplWithdrawalTokenDTO
     */
    veloPoolVersion?: string | null;
}

/**
 * Check if a given object implements the Web3EmplWithdrawalTokenDTO interface.
 */
export function instanceOfWeb3EmplWithdrawalTokenDTO(value: object): value is Web3EmplWithdrawalTokenDTO {
    return true;
}

export function Web3EmplWithdrawalTokenDTOFromJSON(json: any): Web3EmplWithdrawalTokenDTO {
    return Web3EmplWithdrawalTokenDTOFromJSONTyped(json, false);
}

export function Web3EmplWithdrawalTokenDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3EmplWithdrawalTokenDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'id': json['id'] == null ? undefined : json['id'],
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'tokenAddress': json['token_address'] == null ? undefined : json['token_address'],
        'tokenName': json['token_name'] == null ? undefined : json['token_name'],
        'tokenSymbol': json['token_symbol'] == null ? undefined : json['token_symbol'],
        'veloPoolVersion': json['velo_pool_version'] == null ? undefined : json['velo_pool_version'],
    };
}

export function Web3EmplWithdrawalTokenDTOToJSON(json: any): Web3EmplWithdrawalTokenDTO {
    return Web3EmplWithdrawalTokenDTOToJSONTyped(json, false);
}

export function Web3EmplWithdrawalTokenDTOToJSONTyped(value?: Web3EmplWithdrawalTokenDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'id': value['id'],
        'image_url': value['imageUrl'],
        'token_address': value['tokenAddress'],
        'token_name': value['tokenName'],
        'token_symbol': value['tokenSymbol'],
        'velo_pool_version': value['veloPoolVersion'],
    };
}

