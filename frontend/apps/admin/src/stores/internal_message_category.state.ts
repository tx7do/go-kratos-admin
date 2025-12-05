import { defineStore } from 'pinia';

import { createInternalMessageCategoryServiceClient } from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import { requestClientRequestHandler } from '#/utils/request';

export const useInternalMessageCategoryStore = defineStore(
  'internal_message_category',
  () => {
    const service = createInternalMessageCategoryServiceClient(
      requestClientRequestHandler,
    );

    /**
     * 查询通知消息列表
     */
    async function listInternalMessageCategory(
      noPaging: boolean = false,
      page?: null | number,
      pageSize?: null | number,
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
     * 获取通知消息
     */
    async function getInternalMessageCategory(id: number) {
      return await service.Get({ id });
    }

    /**
     * 创建通知消息
     */
    async function createInternalMessageCategory(values: object) {
      return await service.Create({
        // @ts-ignore proto generated code is error.
        data: {
          ...values,
        },
      });
    }

    /**
     * 更新通知消息
     */
    async function updateInternalMessageCategory(id: number, values: object) {
      return await service.Update({
        // @ts-ignore proto generated code is error.
        data: {
          id,
          ...values,
        },
        // @ts-ignore proto generated code is error.
        updateMask: makeUpdateMask(Object.keys(values ?? [])),
      });
    }

    /**
     * 删除通知消息
     */
    async function deleteInternalMessageCategory(id: number) {
      return await service.Delete({
        id,
      });
    }

    function $reset() {}

    return {
      $reset,
      listInternalMessageCategory,
      getInternalMessageCategory,
      createInternalMessageCategory,
      updateInternalMessageCategory,
      deleteInternalMessageCategory,
    };
  },
);
