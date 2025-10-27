import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  TenantAuditStatus,
  TenantStatus,
  TenantType,
} from '#/generated/api/user/service/v1/tenant.pb';
import { defTenantService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useTenantStore = defineStore('tenant', () => {
  /**
   * 查询租户列表
   */
  async function listTenant(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defTenantService.List({
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
   * 获取租户
   */
  async function getTenant(id: number) {
    return await defTenantService.Get({ id });
  }

  /**
   * 创建租户
   */
  async function createTenant(values: object) {
    return await defTenantService.Create({
      // @ts-ignore proto generated code is error.
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新租户
   */
  async function updateTenant(id: number, values: object) {
    return await defTenantService.Update({
      // @ts-ignore proto generated code is error.
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除租户
   */
  async function deleteTenant(id: number) {
    return await defTenantService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listTenant,
    getTenant,
    createTenant,
    updateTenant,
    deleteTenant,
  };
});

export const tenantTypeList = computed(() => [
  {
    value: TenantType.TENANT_TYPE_TRIAL,
    label: $t('enum.tenantType.TENANT_TYPE_TRIAL'),
  },
  {
    value: TenantType.TENANT_TYPE_PAID,
    label: $t('enum.tenantType.TENANT_TYPE_PAID'),
  },
  {
    value: TenantType.TENANT_TYPE_INTERNAL,
    label: $t('enum.tenantType.TENANT_TYPE_INTERNAL'),
  },
  {
    value: TenantType.TENANT_TYPE_PARTNER,
    label: $t('enum.tenantType.TENANT_TYPE_PARTNER'),
  },
  {
    value: TenantType.TENANT_TYPE_CUSTOM,
    label: $t('enum.tenantType.TENANT_TYPE_CUSTOM'),
  },
]);

export function tenantTypeToName(tenantType: any) {
  switch (tenantType) {
    case TenantType.TENANT_TYPE_CUSTOM: {
      return $t('enum.tenantType.TENANT_TYPE_CUSTOM');
    }
    case TenantType.TENANT_TYPE_INTERNAL: {
      return $t('enum.tenantType.TENANT_TYPE_INTERNAL');
    }
    case TenantType.TENANT_TYPE_PAID: {
      return $t('enum.tenantType.TENANT_TYPE_PAID');
    }
    case TenantType.TENANT_TYPE_PARTNER: {
      return $t('enum.tenantType.TENANT_TYPE_PARTNER');
    }
    case TenantType.TENANT_TYPE_TRIAL: {
      return $t('enum.tenantType.TENANT_TYPE_TRIAL');
    }
    default: {
      return '';
    }
  }
}

export function tenantTypeToColor(tenantType: any) {
  switch (tenantType) {
    // 定制租户：通常为深度合作的定制化客户，用深蓝色体现专业感
    case TenantType.TENANT_TYPE_CUSTOM: {
      return '#0050B3';
    }
    // 内部租户：企业内部自用租户，用官方主色调体现正式性
    case TenantType.TENANT_TYPE_INTERNAL: {
      return '#1890FF';
    }
    // 付费租户：核心付费客户，用绿色体现价值与活跃
    case TenantType.TENANT_TYPE_PAID: {
      return '#52C41A';
    }
    // 合作伙伴租户：合作关系，用紫色体现协作与独特性
    case TenantType.TENANT_TYPE_PARTNER: {
      return '#722ED1';
    }
    // 试用租户：临时试用状态，用橙色体现提醒与过渡性
    case TenantType.TENANT_TYPE_TRIAL: {
      return '#FF7D00';
    }
    // 默认值：用中性灰避免UI异常
    default: {
      return '#8C8C8C';
    }
  }
}

export const tenantStatusList = computed(() => [
  {
    value: TenantStatus.TENANT_STATUS_ON,
    label: $t('enum.tenantStatus.TENANT_STATUS_ON'),
  },
  {
    value: TenantStatus.TENANT_STATUS_OFF,
    label: $t('enum.tenantStatus.TENANT_STATUS_OFF'),
  },
  {
    value: TenantStatus.TENANT_STATUS_EXPIRED,
    label: $t('enum.tenantStatus.TENANT_STATUS_EXPIRED'),
  },
  {
    value: TenantStatus.TENANT_STATUS_FREEZE,
    label: $t('enum.tenantStatus.TENANT_STATUS_FREEZE'),
  },
]);

export function tenantStatusToName(tenantStatus: any) {
  switch (tenantStatus) {
    case TenantStatus.TENANT_STATUS_EXPIRED: {
      return $t('enum.tenantStatus.TENANT_STATUS_EXPIRED');
    }
    case TenantStatus.TENANT_STATUS_FREEZE: {
      return $t('enum.tenantStatus.TENANT_STATUS_FREEZE');
    }
    case TenantStatus.TENANT_STATUS_OFF: {
      return $t('enum.tenantStatus.TENANT_STATUS_OFF');
    }
    case TenantStatus.TENANT_STATUS_ON: {
      return $t('enum.tenantStatus.TENANT_STATUS_ON');
    }
    default: {
      return '';
    }
  }
}

export function tenantStatusToColor(tenantStatus: any) {
  switch (tenantStatus) {
    // 过期状态：租户订阅/有效期已结束，用红色体现失效
    case TenantStatus.TENANT_STATUS_EXPIRED: {
      return '#F5222D';
    }
    // 冻结状态：临时限制使用（如违规待处理），用橙色体现警告
    case TenantStatus.TENANT_STATUS_FREEZE: {
      return '#FAAD14';
    }
    // 禁用状态：主动关闭/未启用，用灰色体现非活跃
    case TenantStatus.TENANT_STATUS_OFF: {
      return '#8C8C8C';
    }
    // 正常状态：租户可正常使用，用绿色体现活跃
    case TenantStatus.TENANT_STATUS_ON: {
      return '#52C41A';
    }
    // 默认值：用中性灰避免UI异常
    default: {
      return '#8C8C8C';
    }
  }
}

export const tenantAuditStatusList = computed(() => [
  {
    value: TenantAuditStatus.TENANT_AUDIT_STATUS_PENDING,
    label: $t('enum.tenantAuditStatus.TENANT_AUDIT_STATUS_PENDING'),
  },
  {
    value: TenantAuditStatus.TENANT_AUDIT_STATUS_APPROVED,
    label: $t('enum.tenantAuditStatus.TENANT_AUDIT_STATUS_APPROVED'),
  },
  {
    value: TenantAuditStatus.TENANT_AUDIT_STATUS_REJECTED,
    label: $t('enum.tenantAuditStatus.TENANT_AUDIT_STATUS_REJECTED'),
  },
]);

export function tenantAuditStatusToName(tenantAuditStatus: any) {
  switch (tenantAuditStatus) {
    case TenantAuditStatus.TENANT_AUDIT_STATUS_APPROVED: {
      return $t('enum.tenantAuditStatus.TENANT_AUDIT_STATUS_APPROVED');
    }
    case TenantAuditStatus.TENANT_AUDIT_STATUS_PENDING: {
      return $t('enum.tenantAuditStatus.TENANT_AUDIT_STATUS_PENDING');
    }
    case TenantAuditStatus.TENANT_AUDIT_STATUS_REJECTED: {
      return $t('enum.tenantAuditStatus.TENANT_AUDIT_STATUS_REJECTED');
    }
    default: {
      return '';
    }
  }
}

export function tenantAuditStatusToColor(tenantAuditStatus: any) {
  switch (tenantAuditStatus) {
    // 已批准：审核通过，用绿色体现成功状态
    case TenantAuditStatus.TENANT_AUDIT_STATUS_APPROVED: {
      return '#52C41A';
    }
    // 待审核：审核中，用蓝色体现处理中的过渡状态
    case TenantAuditStatus.TENANT_AUDIT_STATUS_PENDING: {
      return '#1890FF';
    }
    // 已拒绝：审核未通过，用红色体现驳回状态
    case TenantAuditStatus.TENANT_AUDIT_STATUS_REJECTED: {
      return '#F5222D';
    }
    // 默认值：用中性灰避免UI异常
    default: {
      return '#8C8C8C';
    }
  }
}
