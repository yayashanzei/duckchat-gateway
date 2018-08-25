# DuckChat-Gateway

> Version: Gaga抢先版
>
> Duckchat主程序下载地址：https://github.com/duckchat/gaga
>
> 官网：http://duck.chat
>
> QQ群：131466613

DuckChat-Gateway是DuckChat的配套软件，用以提供消息长链接服务，支持Websocket、Zaly两种协议。


## 启动方式

### 安装程序

快速启动：`duckchat-gateway -httpBackend=127.0.0.1:8888`

其中，`-httpBackend` 指的是后端gaga程序的地址。

支持的配置参数有：

```bash
  -debug
    	enable the debug mode.
  -httpBackend string
    	the http backend server http://ip:port, default: http://127.0.0.1:8888/ (default "http://127.0.0.1:8888/")
  -websocket string
    	the websocket server listen address [ip]:port, default: 2031: (default ":2031")
  -zaly string
    	the http backend server ip:port, default: :2021 (default ":2021")
```