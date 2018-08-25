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
		"the http backend server http://ip:port, default: http://127.0.0.1:8888/",
	)

	flag.StringVar(
		&server.WebsocketServerAddr,
		"websocket",
		server.WebsocketServerAddr,
		"the websocket server listen address [ip]:port, default: 2031:",
	)

	flag.StringVar(
		&server.ZalyServerAddr,
		"zaly",
		server.ZalyServerAddr,
		"the http backend server ip:port, default: :2021",
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
	server.StartWebsocketServer()
	//	server.StartZalyServer()
}
