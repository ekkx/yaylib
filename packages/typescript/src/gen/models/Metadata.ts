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
 * @interface Metadata
 */
export interface Metadata {
    /**
     * 
     * @type {boolean}
     * @memberof Metadata
     */
    anonymous?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof Metadata
     */
    body?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof Metadata
     */
    bulkInvitation?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof Metadata
     */
    contentPreview?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Metadata
     */
    groupTopic?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Metadata
     */
    title?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Metadata
     */
    url?: string | null;
}

/**
 * Check if a given object implements the Metadata interface.
 */
export function instanceOfMetadata(value: object): value is Metadata {
    return true;
}

export function MetadataFromJSON(json: any): Metadata {
    return MetadataFromJSONTyped(json, false);
}

export function MetadataFromJSONTyped(json: any, ignoreDiscriminator: boolean): Metadata {
    if (json == null) {
        return json;
    }
    return {
        
        'anonymous': json['anonymous'] == null ? undefined : json['anonymous'],
        'body': json['body'] == null ? undefined : json['body'],
        'bulkInvitation': json['bulk_invitation'] == null ? undefined : json['bulk_invitation'],
        'contentPreview': json['content_preview'] == null ? undefined : json['content_preview'],
        'groupTopic': json['group_topic'] == null ? undefined : json['group_topic'],
        'title': json['title'] == null ? undefined : json['title'],
        'url': json['url'] == null ? undefined : json['url'],
    };
}

export function MetadataToJSON(json: any): Metadata {
    return MetadataToJSONTyped(json, false);
}

export function MetadataToJSONTyped(value?: Metadata | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'anonymous': value['anonymous'],
        'body': value['body'],
        'bulk_invitation': value['bulkInvitation'],
        'content_preview': value['contentPreview'],
        'group_topic': value['groupTopic'],
        'title': value['title'],
        'url': value['url'],
    };
}

