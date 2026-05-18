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
import type { BridgeDTO } from './BridgeDTO.js';
import {
    BridgeDTOFromJSON,
    BridgeDTOFromJSONTyped,
    BridgeDTOToJSON,
    BridgeDTOToJSONTyped,
} from './BridgeDTO.js';
import type { PriceDTO } from './PriceDTO.js';
import {
    PriceDTOFromJSON,
    PriceDTOFromJSONTyped,
    PriceDTOToJSON,
    PriceDTOToJSONTyped,
} from './PriceDTO.js';
import type { SwapDTO } from './SwapDTO.js';
import {
    SwapDTOFromJSON,
    SwapDTOFromJSONTyped,
    SwapDTOToJSON,
    SwapDTOToJSONTyped,
} from './SwapDTO.js';

/**
 * 
 * @export
 * @interface TokenDTO
 */
export interface TokenDTO {
    /**
     * 
     * @type {string}
     * @memberof TokenDTO
     */
    balance?: string | null;
    /**
     * 
     * @type {BridgeDTO}
     * @memberof TokenDTO
     */
    bridge?: BridgeDTO | null;
    /**
     * 
     * @type {string}
     * @memberof TokenDTO
     */
    contractAddress?: string | null;
    /**
     * 
     * @type {number}
     * @memberof TokenDTO
     */
    decimals?: number | null;
    /**
     * 
     * @type {string}
     * @memberof TokenDTO
     */
    imageUrl?: string | null;
    /**
     * 
     * @type {PriceDTO}
     * @memberof TokenDTO
     */
    price?: PriceDTO | null;
    /**
     * 
     * @type {number}
     * @memberof TokenDTO
     */
    priority?: number | null;
    /**
     * 
     * @type {string}
     * @memberof TokenDTO
     */
    status?: string | null;
    /**
     * 
     * @type {SwapDTO}
     * @memberof TokenDTO
     */
    swap?: SwapDTO | null;
    /**
     * 
     * @type {string}
     * @memberof TokenDTO
     */
    symbol?: string | null;
    /**
     * 
     * @type {string}
     * @memberof TokenDTO
     */
    tokenType?: string | null;
}

/**
 * Check if a given object implements the TokenDTO interface.
 */
export function instanceOfTokenDTO(value: object): value is TokenDTO {
    return true;
}

export function TokenDTOFromJSON(json: any): TokenDTO {
    return TokenDTOFromJSONTyped(json, false);
}

export function TokenDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): TokenDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'balance': json['balance'] == null ? undefined : json['balance'],
        'bridge': json['bridge'] == null ? undefined : BridgeDTOFromJSON(json['bridge']),
        'contractAddress': json['contract_address'] == null ? undefined : json['contract_address'],
        'decimals': json['decimals'] == null ? undefined : json['decimals'],
        'imageUrl': json['image_url'] == null ? undefined : json['image_url'],
        'price': json['price'] == null ? undefined : PriceDTOFromJSON(json['price']),
        'priority': json['priority'] == null ? undefined : json['priority'],
        'status': json['status'] == null ? undefined : json['status'],
        'swap': json['swap'] == null ? undefined : SwapDTOFromJSON(json['swap']),
        'symbol': json['symbol'] == null ? undefined : json['symbol'],
        'tokenType': json['token_type'] == null ? undefined : json['token_type'],
    };
}

export function TokenDTOToJSON(json: any): TokenDTO {
    return TokenDTOToJSONTyped(json, false);
}

export function TokenDTOToJSONTyped(value?: TokenDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'balance': value['balance'],
        'bridge': BridgeDTOToJSON(value['bridge']),
        'contract_address': value['contractAddress'],
        'decimals': value['decimals'],
        'image_url': value['imageUrl'],
        'price': PriceDTOToJSON(value['price']),
        'priority': value['priority'],
        'status': value['status'],
        'swap': SwapDTOToJSON(value['swap']),
        'symbol': value['symbol'],
        'token_type': value['tokenType'],
    };
}

