# k8s 配置文件


## 配置：ConfigMap

> Redis, MySQL, Kafka, email 公开配置

```text
go-anything-configMap.yaml
```

## 密钥：Secret

> Redis, MySQL, Kafka, email 私有配置

```text
go-anything-secret.yaml
```

## Job

> 服务启动时数据库迁移

```text
go-anything-job.yaml
```

## CronJob

> 定时任务: 
> 每天的邮件 19:00; 
> 每天的基金指数现状 14:45; 
> 健康检查每 5 分钟检测系统运行状态
> Bing 壁纸：先下载，再转存七牛云，最后邮件

```text
go-anything-cronJob.yaml
```

## Deployment

> 组合容器服务

