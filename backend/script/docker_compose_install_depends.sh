app_root=/root/app
deps=["postgres", "redis", "consul", "minio", "jaeger"]

# 为三方服务创建文件夹，并赋权。
for dep in ${deps[@]}
do
  echo $dep
  # 创建文件夹
  mkdir -p $app_root/$dep
  # 为了避免权限问题，创建文件夹时使用1001用户
  sudo chown -R 1001:1001 $app_root/$dep/
done

# 进入到后端项目的根目录
cd ../

# 部署
docker-compose -f docker-compose-without-services.yaml up -d --force-recreate
