import type { InternalMessageCategoryService } from '#/generated/api/admin/service/v1/i_internal_message_category.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type {
  CreateInternalMessageCategoryRequest,
  DeleteInternalMessageCategoryRequest,
  GetInternalMessageCategoryRequest,
  InternalMessageCategory,
  ListInternalMessageCategoryResponse,
  UpdateInternalMessageCategoryRequest,
} from '#/generated/api/internal_message/service/v1/internal_message_category.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 通知消息分类管理服务 */
class InternalMessageCategoryServiceImpl
  implements InternalMessageCategoryService
{
  async Create(request: CreateInternalMessageCategoryRequest): Promise<Empty> {
    return await requestClient.post<Empty>(
      '/internal-message/categories',
      request,
    );
  }

  async Delete(request: DeleteInternalMessageCategoryRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(
      `/internal-message/categories/${request.id}`,
    );
  }

  async Get(
    request: GetInternalMessageCategoryRequest,
  ): Promise<InternalMessageCategory> {
    return await requestClient.get<InternalMessageCategory>(
      `/internal-message/categories/${request.id}`,
    );
  }

  async List(
    request: PagingRequest,
  ): Promise<ListInternalMessageCategoryResponse> {
    return await requestClient.get<ListInternalMessageCategoryResponse>(
      '/internal-message/categories',
      {
        params: request,
      },
    );
  }

  async Update(request: UpdateInternalMessageCategoryRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(
      `/internal-message/categories/${id}`,
      request,
    );
  }
}

export const defInternalMessageCategoryService =
  new InternalMessageCategoryServiceImpl();
