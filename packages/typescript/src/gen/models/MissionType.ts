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
export const MissionType = {
    Daily: 'daily',
    Special: 'special',
    Endless: 'unlimited',
    Normal: 'once'
} as const;
export type MissionType = typeof MissionType[keyof typeof MissionType];


export function instanceOfMissionType(value: any): boolean {
    for (const key in MissionType) {
        if (Object.prototype.hasOwnProperty.call(MissionType, key)) {
            if (MissionType[key as keyof typeof MissionType] === value) {
                return true;
            }
        }
    }
    return false;
}

export function MissionTypeFromJSON(json: any): MissionType {
    return MissionTypeFromJSONTyped(json, false);
}

export function MissionTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): MissionType {
    return json as MissionType;
}

export function MissionTypeToJSON(value?: MissionType | null): any {
    return value as any;
}

export function MissionTypeToJSONTyped(value: any, ignoreDiscriminator: boolean): MissionType {
    return value as MissionType;
}

