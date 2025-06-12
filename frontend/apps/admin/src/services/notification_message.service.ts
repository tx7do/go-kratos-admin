import type { NotificationMessageService } from '#/generated/api/admin/service/v1/i_notification_message.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type {
  CreateNotificationMessageRequest,
  DeleteNotificationMessageRequest,
  GetNotificationMessageRequest,
  ListNotificationMessageResponse,
  NotificationMessage,
  UpdateNotificationMessageRequest,
} from '#/generated/api/internal_message/service/v1/notification_message.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 通知消息管理服务 */
class NotificationMessageServiceImpl implements NotificationMessageService {
  async Create(request: CreateNotificationMessageRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/notifications', request);
  }

  async Delete(request: DeleteNotificationMessageRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/notifications/${request.id}`);
  }

  async Get(
    request: GetNotificationMessageRequest,
  ): Promise<NotificationMessage> {
    return await requestClient.get<NotificationMessage>(
      `/notifications/${request.id}`,
    );
  }

  async List(request: PagingRequest): Promise<ListNotificationMessageResponse> {
    return await requestClient.get<ListNotificationMessageResponse>(
      '/notifications',
      {
        params: request,
      },
    );
  }

  async Update(request: UpdateNotificationMessageRequest): Promise<Empty> {
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
