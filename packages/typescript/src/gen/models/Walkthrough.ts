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
 * @interface Walkthrough
 */
export interface Walkthrough {
    /**
     * 
     * @type {string}
     * @memberof Walkthrough
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Walkthrough
     */
    url?: string | null;
}

/**
 * Check if a given object implements the Walkthrough interface.
 */
export function instanceOfWalkthrough(value: object): value is Walkthrough {
    return true;
}

export function WalkthroughFromJSON(json: any): Walkthrough {
    return WalkthroughFromJSONTyped(json, false);
}

export function WalkthroughFromJSONTyped(json: any, ignoreDiscriminator: boolean): Walkthrough {
    if (json == null) {
        return json;
    }
    return {
        
        'title': json['title'] == null ? undefined : json['title'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function WalkthroughToJSON(json: any): Walkthrough {
    return WalkthroughToJSONTyped(json, false);
}

export function WalkthroughToJSONTyped(value?: Walkthrough | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'title': value['title'],
        'url': value['url'],
    };
}

