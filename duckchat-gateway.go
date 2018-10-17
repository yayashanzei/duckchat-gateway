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
		"the websocket server listen address [ip]:port",
	)

	flag.StringVar(
		&server.WebsocketServerSslCertFile,
		"websocketSslCertFile",
		server.WebsocketServerSslCertFile,
		"if not empty, websocket-ssl will be enabled. Format: pem",
	)

	flag.StringVar(
		&server.WebsocketServerSslKeyFile,
		"websocketSslKeyFile",
		server.WebsocketServerSslKeyFile,
		"if not empty, websocket-ssl will be enabled. Format: pem",
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
	server.StartZalyServer()
}
