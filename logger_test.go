package dasargolang2_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	// make logger
	logger := logrus.New()

	// print logger, yg ditampilkan adlh time="2022-08-11T08:28:45+07:00" level=info msg="Hallo logger"
	logger.Println("Hallo logger")
}
