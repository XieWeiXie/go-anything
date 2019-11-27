# Docker-Compose 版本

go-anything 系统依赖 mysql, redis, kafka。

思路1： 

所有镜像编写在 docker-compose 文件中，组件之间相互通过自定义的一个网络互联，意味者 go-anything 连接
其他服务只要通过服务名称即可代替 Host。

思路2：

go-anything 系统和底层 mysql, redis, kafka 相互隔离，互联通过显式的 host 连接。



