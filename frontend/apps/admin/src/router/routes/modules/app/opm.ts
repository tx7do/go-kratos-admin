import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const opm: RouteRecordRaw[] = [
  {
    path: '/opm',
    name: 'OrganizationalPersonnelManagement',
    component: BasicLayout,
    meta: {
      order: 2001,
      icon: 'lucide:users',
      title: $t('menu.opm.moduleName'),
      keepAlive: true,
      authority: ['super', 'admin'],
    },
    children: [
      {
        path: 'users',
        name: 'UserManagement',
        meta: {
          icon: 'lucide:user',
          title: $t('menu.opm.user'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/opm/user/index.vue'),
      },
      {
        path: 'users/detail/:id',
        name: 'UserDetail',
        meta: {
          hideInTab: false,
          hideInMenu: true,
          title: $t('menu.opm.userDetail'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/opm/user/detail/index.vue'),
      },

      {
        path: 'organizations',
        name: 'OrganizationManagement',
        meta: {
          icon: 'lucide:building-2',
          title: $t('menu.opm.org'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/opm/org/index.vue'),
      },

      {
        path: 'departments',
        name: 'DepartmentManagement',
        meta: {
          icon: 'lucide:network',
          title: $t('menu.opm.dept'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/opm/dept/index.vue'),
      },

      {
        path: 'positions',
        name: 'PositionManagement',
        meta: {
          icon: 'lucide:id-card',
          title: $t('menu.opm.position'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/opm/position/index.vue'),
      },
    ],
  },
];

export default opm;
