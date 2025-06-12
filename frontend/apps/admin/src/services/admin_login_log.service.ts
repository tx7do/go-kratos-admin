import type {
  AdminLoginLog,
  AdminLoginLogService,
  GetAdminLoginLogRequest,
  ListAdminLoginLogResponse,
} from '#/generated/api/admin/service/v1/i_admin_login_log.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 后台登录日志管理服务 */
class AdminLoginLogServiceImpl implements AdminLoginLogService {
  async Get(request: GetAdminLoginLogRequest): Promise<AdminLoginLog> {
    return await requestClient.get<AdminLoginLog>(
      `/admin_login_logs/${request.id}`,
    );
  }

  async List(request: PagingRequest): Promise<ListAdminLoginLogResponse> {
    return await requestClient.get<ListAdminLoginLogResponse>(
      '/admin_login_logs',
      {
        params: request,
      },
    );
  }
}

export const defAdminLoginLogService = new AdminLoginLogServiceImpl();
