package common

import "github.com/sirupsen/logrus"

// Check is shorthand for dieing if e contains a value
func Check(e error) {
	if e != nil {
		logrus.Fatal(e)
	}
}
