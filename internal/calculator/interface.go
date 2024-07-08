package calculator

import (
	"pparse/internal/packet"
	"time"
)

type CalculatorI interface {
	ExtractPacketValues(p packet.ParsedPacketI)
	StatsAsMap(at time.Time) map[string]any
	Stats(at time.Time) any
	Cleanup()
}
