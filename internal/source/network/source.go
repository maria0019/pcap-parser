package network

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	// The same default as tcpdump.
	defaultSnapLen = 262144
	bpfFilter      = "port 80"
	TypeNetwork    = "network"
)

type Source struct {
	SrcPath string // name of the network interface
}

func NewSource(path string) Source {
	return Source{SrcPath: path}
}

func (s Source) Path() string {
	return s.SrcPath
}

func (s Source) Type() string {
	return TypeNetwork
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
	handle, err := pcap.OpenLive(s.Path(), defaultSnapLen, true, pcap.BlockForever)
	if err != nil {
		return handle, err
	}

	if err := handle.SetBPFFilter(bpfFilter); err != nil { // TODO
		return handle, err
	}

	return handle, nil
}
