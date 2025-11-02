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
        path: 'messages',
        name: 'InternalMessageList',
        meta: {
          icon: 'lucide:message-circle-more',
          title: $t('menu.internalMessage.internalMessage'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/internal_message/message/index.vue'),
      },

      {
        path: 'categories',
        name: 'InternalMessageCategoryManagement',
        meta: {
          icon: 'lucide:calendar-check',
          title: $t('menu.internalMessage.internalMessageCategory'),
          authority: ['super', 'admin'],
        },
        component: () =>
          import('#/views/app/internal_message/category/index.vue'),
      },
    ],
  },
];

export default internal_message;
