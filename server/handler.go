package server

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"sync"

	"github.com/yayashanzei/duckchat-gateway/proto/core"

	_ "github.com/yayashanzei/duckchat-gateway/proto/client"
	_ "github.com/yayashanzei/duckchat-gateway/proto/site"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var handlerInstance *serverHandler

// TODO replace with sync.Map
var connMap map[string]gatewayConnHelper
var connMapLock sync.Mutex

func connMapSet(key string, handler gatewayConnHelper) {
	connMapLock.Lock()
	connMap[key] = handler
	connMapLock.Unlock()
}

func connMapDel(key string) {
	connMapLock.Lock()
	delete(connMap, key)
	connMapLock.Unlock()
}

func connMapGet(key string) (gatewayConnHelper, bool) {
	val, ok := connMap[key]
	return val, ok
}

type gatewayConnHelper interface {
	addr() string
	writeMessage(content []byte) error
	readMessage() ([]byte, error)
	close()
	remoteAddr() string
	bodyFormat() string

	setSocketId(socketId string)
	getSocketId() string
}

type gatewayConn struct {
	socketId string
}

func (conn *gatewayConn) setSocketId(socketId string) {
	conn.socketId = socketId
}

func (conn *gatewayConn) getSocketId() string {
	return conn.socketId
}

type serverHandler struct {
	firstMessageTime int64
}

func (app *serverHandler) serve(conn gatewayConnHelper) {
	socketId := app.newSocketId()
	conn.setSocketId(socketId)
	connMapSet(socketId, conn)

	defer func() {
		conn.close()
	}()

	for {
		message, err := conn.readMessage()

		if nil != err {
			conn.close()
			connMapDel(socketId)
			break
		}
		go app.handleMessage(conn, message)
	}
}

func (app *serverHandler) newSocketId() string {
	return strconv.FormatInt(rand.Int63(), 10)
}

// 保留originData字段，防止gateway没有一些proto，导致一些字段无法被序列化为json
func (app *serverHandler) bytesToJson(bodyFormat string, originData []byte, transportData *core.TransportData) string {
	var debugMessage string
	var error error
	if bodyFormat == "json" {
		debugMessage = string(originData)
	} else if bodyFormat == "pb" {
		m := &jsonpb.Marshaler{}
		debugMessage, error = m.MarshalToString(transportData)
		if nil != error {
			logger.Debugf("parse pb error " + error.Error())
		}
	}
	return debugMessage
}
func (app *serverHandler) handleMessage(conn gatewayConnHelper, data []byte) {

	var keepSocket bool = false
	var action string = ""
	var requestSize int = len(data)
	var responseSize int = 0
	var bodyFormat string = ""
	var err error = nil

	defer func() {
		if false == keepSocket {
			conn.close()
			connMapDel(conn.getSocketId())
		}

		var errorMsg = ""
		if nil != err {
			errorMsg = err.Error()
		}

		logger.Infof("action: %s client: %-21s request: %d response: %d keepSocket: %v bodyFormat: %s %s", action, conn.remoteAddr(), requestSize, responseSize, keepSocket, bodyFormat, errorMsg)
	}()

	bodyFormat = conn.bodyFormat()
	transportData := app.getTransportData(bodyFormat, data)
	action = transportData.Action

	if EnableDebug {
		debugMessage := app.bytesToJson(bodyFormat, data, transportData)

		if len(debugMessage) > 5*1024 {
			debugMessage = debugMessage[0:5*1024] + "......"
		}

		debugMessage = fmt.Sprintf("DEBUG =====client.Request ViewInJson From %s=====\n%s\n", conn.remoteAddr(), debugMessage)
		logger.Debug(debugMessage)
	}

	if "" == action {
		err = errors.New("noActionInMessage")
		return
	}

	var bodyData []byte
	bodyData, keepSocket, _ = app.requestHttpBackend(conn.addr(), action, bodyFormat, data, conn.getSocketId())

	responseSize = len(bodyData)
	if len(bodyData) == 0 {
		return
	}

	var writeError error
	writeError = conn.writeMessage(bodyData)

	if EnableDebug {
		responseTransportData := app.getTransportData(bodyFormat, bodyData)
		debugMessage := app.bytesToJson(bodyFormat, bodyData, responseTransportData)
		debugMessageBase64 := base64.StdEncoding.EncodeToString(bodyData)

		if len(debugMessage) > 5*1024 {
			debugMessage = debugMessage[0:5*1024] + "......"
			debugMessageBase64 = "_too_large_data_"
		}

		var viewInJson = fmt.Sprintf("DEBUG =====httpBackend.Response ViewInJson=====\n%s\n", debugMessage)
		var originBody = fmt.Sprintf("DEBUG =====httpBackend.Response ViewInBase64=====\n%s\n", debugMessageBase64)
		logger.Debugf(
			"websocket client:%-21s writeToClient writeError: %v\n%s\n%s\n",
			conn.remoteAddr(),
			writeError,
			viewInJson,
			originBody,
		)
	}
}

func (app *serverHandler) getTransportData(bodyFormat string, tData []byte) *core.TransportData {

	responseTransportData := &core.TransportData{}
	var err error
	if bodyFormat == "json" {
		buf := bytes.NewReader(tData)
		err = jsonpb.Unmarshal(buf, responseTransportData)
	} else if bodyFormat == "pb" {
		err = proto.Unmarshal(tData, responseTransportData)
	} else if bodyFormat == "base64pb" {

	}

	if nil != err {
		logger.Warning(err)
	}
	return responseTransportData
}

func (app *serverHandler) requestHttpBackend(addr string, action string, bodyFormat string, data []byte, socketId string) ([]byte, bool, error) {
	// copy from zaly-gw
	// to http-backend

	url := fmt.Sprintf(
		"%s?body_format=%s&gw-host=%s&gw-port=%d&gw-socket-id=%s&proxy_remote_addr=127.0.0.1:27361&action=%s",
		HttpBackendAddr,
		bodyFormat,
		HttpServerIp,
		httpServerPort,
		socketId,
		action,
	)

	resp, err := http.Post(url, "text/protobuf", bytes.NewReader(data))
	defer func() {
		if nil != resp && nil != resp.Body {
			resp.Body.Close()
		}
	}()

	if nil != err {
		logger.Warning(err)
		return []byte{}, false, err
	}

	keepSocket := resp.Header.Get("KeepSocket")
	bodyData, _ := ioutil.ReadAll(resp.Body)

	return bodyData, keepSocket == "true", nil
}

func init() {
	handlerInstance = &serverHandler{}
	connMap = make(map[string]gatewayConnHelper)
}
