package file

import (
	calculatorMock "pparse/internal/calculator/mock"
	packetMock "pparse/internal/packet/mock"
	"testing"
)

func Test_Counter_Packet(t *testing.T) {
	t.Run("Process valid packet", func(t *testing.T) {
		calc := calculatorMock.Calculator{}
		p := packetMock.ParsedPacketValid{}
		timeInterval := 1

		c := NewCounter(calc, timeInterval)
		c.Init()
		c.ProcessPacket(p)
	})

	t.Run("Process invalid packet", func(t *testing.T) {
		calc := calculatorMock.Calculator{}
		p := packetMock.ParsedPacketInvalid{}
		timeInterval := 1

		c := NewCounter(calc, timeInterval)
		c.Init()
		c.ProcessPacket(p)
	})
}
