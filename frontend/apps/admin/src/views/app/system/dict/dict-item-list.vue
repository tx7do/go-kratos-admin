<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { DictItem } from '#/generated/api/dict/service/v1/dict.pb';

import { h, watch } from 'vue';

import { useVbenDrawer, type VbenFormProps } from '@vben/common-ui';
import { LucideFilePenLine, LucideTrash2 } from '@vben/icons';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { statusToColor, statusToName, useDictStore } from '#/stores';
import { useDictViewStore } from '#/views/app/system/dict/dict_view.state';

import DictItemDrawer from './dict-item-drawer.vue';

const dictStore = useDictStore();
const dictViewStore = useDictViewStore();

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
      fieldName: 'code',
      label: $t('page.dict.code'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<DictItem> = {
  toolbarConfig: {
    custom: false,
    export: true,
    import: true,
    refresh: true,
    zoom: false,
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
        return await dictViewStore.fetchItemList(
          dictViewStore.currentMainId,
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: $t('page.dict.name'), field: 'name' },
    { title: $t('page.dict.code'), field: 'code' },
    { title: $t('page.dict.value'), field: 'value' },
    { title: $t('ui.table.sortId'), field: 'sortId' },
    {
      title: $t('ui.table.status'),
      field: 'status',
      slots: { default: 'status' },
      width: 95,
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
  connectedComponent: DictItemDrawer,

  onOpenChange(isOpen: boolean) {
    if (!isOpen) {
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
    await dictStore.deleteDictItem([row.id]);

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

watch(
  () => dictViewStore.currentMainId,
  () => {
    gridApi.reload();
  },
);
</script>

<template>
  <Grid :table-title="$t('page.dict.dictItemList')">
    <template #toolbar-tools>
      <a-button type="primary" @click="handleCreate">
        {{ $t('page.dict.button.create') }}
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
            moduleName: $t('page.dict.moduleName'),
          })
        "
        @confirm="handleDelete(row)"
      >
        <a-button danger type="link" :icon="h(LucideTrash2)" />
      </a-popconfirm>
    </template>
  </Grid>
  <Drawer />
</template>
