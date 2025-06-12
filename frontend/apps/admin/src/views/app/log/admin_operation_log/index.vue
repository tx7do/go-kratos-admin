<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { AdminOperationLog } from '#/generated/api/admin/service/v1/i_admin_operation_log.pb';

import { Page, type VbenFormProps } from '@vben/common-ui';

import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  successToColor,
  successToName,
  useAdminOperationLogStore,
} from '#/stores';

const adminOperationLogStore = useAdminOperationLogStore();

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
      label: $t('page.adminOperationLog.username'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'RangePicker',
      fieldName: 'createTime',
      label: $t('page.adminOperationLog.createTime'),
      componentProps: {
        showTime: true,
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<AdminOperationLog> = {
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
        console.log('query:', formValues);

        let startTime: any;
        let endTime: any;
        if (
          formValues.createTime !== undefined &&
          formValues.createTime.length === 2
        ) {
          startTime = dayjs(formValues.createTime[0]).format(
            'YYYY-MM-DD HH:mm:ss',
          );
          endTime = dayjs(formValues.createTime[1]).format(
            'YYYY-MM-DD HH:mm:ss',
          );
          console.log(startTime, endTime);
        }

        return await adminOperationLogStore.listAdminOperationLog(
          false,
          page.currentPage,
          page.pageSize,
          {
            username: formValues.username,
            create_time__gte: startTime,
            create_time__lte: endTime,
          },
        );
      },
    },
  },

  columns: [
    { title: $t('ui.table.seq'), type: 'seq', width: 50 },
    { title: $t('page.adminOperationLog.username'), field: 'username' },
    {
      title: $t('page.adminOperationLog.success'),
      field: 'success',
      slots: { default: 'success' },
      width: 80,
    },
    {
      title: $t('page.adminOperationLog.createTime'),
      field: 'createTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    {
      title: $t('page.adminOperationLog.method'),
      field: 'method',
      width: 80,
    },
    { title: $t('page.adminOperationLog.path'), field: 'path' },
    { title: $t('page.adminOperationLog.location'), field: 'location' },
    {
      title: $t('page.adminOperationLog.clientName'),
      field: 'clientName',
      slots: { default: 'platform' },
    },
    {
      title: $t('page.adminOperationLog.clientIp'),
      field: 'clientIp',
      width: 140,
    },
  ],
};

const [Grid] = useVbenVxeGrid({ gridOptions, formOptions });
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.log.adminOperationLog')">
      <template #success="{ row }">
        <a-tag :color="successToColor(row.success)">
          {{ successToName(row.success, row.statusCode) }}
        </a-tag>
      </template>
      <template #platform="{ row }">
        <span> {{ row.osName }} {{ row.browserName }}</span>
      </template>
    </Grid>
  </Page>
</template>
