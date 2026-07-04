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
export const CallType = {
    Video: 'vdo',
    Voice: 'voice'
} as const;
export type CallType = typeof CallType[keyof typeof CallType];


export function instanceOfCallType(value: any): boolean {
    for (const key in CallType) {
        if (Object.prototype.hasOwnProperty.call(CallType, key)) {
            if (CallType[key as keyof typeof CallType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function CallTypeFromJSON(json: any): CallType {
    return CallTypeFromJSONTyped(json, false);
}

export function CallTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): CallType {
    return json as CallType;
}

export function CallTypeToJSON(value?: CallType | null): any {
    return value as any;
}

export function CallTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): CallType {
    return value as CallType;
}

