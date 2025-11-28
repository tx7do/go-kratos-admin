#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# 环境
SUDO=${SUDO:-}
if [ "$EUID" -ne 0 ]; then
  SUDO='sudo'
fi
TARGET_USER=${SUDO_USER:-$(whoami)}
TARGET_HOME=${HOME:-/home/${TARGET_USER}}
export DEBIAN_FRONTEND=noninteractive

log() { echo "==> $*"; }
err_trap() { echo "Error occurred on line $1"; exit 1; }
trap 'err_trap $LINENO' ERR

log "更新软件包索引并升级"
${SUDO} apt-get update -y
${SUDO} apt-get upgrade -y

log "安装常用工具"
${SUDO} apt-get install -y htop wget unzip git jq ca-certificates curl gnupg lsb-release apt-transport-https software-properties-common

log "安装 Node.js (使用 NodeSource)"
# 使用 Node.js 18.x，必要时可改为 20.x 或其他
curl -fsSL https://deb.nodesource.com/setup_18.x | ${SUDO} -E bash -
${SUDO} apt-get install -y nodejs
node -v || true
npm -v || true

log "安装 pm2 并配置开机自启（为正确用户创建 systemd 单元）"
${SUDO} npm install -g pm2
pm2 --version || true
# 为目标用户生成 systemd 启动命令并执行（确保 PATH 可用）
${SUDO} env PATH="${PATH}:/usr/bin" pm2 startup systemd -u "${TARGET_USER}" --hp "${TARGET_HOME}" || true
# systemd 单元名通常为 pm2-TARGET_USER.service
${SUDO} systemctl daemon-reload || true
${SUDO} systemctl enable --now "pm2-${TARGET_USER}" || true

log "清理可能残留的旧 Docker 包"
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do
  ${SUDO} apt-get remove -y $pkg || true
done

log "配置 Docker 官方仓库并安装"
${SUDO} mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | ${SUDO} gpg --dearmor -o /etc/apt/keyrings/docker.gpg
${SUDO} chmod a+r /etc/apt/keyrings/docker.gpg
ARCH=$(dpkg --print-architecture)
CODENAME=$(. /etc/os-release && echo "$VERSION_CODENAME")
echo "deb [arch=${ARCH} signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu ${CODENAME} stable" | ${SUDO} tee /etc/apt/sources.list.d/docker.list > /dev/null

${SUDO} apt-get update -y
${SUDO} apt-get install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

log "将用户加入 docker 组（需要重新登录以生效）"
${SUDO} groupadd -f docker
${SUDO} usermod -aG docker "${TARGET_USER}" || true

log "启用并启动 Docker"
${SUDO} systemctl enable --now docker

log "运行项目内的 Golang 安装脚本（如果存在且可执行）"
if [ -x ./install_golang.sh ]; then
  ./install_golang.sh
else
  log "未找到可执行的 ./install_golang.sh，跳过"
fi

log "清理"
${SUDO} apt-get autoremove -y
${SUDO} apt-get autoclean -y

log "完成。请注意：如果将用户加入 docker 组，需要重新登录；pm2 的 systemd 单元已为用户 ${TARGET_USER} 尝试启用。"
