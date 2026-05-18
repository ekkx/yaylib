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
import type { LocalizedStringDTO } from './LocalizedStringDTO.js';
import {
    LocalizedStringDTOFromJSON,
    LocalizedStringDTOFromJSONTyped,
    LocalizedStringDTOToJSON,
    LocalizedStringDTOToJSONTyped,
} from './LocalizedStringDTO.js';

/**
 * 
 * @export
 * @interface DAppDTO
 */
export interface DAppDTO {
    /**
     * 
     * @type {LocalizedStringDTO}
     * @memberof DAppDTO
     */
    descriptionLocalized?: LocalizedStringDTO | null;
    /**
     * 
     * @type {boolean}
     * @memberof DAppDTO
     */
    enabled?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof DAppDTO
     */
    iconUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof DAppDTO
     */
    id?: string | null;
    /**
     * 
     * @type {number}
     * @memberof DAppDTO
     */
    order?: number | null;
    /**
     * 
     * @type {Array<string>}
     * @memberof DAppDTO
     */
    searchKeywords?: Array<string> | null;
    /**
     * 
     * @type {boolean}
     * @memberof DAppDTO
     */
    showInRecommendDApps?: boolean | null;
    /**
     * 
     * @type {LocalizedStringDTO}
     * @memberof DAppDTO
     */
    titleLocalized?: LocalizedStringDTO | null;
    /**
     * 
     * @type {LocalizedStringDTO}
     * @memberof DAppDTO
     */
    urlLocalized?: LocalizedStringDTO | null;
}

/**
 * Check if a given object implements the DAppDTO interface.
 */
export function instanceOfDAppDTO(value: object): value is DAppDTO {
    return true;
}

export function DAppDTOFromJSON(json: any): DAppDTO {
    return DAppDTOFromJSONTyped(json, false);
}

export function DAppDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): DAppDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'descriptionLocalized': json['description_localized'] == null ? undefined : LocalizedStringDTOFromJSON(json['description_localized']),
        'enabled': json['enabled'] == null ? undefined : json['enabled'],
        'iconUrl': json['icon_url'] == null ? undefined : json['icon_url'],
        'id': json['id'] == null ? undefined : json['id'],
        'order': json['order'] == null ? undefined : json['order'],
        'searchKeywords': json['search_keywords'] == null ? undefined : json['search_keywords'],
        'showInRecommendDApps': json['show_in_recommend_d_apps'] == null ? undefined : json['show_in_recommend_d_apps'],
        'titleLocalized': json['title_localized'] == null ? undefined : LocalizedStringDTOFromJSON(json['title_localized']),
        'urlLocalized': json['url_localized'] == null ? undefined : LocalizedStringDTOFromJSON(json['url_localized']),
    };
}

export function DAppDTOToJSON(json: any): DAppDTO {
    return DAppDTOToJSONTyped(json, false);
}

export function DAppDTOToJSONTyped(value?: DAppDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'description_localized': LocalizedStringDTOToJSON(value['descriptionLocalized']),
        'enabled': value['enabled'],
        'icon_url': value['iconUrl'],
        'id': value['id'],
        'order': value['order'],
        'search_keywords': value['searchKeywords'],
        'show_in_recommend_d_apps': value['showInRecommendDApps'],
        'title_localized': LocalizedStringDTOToJSON(value['titleLocalized']),
        'url_localized': LocalizedStringDTOToJSON(value['urlLocalized']),
    };
}

