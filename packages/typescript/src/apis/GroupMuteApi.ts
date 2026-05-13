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
  GroupMuteUsersResponse,
} from '../models/index';
import {
    GroupMuteUsersResponseFromJSON,
    GroupMuteUsersResponseToJSON,
} from '../models/index';

export interface ListMutedGroupUsersRequest {
    id: number;
    keyword?: string | null;
    cursor?: string;
    size?: number;
}

export interface MuteGroupUserRequest {
    id: number;
    userId: number;
}

export interface UnmuteGroupUserRequest {
    id: number;
    userId: number;
}

/**
 * 
 */
export class GroupMuteApi extends runtime.BaseAPI {

    /**
     */
    async listMutedGroupUsersRaw(requestParameters: ListMutedGroupUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupMuteUsersResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling listMutedGroupUsers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['keyword'] != null) {
            queryParameters['keyword'] = requestParameters['keyword'];
        }

        if (requestParameters['cursor'] != null) {
            queryParameters['cursor'] = requestParameters['cursor'];
        }

        if (requestParameters['size'] != null) {
            queryParameters['size'] = requestParameters['size'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/group_mute/{id}/muted_users`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupMuteUsersResponseFromJSON(jsonValue));
    }

    /**
     */
    async listMutedGroupUsers(requestParameters: ListMutedGroupUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupMuteUsersResponse> {
        const response = await this.listMutedGroupUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async muteGroupUserRaw(requestParameters: MuteGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling muteGroupUser().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling muteGroupUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/group_mute/{id}/mute/{user_id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async muteGroupUser(requestParameters: MuteGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.muteGroupUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async unmuteGroupUserRaw(requestParameters: UnmuteGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling unmuteGroupUser().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling unmuteGroupUser().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/group_mute/{id}/unmute/{user_id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unmuteGroupUser(requestParameters: UnmuteGroupUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unmuteGroupUserRaw(requestParameters, initOverrides);
    }

}
