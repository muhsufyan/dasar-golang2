package dasargolang2_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestChangeFormatter(t *testing.T) {
	// make logger
	logger := logrus.New()

	// change TextFormatter (default) to JSONFormatter
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hallo logging from info")
	logger.Warn("hallo logging from Warn")
	logger.Error("hallo logging from Error")
}
