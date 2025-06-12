import type { PositionService } from '#/generated/api/admin/service/v1/i_position.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';
import type {
  CreatePositionRequest,
  DeletePositionRequest,
  GetPositionRequest,
  ListPositionResponse,
  Position,
  UpdatePositionRequest,
} from '#/generated/api/user/service/v1/position.pb';

import { requestClient } from '#/utils/request';

/** 职位管理服务 */
class PositionServiceImpl implements PositionService {
  async Create(request: CreatePositionRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/positions', request);
  }

  async Delete(request: DeletePositionRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/positions/${request.id}`);
  }

  async Get(request: GetPositionRequest): Promise<Position> {
    return await requestClient.get<Position>(`/positions/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListPositionResponse> {
    return await requestClient.get<ListPositionResponse>('/positions', {
      params: request,
    });
  }

  async Update(request: UpdatePositionRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/positions/${id}`, request);
  }
}

export const defPositionService = new PositionServiceImpl();
