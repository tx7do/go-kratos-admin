<script setup lang="ts">
import { computed, ref } from 'vue';

import { Page } from '@vben/common-ui';
import { $t } from '@vben/locales';
import { formatDateTime } from '@vben/utils';

import { Avatar, Descriptions, DescriptionsItem } from 'ant-design-vue';

import {
  type User,
  User_Status,
} from '#/generated/api/user/service/v1/user.pb';
import { genderToColor, genderToName, useUserStore } from '#/stores';

const props = defineProps({
  userId: { type: Number, default: undefined },
});

const userStore = useUserStore();

const data = ref<User>();

// 获取首字母（默认用'?'）
const getFirstChar = computed(() => {
  if (!data.value?.username) return '?';
  return data.value.username.slice(0, 1).toUpperCase();
});

// 根据首字母生成固定随机色
const getAvatarColor = () => {
  const char = getFirstChar.value;
  // 1. 将字符转换为哈希值（确保同一字符结果固定）
  let hash = 0;
  for (let i = 0; i < char.length; i++) {
    hash = char.charCodeAt(i) + ((hash << 5) - hash);
  }
  // 2. 哈希值映射到HSL色相（0-360）
  const hue = Math.abs(hash % 360);
  // 3. 固定饱和度和亮度（确保颜色美观且文字清晰）
  const saturation = 60; // 60% 饱和度（鲜艳但不刺眼）
  const lightness = 45; // 45% 亮度（深色背景，白色文字更清晰）
  return `hsl(${hue}, ${saturation}%, ${lightness}%)`;
};

/**
 * 重新加载用户信息
 */
async function reload() {
  if (props.userId) {
    data.value = await userStore.getUser(props.userId);
  }
}

reload();
</script>

<template>
  <Page>
    <!-- 基本信息卡片 -->
    <a-card class="mt-4">
      <div class="basic-info-container">
        <!-- 头像与状态 -->
        <div class="avatar-section">
          <Avatar
            :src="data?.avatar ?? ''"
            class="avatar"
            :style="!data?.avatar ? { backgroundColor: getAvatarColor() } : {}"
          >
            <!-- 头像加载失败/无头像时显示姓名首字母，添加占位样式 -->
            <span class="avatar-placeholder">
              {{ data?.username?.substring(0, 1) || '?' }}
            </span>
          </Avatar>
          <a-tag
            class="status-badge"
            :color="data?.status === User_Status.ON ? 'success' : 'error'"
          >
            {{
              data?.status === User_Status.ON
                ? $t('enum.status.ON')
                : $t('enum.status.OFF')
            }}
          </a-tag>
        </div>

        <!-- 详细信息列表 -->
        <Descriptions class="info-list">
          <DescriptionsItem :label="$t('page.user.detail.desc.username')">
            {{ data?.username }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.realname')">
            {{ data?.realname }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.nickname')">
            {{ data?.nickname }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.gender')">
            <a-tag :color="genderToColor(data?.gender)">
              {{ genderToName(data?.gender) }}
            </a-tag>
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.mobile')">
            {{ data?.mobile }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.email')">
            {{ data?.email }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.orgName')">
            {{ data?.orgName }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.departmentName')">
            {{ data?.departmentName }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.positionName')">
            {{ data?.positionName }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('ui.table.createdAt')">
            {{ formatDateTime(data?.createdAt ?? '') }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.lastLoginTime')">
            {{ data?.lastLoginTime }}
          </DescriptionsItem>
          <DescriptionsItem :label="$t('page.user.detail.desc.lastLoginIp')">
            {{ data?.lastLoginIp }}
          </DescriptionsItem>
        </Descriptions>
      </div>
    </a-card>
  </Page>
</template>

<style scoped>
.basic-info-container {
  display: flex;
  gap: 32px; /* 头像与信息的间距 */
  padding: 24px;
  flex-wrap: wrap; /* 小屏幕自动换行 */
}

.avatar-section {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 16px;
}

.avatar {
  width: 140px;
  height: 140px;
  border-radius: 50%; /* 确保是正圆形（部分组件可能默认非圆形） */
  display: inline-flex;
  align-items: center;
  justify-content: center;
  overflow: hidden; /* 防止头像或文字溢出圆形 */
}

/* 首字母占位样式：占满容器并居中 */
.avatar-placeholder {
  /* 充满整个头像容器 */
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 85px;
  font-weight: 700;
  color: #fff;
  line-height: 1; /* 消除行高带来的垂直偏移 */
  text-transform: uppercase; /* 统一转为大写，视觉更规整 */
}

.status-badge {
  padding: 4px 12px;
  font-size: 14px;
}

.info-list {
  flex: 1;
  min-width: 400px; /* 确保小屏幕不挤压 */
}

/* 描述项样式优化 */
:deep(.ant-descriptions-item) {
  padding: 12px 0;
}

:deep(.ant-descriptions-item-label) {
  font-weight: 500;
  width: 120px;
}

.mt-4 {
  margin-top: 16px;
}
</style>
