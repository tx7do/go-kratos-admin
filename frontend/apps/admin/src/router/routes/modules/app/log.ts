import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const log: RouteRecordRaw[] = [
  {
    path: '/log',
    name: 'LogAuditManagement',
    component: BasicLayout,
    meta: {
      order: 2003,
      icon: 'lucide:logs',
      title: $t('menu.log.moduleName'),
      keepAlive: true,
      authority: ['super'],
    },
    children: [
      {
        path: 'login',
        name: 'AdminLoginLog',
        meta: {
          icon: 'lucide:file-symlink',
          title: $t('menu.log.adminLoginLog'),
          authority: ['super'],
        },
        component: () => import('#/views/app/log/admin_login_log/index.vue'),
      },

      {
        path: 'operation',
        name: 'AdminOperationLog',
        meta: {
          icon: 'lucide:file-sliders',
          title: $t('menu.log.adminOperationLog'),
          authority: ['super'],
        },
        component: () =>
          import('#/views/app/log/admin_operation_log/index.vue'),
      },
    ],
  },
];

export default log;
