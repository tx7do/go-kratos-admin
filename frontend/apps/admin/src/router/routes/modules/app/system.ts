import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const system: RouteRecordRaw[] = [
  {
    path: '/system',
    name: 'System',
    component: BasicLayout,
    meta: {
      order: 2004,
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
          icon: 'lucide:file-stack',
          title: $t('menu.system.file'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/system/file/index.vue'),
      },

      {
        path: 'tasks',
        name: 'TaskManagement',
        meta: {
          icon: 'lucide:calendar-clock',
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
        path: 'notifications',
        name: 'NotificationMessageManagement',
        meta: {
          icon: 'lucide:bell',
          title: $t('menu.system.notificationMessage'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/system/notification_message/index.vue'),
      },

      {
        path: 'notification_categories',
        name: 'NotificationMessageCategoryManagement',
        meta: {
          icon: 'lucide:bell-dot',
          title: $t('menu.system.notificationMessageCategory'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/system/notification_message_category/index.vue'),
      },

      {
        path: 'private_messages',
        name: 'PrivateMessageManagement',
        meta: {
          icon: 'lucide:message-circle-more',
          title: $t('menu.system.privateMessage'),
          authority: ['super', 'admin'],
        },
        component: () => import('#/views/app/system/private_message/index.vue'),
      },

      {
        path: 'admin_login_restriction',
        name: 'AdminLoginRestrictionManagement',
        meta: {
          icon: 'lucide:door-open',
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
