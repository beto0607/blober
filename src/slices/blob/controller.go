package blob_slice

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"beto0607.com/blober/src/utils"
)

func GetBlob(w http.ResponseWriter, r *http.Request) {
	blobId := r.PathValue("id")
	if len(blobId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	blob, err := FindBlobEntity(blobId)
	if err != nil {
		log.Printf("Couldn't find %s\n", blobId)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fileData, err := utils.ReadFile(blob.Path)

	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(fileData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", blob.MimeType)
}

func GetBlobMetadata(w http.ResponseWriter, r *http.Request) {
	blobId := r.PathValue("id")
	if len(blobId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	blob, err := FindBlobEntity(blobId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&blob)
}

func PostBlob(w http.ResponseWriter, r *http.Request) {
	entity, err := CreateBlobEntity()
	if err != nil {
		log.Println("Problems with creating file")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
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

	entity.Status = "Created"
	SaveBlobEntity(entity)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&entity)
}

func PutBlob(w http.ResponseWriter, r *http.Request) {
}

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
}
