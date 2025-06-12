import type { UserService } from '#/generated/api/admin/service/v1/i_user.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';
import type {
  CreateUserRequest,
  DeleteUserRequest,
  GetUserRequest,
  ListUserResponse,
  UpdateUserRequest,
  User,
} from '#/generated/api/user/service/v1/user.pb';

import { requestClient } from '#/utils/request';

/** 用户管理服务 */
class UserServiceImpl implements UserService {
  async Create(request: CreateUserRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/users', request);
  }

  async Delete(request: DeleteUserRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/users/${request.id}`);
  }

  async Get(request: GetUserRequest): Promise<User> {
    return await requestClient.get<User>(`/users/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListUserResponse> {
    return await requestClient.get<ListUserResponse>('/users', {
      params: request,
    });
  }

  async Update(request: UpdateUserRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/users/${id}`, request);
  }
}

export const defUserService = new UserServiceImpl();
