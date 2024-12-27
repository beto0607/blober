package utils

import (
	"os"
	"path"

	"beto0607.com/blober/src/config"
)

func ReadFile(subpath string) ([]byte, error) {
	path := config.RootFolder + "/" + subpath
	return os.ReadFile(path)
}

func WriteFile(subpath string, data []byte) error {
	path := config.RootFolder + "/" + subpath

	return os.WriteFile(path, data, 0600)
}

func CreateFile(directory string, fileName string) (*os.File, error) {
	directoryPath := path.Join(config.RootFolder, directory)
	err := os.MkdirAll(directoryPath, 0700)
	if err != nil {
		return nil, err
	}
	filePath := path.Join(config.RootFolder, directory, fileName)
	return os.Create(filePath)
}
