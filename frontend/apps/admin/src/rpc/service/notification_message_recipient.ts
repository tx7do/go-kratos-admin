import type { NotificationMessageRecipientService } from '#/rpc/api/admin/service/v1/i_notification_message_recipient.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type {
  CreateNotificationMessageRecipientRequest,
  DeleteNotificationMessageRecipientRequest,
  GetNotificationMessageRecipientRequest,
  ListNotificationMessageRecipientResponse,
  NotificationMessageRecipient,
  UpdateNotificationMessageRecipientRequest,
} from '#/rpc/api/internal_message/service/v1/notification_message_recipient.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';

import { requestClient } from '#/rpc/request';

/** 通知消息接收者管理服务 */
class NotificationMessageRecipientServiceImpl
  implements NotificationMessageRecipientService
{
  async Create(
    request: CreateNotificationMessageRecipientRequest,
  ): Promise<Empty> {
    return await requestClient.post<Empty>(
      '/notifications:recipients',
      request,
    );
  }

  async Delete(
    request: DeleteNotificationMessageRecipientRequest,
  ): Promise<Empty> {
    return await requestClient.delete<Empty>(
      `/notifications:recipients/${request.id}`,
    );
  }

  async Get(
    request: GetNotificationMessageRecipientRequest,
  ): Promise<NotificationMessageRecipient> {
    return await requestClient.get<NotificationMessageRecipient>(
      `/notifications:recipients/${request.id}`,
    );
  }

  async List(
    request: PagingRequest,
  ): Promise<ListNotificationMessageRecipientResponse> {
    return await requestClient.get<ListNotificationMessageRecipientResponse>(
      '/notifications:recipients',
      {
        params: request,
      },
    );
  }

  async Update(
    request: UpdateNotificationMessageRecipientRequest,
  ): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateNotificationMessageRecipient', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(
      `/notifications:recipients/${id}`,
      request,
    );
  }
}

export const defNotificationMessageRecipientService =
  new NotificationMessageRecipientServiceImpl();
