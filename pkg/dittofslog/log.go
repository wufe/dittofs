package dittofslog

import (
	"log/slog"

	"github.com/marmos91/dittofs/internal/logger"
)

// SetLevel sets the minimum log level for the dittofs internal logger.
// Valid values: "DEBUG", "INFO", "WARN", "ERROR" (case-insensitive).
func SetLevel(level string) {
	logger.SetLevel(level)
}

// SetHandler replaces the dittofs logger's slog.Handler entirely.
// Use this to route dittofs logs through any external logger — zerolog, zap,
// a discard handler, etc. Level filtering is delegated to the provided handler.
func SetHandler(h slog.Handler) {
	logger.SetHandler(h)
}
