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
import type { NFTMetadataDTO } from './NFTMetadataDTO.js';
import {
    NFTMetadataDTOFromJSON,
    NFTMetadataDTOFromJSONTyped,
    NFTMetadataDTOToJSON,
    NFTMetadataDTOToJSONTyped,
} from './NFTMetadataDTO.js';

/**
 * 
 * @export
 * @interface Web3NftCollectionItemDTO
 */
export interface Web3NftCollectionItemDTO {
    /**
     * 
     * @type {string}
     * @memberof Web3NftCollectionItemDTO
     */
    balance?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Web3NftCollectionItemDTO
     */
    chainId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftCollectionItemDTO
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {NFTMetadataDTO}
     * @memberof Web3NftCollectionItemDTO
     */
    extraMetadata?: NFTMetadataDTO | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftCollectionItemDTO
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftCollectionItemDTO
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftCollectionItemDTO
     */
    tokenId?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Web3NftCollectionItemDTO
     */
    tokenType?: string | null;
}

/**
 * Check if a given object implements the Web3NftCollectionItemDTO interface.
 */
export function instanceOfWeb3NftCollectionItemDTO(value: object): value is Web3NftCollectionItemDTO {
    return true;
}

export function Web3NftCollectionItemDTOFromJSON(json: any): Web3NftCollectionItemDTO {
    return Web3NftCollectionItemDTOFromJSONTyped(json, false);
}

export function Web3NftCollectionItemDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3NftCollectionItemDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'balance': json['balance'] == null ? undefined : json['balance'],
        'chainId': json['chain_id'] == null ? undefined : json['chain_id'],
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'extraMetadata': json['extra_metadata'] == null ? undefined : NFTMetadataDTOFromJSON(json['extra_metadata']),
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'name': json['name'] == null ? undefined : json['name'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
        'tokenType': json['token_type'] == null ? undefined : json['token_type'],
    };
}

export function Web3NftCollectionItemDTOToJSON(json: any): Web3NftCollectionItemDTO {
    return Web3NftCollectionItemDTOToJSONTyped(json, false);
}

export function Web3NftCollectionItemDTOToJSONTyped(value?: Web3NftCollectionItemDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'balance': value['balance'],
        'chain_id': value['chainId'],
        'contract_address': value['contractAddress'],
        'extra_metadata': NFTMetadataDTOToJSON(value['extraMetadata']),
        'image_url': value['imageUrl'],
        'name': value['name'],
        'token_id': value['tokenId'],
        'token_type': value['tokenType'],
    };
}

