package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/erzha/elog"
)

var HttpServerIp string = "127.0.0.1"
var httpServerPort uint16

var ZalyServerAddr string = ":2021"
var WebsocketServerAddr string = ":2031"
var HttpBackendAddr string = "http://127.0.0.1:8888/"
var EnableDebug bool = false

var logger *elog.Logger

func requestHttpBackend(action string, bodyFormat string, data []byte) ([]byte, bool, error) {

	// copy from zaly-gw
	// to http-backend
	url := fmt.Sprintf(
		"%s?body_format=%s&proxy-id=zaly&proxy_remote_addr=127.0.0.1:27361&action=%s",
		HttpBackendAddr,
		bodyFormat,
		action,
	)

	resp, _ := http.Post(url, "text/protobuf", bytes.NewReader(data))
	defer resp.Body.Close()

	keepSocket := resp.Header.Get("KeepSocket")
	bodyData, _ := ioutil.ReadAll(resp.Body)

	return bodyData, keepSocket == "true", nil
}

func SetLoggerLevel() {
	logger = elog.NewLogger()
	if EnableDebug {
		logger.SetMinLogLevel(elog.LEVEL_DEBUG)
	} else {
		logger.SetMinLogLevel(elog.LEVEL_INFO)
	}
}
