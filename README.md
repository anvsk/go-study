# go拿来练手的好项目、go简易框架、go学习

#### 自己封装的go项目架构

### 目录说明
- cmd          功能模块，每个模块打包成独立的执行文件，之间毫无关联

    - ticket   DEMO: 抢票脚本
    - api      GIN/ECHO API服务模块
    - rpc      rpc 服务模块
    - cron     定时脚本模块
    - test     测试用例

- pkg          可公用的工具类、与业务毫无关联
    - store    mysql、redis存储相关
    - util     lib函数、httpclient、config配置加载

- service       也相当于internal(internal命名是go标准规范)、存放各个模块逻辑代码

config.yml      yml格式配置文件

> 目前仅完成ticket模块，下一步完成api、rpc






# go-study
