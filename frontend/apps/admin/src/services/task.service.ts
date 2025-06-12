import type {
  ControlTaskRequest,
  CreateTaskRequest,
  DeleteTaskRequest,
  GetTaskByTypeNameRequest,
  GetTaskRequest,
  ListTaskResponse,
  RestartAllTaskResponse,
  Task,
  TaskService,
  UpdateTaskRequest,
} from '#/generated/api/admin/service/v1/i_task.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import * as console from 'node:console';

import { requestClient } from '#/utils/request';

/** 调度任务管理服务 */
class TaskServiceImpl implements TaskService {
  async ControlTask(request: ControlTaskRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/tasks:control', request);
  }

  async Create(request: CreateTaskRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/tasks', request);
  }

  async Delete(request: DeleteTaskRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/tasks/${request.id}`);
  }

  async Get(request: GetTaskRequest): Promise<Task> {
    return await requestClient.get<Task>(`/tasks/${request.id}`);
  }

  GetTaskByTypeName(_request: GetTaskByTypeNameRequest): Promise<Task> {
    return Promise.resolve({} as Task);
  }

  async List(request: PagingRequest): Promise<ListTaskResponse> {
    return await requestClient.get<ListTaskResponse>('/tasks', {
      params: request,
    });
  }

  async RestartAllTask(request: Empty): Promise<RestartAllTaskResponse> {
    return await requestClient.post<RestartAllTaskResponse>(
      '/tasks:restart',
      request,
    );
  }

  async StopAllTask(_request: Empty): Promise<Empty> {
    return await requestClient.post<Empty>('/tasks:stop', _request);
  }

  async Update(request: UpdateTaskRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateTask', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/tasks/${id}`, request);
  }
}

export const defTaskService = new TaskServiceImpl();
