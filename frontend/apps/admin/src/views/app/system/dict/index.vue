<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { Dict } from '#/rpc/api/system/service/v1/dict.pb';

import { Page, useVbenModal, type VbenFormProps } from '@vben/common-ui';

import { notification } from 'ant-design-vue';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { useDictStore } from '#/store';

import DictModal from './dict-modal.vue';

const dictStore = useDictStore();

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
      fieldName: 'key',
      label: '字典键',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'categoryDesc',
      label: '字典类型名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<Dict> = {
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

        return await dictStore.listDict(
          page.currentPage,
          page.pageSize,
          formValues,
        );
      },
    },
  },

  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { title: '字典键', field: 'key' },
    { title: '字典类型', field: 'category' },
    { title: '字典类型名称', field: 'categoryDesc' },
    { title: '字典值', field: 'value' },
    { title: '字典值名称', field: 'valueDesc' },
    { title: '排序', field: 'sortId' },
    { title: '描述', field: 'remark' },
    { title: '状态', field: 'status', slots: { default: 'status' }, width: 80 },
    {
      title: '创建时间',
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 140,
    },
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
  connectedComponent: DictModal,
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
    await dictStore.deleteDict(row.id);

    notification.success({
      message: '删除字典成功',
    });

    await gridApi.reload();
  } catch {
    notification.error({
      message: '删除字典失败',
    });
  }
}

/* 修改字典状态 */
async function handleStatusChanged(row: any, checked: boolean) {
  console.log('handleStatusChanged', row.status, checked);

  row.pending = true;
  row.status = checked ? 'ON' : 'OFF';

  try {
    await dictStore.updateDict(row.id, { status: row.status });

    notification.success({
      message: '更新字典状态成功',
    });
  } catch {
    notification.error({
      message: '更新字典状态失败',
    });
  } finally {
    row.pending = false;
  }
}
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.system.dict')">
      <template #toolbar-tools>
        <a-button type="primary" @click="handleCreate">创建条目</a-button>
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
      <template #action="{ row }">
        <a-button type="link" @click="() => handleEdit(row)">编辑</a-button>
        <a-popconfirm
          cancel-text="不要"
          ok-text="是的"
          title="你是否要删除掉该条目？"
          @confirm="() => handleDelete(row)"
        >
          <a-button danger type="link">删除</a-button>
        </a-popconfirm>
      </template>
    </Grid>
    <Modal />
  </Page>
</template>
