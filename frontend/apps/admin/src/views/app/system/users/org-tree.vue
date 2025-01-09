<script lang="ts" setup>
import type { TreeProps } from 'ant-design-vue';

import { onMounted, ref } from 'vue';

import { useOrganizationStore } from '#/store';

const emit = defineEmits(['select']);

const orgStore = useOrganizationStore();

const expandedKeys = ref<(number | string)[]>([]);
const searchValue = ref<string>('');
const autoExpandParent = ref<boolean>(true);
const treeData = ref<TreeProps['treeData']>([]);
const dataList: TreeProps['treeData'] = [];

async function fetch() {
  const orgData = await orgStore.listOrganization(
    null,
    null,
    null,
    null,
    null,
    true,
  );

  for (let i = 0; i < orgData.items.length; i++) {
    const node = orgData.items[i];
    if (node === null) continue;

    const key = node?.id ?? '';
    dataList?.push({ key, title: key });
  }
}

const handleExpand = (keys: string[]) => {
  expandedKeys.value = keys;
  autoExpandParent.value = false;
};

function handleSelect(keys: any[]) {
  emit('select', keys[0]);
}

onMounted(() => {
  fetch();
});
</script>

<template>
  <div class="m-4 mr-0 overflow-hidden bg-white">
    <a-input-search
      v-model:value="searchValue"
      style="margin-bottom: 8px"
      placeholder="Search"
    />
    <a-tree
      :expanded-keys="expandedKeys"
      :auto-expand-parent="autoExpandParent"
      :tree-data="treeData"
      @expand="handleExpand"
      @select="handleSelect"
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
