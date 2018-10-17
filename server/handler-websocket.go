package server

import (
	"crypto/tls"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{}

//
// implement net/http handler
//

type wsGWHandler struct{}

func (h *wsGWHandler) ServeHTTP(rw http.ResponseWriter, request *http.Request) {

	c, err := wsUpgrader.Upgrade(rw, request, nil)
	if err != nil {
		logger.Warningf("websocket upgradeConnError %s", err)
		return
	}

	var gatewayConn = &gatewayConnWebsocket{}
	gatewayConn.ws = c
	gatewayConn.request = request
	gatewayConn.rw = rw

	defer func() {
		gatewayConn.close()
		logger.Debugf("websocket connClosed %s", request.RemoteAddr)
	}()

	handlerInstance.serve(gatewayConn)
}

//
// implement gatewayConnHelper
//

type gatewayConnWebsocket struct {
	gatewayConn

	ws        *websocket.Conn
	rw        http.ResponseWriter
	request   *http.Request
	writeLock sync.Mutex
	readLock  sync.Mutex
}

func (conn *gatewayConnWebsocket) addr() string {
	return conn.ws.LocalAddr().String()
}

func (conn *gatewayConnWebsocket) writeMessage(content []byte) error {
	var writeError error
	conn.writeLock.Lock()
	defer func() {
		conn.writeLock.Unlock()
	}()

	if conn.bodyFormat() == "pb" {
		writeError = conn.ws.WriteMessage(websocket.BinaryMessage, content)
	} else {
		writeError = conn.ws.WriteMessage(websocket.TextMessage, content)
	}
	return writeError
}

func (conn *gatewayConnWebsocket) close() {
	conn.ws.Close()
}

func (conn *gatewayConnWebsocket) remoteAddr() string {
	return conn.request.RemoteAddr
}

func (conn *gatewayConnWebsocket) bodyFormat() string {
	return conn.request.URL.Query().Get("body_format")
}

func (conn *gatewayConnWebsocket) readMessage() ([]byte, error) {
	conn.readLock.Lock()
	defer func() {
		conn.readLock.Unlock()
	}()

	_, message, err := conn.ws.ReadMessage()
	if err != nil {
		return nil, err
	}
	return message, nil
}

//
// start the server
//
func StartWebsocketServer() {
	wsUpgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	handler := &wsGWHandler{}

	if WebsocketServerSsl != "" && WebsocketServerSsl != "0" {

		cert, err := tls.LoadX509KeyPair(WebsocketServerSslCertFile, WebsocketServerSslKeyFile)
		if nil != err {
			logger.Fatalf("start_websocket_ssl_server failed LoadX509KeyPair %s", err)
			return
		}

		host, _, errSplit := net.SplitHostPort(WebsocketServerSsl)
		if nil != err {
			logger.Fatalf("start_websocket_ssl_server failed Invalid WebsocketServerSsl %s", errSplit)
			return
		}

		var tlsConfig = &tls.Config{
			ServerName:   host,
			Certificates: []tls.Certificate{cert},
		}

		l, errListen := tls.Listen("tcp", WebsocketServerSsl, tlsConfig)
		if nil != errListen {
			logger.Fatalf("start_websocket_ssl_server failed  %s", errSplit)
			return
		}

		go http.Serve(l, handler)
		logger.Infof("start_websocket_ssl_server success")
	}

	if WebsocketServerAddr != "0" {

		l, errListen := net.Listen("tcp", WebsocketServerSsl)
		if nil != errListen {
			logger.Fatalf("start_websocket_server failed  %s", errListen)
			return
		}

		go http.Serve(l, handler)
		logger.Infof("start_websocket_server success")
	}
}
