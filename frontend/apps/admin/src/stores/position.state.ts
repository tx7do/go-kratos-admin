import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { createPositionServiceClient } from '#/generated/api/admin/service/v1';
import { makeQueryString, makeUpdateMask } from '#/utils/query';
import {
  type userservicev1_Position as Position,
  type userservicev1_Position_Status as Position_Status,
  requestClientRequestHandler,
} from '#/utils/request';

export const usePositionStore = defineStore('position', () => {
  const service = createPositionServiceClient(requestClientRequestHandler);

  /**
   * 查询职位列表
   */
  async function listPosition(
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
   * 获取职位
   */
  async function getPosition(id: number) {
    return await service.Get({ id });
  }

  /**
   * 创建职位
   */
  async function createPosition(values: object) {
    return await service.Create({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新职位
   */
  async function updatePosition(id: number, values: object) {
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
   * 删除职位
   */
  async function deletePosition(id: number) {
    return await service.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listPosition,
    getPosition,
    createPosition,
    updatePosition,
    deletePosition,
  };
});

export const positionStatusList = computed(() => [
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
export function positionStatusToName(status: Position_Status) {
  const values = positionStatusList.value;
  const matchedItem = values.find((item) => item.value === status);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 状态转颜色值
 * @param status 状态值
 */
export function positionStatusToColor(status: Position_Status) {
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

export const findPosition = (
  list: Position[],
  id: number,
): null | Position | undefined => {
  for (const item of list) {
    if (item.id == id) {
      return item;
    }

    if (item.children && item.children.length > 0) {
      const found = findPosition(item.children, id);
      if (found) return found;
    }
  }

  return null;
};
