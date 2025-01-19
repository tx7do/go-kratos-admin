import { defineStore } from 'pinia';

import { defDictService, makeQueryString, makeUpdateMask } from '#/rpc';

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
    return await defDictService.ListDict({
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
    return await defDictService.GetDict({ id });
  }

  /**
   * 创建字典
   */
  async function createDict(values: object) {
    return await defDictService.CreateDict({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新字典
   */
  async function updateDict(id: number, values: object) {
    return await defDictService.UpdateDict({
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
    return await defDictService.DeleteDict({ id });
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
