package blob_slice

import (
	"beto0607.com/blober/src/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlobModel struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Name        string             `json:"name" validate:"required" bson:"name"`
	Filename    string             `json:"filename" validate:"required" bson:"filename"`
	MimeType    string             `json:"mimeType" bson:"mimeType"`
	Status      string             `json:"status" bson:"status"`
	SizeInBytes int64              `json:"size" bson:"size"`
	Path        string             `json:"path,omitempty" bson:"path"`
	Parent      string             `json:"parent,omitempty" bson:"parent"`
	CreatedAt   string             `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt   string             `json:"updatedAt,omitempty" bson:"updated_at"`
	DeletedAt   string             `json:"deletedAt,omitempty" bson:"deleted_at"`
}

func NewBlobModel() BlobModel {
	return BlobModel{
		Id:          primitive.NewObjectID(),
		Name:        "",
		Filename:    "",
		MimeType:    "",
		Status:      "Creating",
		SizeInBytes: 0,
		Path:        "",
		Parent:      "",
		CreatedAt:   utils.UTCTimestamp(),
	}
}
