package httppack

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	log "github.com/gookit/slog"
	"net/http"
	"pparse/config"
	"pparse/internal/calculator"
	"pparse/internal/source"
	"time"
)

const http11 = "HTTP/1.1"

type Parser struct {
	Config     config.ParserConfig
	DataSource source.DataSourceI
	Counter    source.CounterI
	Calculator calculator.CalculatorI
}

func (p Parser) Run() (int, error) {
	p.Counter.Init()

	packets, err := p.DataSource.Packets()
	if err != nil {
		return 0, err
	}

	var count int
	for item := range packets {
		parsed, err := extractPacketData(item)
		if err != nil {
			log.WithData(log.M{"error": err.Error()}).Error("Extract parser data error")
		}
		if parsed.IsValid() {
			p.Counter.ProcessPacket(parsed)
			count++
		}
	}

	return count, nil
}

func extractPacketData(packet gopacket.Packet) (ParsedPacket, error) {
	var data ParsedPacket

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return ParsedPacket{}, nil
	}

	if err := packet.ErrorLayer(); err != nil {
		return ParsedPacket{}, err.Error()
	}

	tcp, _ := tcpLayer.(*layers.TCP)
	if len(tcp.Payload) != 0 {
		reqReader := bufio.NewReader(bytes.NewReader(tcp.Payload))
		// err cases need more investigation. Currently, it says "Error in HTTP request: invalid method "HTTP/1.1""
		req, _ := http.ReadRequest(reqReader)

		if req != nil && req.Proto == http11 { // we can use switch if more Proto are needed
			data.URL = parseHttpRequestUrl(req)
			data.At = parsePacketTimestamp(packet)
			data.Uid = parseTcpRequestUid(tcp)

			return data, nil
		}

		respReader := bufio.NewReader(bytes.NewReader(tcp.Payload))
		resp, _ := http.ReadResponse(respReader, nil) // err should be processed same as for the ReadRequest

		if resp != nil && resp.Proto == http11 { // we can use switch if more Proto are needed
			data.At = parsePacketTimestamp(packet)
			data.Uid = parseTcpResponseUid(tcp)

			return data, nil
		}
	}

	return ParsedPacket{}, nil
}

func parsePacketTimestamp(packet gopacket.Packet) time.Time {
	return packet.Metadata().CaptureInfo.Timestamp
}

func parseTcpRequestUid(tcp *layers.TCP) int {
	for _, opt := range tcp.Options {
		if opt.OptionType == layers.TCPOptionKindTimestamps {
			return int(binary.BigEndian.Uint32(opt.OptionData[:4]))
		}
	}

	return 0
}

func parseTcpResponseUid(tcp *layers.TCP) int {
	for _, opt := range tcp.Options {
		if opt.OptionType == layers.TCPOptionKindTimestamps {

			return int(binary.BigEndian.Uint32(opt.OptionData[4:8]))
		}
	}

	return 0
}

func parseHttpRequestUrl(req *http.Request) string {
	var url string

	if l := req.Host; l != "" {
		url = l
		if u := req.URL; u != nil {
			url += u.String()
		}
	}

	return url
}
