package request

import (
	"net/http"
)

// SimpleGet Simple Get
func SimpleGet(requestUrl string) (*HttpResultDto, error) {
	return Get(requestUrl, NewOptions())
}

// Get with options
func Get(requestUrl string, options *Options) (*HttpResultDto, error) {
	if options.Header == nil {
		options.Header = &HeaderDto{UserAgent: HTTP_USER_AGENT}
	}
	req, _ := http.NewRequest("GET", requestUrl, nil)
	req = addHeader(req, options.Header)
	req = addCookie(req, options.Cookie)
	if options.Data != nil {
		q := req.URL.Query()
		for k, v := range options.Data {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}
	return request(req, options.Proxy, options.Timeout)
}
