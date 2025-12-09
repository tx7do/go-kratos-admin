import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  createTenantServiceClient,
  type CreateTenantWithAdminUserRequest,
  type userservicev1_Tenant_AuditStatus as Tenant_AuditStatus,
  type userservicev1_Tenant_Status as Tenant_Status,
  type userservicev1_Tenant_Type as Tenant_Type,
} from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import { requestClientRequestHandler } from '#/utils/request';

export const useTenantStore = defineStore('tenant', () => {
  const service = createTenantServiceClient(requestClientRequestHandler);

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
   * 获取租户
   */
  async function getTenant(id: number) {
    return await service.Get({ id });
  }

  /**
   * 创建租户
   */
  async function createTenant(values: object) {
    return await service.Create({
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
    return await service.CreateTenantWithAdminUser(
      <CreateTenantWithAdminUserRequest>values,
    );
  }

  /**
   * 更新租户
   */
  async function updateTenant(id: number, values: object) {
    return await service.Update({
      id,
      // @ts-ignore proto generated code is error.
      data: {
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
    return await service.Delete({ id });
  }

  /**
   * 租户是否存在
   * @param code 租户编码
   */
  async function tenantExists(code: string) {
    return await service.TenantExists({ code });
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
    value: 'TRIAL',
    label: $t('enum.tenantType.TRIAL'),
  },
  {
    value: 'PAID',
    label: $t('enum.tenantType.PAID'),
  },
  {
    value: 'INTERNAL',
    label: $t('enum.tenantType.INTERNAL'),
  },
  {
    value: 'PARTNER',
    label: $t('enum.tenantType.PARTNER'),
  },
  {
    value: 'CUSTOM',
    label: $t('enum.tenantType.CUSTOM'),
  },
]);

export function tenantTypeToName(tenantType: Tenant_Type) {
  const values = tenantTypeList.value;
  const matchedItem = values.find((item) => item.value === tenantType);
  return matchedItem ? matchedItem.label : '';
}

export function tenantTypeToColor(tenantType: Tenant_Type) {
  switch (tenantType) {
    // 定制租户：通常为深度合作的定制化客户，用深蓝色体现专业感
    case 'CUSTOM': {
      return '#0050B3';
    }
    // 内部租户：企业内部自用租户，用官方主色调体现正式性
    case 'INTERNAL': {
      return '#1890FF';
    }
    // 付费租户：核心付费客户，用绿色体现价值与活跃
    case 'PAID': {
      return '#52C41A';
    }
    // 合作伙伴租户：合作关系，用紫色体现协作与独特性
    case 'PARTNER': {
      return '#722ED1';
    }
    // 试用租户：临时试用状态，用橙色体现提醒与过渡性
    case 'TRIAL': {
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
    value: 'ON',
    label: $t('enum.tenantStatus.ON'),
  },
  {
    value: 'OFF',
    label: $t('enum.tenantStatus.OFF'),
  },
  {
    value: 'EXPIRED',
    label: $t('enum.tenantStatus.EXPIRED'),
  },
  {
    value: 'FREEZE',
    label: $t('enum.tenantStatus.FREEZE'),
  },
]);

export function tenantStatusToName(tenantStatus: Tenant_Status) {
  const values = tenantStatusList.value;
  const matchedItem = values.find((item) => item.value === tenantStatus);
  return matchedItem ? matchedItem.label : '';
}

export function tenantStatusToColor(tenantStatus: Tenant_Status) {
  switch (tenantStatus) {
    // 过期状态：租户订阅/有效期已结束，用红色体现失效
    case 'EXPIRED': {
      return '#F5222D';
    }
    // 冻结状态：临时限制使用（如违规待处理），用橙色体现警告
    case 'FREEZE': {
      return '#FAAD14';
    }
    // 禁用状态：主动关闭/未启用，用灰色体现非活跃
    case 'OFF': {
      return '#8C8C8C';
    }
    // 正常状态：租户可正常使用，用绿色体现活跃
    case 'ON': {
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
    value: 'PENDING',
    label: $t('enum.tenantAuditStatus.PENDING'),
  },
  {
    value: 'APPROVED',
    label: $t('enum.tenantAuditStatus.APPROVED'),
  },
  {
    value: 'REJECTED',
    label: $t('enum.tenantAuditStatus.REJECTED'),
  },
]);

export function tenantAuditStatusToName(tenantAuditStatus: Tenant_AuditStatus) {
  const values = tenantAuditStatusList.value;
  const matchedItem = values.find((item) => item.value === tenantAuditStatus);
  return matchedItem ? matchedItem.label : '';
}

export function tenantAuditStatusToColor(
  tenantAuditStatus: Tenant_AuditStatus,
) {
  switch (tenantAuditStatus) {
    // 已批准：审核通过，用绿色体现成功状态
    case 'APPROVED': {
      return '#52C41A';
    }
    // 待审核：审核中，用蓝色体现处理中的过渡状态
    case 'PENDING': {
      return '#1890FF';
    }
    // 已拒绝：审核未通过，用红色体现驳回状态
    case 'REJECTED': {
      return '#F5222D';
    }
    // 默认值：用中性灰避免UI异常
    default: {
      return '#8C8C8C';
    }
  }
}
