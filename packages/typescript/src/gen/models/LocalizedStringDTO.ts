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
 * @interface LocalizedStringDTO
 */
export interface LocalizedStringDTO {
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    en?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    id?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    ja?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    ko?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    th?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    zhCn?: string | null;
    /**
     * 
     * @type {string}
     * @memberof LocalizedStringDTO
     */
    zhHk?: string | null;
}

/**
 * Check if a given object implements the LocalizedStringDTO interface.
 */
export function instanceOfLocalizedStringDTO(value: object): value is LocalizedStringDTO {
    return true;
}

export function LocalizedStringDTOFromJSON(json: any): LocalizedStringDTO {
    return LocalizedStringDTOFromJSONTyped(json, false);
}

export function LocalizedStringDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): LocalizedStringDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'en': json['en'] == null ? undefined : json['en'],
        'id': json['id'] == null ? undefined : json['id'],
        'ja': json['ja'] == null ? undefined : json['ja'],
        'ko': json['ko'] == null ? undefined : json['ko'],
        'th': json['th'] == null ? undefined : json['th'],
        'zhCn': json['zh_cn'] == null ? undefined : json['zh_cn'],
        'zhHk': json['zh_hk'] == null ? undefined : json['zh_hk'],
    };
}

export function LocalizedStringDTOToJSON(json: any): LocalizedStringDTO {
    return LocalizedStringDTOToJSONTyped(json, false);
}

export function LocalizedStringDTOToJSONTyped(value?: LocalizedStringDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'en': value['en'],
        'id': value['id'],
        'ja': value['ja'],
        'ko': value['ko'],
        'th': value['th'],
        'zh_cn': value['zhCn'],
        'zh_hk': value['zhHk'],
    };
}

