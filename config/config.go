package config

import (
	"errors"
	"flag"
	"github.com/joho/godotenv"
	"os"
)

const (
	DefaultMetricsInterval = 60
)

type ParserConfig struct {
	MetricsInterval int // seconds
	FilePath        string
	NetInterface    string
	Protocol        string
}

func (c ParserConfig) Validate() error {
	if c.FilePath == "" && c.NetInterface == "" {
		return errors.New("file path or network interface is required")
	}

	return nil
}

func Init() (ParserConfig, error) {
	filePath := flag.String("file", "", "File name to parse")
	netInterface := flag.String("net", "", "Network interface name to receive packets")
	metricsInterval := flag.Int("interval", DefaultMetricsInterval, "Events aggregation interval in seconds")
	flag.Parse()

	if *filePath == "" && *netInterface == "" {
		return ParserConfig{}, errors.New("file path or network interface is required")
	}

	if err := godotenv.Load(); err != nil {
		return ParserConfig{}, errors.New("error loading .env file")
	}

	protocol, ok := os.LookupEnv("PROTOCOL")
	if !ok {
		return ParserConfig{}, errors.New("cannot get PROTOCOL setting from .env")
	}

	return ParserConfig{
		MetricsInterval: *metricsInterval,
		FilePath:        *filePath,
		NetInterface:    *netInterface,
		Protocol:        protocol,
	}, nil
}
