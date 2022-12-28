package common

import "github.com/sirupsen/logrus"

func Check(e error) {
	if e != nil {
		logrus.Fatal(e)
	}
}
