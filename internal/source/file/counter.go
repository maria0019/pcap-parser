package file

import (
	"pparse/internal/calculator"
	"pparse/internal/packet"
	"pparse/internal/sender"
	"time"
)

type Counter struct {
	calculator calculator.CalculatorI
	interval   time.Duration
	timeBorder time.Time
}

func NewCounter(calculator calculator.CalculatorI, timeGap int) Counter {
	return Counter{
		calculator: calculator,
		interval:   time.Second * time.Duration(timeGap),
		timeBorder: time.Time{},
	}
}

func (c *Counter) Init() {}

func (c *Counter) ProcessPacket(parsed packet.ParsedPacketI) {
	if c.timeBorder.IsZero() {
		c.timeBorder = parsed.Timestamp().Add(c.interval)
	}

	if parsed.Timestamp().Before(c.timeBorder) {
		c.calculator.ExtractPacketValues(parsed) // collect statistics while time border is not reached
	}

	if parsed.Timestamp().Equal(c.timeBorder) || parsed.Timestamp().After(c.timeBorder) {
		// print statistics when the time border is reached
		sender.ToStdout("Stats", c.calculator.StatsAsMap(c.timeBorder))

		c.calculator.Cleanup()                      // reset stats calculator when the time border is reached
		c.timeBorder = c.timeBorder.Add(c.interval) // set next time step
	}
}
