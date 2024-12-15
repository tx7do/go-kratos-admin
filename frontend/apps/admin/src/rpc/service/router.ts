import type {
  ListPermissionCodeResponse,
  ListRouteResponse,
  RouterService,
} from '#/rpc/api/admin/service/v1/i_router.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';

import { requestClient } from '#/rpc/request';

/** 网站后台动态路由服务 */
export class RouterServiceImpl implements RouterService {
  async ListPermissionCode(
    _request: Empty,
  ): Promise<ListPermissionCodeResponse> {
    return await requestClient.get<ListPermissionCodeResponse>(`/perm-codes`);
  }

  async ListRoute(_request: Empty): Promise<ListRouteResponse> {
    return await requestClient.get<ListRouteResponse>(`/routes`);
  }
}

export const defRouterService = new RouterServiceImpl();
