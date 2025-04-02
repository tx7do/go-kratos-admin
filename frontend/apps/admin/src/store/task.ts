import { defineStore } from 'pinia';

import { defTaskService, makeQueryString, makeUpdateMask } from '#/rpc';

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
