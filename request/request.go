package request

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tongsq/go-lib/ecode"
)

/**
get request use proxy
*/
func WebGetProxy(requestUrl string, header *RequestHeaderDto, host string, port string) (*HttpResultDto, error) {
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, header)
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	return request(client, req)
}

/**
get with headers and cookies
*/
func WebGet(requestUrl string, header *RequestHeaderDto, cookie map[string]string) (*HttpResultDto, error) {
	if header == nil {
		header = &RequestHeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, header)
	req = addCookie(req, cookie)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return request(client, req)
}

/**
post with cookies and headers
*/
func WebPost(requestUrl string, data map[string]string, header *RequestHeaderDto, cookie map[string]string) (*HttpResultDto, error) {
	if header == nil {
		header = &RequestHeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	if header.ContentType == "" {
		header.ContentType = CONTENT_TYPE_FORM
	}
	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(GetReqData(data)))
	req = addHeader(req, header)
	req = addCookie(req, cookie)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return request(client, req)
}

/**
post with cookies and headers
*/
func WebPostJson(requestUrl string, data map[string]string, header *RequestHeaderDto, cookie map[string]string) (*HttpResultDto, error) {
	if header == nil {
		header = &RequestHeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	if header.ContentType == "" {
		header.ContentType = CONTENT_TYPE_JSON
	}
	jsonData, _ := json.Marshal(data)
	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(string(jsonData)))
	req = addHeader(req, header)
	req = addCookie(req, cookie)
	client := &http.Client{
		Timeout: time.Second * 5,
	}
	return request(client, req)
}

/**
post without header cookie
*/
func ApiPost(requestUrl string, data map[string]string) (*HttpResultDto, error) {
	return WebPost(requestUrl, data, &RequestHeaderDto{}, map[string]string{})
}

/**
format query params
*/
func GetReqData(d map[string]string) string {
	s := ""
	for k, v := range d {
		s = s + fmt.Sprintf("%s=%s&", k, url.QueryEscape(v))
	}
	return s
}

func WebRequest(req *http.Request) (*HttpResultDto, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	return request(client, req)
}

func WebRequestProxy(req *http.Request, host string, port string) (*HttpResultDto, error) {
	proxyServer := fmt.Sprintf("http://%s:%s", host, port)
	proxyUrl, _ := url.Parse(proxyServer)
	client := &http.Client{
		Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)},
		Timeout:   time.Second * 5,
	}
	return request(client, req)
}

/**
http request
*/
func request(client *http.Client, req *http.Request) (*HttpResultDto, error) {
	result := NewHttpResultDto()
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	saveResponse(result, resp)
	defer resp.Body.Close()
	data := resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		data, err = gzip.NewReader(resp.Body)
		if err != nil {
			return result, ecode.HTTP_GZIP_DECODE_ERROR
		}
		defer data.Close()
	}
	body, err := ioutil.ReadAll(data)
	if err != nil {
		return result, ecode.HTTP_READ_ERROR
	}
	result.Body = string(body)
	if resp.StatusCode != HTTP_CODE_OK {
		return result, ecode.HTTP_CODE_ERROR
	}
	return result, nil
}

/**
add request header
*/
func addHeader(req *http.Request, h *RequestHeaderDto) *http.Request {
	if h.Host != "" {
		req.Header.Set("Host", h.Host)
		req.Host = h.Host
	}
	if h.Accept != "" {
		req.Header.Set("Accept", h.Accept)
	}
	if h.AcceptEncoding != "" {
		req.Header.Set("Accept-Encoding", h.AcceptEncoding)
	}
	if h.Referer != "" {
		req.Header.Set("Referer", h.Referer)
	}
	if h.UpgradeInsecureRequests != "" {
		req.Header.Set("Upgrade-Insecure-Requests", h.UpgradeInsecureRequests)
	}
	if h.UserAgent != "" {
		req.Header.Set("User-Agent", h.UserAgent)
	}
	if h.AcceptLanguage != "" {
		req.Header.Set("Accept-Language", h.AcceptLanguage)
	}
	if h.SecFetchDest != "" {
		req.Header.Set("Sec-Fetch-Dest", h.SecFetchDest)
	}
	if h.SecFetchMode != "" {
		req.Header.Set("Sec-Fetch-Mode", h.SecFetchMode)
	}
	if h.XRequestedWith != "" {
		req.Header.Set("X-Requested-With", h.XRequestedWith)
	}
	if h.ContentType != "" {
		req.Header.Set("Content-Type", h.ContentType)
	} else {
		req.Header.Set("Content-Type", CONTENT_TYPE_FORM)
	}
	return req
}

/**
add request cookie
*/
func addCookie(req *http.Request, c map[string]string) *http.Request {
	for k, v := range c {
		cookie := &http.Cookie{Name: k, Value: v, HttpOnly: true}
		req.AddCookie(cookie)
	}
	return req
}

/**
add http response info to dto
*/
func saveResponse(data *HttpResultDto, resp *http.Response) {
	data.HttpCode = resp.StatusCode
	data.ContentLength = resp.ContentLength
	data.Header = resp.Header
	for _, c := range resp.Cookies() {
		data.Cookies[c.Name] = c.Value
	}
}
