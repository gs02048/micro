
# micro 拷贝自bilibili kratos框架v0.3.3版本

一套Go微服务框架，包含大量微服务相关框架及工具。

## Goals

致力于提供完整的微服务研发体验，整合相关框架及工具后，微服务治理相关部分可对整体业务开发周期无感，从而更加聚焦于业务交付。对每位开发者而言，整套micro框架也是不错的学习仓库，可以了解和参考到在微服务方面的技术积累和经验。

## Features
* HTTP Blademaster：核心基于[gin](https://github.com/gin-gonic/gin)进行模块化设计，简单易用、核心足够轻量；
* GRPC Warden：基于官方gRPC开发，集成服务发现，并融合P2C负载均衡；
* Cache：优雅的接口化设计，非常方便的缓存序列化，推荐结合代理模式
* Database：集成MySQL/HBase/TiDB，添加熔断保护和统计支持，可快速发现数据层压力；
* Config：方便易用的[paladin sdk](doc/wiki-cn/config.md)，可配合远程配置中心，实现配置版本管理和更新；
* Log：类似[zap](https://github.com/uber-go/zap)的field实现高性能日志库，并结合log-agent实现远程日志管理；
* Trace：基于opentracing，集成了全链路trace支持（gRPC/HTTP/MySQL/Redis/Memcached）；
* micro Tool：工具链，可快速生成标准项目，或者通过Protobuf生成代码，非常便捷使用gRPC、HTTP、swagger文档；

## Quick start

### Requirments

Go version>=1.13
安装protoc 3
安装etcd v3

### Installation
```shell
export GOPRIVATE=github.com && export GO111MODULE=on && go get -u github.com/gs02048/micro/tool/micro
cd $GOPATH/src
micro new micro-demo
```

通过 `micro new` 会快速生成基于micro库的脚手架代码，如生成 [micro-demo](https://github.com/gs02048/micro-demo)

### Build & Run

```shell
cd micro-demo/cmd
go build
./cmd -conf ../configs -etcd.endpoints=127.0.0.1:2379
```

打开浏览器访问：[http://localhost:8000/micro-demo/start](http://localhost:8000/micro-demo/start)，你会看到输出了`Golang 大法好 ！！！`

[快速开始](doc/wiki-cn/quickstart.md)  [micro工具](doc/wiki-cn/micro-tool.md)

## Documentation

> [简体中文](doc/wiki-cn/summary.md)  
> [FAQ](doc/wiki-cn/FAQ.md)
