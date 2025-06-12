import { defineStore } from 'pinia';

import { defPositionService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const usePositionStore = defineStore('position', () => {
  /**
   * 查询职位列表
   */
  async function listPosition(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defPositionService.List({
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
   * 获取职位
   */
  async function getPosition(id: number) {
    return await defPositionService.Get({ id });
  }

  /**
   * 创建职位
   */
  async function createPosition(values: object) {
    return await defPositionService.Create({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新职位
   */
  async function updatePosition(id: number, values: object) {
    return await defPositionService.Update({
      data: {
        id,
        ...values,
        children: [],
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除职位
   */
  async function deletePosition(id: number) {
    return await defPositionService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listPosition,
    getPosition,
    createPosition,
    updatePosition,
    deletePosition,
  };
});
