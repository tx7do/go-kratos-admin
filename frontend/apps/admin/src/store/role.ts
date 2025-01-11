import { defineStore } from 'pinia';

import { defRoleService, makeQueryString, makeUpdateMask } from '#/rpc';

export const useRoleStore = defineStore('role', () => {
  /**
   * 查询职位列表
   */
  async function listRole(
    page: number,
    pageSize: number,
    formValues: object,
    fieldMask: null | string = null,
    orderBy: string[] = [],
    noPaging: boolean = false,
  ) {
    return await defRoleService.ListRole({
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
  async function getRole(id: number) {
    return await defRoleService.GetRole({ id });
  }

  /**
   * 创建职位
   */
  async function createRole(values: object) {
    return await defRoleService.CreateRole({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新职位
   */
  async function updateRole(id: number, values: object) {
    return await defRoleService.UpdateRole({
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
  async function deleteRole(id: number) {
    return await defRoleService.DeleteRole({ id });
  }

  return {
    listRole,
    getRole,
    createRole,
    updateRole,
    deleteRole,
  };
});
