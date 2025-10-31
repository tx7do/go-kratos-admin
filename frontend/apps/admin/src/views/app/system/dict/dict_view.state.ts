import type {
  ListDictItemResponse,
  ListDictMainResponse,
} from '#/generated/api/dict/service/v1/dict.pb';

import { defineStore } from 'pinia';

import { useDictStore } from '#/stores';

const dictStore = useDictStore();

interface DictViewState {
  currentMainId: null | number; // 当前选中的主字典ID
  loading: boolean; // 加载状态

  mainList: ListDictMainResponse; // 主字典列表
  itemList: ListDictItemResponse; // 子字典列表
}

/**
 * 字典视图状态
 */
export const useDictViewStore = defineStore('dict-view', {
  state: (): DictViewState => ({
    currentMainId: null,
    loading: false,
    mainList: { items: [], total: 0 },
    itemList: { items: [], total: 0 },
  }),

  actions: {
    /**
     * 获取主字典列表
     */
    async fetchMainList(
      currentPage: number,
      pageSize: number,
      formValues: any,
    ) {
      this.loading = true;
      try {
        this.mainList = await dictStore.listDictMain(
          false,
          currentPage,
          pageSize,
          formValues,
        );
        return this.mainList;
      } catch (error) {
        console.error('获取主字典失败:', error);
        this.resetMainList();
      } finally {
        this.loading = false;
      }

      return this.mainList;
    },

    /**
     * 根据主字典ID获取子字典列表
     * @param mainId 主字典ID
     * @param currentPage
     * @param pageSize
     * @param formValues
     */
    async fetchItemList(
      mainId: null | number,
      currentPage: number,
      pageSize: number,
      formValues: any,
    ) {
      if (!mainId) {
        this.resetItemList(); // 无主字典ID时清空子列表
        return this.itemList;
      }

      this.loading = true;
      try {
        this.itemList = await dictStore.listDictItem(
          false,
          currentPage,
          pageSize,
          {
            ...formValues,
            main_id: mainId.toString(),
          },
        );
      } catch (error) {
        console.error(`获取主字典[${mainId}]的子项失败:`, error);
        this.resetItemList();
      } finally {
        this.loading = false;
      }

      return this.itemList;
    },

    /**
     * 点击主字典时触发：设置当前主字典ID + 刷新子字典列表
     * @param mainId 主字典ID
     */
    async setCurrentMain(mainId: number) {
      this.currentMainId = mainId; // 更新当前选中的主字典ID
      await this.fetchItemList(mainId, 0, 10, null); // 联动刷新子字典
    },

    resetMainList() {
      this.mainList = { items: [], total: 0 };
    },

    resetItemList() {
      this.itemList = { items: [], total: 0 };
    },
  },
});
