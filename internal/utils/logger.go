package utils

import (
	"fmt"
	"github.com/gogap/logrus_mate"
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func SetupLogger() {
	mate, err := logrus_mate.NewLogrusMate(
		logrus_mate.ConfigFile(
			"./logger.conf",
		),
	)
	if err != nil {
		fmt.Printf("Error loading logger configuration: %v", err)
		os.Exit(1)
	}

	Log = mate.Logger("echo_logger")

	Log.Trace("This is a trace message")
	Log.Debug("This is a debug message")
	Log.Info("This is an info message")
	Log.Warn("This is a warning message")
	Log.Error("This is an error message")
}
