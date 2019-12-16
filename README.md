<h1 align="center">Go-Anything</h1>
<p align="center">
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/Author-wuxiaoxiaoshen-green" alt="Author">
    </a>
    <a href="https://github.com/wuxiaoxiaoshen">
        <img src="https://img.shields.io/badge/progressing-5%25-green" alt="Author">
    </a>
</p>


> 万物皆可爬

1. 微信平台
2. 番号查询
3. 火车票查询
4. Bing 每日壁纸
5. 天天基金
6. 定时任务
7. 知乎热榜
8. 节假日官方数据来源


## 技术栈

- mysql
- redis
- docker / docker-compose
- go / iris
- k8s
- etcd
- kafka
- elasticSearch

## CHANGELOG

- 增加节假日官方通知数据来源
- 增加：docker-compose ci
- 增加组件：elk
- githubAction: 自动编译，自动构建镜像推送至 Docker Hub
- k8s 调度脚本：yaml : 使用 kubectl 启动服务
- docker-compose 版本：v2.0:  本地环境一键启动服务