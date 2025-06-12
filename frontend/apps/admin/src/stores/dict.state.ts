import { defineStore } from 'pinia';

import { defDictService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useDictStore = defineStore('dict', () => {
  /**
   * 查询字典列表
   */
  async function listDict(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDictService.List({
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
   * 获取字典
   */
  async function getDict(id: number) {
    return await defDictService.Get({ id });
  }

  /**
   * 创建字典
   */
  async function createDict(values: object) {
    return await defDictService.Create({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新字典
   */
  async function updateDict(id: number, values: object) {
    return await defDictService.Update({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除字典
   */
  async function deleteDict(id: number) {
    return await defDictService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listDict,
    getDict,
    createDict,
    updateDict,
    deleteDict,
  };
});
