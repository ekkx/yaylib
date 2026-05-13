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
export const OnlineStatusEnum = {
    Offline: 'offline',
    Hidden: 'hidden'
} as const;
export type OnlineStatusEnum = typeof OnlineStatusEnum[keyof typeof OnlineStatusEnum];


export function instanceOfOnlineStatusEnum(value: any): boolean {
    for (const key in OnlineStatusEnum) {
        if (Object.prototype.hasOwnProperty.call(OnlineStatusEnum, key)) {
            if (OnlineStatusEnum[key as keyof typeof OnlineStatusEnum] === value) {
                return true;
            }
        }
    }
    return false;
}

export function OnlineStatusEnumFromJSON(json: any): OnlineStatusEnum {
    return OnlineStatusEnumFromJSONTyped(json, false);
}

export function OnlineStatusEnumFromJSONTyped(json: any, ignoreDiscriminator: boolean): OnlineStatusEnum {
    return json as OnlineStatusEnum;
}

export function OnlineStatusEnumToJSON(value?: OnlineStatusEnum | null): any {
    return value as any;
}

export function OnlineStatusEnumToJSONTyped(value: any, ignoreDiscriminator: boolean): OnlineStatusEnum {
    return value as OnlineStatusEnum;
}

