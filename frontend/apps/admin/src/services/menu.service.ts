import type {
  CreateMenuRequest,
  DeleteMenuRequest,
  GetMenuRequest,
  ListMenuResponse,
  Menu,
  MenuService,
  UpdateMenuRequest,
} from '#/generated/api/admin/service/v1/i_menu.pb';
import type { Empty } from '#/generated/api/google/protobuf/empty.pb';
import type { PagingRequest } from '#/generated/api/pagination/v1/pagination.pb';

import { requestClient } from '#/utils/request';

/** 后台菜单管理服务 */
class MenuServiceImpl implements MenuService {
  async Create(request: CreateMenuRequest): Promise<Empty> {
    return await requestClient.post<Empty>('/menus', request);
  }

  async Delete(request: DeleteMenuRequest): Promise<Empty> {
    return await requestClient.delete<Empty>(`/menus/${request.id}`);
  }

  async Get(request: GetMenuRequest): Promise<Menu> {
    return await requestClient.get<Menu>(`/menus/${request.id}`);
  }

  async List(request: PagingRequest): Promise<ListMenuResponse> {
    return await requestClient.get<ListMenuResponse>('/menus', {
      params: request,
    });
  }

  async Update(request: UpdateMenuRequest): Promise<Empty> {
    const id = request.data?.id;
    if (request.data !== null && request.data !== undefined) {
      request.data.id = undefined;
    }
    return await requestClient.put<Empty>(`/menus/${id}`, request);
  }
}

export const defMenuService = new MenuServiceImpl();
