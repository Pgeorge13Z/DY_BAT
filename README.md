# runedance_douyin

第五届青训营-抖音项目

## 项目结构

### 目录结构

暂时如下：
```go
.
|__idl
| |__user.thrift
|__cmd
| |__user
| | |__biz
| | |__hertz_gen
| | |__hertz_client
|__kitex_gen
| |__user
|__middleware
|__sql

```

### 说明

cmd里放不同的微服务，比如我做user相关的接口，把user相关的微服务生成后的代码放入cmd的user中

kitex生成的代码都集中放到kitex_gen目录下

idl放thrift文件

sql放sql相关文件

middleware放中间件

## 编写微服务的几个步骤

> 从编写idl到启动微服务

*kitex生成代码命令*

```bash
kitex -module example -service example echo.thrift
比如我的user服务，kitex -module DY_BAT -service user ../idl/user.thrift 
```

hz生成server命令

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
3. 把生成的kitex_gen目录里的内容**移动**到最外层的kitex_gen目录中，其他内容保留在当前目录中即可。
4. cmd/xx目录中的main.go就是我们的服务端启动的入口。(如果是goland编译器，可以直接点三角号)
5. 实现业务逻辑。

