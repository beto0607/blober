package utils

import (
	"log"

	"github.com/gabriel-vasile/mimetype"
)

func GetMimeType(subpath string) string {
	pathToFile := RootFolder + "/" + subpath
	mtype, err := mimetype.DetectFile(pathToFile)
	if err != nil {
		return "application/octet-stream"
	}
	log.Printf("MimeType found: %s", mtype.String())
	return mtype.String()
}
