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
import type { LocalizedStringDTO } from './LocalizedStringDTO.js';
import {
    LocalizedStringDTOFromJSON,
    LocalizedStringDTOFromJSONTyped,
    LocalizedStringDTOToJSON,
    LocalizedStringDTOToJSONTyped,
} from './LocalizedStringDTO.js';

/**
 * 
 * @export
 * @interface Web3NftInfoDTO
 */
export interface Web3NftInfoDTO {
    /**
     * 
     * @type {number}
     * @memberof Web3NftInfoDTO
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftInfoDTO
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {LocalizedStringDTO}
     * @memberof Web3NftInfoDTO
     */
    localizedUrls?: LocalizedStringDTO | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftInfoDTO
     */
    name?: string | null;
}

/**
 * Check if a given object implements the Web3NftInfoDTO interface.
 */
export function instanceOfWeb3NftInfoDTO(value: object): value is Web3NftInfoDTO {
    return true;
}

export function Web3NftInfoDTOFromJSON(json: any): Web3NftInfoDTO {
    return Web3NftInfoDTOFromJSONTyped(json, false);
}

export function Web3NftInfoDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3NftInfoDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'localizedUrls': json['localized_urls'] == null ? undefined : LocalizedStringDTOFromJSON(json['localized_urls']),
        'name': json['name'] == null ? undefined : json['name'],
    };
}

export function Web3NftInfoDTOToJSON(json: any): Web3NftInfoDTO {
    return Web3NftInfoDTOToJSONTyped(json, false);
}

export function Web3NftInfoDTOToJSONTyped(value?: Web3NftInfoDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'contract_address': value['contractAddress'],
        'localized_urls': LocalizedStringDTOToJSON(value['localizedUrls']),
        'name': value['name'],
    };
}

