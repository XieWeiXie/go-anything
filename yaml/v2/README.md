# 问题收集

- namespace

命名规则：字母、数字、- 组合

```text
// OK
xw-example
// !OK
xw_example
```

- configMap

key: 基本符合编程语言中变量的命名规则
value: 都处理称字符串，即带上双引号

```text
// OK

mysql.HOST: "192.168.1.1"

// !OK

mysql.PORT: 3306 

error: cannot convert int64 to string

```

- service

ClusterIP: None
nodePort 不可设置

type: NodePort
nodePort: (30000-32767) 端口范围

- deployment

template 一定需要编写 metadata: labels

selector 中 matchLabels 约定了几个字段，template 中的 metadata labels 就需要几个字段

```text
spec:
  selector:
    matchLabels:
      app: xw-example-kafka-3
      role: kafka
  template:
    metadata:
      labels:
        app: xw-example-kafka-3
        role: kafka
``` 

### 安装步骤

- namespace
```text
>> kubectl apply -f 1namespace.yml
>> kubectl get namespace
NAME           STATUS   AGE
default        Active   260d
istio-system   Active   259d
kube-public    Active   260d
kube-system    Active   260d
production     Active   259d
xw             Active   3d
xw-example     Active   20h
```

- configMap

```text
>> kubectl apply -f 2configmap.yml
>> kubectl get configmap --namespace xw-example
NAME                   DATA   AGE
xw-example-configmap   29     20h

>> kubectl describe configmap xw-example-configmap --namespace xw-example
...
```

- service

```text
>> kubectl apply -f 3service.yml
>> kubectl get svc --namespace xw-example
NAME                             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)                                        AGE
xw-example-etcd-1-service        ClusterIP   None            <none>        2379/TCP,2378/TCP                              1h
xw-example-etcd-2-service        ClusterIP   None            <none>        2379/TCP,2378/TCP                              1h
xw-example-etcd-3-service        ClusterIP   None            <none>        2379/TCP,2378/TCP                              1h
xw-example-go-anything-service   NodePort    10.247.82.233   <none>        18089:32088/TCP                                1h
xw-example-kafka-1-service       ClusterIP   None            <none>        9092/TCP                                       1h
xw-example-kafka-2-service       ClusterIP   None            <none>        9092/TCP                                       1h
xw-example-kafka-3-service       ClusterIP   None            <none>        9092/TCP                                       1h
xw-example-mysql-service         ClusterIP   None            <none>        3306/TCP                                       1h
xw-example-redis-service         ClusterIP   None            <none>        6379/TCP                                       1h
xw-example-zookeeper-1-service   NodePort    10.247.73.215   <none>        2181:32105/TCP,2888:30640/TCP,3888:32457/TCP   19s
xw-example-zookeeper-2-service   NodePort    10.247.47.249   <none>        2181:30386/TCP,2888:30748/TCP,3888:30040/TCP   19s
xw-example-zookeeper-3-service   NodePort    10.247.93.179   <none>        2181:32747/TCP,2888:31721/TCP,3888:32731/TCP   19s
```

- deployment

mysql:
```text
>> kubectl apply -f 4mysql.yml
>> kubectl get pods --namespace xw-example

NAME                                           READY   STATUS    RESTARTS   AGE
xw-example-mysql-deployment-697b744678-jf4st   1/1     Running   0          1m
xw-example-mysql-deployment-697b744678-ls2r6   1/1     Running   0          1m

>> kubectl logs -f xw-example-mysql-deployment-697b744678-jf4st --namespace xw-example
>> kubectl describe pod xw-example-mysql-deployment-697b744678-jf4st --namespace xw-example
```

redis:
```text
>> kubectl apply -f 4redis.yml
>> kubectl get pods --namespace xw-example

NAME                                           READY   STATUS    RESTARTS   AGE
xw-example-mysql-deployment-697b744678-jf4st   1/1     Running   0          7m
xw-example-mysql-deployment-697b744678-ls2r6   1/1     Running   0          7m
xw-example-redis-deployment-7cb8bc5c6f-45q4j   1/1     Running   0          1m
xw-example-redis-deployment-7cb8bc5c6f-wv6vw   1/1     Running   0          1m

>> kubectl logs -f xw-example-redis-deployment-7cb8bc5c6f-45q4j --namespace xw-example
>> kubectl describe pod xw-example-redis-deployment-7cb8bc5c6f-45q4j --namespace xw-example
```

