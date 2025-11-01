import type {
  ListDictEntryResponse,
  ListDictTypeResponse,
} from '#/generated/api/dict/service/v1/dict.pb';

import { defineStore } from 'pinia';

import { useDictStore } from '#/stores';

const dictStore = useDictStore();

interface DictViewState {
  currentTypeId: null | number; // 当前选中的字典类型ID
  loading: boolean; // 加载状态

  typeList: ListDictTypeResponse; // 字典类型列表
  entryList: ListDictEntryResponse; // 字典条目列表
}

/**
 * 字典视图状态
 */
export const useDictViewStore = defineStore('dict-view', {
  state: (): DictViewState => ({
    currentTypeId: null,
    loading: false,
    typeList: { items: [], total: 0 },
    entryList: { items: [], total: 0 },
  }),

  actions: {
    /**
     * 获取字典类型列表
     */
    async fetchTypeList(
      currentPage: number,
      pageSize: number,
      formValues: any,
    ) {
      this.loading = true;
      try {
        this.typeList = await dictStore.listDictType(
          false,
          currentPage,
          pageSize,
          formValues,
        );
        return this.typeList;
      } catch (error) {
        console.error('获取字典类型失败:', error);
        this.resetTypeList();
      } finally {
        this.loading = false;
      }

      return this.typeList;
    },

    /**
     * 根据字典类型ID获取字典条目列表
     * @param typeId 字典类型ID
     * @param currentPage
     * @param pageSize
     * @param formValues
     */
    async fetchEntryList(
      typeId: null | number,
      currentPage: number,
      pageSize: number,
      formValues: any,
    ) {
      if (!typeId) {
        this.resetEntryList(); // 无字典类型ID时清空子列表
        this.resetEntryList(); // 无字典类型ID时清空子列表
        return this.entryList;
      }

      this.loading = true;
      try {
        this.entryList = await dictStore.listDictEntry(
          false,
          currentPage,
          pageSize,
          {
            ...formValues,
            type_id: typeId.toString(),
          },
        );
      } catch (error) {
        console.error(`获取字典类型[${typeId}]的条目失败:`, error);
        this.resetEntryList();
      } finally {
        this.loading = false;
      }

      return this.entryList;
    },

    /**
     * 点击字典类型时触发：设置当前字典类型ID + 刷新字典条目列表
     * @param typeId 字典类型ID
     */
    async setCurrentTypeId(typeId: number) {
      this.currentTypeId = typeId; // 更新当前选中的字典类型ID
      await this.fetchEntryList(typeId, 0, 10, null); // 联动刷新字典条目
    },

    resetTypeList() {
      this.typeList = { items: [], total: 0 };
    },

    resetEntryList() {
      this.entryList = { items: [], total: 0 };
    },
  },
});
