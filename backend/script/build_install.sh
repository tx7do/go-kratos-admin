#!/usr/bin/env bash

cd ..

# 编译服务
make build

project_name=go_wind_admin
install_root=~/app/$project_name
project_root=$(cd $(dirname $0);pwd)
app_root=$project_root/app

function list_dir() {
#    echo $1
    for file in $1/*
    do
    if test -d $file
    then
#        echo $file
        arr=(${arr[*]} $file)
    fi
    done
}

list_dir $app_root
#echo ${arr[@]}

# 创建安装文件夹
mkdir -p $install_root

for v in ${arr[@]}
do
		app=${v##*$app_root/}
		echo $app service installing...

		app_install_root=$install_root/$app
		echo $app_install_root

    # 创建二进制存放路径
    mkdir -p $app_install_root/service/bin/
    # 创建配置存放路径
    mkdir -p $app_install_root/service/configs/

    # 安装二进制程序
    mv -f $app_root/$app/service/bin/server $app_install_root/service/bin/server
    # 拷贝配置文件
    cp -rf $app_root/$app/service/configs/*.yaml $app_install_root/service/configs/
done

# 加入PM2监控运行
for v in ${arr[@]}
do
		app=${v##*$app_root/}
		echo $app service starting...

    app_install_root=$install_root/$app
    echo $app_install_root

    # 加入PM2监控运行
    cd $app_install_root/service/bin/
    pm2 start --namespace $project_name --name $app server -- -conf ../configs/
done

pm2 save
pm2 restart $project_name
