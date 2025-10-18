import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const system: RouteRecordRaw[] = [
  {
    path: '/system',
    name: 'System',
    component: BasicLayout,
    meta: {
      order: 2005,
      icon: 'lucide:settings',
      title: $t('menu.system.moduleName'),
      keepAlive: true,
      authority: ['super', 'admin'],
    },
    children: [
      {
        path: 'dict',
        name: 'DictManagement',
        meta: {
          icon: 'lucide:library-big',
          title: $t('menu.system.dict'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/system/dict/index.vue'),
      },

      {
        path: 'files',
        name: 'FileManagement',
        meta: {
          icon: 'lucide:file-search',
          title: $t('menu.system.file'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/system/file/index.vue'),
      },

      {
        path: 'tasks',
        name: 'TaskManagement',
        meta: {
          icon: 'lucide:list-todo',
          title: $t('menu.system.task'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/system/task/index.vue'),
      },

      {
        path: 'apis',
        name: 'APIResourceManagement',
        meta: {
          icon: 'lucide:route',
          title: $t('menu.system.apiResource'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/system/api_resource/index.vue'),
      },

      {
        path: 'admin_login_restriction',
        name: 'AdminLoginRestrictionManagement',
        meta: {
          icon: 'lucide:shield-x',
          title: $t('menu.system.adminLoginRestriction'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/system/admin_login_restriction/index.vue'),
      },
    ],
  },
];

export default system;
