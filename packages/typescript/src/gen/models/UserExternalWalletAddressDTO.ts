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
 * @interface UserExternalWalletAddressDTO
 */
export interface UserExternalWalletAddressDTO {
    /**
     * 
     * @type {string}
     * @memberof UserExternalWalletAddressDTO
     */
    externalWalletAddress?: string | null;
}

/**
 * Check if a given object implements the UserExternalWalletAddressDTO interface.
 */
export function instanceOfUserExternalWalletAddressDTO(value: object): value is UserExternalWalletAddressDTO {
    return true;
}

export function UserExternalWalletAddressDTOFromJSON(json: any): UserExternalWalletAddressDTO {
    return UserExternalWalletAddressDTOFromJSONTyped(json, false);
}

export function UserExternalWalletAddressDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserExternalWalletAddressDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'externalWalletAddress': json['external_wallet_address'] == null ? undefined : json['external_wallet_address'],
    };
}

export function UserExternalWalletAddressDTOToJSON(json: any): UserExternalWalletAddressDTO {
    return UserExternalWalletAddressDTOToJSONTyped(json, false);
}

export function UserExternalWalletAddressDTOToJSONTyped(value?: UserExternalWalletAddressDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'external_wallet_address': value['externalWalletAddress'],
    };
}

