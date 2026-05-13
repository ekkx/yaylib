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
 * @interface ProfileImage
 */
export interface ProfileImage {
    /**
     * 
     * @type {string}
     * @memberof ProfileImage
     */
    image?: string | null;
}

/**
 * Check if a given object implements the ProfileImage interface.
 */
export function instanceOfProfileImage(value: object): value is ProfileImage {
    return true;
}

export function ProfileImageFromJSON(json: any): ProfileImage {
    return ProfileImageFromJSONTyped(json, false);
}

export function ProfileImageFromJSONTyped(json: any, ignoreDiscriminator: boolean): ProfileImage {
    if (json == null) {
        return json;
    }
    return {
        
        'image': json['image'] == null ? undefined : json['image'],
    };
}

export function ProfileImageToJSON(json: any): ProfileImage {
    return ProfileImageToJSONTyped(json, false);
}

export function ProfileImageToJSONTyped(value?: ProfileImage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'image': value['image'],
    };
}

