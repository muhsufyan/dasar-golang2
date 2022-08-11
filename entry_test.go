package dasargolang2_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})

	// make entry
	entry := logrus.NewEntry(logger)

	// entry have properties, we add properties
	entry.WithField("id", "4")

	entry.Info("hi")

}
