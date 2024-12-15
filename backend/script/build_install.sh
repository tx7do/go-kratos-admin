#!/usr/bin/env bash

cd ..
make build

mkdir -p ~/app/kratos_admin

mkdir -p ~/app/kratos_admin/admin/service/bin/

mkdir -p ~/app/kratos_admin/admin/service/configs/

mv -f ./app/admin/service/bin/server ~/app/kratos_admin/admin/service/bin/server

cp -rf ./app/admin/service/configs/*.yaml ~/app/kratos_admin/admin/service/configs/

cd ~/app/kratos_admin/admin/service/bin/
pm2 start --namespace kratos_admin --name admin server -- -conf ../configs/

pm2 save

pm2 restart kratos_admin
