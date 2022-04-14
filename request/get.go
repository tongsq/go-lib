package request

import (
	"net/http"
)

func get(requestUrl string, header *HeaderDto, cookie map[string]string, proxy *ProxyDto) (*HttpResultDto, error) {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, header)
	req = addCookie(req, cookie)

	return request(req, proxy)
}

/**
get request use proxy
*/
func WebGetProxy(requestUrl string, header *HeaderDto, cookie map[string]string, proxy *ProxyDto) (*HttpResultDto, error) {
	if header == nil {
		header = &HeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	return get(requestUrl, header, cookie, proxy)
}

/**
get with headers and cookies
*/
func WebGet(requestUrl string, header *HeaderDto, cookie map[string]string) (*HttpResultDto, error) {
	return WebGetProxy(requestUrl, header, cookie, nil)
}

//Simple get
func Get(requestUrl string) (*HttpResultDto, error) {
	return WebGet(requestUrl, nil, nil)
}

//Simple get use proxy
func GetProxy(requestUrl string, proxy *ProxyDto) (*HttpResultDto, error) {
	return WebGetProxy(requestUrl, nil, nil, proxy)
}
