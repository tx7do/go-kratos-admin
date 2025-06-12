import type { TenantService } from '#/generated/api/admin/service/v1/i_tenant.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';
import type {
  CreateTenantRequest,
  DeleteTenantRequest,
  GetTenantRequest,
  ListTenantResponse,
  Tenant,
  UpdateTenantRequest,
} from '#/generated/api/user/service/v1/tenant.pb';

import { requestClient } from '#/utils/request';

/** 租户管理服务 */
class TenantServiceImpl implements TenantService {
  async Create(request: CreateTenantRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/tenants', request);
  }

  async Delete(request: DeleteTenantRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/tenants/${request.id}`);
  }

  async Get(request: GetTenantRequest): Promise<Tenant> {
    return await requestClient.get<Tenant>(`/tenants/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListTenantResponse> {
    return await requestClient.get<ListTenantResponse>('/tenants', {
      params: request,
    });
  }

  async Update(request: UpdateTenantRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateTenant', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/tenants/${id}`, request);
  }
}

export const defTenantService = new TenantServiceImpl();
