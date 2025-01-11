import { defineStore } from 'pinia';

import { defDictService, makeQueryString, makeUpdateMask } from '#/rpc';

export const useDictStore = defineStore('dict', () => {
  /**
   * 查询字典列表
   */
  async function listDict(
    page: number,
    pageSize: number,
    formValues: object,
    fieldMask: null | string = null,
    orderBy: string[] = [],
    noPaging: boolean = false,
  ) {
    return await defDictService.ListDict({
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

  return {
    listDict,
    getDict,
    createDict,
    updateDict,
    deleteDict,
  };
});
