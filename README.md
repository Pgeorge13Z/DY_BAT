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
| | |___dal
| | |___kitex_gen
| | |___handler.go main.go ...
|__pkg
| |__consts
| |__errno
| |__tools
|__sql
| |__script
| |__sqlcreate




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

