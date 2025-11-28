#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# 可sudo执行命令
SUDO=${SUDO:-}
if [ "$EUID" -ne 0 ]; then
  SUDO='sudo'
fi

# 目标用户与主目录
TARGET_USER=${SUDO_USER:-$(whoami)}
TARGET_HOME=$(eval echo "~${TARGET_USER}")

log() { printf '==> %s\n' "$*"; }
err_trap() { printf '错误：第 %s 行发生错误\n' "$1" >&2; exit 1; }
trap 'err_trap $LINENO' ERR

# 选择包管理器（centos7 可能只有 yum，centos8+ 有 dnf）
if command -v dnf >/dev/null 2>&1; then
  PKG=dnf
else
  PKG=yum
fi

log "使用包管理器：${PKG}"
export LC_ALL=C

log "更新系统包索引并升级"
if [ "${PKG}" = "dnf" ]; then
  ${SUDO} ${PKG} -y upgrade --refresh || true
else
  ${SUDO} ${PKG} -y update || true
  ${SUDO} ${PKG} -y upgrade || true
fi

log "安装基础工具与 EPEL"
${SUDO} ${PKG} -y install epel-release || true
${SUDO} ${PKG} -y install -y htop wget unzip git jq curl gnupg2 yum-utils || true

log "安装 Node.js (NodeSource 18.x) 并安装 pm2"
curl -fsSL https://rpm.nodesource.com/setup_18.x | ${SUDO} bash - || true
${SUDO} ${PKG} -y install nodejs || true

# 全局安装 pm2（以 root 安装以保证 system-wide 可用）
${SUDO} npm install -g pm2@latest --no-fund --no-audit || true

# 生成并启用 pm2 的 systemd 单元为目标用户
log "为用户 ${TARGET_USER} 生成 pm2 systemd 单元并尝试启用"
# 使用 env 明确 HOME 与 PM2_HOME
${SUDO} env HOME="${TARGET_HOME}" USER="${TARGET_USER}" PM2_HOME="${TARGET_HOME}/.pm2" PATH="/usr/bin:${PATH}" pm2 startup systemd -u "${TARGET_USER}" --hp "${TARGET_HOME}" || true
${SUDO} systemctl daemon-reload || true
# 尝试启用可能的单元名
${SUDO} systemctl enable --now "pm2-${TARGET_USER}.service" 2>/dev/null || ${SUDO} systemctl enable --now "pm2-${TARGET_USER}" 2>/dev/null || true

# 尝试安装 pm2 tab 补全（为目标用户）
${SUDO} env HOME="${TARGET_HOME}" USER="${TARGET_USER}" PATH="/usr/bin:${PATH}" bash -lc 'pm2 completion install >/dev/null 2>&1 || true'

log "清理可能残留的旧 Docker 包"
for pkg in docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-engine podman-docker containerd runc; do
  ${SUDO} ${PKG} -y remove $pkg >/dev/null 2>&1 || true
done

log "配置 Docker 官方仓库并安装"
${SUDO} yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo || true
${SUDO} ${PKG} -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin || true

log "启用并启动 Docker，添加用户到 docker 组"
${SUDO} systemctl enable --now docker || true
${SUDO} groupadd -f docker || true
${SUDO} usermod -aG docker "${TARGET_USER}" || true

log "运行项目内的 Golang 安装脚本（如果存在且可执行）"
if [ -x ./install_golang.sh ]; then
  ./install_golang.sh
else
  log "未找到可执行的 ./install_golang.sh，跳过"
fi

log "清理与完成"
${SUDO} ${PKG} -y autoremove || true
${SUDO} ${PKG} clean all || true

log "完成。注意：将用户加入 docker 组后需要重新登录以生效；pm2 的 systemd 单元已尝试为用户 ${TARGET_USER} 启用。"