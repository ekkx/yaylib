// @ts-nocheck
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


import * as runtime from '../runtime.js';
import type {
  VoteSurveyResponse,
} from '../models/index.js';
import {
    VoteSurveyResponseFromJSON,
    VoteSurveyResponseToJSON,
} from '../models/index.js';

export interface VoteSurveyRequest {
    id: number;
    choiceId?: number;
}

/**
 * 
 */
export class SurveysApi extends runtime.BaseAPI {

    /**
     */
    async voteSurveyRaw(requestParameters: VoteSurveyRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<VoteSurveyResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling voteSurvey().'
            );
        }

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

        if (requestParameters['choiceId'] != null) {
            formParams.append('choice_id', requestParameters['choiceId'] as any);
        }

        const response = await this.request({
            path: `/v2/surveys/{id}/vote`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => VoteSurveyResponseFromJSON(jsonValue));
    }

    /**
     */
    async voteSurvey(requestParameters: VoteSurveyRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<VoteSurveyResponse> {
        const response = await this.voteSurveyRaw(requestParameters, initOverrides);
        return await response.value();
    }

}