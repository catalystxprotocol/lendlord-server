package error

import (
	"fmt"
	"net/http"
)

type Error interface {
	Error() string
	Code() int
	Msg() string
	StatusCode() int
}

type Err struct {
	code int
	msg  string
}

func (e Err) Error() string {
	return fmt.Sprintf("err_code: %d, err_msg: %s\n", e.code, e.msg)
}

func (e Err) Code() int {
	return e.code
}

func (e Err) Msg() string {
	return e.msg
}

func WrapErrorDetail(code int, msg string) Error {
	return Err{
		code: code,
		msg:  msg,
	}
}

func (e Err) StatusCode() int {
	switch e.Code() {
	case CodeInvalidParam:
		return http.StatusBadRequest
	case CodeInternalError:
		return http.StatusInternalServerError
	case CodeNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
