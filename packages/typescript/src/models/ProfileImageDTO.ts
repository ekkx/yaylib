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
/**
 * 
 * @export
 * @interface ProfileImageDTO
 */
export interface ProfileImageDTO {
    /**
     * 
     * @type {string}
     * @memberof ProfileImageDTO
     */
    image?: string | null;
}

/**
 * Check if a given object implements the ProfileImageDTO interface.
 */
export function instanceOfProfileImageDTO(value: object): value is ProfileImageDTO {
    return true;
}

export function ProfileImageDTOFromJSON(json: any): ProfileImageDTO {
    return ProfileImageDTOFromJSONTyped(json, false);
}

export function ProfileImageDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileImageDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'image': json['image'] == null ? undefined : json['image'],
    };
}

export function ProfileImageDTOToJSON(json: any): ProfileImageDTO {
    return ProfileImageDTOToJSONTyped(json, false);
}

export function ProfileImageDTOToJSONTyped(value?: ProfileImageDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'image': value['image'],
    };
}

