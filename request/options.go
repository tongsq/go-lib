package request

import "time"

type DataType string

const JSON DataType = "json"
const FORM DataType = "form"
const DEFAULT DataType = ""

type Options struct {
	Header   *HeaderDto
	Cookie   map[string]string
	Proxy    *ProxyDto
	Timeout  time.Duration
	Data     map[string]string
	JsonData interface{}
	DataType DataType
}

func NewOptions() *Options {
	return &Options{
		Timeout:  DefaultTimeout,
		DataType: DEFAULT,
	}
}

func (o *Options) WithHeader(header *HeaderDto) *Options {
	o.Header = header
	return o
}

func (o *Options) WithCookie(cookie map[string]string) *Options {
	o.Cookie = cookie
	return o
}

func (o *Options) WithProxy(proxy *ProxyDto) *Options {
	o.Proxy = proxy
	return o
}

func (o *Options) WithJsonData(data interface{}) *Options {
	o.JsonData = data
	return o
}

func (o *Options) WithData(data map[string]string) *Options {
	o.Data = data
	return o
}

func (o *Options) WithDataType(t DataType) *Options {
	o.DataType = t
	return o
}

func (o *Options) WithTimeout(timeout time.Duration) *Options {
	o.Timeout = timeout
	return o
}
