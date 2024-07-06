package sender

import (
	log "github.com/sirupsen/logrus"
)

func ToStdout(title string, params map[string]interface{}) {
	log.WithFields(params).Info(title)
}
