<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm, z } from '#/adapter/form';
import {
  type userservicev1_Department as Department,
  type userservicev1_Organization as Organization,
  type userservicev1_Position as Position,
} from '#/generated/api/admin/service/v1';
import {
  authorityList,
  findDepartment,
  findPosition,
  genderList,
  statusList,
  useDepartmentStore,
  useOrganizationStore,
  usePositionStore,
  useRoleStore,
  useUserStore,
} from '#/stores';

const userStore = useUserStore();
const roleStore = useRoleStore();
const orgStore = useOrganizationStore();
const positionStore = usePositionStore();
const deptStore = useDepartmentStore();

const data = ref();

const orgList = ref<Organization[]>([]);
const deptList = ref<Department[]>([]);
const positionList = ref<Position[]>([]);

const getTitle = computed(() =>
  data.value?.create
    ? $t('ui.modal.create', { moduleName: $t('page.user.moduleName') })
    : $t('ui.modal.update', { moduleName: $t('page.user.moduleName') }),
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
      fieldName: 'username',
      label: $t('page.user.table.username'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: z.string().min(1, { message: $t('ui.formRules.required') }),
      dependencies: {
        disabled: () => !data.value?.create,
        triggerFields: ['username'],
      },
    },
    {
      component: 'VbenInputPassword',
      fieldName: 'password',
      label: $t('page.user.table.password'),
      componentProps: {
        passwordStrength: true,
        placeholder: $t('ui.placeholder.input'),
      },
      // rules: 'required',
    },
    {
      component: 'Select',
      fieldName: 'authority',
      label: $t('page.user.table.authority'),
      defaultValue: 'CUSTOMER_USER',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        options: authorityList,
        filterOption: (input: string, option: any) =>
          option.label.toLowerCase().includes(input.toLowerCase()),
        allowClear: true,
        showSearch: true,
      },
      rules: 'selectRequired',
      dependencies: {
        disabled: () => !data.value?.create,
        triggerFields: ['authority'],
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'roleIds',
      label: $t('page.user.form.role'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        showSearch: true,
        multiple: true,
        treeDefaultExpandAll: false,
        allowClear: true,
        loadingSlot: 'suffixIcon',
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        treeNodeFilterProp: 'label',
        api: async () => {
          const result = await roleStore.listRole(undefined, {
            // parent_id: 0,
            status: 'ON',
          });

          return result.items;
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'orgId',
      label: $t('page.user.form.org'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        treeNodeFilterProp: 'label',
        api: async () => {
          const result = await orgStore.listOrganization(undefined, {
            status: 'ON',
          });
          orgList.value = result.items ?? [];
          return result.items;
        },
        onChange: async (orgId: any) => {
          console.log('org onChange:', orgId);

          if (!orgId) {
            await baseFormApi.setValues(
              {
                orgId: undefined,
                departmentId: undefined,
                positionId: undefined,
              },
              false,
            );
          }
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'departmentId',
      label: $t('page.user.form.department'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        treeNodeFilterProp: 'label',
        api: async () => {
          const values = await baseFormApi.getValues();
          // console.log('values', values);

          const result = await deptStore.listDepartment(undefined, {
            status: 'ON',
            organizationId: values.orgId,
          });
          deptList.value = result.items ?? [];
          return result.items;
        },
        onChange: async (deptId: any) => {
          // console.log('department onChange:', deptId);

          if (!deptId) {
            await baseFormApi.setValues(
              {
                positionId: undefined,
              },
              false,
            );
          }

          const selectedDept = findDepartment(deptList.value, deptId);
          console.log('selectedDept:', selectedDept);
          if (selectedDept) {
            await baseFormApi.setValues(
              {
                orgId: selectedDept.organizationId || undefined,
                positionId: undefined,
              },
              false,
            );
          }
        },
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'positionId',
      label: $t('page.user.form.position'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        numberToString: true,
        showSearch: true,
        treeDefaultExpandAll: true,
        allowClear: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        treeNodeFilterProp: 'label',
        api: async () => {
          const result = await positionStore.listPosition(undefined, {
            status: 'ON',
          });
          positionList.value = result.items ?? [];
          return result.items;
        },
        onChange: async (positionId: any) => {
          console.log('position onChange:', positionId);

          if (!positionId) {
            await baseFormApi.setValues(
              {
                orgId: undefined,
                departmentId: undefined,
              },
              false,
            );
          }

          const selectedPosition = findPosition(positionList.value, positionId);
          console.log('selectedPosition:', selectedPosition);
          if (selectedPosition) {
            await baseFormApi.setValues(
              {
                orgId: selectedPosition.organizationId || undefined,
                departmentId: selectedPosition.departmentId || undefined,
              },
              false,
            );
          }
        },
      },
    },

    {
      component: 'Select',
      fieldName: 'gender',
      label: $t('page.user.table.gender'),
      defaultValue: 'SECRET',
      componentProps: {
        filterOption: (input: string, option: any) =>
          option.label.toLowerCase().includes(input.toLowerCase()),
        allowClear: true,
        showSearch: true,
        options: genderList,
        placeholder: $t('ui.placeholder.select'),
      },
    },

    {
      component: 'Input',
      fieldName: 'nickname',
      label: $t('page.user.table.nickname'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'realname',
      label: $t('page.user.table.realname'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'Input',
      fieldName: 'email',
      label: $t('page.user.table.email'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'mobile',
      label: $t('page.user.table.mobile'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },

    {
      component: 'RadioGroup',
      fieldName: 'status',
      label: $t('ui.table.status'),
      defaultValue: 'ON',
      rules: 'selectRequired',
      componentProps: {
        optionType: 'button',
        buttonStyle: 'solid',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: statusList,
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

    // 加载条设置为加载状态
    setLoading(true);

    // 获取表单数据
    const values = await baseFormApi.getValues();

    console.log(getTitle.value, Object.keys(values));

    try {
      await (data.value?.create
        ? userStore.createUser(values)
        : userStore.updateUser(data.value.row.id, values));

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
      if (data.value.row !== undefined) {
        if (data.value?.row?.orgId !== undefined) {
          data.value.row.orgId = data.value?.row?.orgId.toString();
        }
        baseFormApi.setValues(data.value?.row);
      }

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
