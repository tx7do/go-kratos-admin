import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { defAdminLoginLogService } from '#/services';
import { makeQueryString } from '#/utils/query';

export const useAdminLoginLogStore = defineStore('admin_login_log', () => {
  /**
   * 查询登录日志列表
   */
  async function listAdminLoginLog(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defAdminLoginLogService.List({
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
   * 查询登录日志
   */
  async function getAdminLoginLog(id: number) {
    return await defAdminLoginLogService.Get({ id });
  }

  function $reset() {}

  return {
    $reset,
    listAdminLoginLog,
    getAdminLoginLog,
  };
});

/**
 * 成功失败的颜色
 * @param success
 */
export function successToColor(success: boolean) {
  // 成功用柔和的绿色，失败用柔和的红色，兼顾视觉舒适度与直观性
  return success ? 'limegreen' : 'crimson';
}

export function successToName(success: boolean, statusCode: number) {
  return success
    ? $t('enum.successStatus.success')
    : ` ${$t('enum.successStatus.failed')} (${statusCode})`;
}
