import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  createDepartmentServiceClient,
  type userservicev1_Department as Department,
  type userservicev1_Department_Status as Department_Status,
} from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import { requestClientRequestHandler } from '#/utils/request';

export const useDepartmentStore = defineStore('department', () => {
  const service = createDepartmentServiceClient(requestClientRequestHandler);

  /**
   * 查询部门列表
   */
  async function listDepartment(
    noPaging: boolean = false,
    page?: number,
    pageSize?: number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await service.List({
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
   * 获取部门
   */
  async function getDepartment(id: number) {
    return await service.Get({ id });
  }

  /**
   * 创建部门
   */
  async function createDepartment(values: object) {
    return await service.Create({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新部门
   */
  async function updateDepartment(id: number, values: object) {
    return await service.Update({
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
   * 删除部门
   */
  async function deleteDepartment(id: number) {
    return await service.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listDepartment,
    getDepartment,
    createDepartment,
    updateDepartment,
    deleteDepartment,
  };
});

export const departmentStatusList = computed(() => [
  { value: 'ON', label: $t('enum.status.ON') },
  {
    value: 'OFF',
    label: $t('enum.status.OFF'),
  },
]);

/**
 * 状态转名称
 * @param status 状态值
 */
export function departmentStatusToName(status: Department_Status) {
  const values = departmentStatusList.value;
  const matchedItem = values.find((item) => item.value === status);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 状态转颜色值
 * @param status 状态值
 */
export function departmentStatusToColor(status: Department_Status) {
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

export const findDepartment = (
  list: Department[],
  id: number,
): Department | null | undefined => {
  for (const item of list) {
    if (item.id == id) {
      return item;
    }

    if (item.children && item.children.length > 0) {
      const found = findDepartment(item.children, id);
      if (found) return found;
    }
  }

  return null;
};
