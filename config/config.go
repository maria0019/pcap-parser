package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

const (
	DefaultMetricsInterval = 60
	FileExtension          = ".pcap"
)

type ParserConfig struct {
	MetricsInterval int // seconds
	FilePath        string
	NetInterface    string
	Protocol        string
}

func Init() (ParserConfig, error) {
	filePath := flag.String("file", "", "File name to parse")
	netInterface := flag.String("net", "", "Network interface name to receive packets")
	metricsInterval := flag.Int("interval", DefaultMetricsInterval, "Events aggregation interval in seconds")
	flag.Parse()

	conf := ParserConfig{
		MetricsInterval: *metricsInterval,
		FilePath:        *filePath,
		NetInterface:    *netInterface,
		Protocol:        "HTTP",
	}

	if err := godotenv.Load(); err != nil {
		return ParserConfig{}, errors.New("error loading .env file")
	}

	conf.Protocol, _ = os.LookupEnv("PROTOCOL")

	return conf, nil
}

func (conf ParserConfig) Validate() error {
	if conf.FilePath == "" && conf.NetInterface == "" {
		return errors.New("file path or network interface is required")
	}

	if conf.FilePath != "" && conf.FilePath[len(conf.FilePath)-5:] != FileExtension {
		return fmt.Errorf("file should have %s extension", FileExtension)
	}

	if conf.Protocol == "" {
		return errors.New("cannot get PROTOCOL setting from .env")
	}

	return nil
}
