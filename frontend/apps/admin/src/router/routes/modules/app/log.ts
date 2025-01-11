import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const log: RouteRecordRaw[] = [
  {
    path: '/log',
    name: 'Log',
    component: BasicLayout,
    meta: {
      order: 2001,
      icon: 'lucide:logs',
      title: $t('menu.log.moduleName'),
      keepAlive: true,
    },
    children: [
      {
        path: 'login',
        name: 'AdminLoginLog',
        meta: {
          icon: 'lucide:log-in',
          title: $t('menu.log.admin_login_log'),
        },
        component: () => import('#/views/app/log/admin_login_log/index.vue'),
      },

      {
        path: 'operation',
        name: 'AdminOperationLog',
        meta: {
          icon: 'lucide:arrow-up-down',
          title: $t('menu.log.admin_operation_log'),
        },
        component: () =>
          import('#/views/app/log/admin_operation_log/index.vue'),
      },
    ],
  },
];

export default log;
