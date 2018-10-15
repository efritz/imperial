package imperial

import "github.com/efritz/imperial/red"

type (
	REDReporter      = red.Reporter
	ErrorInterpreter = red.ErrorInterpreter
	HTTPStatusError  = red.HTTPStatusError
)

var (
	NewREDReporter           = red.NewReporter
	DefaultErrorInterpreter  = red.DefaultErrorInterpreter
	HTTPErrorInterpreter     = red.HTTPErrorInterpreter
	GRPCErrorInterpreter     = red.GRPCErrorInterpreter
	NewMultiErrorInterpreter = red.NewMultiErrorInterpreter
	NewHTTPStatusError       = red.NewHTTPStatusError
)
