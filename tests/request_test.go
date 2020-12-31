package tests

import (
	"github.com/tongsq/go-lib/request"
	"testing"
)

func TestGetRequest(t *testing.T) {
	u := "https://www.baidu.com"
	data, err := request.WebGet(u, nil, nil)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test get request fail", data, err)
	} else {
		t.Log("test get request success", data, err)
	}
}

func TestPostRequest(t *testing.T) {
	u := "https://www.baidu.com"
	data, err := request.WebPost(u, nil, nil, nil)
	if err != nil || data.HttpCode != request.HTTP_CODE_OK {
		t.Fatal("test post request fail", data, err)
	} else {
		t.Log("test post request success", data, err)
	}
}
