import { defineStore } from 'pinia';

import { defAdminOperationLogService, makeQueryString } from '#/rpc';

export const useAdminOperationLogStore = defineStore(
  'admin_operation_log',
  () => {
    /**
     * 查询操作日志列表
     */
    async function listAdminOperationLog(
      page: number,
      pageSize: number,
      formValues: object,
      fieldMask: null | string = null,
      orderBy: string[] = [],
      noPaging: boolean = false,
    ) {
      return await defAdminOperationLogService.ListAdminOperationLog({
        // @ts-ignore proto generated code is error.
        fieldMask,
        orderBy,
        query: makeQueryString(formValues),
        page,
        pageSize,
        noPaging,
      });
    }

    /**
     * 查询操作日志
     */
    async function getAdminOperationLog(id: number) {
      return await defAdminOperationLogService.GetAdminOperationLog({ id });
    }

    return {
      listAdminOperationLog,
      getAdminOperationLog,
    };
  },
);
