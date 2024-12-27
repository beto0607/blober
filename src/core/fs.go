package core

import (
	"errors"
	"fmt"
	"log"
	"os"

	"beto0607.com/blober/src/config"
	"beto0607.com/blober/src/utils"
)

func InitFS() error {
	rootFolder, err := config.GetEnvVar("ROOT_FOLDER")

	if err != nil {
		rootFolder = "/tmp/blober"
	}

	err = os.MkdirAll(rootFolder, os.ModePerm)

	if err != nil {
		errorMessage := fmt.Sprintf("Couldn't create folder")
		return errors.New(errorMessage)
	}
	log.Printf("%s created\n", rootFolder)

	utils.RootFolder = rootFolder

	return nil
}
