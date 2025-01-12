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

export const authorityList = [
  { value: UserAuthority.GUEST_USER, label: '游客' },
  { value: UserAuthority.CUSTOMER_USER, label: '普通用户' },
  { value: UserAuthority.SYS_MANAGER, label: '普通管理' },
  { value: UserAuthority.SYS_ADMIN, label: '超级管理' },
];

/**
 * 权限转名称
 * @param authority
 */
export function authorityToName(authority: any) {
  switch (authority) {
    case UserAuthority.CUSTOMER_USER: {
      return '普通用户';
    }
    case UserAuthority.GUEST_USER: {
      return '游客';
    }
    case UserAuthority.SYS_ADMIN: {
      return '超级管理';
    }
    case UserAuthority.SYS_MANAGER: {
      return '普通管理';
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

export const statusList = [
  { value: 'ON', label: '正常' },
  { value: 'OFF', label: '停用' },
];
