package source

import (
	"fmt"
	"github.com/google/gopacket"
	"pparse/internal/source/file"
	"pparse/internal/source/network"
)

type DataSourceI interface {
	Path() string
	Type() string
	Packets() (chan gopacket.Packet, error)
}

func NewDataSource(filePath, netInterface string) (DataSourceI, error) {
	switch {
	case filePath != "":
		return file.NewSource(filePath), nil
	case netInterface != "":
		return network.NewSource(netInterface), nil
	default:
		return nil, fmt.Errorf("data source is not provided")
	}
}
