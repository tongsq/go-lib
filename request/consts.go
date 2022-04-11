package request

const HTTP_USER_AGENT = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"
const HTTP_CODE_OK = 200
const CONTENT_TYPE_JSON = "application/json"
const CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"

const PROTO_HTTP = "http"
const PROTO_HTTPS = "https"
const PROTO_SOCKS4 = "socks4"
const PROTO_SOCKS5 = "socks5"

// protocols
var PROTO_LIST = []string{
	PROTO_HTTP,
	PROTO_SOCKS5,
	PROTO_HTTPS,
	PROTO_SOCKS4,
}
