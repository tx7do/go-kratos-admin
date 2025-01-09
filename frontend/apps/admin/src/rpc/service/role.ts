import type { RoleService } from '#/rpc/api/admin/service/v1/i_role.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  CreateRoleRequest,
  DeleteRoleRequest,
  GetRoleRequest,
  ListRoleResponse,
  Role,
  UpdateRoleRequest,
} from '#/rpc/api/user/service/v1/role.pb';

import { requestClient } from '#/rpc/request';

/** 角色管理服务 */
class RoleServiceImpl implements RoleService {
  async CreateRole(request: CreateRoleRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/roles', request);
  }

  async DeleteRole(request: DeleteRoleRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/roles/${request.id}`);
  }

  async GetRole(request: GetRoleRequest): Promise<Role> {
    return await requestClient.get<Role>(`/roles/${request.id}`);
  }

  async ListRole(request: PagingRequest): Promise<ListRoleResponse> {
    return await requestClient.get<ListRoleResponse>('/roles', {
      params: request,
    });
  }

  async UpdateRole(request: UpdateRoleRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateRole', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/roles/${id}`, request);
  }
}

export const defRoleService = new RoleServiceImpl();
