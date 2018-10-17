package main

import (
	"flag"

	"github.com/duckchat/duckchat-gateway/server"
)

func main() {

	flag.StringVar(
		&server.HttpBackendAddr,
		"httpBackend",
		server.HttpBackendAddr,
		"the http backend server http://ip:port",
	)

	flag.StringVar(
		&server.WebsocketServerAddr,
		"websocket",
		server.WebsocketServerAddr,
		"the websocket server listen address, format: [ip]:port  .  set 0 to disabled.",
	)

	flag.StringVar(
		&server.WebsocketServerSsl,
		"websocketServerSsl",
		server.WebsocketServerSsl,
		"if not empty, websocket-ssl will be enabled. ip:port, ex: 127.0.0.1:2031, :2031",
	)

	flag.StringVar(
		&server.WebsocketServerSslCertFile,
		"websocketSslCertFile",
		server.WebsocketServerSslCertFile,
		"Format: pem",
	)

	flag.StringVar(
		&server.WebsocketServerSslKeyFile,
		"websocketSslKeyFile",
		server.WebsocketServerSslKeyFile,
		"Format: pem",
	)

	flag.StringVar(
		&server.ZalyServerAddr,
		"zaly",
		server.ZalyServerAddr,
		"the http backend server ip:port",
	)

	flag.BoolVar(
		&server.EnableDebug,
		"debug",
		server.EnableDebug,
		"enable the debug mode.",
	)

	flag.Parse()
	server.SetLoggerLevel()

	go server.StartHttpServer()
	go server.StartWebsocketServer()
	go server.StartZalyServer()

	var c chan int
	<-c
}
