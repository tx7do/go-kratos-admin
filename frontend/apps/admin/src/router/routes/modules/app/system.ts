import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const system: RouteRecordRaw[] = [
  {
    path: '/system',
    name: 'System',
    component: BasicLayout,
    meta: {
      order: 2000,
      icon: 'lucide:settings',
      title: $t('menu.system.moduleName'),
      keepAlive: true,
    },
    children: [
      {
        path: 'users',
        name: 'UserManagement',
        meta: {
          icon: 'lucide:users',
          title: $t('menu.system.user'),
        },
        component: () => import('#/views/app/system/users/index.vue'),
      },
      {
        path: 'users/detail/:id',
        name: 'UserDetail',
        meta: {
          hideInTab: false,
          hideInMenu: true,
          title: $t('menu.system.user_detail'),
        },
        component: () => import('#/views/app/system/users/detail/index.vue'),
      },

      {
        path: 'menus',
        name: 'MenuManagement',
        meta: {
          icon: 'lucide:square-menu',
          title: $t('menu.system.menu'),
        },
        component: () => import('#/views/app/system/menu/index.vue'),
      },
      {
        path: 'organizations',
        name: 'OrganizationManagement',
        meta: {
          icon: 'lucide:network',
          title: $t('menu.system.org'),
        },
        component: () => import('#/views/app/system/org/index.vue'),
      },

      {
        path: 'roles',
        name: 'RoleManagement',
        meta: {
          icon: 'lucide:shield-check',
          title: $t('menu.system.role'),
          hideInTab: false,
        },
        component: () => import('#/views/app/system/role/index.vue'),
      },
      {
        path: 'positions',
        name: 'PositionManagement',
        meta: {
          icon: 'lucide:id-card',
          title: $t('menu.system.position'),
          hideInTab: false,
        },
        component: () => import('#/views/app/system/position/index.vue'),
      },
    ],
  },
];

export default system;
