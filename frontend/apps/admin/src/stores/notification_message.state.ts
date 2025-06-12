import { defineStore } from 'pinia';

import { defNotificationMessageService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useNotificationMessageStore = defineStore(
  'notification_message',
  () => {
    /**
     * 查询通知消息列表
     */
    async function listNotificationMessage(
      noPaging: boolean = false,
      page?: null | number,
      pageSize?: null | number,
      formValues?: null | object,
      fieldMask?: null | string,
      orderBy?: null | string[],
    ) {
      return await defNotificationMessageService.List({
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
    async function getNotificationMessage(id: number) {
      return await defNotificationMessageService.Get({ id });
    }

    /**
     * 创建通知消息
     */
    async function createNotificationMessage(values: object) {
      return await defNotificationMessageService.Create({
        data: {
          ...values,
        },
      });
    }

    /**
     * 更新通知消息
     */
    async function updateNotificationMessage(id: number, values: object) {
      return await defNotificationMessageService.Update({
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
    async function deleteNotificationMessage(id: number) {
      return await defNotificationMessageService.Delete({
        id,
      });
    }

    function $reset() {}

    return {
      $reset,
      listNotificationMessage,
      getNotificationMessage,
      createNotificationMessage,
      updateNotificationMessage,
      deleteNotificationMessage,
    };
  },
);
