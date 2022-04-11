package request

import (
	"net/http"
	"net/url"
	"time"
)

func get(requestUrl string, header *HeaderDto, cookie map[string]string, proxy *ProxyDto) (*HttpResultDto, error) {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, header)
	req = addCookie(req, cookie)
	var client *http.Client
	if proxy != nil {
		proxyServer := getProxyUrl(proxy)
		proxyUrl, _ := url.Parse(proxyServer)
		client = &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
			Timeout:   time.Second * 5,
		}
	} else {
		client = &http.Client{
			Timeout: time.Second * 5,
		}
	}
	return request(client, req)
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
