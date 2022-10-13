#### 调试神器

> 一般程度debug用fmt打印即可、深度debug或查看源码执行流程要用debug工具

1. 安装万能编辑器VsCode
2. 安装相关插件 cmd+p 输入 install 找到Go: install/update tools 全部勾选安装
3. 安装debug ```go get -u github.com/go-delve/delve/cmd/dlv```
4. 配置launch.json
   ```json
   {
       "configurations": [
           
       {
           "name": "API",
           "type": "go",
           "request": "launch",
           "mode": "debug",
           "program": "${workspaceFolder}/cmd/api/main.go",
           "args": [
               "-config",
               "../../config.yml"
           ],
           "showLog": true,
       }
       ]
   }
   ```
5. F5开始调试之旅:打印调用栈、结构体变量、方便看源码、调用流程
6. 不适合和Air同时使用


#### Air热编译插件
> air.toml 文件配置编译参数
```sh
go get github.com/cosmtrek/air
air -c air.toml
```

#### 接口调试推荐POSTMAN
> PostMan设置登录后刷新变量token
Tests里面写入
```javascript
postman.setGlobalVariable("gostudy-token", JSON.parse(responseBody).token);
```


#### Test 测试用例
* gotests 生成表格驱动代码块，直接填入测试数据即可
go get -u github.com/cweill/gotests/
gotests -all -w split.go #生成split.go的测试用例
* assert 利用该包assert.Equal() 格式化输出
go get -u github.com/stretchr/testify/assert
* 运行benchmark
go test -bench=.  -run=^# -benchtime 5s 
结果输出
第一列是执行的函数名称，第二列是总的执行次数，第三列是平均运行时间；


    