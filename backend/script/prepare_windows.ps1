####################################
## 更新系统和软件
####################################

# 安装 Scoop (如果尚未安装)
if (!(Get-Command scoop -ErrorAction SilentlyContinue)) {
    Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
    Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
}

# 更新 Scoop 包
scoop update *

####################################
## 安装工具软件
####################################

# 添加常用 bucket
scoop bucket add extras
scoop bucket add main

# 安装工具
scoop install wget unzip git jq

scoop install make
scoop install grep gawk sed touch

scoop install mingw

####################################
## 安装PM2
####################################

# 安装 Node.js 和 npm
scoop install nodejs

node -v
npm -v

# 安装 pm2
npm install pm2 -g
# 查看 pm2 的版本
pm2 --version

# tab 补全
pm2 completion install

# 安装服务化工具
npm install pm2-windows-service -g
# 安装PM2服务（会提示环境设置）
pm2-service-install

####################################
## 安装Golang
####################################

# 使用 Scoop 安装 Golang
scoop install go

####################################
## 安装Docker
####################################

# 安装 Docker
scoop install docker

# 启动 Docker 服务 (需要 Docker Desktop 或 WSL2)
Start-Service docker

# 设置 Docker 开机自启动
Set-Service -Name docker -StartupType Automatic
