import { defineStore } from 'pinia';

import { createAdminLoginLogServiceClient } from '#/generated/api/admin/service/v1';
import { makeQueryString } from '#/utils/query';
import { requestClientRequestHandler } from '#/utils/request';

export const useAdminOperationLogStore = defineStore(
  'admin_operation_log',
  () => {
    const service = createAdminLoginLogServiceClient(
      requestClientRequestHandler,
    );

    /**
     * 查询操作日志列表
     */
    async function listAdminOperationLog(
      noPaging: boolean = false,
      page?: number,
      pageSize?: number,
      formValues?: null | object,
      fieldMask?: null | string,
      orderBy?: null | string[],
    ) {
      return await service.List({
        // @ts-ignore proto generated code is error.
        fieldMask,
        orderBy: orderBy ?? [],
        query: makeQueryString(formValues ?? null),
        page,
        pageSize,
        noPaging,
      });
    }

    /**
     * 查询操作日志
     */
    async function getAdminOperationLog(id: number) {
      return await service.Get({ id });
    }

    function $reset() {}

    return {
      $reset,
      listAdminOperationLog,
      getAdminOperationLog,
    };
  },
);
