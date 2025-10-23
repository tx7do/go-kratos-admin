import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const permission: RouteRecordRaw[] = [
  {
    path: '/permission',
    name: 'PermissionManagement',
    component: BasicLayout,
    redirect: '/permission/roles',
    meta: {
      order: 2002,
      icon: 'lucide:shield-check',
      title: $t('menu.permission.moduleName'),
      keepAlive: true,
      authority: ['super', 'admin'],
    },
    children: [
      {
        path: 'roles',
        name: 'RoleManagement',
        meta: {
          order: 1,
          icon: 'lucide:shield-user',
          title: $t('menu.permission.role'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/permission/role/index.vue'),
      },

      {
        path: 'menus',
        name: 'MenuManagement',
        meta: {
          order: 2,
          icon: 'lucide:square-menu',
          title: $t('menu.permission.menu'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/permission/menu/index.vue'),
      },
    ],
  },
];

export default permission;
