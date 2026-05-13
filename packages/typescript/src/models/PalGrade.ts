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
export const PalGrade = {
    Egg: 'Egg',
    Normal: 'Pal',
    Super: 'Super Pal',
    UltimatePal: 'Ultimate Pal'
} as const;
export type PalGrade = typeof PalGrade[keyof typeof PalGrade];


export function instanceOfPalGrade(value: any): boolean {
    for (const key in PalGrade) {
        if (Object.prototype.hasOwnProperty.call(PalGrade, key)) {
            if (PalGrade[key as keyof typeof PalGrade] === value) {
                return true;
            }
        }
    }
    return false;
}

export function PalGradeFromJSON(json: any): PalGrade {
    return PalGradeFromJSONTyped(json, false);
}

export function PalGradeFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalGrade {
    return json as PalGrade;
}

export function PalGradeToJSON(value?: PalGrade | null): any {
    return value as any;
}

export function PalGradeToJSONTyped(value: any, ignoreDiscriminator: boolean): PalGrade {
    return value as PalGrade;
}

