# 开发者文档

### docker

所有的容器启动，显式的使用 docker-compose 编写yaml 文件的形式启动。
好处是：
- 明确具体的执行步骤和命令

参数遇到的：advertise 带这种关键字的参数，指的是建议访问地址，一般指定服务名称和端口

```text
// etcd
      -advertise-client-urls http://etcd1:2379
      -initial-advertise-peer-urls http://etcd1:2380
// kafka
      KAFKA_ADVERTISED_LISTENERS: "PLAINTEXT://kafka:9092"


```