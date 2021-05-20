# go-api


### api服务端
- com           #公用代码块
- enum          #枚举值、缓存key、常量定义
- internal      #按不同服务区分，路由|端口隔离，按需加载
  - admin       #后台模块
  - shop        #前台商城模块
    - goods     #类似controller、核心逻辑调用service里面代码或后期rpc服务
    - router    #路由、中间件绑定定义
- middleware    #自定义中间件
  - jwt
  - ratelimit
  - format


```sh
# 热启动
air -c config.yml

# 登录
curl --location --request POST 'localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"admin",
    "password":"admin"
}'

# 用户信息
curl --location --request GET 'localhost:8080/user/info' \
--header '{header}'

# 商品列表
curl --location --request GET 'localhost:8080/goods' \
--header '{header}' \

```
