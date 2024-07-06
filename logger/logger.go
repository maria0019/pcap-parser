package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}
