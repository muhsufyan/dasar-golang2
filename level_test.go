package dasargolang2_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	// func jenis level, semakin kebawah semakin fatal. paramnya kita isi semua dg info
	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
	// untuk fatal & panic akan otomatis exit
	logger.Fatal("This is fatal")
	logger.Panic("This is panic")
}

func TestLoggingLevelTrace(t *testing.T) {
	logger := logrus.New()

	// ubah level dari defaultnya info level ke trace
	logger.SetLevel(logrus.TraceLevel)

	// func jenis level, semakin kebawah semakin fatal. paramnya kita isi semua dg info
	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
	// untuk fatal & panic akan otomatis exit
	logger.Fatal("This is fatal")
	logger.Panic("This is panic")
}

func TestLoggingLevelWarn(t *testing.T) {
	logger := logrus.New()

	// ubah level dari defaultnya info level ke trace
	logger.SetLevel(logrus.WarnLevel)

	// func jenis level, semakin kebawah semakin fatal. paramnya kita isi semua dg info
	logger.Trace("This is trace")
	logger.Debug("This is debug")
	logger.Info("This is info")
	logger.Warn("This is warn")
	logger.Error("This is error")
	// untuk fatal & panic akan otomatis exit
	logger.Fatal("This is fatal")
	logger.Panic("This is panic")
}
