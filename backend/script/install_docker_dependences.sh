# 创建文件夹
mkdir -p /root/app/kafka
mkdir -p /root/app/postgresql
mkdir -p /root/app/redis
mkdir -p /root/app/consul
mkdir -p /root/app/doris

# 赋权
sudo chown -R 1001:1001 /root/app/kafka/
sudo chown -R 1001:1001 /root/app/postgresql/
sudo chown -R 1001:1001 /root/app/redis/
sudo chown -R 1001:1001 /root/app/consul/
sudo chown -R 1001:1001 /root/app/doris/

# Doris JVM
sudo sysctl -w vm.max_map_count=2000000

# 部署
cd ../.docker/compose/
docker-compose up -d --force-recreate
