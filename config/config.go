package config

var (
	logger *Logger
)

func GetLogger(prefix string) *Logger {
	// Initialize Logger
	logger = NewLogger(prefix)
	return logger
}
