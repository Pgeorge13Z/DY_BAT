# runedance_douyin

第五届青训营-抖音项目

## 项目结构

### 目录结构

暂时如下：
```go
.
├── cmd
│   ├── api
│   │   ├── biz
│   │   │   ├── handler
│   │   │   │   ├── api
│   │   │   │   │   └── api_service.go
│   │   │   │   └── ping.go
│   │   │   ├── model
│   │   │   │   └── api
│   │   │   │       └── api.go
│   │   │   ├── router
│   │   │   │   ├── api
│   │   │   │   │   ├── api.go
│   │   │   │   │   └── middleware.go
│   │   │   │   └── register.go
│   │   │   └── rpc
│   │   │       └── user.go
│   │   ├── main.go
│   │   ├── router_gen.go
│   │   └── router.go
│   ├── follow
│   │   ├── build.sh
│   │   ├── dao
│   │   │   ├── follow.go
│   │   │   ├── init.go
│   │   │   └── test.db
│   │   ├── follow.thrift
│   │   ├── handler.go
│   │   ├── kitex_gen
│   │   │   └── follow
│   │   │       ├── follow.go
│   │   │       ├── followservice
│   │   │       │   ├── client.go
│   │   │       │   ├── followservice.go
│   │   │       │   ├── invoker.go
│   │   │       │   └── server.go
│   │   │       ├── k-consts.go
│   │   │       └── k-follow.go
│   │   ├── kitex.yaml
│   │   ├── main.go
│   │   └── script
│   │       └── bootstrap.sh
│   └── user
│       ├── build.sh
│       ├── client
│       │   └── client.go
│       ├── dal
│       │   ├── db_mysql
│       │   │   ├── test
│       │   │   │   └── test.go
│       │   │   ├── user_dao.go
│       │   │   ├── user_model.go
│       │   │   └── user_service.go
│       │   └── init.go
│       ├── handler.go
│       ├── kitex_gen
│       │   └── user
│       │       ├── k-consts.go
│       │       ├── k-user.go
│       │       ├── user.go
│       │       └── userservice
│       │           ├── client.go
│       │           ├── invoker.go
│       │           ├── server.go
│       │           └── userservice.go
│       ├── kitex.yaml
│       ├── main.go
│       ├── output
│       │   ├── bin
│       │   │   └── user
│       │   └── bootstrap.sh
│       └── script
│           └── bootstrap.sh
├── docker-compose.yml
├── go.mod
├── go.sum
├── idl
│   ├── api.thrift
│   ├── follow.thrift
│   └── user.thrift
├── middleware
│   └── middleware.md
├── pkg
│   ├── consts
│   │   └── constants.go
│   ├── errno
│   │   └── errnos.go
│   └── tools
│       ├── encode.go
│       ├── JWT.go
│       └── random_string.go
├── README.md
└── sql
    ├── script
    │   └── db_init.go
    └── sqlcreate
        ├── database.sql
        └── user.sql



```


## 注意
1. Kitex 依赖 Thrift v0.13 ,否则会出现not enough arguments in call to iprot.ReadStructBegin错误

## 编写微服务的几个步骤

> 从编写idl到启动微服务

*kitex生成代码命令*

```bash
kitex -module example -service example echo.thrift
比如我的user服务，kitex -module DY_BAT -service user ../idl/user.thrift 
```

hz生成server命令（hz代码可以后面再写，先写kitex的逻辑代码）

```go
hz new --idl=../idl/psm.thrift --handler_by_method -t=template=slim
比如我的user服务，hz new -module=DY_BAT --idl=../../idl/user.thrift --handler_by_method -t=template=slim
```

hz生成client命令

```go
hz client --idl=../idl/psm.thrift --model_dir=hertz_gen -t=template=slim --client_dir=hz_client
比如我的user服务， hz client -module=DY_BAT --idl=../../idl/user.thrift --model_dir=hertz_gen -t=template=slim --client_dir=hz_client
```



1. 编写idl（idl官方已经给了，见https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#boskXk的三、接口说明)
2. 假设服务是xx，则新建cmd/xx目录，并在目录下调用kitex代码生成工具。（注意版本需要为v0.4.4.）
3. cmd/xx目录中的main.go就是我们的服务端启动的入口。(如果是goland编译器，可以直接点三角号)
4. 在handler.go实现业务逻辑。
5. 运行build.sh会生成output


## 编写http服务的几个步骤
1. 修改idl/api.thrift文件，将自己的结构体和方法写进去

2. ~~进入cmd/api, hz update -idl ../../idl/api.thrift~~   
    测试update命令不成功，我是重新生成然后修改的。

3. 进入cmd/api/biz/rpc,写自己的http请求

4. 进入cmd/api/biz/handler/api,在api_service里更改自己的方法部分

**先在自己的分支上进行update,测试成功后再提交到主分支**