package ecode

var HTTP_CODE_ERROR = NewError(-900, "http返回码不为200")
var HTTP_GZIP_DECODE_ERROR = NewError(-901, "gzip解码失败")
var HTTP_READ_ERROR = NewError(-902, "请求response读取失败")

type Error struct {
	ErrCode int
	ErrMsg  string
}

func NewError(code int, msg string) *Error {
	return &Error{ErrMsg: msg, ErrCode: code}
}

func (err *Error) Error() string {
	return err.ErrMsg
}
