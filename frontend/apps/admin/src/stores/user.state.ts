import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  User_Authority,
  User_Gender,
  User_Status,
} from '#/generated/api/user/service/v1/user.pb';
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

  /**
   * 用户是否存在
   * @param username 用户名
   */
  async function userExists(username: string) {
    return await defUserService.UserExists({ username });
  }

  /**
   * 修改用户密码
   * @param id 用户ID
   * @param password 用户新密码
   */
  async function editUserPassword(id: number, password: string) {
    return await defUserService.EditUserPassword({
      userId: id,
      newPassword: password,
    });
  }

  function $reset() {}

  return {
    $reset,
    listUser,
    getUser,
    createUser,
    updateUser,
    deleteUser,
    editUserPassword,
    userExists,
  };
});

export const authorityList = computed(() => [
  {
    value: User_Authority.GUEST,
    label: $t('enum.authority.GUEST'),
  },
  {
    value: User_Authority.CUSTOMER_USER,
    label: $t('enum.authority.CUSTOMER_USER'),
  },
  {
    value: User_Authority.TENANT_ADMIN,
    label: $t('enum.authority.TENANT_ADMIN'),
  },
  {
    value: User_Authority.SYS_ADMIN,
    label: $t('enum.authority.SYS_ADMIN'),
  },
]);

/**
 * 权限转名称
 * @param authority 权限值
 */
export function authorityToName(authority: any) {
  const values = authorityList.value;
  const matchedItem = values.find((item) => item.value === authority);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 权限转颜色值
 * @param authority 权限值
 */
export function authorityToColor(authority: any) {
  switch (authority) {
    case User_Authority.CUSTOMER_USER: {
      // 普通客户用户：基础权限，友好绿色
      return '#52C41A';
    } // 柔和绿色（Antd success色，体现“正常、常规”）
    case User_Authority.GUEST: {
      // 访客用户：最低权限，浅灰弱化
      return '#b0b0b0';
    } // 浅灰色（视觉上弱化，体现“临时、受限”）
    case User_Authority.SYS_ADMIN: {
      // 系统管理员：最高权限，深蓝稳重
      return '#1890FF';
    } // 深蓝色（Antd primary色，体现“全局控制、专业”）
    case User_Authority.TENANT_ADMIN: {
      // 租户管理员：中等权限，温和橙色
      return '#FAAD14';
    } // 柔和橙色（体现“租户内管理，范围有限”）
    default: {
      // 未知权限：中灰色
      return '#8C8C8C';
    }
  }
}

export const statusList = computed(() => [
  { value: User_Status.ON, label: $t('enum.status.ON') },
  { value: User_Status.OFF, label: $t('enum.status.OFF') },
]);

/**
 * 状态转名称
 * @param status 状态值
 */
export function statusToName(status: any) {
  const values = statusList.value;
  const matchedItem = values.find((item) => item.value === status);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 状态转颜色值
 * @param status 状态值
 */
export function statusToColor(status: any) {
  switch (status) {
    case User_Status.OFF: {
      // 关闭/停用：深灰色，明确非激活状态
      return '#8C8C8C';
    } // 中深灰色，与“关闭”语义匹配，区别于浅灰的“未知”
    case User_Status.ON: {
      // 开启/激活：标准成功绿，体现正常运行
      return '#52C41A';
    } // 对应Element Plus的success色，大众认知中的“正常”色
    default: {
      // 异常状态：浅灰色，代表未定义状态
      return '#C9CDD4';
    }
  }
}

export const genderList = computed(() => [
  { value: User_Gender.SECRET, label: $t('enum.gender.SECRET') },
  { value: User_Gender.MALE, label: $t('enum.gender.MALE') },
  { value: User_Gender.FEMALE, label: $t('enum.gender.FEMALE') },
]);

/**
 * 性别转名称
 * @param gender 性别值
 */
export function genderToName(gender: any) {
  const values = genderList.value;
  const matchedItem = values.find((item) => item.value === gender);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 性别转颜色值
 * @param gender 性别值
 */
export function genderToColor(gender: any) {
  switch (gender) {
    case User_Gender.FEMALE: {
      // 女性：温和粉色，符合大众视觉认知
      return '#F77272';
    } // 柔和粉色
    case User_Gender.MALE: {
      // 男性：专业蓝色，体现沉稳感
      return '#4096FF';
    } // 浅蓝色
    case User_Gender.SECRET: {
      // 保密：中性灰色，代表未知
      return '#86909C';
    } // 浅灰色
    default: {
      // 异常情况：默认中性色
      return '#C9CDD4';
    } // 极浅灰色
  }
}
