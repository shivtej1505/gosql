package utils

import (
	"errors"
	"os"
)

func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	// dir exists, nothing to do
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		if err := os.Mkdir(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
