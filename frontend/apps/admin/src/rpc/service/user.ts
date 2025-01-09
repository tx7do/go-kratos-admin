import type { UserService } from '#/rpc/api/admin/service/v1/i_user.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  CreateUserRequest,
  DeleteUserRequest,
  GetUserRequest,
  ListUserResponse,
  UpdateUserRequest,
  User,
} from '#/rpc/api/user/service/v1/user.pb';

import { UserAuthority } from '#/rpc/api/user/service/v1/user.pb';
import { requestClient } from '#/rpc/request';

/** 用户管理服务 */
class UserServiceImpl implements UserService {
  async CreateUser(request: CreateUserRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/users', request);
  }

  async DeleteUser(request: DeleteUserRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/users/${request.id}`);
  }

  async GetUser(request: GetUserRequest): Promise<User> {
    return await requestClient.get<User>(`/users/${request.id}`);
  }

  async ListUser(request: PagingRequest): Promise<ListUserResponse> {
    return await requestClient.get<ListUserResponse>('/users', {
      params: request,
    });
  }

  async UpdateUser(request: UpdateUserRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null) request.data.id = undefined;
    return await requestClient.put<Empty>(`/users/${id}`, request);
  }
}

export const authorityList = [
  { value: UserAuthority.GUEST_USER, label: '游客' },
  { value: UserAuthority.CUSTOMER_USER, label: '普通用户' },
  { value: UserAuthority.SYS_MANAGER, label: '普通管理' },
  { value: UserAuthority.SYS_ADMIN, label: '超级管理' },
];

/**
 * 权限转名称
 * @param authority
 */
export function authorityToName(authority: any) {
  switch (authority) {
    case UserAuthority.CUSTOMER_USER: {
      return '普通用户';
    }
    case UserAuthority.GUEST_USER: {
      return '游客';
    }
    case UserAuthority.SYS_ADMIN: {
      return '超级管理';
    }
    case UserAuthority.SYS_MANAGER: {
      return '普通管理';
    }
    default: {
      return '';
    }
  }
}

/**
 * 权限转颜色值
 * @param authority
 */
export function authorityToColor(authority: any) {
  switch (authority) {
    case UserAuthority.CUSTOMER_USER: {
      return 'green';
    }
    case UserAuthority.GUEST_USER: {
      return 'green';
    }
    case UserAuthority.SYS_ADMIN: {
      return 'orange';
    }
    case UserAuthority.SYS_MANAGER: {
      return 'red';
    }
    default: {
      return 'black';
    }
  }
}

export const defUserService = new UserServiceImpl();
