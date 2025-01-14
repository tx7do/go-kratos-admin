import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import { defMenuService, makeQueryString, makeUpdateMask } from '#/rpc';
import { MenuType } from '#/rpc/api/system/service/v1/menu.pb';

export const useMenuStore = defineStore('menu', () => {
  /**
   * 查询菜单列表
   */
  async function listMenu(
    noPaging: boolean = false,
    page?: null | number,
    pageSize?: null | number,
    formValues?: null | object,
    fieldMask?: null | string,
    orderBy?: null | string[],
  ) {
    return await defMenuService.ListMenu({
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
   * 获取菜单
   */
  async function getMenu(id: number) {
    return await defMenuService.GetMenu({ id });
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
    getMenu,
    createMenu,
    updateMenu,
    deleteMenu,
  };
});

export const menuTypeList = [
  { value: MenuType.FOLDER, label: $t('enum.menuType.FOLDER') },
  { value: MenuType.MENU, label: $t('enum.menuType.FOLDER') },
  { value: MenuType.BUTTON, label: $t('enum.menuType.FOLDER') },
];
