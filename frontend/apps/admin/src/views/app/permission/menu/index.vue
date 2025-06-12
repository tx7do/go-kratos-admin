<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { Menu } from '#/generated/api/admin/service/v1/i_menu.pb';

import { h } from 'vue';

import { Page, useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { Icon } from '@iconify/vue';
import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { statusList, useMenuStore } from '#/stores';

import MenuDrawer from './menu-drawer.vue';

const menuStore = useMenuStore();

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
      label: $t('page.menu.name'),
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
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<Menu> = {
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

  stripe: true,

  treeConfig: {
    parentField: 'parentId',
    // childrenField: 'children',
    rowField: 'id',
    transform: true,
  },

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);

        return await menuStore.listMenu(true, page.currentPage, page.pageSize, {
          'meta.title': formValues.name,
          status: formValues.status,
        });
      },
    },
  },

  columns: [
    { title: $t('ui.table.seq'), type: 'seq', width: 50 },
    {
      title: $t('page.menu.name'),
      field: 'meta.title',
      slots: { default: 'title' },
      width: 200,
      treeNode: true,
    },
    {
      title: $t('page.menu.icon'),
      field: 'meta.icon',
      slots: { default: 'icon' },
      width: 50,
    },
    { title: $t('ui.table.sortId'), field: 'meta.order', width: 70 },
    // { title: '权限标识', field: 'permissionCode', width: 50 },
    { title: $t('page.menu.path'), field: 'path' },
    { title: $t('page.menu.component'), field: 'component' },
    {
      title: $t('ui.table.status'),
      field: 'status',
      slots: { default: 'status' },
      width: 95,
    },
    {
      title: $t('ui.table.updateTime'),
      field: 'updateTime',
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
async function handleDelete(row: any) {
  console.log('删除', row);

  try {
    await menuStore.deleteMenu(row.id);

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

/* 修改菜单状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await menuStore.updateMenu(row.id, { status: row.status });

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
    <Grid :table-title="$t('menu.permission.menu')">
      <template #toolbar-tools>
        <a-button class="mr-2" type="primary" @click="handleCreate">
          {{ $t('page.menu.button.create') }}
        </a-button>
        <a-button class="mr-2" @click="expandAll">
          {{ $t('ui.tree.expand_all') }}
        </a-button>
        <a-button class="mr-2" @click="collapseAll">
          {{ $t('ui.tree.collapse_all') }}
        </a-button>
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
          :checked-children="$t('ui.switch.active')"
          :un-checked-children="$t('ui.switch.inactive')"
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
          :cancel-text="$t('ui.button.cancel')"
          :ok-text="$t('ui.button.ok')"
          :title="
            $t('ui.text.do_you_want_delete', {
              moduleName: $t('page.menu.moduleName'),
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
