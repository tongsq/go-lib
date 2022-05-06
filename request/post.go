package request

import (
	"encoding/json"
	"net/http"
	"strings"
)

// SimplePost Post
func SimplePost(requestUrl string, data map[string]string) (*HttpResultDto, error) {
	return Post(requestUrl, NewOptions().WithData(data))
}

// Post Post with options
func Post(requestUrl string, options *Options) (*HttpResultDto, error) {
	if options.Header == nil {
		options.Header = &HeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	var param string
	if options.JsonData != nil || options.DataType == JSON {
		if options.Header.ContentType == "" {
			options.Header.ContentType = CONTENT_TYPE_JSON
		}
		var jsonData []byte
		if options.JsonData != nil {
			jsonData, _ = json.Marshal(options.JsonData)
		} else if options.Data != nil {
			jsonData, _ = json.Marshal(options.Data)
		}
		param = string(jsonData)
	} else {
		if options.Header.ContentType == "" {
			options.Header.ContentType = CONTENT_TYPE_FORM
		}
		param = GetReqData(options.Data)
	}
	req, _ := http.NewRequest("POST", requestUrl, strings.NewReader(param))
	req = addHeader(req, options.Header)
	req = addCookie(req, options.Cookie)
	return request(req, options.Proxy, options.Timeout)
}
