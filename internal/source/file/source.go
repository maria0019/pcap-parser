package file

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const TypeFile = "file"

type Source struct {
	SrcPath string // path to file
}

func NewSource(path string) Source {
	return Source{SrcPath: path}
}

func (s Source) Path() string {
	return s.SrcPath
}

func (s Source) Type() string {
	return TypeFile
}

func (s Source) Packets() (chan gopacket.Packet, error) {
	handle, err := s.pcapHandle()
	if err != nil {
		return nil, err
	}

	h := gopacket.NewPacketSource(handle, handle.LinkType())

	return h.Packets(), nil
}

func (s Source) pcapHandle() (*pcap.Handle, error) {
	return pcap.OpenOffline(s.Path())
}
