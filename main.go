package main

import (
	log "github.com/sirupsen/logrus"
	"pparse/config"
	"pparse/internal/parser"
	"pparse/logger"
)

func main() {
	logger.Init()
	log.Info("Parser app started")

	var count int
	defer func() {
		log.WithField("packets", count).Info("Done") // TODO make it work for network, add graceful shutdown
	}()

	conf, err := config.Init()
	if err != nil {
		log.WithError(err).Fatal("Config error")
	}

	log.WithFields(log.Fields{
		"protocol":        conf.Protocol,
		"filePath":        conf.FilePath,
		"netInterface":    conf.NetInterface,
		"metricsInterval": conf.MetricsInterval,
	}).Info("Run parser")

	app := parser.New(conf)
	count, err = app.Run()
	if err != nil {
		log.WithError(err).Fatal("Parser failed")
	}
}
