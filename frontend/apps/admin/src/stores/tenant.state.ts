import type { CreateTenantWithAdminUserRequest } from '#/generated/api/admin/service/v1/i_tenant.pb';

import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  Tenant_AuditStatus,
  Tenant_Status,
  Tenant_Type,
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
   * 创建租户及管理员用户
   * @param values
   */
  async function createTenantWithAdminUser(values: object) {
    return await defTenantService.CreateTenantWithAdminUser(
      <CreateTenantWithAdminUserRequest>values,
    );
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

  /**
   * 租户是否存在
   * @param code 租户编码
   */
  async function tenantExists(code: string) {
    return await defTenantService.TenantExists({ code });
  }

  function $reset() {}

  return {
    $reset,
    listTenant,
    getTenant,
    createTenant,
    createTenantWithAdminUser,
    updateTenant,
    deleteTenant,
    tenantExists,
  };
});

export const tenantTypeList = computed(() => [
  {
    value: Tenant_Type.TRIAL,
    label: $t('enum.tenantType.TRIAL'),
  },
  {
    value: Tenant_Type.PAID,
    label: $t('enum.tenantType.PAID'),
  },
  {
    value: Tenant_Type.INTERNAL,
    label: $t('enum.tenantType.INTERNAL'),
  },
  {
    value: Tenant_Type.PARTNER,
    label: $t('enum.tenantType.PARTNER'),
  },
  {
    value: Tenant_Type.CUSTOM,
    label: $t('enum.tenantType.CUSTOM'),
  },
]);

export function tenantTypeToName(tenantType: any) {
  switch (tenantType) {
    case Tenant_Type.CUSTOM: {
      return $t('enum.tenantType.CUSTOM');
    }
    case Tenant_Type.INTERNAL: {
      return $t('enum.tenantType.INTERNAL');
    }
    case Tenant_Type.PAID: {
      return $t('enum.tenantType.PAID');
    }
    case Tenant_Type.PARTNER: {
      return $t('enum.tenantType.PARTNER');
    }
    case Tenant_Type.TRIAL: {
      return $t('enum.tenantType.TRIAL');
    }
    default: {
      return '';
    }
  }
}

export function tenantTypeToColor(tenantType: any) {
  switch (tenantType) {
    // 定制租户：通常为深度合作的定制化客户，用深蓝色体现专业感
    case Tenant_Type.CUSTOM: {
      return '#0050B3';
    }
    // 内部租户：企业内部自用租户，用官方主色调体现正式性
    case Tenant_Type.INTERNAL: {
      return '#1890FF';
    }
    // 付费租户：核心付费客户，用绿色体现价值与活跃
    case Tenant_Type.PAID: {
      return '#52C41A';
    }
    // 合作伙伴租户：合作关系，用紫色体现协作与独特性
    case Tenant_Type.PARTNER: {
      return '#722ED1';
    }
    // 试用租户：临时试用状态，用橙色体现提醒与过渡性
    case Tenant_Type.TRIAL: {
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
    value: Tenant_Status.ON,
    label: $t('enum.tenantStatus.ON'),
  },
  {
    value: Tenant_Status.OFF,
    label: $t('enum.tenantStatus.OFF'),
  },
  {
    value: Tenant_Status.EXPIRED,
    label: $t('enum.tenantStatus.EXPIRED'),
  },
  {
    value: Tenant_Status.FREEZE,
    label: $t('enum.tenantStatus.FREEZE'),
  },
]);

export function tenantStatusToName(tenantStatus: any) {
  switch (tenantStatus) {
    case Tenant_Status.EXPIRED: {
      return $t('enum.tenantStatus.EXPIRED');
    }
    case Tenant_Status.FREEZE: {
      return $t('enum.tenantStatus.FREEZE');
    }
    case Tenant_Status.OFF: {
      return $t('enum.tenantStatus.OFF');
    }
    case Tenant_Status.ON: {
      return $t('enum.tenantStatus.ON');
    }
    default: {
      return '';
    }
  }
}

export function tenantStatusToColor(tenantStatus: any) {
  switch (tenantStatus) {
    // 过期状态：租户订阅/有效期已结束，用红色体现失效
    case Tenant_Status.EXPIRED: {
      return '#F5222D';
    }
    // 冻结状态：临时限制使用（如违规待处理），用橙色体现警告
    case Tenant_Status.FREEZE: {
      return '#FAAD14';
    }
    // 禁用状态：主动关闭/未启用，用灰色体现非活跃
    case Tenant_Status.OFF: {
      return '#8C8C8C';
    }
    // 正常状态：租户可正常使用，用绿色体现活跃
    case Tenant_Status.ON: {
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
    value: Tenant_AuditStatus.PENDING,
    label: $t('enum.tenantAuditStatus.PENDING'),
  },
  {
    value: Tenant_AuditStatus.APPROVED,
    label: $t('enum.tenantAuditStatus.APPROVED'),
  },
  {
    value: Tenant_AuditStatus.REJECTED,
    label: $t('enum.tenantAuditStatus.REJECTED'),
  },
]);

export function tenantAuditStatusToName(tenantAuditStatus: any) {
  switch (tenantAuditStatus) {
    case Tenant_AuditStatus.APPROVED: {
      return $t('enum.tenantAuditStatus.APPROVED');
    }
    case Tenant_AuditStatus.PENDING: {
      return $t('enum.tenantAuditStatus.PENDING');
    }
    case Tenant_AuditStatus.REJECTED: {
      return $t('enum.tenantAuditStatus.REJECTED');
    }
    default: {
      return '';
    }
  }
}

export function tenantAuditStatusToColor(tenantAuditStatus: any) {
  switch (tenantAuditStatus) {
    // 已批准：审核通过，用绿色体现成功状态
    case Tenant_AuditStatus.APPROVED: {
      return '#52C41A';
    }
    // 待审核：审核中，用蓝色体现处理中的过渡状态
    case Tenant_AuditStatus.PENDING: {
      return '#1890FF';
    }
    // 已拒绝：审核未通过，用红色体现驳回状态
    case Tenant_AuditStatus.REJECTED: {
      return '#F5222D';
    }
    // 默认值：用中性灰避免UI异常
    default: {
      return '#8C8C8C';
    }
  }
}
