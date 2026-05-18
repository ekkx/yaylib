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
 * @interface Web3TokenInfoDTO
 */
export interface Web3TokenInfoDTO {
    /**
     * 
     * @type {number}
     * @memberof Web3TokenInfoDTO
     */
    chainId?: number | null;
    /**
     * 
     * @type {LocalizedStringDTO}
     * @memberof Web3TokenInfoDTO
     */
    localizedUrls?: LocalizedStringDTO | null;
    /**
     * 
     * @type {string}
     * @memberof Web3TokenInfoDTO
     */
    symbol?: string | null;
}

/**
 * Check if a given object implements the Web3TokenInfoDTO interface.
 */
export function instanceOfWeb3TokenInfoDTO(value: object): value is Web3TokenInfoDTO {
    return true;
}

export function Web3TokenInfoDTOFromJSON(json: any): Web3TokenInfoDTO {
    return Web3TokenInfoDTOFromJSONTyped(json, false);
}

export function Web3TokenInfoDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3TokenInfoDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'localizedUrls': json['localized_urls'] == null ? undefined : LocalizedStringDTOFromJSON(json['localized_urls']),
        'symbol': json['symbol'] == null ? undefined : json['symbol'],
    };
}

export function Web3TokenInfoDTOToJSON(json: any): Web3TokenInfoDTO {
    return Web3TokenInfoDTOToJSONTyped(json, false);
}

export function Web3TokenInfoDTOToJSONTyped(value?: Web3TokenInfoDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chain_id': value['chainId'],
        'localized_urls': LocalizedStringDTOToJSON(value['localizedUrls']),
        'symbol': value['symbol'],
    };
}

