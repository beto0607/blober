package utils

import (
	"log"

	"github.com/gabriel-vasile/mimetype"
)

func GetMimeType(fileHeader []byte) string {
	mtype := mimetype.Detect(fileHeader)
	log.Printf("MimeType found: %s", mtype.String())
	return mtype.String()
}
