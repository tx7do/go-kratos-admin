import type { AdminOperationLogService } from '#/rpc/api/admin/service/v1/i_admin_operation_log.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  AdminOperationLog,
  GetAdminOperationLogRequest,
  ListAdminOperationLogResponse,
} from '#/rpc/api/system/service/v1/admin_operation_log.pb';

import { requestClient } from '#/rpc/request';

/** 后台操作日志管理服务 */
class AdminOperationLogServiceImpl implements AdminOperationLogService {
  async GetAdminOperationLog(
    request: GetAdminOperationLogRequest,
  ): Promise<AdminOperationLog> {
    return await requestClient.get<AdminOperationLog>(
      `/admin_operation_logs/${request.id}`,
    );
  }

  async ListAdminOperationLog(
    request: PagingRequest,
  ): Promise<ListAdminOperationLogResponse> {
    return await requestClient.get<ListAdminOperationLogResponse>(
      '/admin_operation_logs',
      {
        params: request,
      },
    );
  }
}

export const defAdminOperationLogService = new AdminOperationLogServiceImpl();
