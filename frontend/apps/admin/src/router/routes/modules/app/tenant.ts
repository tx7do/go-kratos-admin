import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const tenant: RouteRecordRaw[] = [
  {
    path: '/tenant',
    name: 'TenantManagement',
    component: BasicLayout,
    meta: {
      order: 2000,
      icon: 'lucide:earth',
      title: $t('menu.tenant.moduleName'),
      keepAlive: true,
      authority: ['super', 'admin'],
    },
    children: [
      {
        path: 'members',
        name: 'TenantMemberManagement',
        meta: {
          icon: 'lucide:book-user',
          title: $t('menu.tenant.member'),
          hideInTab: false,
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/tenant/tenant/index.vue'),
      },
    ],
  },
];

export default tenant;
