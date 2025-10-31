import { defineStore } from 'pinia';

import { defDictService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useDictStore = defineStore('dict', () => {
  /**
   * 查询主字典列表
   */
  async function listDictMain(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDictService.ListDictMain({
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
   * 查询子字典列表
   */
  async function listDictItem(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDictService.ListDictItem({
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
   * 获取主字典
   */
  async function getDictMain(id: number) {
    return await defDictService.GetDictMain({
      queryBy: { $case: 'id', id },
    });
  }

  /**
   * 获取主字典
   */
  async function getDictMainByCode(code: string) {
    return await defDictService.GetDictMain({
      queryBy: { $case: 'code', code },
    });
  }

  /**
   * 创建主字典
   */
  async function createDictMain(values: object) {
    return await defDictService.CreateDictMain({
      data: {
        ...values,
      },
    });
  }

  /**
   * 创建子字典
   */
  async function createDictItem(values: object) {
    return await defDictService.CreateDictItem({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新主字典
   */
  async function updateDictMain(id: number, values: object) {
    return await defDictService.UpdateDictMain({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 更新子字典
   */
  async function updateDictItem(id: number, values: object) {
    return await defDictService.UpdateDictItem({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除主字典
   */
  async function deleteDictMain(ids: number[]) {
    return await defDictService.DeleteDictMain({ ids });
  }

  /**
   * 删除子字典
   */
  async function deleteDictItem(ids: number[]) {
    return await defDictService.DeleteDictItem({ ids });
  }

  function $reset() {}

  return {
    $reset,
    listDictMain,
    listDictItem,
    getDictMain,
    getDictMainByCode,
    createDictMain,
    createDictItem,
    updateDictMain,
    updateDictItem,
    deleteDictMain,
    deleteDictItem,
  };
});
