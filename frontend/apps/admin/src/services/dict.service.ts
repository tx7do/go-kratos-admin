import type {
  BatchDeleteDictRequest,
  CreateDictItemRequest,
  CreateDictMainRequest,
  DictMain,
  DictService,
  GetDictMainRequest,
  ListDictItemResponse,
  ListDictMainResponse,
  UpdateDictItemRequest,
  UpdateDictMainRequest,
} from '#/generated/api/admin/service/v1/i_dict.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 字典管理服务 */
class DictServiceImpl implements DictService {
  async CreateDictItem(request: CreateDictItemRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/dict-items', request);
  }

  async CreateDictMain(request: CreateDictMainRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/dict-mains', request);
  }

  async DeleteDictItem(request: BatchDeleteDictRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/dict-items`, {
      params: request,
    });
  }

  async DeleteDictMain(request: BatchDeleteDictRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/dict-mains`, {
      params: request,
    });
  }

  async GetDictMain(request: GetDictMainRequest): Promise<DictMain> {
    switch (request.queryBy?.$case) {
      case 'code': {
        return await requestClient.get<DictMain>(
          `/dict-mains/code/${request.queryBy.code}`,
        );
      }
      case 'id': {
        return await requestClient.get<DictMain>(
          `/dict-mains/${request.queryBy.id}`,
        );
      }
    }
    throw new Error('GetDictMain must set queryBy');
  }

  async ListDictItem(request: PagingRequest): Promise<ListDictItemResponse> {
    return await requestClient.get<ListDictItemResponse>('/dict-items', {
      params: request,
    });
  }

  async ListDictMain(request: PagingRequest): Promise<ListDictMainResponse> {
    return await requestClient.get<ListDictMainResponse>('/dict-mains', {
      params: request,
    });
  }

  async UpdateDictItem(request: UpdateDictItemRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/dict-items/${id}`, request);
  }

  async UpdateDictMain(request: UpdateDictMainRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/dict-mains/${id}`, request);
  }
}

export const defDictService = new DictServiceImpl();
