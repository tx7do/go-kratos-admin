import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { UserAuthority } from '#/generated/api/user/service/v1/user.pb';
import { defUserService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useUserStore = defineStore('user', () => {
  /**
   * 查询用户列表
   */
  async function listUser(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defUserService.List({
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
   * 获取用户
   */
  async function getUser(id: number) {
    return await defUserService.Get({ id });
  }

  /**
   * 创建用户
   */
  async function createUser(values: object) {
    return await defUserService.Create({
      // @ts-ignore proto generated code is error.
      data: {
        ...values,
      },
      // @ts-ignore proto generated code is error.
      password: values.password ?? null,
    });
  }

  /**
   * 更新用户
   */
  async function updateUser(id: number, values: object) {
    const updateMask = makeUpdateMask(Object.keys(values ?? []));
    return await defUserService.Update({
      // @ts-ignore proto generated code is error.
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      password: values.password ?? null,
      // @ts-ignore proto generated code is error.
      updateMask,
    });
  }

  /**
   * 删除用户
   */
  async function deleteUser(id: number) {
    return await defUserService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listUser,
    getUser,
    createUser,
    updateUser,
    deleteUser,
  };
});

export const authorityList = computed(() => [
  {
    value: UserAuthority.GUEST,
    label: $t('enum.authority.GUEST'),
  },
  {
    value: UserAuthority.CUSTOMER_USER,
    label: $t('enum.authority.CUSTOMER_USER'),
  },
  {
    value: UserAuthority.TENANT_ADMIN,
    label: $t('enum.authority.TENANT_ADMIN'),
  },
  {
    value: UserAuthority.SYS_ADMIN,
    label: $t('enum.authority.SYS_ADMIN'),
  },
]);

/**
 * 权限转名称
 * @param authority
 */
export function authorityToName(authority: any) {
  switch (authority) {
    case UserAuthority.CUSTOMER_USER: {
      return $t('enum.authority.CUSTOMER_USER');
    }
    case UserAuthority.GUEST: {
      return $t('enum.authority.GUEST');
    }
    case UserAuthority.SYS_ADMIN: {
      return $t('enum.authority.SYS_ADMIN');
    }
    case UserAuthority.TENANT_ADMIN: {
      return $t('enum.authority.TENANT_ADMIN');
    }
    default: {
      return '';
    }
  }
}

/**
 * 权限转颜色值
 * @param authority
 */
export function authorityToColor(authority: any) {
  switch (authority) {
    case UserAuthority.CUSTOMER_USER: {
      return 'green'; // 普通客户用户：绿色（常规、正常权限）
    }
    case UserAuthority.GUEST: {
      return 'gray'; // 访客用户：灰色（最低权限，临时访问）
    }
    case UserAuthority.SYS_ADMIN: {
      return 'red'; // 系统管理员：红色（最高权限，需突出显示）
    }
    case UserAuthority.TENANT_ADMIN: {
      return 'orange'; // 租户管理员：橙色（中等权限，管理租户内资源）
    }
    default: {
      return 'gray'; // 未知权限：灰色（默认中性色）
    }
  }
}

export const statusList = computed(() => [
  { value: 'ON', label: $t('enum.status.ON') },
  { value: 'OFF', label: $t('enum.status.OFF') },
]);

export const genderList = computed(() => [
  { value: 'SECRET', label: $t('enum.gender.SECRET') },
  { value: 'MALE', label: $t('enum.gender.MALE') },
  { value: 'FEMALE', label: $t('enum.gender.FEMALE') },
]);
