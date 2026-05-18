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
import type { Web3TokenInfoDTO } from './Web3TokenInfoDTO.js';
import {
    Web3TokenInfoDTOFromJSON,
    Web3TokenInfoDTOFromJSONTyped,
    Web3TokenInfoDTOToJSON,
    Web3TokenInfoDTOToJSONTyped,
} from './Web3TokenInfoDTO.js';

/**
 * 
 * @export
 * @interface Web3TokenInfosDTO
 */
export interface Web3TokenInfosDTO {
    /**
     * 
     * @type {Array<Web3TokenInfoDTO>}
     * @memberof Web3TokenInfosDTO
     */
    tokenInfos?: Array<Web3TokenInfoDTO> | null;
}

/**
 * Check if a given object implements the Web3TokenInfosDTO interface.
 */
export function instanceOfWeb3TokenInfosDTO(value: object): value is Web3TokenInfosDTO {
    return true;
}

export function Web3TokenInfosDTOFromJSON(json: any): Web3TokenInfosDTO {
    return Web3TokenInfosDTOFromJSONTyped(json, false);
}

export function Web3TokenInfosDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3TokenInfosDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'tokenInfos': json['token_infos'] == null ? undefined : ((json['token_infos'] as Array<any>).map(Web3TokenInfoDTOFromJSON)),
    };
}

export function Web3TokenInfosDTOToJSON(json: any): Web3TokenInfosDTO {
    return Web3TokenInfosDTOToJSONTyped(json, false);
}

export function Web3TokenInfosDTOToJSONTyped(value?: Web3TokenInfosDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'token_infos': value['tokenInfos'] == null ? undefined : ((value['tokenInfos'] as Array<any>).map(Web3TokenInfoDTOToJSON)),
    };
}

