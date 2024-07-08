package mock

import (
	"pparse/internal/packet"
	"time"
)

type Calculator struct{}

func (c Calculator) ExtractPacketValues(p packet.ParsedPacketI) {}

func (c Calculator) StatsAsMap(at time.Time) map[string]any {
	return map[string]any{"avg": 3}
}

func (c Calculator) Stats(at time.Time) any {
	return map[string]any{"avg": 3}
}

func (c Calculator) Cleanup() {}
