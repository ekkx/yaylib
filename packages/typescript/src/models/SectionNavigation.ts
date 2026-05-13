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
 * @interface SectionNavigation
 */
export interface SectionNavigation {
    /**
     * 
     * @type {number}
     * @memberof SectionNavigation
     */
    iconId?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof SectionNavigation
     */
    isEnabled?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof SectionNavigation
     */
    sectionIndex?: number | null;
    /**
     * 
     * @type {number}
     * @memberof SectionNavigation
     */
    titleId?: number | null;
}

/**
 * Check if a given object implements the SectionNavigation interface.
 */
export function instanceOfSectionNavigation(value: object): value is SectionNavigation {
    return true;
}

export function SectionNavigationFromJSON(json: any): SectionNavigation {
    return SectionNavigationFromJSONTyped(json, false);
}

export function SectionNavigationFromJSONTyped(json: any, ignoreDiscriminator: boolean): SectionNavigation {
    if (json == null) {
        return json;
    }
    return {
        
        'iconId': json['icon_id'] == null ? undefined : json['icon_id'],
        'isEnabled': json['is_enabled'] == null ? undefined : json['is_enabled'],
        'sectionIndex': json['section_index'] == null ? undefined : json['section_index'],
        'titleId': json['title_id'] == null ? undefined : json['title_id'],
    };
}

export function SectionNavigationToJSON(json: any): SectionNavigation {
    return SectionNavigationToJSONTyped(json, false);
}

export function SectionNavigationToJSONTyped(value?: SectionNavigation | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'icon_id': value['iconId'],
        'is_enabled': value['isEnabled'],
        'section_index': value['sectionIndex'],
        'title_id': value['titleId'],
    };
}

