<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { AdminLoginLog } from '#/generated/api/admin/service/v1/i_admin_login_log.pb';

import { Page, type VbenFormProps } from '@vben/common-ui';

import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import { successToColor, successToName, useAdminLoginLogStore } from '#/stores';

const adminLoginLogStore = useAdminLoginLogStore();

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
      label: $t('page.adminLoginLog.username'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'RangePicker',
      fieldName: 'loginTime',
      label: $t('page.adminLoginLog.loginTime'),
      componentProps: {
        showTime: true,
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<AdminLoginLog> = {
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
          formValues.loginTime !== undefined &&
          formValues.loginTime.length === 2
        ) {
          startTime = dayjs(formValues.loginTime[0]).format(
            'YYYY-MM-DD HH:mm:ss',
          );
          endTime = dayjs(formValues.loginTime[1]).format(
            'YYYY-MM-DD HH:mm:ss',
          );
          console.log(startTime, endTime);
        }

        return await adminLoginLogStore.listAdminLoginLog(
          false,
          page.currentPage,
          page.pageSize,
          {
            username: formValues.username,
            login_time__gte: startTime,
            login_time__lte: endTime,
          },
        );
      },
    },
  },

  columns: [
    { title: $t('ui.table.seq'), type: 'seq', width: 50 },
    { title: $t('page.adminLoginLog.username'), field: 'username' },
    {
      title: $t('page.adminLoginLog.success'),
      field: 'success',
      slots: { default: 'success' },
    },
    {
      title: $t('page.adminLoginLog.loginTime'),
      field: 'loginTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: $t('page.adminLoginLog.location'), field: 'location' },
    {
      title: $t('page.adminLoginLog.clientName'),
      field: 'clientName',
      slots: { default: 'platform' },
    },
    { title: $t('page.adminLoginLog.loginIp'), field: 'loginIp', width: 140 },
  ],
};

const [Grid] = useVbenVxeGrid({ gridOptions, formOptions });
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.log.adminLoginLog')">
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
