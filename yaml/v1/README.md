# k8s 部署方法

### namespace

创建新的 namespace, 否则使用的默认的 default。

原因是：进行隔离（逻辑上）

命名规则：小写字母，数字， "-"

### configMap/secret/persistentVolume/persistentVolumeClaim

对项目进行配置文件，比如某些容器启动需要参数，都规则到相应的配置中来。
对项目进行持久化存储约束，比如需要映射文件到本地环境持久化。

### service

优先创建服务和 POD 之间的强连接。

这个时候需要对 POD 的 label 以及暴露的端口明确。依据项目，看是否使用 NodePort 还是 ClusterIP 的形式对外访问。

> 集群内使用这种形式访问 <pod-service-name>.<namespace>.svc.cluster.local
### Preset

预先设置 pod 参数，以注入的形式注入 pod, 精简 pod 配置文件。

预设置的参数有：

- env
- volume

### 控制器

- deployment
- statefulSet
- job
- cronJob
- daemon
- replicaSet

等控制器约束，核心是编写 Yaml 文件，描述对象。

要注意这些字段：labels 和 service 中定义的一致。这样 service 才能和对应的 POD 自动关联起来。



