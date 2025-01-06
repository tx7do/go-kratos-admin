#!/usr/bin/env bash

####################################
## 更新软件源和软件
####################################

sudo dnf check-update
sudo dnf -y update

####################################
## 启用 CRB 存储库
####################################

sudo dnf config-manager --set-enabled crb

####################################
## 安装工具软件
####################################

sudo dnf -y install \
    https://dl.fedoraproject.org/pub/epel/epel-release-latest-9.noarch.rpm \
    https://dl.fedoraproject.org/pub/epel/epel-next-release-latest-9.noarch.rpm

sudo dnf -y install epel-release htop wget unzip git jq

####################################
## 安装PM2
####################################

# 安装nodejs和npm
sudo dnf -y install nodejs npm

node -v
npm -v

# 安装pm2
npm install -g pm2
# 查看pm2的版本
pm2 --version
# tab补全
pm2 completion install
# 创建pm2开机启动脚本
pm2 startup
# 设置pm2的开机启动
sudo systemctl enable pm2-${USER}

####################################
## 安装Golang
####################################

./install_golang.sh

####################################
## 安装Docker
####################################

sudo dnf config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo

sudo yum -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

sudo systemctl start docker
