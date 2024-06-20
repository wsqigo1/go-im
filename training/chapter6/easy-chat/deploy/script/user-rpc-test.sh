#! /bin/bash
repo_addr='registry.cn-hangzhou.aliyuncs.com/go-easy-im/user-rpc-dev'
tag='latest'

pod_ip="192.168.2.6"

container_name="easy-chat-user-rpc-test"

docker stop ${container_name}
docker rm ${container_name}
docker rmi ${repo_addr}:${tag}
docker pull ${repo_addr}:${tag}

# 如果需要指定配置文件的
# docker run -p 10001:8080 --network easy-chat -v /easy-im/config/user-rpc:/user/conf/ --name=${container_name} -d ${reso_addr}:${tag}
docker run -p 10000:10000 -e POD_IP=${pod_ip} --name=${container_name} -d ${repo_addr}:${tag}