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
 * @interface PolicyAgreementsResponse
 */
export interface PolicyAgreementsResponse {
    /**
     * 
     * @type {boolean}
     * @memberof PolicyAgreementsResponse
     */
    latestDotmoneyAgreed?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PolicyAgreementsResponse
     */
    latestEmpl2Agreed?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PolicyAgreementsResponse
     */
    latestPrivacyPolicyAgreed?: boolean | null;
    /**
     * 
     * @type {boolean}
     * @memberof PolicyAgreementsResponse
     */
    latestTermsOfUseAgreed?: boolean | null;
}

/**
 * Check if a given object implements the PolicyAgreementsResponse interface.
 */
export function instanceOfPolicyAgreementsResponse(value: object): value is PolicyAgreementsResponse {
    return true;
}

export function PolicyAgreementsResponseFromJSON(json: any): PolicyAgreementsResponse {
    return PolicyAgreementsResponseFromJSONTyped(json, false);
}

export function PolicyAgreementsResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PolicyAgreementsResponse {
    if (json == null) {
        return json;
    }
    return {
        
        'latestDotmoneyAgreed': json['latest_dotmoney_agreed'] == null ? undefined : json['latest_dotmoney_agreed'],
        'latestEmpl2Agreed': json['latest_empl2_agreed'] == null ? undefined : json['latest_empl2_agreed'],
        'latestPrivacyPolicyAgreed': json['latest_privacy_policy_agreed'] == null ? undefined : json['latest_privacy_policy_agreed'],
        'latestTermsOfUseAgreed': json['latest_terms_of_use_agreed'] == null ? undefined : json['latest_terms_of_use_agreed'],
    };
}

export function PolicyAgreementsResponseToJSON(json: any): PolicyAgreementsResponse {
    return PolicyAgreementsResponseToJSONTyped(json, false);
}

export function PolicyAgreementsResponseToJSONTyped(value?: PolicyAgreementsResponse | null, ignoreDiscriminator: boolean = false): any {
    if (value == null) {
        return value;
    }

    return {
        
        'latest_dotmoney_agreed': value['latestDotmoneyAgreed'],
        'latest_empl2_agreed': value['latestEmpl2Agreed'],
        'latest_privacy_policy_agreed': value['latestPrivacyPolicyAgreed'],
        'latest_terms_of_use_agreed': value['latestTermsOfUseAgreed'],
    };
}

