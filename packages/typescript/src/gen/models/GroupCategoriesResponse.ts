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
import type { GroupCategory } from './GroupCategory';
import {
    GroupCategoryFromJSON,
    GroupCategoryFromJSONTyped,
    GroupCategoryToJSON,
    GroupCategoryToJSONTyped,
} from './GroupCategory';

/**
 * 
 * @export
 * @interface GroupCategoriesResponse
 */
export interface GroupCategoriesResponse {
    /**
     * 
     * @type {Array<GroupCategory>}
     * @memberof GroupCategoriesResponse
     */
    groupCategories?: Array<GroupCategory> | null;
}

/**
 * Check if a given object implements the GroupCategoriesResponse interface.
 */
export function instanceOfGroupCategoriesResponse(value: object): value is GroupCategoriesResponse {
    return true;
}

export function GroupCategoriesResponseFromJSON(json: any): GroupCategoriesResponse {
    return GroupCategoriesResponseFromJSONTyped(json, false);
}

export function GroupCategoriesResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): GroupCategoriesResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'groupCategories': json['group_categories'] == null ? undefined : ((json['group_categories'] as Array<any>).map(GroupCategoryFromJSON)),
    };
}

export function GroupCategoriesResponseToJSON(json: any): GroupCategoriesResponse {
    return GroupCategoriesResponseToJSONTyped(json, false);
}

export function GroupCategoriesResponseToJSONTyped(value?: GroupCategoriesResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'group_categories': value['groupCategories'] == null ? undefined : ((value['groupCategories'] as Array<any>).map(GroupCategoryToJSON)),
    };
}

