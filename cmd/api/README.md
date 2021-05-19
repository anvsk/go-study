# go-api


### api服务端
- common        #公用代码块
- internal      #按不同服务区分，路由|端口隔离，按需加载
  - admin       #后台模块
  - shop        #前台商城模块
    - goods     #类似controller、核心逻辑调用service里面代码或后期rpc服务
    - router    #路由、中间件绑定定义
- middleware    #自定义中间件
  - jwt
  - ratelimit
  - exectime
  - format


```sh
# 热启动
air -c config.yml

```

