<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { NotificationMessageCategory } from '#/generated/api/internal_message/service/v1/notification_message_category.pb';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  statusToColor,
  statusToName,
  useNotificationMessageCategoryStore,
} from '#/stores';

import NotificationMessageCategoryCategoryDrawer from './notification-message-category-drawer.vue';

const notificationMessageCategoryStore = useNotificationMessageCategoryStore();

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: true,
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.notificationMessageCategory.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<NotificationMessageCategory> = {
  toolbarConfig: {
    custom: true,
    export: true,
    // import: true,
    refresh: true,
    zoom: true,
  },
  height: 'auto',
  exportConfig: {},
  pagerConfig: {
    enabled: false,
  },
  rowConfig: {
    isHover: true,
  },

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);

        return await notificationMessageCategoryStore.listNotificationMessageCategory(
          true,
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: $t('ui.table.seq'), type: 'seq', width: 50 },
    {
      title: $t('page.notificationMessageCategory.name'),
      field: 'name',
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: $t('ui.table.remark'), field: 'remark' },
    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      width: 90,
    },
  ],
};

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

const [Drawer, drawerApi] = useVbenDrawer({
  // 连接抽离的组件
  connectedComponent: NotificationMessageCategoryCategoryDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      // 关闭时，重载表格数据
      gridApi.reload();
    }
  },
});

/* 打开模态窗口 */
function openModal(create: boolean, row?: any) {
  drawerApi.setData({
    create,
    row,
  });

  drawerApi.open();
}

/* 创建 */
function handleCreate() {
  console.log('创建');

  openModal(true);
}

/* 编辑 */
function handleEdit(row: any) {
  console.log('编辑', row);
  openModal(false, row);
}

/* 删除 */
async function handleDelete(row: any) {
  console.log('删除', row);

  try {
    await notificationMessageCategoryStore.deleteNotificationMessageCategory(
      row.id,
    );

    notification.success({
      message: $t('ui.notification.delete_success'),
    });

    await gridApi.reload();
  } catch {
    notification.error({
      message: $t('ui.notification.delete_failed'),
    });
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.internalMessage.notificationMessageCategory')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.notificationMessageCategory.button.create') }}
        </a-button>
      </template>
      <template #status="{ row }">
        <a-tag :color="statusToColor(row.status)">
          {{ statusToName(row.status) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideFilePenLine)"
          @click.stop="handleEdit(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('page.notificationMessageCategory.moduleName'),
            })
          "
          @confirm="handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
