package mock

import (
	"github.com/google/gopacket"
	"pparse/internal/packet/mock"
)

const TypeFile = "file"

type SourceMock struct {
	SrcPath string // path to file
}

func NewSourceMock(path string) SourceMock {
	return SourceMock{SrcPath: path}
}

func (s SourceMock) Path() string {
	return s.SrcPath
}

func (s SourceMock) Type() string {
	return TypeFile
}

func (s SourceMock) Packets() (chan gopacket.Packet, error) {
	c := make(chan gopacket.Packet, 2)

	pack1 := mock.PacketHTTPReqMock{}
	c <- pack1

	pack2 := mock.PacketHTTPRespMock{}
	c <- pack2

	close(c)

	return c, nil
}
