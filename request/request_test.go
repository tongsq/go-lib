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
