import type { NotificationMessageService } from '#/rpc/api/admin/service/v1/i_notification_message.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type {
  CreateNotificationMessageRequest,
  DeleteNotificationMessageRequest,
  GetNotificationMessageRequest,
  ListNotificationMessageResponse,
  NotificationMessage,
  UpdateNotificationMessageRequest,
} from '#/rpc/api/internal_message/service/v1/notification_message.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';

import { requestClient } from '#/rpc/request';

/** 通知消息管理服务 */
class NotificationMessageServiceImpl implements NotificationMessageService {
  async CreateNotificationMessage(
    request: CreateNotificationMessageRequest,
  ): Promise<Empty> {
    return await requestClient.post<Empty>('/notifications', request);
  }

  async DeleteNotificationMessage(
    request: DeleteNotificationMessageRequest,
  ): Promise<Empty> {
    return await requestClient.delete<Empty>(`/notifications/${request.id}`);
  }

  async GetNotificationMessage(
    request: GetNotificationMessageRequest,
  ): Promise<NotificationMessage> {
    return await requestClient.get<NotificationMessage>(
      `/notifications/${request.id}`,
    );
  }

  async ListNotificationMessage(
    request: PagingRequest,
  ): Promise<ListNotificationMessageResponse> {
    return await requestClient.get<ListNotificationMessageResponse>(
      '/notifications',
      {
        params: request,
      },
    );
  }

  async UpdateNotificationMessage(
    request: UpdateNotificationMessageRequest,
  ): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateNotificationMessage', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/notifications/${id}`, request);
  }
}

export const defNotificationMessageService =
  new NotificationMessageServiceImpl();
