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
  BgmsResponse,
  CallActionSignatureResponse,
  CallGiftHistoryResponse,
  CallStatusResponse,
  ConferenceCallResponse,
  RtmTokenResponse,
  UsersByTimestampResponse,
} from '../models/index.js';
import {
    BgmsResponseFromJSON,
    BgmsResponseToJSON,
    CallActionSignatureResponseFromJSON,
    CallActionSignatureResponseToJSON,
    CallGiftHistoryResponseFromJSON,
    CallGiftHistoryResponseToJSON,
    CallStatusResponseFromJSON,
    CallStatusResponseToJSON,
    ConferenceCallResponseFromJSON,
    ConferenceCallResponseToJSON,
    RtmTokenResponseFromJSON,
    RtmTokenResponseToJSON,
    UsersByTimestampResponseFromJSON,
    UsersByTimestampResponseToJSON,
} from '../models/index.js';

export interface BulkInviteToCallRequest {
    callId: number;
    groupId?: number | null;
}

export interface BumpCallRequest {
    callId: number;
    participantLimit?: number | null;
}

export interface GenerateCallActionSignatureRequest {
    action?: string;
    conferenceId?: number;
    targetUserUuid?: string;
}

export interface GetAgoraRtmTokenRequest {
    callId: number;
}

export interface GetCallGiftHistoryRequest {
    callId: number;
    from?: string | null;
    number?: number | null;
}

export interface GetConferenceCallRequest {
    callId: number;
}

export interface GetInvitableCallUsersRequest {
    callId: number;
    fromTimestamp?: number | null;
    userNickname?: string | null;
}

export interface GetPhoneStatusRequest {
    opponentId: number;
}

export interface InviteToCallRequest {
    chatRoomId?: number;
    roomId?: number;
    roomUrl?: string;
}

export interface InviteToConferenceCallRequest {
    callId: number;
    userIds?: Array<number>;
}

export interface KickFromConferenceCallRequest {
    callId: number;
    ban?: boolean;
    uuid?: string;
}

export interface LeaveAgoraChannelRequest {
    conferenceId?: number;
    userId?: number;
}

export interface LeaveConferenceCallRequest {
    callSid?: string;
    conferenceId?: number;
    duration?: number | null;
    totalUsersInCall?: number | null;
}

export interface StartConferenceCallRequest {
    callSid?: string;
    conferenceId?: number;
}

export interface UpdateCallRequest {
    callId: number;
    categoryId?: string | null;
    gameTitle?: string | null;
    joinableBy?: string;
}

export interface UpdateCallUserRequest {
    callId: number;
    userId: number;
    role?: string;
}

export interface ValidateCallActionSignatureRequest {
    callId?: number;
    senderUuid?: string;
    receiverUuid?: string;
    action?: string;
    timestamp?: number;
    signature?: string;
}

/**
 * 
 */
export class CallsApi extends runtime.BaseAPI {

    /**
     */
    async bulkInviteToCallRaw(requestParameters: BulkInviteToCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling bulkInviteToCall().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['groupId'] != null) {
            queryParameters['group_id'] = requestParameters['groupId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/calls/{call_id}/bulk_invite`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async bulkInviteToCall(requestParameters: BulkInviteToCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.bulkInviteToCallRaw(requestParameters, initOverrides);
    }

    /**
     */
    async bumpCallRaw(requestParameters: BumpCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling bumpCall().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['participantLimit'] != null) {
            queryParameters['participant_limit'] = requestParameters['participantLimit'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/calls/{call_id}/bump`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async bumpCall(requestParameters: BumpCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.bumpCallRaw(requestParameters, initOverrides);
    }

    /**
     */
    async generateCallActionSignatureRaw(requestParameters: GenerateCallActionSignatureRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CallActionSignatureResponse>> {
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

        if (requestParameters['action'] != null) {
            formParams.append('action', requestParameters['action'] as any);
        }

        if (requestParameters['conferenceId'] != null) {
            formParams.append('conference_id', requestParameters['conferenceId'] as any);
        }

        if (requestParameters['targetUserUuid'] != null) {
            formParams.append('target_user_uuid', requestParameters['targetUserUuid'] as any);
        }

        const response = await this.request({
            path: `/v1/calls/action_signature/generate`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CallActionSignatureResponseFromJSON(jsonValue));
    }

    /**
     */
    async generateCallActionSignature(requestParameters: GenerateCallActionSignatureRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CallActionSignatureResponse> {
        const response = await this.generateCallActionSignatureRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getAgoraRtmTokenRaw(requestParameters: GetAgoraRtmTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<RtmTokenResponse>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling getAgoraRtmToken().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v3/calls/{call_id}/agora_rtm_token`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => RtmTokenResponseFromJSON(jsonValue));
    }

    /**
     */
    async getAgoraRtmToken(requestParameters: GetAgoraRtmTokenRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<RtmTokenResponse> {
        const response = await this.getAgoraRtmTokenRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getCallBgmsRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<BgmsResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/calls/bgm`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => BgmsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getCallBgms(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<BgmsResponse> {
        const response = await this.getCallBgmsRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getCallGiftHistoryRaw(requestParameters: GetCallGiftHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CallGiftHistoryResponse>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling getCallGiftHistory().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['from'] != null) {
            queryParameters['from'] = requestParameters['from'];
        }

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/calls/{call_id}/gift_transactions`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CallGiftHistoryResponseFromJSON(jsonValue));
    }

    /**
     */
    async getCallGiftHistory(requestParameters: GetCallGiftHistoryRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CallGiftHistoryResponse> {
        const response = await this.getCallGiftHistoryRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getConferenceCallRaw(requestParameters: GetConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ConferenceCallResponse>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling getConferenceCall().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/calls/conferences/{call_id}`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ConferenceCallResponseFromJSON(jsonValue));
    }

    /**
     */
    async getConferenceCall(requestParameters: GetConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ConferenceCallResponse> {
        const response = await this.getConferenceCallRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getInvitableCallUsersRaw(requestParameters: GetInvitableCallUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UsersByTimestampResponse>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling getInvitableCallUsers().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        if (requestParameters['userNickname'] != null) {
            queryParameters['user[nickname]'] = requestParameters['userNickname'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/calls/{call_id}/users/invitable`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UsersByTimestampResponseFromJSON(jsonValue));
    }

    /**
     */
    async getInvitableCallUsers(requestParameters: GetInvitableCallUsersRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UsersByTimestampResponse> {
        const response = await this.getInvitableCallUsersRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getPhoneStatusRaw(requestParameters: GetPhoneStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CallStatusResponse>> {
        if (requestParameters['opponentId'] == null) {
            throw new runtime.RequiredError(
                'opponentId',
                'Required parameter "opponentId" was null or undefined when calling getPhoneStatus().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/calls/phone_status/{opponent_id}`.replace(`{${"opponent_id"}}`, encodeURIComponent(String(requestParameters['opponentId']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CallStatusResponseFromJSON(jsonValue));
    }

    /**
     */
    async getPhoneStatus(requestParameters: GetPhoneStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CallStatusResponse> {
        const response = await this.getPhoneStatusRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async inviteToCallRaw(requestParameters: InviteToCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['chatRoomId'] != null) {
            formParams.append('chat_room_id', requestParameters['chatRoomId'] as any);
        }

        if (requestParameters['roomId'] != null) {
            formParams.append('room_id', requestParameters['roomId'] as any);
        }

        if (requestParameters['roomUrl'] != null) {
            formParams.append('room_url', requestParameters['roomUrl'] as any);
        }

        const response = await this.request({
            path: `/v2/calls/invite`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async inviteToCall(requestParameters: InviteToCallRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.inviteToCallRaw(requestParameters, initOverrides);
    }

    /**
     */
    async inviteToConferenceCallRaw(requestParameters: InviteToConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling inviteToConferenceCall().'
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

        if (requestParameters['userIds'] != null) {
            formParams.append('user_ids[]', requestParameters['userIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v1/calls/conference_calls/{call_id}/invite`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async inviteToConferenceCall(requestParameters: InviteToConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.inviteToConferenceCallRaw(requestParameters, initOverrides);
    }

    /**
     */
    async kickFromConferenceCallRaw(requestParameters: KickFromConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CallActionSignatureResponse>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling kickFromConferenceCall().'
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

        if (requestParameters['ban'] != null) {
            formParams.append('ban', requestParameters['ban'] as any);
        }

        if (requestParameters['uuid'] != null) {
            formParams.append('uuid', requestParameters['uuid'] as any);
        }

        const response = await this.request({
            path: `/v3/calls/conference_calls/{call_id}/kick`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CallActionSignatureResponseFromJSON(jsonValue));
    }

    /**
     */
    async kickFromConferenceCall(requestParameters: KickFromConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CallActionSignatureResponse> {
        const response = await this.kickFromConferenceCallRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async leaveAgoraChannelRaw(requestParameters: LeaveAgoraChannelRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['conferenceId'] != null) {
            formParams.append('conference_id', requestParameters['conferenceId'] as any);
        }

        if (requestParameters['userId'] != null) {
            formParams.append('user_id', requestParameters['userId'] as any);
        }

        const response = await this.request({
            path: `/v1/calls/leave_agora_channel`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async leaveAgoraChannel(requestParameters: LeaveAgoraChannelRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.leaveAgoraChannelRaw(requestParameters, initOverrides);
    }

    /**
     */
    async leaveConferenceCallRaw(requestParameters: LeaveConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['callSid'] != null) {
            formParams.append('call_sid', requestParameters['callSid'] as any);
        }

        if (requestParameters['conferenceId'] != null) {
            formParams.append('conference_id', requestParameters['conferenceId'] as any);
        }

        if (requestParameters['duration'] != null) {
            formParams.append('duration', requestParameters['duration'] as any);
        }

        if (requestParameters['totalUsersInCall'] != null) {
            formParams.append('total_users_in_call', requestParameters['totalUsersInCall'] as any);
        }

        const response = await this.request({
            path: `/v1/calls/leave_conference_call`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async leaveConferenceCall(requestParameters: LeaveConferenceCallRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.leaveConferenceCallRaw(requestParameters, initOverrides);
    }

    /**
     */
    async startConferenceCallRaw(requestParameters: StartConferenceCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ConferenceCallResponse>> {
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

        if (requestParameters['callSid'] != null) {
            formParams.append('call_sid', requestParameters['callSid'] as any);
        }

        if (requestParameters['conferenceId'] != null) {
            formParams.append('conference_id', requestParameters['conferenceId'] as any);
        }

        const response = await this.request({
            path: `/v2/calls/start_conference_call`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ConferenceCallResponseFromJSON(jsonValue));
    }

    /**
     */
    async startConferenceCall(requestParameters: StartConferenceCallRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ConferenceCallResponse> {
        const response = await this.startConferenceCallRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async updateCallRaw(requestParameters: UpdateCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling updateCall().'
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

        if (requestParameters['categoryId'] != null) {
            formParams.append('category_id', requestParameters['categoryId'] as any);
        }

        if (requestParameters['gameTitle'] != null) {
            formParams.append('game_title', requestParameters['gameTitle'] as any);
        }

        if (requestParameters['joinableBy'] != null) {
            formParams.append('joinable_by', requestParameters['joinableBy'] as any);
        }

        const response = await this.request({
            path: `/v1/calls/{call_id}`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateCall(requestParameters: UpdateCallRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateCallRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateCallUserRaw(requestParameters: UpdateCallUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['callId'] == null) {
            throw new runtime.RequiredError(
                'callId',
                'Required parameter "callId" was null or undefined when calling updateCallUser().'
            );
        }

        if (requestParameters['userId'] == null) {
            throw new runtime.RequiredError(
                'userId',
                'Required parameter "userId" was null or undefined when calling updateCallUser().'
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

        if (requestParameters['role'] != null) {
            formParams.append('role', requestParameters['role'] as any);
        }

        const response = await this.request({
            path: `/v1/calls/{call_id}/users/{user_id}`.replace(`{${"call_id"}}`, encodeURIComponent(String(requestParameters['callId']))).replace(`{${"user_id"}}`, encodeURIComponent(String(requestParameters['userId']))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateCallUser(requestParameters: UpdateCallUserRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateCallUserRaw(requestParameters, initOverrides);
    }

    /**
     */
    async validateCallActionSignatureRaw(requestParameters: ValidateCallActionSignatureRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        if (requestParameters['callId'] != null) {
            queryParameters['call_id'] = requestParameters['callId'];
        }

        if (requestParameters['senderUuid'] != null) {
            queryParameters['sender_uuid'] = requestParameters['senderUuid'];
        }

        if (requestParameters['receiverUuid'] != null) {
            queryParameters['receiver_uuid'] = requestParameters['receiverUuid'];
        }

        if (requestParameters['action'] != null) {
            queryParameters['action'] = requestParameters['action'];
        }

        if (requestParameters['timestamp'] != null) {
            queryParameters['timestamp'] = requestParameters['timestamp'];
        }

        if (requestParameters['signature'] != null) {
            queryParameters['signature'] = requestParameters['signature'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/calls/action_signature/validate`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async validateCallActionSignature(requestParameters: ValidateCallActionSignatureRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.validateCallActionSignatureRaw(requestParameters, initOverrides);
    }

}