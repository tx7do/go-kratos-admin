import type { AdminLoginLogService } from '#/rpc/api/admin/service/v1/i_admin_login_log.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  AdminLoginLog,
  GetAdminLoginLogRequest,
  ListAdminLoginLogResponse,
} from '#/rpc/api/system/service/v1/admin_login_log.pb';

import { requestClient } from '#/rpc/request';

/** 后台登录日志管理服务 */
class AdminLoginLogServiceImpl implements AdminLoginLogService {
  async GetAdminLoginLog(
    request: GetAdminLoginLogRequest,
  ): Promise<AdminLoginLog> {
    return await requestClient.get<AdminLoginLog>(
      `/admin_login_logs/${request.id}`,
    );
  }

  async ListAdminLoginLog(
    request: PagingRequest,
  ): Promise<ListAdminLoginLogResponse> {
    return await requestClient.get<ListAdminLoginLogResponse>(
      '/admin_login_logs',
      {
        params: request,
      },
    );
  }
}

export const defAdminLoginLogService = new AdminLoginLogServiceImpl();
