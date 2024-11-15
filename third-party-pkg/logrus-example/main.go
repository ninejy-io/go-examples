package main

import (
	"github.com/sirupsen/logrus"

	"logrus-example/log"
)

func main() {
	log.Log.WithFields(logrus.Fields{"animal": "dog"}).Info("this is a test log")
}
