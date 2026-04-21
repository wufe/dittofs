package dittofslog

import "github.com/marmos91/dittofs/internal/logger"

// SetLevel sets the minimum log level for the dittofs internal logger.
// Valid values: "DEBUG", "INFO", "WARN", "ERROR" (case-insensitive).
func SetLevel(level string) {
	logger.SetLevel(level)
}
