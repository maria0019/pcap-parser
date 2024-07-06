package network

import (
	"pparse/internal/calculator"
	"pparse/internal/packet"
	"pparse/internal/sender"
	"time"
)

type Counter struct {
	calculator calculator.CalculatorI
	interval   time.Duration
	ticker     *time.Ticker
}

func NewCounter(calculator calculator.CalculatorI, timeGap int) Counter {
	return Counter{
		calculator: calculator,
		ticker:     time.NewTicker(time.Second * time.Duration(timeGap)),
	}
}

func (c *Counter) Init() {
	go func() {
		for t := range c.ticker.C {
			// print stats when the time border is reached
			sender.ToStdout("Stats", c.calculator.StatsAsMap(t))

			c.calculator.Cleanup() // reset stats calculator when the time border is reached
		}
	}()
}

func (c *Counter) ProcessPacket(parsed packet.ParsedPacketI) {
	c.calculator.ExtractPacketValues(parsed) // collect statistics while time border is not reached
}
