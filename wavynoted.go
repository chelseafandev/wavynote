package main

import (
	"github.com/wavynote/internal/gateway/http"
	"github.com/wavynote/internal/wavynote"
)

func main() {
	httpInfo := wavynote.HTTPServerInfo{
		Ip:       wavynote.HTTPSERVER_IP,
		Port:     wavynote.HTTPSERVER_PORT,
		Cert:     wavynote.HTTPSERVER_CERT_PATH,
		Pkey:     wavynote.HTTPSERVER_PKEY_PATH,
		Rtimeout: wavynote.HTTPSERVER_RTIMEOUT,
		Wtimeout: wavynote.HTTPSERVER_WTIMEOUT,
		UseHttps: wavynote.HTTPSERVER_USE_HTTPS,
	}

	dbInfo := wavynote.DataBaseInfo{
		Host:     wavynote.DATABASE_HOST,
		Port:     wavynote.DATABASE_PORT,
		Login:    wavynote.DATABASE_LOGIN,
		Password: wavynote.DATABASE_PASSWORD,
		Database: wavynote.DATABASE_DB,
		SSLMode:  wavynote.DATABASE_SSLMODE,
		AppName:  wavynote.DATABASE_APPNAME,
	}

	httpServer := http.NewHTTPServer(httpInfo, dbInfo, httpInfo.UseHttps)
	httpServer.StartServer()
}
