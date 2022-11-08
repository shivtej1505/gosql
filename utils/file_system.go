package utils

import "os"

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
