package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(varName string) (string, error) {
	if err := godotenv.Load(); err != nil {
		return "", err
	}
	value, ok := os.LookupEnv(varName)
	if !ok {
		erroMessage := fmt.Sprintf("%s not found", varName)
		return "", errors.New(erroMessage)
	}
	return value, nil
}
