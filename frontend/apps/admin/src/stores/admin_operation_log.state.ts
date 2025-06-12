import { defineStore } from 'pinia';

import { defAdminOperationLogService } from '#/services';
import { makeQueryString } from '#/utils/query';

export const useAdminOperationLogStore = defineStore(
  'admin_operation_log',
  () => {
    /**
     * 查询操作日志列表
     */
    async function listAdminOperationLog(
      noPaging: boolean = false,
      page?: null | number,
      pageSize?: null | number,
      formValues?: null | object,
      fieldMask?: null | string,
      orderBy?: null | string[],
    ) {
      return await defAdminOperationLogService.List({
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
      return await defAdminOperationLogService.Get({ id });
    }

    function $reset() {}

    return {
      $reset,
      listAdminOperationLog,
      getAdminOperationLog,
    };
  },
);
