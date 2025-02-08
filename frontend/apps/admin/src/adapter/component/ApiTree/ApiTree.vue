<script setup lang="ts">
import type { TreeProps } from 'ant-design-vue';

import { computed, h, ref, unref, watch } from 'vue';

import { LucideEllipsisVertical } from '@vben/icons';
import { $t } from '@vben/locales';
import { get, isEqual, isFunction } from '@vben-core/shared/utils';

import { objectOmit } from '@vueuse/core';

import {
  type MenuInfo,
  type OptionsItem,
  type Props,
  ToolbarEnum,
  type TreeEmits,
} from './types';

const props = withDefaults(defineProps<Props>(), {
  title: '',
  toolbar: true,
  checkable: true,
  search: true,
  searchText: '',

  labelField: 'label',
  valueField: 'value',
  childrenField: '',
  resultField: '',
  numberToString: false,
  params: () => ({}),
  immediate: true,
  alwaysLoad: false,
  loadingSlot: '',
  beforeFetch: undefined,
  afterFetch: undefined,
  modelPropName: 'modelValue',
  api: undefined,
  options: () => [],
});

const emit = defineEmits<TreeEmits>();

const refOptions = ref<OptionsItem[]>([]);

// 首次是否加载过了
const isFirstLoaded = ref(false);
const loading = ref(false);

const treeData = ref<TreeProps['treeData']>([]);
const expandedKeys = ref<string[]>();
const selectedKeys = ref<string[]>();
const checkedKeys = ref<string[]>();
const checkStrictly = ref(false);

const toolbarList = computed(() => {
  const { checkable } = props;
  const defaultToolbarList = [
    { label: $t('ui.tree.expand_all'), value: ToolbarEnum.EXPAND_ALL },
    {
      label: $t('ui.tree.collapse_all'),
      value: ToolbarEnum.UN_EXPAND_ALL,
      divider: checkable,
    },
  ];

  return checkable
    ? [
        { label: $t('ui.tree.select_all'), value: ToolbarEnum.SELECT_ALL },
        {
          label: $t('ui.tree.unselect_all'),
          value: ToolbarEnum.UN_SELECT_ALL,
          divider: checkable,
        },
        ...defaultToolbarList,
        {
          label: $t('ui.tree.hierarchical_association'),
          value: ToolbarEnum.CHECK_STRICTLY,
        },
        {
          label: $t('ui.tree.hierarchical_independence'),
          value: ToolbarEnum.CHECK_UN_STRICTLY,
        },
      ]
    : defaultToolbarList;
});

function handleMenuClick(e: MenuInfo) {
  const { key } = e;
  switch (key) {
    case ToolbarEnum.CHECK_STRICTLY: {
      checkStrictly.value = false;
      break;
    }
    case ToolbarEnum.CHECK_UN_STRICTLY: {
      checkStrictly.value = true;
      break;
    }
    case ToolbarEnum.EXPAND_ALL: {
      expandAll();
      break;
    }
    case ToolbarEnum.SELECT_ALL: {
      checkAll();
      break;
    }
    case ToolbarEnum.UN_EXPAND_ALL: {
      collapseAll();
      break;
    }
    case ToolbarEnum.UN_SELECT_ALL: {
      uncheckAll();
      break;
    }
  }
}

const getOptions = computed(() => {
  const { labelField, valueField, childrenField, numberToString } = props;

  const refOptionsData = unref(refOptions);

  function transformData(data: OptionsItem[]): OptionsItem[] {
    return data.map((item) => {
      const value = get(item, valueField);
      return {
        ...objectOmit(item, [labelField, valueField, childrenField]),
        title: get(item, labelField),
        key: numberToString ? `${value}` : value,
        ...(childrenField && item[childrenField]
          ? { children: transformData(item[childrenField]) }
          : {}),
      };
    });
  }

  const data: OptionsItem[] = transformData(refOptionsData);

  return data.length > 0 ? data : props.options;
});

function emitChange() {
  treeData.value = unref(getOptions) as TreeProps['treeData'];
  emit('optionsChange', unref(getOptions));
}

async function fetchApi() {
  let { api, beforeFetch, afterFetch, params, resultField } = props;

  if (!api || !isFunction(api) || loading.value) {
    return;
  }

  refOptions.value = [];
  try {
    loading.value = true;

    if (beforeFetch && isFunction(beforeFetch)) {
      params = (await beforeFetch(params)) || params;
    }

    let res = await api(params);

    if (afterFetch && isFunction(afterFetch)) {
      res = (await afterFetch(res)) || res;
    }

    isFirstLoaded.value = true;

    if (Array.isArray(res)) {
      refOptions.value = res;
      emitChange();
      return;
    }

    if (resultField) {
      refOptions.value = get(res, resultField) || [];
    }
    emitChange();
  } catch (error) {
    console.warn(error);
    // reset status
    isFirstLoaded.value = false;
  } finally {
    loading.value = false;
  }
}

async function handleFetchForVisible(visible: boolean) {
  if (visible) {
    if (props.alwaysLoad) {
      await fetchApi();
    } else if (!props.immediate && !unref(isFirstLoaded)) {
      await fetchApi();
    }
  }
}

function getAllKeys() {
  const keys: string[] = [];
  function getKeys(data: OptionsItem[]) {
    data.forEach((item) => {
      keys.push(item.key);
      if (item.children) {
        getKeys(item.children);
      }
    });
  }
  getKeys(unref(getOptions));
  return keys;
}

/**
 * 展开所有节点
 */
function expandAll() {
  expandedKeys.value = getAllKeys();
}

/**
 * 收起所有节点
 */
function collapseAll() {
  expandedKeys.value = [];
}

/**
 * 全选
 */
function checkAll() {
  checkedKeys.value = getAllKeys();
}

/**
 * 全不选
 */
function uncheckAll() {
  checkedKeys.value = [];
}

watch(
  () => props.params,
  (value, oldValue) => {
    if (isEqual(value, oldValue)) {
      return;
    }
    fetchApi();
  },
  { deep: true, immediate: props.immediate },
);

handleFetchForVisible(true);
</script>

<template>
  <a-space direction="vertical">
    <a-space>
      <div>{{ props.title }}</div>
      <a-dropdown>
        <a-button type="link" :icon="h(LucideEllipsisVertical)" />
        <template #overlay>
          <a-menu @click="handleMenuClick">
            <template v-for="item in toolbarList" :key="item.value">
              <a-menu-item v-bind="{ key: item.value }">
                {{ item.label }}
              </a-menu-item>
              <a-menu-divider v-if="item.divider" />
            </template>
          </a-menu>
        </template>
      </a-dropdown>
    </a-space>
    <a-tree
      checkable
      v-model:expanded-keys="expandedKeys"
      v-model:selected-keys="selectedKeys"
      v-model:checked-keys="checkedKeys"
      :tree-data="treeData"
      :check-strictly="checkStrictly"
    />
  </a-space>
</template>

<style scoped></style>
