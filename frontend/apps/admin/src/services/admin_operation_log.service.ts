import type {
  AdminOperationLog,
  AdminOperationLogService,
  GetAdminOperationLogRequest,
  ListAdminOperationLogResponse,
} from '#/generated/api/admin/service/v1/i_admin_operation_log.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 后台操作日志管理服务 */
class AdminOperationLogServiceImpl implements AdminOperationLogService {
  async Get(request: GetAdminOperationLogRequest): Promise<AdminOperationLog> {
    return await requestClient.get<AdminOperationLog>(
      `/admin_operation_logs/${request.id}`,
    );
  }

  async List(request: PagingRequest): Promise<ListAdminOperationLogResponse> {
    return await requestClient.get<ListAdminOperationLogResponse>(
      '/admin_operation_logs',
      {
        params: request,
      },
    );
  }
}

export const defAdminOperationLogService = new AdminOperationLogServiceImpl();
