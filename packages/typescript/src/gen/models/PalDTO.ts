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
import type { Attribute } from './Attribute.js';
import {
    AttributeFromJSON,
    AttributeFromJSONTyped,
    AttributeToJSON,
    AttributeToJSONTyped,
} from './Attribute.js';
import type { MaxAttribute } from './MaxAttribute.js';
import {
    MaxAttributeFromJSON,
    MaxAttributeFromJSONTyped,
    MaxAttributeToJSON,
    MaxAttributeToJSONTyped,
} from './MaxAttribute.js';

/**
 * 
 * @export
 * @interface PalDTO
 */
export interface PalDTO {
    /**
     * 
     * @type {number}
     * @memberof PalDTO
     */
    ageInDays?: number | null;
    /**
     * 
     * @type {Attribute}
     * @memberof PalDTO
     */
    attributes?: Attribute | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    currentEmplEarned?: string | null;
    /**
     * 
     * @type {number}
     * @memberof PalDTO
     */
    currentLevel?: number | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    description?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    emplEarnedLimit?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalDTO
     */
    fromLiquidityPool?: boolean | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    hatchingAnimationUrl?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    image?: string | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalDTO
     */
    inLockedParty?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalDTO
     */
    isAlive?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalDTO
     */
    isPending?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalDTO
     */
    isValidToEvolve?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PalDTO
     */
    isValidToLevelUp?: boolean | null;
    /**
     * 
     * @type {number}
     * @memberof PalDTO
     */
    maxAgeInDays?: number | null;
    /**
     * 
     * @type {MaxAttribute}
     * @memberof PalDTO
     */
    maxAttributes?: MaxAttribute | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    name?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    originalWalletAddress?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    raceStatus?: string | null;
    /**
     * 
     * @type {string}
     * @memberof PalDTO
     */
    tokenId?: string | null;
}

/**
 * Check if a given object implements the PalDTO interface.
 */
export function instanceOfPalDTO(value: object): value is PalDTO {
    return true;
}

export function PalDTOFromJSON(json: any): PalDTO {
    return PalDTOFromJSONTyped(json, false);
}

export function PalDTOFromJSONTyped(json: any, ignoreDiscriminator: boolean): PalDTO {
    if (json == null) {
        return json;
    }
    return {
        
        'ageInDays': json['age_in_days'] == null ? undefined : json['age_in_days'],
        'attributes': json['attributes'] == null ? undefined : AttributeFromJSON(json['attributes']),
        'currentEmplEarned': json['current_empl_earned'] == null ? undefined : json['current_empl_earned'],
        'currentLevel': json['current_level'] == null ? undefined : json['current_level'],
        'description': json['description'] == null ? undefined : json['description'],
        'emplEarnedLimit': json['empl_earned_limit'] == null ? undefined : json['empl_earned_limit'],
        'fromLiquidityPool': json['from_liquidity_pool'] == null ? undefined : json['from_liquidity_pool'],
        'hatchingAnimationUrl': json['hatching_animation_url'] == null ? undefined : json['hatching_animation_url'],
        'image': json['image'] == null ? undefined : json['image'],
        'inLockedParty': json['in_locked_party'] == null ? undefined : json['in_locked_party'],
        'isAlive': json['is_alive'] == null ? undefined : json['is_alive'],
        'isPending': json['is_pending'] == null ? undefined : json['is_pending'],
        'isValidToEvolve': json['is_valid_to_evolve'] == null ? undefined : json['is_valid_to_evolve'],
        'isValidToLevelUp': json['is_valid_to_level_up'] == null ? undefined : json['is_valid_to_level_up'],
        'maxAgeInDays': json['max_age_in_days'] == null ? undefined : json['max_age_in_days'],
        'maxAttributes': json['max_attributes'] == null ? undefined : MaxAttributeFromJSON(json['max_attributes']),
        'name': json['name'] == null ? undefined : json['name'],
        'originalWalletAddress': json['original_wallet_address'] == null ? undefined : json['original_wallet_address'],
        'raceStatus': json['race_status'] == null ? undefined : json['race_status'],
        'tokenId': json['token_id'] == null ? undefined : json['token_id'],
    };
}

export function PalDTOToJSON(json: any): PalDTO {
    return PalDTOToJSONTyped(json, false);
}

export function PalDTOToJSONTyped(value?: PalDTO | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'age_in_days': value['ageInDays'],
        'attributes': AttributeToJSON(value['attributes']),
        'current_empl_earned': value['currentEmplEarned'],
        'current_level': value['currentLevel'],
        'description': value['description'],
        'empl_earned_limit': value['emplEarnedLimit'],
        'from_liquidity_pool': value['fromLiquidityPool'],
        'hatching_animation_url': value['hatchingAnimationUrl'],
        'image': value['image'],
        'in_locked_party': value['inLockedParty'],
        'is_alive': value['isAlive'],
        'is_pending': value['isPending'],
        'is_valid_to_evolve': value['isValidToEvolve'],
        'is_valid_to_level_up': value['isValidToLevelUp'],
        'max_age_in_days': value['maxAgeInDays'],
        'max_attributes': MaxAttributeToJSON(value['maxAttributes']),
        'name': value['name'],
        'original_wallet_address': value['originalWalletAddress'],
        'race_status': value['raceStatus'],
        'token_id': value['tokenId'],
    };
}

