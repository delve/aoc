package gen

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

func createDirectory(path string) error {
	if _, err := os.Stat(path); err == nil || !errors.Is(err, os.ErrNotExist) {
		logrus.Infof("Directory already exists: %s", path)
		return nil
	}

	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}

	logrus.Infof("Created directory: %s", path)
	return nil
}

func Init() error {
	logrus.Info("setting up from gopherholes.yaml")
	return nil
}
