import { defineStore } from 'pinia';

import { defOrganizationService, makeQueryString, makeUpdateMask } from '#/rpc';

export const useOrganizationStore = defineStore('organization', () => {
  /**
   * 查询组织列表
   */
  async function listOrganization(
    page?: null | number,
    pageSize?: null | number,
    queryValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
    noPaging?: boolean,
  ) {
    return await defOrganizationService.ListOrganization({
      // @ts-ignore proto generated code is error.
      fieldMask,
      orderBy: orderBy ?? [],
      query: makeQueryString(queryValues ?? null),
      page,
      pageSize,
      noPaging,
    });
  }

  /**
   * 获取组织
   */
  async function getOrganization(id: number) {
    return await defOrganizationService.GetOrganization({ id });
  }

  /**
   * 创建组织
   */
  async function createOrganization(values: object) {
    return await defOrganizationService.CreateOrganization({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新组织
   */
  async function updateOrganization(id: number, values: object) {
    return await defOrganizationService.UpdateOrganization({
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
   * 删除组织
   */
  async function deleteOrganization(id: number) {
    return await defOrganizationService.DeleteOrganization({ id });
  }

  return {
    listOrganization,
    getOrganization,
    createOrganization,
    updateOrganization,
    deleteOrganization,
  };
});
