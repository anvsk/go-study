#### redis

* 非阻塞io
建立连接，接收请求的socket可以设置nonblock，现代很多操作系统都支持
* io多路复用(一个线程处理多个io流)
得益于事件通知，epoll,select,poll区别
poll也是轮训、主要是拷贝到内核态，然后没有最大连接限制（连表存储）


##### 列表
redis4.0 后统一使用快速列表（双向链表，元素为压缩列表）
type quickNode struct{
    prevNode
    nextNode
    *ziplist
}

* set 
分为两种intset和hash表
intset仅用作全是整数的情况，内部会做一个排序，便于快速的二分查找

* zset
ziplist 和 字典+跳跃表(层级索引)
https://laravelacademy.org/post/22214

* dict
压缩列表和hash，数量超过一定阀值用hash

* 主从同步



#### mysql

* 引擎区别
https://www.runoob.com/w3cnote/mysql-different-nnodb-myisam.html

* 二叉树、B树、B+树
https://laravelacademy.org/post/21942
B树也叫多路树，每个节点可以有多个子节点【二叉树只有两个】
B+树非叶子结点不存数据，只有叶子节点存，进一步降低树的高度
非叶子节点只有导航信息，数据更加紧凑，单位数据
todo 红黑树

* 执行顺序
from joinon where group having  select order

* 事物隔离级别
读一提交，读未提交，可重复读，序列化
表格化
从左到右title 脏读 、不可重复读、幻读

* buffer pool
 

* undo redo日志



