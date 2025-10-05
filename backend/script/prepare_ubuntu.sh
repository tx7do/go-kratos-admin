#!/usr/bin/env bash

####################################
## 更新软件源和软件
####################################

sudo apt update -y && sudo apt upgrade -y

####################################
## 安装工具软件
####################################

sudo apt install -y htop wget unzip git jq

####################################
## 安装PM2
####################################

# 安装nodejs和npm
sudo apt install -y nodejs npm

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
## 安装Docker
####################################

for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done

sudo apt install -y ca-certificates curl gnupg
sudo install -m 0755 -d /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
sudo chmod a+r /etc/apt/keyrings/docker.gpg

echo \
"deb [arch="$(dpkg --print-architecture)" signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
"$(. /etc/os-release && echo "$VERSION_CODENAME")" stable" | \
sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update -y
sudo apt install -y docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
#sudo apt install -y docker-compose

sudo systemctl enable docker
sudo systemctl start docker

####################################
## 安装Golang
####################################

./install_golang.sh
