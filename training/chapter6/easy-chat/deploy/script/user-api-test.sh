#! /bin/bash
repo_addr='registry.cn-hangzhou.aliyuncs.com/go-easy-im/user-api-dev'
tag='latest'

container_name="easy-chat-user-api-test"

docker stop ${container_name}
docker rm ${container_name}
docker rmi ${repo_addr}:${tag}
docker pull ${repo_addr}:${tag}

# 如果需要指定配置文件的
# docker run -p 8888:8888 --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 8888:8888 --name=${container_name} -d ${repo_addr}:${tag}