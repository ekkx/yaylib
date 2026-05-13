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
import type { Web3NftInfoDTO } from './Web3NftInfoDTO';
import {
    Web3NftInfoDTOFromJSON,
    Web3NftInfoDTOFromJSONTyped,
    Web3NftInfoDTOToJSON,
    Web3NftInfoDTOToJSONTyped,
} from './Web3NftInfoDTO';

/**
 * 
 * @export
 * @interface Web3NftInfosDTO
 */
export interface Web3NftInfosDTO {
    /**
     * 
     * @type {Array<Web3NftInfoDTO>}
     * @memberof Web3NftInfosDTO
     */
    nftInfos?: Array<Web3NftInfoDTO> | null;
}

/**
 * Check if a given object implements the Web3NftInfosDTO interface.
 */
export function instanceOfWeb3NftInfosDTO(value: object): value is Web3NftInfosDTO {
    return true;
}

export function Web3NftInfosDTOFromJSON(json: any): Web3NftInfosDTO {
    return Web3NftInfosDTOFromJSONTyped(json, false);
}

export function Web3NftInfosDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3NftInfosDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'nftInfos': json['nft_infos'] == null ? undefined : ((json['nft_infos'] as Array<any>).map(Web3NftInfoDTOFromJSON)),
    };
}

export function Web3NftInfosDTOToJSON(json: any): Web3NftInfosDTO {
    return Web3NftInfosDTOToJSONTyped(json, false);
}

export function Web3NftInfosDTOToJSONTyped(value?: Web3NftInfosDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'nft_infos': value['nftInfos'] == null ? undefined : ((value['nftInfos'] as Array<any>).map(Web3NftInfoDTOToJSON)),
    };
}

