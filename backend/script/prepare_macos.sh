#!/usr/bin/env bash
set -euo pipefail

####################################
## macOS 准备脚本
## 保存为: script/prepare_macos.sh
####################################

# 检查并安装 Homebrew
if ! command -v brew >/dev/null 2>&1; then
  echo "Homebrew 未检测到，开始安装 Homebrew..."
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  eval "$(/opt/homebrew/bin/brew shellenv || true)"
fi

echo "更新 Homebrew..."
brew update

####################################
## 安装工具软件
####################################
echo "安装常用工具..."
brew install htop wget git jq unzip || true

####################################
## 安装 Node.js 和 pm2
####################################
echo "安装 Node.js..."
brew install node

echo "Node 版本："
node -v || true
echo "npm 版本："
npm -v || true

echo "全局安装 pm2..."
npm install -g pm2

# 配置 pm2 在 macOS 使用 launchd 开机启动
BREW_PREFIX="$(brew --prefix)"
# 通过 env 保证 PATH，适配 Intel/Apple Silicon
echo "配置 pm2 开机启动（launchd）..."
sudo env PATH="$PATH:${BREW_PREFIX}/bin" pm2 startup launchd -u "$USER" --hp "$HOME" || true
pm2 save || true

####################################
## 安装 Docker Desktop
####################################
echo "安装 Docker Desktop..."
brew install --cask docker || true

echo "尝试启动 Docker Desktop（可能需要手动允许权限）..."
open -a Docker || true

# 等待 Docker 启动（最多等待 120 秒）
echo "等待 Docker 完全启动..."
n=0
until docker info >/dev/null 2>&1; do
  n=$((n+1))
  if [ "$n" -ge 24 ]; then
    echo "等待 Docker 启动超时（120s）。请手动打开 Docker Desktop 并登录后重试。"
    break
  fi
  sleep 5
done
echo "Docker 状态检查完成。"

####################################
## 安装 Golang
####################################
echo "尝试运行本仓库的 install_golang.sh..."
if [ -x "./install_golang.sh" ]; then
  ./install_golang.sh || { echo "install_golang.sh 执行失败，尝试通过 brew 安装 go"; brew install go; }
else
  echo "未找到可执行的 install_golang.sh，使用 brew 安装 go"
  brew install go
fi

echo "所有步骤完成。建议重启终端以加载可能的环境变量更改。"
