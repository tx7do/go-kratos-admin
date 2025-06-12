import { computed } from 'vue';

import { $t } from '@vben/locales';

import { TaskType } from '#/generated/api/admin/service/v1/i_task.pb';
import { defineStore } from 'pinia';

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
    value: TaskType.PERIODIC,
    label: $t('enum.taskType.Periodic'),
  },
  {
    value: TaskType.DELAY,
    label: $t('enum.taskType.Delay'),
  },
  {
    value: TaskType.WAIT_RESULT,
    label: $t('enum.taskType.WaitResult'),
  },
]);

export function taskTypeToName(taskType: any) {
  switch (taskType) {
    case TaskType.DELAY: {
      return $t('enum.taskType.Delay');
    }

    case TaskType.PERIODIC: {
      return $t('enum.taskType.Periodic');
    }

    case TaskType.WAIT_RESULT: {
      return $t('enum.taskType.WaitResult');
    }
  }
}

export function taskTypeToColor(taskType: any) {
  switch (taskType) {
    case TaskType.DELAY: {
      return 'green';
    }

    case TaskType.PERIODIC: {
      return 'orange';
    }

    case TaskType.WAIT_RESULT: {
      return 'red';
    }
  }
}
