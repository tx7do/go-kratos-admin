<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import lucide from '@iconify/json/json/lucide.json';
import { addCollection } from '@iconify/vue';
import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import { menuTypeList, statusList, useMenuStore } from '#/store';
import {MenuType} from "#/rpc/api/system/service/v1/menu.pb";

const menuStore = useMenuStore();

addCollection(lucide);

const data = ref();

const getTitle = computed(() =>
  data.value?.create
    ? $t('ui.modal.create', { moduleName: $t('page.menu.moduleName') })
    : $t('ui.modal.update', { moduleName: $t('page.menu.moduleName') }),
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
      component: 'RadioGroup',
      fieldName: 'type',
      label: $t('page.menu.type'),
      defaultValue: MenuType.FOLDER,
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
      label: $t('page.menu.name'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'TreeSelect',
      fieldName: 'parentId',
      label: $t('page.menu.parentId'),
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
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
    },
    {
      component: 'IconPicker',
      fieldName: 'icon',
      label: $t('page.menu.icon'),
      componentProps: {
        prefix: 'lucide',
      },
    },
    {
      component: 'Input',
      fieldName: 'path',
      label: $t('page.menu.path'),
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      rules: 'required',
    },
    {
      component: 'Input',
      fieldName: 'component',
      label: $t('page.menu.component'),
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
      label: $t('ui.table.status'),
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
      label: $t('page.menu.isExt'),
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'keepAlive',
      label: $t('page.menu.keepAlive'),
      componentProps: {
        class: 'w-auto',
      },
    },
    {
      component: 'Switch',
      fieldName: 'show',
      label: $t('page.menu.show'),
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
        ? menuStore.createMenu(values)
        : menuStore.updateMenu(data.value.row.id, values));

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
