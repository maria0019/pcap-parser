package parser

import (
	"github.com/stretchr/testify/assert"
	"pparse/config"
	"testing"
)

func Test_Parser(t *testing.T) {
	var tests = []struct {
		title           string
		conf            config.ParserConfig
		isErrorExpected bool
	}{
		{"Pass config with HTTP protocol",
			config.ParserConfig{MetricsInterval: 10, FilePath: "abc.pcap", NetInterface: "wlp0s20f3", Protocol: "HTTP"},
			false,
		},
		{"Pass config with empty protocol",
			config.ParserConfig{MetricsInterval: 10, FilePath: "abc.pcap", NetInterface: "wlp0s20f3", Protocol: ""},
			true,
		},
		{"Pass config with empty Filename and empty Net interface",
			config.ParserConfig{MetricsInterval: 10, FilePath: "", NetInterface: "", Protocol: "HTTP"},
			true,
		},
		{"Pass config with empty file name",
			config.ParserConfig{MetricsInterval: 10, FilePath: "", NetInterface: "wlp0s20f3", Protocol: "HTTP"},
			false,
		},
		{"Pass config with empty network interface",
			config.ParserConfig{MetricsInterval: 10, FilePath: "abc.pcap", NetInterface: "", Protocol: "HTTP"},
			false,
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			_, err := New(test.conf)

			if test.isErrorExpected {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
