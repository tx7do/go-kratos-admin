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
        path: 'organizations',
        name: 'OrganizationManagement',
        meta: {
          order: 1,
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
          order: 2,
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
          order: 3,
          icon: 'lucide:id-card',
          title: $t('menu.opm.position'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/opm/position/index.vue'),
      },

      {
        path: 'users',
        name: 'UserManagement',
        meta: {
          order: 4,
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
    ],
  },
];

export default opm;
