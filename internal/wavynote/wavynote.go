// wavynote 모듈에 포함된 모든 패키지에서 공통적으로 사용되는 구조체를 정의
package wavynote

// 옵션화할 대상
const (
	DATABASE_HOST     = "ep-white-morning-a2cks0uu.eu-central-1.pg.koyeb.app"
	DATABASE_PORT     = 5432
	DATABASE_LOGIN    = "wavynote"
	DATABASE_PASSWORD = "9Ru7lqmOwpcn"
	DATABASE_DB       = "wavynote"
	DATABASE_SSLMODE  = "require"
	DATABASE_APPNAME  = "wavynoted"

	HTTPSERVER_IP        = ""
	HTTPSERVER_PORT      = 16770
	HTTPSERVER_CERT_PATH = "server.crt"
	HTTPSERVER_PKEY_PATH = "server.key"
	HTTPSERVER_RTIMEOUT  = 3600
	HTTPSERVER_WTIMEOUT  = 3600
	HTTPSERVER_USE_HTTPS = "off"
)

type HTTPServerInfo struct {
	Ip       string
	Port     int
	Cert     string
	Pkey     string
	Rtimeout int
	Wtimeout int
	UseHttps string
}

type DataBaseInfo struct {
	Host     string
	Port     int
	Login    string
	Password string
	Database string
	SSLMode  string
	AppName  string
}
