import { defineStore } from 'pinia';

import { defDictService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useDictStore = defineStore('dict', () => {
  /**
   * 查询字典类型列表
   */
  async function listDictType(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDictService.ListDictType({
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
   * 查询字典条目列表
   */
  async function listDictEntry(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDictService.ListDictEntry({
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
   * 获取字典类型
   */
  async function getDictType(id: number) {
    return await defDictService.GetDictType({
      queryBy: { $case: 'id', id },
    });
  }

  /**
   * 获取字典类型
   */
  async function getDictTypeByCode(code: string) {
    return await defDictService.GetDictType({
      queryBy: { $case: 'code', code },
    });
  }

  /**
   * 创建字典类型
   */
  async function createDictType(values: object) {
    return await defDictService.CreateDictType({
      data: {
        ...values,
      },
    });
  }

  /**
   * 创建字典条目
   */
  async function createDictEntry(values: object) {
    return await defDictService.CreateDictEntry({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新字典类型
   */
  async function updateDictType(id: number, values: object) {
    return await defDictService.UpdateDictType({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 更新字典条目
   */
  async function updateDictEntry(id: number, values: object) {
    return await defDictService.UpdateDictEntry({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除字典类型
   */
  async function deleteDictType(ids: number[]) {
    return await defDictService.DeleteDictType({ ids });
  }

  /**
   * 删除字典条目
   */
  async function deleteDictEntry(ids: number[]) {
    return await defDictService.DeleteDictEntry({ ids });
  }

  function $reset() {}

  return {
    $reset,
    listDictType,
    listDictEntry,
    getDictType,
    getDictTypeByCode,
    createDictType,
    createDictEntry,
    updateDictType,
    updateDictEntry,
    deleteDictType,
    deleteDictEntry,
  };
});
