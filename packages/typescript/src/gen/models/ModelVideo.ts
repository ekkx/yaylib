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
 * @interface ModelVideo
 */
export interface ModelVideo {
    /**
     * 
     * @type {number}
     * @memberof ModelVideo
     */
    bitrate?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelVideo
     */
    height?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelVideo
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof ModelVideo
     */
    isCompleted?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof ModelVideo
     */
    thumbnailBigUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelVideo
     */
    thumbnailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof ModelVideo
     */
    videoUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof ModelVideo
     */
    viewsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof ModelVideo
     */
    width?: number | null;
}

/**
 * Check if a given object implements the ModelVideo interface.
 */
export function instanceOfModelVideo(value: object): value is ModelVideo {
    return true;
}

export function ModelVideoFromJSON(json: any): ModelVideo {
    return ModelVideoFromJSONTyped(json, false);
}

export function ModelVideoFromJSONTyped(json: any, ignoreDiscriminator: boolean): ModelVideo {
    if (json == null) {
        return json;
    }
    return {
        
        'bitrate': json['bitrate'] == null ? undefined : json['bitrate'],
        'height': json['height'] == null ? undefined : json['height'],
        'id': json['id'] == null ? undefined : json['id'],
        'isCompleted': json['is_completed'] == null ? undefined : json['is_completed'],
        'thumbnailBigUrl': json['thumbnail_big_url'] == null ? undefined : json['thumbnail_big_url'],
        'thumbnailUrl': json['thumbnail_url'] == null ? undefined : json['thumbnail_url'],
        'videoUrl': json['video_url'] == null ? undefined : json['video_url'],
        'viewsCount': json['views_count'] == null ? undefined : json['views_count'],
        'width': json['width'] == null ? undefined : json['width'],
    };
}

export function ModelVideoToJSON(json: any): ModelVideo {
    return ModelVideoToJSONTyped(json, false);
}

export function ModelVideoToJSONTyped(value?: ModelVideo | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'bitrate': value['bitrate'],
        'height': value['height'],
        'id': value['id'],
        'is_completed': value['isCompleted'],
        'thumbnail_big_url': value['thumbnailBigUrl'],
        'thumbnail_url': value['thumbnailUrl'],
        'video_url': value['videoUrl'],
        'views_count': value['viewsCount'],
        'width': value['width'],
    };
}

