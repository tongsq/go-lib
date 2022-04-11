package tests

import (
	"testing"

	"github.com/tongsq/go-lib/request"
)

func TestSimpleGet(t *testing.T) {
	u := "https://www.baidu.com"
	data, err := request.Get(u)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test get request fail", err)
	} else {
		t.Log("test get request success", err)
	}
}

func TestSimplePost(t *testing.T) {
	u := "https://www.baidu.com"
	data, err := request.Post(u, nil)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test post request fail", err)
	} else {
		t.Log("test post request success", err)
	}
}

func TestWebGetProxy(t *testing.T) {
	u := "https://api.ip.sb/ip"
	data, err := request.WebGetProxy(u, nil, nil, &request.ProxyDto{Host: "127.0.0.1", Port: "9999"})
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test WebGetProxy fail", err)
	} else {
		t.Log("test WebGetProxy success: ", data.Body)
	}
}

func TestWebGet(t *testing.T) {
	u := "https://api.ip.sb/ip"
	data, err := request.WebGet(u, &request.HeaderDto{UserAgent: request.HTTP_USER_AGENT, Referer: "abc"}, map[string]string{"session": "abc"})
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test WebGet fail", err)
	} else {
		t.Log("test WebGet success: ", data.Body)
	}
}
