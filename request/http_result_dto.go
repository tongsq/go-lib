package request

import "encoding/json"

func NewHttpResultDto() *HttpResultDto {
	d := new(HttpResultDto)
	d.Cookies = make(map[string]string)
	return d
}

type HttpResultDto struct {
	Body          string
	HttpCode      int
	ContentLength int64
	Header        map[string][]string
	Cookies       map[string]string
}

func (s HttpResultDto) String() string {
	m, _ := json.Marshal(s)
	return string(m)
}
