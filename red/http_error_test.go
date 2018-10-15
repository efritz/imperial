package red

import (
	"net/http"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

type HTTPErrorSuite struct{}

func (s *HTTPErrorSuite) TestBasic(t sweet.T) {
	err := NewHTTPStatusError(http.StatusConflict, "message already exists")
	Expect(err.Code()).To(Equal(http.StatusConflict))
	Expect(err.Error()).To(Equal("http error: code = Conflict desc = message already exists"))
}
