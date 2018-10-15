package red

import (
	"fmt"

	"google.golang.org/grpc/status"
)

type ErrorInterpreter func(err error) (string, bool)

func DefaultErrorInterpreter(err error) (string, bool) {
	return "", err != nil
}

func HTTPErrorInterpreter(err error) (string, bool) {
	if httpErr, ok := err.(*HTTPStatusError); ok {
		return fmt.Sprintf("%d", httpErr.Code()), true
	}

	return "", false
}

func GRPCErrorInterpreter(err error) (string, bool) {
	if st, ok := status.FromError(err); ok {
		return st.Code().String(), true
	}

	return "", false
}

func NewMultiErrorInterpreter(interpreters ...ErrorInterpreter) ErrorInterpreter {
	return func(err error) (string, bool) {
		for _, interpreter := range interpreters {
			if code, ok := interpreter(err); ok {
				return code, true
			}
		}

		return "", false
	}
}
