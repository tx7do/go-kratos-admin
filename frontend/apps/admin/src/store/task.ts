import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { defTaskService, makeQueryString, makeUpdateMask } from '#/rpc';
import { TaskType } from '#/rpc/api/system/service/v1/task.pb';

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
    return await defTaskService.ListTask({
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
    return await defTaskService.GetTask({ id });
  }

  /**
   * 创建任务
   */
  async function createTask(values: object) {
    return await defTaskService.CreateTask({
      data: {
        ...values,
      },
    });
  }

  /**
   * 更新任务
   */
  async function updateTask(id: number, values: object) {
    return await defTaskService.UpdateTask({
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
    return await defTaskService.DeleteTask({ id });
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
    value: TaskType.TaskType_Periodic,
    label: $t('enum.taskType.Periodic'),
  },
  {
    value: TaskType.TaskType_Delay,
    label: $t('enum.taskType.Delay'),
  },
  {
    value: TaskType.TaskType_WaitResult,
    label: $t('enum.taskType.WaitResult'),
  },
]);

export function taskTypeToName(taskType: any) {
  switch (taskType) {
    case TaskType.TaskType_Delay: {
      return $t('enum.taskType.Delay');
    }

    case TaskType.TaskType_Periodic: {
      return $t('enum.taskType.Periodic');
    }

    case TaskType.TaskType_WaitResult: {
      return $t('enum.taskType.WaitResult');
    }
  }
}

export function taskTypeToColor(taskType: any) {
  switch (taskType) {
    case TaskType.TaskType_Delay: {
      return 'green';
    }

    case TaskType.TaskType_Periodic: {
      return 'orange';
    }

    case TaskType.TaskType_WaitResult: {
      return 'red';
    }
  }
}
