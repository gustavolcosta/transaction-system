package log_application

import "log/slog"

func Info(msg string, context string) {
	slog.Info(msg, "context", context)
}

func Error(msg string, err error, context string) {
	slog.Error(msg, "error", err, "context", context)
}
