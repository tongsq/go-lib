package request

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/**
post with cookies and headers
*/
func WebPost(requestUrl string, data map[string]string, header *HeaderDto, cookie map[string]string) (*HttpResultDto, error) {
	return WebPostProxy(requestUrl, data, header, cookie, nil)
}

/**
post with cookies and headers
*/
func WebPostProxy(requestUrl string, data map[string]string, header *HeaderDto, cookie map[string]string, proxy *ProxyDto) (*HttpResultDto, error) {
	if header == nil {
		header = &HeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	if header.ContentType == "" {
		header.ContentType = CONTENT_TYPE_FORM
	}
	param := GetReqData(data)
	return post(requestUrl, param, header, cookie, proxy)
}

/**
post with cookies and headers
*/
func WebPostJson(requestUrl string, data map[string]string, header *HeaderDto, cookie map[string]string) (*HttpResultDto, error) {
	return WebPostJsonProxy(requestUrl, data, header, cookie, nil)
}

/**
post with cookies and headers
*/
func WebPostJsonProxy(requestUrl string, data map[string]string, header *HeaderDto, cookie map[string]string, proxy *ProxyDto) (*HttpResultDto, error) {
	if header == nil {
		header = &HeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	if header.ContentType == "" {
		header.ContentType = CONTENT_TYPE_JSON
	}
	jsonData, _ := json.Marshal(data)
	return post(requestUrl, string(jsonData), header, cookie, proxy)
}

/**
Simple post
*/
func Post(requestUrl string, data map[string]string) (*HttpResultDto, error) {
	return WebPost(requestUrl, data, nil, map[string]string{})
}

/**
Simple post use proxy
*/
func PostJson(requestUrl string, data map[string]string) (*HttpResultDto, error) {
	return WebPostJson(requestUrl, data, nil, map[string]string{})
}

func post(requestUrl string, data string, header *HeaderDto, cookie map[string]string, proxy *ProxyDto) (*HttpResultDto, error) {
	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(data))
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
