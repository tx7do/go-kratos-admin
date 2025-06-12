import { computed } from 'vue';

import { $t } from '@vben/locales';

import {
  AdminLoginRestrictionMethod,
  AdminLoginRestrictionType,
} from '#/generated/api/admin/service/v1/i_admin_login_restriction.pb';
import { defineStore } from 'pinia';

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
  switch (typeName) {
    case AdminLoginRestrictionType.BLACKLIST: {
      return $t('enum.adminLoginRestrictionType.BLACKLIST');
    }

    case AdminLoginRestrictionType.WHITELIST: {
      return $t('enum.adminLoginRestrictionType.WHITELIST');
    }
  }
}

export function adminLoginRestrictionTypeToColor(typeName: any) {
  switch (typeName) {
    case AdminLoginRestrictionType.BLACKLIST: {
      return 'red';
    }

    case AdminLoginRestrictionType.WHITELIST: {
      return 'green';
    }
  }
}

export function adminLoginRestrictionMethodToName(methodName: any) {
  switch (methodName) {
    case AdminLoginRestrictionMethod.DEVICE: {
      return $t('enum.adminLoginRestrictionMethod.DEVICE');
    }

    case AdminLoginRestrictionMethod.IP: {
      return $t('enum.adminLoginRestrictionMethod.IP');
    }

    case AdminLoginRestrictionMethod.MAC: {
      return $t('enum.adminLoginRestrictionMethod.MAC');
    }

    case AdminLoginRestrictionMethod.REGION: {
      return $t('enum.adminLoginRestrictionMethod.REGION');
    }

    case AdminLoginRestrictionMethod.TIME: {
      return $t('enum.adminLoginRestrictionMethod.TIME');
    }
  }
}
