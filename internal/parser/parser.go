package parser

import (
	"pparse/config"
	httppack "pparse/internal/parser/http-pack"
	"pparse/internal/source"
)

const HTTP = "HTTP"

type PacketParserI interface {
	Run() (int, error)
}

func New(conf config.ParserConfig) PacketParserI {
	var entity PacketParserI

	switch conf.Protocol {
	case HTTP:
		c := httppack.NewCalculator()
		entity = httppack.Parser{
			Config:     conf,
			DataSource: source.NewDataSource(conf.FilePath, conf.NetInterface),
			Counter:    source.NewCounter(conf.FilePath, conf.NetInterface, c, conf.MetricsInterval),
			Calculator: c,
		}
	}

	return entity
}
