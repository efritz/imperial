package red

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc/status"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type ErrorInterpreterSuite struct{}

func (s *ErrorInterpreterSuite) TestDefaultErorInterpreter(t sweet.T) {
	err := fmt.Errorf("any error")
	code, ok := DefaultErrorInterpreter(err)
	Expect(ok).To(BeTrue())
	Expect(code).To(Equal(""))
}

func (s *ErrorInterpreterSuite) TestDefaultErorInterpreterNil(t sweet.T) {
	_, ok := DefaultErrorInterpreter(nil)
	Expect(ok).To(BeFalse())
}

func (s *ErrorInterpreterSuite) TestHTTPErrorInterpreter(t sweet.T) {
	err := NewHTTPStatusError(http.StatusGatewayTimeout, "")
	code, ok := HTTPErrorInterpreter(err)
	Expect(ok).To(BeTrue())
	Expect(code).To(Equal("504"))
}

func (s *ErrorInterpreterSuite) TestHTTPErrorInterpreterNonHTTPError(t sweet.T) {
	err := fmt.Errorf("not an http error")
	_, ok := HTTPErrorInterpreter(err)
	Expect(ok).To(BeFalse())
}

func (s *ErrorInterpreterSuite) TestGRPCErrorInterpreter(t sweet.T) {
	err := status.Error(codes.DeadlineExceeded, "")
	code, ok := GRPCErrorInterpreter(err)
	Expect(ok).To(BeTrue())
	Expect(code).To(Equal("DeadlineExceeded"))
}

func (s *ErrorInterpreterSuite) TestGRPCErrorInterpreterNonGRPCError(t sweet.T) {
	err := fmt.Errorf("not a grpc error")
	_, ok := GRPCErrorInterpreter(err)
	Expect(ok).To(BeFalse())
}

func (s *ErrorInterpreterSuite) TestMultiErrorInterpreter(t sweet.T) {
	interpreter := NewMultiErrorInterpreter(
		HTTPErrorInterpreter,
		GRPCErrorInterpreter,
	)

	var err error

	// Match first
	err = NewHTTPStatusError(http.StatusGatewayTimeout, "")
	code, ok := interpreter(err)
	Expect(ok).To(BeTrue())
	Expect(code).To(Equal("504"))

	// Match second
	err = status.Error(codes.DeadlineExceeded, "")
	code, ok = interpreter(err)
	Expect(ok).To(BeTrue())
	Expect(code).To(Equal("DeadlineExceeded"))

	// Match none
	unexpectedErr := fmt.Errorf("any error")
	_, ok = interpreter(unexpectedErr)
	Expect(ok).To(BeFalse())
}
