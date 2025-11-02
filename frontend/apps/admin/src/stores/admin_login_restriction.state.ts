import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { AdminLoginRestriction_Type } from '#/generated/api/admin/service/v1/i_admin_login_restriction.pb';
import { defAdminLoginRestrictionService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

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

export const adminLoginRestrictionTypeList = computed(() => [
  { value: 'BLACKLIST', label: $t('enum.adminLoginRestrictionType.BLACKLIST') },
  { value: 'WHITELIST', label: $t('enum.adminLoginRestrictionType.WHITELIST') },
]);

export const adminLoginRestrictionMethodList = computed(() => [
  { value: 'IP', label: $t('enum.adminLoginRestrictionMethod.IP') },
  { value: 'MAC', label: $t('enum.adminLoginRestrictionMethod.MAC') },
  { value: 'REGION', label: $t('enum.adminLoginRestrictionMethod.REGION') },
  { value: 'TIME', label: $t('enum.adminLoginRestrictionMethod.TIME') },
  { value: 'DEVICE', label: $t('enum.adminLoginRestrictionMethod.DEVICE') },
]);

export function adminLoginRestrictionTypeToName(typeName: any) {
  const values = adminLoginRestrictionTypeList.value;
  const matchedItem = values.find((item) => item.value === typeName);
  return matchedItem ? matchedItem.label : '';
}

export function adminLoginRestrictionTypeToColor(typeName: any) {
  switch (typeName) {
    case AdminLoginRestriction_Type.BLACKLIST: {
      return 'red'; // 黑名单用红色（表示限制/禁止）
    }
    case AdminLoginRestriction_Type.WHITELIST: {
      return 'green'; // 白名单用绿色（表示允许/信任）
    }
    default: {
      // 新增默认分支，处理未知类型，避免返回undefined
      return 'gray'; // 未知类型用灰色（中性默认值）
    }
  }
}

export function adminLoginRestrictionMethodToName(methodName: any) {
  const values = adminLoginRestrictionMethodList.value;
  const matchedItem = values.find((item) => item.value === methodName);
  return matchedItem ? matchedItem.label : '';
}
