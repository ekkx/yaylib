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
import type { BlockchainDTO } from './BlockchainDTO.js';
import {
    BlockchainDTOFromJSON,
    BlockchainDTOFromJSONTyped,
    BlockchainDTOToJSON,
    BlockchainDTOToJSONTyped,
} from './BlockchainDTO.js';

/**
 * 
 * @export
 * @interface BlockchainNetworkInfoResponse
 */
export interface BlockchainNetworkInfoResponse {
    /**
     * 
     * @type {Array<BlockchainDTO>}
     * @memberof BlockchainNetworkInfoResponse
     */
    chains?: Array<BlockchainDTO> | null;
}

/**
 * Check if a given object implements the BlockchainNetworkInfoResponse interface.
 */
export function instanceOfBlockchainNetworkInfoResponse(value: object): value is BlockchainNetworkInfoResponse {
    return true;
}

export function BlockchainNetworkInfoResponseFromJSON(json: any): BlockchainNetworkInfoResponse {
    return BlockchainNetworkInfoResponseFromJSONTyped(json, false);
}

export function BlockchainNetworkInfoResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockchainNetworkInfoResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'chains': json['chains'] == null ? undefined : ((json['chains'] as Array<any>).map(BlockchainDTOFromJSON)),
    };
}

export function BlockchainNetworkInfoResponseToJSON(json: any): BlockchainNetworkInfoResponse {
    return BlockchainNetworkInfoResponseToJSONTyped(json, false);
}

export function BlockchainNetworkInfoResponseToJSONTyped(value?: BlockchainNetworkInfoResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'chains': value['chains'] == null ? undefined : ((value['chains'] as Array<any>).map(BlockchainDTOToJSON)),
    };
}

