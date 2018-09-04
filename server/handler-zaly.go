package server

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	"github.com/duckchat/duckchat-gateway/protocol"
)

var connDB map[string]net.Conn = make(map[string]net.Conn)

func StartZalyServer() {
	ln, err := net.Listen("tcp", ZalyServerAddr)
	if err != nil {
		// handle error
		fmt.Println(err)
		return
	} else {
		logger.Infof("start_zaly_server suucess %s", ZalyServerAddr)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		logger.Debugf("new zaly connection")

		var gatewayConn = &gatewayConnZaly{}
		gatewayConn.conn = conn
		gatewayConn.reader = bufio.NewReader(conn)

		defer func() {
			//			gatewayConn.close()
			//			logger.Debugf("zaly connClosed %s", request.RemoteAddr)
		}()

		go handlerInstance.serve(gatewayConn)
	}
}

//
// implement gatewayConnHelper
//

type gatewayConnZaly struct {
	gatewayConn

	conn      net.Conn
	reader    *bufio.Reader
	writeLock sync.Mutex
	readLock  sync.Mutex
}

func (conn *gatewayConnZaly) addr() string {
	return conn.conn.LocalAddr().String()
}

func (conn *gatewayConnZaly) writeMessage(content []byte) error {
	//	var writeError error
	conn.writeLock.Lock()
	defer func() {
		conn.writeLock.Unlock()
	}()

	bodyInZaly, _ := protocol.PackCommand("0.0", string(content))
	conn.conn.Write(bodyInZaly)
	return nil
}

func (conn *gatewayConnZaly) close() {
	conn.conn.Close()
}

func (conn *gatewayConnZaly) remoteAddr() string {
	return conn.conn.RemoteAddr().String()
}

func (conn *gatewayConnZaly) bodyFormat() string {
	return "pb"
}

func (conn *gatewayConnZaly) readMessage() ([]byte, error) {
	conn.readLock.Lock()
	defer func() {
		conn.readLock.Unlock()
	}()

	message, err := protocol.UnpackFromReader(conn.reader)
	if err != nil {
		logger.Warning(message, err)
		return nil, err
	}

	commands, _ := message.Strings()
	val := commands[1]
	if len(val) <= 0 {
		fmt.Println("too small values")
		return nil, err
	}

	return []byte(val), nil
}
