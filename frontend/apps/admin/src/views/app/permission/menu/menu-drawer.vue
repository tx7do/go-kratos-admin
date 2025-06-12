<script lang="ts" setup>
import { computed, ref } from 'vue';

import { useVbenDrawer } from '@vben/common-ui';
import { $t } from '@vben/locales';

import { MenuType } from '#/generated/api/admin/service/v1/i_menu.pb';
import lucide from '@iconify/json/json/lucide.json';
import { addCollection } from '@iconify/vue';
import { notification } from 'ant-design-vue';

import { useVbenForm } from '#/adapter/form';
import {
  buildMenuTree,
  isButton,
  isFolder,
  isMenu,
  menuTypeList,
  statusList,
  useMenuStore,
} from '#/stores';

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
      label: '菜单类型',
      componentProps: {
        optionType: 'button',
        class: 'flex flex-wrap', // 如果选项过多，可以添加class来自动折叠
        options: menuTypeList,
      },
      defaultValue: MenuType.FOLDER,
      rules: 'selectRequired',
    },

    {
      component: 'Input',
      fieldName: 'meta.title',
      label: '菜单名称',
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'ApiTreeSelect',
      fieldName: 'parentId',
      label: '上级菜单',
      componentProps: {
        placeholder: $t('ui.placeholder.select'),
        api: async () => {
          const fieldValue = baseFormApi.form.values;
          const result = await menuStore.listMenu(true, null, null, {
            parentId: fieldValue.parentId,
            status: 'ON',
          });
          return result.items;
        },
        numberToString: true,
        childrenField: 'children',
        labelField: 'meta.title',
        valueField: 'id',
        afterFetch: (data: any) => {
          return buildMenuTree(data);
        },
      },
    },
    {
      component: 'InputNumber',
      fieldName: 'meta.order',
      label: '排序',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
    },
    {
      component: 'IconPicker',
      fieldName: 'meta.icon',
      label: '图标',
      componentProps: {
        prefix: 'lucide',
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Input',
      fieldName: 'path',
      label: '路由地址',
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Input',
      fieldName: 'component',
      label: '组件路径',
      defaultValue: 'BasicLayout',
      rules: 'required',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      dependencies: {
        show: (values) => isMenu(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Input',
      fieldName: 'meta.authority',
      label: '权限标识',
      componentProps: {
        placeholder: $t('ui.placeholder.input'),
        allowClear: true,
      },
      dependencies: {
        show: (values) => !isFolder(values.type),
        triggerFields: ['type'],
      },
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
      component: 'Switch',
      fieldName: 'meta.affixTab',
      label: '固定标签页',
      componentProps: {
        class: 'w-auto',
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Switch',
      fieldName: 'meta.hideChildrenInMenu',
      label: '子级不展现',
      componentProps: {
        class: 'w-auto',
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Switch',
      fieldName: 'meta.hideInBreadcrumb',
      label: '面包屑中不展现',
      componentProps: {
        class: 'w-auto',
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Switch',
      fieldName: 'meta.hideInMenu',
      label: '菜单中不展现',
      componentProps: {
        class: 'w-auto',
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Switch',
      fieldName: 'meta.hideInTab',
      label: '标签页中不展现',
      componentProps: {
        class: 'w-auto',
      },
      dependencies: {
        show: (values) => !isButton(values.type),
        triggerFields: ['type'],
      },
    },
    {
      component: 'Switch',
      fieldName: 'meta.keepAlive',
      label: '是否缓存',
      componentProps: {
        class: 'w-auto',
      },
      dependencies: {
        show: (values) => isMenu(values.type),
        triggerFields: ['type'],
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
      if (values.meta.authority) {
        values.meta.authority = values.meta.authority.split(',');
      }

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

      if (data.value?.row?.meta && data.value?.row?.meta?.authority) {
        const authority = data.value.row.meta.authority;
        data.value.row.meta.authority = authority.join(',');
      }

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
