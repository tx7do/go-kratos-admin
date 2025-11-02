<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { Department_Status } from '#/generated/api/user/service/v1/department.pb';
import { Organization_Status } from '#/generated/api/user/service/v1/organization.pb';
import {
  type Position,
  Position_Status,
} from '#/generated/api/user/service/v1/position.pb';
import { $t } from '#/locales';
import {
  positionStatusToColor,
  positionStatusToName,
  statusList,
  useDepartmentStore,
  useOrganizationStore,
  usePositionStore,
} from '#/stores';

import PositionDrawer from './position-drawer.vue';

const positionStore = usePositionStore();
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
      fieldName: 'name',
      label: $t('page.position.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: $t('page.position.code'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Select',
      fieldName: 'status',
      label: $t('ui.table.status'),
      componentProps: {
        options: statusList,
        placeholder: $t('ui.placeholder.select'),
        filterOption: (input: string, option: any) =>
          option.label.toLowerCase().includes(input.toLowerCase()),
        allowClear: true,
        showSearch: true,
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'organizationId',
      label: $t('page.position.organization'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        treeNodeFilterProp: 'label',
        api: async () => {
          const result = await orgStore.listOrganization(true, null, null, {
            // parent_id: 0,
            status: Organization_Status.ON,
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
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        treeNodeFilterProp: 'label',
        api: async () => {
          const result = await deptStore.listDepartment(true, null, null, {
            // parent_id: 0,
            status: Department_Status.ON,
          });
          return result.items;
        },
      },
    },
  ],
};

const gridOptions: VxeGridProps<Position> = {
  toolbarConfig: {
    custom: true,
    export: true,
    // import: true,
    refresh: true,
    zoom: true,
  },
  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },
  treeConfig: {
    childrenField: 'children',
    rowField: 'id',
    // transform: true,
  },
  height: 'auto',
  // stripe: true,

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);

        return await positionStore.listPosition(
          true,
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: $t('page.position.name'), field: 'name', treeNode: true },
    { title: $t('page.position.code'), field: 'code' },
    { title: $t('page.position.description'), field: 'description' },
    {
      title: $t('page.position.organizationName'),
      field: 'organizationName',
      width: 150,
    },
    {
      title: $t('page.position.departmentName'),
      field: 'departmentName',
      width: 150,
    },
    { title: $t('page.position.quota'), field: 'quota', width: 80 },
    {
      title: $t('ui.table.status'),
      field: 'status',
      slots: { default: 'status' },
      width: 95,
    },
    { title: $t('ui.table.sortOrder'), field: 'sortOrder', width: 70 },
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
  connectedComponent: PositionDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
      // 关闭时，重载表格数据
      gridApi.reload();
    }
  },
});

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
    await positionStore.deletePosition(row.id);

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

/* 修改职位状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? Position_Status.ON : Position_Status.OFF;

  try {
    await positionStore.updatePosition(row.id, { status: row.status });

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

const expandAll = () => {
  gridApi.grid?.setAllTreeExpand(true);
};

const collapseAll = () => {
  gridApi.grid?.setAllTreeExpand(false);
};
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.opm.position')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.position.button.create') }}
        </a-button>
        <a-button type="default" class="mr-2" @click="expandAll">
          {{ $t('ui.tree.expand_all') }}
        </a-button>
        <a-button type="default" class="mr-2" @click="collapseAll">
          {{ $t('ui.tree.collapse_all') }}
        </a-button>
      </template>
      <template #status="{ row }">
        <a-tag :color="positionStatusToColor(row.status)">
          {{ positionStatusToName(row.status) }}
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
              moduleName: $t('page.position.moduleName'),
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
