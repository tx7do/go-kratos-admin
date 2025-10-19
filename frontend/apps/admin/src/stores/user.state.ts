import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  UserAuthority,
  UserGender,
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
 * @param authority 权限值
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
 * @param authority 权限值
 */
export function authorityToColor(authority: any) {
  switch (authority) {
    case UserAuthority.CUSTOMER_USER: {
      // 普通客户用户：基础权限，友好绿色
      return '#52C41A';
    } // 柔和绿色（Antd success色，体现“正常、常规”）
    case UserAuthority.GUEST: {
      // 访客用户：最低权限，浅灰弱化
      return '#C9CDD4';
    } // 浅灰色（视觉上弱化，体现“临时、受限”）
    case UserAuthority.SYS_ADMIN: {
      // 系统管理员：最高权限，深蓝稳重
      return '#1890FF';
    } // 深蓝色（Antd primary色，体现“全局控制、专业”）
    case UserAuthority.TENANT_ADMIN: {
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
  { value: 'ON', label: $t('enum.status.ON') },
  { value: 'OFF', label: $t('enum.status.OFF') },
]);

/**
 * 状态转名称
 * @param status 状态值
 */
export function statusToName(status: any) {
  switch (status) {
    case 'OFF': {
      return $t('enum.status.OFF');
    }
    case 'ON': {
      return $t('enum.status.ON');
    }
    default: {
      return '';
    }
  }
}

/**
 * 状态转颜色值
 * @param status 状态值
 */
export function statusToColor(status: any) {
  switch (status) {
    case 'OFF': {
      // 关闭/停用：深灰色，明确非激活状态
      return '#8C8C8C';
    } // 中深灰色，与“关闭”语义匹配，区别于浅灰的“未知”
    case 'ON': {
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
  { value: UserGender.SECRET, label: $t('enum.gender.SECRET') },
  { value: UserGender.MALE, label: $t('enum.gender.MALE') },
  { value: UserGender.FEMALE, label: $t('enum.gender.FEMALE') },
]);

/**
 * 性别转名称
 * @param gender 性别值
 */
export function genderToName(gender: any) {
  switch (gender) {
    case UserGender.FEMALE: {
      return $t('enum.gender.FEMALE');
    }
    case UserGender.MALE: {
      return $t('enum.gender.MALE');
    }
    case UserGender.SECRET: {
      return $t('enum.gender.SECRET');
    }
    default: {
      return '';
    }
  }
}

/**
 * 性别转颜色值
 * @param gender 性别值
 */
export function genderToColor(gender: any) {
  switch (gender) {
    case UserGender.FEMALE: {
      // 女性：温和粉色，符合大众视觉认知
      return '#F77272';
    } // 柔和粉色
    case UserGender.MALE: {
      // 男性：专业蓝色，体现沉稳感
      return '#4096FF';
    } // 浅蓝色
    case UserGender.SECRET: {
      // 保密：中性灰色，代表未知
      return '#86909C';
    } // 浅灰色
    default: {
      // 异常情况：默认中性色
      return '#C9CDD4';
    } // 极浅灰色
  }
}
