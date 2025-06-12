<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm, z } from '#/adapter/form';
import {
  authorityList,
  genderList, isButton,
  useOrganizationStore,
  useUserStore,
} from '#/stores';

const userStore = useUserStore();
const orgStore = useOrganizationStore();

const data = ref();

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
      label: '密码',
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
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        options: authorityList,
        allowClear: true,
      },
      rules: 'selectRequired',
      dependencies: {
        disabled: () => !data.value?.create,
        triggerFields: ['authority'],
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'orgId',
      label: $t('page.user.table.orgId'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        api: async () => {
          const result = await orgStore.listOrganization(true);

          return result.items;
        },
        numberToString: true,
        childrenField: 'children',
        labelField: 'name',
        valueField: 'id',
        // afterFetch: (data: any) => {
        //   return data.map((item: any) => ({
        //     label: item.name,
        //     value: item.id,
        //   }));
        // },
      },
      // rules: 'selectRequired',
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
      // rules: 'required',
    },

    {
      component: 'Select',
      fieldName: 'gender',
      label: $t('page.user.table.gender'),
      componentProps: {
        options: genderList,
        placeholder: $t('ui.placeholder.select'),
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
