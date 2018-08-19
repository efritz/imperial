package imperial

import "github.com/efritz/imperial/base"

type (
	Reporter = base.Reporter
	Logger   = base.Logger
)

const (
	UnitCount        = base.UnitCount
	UnitPercent      = base.UnitPercent
	UnitSeconds      = base.UnitSeconds
	UnitMicroseconds = base.UnitMicroseconds
	UnitMilliseconds = base.UnitMilliseconds
	UnitBytes        = base.UnitBytes
	UnitKilobytes    = base.UnitKilobytes
	UnitMegabytes    = base.UnitMegabytes
	UnitGigabytes    = base.UnitGigabytes
	UnitTerabytes    = base.UnitTerabytes
	UnitBits         = base.UnitBits
	UnitKilobits     = base.UnitKilobits
	UnitMegabits     = base.UnitMegabits
	UnitGigabits     = base.UnitGigabits
	UnitTerabits     = base.UnitTerabits
)

var (
	NewSimpleReporterShim = base.NewSimpleReporterShim
	WithNamespace         = base.WithNamespace
	WithSubsystem         = base.WithSubsystem
	WithHelp              = base.WithHelp
	WithUnit              = base.WithUnit
	WithAttributes        = base.WithAttributes
	WithExposedAttributes = base.WithExposedAttributes
	WithBuckets           = base.WithBuckets
	WithQuantiles         = base.WithQuantiles
	NewPrintLogger        = base.NewPrintLogger
	NewNilLogger          = base.NewNilLogger
)
