// logging.go in yourloggingpackage
package logging

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

// InitLogger initializes the logger with desired configurations
func InitLogger() {
	// Set logrus to use the default logrus formatter
	Log.SetFormatter(&logrus.TextFormatter{})

	// You can customize other logrus configurations as needed
	// For example:
	// Log.SetLevel(logrus.DebugLevel)
}

// LogInfo logs an informational message
func LogInfo(message string) {
	Log.Info(message)
}

// LogWarning logs a warning message
func LogWarning(message string) {
	Log.Warn(message)
}

// LogError logs an error message
func LogError(message string) {
	Log.Error(message)
}

// LogFatal logs a fatal message and exits the program
func LogFatal(message string) {
	Log.Fatal(message)
}
