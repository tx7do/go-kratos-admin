import { defineStore } from 'pinia';

import { defTenantService, makeQueryString, makeUpdateMask } from '#/rpc';

export const useTenantStore = defineStore('tenant', () => {
  /**
   * 查询租户列表
   */
  async function listTenant(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defTenantService.ListTenant({
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
   * 获取租户
   */
  async function getTenant(id: number) {
    return await defTenantService.GetTenant({ id });
  }

  /**
   * 创建租户
   */
  async function createTenant(values: object) {
    return await defTenantService.CreateTenant({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新租户
   */
  async function updateTenant(id: number, values: object) {
    return await defTenantService.UpdateTenant({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除租户
   */
  async function deleteTenant(id: number) {
    return await defTenantService.DeleteTenant({ id });
  }

  function $reset() {}

  return {
    $reset,
    listTenant,
    getTenant,
    createTenant,
    updateTenant,
    deleteTenant,
  };
});
