import type { FileService } from '#/generated/api/admin/service/v1/i_file.pb';
import type {
  CreateFileRequest,
  DeleteFileRequest,
  File,
  GetFileRequest,
  ListFileResponse,
  UpdateFileRequest,
} from '#/generated/api/file/service/v1/file.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 文件管理服务 */
class FileServiceImpl implements FileService {
  async Create(request: CreateFileRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/files', request);
  }

  async Delete(request: DeleteFileRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/files/${request.id}`);
  }

  async Get(request: GetFileRequest): Promise<File> {
    return await requestClient.get<File>(`/files/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListFileResponse> {
    return await requestClient.get<ListFileResponse>('/files', {
      params: request,
    });
  }

  async Update(request: UpdateFileRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateFile', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/files/${id}`, request);
  }
}

export const defFileService = new FileServiceImpl();
