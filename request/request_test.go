package request

import (
	"fmt"
	"testing"
)

func TestWebPostJson(t *testing.T) {
	u := "http://220.181.38.148"
	h := RequestHeaderDto{
		//UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "baidu.com",
		ContentType:             CONTENT_TYPE_JSON,
	}
	param := map[string]string{"name": "aa"}
	data, err := WebPostJson(u, param, &h, nil)
	fmt.Printf("%#v", data)
	if err != nil || data.HttpCode != HTTP_CODE_OK {
		t.Fatal("test post request fail", err)
	} else {
		t.Log("test post request success", err)
	}
}

func TestWebGet(t *testing.T) {
	u := "https://httpbin.org/anything"
	h := RequestHeaderDto{
		//UserAgent:               consts.USER_AGENT,
		UpgradeInsecureRequests: "1",
		Host:                    "example.com",
		ContentType:             CONTENT_TYPE_JSON,
	}
	data, err := WebGet(u, &h, map[string]string{"ck": "cookie demo"})
	fmt.Printf("%#v\n", data)
	if err != nil || data.HttpCode != HTTP_CODE_OK {
		t.Fatal("test get request fail", err)
	} else {
		t.Log("test get request success:", data.Body)
	}
}

func TestWebGetGbk(t *testing.T) {
	u := "http://qt.gtimg.cn/q=sh601318,sh000107"
	h := RequestHeaderDto{
		UserAgent: HTTP_USER_AGENT,
	}
	data, err := WebGet(u, &h, nil)
	fmt.Printf("%#v\n", data)
	if err != nil || data.HttpCode != HTTP_CODE_OK {
		t.Fatal("test get request fail", err)
	} else {
		t.Log("test get request success:", data.Body)
	}
}
