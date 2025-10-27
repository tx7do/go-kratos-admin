import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { Task_Type } from '#/generated/api/admin/service/v1/i_task.pb';
import { defTaskService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

export const useTaskStore = defineStore('task', () => {
  /**
   * 查询任务列表
   */
  async function listTask(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defTaskService.List({
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
   * 获取任务
   */
  async function getTask(id: number) {
    return await defTaskService.Get({ id });
  }

  /**
   * 创建任务
   */
  async function createTask(values: object) {
    return await defTaskService.Create({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新任务
   */
  async function updateTask(id: number, values: object) {
    return await defTaskService.Update({
      data: {
        id,
        ...values,
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除任务
   */
  async function deleteTask(id: number) {
    return await defTaskService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listTask,
    getTask,
    createTask,
    updateTask,
    deleteTask,
  };
});

export const enableList = computed(() => [
  { value: 'true', label: $t('enum.enable.true') },
  { value: 'false', label: $t('enum.enable.false') },
]);

export const enableBoolList = computed(() => [
  { value: true, label: $t('enum.enable.true') },
  { value: false, label: $t('enum.enable.false') },
]);

export const taskTypeList = computed(() => [
  {
    value: Task_Type.PERIODIC,
    label: $t('enum.taskType.Periodic'),
  },
  {
    value: Task_Type.DELAY,
    label: $t('enum.taskType.Delay'),
  },
  {
    value: Task_Type.WAIT_RESULT,
    label: $t('enum.taskType.WaitResult'),
  },
]);

export function taskTypeToName(taskType: any) {
  switch (taskType) {
    case Task_Type.DELAY: {
      return $t('enum.taskType.Delay');
    }

    case Task_Type.PERIODIC: {
      return $t('enum.taskType.Periodic');
    }

    case Task_Type.WAIT_RESULT: {
      return $t('enum.taskType.WaitResult');
    }
  }
}

export function taskTypeToColor(taskType: any) {
  switch (taskType) {
    case Task_Type.DELAY: {
      return 'blue'; // 延迟任务：蓝色（表示计划中、待执行的状态）
    }
    case Task_Type.PERIODIC: {
      return 'orange'; // 周期性任务：橙色（表示循环执行、持续运行的特性）
    }
    case Task_Type.WAIT_RESULT: {
      return 'purple'; // 等待结果任务：紫色（表示过渡状态、等待响应）
    }
    default: {
      return 'gray'; // 未知任务类型：灰色（默认中性色，避免返回undefined）
    }
  }
}
