import { computed } from 'vue';

import { $t } from '@vben/locales';

import { defineStore } from 'pinia';

import {
  type Menu,
  Menu_Type,
} from '#/generated/api/admin/service/v1/i_menu.pb';
import { defMenuService } from '#/services';
import { makeQueryString, makeUpdateMask } from '#/utils/query';

const parseToArray = (str: string): string[] => {
  if (!str.trim()) {
    return []; // 空输入返回空数组
  }
  // 按逗号分割，去除每个元素的前后空格，过滤空字符串
  return str
    .split(',')
    .map((item) => item.trim())
    .filter(Boolean); // 排除空字符串（如连续逗号产生的空项）
};

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

  function prepareMenuData(values: object): Menu {
    // eslint-disable-next-line unicorn/prefer-structured-clone
    const copyData: Menu = JSON.parse(JSON.stringify(values));

    // noinspection TypeScriptUnresolvedReference
    delete copyData.divider1;
    if (
      copyData.meta?.authority !== undefined &&
      copyData.meta?.authority !== null &&
      copyData.meta?.authority !== ''
    ) {
      copyData.meta.authority = parseToArray(copyData.meta?.authority);
    }

    return copyData;
  }

  /**
   * 创建菜单
   */
  async function createMenu(values: object) {
    const copyData = prepareMenuData(values);

    return await defMenuService.Create({
      data: {
        ...copyData,
        children: [],
      },
    });
  }

  /**
   * 更新菜单
   */
  async function updateMenu(id: number, values: object) {
    const copyData = prepareMenuData(values);

    console.log('updateMenu', copyData);

    return await defMenuService.Update({
      data: {
        id,
        ...copyData,
        children: [],
      },
      // @ts-ignore proto generated code is error.
      updateMask: makeUpdateMask(Object.keys(copyData ?? [])),
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
  { value: Menu_Type.FOLDER, label: $t('enum.menuType.FOLDER') },
  { value: Menu_Type.MENU, label: $t('enum.menuType.MENU') },
  { value: Menu_Type.BUTTON, label: $t('enum.menuType.BUTTON') },
  { value: Menu_Type.EMBEDDED, label: $t('enum.menuType.EMBEDDED') },
  { value: Menu_Type.LINK, label: $t('enum.menuType.LINK') },
]);

/**
 * 目录类型转名称
 * @param menuType 目录类型
 */
export function menuTypeToName(menuType: any): string {
  const values = menuTypeList.value;
  const matchedItem = values.find((item) => item.value === menuType);
  return matchedItem ? matchedItem.label : '';
}

/**
 * 菜单类型转颜色值
 * @param menuType 菜单类型枚举
 * @returns 十六进制颜色值（兼容所有UI框架）
 */
export function menuTypeToColor(menuType: any) {
  switch (menuType) {
    case Menu_Type.BUTTON: {
      // 按钮：操作型元素，醒目柔和
      return '#F56C6C';
    } // 柔和红色
    case Menu_Type.EMBEDDED: {
      // 嵌入式菜单：融合科技感
      return '#4096FF';
    } // 浅蓝色
    case Menu_Type.FOLDER: {
      // 文件夹：归类属性
      return '#27AE60';
    } // 深绿色
    case Menu_Type.LINK: {
      // 链接菜单：跳转属性
      return '#9B59B6';
    } // 紫色
    case Menu_Type.MENU: {
      // 普通菜单：基础导航
      return '#165DFF';
    } // 深蓝色
    default: {
      // 未知类型：中性色
      return '#86909C';
    } // 浅灰色
  }
}

export const isFolder = (type: string) => type === Menu_Type.FOLDER;
export const isMenu = (type: string) => type === Menu_Type.MENU;
export const isButton = (type: string) => type === Menu_Type.BUTTON;
export const isEmbedded = (type: string) => type === Menu_Type.EMBEDDED;
export const isLink = (type: string) => type === Menu_Type.LINK;

/** 遍历菜单子节点
 * @param nodes 节点列表
 * @param parent 父节点
 * @return 是否找到并添加
 */
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

/**
 * 构建菜单树
 * @param menus 菜单列表
 * @return 菜单树
 */
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
