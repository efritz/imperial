package red

import (
	"fmt"
	"net/http"
	"strings"
)

type HTTPStatusError struct {
	code    int
	message string
}

func NewHTTPStatusError(code int, message string, args ...interface{}) *HTTPStatusError {
	return &HTTPStatusError{
		code:    code,
		message: fmt.Sprintf(message, args...),
	}
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf(
		"http error: code = %s desc = %s",
		strings.Replace(http.StatusText(e.code), " ", "", -1),
		e.message,
	)
}

func (e *HTTPStatusError) Code() int {
	if e == nil {
		return http.StatusOK
	}

	return e.code
}
