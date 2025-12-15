#!/usr/bin/env bash
set -euo pipefail
IFS=$'\n\t'

# 可通过环境变量覆盖目标目录和 compose 文件
APP_ROOT=${APP_ROOT:-/root/app}
COMPOSE_FILE=${COMPOSE_FILE:-docker-compose-without-services.yaml}
deps=(postgres redis consul minio jaeger)

# 切换到脚本所在目录的上一级（项目根）
script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$script_dir/.."
cd "$repo_root" || { echo "Failed to cd to repo root: $repo_root" >&2; exit 1; }

echo "Repo root: $repo_root"
echo "App root: $APP_ROOT"
echo "Compose file: $COMPOSE_FILE"

# 检查 docker-compose 文件存在
if [ ! -f "$COMPOSE_FILE" ]; then
  echo "Compose file not found: $COMPOSE_FILE" >&2
  exit 1
fi

# 检查权限 / sudo
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
  target="$APP_ROOT/$dep"
  echo "Preparing $dep at $target"
  mkdir -p "$target"
  if [ "$is_root" = true ]; then
    chown -R 1001:1001 "$target" || true
  elif [ "$has_sudo" = true ]; then
    sudo chown -R 1001:1001 "$target" || true
  else
    echo "Warning: not root and sudo not available; skipping chown for $target" >&2
  fi
done

# 选择 docker compose 命令（支持 v2 插件或旧的 docker-compose）
DOCKER_COMPOSE_CMD=""
if command -v docker >/dev/null 2>&1 && docker compose version >/dev/null 2>&1; then
  DOCKER_COMPOSE_CMD="docker compose"
elif command -v docker-compose >/dev/null 2>&1; then
  DOCKER_COMPOSE_CMD="docker-compose"
else
  echo "Neither 'docker compose' nor 'docker-compose' found. Please install docker compose." >&2
  exit 1
fi

# 部署
echo "Bringing up services with: $DOCKER_COMPOSE_CMD -f $COMPOSE_FILE up -d --force-recreate"
$DOCKER_COMPOSE_CMD -f "$COMPOSE_FILE" up -d --force-recreate