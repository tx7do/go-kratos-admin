import type { RouteRecordRaw } from 'vue-router';

import { BasicLayout } from '#/layouts';
import { $t } from '#/locales';

const internal_message: RouteRecordRaw[] = [
  {
    path: '/internal_message',
    name: 'InternalMessageManagement',
    redirect: '/internal_message/notifications',
    component: BasicLayout,
    meta: {
      order: 2003,
      icon: 'lucide:mail',
      title: $t('menu.internalMessage.moduleName'),
      keepAlive: true,
      authority: ['super'],
    },
    children: [
      {
        path: 'notifications',
        name: 'NotificationMessageManagement',
        meta: {
          icon: 'lucide:bell',
          title: $t('menu.internalMessage.notificationMessage'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/internal_message/notification_message/index.vue'),
      },

      {
        path: 'notification_categories',
        name: 'NotificationMessageCategoryManagement',
        meta: {
          icon: 'lucide:calendar-check',
          title: $t('menu.internalMessage.notificationMessageCategory'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import(
            '#/views/app/internal_message/notification_message_category/index.vue'
          ),
      },

      {
        path: 'private_messages',
        name: 'PrivateMessageManagement',
        meta: {
          icon: 'lucide:message-circle-more',
          title: $t('menu.internalMessage.privateMessage'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/internal_message/private_message/index.vue'),
      },
    ],
  },
];

export default internal_message;
