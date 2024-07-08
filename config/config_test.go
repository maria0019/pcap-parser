package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_Validate(t *testing.T) {
	var tests = []struct {
		title        string
		fileFlag     string
		netFlag      string
		intervalFlag int
		protocol     string
		errorMessage string
	}{
		{"File and Net flags are passed",
			"file.pcap", "wlp0s20f3", 1, "HTTP", ""},
		{"Protocol is not passed",
			"file.pcap", "wlp0s20f3", 1, "", "cannot get PROTOCOL setting from .env"},
		{"File and Net flags are empty",
			"", "", 1, "HTTP", "file path or network interface is required",
		},
		{"File extension is invalid",
			"abc.txt", "", 1, "HTTP", fmt.Sprintf("file should have %s extension", FileExtension)},
		{"File extension is invalid",
			"file_pcap", "", 1, "HTTP", fmt.Sprintf("file should have %s extension", FileExtension)},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			conf := ParserConfig{
				MetricsInterval: 10,
				FilePath:        test.fileFlag,
				NetInterface:    test.netFlag,
				Protocol:        test.protocol,
			}
			err := conf.Validate()

			if test.errorMessage == "" {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
