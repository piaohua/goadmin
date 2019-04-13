# goadmin based on [beego](https://github.com/astaxie/beego)

    Beego + vue.js + AdminLte

## Installation

```sh
$ go get github.com/astaxie/beego

$ ./ctrl build goadmin

$ ./ctrl start goadmin
$ ./ctrl stop goadmin

$ cp config/goadmin.json .
$ cp config/goadmin.sh .
$ pm2 start goadmin.json

```

# Documentation

```
static 静态资源
conf 配置文件
controllers 控制器
data sql文件
logs 日志文件
libs 方法
routers 路由
views 视图目录
tests 测试目录
tmp session
ctrl 控制脚本
main.go 启动文件
```

## TODOs
