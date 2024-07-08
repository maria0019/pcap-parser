package gopacket_mock

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"pparse/mock"
	"time"
)

type PacketHTTPReqMock struct {
}

// // Functions for outputting the packet as a human-readable string:
// // ------------------------------------------------------------------
// String returns a human-readable string representation of the packet.
// It uses LayerString on each layer to output the layer.
func (m PacketHTTPReqMock) String() string {
	return ""
}

// Dump returns a verbose human-readable string representation of the packet,
// including a hex dump of all layers.  It uses LayerDump on each layer to
// output the layer.
func (m PacketHTTPReqMock) Dump() string {
	return ""
}

func (m PacketHTTPReqMock) Layers() []gopacket.Layer {
	return []gopacket.Layer{}
}

// Layer returns the first layer in this packet of the given type, or nil
func (m PacketHTTPReqMock) Layer(layerType gopacket.LayerType) gopacket.Layer {
	payload := `GET / HTTP/1.1
Host: captive.apple.com

`
	optionData := "11112222"

	l := layers.TCP{
		BaseLayer: layers.BaseLayer{Payload: []byte(payload)},
		Options: []layers.TCPOption{{
			OptionType: layers.TCPOptionKindTimestamps,
			OptionData: []byte(optionData),
		}},
	}

	return &l
}

// LayerClass returns the first layer in this packet of the given class,
// or nil.
func (m PacketHTTPReqMock) LayerClass(class gopacket.LayerClass) gopacket.Layer {
	return gopacket.Payload{}
}

// LinkLayer returns the first link layer in the packet
func (m PacketHTTPReqMock) LinkLayer() gopacket.LinkLayer {
	return &layers.Ethernet{}
}

// NetworkLayer returns the first network layer in the packet
func (m PacketHTTPReqMock) NetworkLayer() gopacket.NetworkLayer {
	return &layers.IPv4{}
}

// TransportLayer returns the first transport layer in the packet
func (m PacketHTTPReqMock) TransportLayer() gopacket.TransportLayer {
	return &layers.TCP{Options: []layers.TCPOption{
		{
			OptionType: layers.TCPOptionKindTimestamps,
		},
	}}
}

// ApplicationLayer returns the first application layer in the packet
func (m PacketHTTPReqMock) ApplicationLayer() gopacket.ApplicationLayer {
	return &layers.TLS{}
}

// ErrorLayer is particularly useful, since it returns nil if the packet
// was fully decoded successfully, and non-nil if an error was encountered
// in decoding and the packet was only partially decoded.  Thus, its output
// can be used to determine if the entire packet was able to be decoded.
func (m PacketHTTPReqMock) ErrorLayer() gopacket.ErrorLayer {
	return nil
}

// Data returns the set of bytes that make up this entire packet.
func (m PacketHTTPReqMock) Data() []byte {
	return []byte{}
}

// Metadata returns packet metadata associated with this packet.
func (m PacketHTTPReqMock) Metadata() *gopacket.PacketMetadata {
	t, _ := time.Parse(time.DateTime, mock.TimeNow)

	return &gopacket.PacketMetadata{
		CaptureInfo: gopacket.CaptureInfo{Timestamp: t},
		Truncated:   false,
	}
}
