package source

import (
	"fmt"
	"pparse/internal/calculator"
	"pparse/internal/packet"
	"pparse/internal/source/file"
	"pparse/internal/source/network"
)

type CounterI interface {
	Init()
	ProcessPacket(parsed packet.ParsedPacketI)
}

func NewCounter(filePath, netInterface string, c calculator.CalculatorI, timeInterval int) (CounterI, error) {
	switch {
	case filePath != "":
		c := file.NewCounter(c, timeInterval)

		return &c, nil
	case netInterface != "":
		c := network.NewCounter(c, timeInterval)

		return &c, nil
	default:
		return nil, fmt.Errorf("counter is not found for data source")
	}
}
