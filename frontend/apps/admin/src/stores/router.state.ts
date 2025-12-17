import { defineStore } from 'pinia';

import { createRouterServiceClient } from '#/generated/api/admin/service/v1';
import { requestClientRequestHandler } from '#/utils/request';

export const useRouterStore = defineStore('router', () => {
  const service = createRouterServiceClient(requestClientRequestHandler);

  /**
   * 查询路由列表
   */
  async function listRouter() {
    return await service.ListRoute({});
  }

  function $reset() {}

  return {
    $reset,
    listRouter,
  };
});
