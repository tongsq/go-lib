package tests

import (
	"fmt"
	"testing"

	"github.com/tongsq/go-lib/request"
)

func TestSimpleGet(t *testing.T) {
	u := "https://www.baidu.com"
	data, err := request.SimpleGet(u)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test get request fail", err)
	} else {
		t.Log("test get request success", err)
	}
}

func TestGetRequest(t *testing.T) {
	u := "https://httpbin.org/anything?k1=v1"
	h := request.HeaderDto{
		//UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "example.com",
		ContentType:             request.CONTENT_TYPE_JSON,
	}
	query := map[string]string{"query": "abce"}
	data, err := request.Get(u, request.NewOptions().WithHeader(&h).WithCookie(map[string]string{"ck": "cookie demo"}).WithData(map[string]string{"k2": "v2"}).WithQuery(query))
	fmt.Printf("%#v\n", data)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test get request fail", err)
	} else {
		t.Log("test get request success:", data.Body)
	}
}

func TestWebGetGbk(t *testing.T) {
	u := "http://qt.gtimg.cn/q=sh601318,sh000107"
	h := request.HeaderDto{
		UserAgent: request.HTTP_USER_AGENT,
	}
	data, err := request.Get(u, request.NewOptions().WithHeader(&h))
	fmt.Printf("%#v\n", data)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test get request fail", err)
	} else {
		t.Log("test get request success:", data.Body)
	}
}

func TestWebGetProxy(t *testing.T) {
	u := "https://httpbin.org/anything"
	data, err := request.Get(u, request.NewOptions().WithProxy(&request.ProxyDto{Proto: request.PROTO_HTTP, Host: "127.0.0.1", Port: "8888", User: "root", Password: "123"}))
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test WebGetProxy fail", err)
	} else {
		t.Log("test WebGetProxy success: ", data.Body)
	}
}
