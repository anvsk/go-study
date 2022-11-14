# go-study
#### go拿来练手的好项目、go简易框架、go学习

##### 包含以下技能点DEMO演示、属于
* Go项目规范、目录文件命名等、基础架构搭建[config、store、cache、toollibs]
* Gin高性能WEB框架的使用、包含中间件JWT、路由定义、目录结构规划和封装
* Gorm Go版MYSQL ORM 使用示例包含自动迁徙、关联模型、预加载
> todo
1. cron 完善
2. rpc [etcd+grpc]
3. kafka消息系统

##### <a href="https://github.com/anvsk/go-study/blob/main/Debug.md">工欲善其事必先利其器、配置DEBUG环境、有助于调试和看源码</a>

### 目录说明
- cmd          功能模块，每个模块打包成独立的执行文件，之间毫无关联

    - api      GIN API服务模块[已加入gin、路由、中间件、gorm示例、具体看该目录README.md]
    - rpc      rpc 服务模块
    - cron     定时脚本模块
    - test     用例demo
    - question 面试相关

- pkg          可公用的工具类、与业务毫无关联
    - store    mysql、redis存储相关
    - util     lib函数、httpclient、config配置加载

- service       也相当于internal(internal命名是go标准规范)、存放各个模块逻辑代码

config.yml      yml格式配置文件
air.toml        api服务热编译

> 目前完成ticket、api、cron模块，下一步完成rpc






