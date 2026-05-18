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
 * @interface VideoProcessedData
 */
export interface VideoProcessedData {
    /**
     * 
     * @type {number}
     * @memberof VideoProcessedData
     */
    id?: number | null;
    /**
     * 
     * @type {boolean}
     * @memberof VideoProcessedData
     */
    videoProcessed?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof VideoProcessedData
     */
    videoThumbnailBigUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof VideoProcessedData
     */
    videoThumbnailUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof VideoProcessedData
     */
    videoUrl?: string | null;
}

/**
 * Check if a given object implements the VideoProcessedData interface.
 */
export function instanceOfVideoProcessedData(value: object): value is VideoProcessedData {
    return true;
}

export function VideoProcessedDataFromJSON(json: any): VideoProcessedData {
    return VideoProcessedDataFromJSONTyped(json, false);
}

export function VideoProcessedDataFromJSONTyped(json: any, ignoreDiscriminator: boolean): VideoProcessedData {
    if (json == null) {
        return json;
    }
    return {
        
        'id': json['id'] == null ? undefined : json['id'],
        'videoProcessed': json['video_processed'] == null ? undefined : json['video_processed'],
        'videoThumbnailBigUrl': json['video_thumbnail_big_url'] == null ? undefined : json['video_thumbnail_big_url'],
        'videoThumbnailUrl': json['video_thumbnail_url'] == null ? undefined : json['video_thumbnail_url'],
        'videoUrl': json['video_url'] == null ? undefined : json['video_url'],
    };
}

export function VideoProcessedDataToJSON(json: any): VideoProcessedData {
    return VideoProcessedDataToJSONTyped(json, false);
}

export function VideoProcessedDataToJSONTyped(value?: VideoProcessedData | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'id': value['id'],
        'video_processed': value['videoProcessed'],
        'video_thumbnail_big_url': value['videoThumbnailBigUrl'],
        'video_thumbnail_url': value['videoThumbnailUrl'],
        'video_url': value['videoUrl'],
    };
}

