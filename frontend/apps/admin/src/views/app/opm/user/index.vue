<script lang="ts" setup>
import type { VxeGridListeners, VxeGridProps } from '#/adapter/vxe-table';
import type { User } from '#/generated/api/user/service/v1/user.pb';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideInfo, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { router } from '#/router';
import {
  authorityList,
  authorityToColor,
  authorityToName,
  genderToColor,
  genderToName,
  statusToColor,
  statusToName,
  useDepartmentStore,
  useOrganizationStore,
  usePositionStore,
  useRoleStore,
  useUserStore,
} from '#/stores';

import UserDrawer from './user-drawer.vue';

const userStore = useUserStore();
const roleStore = useRoleStore();
const deptStore = useDepartmentStore();
const orgStore = useOrganizationStore();

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
      fieldName: 'username',
      label: $t('page.user.form.username'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'realname',
      label: $t('page.user.form.realname'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'mobile',
      label: $t('page.user.form.mobile'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'authority',
      label: $t('page.user.form.authority'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
        options: authorityList,
      },
    },
    {
      component: 'ApiSelect',
      fieldName: 'roleId',
      label: $t('page.user.form.roleId'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        allowClear: true,
        afterFetch: (data: { name: string; path: string }[]) => {
          return data.map((item: any) => ({
            label: item.name,
            value: item.id.toString(),
          }));
        },
        api: async () => {
          const result = await roleStore.listRole(true);
          return result.items;
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'organizationId',
      label: $t('page.position.organization'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        api: async () => {
          const result = await orgStore.listOrganization(true, null, null, {
            // parent_id: 0,
            status: 'ON',
          });
          return result.items;
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'departmentId',
      label: $t('page.position.department'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        api: async () => {
          const result = await deptStore.listDepartment(true, null, null, {
            // parent_id: 0,
            status: 'ON',
          });
          return result.items;
        },
      },
    },
  ],
};

const gridOptions: VxeGridProps<User> = {
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

        return await userStore.listUser(
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
    { title: $t('page.user.table.username'), field: 'username', width: 120 },
    { title: $t('page.user.table.realname'), field: 'realname', width: 100 },
    { title: $t('page.user.table.nickname'), field: 'nickname', width: 100 },
    {
      title: $t('ui.table.status'),
      field: 'status',
      slots: { default: 'status' },
      width: 100,
    },
    { title: $t('page.user.table.email'), field: 'email', width: 160 },
    { title: $t('page.user.table.mobile'), field: 'mobile', width: 130 },
    {
      title: $t('page.user.table.authority'),
      field: 'authority',
      slots: { default: 'authority' },
      width: 110,
    },
    {
      title: $t('page.user.table.lastLoginTime'),
      field: 'lastLoginTime',
      formatter: 'formatDateTime',
      width: 160,
    },
    {
      title: $t('ui.table.createTime'),
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 160,
    },
    { title: $t('ui.table.remark'), field: 'remark' },

    {
      title: $t('ui.table.action'),
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      width: 120,
    },
  ],
};

const gridEvents: VxeGridListeners<User> = {
  cellDblclick: ({ row }) => {
    // console.log(`cell-click: ${row.id}`);
    handleDetail(row);
  },
};

const [Grid, gridApi] = useVbenVxeGrid({
  gridOptions,
  formOptions,
  gridEvents,
});

const [Drawer, drawerApi] = useVbenDrawer({
  // 连接抽离的组件
  connectedComponent: UserDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      // 关闭时，重载表格数据
      gridApi.reload();
    }
  },
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
    await userStore.deleteUser(row.id);

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

/* 详情 */
function handleDetail(row: any) {
  router.push(`/system/users/detail/${row.username}`);
}

/* 修改用户状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await userStore.updateUser(row.id, { status: row.status });

    notification.success({
      message: $t('ui.notification.update_status_success'),
    });
  } catch {
    notification.error({
      message: $t('ui.notification.update_status_failed'),
    });
  } finally {
    row.pending = false;
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.opm.user')">
      <template #toolbar-tools>
        <a-button type="primary" @click="handleCreate">
          {{ $t('page.user.button.create') }}
        </a-button>
      </template>
      <template #status="{ row }">
        <a-tag :color="statusToColor(row.status)">
          {{ statusToName(row.status) }}
        </a-tag>
      </template>
      <template #authority="{ row }">
        <a-tag :color="authorityToColor(row.authority)">
          {{ authorityToName(row.authority) }}
        </a-tag>
      </template>
      <template #gender="{ row }">
        <a-tag :color="genderToColor(row.gender)">
          {{ genderToName(row.gender) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button
          type="link"
          :icon="h(LucideInfo)"
          @click.stop="handleDetail(row)"
        />

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
              moduleName: $t('page.user.moduleName'),
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
