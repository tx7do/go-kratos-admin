import type { App } from 'vue';

import {
  Button,
  Card,
  Input,
  Layout,
  Popconfirm,
  Switch,
  Tabs,
  Tag,
  Tree,
} from 'ant-design-vue';

/**
 * 注册全局组件
 * @param app
 */
export function registerGlobComp(app: App) {
  app
    .use(Input)
    .use(Button)
    .use(Layout)
    .use(Card)
    .use(Switch)
    .use(Popconfirm)
    .use(Tag)
    .use(Tabs)
    .use(Tree);
}
