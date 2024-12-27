package blob_slice

import (
	"io"
	"log"
	"net/http"
	"time"

	"beto0607.com/blober/src/utils"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
}

func GetBlobMetadata(w http.ResponseWriter, r *http.Request) {
}

func PostBlob(w http.ResponseWriter, r *http.Request) {
	entity := BlobModel{
		Name:        "",
		MimeType:    "",
		Status:      "Creating",
		SizeInBytes: 0,
		Path:        "",
		Parent:      "",
		CreatedAt:   time.Now().UTC().String(),
	}

	directoryName := utils.GenerateRandomHexString(5)
	fileName := utils.Uuidv4()

	subpath := directoryName + "/" + fileName
	entity.Path = subpath

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Problems with received file")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	entity.SizeInBytes = header.Size
	entity.Name = header.Filename

	fo, err := utils.CreateFile(directoryName, fileName)
	if err != nil {
		log.Println("Problems with creating file")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer fo.Close()
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if entity.MimeType == "" {
			entity.MimeType = utils.GetMimeType(buf)
		}

		// write a chunk
		if _, err := fo.Write(buf[:n]); err != nil {
			panic(err)
		}
	}

}

func PutBlob(w http.ResponseWriter, r *http.Request) {
}

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
}

