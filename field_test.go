package dasargolang2_test

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()

	// one field (id_account)
	logger.WithField("id_account", "2").Info("this account can not login error")
	// will print time="2022-08-11T10:03:27+07:00" level=info msg="this account can not login error" id_account=2

	// multiple field
	logger.WithField("id_account", "3").
		WithField("username", "mamat").
		Info("this account with username is backlist")
	// will print time="2022-08-11T10:03:27+07:00" level=info msg="this account with username is backlist" id_account=3 username=mamat
}

func TestFieldJSONFormatter(t *testing.T) {
	logger := logrus.New()

	// change formatter
	logger.SetFormatter(&logrus.JSONFormatter{})

	// one field (id_account)
	logger.WithField("id_account", "2").Info("this account can not login error")
	// will print {"id_account":"2","level":"info","msg":"this account can not login error","time":"2022-08-11T10:04:49+07:00"}

	// multiple field
	logger.WithField("id_account", "3").
		WithField("username", "mamat").
		Info("this account with username is backlist")
	// will print {"id_account":"3","level":"info","msg":"this account with username is backlist","time":"2022-08-11T10:04:49+07:00","username":"mamat"}
}

// multiple fields
// fields is alias from map[string]interface
func TestMultipleFieldsJSONFormatter(t *testing.T) {
	logger := logrus.New()

	// change formatter
	logger.SetFormatter(&logrus.JSONFormatter{})

	// multiple fields
	logger.WithFields(logrus.Fields{
		"id_account": "3",
		"username":   "mamat",
	}).Info("this account with username is backlist")
	// will print {"id_account":"3","level":"info","msg":"this account with username is backlist","time":"2022-08-11T10:04:49+07:00","username":"mamat"}
}
