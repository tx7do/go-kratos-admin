import type { DepartmentService } from '#/rpc/api/admin/service/v1/i_department.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  CreateDepartmentRequest,
  DeleteDepartmentRequest,
  Department,
  GetDepartmentRequest,
  ListDepartmentResponse,
  UpdateDepartmentRequest,
} from '#/rpc/api/user/service/v1/department.pb';

import { requestClient } from '#/rpc/request';

/** 部门管理服务 */
class DepartmentServiceImpl implements DepartmentService {
  async CreateDepartment(request: CreateDepartmentRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/departments', request);
  }

  async DeleteDepartment(request: DeleteDepartmentRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/departments/${request.id}`);
  }

  async GetDepartment(request: GetDepartmentRequest): Promise<Department> {
    return await requestClient.get<Department>(`/departments/${request.id}`);
  }

  async ListDepartment(
    request: PagingRequest,
  ): Promise<ListDepartmentResponse> {
    return await requestClient.get<ListDepartmentResponse>('/departments', {
      params: request,
    });
  }

  async UpdateDepartment(request: UpdateDepartmentRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null) request.data.id = undefined;
    return await requestClient.put<Empty>(`/departments/${id}`, request);
  }
}

export const defDepartmentService = new DepartmentServiceImpl();
