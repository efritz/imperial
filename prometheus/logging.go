package prometheus

import (
	"fmt"

	"github.com/efritz/imperial/base"
)

type loggerShim struct {
	base.Logger
}

func (l *loggerShim) Println(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			fmt.Printf(" ")
		}

		fmt.Printf("%v ", arg)
	}

	fmt.Println()
}
