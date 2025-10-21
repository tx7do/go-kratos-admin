<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { DepartmentStatus } from '#/generated/api/user/service/v1/department.pb';
import {
  departmentStatusList,
  useDepartmentStore,
  useOrganizationStore,
  useUserStore,
} from '#/stores';

const deptStore = useDepartmentStore();
const orgStore = useOrganizationStore();
const userStore = useUserStore();

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('ui.modal.create', { moduleName: $t('page.dept.moduleName') })
    : $t('ui.modal.update', { moduleName: $t('page.dept.moduleName') }),
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
      label: $t('page.dept.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'parentId',
      label: $t('page.dept.parentId'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
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
      component: 'ApiTreeSelect',
      fieldName: 'organizationId',
      label: $t('page.dept.organization'),
      rules: 'selectRequired',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
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
      component: 'ApiSelect',
      fieldName: 'managerId',
      label: $t('page.dept.managerId'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        allowClear: true,
        afterFetch: (data: { name: string; path: string }[]) => {
          return data.map((item: any) => ({
            label: item.nickname,
            value: item.id.toString(),
          }));
        },
        api: async () => {
          const result = await userStore.listUser(true, null, null, {
            // parent_id: 0,
            status: 'ON',
          });
          return result.items;
        },
      },
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
      defaultValue: DepartmentStatus.DEPARTMENT_STATUS_ON,
      label: $t('ui.table.status'),
      rules: 'selectRequired',
      componentProps: {
        optionType: 'button',
        buttonStyle: 'solid',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: departmentStatusList,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'description',
      label: $t('page.dept.description'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      label: $t('ui.table.remark'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
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
        ? deptStore.createDepartment(values)
        : deptStore.updateDepartment(data.value.row.id, values));

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
      // 关闭窗口
      drawerApi.close();
      setLoading(false);
    }
  },

  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      // 获取传入的数据
      data.value = drawerApi.getData<Record<string, any>>();

      // 为表单赋值
      baseFormApi.setValues(data.value?.row);

      setLoading(false);

      console.log('onOpenChange', data.value, data.value?.create);
    }
  },
});

function setLoading(loading: boolean) {
  drawerApi.setState({ confirmLoading: loading });
}
</script>

<template>
  <Drawer :title="getTitle">
    <BaseForm />
  </Drawer>
</template>
