import { defineStore } from 'pinia';

import { defPrivateMessageService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const usePrivateMessageStore = defineStore('private_message', () => {
  /**
   * 查询私信消息列表
   */
  async function listPrivateMessage(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defPrivateMessageService.List({
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
   * 获取私信消息
   */
  async function getPrivateMessage(id: number) {
    return await defPrivateMessageService.Get({ id });
  }

  /**
   * 创建私信消息
   */
  async function createPrivateMessage(values: object) {
    return await defPrivateMessageService.Create({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新私信消息
   */
  async function updatePrivateMessage(id: number, values: object) {
    return await defPrivateMessageService.Update({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除私信消息
   */
  async function deletePrivateMessage(id: number) {
    return await defPrivateMessageService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listPrivateMessage,
    getPrivateMessage,
    createPrivateMessage,
    updatePrivateMessage,
    deletePrivateMessage,
  };
});
