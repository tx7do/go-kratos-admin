import type { InternalMessageService } from '#/generated/api/admin/service/v1/i_internal_message.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type {
  DeleteInternalMessageRequest,
  GetInternalMessageRequest,
  InternalMessage,
  ListInternalMessageResponse,
  RevokeMessageRequest,
  SendMessageRequest,
  SendMessageResponse,
  UpdateInternalMessageRequest,
} from '#/generated/api/internal_message/service/v1/internal_message.pb';
import type {
  DeleteNotificationFromInboxRequest,
  ListUserInboxResponse,
  MarkNotificationAsReadRequest,
} from '#/generated/api/internal_message/service/v1/internal_message_recipient.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 通知消息管理服务 */
class InternalMessageServiceImpl implements InternalMessageService {
  async DeleteMessage(request: DeleteInternalMessageRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(
      `/internal-message/messages/${request.id}`,
    );
  }

  async DeleteNotificationFromInbox(
    request: DeleteNotificationFromInboxRequest,
  ): Promise<Empty> {
    return await requestClient.post<Empty>('/internal-message/inbox/delete', {
      params: request,
    });
  }

  async GetMessage(
    request: GetInternalMessageRequest,
  ): Promise<InternalMessage> {
    switch (request.queryBy?.$case) {
      case 'id': {
        return await requestClient.get<InternalMessage>(
          `/internal-message/messages/${request.queryBy.id}`,
        );
      }
    }
    throw new Error('GetInternalMessage must set queryBy');
  }

  async ListMessage(
    request: PagingRequest,
  ): Promise<ListInternalMessageResponse> {
    return await requestClient.get<ListInternalMessageResponse>(
      '/internal-message/messages',
      {
        params: request,
      },
    );
  }

  async ListUserInbox(request: PagingRequest): Promise<ListUserInboxResponse> {
    return await requestClient.get<ListUserInboxResponse>(
      '/internal-message/inbox',
      {
        params: request,
      },
    );
  }

  async MarkNotificationAsRead(
    request: MarkNotificationAsReadRequest,
  ): Promise<Empty> {
    return await requestClient.post<Empty>('/internal-message/read', request);
  }

  async RevokeMessage(request: RevokeMessageRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/internal-message/revoke', request);
  }

  async SendMessage(request: SendMessageRequest): Promise<SendMessageResponse> {
    return await requestClient.post<SendMessageResponse>(
      '/internal-message/send',
      request,
    );
  }

  async UpdateMessage(request: UpdateInternalMessageRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(
      `/internal-message/messages/${id}`,
      request,
    );
  }
}

export const defInternalMessageService = new InternalMessageServiceImpl();
