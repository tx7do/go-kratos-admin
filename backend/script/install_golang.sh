####################################
## 安装Golang
####################################

# 获取当前操作系统和架构
get_os_arch() {
    local os=$(uname | tr '[:upper:]' '[:lower:]')
    local arch="amd64"

    if [ "$os" = "darwin" ]; then
        os="darwin"
    elif [ "$os" = "linux" ]; then
        os="linux"
    else
        echo "Unsupported OS: $os"
        exit 1
    fi

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
