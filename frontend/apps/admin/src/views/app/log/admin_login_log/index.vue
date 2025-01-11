<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';
import type { AdminLoginLog } from '#/rpc/api/system/service/v1/admin_login_log.pb';

import { Page, type VbenFormProps } from '@vben/common-ui';

import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { $t } from '#/locales';
import {
  authorityToColor,
  authorityToName,
  successToColor,
  successToName,
  useAdminLoginLogStore
} from '#/store';

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
      fieldName: 'userName',
      label: '登录名',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'RangePicker',
      fieldName: 'loginTime',
      label: '登录时间',
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
          page.currentPage,
          page.pageSize,
          {
            userName: formValues.userName,
            login_time__gte: startTime,
            login_time__lte: endTime,
          },
        );
      },
    },
  },

  columns: [
    { title: '序号', type: 'seq', width: 50 },
    { title: '登录名', field: 'userName' },
    { title: '登录状态', field: 'success', slots: { default: 'success' } },
    {
      title: '登录时间',
      field: 'loginTime',
      formatter: 'formatDateTime',
      width: 140,
    },
    { title: '登录地', field: 'location' },
    { title: '登录平台', field: 'clientName', slots: { default: 'platform' } },
    { title: '登录地址', field: 'loginIp', width: 140 },
  ],
};

const [Grid] = useVbenVxeGrid({ gridOptions, formOptions });
</script>

<template>
  <Page auto-content-height>
    <Grid :table-title="$t('menu.log.admin_login_log')">
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
