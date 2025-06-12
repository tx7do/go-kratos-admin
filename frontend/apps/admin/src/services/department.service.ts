import type { DepartmentService } from '#/generated/api/admin/service/v1/i_department.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';
import type {
  CreateDepartmentRequest,
  DeleteDepartmentRequest,
  Department,
  GetDepartmentRequest,
  ListDepartmentResponse,
  UpdateDepartmentRequest,
} from '#/generated/api/user/service/v1/department.pb';

import { requestClient } from '#/utils/request';

/** 部门管理服务 */
class DepartmentServiceImpl implements DepartmentService {
  async Create(request: CreateDepartmentRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/departments', request);
  }

  async Delete(request: DeleteDepartmentRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/departments/${request.id}`);
  }

  async Get(request: GetDepartmentRequest): Promise<Department> {
    return await requestClient.get<Department>(`/departments/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListDepartmentResponse> {
    return await requestClient.get<ListDepartmentResponse>('/departments', {
      params: request,
    });
  }

  async Update(request: UpdateDepartmentRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/departments/${id}`, request);
  }
}

export const defDepartmentService = new DepartmentServiceImpl();
