# 如何搭建前端开发环境

## 安装开发工具

需要安装的软件有：

- [Git](https://git-scm.com/)
- [Visual Studio Code](https://code.visualstudio.com/)
- [WebStorm](https://www.jetbrains.com/webstorm/)
- [Node.js](https://nodejs.org/)
- [npm](https://www.npmjs.com/)
- [pnpm](https://pnpm.io/)

### Windows

Windows下安装软件的方法有很多种，这里推荐使用软件包管理工具：[scoop](https://scoop.sh/)。

```shell
scoop bucket add extras
scoop install git vscode webstorm nodejs pnpm
```

### MacOS

MacOS下安装软件的方法有很多种，这里推荐使用软件包管理工具：[Homebrew](https://brew.sh/)。

```shell
brew install git node pnpm
brew install --cask visual-studio-code webstorm
```

## 安装插件

前端需要的插件主要是Protobuf的插件：

- [ts-proto](https://github.com/stephenh/ts-proto)
- [Dart plugin for protobuf compiler](https://pub.dev/packages/protoc_plugin)

安装方法：

### Dart

```shell
flutter pub global activate protoc_plugin
```

### TypeScript

```shell
npm install -g ts-proto
```

## npm/pnpm/yarn切换源

* 国内镜像

| 提供商  | 搜索地址                   | registry地址                                         |
|------|------------------------|----------------------------------------------------|
| 淘宝   | https://npmmirror.com/ | https://registry.npmmirror.com                     |
| 腾讯云  |                        | http://mirrors.cloud.tencent.com/npm/              |
| 华为云  |                        | https://mirrors.huaweicloud.com/repository/npm     |
| 浙江大学 |                        | http://mirrors.zju.edu.cn/npm/                     |
| 南京邮电 |                        | https://mirrors.njupt.edu.cn/nexus/repository/npm/ |

### npm

```shell
# 查看源
npm get registry
npm config get registry

# 临时修改
npm --registry https://registry.npmmirror.com install any-touch

# 永久修改
npm config set registry https://registry.npmmirror.com

# 还原
npm config set registry https://registry.npmjs.org
```

### nrm

```shell
# 安装 nrm
npm install -g nrm

# 列出当前可用的所有镜像源
nrm ls

# 使用淘宝镜像源
nrm use taobao

# 测试访问速度
nrm test taobao
```

### pnpm

```shell
# 查看源
pnpm get registry
pnpm config get registry

# 临时修改
pnpm --registry https://registry.npmmirror.com install any-touch

# 永久修改
pnpm config set registry https://registry.npmmirror.com

# 还原
pnpm config set registry https://registry.npmjs.org
```

### yarn

```shell
# 查看源
yarn config get registry

# 临时修改
yarn add any-touch@latest --registry=https://registry.npmjs.org/

# 永久修改
yarn config set registry https://registry.npmmirror.com/

# 还原
yarn config set registry https://registry.yarnpkg.com
```

### yrm

```shell
# 安装 yrm
npm install -g yrm

# 列出当前可用的所有镜像源
yrm ls

# 使用淘宝镜像源
yrm use taobao

# 测试访问速度
yrm test taobao
```
