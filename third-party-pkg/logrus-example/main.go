package main

import (
	"github.com/sirupsen/logrus"

	"go-examples/third-party-pkg/logrus-example/log"
)

func main() {
	log.Log.WithFields(logrus.Fields{"animal": "dog"}).Info("this is a test log")
}
