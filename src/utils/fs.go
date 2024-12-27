package utils

import (
	"mime/multipart"
	"os"
	"path"
)

var RootFolder string

func ReadFile(subpath string) ([]byte, error) {
	path := RootFolder + "/" + subpath
	return os.ReadFile(path)
}

func WriteFile(subpath string, data []byte) error {
	path := RootFolder + "/" + subpath

	return os.WriteFile(path, data, 0600)
}

func createFile(directory string, fileName string) (*os.File, error) {
	directoryPath := path.Join(RootFolder, directory)
	err := os.MkdirAll(directoryPath, 0700)
	if err != nil {
		return nil, err
	}
	filePath := path.Join(RootFolder, directory, fileName)
	return os.Create(filePath)
}

func CreateFileFrom(directory string, fileName string, input *multipart.File) (*os.File, error) {
	output, err := createFile(directory, fileName)
	if err != nil {
		return nil, err
	}
	output.ReadFrom(*input)

	return output, nil
}

func DeleteFile(filepath string) error {
	pathToFile := path.Join(RootFolder, filepath)
	return os.Remove(pathToFile)
}
