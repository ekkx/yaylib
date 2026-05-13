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
 * @interface ModelWalkthrough
 */
export interface ModelWalkthrough {
    /**
     * 
     * @type {number}
     * @memberof ModelWalkthrough
     */
    articleId?: number | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWalkthrough
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelWalkthrough
     */
    url?: string | null;
}

/**
 * Check if a given object implements the ModelWalkthrough interface.
 */
export function instanceOfModelWalkthrough(value: object): value is ModelWalkthrough {
    return true;
}

export function ModelWalkthroughFromJSON(json: any): ModelWalkthrough {
    return ModelWalkthroughFromJSONTyped(json, false);
}

export function ModelWalkthroughFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelWalkthrough {
    if (json == null) {
        return json;
    }
    return {
        
        'articleId': json['article_id'] == null ? undefined : json['article_id'],
        'title': json['title'] == null ? undefined : json['title'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function ModelWalkthroughToJSON(json: any): ModelWalkthrough {
    return ModelWalkthroughToJSONTyped(json, false);
}

export function ModelWalkthroughToJSONTyped(value?: ModelWalkthrough | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'article_id': value['articleId'],
        'title': value['title'],
        'url': value['url'],
    };
}

