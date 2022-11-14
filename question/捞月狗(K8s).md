K8s架构 各个组件功能
https://blog.51cto.com/u_15069485/4518558

master:
apiserver 提供对外接口服务统一入口  
etcd 存储中心
controllers 
schedulers 集群节点的调度

node:
kubelet 接收master下达的指令，如创建（调用container的接口执行），等生命周期的操作
kube-procy 网络层面代理，有外向内有负载均衡转发、由内向外提供服务
runtime

Pod的启动流程
APIServer(收到启动指令后)将pod信息存入etcd，通知Scheduler；
Scheduler根据调度算法，为pod选择一个节点，然后向APIServer发送更新spec.nodeName；
APIServer更新完毕，通知对应节点的kubelet；
kubelet发现pod调度到本节点，创建并运行pod的容器；

说说副本
自动扩容 缩容的原理
HPA horizontal-pod-autoscaler
Kubernetes中的某个Metrics Server持续采集所有Pod副本的指标数据。HAP控制器通过 Metrics Server的API获取这些数据，基于用户定义的扩容规则进行计算，得到目标Pod的副本数量。当目标Pod副本数量与当前副本数量不同时，HPA控制器就向Pod的副本数量控制器(Deployment、RC或ReplicaSet)发起scale操作，调整Pod的副本数量完成扩缩容操作


讲讲探针
https://kubernetes.io/zh-cn/docs/concepts/workloads/pods/pod-lifecycle/#container-probes
官方三种探针、存活、就绪、启动
就绪，readnessprobe 是否准备好
启动，startupprobe 是否启动

容器平滑重启的原理（滚动发布）
https://www.cnblogs.com/scajy/p/16292210.html
优点
用户无感知，平滑过渡
缺点
部署周期长
发布策略较复杂
不易回滚
* 利用多节点交替停启

#### 启动顺序
https://cloud.tencent.com/developer/news/818503
[]spec.containers
通过poststart事件通知
一下链接超级详细讲解了
https://www.quwenqing.com/archives/1995.html
三种方式可以控制顺序
启动命令、poststart、initcontroller


进程网络隔离原理
https://www.cnblogs.com/hanfan/p/16166275.html

A协程panic b协程能recover到吗
不能

指针使用的场景
Forrange遍历切片，改值影响切片吗
struct里面有个map，和int，分别在函数里
传地址和传值 然后改变他们，会发生什么
Gmp
对一个map的k取地址会怎么样
只写一个channel，永远不读，程序会异常
吗
有缓冲和没缓冲有什么区别，讲讲原理

Rabitma怎么保证消息的准确发送和接受，
重复消费的问题怎么解决
Rabit和kafka区别大概说说
Redis都有些什么类型，订阅发布用过没讲讲
数据库都有什么索引，都讲一讲，索引失效
的场，最左匹配
说一说项目里自己处理过的亮点和难点