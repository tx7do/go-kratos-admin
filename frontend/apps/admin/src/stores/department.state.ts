import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  type Department,
  DepartmentStatus,
} from '#/generated/api/user/service/v1/department.pb';
import { defDepartmentService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useDepartmentStore = defineStore('department', () => {
  /**
   * 查询部门列表
   */
  async function listDepartment(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defDepartmentService.List({
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
    return await defDepartmentService.Get({ id });
  }

  /**
   * 创建部门
   */
  async function createDepartment(values: object) {
    return await defDepartmentService.Create({
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
    return await defDepartmentService.Update({
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
    return await defDepartmentService.Delete({ id });
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
  { value: DepartmentStatus.DEPARTMENT_STATUS_ON, label: $t('enum.status.ON') },
  {
    value: DepartmentStatus.DEPARTMENT_STATUS_OFF,
    label: $t('enum.status.OFF'),
  },
]);

/**
 * 状态转名称
 * @param status 状态值
 */
export function departmentStatusToName(status: any) {
  switch (status) {
    case DepartmentStatus.DEPARTMENT_STATUS_OFF: {
      return $t('enum.status.OFF');
    }
    case DepartmentStatus.DEPARTMENT_STATUS_ON: {
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
export function departmentStatusToColor(status: any) {
  switch (status) {
    case DepartmentStatus.DEPARTMENT_STATUS_OFF: {
      // 关闭/停用：深灰色，明确非激活状态
      return '#8C8C8C';
    } // 中深灰色，与“关闭”语义匹配，区别于浅灰的“未知”
    case DepartmentStatus.DEPARTMENT_STATUS_ON: {
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
