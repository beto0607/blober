package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVar(varName string) (string, error) {
	_ = godotenv.Load()
	value, ok := os.LookupEnv(varName)
	if !ok {
		erroMessage := fmt.Sprintf("%s not found", varName)
		return "", errors.New(erroMessage)
	}
	return value, nil
}
