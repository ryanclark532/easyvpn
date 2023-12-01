package utils

import "os"

func WriteFile(filePath string, contents []byte) error  {
	err := os.WriteFile(filePath, contents, 0644)
	if err != nil {
		return err
	}
	return nil

}