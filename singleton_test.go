package dasargolang2_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("this is info")
	logrus.Warn("this is warn")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	// dg mengubah jd JSONFormatter semuanya akan berubah jd json sehingga dari pada sprti itu (use singleton) lbh baik gunakan logrus.New() terlbh dahulu
	logrus.Info("this is info")
	logrus.Warn("this is warn")
}
