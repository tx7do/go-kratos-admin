import { defineStore } from 'pinia';

import { defOrganizationService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useOrganizationStore = defineStore('organization', () => {
  /**
   * 查询组织列表
   */
  async function listOrganization(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defOrganizationService.List({
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
   * 获取组织
   */
  async function getOrganization(id: number) {
    return await defOrganizationService.Get({ id });
  }

  /**
   * 创建组织
   */
  async function createOrganization(values: object) {
    return await defOrganizationService.Create({
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
    return await defOrganizationService.Update({
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
    return await defOrganizationService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listOrganization,
    getOrganization,
    createOrganization,
    updateOrganization,
    deleteOrganization,
  };
});
