<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import lucide from '@iconify/json/json/lucide.json';
import { addCollection } from '@iconify/vue';
import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  defMenuService,
  makeUpdateMask,
  menuTypeList,
  statusList,
} from '#/rpc';

const data = ref();
addCollection(lucide);

const getTitle = computed(() => (data.value?.create ? '创建菜单' : '编辑菜单'));
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
      component: 'RadioGroup',
      fieldName: 'type',
      label: '菜单类型',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: menuTypeList,
      },
      rules: 'selectRequired',
    },

    {
      component: 'Input',
      fieldName: 'name',
      label: '菜单名称',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'TreeSelect',
      fieldName: 'parentId',
      label: '上级菜单',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'orderNo',
      label: '排序',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'IconPicker',
      fieldName: 'icon',
      label: '图标',
      componentProps: {
        prefix: 'lucide',
      },
    },
    {
      component: 'Input',
      fieldName: 'path',
      label: '路由地址',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'component',
      label: '组件路径',
      defaultValue: 'BasicLayout',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    // {
    //   component: 'Input',
    //   fieldName: 'permissionCode',
    //   label: '权限标识',
    //   componentProps: {
    //     placeholder: $t('ui.placeholder.input'),
    //     allowClear: true,
    //   },
    // },
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
      component: 'Switch',
      fieldName: 'isExt',
      label: '是否外链',
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'keepAlive',
      label: '是否缓存',
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'show',
      label: '是否显示',
      componentProps: {
        class: 'w-auto',
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
        ? defMenuService.CreateMenu({
            menu: {
              ...values,
              children: [],
            },
          })
        : defMenuService.UpdateMenu({
            menu: {
              id: data.value.row.id,
              children: [],
              ...values,
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
