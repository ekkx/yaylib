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
 * @interface Page
 */
export interface Page {
    /**
     * 
     * @type {number}
     * @memberof Page
     */
    id?: number | null;
    /**
     * 
     * @type {Array<object>}
     * @memberof Page
     */
    items?: Array<object> | null;
    /**
     * 
     * @type {string}
     * @memberof Page
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<object>}
     * @memberof Page
     */
    pinnedItems?: Array<object> | null;
    /**
     * 
     * @type {number}
     * @memberof Page
     */
    totalItemCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Page
     */
    totalItemLimit?: number | null;
}

/**
 * Check if a given object implements the Page interface.
 */
export function instanceOfPage(value: object): value is Page {
    return true;
}

export function PageFromJSON(json: any): Page {
    return PageFromJSONTyped(json, false);
}

export function PageFromJSONTyped(json: any, ignoreDiscriminator: boolean): Page {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'items': json['items'] == null ? undefined : json['items'],
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'pinnedItems': json['pinned_items'] == null ? undefined : json['pinned_items'],
        'totalItemCount': json['total_item_count'] == null ? undefined : json['total_item_count'],
        'totalItemLimit': json['total_item_limit'] == null ? undefined : json['total_item_limit'],
    };
}

export function PageToJSON(json: any): Page {
    return PageToJSONTyped(json, false);
}

export function PageToJSONTyped(value?: Page | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'items': value['items'],
        'next_page_value': value['nextPageValue'],
        'pinned_items': value['pinnedItems'],
        'total_item_count': value['totalItemCount'],
        'total_item_limit': value['totalItemLimit'],
    };
}

