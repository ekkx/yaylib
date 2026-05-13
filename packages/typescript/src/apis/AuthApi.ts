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
  TokenResponse,
} from '../models/index';
import {
    TokenResponseFromJSON,
    TokenResponseToJSON,
} from '../models/index';

export interface OauthTokenRequest {
    email?: string | null;
    grantType?: string;
    password?: string | null;
    refreshToken?: string | null;
}

/**
 * 
 */
export class AuthApi extends runtime.BaseAPI {

    /**
     */
    async oauthTokenRaw(requestParameters: OauthTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<TokenResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const consumes: runtime.Consume[] = [
            { contentType: 'application/x-www-form-urlencoded' },
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

        if (requestParameters['email'] != null) {
            formParams.append('email', requestParameters['email'] as any);
        }

        if (requestParameters['grantType'] != null) {
            formParams.append('grant_type', requestParameters['grantType'] as any);
        }

        if (requestParameters['password'] != null) {
            formParams.append('password', requestParameters['password'] as any);
        }

        if (requestParameters['refreshToken'] != null) {
            formParams.append('refresh_token', requestParameters['refreshToken'] as any);
        }

        const response = await this.request({
            path: `/api/v1/oauth/token`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => TokenResponseFromJSON(jsonValue));
    }

    /**
     */
    async oauthToken(requestParameters: OauthTokenRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<TokenResponse> {
        const response = await this.oauthTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
