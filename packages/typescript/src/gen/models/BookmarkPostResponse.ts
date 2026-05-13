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
 * @interface BookmarkPostResponse
 */
export interface BookmarkPostResponse {
    /**
     * 
     * @type {boolean}
     * @memberof BookmarkPostResponse
     */
    isBookmarked?: boolean | null;
}

/**
 * Check if a given object implements the BookmarkPostResponse interface.
 */
export function instanceOfBookmarkPostResponse(value: object): value is BookmarkPostResponse {
    return true;
}

export function BookmarkPostResponseFromJSON(json: any): BookmarkPostResponse {
    return BookmarkPostResponseFromJSONTyped(json, false);
}

export function BookmarkPostResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): BookmarkPostResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'isBookmarked': json['is_bookmarked'] == null ? undefined : json['is_bookmarked'],
    };
}

export function BookmarkPostResponseToJSON(json: any): BookmarkPostResponse {
    return BookmarkPostResponseToJSONTyped(json, false);
}

export function BookmarkPostResponseToJSONTyped(value?: BookmarkPostResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'is_bookmarked': value['isBookmarked'],
    };
}

