package errno

import "fmt"

// 错误信息处理
// 通过 errno.go 来对自定义的错误进行处理

/*
有两个核心函数 New() 和 DecodeErr() ，一个用来新建定制的错误，一 个用来解析定制的错误
同时也提供了 Add() 和 Addf() 函数，如果想对外展示更多的信息可以调用此函 数
*/

//Errno 错误码与错误信息的映射
type Errno struct {
	Code    int
	Message string
}

//错误码详情
func (err Errno) Error() string {
	return err.Message
}

type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{
		Code:    errno.Code,
		Message: errno.Message,
		Err:     err,
	}
}

func (err *Err) Add(message string) error {
	err.Message += " " + message
	return err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	err.Message += " " + fmt.Sprintf(format, args...)
	return err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
