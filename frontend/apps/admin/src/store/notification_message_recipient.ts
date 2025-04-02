import { defineStore } from 'pinia';

import {
  defNotificationMessageRecipientService,
  makeQueryString,
  makeUpdateMask,
} from '#/rpc';

export const useNotificationMessageRecipientStore = defineStore(
  'notification_message_recipient',
  () => {
    /**
     * 查询通知消息类型列表
     */
    async function listNotificationMessageRecipient(
      noPaging: boolean = false,
      page?: null | number,
      pageSize?: null | number,
      formValues?: null | object,
      fieldMask?: null | string,
      orderBy?: null | string[],
    ) {
      return await defNotificationMessageRecipientService.ListNotificationMessageRecipient(
        {
          // @ts-ignore proto generated code is error.
          fieldMask,
          orderBy: orderBy ?? [],
          query: makeQueryString(formValues ?? null),
          page,
          pageSize,
          noPaging,
        },
      );
    }

    /**
     * 获取通知消息类型
     */
    async function getNotificationMessageRecipient(id: number) {
      return await defNotificationMessageRecipientService.GetNotificationMessageRecipient(
        { id },
      );
    }

    /**
     * 创建通知消息类型
     */
    async function createNotificationMessageRecipient(values: object) {
      return await defNotificationMessageRecipientService.CreateNotificationMessageRecipient(
        {
          data: {
            ...values,
          },
        },
      );
    }

    /**
     * 更新通知消息类型
     */
    async function updateNotificationMessageRecipient(
      id: number,
      values: object,
    ) {
      return await defNotificationMessageRecipientService.UpdateNotificationMessageRecipient(
        {
          data: {
            id,
            ...values,
          },
          // @ts-ignore proto generated code is error.
          updateMask: makeUpdateMask(Object.keys(values ?? [])),
        },
      );
    }

    /**
     * 删除通知消息类型
     */
    async function deleteNotificationMessageRecipient(id: number) {
      return await defNotificationMessageRecipientService.DeleteNotificationMessageRecipient(
        {
          id,
        },
      );
    }

    function $reset() {}

    return {
      $reset,
      listNotificationMessageRecipient,
      getNotificationMessageRecipient,
      createNotificationMessageRecipient,
      updateNotificationMessageRecipient,
      deleteNotificationMessageRecipient,
    };
  },
);
