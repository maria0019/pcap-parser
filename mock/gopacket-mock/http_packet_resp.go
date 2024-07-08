package gopacket_mock

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"pparse/mock"
	"time"
)

type PacketHTTPRespMock struct {
}

func (m PacketHTTPRespMock) String() string {
	return ""
}

func (m PacketHTTPRespMock) Dump() string {
	return ""
}

func (m PacketHTTPRespMock) Layers() []gopacket.Layer {
	return []gopacket.Layer{}
}

func (m PacketHTTPRespMock) Layer(layerType gopacket.LayerType) gopacket.Layer {
	payload := `HTTP/1.1 200 OK
Content-Length: 69
Cache-Control: max-age=31536000
Last-Modified: Fri, 21 Jun 2024 00:48:48 GMT
Content-Type: text/html
Date: Fri, 21 Jun 2024 00:48:48 GMT
Age: 1338207
Via: http/1.1 gbmnc1-edge-bx-007.ts.apple.com (acdn/153.14426), http/1.1 gbmnc1-edge-bx-007.ts.apple.com (acdn/252.14441)
X-Cache: miss, hit-fresh
CDNUUID: 7b1d4a56-4b1a-4efe-b0c7-b475c3f148b3-8502813648
Access-Control-Allow-Origin: *
Connection: keep-alive

<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>
`
	optionData := "33331111"

	l := layers.TCP{
		BaseLayer: layers.BaseLayer{Payload: []byte(payload)},
		Options: []layers.TCPOption{{
			OptionType: layers.TCPOptionKindTimestamps,
			OptionData: []byte(optionData)}},
	}

	return &l
}

func (m PacketHTTPRespMock) LayerClass(class gopacket.LayerClass) gopacket.Layer {
	return gopacket.Payload{}
}

func (m PacketHTTPRespMock) LinkLayer() gopacket.LinkLayer {
	return &layers.Ethernet{}
}

func (m PacketHTTPRespMock) NetworkLayer() gopacket.NetworkLayer {
	return &layers.IPv4{}
}

func (m PacketHTTPRespMock) TransportLayer() gopacket.TransportLayer {
	return &layers.TCP{Options: []layers.TCPOption{
		{
			OptionType: layers.TCPOptionKindTimestamps,
		},
	}}
}

func (m PacketHTTPRespMock) ApplicationLayer() gopacket.ApplicationLayer {
	return &layers.TLS{}
}

func (m PacketHTTPRespMock) ErrorLayer() gopacket.ErrorLayer {
	return nil
}

func (m PacketHTTPRespMock) Data() []byte {
	return []byte{}
}

func (m PacketHTTPRespMock) Metadata() *gopacket.PacketMetadata {
	t, _ := time.Parse(time.DateTime, mock.TimeNow)

	return &gopacket.PacketMetadata{
		CaptureInfo: gopacket.CaptureInfo{Timestamp: t},
		Truncated:   false,
	}
}
