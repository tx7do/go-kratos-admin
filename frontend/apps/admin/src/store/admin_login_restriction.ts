import { defineStore } from 'pinia';

import {
  defAdminLoginRestrictionService,
  makeQueryString,
  makeUpdateMask,
} from '#/rpc';

export const useAdminLoginRestrictionStore = defineStore(
  'admin_login_restriction',
  () => {
    /**
     * 查询后台登录限制列表
     */
    async function listAdminLoginRestriction(
      noPaging: boolean = false,
      page?: null | number,
      pageSize?: null | number,
      formValues?: null | object,
      fieldMask?: null | string,
      orderBy?: null | string[],
    ) {
      return await defAdminLoginRestrictionService.List({
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
     * 获取后台登录限制
     */
    async function getAdminLoginRestriction(id: number) {
      return await defAdminLoginRestrictionService.Get({ id });
    }

    /**
     * 创建后台登录限制
     */
    async function createAdminLoginRestriction(values: object) {
      return await defAdminLoginRestrictionService.Create({
        data: {
          ...values,
        },
      });
    }

    /**
     * 更新后台登录限制
     */
    async function updateAdminLoginRestriction(id: number, values: object) {
      return await defAdminLoginRestrictionService.Update({
        data: {
          id,
          ...values,
        },
        // @ts-ignore proto generated code is error.
        updateMask: makeUpdateMask(Object.keys(values ?? [])),
      });
    }

    /**
     * 删除后台登录限制
     */
    async function deleteAdminLoginRestriction(id: number) {
      return await defAdminLoginRestrictionService.Delete({ id });
    }

    function $reset() {}

    return {
      $reset,
      listAdminLoginRestriction,
      getAdminLoginRestriction,
      createAdminLoginRestriction,
      updateAdminLoginRestriction,
      deleteAdminLoginRestriction,
    };
  },
);
