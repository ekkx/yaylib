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
 * @interface Video
 */
export interface Video {
    /**
     * 
     * @type {number}
     * @memberof Video
     */
    bitrate?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof Video
     */
    completed?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof Video
     */
    height?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Video
     */
    id?: number | null;
    /**
     * 
     * @type {string}
     * @memberof Video
     */
    thumbnailBigUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Video
     */
    thumbnailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof Video
     */
    videoUrl?: string | null;
    /**
     * 
     * @type {number}
     * @memberof Video
     */
    viewsCount?: number | null;
    /**
     * 
     * @type {number}
     * @memberof Video
     */
    width?: number | null;
}

/**
 * Check if a given object implements the Video interface.
 */
export function instanceOfVideo(value: object): value is Video {
    return true;
}

export function VideoFromJSON(json: any): Video {
    return VideoFromJSONTyped(json, false);
}

export function VideoFromJSONTyped(json: any, ignoreDiscriminator: boolean): Video {
    if (json == null) {
        return json;
    }
    return {
        
        'bitrate': json['bitrate'] == null ? undefined : json['bitrate'],
        'completed': json['completed'] == null ? undefined : json['completed'],
        'height': json['height'] == null ? undefined : json['height'],
        'id': json['id'] == null ? undefined : json['id'],
        'thumbnailBigUrl': json['thumbnail_big_url'] == null ? undefined : json['thumbnail_big_url'],
        'thumbnailUrl': json['thumbnail_url'] == null ? undefined : json['thumbnail_url'],
        'videoUrl': json['video_url'] == null ? undefined : json['video_url'],
        'viewsCount': json['views_count'] == null ? undefined : json['views_count'],
        'width': json['width'] == null ? undefined : json['width'],
    };
}

export function VideoToJSON(json: any): Video {
    return VideoToJSONTyped(json, false);
}

export function VideoToJSONTyped(value?: Video | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'bitrate': value['bitrate'],
        'completed': value['completed'],
        'height': value['height'],
        'id': value['id'],
        'thumbnail_big_url': value['thumbnailBigUrl'],
        'thumbnail_url': value['thumbnailUrl'],
        'video_url': value['videoUrl'],
        'views_count': value['viewsCount'],
        'width': value['width'],
    };
}

