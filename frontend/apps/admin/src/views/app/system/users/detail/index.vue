<script lang="ts" setup>
import { computed, ref } from 'vue';
import { useRoute } from 'vue-router';

import { Page, useVbenModal } from '@vben/common-ui';
import { LucideArrowLeft } from '@vben/icons';

import { router } from '#/router';

import DetailPage from './detail-page.vue';
import EditPasswordModal from './edit-password-modal.vue';
import LogPage from './log-page.vue';

const activeTab = ref('detail');

const route = useRoute();

const userId = computed(() => {
  return route.params?.id ?? -1;
});

const [Modal, modalApi] = useVbenModal({
  // 连接抽离的组件
  connectedComponent: EditPasswordModal,
});

/* 打开模态窗口 */
function openModal(create: boolean, row?: any) {
  modalApi.setData({
    create,
    row,
  });

  modalApi.open();
}

/**
 * 返回上一级页面
 */
function goBack() {
  router.push('/system/users');
}

/**
 * 禁用账户
 */
function handleBanAccount() {}

/**
 * 编辑密码
 */
function handleEditPassword() {
  openModal(true);
}
</script>

<template>
  <Page content-class="flex flex-col gap-4">
    <template #title>
      <div
        style="
          display: flex;
          justify-content: flex-start;
          align-items: center;
          gap: 10px;
        "
      >
        <a-button type="text" @click="goBack">
          <template #icon>
            <LucideArrowLeft class="text-align:center" />
          </template>
        </a-button>
        <span>{{ $t('page.user.detailTitle', { userId }) }}</span>
      </div>
    </template>
    <template #extra>
      <a-button class="mr-2" danger type="primary" @click="handleBanAccount">
        {{ $t('page.user.button.banAccount') }}
      </a-button>
      <a-button class="mr-2" type="primary" @click="handleEditPassword">
        {{ $t('page.user.button.editPassword') }}
      </a-button>
    </template>
    <template #description>
      <a-tabs
        v-model:active-key="activeTab"
        :tab-bar-style="{ marginBottom: 0 }"
      >
        <a-tab-pane key="detail" :tab="$t('page.user.tab.detail')" />
        <a-tab-pane key="log" :tab="$t('page.user.tab.log')" />
      </a-tabs>
    </template>

    <a-card v-show="activeTab === 'detail'">
      <DetailPage />
    </a-card>
    <a-card v-show="activeTab === 'log'">
      <LogPage />
    </a-card>
    <Modal />
  </Page>
</template>

<style></style>
