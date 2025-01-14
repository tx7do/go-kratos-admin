import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { defUserService, makeQueryString, makeUpdateMask } from '#/rpc';
import { UserAuthority } from '#/rpc/api/user/service/v1/user.pb';

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
    return await defUserService.ListUser({
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
    return await defUserService.GetUser({ id });
  }

  /**
   * 创建用户
   */
  async function createUser(values: object) {
    return await defUserService.CreateUser({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新用户
   */
  async function updateUser(id: number, values: object) {
    return await defUserService.UpdateUser({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除用户
   */
  async function deleteUser(id: number) {
    return await defUserService.DeleteUser({ id });
  }

  return {
    listUser,
    getUser,
    createUser,
    updateUser,
    deleteUser,
  };
});

export const authorityList = computed(() => [
  {
    value: UserAuthority.GUEST_USER,
    label: $t('enum.authority.GUEST_USER'),
  },
  {
    value: UserAuthority.CUSTOMER_USER,
    label: $t('enum.authority.CUSTOMER_USER'),
  },
  {
    value: UserAuthority.SYS_MANAGER,
    label: $t('enum.authority.SYS_MANAGER'),
  },
  { value: UserAuthority.SYS_ADMIN, label: $t('enum.authority.SYS_ADMIN') },
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
    case UserAuthority.GUEST_USER: {
      return $t('enum.authority.GUEST_USER');
    }
    case UserAuthority.SYS_ADMIN: {
      return $t('enum.authority.SYS_ADMIN');
    }
    case UserAuthority.SYS_MANAGER: {
      return $t('enum.authority.SYS_MANAGER');
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
      return 'green';
    }
    case UserAuthority.GUEST_USER: {
      return 'green';
    }
    case UserAuthority.SYS_ADMIN: {
      return 'orange';
    }
    case UserAuthority.SYS_MANAGER: {
      return 'red';
    }
    default: {
      return 'black';
    }
  }
}

export const statusList = computed(() => [
  { value: 'ON', label: $t('enum.status.ON') },
  { value: 'OFF', label: $t('enum.status.OFF') },
]);
