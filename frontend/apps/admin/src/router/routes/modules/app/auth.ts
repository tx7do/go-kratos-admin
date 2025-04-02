import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const auth: RouteRecordRaw[] = [
  {
    path: '/auth',
    name: 'Auth',
    component: BasicLayout,
    meta: {
      order: 2000,
      icon: 'lucide:shield-check',
      title: $t('menu.auth.moduleName'),
      keepAlive: true,
      authority: ['super', 'admin'],
    },
    children: [
      {
        path: 'users',
        name: 'UserManagement',
        meta: {
          icon: 'lucide:users',
          title: $t('menu.auth.user'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/users/index.vue'),
      },
      {
        path: 'users/detail/:id',
        name: 'UserDetail',
        meta: {
          hideInTab: false,
          hideInMenu: true,
          title: $t('menu.auth.userDetail'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/users/detail/index.vue'),
      },

      {
        path: 'tenants',
        name: 'TenantManagement',
        meta: {
          icon: 'lucide:book-user',
          title: $t('menu.auth.tenant'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/tenant/index.vue'),
      },

      {
        path: 'roles',
        name: 'RoleManagement',
        meta: {
          icon: 'lucide:chef-hat',
          title: $t('menu.auth.role'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/role/index.vue'),
      },

      {
        path: 'organizations',
        name: 'OrganizationManagement',
        meta: {
          icon: 'lucide:building-2',
          title: $t('menu.auth.org'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/org/index.vue'),
      },

      {
        path: 'departments',
        name: 'DepartmentManagement',
        meta: {
          icon: 'lucide:network',
          title: $t('menu.auth.dept'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/dept/index.vue'),
      },

      {
        path: 'positions',
        name: 'PositionManagement',
        meta: {
          icon: 'lucide:id-card',
          title: $t('menu.auth.position'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/position/index.vue'),
      },

      {
        path: 'menus',
        name: 'MenuManagement',
        meta: {
          icon: 'lucide:square-menu',
          title: $t('menu.auth.menu'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/auth/menu/index.vue'),
      },
    ],
  },
];

export default auth;
