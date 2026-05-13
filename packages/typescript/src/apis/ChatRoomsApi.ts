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
  ChatRoomResponse,
  ChatRoomsResponse,
  CreateChatRoomResponse,
  MessageResponse,
  MessageType,
  MessagesResponse,
  ReadAttachmentRequest,
  TotalChatRequestResponse,
  UnreadStatusResponse,
} from '../models/index';
import {
    ChatRoomResponseFromJSON,
    ChatRoomResponseToJSON,
    ChatRoomsResponseFromJSON,
    ChatRoomsResponseToJSON,
    CreateChatRoomResponseFromJSON,
    CreateChatRoomResponseToJSON,
    MessageResponseFromJSON,
    MessageResponseToJSON,
    MessageTypeFromJSON,
    MessageTypeToJSON,
    MessagesResponseFromJSON,
    MessagesResponseToJSON,
    ReadAttachmentRequestFromJSON,
    ReadAttachmentRequestToJSON,
    TotalChatRequestResponseFromJSON,
    TotalChatRequestResponseToJSON,
    UnreadStatusResponseFromJSON,
    UnreadStatusResponseToJSON,
} from '../models/index';

export interface AcceptChatRequestRequest {
    chatRoomIds?: Array<number>;
}

export interface CreateChatRoomRequest {
    backgroundFilename?: string | null;
    iconFilename?: string | null;
    name?: string;
    withUserIds?: Array<number>;
}

export interface CreateChatRoomV1Request {
    himaChat?: boolean;
    matchingId?: number | null;
    withUserId?: number;
}

export interface DeleteChatMessageRequest {
    roomId: number;
    messageId: number;
}

export interface DeleteChatRoomsRequest {
    chatRoomIds?: Array<number>;
}

export interface GetChatMessagesRequest {
    id: number;
    number?: number | null;
    fromMessageId?: number | null;
    toMessageId?: number | null;
}

export interface GetChatRequestsRequest {
    fromTimestamp?: number | null;
}

export interface GetChatRoomRequest {
    id: number;
}

export interface GetChatUnreadStatusRequest {
    fromTime?: number | null;
}

export interface GetMainChatRoomsRequest {
    fromTimestamp?: number | null;
}

export interface GetUpdatedChatRoomsRequest {
    fromTime?: number | null;
}

export interface InviteToChatRoomRequest {
    id: number;
    withUserIds?: Array<number>;
}

export interface KickFromChatRoomRequest {
    id: number;
    withUserIds?: Array<number>;
}

export interface PinChatRoomRequest {
    id: number;
}

export interface ReadChatAttachmentsRequest {
    id: number;
    readAttachmentRequest: ReadAttachmentRequest;
}

export interface ReadChatMessageRequest {
    id: number;
    messageId: number;
}

export interface ReadChatVideosRequest {
    id: number;
    videoMsgIds?: Array<number>;
}

export interface RemoveChatRoomBackgroundRequest {
    id: number;
}

export interface ReportChatRoomRequest {
    chatRoomId: number;
    categoryId?: number;
    opponentId?: number;
    reason?: string | null;
    screenshot2Filename?: string | null;
    screenshot3Filename?: string | null;
    screenshot4Filename?: string | null;
    screenshotFilename?: string | null;
}

export interface SendChatMessageRequest {
    id: number;
    attachmentFileName?: string | null;
    callType?: string | null;
    fontSize?: number | null;
    gifImageId?: number | null;
    messageType?: MessageType;
    parentId?: number | null;
    stickerId?: number | null;
    stickerPackId?: number | null;
    text?: string | null;
    videoFileName?: string | null;
}

export interface UnpinChatRoomRequest {
    id: number;
}

export interface UpdateChatRoomRequest {
    id: number;
    backgroundFilename?: string | null;
    iconFilename?: string | null;
    name?: string | null;
}

/**
 * 
 */
export class ChatRoomsApi extends runtime.BaseAPI {

    /**
     */
    async acceptChatRequestRaw(requestParameters: AcceptChatRequestRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['chatRoomIds'] != null) {
            formParams.append('chat_room_ids[]', requestParameters['chatRoomIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v1/chat_rooms/accept_chat_request`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async acceptChatRequest(requestParameters: AcceptChatRequestRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.acceptChatRequestRaw(requestParameters, initOverrides);
    }

    /**
     */
    async createChatRoomRaw(requestParameters: CreateChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateChatRoomResponse>> {
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

        if (requestParameters['backgroundFilename'] != null) {
            formParams.append('background_filename', requestParameters['backgroundFilename'] as any);
        }

        if (requestParameters['iconFilename'] != null) {
            formParams.append('icon_filename', requestParameters['iconFilename'] as any);
        }

        if (requestParameters['name'] != null) {
            formParams.append('name', requestParameters['name'] as any);
        }

        if (requestParameters['withUserIds'] != null) {
            formParams.append('with_user_ids[]', requestParameters['withUserIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v3/chat_rooms/new`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateChatRoomResponseFromJSON(jsonValue));
    }

    /**
     */
    async createChatRoom(requestParameters: CreateChatRoomRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateChatRoomResponse> {
        const response = await this.createChatRoomRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async createChatRoomV1Raw(requestParameters: CreateChatRoomV1Request, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<CreateChatRoomResponse>> {
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

        if (requestParameters['himaChat'] != null) {
            formParams.append('hima_chat', requestParameters['himaChat'] as any);
        }

        if (requestParameters['matchingId'] != null) {
            formParams.append('matching_id', requestParameters['matchingId'] as any);
        }

        if (requestParameters['withUserId'] != null) {
            formParams.append('with_user_id', requestParameters['withUserId'] as any);
        }

        const response = await this.request({
            path: `/v1/chat_rooms/new`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => CreateChatRoomResponseFromJSON(jsonValue));
    }

    /**
     */
    async createChatRoomV1(requestParameters: CreateChatRoomV1Request = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<CreateChatRoomResponse> {
        const response = await this.createChatRoomV1Raw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async deleteChatMessageRaw(requestParameters: DeleteChatMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['roomId'] == null) {
            throw new runtime.RequiredError(
                'roomId',
                'Required parameter "roomId" was null or undefined when calling deleteChatMessage().'
            );
        }

        if (requestParameters['messageId'] == null) {
            throw new runtime.RequiredError(
                'messageId',
                'Required parameter "messageId" was null or undefined when calling deleteChatMessage().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/{room_id}/messages/{message_id}/delete`.replace(`{${"room_id"}}`, encodeURIComponent(String(requestParameters['roomId']))).replace(`{${"message_id"}}`, encodeURIComponent(String(requestParameters['messageId']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteChatMessage(requestParameters: DeleteChatMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteChatMessageRaw(requestParameters, initOverrides);
    }

    /**
     */
    async deleteChatRoomsRaw(requestParameters: DeleteChatRoomsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
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

        if (requestParameters['chatRoomIds'] != null) {
            formParams.append('chat_room_ids[]', requestParameters['chatRoomIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v1/chat_rooms/mass_destroy`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async deleteChatRooms(requestParameters: DeleteChatRoomsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteChatRoomsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async getChatMessagesRaw(requestParameters: GetChatMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<MessagesResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getChatMessages().'
            );
        }

        const queryParameters: any = {};

        if (requestParameters['number'] != null) {
            queryParameters['number'] = requestParameters['number'];
        }

        if (requestParameters['fromMessageId'] != null) {
            queryParameters['from_message_id'] = requestParameters['fromMessageId'];
        }

        if (requestParameters['toMessageId'] != null) {
            queryParameters['to_message_id'] = requestParameters['toMessageId'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/chat_rooms/{id}/messages`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => MessagesResponseFromJSON(jsonValue));
    }

    /**
     */
    async getChatMessages(requestParameters: GetChatMessagesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<MessagesResponse> {
        const response = await this.getChatMessagesRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getChatRequestCountRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<TotalChatRequestResponse>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/total_chat_request`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => TotalChatRequestResponseFromJSON(jsonValue));
    }

    /**
     */
    async getChatRequestCount(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<TotalChatRequestResponse> {
        const response = await this.getChatRequestCountRaw(initOverrides);
        return await response.value();
    }

    /**
     */
    async getChatRequestsRaw(requestParameters: GetChatRequestsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatRoomsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/request_list`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatRoomsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getChatRequests(requestParameters: GetChatRequestsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatRoomsResponse> {
        const response = await this.getChatRequestsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getChatRoomRaw(requestParameters: GetChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatRoomResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling getChatRoom().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/chat_rooms/{id}`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatRoomResponseFromJSON(jsonValue));
    }

    /**
     */
    async getChatRoom(requestParameters: GetChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatRoomResponse> {
        const response = await this.getChatRoomRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getChatUnreadStatusRaw(requestParameters: GetChatUnreadStatusRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<UnreadStatusResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTime'] != null) {
            queryParameters['from_time'] = requestParameters['fromTime'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/unread_status`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => UnreadStatusResponseFromJSON(jsonValue));
    }

    /**
     */
    async getChatUnreadStatus(requestParameters: GetChatUnreadStatusRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<UnreadStatusResponse> {
        const response = await this.getChatUnreadStatusRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getMainChatRoomsRaw(requestParameters: GetMainChatRoomsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatRoomsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTimestamp'] != null) {
            queryParameters['from_timestamp'] = requestParameters['fromTimestamp'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/main_list`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatRoomsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getMainChatRooms(requestParameters: GetMainChatRoomsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatRoomsResponse> {
        const response = await this.getMainChatRoomsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async getUpdatedChatRoomsRaw(requestParameters: GetUpdatedChatRoomsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<ChatRoomsResponse>> {
        const queryParameters: any = {};

        if (requestParameters['fromTime'] != null) {
            queryParameters['from_time'] = requestParameters['fromTime'];
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/chat_rooms/update`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => ChatRoomsResponseFromJSON(jsonValue));
    }

    /**
     */
    async getUpdatedChatRooms(requestParameters: GetUpdatedChatRoomsRequest = {}, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<ChatRoomsResponse> {
        const response = await this.getUpdatedChatRoomsRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async inviteToChatRoomRaw(requestParameters: InviteToChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling inviteToChatRoom().'
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

        if (requestParameters['withUserIds'] != null) {
            formParams.append('with_user_ids[]', requestParameters['withUserIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v2/chat_rooms/{id}/invite`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async inviteToChatRoom(requestParameters: InviteToChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.inviteToChatRoomRaw(requestParameters, initOverrides);
    }

    /**
     */
    async kickFromChatRoomRaw(requestParameters: KickFromChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling kickFromChatRoom().'
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

        if (requestParameters['withUserIds'] != null) {
            formParams.append('with_user_ids[]', requestParameters['withUserIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v2/chat_rooms/{id}/kick`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async kickFromChatRoom(requestParameters: KickFromChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.kickFromChatRoomRaw(requestParameters, initOverrides);
    }

    /**
     */
    async pinChatRoomRaw(requestParameters: PinChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling pinChatRoom().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/{id}/pinned`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async pinChatRoom(requestParameters: PinChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.pinChatRoomRaw(requestParameters, initOverrides);
    }

    /**
     */
    async readChatAttachmentsRaw(requestParameters: ReadChatAttachmentsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling readChatAttachments().'
            );
        }

        if (requestParameters['readAttachmentRequest'] == null) {
            throw new runtime.RequiredError(
                'readAttachmentRequest',
                'Required parameter "readAttachmentRequest" was null or undefined when calling readChatAttachments().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/v1/chat_rooms/{id}/attachments_read`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: ReadAttachmentRequestToJSON(requestParameters['readAttachmentRequest']),
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async readChatAttachments(requestParameters: ReadChatAttachmentsRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.readChatAttachmentsRaw(requestParameters, initOverrides);
    }

    /**
     */
    async readChatMessageRaw(requestParameters: ReadChatMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling readChatMessage().'
            );
        }

        if (requestParameters['messageId'] == null) {
            throw new runtime.RequiredError(
                'messageId',
                'Required parameter "messageId" was null or undefined when calling readChatMessage().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/chat_rooms/{id}/messages/{message_id}/read`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))).replace(`{${"message_id"}}`, encodeURIComponent(String(requestParameters['messageId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async readChatMessage(requestParameters: ReadChatMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.readChatMessageRaw(requestParameters, initOverrides);
    }

    /**
     */
    async readChatVideosRaw(requestParameters: ReadChatVideosRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling readChatVideos().'
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

        if (requestParameters['videoMsgIds'] != null) {
            formParams.append('video_msg_ids[]', requestParameters['videoMsgIds']!.join(runtime.COLLECTION_FORMATS["csv"]));
        }

        const response = await this.request({
            path: `/v1/chat_rooms/{id}/videos_read`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async readChatVideos(requestParameters: ReadChatVideosRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.readChatVideosRaw(requestParameters, initOverrides);
    }

    /**
     */
    async removeChatRoomBackgroundRaw(requestParameters: RemoveChatRoomBackgroundRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling removeChatRoomBackground().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v2/chat_rooms/{id}/background`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async removeChatRoomBackground(requestParameters: RemoveChatRoomBackgroundRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.removeChatRoomBackgroundRaw(requestParameters, initOverrides);
    }

    /**
     */
    async reportChatRoomRaw(requestParameters: ReportChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['chatRoomId'] == null) {
            throw new runtime.RequiredError(
                'chatRoomId',
                'Required parameter "chatRoomId" was null or undefined when calling reportChatRoom().'
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

        if (requestParameters['opponentId'] != null) {
            formParams.append('opponent_id', requestParameters['opponentId'] as any);
        }

        if (requestParameters['reason'] != null) {
            formParams.append('reason', requestParameters['reason'] as any);
        }

        if (requestParameters['screenshot2Filename'] != null) {
            formParams.append('screenshot_2_filename', requestParameters['screenshot2Filename'] as any);
        }

        if (requestParameters['screenshot3Filename'] != null) {
            formParams.append('screenshot_3_filename', requestParameters['screenshot3Filename'] as any);
        }

        if (requestParameters['screenshot4Filename'] != null) {
            formParams.append('screenshot_4_filename', requestParameters['screenshot4Filename'] as any);
        }

        if (requestParameters['screenshotFilename'] != null) {
            formParams.append('screenshot_filename', requestParameters['screenshotFilename'] as any);
        }

        const response = await this.request({
            path: `/v3/chat_rooms/{chat_room_id}/report`.replace(`{${"chat_room_id"}}`, encodeURIComponent(String(requestParameters['chatRoomId']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async reportChatRoom(requestParameters: ReportChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.reportChatRoomRaw(requestParameters, initOverrides);
    }

    /**
     */
    async sendChatMessageRaw(requestParameters: SendChatMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<MessageResponse>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling sendChatMessage().'
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

        if (requestParameters['attachmentFileName'] != null) {
            formParams.append('attachment_file_name', requestParameters['attachmentFileName'] as any);
        }

        if (requestParameters['callType'] != null) {
            formParams.append('call_type', requestParameters['callType'] as any);
        }

        if (requestParameters['fontSize'] != null) {
            formParams.append('font_size', requestParameters['fontSize'] as any);
        }

        if (requestParameters['gifImageId'] != null) {
            formParams.append('gif_image_id', requestParameters['gifImageId'] as any);
        }

        if (requestParameters['messageType'] != null) {
            formParams.append('message_type', requestParameters['messageType'] as any);
        }

        if (requestParameters['parentId'] != null) {
            formParams.append('parent_id', requestParameters['parentId'] as any);
        }

        if (requestParameters['stickerId'] != null) {
            formParams.append('sticker_id', requestParameters['stickerId'] as any);
        }

        if (requestParameters['stickerPackId'] != null) {
            formParams.append('sticker_pack_id', requestParameters['stickerPackId'] as any);
        }

        if (requestParameters['text'] != null) {
            formParams.append('text', requestParameters['text'] as any);
        }

        if (requestParameters['videoFileName'] != null) {
            formParams.append('video_file_name', requestParameters['videoFileName'] as any);
        }

        const response = await this.request({
            path: `/v3/chat_rooms/{id}/messages/new`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => MessageResponseFromJSON(jsonValue));
    }

    /**
     */
    async sendChatMessage(requestParameters: SendChatMessageRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<MessageResponse> {
        const response = await this.sendChatMessageRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     */
    async unpinChatRoomRaw(requestParameters: UnpinChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling unpinChatRoom().'
            );
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/chat_rooms/{id}/pinned`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async unpinChatRoom(requestParameters: UnpinChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.unpinChatRoomRaw(requestParameters, initOverrides);
    }

    /**
     */
    async updateChatRoomRaw(requestParameters: UpdateChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters['id'] == null) {
            throw new runtime.RequiredError(
                'id',
                'Required parameter "id" was null or undefined when calling updateChatRoom().'
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

        if (requestParameters['backgroundFilename'] != null) {
            formParams.append('background_filename', requestParameters['backgroundFilename'] as any);
        }

        if (requestParameters['iconFilename'] != null) {
            formParams.append('icon_filename', requestParameters['iconFilename'] as any);
        }

        if (requestParameters['name'] != null) {
            formParams.append('name', requestParameters['name'] as any);
        }

        const response = await this.request({
            path: `/v3/chat_rooms/{id}/edit`.replace(`{${"id"}}`, encodeURIComponent(String(requestParameters['id']))),
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: formParams,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     */
    async updateChatRoom(requestParameters: UpdateChatRoomRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.updateChatRoomRaw(requestParameters, initOverrides);
    }

}
