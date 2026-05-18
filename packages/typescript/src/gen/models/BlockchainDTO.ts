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
import type { AdditionGasPercentDTO } from './AdditionGasPercentDTO.js';
import {
    AdditionGasPercentDTOFromJSON,
    AdditionGasPercentDTOFromJSONTyped,
    AdditionGasPercentDTOToJSON,
    AdditionGasPercentDTOToJSONTyped,
} from './AdditionGasPercentDTO.js';
import type { TokenDTO } from './TokenDTO.js';
import {
    TokenDTOFromJSON,
    TokenDTOFromJSONTyped,
    TokenDTOToJSON,
    TokenDTOToJSONTyped,
} from './TokenDTO.js';
import type { NftCollectionDTO } from './NftCollectionDTO.js';
import {
    NftCollectionDTOFromJSON,
    NftCollectionDTOFromJSONTyped,
    NftCollectionDTOToJSON,
    NftCollectionDTOToJSONTyped,
} from './NftCollectionDTO.js';

/**
 * 
 * @export
 * @interface BlockchainDTO
 */
export interface BlockchainDTO {
    /**
     * 
     * @type {AdditionGasPercentDTO}
     * @memberof BlockchainDTO
     */
    additionGasPercent?: AdditionGasPercentDTO | null;
    /**
     * 
     * @type {string}
     * @memberof BlockchainDTO
     */
    blockExplorer?: string | null;
    /**
     * 
     * @type {number}
     * @memberof BlockchainDTO
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof BlockchainDTO
     */
    chainName?: string | null;
    /**
     * 
     * @type {string}
     * @memberof BlockchainDTO
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {Array<NftCollectionDTO>}
     * @memberof BlockchainDTO
     */
    nftCollections?: Array<NftCollectionDTO> | null;
    /**
     * 
     * @type {number}
     * @memberof BlockchainDTO
     */
    priority?: number | null;
    /**
     * 
     * @type {string}
     * @memberof BlockchainDTO
     */
    rpcUrl?: string | null;
    /**
     * 
     * @type {Array<TokenDTO>}
     * @memberof BlockchainDTO
     */
    tokens?: Array<TokenDTO> | null;
}

/**
 * Check if a given object implements the BlockchainDTO interface.
 */
export function instanceOfBlockchainDTO(value: object): value is BlockchainDTO {
    return true;
}

export function BlockchainDTOFromJSON(json: any): BlockchainDTO {
    return BlockchainDTOFromJSONTyped(json, false);
}

export function BlockchainDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): BlockchainDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'additionGasPercent': json['addition_gas_percent'] == null ? undefined : AdditionGasPercentDTOFromJSON(json['addition_gas_percent']),
        'blockExplorer': json['block_explorer'] == null ? undefined : json['block_explorer'],
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'chainName': json['chain_name'] == null ? undefined : json['chain_name'],
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'nftCollections': json['nft_collections'] == null ? undefined : ((json['nft_collections'] as Array<any>).map(NftCollectionDTOFromJSON)),
        'priority': json['priority'] == null ? undefined : json['priority'],
        'rpcUrl': json['rpc_url'] == null ? undefined : json['rpc_url'],
        'tokens': json['tokens'] == null ? undefined : ((json['tokens'] as Array<any>).map(TokenDTOFromJSON)),
    };
}

export function BlockchainDTOToJSON(json: any): BlockchainDTO {
    return BlockchainDTOToJSONTyped(json, false);
}

export function BlockchainDTOToJSONTyped(value?: BlockchainDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'addition_gas_percent': AdditionGasPercentDTOToJSON(value['additionGasPercent']),
        'block_explorer': value['blockExplorer'],
        'chain_id': value['chainId'],
        'chain_name': value['chainName'],
        'image_url': value['imageUrl'],
        'nft_collections': value['nftCollections'] == null ? undefined : ((value['nftCollections'] as Array<any>).map(NftCollectionDTOToJSON)),
        'priority': value['priority'],
        'rpc_url': value['rpcUrl'],
        'tokens': value['tokens'] == null ? undefined : ((value['tokens'] as Array<any>).map(TokenDTOToJSON)),
    };
}

