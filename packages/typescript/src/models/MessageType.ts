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
export const MessageType = {
    Text: 'text',
    Image: 'image',
    EternalImage: 'eternal_image',
    Video: 'video',
    EternalVideo: 'eternal_video',
    Screenshot: 'screenshot_warning',
    Gif: 'gif',
    Sticker: 'sticker',
    PrivateCall: 'individual_call',
    Deleted: 'deleted',
    UserJoined: 'user_joined',
    UserLeave: 'user_leave',
    OwnerKick: 'owner_kick'
} as const;
export type MessageType = typeof MessageType[keyof typeof MessageType];


export function instanceOfMessageType(value: any): boolean {
    for (const key in MessageType) {
        if (Object.prototype.hasOwnProperty.call(MessageType, key)) {
            if (MessageType[key as keyof typeof MessageType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function MessageTypeFromJSON(json: any): MessageType {
    return MessageTypeFromJSONTyped(json, false);
}

export function MessageTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MessageType {
    return json as MessageType;
}

export function MessageTypeToJSON(value?: MessageType | null): any {
    return value as any;
}

export function MessageTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): MessageType {
    return value as MessageType;
}

