import { defineStore } from 'pinia';

import { defRoleService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useRoleStore = defineStore('role', () => {
  /**
   * 查询角色列表
   */
  async function listRole(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defRoleService.List({
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
   * 获取角色
   */
  async function getRole(id: number) {
    return await defRoleService.Get({ id });
  }

  /**
   * 创建角色
   */
  async function createRole(values: object) {
    return await defRoleService.Create({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新角色
   */
  async function updateRole(id: number, values: object) {
    return await defRoleService.Update({
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
   * 删除角色
   */
  async function deleteRole(id: number) {
    return await defRoleService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listRole,
    getRole,
    createRole,
    updateRole,
    deleteRole,
  };
});
