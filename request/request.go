package request

import (
	"compress/gzip"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/shadowsocks/shadowsocks-go/shadowsocks"
	"github.com/tongsq/go-lib/ecode"
	"github.com/tongsq/go-lib/util"
	"h12.io/socks"
)

var DefaultTimeout = time.Second * 5

func SetTimeout(t time.Duration) {
	DefaultTimeout = t
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

/**
http request
*/
func request(req *http.Request, proxy *ProxyDto, timeout time.Duration) (*HttpResultDto, error) {
	var client *http.Client
	var tran *http.Transport

	if proxy != nil {
		proxyServer := getProxyUrl(proxy)
		switch proxy.Proto {
		case PROTO_SS:
			cipher, err := shadowsocks.NewCipher(proxy.User, proxy.Password)
			if err != nil {
				return nil, err
			}
			tran = &http.Transport{
				Dial: func(_, addr string) (net.Conn, error) {
					return shadowsocks.Dial(addr, fmt.Sprintf("%s:%s", proxy.Host, proxy.Port), cipher.Copy())
				},
			}
		case PROTO_SOCKS4, PROTO_SOCKS4A:
			dial := socks.Dial(proxyServer)
			tran = &http.Transport{
				Dial: dial,
			}
		default:
			proxyUrl, err := url.Parse(proxyServer)
			if err != nil {
				return nil, err
			}
			tran = &http.Transport{
				Proxy:           http.ProxyURL(proxyUrl),
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
		}
	}
	if tran != nil {
		client = &http.Client{
			Transport: tran,
			Timeout:   timeout,
		}
	} else {
		client = &http.Client{
			Timeout: timeout,
		}
	}
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
	types, ok := result.Header["Content-Type"]
	bodyStr := string(body)
	if ok {
		for _, t := range types {
			if util.Contains(t, []string{"GBK", "gbk", "gb2312", "GB2312"}) {
				b, err := util.GbkToUtf8(bodyStr)
				if err == nil {
					bodyStr = b
				}
				break
			}
		}
	}
	result.Body = bodyStr
	if resp.StatusCode != HTTP_CODE_OK {
		return result, ecode.HTTP_CODE_ERROR
	}
	return result, nil
}

/**
add request header
*/
func addHeader(req *http.Request, h *HeaderDto) *http.Request {
	if h == nil {
		return req
	}
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
	if h.Other != nil {
		for k, v := range h.Other {
			req.Header.Set(k, v)
		}
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

func getProxyUrl(p *ProxyDto) string {
	proto := p.Proto
	if proto == "" {
		proto = PROTO_HTTP
	}
	if p.User == "" {
		return fmt.Sprintf("%s://%s:%s", proto, p.Host, p.Port)
	} else {
		return fmt.Sprintf("%s://%s:%s@%s:%s", proto, p.User, p.Password, p.Host, p.Port)
	}
}
