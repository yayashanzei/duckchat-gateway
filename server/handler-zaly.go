package server

import (
	"bufio"
	"fmt"
	"net"

	"github.com/duckchat/duckchat-gateway/proto/core"
	"github.com/duckchat/duckchat-gateway/protocol"
	"github.com/golang/protobuf/proto"
)

var connDB map[string]net.Conn = make(map[string]net.Conn)

func StartZalyServer() {
	ln, err := net.Listen("tcp", ZalyServerAddr)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleZalyRequest(conn)
	}
}

func handleZalyRequest(conn net.Conn) {

	var socketId = ""

	defer func() {
		delete(connDB, socketId)
		conn.Close()
	}()

	// parse
	reader := bufio.NewReader(conn)
	for {
		message, _ := protocol.UnpackFromReader(reader)
		commands, _ := message.Strings()
		val := commands[1]
		if len(val) <= 0 {
			fmt.Println("too small values")
			return
		}
		td := &core.TransportData{}
		_ = proto.Unmarshal([]byte(val), td)
		sessionid := td.Header[string(core.TransportDataHeaderKey_HeaderSessionid)]

		socketId = sessionid
		connDB[socketId] = conn

		fmt.Println(sessionid)

		// to http-backend
		postBody, _ := proto.Marshal(td)

		bodyData, keepSocket, _ := requestHttpBackend(td.Action, "pb", postBody)

		// 处理is_api_request
		fmt.Println("keepSocket", keepSocket)

		bodyInZaly, _ := protocol.PackCommand("1", string(bodyData))

		len, err := conn.Write(bodyInZaly)
		fmt.Println(len, err)

		if false == keepSocket {
			return
		}
	}
}
