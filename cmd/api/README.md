# go-api


### api服务端

- internal      #按不同模块区分，路由文件隔离，按需加载
  - admin       #后台模块
  - shop        #前台商城模块
    - goods     #类似controller
    - router    #路由定义
    