package config

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var RootFolder string

func InitFS() error {
	rootFolder, err := GetEnvVar("ROOT_FOLDER")

	if err != nil {
		rootFolder = "/tmp/blober"
	}

	err = os.MkdirAll(rootFolder, os.ModePerm)

	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't create folder")
		return errors.New(errorMessage)
	}
	log.Printf("%s created\n", rootFolder)

	RootFolder = rootFolder

	return nil
}
