<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenModal } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { notification } from 'ant-design-vue';

import { useVbenForm, z } from '#/adapter/form';
import { useDictStore } from '#/store';

const dictStore = useDictStore();

const data = ref();

const getTitle = computed(() => (data.value?.create ? '创建条目' : '编辑条目'));
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
      fieldName: 'key',
      label: '字典键',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: z.string().min(1, { message: $t('ui.formRules.required') }),
    },
    {
      component: 'Input',
      fieldName: 'category',
      label: '字典类型',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'categoryDesc',
      label: '字典类型名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'value',
      label: '字典值',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'valueDesc',
      label: '字典值名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },

    {
      component: 'InputNumber',
      fieldName: 'sortId',
      label: '排序',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
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
        ? dictStore.createDict(values)
        : dictStore.updateDict(data.value.row.id, values));

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
