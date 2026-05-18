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
  AdditionalSettingsResponse,
  GroupNotificationSettingsResponse,
  NotificationSettingResponse,
} from '../models/index.js';
import {
    AdditionalSettingsResponseFromJSON,
    AdditionalSettingsResponseToJSON,
    GroupNotificationSettingsResponseFromJSON,
    GroupNotificationSettingsResponseToJSON,
    NotificationSettingResponseFromJSON,
    NotificationSettingResponseToJSON,
} from '../models/index.js';

export interface GetGroupNotificationSettingsRequest {
    id: number;
}

export interface UpdateChatRoomNotificationSettingsRequest {
    id: number;
    notificationChat?: number;
}

export interface UpdateGroupNotificationSettingsRequest {
    id: number;
    notificationGroupJoin?: number | null;
    notificationGroupMessageTagAll?: number | null;
    notificationGroupPost?: number | null;
    notificationGroupRequest?: number | null;
}

/**
 * 
 */
export class NotificationSettingsApi extends runtime.BaseAPI {

    /**
     */
    async getGroupNotificationSettingsRaw(requestParameters: GetGroupNotificationSettingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<GroupNotificationSettingsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getGroupNotificationSettings().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/notification_settings/groups/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => GroupNotificationSettingsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getGroupNotificationSettings(requestParameters: GetGroupNotificationSettingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<GroupNotificationSettingsResponse> {
        const response = await this.getGroupNotificationSettingsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async updateChatRoomNotificationSettingsRaw(requestParameters: UpdateChatRoomNotificationSettingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<NotificationSettingResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updateChatRoomNotificationSettings().'
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

        if (requestParameters['notificationChat'] != null) {
            formParams.append('notification_chat', requestParameters['notificationChat'] as any);
        }

        const response = await this.request({
            path: `/v2/notification_settings/chat_rooms/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => NotificationSettingResponseFromJSON(jsonValue));
    }

    /**
     */
    async updateChatRoomNotificationSettings(requestParameters: UpdateChatRoomNotificationSettingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<NotificationSettingResponse> {
        const response = await this.updateChatRoomNotificationSettingsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async updateGroupNotificationSettingsRaw(requestParameters: UpdateGroupNotificationSettingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<AdditionalSettingsResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updateGroupNotificationSettings().'
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

        if (requestParameters['notificationGroupJoin'] != null) {
            formParams.append('notification_group_join', requestParameters['notificationGroupJoin'] as any);
        }

        if (requestParameters['notificationGroupMessageTagAll'] != null) {
            formParams.append('notification_group_message_tag_all', requestParameters['notificationGroupMessageTagAll'] as any);
        }

        if (requestParameters['notificationGroupPost'] != null) {
            formParams.append('notification_group_post', requestParameters['notificationGroupPost'] as any);
        }

        if (requestParameters['notificationGroupRequest'] != null) {
            formParams.append('notification_group_request', requestParameters['notificationGroupRequest'] as any);
        }

        const response = await this.request({
            path: `/v2/notification_settings/groups/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => AdditionalSettingsResponseFromJSON(jsonValue));
    }

    /**
     */
    async updateGroupNotificationSettings(requestParameters: UpdateGroupNotificationSettingsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<AdditionalSettingsResponse> {
        const response = await this.updateGroupNotificationSettingsRaw(requestParameters, initOverrides);
        return await response.value();
    }

}