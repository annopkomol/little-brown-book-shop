package domain

import "net/http"

type Error interface {
	Error() string
	GetHttpStatus() int
}

type appError struct {
	httpStatus int
	message    string
	err        error
}

func NewError() *appError {
	return &appError{
		httpStatus: http.StatusInternalServerError,
		message:    "internal server error",
	}
}

func NewErrorWithConfig(httpStatus int, msg string) *appError {
	return &appError{
		httpStatus: httpStatus,
		message:    msg,
	}
}

func (a *appError) Error() string {
	return a.message
}

func (a *appError) GetHttpStatus() int {
	return a.httpStatus
}

func (a *appError) WithError(err error) *appError {
	a.err = err
	return a
}
func (a *appError) WithMsg(msg string) *appError {
	a.message = msg
	return a
}
func (a *appError) WithHTTP(httpStatus int) *appError {
	a.httpStatus = httpStatus
	return a
}
