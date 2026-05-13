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


/**
 * 
 * @export
 */
export const PostType = {
    Call: 'call',
    Image: 'image',
    Video: 'video',
    Survey: 'survey',
    Repost: 'repost',
    Thread: 'thread',
    ShareableGroup: 'shareable_group',
    ShareableUrl: 'shareable_url',
    Youtube: 'youtube',
    ShareableThread: 'shareable_thread',
    SharePal: 'sharepal',
    Undefined: 'undefined',
    Text: 'text'
} as const;
export type PostType = typeof PostType[keyof typeof PostType];


export function instanceOfPostType(value: any): boolean {
    for (const key in PostType) {
        if (Object.prototype.hasOwnProperty.call(PostType, key)) {
            if (PostType[key as keyof typeof PostType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PostTypeFromJSON(json: any): PostType {
    return PostTypeFromJSONTyped(json, false);
}

export function PostTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PostType {
    return json as PostType;
}

export function PostTypeToJSON(value?: PostType | null): any {
    return value as any;
}

export function PostTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): PostType {
    return value as PostType;
}

