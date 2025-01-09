<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { defDepartmentService, makeUpdateMask, statusList } from '#/rpc';

const data = ref();

const getTitle = computed(() => (data.value?.create ? '创建部门' : '编辑部门'));
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
      label: '部门名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'TreeSelect',
      fieldName: 'parentId',
      label: '上级部门',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
      },
      rules: 'selectRequired',
    },
    {
      component: 'InputNumber',
      fieldName: 'sortId',
      label: '排序',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'RadioGroup',
      fieldName: 'status',
      defaultValue: 'ON',
      label: '状态',
      rules: 'selectRequired',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: statusList,
      },
    },
    {
      component: 'Textarea',
      fieldName: 'remark',
      label: '备注',
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

    setLoading(true);

    // 获取表单数据
    const values = await baseFormApi.getValues();

    console.log(getTitle.value, values);

    try {
      await (data.value?.create
        ? defDepartmentService.CreateDepartment({
            data: {
              ...values,
              children: [],
            },
          })
        : defDepartmentService.UpdateDepartment({
            data: {
              id: data.value.row.id,
              ...values,
              children: [],
            },
            updateMask: makeUpdateMask(Object.keys(values)),
          }));

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
      baseFormApi.setValues(data.value?.row);

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
