package source

import (
	"github.com/google/gopacket"
	"pparse/internal/source/file"
	"pparse/internal/source/network"
)

type DataSourceI interface {
	Path() string
	Type() string
	Packets() (chan gopacket.Packet, error)
}

func NewDataSource(filePath, netInterface string) DataSourceI {
	var dataSource DataSourceI

	switch {
	case filePath != "":
		return file.NewSource(filePath)
	case netInterface != "":
		return network.NewSource(netInterface)
	}

	return dataSource
}
