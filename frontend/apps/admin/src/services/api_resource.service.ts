import type {
  ApiResource,
  ApiResourceService,
  CreateApiResourceRequest,
  DeleteApiResourceRequest,
  GetApiResourceRequest,
  ListApiResourceResponse,
  UpdateApiResourceRequest,
} from '#/generated/api/admin/service/v1/i_api_resource.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** API资源管理服务 */
class ApiResourceServiceImpl implements ApiResourceService {
  async Create(request: CreateApiResourceRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/api-resources', request);
  }

  async Delete(request: DeleteApiResourceRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/api-resources/${request.id}`);
  }

  async Get(request: GetApiResourceRequest): Promise<ApiResource> {
    return await requestClient.get<ApiResource>(`/api-resources/${request.id}`);
  }

  GetWalkRouteData(request: Empty): Promise<ListApiResourceResponse> {
    return requestClient.get<ListApiResourceResponse>(
      '/api-resources/walk-route',
      {
        params: request,
      },
    );
  }

  async List(request: PagingRequest): Promise<ListApiResourceResponse> {
    return await requestClient.get<ListApiResourceResponse>('/api-resources', {
      params: request,
    });
  }

  SyncApiResources(request: Empty): Promise<Empty> {
    return requestClient.post<Empty>('/api-resources/sync', request);
  }

  async Update(request: UpdateApiResourceRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/api-resources/${id}`, request);
  }
}

export const defApiResourceService = new ApiResourceServiceImpl();
