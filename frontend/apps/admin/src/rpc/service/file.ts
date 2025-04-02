import type { FileService } from '#/rpc/api/admin/service/v1/i_file.pb';
import type {
  CreateFileRequest,
  DeleteFileRequest,
  File,
  GetFileRequest,
  ListFileResponse,
  UpdateFileRequest,
} from '#/rpc/api/file/service/v1/file.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';

import { requestClient } from '#/rpc/request';

/** 文件管理服务 */
class FileServiceImpl implements FileService {
  async CreateFile(request: CreateFileRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/files', request);
  }

  async DeleteFile(request: DeleteFileRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/files/${request.id}`);
  }

  async GetFile(request: GetFileRequest): Promise<File> {
    return await requestClient.get<File>(`/files/${request.id}`);
  }

  async ListFile(request: PagingRequest): Promise<ListFileResponse> {
    return await requestClient.get<ListFileResponse>('/files', {
      params: request,
    });
  }

  async UpdateFile(request: UpdateFileRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateFile', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/files/${id}`, request);
  }
}

export const defFileService = new FileServiceImpl();
