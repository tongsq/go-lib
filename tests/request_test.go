package tests

import (
	"testing"

	"github.com/tongsq/go-lib/request"
)

func TestWebPostJson(t *testing.T) {
	u := "https://httpbin.org/anything"
	h := request.HeaderDto{
		//UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "example.com",
		ContentType:             request.CONTENT_TYPE_JSON,
	}
	param := map[string]string{"name": "aa"}
	query := map[string]string{"key": "areqfdq456", "va": "feq"}
	data, err := request.Post(u, request.NewOptions().WithDataType(request.JSON).WithData(param).WithHeader(&h).WithQuery(query))
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test post request fail", err)
	} else {
		t.Log("test post request success", data.Body)
	}
}

func TestSimplePost(t *testing.T) {
	u := "https://www.baidu.com"
	data, err := request.SimplePost(u, nil)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test post request fail", err)
	} else {
		t.Log("test post request success", err)
	}
}

func TestSocks4Proxy(t *testing.T) {
	u := "https://httpbin.org/anything"
	data, err := request.Get(u, request.NewOptions().WithProxy(&request.ProxyDto{Host: "218.75.69.50", Port: "56430", Proto: request.PROTO_SOCKS4}))
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test WebGetProxy fail", err)
	} else {
		t.Log("test WebGetProxy success: ", data.Body)
	}
}

func TestSocks5Proxy(t *testing.T) {
	u := "https://httpbin.org/anything"
	data, err := request.Get(u, request.NewOptions().WithProxy(&request.ProxyDto{Host: "38.142.63.146", Port: "31596", Proto: request.PROTO_SOCKS5}))
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test WebGetProxy fail", err)
	} else {
		t.Log("test WebGetProxy success: ", data.Body)
	}
}

func TestSsProxy(t *testing.T) {
	u := "https://api.ip.sb/ip"
	data, err := request.Get(u, request.NewOptions().WithProxy(&request.ProxyDto{Host: "127.0.0.1", Port: "1080", Proto: request.PROTO_SS, User: "chacha20", Password: "123"}))
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test WebGetProxy fail", err)
	} else {
		t.Log("test WebGetProxy success: ", data.Body)
	}
}
