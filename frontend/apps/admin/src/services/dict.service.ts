import type {
  BatchDeleteDictRequest,
  CreateDictEntryRequest,
  CreateDictTypeRequest,
  DictService,
  DictType,
  GetDictTypeRequest,
  ListDictEntryResponse,
  ListDictTypeResponse,
  UpdateDictEntryRequest,
  UpdateDictTypeRequest,
} from '#/generated/api/dict/service/v1/dict.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 字典管理服务 */
class DictServiceImpl implements DictService {
  async CreateDictEntry(request: CreateDictEntryRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/dict-entries', request);
  }

  async CreateDictType(request: CreateDictTypeRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/dict-types', request);
  }

  async DeleteDictEntry(request: BatchDeleteDictRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/dict-entries`, {
      params: request,
    });
  }

  async DeleteDictType(request: BatchDeleteDictRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/dict-types`, {
      params: request,
    });
  }

  async GetDictType(request: GetDictTypeRequest): Promise<DictType> {
    switch (request.queryBy?.$case) {
      case 'code': {
        return await requestClient.get<DictType>(
          `/dict-types/code/${request.queryBy.code}`,
        );
      }
      case 'id': {
        return await requestClient.get<DictType>(
          `/dict-types/${request.queryBy.id}`,
        );
      }
    }
    throw new Error('GetDictType must set queryBy');
  }

  async ListDictEntry(request: PagingRequest): Promise<ListDictEntryResponse> {
    return await requestClient.get<ListDictEntryResponse>('/dict-entries', {
      params: request,
    });
  }

  async ListDictType(request: PagingRequest): Promise<ListDictTypeResponse> {
    return await requestClient.get<ListDictTypeResponse>('/dict-types', {
      params: request,
    });
  }

  async UpdateDictEntry(request: UpdateDictEntryRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/dict-entries/${id}`, request);
  }

  async UpdateDictType(request: UpdateDictTypeRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/dict-types/${id}`, request);
  }
}

export const defDictService = new DictServiceImpl();
