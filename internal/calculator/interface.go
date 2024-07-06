package calculator

import (
	"pparse/internal/packet"
	"time"
)

type CalculatorI interface {
	ExtractPacketValues(p packet.ParsedPacketI)
	StatsAsMap(at time.Time) map[string]interface{}
	Stats(at time.Time) interface{}
	Cleanup()
}
