package server

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/yayashanzei/duckchat-gateway/proto/gateway"

	"github.com/golang/protobuf/proto"
)

func socketWriteHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	var numMessages int
	var numMessagesWrite int

	defer func() {
		logger.Debugf("/gw/socket/write %d %d", numMessages, numMessagesWrite)
	}()

	gwSocketWriteRequest := &gateway.GwSocketWriteRequest{}
	err := proto.Unmarshal(data, gwSocketWriteRequest)
	if nil != err {
		logger.Warningf("proto_parse_error " + err.Error())
		return
	}

	for _, pack := range gwSocketWriteRequest.Packages {
		for _, socketId := range pack.SocketIds {
			numMessages = numMessages + 1
			conn, _ := connMapGet(socketId)
			if nil == conn {
				continue
			}

			conn.writeMessage(pack.Content)
			numMessagesWrite = numMessagesWrite + 1
		}
	}

	gwSocketWriteResponse := &gateway.GwSocketWriteResponse{}
	gwSocketWriteResponse.Length = 99998888
	d, _ := proto.Marshal(gwSocketWriteResponse)
	w.Write(d)
}

func socketCloseHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)

	gwSocketCloseRequest := &gateway.GwSocketCloseRequest{}
	err := proto.Unmarshal(data, gwSocketCloseRequest)

	if nil != err {
		logger.Warningf("proto_parse_error " + err.Error())
		return
	}

	for _, socketId := range gwSocketCloseRequest.SocketIds {
		conn, _ := connMapGet(socketId)
		if nil != conn {
			conn.close()
		}
	}
}

func StartHttpServer() {

	http.HandleFunc("/gw/socket/write", socketWriteHandler)
	http.HandleFunc("/gw/socket/close", socketCloseHandler)

	var host string
	for {
		httpServerPort = uint16(10000 + rand.Int()%20000)
		host = fmt.Sprintf("%s:%d", HttpServerIp, httpServerPort)
		if nil == http.ListenAndServe(host, nil) {
			logger.Infof("start_inner_http_server_success port: %d", httpServerPort)
			break
		}
	}
}
