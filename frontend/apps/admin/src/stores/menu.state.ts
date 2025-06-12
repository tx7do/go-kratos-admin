import { computed } from 'vue';

import { $t } from '@vben/locales';

import { type Menu, MenuType } from '#/generated/api/admin/service/v1/i_menu.pb';
import { defineStore } from 'pinia';

import { defMenuService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

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
    return await defMenuService.List({
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
    return await defMenuService.Get({ id });
  }

  /**
   * 创建菜单
   */
  async function createMenu(values: object) {
    return await defMenuService.Create({
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
    return await defMenuService.Update({
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
    return await defMenuService.Delete({ id });
  }

  function $reset() {}

  return {
    $reset,
    listMenu,
    getMenu,
    createMenu,
    updateMenu,
    deleteMenu,
  };
});

export const menuTypeList = computed(() => [
  { value: MenuType.FOLDER, label: $t('enum.menuType.FOLDER') },
  { value: MenuType.MENU, label: $t('enum.menuType.MENU') },
  { value: MenuType.BUTTON, label: $t('enum.menuType.BUTTON') },
]);

export const isFolder = (type: string) => type === MenuType.FOLDER;
export const isMenu = (type: string) => type === MenuType.MENU;
export const isButton = (type: string) => type === MenuType.BUTTON;

export function travelMenuChild(nodes: Menu[], parent: Menu): boolean {
  if (nodes === undefined) {
    return false;
  }

  if (parent.parentId === 0 || parent.parentId === undefined) {
    if (parent?.meta?.title) {
      parent.meta.title = $t(parent?.meta?.title ?? '');
    }
    nodes.push(parent);
    return true;
  }

  for (const node of nodes) {
    if (node.id === parent.parentId) {
      if (parent?.meta?.title) {
        parent.meta.title = $t(parent?.meta?.title ?? '');
      }
      node.children.push(parent);
      return true;
    }

    if (travelMenuChild(node.children, parent)) {
      return true;
    }
  }

  return false;
}

export function buildMenuTree(menus: Menu[]): Menu[] {
  const tree: Menu[] = [];

  for (const menu of menus) {
    if (!menu) {
      continue;
    }

    if (menu.parentId !== 0 && menu.parentId !== undefined) {
      continue;
    }

    if (menu?.meta?.title) {
      menu.meta.title = $t(menu?.meta?.title ?? '');
    }
    tree.push(menu);
  }

  for (const menu of menus) {
    if (!menu) {
      continue;
    }

    if (menu.parentId === 0 || menu.parentId === undefined) {
      continue;
    }

    if (travelMenuChild(tree, menu)) {
      continue;
    }

    if (menu?.meta?.title) {
      menu.meta.title = $t(menu?.meta?.title ?? '');
    }
    tree.push(menu);
  }

  return tree;
}
