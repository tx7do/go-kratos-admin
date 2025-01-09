<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { User } from '#/rpc/api/user/service/v1/user.pb';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { Icon } from '@iconify/vue';
import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { defMenuService, makeQueryString, statusList } from '#/rpc';

import MenuDrawer from './menu-drawer.vue';

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
      label: '菜单名称',
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
  pagerConfig: {
    enabled: false,
  },
  rowConfig: {
    isHover: true,
  },

  // stripe: true,

  treeConfig: {
    childrenField: 'children',
    rowField: 'id',
    // transform: true,
  },

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);
        return await defMenuService.ListMenu({
          fieldMask: null,
          orderBy: [],
          query: makeQueryString(formValues),
          page: page.currentPage,
          pageSize: page.pageSize,
        });
      },
    },
  },

  columns: [
    { title: '序号', type: 'seq', width: 50 },
    {
      title: '菜单名称',
      field: 'meta.title',
      slots: { default: 'title' },
      width: 200,
      treeNode: true,
    },
    { title: '图标', field: 'icon', slots: { default: 'icon' }, width: 50 },
    { title: '排序', field: 'meta.order', width: 50 },
    // { title: '权限标识', field: 'permissionCode', width: 50 },
    { title: '路由地址', field: 'path' },
    { title: '组件路径', field: 'component' },
    { title: '状态', field: 'status', slots: { default: 'status' }, width: 80 },
    {
      title: '更新时间',
      field: 'updateTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    {
      title: '操作',
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
  connectedComponent: MenuDrawer,
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
function handleDelete(row: any) {
  console.log('删除', row);

  try {
    defMenuService.DeleteMenu({ id: row.id });

    notification.success({
      message: '删除菜单成功',
    });

    gridApi.reload();
  } catch {
    notification.error({
      message: '删除菜单失败',
    });
  }
}

/* 修改菜单状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await defMenuService.UpdateMenu({
      data: { id: row.id, status: row.status },
      updateMask: 'id,status',
    });

    notification.success({
      message: '更新菜单状态成功',
    });
  } catch {
    notification.error({
      message: '更新菜单状态失败',
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
    <Grid :table-title="$t('menu.system.menu')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          创建菜单
        </a-button>
        <a-button class="mr-2" @click="expandAll"> 展开全部 </a-button>
        <a-button class="mr-2" @click="collapseAll"> 折叠全部 </a-button>
      </template>
      <template #title="{ row }">
        <span :style="{ marginRight: '15px' }">{{ $t(row.meta.title) }}</span>
      </template>
      <template #icon="{ row }">
        <Icon
          v-if="row.meta.icon !== undefined"
          :icon="row.meta.icon"
          class="mr-1 size-4 flex-shrink-0"
        />
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
        <a-button
          type="link"
          :icon="h(LucideFilePenLine)"
          @click="() => handleEdit(row)"
        />
        <a-popconfirm
          cancel-text="不要"
          ok-text="是的"
          title="你是否要删除掉该菜单？"
          @confirm="() => handleDelete(row)"
        >
          <a-button danger type="link" :icon="h(LucideTrash2)" />
        </a-popconfirm>
      </template>
    </Grid>
    <Drawer />
  </Page>
</template>
