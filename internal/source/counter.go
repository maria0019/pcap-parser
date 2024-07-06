package source

import (
	"pparse/internal/calculator"
	"pparse/internal/packet"
	"pparse/internal/source/file"
	"pparse/internal/source/network"
)

type CounterI interface {
	Init()
	ProcessPacket(parsed packet.ParsedPacketI)
}

func NewCounter(filePath, netInterface string, c calculator.CalculatorI, timeGap int) CounterI {
	var counter CounterI

	switch {
	case filePath != "":
		c := file.NewCounter(c, timeGap)

		return &c
	case netInterface != "":
		c := network.NewCounter(c, timeGap)

		return &c
	}

	return counter
}
