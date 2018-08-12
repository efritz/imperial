package imperial

//go:generate go-mockgen github.com/efritz/imperial -i Reporter -o mock_reporter_test.go -f

import (
	"testing"

	"github.com/aphistic/sweet"
	"github.com/aphistic/sweet-junit"
	. "github.com/onsi/gomega"
)

func TestMain(m *testing.M) {
	RegisterFailHandler(sweet.GomegaFail)

	sweet.Run(m, func(s *sweet.S) {
		s.RegisterPlugin(junit.NewPlugin())

		s.AddSuite(&RiemannSuite{})
		s.AddSuite(&MultiReporterSuite{})
	})
}
