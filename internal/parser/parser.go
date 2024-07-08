package parser

import (
	"fmt"
	"pparse/config"
	httppack "pparse/internal/parser/http-pack"
	"pparse/internal/source"
)

const HTTP = "HTTP"

type PacketParserI interface {
	Run() (int, error)
}

func New(conf config.ParserConfig) (PacketParserI, error) {
	switch conf.Protocol {
	case HTTP:
		c := httppack.NewCalculator()
		src, err := source.NewDataSource(conf.FilePath, conf.NetInterface)
		if err != nil {
			return nil, err
		}
		counter, err := source.NewCounter(conf.FilePath, conf.NetInterface, c, conf.MetricsInterval)
		if err != nil {
			return nil, err
		}
		return httppack.Parser{
			Config:     conf,
			DataSource: src,
			Counter:    counter,
			Calculator: c,
		}, nil
	default:
		return nil, fmt.Errorf("protocol [%s] is not supported", conf.Protocol)
	}
}
