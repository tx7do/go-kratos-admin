import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { createOrganizationServiceClient } from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import {
  type userservicev1_Organization as Organization,
  type userservicev1_Organization_Status as Organization_Status,
  type userservicev1_Organization_Type as Organization_Type,
  requestClientRequestHandler,
} from '#/utils/request';

export const useOrganizationStore = defineStore('organization', () => {
  const service = createOrganizationServiceClient(requestClientRequestHandler);

  /**
   * 查询组织列表
   */
  async function listOrganization(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
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
   * 获取组织
   */
  async function getOrganization(id: number) {
    return await service.Get({ id });
  }

  /**
   * 创建组织
   */
  async function createOrganization(values: object) {
    return await service.Create({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新组织
   */
  async function updateOrganization(id: number, values: object) {
    return await service.Update({
      id,
      data: {
        ...values,
        children: [],
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除组织
   */
  async function deleteOrganization(id: number) {
    return await service.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listOrganization,
    getOrganization,
    createOrganization,
    updateOrganization,
    deleteOrganization,
  };
});

export const organizationStatusList = computed(() => [
  {
    value: 'ON',
    label: $t('enum.status.ON'),
  },
  {
    value: 'OFF',
    label: $t('enum.status.OFF'),
  },
]);

/**
 * 状态转名称
 * @param status 状态值
 */
export function organizationStatusToName(status: Organization_Status) {
  const values = organizationStatusList.value;
  const matchedItem = values.find((item) => item.value === status);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 状态转颜色值
 * @param status 状态值
 */
export function organizationStatusToColor(status: Organization_Status) {
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

export const organizationTypeList = computed(() => [
  {
    value: 'GROUP',
    label: $t('enum.organizationType.GROUP'),
  },
  {
    value: 'SUBSIDIARY',
    label: $t('enum.organizationType.SUBSIDIARY'),
  },
  {
    value: 'FILIALE',
    label: $t('enum.organizationType.FILIALE'),
  },
  {
    value: 'DIVISION',
    label: $t('enum.organizationType.DIVISION'),
  },
]);

export const organizationTypeListForQuery = computed(() => [
  {
    value: 'GROUP',
    label: $t('enum.organizationType.GROUP'),
  },
  {
    value: 'SUBSIDIARY',
    label: $t('enum.organizationType.SUBSIDIARY'),
  },
  {
    value: 'FILIALE',
    label: $t('enum.organizationType.FILIALE'),
  },
  {
    value: 'DIVISION',
    label: $t('enum.organizationType.DIVISION'),
  },
]);

/**
 * 组织类型转名称
 * @param organizationType
 */
export function organizationTypeToName(organizationType: Organization_Type) {
  const values = organizationTypeList.value;
  const matchedItem = values.find((item) => item.value === organizationType);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 组织类型转颜色值
 * @param organizationType
 */
export function organizationTypeToColor(organizationType: Organization_Type) {
  switch (organizationType) {
    case 'DIVISION': {
      // 事业部
      return '#FF7D00';
    } // 橙色（活力，业务线特性）
    case 'FILIALE': {
      // 分公司
      return '#4096FF';
    } // 浅蓝色（从属集团，区域分支）
    case 'GROUP': {
      // 集团
      return '#165DFF';
    } // 深蓝色（核心，权威）
    case 'SUBSIDIARY': {
      // 子公司
      return '#722ED1';
    } // 紫色（独立法人，专业属性）
    default: {
      return 'gray'; // 未知权限：灰色（默认中性色）
    }
  }
}

export const findOrganization = (
  list: Organization[],
  id: number,
): null | Organization | undefined => {
  for (const item of list) {
    if (item.id == id) {
      return item;
    }

    if (item.children && item.children.length > 0) {
      const found = findOrganization(item.children, id);
      if (found) return found;
    }
  }

  return null;
};
