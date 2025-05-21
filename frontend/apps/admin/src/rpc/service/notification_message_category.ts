import type { NotificationMessageCategoryService } from '#/rpc/api/admin/service/v1/i_notification_message_category.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type {
  CreateNotificationMessageCategoryRequest,
  DeleteNotificationMessageCategoryRequest,
  GetNotificationMessageCategoryRequest,
  ListNotificationMessageCategoryResponse,
  NotificationMessageCategory,
  UpdateNotificationMessageCategoryRequest,
} from '#/rpc/api/internal_message/service/v1/notification_message_category.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';

import { requestClient } from '#/rpc/request';

/** 通知消息分类管理服务 */
class NotificationMessageCategoryServiceImpl
  implements NotificationMessageCategoryService
{
  async Create(
    request: CreateNotificationMessageCategoryRequest,
  ): Promise<Empty> {
    return await requestClient.post<Empty>(
      '/notifications:categories',
      request,
    );
  }

  async Delete(
    request: DeleteNotificationMessageCategoryRequest,
  ): Promise<Empty> {
    return await requestClient.delete<Empty>(
      `/notifications:categories/${request.id}`,
    );
  }

  async Get(
    request: GetNotificationMessageCategoryRequest,
  ): Promise<NotificationMessageCategory> {
    return await requestClient.get<NotificationMessageCategory>(
      `/notifications:categories/${request.id}`,
    );
  }

  async List(
    request: PagingRequest,
  ): Promise<ListNotificationMessageCategoryResponse> {
    return await requestClient.get<ListNotificationMessageCategoryResponse>(
      '/notifications:categories',
      {
        params: request,
      },
    );
  }

  async Update(
    request: UpdateNotificationMessageCategoryRequest,
  ): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateNotificationMessageCategory', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(
      `/notifications:categories/${id}`,
      request,
    );
  }
}

export const defNotificationMessageCategoryService =
  new NotificationMessageCategoryServiceImpl();
