package godi

import (
	"fmt"
	"github.com/lvyahui8/goenum"
)

type DIError struct {
	goenum.Enum
	ErrorMsg string
	// Cause 引起错误的根本error
	Cause *error
}

func (e DIError) Error() string {
	return e.ErrorMsg
}

func (e DIError) CreateError(cause *error, args ...any) *DIError {
	e.Cause = cause
	if len(args) > 0 {
		e.ErrorMsg = fmt.Sprintf(e.ErrorMsg, args)
	}
	// 将格式化好的新对象返回
	return &e
}

var (
	ErrSameBeanName = goenum.NewEnum[DIError]("ErrSameBeanName",
		DIError{ErrorMsg: "Beans with the same name are not allowed. beanName:%s,existType:%s,newType:%s"})
	ErrTagParseFailed = goenum.NewEnum[DIError]("ErrTagParseFailed",
		DIError{ErrorMsg: "Tag parsing exception. filedName:%s,tag:%s"})
	ErrBeanNotFound = goenum.NewEnum[DIError]("ErrBeanNotFound",
		DIError{ErrorMsg: "Bean [%s] not found"})
	ErrTypeMisMatch = goenum.NewEnum[DIError]("ErrTypeMisMatch",
		DIError{ErrorMsg: "Injection type mismatch, expected type %s, actual type %s"})
)
