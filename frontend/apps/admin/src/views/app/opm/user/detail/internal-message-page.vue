<script lang="ts" setup>
import type { VxeGridProps } from '#/adapter/vxe-table';

import { Page, type VbenFormProps } from '@vben/common-ui';

import dayjs from 'dayjs';

import { useVbenVxeGrid } from '#/adapter/vxe-table';
import { type internal_messageservicev1_InternalMessageRecipient as InternalMessageRecipient } from '#/generated/api/admin/service/v1';
import { $t } from '#/locales';
import {
  internalMessageRecipientStatusColor,
  internalMessageRecipientStatusLabel,
  useInternalMessageStore,
} from '#/stores';

const props = defineProps({
  userId: { type: Number, default: undefined },
});

const internalMessageStore = useInternalMessageStore();

const formOptions: VbenFormProps = {
  // 默认展开
  collapsed: false,
  // 控制表单是否显示折叠按钮
  showCollapseButton: false,
  // 按下回车时是否提交表单
  submitOnEnter: true,
  schema: [
    {
      component: 'RangePicker',
      fieldName: 'createdAt',
      label: $t('ui.table.createdAt'),
      componentProps: {
        showTime: true,
        allowClear: true,
      },
    },
  ],
};

const gridOptions: VxeGridProps<InternalMessageRecipient> = {
  height: 'auto',
  stripe: true,

  exportConfig: {},
  pagerConfig: {},
  rowConfig: {
    isHover: true,
  },

  proxyConfig: {
    ajax: {
      query: async ({ page }, formValues) => {
        console.log('query:', formValues);

        let startTime: any;
        let endTime: any;
        if (
          formValues.createdAt !== undefined &&
          formValues.createdAt.length === 2
        ) {
          startTime = dayjs(formValues.createdAt[0]).format(
            'YYYY-MM-DD HH:mm:ss',
          );
          endTime = dayjs(formValues.createdAt[1]).format(
            'YYYY-MM-DD HH:mm:ss',
          );
          console.log(startTime, endTime);
        }

        return await internalMessageStore.listUserInbox(
          {
            page: page.currentPage,
            pageSize: page.pageSize,
          },
          {
            recipient_user_id: props.userId?.toString(),
            create_time__gte: startTime,
            create_time__lte: endTime,
          },
        );
      },
    },
  },

  columns: [
    { title: $t('page.internalMessage.title'), field: 'title' },
    {
      title: $t('page.internalMessage.status'),
      field: 'status',
      slots: { default: 'status' },
      width: 100,
    },
    {
      title: $t('page.internalMessage.readAt'),
      field: 'readAt',
      formatter: 'formatDateTime',
      width: 140,
    },
    {
      title: $t('ui.table.createdAt'),
      field: 'createdAt',
      formatter: 'formatDateTime',
      width: 140,
    },
  ],
};

const [Grid] = useVbenVxeGrid({ gridOptions, formOptions });
</script>

<template>
  <Page auto-content-height>
    <Grid>
      <template #status="{ row }">
        <a-tag :color="internalMessageRecipientStatusColor(row.status)">
          {{ internalMessageRecipientStatusLabel(row.status) }}
        </a-tag>
      </template>
      <template #platform="{ row }">
        <span> {{ row.osName }} {{ row.browserName }}</span>
      </template>
    </Grid>
  </Page>
</template>
