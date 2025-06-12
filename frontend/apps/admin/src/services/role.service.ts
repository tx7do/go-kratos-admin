import type { RoleService } from '#/generated/api/admin/service/v1/i_role.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';
import type {
  CreateRoleRequest,
  DeleteRoleRequest,
  GetRoleRequest,
  ListRoleResponse,
  Role,
  UpdateRoleRequest,
} from '#/generated/api/user/service/v1/role.pb';

import { requestClient } from '#/utils/request';

/** 角色管理服务 */
class RoleServiceImpl implements RoleService {
  async Create(request: CreateRoleRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/roles', request);
  }

  async Delete(request: DeleteRoleRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/roles/${request.id}`);
  }

  async Get(request: GetRoleRequest): Promise<Role> {
    return await requestClient.get<Role>(`/roles/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListRoleResponse> {
    return await requestClient.get<ListRoleResponse>('/roles', {
      params: request,
    });
  }

  async Update(request: UpdateRoleRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateRole', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/roles/${id}`, request);
  }
}

export const defRoleService = new RoleServiceImpl();
