#### 调试神器

1. 安装万能编辑器VsCode
2. 安装相关插件 cmd+p 输入 install 找到Go: install/update tools 全部勾选
3. 安装debug ```go get -u github.com/go-delve/delve/cmd/dlv```
4. 配置launch.json
   ```json
   {
       "configurations": [
           
       {
           "name": "Launch file",
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


#### Air热编译插件
```go get github.com/cosmtrek/air```
air -c air.toml

#### PostMan设置登录后刷新变量token
Tests里面写入
```postman.setGlobalVariable("gostudy-token", JSON.parse(responseBody).token);```


    