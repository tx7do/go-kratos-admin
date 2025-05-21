import type {
  CreateDictRequest,
  DeleteDictRequest,
  Dict,
  DictService,
  GetDictRequest,
  ListDictResponse,
  UpdateDictRequest,
} from '#/rpc/api/admin/service/v1/i_dict.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';

import { requestClient } from '#/rpc/request';

/** 职位管理服务 */
class DictServiceImpl implements DictService {
  async Create(request: CreateDictRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/dict', request);
  }

  async Delete(request: DeleteDictRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/dict/${request.id}`);
  }

  async Get(request: GetDictRequest): Promise<Dict> {
    return await requestClient.get<Dict>(`/dict/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListDictResponse> {
    return await requestClient.get<ListDictResponse>('/dict', {
      params: request,
    });
  }

  async Update(request: UpdateDictRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/dict/${id}`, request);
  }
}

export const defDictService = new DictServiceImpl();
