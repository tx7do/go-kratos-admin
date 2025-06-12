import type { PrivateMessageService } from '#/generated/api/admin/service/v1/i_private_message.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type {
  CreatePrivateMessageRequest,
  DeletePrivateMessageRequest,
  GetPrivateMessageRequest,
  ListPrivateMessageResponse,
  PrivateMessage,
  UpdatePrivateMessageRequest,
} from '#/generated/api/internal_message/service/v1/private_message.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 私信消息管理服务 */
class PrivateMessageServiceImpl implements PrivateMessageService {
  async Create(request: CreatePrivateMessageRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/private_messages', request);
  }

  async Delete(request: DeletePrivateMessageRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/private_messages/${request.id}`);
  }

  async Get(request: GetPrivateMessageRequest): Promise<PrivateMessage> {
    return await requestClient.get<PrivateMessage>(
      `/private_messages/${request.id}`,
    );
  }

  async List(request: PagingRequest): Promise<ListPrivateMessageResponse> {
    return await requestClient.get<ListPrivateMessageResponse>(
      '/private_messages',
      {
        params: request,
      },
    );
  }

  async Update(request: UpdatePrivateMessageRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdatePrivateMessage', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/private_messages/${id}`, request);
  }
}

export const defPrivateMessageService = new PrivateMessageServiceImpl();
