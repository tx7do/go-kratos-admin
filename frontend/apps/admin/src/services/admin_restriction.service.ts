import type {
  AdminLoginRestriction,
  AdminLoginRestrictionService,
  CreateAdminLoginRestrictionRequest,
  DeleteAdminLoginRestrictionRequest,
  GetAdminLoginRestrictionRequest,
  ListAdminLoginRestrictionResponse,
  UpdateAdminLoginRestrictionRequest,
} from '#/generated/api/admin/service/v1/i_admin_login_restriction.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 后台登录限制管理服务 */
class AdminLoginRestrictionServiceImpl implements AdminLoginRestrictionService {
  async Create(request: CreateAdminLoginRestrictionRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/login-restrictions', request);
  }

  async Delete(request: DeleteAdminLoginRestrictionRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(
      `/login-restrictions/${request.id}`,
    );
  }

  async Get(
    request: GetAdminLoginRestrictionRequest,
  ): Promise<AdminLoginRestriction> {
    return await requestClient.get<AdminLoginRestriction>(
      `/login-restrictions/${request.id}`,
    );
  }

  async List(
    request: PagingRequest,
  ): Promise<ListAdminLoginRestrictionResponse> {
    return await requestClient.get<ListAdminLoginRestrictionResponse>(
      '/login-restrictions',
      {
        params: request,
      },
    );
  }

  async Update(request: UpdateAdminLoginRestrictionRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/login-restrictions/${id}`, request);
  }
}

export const defAdminLoginRestrictionService =
  new AdminLoginRestrictionServiceImpl();
