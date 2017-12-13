package logger

import (
	"github.com/sirupsen/logrus"
)

func GetLogger() logrus {

	return logrus.New()
}
