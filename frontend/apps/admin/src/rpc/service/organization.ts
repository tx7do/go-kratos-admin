import type { OrganizationService } from '#/rpc/api/admin/service/v1/i_organization.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  CreateOrganizationRequest,
  DeleteOrganizationRequest,
  GetOrganizationRequest,
  ListOrganizationResponse,
  Organization,
  UpdateOrganizationRequest,
} from '#/rpc/api/user/service/v1/organization.pb';

import { requestClient } from '#/rpc/request';

/** 组织管理服务 */
class OrganizationServiceImpl implements OrganizationService {
  async CreateOrganization(request: CreateOrganizationRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/organizations', request);
  }

  async DeleteOrganization(request: DeleteOrganizationRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/organizations/${request.id}`);
  }

  async GetOrganization(
    request: GetOrganizationRequest,
  ): Promise<Organization> {
    return await requestClient.get<Organization>(
      `/organizations/${request.id}`,
    );
  }

  async ListOrganization(
    request: PagingRequest,
  ): Promise<ListOrganizationResponse> {
    return await requestClient.get<ListOrganizationResponse>('/organizations', {
      params: request,
    });
  }

  async UpdateOrganization(request: UpdateOrganizationRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/organizations/${id}`, request);
  }
}

export const defOrganizationService = new OrganizationServiceImpl();
