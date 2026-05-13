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
import type { Web3NftCollectionItemDTO } from './Web3NftCollectionItemDTO';
import {
    Web3NftCollectionItemDTOFromJSON,
    Web3NftCollectionItemDTOFromJSONTyped,
    Web3NftCollectionItemDTOToJSON,
    Web3NftCollectionItemDTOToJSONTyped,
} from './Web3NftCollectionItemDTO';

/**
 * 
 * @export
 * @interface Web3NftCollectionItemsResponse
 */
export interface Web3NftCollectionItemsResponse {
    /**
     * 
     * @type {Array<Web3NftCollectionItemDTO>}
     * @memberof Web3NftCollectionItemsResponse
     */
    data?: Array<Web3NftCollectionItemDTO> | null;
}

/**
 * Check if a given object implements the Web3NftCollectionItemsResponse interface.
 */
export function instanceOfWeb3NftCollectionItemsResponse(value: object): value is Web3NftCollectionItemsResponse {
    return true;
}

export function Web3NftCollectionItemsResponseFromJSON(json: any): Web3NftCollectionItemsResponse {
    return Web3NftCollectionItemsResponseFromJSONTyped(json, false);
}

export function Web3NftCollectionItemsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): Web3NftCollectionItemsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'data': json['data'] == null ? undefined : ((json['data'] as Array<any>).map(Web3NftCollectionItemDTOFromJSON)),
    };
}

export function Web3NftCollectionItemsResponseToJSON(json: any): Web3NftCollectionItemsResponse {
    return Web3NftCollectionItemsResponseToJSONTyped(json, false);
}

export function Web3NftCollectionItemsResponseToJSONTyped(value?: Web3NftCollectionItemsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'data': value['data'] == null ? undefined : ((value['data'] as Array<any>).map(Web3NftCollectionItemDTOToJSON)),
    };
}

