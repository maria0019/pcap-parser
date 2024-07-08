package sender

import (
	log "github.com/gookit/slog"
)

func ToStdout(title string, params map[string]any) {
	log.WithData(params).Info(title)
}
