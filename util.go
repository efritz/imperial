package imperial

import "time"

func boolptr(val bool) *bool           { return &val }
func stringptr(val string) *string     { return &val }
func float64ptr(val float64) *float64  { return &val }
func timeptr(val time.Time) *time.Time { return &val }
