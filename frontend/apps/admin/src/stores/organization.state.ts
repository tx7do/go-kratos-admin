import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  OrganizationStatus,
  OrganizationType,
} from '#/generated/api/user/service/v1/organization.pb';
import { defOrganizationService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useOrganizationStore = defineStore('organization', () => {
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
    return await defOrganizationService.List({
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
    return await defOrganizationService.Get({ id });
  }

  /**
   * 创建组织
   */
  async function createOrganization(values: object) {
    return await defOrganizationService.Create({
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
    return await defOrganizationService.Update({
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
   * 删除组织
   */
  async function deleteOrganization(id: number) {
    return await defOrganizationService.Delete({ id });
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
    value: OrganizationStatus.ORGANIZATION_STATUS_ON,
    label: $t('enum.status.ON'),
  },
  {
    value: OrganizationStatus.ORGANIZATION_STATUS_OFF,
    label: $t('enum.status.OFF'),
  },
]);

/**
 * 状态转名称
 * @param status 状态值
 */
export function organizationStatusToName(status: any) {
  switch (status) {
    case OrganizationStatus.ORGANIZATION_STATUS_OFF: {
      return $t('enum.status.OFF');
    }
    case OrganizationStatus.ORGANIZATION_STATUS_ON: {
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
export function organizationStatusToColor(status: any) {
  switch (status) {
    case OrganizationStatus.ORGANIZATION_STATUS_OFF: {
      // 关闭/停用：深灰色，明确非激活状态
      return '#8C8C8C';
    } // 中深灰色，与“关闭”语义匹配，区别于浅灰的“未知”
    case OrganizationStatus.ORGANIZATION_STATUS_ON: {
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
    value: OrganizationType.ORGANIZATION_TYPE_GROUP,
    label: $t('enum.organizationType.ORGANIZATION_TYPE_GROUP'),
  },
  {
    value: OrganizationType.ORGANIZATION_TYPE_SUBSIDIARY,
    label: $t('enum.organizationType.ORGANIZATION_TYPE_SUBSIDIARY'),
  },
  {
    value: OrganizationType.ORGANIZATION_TYPE_FILIALE,
    label: $t('enum.organizationType.ORGANIZATION_TYPE_FILIALE'),
  },
  {
    value: OrganizationType.ORGANIZATION_TYPE_DIVISION,
    label: $t('enum.organizationType.ORGANIZATION_TYPE_DIVISION'),
  },
]);

export const organizationTypeListForQuery = computed(() => [
  {
    value: 'GROUP',
    label: $t('enum.organizationType.ORGANIZATION_TYPE_GROUP'),
  },
  {
    value: 'SUBSIDIARY',
    label: $t('enum.organizationType.ORGANIZATION_TYPE_SUBSIDIARY'),
  },
  {
    value: 'FILIALE',
    label: $t('enum.organizationType.ORGANIZATION_TYPE_FILIALE'),
  },
  {
    value: 'DIVISION',
    label: $t('enum.organizationType.ORGANIZATION_TYPE_DIVISION'),
  },
]);


/**
 * 组织类型转名称
 * @param organizationType
 */
export function organizationTypeToName(organizationType: any) {
  switch (organizationType) {
    case OrganizationType.ORGANIZATION_TYPE_DIVISION: {
      return $t('enum.organizationType.ORGANIZATION_TYPE_DIVISION');
    }
    case OrganizationType.ORGANIZATION_TYPE_FILIALE: {
      return $t('enum.organizationType.ORGANIZATION_TYPE_FILIALE');
    }
    case OrganizationType.ORGANIZATION_TYPE_GROUP: {
      return $t('enum.organizationType.ORGANIZATION_TYPE_GROUP');
    }
    case OrganizationType.ORGANIZATION_TYPE_SUBSIDIARY: {
      return $t('enum.organizationType.ORGANIZATION_TYPE_SUBSIDIARY');
    }
    default: {
      return '';
    }
  }
}

/**
 * 组织类型转颜色值
 * @param organizationType
 */
export function organizationTypeToColor(organizationType: any) {
  switch (organizationType) {
    case OrganizationType.ORGANIZATION_TYPE_DIVISION: {
      // 事业部
      return '#FF7D00';
    } // 橙色（活力，业务线特性）
    case OrganizationType.ORGANIZATION_TYPE_FILIALE: {
      // 分公司
      return '#4096FF';
    } // 浅蓝色（从属集团，区域分支）
    case OrganizationType.ORGANIZATION_TYPE_GROUP: {
      // 集团
      return '#165DFF';
    } // 深蓝色（核心，权威）
    case OrganizationType.ORGANIZATION_TYPE_SUBSIDIARY: {
      // 子公司
      return '#722ED1';
    } // 紫色（独立法人，专业属性）
    default: {
      return 'gray'; // 未知权限：灰色（默认中性色）
    }
  }
}
