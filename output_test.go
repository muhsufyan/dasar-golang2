package dasargolang2_test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestChangeOutput(t *testing.T) {
	// make logger
	logger := logrus.New()

	// config file for output
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// change output from console (default) to file
	logger.SetOutput(file)

	logger.Info("hallo logging from info")
	logger.Warn("hallo logging from Warn")
	logger.Error("hallo logging from Error")
}
