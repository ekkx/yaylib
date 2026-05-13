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
import type { SectionNavigation } from './SectionNavigation';
import {
    SectionNavigationFromJSON,
    SectionNavigationFromJSONTyped,
    SectionNavigationToJSON,
    SectionNavigationToJSONTyped,
} from './SectionNavigation';

/**
 * 
 * @export
 * @interface Navigation
 */
export interface Navigation {
    /**
     * 
     * @type {Array<SectionNavigation>}
     * @memberof Navigation
     */
    navigationItemList?: Array<SectionNavigation> | null;
}

/**
 * Check if a given object implements the Navigation interface.
 */
export function instanceOfNavigation(value: object): value is Navigation {
    return true;
}

export function NavigationFromJSON(json: any): Navigation {
    return NavigationFromJSONTyped(json, false);
}

export function NavigationFromJSONTyped(json: any, ignoreDiscriminator: boolean): Navigation {
    if (json == null) {
        return json;
    }
    return {
        
        'navigationItemList': json['navigation_item_list'] == null ? undefined : ((json['navigation_item_list'] as Array<any>).map(SectionNavigationFromJSON)),
    };
}

export function NavigationToJSON(json: any): Navigation {
    return NavigationToJSONTyped(json, false);
}

export function NavigationToJSONTyped(value?: Navigation | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'navigation_item_list': value['navigationItemList'] == null ? undefined : ((value['navigationItemList'] as Array<any>).map(SectionNavigationToJSON)),
    };
}

