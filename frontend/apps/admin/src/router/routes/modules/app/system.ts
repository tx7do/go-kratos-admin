import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const system: RouteRecordRaw[] = [
  {
    path: '/',
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
        path: 'menus',
        name: 'MenuManagement',
        meta: {
          icon: 'lucide:square-menu',
          title: $t('menu.system.menu'),
        },
        component: () => import('#/views/app/system/menu/index.vue'),
      },

      {
        path: 'dict',
        name: 'DictManagement',
        meta: {
          icon: 'lucide:library-big',
          title: $t('menu.system.dict'),
        },
        component: () => import('#/views/app/system/dict/index.vue'),
      },

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
        path: 'roles',
        name: 'RoleManagement',
        meta: {
          icon: 'lucide:shirt',
          title: $t('menu.system.role'),
          hideInTab: false,
        },
        component: () => import('#/views/app/system/role/index.vue'),
      },

      {
        path: 'organizations',
        name: 'OrganizationManagement',
        meta: {
          icon: 'lucide:building-2',
          title: $t('menu.system.org'),
        },
        component: () => import('#/views/app/system/org/index.vue'),
      },

      {
        path: 'departments',
        name: 'DepartmentManagement',
        meta: {
          icon: 'lucide:network',
          title: $t('menu.system.dept'),
        },
        component: () => import('#/views/app/system/dept/index.vue'),
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
