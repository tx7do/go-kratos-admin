<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { PositionStatus } from '#/generated/api/user/service/v1/position.pb';
import {
  positionStatusList,
  statusList,
  useDepartmentStore,
  useOrganizationStore,
  usePositionStore,
} from '#/stores';

const positionStore = usePositionStore();
const deptStore = useDepartmentStore();
const orgStore = useOrganizationStore();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('ui.modal.create', { moduleName: $t('page.position.moduleName') })
    : $t('ui.modal.update', { moduleName: $t('page.position.moduleName') }),
);
// const isCreate = computed(() => data.value?.create);

const [BaseForm, baseFormApi] = useVbenForm({
  showDefaultActions: false,
  // 所有表单项共用，可单独在表单内覆盖
  commonConfig: {
    // 所有表单项
    componentProps: {
      class: 'w-full',
    },
  },
  schema: [
    {
      component: 'Input',
      fieldName: 'name',
      label: $t('page.position.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'code',
      label: $t('page.position.code'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'parentId',
      label: $t('page.position.parentId'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        api: async () => {
          const result = await positionStore.listPosition(true, null, null, {
            // parent_id: 0,
            status: 'ON',
          });
          return result.items;
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'organizationId',
      label: $t('page.position.organization'),
      rules: 'selectRequired',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        api: async () => {
          const result = await orgStore.listOrganization(true, null, null, {
            // parent_id: 0,
            status: 'ON',
          });
          return result.items;
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'departmentId',
      label: $t('page.position.department'),
      rules: 'selectRequired',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        api: async () => {
          const result = await deptStore.listDepartment(true, null, null, {
            // parent_id: 0,
            status: 'ON',
          });
          return result.items;
        },
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'quota',
      label: $t('page.position.quota'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
        defaultValue: 1,
      },
      rules: 'required',
    },
    {
      component: 'InputNumber',
      fieldName: 'sortId',
      label: $t('ui.table.sortId'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      fieldName: 'status',
      label: $t('ui.table.status'),
      defaultValue: PositionStatus.POSITION_STATUS_ON,
      rules: 'selectRequired',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: positionStatusList,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'description',
      label: $t('page.position.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      label: $t('ui.table.remark'),
    },
  ],
});

const [Drawer, drawerApi] = useVbenDrawer({
  onCancel() {
    drawerApi.close();
  },

  async onConfirm() {
    console.log('onConfirm');

    // 校验输入的数据
    const validate = await baseFormApi.validate();
    if (!validate.valid) {
      return;
    }

    setLoading(true);

    // 获取表单数据
    const values = await baseFormApi.getValues();

    console.log(getTitle.value, values);

    try {
      await (data.value?.create
        ? positionStore.createPosition(values)
        : positionStore.updatePosition(data.value.row.id, values));

      notification.success({
        message: data.value?.create
          ? $t('ui.notification.create_success')
          : $t('ui.notification.update_success'),
      });
    } catch {
      notification.error({
        message: data.value?.create
          ? $t('ui.notification.create_failed')
          : $t('ui.notification.update_failed'),
      });
    } finally {
      drawerApi.close();
      setLoading(false);
    }
  },

  onOpenChange(isOpen) {
    if (isOpen) {
      // 获取传入的数据
      data.value = drawerApi.getData<Record<string, any>>();

      // 为表单赋值
      baseFormApi.setValues(data.value?.row);

      setLoading(false);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ loading });
}
</script>

<template>
  <Drawer :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
