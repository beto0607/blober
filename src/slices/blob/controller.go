package blob_slice

import (
	"encoding/json"
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
		log.Println("Problems with creating blob entry")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	directoryName := utils.GenerateRandomHexString(5)
	fileName := utils.Uuidv4()

	subpath := directoryName + "/" + fileName

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println("Problems with received file")
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	fo, err := utils.CreateFileFrom(directoryName, fileName, &file)
	if err != nil {
		log.Println("Problems with creating file")
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer fo.Close()

	entity.SizeInBytes = header.Size
	entity.Name = header.Filename
	entity.Filename = header.Filename
	entity.Path = subpath
	entity.MimeType = utils.GetMimeType(entity.Path)

	entity.Status = "Created"
	SaveBlobEntity(entity)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&entity)
}

func PutBlob(w http.ResponseWriter, r *http.Request) {
}

func DeleteBlob(w http.ResponseWriter, r *http.Request) {
	blobId := r.PathValue("id")
	if len(blobId) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	hardDelete := r.URL.Query().Get("hardDelete") == "true"
	if !hardDelete {
		blob, err := FindBlobEntity(blobId)
		if err != nil {
			log.Printf("Couldn't find %s\n", blobId)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		err = DeleteBlobEntity(blob, false)

		if err != nil {
			log.Printf("Couldn't delete %s\n", blobId)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		return
	}

	blob, err := FindDeletedBlobEntity(blobId)
	if err != nil {
		log.Printf("Couldn't find %s\n", blobId)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = utils.DeleteFile(blob.Path)

	if err != nil {
		log.Printf("Couldn't delete file for %s\n", blobId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = DeleteBlobEntity(blob, true)

	if err != nil {
		log.Printf("Couldn't hard delete %s from DB\n", blobId)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
