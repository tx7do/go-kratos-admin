import { defineStore } from 'pinia';

import { defMenuService, makeQueryString, makeUpdateMask } from '#/rpc';
import { MenuType } from '#/rpc/api/system/service/v1/menu.pb';

export const useMenuStore = defineStore('menu', () => {
  /**
   * 查询菜单列表
   */
  async function listMenu(
    page: number,
    pageSize: number,
    formValues: object,
    fieldMask: null | string = null,
    orderBy: string[] = [],
    noPaging: boolean = false,
  ) {
    return await defMenuService.ListMenu({
      // @ts-ignore proto generated code is error.
      fieldMask,
      orderBy,
      query: makeQueryString(formValues),
      page,
      pageSize,
      noPaging,
    });
  }

  /**
   * 创建菜单
   */
  async function createMenu(values: object) {
    return await defMenuService.CreateMenu({
      data: {
        ...values,
        children: [],
      },
    });
  }

  /**
   * 更新菜单
   */
  async function updateMenu(id: number, values: object) {
    return await defMenuService.UpdateMenu({
      data: {
        id,
        ...values,
        children: [],
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(values ?? [])),
    });
  }

  /**
   * 删除菜单
   */
  async function deleteMenu(id: number) {
    return await defMenuService.DeleteMenu({ id });
  }

  return {
    listMenu,
    createMenu,
    updateMenu,
    deleteMenu,
  };
});

export const menuTypeList = [
  { value: MenuType.FOLDER, label: '目录' },
  { value: MenuType.MENU, label: '菜单' },
  { value: MenuType.BUTTON, label: '按钮' },
];
