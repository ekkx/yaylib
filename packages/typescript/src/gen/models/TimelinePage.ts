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
 * @interface TimelinePage
 */
export interface TimelinePage {
    /**
     * 
     * @type {number}
     * @memberof TimelinePage
     */
    id?: number | null;
    /**
     * 
     * @type {Array<object>}
     * @memberof TimelinePage
     */
    items?: Array<object> | null;
    /**
     * 
     * @type {string}
     * @memberof TimelinePage
     */
    nextPageValue?: string | null;
    /**
     * 
     * @type {Array<object>}
     * @memberof TimelinePage
     */
    pinnedItems?: Array<object> | null;
    /**
     * 
     * @type {boolean}
     * @memberof TimelinePage
     */
    showMoreHotPostsButton?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof TimelinePage
     */
    totalItemCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof TimelinePage
     */
    totalItemLimit?: number | null;
}

/**
 * Check if a given object implements the TimelinePage interface.
 */
export function instanceOfTimelinePage(value: object): value is TimelinePage {
    return true;
}

export function TimelinePageFromJSON(json: any): TimelinePage {
    return TimelinePageFromJSONTyped(json, false);
}

export function TimelinePageFromJSONTyped(json: any, ignoreDiscriminator: boolean): TimelinePage {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'items': json['items'] == null ? undefined : json['items'],
        'nextPageValue': json['next_page_value'] == null ? undefined : String(json['next_page_value']),
        'pinnedItems': json['pinned_items'] == null ? undefined : json['pinned_items'],
        'showMoreHotPostsButton': json['show_more_hot_posts_button'] == null ? undefined : json['show_more_hot_posts_button'],
        'totalItemCount': json['total_item_count'] == null ? undefined : json['total_item_count'],
        'totalItemLimit': json['total_item_limit'] == null ? undefined : json['total_item_limit'],
    };
}

export function TimelinePageToJSON(json: any): TimelinePage {
    return TimelinePageToJSONTyped(json, false);
}

export function TimelinePageToJSONTyped(value?: TimelinePage | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'items': value['items'],
        'next_page_value': value['nextPageValue'],
        'pinned_items': value['pinnedItems'],
        'show_more_hot_posts_button': value['showMoreHotPostsButton'],
        'total_item_count': value['totalItemCount'],
        'total_item_limit': value['totalItemLimit'],
    };
}

