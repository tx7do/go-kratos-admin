<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { AdminLoginRestriction } from '#/generated/api/admin/service/v1/i_admin_login_restriction.pb';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  adminLoginRestrictionMethodList,
  adminLoginRestrictionMethodToName,
  adminLoginRestrictionTypeList,
  adminLoginRestrictionTypeToColor,
  adminLoginRestrictionTypeToName,
  useAdminLoginRestrictionStore,
} from '#/stores';

import AdminLoginRestrictionDrawer from './admin-login-restriction-drawer.vue';

const adminLoginRestrictionStore = useAdminLoginRestrictionStore();

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: true,
  schema: [
    {
      component: 'Select',
      fieldName: 'type',
      label: $t('page.adminLoginRestriction.type'),
      componentProps: {
        options: adminLoginRestrictionTypeList,
        placeholder: $t('ui.placeholder.select'),
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'method',
      label: $t('page.adminLoginRestriction.method'),
      componentProps: {
        options: adminLoginRestrictionMethodList,
        placeholder: $t('ui.placeholder.select'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<AdminLoginRestriction> = {
  toolbarConfig: {
    custom: true,
    export: true,
    // import: true,
    refresh: true,
    zoom: true,
  },
  height: 'auto',
  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },
  stripe: true,

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        // console.log('query:', filters, form, formValues);

        return await adminLoginRestrictionStore.listAdminLoginRestriction(
          false,
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: $t('ui.table.seq'), type: 'seq', width: 50 },
    { title: $t('page.adminLoginRestriction.targetId'), field: 'targetId' },
    {
      title: $t('page.adminLoginRestriction.type'),
      field: 'type',
      slots: { default: 'type' },
    },
    {
      title: $t('page.adminLoginRestriction.method'),
      field: 'method',
      slots: { default: 'method' },
    },
    { title: $t('page.adminLoginRestriction.value'), field: 'value' },
    { title: $t('page.adminLoginRestriction.reason'), field: 'reason' },
    {
      title: $t('ui.table.createTime'),
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 140,
    },
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
  connectedComponent: AdminLoginRestrictionDrawer,
});

/* 打开模态窗口 */
function openDrawer(create: boolean, row?: any) {
  drawerApi.setData({
    create,
    row,
  });

  drawerApi.open();
}

/* 创建 */
function handleCreate() {
  console.log('创建');
  openDrawer(true);
}

/* 编辑 */
function handleEdit(row: any) {
  console.log('编辑', row);
  openDrawer(false, row);
}

/* 删除 */
async function handleDelete(row: any) {
  console.log('删除', row);

  try {
    await adminLoginRestrictionStore.deleteAdminLoginRestriction(row.id);

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
    <Grid :table-title="$t('menu.system.adminLoginRestriction')">
      <template #toolbar-tools>
        <a-button type="primary" class="mr-2" @click="handleCreate">
          {{ $t('page.adminLoginRestriction.button.create') }}
        </a-button>
      </template>
      <template #type="{ row }">
        <a-tag :color="adminLoginRestrictionTypeToColor(row.type)">
          {{ adminLoginRestrictionTypeToName(row.type) }}
        </a-tag>
      </template>
      <template #method="{ row }">
        <a-tag color="cyan">
          {{ adminLoginRestrictionMethodToName(row.method) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideFilePenLine)"
          @click="() => handleEdit(row)"
        />
        <a-popconfirm
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('page.adminLoginRestriction.moduleName'),
            })
          "
          @confirm="() => handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
