<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { Department } from '#/rpc/api/user/service/v1/department.pb';

import { Page, useVbenModal, type VbenFormProps } from '@vben/common-ui';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { statusList, useDepartmentStore } from '#/store';

import DeptModal from './dept-modal.vue';

const deptStore = useDepartmentStore();

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
      label: '部门名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'status',
      label: '状态',
      componentProps: {
        options: statusList,
        placeholder: $t('ui.placeholder.select'),
      },
    },
  ],
};

const gridOptions: VxeGridProps<Department> = {
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

  treeConfig: {
    childrenField: 'children',
    rowField: 'id',
    // transform: true,
  },

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);

        return await deptStore.listDepartment(
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { title: '部门名称', field: 'name', treeNode: true },
    { title: '排序', field: 'sortId', width: 50 },
    { title: '状态', field: 'status', slots: { default: 'status' }, width: 80 },
    {
      title: '创建时间',
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: '备注', field: 'remark' },
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
  connectedComponent: DeptModal,
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
    await deptStore.deleteDepartment(row.id);

    notification.success({
      message: '删除部门成功',
    });

    await gridApi.reload();
  } catch {
    notification.error({
      message: '删除部门失败',
    });
  }
}

/* 修改部门状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await deptStore.updateDepartment(row.id, { status: row.status });

    notification.success({
      message: '更新部门状态成功',
    });
  } catch {
    notification.error({
      message: '更新部门状态失败',
    });
  } finally {
    row.pending = false;
  }
}

const expandAll = () => {
  gridApi.grid?.setAllTreeExpand(true);
};

const collapseAll = () => {
  gridApi.grid?.setAllTreeExpand(false);
};
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.system.org')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          创建部门
        </a-button>
        <a-button class="mr-2" @click="expandAll"> 展开全部</a-button>
        <a-button class="mr-2" @click="collapseAll"> 折叠全部</a-button>
      </template>
      <template #status="{ row }">
        <a-switch
          :checked="row.status === 'ON'"
          :loading="row.pending"
          checked-children="正常"
          un-checked-children="停用"
          @change="
            (checked: any) => handleStatusChanged(row, checked as boolean)
          "
        />
      </template>
      <template #action="{ row }">
        <a-button type="link" @click="() => handleEdit(row)">编辑</a-button>
        <a-popconfirm
          cancel-text="不要"
          ok-text="是的"
          title="你是否要删除掉该部门？"
          @confirm="() => handleDelete(row)"
        >
          <a-button danger type="link">删除</a-button>
        </a-popconfirm>
      </template>
    </Grid>
    <Modal />
  </Page>
</template>
