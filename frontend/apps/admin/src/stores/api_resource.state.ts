import { defineStore } from 'pinia';

import { defApiResourceService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useApiResourceStore = defineStore('api-resource', () => {
  /**
   * 查询API列表
   */
  async function listApiResource(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defApiResourceService.List({
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
   * 获取API
   */
  async function getApiResource(id: number) {
    return await defApiResourceService.Get({ id });
  }

  /**
   * 创建API
   */
  async function createApiResource(values: object) {
    return await defApiResourceService.Create({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新API
   */
  async function updateApiResource(id: number, values: object) {
    return await defApiResourceService.Update({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除API
   */
  async function deleteApiResource(id: number) {
    return await defApiResourceService.Delete({ id });
  }

  async function getWalkRouteData() {
    return await defApiResourceService.GetWalkRouteData({});
  }

  async function syncApiResources() {
    return await defApiResourceService.SyncApiResources({});
  }

  function $reset() {}

  return {
    $reset,
    listApiResource,
    getApiResource,
    createApiResource,
    updateApiResource,
    deleteApiResource,
    getWalkRouteData,
    syncApiResources,
  };
});

export const methodList = [
  { value: 'GET', label: 'GET' },
  { value: 'POST', label: 'POST' },
  { value: 'PUT', label: 'PUT' },
  { value: 'DELETE', label: 'DELETE' },
];
