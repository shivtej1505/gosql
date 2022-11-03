package command

import "os"

// Give table instance
func Insert(table string, value string) {
	dir := "data"
	err := createDir(dir)
	if err != nil {
		panic(err)
	}

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(dir+"/"+table, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	_, err = f.Write([]byte(value))
	if err != nil {
		panic(err)
	}
}

func createDir(dir string) error {
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
