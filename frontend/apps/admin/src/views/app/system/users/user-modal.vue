<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm, z } from '#/adapter/form';
import { authorityList, useOrganizationStore, useUserStore } from '#/store';

const userStore = useUserStore();
const orgStore = useOrganizationStore();

const data = ref();

const getTitle = computed(() => (data.value?.create ? '创建账号' : '编辑账号'));
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
      fieldName: 'userName',
      label: '用户名',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: z.string().min(1, { message: $t('ui.formRules.required') }),
    },
    // {
    //   component: 'VbenInputPassword',
    //   fieldName: 'password',
    //   label: '密码',
    //   componentProps: {
    //     passwordStrength: true,
    //     placeholder: $t('ui.placeholder.input'),
    //   },
    //   rules: 'required',
    // },
    {
      component: 'Select',
      fieldName: 'authority',
      label: '权限',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        options: authorityList,
      },
      rules: 'selectRequired',
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'orgId',
      label: '所属部门',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        api: async () => {
          const result = await orgStore.listOrganization(
            null,
            null,
            null,
            null,
            null,
            true,
          );

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
      rules: 'selectRequired',
    },
    {
      component: 'Input',
      fieldName: 'nickName',
      label: '昵称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'email',
      label: '邮箱',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },

    {
      component: 'Textarea',
      fieldName: 'remark',
      label: '备注',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
  ],
});

const [Modal, modalApi] = useVbenModal({
  onCancel() {
    modalApi.close();
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
        message: `${getTitle.value}成功`,
      });
    } catch {
      notification.success({
        message: `${getTitle.value}失败`,
      });
    } finally {
      // 关闭窗口
      modalApi.close();
      setLoading(false);
    }
  },

  onOpenChange(isOpen: boolean) {
    if (isOpen) {
      // 获取传入的数据
      data.value = modalApi.getData<Record<string, any>>();

      // 为表单赋值
      if (data.value.row !== undefined) {
        data.value.row.orgId = data.value?.row?.orgId.toString();
        baseFormApi.setValues(data.value?.row);
      }

      setLoading(false);

      console.log('onOpenChange', data.value, data.value?.create);
    }
  },
});

function setLoading(loading: boolean) {
  modalApi.setState({ confirmLoading: loading });
}
</script>

<template>
  <Modal :title="getTitle">
    <BaseForm />
  </Modal>
</template>
