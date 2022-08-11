package dasargolang2_test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

// make implementasi dari interface hook
type SampleHook struct {
}

// implementasi func untuk Hook (levels & fire)
func (s *SampleHook) Levels() []logrus.Level {
	// ketika ada event dg level error &/ warn maka hook(callback yaitu func fire) akan dieksekusi
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}
func (s *SampleHook) Fire(entry *logrus.Entry) error {
	// ini adalah callback yg akan di run, misal advance nya kirim notif ke email/chat tp now ckp print saja
	fmt.Println("sample hook, callback is run", entry.Level, entry.Message)
	return nil
}
func TestHook(t *testing.T) {
	logger := logrus.New()

	// buat hook
	logger.AddHook(&SampleHook{})

	logger.Info("Fire/callback tdk run")
}

func TestHookRunCallback(t *testing.T) {
	logger := logrus.New()

	// buat hook
	logger.AddHook(&SampleHook{})

	logger.Info("Fire/callback tdk run")
	logger.Warn("Fire/callback run")
}
