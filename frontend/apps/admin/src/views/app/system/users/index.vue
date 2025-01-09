<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { User } from '#/rpc/api/user/service/v1/user.pb';

import { Page, useVbenModal, type VbenFormProps } from '@vben/common-ui';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { router } from '#/router';
import {
  authorityToColor,
  authorityToName,
  useOrganizationStore,
  useUserStore,
} from '#/store';

import UserModal from './user-modal.vue';

const userStore = useUserStore();
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
      fieldName: 'realName',
      label: '用户名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'phone',
      label: '手机号码',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
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
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { title: '用户名', field: 'userName' },
    { title: '昵称', field: 'nickName' },
    { title: '姓名', field: 'realName' },
    { title: '邮箱', field: 'email' },
    { title: '手机', field: 'phone' },
    {
      title: '创建时间',
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: '权限', field: 'authority', slots: { default: 'authority' } },
    { title: '状态', field: 'status', slots: { default: 'status' }, width: 80 },
    {
      title: '操作',
      field: 'action',
      fixed: 'right',
      slots: { default: 'action' },
      width: 210,
    },
  ],
};

const [Grid, gridApi] = useVbenVxeGrid({ gridOptions, formOptions });

const [Modal, modalApi] = useVbenModal({
  // 连接抽离的组件
  connectedComponent: UserModal,
});

/* 打开模态窗口 */
function openModal(create: boolean, row?: any) {
  modalApi.setData({
    create,
    row,
  });

  modalApi.open();
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
    await userStore.deleteUser(row.id);

    notification.success({
      message: '删除用户成功',
    });

    await gridApi.reload();
  } catch {
    notification.error({
      message: '删除用户失败',
    });
  }
}

/* 详情 */
function handleDetail(row: any) {
  router.push(`/system/users/detail/${row.userName}`);
}

/* 修改用户状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await userStore.updateUser(row.id, { status: row.status });

    notification.success({
      message: '更新用户状态成功',
    });
  } catch {
    notification.error({
      message: '更新用户状态失败',
    });
  } finally {
    row.pending = false;
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.system.user')">
      <template #toolbar-tools>
        <a-button type="primary" @click="handleCreate">创建账号</a-button>
      </template>
      <template #status="{ row }">
        <a-switch
          :checked="row.status === 'ON'"
          :loading="row.pending"
          checked-children="正常"
          un-checked-children="停用"
          @change="
            (checked: boolean) => handleStatusChanged(row, checked as boolean)
          "
        />
      </template>
      <template #authority="{ row }">
        <a-tag :color="authorityToColor(row.authority)">
          {{ authorityToName(row.authority) }}
        </a-tag>
      </template>
      <template #action="{ row }">
        <a-button type="link" @click="() => handleDetail(row)">详情</a-button>

        <a-button type="link" @click="() => handleEdit(row)">编辑</a-button>
        <a-popconfirm
          cancel-text="不要"
          ok-text="是的"
          title="你是否要删除掉该用户？"
          @confirm="() => handleDelete(row)"
        >
          <a-button danger type="link">删除</a-button>
        </a-popconfirm>
      </template>
    </Grid>
    <Modal />
  </Page>
</template>
