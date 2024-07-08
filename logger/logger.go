package logger

import "github.com/gookit/slog"

func Init() {
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.SetTemplate(slog.NamedTemplate)
	})
}
