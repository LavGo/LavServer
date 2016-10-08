# LavServer
A Static File Server base on golang(一个基于Golang的Web静态服务器)

#### Install
首先自行安装Golang环境
```
Go get -u github.com/LavGo/LavServer/
go install -u github.com/LavGo/LavServer/
//Notice:此种方式获取的是默认分支 branch:master tag:go1上的代码

```

```
-----conf
|     |
|    config.ini(配置文件)
|
-----src（代码目录）
|     | 
|     ------logs(日志)
|     |
|     ------model
-----view（静态文件）
|     |
|     -------static
|              |
|              ------html（html文件）
|              |
|              ------css（Css文件）
|              |
|              ------js(Js文件)
|              |
|              ------img（图片文件）
------LavServer.go

```

####Run
```
cd LavServer 目录
go build LavServer.go
go run LavServer.go
```
