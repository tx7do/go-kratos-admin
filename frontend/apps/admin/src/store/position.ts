import { defineStore } from 'pinia';

import { defPositionService, makeQueryString, makeUpdateMask } from '#/rpc';

export const usePositionStore = defineStore('position', () => {
  /**
   * 查询职位列表
   */
  async function listPosition(
    page: number,
    pageSize: number,
    formValues: object,
    fieldMask: null | string = null,
    orderBy: string[] = [],
    noPaging: boolean = false,
  ) {
    return await defPositionService.ListPosition({
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
   * 获取职位
   */
  async function getPosition(id: number) {
    return await defPositionService.GetPosition({ id });
  }

  /**
   * 创建职位
   */
  async function createPosition(values: object) {
    return await defPositionService.CreatePosition({
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
    return await defPositionService.UpdatePosition({
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
    return await defPositionService.DeletePosition({ id });
  }

  return {
    listPosition,
    getPosition,
    createPosition,
    updatePosition,
    deletePosition,
  };
});
