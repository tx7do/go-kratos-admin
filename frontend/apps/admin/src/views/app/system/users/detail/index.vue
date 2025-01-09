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
        <span>用户{{ userId }}的资料</span>
      </div>
    </template>
    <template #extra>
      <a-button class="mr-2" danger type="primary" @click="handleBanAccount">
        禁用账号
      </a-button>
      <a-button class="mr-2" type="primary" @click="handleEditPassword">
        修改密码
      </a-button>
    </template>
    <template #description>
      <a-tabs
        v-model:active-key="activeTab"
        :tab-bar-style="{ marginBottom: 0 }"
      >
        <a-tab-pane key="detail" tab="用户资料" />
        <a-tab-pane key="log" tab="操作日志" />
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
