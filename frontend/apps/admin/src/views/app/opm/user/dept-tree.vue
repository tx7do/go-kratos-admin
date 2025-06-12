<script lang="ts" setup>
import type { TreeProps } from 'ant-design-vue';

import type { Department } from '#/generated/api/user/service/v1/department.pb';

import { onMounted, ref } from 'vue';

import { LucideEllipsisVertical } from '@vben/icons';
import { $t } from '@vben/locales';
import { VbenDropdownMenu, VbenIconButton } from '@vben-core/shadcn-ui';
import { mapTree } from '@vben-core/shared/utils';

import { TreeActionEnum } from '#/constants/tree';
import { useDepartmentStore } from '#/stores';

const emit = defineEmits(['select']);

// const orgStore = useOrganizationStore();
const deptStore = useDepartmentStore();

const toolbarList = [
  {
    value: TreeActionEnum.EXPAND_ALL,
    label: $t('ui.tree.expand_all'),
    handler: handleMenuExpandAll,
  },
  {
    value: TreeActionEnum.COLLAPSE_ALL,
    label: $t('ui.tree.collapse_all'),
    handler: handleMenuCollapseAll,
  },
];

const expandedKeys = ref<(number | string)[]>([]);
const searchValue = ref<string>('');
const autoExpandParent = ref<boolean>(true);

const treeData = ref<TreeProps['treeData']>([]);

async function fetch() {
  try {
    const response = await deptStore.listDepartment(true);

    const newTree = mapTree(response.items, (node: Department) => ({
      ...node,
      key: `${node.parentId}-${node.id}`,
      title: node.name,
      isLeaf: !node.children || node.children.length === 0,
    }));
    console.log('newTree', newTree);
    treeData.value = newTree ?? [];
  } catch (error) {
    console.error(error);
  }
}

/**
 * 展开所有节点
 * @param data
 */
function handleMenuExpandAll(data: any) {
  console.log('handleMenuExpandAll', data);
}

/**
 * 折叠所有节点
 * @param data
 */
function handleMenuCollapseAll(data: any) {
  console.log('handleMenuCollapseAll', data);
}

/**
 * 展开单个节点
 * @param keys
 */
const handleExpandNode = (keys: string[]) => {
  expandedKeys.value = keys;
  autoExpandParent.value = false;
};

/**
 * 选中单个节点
 * @param keys
 */
function handleSelectNode(keys: any[]) {
  emit('select', keys[0]);
}

onMounted(() => {
  fetch();
});
</script>

<template>
  <div class="dept-container m-4 mb-0 ml-0 mt-0 h-full">
    <div class="dept-header">
      <a-space>
        <a-input-search
          allow-clear
          size="middle"
          v-model:value="searchValue"
          :placeholder="$t('ui.input-search.placeholder')"
        />
        <VbenDropdownMenu :modal="false" :menus="toolbarList">
          <VbenIconButton>
            <LucideEllipsisVertical />
          </VbenIconButton>
        </VbenDropdownMenu>
      </a-space>
    </div>

    <a-tree
      :expanded-keys="expandedKeys"
      :auto-expand-parent="autoExpandParent"
      :tree-data="treeData"
      @expand="handleExpandNode"
      @select="handleSelectNode"
      class="h-full w-full"
    >
      <template #title="{ title }">
        <span v-if="title.indexOf(searchValue) > -1">
          {{ title.substring(0, title.indexOf(searchValue)) }}
          <span style="color: #f50">{{ searchValue }}</span>
          {{ title.substring(title.indexOf(searchValue) + searchValue.length) }}
        </span>
        <span v-else>{{ title }}</span>
      </template>
    </a-tree>
  </div>
</template>

<style lang="less" scoped>
.dept-container {
  display: flex;
  flex-direction: column;
}

.dept-header {
  margin-bottom: 2px;
}
</style>
