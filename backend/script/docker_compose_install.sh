#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# 可通过环境变量覆盖，默认 /root/app
APP_ROOT=${APP_ROOT:-/root/app}
deps=(postgres redis consul minio jaeger)

# 切换到脚本所在目录的上一级（项目根）
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$script_dir/.." || { echo "Failed to cd to repo root ($script_dir/..)" >&2; exit 1; }

# 检查 docker-compose 文件
if [ ! -f docker-compose.yml ] && [ ! -f docker-compose.yaml ]; then
  echo "No docker-compose.yml or docker-compose.yaml found in project root." >&2
  exit 1
fi

# 检查是否为 root，或是否有 sudo
is_root=false
if [ "$(id -u)" -eq 0 ]; then
  is_root=true
fi
has_sudo=false
if command -v sudo >/dev/null 2>&1; then
  has_sudo=true
fi

# 创建目录并设置权限（尽量使用 sudo 当非 root 时）
mkdir -p "$APP_ROOT"
for dep in "${deps[@]}"; do
  echo "Preparing $dep..."
  mkdir -p "$APP_ROOT/$dep"
  if [ "$is_root" = true ]; then
    chown -R 1001:1001 "$APP_ROOT/$dep"
  elif [ "$has_sudo" = true ]; then
    sudo chown -R 1001:1001 "$APP_ROOT/$dep"
  else
    echo "Warning: not root and sudo not available; skipping chown for $APP_ROOT/$dep" >&2
  fi
done

# 选择 docker compose 命令（支持 v2 插件或旧的 docker-compose 可执行文件）
DOCKER_COMPOSE_CMD=""
if command -v docker >/dev/null 2>&1 && docker compose version >/dev/null 2>&1; then
  DOCKER_COMPOSE_CMD="docker compose"
elif command -v docker-compose >/dev/null 2>&1; then
  DOCKER_COMPOSE_CMD="docker-compose"
else
  echo "Neither 'docker compose' plugin nor 'docker-compose' found. Please install docker compose." >&2
  exit 1
fi

# 启动服务
echo "Bringing up services with: $DOCKER_COMPOSE_CMD up -d --force-recreate"
$DOCKER_COMPOSE_CMD up -d --force-recreate