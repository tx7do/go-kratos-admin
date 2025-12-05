import { defineStore } from 'pinia';

import { createRoleServiceClient } from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import { requestClientRequestHandler } from '#/utils/request';

export const useRoleStore = defineStore('role', () => {
  const service = createRoleServiceClient(requestClientRequestHandler);

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
    return await service.List({
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
    return await service.Get({ id });
  }

  /**
   * 创建角色
   */
  async function createRole(values: object) {
    return await service.Create({
      // @ts-ignore proto generated code is error.
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
    return await service.Update({
      // @ts-ignore proto generated code is error.
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
    return await service.Delete({ id });
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

export const findRole = (list: Role[], id: number): null | Role | undefined => {
  for (const item of list) {
    if (item.id == id) {
      return item;
    }

    if (item.children && item.children.length > 0) {
      const found = findRole(item.children, id);
      if (found) return found;
    }
  }

  return null;
};
