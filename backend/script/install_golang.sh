####################################
## 安装Golang
####################################

# 获取当前操作系统和架构
get_os_arch() {
  local os=$(uname | tr '[:upper:]' '[:lower:]')
  local arch_raw=$(uname -m)
  local arch

  # 处理操作系统类型
  case "$os" in
      darwin|linux|freebsd|aix|dragonfly|illumos|openbsd|netbsd|plan9|solaris)
          # 支持的操作系统，继续处理架构
          ;;
      *)
          echo "Unsupported OS: $os"
          exit 1
          ;;
  esac

  # 处理架构类型（映射为常见的架构名称）
  case "$arch_raw" in
      i386|i486|i586|i686|i786|386)
          arch="386"      # 32 位 x86 架构（老旧 Intel/AMD CPU）
        ;;
      x86_64|amd64)       # 64 位 x86 架构（Intel/AMD 主流 CPU）
          arch="amd64"
          ;;
      aarch64|arm64)
          arch="arm64"    # 64 位 ARM 架构（Apple Silicon、华为鲲鹏等）
          ;;
      armv7l|armv6l|arm)
          arch="arm"      # 32 位 ARM 架构（armv6/armv7，如树莓派早期型号）
          ;;
      s390x)
          arch="s390x"    # 64 位 IBM Z 架构（大型机）
          ;;
      ppc64)
          arch="ppc64"    # 64 位 PowerPC 架构（大端模式）
          ;;
      ppc64le)
          arch="ppc64le"  # 64 位 PowerPC 架构（小端模式）
          ;;
      loongarch64)
          arch="loong64"  # 64 位龙芯架构（国产龙芯 CPU）
          ;;
      riscv64)
          arch="riscv64"  # 64 位 RISC-V 架构（开源指令集，如各类 RISC-V 开发板）
          ;;
      mips64)
          arch="mips64"   # MIPS 64位（大端模式）
          ;;
      mips64el)
          arch="mips64el" # MIPS 64位（小端模式）
          ;;
      mips)
          arch="mips"     # MIPS 32位（大端模式）
          ;;
      mipsel|mipsle)
          arch="mipsle"   # MIPS 32位（小端模式）
          ;;
      *)
          echo "Unsupported architecture: $arch_raw"
          exit 1
          ;;
  esac

  echo "$os $arch"
}

# 安装指定版本的Go
install_go() {
    local version=$1
    local os=$2
    local arch=$3

    echo "Downloading Go $version for $os-$arch..."
    wget -q https://go.dev/dl/$version.$os-$arch.tar.gz

    echo "Removing old Go version..."
    sudo rm -rf /usr/local/go
    echo "Installing..."
    sudo tar -C /usr/local -xzf $version.$os-$arch.tar.gz

    echo "Cleaning up..."
    rm $version.$os-$arch.tar.gz

    echo "Adding Go to PATH..."
    export PATH=$PATH:/usr/local/go/bin
    if [[ "$SHELL" == *"zsh"* ]]; then
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.zshrc
        source ~/.zshrc
    else
        echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
        source ~/.bashrc
    fi

    echo "Go $version installed!"

    go version
}

# 获取本地Go的版本
get_go_local_version() {
    if ! command -v go &> /dev/null; then
        echo "Go is not installed!"
        return
    fi
    echo $(go version | awk '{print $3}')
}

# 获取最新版本的Go
get_go_latest_version() {
    echo $(curl -s https://go.dev/dl/?mode=json | jq -r '.[0].version')
}

GO_LOCAL_VERSION=$(get_go_local_version)
GO_LATEST_VERSION=$(get_go_latest_version)

echo "Current Go version: $GO_LOCAL_VERSION"

read OS ARCH <<< $(get_os_arch)

if [ "$GO_LOCAL_VERSION" = "$GO_LATEST_VERSION" ]; then
    echo "Go is already up to date!"
else
    echo "Installing Go $GO_LATEST_VERSION..."
    install_go $GO_LATEST_VERSION $OS $ARCH
fi
