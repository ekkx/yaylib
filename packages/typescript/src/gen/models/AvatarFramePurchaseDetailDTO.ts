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
 * @interface AvatarFramePurchaseDetailDTO
 */
export interface AvatarFramePurchaseDetailDTO {
    /**
     * 
     * @type {string}
     * @memberof AvatarFramePurchaseDetailDTO
     */
    frame?: string | null;
    /**
     * 
     * @type {string}
     * @memberof AvatarFramePurchaseDetailDTO
     */
    frameThumbnail?: string | null;
    /**
     * 
     * @type {string}
     * @memberof AvatarFramePurchaseDetailDTO
     */
    frameType?: string | null;
    /**
     * 
     * @type {number}
     * @memberof AvatarFramePurchaseDetailDTO
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof AvatarFramePurchaseDetailDTO
     */
    vipOnly?: boolean | null;
}

/**
 * Check if a given object implements the AvatarFramePurchaseDetailDTO interface.
 */
export function instanceOfAvatarFramePurchaseDetailDTO(value: object): value is AvatarFramePurchaseDetailDTO {
    return true;
}

export function AvatarFramePurchaseDetailDTOFromJSON(json: any): AvatarFramePurchaseDetailDTO {
    return AvatarFramePurchaseDetailDTOFromJSONTyped(json, false);
}

export function AvatarFramePurchaseDetailDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): AvatarFramePurchaseDetailDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'frame': json['frame'] == null ? undefined : json['frame'],
        'frameThumbnail': json['frame_thumbnail'] == null ? undefined : json['frame_thumbnail'],
        'frameType': json['frame_type'] == null ? undefined : json['frame_type'],
        'id': json['id'] == null ? undefined : json['id'],
        'vipOnly': json['vip_only'] == null ? undefined : json['vip_only'],
    };
}

export function AvatarFramePurchaseDetailDTOToJSON(json: any): AvatarFramePurchaseDetailDTO {
    return AvatarFramePurchaseDetailDTOToJSONTyped(json, false);
}

export function AvatarFramePurchaseDetailDTOToJSONTyped(value?: AvatarFramePurchaseDetailDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'frame': value['frame'],
        'frame_thumbnail': value['frameThumbnail'],
        'frame_type': value['frameType'],
        'id': value['id'],
        'vip_only': value['vipOnly'],
    };
}

