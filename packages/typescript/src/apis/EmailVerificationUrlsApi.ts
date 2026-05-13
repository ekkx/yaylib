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


import * as runtime from '../runtime';
import type {
  CommonUrlResponse,
} from '../models/index';
import {
    CommonUrlResponseFromJSON,
    CommonUrlResponseToJSON,
} from '../models/index';

export interface RequestEmailVerificationRequest {
    deviceUuid?: string;
    email?: string;
    intent?: string | null;
    locale?: string;
}

/**
 * 
 */
export class EmailVerificationUrlsApi extends runtime.BaseAPI {

    /**
     */
    async requestEmailVerificationRaw(requestParameters: RequestEmailVerificationRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CommonUrlResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const consumes: runtime.Consume[] = [
            { contentType: 'multipart/form-data' },
        ];
        // @ts-ignore: canConsumeForm may be unused
        const canConsumeForm = runtime.canConsumeForm(consumes);

        let formParams: { append(param: string, value: any): any };
        let useForm = false;
        if (useForm) {
            formParams = new FormData();
        } else {
            formParams = new URLSearchParams();
        }

        if (requestParameters['deviceUuid'] != null) {
            formParams.append('device_uuid', requestParameters['deviceUuid'] as any);
        }

        if (requestParameters['email'] != null) {
            formParams.append('email', requestParameters['email'] as any);
        }

        if (requestParameters['intent'] != null) {
            formParams.append('intent', requestParameters['intent'] as any);
        }

        if (requestParameters['locale'] != null) {
            formParams.append('locale', requestParameters['locale'] as any);
        }

        const response = await this.request({
            path: `/v1/email_verification_urls`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CommonUrlResponseFromJSON(jsonValue));
    }

    /**
     */
    async requestEmailVerification(requestParameters: RequestEmailVerificationRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CommonUrlResponse> {
        const response = await this.requestEmailVerificationRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
