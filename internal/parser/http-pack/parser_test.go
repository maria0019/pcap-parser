package httppack

import (
	"github.com/stretchr/testify/assert"
	"pparse/config"
	"pparse/internal/calculator"
	"pparse/internal/source"
	"pparse/internal/source/file"
	sourceMock "pparse/internal/source/mock"
	"testing"
)

// Mocked the data source only (used HTTP packets mock). Packets counter and stats calculator aren't mocked
func Test_Parser_File(t *testing.T) {
	conf := config.ParserConfig{
		MetricsInterval: 10,
		FilePath:        "any.pcap",
		NetInterface:    "",
		Protocol:        "HTTP",
	}
	dataSource := sourceMock.NewSourceMockHttpPacket(conf.FilePath)
	calc := NewCalculator()
	counter := file.NewCounter(calc, conf.MetricsInterval)

	var tests = []struct {
		title      string
		Config     config.ParserConfig
		DataSource source.DataSourceI
		Counter    source.CounterI
		Calculator calculator.CalculatorI
	}{
		{
			"File parser",
			conf,
			dataSource,
			&counter,
			calc,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			p := Parser{
				Config:     conf,
				DataSource: dataSource,
				Counter:    &counter,
				Calculator: calc,
			}
			count, err := p.Run()

			assert.Equal(t, 2, count)
			assert.NoError(t, err)
		})
	}
}
