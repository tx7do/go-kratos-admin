<script lang="ts" setup>
import type { NotificationItem } from '@vben/layouts';

import { computed, ref, watch } from 'vue';

import { AuthenticationLoginExpiredModal } from '@vben/common-ui';
import { VBEN_DOC_URL, VBEN_GITHUB_URL } from '@vben/constants';
import { useWatermark } from '@vben/hooks';
import { BookOpenText, CircleHelp, MdiGithub } from '@vben/icons';
import {
  BasicLayout,
  LockScreen,
  Notification,
  UserDropdown,
} from '@vben/layouts';
import { preferences } from '@vben/preferences';
import { useAccessStore, useUserStore } from '@vben/stores';
import { dateUtil, openWindow } from '@vben/utils';

import { notification } from 'ant-design-vue';

import { InternalMessageRecipient_Status } from '#/generated/api/internal_message/service/v1/internal_message.pb';
import { $t } from '#/locales';
import {
  authorityToName,
  useAuthStore,
  useInternalMessageStore,
} from '#/stores';
import LoginForm from '#/views/_core/authentication/login.vue';

const userStore = useUserStore();
const authStore = useAuthStore();
const accessStore = useAccessStore();
const internalMessageStore = useInternalMessageStore();

const notifications = ref<NotificationItem[]>([]);

const { destroyWatermark, updateWatermark } = useWatermark();
const showDot = computed(() =>
  notifications.value.some((item) => !item.isRead),
);

const menus = computed(() => [
  {
    handler: () => {
      openWindow(VBEN_DOC_URL, {
        target: '_blank',
      });
    },
    icon: BookOpenText,
    text: $t('ui.widgets.document'),
  },
  {
    handler: () => {
      openWindow(VBEN_GITHUB_URL, {
        target: '_blank',
      });
    },
    icon: MdiGithub,
    text: 'GitHub',
  },
  {
    handler: () => {
      openWindow(`${VBEN_GITHUB_URL}/issues`, {
        target: '_blank',
      });
    },
    icon: CircleHelp,
    text: $t('ui.widgets.qa'),
  },
]);

const avatar = computed(() => {
  return userStore.userInfo?.avatar ?? preferences.app.defaultAvatar;
});

/**
 * 重载用户收件箱列表
 */
async function reloadMessages() {
  const resp = await internalMessageStore.listUserInbox(
    1,
    5,
    {
      recipient_user_id: userStore.userInfo?.id.toString(),
    },
    null,
    ['-created_at'],
  );

  for (const item of resp.items) {
    const date = dateUtil(item.createdAt as string).fromNow();
    notifications.value.push({
      id: item.id ?? 0,
      avatar: preferences.app.defaultAvatar,
      date,
      isRead: item.status === InternalMessageRecipient_Status.READ,
      message: item.content || '',
      title: item.title || '',
    });
  }
}

/**
 * 登出账号
 */
async function handleLogout() {
  await authStore.logout(false);
}

/**
 * 清空通知
 */
function handleNoticeClear() {
  notifications.value = [];
}

/**
 * 标记为已读
 * @param item
 */
function handleMarkAsRead(item: NotificationItem) {
  if (item.isRead) {
    return;
  }

  try {
    internalMessageStore.markNotificationAsRead(userStore.userInfo?.id ?? 0, [
      item.id,
    ]);

    notification.success({
      message: $t('ui.notification.update_success'),
    });
  } catch {
    notification.error({
      message: $t('ui.notification.update_failed'),
    });
  } finally {
    for (const n of notifications.value) {
      if (n.id === item.id) {
        n.isRead = true;
      }
    }
  }
}

/**
 * 全部通知标识为已读
 */
function handleMakeAll() {
  const ids: number[] = [];
  for (const item of notifications.value) {
    if (!item.isRead) {
      ids.push(item.id);
    }
  }

  if (ids.length === 0) {
    return;
  }

  try {
    internalMessageStore.markNotificationAsRead(
      userStore.userInfo?.id ?? 0,
      ids,
    );

    notification.success({
      message: $t('ui.notification.update_success'),
    });
  } catch {
    notification.error({
      message: $t('ui.notification.update_failed'),
    });
  } finally {
    notifications.value.forEach((item) => (item.isRead = true));
  }
}

// setDemoData();
reloadMessages();

watch(
  () => preferences.app.watermark,
  async (enable) => {
    if (enable) {
      await updateWatermark({
        content: `${userStore.userInfo?.username}`,
      });
    } else {
      destroyWatermark();
    }
  },
  {
    immediate: true,
  },
);
</script>

<template>
  <BasicLayout @clear-preferences-and-logout="handleLogout">
    <template #user-dropdown>
      <UserDropdown
        :avatar
        :menus
        :text="userStore.userInfo?.realname"
        :description="userStore.userInfo?.email"
        :tag-text="authorityToName(userStore.userInfo?.authority)"
        @logout="handleLogout"
      />
    </template>
    <template #notification>
      <Notification
        :dot="showDot"
        :notifications="notifications"
        @clear="handleNoticeClear"
        @make-all="handleMakeAll"
        @read="handleMarkAsRead"
      />
    </template>
    <template #extra>
      <AuthenticationLoginExpiredModal
        v-model:open="accessStore.loginExpired"
        :avatar
      >
        <LoginForm />
      </AuthenticationLoginExpiredModal>
    </template>
    <template #lock-screen>
      <LockScreen :avatar @to-login="handleLogout" />
    </template>
  </BasicLayout>
</template>
