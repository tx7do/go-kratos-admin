import type { TaskService } from '#/rpc/api/admin/service/v1/i_task.pb';
import type { Empty } from '#/rpc/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/rpc/api/pagination/v1/pagination.pb';
import type {
  ControlTaskRequest,
  CreateTaskRequest,
  DeleteTaskRequest,
  GetTaskRequest,
  ListTaskResponse,
  RestartAllTaskResponse,
  Task,
  UpdateTaskRequest,
} from '#/rpc/api/system/service/v1/task.pb';

import { requestClient } from '#/rpc/request';

/** 调度任务管理服务 */
class TaskServiceImpl implements TaskService {
  async ControlTask(request: ControlTaskRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/tasks:control', request);
  }

  async CreateTask(request: CreateTaskRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/tasks', request);
  }

  async DeleteTask(request: DeleteTaskRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/tasks/${request.id}`);
  }

  async GetTask(request: GetTaskRequest): Promise<Task> {
    return await requestClient.get<Task>(`/tasks/${request.id}`);
  }

  async ListTask(request: PagingRequest): Promise<ListTaskResponse> {
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

  async UpdateTask(request: UpdateTaskRequest): Promise<Empty> {
    const id = request.data?.id;

    console.log('UpdateTask', request.data);
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/tasks/${id}`, request);
  }
}

export const defTaskService = new TaskServiceImpl();
