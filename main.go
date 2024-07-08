package main

import (
	log "github.com/gookit/slog"
	"pparse/config"
	"pparse/internal/parser"
	"pparse/logger"
)

func main() {
	logger.Init()
	log.Info("Parser app started")

	conf, err := config.Init()
	if err != nil {
		log.WithData(log.M{"error": err.Error()}).Error("Config init error")
	}
	if err := conf.Validate(); err != nil {
		log.WithData(log.M{"error": err.Error()}).Error("Config validation error")
	}

	log.WithData(log.M{
		"protocol":        conf.Protocol,
		"filePath":        conf.FilePath,
		"netInterface":    conf.NetInterface,
		"metricsInterval": conf.MetricsInterval,
	}).Info("Run parser")

	app, err := parser.New(conf)
	if err != nil {
		log.WithData(log.M{"error": err.Error()}).Error("App init failed")
	}

	_, err = app.Run()
	if err != nil {
		log.WithData(log.M{"error": err.Error()}).Error("App run failed")
	}
}
